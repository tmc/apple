//go:build darwin

package espresso

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pe "github.com/tmc/apple/private/espresso"
)

// ModelDef describes a network topology for training or inference.
type ModelDef struct {
	raw pe.ETModelDefinition
	mlp pe.ETModelDefMLP // set for MLP models
}

// FromNetworkPath creates a ModelDef for inference from an Espresso
// network file path (typically a .espresso.net JSON file).
func FromNetworkPath(path string) (*ModelDef, error) {
	nsPath := objectivec.Object{ID: objc.String(path)}
	raw, err := pe.NewETModelDefinitionWithInferenceNetworkPathError(nsPath)
	if err != nil {
		return nil, fmt.Errorf("espresso modeldef: %w", err)
	}
	if raw.GetID() == 0 {
		return nil, fmt.Errorf("espresso modeldef: nil model definition for %s", path)
	}
	return &ModelDef{raw: raw}, nil
}

// FromNetworkPathWithIO creates a ModelDef with explicit input/output names.
func FromNetworkPathWithIO(path string, inputs, outputs []string) (*ModelDef, error) {
	nsPath := objectivec.Object{ID: objc.String(path)}
	nsInputs := stringsToNSArray(inputs)
	nsOutputs := stringsToNSArray(outputs)
	raw, err := pe.NewETModelDefinitionWithInferenceNetworkPathInferenceInputsInferenceOutputsError(
		nsPath, nsInputs, nsOutputs,
	)
	if err != nil {
		return nil, fmt.Errorf("espresso modeldef: %w", err)
	}
	if raw.GetID() == 0 {
		return nil, fmt.Errorf("espresso modeldef: nil model definition for %s", path)
	}
	return &ModelDef{raw: raw}, nil
}

// MLPConfig describes a multi-layer perceptron architecture.
type MLPConfig struct {
	InputSize  int
	HiddenSize int
	OutputSize int
}

// NewMLP creates a predefined MLP (multi-layer perceptron) model.
//
// Note: BuildNetwork requires the Espresso IR platform and may fail
// at runtime if the framework does not support gradient builder creation.
// Use FromNetworkPath with a pre-built espresso.net file as a more
// reliable alternative.
func NewMLP(cfg MLPConfig) (*ModelDef, error) {
	mlp := pe.NewETModelDefMLP()
	if mlp.GetID() == 0 {
		return nil, fmt.Errorf("espresso modeldef: failed to allocate ETModelDefMLP")
	}
	mlp.SetInput_size(cfg.InputSize)
	mlp.SetHidden_size(cfg.HiddenSize)
	mlp.SetOutput_size(cfg.OutputSize)
	// BuildNetwork is deferred — it requires the IR platform and is called
	// internally when the model is used in a training task. Calling it
	// eagerly causes an Espresso::invalid_state_error.
	return &ModelDef{mlp: mlp}, nil
}

// Inputs returns the inference input tensor names.
func (m *ModelDef) Inputs() []string {
	return nsArrayToStrings(m.raw.Inputs())
}

// Outputs returns the inference output tensor names.
func (m *ModelDef) Outputs() []string {
	return nsArrayToStrings(m.raw.Outputs())
}

// LayerNames returns the layer names in the network.
func (m *ModelDef) LayerNames() []string {
	return nsArrayToStrings(m.raw.LayerNames())
}

func stringsToNSArray(ss []string) objectivec.Object {
	arr := foundation.NewNSMutableArray()
	for _, s := range ss {
		arr.AddObject(objectivec.Object{ID: objc.String(s)})
	}
	return objectivec.Object{ID: arr.GetID()}
}
