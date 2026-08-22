//go:build darwin

package texttospeech_test

import (
	"fmt"

	"github.com/tmc/apple/private/texttospeech"
)

func ExampleGetTTSSpeechManagerClass() {
	cls := texttospeech.GetTTSSpeechManagerClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: texttospeech.TTSSpeechManagerClass
}
