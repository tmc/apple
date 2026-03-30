// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCustomLayerWrapper] class.
var (
	_MLCustomLayerWrapperClass     MLCustomLayerWrapperClass
	_MLCustomLayerWrapperClassOnce sync.Once
)

func getMLCustomLayerWrapperClass() MLCustomLayerWrapperClass {
	_MLCustomLayerWrapperClassOnce.Do(func() {
		_MLCustomLayerWrapperClass = MLCustomLayerWrapperClass{class: objc.GetClass("MLCustomLayerWrapper")}
	})
	return _MLCustomLayerWrapperClass
}

// GetMLCustomLayerWrapperClass returns the class object for MLCustomLayerWrapper.
func GetMLCustomLayerWrapperClass() MLCustomLayerWrapperClass {
	return getMLCustomLayerWrapperClass()
}

type MLCustomLayerWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCustomLayerWrapperClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCustomLayerWrapperClass) Alloc() MLCustomLayerWrapper {
	rv := objc.Send[MLCustomLayerWrapper](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCustomLayerWrapper.ClassName]
//   - [MLCustomLayerWrapper.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLCustomLayerWrapper.CustomImpl]
//   - [MLCustomLayerWrapper.SetCustomImpl]
//   - [MLCustomLayerWrapper.EncodeToMetalCommandBufferInputTensorsOutputTensors]
//   - [MLCustomLayerWrapper.HasGPUSupport]
//   - [MLCustomLayerWrapper.NdMode]
//   - [MLCustomLayerWrapper.SetMappedWeightsSizeInBytes]
//   - [MLCustomLayerWrapper.SetupForInputShapesWithParameters]
//   - [MLCustomLayerWrapper.InitWithParameters]
//   - [MLCustomLayerWrapper.DebugDescription]
//   - [MLCustomLayerWrapper.Description]
//   - [MLCustomLayerWrapper.Hash]
//   - [MLCustomLayerWrapper.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper
type MLCustomLayerWrapper struct {
	objectivec.Object
}

// MLCustomLayerWrapperFromID constructs a [MLCustomLayerWrapper] from an objc.ID.
func MLCustomLayerWrapperFromID(id objc.ID) MLCustomLayerWrapper {
	return MLCustomLayerWrapper{objectivec.Object{ID: id}}
}

// Ensure MLCustomLayerWrapper implements IMLCustomLayerWrapper.
var _ IMLCustomLayerWrapper = MLCustomLayerWrapper{}

// An interface definition for the [MLCustomLayerWrapper] class.
//
// # Methods
//
//   - [IMLCustomLayerWrapper.ClassName]
//   - [IMLCustomLayerWrapper.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLCustomLayerWrapper.CustomImpl]
//   - [IMLCustomLayerWrapper.SetCustomImpl]
//   - [IMLCustomLayerWrapper.EncodeToMetalCommandBufferInputTensorsOutputTensors]
//   - [IMLCustomLayerWrapper.HasGPUSupport]
//   - [IMLCustomLayerWrapper.NdMode]
//   - [IMLCustomLayerWrapper.SetMappedWeightsSizeInBytes]
//   - [IMLCustomLayerWrapper.SetupForInputShapesWithParameters]
//   - [IMLCustomLayerWrapper.InitWithParameters]
//   - [IMLCustomLayerWrapper.DebugDescription]
//   - [IMLCustomLayerWrapper.Description]
//   - [IMLCustomLayerWrapper.Hash]
//   - [IMLCustomLayerWrapper.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper
type IMLCustomLayerWrapper interface {
	objectivec.IObject

	// Topic: Methods

	ClassName() string
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	CustomImpl() objectivec.Object
	SetCustomImpl(value objectivec.Object)
	EncodeToMetalCommandBufferInputTensorsOutputTensors(buffer objectivec.IObject, tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	NdMode() bool
	SetMappedWeightsSizeInBytes(weights unsafe.Pointer, bytes uint64)
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLCustomLayerWrapper
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLCustomLayerWrapper) Init() MLCustomLayerWrapper {
	rv := objc.Send[MLCustomLayerWrapper](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCustomLayerWrapper) Autorelease() MLCustomLayerWrapper {
	rv := objc.Send[MLCustomLayerWrapper](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCustomLayerWrapper creates a new MLCustomLayerWrapper instance.
func NewMLCustomLayerWrapper() MLCustomLayerWrapper {
	class := getMLCustomLayerWrapperClass()
	rv := objc.Send[MLCustomLayerWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/initWithParameters:
func NewCustomLayerWrapperWithParameters(parameters objectivec.IObject) MLCustomLayerWrapper {
	instance := getMLCustomLayerWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLCustomLayerWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/computeOnCPUWithInputTensors:outputTensors:
func (c MLCustomLayerWrapper) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/encodeToMetalCommandBuffer:inputTensors:outputTensors:
func (c MLCustomLayerWrapper) EncodeToMetalCommandBufferInputTensorsOutputTensors(buffer objectivec.IObject, tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToMetalCommandBuffer:inputTensors:outputTensors:"), buffer, tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/hasGPUSupport
func (c MLCustomLayerWrapper) HasGPUSupport() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/setMappedWeights:sizeInBytes:
func (c MLCustomLayerWrapper) SetMappedWeightsSizeInBytes(weights unsafe.Pointer, bytes uint64) {
	objc.Send[objc.ID](c.ID, objc.Sel("setMappedWeights:sizeInBytes:"), weights, bytes)
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/setupForInputShapes:withParameters:
func (c MLCustomLayerWrapper) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/initWithParameters:
func (c MLCustomLayerWrapper) InitWithParameters(parameters objectivec.IObject) MLCustomLayerWrapper {
	rv := objc.Send[MLCustomLayerWrapper](c.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/coremlShapeToEspressoShape:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) CoremlShapeToEspressoShapeNdMode(shape objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("coremlShapeToEspressoShape:ndMode:"), shape, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/coremlShapesToEspressoShapes:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) CoremlShapesToEspressoShapesNdMode(shapes objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("coremlShapesToEspressoShapes:ndMode:"), shapes, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/espressoShapeToCoremlShape:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoShapeToCoremlShapeNdMode(shape objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoShapeToCoremlShape:ndMode:"), shape, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/espressoShapesToCoremlShapes:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoShapesToCoremlShapesNdMode(shapes objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoShapesToCoremlShapes:ndMode:"), shapes, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/espressoTensorToCoremlTensor:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorToCoremlTensorNdMode(tensor objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorToCoremlTensor:ndMode:"), tensor, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/espressoTensorsToCoremlTensors:ndMode:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorsToCoremlTensorsNdMode(tensors objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorsToCoremlTensors:ndMode:"), tensors, mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/espressoTensorsToCoremlTensorsGPU:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorsToCoremlTensorsGPU(gpu objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorsToCoremlTensorsGPU:"), gpu)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/factory
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) Factory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("factory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/getStrides:
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) GetStrides(strides objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("getStrides:"), strides)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/className
func (c MLCustomLayerWrapper) ClassName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("className"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/customImpl
func (c MLCustomLayerWrapper) CustomImpl() objectivec.Object {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("customImpl"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (c MLCustomLayerWrapper) SetCustomImpl(value objectivec.Object) {
	objc.Send[struct{}](c.ID, objc.Sel("setCustomImpl:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/debugDescription
func (c MLCustomLayerWrapper) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/description
func (c MLCustomLayerWrapper) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/hash
func (c MLCustomLayerWrapper) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/ndMode
func (c MLCustomLayerWrapper) NdMode() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("ndMode"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomLayerWrapper/superclass
func (c MLCustomLayerWrapper) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
