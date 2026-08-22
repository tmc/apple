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
	rv := objc.SendIfResponds[MLCustomLayerWrapper](objc.ID(mc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLCustomLayerWrapper) Init() MLCustomLayerWrapper {
	rv := objc.SendIfResponds[MLCustomLayerWrapper](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCustomLayerWrapper) Autorelease() MLCustomLayerWrapper {
	rv := objc.SendIfResponds[MLCustomLayerWrapper](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCustomLayerWrapper creates a new MLCustomLayerWrapper instance.
func NewMLCustomLayerWrapper() MLCustomLayerWrapper {
	class := getMLCustomLayerWrapperClass()
	rv := objc.SendIfResponds[MLCustomLayerWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCustomLayerWrapperWithParameters(parameters objectivec.IObject) MLCustomLayerWrapper {
	instance := getMLCustomLayerWrapperClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLCustomLayerWrapperFromID(rv)
}

func (m MLCustomLayerWrapper) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLCustomLayerWrapper) EncodeToMetalCommandBufferInputTensorsOutputTensors(buffer objectivec.IObject, tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encodeToMetalCommandBuffer:inputTensors:outputTensors:"), buffer, tensors, tensors2)
}
func (m MLCustomLayerWrapper) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLCustomLayerWrapper) SetMappedWeightsSizeInBytes(weights unsafe.Pointer, bytes uint64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setMappedWeights:sizeInBytes:"), weights, bytes)
}
func (m MLCustomLayerWrapper) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLCustomLayerWrapper) InitWithParameters(parameters objectivec.IObject) MLCustomLayerWrapper {
	rv := objc.SendIfResponds[MLCustomLayerWrapper](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) CoremlShapeToEspressoShapeNdMode(shape objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("coremlShapeToEspressoShape:ndMode:"), shape, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) CoremlShapesToEspressoShapesNdMode(shapes objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("coremlShapesToEspressoShapes:ndMode:"), shapes, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoShapeToCoremlShapeNdMode(shape objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoShapeToCoremlShape:ndMode:"), shape, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoShapesToCoremlShapesNdMode(shapes objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoShapesToCoremlShapes:ndMode:"), shapes, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorToCoremlTensorNdMode(tensor objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorToCoremlTensor:ndMode:"), tensor, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorsToCoremlTensorsNdMode(tensors objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorsToCoremlTensors:ndMode:"), tensors, mode)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) EspressoTensorsToCoremlTensorsGPU(gpu objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("espressoTensorsToCoremlTensorsGPU:"), gpu)
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) Factory() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("factory"))
	return objectivec.Object{ID: rv}
}
func (_MLCustomLayerWrapperClass MLCustomLayerWrapperClass) GetStrides(strides objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCustomLayerWrapperClass.class), objc.Sel("getStrides:"), strides)
	return objectivec.Object{ID: rv}
}

func (m MLCustomLayerWrapper) ClassName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("className"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCustomLayerWrapper) CustomImpl() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("customImpl"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLCustomLayerWrapper) SetCustomImpl(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setCustomImpl:"), value)
}
func (m MLCustomLayerWrapper) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCustomLayerWrapper) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCustomLayerWrapper) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLCustomLayerWrapper) NdMode() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("ndMode"))
	return rv
}
func (m MLCustomLayerWrapper) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
