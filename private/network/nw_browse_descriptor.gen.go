// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWBrowseDescriptor] class.
var (
	_NWBrowseDescriptorClass     NWBrowseDescriptorClass
	_NWBrowseDescriptorClassOnce sync.Once
)

func getNWBrowseDescriptorClass() NWBrowseDescriptorClass {
	_NWBrowseDescriptorClassOnce.Do(func() {
		_NWBrowseDescriptorClass = NWBrowseDescriptorClass{class: objc.GetClass("NWBrowseDescriptor")}
	})
	return _NWBrowseDescriptorClass
}

// GetNWBrowseDescriptorClass returns the class object for NWBrowseDescriptor.
func GetNWBrowseDescriptorClass() NWBrowseDescriptorClass {
	return getNWBrowseDescriptorClass()
}

type NWBrowseDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWBrowseDescriptorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWBrowseDescriptorClass) Alloc() NWBrowseDescriptor {
	rv := objc.Send[NWBrowseDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWBrowseDescriptor.BonjourServiceDomain]
//   - [NWBrowseDescriptor.BonjourServiceType]
//   - [NWBrowseDescriptor.BrowseWithCompletionHandler]
//   - [NWBrowseDescriptor.CreateProtocolBufferObject]
//   - [NWBrowseDescriptor.DescriptionWithIndentShowFullContent]
//   - [NWBrowseDescriptor.EncodedData]
//   - [NWBrowseDescriptor.InternalDescriptor]
//   - [NWBrowseDescriptor.SetInternalDescriptor]
//   - [NWBrowseDescriptor.PrivateDescription]
//   - [NWBrowseDescriptor.InitWithDescriptor]
//   - [NWBrowseDescriptor.InitWithEncodedData]
type NWBrowseDescriptor struct {
	objectivec.Object
}

// NWBrowseDescriptorFromID constructs a [NWBrowseDescriptor] from an objc.ID.
func NWBrowseDescriptorFromID(id objc.ID) NWBrowseDescriptor {
	return NWBrowseDescriptor{objectivec.Object{ID: id}}
}

// Ensure NWBrowseDescriptor implements INWBrowseDescriptor.
var _ INWBrowseDescriptor = NWBrowseDescriptor{}

// An interface definition for the [NWBrowseDescriptor] class.
//
// # Methods
//
//   - [INWBrowseDescriptor.BonjourServiceDomain]
//   - [INWBrowseDescriptor.BonjourServiceType]
//   - [INWBrowseDescriptor.BrowseWithCompletionHandler]
//   - [INWBrowseDescriptor.CreateProtocolBufferObject]
//   - [INWBrowseDescriptor.DescriptionWithIndentShowFullContent]
//   - [INWBrowseDescriptor.EncodedData]
//   - [INWBrowseDescriptor.InternalDescriptor]
//   - [INWBrowseDescriptor.SetInternalDescriptor]
//   - [INWBrowseDescriptor.PrivateDescription]
//   - [INWBrowseDescriptor.InitWithDescriptor]
//   - [INWBrowseDescriptor.InitWithEncodedData]
type INWBrowseDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	BonjourServiceDomain() string
	BonjourServiceType() string
	BrowseWithCompletionHandler(handler ErrorHandler)
	CreateProtocolBufferObject() objectivec.IObject
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	EncodedData() objectivec.IObject
	InternalDescriptor() objectivec.Object
	SetInternalDescriptor(value objectivec.Object)
	PrivateDescription() string
	InitWithDescriptor(descriptor objectivec.IObject) NWBrowseDescriptor
	InitWithEncodedData(data objectivec.IObject) NWBrowseDescriptor
}

// Init initializes the instance.
func (n NWBrowseDescriptor) Init() NWBrowseDescriptor {
	rv := objc.Send[NWBrowseDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWBrowseDescriptor) Autorelease() NWBrowseDescriptor {
	rv := objc.Send[NWBrowseDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWBrowseDescriptor creates a new NWBrowseDescriptor instance.
func NewNWBrowseDescriptor() NWBrowseDescriptor {
	class := getNWBrowseDescriptorClass()
	rv := objc.Send[NWBrowseDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWBrowseDescriptorWithDescriptor(descriptor objectivec.IObject) NWBrowseDescriptor {
	instance := getNWBrowseDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescriptor:"), descriptor)
	return NWBrowseDescriptorFromID(rv)
}

func NewNWBrowseDescriptorWithEncodedData(data objectivec.IObject) NWBrowseDescriptor {
	instance := getNWBrowseDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEncodedData:"), data)
	return NWBrowseDescriptorFromID(rv)
}

func (n NWBrowseDescriptor) BrowseWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](n.ID, objc.Sel("browseWithCompletionHandler:"), _block0)
}
func (n NWBrowseDescriptor) CreateProtocolBufferObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createProtocolBufferObject"))
	return objectivec.Object{ID: rv}
}
func (n NWBrowseDescriptor) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}
func (n NWBrowseDescriptor) EncodedData() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodedData"))
	return objectivec.Object{ID: rv}
}
func (n NWBrowseDescriptor) InitWithDescriptor(descriptor objectivec.IObject) NWBrowseDescriptor {
	rv := objc.Send[NWBrowseDescriptor](n.ID, objc.Sel("initWithDescriptor:"), descriptor)
	return rv
}
func (n NWBrowseDescriptor) InitWithEncodedData(data objectivec.IObject) NWBrowseDescriptor {
	rv := objc.Send[NWBrowseDescriptor](n.ID, objc.Sel("initWithEncodedData:"), data)
	return rv
}

func (_NWBrowseDescriptorClass NWBrowseDescriptorClass) CopyClassForDescriptorType(type_ int) objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_NWBrowseDescriptorClass.class), objc.Sel("copyClassForDescriptorType:"), type_)
	return objectivec.Class(rv)
}
func (_NWBrowseDescriptorClass NWBrowseDescriptorClass) DescriptorType() uint32 {
	rv := objc.Send[uint32](objc.ID(_NWBrowseDescriptorClass.class), objc.Sel("descriptorType"))
	return rv
}
func (_NWBrowseDescriptorClass NWBrowseDescriptorClass) DescriptorWithInternalDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWBrowseDescriptorClass.class), objc.Sel("descriptorWithInternalDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}
func (_NWBrowseDescriptorClass NWBrowseDescriptorClass) DescriptorWithProtocolBufferData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWBrowseDescriptorClass.class), objc.Sel("descriptorWithProtocolBufferData:"), data)
	return objectivec.Object{ID: rv}
}
func (_NWBrowseDescriptorClass NWBrowseDescriptorClass) SupportsBrowseCallback() bool {
	rv := objc.Send[bool](objc.ID(_NWBrowseDescriptorClass.class), objc.Sel("supportsBrowseCallback"))
	return rv
}

func (n NWBrowseDescriptor) BonjourServiceDomain() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("bonjourServiceDomain"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWBrowseDescriptor) BonjourServiceType() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("bonjourServiceType"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWBrowseDescriptor) InternalDescriptor() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalDescriptor"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWBrowseDescriptor) SetInternalDescriptor(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalDescriptor:"), value)
}
func (n NWBrowseDescriptor) PrivateDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("privateDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Browse is a synchronous wrapper around [NWBrowseDescriptor.BrowseWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (n NWBrowseDescriptor) Browse(ctx context.Context) error {
	done := make(chan error, 1)
	n.BrowseWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
