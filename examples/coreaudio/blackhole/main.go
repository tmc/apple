// Command blackhole demonstrates the BlackHole virtual audio loopback driver in Go.
//
// BlackHole is a zero-latency virtual audio device that routes audio output
// from one application to audio input of another application.
//
// Run with:
//
//	go run ./examples/coreaudio/blackhole
package main

import (
	"fmt"
	"math"

	"github.com/tmc/apple/coreaudio"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func run() error {
	fmt.Println("Initializing BlackHole Virtual Audio Driver...")

	driver, err := NewDriver(48000.0, 2)
	if err != nil {
		return fmt.Errorf("initialize driver: %w", err)
	}

	fmt.Printf("Driver: %s\n", driver)
	fmt.Printf("Manufacturer: %s\n", ManufacturerName)
	fmt.Printf("Bundle ID: %s\n", PlugInBundleID)

	// 1. Device and Mirror Device property queries
	devName, _ := driver.GetPropertyData(ObjectIDDevice, AudioObjectPropertyAddress{Selector: PropName})
	dev2Name, _ := driver.GetPropertyData(ObjectIDDevice2, AudioObjectPropertyAddress{Selector: PropName})
	dev2Hidden, _ := driver.GetPropertyData(ObjectIDDevice2, AudioObjectPropertyAddress{Selector: PropIsHidden})
	fmt.Printf("Primary Device: %v\n", devName)
	fmt.Printf("Mirror Device: %v (Hidden: %v)\n", dev2Name, dev2Hidden)

	asbd := driver.AudioStreamBasicDescription()
	fmt.Printf("Format: %d channels, %.0f Hz, %d bits/channel\n",
		asbd.MChannelsPerFrame, asbd.MSampleRate, asbd.MBitsPerChannel)

	// 2. Client tracking test
	driver.AddClient(101, 1234)
	fmt.Printf("Attached Audio Clients: %d\n", driver.ClientCount())
	driver.RemoveClient(101)

	fmt.Println("\nStarting IO loopback engine...")
	if err := driver.StartIO(); err != nil {
		return fmt.Errorf("start IO: %w", err)
	}
	defer driver.StopIO()

	// Generate 480 frames of a 440 Hz test tone for 2 channels
	const numFrames = 480
	const freq = 440.0
	const sampleRate = 48000.0

	writeBuf := make([]float32, numFrames*2)
	for i := 0; i < numFrames; i++ {
		t := float64(i) / sampleRate
		val := float32(math.Sin(2 * math.Pi * freq * t))
		writeBuf[i*2] = val   // Left channel
		writeBuf[i*2+1] = val // Right channel
	}

	// 3. Pass-through audio loopback test
	sampleTime := uint64(1000)
	if err := driver.WriteMix(sampleTime, numFrames, writeBuf); err != nil {
		return fmt.Errorf("write mix: %w", err)
	}

	readBuf := make([]float32, numFrames*2)
	if err := driver.ReadInput(sampleTime, numFrames, readBuf); err != nil {
		return fmt.Errorf("read input: %w", err)
	}

	match := true
	for i := range writeBuf {
		if math.Abs(float64(writeBuf[i]-readBuf[i])) > 1e-6 {
			match = false
			break
		}
	}
	fmt.Printf("Audio Loopback 100%% Match: %v\n", match)

	// 4. Volume control test (set volume to 50%)
	driver.SetVolume(0.5)
	fmt.Printf("Set Volume: %.1f (dB: %.1f)\n", driver.Volume(), VolumeToDecibel(driver.Volume()))

	sampleTime += uint64(numFrames)
	if err := driver.WriteMix(sampleTime, numFrames, writeBuf); err != nil {
		return fmt.Errorf("write mix volume test: %w", err)
	}
	if err := driver.ReadInput(sampleTime, numFrames, readBuf); err != nil {
		return fmt.Errorf("read input volume test: %w", err)
	}

	fmt.Printf("Volume Scaled Sample [0] (written %.2f -> read %.2f): %v\n",
		writeBuf[2], readBuf[2], math.Abs(float64(readBuf[2]-writeBuf[2]*0.5)) < 1e-5)

	// 5. Pitch adjustment & Clock Source test
	driver.SetPitchAdjust(0.75)
	if err := driver.SetClockSource(ClockSourceInternalAdjustable); err != nil {
		return fmt.Errorf("set clock source: %w", err)
	}
	clockVal, _ := driver.GetPropertyData(ObjectIDClockSource, AudioObjectPropertyAddress{Selector: PropSelectorControlValue})
	fmt.Printf("Clock Source: %v (Pitch Adjust: %.2f)\n", clockVal, driver.PitchAdjust())

	// 6. Master Mute test
	driver.SetMute(true)
	fmt.Printf("Set Mute: %v\n", driver.Mute())

	sampleTime += uint64(numFrames)
	if err := driver.WriteMix(sampleTime, numFrames, writeBuf); err != nil {
		return fmt.Errorf("write mix mute test: %w", err)
	}
	if err := driver.ReadInput(sampleTime, numFrames, readBuf); err != nil {
		return fmt.Errorf("read input mute test: %w", err)
	}

	mutedAllZero := true
	for _, v := range readBuf {
		if v != 0 {
			mutedAllZero = false
			break
		}
	}
	fmt.Printf("Muted Buffer Zeroed Out: %v\n", mutedAllZero)

	// Reset mute & volume
	driver.SetMute(false)
	driver.SetVolume(1.0)

	// 7. Sample Rate Switch test
	if err := driver.SetSampleRate(96000.0); err != nil {
		return fmt.Errorf("set sample rate: %w", err)
	}
	fmt.Printf("Updated Sample Rate: %.0f Hz\n", driver.SampleRate())

	st, ht := driver.GetZeroTimeStamp()
	fmt.Printf("ZeroTimeStamp: sampleTime=%.0f (hostTime positive: %v)\n",
		st, ht > 0)

	// 8. CoreAudio Property Address Query
	addr := AudioObjectPropertyAddress{
		Selector: PropNominalSampleRate,
		Scope:    0x676c6f62, // 'glob' (Global)
		Element:  0,
	}
	fmt.Printf("Supports NominalSampleRate Property Query: %v\n",
		driver.HasProperty(coreaudio.AudioObjectID(ObjectIDDevice), addr))

	fmt.Println("BlackHole Virtual Audio Driver simulation finished successfully.")
	return nil
}
