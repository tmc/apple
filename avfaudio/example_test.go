//go:build darwin

package avfaudio_test

import (
	"fmt"

	"github.com/tmc/apple/avfaudio"
)

func ExampleAVAudioEngine() {
	engine := avfaudio.NewAVAudioEngine()
	player := avfaudio.NewAVAudioPlayerNode()

	engine.AttachNode(player)

	fmt.Println("Attached nodes:", engine.AttachedNodes().Count())
	fmt.Println("Engine running:", engine.IsRunning())
	fmt.Println("Player playing:", player.IsPlaying())
	// Output:
	// Attached nodes: 1
	// Engine running: false
	// Player playing: false
}

func ExampleAVAudioFormat() {
	format := avfaudio.NewAudioFormatStandardFormatWithSampleRateChannels(44100, 2)

	fmt.Printf("Sample rate: %.1f\n", format.SampleRate())
	fmt.Println("Channels:", format.ChannelCount())
	fmt.Println("Is interleaved:", format.IsInterleaved())
	// Output:
	// Sample rate: 44100.0
	// Channels: 2
	// Is interleaved: false
}

func ExampleAVAudioPCMBuffer() {
	format := avfaudio.NewAudioFormatStandardFormatWithSampleRateChannels(44100, 2)
	buffer := avfaudio.NewAudioPCMBufferWithPCMFormatFrameCapacity(format, 1024)
	buffer.SetFrameLength(512)

	fmt.Println("Frame capacity:", buffer.FrameCapacity())
	fmt.Println("Frame length:", buffer.FrameLength())
	// Output:
	// Frame capacity: 1024
	// Frame length: 512
}
