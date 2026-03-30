// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEDeviceController] class.
var (
	_ANEDeviceControllerClass     ANEDeviceControllerClass
	_ANEDeviceControllerClassOnce sync.Once
)

func getANEDeviceControllerClass() ANEDeviceControllerClass {
	_ANEDeviceControllerClassOnce.Do(func() {
		_ANEDeviceControllerClass = ANEDeviceControllerClass{class: objc.GetClass("_ANEDeviceController")}
	})
	return _ANEDeviceControllerClass
}

// GetANEDeviceControllerClass returns the class object for _ANEDeviceController.
func GetANEDeviceControllerClass() ANEDeviceControllerClass {
	return getANEDeviceControllerClass()
}

type ANEDeviceControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEDeviceControllerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEDeviceControllerClass) Alloc() ANEDeviceController {
	rv := objc.Send[ANEDeviceController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEDeviceController.Device]
//   - [ANEDeviceController.SetDevice]
//   - [ANEDeviceController.IsPrivileged]
//   - [ANEDeviceController.ProgramHandle]
//   - [ANEDeviceController.Start]
//   - [ANEDeviceController.Stop]
//   - [ANEDeviceController.Usecount]
//   - [ANEDeviceController.SetUsecount]
//   - [ANEDeviceController.InitWithANEPrivilegedVM]
//   - [ANEDeviceController.InitWithProgramHandlePriviledged]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController
type ANEDeviceController struct {
	objectivec.Object
}

// ANEDeviceControllerFromID constructs a [ANEDeviceController] from an objc.ID.
func ANEDeviceControllerFromID(id objc.ID) ANEDeviceController {
	return ANEDeviceController{objectivec.Object{ID: id}}
}

// Ensure ANEDeviceController implements IANEDeviceController.
var _ IANEDeviceController = ANEDeviceController{}

// An interface definition for the [ANEDeviceController] class.
//
// # Methods
//
//   - [IANEDeviceController.Device]
//   - [IANEDeviceController.SetDevice]
//   - [IANEDeviceController.IsPrivileged]
//   - [IANEDeviceController.ProgramHandle]
//   - [IANEDeviceController.Start]
//   - [IANEDeviceController.Stop]
//   - [IANEDeviceController.Usecount]
//   - [IANEDeviceController.SetUsecount]
//   - [IANEDeviceController.InitWithANEPrivilegedVM]
//   - [IANEDeviceController.InitWithProgramHandlePriviledged]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController
type IANEDeviceController interface {
	objectivec.IObject

	// Topic: Methods

	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	IsPrivileged() bool
	ProgramHandle() uint64
	Start()
	Stop()
	Usecount() int64
	SetUsecount(value int64)
	InitWithANEPrivilegedVM(vm bool) ANEDeviceController
	InitWithProgramHandlePriviledged(handle uint64, priviledged bool) ANEDeviceController
}

// Init initializes the instance.
func (a ANEDeviceController) Init() ANEDeviceController {
	rv := objc.Send[ANEDeviceController](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEDeviceController) Autorelease() ANEDeviceController {
	rv := objc.Send[ANEDeviceController](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEDeviceController creates a new ANEDeviceController instance.
func NewANEDeviceController() ANEDeviceController {
	class := getANEDeviceControllerClass()
	rv := objc.Send[ANEDeviceController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/initWithANEPrivilegedVM:
func NewANEDeviceControllerWithANEPrivilegedVM(vm bool) ANEDeviceController {
	instance := getANEDeviceControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithANEPrivilegedVM:"), vm)
	return ANEDeviceControllerFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/initWithProgramHandle:priviledged:
func NewANEDeviceControllerWithProgramHandlePriviledged(handle uint64, priviledged bool) ANEDeviceController {
	instance := getANEDeviceControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramHandle:priviledged:"), handle, priviledged)
	return ANEDeviceControllerFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/start
func (a ANEDeviceController) Start() {
	objc.Send[objc.ID](a.ID, objc.Sel("start"))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/stop
func (a ANEDeviceController) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/initWithANEPrivilegedVM:
func (a ANEDeviceController) InitWithANEPrivilegedVM(vm bool) ANEDeviceController {
	rv := objc.Send[ANEDeviceController](a.ID, objc.Sel("initWithANEPrivilegedVM:"), vm)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/initWithProgramHandle:priviledged:
func (a ANEDeviceController) InitWithProgramHandlePriviledged(handle uint64, priviledged bool) ANEDeviceController {
	rv := objc.Send[ANEDeviceController](a.ID, objc.Sel("initWithProgramHandle:priviledged:"), handle, priviledged)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/controllerWithPrivilegedVM:
func (_ANEDeviceControllerClass ANEDeviceControllerClass) ControllerWithPrivilegedVM(vm bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceControllerClass.class), objc.Sel("controllerWithPrivilegedVM:"), vm)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/controllerWithProgramHandle:
func (_ANEDeviceControllerClass ANEDeviceControllerClass) ControllerWithProgramHandle(handle uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceControllerClass.class), objc.Sel("controllerWithProgramHandle:"), handle)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/sharedPrivilegedConnection
func (_ANEDeviceControllerClass ANEDeviceControllerClass) SharedPrivilegedConnection() *ANEDeviceController {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceControllerClass.class), objc.Sel("sharedPrivilegedConnection"))
	if rv == 0 {
		return nil
	}
	val := ANEDeviceControllerFromID(rv)
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/device
func (a ANEDeviceController) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("device"))
	return rv
}
func (a ANEDeviceController) SetDevice(value unsafe.Pointer) {
	objc.Send[struct{}](a.ID, objc.Sel("setDevice:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/isPrivileged
func (a ANEDeviceController) IsPrivileged() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPrivileged"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/programHandle
func (a ANEDeviceController) ProgramHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceController/usecount
func (a ANEDeviceController) Usecount() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("usecount"))
	return rv
}
func (a ANEDeviceController) SetUsecount(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setUsecount:"), value)
}
