// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSTouchEventServer] class.
var (
	_WSTouchEventServerClass     WSTouchEventServerClass
	_WSTouchEventServerClassOnce sync.Once
)

func getWSTouchEventServerClass() WSTouchEventServerClass {
	_WSTouchEventServerClassOnce.Do(func() {
		_WSTouchEventServerClass = WSTouchEventServerClass{class: objc.GetClass("WSTouchEventServer")}
	})
	return _WSTouchEventServerClass
}

// GetWSTouchEventServerClass returns the class object for WSTouchEventServer.
func GetWSTouchEventServerClass() WSTouchEventServerClass {
	return getWSTouchEventServerClass()
}

type WSTouchEventServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSTouchEventServerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSTouchEventServerClass) Alloc() WSTouchEventServer {
	rv := objc.SendIfResponds[WSTouchEventServer](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSTouchEventServer._initWithServerClass]
//   - [WSTouchEventServer.Activate]
//   - [WSTouchEventServer.BkServer]
type WSTouchEventServer struct {
	objectivec.Object
}

// WSTouchEventServerFromID constructs a [WSTouchEventServer] from an objc.ID.
func WSTouchEventServerFromID(id objc.ID) WSTouchEventServer {
	return WSTouchEventServer{objectivec.Object{ID: id}}
}

// Ensure WSTouchEventServer implements IWSTouchEventServer.
var _ IWSTouchEventServer = WSTouchEventServer{}

// An interface definition for the [WSTouchEventServer] class.
//
// # Methods
//
//   - [IWSTouchEventServer._initWithServerClass]
//   - [IWSTouchEventServer.Activate]
//   - [IWSTouchEventServer.BkServer]
type IWSTouchEventServer interface {
	objectivec.IObject

	// Topic: Methods

	_initWithServerClass(class objectivec.Class) objectivec.IObject
	Activate()
	BkServer() objectivec.IObject
}

// Init initializes the instance.
func (w WSTouchEventServer) Init() WSTouchEventServer {
	rv := objc.SendIfResponds[WSTouchEventServer](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSTouchEventServer) Autorelease() WSTouchEventServer {
	rv := objc.SendIfResponds[WSTouchEventServer](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSTouchEventServer creates a new WSTouchEventServer instance.
func NewWSTouchEventServer() WSTouchEventServer {
	class := getWSTouchEventServerClass()
	rv := objc.SendIfResponds[WSTouchEventServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSTouchEventServer) _initWithServerClass(class objectivec.Class) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_initWithServerClass:"), class)
	return objectivec.Object{ID: rv}
}

// InitWithServerClass is an exported wrapper for the private method _initWithServerClass.
func (w WSTouchEventServer) InitWithServerClass(class objectivec.Class) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_initWithServerClass:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithServerClass:"}
		return nil, err
	}
	return w._initWithServerClass(class), nil
}

// CanInitWithServerClass reports whether the receiver responds to the private selector _initWithServerClass:.
func (w WSTouchEventServer) CanInitWithServerClass() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_initWithServerClass:"))
}
func (w WSTouchEventServer) Activate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("activate"))
}

func (_WSTouchEventServerClass WSTouchEventServerClass) SharedInstance() WSTouchEventServer {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSTouchEventServerClass.class), objc.Sel("sharedInstance"))
	return WSTouchEventServerFromID(rv)
}

func (w WSTouchEventServer) BkServer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("bkServer"))
	return objectivec.Object{ID: rv}
}
