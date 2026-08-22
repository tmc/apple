// Command padstate prints the state of connected game controllers.
//
// It lists every controller the Game Controller framework knows about,
// reports battery, motion and haptics availability, and then streams
// thumbstick, trigger and button state to the terminal until the deadline
// expires or the process is interrupted.
//
// Usage: padstate [-d duration] [-r rate]
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/gamecontroller"
)

func main() {
	// The Game Controller framework delivers input on the run loop of the
	// thread that reads it, so keep this goroutine pinned.
	runtime.LockOSThread()

	duration := flag.Duration("d", 10*time.Second, "how long to stream controller state")
	rate := flag.Duration("r", 250*time.Millisecond, "sampling interval")
	flag.Parse()

	controllers := gamecontroller.GetGCControllerClass().Controllers()
	if len(controllers) == 0 {
		fmt.Fprintf(os.Stderr, "no game controllers found; pair or connect one and re-run\n")
		os.Exit(1)
	}

	for i, c := range controllers {
		describe(i, c)
	}

	if *duration <= 0 {
		return
	}
	fmt.Printf("\nstreaming for %v (interval %v)\n\n", *duration, *rate)
	deadline := time.Now().Add(*duration)
	for time.Now().Before(deadline) {
		// Pump the run loop so the framework can refresh element values.
		corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, corefoundation.CFTimeInterval(rate.Seconds()), false)
		for i, c := range controllers {
			line := sample(c)
			if line == "" {
				continue
			}
			fmt.Printf("[%d] %s\n", i, line)
		}
	}
}

// describe prints the static capabilities of a controller.
func describe(i int, c gamecontroller.GCController) {
	fmt.Printf("[%d] %s (%s)\n", i, c.VendorName(), c.ProductCategory())
	fmt.Printf("     player index: %v  attached to device: %t\n", c.PlayerIndex(), c.IsAttachedToDevice())

	if b := c.Battery(); b.GetID() != 0 {
		fmt.Printf("     battery: %.0f%% (%v)\n", b.BatteryLevel()*100, b.BatteryState())
	} else {
		fmt.Printf("     battery: not reported\n")
	}

	if m := c.Motion(); m.GetID() != 0 {
		fmt.Printf("     motion: attitude=%t rotationRate=%t gravity=%t\n",
			m.HasAttitude(), m.HasRotationRate(), m.HasGravityAndUserAcceleration())
	} else {
		fmt.Printf("     motion: unavailable\n")
	}

	fmt.Printf("     haptics: %t\n", c.Haptics().GetID() != 0)
	fmt.Printf("     light: %t\n", c.Light().GetID() != 0)

	if c.ExtendedGamepad().GetID() == 0 {
		fmt.Printf("     profile: not an extended gamepad; stick and button state unavailable\n")
	}
}

// sample formats one line of live state for a controller, or "" if the
// controller does not expose the extended gamepad profile.
func sample(c gamecontroller.GCController) string {
	g := c.ExtendedGamepad()
	if g.GetID() == 0 {
		return ""
	}
	ls, rs := g.LeftThumbstick(), g.RightThumbstick()
	var pressed []string
	for _, b := range []struct {
		name   string
		button gamecontroller.IGCControllerButtonInput
	}{
		{"A", g.ButtonA()},
		{"B", g.ButtonB()},
		{"X", g.ButtonX()},
		{"Y", g.ButtonY()},
		{"LS", g.LeftShoulder()},
		{"RS", g.RightShoulder()},
		{"Menu", g.ButtonMenu()},
		{"Options", g.ButtonOptions()},
		{"L3", g.LeftThumbstickButton()},
		{"R3", g.RightThumbstickButton()},
	} {
		if b.button.GetID() != 0 && b.button.IsPressed() {
			pressed = append(pressed, b.name)
		}
	}
	dpad := g.Dpad()
	return fmt.Sprintf("L(%+.2f,%+.2f) R(%+.2f,%+.2f) LT=%.2f RT=%.2f dpad(%+.0f,%+.0f) buttons=[%s]",
		ls.XAxis().Value(), ls.YAxis().Value(),
		rs.XAxis().Value(), rs.YAxis().Value(),
		g.LeftTrigger().Value(), g.RightTrigger().Value(),
		dpad.XAxis().Value(), dpad.YAxis().Value(),
		strings.Join(pressed, " "))
}
