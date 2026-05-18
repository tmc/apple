// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWPath] class.
var (
	_NWPathClass     NWPathClass
	_NWPathClassOnce sync.Once
)

func getNWPathClass() NWPathClass {
	_NWPathClassOnce.Do(func() {
		_NWPathClass = NWPathClass{class: objc.GetClass("NWPath")}
	})
	return _NWPathClass
}

// GetNWPathClass returns the class object for NWPath.
func GetNWPathClass() NWPathClass {
	return getNWPathClass()
}

type NWPathClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWPathClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWPathClass) Alloc() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWPath.AdvertiseDescriptor]
//   - [NWPath.BrowseDescriptor]
//   - [NWPath.CPath]
//   - [NWPath.ClientID]
//   - [NWPath.ConnectedInterface]
//   - [NWPath.CopyDNSSearchDomains]
//   - [NWPath.CopyDNSServerEndpoints]
//   - [NWPath.CopyDNSServersStrings]
//   - [NWPath.CopyDataFromNetworkAgentWithDomainType]
//   - [NWPath.CopyFlowDivertToken]
//   - [NWPath.CreateProtocolBufferObject]
//   - [NWPath.DelegateInterface]
//   - [NWPath.DerivedParameters]
//   - [NWPath.DescriptionWithIndentShowFullContent]
//   - [NWPath.DnsSearchDomains]
//   - [NWPath.DnsServers]
//   - [NWPath.DnsServersAsStrings]
//   - [NWPath.DnsServiceID]
//   - [NWPath.EffectiveLocalEndpoint]
//   - [NWPath.EffectiveRemoteEndpoint]
//   - [NWPath.Endpoint]
//   - [NWPath.FallbackEligible]
//   - [NWPath.FallbackInterface]
//   - [NWPath.FallbackInterfaceIndex]
//   - [NWPath.FallbackIsPreferred]
//   - [NWPath.FallbackIsWeak]
//   - [NWPath.FilterControlUnit]
//   - [NWPath.FlowDivertAggregateUnit]
//   - [NWPath.FlowDivertControlUnit]
//   - [NWPath.Flows]
//   - [NWPath.Gateways]
//   - [NWPath.GenericNetworkAgentsWithDomainType]
//   - [NWPath.GroupMembers]
//   - [NWPath.HasAdvertiseDescriptor]
//   - [NWPath.HasApplicationLevelFirewall]
//   - [NWPath.HasBrowseDescriptor]
//   - [NWPath.HasCustomPFRules]
//   - [NWPath.HasKernelExtensionFilter]
//   - [NWPath.HasParentalControls]
//   - [NWPath.HasProxySettings]
//   - [NWPath.HasUnsatisfiedFallbackAgent]
//   - [NWPath.InactiveNetworkAgentUUIDsOnlyVoluntary]
//   - [NWPath.Interface]
//   - [NWPath.InternalPath]
//   - [NWPath.IsConstrained]
//   - [NWPath.IsDirect]
//   - [NWPath.IsEligibleForCrazyIvan46]
//   - [NWPath.IsEqualToPath]
//   - [NWPath.IsExpensive]
//   - [NWPath.IsFiltered]
//   - [NWPath.IsFlowDivert]
//   - [NWPath.IsLinkQualityAbort]
//   - [NWPath.IsListener]
//   - [NWPath.IsListenerInterfaceSpecific]
//   - [NWPath.IsLocal]
//   - [NWPath.IsPerAppVPN]
//   - [NWPath.IsRoaming]
//   - [NWPath.IsUltraConstrained]
//   - [NWPath.IsViable]
//   - [NWPath.MaximumDatagramSize]
//   - [NWPath.Mtu]
//   - [NWPath.NetworkAgentsOfType]
//   - [NWPath.OverrideDNSSearchDomains]
//   - [NWPath.OverrideDNSServers]
//   - [NWPath.OverrideDNSServersAsStrings]
//   - [NWPath.Parameters]
//   - [NWPath.PolicyID]
//   - [NWPath.PrivateDescription]
//   - [NWPath.ProxySettings]
//   - [NWPath.Reason]
//   - [NWPath.ReasonDescription]
//   - [NWPath.ScopedInterface]
//   - [NWPath.SecondsSinceInterfaceChange]
//   - [NWPath.ShouldProbeConnectivity]
//   - [NWPath.Status]
//   - [NWPath.StatusAsString]
//   - [NWPath.SupportsDNS]
//   - [NWPath.SupportsIPv4]
//   - [NWPath.SupportsIPv6]
//   - [NWPath.UnsatisfiedVoluntaryAgentMatchesAddressTriggerImmediately]
//   - [NWPath.UsesCompanion]
//   - [NWPath.UsesInterfaceType]
//   - [NWPath.UsesNetworkAgent]
//   - [NWPath.UsesNetworkAgentType]
//   - [NWPath.InitWithPath]
//   - [NWPath.Constrained]
//   - [NWPath.Direct]
//   - [NWPath.EligibleForCrazyIvan46]
//   - [NWPath.Expensive]
//   - [NWPath.Filtered]
//   - [NWPath.FlowDivert]
//   - [NWPath.Listener]
//   - [NWPath.Local]
//   - [NWPath.PerAppVPN]
//   - [NWPath.Roaming]
//   - [NWPath.Viable]
//
// See: https://developer.apple.com/documentation/Network/NWPath
type NWPath struct {
	objectivec.Object
}

// NWPathFromID constructs a [NWPath] from an objc.ID.
func NWPathFromID(id objc.ID) NWPath {
	return NWPath{objectivec.Object{ID: id}}
}

// Ensure NWPath implements INWPath.
var _ INWPath = NWPath{}

// An interface definition for the [NWPath] class.
//
// # Methods
//
//   - [INWPath.AdvertiseDescriptor]
//   - [INWPath.BrowseDescriptor]
//   - [INWPath.CPath]
//   - [INWPath.ClientID]
//   - [INWPath.ConnectedInterface]
//   - [INWPath.CopyDNSSearchDomains]
//   - [INWPath.CopyDNSServerEndpoints]
//   - [INWPath.CopyDNSServersStrings]
//   - [INWPath.CopyDataFromNetworkAgentWithDomainType]
//   - [INWPath.CopyFlowDivertToken]
//   - [INWPath.CreateProtocolBufferObject]
//   - [INWPath.DelegateInterface]
//   - [INWPath.DerivedParameters]
//   - [INWPath.DescriptionWithIndentShowFullContent]
//   - [INWPath.DnsSearchDomains]
//   - [INWPath.DnsServers]
//   - [INWPath.DnsServersAsStrings]
//   - [INWPath.DnsServiceID]
//   - [INWPath.EffectiveLocalEndpoint]
//   - [INWPath.EffectiveRemoteEndpoint]
//   - [INWPath.Endpoint]
//   - [INWPath.FallbackEligible]
//   - [INWPath.FallbackInterface]
//   - [INWPath.FallbackInterfaceIndex]
//   - [INWPath.FallbackIsPreferred]
//   - [INWPath.FallbackIsWeak]
//   - [INWPath.FilterControlUnit]
//   - [INWPath.FlowDivertAggregateUnit]
//   - [INWPath.FlowDivertControlUnit]
//   - [INWPath.Flows]
//   - [INWPath.Gateways]
//   - [INWPath.GenericNetworkAgentsWithDomainType]
//   - [INWPath.GroupMembers]
//   - [INWPath.HasAdvertiseDescriptor]
//   - [INWPath.HasApplicationLevelFirewall]
//   - [INWPath.HasBrowseDescriptor]
//   - [INWPath.HasCustomPFRules]
//   - [INWPath.HasKernelExtensionFilter]
//   - [INWPath.HasParentalControls]
//   - [INWPath.HasProxySettings]
//   - [INWPath.HasUnsatisfiedFallbackAgent]
//   - [INWPath.InactiveNetworkAgentUUIDsOnlyVoluntary]
//   - [INWPath.Interface]
//   - [INWPath.InternalPath]
//   - [INWPath.IsConstrained]
//   - [INWPath.IsDirect]
//   - [INWPath.IsEligibleForCrazyIvan46]
//   - [INWPath.IsEqualToPath]
//   - [INWPath.IsExpensive]
//   - [INWPath.IsFiltered]
//   - [INWPath.IsFlowDivert]
//   - [INWPath.IsLinkQualityAbort]
//   - [INWPath.IsListener]
//   - [INWPath.IsListenerInterfaceSpecific]
//   - [INWPath.IsLocal]
//   - [INWPath.IsPerAppVPN]
//   - [INWPath.IsRoaming]
//   - [INWPath.IsUltraConstrained]
//   - [INWPath.IsViable]
//   - [INWPath.MaximumDatagramSize]
//   - [INWPath.Mtu]
//   - [INWPath.NetworkAgentsOfType]
//   - [INWPath.OverrideDNSSearchDomains]
//   - [INWPath.OverrideDNSServers]
//   - [INWPath.OverrideDNSServersAsStrings]
//   - [INWPath.Parameters]
//   - [INWPath.PolicyID]
//   - [INWPath.PrivateDescription]
//   - [INWPath.ProxySettings]
//   - [INWPath.Reason]
//   - [INWPath.ReasonDescription]
//   - [INWPath.ScopedInterface]
//   - [INWPath.SecondsSinceInterfaceChange]
//   - [INWPath.ShouldProbeConnectivity]
//   - [INWPath.Status]
//   - [INWPath.StatusAsString]
//   - [INWPath.SupportsDNS]
//   - [INWPath.SupportsIPv4]
//   - [INWPath.SupportsIPv6]
//   - [INWPath.UnsatisfiedVoluntaryAgentMatchesAddressTriggerImmediately]
//   - [INWPath.UsesCompanion]
//   - [INWPath.UsesInterfaceType]
//   - [INWPath.UsesNetworkAgent]
//   - [INWPath.UsesNetworkAgentType]
//   - [INWPath.InitWithPath]
//   - [INWPath.Constrained]
//   - [INWPath.Direct]
//   - [INWPath.EligibleForCrazyIvan46]
//   - [INWPath.Expensive]
//   - [INWPath.Filtered]
//   - [INWPath.FlowDivert]
//   - [INWPath.Listener]
//   - [INWPath.Local]
//   - [INWPath.PerAppVPN]
//   - [INWPath.Roaming]
//   - [INWPath.Viable]
//
// See: https://developer.apple.com/documentation/Network/NWPath
type INWPath interface {
	objectivec.IObject

	// Topic: Methods

	AdvertiseDescriptor() INWAdvertiseDescriptor
	BrowseDescriptor() INWBrowseDescriptor
	CPath() objectivec.Object
	ClientID() foundation.NSUUID
	ConnectedInterface() INWInterface
	CopyDNSSearchDomains(domains bool) objectivec.IObject
	CopyDNSServerEndpoints(endpoints bool) objectivec.IObject
	CopyDNSServersStrings(strings objectivec.IObject) objectivec.IObject
	CopyDataFromNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) objectivec.IObject
	CopyFlowDivertToken() objectivec.IObject
	CreateProtocolBufferObject() objectivec.IObject
	DelegateInterface() objectivec.IObject
	DerivedParameters() INWParameters
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	DnsSearchDomains() foundation.INSArray
	DnsServers() foundation.INSArray
	DnsServersAsStrings() foundation.INSArray
	DnsServiceID() int
	EffectiveLocalEndpoint() INWEndpoint
	EffectiveRemoteEndpoint() INWEndpoint
	Endpoint() INWEndpoint
	FallbackEligible() bool
	FallbackInterface() INWInterface
	FallbackInterfaceIndex() uint32
	FallbackIsPreferred() bool
	FallbackIsWeak() bool
	FilterControlUnit() uint32
	FlowDivertAggregateUnit() uint32
	FlowDivertControlUnit() uint32
	Flows() foundation.INSArray
	Gateways() foundation.INSArray
	GenericNetworkAgentsWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) objectivec.IObject
	GroupMembers() foundation.INSArray
	HasAdvertiseDescriptor() bool
	HasApplicationLevelFirewall() bool
	HasBrowseDescriptor() bool
	HasCustomPFRules() bool
	HasKernelExtensionFilter() bool
	HasParentalControls() bool
	HasProxySettings() bool
	HasUnsatisfiedFallbackAgent() bool
	InactiveNetworkAgentUUIDsOnlyVoluntary(voluntary bool) objectivec.IObject
	Interface() INWInterface
	InternalPath() objectivec.Object
	IsConstrained() bool
	IsDirect() bool
	IsEligibleForCrazyIvan46() bool
	IsEqualToPath(path objectivec.IObject) bool
	IsExpensive() bool
	IsFiltered() bool
	IsFlowDivert() bool
	IsLinkQualityAbort() bool
	IsListener() bool
	IsListenerInterfaceSpecific() bool
	IsLocal() bool
	IsPerAppVPN() bool
	IsRoaming() bool
	IsUltraConstrained() bool
	IsViable() bool
	MaximumDatagramSize() int64
	Mtu() int64
	NetworkAgentsOfType(type_ objc.Class) objectivec.IObject
	OverrideDNSSearchDomains() foundation.INSArray
	OverrideDNSServers() foundation.INSArray
	OverrideDNSServersAsStrings() foundation.INSArray
	Parameters() INWParameters
	PolicyID() uint32
	PrivateDescription() string
	ProxySettings() foundation.INSArray
	Reason() int64
	ReasonDescription() string
	ScopedInterface() INWInterface
	SecondsSinceInterfaceChange() uint64
	ShouldProbeConnectivity() bool
	Status() int64
	StatusAsString() string
	SupportsDNS() bool
	SupportsIPv4() bool
	SupportsIPv6() bool
	UnsatisfiedVoluntaryAgentMatchesAddressTriggerImmediately(address objectivec.IObject) (bool, bool)
	UsesCompanion() bool
	UsesInterfaceType(type_ int64) bool
	UsesNetworkAgent(agent objectivec.IObject) bool
	UsesNetworkAgentType(type_ objc.Class) bool
	InitWithPath(path objectivec.IObject) NWPath
	Constrained() bool
	Direct() bool
	EligibleForCrazyIvan46() bool
	Expensive() bool
	Filtered() bool
	FlowDivert() bool
	Listener() bool
	Local() bool
	PerAppVPN() bool
	Roaming() bool
	Viable() bool
}

// Init initializes the instance.
func (n NWPath) Init() NWPath {
	rv := objc.Send[NWPath](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWPath) Autorelease() NWPath {
	rv := objc.Send[NWPath](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWPath creates a new NWPath instance.
func NewNWPath() NWPath {
	class := getNWPathClass()
	rv := objc.Send[NWPath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/initWithPath:
func NewNWPathWithPath(path objectivec.IObject) NWPath {
	instance := getNWPathClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:"), path)
	return NWPathFromID(rv)
}

// See: https://developer.apple.com/documentation/Network/NWPath/copyDNSSearchDomains:
func (n NWPath) CopyDNSSearchDomains(domains bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyDNSSearchDomains:"), domains)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/copyDNSServerEndpoints:
func (n NWPath) CopyDNSServerEndpoints(endpoints bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyDNSServerEndpoints:"), endpoints)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/copyDNSServersStrings:
func (n NWPath) CopyDNSServersStrings(strings objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyDNSServersStrings:"), strings)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/copyDataFromNetworkAgentWithDomain:type:
func (n NWPath) CopyDataFromNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyDataFromNetworkAgentWithDomain:type:"), domain, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/copyFlowDivertToken
func (n NWPath) CopyFlowDivertToken() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyFlowDivertToken"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/createProtocolBufferObject
func (n NWPath) CreateProtocolBufferObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createProtocolBufferObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/delegateInterface
func (n NWPath) DelegateInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("delegateInterface"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/descriptionWithIndent:showFullContent:
func (n NWPath) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/genericNetworkAgentsWithDomain:type:
func (n NWPath) GenericNetworkAgentsWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("genericNetworkAgentsWithDomain:type:"), domain, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasUnsatisfiedFallbackAgent
func (n NWPath) HasUnsatisfiedFallbackAgent() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasUnsatisfiedFallbackAgent"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/inactiveNetworkAgentUUIDsOnlyVoluntary:
func (n NWPath) InactiveNetworkAgentUUIDsOnlyVoluntary(voluntary bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("inactiveNetworkAgentUUIDsOnlyVoluntary:"), voluntary)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/isConstrained
func (n NWPath) IsConstrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isConstrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isDirect
func (n NWPath) IsDirect() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isDirect"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isEligibleForCrazyIvan46
func (n NWPath) IsEligibleForCrazyIvan46() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isEligibleForCrazyIvan46"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isEqualToPath:
func (n NWPath) IsEqualToPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isEqualToPath:"), path)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isExpensive
func (n NWPath) IsExpensive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isExpensive"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isFiltered
func (n NWPath) IsFiltered() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isFiltered"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isFlowDivert
func (n NWPath) IsFlowDivert() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isFlowDivert"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isLinkQualityAbort
func (n NWPath) IsLinkQualityAbort() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isLinkQualityAbort"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isListener
func (n NWPath) IsListener() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isListener"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isListenerInterfaceSpecific
func (n NWPath) IsListenerInterfaceSpecific() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isListenerInterfaceSpecific"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isLocal
func (n NWPath) IsLocal() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isLocal"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isPerAppVPN
func (n NWPath) IsPerAppVPN() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isPerAppVPN"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isRoaming
func (n NWPath) IsRoaming() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isRoaming"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isUltraConstrained
func (n NWPath) IsUltraConstrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isUltraConstrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/isViable
func (n NWPath) IsViable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isViable"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/networkAgentsOfType:
func (n NWPath) NetworkAgentsOfType(type_ objc.Class) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("networkAgentsOfType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/shouldProbeConnectivity
func (n NWPath) ShouldProbeConnectivity() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("shouldProbeConnectivity"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/unsatisfiedVoluntaryAgentMatchesAddress:triggerImmediately:
func (n NWPath) UnsatisfiedVoluntaryAgentMatchesAddressTriggerImmediately(address objectivec.IObject) (bool, bool) {
	var immediately bool
	rv := objc.Send[bool](n.ID, objc.Sel("unsatisfiedVoluntaryAgentMatchesAddress:triggerImmediately:"), address, unsafe.Pointer(&immediately))
	return immediately, rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/usesCompanion
func (n NWPath) UsesCompanion() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesCompanion"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/usesInterfaceType:
func (n NWPath) UsesInterfaceType(type_ int64) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesInterfaceType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/usesNetworkAgent:
func (n NWPath) UsesNetworkAgent(agent objectivec.IObject) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesNetworkAgent:"), agent)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/usesNetworkAgentType:
func (n NWPath) UsesNetworkAgentType(type_ objc.Class) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("usesNetworkAgentType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/initWithPath:
func (n NWPath) InitWithPath(path objectivec.IObject) NWPath {
	rv := objc.Send[NWPath](n.ID, objc.Sel("initWithPath:"), path)
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/allClientIDs
func (_NWPathClass NWPathClass) AllClientIDs() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWPathClass.class), objc.Sel("allClientIDs"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/createStringFromStatus:
func (_NWPathClass NWPathClass) CreateStringFromStatus(status int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWPathClass.class), objc.Sel("createStringFromStatus:"), status)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/pathForClientID:
func (_NWPathClass NWPathClass) PathForClientID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWPathClass.class), objc.Sel("pathForClientID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/pathForClientID:parametersTLV:pathResultTLV:
func (_NWPathClass NWPathClass) PathForClientIDParametersTLVPathResultTLV(id objectivec.IObject, tlv objectivec.IObject, tlv2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWPathClass.class), objc.Sel("pathForClientID:parametersTLV:pathResultTLV:"), id, tlv, tlv2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/pathWithProtocolBufferData:
func (_NWPathClass NWPathClass) PathWithProtocolBufferData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWPathClass.class), objc.Sel("pathWithProtocolBufferData:"), data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Network/NWPath/advertiseDescriptor
func (n NWPath) AdvertiseDescriptor() INWAdvertiseDescriptor {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("advertiseDescriptor"))
	return NWAdvertiseDescriptorFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/browseDescriptor
func (n NWPath) BrowseDescriptor() INWBrowseDescriptor {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("browseDescriptor"))
	return NWBrowseDescriptorFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/cPath
func (n NWPath) CPath() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("cPath"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/clientID
func (n NWPath) ClientID() foundation.NSUUID {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("clientID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/connectedInterface
func (n NWPath) ConnectedInterface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("connectedInterface"))
	return NWInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/constrained
func (n NWPath) Constrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("constrained"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/derivedParameters
func (n NWPath) DerivedParameters() INWParameters {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("derivedParameters"))
	return NWParametersFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/direct
func (n NWPath) Direct() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("direct"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/dnsSearchDomains
func (n NWPath) DnsSearchDomains() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("dnsSearchDomains"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/dnsServers
func (n NWPath) DnsServers() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("dnsServers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/dnsServersAsStrings
func (n NWPath) DnsServersAsStrings() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("dnsServersAsStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/dnsServiceID
func (n NWPath) DnsServiceID() int {
	rv := objc.Send[int](n.ID, objc.Sel("dnsServiceID"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/effectiveLocalEndpoint
func (n NWPath) EffectiveLocalEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("effectiveLocalEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/effectiveRemoteEndpoint
func (n NWPath) EffectiveRemoteEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("effectiveRemoteEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/eligibleForCrazyIvan46
func (n NWPath) EligibleForCrazyIvan46() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("eligibleForCrazyIvan46"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/endpoint
func (n NWPath) Endpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("endpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/expensive
func (n NWPath) Expensive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("expensive"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/fallbackEligible
func (n NWPath) FallbackEligible() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("fallbackEligible"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/fallbackInterface
func (n NWPath) FallbackInterface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("fallbackInterface"))
	return NWInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/fallbackInterfaceIndex
func (n NWPath) FallbackInterfaceIndex() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("fallbackInterfaceIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/fallbackIsPreferred
func (n NWPath) FallbackIsPreferred() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("fallbackIsPreferred"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/fallbackIsWeak
func (n NWPath) FallbackIsWeak() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("fallbackIsWeak"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/filterControlUnit
func (n NWPath) FilterControlUnit() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("filterControlUnit"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/filtered
func (n NWPath) Filtered() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("filtered"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/flowDivert
func (n NWPath) FlowDivert() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("flowDivert"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/flowDivertAggregateUnit
func (n NWPath) FlowDivertAggregateUnit() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("flowDivertAggregateUnit"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/flowDivertControlUnit
func (n NWPath) FlowDivertControlUnit() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("flowDivertControlUnit"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/flows
func (n NWPath) Flows() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("flows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/gateways
func (n NWPath) Gateways() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("gateways"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/groupMembers
func (n NWPath) GroupMembers() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("groupMembers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasAdvertiseDescriptor
func (n NWPath) HasAdvertiseDescriptor() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasAdvertiseDescriptor"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasApplicationLevelFirewall
func (n NWPath) HasApplicationLevelFirewall() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasApplicationLevelFirewall"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasBrowseDescriptor
func (n NWPath) HasBrowseDescriptor() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasBrowseDescriptor"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasCustomPFRules
func (n NWPath) HasCustomPFRules() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasCustomPFRules"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasKernelExtensionFilter
func (n NWPath) HasKernelExtensionFilter() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasKernelExtensionFilter"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasParentalControls
func (n NWPath) HasParentalControls() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasParentalControls"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/hasProxySettings
func (n NWPath) HasProxySettings() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasProxySettings"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/interface
func (n NWPath) Interface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("interface"))
	return NWInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/internalPath
func (n NWPath) InternalPath() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalPath"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/listener
func (n NWPath) Listener() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("listener"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/local
func (n NWPath) Local() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("local"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/maximumDatagramSize
func (n NWPath) MaximumDatagramSize() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("maximumDatagramSize"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/mtu
func (n NWPath) Mtu() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("mtu"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/overrideDNSSearchDomains
func (n NWPath) OverrideDNSSearchDomains() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("overrideDNSSearchDomains"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/overrideDNSServers
func (n NWPath) OverrideDNSServers() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("overrideDNSServers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/overrideDNSServersAsStrings
func (n NWPath) OverrideDNSServersAsStrings() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("overrideDNSServersAsStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/parameters
func (n NWPath) Parameters() INWParameters {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameters"))
	return NWParametersFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/perAppVPN
func (n NWPath) PerAppVPN() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("perAppVPN"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/policyID
func (n NWPath) PolicyID() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("policyID"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/privateDescription
func (n NWPath) PrivateDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("privateDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Network/NWPath/proxySettings
func (n NWPath) ProxySettings() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("proxySettings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/reason
func (n NWPath) Reason() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("reason"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/reasonDescription
func (n NWPath) ReasonDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reasonDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Network/NWPath/roaming
func (n NWPath) Roaming() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("roaming"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/scopedInterface
func (n NWPath) ScopedInterface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("scopedInterface"))
	return NWInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Network/NWPath/secondsSinceInterfaceChange
func (n NWPath) SecondsSinceInterfaceChange() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("secondsSinceInterfaceChange"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/status
func (n NWPath) Status() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("status"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/statusAsString
func (n NWPath) StatusAsString() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("statusAsString"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Network/NWPath/supportsDNS
func (n NWPath) SupportsDNS() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("supportsDNS"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/supportsIPv4
func (n NWPath) SupportsIPv4() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("supportsIPv4"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/supportsIPv6
func (n NWPath) SupportsIPv6() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("supportsIPv6"))
	return rv
}

// See: https://developer.apple.com/documentation/Network/NWPath/viable
func (n NWPath) Viable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("viable"))
	return rv
}
