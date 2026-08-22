//go:build darwin

package coreaudio_test

import (
	"fmt"

	"github.com/tmc/apple/coreaudio"
)

func ExampleAudioGetCurrentHostTime() {
	now := coreaudio.AudioGetCurrentHostTime()
	freq := coreaudio.AudioGetHostClockFrequency()

	fmt.Println("Host time positive:", now > 0)
	fmt.Printf("Clock frequency: %.0f\n", freq)
	// Output:
	// Host time positive: true
	// Clock frequency: 24000000
}

func ExampleAudioConvertHostTimeToNanos() {
	nanos := uint64(1_000_000_000) // 1 second
	hostTime := coreaudio.AudioConvertNanosToHostTime(nanos)
	convertedNanos := coreaudio.AudioConvertHostTimeToNanos(hostTime)

	fmt.Println("Input nanos:", nanos)
	fmt.Println("Converted nanos:", convertedNanos)
	// Output:
	// Input nanos: 1000000000
	// Converted nanos: 1000000000
}

func ExampleAudioObjectPropertyAddress() {
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: 0x676c6f62, // 'glob'
		MScope:    0x676c6f62, // 'glob'
		MElement:  0,
	}

	fmt.Printf("Selector: 0x%x\n", addr.MSelector)
	fmt.Printf("Scope: 0x%x\n", addr.MScope)
	fmt.Println("Element:", addr.MElement)
	// Output:
	// Selector: 0x676c6f62
	// Scope: 0x676c6f62
	// Element: 0
}
