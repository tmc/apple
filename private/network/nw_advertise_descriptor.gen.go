// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWAdvertiseDescriptor] class.
var (
	_NWAdvertiseDescriptorClass     NWAdvertiseDescriptorClass
	_NWAdvertiseDescriptorClassOnce sync.Once
)

func getNWAdvertiseDescriptorClass() NWAdvertiseDescriptorClass {
	_NWAdvertiseDescriptorClassOnce.Do(func() {
		_NWAdvertiseDescriptorClass = NWAdvertiseDescriptorClass{class: objc.GetClass("NWAdvertiseDescriptor")}
	})
	return _NWAdvertiseDescriptorClass
}

// GetNWAdvertiseDescriptorClass returns the class object for NWAdvertiseDescriptor.
func GetNWAdvertiseDescriptorClass() NWAdvertiseDescriptorClass {
	return getNWAdvertiseDescriptorClass()
}

type NWAdvertiseDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWAdvertiseDescriptorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWAdvertiseDescriptorClass) Alloc() NWAdvertiseDescriptor {
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWAdvertiseDescriptor.BonjourServiceDomain]
//   - [NWAdvertiseDescriptor.BonjourServiceName]
//   - [NWAdvertiseDescriptor.BonjourServiceType]
//   - [NWAdvertiseDescriptor.DescriptionWithIndentShowFullContent]
//   - [NWAdvertiseDescriptor.InternalDescriptor]
//   - [NWAdvertiseDescriptor.SetInternalDescriptor]
//   - [NWAdvertiseDescriptor.PrivateDescription]
//   - [NWAdvertiseDescriptor.TxtRecord]
//   - [NWAdvertiseDescriptor.SetTxtRecord]
//   - [NWAdvertiseDescriptor.InitWithDescriptor]
//   - [NWAdvertiseDescriptor.InitWithNameTypeDomain]
type NWAdvertiseDescriptor struct {
	objectivec.Object
}

// NWAdvertiseDescriptorFromID constructs a [NWAdvertiseDescriptor] from an objc.ID.
func NWAdvertiseDescriptorFromID(id objc.ID) NWAdvertiseDescriptor {
	return NWAdvertiseDescriptor{objectivec.Object{ID: id}}
}

// Ensure NWAdvertiseDescriptor implements INWAdvertiseDescriptor.
var _ INWAdvertiseDescriptor = NWAdvertiseDescriptor{}

// An interface definition for the [NWAdvertiseDescriptor] class.
//
// # Methods
//
//   - [INWAdvertiseDescriptor.BonjourServiceDomain]
//   - [INWAdvertiseDescriptor.BonjourServiceName]
//   - [INWAdvertiseDescriptor.BonjourServiceType]
//   - [INWAdvertiseDescriptor.DescriptionWithIndentShowFullContent]
//   - [INWAdvertiseDescriptor.InternalDescriptor]
//   - [INWAdvertiseDescriptor.SetInternalDescriptor]
//   - [INWAdvertiseDescriptor.PrivateDescription]
//   - [INWAdvertiseDescriptor.TxtRecord]
//   - [INWAdvertiseDescriptor.SetTxtRecord]
//   - [INWAdvertiseDescriptor.InitWithDescriptor]
//   - [INWAdvertiseDescriptor.InitWithNameTypeDomain]
type INWAdvertiseDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	BonjourServiceDomain() string
	BonjourServiceName() string
	BonjourServiceType() string
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	InternalDescriptor() objectivec.Object
	SetInternalDescriptor(value objectivec.Object)
	PrivateDescription() objectivec.IObject
	TxtRecord() foundation.NSData
	SetTxtRecord(value foundation.NSData)
	InitWithDescriptor(descriptor objectivec.IObject) NWAdvertiseDescriptor
	InitWithNameTypeDomain(name objectivec.IObject, type_ objectivec.IObject, domain objectivec.IObject) NWAdvertiseDescriptor
}

// Init initializes the instance.
func (n NWAdvertiseDescriptor) Init() NWAdvertiseDescriptor {
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWAdvertiseDescriptor) Autorelease() NWAdvertiseDescriptor {
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWAdvertiseDescriptor creates a new NWAdvertiseDescriptor instance.
func NewNWAdvertiseDescriptor() NWAdvertiseDescriptor {
	class := getNWAdvertiseDescriptorClass()
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWAdvertiseDescriptorWithDescriptor(descriptor objectivec.IObject) NWAdvertiseDescriptor {
	instance := getNWAdvertiseDescriptorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescriptor:"), descriptor)
	return NWAdvertiseDescriptorFromID(rv)
}

func NewNWAdvertiseDescriptorWithNameTypeDomain(name objectivec.IObject, type_ objectivec.IObject, domain objectivec.IObject) NWAdvertiseDescriptor {
	instance := getNWAdvertiseDescriptorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:type:domain:"), name, type_, domain)
	return NWAdvertiseDescriptorFromID(rv)
}

func (n NWAdvertiseDescriptor) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}
func (n NWAdvertiseDescriptor) PrivateDescription() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("privateDescription"))
	return objectivec.Object{ID: rv}
}
func (n NWAdvertiseDescriptor) InitWithDescriptor(descriptor objectivec.IObject) NWAdvertiseDescriptor {
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](n.ID, objc.Sel("initWithDescriptor:"), descriptor)
	return rv
}
func (n NWAdvertiseDescriptor) InitWithNameTypeDomain(name objectivec.IObject, type_ objectivec.IObject, domain objectivec.IObject) NWAdvertiseDescriptor {
	rv := objc.SendIfResponds[NWAdvertiseDescriptor](n.ID, objc.Sel("initWithName:type:domain:"), name, type_, domain)
	return rv
}

func (n NWAdvertiseDescriptor) BonjourServiceDomain() string {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("bonjourServiceDomain"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWAdvertiseDescriptor) BonjourServiceName() string {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("bonjourServiceName"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWAdvertiseDescriptor) BonjourServiceType() string {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("bonjourServiceType"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWAdvertiseDescriptor) InternalDescriptor() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("internalDescriptor"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWAdvertiseDescriptor) SetInternalDescriptor(value objectivec.Object) {
	objc.SendIfResponds[struct{}](n.ID, objc.Sel("setInternalDescriptor:"), value)
}
func (n NWAdvertiseDescriptor) TxtRecord() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](n.ID, objc.Sel("txtRecord"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n NWAdvertiseDescriptor) SetTxtRecord(value foundation.NSData) {
	objc.SendIfResponds[struct{}](n.ID, objc.Sel("setTxtRecord:"), value)
}
