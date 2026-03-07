package main

import (
	"fmt"

	"github.com/tmc/apple/foundation"
)

func main() {
	pi := foundation.GetProcessInfoClass().ProcessInfo()
	fmt.Printf("pid=%d\n", pi.ProcessIdentifier())
	fmt.Printf("name=%s\n", pi.ProcessName())
	fmt.Printf("host=%s\n", pi.HostName())
	fmt.Printf("os=%s\n", pi.OperatingSystemVersionString())
	fmt.Printf("args=%d\n", len(pi.Arguments()))
}
