// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXLaunchServicesInterface] class.
var (
	_CPXLaunchServicesInterfaceClass     CPXLaunchServicesInterfaceClass
	_CPXLaunchServicesInterfaceClassOnce sync.Once
)

func getCPXLaunchServicesInterfaceClass() CPXLaunchServicesInterfaceClass {
	_CPXLaunchServicesInterfaceClassOnce.Do(func() {
		_CPXLaunchServicesInterfaceClass = CPXLaunchServicesInterfaceClass{class: objc.GetClass("CPXLaunchServicesInterface")}
	})
	return _CPXLaunchServicesInterfaceClass
}

// GetCPXLaunchServicesInterfaceClass returns the class object for CPXLaunchServicesInterface.
func GetCPXLaunchServicesInterfaceClass() CPXLaunchServicesInterfaceClass {
	return getCPXLaunchServicesInterfaceClass()
}

type CPXLaunchServicesInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXLaunchServicesInterfaceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXLaunchServicesInterfaceClass) Alloc() CPXLaunchServicesInterface {
	rv := objc.SendIfResponds[CPXLaunchServicesInterface](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXLaunchServicesInterface.AddToPermittedFrontASNList]
//   - [CPXLaunchServicesInterface.ApplicationInformationSeed]
//   - [CPXLaunchServicesInterface.ApplicationType]
//   - [CPXLaunchServicesInterface.FrontApplication]
//   - [CPXLaunchServicesInterface.FrontApplicationSeed]
//   - [CPXLaunchServicesInterface.PermittedFrontApplications]
//   - [CPXLaunchServicesInterface.RemoveFromPermittedFrontASNList]
//   - [CPXLaunchServicesInterface.RunningApplications]
//   - [CPXLaunchServicesInterface.SessionID]
//   - [CPXLaunchServicesInterface.InitWithSessionID]
type CPXLaunchServicesInterface struct {
	objectivec.Object
}

// CPXLaunchServicesInterfaceFromID constructs a [CPXLaunchServicesInterface] from an objc.ID.
func CPXLaunchServicesInterfaceFromID(id objc.ID) CPXLaunchServicesInterface {
	return CPXLaunchServicesInterface{objectivec.Object{ID: id}}
}

// Ensure CPXLaunchServicesInterface implements ICPXLaunchServicesInterface.
var _ ICPXLaunchServicesInterface = CPXLaunchServicesInterface{}

// An interface definition for the [CPXLaunchServicesInterface] class.
//
// # Methods
//
//   - [ICPXLaunchServicesInterface.AddToPermittedFrontASNList]
//   - [ICPXLaunchServicesInterface.ApplicationInformationSeed]
//   - [ICPXLaunchServicesInterface.ApplicationType]
//   - [ICPXLaunchServicesInterface.FrontApplication]
//   - [ICPXLaunchServicesInterface.FrontApplicationSeed]
//   - [ICPXLaunchServicesInterface.PermittedFrontApplications]
//   - [ICPXLaunchServicesInterface.RemoveFromPermittedFrontASNList]
//   - [ICPXLaunchServicesInterface.RunningApplications]
//   - [ICPXLaunchServicesInterface.SessionID]
//   - [ICPXLaunchServicesInterface.InitWithSessionID]
type ICPXLaunchServicesInterface interface {
	objectivec.IObject

	// Topic: Methods

	AddToPermittedFrontASNList(aSNList LSASNRef) int
	ApplicationInformationSeed(seed LSASNRef) int
	ApplicationType(type_ LSASNRef) byte
	FrontApplication() objectivec.IObject
	FrontApplicationSeed() uint32
	PermittedFrontApplications() foundation.INSArray
	RemoveFromPermittedFrontASNList(aSNList LSASNRef) int
	RunningApplications() foundation.INSArray
	SessionID() int
	InitWithSessionID(id int) CPXLaunchServicesInterface
}

// Init initializes the instance.
func (c CPXLaunchServicesInterface) Init() CPXLaunchServicesInterface {
	rv := objc.SendIfResponds[CPXLaunchServicesInterface](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXLaunchServicesInterface) Autorelease() CPXLaunchServicesInterface {
	rv := objc.SendIfResponds[CPXLaunchServicesInterface](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXLaunchServicesInterface creates a new CPXLaunchServicesInterface instance.
func NewCPXLaunchServicesInterface() CPXLaunchServicesInterface {
	class := getCPXLaunchServicesInterfaceClass()
	rv := objc.SendIfResponds[CPXLaunchServicesInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXLaunchServicesInterfaceWithSessionID(id int) CPXLaunchServicesInterface {
	instance := getCPXLaunchServicesInterfaceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSessionID:"), id)
	return CPXLaunchServicesInterfaceFromID(rv)
}

func (c CPXLaunchServicesInterface) AddToPermittedFrontASNList(aSNList LSASNRef) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("addToPermittedFrontASNList:"), aSNList)
	return rv
}
func (c CPXLaunchServicesInterface) ApplicationInformationSeed(seed LSASNRef) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("applicationInformationSeed:"), seed)
	return rv
}
func (c CPXLaunchServicesInterface) ApplicationType(type_ LSASNRef) byte {
	rv := objc.SendIfResponds[byte](c.ID, objc.Sel("applicationType:"), type_)
	return rv
}
func (c CPXLaunchServicesInterface) RemoveFromPermittedFrontASNList(aSNList LSASNRef) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("removeFromPermittedFrontASNList:"), aSNList)
	return rv
}
func (c CPXLaunchServicesInterface) InitWithSessionID(id int) CPXLaunchServicesInterface {
	rv := objc.SendIfResponds[CPXLaunchServicesInterface](c.ID, objc.Sel("initWithSessionID:"), id)
	return rv
}

func (c CPXLaunchServicesInterface) FrontApplication() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("frontApplication"))
	return objectivec.Object{ID: rv}
}
func (c CPXLaunchServicesInterface) FrontApplicationSeed() uint32 {
	rv := objc.SendIfResponds[uint32](c.ID, objc.Sel("frontApplicationSeed"))
	return rv
}
func (c CPXLaunchServicesInterface) PermittedFrontApplications() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("permittedFrontApplications"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (c CPXLaunchServicesInterface) RunningApplications() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("runningApplications"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (c CPXLaunchServicesInterface) SessionID() int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("sessionID"))
	return rv
}
