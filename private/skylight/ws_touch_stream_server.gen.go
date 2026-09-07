// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSTouchStreamServer] class.
var (
	_WSTouchStreamServerClass     WSTouchStreamServerClass
	_WSTouchStreamServerClassOnce sync.Once
)

func getWSTouchStreamServerClass() WSTouchStreamServerClass {
	_WSTouchStreamServerClassOnce.Do(func() {
		_WSTouchStreamServerClass = WSTouchStreamServerClass{class: objc.GetClass("WSTouchStreamServer")}
	})
	return _WSTouchStreamServerClass
}

// GetWSTouchStreamServerClass returns the class object for WSTouchStreamServer.
func GetWSTouchStreamServerClass() WSTouchStreamServerClass {
	return getWSTouchStreamServerClass()
}

type WSTouchStreamServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSTouchStreamServerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSTouchStreamServerClass) Alloc() WSTouchStreamServer {
	rv := objc.SendIfResponds[WSTouchStreamServer](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSTouchStreamServer._init]
//   - [WSTouchStreamServer.Activate]
//   - [WSTouchStreamServer.AppendDescriptionToStream]
//   - [WSTouchStreamServer.BkServer]
//   - [WSTouchStreamServer.DebugDescription]
//   - [WSTouchStreamServer.Description]
//   - [WSTouchStreamServer.Hash]
//   - [WSTouchStreamServer.Superclass]
type WSTouchStreamServer struct {
	objectivec.Object
}

// WSTouchStreamServerFromID constructs a [WSTouchStreamServer] from an objc.ID.
func WSTouchStreamServerFromID(id objc.ID) WSTouchStreamServer {
	return WSTouchStreamServer{objectivec.Object{ID: id}}
}

// Ensure WSTouchStreamServer implements IWSTouchStreamServer.
var _ IWSTouchStreamServer = WSTouchStreamServer{}

// An interface definition for the [WSTouchStreamServer] class.
//
// # Methods
//
//   - [IWSTouchStreamServer._init]
//   - [IWSTouchStreamServer.Activate]
//   - [IWSTouchStreamServer.AppendDescriptionToStream]
//   - [IWSTouchStreamServer.BkServer]
//   - [IWSTouchStreamServer.DebugDescription]
//   - [IWSTouchStreamServer.Description]
//   - [IWSTouchStreamServer.Hash]
//   - [IWSTouchStreamServer.Superclass]
type IWSTouchStreamServer interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	Activate()
	AppendDescriptionToStream(stream objectivec.IObject)
	BkServer() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSTouchStreamServer) Init() WSTouchStreamServer {
	rv := objc.SendIfResponds[WSTouchStreamServer](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSTouchStreamServer) Autorelease() WSTouchStreamServer {
	rv := objc.SendIfResponds[WSTouchStreamServer](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSTouchStreamServer creates a new WSTouchStreamServer instance.
func NewWSTouchStreamServer() WSTouchStreamServer {
	class := getWSTouchStreamServerClass()
	rv := objc.SendIfResponds[WSTouchStreamServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSTouchStreamServer) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (w WSTouchStreamServer) Activate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("activate"))
}
func (w WSTouchStreamServer) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

func (_WSTouchStreamServerClass WSTouchStreamServerClass) SharedInstance() WSTouchStreamServer {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSTouchStreamServerClass.class), objc.Sel("sharedInstance"))
	return WSTouchStreamServerFromID(rv)
}

func (w WSTouchStreamServer) BkServer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("bkServer"))
	return objectivec.Object{ID: rv}
}
func (w WSTouchStreamServer) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSTouchStreamServer) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSTouchStreamServer) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSTouchStreamServer) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
