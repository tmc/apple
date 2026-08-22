//go:build darwin

package quartzcore_test

import (
	"fmt"

	"github.com/tmc/apple/quartzcore"
)

func ExampleCACurrentMediaTime() {
	fr := quartzcore.CAFrameRateRangeMake(30.0, 120.0, 60.0)
	fmt.Printf("FrameRateRange: min=%.0f max=%.0f pref=%.0f\n", fr.Minimum, fr.Maximum, fr.Preferred)

	t := quartzcore.CACurrentMediaTime()
	fmt.Printf("CACurrentMediaTime > 0: %t\n", t > 0)

	// Output:
	// FrameRateRange: min=30 max=120 pref=60
	// CACurrentMediaTime > 0: true
}

func ExampleNewCALayer() {
	layer := quartzcore.NewCALayer()
	layer.SetOpacity(0.8)
	layer.SetCornerRadius(5.0)
	layer.SetHidden(false)
	layer.SetMasksToBounds(true)

	fmt.Printf("Opacity: %.1f\n", layer.Opacity())
	fmt.Printf("CornerRadius: %.1f\n", layer.CornerRadius())
	fmt.Printf("Hidden: %t\n", layer.IsHidden())
	fmt.Printf("MasksToBounds: %t\n", layer.MasksToBounds())

	// Output:
	// Opacity: 0.8
	// CornerRadius: 5.0
	// Hidden: false
	// MasksToBounds: true
}

func ExampleCATransform3DIsIdentity() {
	identity := quartzcore.CATransform3DIdentity
	scale := quartzcore.CATransform3DMakeScale(2.0, 3.0, 1.0)

	fmt.Printf("Identity is identity: %t\n", quartzcore.CATransform3DIsIdentity(identity))
	fmt.Printf("Scale is identity: %t\n", quartzcore.CATransform3DIsIdentity(scale))
	fmt.Printf("Scale diagonal: %.1f, %.1f, %.1f, %.1f\n", scale.M11, scale.M22, scale.M33, scale.M44)

	// Output:
	// Identity is identity: true
	// Scale is identity: false
	// Scale diagonal: 2.0, 3.0, 1.0, 1.0
}
