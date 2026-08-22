//go:build darwin

package avfoundation_test

import (
	"fmt"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/foundation"
)

func ExampleAVPlayer() {
	url := foundation.NewURLWithString("file:///tmp/sample.mp4")
	player := avfoundation.NewPlayerWithURL(url)
	player.SetVolume(0.8)

	fmt.Printf("Volume: %.1f\n", player.Volume())
	fmt.Printf("Rate: %.1f\n", player.Rate())
	// Output:
	// Volume: 0.8
	// Rate: 0.0
}

func ExampleAVAsset() {
	url := foundation.NewURLWithString("file:///tmp/sample.mp4")
	asset := avfoundation.NewURLAssetWithURLOptions(url, nil)

	fmt.Println("URL:", asset.URL().AbsoluteString())
	fmt.Println("Reference restrictions:", asset.ReferenceRestrictions())
	// Output:
	// URL: file:///tmp/sample.mp4
	// Reference restrictions: AVAssetReferenceRestrictionForbidNone
}

func ExampleAVCaptureSession() {
	session := avfoundation.NewAVCaptureSession()
	session.SetSessionPreset(avfoundation.AVCaptureSessionPresets.High)
	session.BeginConfiguration()
	running := session.IsRunning()
	session.CommitConfiguration()

	fmt.Println("Preset:", session.SessionPreset())
	fmt.Println("Running:", running)
	// Output:
	// Preset: AVCaptureSessionPresetHigh
	// Running: false
}
