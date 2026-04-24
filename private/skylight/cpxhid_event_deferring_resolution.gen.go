// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXHIDEventDeferringResolution] class.
var (
	_CPXHIDEventDeferringResolutionClass     CPXHIDEventDeferringResolutionClass
	_CPXHIDEventDeferringResolutionClassOnce sync.Once
)

func getCPXHIDEventDeferringResolutionClass() CPXHIDEventDeferringResolutionClass {
	_CPXHIDEventDeferringResolutionClassOnce.Do(func() {
		_CPXHIDEventDeferringResolutionClass = CPXHIDEventDeferringResolutionClass{class: objc.GetClass("CPXHIDEventDeferringResolution")}
	})
	return _CPXHIDEventDeferringResolutionClass
}

// GetCPXHIDEventDeferringResolutionClass returns the class object for CPXHIDEventDeferringResolution.
func GetCPXHIDEventDeferringResolutionClass() CPXHIDEventDeferringResolutionClass {
	return getCPXHIDEventDeferringResolutionClass()
}

type CPXHIDEventDeferringResolutionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXHIDEventDeferringResolutionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXHIDEventDeferringResolutionClass) Alloc() CPXHIDEventDeferringResolution {
	rv := objc.Send[CPXHIDEventDeferringResolution](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXHIDEventDeferringResolution.ConnectionID]
//   - [CPXHIDEventDeferringResolution.Environment]
//   - [CPXHIDEventDeferringResolution.Pid]
//   - [CPXHIDEventDeferringResolution.ProcessRecord]
//   - [CPXHIDEventDeferringResolution.Token]
//   - [CPXHIDEventDeferringResolution.InitWithProcess]
//   - [CPXHIDEventDeferringResolution.InitWithProcessConnectionID]
//   - [CPXHIDEventDeferringResolution.DebugDescription]
//   - [CPXHIDEventDeferringResolution.Description]
//   - [CPXHIDEventDeferringResolution.Hash]
//   - [CPXHIDEventDeferringResolution.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution
type CPXHIDEventDeferringResolution struct {
	objectivec.Object
}

// CPXHIDEventDeferringResolutionFromID constructs a [CPXHIDEventDeferringResolution] from an objc.ID.
func CPXHIDEventDeferringResolutionFromID(id objc.ID) CPXHIDEventDeferringResolution {
	return CPXHIDEventDeferringResolution{objectivec.Object{ID: id}}
}

// Ensure CPXHIDEventDeferringResolution implements ICPXHIDEventDeferringResolution.
var _ ICPXHIDEventDeferringResolution = CPXHIDEventDeferringResolution{}

// An interface definition for the [CPXHIDEventDeferringResolution] class.
//
// # Methods
//
//   - [ICPXHIDEventDeferringResolution.ConnectionID]
//   - [ICPXHIDEventDeferringResolution.Environment]
//   - [ICPXHIDEventDeferringResolution.Pid]
//   - [ICPXHIDEventDeferringResolution.ProcessRecord]
//   - [ICPXHIDEventDeferringResolution.Token]
//   - [ICPXHIDEventDeferringResolution.InitWithProcess]
//   - [ICPXHIDEventDeferringResolution.InitWithProcessConnectionID]
//   - [ICPXHIDEventDeferringResolution.DebugDescription]
//   - [ICPXHIDEventDeferringResolution.Description]
//   - [ICPXHIDEventDeferringResolution.Hash]
//   - [ICPXHIDEventDeferringResolution.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution
type ICPXHIDEventDeferringResolution interface {
	objectivec.IObject

	// Topic: Methods

	ConnectionID() uint32
	Environment() unsafe.Pointer
	Pid() int
	ProcessRecord() *CPSProcessRecRef
	Token() unsafe.Pointer
	InitWithProcess(process *CPSProcessRecRef) CPXHIDEventDeferringResolution
	InitWithProcessConnectionID(process *CPSProcessRecRef, id uint32) CPXHIDEventDeferringResolution
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXHIDEventDeferringResolution) Init() CPXHIDEventDeferringResolution {
	rv := objc.Send[CPXHIDEventDeferringResolution](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXHIDEventDeferringResolution) Autorelease() CPXHIDEventDeferringResolution {
	rv := objc.Send[CPXHIDEventDeferringResolution](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXHIDEventDeferringResolution creates a new CPXHIDEventDeferringResolution instance.
func NewCPXHIDEventDeferringResolution() CPXHIDEventDeferringResolution {
	class := getCPXHIDEventDeferringResolutionClass()
	rv := objc.Send[CPXHIDEventDeferringResolution](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/initWithProcess:
func NewCPXHIDEventDeferringResolutionWithProcess(process *CPSProcessRecRef) CPXHIDEventDeferringResolution {
	instance := getCPXHIDEventDeferringResolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcess:"), process)
	return CPXHIDEventDeferringResolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/initWithProcess:connectionID:
func NewCPXHIDEventDeferringResolutionWithProcessConnectionID(process *CPSProcessRecRef, id uint32) CPXHIDEventDeferringResolution {
	instance := getCPXHIDEventDeferringResolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcess:connectionID:"), process, id)
	return CPXHIDEventDeferringResolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/initWithProcess:
func (c CPXHIDEventDeferringResolution) InitWithProcess(process *CPSProcessRecRef) CPXHIDEventDeferringResolution {
	rv := objc.Send[CPXHIDEventDeferringResolution](c.ID, objc.Sel("initWithProcess:"), process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/initWithProcess:connectionID:
func (c CPXHIDEventDeferringResolution) InitWithProcessConnectionID(process *CPSProcessRecRef, id uint32) CPXHIDEventDeferringResolution {
	rv := objc.Send[CPXHIDEventDeferringResolution](c.ID, objc.Sel("initWithProcess:connectionID:"), process, id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/connectionID
func (c CPXHIDEventDeferringResolution) ConnectionID() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("connectionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/debugDescription
func (c CPXHIDEventDeferringResolution) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/description
func (c CPXHIDEventDeferringResolution) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/environment
func (c CPXHIDEventDeferringResolution) Environment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("environment"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/hash
func (c CPXHIDEventDeferringResolution) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/pid
func (c CPXHIDEventDeferringResolution) Pid() int {
	rv := objc.Send[int](c.ID, objc.Sel("pid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/processRecord
func (c CPXHIDEventDeferringResolution) ProcessRecord() *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processRecord"))
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/superclass
func (c CPXHIDEventDeferringResolution) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolution/token
func (c CPXHIDEventDeferringResolution) Token() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("token"))
	return rv
}
