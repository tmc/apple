// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWEndpoint] class.
var (
	_NWEndpointClass     NWEndpointClass
	_NWEndpointClassOnce sync.Once
)

func getNWEndpointClass() NWEndpointClass {
	_NWEndpointClassOnce.Do(func() {
		_NWEndpointClass = NWEndpointClass{class: objc.GetClass("NWEndpoint")}
	})
	return _NWEndpointClass
}

// GetNWEndpointClass returns the class object for NWEndpoint.
func GetNWEndpointClass() NWEndpointClass {
	return getNWEndpointClass()
}

type NWEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWEndpointClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWEndpointClass) Alloc() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWEndpoint.AlternatePort]
//   - [NWEndpoint.SetAlternatePort]
//   - [NWEndpoint.CopyCEndpoint]
//   - [NWEndpoint.CreateProtocolBufferObject]
//   - [NWEndpoint.DescriptionWithIndentShowFullContent]
//   - [NWEndpoint.EncodeWithCoder]
//   - [NWEndpoint.EncodedData]
//   - [NWEndpoint.Interface]
//   - [NWEndpoint.SetInterface]
//   - [NWEndpoint.InternalEndpoint]
//   - [NWEndpoint.SetInternalEndpoint]
//   - [NWEndpoint.ParentEndpointDomain]
//   - [NWEndpoint.PrivateDescription]
//   - [NWEndpoint.RemoteInterfaceType]
//   - [NWEndpoint.SetRemoteInterfaceType]
//   - [NWEndpoint.TxtRecord]
//   - [NWEndpoint.SetTxtRecord]
//   - [NWEndpoint.InitWithCoder]
//   - [NWEndpoint.InitWithEndpoint]
type NWEndpoint struct {
	objectivec.Object
}

// NWEndpointFromID constructs a [NWEndpoint] from an objc.ID.
func NWEndpointFromID(id objc.ID) NWEndpoint {
	return NWEndpoint{objectivec.Object{ID: id}}
}

// Ensure NWEndpoint implements INWEndpoint.
var _ INWEndpoint = NWEndpoint{}

// An interface definition for the [NWEndpoint] class.
//
// # Methods
//
//   - [INWEndpoint.AlternatePort]
//   - [INWEndpoint.SetAlternatePort]
//   - [INWEndpoint.CopyCEndpoint]
//   - [INWEndpoint.CreateProtocolBufferObject]
//   - [INWEndpoint.DescriptionWithIndentShowFullContent]
//   - [INWEndpoint.EncodeWithCoder]
//   - [INWEndpoint.EncodedData]
//   - [INWEndpoint.Interface]
//   - [INWEndpoint.SetInterface]
//   - [INWEndpoint.InternalEndpoint]
//   - [INWEndpoint.SetInternalEndpoint]
//   - [INWEndpoint.ParentEndpointDomain]
//   - [INWEndpoint.PrivateDescription]
//   - [INWEndpoint.RemoteInterfaceType]
//   - [INWEndpoint.SetRemoteInterfaceType]
//   - [INWEndpoint.TxtRecord]
//   - [INWEndpoint.SetTxtRecord]
//   - [INWEndpoint.InitWithCoder]
//   - [INWEndpoint.InitWithEndpoint]
type INWEndpoint interface {
	objectivec.IObject

	// Topic: Methods

	AlternatePort() uint16
	SetAlternatePort(value uint16)
	CopyCEndpoint() objectivec.IObject
	CreateProtocolBufferObject() objectivec.IObject
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	EncodedData() objectivec.IObject
	Interface() INWInterface
	SetInterface(value INWInterface)
	InternalEndpoint() objectivec.Object
	SetInternalEndpoint(value objectivec.Object)
	ParentEndpointDomain() string
	PrivateDescription() string
	RemoteInterfaceType() int64
	SetRemoteInterfaceType(value int64)
	TxtRecord() foundation.NSData
	SetTxtRecord(value foundation.NSData)
	InitWithCoder(coder foundation.INSCoder) NWEndpoint
	InitWithEndpoint(endpoint objectivec.IObject) NWEndpoint
}

// Init initializes the instance.
func (n NWEndpoint) Init() NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWEndpoint) Autorelease() NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWEndpoint creates a new NWEndpoint instance.
func NewNWEndpoint() NWEndpoint {
	class := getNWEndpointClass()
	rv := objc.Send[NWEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWEndpointWithCoder(coder objectivec.IObject) NWEndpoint {
	instance := getNWEndpointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NWEndpointFromID(rv)
}

func NewNWEndpointWithEndpoint(endpoint objectivec.IObject) NWEndpoint {
	instance := getNWEndpointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEndpoint:"), endpoint)
	return NWEndpointFromID(rv)
}

func (n NWEndpoint) CopyCEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyCEndpoint"))
	return objectivec.Object{ID: rv}
}
func (n NWEndpoint) CreateProtocolBufferObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createProtocolBufferObject"))
	return objectivec.Object{ID: rv}
}
func (n NWEndpoint) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}
func (n NWEndpoint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (n NWEndpoint) EncodedData() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodedData"))
	return objectivec.Object{ID: rv}
}
func (n NWEndpoint) InitWithCoder(coder foundation.INSCoder) NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (n NWEndpoint) InitWithEndpoint(endpoint objectivec.IObject) NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("initWithEndpoint:"), endpoint)
	return rv
}

func (_NWEndpointClass NWEndpointClass) CopyClassForEndpointType(type_ int) objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_NWEndpointClass.class), objc.Sel("copyClassForEndpointType:"), type_)
	return objectivec.Class(rv)
}
func (_NWEndpointClass NWEndpointClass) EndpointType() uint32 {
	rv := objc.Send[uint32](objc.ID(_NWEndpointClass.class), objc.Sel("endpointType"))
	return rv
}
func (_NWEndpointClass NWEndpointClass) EndpointWithCEndpoint(cEndpoint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWEndpointClass.class), objc.Sel("endpointWithCEndpoint:"), cEndpoint)
	return objectivec.Object{ID: rv}
}
func (_NWEndpointClass NWEndpointClass) EndpointWithInternalEndpoint(endpoint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWEndpointClass.class), objc.Sel("endpointWithInternalEndpoint:"), endpoint)
	return objectivec.Object{ID: rv}
}
func (_NWEndpointClass NWEndpointClass) EndpointWithProtocolBufferData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWEndpointClass.class), objc.Sel("endpointWithProtocolBufferData:"), data)
	return objectivec.Object{ID: rv}
}
func (_NWEndpointClass NWEndpointClass) SupportsResolverCallback() bool {
	rv := objc.Send[bool](objc.ID(_NWEndpointClass.class), objc.Sel("supportsResolverCallback"))
	return rv
}
func (_NWEndpointClass NWEndpointClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_NWEndpointClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (n NWEndpoint) AlternatePort() uint16 {
	rv := objc.Send[uint16](n.ID, objc.Sel("alternatePort"))
	return rv
}
func (n NWEndpoint) SetAlternatePort(value uint16) {
	objc.Send[struct{}](n.ID, objc.Sel("setAlternatePort:"), value)
}
func (n NWEndpoint) Interface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("interface"))
	return NWInterfaceFromID(objc.ID(rv))
}
func (n NWEndpoint) SetInterface(value INWInterface) {
	objc.Send[struct{}](n.ID, objc.Sel("setInterface:"), value)
}
func (n NWEndpoint) InternalEndpoint() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalEndpoint"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWEndpoint) SetInternalEndpoint(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalEndpoint:"), value)
}
func (n NWEndpoint) ParentEndpointDomain() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parentEndpointDomain"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWEndpoint) PrivateDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("privateDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWEndpoint) RemoteInterfaceType() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("remoteInterfaceType"))
	return rv
}
func (n NWEndpoint) SetRemoteInterfaceType(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setRemoteInterfaceType:"), value)
}
func (n NWEndpoint) TxtRecord() foundation.NSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("txtRecord"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n NWEndpoint) SetTxtRecord(value foundation.NSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setTxtRecord:"), value)
}
