// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphSDPADescriptor] class.
var (
	_MPSGraphSDPADescriptorClass     MPSGraphSDPADescriptorClass
	_MPSGraphSDPADescriptorClassOnce sync.Once
)

func getMPSGraphSDPADescriptorClass() MPSGraphSDPADescriptorClass {
	_MPSGraphSDPADescriptorClassOnce.Do(func() {
		_MPSGraphSDPADescriptorClass = MPSGraphSDPADescriptorClass{class: objc.GetClass("MPSGraphSDPADescriptor")}
	})
	return _MPSGraphSDPADescriptorClass
}

// GetMPSGraphSDPADescriptorClass returns the class object for MPSGraphSDPADescriptor.
func GetMPSGraphSDPADescriptorClass() MPSGraphSDPADescriptorClass {
	return getMPSGraphSDPADescriptorClass()
}

type MPSGraphSDPADescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphSDPADescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphSDPADescriptorClass) Alloc() MPSGraphSDPADescriptor {
	rv := objc.Send[MPSGraphSDPADescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A descriptor that configures a scaled dot product attention (SDPA)
// operation.
//
// # Overview
//
// Use this descriptor with
// [MPSGraph.ScaledDotProductAttentionWithQueryTensorKeyTensorValueTensorDescriptorName]
// to specify optional features such as an attention mask, causal masking, and
// attention sinks.
//
// # Instance Properties
//
//   - [MPSGraphSDPADescriptor.IsCausal]: When YES, a causal (lower-triangular) mask is applied so that each query position attends only to key positions at or before it. Mutually exclusive with [maskTensor](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/maskTensor>).
//   - [MPSGraphSDPADescriptor.SetIsCausal]
//   - [MPSGraphSDPADescriptor.MaskTensor]: An optional additive mask tensor applied to the scaled QK^T scores before softmax. Must be broadcast-compatible with shape `[batch, heads, T_q, T_kv]`. Mutually exclusive with [isCausal](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/isCausal>).
//   - [MPSGraphSDPADescriptor.SetMaskTensor]
//   - [MPSGraphSDPADescriptor.Scale]: The scale applied to the result of the query–key matrix multiply before softmax. Typically set to `1/sqrt(headDimension)`.
//   - [MPSGraphSDPADescriptor.SetScale]
//   - [MPSGraphSDPADescriptor.SinksTensor]: An optional attention-sinks tensor of shape `[nHeads]`. Each element seeds the online-softmax accumulator for the corresponding query head with a virtual token logit, causing real-token attention weights to sum to less than one.
//   - [MPSGraphSDPADescriptor.SetSinksTensor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor
type MPSGraphSDPADescriptor struct {
	MPSGraphObject
}

// MPSGraphSDPADescriptorFromID constructs a [MPSGraphSDPADescriptor] from an objc.ID.
//
// A descriptor that configures a scaled dot product attention (SDPA)
// operation.
func MPSGraphSDPADescriptorFromID(id objc.ID) MPSGraphSDPADescriptor {
	return MPSGraphSDPADescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphSDPADescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphSDPADescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphSDPADescriptor.IsCausal]: When YES, a causal (lower-triangular) mask is applied so that each query position attends only to key positions at or before it. Mutually exclusive with [maskTensor](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/maskTensor>).
//   - [IMPSGraphSDPADescriptor.SetIsCausal]
//   - [IMPSGraphSDPADescriptor.MaskTensor]: An optional additive mask tensor applied to the scaled QK^T scores before softmax. Must be broadcast-compatible with shape `[batch, heads, T_q, T_kv]`. Mutually exclusive with [isCausal](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/isCausal>).
//   - [IMPSGraphSDPADescriptor.SetMaskTensor]
//   - [IMPSGraphSDPADescriptor.Scale]: The scale applied to the result of the query–key matrix multiply before softmax. Typically set to `1/sqrt(headDimension)`.
//   - [IMPSGraphSDPADescriptor.SetScale]
//   - [IMPSGraphSDPADescriptor.SinksTensor]: An optional attention-sinks tensor of shape `[nHeads]`. Each element seeds the online-softmax accumulator for the corresponding query head with a virtual token logit, causing real-token attention weights to sum to less than one.
//   - [IMPSGraphSDPADescriptor.SetSinksTensor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor
type IMPSGraphSDPADescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// When YES, a causal (lower-triangular) mask is applied so that each query position attends only to key positions at or before it. Mutually exclusive with [maskTensor](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/maskTensor>).
	IsCausal() bool
	SetIsCausal(value bool)
	// An optional additive mask tensor applied to the scaled QK^T scores before softmax. Must be broadcast-compatible with shape `[batch, heads, T_q, T_kv]`. Mutually exclusive with [isCausal](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/isCausal>).
	MaskTensor() IMPSGraphTensor
	SetMaskTensor(value IMPSGraphTensor)
	// The scale applied to the result of the query–key matrix multiply before softmax. Typically set to `1/sqrt(headDimension)`.
	Scale() float32
	SetScale(value float32)
	// An optional attention-sinks tensor of shape `[nHeads]`. Each element seeds the online-softmax accumulator for the corresponding query head with a virtual token logit, causing real-token attention weights to sum to less than one.
	SinksTensor() IMPSGraphTensor
	SetSinksTensor(value IMPSGraphTensor)
}

// Init initializes the instance.
func (g MPSGraphSDPADescriptor) Init() MPSGraphSDPADescriptor {
	rv := objc.Send[MPSGraphSDPADescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphSDPADescriptor) Autorelease() MPSGraphSDPADescriptor {
	rv := objc.Send[MPSGraphSDPADescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphSDPADescriptor creates a new MPSGraphSDPADescriptor instance.
func NewMPSGraphSDPADescriptor() MPSGraphSDPADescriptor {
	class := getMPSGraphSDPADescriptorClass()
	rv := objc.Send[MPSGraphSDPADescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a descriptor with the given scale and all other properties set to
// their defaults (no mask, isCausal = NO, no sinks).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/init(scale:)
func NewGraphSDPADescriptorWithScale(scale float32) MPSGraphSDPADescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphSDPADescriptorClass().class), objc.Sel("descriptorWithScale:"), scale)
	return MPSGraphSDPADescriptorFromID(rv)
}

// When YES, a causal (lower-triangular) mask is applied so that each query
// position attends only to key positions at or before it. Mutually exclusive
// with [MPSGraphSDPADescriptor.MaskTensor].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/isCausal
func (g MPSGraphSDPADescriptor) IsCausal() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isCausal"))
	return rv
}
func (g MPSGraphSDPADescriptor) SetIsCausal(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setIsCausal:"), value)
}

// An optional additive mask tensor applied to the scaled QK^T scores before
// softmax. Must be broadcast-compatible with shape `[batch, heads, T_q,
// T_kv]`. Mutually exclusive with [MPSGraphSDPADescriptor.IsCausal].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/maskTensor
func (g MPSGraphSDPADescriptor) MaskTensor() IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("maskTensor"))
	return MPSGraphTensorFromID(objc.ID(rv))
}
func (g MPSGraphSDPADescriptor) SetMaskTensor(value IMPSGraphTensor) {
	objc.Send[struct{}](g.ID, objc.Sel("setMaskTensor:"), value)
}

// The scale applied to the result of the query–key matrix multiply before
// softmax. Typically set to `1/sqrt(headDimension)`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/scale
func (g MPSGraphSDPADescriptor) Scale() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("scale"))
	return rv
}
func (g MPSGraphSDPADescriptor) SetScale(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setScale:"), value)
}

// An optional attention-sinks tensor of shape `[nHeads]`. Each element seeds
// the online-softmax accumulator for the corresponding query head with a
// virtual token logit, causing real-token attention weights to sum to less
// than one.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSDPADescriptor/sinksTensor
func (g MPSGraphSDPADescriptor) SinksTensor() IMPSGraphTensor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sinksTensor"))
	return MPSGraphTensorFromID(objc.ID(rv))
}
func (g MPSGraphSDPADescriptor) SetSinksTensor(value IMPSGraphTensor) {
	objc.Send[struct{}](g.ID, objc.Sel("setSinksTensor:"), value)
}
