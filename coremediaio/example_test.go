//go:build darwin

package coremediaio_test

import (
	"fmt"

	"github.com/tmc/apple/coremediaio"
)

func ExampleCMIOObjectPropertyAddress() {
	address := coremediaio.CMIOObjectPropertyAddress{
		MSelector: 0x676c6f62, // 'glob'
		MScope:    0x676c6f62, // 'glob'
		MElement:  0,
	}

	fmt.Printf("Selector: 0x%x, Scope: 0x%x, Element: %d\n",
		address.MSelector, address.MScope, address.MElement)

	// Output:
	// Selector: 0x676c6f62, Scope: 0x676c6f62, Element: 0
}

func ExampleCMIODeviceStreamConfiguration() {
	config := coremediaio.CMIODeviceStreamConfiguration{
		MNumberStreams: 2,
	}

	fmt.Printf("Streams: %d\n", config.MNumberStreams)

	// Output:
	// Streams: 2
}

func ExampleCMIOStreamDeck() {
	deck := coremediaio.CMIOStreamDeck{
		MStatus: 1,
		MState:  0,
		MState2: 0,
	}

	fmt.Printf("Deck Status: %d, State: %d\n", deck.MStatus, deck.MState)

	// Output:
	// Deck Status: 1, State: 0
}
