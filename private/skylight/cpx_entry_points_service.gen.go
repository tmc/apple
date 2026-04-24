// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXEntryPointsService] class.
var (
	_CPXEntryPointsServiceClass     CPXEntryPointsServiceClass
	_CPXEntryPointsServiceClassOnce sync.Once
)

func getCPXEntryPointsServiceClass() CPXEntryPointsServiceClass {
	_CPXEntryPointsServiceClassOnce.Do(func() {
		_CPXEntryPointsServiceClass = CPXEntryPointsServiceClass{class: objc.GetClass("CPXEntryPointsService")}
	})
	return _CPXEntryPointsServiceClass
}

// GetCPXEntryPointsServiceClass returns the class object for CPXEntryPointsService.
func GetCPXEntryPointsServiceClass() CPXEntryPointsServiceClass {
	return getCPXEntryPointsServiceClass()
}

type CPXEntryPointsServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXEntryPointsServiceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXEntryPointsServiceClass) Alloc() CPXEntryPointsService {
	rv := objc.Send[CPXEntryPointsService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXEntryPointsService.ClientAddToPermittedFrontList]
//   - [CPXEntryPointsService.ClientRemoveFromPermittedFrontList]
//   - [CPXEntryPointsService.InitWithFocusControllerProcessManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService
type CPXEntryPointsService struct {
	objectivec.Object
}

// CPXEntryPointsServiceFromID constructs a [CPXEntryPointsService] from an objc.ID.
func CPXEntryPointsServiceFromID(id objc.ID) CPXEntryPointsService {
	return CPXEntryPointsService{objectivec.Object{ID: id}}
}

// Ensure CPXEntryPointsService implements ICPXEntryPointsService.
var _ ICPXEntryPointsService = CPXEntryPointsService{}

// An interface definition for the [CPXEntryPointsService] class.
//
// # Methods
//
//   - [ICPXEntryPointsService.ClientAddToPermittedFrontList]
//   - [ICPXEntryPointsService.ClientRemoveFromPermittedFrontList]
//   - [ICPXEntryPointsService.InitWithFocusControllerProcessManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService
type ICPXEntryPointsService interface {
	objectivec.IObject

	// Topic: Methods

	ClientAddToPermittedFrontList(client unsafe.Pointer, list objectivec.IObject) int
	ClientRemoveFromPermittedFrontList(client unsafe.Pointer, list objectivec.IObject) int
	InitWithFocusControllerProcessManager(controller objectivec.IObject, manager objectivec.IObject) CPXEntryPointsService
}

// Init initializes the instance.
func (c CPXEntryPointsService) Init() CPXEntryPointsService {
	rv := objc.Send[CPXEntryPointsService](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXEntryPointsService) Autorelease() CPXEntryPointsService {
	rv := objc.Send[CPXEntryPointsService](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXEntryPointsService creates a new CPXEntryPointsService instance.
func NewCPXEntryPointsService() CPXEntryPointsService {
	class := getCPXEntryPointsServiceClass()
	rv := objc.Send[CPXEntryPointsService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService/initWithFocusController:processManager:
func NewCPXEntryPointsServiceWithFocusControllerProcessManager(controller objectivec.IObject, manager objectivec.IObject) CPXEntryPointsService {
	instance := getCPXEntryPointsServiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFocusController:processManager:"), controller, manager)
	return CPXEntryPointsServiceFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService/client:addToPermittedFrontList:
func (c CPXEntryPointsService) ClientAddToPermittedFrontList(client unsafe.Pointer, list objectivec.IObject) int {
	rv := objc.Send[int](c.ID, objc.Sel("client:addToPermittedFrontList:"), client, list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService/client:removeFromPermittedFrontList:
func (c CPXEntryPointsService) ClientRemoveFromPermittedFrontList(client unsafe.Pointer, list objectivec.IObject) int {
	rv := objc.Send[int](c.ID, objc.Sel("client:removeFromPermittedFrontList:"), client, list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEntryPointsService/initWithFocusController:processManager:
func (c CPXEntryPointsService) InitWithFocusControllerProcessManager(controller objectivec.IObject, manager objectivec.IObject) CPXEntryPointsService {
	rv := objc.Send[CPXEntryPointsService](c.ID, objc.Sel("initWithFocusController:processManager:"), controller, manager)
	return rv
}
