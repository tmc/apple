// Command ane_fwtofw demonstrates configuring firmware-to-firmware (FWToFW) signal options
// and hardware cross-synchronization options between ANE and GPU.
package main

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/private/appleneuralengine"
	"github.com/tmc/apple/x/ane"
)

func main() {
	fmt.Println("=== Apple Neural Engine (ANE) FWToFW Hardware Signaling Demo ===")

	// 1. Inspect global key KANEFEnableFWToFWSignal
	fwKey := appleneuralengine.KANEFEnableFWToFWSignal
	fmt.Printf("[1/3] ANE Private Framework Global Symbol kANEFEnableFWToFWSignal: %v\n", fwKey)

	// 2. Build ANE Option Dictionary with FWToFW Signal enabled
	fmt.Println("[2/3] Constructing ANE Evaluation Option Dictionary...")
	opts := foundation.NewNSMutableDictionary()

	// Enable firmware-to-firmware hardware signaling flag
	opts.SetObjectForKey(
		foundation.GetNSNumberClass().NumberWithBool(true),
		appleneuralengine.KANEFEnableFWToFWSignal,
	)

	fmt.Printf("      -> Configured kANEFEnableFWToFWSignal = true in ANE request options.\n")

	// 3. Create Hardware Shared Signal/Wait Events for ANE <-> Metal GPU sync
	fmt.Println("[3/3] Setting up Hardware Synchronization Events...")
	dev := metal.MTLCreateSystemDefaultDevice()
	if dev.ID != 0 {
		sharedEvent := dev.NewSharedEvent()
		fmt.Printf("      -> Created Metal MTLSharedEvent: %v\n", sharedEvent)

		// ANE Shared Signal Event creation
		sigEventClass := appleneuralengine.GetANESharedSignalEventClass()
		fmt.Printf("      -> _ANESharedSignalEvent ObjC Class: %v\n", sigEventClass.Class())
	}

	// 4. Reference high-level x/ane package integration options
	sharedOpts := ane.SharedEventEvalOptions{
		EnableFWToFWSignal:             true,
		DisableIOFencesUseSharedEvents: true,
	}
	fmt.Printf("\n[x/ane] Configured SharedEventEvalOptions: EnableFWToFWSignal=%v, DisableIOFencesUseSharedEvents=%v\n",
		sharedOpts.EnableFWToFWSignal, sharedOpts.DisableIOFencesUseSharedEvents)

	fmt.Println("\n=== FWToFW Hardware Signaling Option Setup Complete ===")
}
