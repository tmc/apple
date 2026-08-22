//go:build darwin

package metalkit_test

import (
	"fmt"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/metalkit"
)

func ExampleMTKView() {
	view := metalkit.NewMTKView()
	view.SetPaused(true)
	view.SetEnableSetNeedsDisplay(true)
	view.SetPreferredFramesPerSecond(60)
	view.SetClearColor(metal.MTLClearColor{Red: 0.0, Green: 0.5, Blue: 1.0, Alpha: 1.0})
	view.SetClearDepth(1.0)

	fmt.Printf("Paused: %t\n", view.IsPaused())
	fmt.Printf("EnableSetNeedsDisplay: %t\n", view.EnableSetNeedsDisplay())
	fmt.Printf("PreferredFPS: %d\n", view.PreferredFramesPerSecond())
	color := view.ClearColor()
	fmt.Printf("ClearColor: Red=%.1f Green=%.1f Blue=%.1f Alpha=%.1f\n", color.Red, color.Green, color.Blue, color.Alpha)
	fmt.Printf("ClearDepth: %.1f\n", view.ClearDepth())

	// Output:
	// Paused: true
	// EnableSetNeedsDisplay: true
	// PreferredFPS: 60
	// ClearColor: Red=0.0 Green=0.5 Blue=1.0 Alpha=1.0
	// ClearDepth: 1.0
}

func ExampleMTKTextureLoaderOptions() {
	fmt.Println(metalkit.MTKTextureLoaderOptions.AllocateMipmaps)
	fmt.Println(metalkit.MTKTextureLoaderOrigins.BottomLeft)

	// Output:
	// MTKTextureLoaderOptionAllocateMipmaps
	// MTKTextureLoaderOriginBottomLeft
}
