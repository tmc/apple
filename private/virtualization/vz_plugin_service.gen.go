// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPluginService] class.
var (
	_VZPluginServiceClass     VZPluginServiceClass
	_VZPluginServiceClassOnce sync.Once
)

func getVZPluginServiceClass() VZPluginServiceClass {
	_VZPluginServiceClassOnce.Do(func() {
		_VZPluginServiceClass = VZPluginServiceClass{class: objc.GetClass("_VZPluginService")}
	})
	return _VZPluginServiceClass
}

// GetVZPluginServiceClass returns the class object for _VZPluginService.
func GetVZPluginServiceClass() VZPluginServiceClass {
	return getVZPluginServiceClass()
}

type VZPluginServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPluginServiceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPluginServiceClass) Alloc() VZPluginService {
	rv := objc.Send[VZPluginService](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPluginService
type VZPluginService struct {
	objectivec.Object
}

// VZPluginServiceFromID constructs a [VZPluginService] from an objc.ID.
func VZPluginServiceFromID(id objc.ID) VZPluginService {
	return VZPluginService{objectivec.Object{ID: id}}
}
// Ensure VZPluginService implements IVZPluginService.
var _ IVZPluginService = VZPluginService{}

// An interface definition for the [VZPluginService] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPluginService
type IVZPluginService interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZPluginService) Init() VZPluginService {
	rv := objc.Send[VZPluginService](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPluginService) Autorelease() VZPluginService {
	rv := objc.Send[VZPluginService](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPluginService creates a new VZPluginService instance.
func NewVZPluginService() VZPluginService {
	class := getVZPluginServiceClass()
	rv := objc.Send[VZPluginService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPluginService/xpcMain
func (_VZPluginServiceClass VZPluginServiceClass) XpcMain() int {
	rv := objc.Send[int](objc.ID(_VZPluginServiceClass.class), objc.Sel("xpcMain"))
	return rv
}

