//go:build darwin

package ane

import "testing"

// TestShapeDoesNotExposeTensorState mirrors
// TestSurfaceSlicesDoNotExposeModelState (model_test.go): Shape() must
// return a defensive copy, not the tensor's live internal shape slice,
// matching the same fix already applied to Model.InputSurfaces()/
// OutputSurfaces(). Constructed via unexported fields (like model_test.go
// does for Model) rather than NewIOSurfaceTensor, since the latter requires
// a real IOSurface-backed CVPixelBuffer that this host may not provide.
func TestShapeDoesNotExposeTensorState(t *testing.T) {
	tensor := &IOSurfaceTensor{shape: []int{1, 3, 4, 4}}

	shape := tensor.Shape()
	shape[0] = 99
	shape = append(shape, 100)

	got := tensor.Shape()
	if len(got) != 4 || got[0] != 1 || got[1] != 3 || got[2] != 4 || got[3] != 4 {
		t.Fatalf("Shape() exposed internal state: %v", got)
	}
}
