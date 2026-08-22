// Command voltap demonstrates per-process audio volume control and muting using Core Audio process taps.
//
// Usage:
//
//	voltap [-bundle com.brave.Browser | -pid N] [-mute|-unmute] [-d 3s]
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/coreaudio"
	"github.com/tmc/apple/foundation"
)

func main() {
	pid := flag.Int("pid", 0, "Target process ID to control volume/mute")
	bundleID := flag.String("bundle", "", "Target bundle ID to control volume/mute (e.g. 'com.brave.Browser')")
	mute := flag.Bool("mute", false, "Mute the process audio output")
	unmute := flag.Bool("unmute", false, "Unmute the process audio output")
	duration := flag.Duration("d", 3*time.Second, "Hold duration for process tap before cleanup")
	flag.Parse()

	if *pid <= 0 && *bundleID == "" {
		*bundleID = "com.brave.Browser"
	}

	if err := run(*pid, *bundleID, *mute, *unmute, *duration); err != nil {
		fmt.Fprintf(os.Stderr, "voltap: %v\n", err)
		os.Exit(1)
	}
}

func run(pid int, bundleID string, mute, unmute bool, holdDuration time.Duration) error {
	pool := foundation.NewNSAutoreleasePool()
	defer pool.Drain()

	var tapDesc coreaudio.CATapDescription
	if bundleID != "" {
		fmt.Printf("Creating process tap for bundle %s...\n", bundleID)
		tapDesc = coreaudio.NewTapDescriptionStereoGlobalTapButExcludeProcesses(nil)
		tapDesc.SetBundleIDs([]string{bundleID})
	} else {
		fmt.Printf("Creating process tap for PID %d...\n", pid)
		procNum := foundation.NewNumberWithInt(int32(pid))
		tapDesc = coreaudio.NewTapDescriptionStereoMixdownOfProcesses([]foundation.NSNumber{procNum})
	}

	// Set mute behavior
	if mute {
		tapDesc.SetMuteBehavior(coreaudio.CATapMuted)
		fmt.Println("Setting mute behavior: Muted")
	} else if unmute {
		tapDesc.SetMuteBehavior(coreaudio.CATapUnmuted)
		fmt.Println("Setting mute behavior: Unmuted")
	} else {
		tapDesc.SetMuteBehavior(coreaudio.CATapUnmuted)
		fmt.Printf("Mute behavior: %v\n", tapDesc.MuteBehavior())
	}

	var tapID uint32
	status := coreaudio.AudioHardwareCreateProcessTap(&tapDesc, &tapID)
	if status != 0 {
		return fmt.Errorf("AudioHardwareCreateProcessTap failed: 0x%x (%d)", status, status)
	}
	defer func() {
		coreaudio.AudioHardwareDestroyProcessTap(tapID)
		fmt.Println("Process tap destroyed.")
	}()

	fmt.Printf("Active Process Tap ID: %d (UUID: %s)\n", tapID, tapDesc.UUID().UUIDString())
	fmt.Printf("Holding process audio tap for %v...\n", holdDuration)
	time.Sleep(holdDuration)

	return nil
}
