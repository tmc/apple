//go:build darwin

package os_test

import (
	"fmt"

	"github.com/tmc/apple/os"
)

func ExampleOSLogCreate() {
	log := os.OSLogCreate("com.example.apple", "network")
	enabled := os.OSLogTypeEnabled(log, os.OSLogTypeDefault)
	fmt.Printf("Default log level enabled: %t\n", enabled)
	// Output:
	// Default log level enabled: true
}

func ExampleOSLogType() {
	fmt.Println(os.OSLogTypeDefault)
	fmt.Println(os.OSLogTypeInfo)
	fmt.Println(os.OSLogTypeDebug)
	fmt.Println(os.OSLogTypeError)
	fmt.Println(os.OSLogTypeFault)
	// Output:
	// OSLogTypeDefault
	// OSLogTypeInfo
	// OSLogTypeDebug
	// OSLogTypeError
	// OSLogTypeFault
}

func ExampleOSSignpostType() {
	fmt.Println(os.OSSignpostEvent)
	fmt.Println(os.OSSignpostIntervalBegin)
	fmt.Println(os.OSSignpostIntervalEnd)
	// Output:
	// OSSignpostEvent
	// OSSignpostIntervalBegin
	// OSSignpostIntervalEnd
}

func ExampleNewOSLog() {
	log := os.NewOSLog()
	fmt.Printf("OSLog zero-value handle: %t\n", log.GetID() == 0)
	// Output:
	// OSLog zero-value handle: true
}
