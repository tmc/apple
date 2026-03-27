//go:build darwin

package espresso

import (
	pe "github.com/tmc/apple/private/espresso"
)

// OptimizerDef wraps an Espresso optimizer definition for training.
type OptimizerDef struct {
	raw pe.ETOptimizerDefSGD // SGD is the concrete type; all optimizers embed ETOptimizerDef
}

// SGDConfig configures a stochastic gradient descent optimizer.
type SGDConfig struct {
	LearningRate float32 // learning rate (default 0.01)
	Momentum     float32 // momentum factor (default 0)
	WeightDecay  float32 // L2 regularization (default 0)
	LRDecayEpoch float32 // epoch at which to decay LR (default 0, no decay)
}

// NewSGD creates an SGD optimizer with the given configuration.
func NewSGD(cfg SGDConfig) *OptimizerDef {
	raw := pe.NewETOptimizerDefSGD()
	if raw.GetID() == 0 {
		return nil
	}
	if cfg.LearningRate > 0 {
		raw.SetLr(cfg.LearningRate)
	} else {
		raw.SetLr(0.01)
	}
	if cfg.Momentum > 0 {
		raw.SetMomentum(cfg.Momentum)
	}
	if cfg.WeightDecay > 0 {
		raw.SetWeight_decay(cfg.WeightDecay)
	}
	if cfg.LRDecayEpoch > 0 {
		raw.SetLr_decay_epoch(cfg.LRDecayEpoch)
	}
	return &OptimizerDef{raw: raw}
}

// LearningRate returns the current learning rate.
func (o *OptimizerDef) LearningRate() float32 { return o.raw.Lr() }

// Momentum returns the momentum factor.
func (o *OptimizerDef) Momentum() float32 { return o.raw.Momentum() }

// WeightDecay returns the L2 regularization strength.
func (o *OptimizerDef) WeightDecay() float32 { return o.raw.Weight_decay() }

// SetLearningRate updates the learning rate.
func (o *OptimizerDef) SetLearningRate(lr float32) { o.raw.SetLr(lr) }
