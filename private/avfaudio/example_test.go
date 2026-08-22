//go:build darwin

package avfaudio_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/avfaudio"
)

func ExampleGetAVSpeechSynthesizerClass() {
	cls := avfaudio.GetAVSpeechSynthesizerClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: AVSpeechSynthesizer
}

func ExampleGetAVAudioEngineClass() {
	cls := avfaudio.GetAVAudioEngineClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: AVAudioEngine
}
