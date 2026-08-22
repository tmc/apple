//go:build darwin

package speechobjects_test

import (
	"fmt"

	"github.com/tmc/apple/private/speechobjects"
)

func ExampleGetSOSpeechItemClass() {
	cls := speechobjects.GetSOSpeechItemClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: speechobjects.SOSpeechItemClass
}
