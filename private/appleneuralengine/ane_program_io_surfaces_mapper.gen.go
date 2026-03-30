// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEProgramIOSurfacesMapper] class.
var (
	_ANEProgramIOSurfacesMapperClass     ANEProgramIOSurfacesMapperClass
	_ANEProgramIOSurfacesMapperClassOnce sync.Once
)

func getANEProgramIOSurfacesMapperClass() ANEProgramIOSurfacesMapperClass {
	_ANEProgramIOSurfacesMapperClassOnce.Do(func() {
		_ANEProgramIOSurfacesMapperClass = ANEProgramIOSurfacesMapperClass{class: objc.GetClass("_ANEProgramIOSurfacesMapper")}
	})
	return _ANEProgramIOSurfacesMapperClass
}

// GetANEProgramIOSurfacesMapperClass returns the class object for _ANEProgramIOSurfacesMapper.
func GetANEProgramIOSurfacesMapperClass() ANEProgramIOSurfacesMapperClass {
	return getANEProgramIOSurfacesMapperClass()
}

type ANEProgramIOSurfacesMapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEProgramIOSurfacesMapperClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEProgramIOSurfacesMapperClass) Alloc() ANEProgramIOSurfacesMapper {
	rv := objc.Send[ANEProgramIOSurfacesMapper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEProgramIOSurfacesMapper.Controller]
//   - [ANEProgramIOSurfacesMapper.DeviceController]
//   - [ANEProgramIOSurfacesMapper.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [ANEProgramIOSurfacesMapper.PrepareANEMemoryMappingParamsRequest]
//   - [ANEProgramIOSurfacesMapper.ProgramHandle]
//   - [ANEProgramIOSurfacesMapper.UnmapIOSurfacesWithModelRequestError]
//   - [ANEProgramIOSurfacesMapper.ValidateRequestModel]
//   - [ANEProgramIOSurfacesMapper.InitWithController]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper
type ANEProgramIOSurfacesMapper struct {
	objectivec.Object
}

// ANEProgramIOSurfacesMapperFromID constructs a [ANEProgramIOSurfacesMapper] from an objc.ID.
func ANEProgramIOSurfacesMapperFromID(id objc.ID) ANEProgramIOSurfacesMapper {
	return ANEProgramIOSurfacesMapper{objectivec.Object{ID: id}}
}

// Ensure ANEProgramIOSurfacesMapper implements IANEProgramIOSurfacesMapper.
var _ IANEProgramIOSurfacesMapper = ANEProgramIOSurfacesMapper{}

// An interface definition for the [ANEProgramIOSurfacesMapper] class.
//
// # Methods
//
//   - [IANEProgramIOSurfacesMapper.Controller]
//   - [IANEProgramIOSurfacesMapper.DeviceController]
//   - [IANEProgramIOSurfacesMapper.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [IANEProgramIOSurfacesMapper.PrepareANEMemoryMappingParamsRequest]
//   - [IANEProgramIOSurfacesMapper.ProgramHandle]
//   - [IANEProgramIOSurfacesMapper.UnmapIOSurfacesWithModelRequestError]
//   - [IANEProgramIOSurfacesMapper.ValidateRequestModel]
//   - [IANEProgramIOSurfacesMapper.InitWithController]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper
type IANEProgramIOSurfacesMapper interface {
	objectivec.IObject

	// Topic: Methods

	Controller() *ANEDeviceController
	DeviceController() *ANEDeviceController
	MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error)
	PrepareANEMemoryMappingParamsRequest(params unsafe.Pointer, request objectivec.IObject)
	ProgramHandle() uint64
	UnmapIOSurfacesWithModelRequestError(model objectivec.IObject, request objectivec.IObject) (bool, error)
	ValidateRequestModel(request objectivec.IObject, model objectivec.IObject) bool
	InitWithController(controller objectivec.IObject) ANEProgramIOSurfacesMapper
}

// Init initializes the instance.
func (a ANEProgramIOSurfacesMapper) Init() ANEProgramIOSurfacesMapper {
	rv := objc.Send[ANEProgramIOSurfacesMapper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEProgramIOSurfacesMapper) Autorelease() ANEProgramIOSurfacesMapper {
	rv := objc.Send[ANEProgramIOSurfacesMapper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEProgramIOSurfacesMapper creates a new ANEProgramIOSurfacesMapper instance.
func NewANEProgramIOSurfacesMapper() ANEProgramIOSurfacesMapper {
	class := getANEProgramIOSurfacesMapperClass()
	rv := objc.Send[ANEProgramIOSurfacesMapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/initWithController:
func NewANEProgramIOSurfacesMapperWithController(controller objectivec.IObject) ANEProgramIOSurfacesMapper {
	instance := getANEProgramIOSurfacesMapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithController:"), controller)
	return ANEProgramIOSurfacesMapperFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/mapIOSurfacesWithModel:request:cacheInference:error:
func (a ANEProgramIOSurfacesMapper) MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("mapIOSurfacesWithModel:request:cacheInference:error:"), model, request, inference, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("mapIOSurfacesWithModel:request:cacheInference:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/prepareANEMemoryMappingParams:request:
func (a ANEProgramIOSurfacesMapper) PrepareANEMemoryMappingParamsRequest(params unsafe.Pointer, request objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("prepareANEMemoryMappingParams:request:"), params, request)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/unmapIOSurfacesWithModel:request:error:
func (a ANEProgramIOSurfacesMapper) UnmapIOSurfacesWithModelRequestError(model objectivec.IObject, request objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("unmapIOSurfacesWithModel:request:error:"), model, request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unmapIOSurfacesWithModel:request:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/validateRequest:model:
func (a ANEProgramIOSurfacesMapper) ValidateRequestModel(request objectivec.IObject, model objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validateRequest:model:"), request, model)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/initWithController:
func (a ANEProgramIOSurfacesMapper) InitWithController(controller objectivec.IObject) ANEProgramIOSurfacesMapper {
	rv := objc.Send[ANEProgramIOSurfacesMapper](a.ID, objc.Sel("initWithController:"), controller)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/mapperWithController:
func (_ANEProgramIOSurfacesMapperClass ANEProgramIOSurfacesMapperClass) MapperWithController(controller objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEProgramIOSurfacesMapperClass.class), objc.Sel("mapperWithController:"), controller)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/mapperWithProgramHandle:
func (_ANEProgramIOSurfacesMapperClass ANEProgramIOSurfacesMapperClass) MapperWithProgramHandle(handle uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEProgramIOSurfacesMapperClass.class), objc.Sel("mapperWithProgramHandle:"), handle)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/controller
func (a ANEProgramIOSurfacesMapper) Controller() *ANEDeviceController {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("controller"))
	if rv == 0 {
		return nil
	}
	val := ANEDeviceControllerFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/deviceController
func (a ANEProgramIOSurfacesMapper) DeviceController() *ANEDeviceController {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("deviceController"))
	if rv == 0 {
		return nil
	}
	val := ANEDeviceControllerFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramIOSurfacesMapper/programHandle
func (a ANEProgramIOSurfacesMapper) ProgramHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
