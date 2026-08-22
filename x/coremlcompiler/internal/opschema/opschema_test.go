package opschema

import (
	"reflect"
	"testing"
)

// TestLookupConv pins one op exactly. A table that regenerated empty or
// truncated would otherwise make every program validate clean.
func TestLookupConv(t *testing.T) {
	op, ok := Lookup("ios18", "conv")
	if !ok {
		t.Fatal("conv not found in ios18")
	}
	want := []string{"x", "weight", "bias", "strides", "pad_type", "pad", "dilations", "groups"}
	var got []string
	for _, p := range op.Params {
		got = append(got, p.Name)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("conv params = %v, want %v", got, want)
	}
	if p, _ := op.Param("pad_type"); p == nil || !p.Const || !p.Optional {
		t.Errorf("pad_type = %+v, want const optional", p)
	}
	if p, _ := op.Param("x"); p == nil || p.Const {
		t.Errorf("x = %+v, want non-const", p)
	}
}

// TestConvWeightDomainByOpset guards the opset-sensitivity of the table: conv
// shares one type domain between x and weight at iOS15 (ops/defs/iOS15/conv.py:129)
// and splits weight into its own domain from iOS17 on.
func TestConvWeightDomainByOpset(t *testing.T) {
	for _, tc := range []struct{ opset, want string }{
		{"ios15", "T"},
		{"ios16", "T"},
		{"ios17", "U"},
		{"ios18", "U"},
	} {
		op, ok := Lookup(tc.opset, "conv")
		if !ok {
			t.Fatalf("conv not found in %s", tc.opset)
		}
		p, _ := op.Param("weight")
		if p == nil || p.DomainID != tc.want {
			t.Errorf("%s conv weight domain = %+v, want %q", tc.opset, p, tc.want)
		}
	}
}

// TestKnown asserts the opsets we emit resolve, so the validator's skip path
// for unknown opsets cannot quietly become universal.
func TestKnown(t *testing.T) {
	for _, opset := range []string{"ios15", "ios16", "ios17", "ios18"} {
		if !Known(opset) {
			t.Errorf("Known(%q) = false, want true", opset)
		}
	}
	// The dump stops at iOS18 (ops/registry.py:56-62), so ios26 has no schema.
	if Known("ios26") {
		t.Error(`Known("ios26") = true, want false`)
	}
}

func TestRegistered(t *testing.T) {
	if !Registered("conv") {
		t.Error(`Registered("conv") = false, want true`)
	}
	if Registered("linear") {
		t.Error(`Registered("linear") = true; linear is outside the dump, want false`)
	}
}

func TestLookupUnknown(t *testing.T) {
	if _, ok := Lookup("ios26", "conv"); ok {
		t.Error("Lookup(ios26, conv) ok, want not ok")
	}
	if _, ok := Lookup("ios18", "no_such_op"); ok {
		t.Error("Lookup(ios18, no_such_op) ok, want not ok")
	}
}
