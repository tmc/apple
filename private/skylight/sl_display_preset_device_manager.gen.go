// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDisplayPresetDeviceManager] class.
var (
	_SLDisplayPresetDeviceManagerClass     SLDisplayPresetDeviceManagerClass
	_SLDisplayPresetDeviceManagerClassOnce sync.Once
)

func getSLDisplayPresetDeviceManagerClass() SLDisplayPresetDeviceManagerClass {
	_SLDisplayPresetDeviceManagerClassOnce.Do(func() {
		_SLDisplayPresetDeviceManagerClass = SLDisplayPresetDeviceManagerClass{class: objc.GetClass("SLDisplayPresetDeviceManager")}
	})
	return _SLDisplayPresetDeviceManagerClass
}

// GetSLDisplayPresetDeviceManagerClass returns the class object for SLDisplayPresetDeviceManager.
func GetSLDisplayPresetDeviceManagerClass() SLDisplayPresetDeviceManagerClass {
	return getSLDisplayPresetDeviceManagerClass()
}

type SLDisplayPresetDeviceManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDisplayPresetDeviceManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDisplayPresetDeviceManagerClass) Alloc() SLDisplayPresetDeviceManager {
	rv := objc.Send[SLDisplayPresetDeviceManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDisplayPresetDeviceManager._serviceAdded]
//   - [SLDisplayPresetDeviceManager._serviceRemoved]
//   - [SLDisplayPresetDeviceManager.CopyDeviceForContainer]
//   - [SLDisplayPresetDeviceManager.CopyDevices]
//   - [SLDisplayPresetDeviceManager.ServiceIsValidFor]
//   - [SLDisplayPresetDeviceManager.StartWithBlockOnQueue]
//   - [SLDisplayPresetDeviceManager.Stop]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager
type SLDisplayPresetDeviceManager struct {
	objectivec.Object
}

// SLDisplayPresetDeviceManagerFromID constructs a [SLDisplayPresetDeviceManager] from an objc.ID.
func SLDisplayPresetDeviceManagerFromID(id objc.ID) SLDisplayPresetDeviceManager {
	return SLDisplayPresetDeviceManager{objectivec.Object{ID: id}}
}

// Ensure SLDisplayPresetDeviceManager implements ISLDisplayPresetDeviceManager.
var _ ISLDisplayPresetDeviceManager = SLDisplayPresetDeviceManager{}

// An interface definition for the [SLDisplayPresetDeviceManager] class.
//
// # Methods
//
//   - [ISLDisplayPresetDeviceManager._serviceAdded]
//   - [ISLDisplayPresetDeviceManager._serviceRemoved]
//   - [ISLDisplayPresetDeviceManager.CopyDeviceForContainer]
//   - [ISLDisplayPresetDeviceManager.CopyDevices]
//   - [ISLDisplayPresetDeviceManager.ServiceIsValidFor]
//   - [ISLDisplayPresetDeviceManager.StartWithBlockOnQueue]
//   - [ISLDisplayPresetDeviceManager.Stop]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager
type ISLDisplayPresetDeviceManager interface {
	objectivec.IObject

	// Topic: Methods

	_serviceAdded(added uint32)
	_serviceRemoved(removed uint32)
	CopyDeviceForContainer(container objectivec.IObject) objectivec.IObject
	CopyDevices() objectivec.IObject
	ServiceIsValidFor(for_ objectivec.IObject) bool
	StartWithBlockOnQueue(block VoidHandler, queue objectivec.IObject)
	Stop()
}

// Init initializes the instance.
func (s SLDisplayPresetDeviceManager) Init() SLDisplayPresetDeviceManager {
	rv := objc.Send[SLDisplayPresetDeviceManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDisplayPresetDeviceManager) Autorelease() SLDisplayPresetDeviceManager {
	rv := objc.Send[SLDisplayPresetDeviceManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDisplayPresetDeviceManager creates a new SLDisplayPresetDeviceManager instance.
func NewSLDisplayPresetDeviceManager() SLDisplayPresetDeviceManager {
	class := getSLDisplayPresetDeviceManagerClass()
	rv := objc.Send[SLDisplayPresetDeviceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/_serviceAdded:
func (s SLDisplayPresetDeviceManager) _serviceAdded(added uint32) {
	objc.Send[objc.ID](s.ID, objc.Sel("_serviceAdded:"), added)
}

// ServiceAdded is an exported wrapper for the private method _serviceAdded.
func (s SLDisplayPresetDeviceManager) ServiceAdded(added uint32) {
	s._serviceAdded(added)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/_serviceRemoved:
func (s SLDisplayPresetDeviceManager) _serviceRemoved(removed uint32) {
	objc.Send[objc.ID](s.ID, objc.Sel("_serviceRemoved:"), removed)
}

// ServiceRemoved is an exported wrapper for the private method _serviceRemoved.
func (s SLDisplayPresetDeviceManager) ServiceRemoved(removed uint32) {
	s._serviceRemoved(removed)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/copyDeviceForContainer:
func (s SLDisplayPresetDeviceManager) CopyDeviceForContainer(container objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyDeviceForContainer:"), container)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/copyDevices
func (s SLDisplayPresetDeviceManager) CopyDevices() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyDevices"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/serviceIsValidFor:
func (s SLDisplayPresetDeviceManager) ServiceIsValidFor(for_ objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("serviceIsValidFor:"), for_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/startWithBlock:onQueue:
func (s SLDisplayPresetDeviceManager) StartWithBlockOnQueue(block VoidHandler, queue objectivec.IObject) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("startWithBlock:onQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/stop
func (s SLDisplayPresetDeviceManager) Stop() {
	objc.Send[objc.ID](s.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/debugDeferArrivalSeconds
func (_SLDisplayPresetDeviceManagerClass SLDisplayPresetDeviceManagerClass) DebugDeferArrivalSeconds() float32 {
	rv := objc.Send[float32](objc.ID(_SLDisplayPresetDeviceManagerClass.class), objc.Sel("debugDeferArrivalSeconds"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/instance
func (_SLDisplayPresetDeviceManagerClass SLDisplayPresetDeviceManagerClass) Instance() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLDisplayPresetDeviceManagerClass.class), objc.Sel("instance"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDeviceManager/setDebugDeferArrivalSeconds:
func (_SLDisplayPresetDeviceManagerClass SLDisplayPresetDeviceManagerClass) SetDebugDeferArrivalSeconds(seconds float32) {
	objc.Send[objc.ID](objc.ID(_SLDisplayPresetDeviceManagerClass.class), objc.Sel("setDebugDeferArrivalSeconds:"), seconds)
}
