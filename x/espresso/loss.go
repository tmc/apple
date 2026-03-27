//go:build darwin

package espresso

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pe "github.com/tmc/apple/private/espresso"
)

// LossType identifies the loss function.
type LossType int

const (
	LossCrossEntropy LossType = iota
	LossL2
)

// LossDef wraps an Espresso loss definition for training.
type LossDef struct {
	raw pe.ETLossDefinition
}

// LossConfig specifies loss function parameters.
type LossConfig struct {
	Type           LossType
	InputName      string // network output to compare
	TargetName     string // ground truth label name
	LossOutputName string // name for the loss output tensor
}

// NewLoss creates a loss definition from the given config.
func NewLoss(cfg LossConfig) *LossDef {
	cls := pe.GetETLossDefinitionClass()
	nsInput := objectivec.Object{ID: objc.String(cfg.InputName)}
	nsTarget := objectivec.Object{ID: objc.String(cfg.TargetName)}
	nsOutput := objectivec.Object{ID: objc.String(cfg.LossOutputName)}

	var raw objectivec.IObject
	switch cfg.Type {
	case LossCrossEntropy:
		raw = cls.CrossEntropyLossWithInputNameTargetInputNameLossOutputName(nsInput, nsTarget, nsOutput)
	case LossL2:
		raw = cls.L2LossWithInputNameTargetInputNameLossOutputName(nsInput, nsTarget, nsOutput)
	default:
		raw = cls.CrossEntropyLossWithInputNameTargetInputNameLossOutputName(nsInput, nsTarget, nsOutput)
	}
	return &LossDef{raw: pe.ETLossDefinitionFromID(raw.GetID())}
}

// CrossEntropyLoss creates a cross-entropy loss function.
func CrossEntropyLoss(inputName, targetName, lossOutputName string) *LossDef {
	return NewLoss(LossConfig{
		Type:           LossCrossEntropy,
		InputName:      inputName,
		TargetName:     targetName,
		LossOutputName: lossOutputName,
	})
}

// L2Loss creates an L2 (mean squared error) loss function.
func L2Loss(inputName, targetName, lossOutputName string) *LossDef {
	return NewLoss(LossConfig{
		Type:           LossL2,
		InputName:      inputName,
		TargetName:     targetName,
		LossOutputName: lossOutputName,
	})
}

// InputName returns the network output tensor name.
func (l *LossDef) InputName() string { return l.raw.InputName() }

// TargetName returns the ground truth tensor name.
func (l *LossDef) TargetName() string { return l.raw.TargetInputName() }

// LossOutputName returns the loss output tensor name.
func (l *LossDef) LossOutputName() string { return l.raw.LossOutputName() }
