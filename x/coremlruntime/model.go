//go:build darwin

package coremlruntime

import (
	"fmt"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	privatecoreml "github.com/tmc/apple/private/coreml"
	"github.com/tmc/apple/x/coremlconfig"
)

// Model wraps a loaded CoreML model with private runtime helpers.
type Model struct {
	raw coreml.MLModel
}

// Prediction retains a CoreML prediction result so callers can inspect outputs
// without immediately materializing them.
type Prediction struct {
	provider coreml.MLFeatureProviderObject
	owner    objectivec.Object
}

// Load opens a compiled .mlmodelc bundle.
func Load(path string, cfg *coremlconfig.Config) (*Model, error) {
	url := foundation.NewURLFileURLWithPathIsDirectory(path, true)
	if url.GetID() == 0 {
		return nil, fmt.Errorf("coremlruntime: invalid path %q", path)
	}
	var (
		raw coreml.MLModel
		err error
	)
	if cfg != nil {
		raw, err = coreml.NewModelWithContentsOfURLConfigurationError(url, cfg.Raw())
	} else {
		raw, err = coreml.NewModelWithContentsOfURLError(url)
	}
	if err != nil {
		return nil, fmt.Errorf("coremlruntime: load model: %w", err)
	}
	if raw.ID == 0 {
		return nil, fmt.Errorf("coremlruntime: load model returned nil for %q", path)
	}
	return &Model{raw: raw}, nil
}

// Wrap exposes an existing CoreML model through the runtime helper surface.
func Wrap(raw coreml.MLModel) *Model {
	if raw.ID == 0 {
		return nil
	}
	return &Model{raw: raw}
}

// Raw returns the underlying MLModel.
func (m *Model) Raw() coreml.MLModel {
	if m == nil {
		return coreml.MLModel{}
	}
	return m.raw
}

func (m *Model) priv() privatecoreml.MLModel {
	if m == nil {
		return privatecoreml.MLModel{}
	}
	return privatecoreml.MLModelFromID(m.raw.ID)
}

// NewState allocates a new CoreML state handle for a stateful model.
func (m *Model) NewState() coreml.MLState {
	if m == nil {
		return coreml.MLState{}
	}
	return coreml.MLStateFromID(m.priv().NewState().GetID())
}

// NewStateWithClientBuffers allocates a new CoreML state handle backed by
// client-supplied buffers.
func (m *Model) NewStateWithClientBuffers(buffers objectivec.IObject) coreml.MLState {
	if m == nil || buffers == nil || buffers.GetID() == 0 {
		return coreml.MLState{}
	}
	return coreml.MLStateFromID(m.priv().NewStateWithClientBuffers(buffers).GetID())
}

// PrepareWithConcurrencyHint forwards the private preparation hint to the model.
func (m *Model) PrepareWithConcurrencyHint(hint int64) {
	if m == nil {
		return
	}
	m.priv().PrepareWithConcurrencyHint(hint)
}

// SupportsConcurrentSubmissions reports whether the model runtime advertises
// concurrent submission support.
func (m *Model) SupportsConcurrentSubmissions() bool {
	if m == nil {
		return false
	}
	return m.priv().SupportsConcurrentSubmissions()
}

// PredictProvider runs inference using a caller-supplied feature provider.
func (m *Model) PredictProvider(provider coreml.MLFeatureProvider) (*Prediction, error) {
	return m.PredictProviderWithStateOptions(provider, nil, nil)
}

// PredictFeatureValues runs inference using a map of prebuilt feature values.
func (m *Model) PredictFeatureValues(values map[string]coreml.MLFeatureValue) (*Prediction, error) {
	provider, err := NewFeatureProvider(values)
	if err != nil {
		return nil, err
	}
	return m.PredictProviderWithStateOptions(provider, nil, nil)
}

// PredictProviderWithState runs inference using a caller-supplied state handle.
func (m *Model) PredictProviderWithState(provider coreml.MLFeatureProvider, state coreml.IMLState) (*Prediction, error) {
	return m.PredictProviderWithStateOptions(provider, state, nil)
}

// PredictProviderWithStateOptions runs inference with optional state and prediction options.
func (m *Model) PredictProviderWithStateOptions(provider coreml.MLFeatureProvider, state coreml.IMLState, options coreml.IMLPredictionOptions) (*Prediction, error) {
	if m == nil || m.raw.ID == 0 {
		return nil, fmt.Errorf("coremlruntime: model is nil")
	}
	if provider == nil || provider.GetID() == 0 {
		return nil, fmt.Errorf("coremlruntime: feature provider is nil")
	}

	var (
		result objectivec.IObject
		err    error
	)
	switch {
	case state != nil && state.GetID() != 0 && options != nil && options.GetID() != 0:
		result, err = m.raw.PredictionFromFeaturesUsingStateOptionsError(provider, state, options)
	case state != nil && state.GetID() != 0:
		result, err = m.raw.PredictionFromFeaturesUsingStateError(provider, state)
	case options != nil && options.GetID() != 0:
		result, err = m.raw.PredictionFromFeaturesOptionsError(provider, options)
	default:
		result, err = m.raw.PredictionFromFeaturesError(provider)
	}
	if err != nil {
		return nil, fmt.Errorf("coremlruntime: prediction: %w", err)
	}
	if result == nil || result.GetID() == 0 {
		return nil, fmt.Errorf("coremlruntime: prediction returned nil")
	}
	owner := objectivec.ObjectFromID(result.GetID())
	owner.Retain()
	return &Prediction{
		provider: coreml.MLFeatureProviderObjectFromID(owner.ID),
		owner:    owner,
	}, nil
}

// NewFeatureProvider builds an MLDictionaryFeatureProvider from prebuilt feature values.
func NewFeatureProvider(values map[string]coreml.MLFeatureValue) (coreml.MLDictionaryFeatureProvider, error) {
	keys := make([]objectivec.IObject, 0, len(values))
	vals := make([]objectivec.IObject, 0, len(values))
	for name, value := range values {
		if value.ID == 0 {
			return coreml.MLDictionaryFeatureProvider{}, fmt.Errorf("coremlruntime: nil feature value for %q", name)
		}
		keys = append(keys, foundation.NewStringWithString(name))
		vals = append(vals, value)
	}
	dict := foundation.GetNSDictionaryClass().DictionaryWithObjectsForKeys(vals, keys)
	provider, err := coreml.NewDictionaryFeatureProviderWithDictionaryError(dict)
	if err != nil {
		return coreml.MLDictionaryFeatureProvider{}, fmt.Errorf("coremlruntime: feature provider: %w", err)
	}
	return provider, nil
}

// FeatureNames returns the prediction's feature names.
func (p *Prediction) FeatureNames() []string {
	if p == nil || p.owner.ID == 0 {
		return nil
	}
	names := p.provider.FeatureNames().AllObjects()
	out := make([]string, 0, len(names))
	for _, name := range names {
		out = append(out, foundation.NSStringFromID(name.GetID()).String())
	}
	return out
}

// FeatureValue returns a named output feature value.
func (p *Prediction) FeatureValue(name string) coreml.MLFeatureValue {
	if p == nil || p.owner.ID == 0 {
		return coreml.MLFeatureValue{}
	}
	value := p.provider.FeatureValueForName(name)
	if value == nil {
		return coreml.MLFeatureValue{}
	}
	return coreml.MLFeatureValueFromID(value.GetID())
}

// MultiArrays returns the prediction's multi-array outputs.
func (p *Prediction) MultiArrays() (map[string]coreml.MLMultiArray, error) {
	if p == nil || p.owner.ID == 0 {
		return nil, fmt.Errorf("coremlruntime: prediction is closed")
	}
	names := p.provider.FeatureNames().AllObjects()
	out := make(map[string]coreml.MLMultiArray, len(names))
	for _, nameObj := range names {
		name := foundation.NSStringFromID(nameObj.GetID()).String()
		value := p.provider.FeatureValueForName(name)
		if value == nil || value.GetID() == 0 {
			continue
		}
		arr := coreml.MLMultiArrayFromID(value.MultiArrayValue().GetID())
		if arr.ID == 0 {
			continue
		}
		out[name] = arr
	}
	return out, nil
}

// Provider returns the retained prediction provider.
func (p *Prediction) Provider() coreml.MLFeatureProviderObject {
	if p == nil {
		return coreml.MLFeatureProviderObject{}
	}
	return p.provider
}

// Close releases the retained prediction result.
func (p *Prediction) Close() error {
	if p == nil || p.owner.ID == 0 {
		return nil
	}
	p.owner.Release()
	p.owner.ID = 0
	p.provider = coreml.MLFeatureProviderObject{}
	return nil
}
