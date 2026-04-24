// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXConnectionManager] class.
var (
	_CPXConnectionManagerClass     CPXConnectionManagerClass
	_CPXConnectionManagerClassOnce sync.Once
)

func getCPXConnectionManagerClass() CPXConnectionManagerClass {
	_CPXConnectionManagerClassOnce.Do(func() {
		_CPXConnectionManagerClass = CPXConnectionManagerClass{class: objc.GetClass("CPXConnectionManager")}
	})
	return _CPXConnectionManagerClass
}

// GetCPXConnectionManagerClass returns the class object for CPXConnectionManager.
func GetCPXConnectionManagerClass() CPXConnectionManagerClass {
	return getCPXConnectionManagerClass()
}

type CPXConnectionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXConnectionManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXConnectionManagerClass) Alloc() CPXConnectionManager {
	rv := objc.Send[CPXConnectionManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXConnectionManager.ConnectionForID]
//   - [CPXConnectionManager.PidForConnection]
//   - [CPXConnectionManager.InitWithSession]
//   - [CPXConnectionManager.DebugDescription]
//   - [CPXConnectionManager.Description]
//   - [CPXConnectionManager.Hash]
//   - [CPXConnectionManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager
type CPXConnectionManager struct {
	objectivec.Object
}

// CPXConnectionManagerFromID constructs a [CPXConnectionManager] from an objc.ID.
func CPXConnectionManagerFromID(id objc.ID) CPXConnectionManager {
	return CPXConnectionManager{objectivec.Object{ID: id}}
}

// Ensure CPXConnectionManager implements ICPXConnectionManager.
var _ ICPXConnectionManager = CPXConnectionManager{}

// An interface definition for the [CPXConnectionManager] class.
//
// # Methods
//
//   - [ICPXConnectionManager.ConnectionForID]
//   - [ICPXConnectionManager.PidForConnection]
//   - [ICPXConnectionManager.InitWithSession]
//   - [ICPXConnectionManager.DebugDescription]
//   - [ICPXConnectionManager.Description]
//   - [ICPXConnectionManager.Hash]
//   - [ICPXConnectionManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager
type ICPXConnectionManager interface {
	objectivec.IObject

	// Topic: Methods

	ConnectionForID(id uint32) unsafe.Pointer
	PidForConnection(connection unsafe.Pointer) int
	InitWithSession(session uintptr) CPXConnectionManager
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXConnectionManager) Init() CPXConnectionManager {
	rv := objc.Send[CPXConnectionManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXConnectionManager) Autorelease() CPXConnectionManager {
	rv := objc.Send[CPXConnectionManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXConnectionManager creates a new CPXConnectionManager instance.
func NewCPXConnectionManager() CPXConnectionManager {
	class := getCPXConnectionManagerClass()
	rv := objc.Send[CPXConnectionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/initWithSession:
func NewCPXConnectionManagerWithSession(session uintptr) CPXConnectionManager {
	instance := getCPXConnectionManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return CPXConnectionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/connectionForID:
func (c CPXConnectionManager) ConnectionForID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("connectionForID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/pidForConnection:
func (c CPXConnectionManager) PidForConnection(connection unsafe.Pointer) int {
	rv := objc.Send[int](c.ID, objc.Sel("pidForConnection:"), connection)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/initWithSession:
func (c CPXConnectionManager) InitWithSession(session uintptr) CPXConnectionManager {
	rv := objc.Send[CPXConnectionManager](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/debugDescription
func (c CPXConnectionManager) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/description
func (c CPXConnectionManager) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/hash
func (c CPXConnectionManager) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManager/superclass
func (c CPXConnectionManager) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
