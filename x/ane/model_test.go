//go:build darwin

package ane

import (
	"testing"

	"github.com/tmc/apple/coregraphics"
)

func TestSurfaceSlicesDoNotExposeModelState(t *testing.T) {
	m := &Model{
		inputs:  []coregraphics.IOSurfaceRef{1, 2},
		outputs: []coregraphics.IOSurfaceRef{3, 4},
	}

	inputs := m.InputSurfaces()
	inputs[0] = 9
	inputs = append(inputs, 10)
	outputs := m.OutputSurfaces()
	outputs[0] = 11
	outputs = append(outputs, 12)

	if got := m.InputSurfaces(); len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("input surfaces exposed internal state: %v", got)
	}
	if got := m.OutputSurfaces(); len(got) != 2 || got[0] != 3 || got[1] != 4 {
		t.Fatalf("output surfaces exposed internal state: %v", got)
	}
}
