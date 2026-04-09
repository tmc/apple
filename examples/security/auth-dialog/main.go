// Command auth-dialog triggers the macOS authentication dialog using
// LocalAuthentication's LAContext. The system handles the entire UI:
// Touch ID prompt, password fallback, icon, and button styling.
//
// The dialog title reflects the executable name. To match the system
// Developer Tools Access prompt exactly, build with a custom output name:
//
//	go build -o "Developer Tools Access" ./examples/security/auth-dialog/
//	./"Developer Tools Access"
package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/tmc/apple/localauthentication"
)

func init() { runtime.LockOSThread() }

func main() {
	ctx := localauthentication.NewLAContext()

	ok, err := ctx.CanEvaluatePolicyError(localauthentication.LAPolicyDeviceOwnerAuthentication)
	if !ok {
		fmt.Fprintf(os.Stderr, "cannot evaluate policy: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	ctx.EvaluatePolicyLocalizedReasonReply(
		localauthentication.LAPolicyDeviceOwnerAuthentication,
		"take control of another process for debugging",
		func(success bool, err error) {
			defer wg.Done()
			if success {
				fmt.Println("authenticated")
			} else {
				fmt.Fprintf(os.Stderr, "authentication failed: %v\n", err)
			}
		},
	)

	wg.Wait()
}
