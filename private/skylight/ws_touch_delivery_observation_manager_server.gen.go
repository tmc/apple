// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSTouchDeliveryObservationManagerServer] class.
var (
	_WSTouchDeliveryObservationManagerServerClass     WSTouchDeliveryObservationManagerServerClass
	_WSTouchDeliveryObservationManagerServerClassOnce sync.Once
)

func getWSTouchDeliveryObservationManagerServerClass() WSTouchDeliveryObservationManagerServerClass {
	_WSTouchDeliveryObservationManagerServerClassOnce.Do(func() {
		_WSTouchDeliveryObservationManagerServerClass = WSTouchDeliveryObservationManagerServerClass{class: objc.GetClass("WSTouchDeliveryObservationManagerServer")}
	})
	return _WSTouchDeliveryObservationManagerServerClass
}

// GetWSTouchDeliveryObservationManagerServerClass returns the class object for WSTouchDeliveryObservationManagerServer.
func GetWSTouchDeliveryObservationManagerServerClass() WSTouchDeliveryObservationManagerServerClass {
	return getWSTouchDeliveryObservationManagerServerClass()
}

type WSTouchDeliveryObservationManagerServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSTouchDeliveryObservationManagerServerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSTouchDeliveryObservationManagerServerClass) Alloc() WSTouchDeliveryObservationManagerServer {
	rv := objc.SendIfResponds[WSTouchDeliveryObservationManagerServer](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSTouchDeliveryObservationManagerServer._initWithServerClass]
//   - [WSTouchDeliveryObservationManagerServer.Activate]
//   - [WSTouchDeliveryObservationManagerServer.BkServer]
type WSTouchDeliveryObservationManagerServer struct {
	objectivec.Object
}

// WSTouchDeliveryObservationManagerServerFromID constructs a [WSTouchDeliveryObservationManagerServer] from an objc.ID.
func WSTouchDeliveryObservationManagerServerFromID(id objc.ID) WSTouchDeliveryObservationManagerServer {
	return WSTouchDeliveryObservationManagerServer{objectivec.Object{ID: id}}
}

// Ensure WSTouchDeliveryObservationManagerServer implements IWSTouchDeliveryObservationManagerServer.
var _ IWSTouchDeliveryObservationManagerServer = WSTouchDeliveryObservationManagerServer{}

// An interface definition for the [WSTouchDeliveryObservationManagerServer] class.
//
// # Methods
//
//   - [IWSTouchDeliveryObservationManagerServer._initWithServerClass]
//   - [IWSTouchDeliveryObservationManagerServer.Activate]
//   - [IWSTouchDeliveryObservationManagerServer.BkServer]
type IWSTouchDeliveryObservationManagerServer interface {
	objectivec.IObject

	// Topic: Methods

	_initWithServerClass(class objectivec.Class) objectivec.IObject
	Activate()
	BkServer() objectivec.IObject
}

// Init initializes the instance.
func (w WSTouchDeliveryObservationManagerServer) Init() WSTouchDeliveryObservationManagerServer {
	rv := objc.SendIfResponds[WSTouchDeliveryObservationManagerServer](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSTouchDeliveryObservationManagerServer) Autorelease() WSTouchDeliveryObservationManagerServer {
	rv := objc.SendIfResponds[WSTouchDeliveryObservationManagerServer](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSTouchDeliveryObservationManagerServer creates a new WSTouchDeliveryObservationManagerServer instance.
func NewWSTouchDeliveryObservationManagerServer() WSTouchDeliveryObservationManagerServer {
	class := getWSTouchDeliveryObservationManagerServerClass()
	rv := objc.SendIfResponds[WSTouchDeliveryObservationManagerServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSTouchDeliveryObservationManagerServer) _initWithServerClass(class objectivec.Class) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_initWithServerClass:"), class)
	return objectivec.Object{ID: rv}
}

// InitWithServerClass is an exported wrapper for the private method _initWithServerClass.
func (w WSTouchDeliveryObservationManagerServer) InitWithServerClass(class objectivec.Class) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_initWithServerClass:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithServerClass:"}
		return nil, err
	}
	return w._initWithServerClass(class), nil
}

// CanInitWithServerClass reports whether the receiver responds to the private selector _initWithServerClass:.
func (w WSTouchDeliveryObservationManagerServer) CanInitWithServerClass() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_initWithServerClass:"))
}
func (w WSTouchDeliveryObservationManagerServer) Activate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("activate"))
}

func (_WSTouchDeliveryObservationManagerServerClass WSTouchDeliveryObservationManagerServerClass) SharedInstance() WSTouchDeliveryObservationManagerServer {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSTouchDeliveryObservationManagerServerClass.class), objc.Sel("sharedInstance"))
	return WSTouchDeliveryObservationManagerServerFromID(rv)
}

func (w WSTouchDeliveryObservationManagerServer) BkServer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("bkServer"))
	return objectivec.Object{ID: rv}
}
