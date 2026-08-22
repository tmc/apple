//go:build darwin

package appkit_test

import (
	"fmt"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
)

func ExampleNSColor() {
	color := appkit.NewColorWithRedGreenBlueAlpha(1.0, 0.0, 0.0, 1.0)
	fmt.Printf("R: %.1f, G: %.1f, B: %.1f, A: %.1f\n",
		color.RedComponent(),
		color.GreenComponent(),
		color.BlueComponent(),
		color.AlphaComponent(),
	)

	// Output:
	// R: 1.0, G: 0.0, B: 0.0, A: 1.0
}

func ExampleNSWorkspace() {
	ws := appkit.GetNSWorkspaceClass().SharedWorkspace()
	fmt.Printf("IsFilePackageAtPath(/System): %t\n", ws.IsFilePackageAtPath("/System"))

	// Output:
	// IsFilePackageAtPath(/System): false
}

func ExampleNSView() {
	rect := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 10, Y: 20},
		Size:   corefoundation.CGSize{Width: 200, Height: 100},
	}
	view := appkit.NewViewWithFrame(rect)
	frame := view.Frame()
	fmt.Printf("View frame: origin=(%.1f, %.1f), size=(%.1f, %.1f)\n",
		frame.Origin.X, frame.Origin.Y, frame.Size.Width, frame.Size.Height)

	// Output:
	// View frame: origin=(10.0, 20.0), size=(200.0, 100.0)
}
