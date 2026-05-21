// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWBrowser] class.
var (
	_NWBrowserClass     NWBrowserClass
	_NWBrowserClassOnce sync.Once
)

func getNWBrowserClass() NWBrowserClass {
	_NWBrowserClassOnce.Do(func() {
		_NWBrowserClass = NWBrowserClass{class: objc.GetClass("NWBrowser")}
	})
	return _NWBrowserClass
}

// GetNWBrowserClass returns the class object for NWBrowser.
func GetNWBrowserClass() NWBrowserClass {
	return getNWBrowserClass()
}

type NWBrowserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWBrowserClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWBrowserClass) Alloc() NWBrowser {
	rv := objc.Send[NWBrowser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWBrowser.Cancel]
//   - [NWBrowser.CopyDiscoveredEndpoints]
//   - [NWBrowser.Descriptor]
//   - [NWBrowser.DiscoveredEndpoints]
//   - [NWBrowser.InternalBrowser]
//   - [NWBrowser.InternalDiscoveredEndpoints]
//   - [NWBrowser.SetInternalDiscoveredEndpoints]
//   - [NWBrowser.Parameters]
//   - [NWBrowser.SetUpdateHandler]
//   - [NWBrowser.InitWithDescriptorParameters]
type NWBrowser struct {
	objectivec.Object
}

// NWBrowserFromID constructs a [NWBrowser] from an objc.ID.
func NWBrowserFromID(id objc.ID) NWBrowser {
	return NWBrowser{objectivec.Object{ID: id}}
}

// Ensure NWBrowser implements INWBrowser.
var _ INWBrowser = NWBrowser{}

// An interface definition for the [NWBrowser] class.
//
// # Methods
//
//   - [INWBrowser.Cancel]
//   - [INWBrowser.CopyDiscoveredEndpoints]
//   - [INWBrowser.Descriptor]
//   - [INWBrowser.DiscoveredEndpoints]
//   - [INWBrowser.InternalBrowser]
//   - [INWBrowser.InternalDiscoveredEndpoints]
//   - [INWBrowser.SetInternalDiscoveredEndpoints]
//   - [INWBrowser.Parameters]
//   - [INWBrowser.SetUpdateHandler]
//   - [INWBrowser.InitWithDescriptorParameters]
type INWBrowser interface {
	objectivec.IObject

	// Topic: Methods

	Cancel()
	CopyDiscoveredEndpoints() objectivec.IObject
	Descriptor() INWBrowseDescriptor
	DiscoveredEndpoints() foundation.INSSet
	InternalBrowser() objectivec.Object
	InternalDiscoveredEndpoints() foundation.INSSet
	SetInternalDiscoveredEndpoints(value foundation.INSSet)
	Parameters() INWParameters
	SetUpdateHandler()
	InitWithDescriptorParameters(descriptor objectivec.IObject, parameters objectivec.IObject) NWBrowser
}

// Init initializes the instance.
func (n NWBrowser) Init() NWBrowser {
	rv := objc.Send[NWBrowser](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWBrowser) Autorelease() NWBrowser {
	rv := objc.Send[NWBrowser](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWBrowser creates a new NWBrowser instance.
func NewNWBrowser() NWBrowser {
	class := getNWBrowserClass()
	rv := objc.Send[NWBrowser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWBrowserWithDescriptorParameters(descriptor objectivec.IObject, parameters objectivec.IObject) NWBrowser {
	instance := getNWBrowserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescriptor:parameters:"), descriptor, parameters)
	return NWBrowserFromID(rv)
}

func (n NWBrowser) Cancel() {
	objc.Send[objc.ID](n.ID, objc.Sel("cancel"))
}
func (n NWBrowser) CopyDiscoveredEndpoints() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyDiscoveredEndpoints"))
	return objectivec.Object{ID: rv}
}
func (n NWBrowser) SetUpdateHandler() {
	objc.Send[objc.ID](n.ID, objc.Sel("setUpdateHandler"))
}
func (n NWBrowser) InitWithDescriptorParameters(descriptor objectivec.IObject, parameters objectivec.IObject) NWBrowser {
	rv := objc.Send[NWBrowser](n.ID, objc.Sel("initWithDescriptor:parameters:"), descriptor, parameters)
	return rv
}

func (_NWBrowserClass NWBrowserClass) AutomaticallyNotifiesObserversForKey(key objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_NWBrowserClass.class), objc.Sel("automaticallyNotifiesObserversForKey:"), key)
	return rv
}

func (n NWBrowser) Descriptor() INWBrowseDescriptor {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptor"))
	return NWBrowseDescriptorFromID(objc.ID(rv))
}
func (n NWBrowser) DiscoveredEndpoints() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("discoveredEndpoints"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n NWBrowser) InternalBrowser() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalBrowser"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWBrowser) InternalDiscoveredEndpoints() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalDiscoveredEndpoints"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n NWBrowser) SetInternalDiscoveredEndpoints(value foundation.INSSet) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalDiscoveredEndpoints:"), value)
}
func (n NWBrowser) Parameters() INWParameters {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameters"))
	return NWParametersFromID(objc.ID(rv))
}
