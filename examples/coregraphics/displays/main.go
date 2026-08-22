// Command displays lists the attached displays and their current
// configuration using the Core Graphics display services API.
//
// Usage: displays
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/coregraphics"
)

// maxDisplays is the size of the display list buffer. Systems with more
// displays than this are not expected.
const maxDisplays = 32

func main() {
	ids := make([]uint32, maxDisplays)
	var count uint32
	if err := coregraphics.CGGetOnlineDisplayList(maxDisplays, &ids[0], &count); err != coregraphics.KCGErrorSuccess {
		fmt.Fprintf(os.Stderr, "CGGetOnlineDisplayList: %v\n", err)
		os.Exit(1)
	}
	mainID := coregraphics.CGMainDisplayID()
	for _, id := range ids[:count] {
		b := coregraphics.CGDisplayBounds(id)
		mm := coregraphics.CGDisplayScreenSize(id)
		fmt.Printf("display %d\n", id)
		fmt.Printf("\tbounds:    %.0f,%.0f %.0fx%.0f\n", b.Origin.X, b.Origin.Y, b.Size.Width, b.Size.Height)
		fmt.Printf("\tpixels:    %dx%d\n", coregraphics.CGDisplayPixelsWide(id), coregraphics.CGDisplayPixelsHigh(id))
		fmt.Printf("\tphysical:  %.0fx%.0f mm\n", mm.Width, mm.Height)
		fmt.Printf("\trotation:  %g degrees\n", coregraphics.CGDisplayRotation(id))
		fmt.Printf("\tmain:      %v\n", id == mainID)
		fmt.Printf("\tbuiltin:   %v\n", coregraphics.CGDisplayIsBuiltin(id) != 0)
		fmt.Printf("\tactive:    %v\n", coregraphics.CGDisplayIsActive(id) != 0)
		fmt.Printf("\tasleep:    %v\n", coregraphics.CGDisplayIsAsleep(id) != 0)
		fmt.Printf("\tvendor:    0x%x model 0x%x serial 0x%x\n",
			coregraphics.CGDisplayVendorNumber(id),
			coregraphics.CGDisplayModelNumber(id),
			coregraphics.CGDisplaySerialNumber(id))
		if mode := coregraphics.CGDisplayCopyDisplayMode(id); mode != 0 {
			fmt.Printf("\tmode:      %dx%d (%dx%d pixels) at %g Hz\n",
				coregraphics.CGDisplayModeGetWidth(mode),
				coregraphics.CGDisplayModeGetHeight(mode),
				coregraphics.CGDisplayModeGetPixelWidth(mode),
				coregraphics.CGDisplayModeGetPixelHeight(mode),
				coregraphics.CGDisplayModeGetRefreshRate(mode))
			coregraphics.CGDisplayModeRelease(mode)
		}
	}
}
