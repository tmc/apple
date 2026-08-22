//go:build darwin

package mil_test

import (
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// TestPublicStateProgramsDoNotRun records the public-model API's state
// limitation and fails if that changes.
//
// It is a tripwire rather than a check on the generators. The text they emit is
// valid MIL. The lower-level e5rt route binds the state as a named inout port;
// that execution path is covered by x/ane/e5rt's subprocess probe. This test
// records the distinct public API limitation, which does not expose the port.
//
// See the doc comments on [mil.GenReadState] and [mil.GenUpdateState].
func TestPublicStateProgramsDoNotRun(t *testing.T) {
	c, err := ane.Open()
	if err != nil {
		t.Skipf("no ANE client: %v", err)
	}
	defer c.Close()

	shape := [4]int{1, 8, 1, 16}

	// A program whose only parameter is a state has no live inputs, and the
	// layout parser rejects it before the compiler sees it.
	if m, err := c.Compile(ane.CompileOptions{
		ModelType: ane.ModelTypeMIL,
		MILText:   []byte(mil.GenReadState("kv", shape)),
	}); err == nil {
		m.Close()
		t.Error("GenReadState now compiles through the public API; state models may have become usable, so update the doc comments on GenReadState and GenUpdateState and consider writing the KV cache example")
	} else if !strings.Contains(err.Error(), "LiveInputList") {
		t.Logf("GenReadState still fails to compile, but with a different error than recorded: %v", err)
	}

	// A program with a tensor input as well compiles, but the state parameter
	// is not among the model's inputs and evaluation fails in the driver.
	m, err := c.Compile(ane.CompileOptions{
		ModelType: ane.ModelTypeMIL,
		MILText:   []byte(mil.GenUpdateState("kv", shape)),
	})
	if err != nil {
		t.Logf("GenUpdateState no longer compiles through the public API: %v", err)
		return
	}
	defer m.Close()

	if got := m.NumInputs(); got != 1 {
		t.Errorf("GenUpdateState model has %d inputs, want 1 (the state parameter is not exposed as one); if the state is now an input, state models may have become usable", got)
	}
	if err := m.Eval(); err == nil {
		t.Error("GenUpdateState now evaluates; state models may have become usable, so update the doc comments on GenReadState and GenUpdateState and consider writing the KV cache example")
	}
}
