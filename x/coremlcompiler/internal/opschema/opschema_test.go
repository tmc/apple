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
	// batch_norm is a real MIL op (ops/defs/iOS15/normalization.py) that no
	// emitter here produces, so it stands in for "outside the dump". Registered
	// must stay false for it: if this ever fires because the op was added to
	// gen/dumpops.py, pick another op rather than deleting the check, or
	// Registered loses its only negative control.
	if Registered("batch_norm") {
		t.Error(`Registered("batch_norm") = true; batch_norm is outside the dump, want false`)
	}
}

// TestRegisteredStateOps covers the two state ops that are not in coremltools'
// core op registry: coreml_update_state lives in the "coreml" dialect namespace
// and write_state has no Operation subclass at all, so both reach the table by a
// path gen/dumpops.py has to take deliberately.
func TestRegisteredStateOps(t *testing.T) {
	for _, opType := range []string{"write_state", "coreml_update_state"} {
		op, ok := Lookup("ios18", opType)
		if !ok {
			t.Errorf("Lookup(ios18, %q) not ok, want ok", opType)
			continue
		}
		// State ops arrived with iOS18; the older opsets must not claim them.
		if _, ok := Lookup("ios17", opType); ok {
			t.Errorf("Lookup(ios17, %q) ok, want not ok", opType)
		}
		if len(op.Params) != 2 {
			t.Errorf("%s has %d params, want 2", opType, len(op.Params))
		}
		if op.Params[0].Kind != KindState {
			t.Errorf("%s param %q kind = %v, want KindState", opType, op.Params[0].Name, op.Params[0].Kind)
		}
	}
	// write_state's input names come from the MIL backend's lowering of
	// coreml_update_state (backend/mil/load.py:313-322), not from its source
	// op's names; a rename that silently no-oped would leave state/value here.
	ws, _ := Lookup("ios18", "write_state")
	for _, name := range []string{"input", "data"} {
		if _, ok := ws.Param(name); !ok {
			t.Errorf("write_state has no input %q; params are %v", name, paramNames(ws))
		}
	}
	for _, name := range []string{"state", "value"} {
		if _, ok := ws.Param(name); ok {
			t.Errorf("write_state has input %q; it should have been renamed", name)
		}
	}
}

func paramNames(op *Op) []string {
	names := make([]string, len(op.Params))
	for i, p := range op.Params {
		names[i] = p.Name
	}
	return names
}

func TestLookupUnknown(t *testing.T) {
	if _, ok := Lookup("ios26", "conv"); ok {
		t.Error("Lookup(ios26, conv) ok, want not ok")
	}
	if _, ok := Lookup("ios18", "no_such_op"); ok {
		t.Error("Lookup(ios18, no_such_op) ok, want not ok")
	}
}
