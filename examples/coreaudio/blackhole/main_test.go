package main

import (
	"fmt"
	"math"
	"testing"
)

func TestVolumeConversions(t *testing.T) {
	tests := []struct {
		name       string
		scalar     float32
		wantLinear float32
		wantDB     float32
	}{
		{
			name:       "zero volume",
			scalar:     0.0,
			wantLinear: 0.0,
			wantDB:     -64.0,
		},
		{
			name:       "full volume",
			scalar:     1.0,
			wantLinear: 1.0,
			wantDB:     0.0,
		},
		{
			name:       "half scalar volume",
			scalar:     0.5,
			wantLinear: VolumeFromScalar(0.5),
			wantDB:     -32.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLinear := VolumeFromScalar(tt.scalar)
			if math.Abs(float64(gotLinear-tt.wantLinear)) > 1e-4 {
				t.Errorf("VolumeFromScalar(%v) = %v; want %v", tt.scalar, gotLinear, tt.wantLinear)
			}

			gotScalar := VolumeToScalar(gotLinear)
			if math.Abs(float64(gotScalar-tt.scalar)) > 1e-4 {
				t.Errorf("VolumeToScalar(%v) = %v; want %v", gotLinear, gotScalar, tt.scalar)
			}

			gotDB := VolumeToDecibel(gotLinear)
			if math.Abs(float64(gotDB-tt.wantDB)) > 1e-4 {
				t.Errorf("VolumeToDecibel(%v) = %v; want %v", gotLinear, gotDB, tt.wantDB)
			}
		})
	}
}

func TestDriverLoopback(t *testing.T) {
	driver, err := NewDriver(48000.0, 2)
	if err != nil {
		t.Fatalf("NewDriver failed: %v", err)
	}

	if err := driver.StartIO(); err != nil {
		t.Fatalf("StartIO failed: %v", err)
	}
	defer driver.StopIO()

	frames := uint32(100)
	writeBuf := make([]float32, frames*2)
	for i := range writeBuf {
		writeBuf[i] = float32(i) * 0.001
	}

	if err := driver.WriteMix(0, frames, writeBuf); err != nil {
		t.Fatalf("WriteMix failed: %v", err)
	}

	readBuf := make([]float32, frames*2)
	if err := driver.ReadInput(0, frames, readBuf); err != nil {
		t.Fatalf("ReadInput failed: %v", err)
	}

	for i := range writeBuf {
		if math.Abs(float64(writeBuf[i]-readBuf[i])) > 1e-6 {
			t.Fatalf("Sample mismatch at index %d: wrote %f, read %f", i, writeBuf[i], readBuf[i])
		}
	}
}

func TestDriverFeatures(t *testing.T) {
	driver, err := NewDriver(48000.0, 2)
	if err != nil {
		t.Fatalf("NewDriver failed: %v", err)
	}

	// Client tracking
	driver.AddClient(1, 100)
	if driver.ClientCount() != 1 {
		t.Errorf("ClientCount() = %d; want 1", driver.ClientCount())
	}
	driver.RemoveClient(1)
	if driver.ClientCount() != 0 {
		t.Errorf("ClientCount() after remove = %d; want 0", driver.ClientCount())
	}

	// Pitch and Clock source
	driver.SetPitchAdjust(0.8)
	if driver.PitchAdjust() != 0.8 {
		t.Errorf("PitchAdjust() = %f; want 0.8", driver.PitchAdjust())
	}

	if err := driver.SetClockSource(ClockSourceInternalAdjustable); err != nil {
		t.Errorf("SetClockSource failed: %v", err)
	}
	if driver.ClockSource() != ClockSourceInternalAdjustable {
		t.Errorf("ClockSource() = %d; want %d", driver.ClockSource(), ClockSourceInternalAdjustable)
	}

	// Property querying
	devName, err := driver.GetPropertyData(ObjectIDDevice, AudioObjectPropertyAddress{Selector: PropName})
	if err != nil || devName != "BlackHole 2ch" {
		t.Errorf("GetPropertyData PropName = %v, err = %v; want 'BlackHole 2ch'", devName, err)
	}

	dev2Name, err := driver.GetPropertyData(ObjectIDDevice2, AudioObjectPropertyAddress{Selector: PropName})
	if err != nil || dev2Name != "BlackHole 2ch 2" {
		t.Errorf("GetPropertyData PropName Device2 = %v, err = %v; want 'BlackHole 2ch 2'", dev2Name, err)
	}
}

func Example() {
	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Output:
	// Initializing BlackHole Virtual Audio Driver...
	// Driver: BlackHole 2ch (2 channels, 48000 Hz, status: stopped, volume: 1.00, muted: false)
	// Manufacturer: Existential Audio Inc.
	// Bundle ID: audio.existential.BlackHole2ch
	// Primary Device: BlackHole 2ch
	// Mirror Device: BlackHole 2ch 2 (Hidden: true)
	// Format: 2 channels, 48000 Hz, 32 bits/channel
	// Attached Audio Clients: 1
	//
	// Starting IO loopback engine...
	// Audio Loopback 100% Match: true
	// Set Volume: 0.5 (dB: -6.0)
	// Volume Scaled Sample [0] (written 0.06 -> read 0.03): true
	// Clock Source: 1 (Pitch Adjust: 0.75)
	// Set Mute: true
	// Muted Buffer Zeroed Out: true
	// Updated Sample Rate: 96000 Hz
	// ZeroTimeStamp: sampleTime=0 (hostTime positive: true)
	// Supports NominalSampleRate Property Query: true
	// BlackHole Virtual Audio Driver simulation finished successfully.
}
