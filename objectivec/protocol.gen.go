// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [Protocol] class.
var (
	_ProtocolClass     ProtocolClass
	_ProtocolClassOnce sync.Once
)

func getProtocolClass() ProtocolClass {
	_ProtocolClassOnce.Do(func() {
		_ProtocolClass = ProtocolClass{class: objc.GetClass("Protocol")}
	})
	return _ProtocolClass
}

// GetProtocolClass returns the class object for Protocol.
func GetProtocolClass() ProtocolClass {
	return getProtocolClass()
}

type ProtocolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc ProtocolClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc ProtocolClass) Alloc() Protocol {
	rv := objc.Send[Protocol](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/ObjectiveC/Protocol
type Protocol struct {
	Object
}

// ProtocolFromID constructs a [Protocol] from an objc.ID.
func ProtocolFromID(id objc.ID) Protocol {
	return Protocol{Object{ID: id}}
}

// Ensure Protocol implements IProtocol.
var _ IProtocol = Protocol{}

// An interface definition for the [Protocol] class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/Protocol
type IProtocol interface {
	IObject
}

// Init initializes the instance.
func (p Protocol) Init() Protocol {
	rv := objc.Send[Protocol](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Protocol) Autorelease() Protocol {
	rv := objc.Send[Protocol](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewProtocol creates a new Protocol instance.
func NewProtocol() Protocol {
	class := getProtocolClass()
	rv := objc.Send[Protocol](objc.ID(class.class), objc.Sel("new"))
	return rv
}
