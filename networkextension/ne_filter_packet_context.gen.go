// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterPacketContext] class.
var (
	_NEFilterPacketContextClass     NEFilterPacketContextClass
	_NEFilterPacketContextClassOnce sync.Once
)

func getNEFilterPacketContextClass() NEFilterPacketContextClass {
	_NEFilterPacketContextClassOnce.Do(func() {
		_NEFilterPacketContextClass = NEFilterPacketContextClass{class: objc.GetClass("NEFilterPacketContext")}
	})
	return _NEFilterPacketContextClass
}

// GetNEFilterPacketContextClass returns the class object for NEFilterPacketContext.
func GetNEFilterPacketContextClass() NEFilterPacketContextClass {
	return getNEFilterPacketContextClass()
}

type NEFilterPacketContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterPacketContextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterPacketContextClass) Alloc() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The context object provided to the filter packet handler.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketContext
type NEFilterPacketContext struct {
	objectivec.Object
}

// NEFilterPacketContextFromID constructs a [NEFilterPacketContext] from an objc.ID.
//
// The context object provided to the filter packet handler.
func NEFilterPacketContextFromID(id objc.ID) NEFilterPacketContext {
	return NEFilterPacketContext{objectivec.Object{ID: id}}
}
// NOTE: NEFilterPacketContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterPacketContext] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketContext
type INEFilterPacketContext interface {
	objectivec.IObject

	// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
	PacketHandler() NEFilterPacketHandler
	SetPacketHandler(value NEFilterPacketHandler)
}

// Init initializes the instance.
func (f NEFilterPacketContext) Init() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterPacketContext) Autorelease() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketContext creates a new NEFilterPacketContext instance.
func NewNEFilterPacketContext() NEFilterPacketContext {
	class := getNEFilterPacketContextClass()
	rv := objc.Send[NEFilterPacketContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Swift closure or an ObjectiveC block that handles each packet received by
// the filter.
//
// See: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/packethandler
func (f NEFilterPacketContext) PacketHandler() NEFilterPacketHandler {
	rv := objc.Send[NEFilterPacketHandler](f.ID, objc.Sel("packetHandler"))
	return NEFilterPacketHandler(rv)
}
func (f NEFilterPacketContext) SetPacketHandler(value NEFilterPacketHandler) {
	objc.Send[struct{}](f.ID, objc.Sel("setPacketHandler:"), value)
}

