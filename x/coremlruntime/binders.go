//go:build darwin

package coremlruntime

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	privatecoreml "github.com/tmc/apple/private/coreml"
)

// PixelBufferPool wraps CoreML's private pixel-buffer pool helper.
type PixelBufferPool struct {
	raw privatecoreml.MLPixelBufferPool
}

// InputPortBinder wraps CoreML's private input-port binder.
type InputPortBinder struct {
	raw privatecoreml.MLE5InputPortBinder
}

// OutputPortBinder wraps CoreML's private output-port binder.
type OutputPortBinder struct {
	raw privatecoreml.MLE5OutputPortBinder
}

// OutputBackingsVerifier wraps CoreML's private output-backing verifier.
type OutputBackingsVerifier struct {
	raw privatecoreml.MLOutputBackingsVerifier
}

// NewPixelBufferPool creates a private CoreML pixel-buffer pool wrapper.
func NewPixelBufferPool() PixelBufferPool {
	return PixelBufferPool{raw: privatecoreml.NewMLPixelBufferPool()}
}

// Raw returns the underlying private pixel-buffer pool object.
func (p PixelBufferPool) Raw() privatecoreml.MLPixelBufferPool {
	return p.raw
}

// CreatePixelBuffer allocates a pixel buffer from the pool.
func (p PixelBufferPool) CreatePixelBuffer(size corefoundation.CGSize, pixelFormat uint32) (corevideo.CVImageBufferRef, error) {
	if p.raw.ID == 0 {
		return 0, fmt.Errorf("coremlruntime: pixel buffer pool is nil")
	}
	buf, err := p.raw.CreatePixelBufferWithSizePixelFormatTypeError(size, pixelFormat)
	if err != nil {
		return 0, fmt.Errorf("coremlruntime: create pixel buffer: %w", err)
	}
	return buf, nil
}

// Cache returns the pool's internal cache dictionary.
func (p PixelBufferPool) Cache() foundation.INSDictionary {
	return p.raw.PixelBufferPoolCache()
}

// NewInputPortBinder constructs an input-port binder for a port/feature pair.
func NewInputPortBinder(port, featureDescription objectivec.IObject) InputPortBinder {
	return InputPortBinder{raw: privatecoreml.NewE5InputPortBinderWithPortFeatureDescription(port, featureDescription)}
}

// Raw returns the underlying private input-port binder object.
func (b InputPortBinder) Raw() privatecoreml.MLE5InputPortBinder {
	return b.raw
}

// SetPixelBufferPool routes image binding through the given pixel-buffer pool.
func (b InputPortBinder) SetPixelBufferPool(pool PixelBufferPool) {
	b.raw.SetPixelBufferPool(pool.raw)
}

// Reset resets binder state between submissions.
func (b InputPortBinder) Reset() {
	b.raw.Reset()
}

// CopyFeatureValue copies a feature value into the bound port.
func (b InputPortBinder) CopyFeatureValue(value coreml.MLFeatureValue) error {
	if _, err := b.raw.CopyFeatureValueError(value); err != nil {
		return fmt.Errorf("coremlruntime: copy feature value: %w", err)
	}
	return nil
}

// BindMemoryObject binds a feature value directly when the runtime supports it.
func (b InputPortBinder) BindMemoryObject(value coreml.MLFeatureValue) error {
	if _, err := b.raw.BindMemoryObjectForFeatureValueError(value); err != nil {
		return fmt.Errorf("coremlruntime: bind memory object: %w", err)
	}
	return nil
}

// ReusableForFeatureValue reports whether the binder can be reused for the value
// and whether the runtime intends to bind it directly.
func (b InputPortBinder) ReusableForFeatureValue(value coreml.MLFeatureValue) (reusable bool, direct bool) {
	direct, reusable = b.raw.ReusableForFeatureValueWillBindDirectly(value)
	return reusable, direct
}

// NewOutputPortBinder constructs an output-port binder for a port/feature pair.
func NewOutputPortBinder(port, featureDescription objectivec.IObject) OutputPortBinder {
	return OutputPortBinder{raw: privatecoreml.NewE5OutputPortBinderWithPortFeatureDescription(port, featureDescription)}
}

// Raw returns the underlying private output-port binder object.
func (b OutputPortBinder) Raw() privatecoreml.MLE5OutputPortBinder {
	return b.raw
}

// SetPixelBufferPool routes image output binding through the given pool.
func (b OutputPortBinder) SetPixelBufferPool(pool PixelBufferPool) {
	b.raw.SetPixelBufferPool(pool.raw)
}

// SetOutputBacking sets the output backing object on the binder.
func (b OutputPortBinder) SetOutputBacking(backing objectivec.IObject) {
	b.raw.SetOutputBacking(backing)
}

// Bind binds the currently configured output backing.
func (b OutputPortBinder) Bind() error {
	if _, err := b.raw.BindAndReturnError(); err != nil {
		return fmt.Errorf("coremlruntime: bind output backing: %w", err)
	}
	return nil
}

// DirectlyBindOutputBacking asks the runtime to directly bind the output backing.
func (b OutputPortBinder) DirectlyBindOutputBacking(backing objectivec.IObject) error {
	if _, err := b.raw.DirectlyBindOutputBackingError(backing); err != nil {
		return fmt.Errorf("coremlruntime: directly bind output backing: %w", err)
	}
	return nil
}

// MakeFeatureValue materializes a feature value from the current output binding.
func (b OutputPortBinder) MakeFeatureValue() (coreml.MLFeatureValue, error) {
	value, err := b.raw.MakeFeatureValueAndReturnError()
	if err != nil {
		return coreml.MLFeatureValue{}, fmt.Errorf("coremlruntime: make feature value: %w", err)
	}
	return coreml.MLFeatureValueFromID(value.GetID()), nil
}

// MakeFeatureValueFromOutputBacking materializes a feature value from a backing object.
func (b OutputPortBinder) MakeFeatureValueFromOutputBacking(backing objectivec.IObject) (coreml.MLFeatureValue, error) {
	value, err := b.raw.MakeFeatureValueFromOutputBackingError(backing)
	if err != nil {
		return coreml.MLFeatureValue{}, fmt.Errorf("coremlruntime: make feature value from output backing: %w", err)
	}
	return coreml.MLFeatureValueFromID(value.GetID()), nil
}

// Reset resets binder state between submissions.
func (b OutputPortBinder) Reset() {
	b.raw.Reset()
}

// ReusableForOutputBacking reports whether the binder can be reused for the backing
// and whether the runtime intends to bind it directly.
func (b OutputPortBinder) ReusableForOutputBacking(backing objectivec.IObject) (reusable bool, direct bool) {
	direct, reusable = b.raw.ReusableForOutputBackingWillBindDirectly(backing)
	return reusable, direct
}

// NewOutputBackingsVerifier constructs a verifier for a model's output descriptions.
func NewOutputBackingsVerifier(descriptions objectivec.IObject) OutputBackingsVerifier {
	return OutputBackingsVerifier{raw: privatecoreml.NewOutputBackingsVerifierWithOutputDescriptions(descriptions)}
}

// VerifyOutputBackings verifies a dictionary of output backings against the descriptions.
func (v OutputBackingsVerifier) VerifyOutputBackings(backings foundation.INSDictionary, predictionUsesBatch bool) error {
	if _, err := v.raw.VerifyOutputBackingsPredictionUsesBatchError(backings, predictionUsesBatch); err != nil {
		return fmt.Errorf("coremlruntime: verify output backings: %w", err)
	}
	return nil
}

// VerifyOutputBacking verifies a single output backing object.
func (v OutputBackingsVerifier) VerifyOutputBacking(backing, feature objectivec.IObject) error {
	if _, err := v.raw.VerifyOutputBackingForFeatureError(backing, feature); err != nil {
		return fmt.Errorf("coremlruntime: verify output backing: %w", err)
	}
	return nil
}

// VerifyPixelBufferOutputBacking verifies a pixel-buffer output backing.
func (v OutputBackingsVerifier) VerifyPixelBufferOutputBacking(backing corevideo.CVImageBufferRef, feature objectivec.IObject) error {
	if _, err := v.raw.VerifyPixelBufferOutputBackingForFeatureError(backing, feature); err != nil {
		return fmt.Errorf("coremlruntime: verify pixel-buffer output backing: %w", err)
	}
	return nil
}

// OutputBackingsDictionary builds an NSDictionary suitable for MLPredictionOptions output backings.
func OutputBackingsDictionary(backings map[string]objectivec.IObject) foundation.NSDictionary {
	keys := make([]objectivec.IObject, 0, len(backings))
	values := make([]objectivec.IObject, 0, len(backings))
	for name, backing := range backings {
		keys = append(keys, foundation.NewStringWithString(name))
		values = append(values, backing)
	}
	return foundation.GetNSDictionaryClass().DictionaryWithObjectsForKeys(values, keys)
}
