//go:build darwin

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWAVHeader(t *testing.T) {
	var buf bytes.Buffer
	err := writeWAVHeader(&buf, 48000, 2, 96000)
	if err != nil {
		t.Fatalf("writeWAVHeader: %v", err)
	}

	data := buf.Bytes()
	if len(data) != 44 {
		t.Fatalf("expected header length 44, got %d", len(data))
	}

	if string(data[0:4]) != "RIFF" {
		t.Errorf("expected RIFF, got %s", string(data[0:4]))
	}

	if string(data[8:12]) != "WAVE" {
		t.Errorf("expected WAVE, got %s", string(data[8:12]))
	}

	channels := binary.LittleEndian.Uint16(data[22:24])
	if channels != 2 {
		t.Errorf("expected 2 channels, got %d", channels)
	}

	sampleRate := binary.LittleEndian.Uint32(data[24:28])
	if sampleRate != 48000 {
		t.Errorf("expected 48000 sample rate, got %d", sampleRate)
	}
}

func TestCalculateRMS(t *testing.T) {
	tests := []struct {
		name        string
		samples     []int16
		wantZero    bool
		minExpected float64
	}{
		{
			name:     "silence",
			samples:  []int16{0, 0, 0, 0, 0, 0},
			wantZero: true,
		},
		{
			name:        "sine-like wave",
			samples:     []int16{0, 16384, 32767, 16384, 0, -16384, -32767, -16384},
			wantZero:    false,
			minExpected: 0.4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]byte, len(tt.samples)*2)
			for i, s := range tt.samples {
				binary.LittleEndian.PutUint16(buf[i*2:], uint16(s))
			}
			rms := CalculateRMS(buf)
			if tt.wantZero && rms != 0 {
				t.Errorf("expected RMS 0, got %f", rms)
			}
			if !tt.wantZero && rms < tt.minExpected {
				t.Errorf("expected RMS >= %f, got %f", tt.minExpected, rms)
			}
		})
	}
}

func TestRecorderCapture(t *testing.T) {
	if os.Getenv("APPLE_SYSTAP_LIVE") != "1" {
		t.Skip("set APPLE_SYSTAP_LIVE=1 to exercise the permission- and hardware-dependent process tap")
	}
	tmpDir := t.TempDir()
	outPath := filepath.Join(tmpDir, "test.wav")

	opts := Options{
		OutputFile: outPath,
		SampleRate: 48000,
		Channels:   2,
		Verbose:    false,
	}

	rec, err := NewRecorder(opts)
	if err != nil {
		t.Fatalf("NewRecorder: %v", err)
	}

	if err := rec.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	if err := rec.Stop(); err != nil {
		t.Fatalf("Stop: %v", err)
	}

	if rec.CallbacksCount() == 0 {
		t.Errorf("expected > 0 callbacks, got 0")
	}

	fi, err := os.Stat(outPath)
	if err != nil {
		t.Fatalf("stat output file: %v", err)
	}

	if fi.Size() <= 44 {
		t.Errorf("expected WAV file size > 44 bytes, got %d", fi.Size())
	}
}

func Example() {
	opts := Options{
		SampleRate: 48000,
		Channels:   2,
	}
	rec, err := NewRecorder(opts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Recorder initialized:", rec != nil)
	// Output:
	// Recorder initialized: true
}
