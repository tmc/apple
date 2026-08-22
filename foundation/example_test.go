//go:build darwin

package foundation_test

import (
	"fmt"

	"github.com/tmc/apple/foundation"
)

func ExampleNSString() {
	str := foundation.NewStringWithString("Apple Frameworks")
	fmt.Println(str.Length())
	fmt.Println(str.SubstringToIndex(5))
	fmt.Println(str.SubstringFromIndex(6))
	fmt.Println(str.StringByAppendingString(" for Go"))

	// Output:
	// 16
	// Apple
	// Frameworks
	// Apple Frameworks for Go
}

func ExampleProcessInfo() {
	info := foundation.GetProcessInfoClass().ProcessInfo()
	origName := info.ProcessName()
	info.SetProcessName("ExampleApp")
	fmt.Println(info.ProcessName())
	info.SetProcessName(origName)

	// Output:
	// ExampleApp
}
