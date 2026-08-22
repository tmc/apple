//go:build darwin

package speech_test

import (
	"fmt"

	"github.com/tmc/apple/speech"
)

func ExampleSFSpeechRecognizer() {
	recognizer := speech.NewSFSpeechRecognizer()
	recognizer.SetSupportsOnDeviceRecognition(true)
	recognizer.SetDefaultTaskHint(speech.SFSpeechRecognitionTaskHintDictation)

	fmt.Println("SupportsOnDeviceRecognition:", recognizer.SupportsOnDeviceRecognition())
	fmt.Println("DefaultTaskHint:", recognizer.DefaultTaskHint())

	// Output:
	// SupportsOnDeviceRecognition: true
	// DefaultTaskHint: SFSpeechRecognitionTaskHintDictation
}

func ExampleSFSpeechAudioBufferRecognitionRequest() {
	req := speech.NewSFSpeechAudioBufferRecognitionRequest()
	req.SetShouldReportPartialResults(true)
	req.SetAddsPunctuation(true)
	req.SetTaskHint(speech.SFSpeechRecognitionTaskHintSearch)

	fmt.Println("ShouldReportPartialResults:", req.ShouldReportPartialResults())
	fmt.Println("AddsPunctuation:", req.AddsPunctuation())
	fmt.Println("TaskHint:", req.TaskHint())

	// Output:
	// ShouldReportPartialResults: true
	// AddsPunctuation: true
	// TaskHint: SFSpeechRecognitionTaskHintSearch
}
