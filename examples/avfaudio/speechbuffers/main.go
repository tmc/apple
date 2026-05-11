// Command speechbuffers renders an utterance into AVAudioBuffer callbacks.
//
// It demonstrates AVSpeechSynthesizer.WriteUtteranceToBufferCallback and pumps
// the CoreFoundation run loop until the first generated buffer arrives.
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	text := "hello from AVFAudio"
	if len(os.Args) > 1 {
		text = os.Args[1]
	}

	synth := avfaudio.NewAVSpeechSynthesizer()
	utterance := avfaudio.NewSpeechUtteranceWithString(text)
	done := make(chan avfaudio.AVAudioBuffer, 1)

	synth.WriteUtteranceToBufferCallback(utterance, func(buffer avfaudio.AVAudioBuffer) {
		select {
		case done <- buffer:
		default:
		}
	})

	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case buffer := <-done:
			fmt.Printf("buffer id=%d\n", buffer.GetID())
			return
		default:
			corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, 0.05, false)
		}
	}
	fmt.Fprintln(os.Stderr, "speech buffer callback did not fire")
	os.Exit(1)
}
