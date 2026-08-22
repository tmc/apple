//go:build darwin

package mil_test

import (
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// TestStateProgramsDoNotRun records that the two state generators produce
// programs nothing in this module can execute, and fails if that changes.
//
// It is a tripwire rather than a check on the generators. The text they emit is
// valid MIL and the ANE compiler accepts one of the two; what is missing is any
// way to bind the state buffer, on either route. Of the 292 e5rt symbols the
// framework exports, the only two mentioning state are a compiler option and
// its getter. So the day a macOS release supplies one, these expectations break
// and the KV cache that has been waiting on them becomes writable — which is
// worth being told about, and is why this asserts the failure rather than
// skipping.
//
// The e5rt half of the story is not exercised here: encoding a state operation
// on that route fails with status 2, and one nearby program shape faults inside
// the compiler outright, which would take the test binary down with it. See the
// doc comments on [mil.GenReadState] and [mil.GenUpdateState].
func TestStateProgramsDoNotRun(t *testing.T) {
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
