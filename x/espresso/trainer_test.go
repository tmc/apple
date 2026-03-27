//go:build darwin

package espresso

import (
	"testing"
)

func TestNewSGD(t *testing.T) {
	opt := NewSGD(SGDConfig{
		LearningRate: 0.001,
		Momentum:     0.9,
		WeightDecay:  0.0001,
	})
	if opt == nil {
		t.Fatal("NewSGD returned nil")
	}
	if got := opt.LearningRate(); got != 0.001 {
		t.Errorf("LearningRate() = %v, want 0.001", got)
	}
	if got := opt.Momentum(); got != 0.9 {
		t.Errorf("Momentum() = %v, want 0.9", got)
	}
	if got := opt.WeightDecay(); got != 0.0001 {
		t.Errorf("WeightDecay() = %v, want 0.0001", got)
	}
}

func TestNewSGDDefaults(t *testing.T) {
	opt := NewSGD(SGDConfig{})
	if opt == nil {
		t.Fatal("NewSGD returned nil")
	}
	if got := opt.LearningRate(); got != 0.01 {
		t.Errorf("default LearningRate() = %v, want 0.01", got)
	}
}

func TestCrossEntropyLoss(t *testing.T) {
	loss := CrossEntropyLoss("output", "label", "loss")
	if loss == nil {
		t.Fatal("CrossEntropyLoss returned nil")
	}
	if got := loss.InputName(); got != "output" {
		t.Errorf("InputName() = %q, want %q", got, "output")
	}
	if got := loss.TargetName(); got != "label" {
		t.Errorf("TargetName() = %q, want %q", got, "label")
	}
	if got := loss.LossOutputName(); got != "loss" {
		t.Errorf("LossOutputName() = %q, want %q", got, "loss")
	}
}

func TestL2Loss(t *testing.T) {
	loss := L2Loss("pred", "target", "mse")
	if loss == nil {
		t.Fatal("L2Loss returned nil")
	}
}

func TestNewMLP(t *testing.T) {
	model, err := NewMLP(MLPConfig{
		InputSize:  784,
		HiddenSize: 128,
		OutputSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}
	if model == nil {
		t.Fatal("NewMLP returned nil model")
	}
	// Verify dimension setters took effect.
	if got := model.mlp.Input_size(); got != 784 {
		t.Errorf("input_size = %d, want 784", got)
	}
	if got := model.mlp.Hidden_size(); got != 128 {
		t.Errorf("hidden_size = %d, want 128", got)
	}
	if got := model.mlp.Output_size(); got != 10 {
		t.Errorf("output_size = %d, want 10", got)
	}
}

func TestNewBufferDataSource(t *testing.T) {
	ds := NewBufferDataSource()
	if ds == nil {
		t.Fatal("NewBufferDataSource returned nil")
	}
	if got := ds.Count(); got != 0 {
		t.Errorf("Count() = %d, want 0", got)
	}
}

// TestTrainerInit requires a fully-formed Espresso IR network.
// MLP models created via NewMLP do not call BuildNetwork (it requires
// the IR platform), so trainer init with raw MLPs crashes. This test
// is skipped until a pre-built .espresso.net file is available.
func TestTrainerInit(t *testing.T) {
	t.Skip("requires pre-built espresso.net file (MLP init does not call BuildNetwork)")
}

func TestTrainerCloseIdempotent(t *testing.T) {
	t.Skip("requires pre-built espresso.net file (MLP init does not call BuildNetwork)")
}
