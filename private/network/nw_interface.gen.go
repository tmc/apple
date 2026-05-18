// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWInterface] class.
var (
	_NWInterfaceClass     NWInterfaceClass
	_NWInterfaceClassOnce sync.Once
)

func getNWInterfaceClass() NWInterfaceClass {
	_NWInterfaceClassOnce.Do(func() {
		_NWInterfaceClass = NWInterfaceClass{class: objc.GetClass("NWInterface")}
	})
	return _NWInterfaceClass
}

// GetNWInterfaceClass returns the class object for NWInterface.
func GetNWInterfaceClass() NWInterfaceClass {
	return getNWInterfaceClass()
}

type NWInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWInterfaceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWInterfaceClass) Alloc() NWInterface {
	rv := objc.Send[NWInterface](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWInterface.CInterface]
//   - [NWInterface.CopyLocalAddressForDefaultIPv4]
//   - [NWInterface.CopyLocalAddressForDefaultIPv6]
//   - [NWInterface.CopyLocalAddressForRemoteAddress]
//   - [NWInterface.CreateProtocolBufferObject]
//   - [NWInterface.DelegateInterface]
//   - [NWInterface.DescriptionWithIndentShowFullContent]
//   - [NWInterface.EncodeWithCoder]
//   - [NWInterface.Generation]
//   - [NWInterface.HasDNS]
//   - [NWInterface.HasNAT64]
//   - [NWInterface.InterfaceIndex]
//   - [NWInterface.InterfaceName]
//   - [NWInterface.InternalInterface]
//   - [NWInterface.SetInternalInterface]
//   - [NWInterface.Ipv4Broadcast]
//   - [NWInterface.Ipv4Netmask]
//   - [NWInterface.IsConstrained]
//   - [NWInterface.IsDeepEqual]
//   - [NWInterface.IsExpensive]
//   - [NWInterface.IsIPv4Routable]
//   - [NWInterface.IsIPv6Routable]
//   - [NWInterface.IsShallowEqual]
//   - [NWInterface.IsUltraConstrained]
//   - [NWInterface.Mtu]
//   - [NWInterface.PrivateDescription]
//   - [NWInterface.Subtype]
//   - [NWInterface.SupportsMulticast]
//   - [NWInterface.Type]
//   - [NWInterface.TypeString]
//   - [NWInterface.InitWithCoder]
//   - [NWInterface.InitWithInterface]
//   - [NWInterface.InitWithInterfaceIndex]
//   - [NWInterface.InitWithInterfaceIndexInterfaceName]
//   - [NWInterface.InitWithInterfaceName]
//   - [NWInterface.Constrained]
//   - [NWInterface.Expensive]
//   - [NWInterface.Ipv4Routable]
//   - [NWInterface.Ipv6Routable]
//
// See: https://developer.apple.com/documentation/Network/NWInterface
type NWInterface struct {
	objectivec.Object
}

// NWInterfaceFromID constructs a [NWInterface] from an objc.ID.
func NWInterfaceFromID(id objc.ID) NWInterface {
	return NWInterface{objectivec.Object{ID: id}}
}

// Ensure NWInterface implements INWInterface.
var _ INWInterface = NWInterface{}

// An interface definition for the [NWInterface] class.
//
// # Methods
//
//   - [INWInterface.CInterface]
//   - [INWInterface.CopyLocalAddressForDefaultIPv4]
//   - [INWInterface.CopyLocalAddressForDefaultIPv6]
//   - [INWInterface.CopyLocalAddressForRemoteAddress]
//   - [INWInterface.CreateProtocolBufferObject]
//   - [INWInterface.DelegateInterface]
//   - [INWInterface.DescriptionWithIndentShowFullContent]
//   - [INWInterface.EncodeWithCoder]
//   - [INWInterface.Generation]
//   - [INWInterface.HasDNS]
//   - [INWInterface.HasNAT64]
//   - [INWInterface.InterfaceIndex]
//   - [INWInterface.InterfaceName]
//   - [INWInterface.InternalInterface]
//   - [INWInterface.SetInternalInterface]
//   - [INWInterface.Ipv4Broadcast]
//   - [INWInterface.Ipv4Netmask]
//   - [INWInterface.IsConstrained]
//   - [INWInterface.IsDeepEqual]
//   - [INWInterface.IsExpensive]
//   - [INWInterface.IsIPv4Routable]
//   - [INWInterface.IsIPv6Routable]
//   - [INWInterface.IsShallowEqual]
//   - [INWInterface.IsUltraConstrained]
//   - [INWInterface.Mtu]
//   - [INWInterface.PrivateDescription]
//   - [INWInterface.Subtype]
//   - [INWInterface.SupportsMulticast]
//   - [INWInterface.Type]
//   - [INWInterface.TypeString]
//   - [INWInterface.InitWithCoder]
//   - [INWInterface.InitWithInterface]
//   - [INWInterface.InitWithInterfaceIndex]
//   - [INWInterface.InitWithInterfaceIndexInterfaceName]
//   - [INWInterface.InitWithInterfaceName]
//   - [INWInterface.Constrained]
//   - [INWInterface.Expensive]
//   - [INWInterface.Ipv4Routable]
//   - [INWInterface.Ipv6Routable]
//
// See: https://developer.apple.com/documentation/Network/NWInterface
type INWInterface interface {
	objectivec.IObject

	// Topic: Methods

	CInterface() objectivec.Object
	CopyLocalAddressForDefaultIPv4() objectivec.IObject
	CopyLocalAddressForDefaultIPv6() objectivec.IObject
	CopyLocalAddressForRemoteAddress(address objectivec.IObject) objectivec.IObject
	CreateProtocolBufferObject() objectivec.IObject
	DelegateInterface() INWInterface
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	Generation() uint64
	HasDNS() bool
	HasNAT64() bool
	InterfaceIndex() uint64
	InterfaceName() string
	InternalInterface() objectivec.Object
	SetInternalInterface(value objectivec.Object)
	Ipv4Broadcast() objectivec.IObject
	Ipv4Netmask() objectivec.IObject
	IsConstrained() bool
	IsDeepEqual(equal objectivec.IObject) bool
	IsExpensive() bool
	IsIPv4Routable() bool
	IsIPv6Routable() bool
	IsShallowEqual(equal objectivec.IObject) bool
	IsUltraConstrained() bool
	Mtu() int64
	PrivateDescription() string
	Subtype() int64
	SupportsMulticast() bool
	Type() int64
	TypeString() string
	InitWithCoder(coder foundation.INSCoder) NWInterface
	InitWithInterface(interface_ objectivec.IObject) NWInterface
	InitWithInterfaceIndex(index uint64) NWInterface
	InitWithInterfaceIndexInterfaceName(index uint64, name objectivec.IObject) NWInterface
	InitWithInterfaceName(name objectivec.IObject) NWInterface
	Constrained() bool
	Expensive() bool
	Ipv4Routable() bool
	Ipv6Routable() bool
}

// Init initializes the instance.
func (n NWInterface) Init() NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWInterface) Autorelease() NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWInterface creates a new NWInterface instance.
func NewNWInterface() NWInterface {
	class := getNWInterfaceClass()
	rv := objc.Send[NWInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithCoder:
func NewNWInterfaceWithCoder(coder objectivec.IObject) NWInterface {
	instance := getNWInterfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NWInterfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterface:
func NewNWInterfaceWithInterface(interface_ objectivec.IObject) NWInterface {
	instance := getNWInterfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterface:"), interface_)
	return NWInterfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceIndex:
func NewNWInterfaceWithInterfaceIndex(index uint64) NWInterface {
	instance := getNWInterfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterfaceIndex:"), index)
	return NWInterfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceIndex:interfaceName:
func NewNWInterfaceWithInterfaceIndexInterfaceName(index uint64, name objectivec.IObject) NWInterface {
	instance := getNWInterfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterfaceIndex:interfaceName:"), index, name)
	return NWInterfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceName:
func NewNWInterfaceWithInterfaceName(name objectivec.IObject) NWInterface {
	instance := getNWInterfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterfaceName:"), name)
	return NWInterfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/copyLocalAddressForDefaultIPv4
func (n NWInterface) CopyLocalAddressForDefaultIPv4() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyLocalAddressForDefaultIPv4"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/copyLocalAddressForDefaultIPv6
func (n NWInterface) CopyLocalAddressForDefaultIPv6() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyLocalAddressForDefaultIPv6"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/copyLocalAddressForRemoteAddress:
func (n NWInterface) CopyLocalAddressForRemoteAddress(address objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyLocalAddressForRemoteAddress:"), address)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/createProtocolBufferObject
func (n NWInterface) CreateProtocolBufferObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createProtocolBufferObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/descriptionWithIndent:showFullContent:
func (n NWInterface) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/encodeWithCoder:
func (n NWInterface) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/ipv4Broadcast
func (n NWInterface) Ipv4Broadcast() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("ipv4Broadcast"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/ipv4Netmask
func (n NWInterface) Ipv4Netmask() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("ipv4Netmask"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isConstrained
func (n NWInterface) IsConstrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isConstrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isDeepEqual:
func (n NWInterface) IsDeepEqual(equal objectivec.IObject) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isDeepEqual:"), equal)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isExpensive
func (n NWInterface) IsExpensive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isExpensive"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isIPv4Routable
func (n NWInterface) IsIPv4Routable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isIPv4Routable"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isIPv6Routable
func (n NWInterface) IsIPv6Routable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isIPv6Routable"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isShallowEqual:
func (n NWInterface) IsShallowEqual(equal objectivec.IObject) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isShallowEqual:"), equal)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/isUltraConstrained
func (n NWInterface) IsUltraConstrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isUltraConstrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithCoder:
func (n NWInterface) InitWithCoder(coder foundation.INSCoder) NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterface:
func (n NWInterface) InitWithInterface(interface_ objectivec.IObject) NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("initWithInterface:"), interface_)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceIndex:
func (n NWInterface) InitWithInterfaceIndex(index uint64) NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("initWithInterfaceIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceIndex:interfaceName:
func (n NWInterface) InitWithInterfaceIndexInterfaceName(index uint64, name objectivec.IObject) NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("initWithInterfaceIndex:interfaceName:"), index, name)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/initWithInterfaceName:
func (n NWInterface) InitWithInterfaceName(name objectivec.IObject) NWInterface {
	rv := objc.Send[NWInterface](n.ID, objc.Sel("initWithInterfaceName:"), name)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/descriptionForSubtype:
func (_NWInterfaceClass NWInterfaceClass) DescriptionForSubtype(subtype int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWInterfaceClass.class), objc.Sel("descriptionForSubtype:"), subtype)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/descriptionForType:
func (_NWInterfaceClass NWInterfaceClass) DescriptionForType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWInterfaceClass.class), objc.Sel("descriptionForType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/interfaceWithProtocolBufferData:
func (_NWInterfaceClass NWInterfaceClass) InterfaceWithProtocolBufferData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWInterfaceClass.class), objc.Sel("interfaceWithProtocolBufferData:"), data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWInterface/supportsSecureCoding
func (_NWInterfaceClass NWInterfaceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_NWInterfaceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/cInterface
func (n NWInterface) CInterface() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("cInterface"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWInterface/constrained
func (n NWInterface) Constrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("constrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/delegateInterface
func (n NWInterface) DelegateInterface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("delegateInterface"))
	return NWInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWInterface/expensive
func (n NWInterface) Expensive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("expensive"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/generation
func (n NWInterface) Generation() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("generation"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/hasDNS
func (n NWInterface) HasDNS() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasDNS"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/hasNAT64
func (n NWInterface) HasNAT64() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasNAT64"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/interfaceIndex
func (n NWInterface) InterfaceIndex() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("interfaceIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/interfaceName
func (n NWInterface) InterfaceName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("interfaceName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Network/NWInterface/internalInterface
func (n NWInterface) InternalInterface() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalInterface"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWInterface) SetInternalInterface(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalInterface:"), value)
}

// See: https://developer.apple.com/documentation/Network/NWInterface/ipv4Routable
func (n NWInterface) Ipv4Routable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("ipv4Routable"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/ipv6Routable
func (n NWInterface) Ipv6Routable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("ipv6Routable"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/mtu
func (n NWInterface) Mtu() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("mtu"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/privateDescription
func (n NWInterface) PrivateDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("privateDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Network/NWInterface/subtype
func (n NWInterface) Subtype() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("subtype"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/supportsMulticast
func (n NWInterface) SupportsMulticast() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("supportsMulticast"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/type
func (n NWInterface) Type() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("type"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWInterface/typeString
func (n NWInterface) TypeString() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("typeString"))
	return foundation.NSStringFromID(rv).String()
}
