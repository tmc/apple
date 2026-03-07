// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLSharedEventHandle] class.
var (
	_MTLSharedEventHandleClass     MTLSharedEventHandleClass
	_MTLSharedEventHandleClassOnce sync.Once
)

func getMTLSharedEventHandleClass() MTLSharedEventHandleClass {
	_MTLSharedEventHandleClassOnce.Do(func() {
		_MTLSharedEventHandleClass = MTLSharedEventHandleClass{class: objc.GetClass("MTLSharedEventHandle")}
	})
	return _MTLSharedEventHandleClass
}

// GetMTLSharedEventHandleClass returns the class object for MTLSharedEventHandle.
func GetMTLSharedEventHandleClass() MTLSharedEventHandleClass {
	return getMTLSharedEventHandleClass()
}

type MTLSharedEventHandleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLSharedEventHandleClass) Alloc() MTLSharedEventHandle {
	rv := objc.Send[MTLSharedEventHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An instance you use to recreate a shareable event.
//
// # Overview
// 
// To create a [MTLSharedEventHandle] instance, call the
// [NewSharedEventHandle] method on an [MTLSharedEvent] instance. Use an XPC
// conection to pass a [MTLSharedEventHandle] instance to another process. To
// recreate the event, call the [NewSharedEventWithHandle] on an [MTLDevice]
// instance.
//
// # Identifying the shareable event handle
//
//   - [MTLSharedEventHandle.Label]: A string that identifies the shareable event.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventHandle
type MTLSharedEventHandle struct {
	objectivec.Object
}

// MTLSharedEventHandleFromID constructs a [MTLSharedEventHandle] from an objc.ID.
//
// An instance you use to recreate a shareable event.
func MTLSharedEventHandleFromID(id objc.ID) MTLSharedEventHandle {
	return MTLSharedEventHandle{objectivec.Object{id}}
}
// NOTE: MTLSharedEventHandle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLSharedEventHandle] class.
//
// # Identifying the shareable event handle
//
//   - [IMTLSharedEventHandle.Label]: A string that identifies the shareable event.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventHandle
type IMTLSharedEventHandle interface {
	objectivec.IObject

	// Topic: Identifying the shareable event handle

	// A string that identifies the shareable event.
	Label() string

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s MTLSharedEventHandle) Init() MTLSharedEventHandle {
	rv := objc.Send[MTLSharedEventHandle](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLSharedEventHandle) Autorelease() MTLSharedEventHandle {
	rv := objc.Send[MTLSharedEventHandle](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLSharedEventHandle creates a new MTLSharedEventHandle instance.
func NewMTLSharedEventHandle() MTLSharedEventHandle {
	class := getMTLSharedEventHandleClass()
	rv := objc.Send[MTLSharedEventHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (s MTLSharedEventHandle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A string that identifies the shareable event.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventHandle/label
func (s MTLSharedEventHandle) Label() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}



























