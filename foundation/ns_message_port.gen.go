// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MessagePort] class.
var (
	_MessagePortClass     MessagePortClass
	_MessagePortClassOnce sync.Once
)

func getMessagePortClass() MessagePortClass {
	_MessagePortClassOnce.Do(func() {
		_MessagePortClass = MessagePortClass{class: objc.GetClass("NSMessagePort")}
	})
	return _MessagePortClass
}

// GetMessagePortClass returns the class object for NSMessagePort.
func GetMessagePortClass() MessagePortClass {
	return getMessagePortClass()
}

type MessagePortClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MessagePortClass) Alloc() MessagePort {
	rv := objc.Send[MessagePort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A port that can be used as an endpoint for distributed object connections
// (or raw messaging).
//
// # Overview
// 
// [NSMessagePort] is a subclass of [NSPort] that allows for local (on the
// same machine) communication only. A companion class, [NSSocketPort], allows
// for both local and remote communication, but may be more expensive than
// [NSMessagePort] for the local case.
// 
// [NSMessagePort] defines no additional methods over those already defined by
// [NSPort].
//
// See: https://developer.apple.com/documentation/Foundation/MessagePort
type MessagePort struct {
	NSPort
}

// MessagePortFromID constructs a [MessagePort] from an objc.ID.
//
// A port that can be used as an endpoint for distributed object connections
// (or raw messaging).
func MessagePortFromID(id objc.ID) MessagePort {
	return NSMessagePort{NSPort: NSPortFromID(id)}
}

// NSMessagePortFromID is an alias for [MessagePortFromID] for cross-framework compatibility.
func NSMessagePortFromID(id objc.ID) MessagePort { return MessagePortFromID(id) }
// NOTE: MessagePort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MessagePort] class.
//
// See: https://developer.apple.com/documentation/Foundation/MessagePort
type IMessagePort interface {
	INSPort
}

// Init initializes the instance.
func (m MessagePort) Init() MessagePort {
	rv := objc.Send[MessagePort](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MessagePort) Autorelease() MessagePort {
	rv := objc.Send[MessagePort](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMessagePort creates a new MessagePort instance.
func NewMessagePort() MessagePort {
	class := getMessagePortClass()
	rv := objc.Send[MessagePort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMessagePortWithCoder(coder INSCoder) MessagePort {
	instance := getMessagePortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MessagePortFromID(rv)
}

