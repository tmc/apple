// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VirtualizationMockEventService] class.
var (
	_VirtualizationMockEventServiceClass     VirtualizationMockEventServiceClass
	_VirtualizationMockEventServiceClassOnce sync.Once
)

func getVirtualizationMockEventServiceClass() VirtualizationMockEventServiceClass {
	_VirtualizationMockEventServiceClassOnce.Do(func() {
		_VirtualizationMockEventServiceClass = VirtualizationMockEventServiceClass{class: objc.GetClass("Virtualization.MockEventService")}
	})
	return _VirtualizationMockEventServiceClass
}

// GetVirtualizationMockEventServiceClass returns the class object for Virtualization.MockEventService.
func GetVirtualizationMockEventServiceClass() VirtualizationMockEventServiceClass {
	return getVirtualizationMockEventServiceClass()
}

type VirtualizationMockEventServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VirtualizationMockEventServiceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VirtualizationMockEventServiceClass) Alloc() VirtualizationMockEventService {
	rv := objc.SendIfResponds[VirtualizationMockEventService](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VirtualizationMockEventService struct {
	objectivec.Object
}

// VirtualizationMockEventServiceFromID constructs a [VirtualizationMockEventService] from an objc.ID.
func VirtualizationMockEventServiceFromID(id objc.ID) VirtualizationMockEventService {
	return VirtualizationMockEventService{objectivec.Object{ID: id}}
}

// Ensure VirtualizationMockEventService implements IVirtualizationMockEventService.
var _ IVirtualizationMockEventService = VirtualizationMockEventService{}

// An interface definition for the [VirtualizationMockEventService] class.
type IVirtualizationMockEventService interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VirtualizationMockEventService) Init() VirtualizationMockEventService {
	rv := objc.SendIfResponds[VirtualizationMockEventService](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VirtualizationMockEventService) Autorelease() VirtualizationMockEventService {
	rv := objc.SendIfResponds[VirtualizationMockEventService](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVirtualizationMockEventService creates a new VirtualizationMockEventService instance.
func NewVirtualizationMockEventService() VirtualizationMockEventService {
	class := getVirtualizationMockEventServiceClass()
	rv := objc.SendIfResponds[VirtualizationMockEventService](objc.ID(class.class), objc.Sel("new"))
	return rv
}
