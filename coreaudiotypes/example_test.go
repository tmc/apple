//go:build darwin

package coreaudiotypes_test

import (
	"fmt"

	"github.com/tmc/apple/coreaudiotypes"
)

func ExampleAudioStreamBasicDescription() {
	// 44.1 kHz, 16-bit stereo LPCM audio stream description
	asbd := coreaudiotypes.AudioStreamBasicDescription{
		MSampleRate:       44100.0,
		MFormatID:         0x6c70636d, // 'lpcm'
		MFormatFlags:      12,         // kAudioFormatFlagIsSignedInteger | kAudioFormatFlagIsPacked
		MBytesPerPacket:   4,
		MFramesPerPacket:  1,
		MBytesPerFrame:    4,
		MChannelsPerFrame: 2,
		MBitsPerChannel:   16,
	}

	fmt.Printf("Sample Rate: %.1f Hz, Channels: %d, Bits/Channel: %d\n",
		asbd.MSampleRate, asbd.MChannelsPerFrame, asbd.MBitsPerChannel)

	// Output:
	// Sample Rate: 44100.0 Hz, Channels: 2, Bits/Channel: 16
}

func ExampleAudioValueRange() {
	valRange := coreaudiotypes.AudioValueRange{
		MMinimum: 20.0,
		MMaximum: 20000.0,
	}

	fmt.Printf("Frequency Range: %.0f Hz - %.0f Hz\n",
		valRange.MMinimum, valRange.MMaximum)

	// Output:
	// Frequency Range: 20 Hz - 20000 Hz
}

func ExampleAudioTimeStamp() {
	ts := coreaudiotypes.AudioTimeStamp{
		MSampleTime: 44100.0,
		MHostTime:   987654321,
		MRateScalar: 1.0,
	}

	fmt.Printf("SampleTime: %.1f, HostTime: %d, RateScalar: %.1f\n",
		ts.MSampleTime, ts.MHostTime, ts.MRateScalar)

	// Output:
	// SampleTime: 44100.0, HostTime: 987654321, RateScalar: 1.0
}
