package coremlcompiler

import "testing"

// TestMILOpSet pins the CoreMLN → iosNN pairing to coremltools' _OPSET table
// (coremltools/__init__.py), where CoreMLN is the opset of spec version N+1,
// i.e. the iOS N+10 release.
func TestMILOpSet(t *testing.T) {
	tests := []struct {
		opset string
		want  string
	}{
		{"CoreML3", "ios13"},
		{"CoreML4", "ios14"},
		{"CoreML5", "ios15"},
		{"CoreML6", "ios16"},
		{"CoreML7", "ios17"},
		{"CoreML8", "ios18"},
		{"CoreML9", "ios26"},
	}
	for _, tt := range tests {
		if got := milOpSet(tt.opset); got != tt.want {
			t.Errorf("milOpSet(%q) = %q, want %q", tt.opset, got, tt.want)
		}
	}
}
