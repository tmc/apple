//go:build darwin

package model

import (
	"testing"

	xane "github.com/tmc/apple/x/ane"
)

func TestSymbolIndexAt(t *testing.T) {
	layouts := func(idx ...int) []xane.TensorLayout {
		out := make([]xane.TensorLayout, len(idx))
		for i, v := range idx {
			out[i] = xane.TensorLayout{SymbolIndex: v}
		}
		return out
	}

	tests := []struct {
		name    string
		layouts []xane.TensorLayout
		i       int
		want    int
	}{
		{"resolved index is used", layouts(2, 0, 1), 0, 2},
		{"resolved index differing from position", layouts(2, 0, 1), 2, 1},
		{"positional layouts are unchanged", layouts(0, 1, 2), 1, 1},
		{"unresolved falls back to position", layouts(0, -1, 2), 1, 1},
		{"index past the layouts falls back", layouts(5), 3, 3},
		{"no layouts falls back", nil, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := symbolIndexAt(tt.layouts, tt.i); got != tt.want {
				t.Errorf("symbolIndexAt(%v, %d) = %d, want %d", tt.layouts, tt.i, got, tt.want)
			}
		})
	}
}
