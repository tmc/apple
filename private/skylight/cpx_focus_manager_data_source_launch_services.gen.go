// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXFocusManagerDataSourceLaunchServices] class.
var (
	_CPXFocusManagerDataSourceLaunchServicesClass     CPXFocusManagerDataSourceLaunchServicesClass
	_CPXFocusManagerDataSourceLaunchServicesClassOnce sync.Once
)

func getCPXFocusManagerDataSourceLaunchServicesClass() CPXFocusManagerDataSourceLaunchServicesClass {
	_CPXFocusManagerDataSourceLaunchServicesClassOnce.Do(func() {
		_CPXFocusManagerDataSourceLaunchServicesClass = CPXFocusManagerDataSourceLaunchServicesClass{class: objc.GetClass("CPXFocusManagerDataSourceLaunchServices")}
	})
	return _CPXFocusManagerDataSourceLaunchServicesClass
}

// GetCPXFocusManagerDataSourceLaunchServicesClass returns the class object for CPXFocusManagerDataSourceLaunchServices.
func GetCPXFocusManagerDataSourceLaunchServicesClass() CPXFocusManagerDataSourceLaunchServicesClass {
	return getCPXFocusManagerDataSourceLaunchServicesClass()
}

type CPXFocusManagerDataSourceLaunchServicesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXFocusManagerDataSourceLaunchServicesClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXFocusManagerDataSourceLaunchServicesClass) Alloc() CPXFocusManagerDataSourceLaunchServices {
	rv := objc.Send[CPXFocusManagerDataSourceLaunchServices](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXFocusManagerDataSourceLaunchServices.AddToPermittedFrontList]
//   - [CPXFocusManagerDataSourceLaunchServices.AppendDescriptionToStream]
//   - [CPXFocusManagerDataSourceLaunchServices.FrontmostProcess]
//   - [CPXFocusManagerDataSourceLaunchServices.GetProcessToBringForwardAtNextCheckin]
//   - [CPXFocusManagerDataSourceLaunchServices.IsProcessPermittedToBeFrontmost]
//   - [CPXFocusManagerDataSourceLaunchServices.IsProcessToBringForwardAtNextCheckin]
//   - [CPXFocusManagerDataSourceLaunchServices.KeyThiefConnectionID]
//   - [CPXFocusManagerDataSourceLaunchServices.RemoveFromPermittedFrontList]
//   - [CPXFocusManagerDataSourceLaunchServices.SetKeyThiefConnectionID]
//   - [CPXFocusManagerDataSourceLaunchServices.SetProcessToBringForwardAtNextCheckin]
//   - [CPXFocusManagerDataSourceLaunchServices.InitWithLaunchServicesProviderProcessManager]
//   - [CPXFocusManagerDataSourceLaunchServices.DebugDescription]
//   - [CPXFocusManagerDataSourceLaunchServices.Description]
//   - [CPXFocusManagerDataSourceLaunchServices.Hash]
//   - [CPXFocusManagerDataSourceLaunchServices.Superclass]
type CPXFocusManagerDataSourceLaunchServices struct {
	objectivec.Object
}

// CPXFocusManagerDataSourceLaunchServicesFromID constructs a [CPXFocusManagerDataSourceLaunchServices] from an objc.ID.
func CPXFocusManagerDataSourceLaunchServicesFromID(id objc.ID) CPXFocusManagerDataSourceLaunchServices {
	return CPXFocusManagerDataSourceLaunchServices{objectivec.Object{ID: id}}
}

// Ensure CPXFocusManagerDataSourceLaunchServices implements ICPXFocusManagerDataSourceLaunchServices.
var _ ICPXFocusManagerDataSourceLaunchServices = CPXFocusManagerDataSourceLaunchServices{}

// An interface definition for the [CPXFocusManagerDataSourceLaunchServices] class.
//
// # Methods
//
//   - [ICPXFocusManagerDataSourceLaunchServices.AddToPermittedFrontList]
//   - [ICPXFocusManagerDataSourceLaunchServices.AppendDescriptionToStream]
//   - [ICPXFocusManagerDataSourceLaunchServices.FrontmostProcess]
//   - [ICPXFocusManagerDataSourceLaunchServices.GetProcessToBringForwardAtNextCheckin]
//   - [ICPXFocusManagerDataSourceLaunchServices.IsProcessPermittedToBeFrontmost]
//   - [ICPXFocusManagerDataSourceLaunchServices.IsProcessToBringForwardAtNextCheckin]
//   - [ICPXFocusManagerDataSourceLaunchServices.KeyThiefConnectionID]
//   - [ICPXFocusManagerDataSourceLaunchServices.RemoveFromPermittedFrontList]
//   - [ICPXFocusManagerDataSourceLaunchServices.SetKeyThiefConnectionID]
//   - [ICPXFocusManagerDataSourceLaunchServices.SetProcessToBringForwardAtNextCheckin]
//   - [ICPXFocusManagerDataSourceLaunchServices.InitWithLaunchServicesProviderProcessManager]
//   - [ICPXFocusManagerDataSourceLaunchServices.DebugDescription]
//   - [ICPXFocusManagerDataSourceLaunchServices.Description]
//   - [ICPXFocusManagerDataSourceLaunchServices.Hash]
//   - [ICPXFocusManagerDataSourceLaunchServices.Superclass]
type ICPXFocusManagerDataSourceLaunchServices interface {
	objectivec.IObject

	// Topic: Methods

	AddToPermittedFrontList(list CPSProcessSerNum) int16
	AppendDescriptionToStream(stream objectivec.IObject)
	FrontmostProcess() unsafe.Pointer
	GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool
	IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool
	KeyThiefConnectionID() uint32
	RemoveFromPermittedFrontList(list CPSProcessSerNum) int16
	SetKeyThiefConnectionID(id uint32)
	SetProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) int
	InitWithLaunchServicesProviderProcessManager(provider objectivec.IObject, manager objectivec.IObject) CPXFocusManagerDataSourceLaunchServices
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CPXFocusManagerDataSourceLaunchServices) Init() CPXFocusManagerDataSourceLaunchServices {
	rv := objc.Send[CPXFocusManagerDataSourceLaunchServices](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXFocusManagerDataSourceLaunchServices) Autorelease() CPXFocusManagerDataSourceLaunchServices {
	rv := objc.Send[CPXFocusManagerDataSourceLaunchServices](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXFocusManagerDataSourceLaunchServices creates a new CPXFocusManagerDataSourceLaunchServices instance.
func NewCPXFocusManagerDataSourceLaunchServices() CPXFocusManagerDataSourceLaunchServices {
	class := getCPXFocusManagerDataSourceLaunchServicesClass()
	rv := objc.Send[CPXFocusManagerDataSourceLaunchServices](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXFocusManagerDataSourceLaunchServicesWithLaunchServicesProviderProcessManager(provider objectivec.IObject, manager objectivec.IObject) CPXFocusManagerDataSourceLaunchServices {
	instance := getCPXFocusManagerDataSourceLaunchServicesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLaunchServicesProvider:processManager:"), provider, manager)
	return CPXFocusManagerDataSourceLaunchServicesFromID(rv)
}

func (c CPXFocusManagerDataSourceLaunchServices) AddToPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("appendDescriptionToStream:"), stream)
}
func (c CPXFocusManagerDataSourceLaunchServices) GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) RemoveFromPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) SetKeyThiefConnectionID(id uint32) {
	objc.Send[objc.ID](c.ID, objc.Sel("setKeyThiefConnectionID:"), id)
}
func (c CPXFocusManagerDataSourceLaunchServices) SetProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) int {
	rv := objc.Send[int](c.ID, objc.Sel("setProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) InitWithLaunchServicesProviderProcessManager(provider objectivec.IObject, manager objectivec.IObject) CPXFocusManagerDataSourceLaunchServices {
	rv := objc.Send[CPXFocusManagerDataSourceLaunchServices](c.ID, objc.Sel("initWithLaunchServicesProvider:processManager:"), provider, manager)
	return rv
}

func (c CPXFocusManagerDataSourceLaunchServices) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXFocusManagerDataSourceLaunchServices) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXFocusManagerDataSourceLaunchServices) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontmostProcess"))
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXFocusManagerDataSourceLaunchServices) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
