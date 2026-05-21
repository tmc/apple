// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWParameters] class.
var (
	_NWParametersClass     NWParametersClass
	_NWParametersClassOnce sync.Once
)

func getNWParametersClass() NWParametersClass {
	_NWParametersClassOnce.Do(func() {
		_NWParametersClass = NWParametersClass{class: objc.GetClass("NWParameters")}
	})
	return _NWParametersClass
}

// GetNWParametersClass returns the class object for NWParameters.
func GetNWParametersClass() NWParametersClass {
	return getNWParametersClass()
}

type NWParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWParametersClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWParametersClass) Alloc() NWParameters {
	rv := objc.Send[NWParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWParameters.SSLCipherSuites]
//   - [NWParameters.SetSSLCipherSuites]
//   - [NWParameters.SSLCipherSuitesInternal]
//   - [NWParameters.SetSSLCipherSuitesInternal]
//   - [NWParameters.TLSSessionID]
//   - [NWParameters.SetTLSSessionID]
//   - [NWParameters.Account]
//   - [NWParameters.SetAccount]
//   - [NWParameters.AllowDuplicateStateUpdates]
//   - [NWParameters.SetAllowDuplicateStateUpdates]
//   - [NWParameters.AllowJoiningConnectedFd]
//   - [NWParameters.SetAllowJoiningConnectedFd]
//   - [NWParameters.AllowSocketAccess]
//   - [NWParameters.SetAllowSocketAccess]
//   - [NWParameters.AllowUnusableAddresses]
//   - [NWParameters.SetAllowUnusableAddresses]
//   - [NWParameters.AttachProtocolListener]
//   - [NWParameters.SetAttachProtocolListener]
//   - [NWParameters.AvoidNetworkAgentWithDomainType]
//   - [NWParameters.AvoidNetworkAgentWithUUID]
//   - [NWParameters.CopyCParameters]
//   - [NWParameters.CopyRequiredAgentsDescription]
//   - [NWParameters.CreateProtocolBufferObject]
//   - [NWParameters.DataMode]
//   - [NWParameters.SetDataMode]
//   - [NWParameters.DescriptionWithIndentShowFullContent]
//   - [NWParameters.DisableNagleAlgorithm]
//   - [NWParameters.SetDisableNagleAlgorithm]
//   - [NWParameters.EffectiveBundleID]
//   - [NWParameters.SetEffectiveBundleID]
//   - [NWParameters.EffectiveProcessUUID]
//   - [NWParameters.SetEffectiveProcessUUID]
//   - [NWParameters.EnableTFO]
//   - [NWParameters.SetEnableTFO]
//   - [NWParameters.EnableTFONoCookie]
//   - [NWParameters.SetEnableTFONoCookie]
//   - [NWParameters.EnableTLS]
//   - [NWParameters.SetEnableTLS]
//   - [NWParameters.EncodeWithCoder]
//   - [NWParameters.FastOpenForceEnable]
//   - [NWParameters.SetFastOpenForceEnable]
//   - [NWParameters.HasDelegatedPIDForOriginatingPID]
//   - [NWParameters.HasDelegatedProcessUUID]
//   - [NWParameters.HasNonEmptyProxyConfiguration]
//   - [NWParameters.HasPreferredNetworkAgents]
//   - [NWParameters.HasProhibitedNetworkAgents]
//   - [NWParameters.HasRequiredNetworkAgents]
//   - [NWParameters.HttpsProxyIsOpaque]
//   - [NWParameters.SetHttpsProxyIsOpaque]
//   - [NWParameters.HttpsProxyOverTLS]
//   - [NWParameters.SetHttpsProxyOverTLS]
//   - [NWParameters.IgnoreResolverStats]
//   - [NWParameters.SetIgnoreResolverStats]
//   - [NWParameters.Indefinite]
//   - [NWParameters.SetIndefinite]
//   - [NWParameters.InternalParameters]
//   - [NWParameters.SetInternalParameters]
//   - [NWParameters.IpProtocol]
//   - [NWParameters.IsDiscretionary]
//   - [NWParameters.IsDryRun]
//   - [NWParameters.IsValid]
//   - [NWParameters.KeepAlive]
//   - [NWParameters.SetKeepAlive]
//   - [NWParameters.KeepAliveIdleTime]
//   - [NWParameters.SetKeepAliveIdleTime]
//   - [NWParameters.KeepAliveInterval]
//   - [NWParameters.SetKeepAliveInterval]
//   - [NWParameters.KeepAliveOffload]
//   - [NWParameters.SetKeepAliveOffload]
//   - [NWParameters.LocalAddress]
//   - [NWParameters.SetLocalAddress]
//   - [NWParameters.MaximumSSLProtocolVersion]
//   - [NWParameters.SetMaximumSSLProtocolVersion]
//   - [NWParameters.MaximumSSLProtocolVersionInternal]
//   - [NWParameters.SetMaximumSSLProtocolVersionInternal]
//   - [NWParameters.Metadata]
//   - [NWParameters.SetMetadata]
//   - [NWParameters.MinimumSSLProtocolVersion]
//   - [NWParameters.SetMinimumSSLProtocolVersion]
//   - [NWParameters.MinimumSSLProtocolVersionInternal]
//   - [NWParameters.SetMinimumSSLProtocolVersionInternal]
//   - [NWParameters.Multipath]
//   - [NWParameters.SetMultipath]
//   - [NWParameters.MultipathForceEnable]
//   - [NWParameters.SetMultipathForceEnable]
//   - [NWParameters.MultipathService]
//   - [NWParameters.SetMultipathService]
//   - [NWParameters.NoProxy]
//   - [NWParameters.SetNoProxy]
//   - [NWParameters.ParentID]
//   - [NWParameters.SetParentID]
//   - [NWParameters.ParentIDs]
//   - [NWParameters.Pid]
//   - [NWParameters.SetPid]
//   - [NWParameters.PreferNetworkAgentWithDomainType]
//   - [NWParameters.PreferNetworkAgentWithUUID]
//   - [NWParameters.PreferNoProxy]
//   - [NWParameters.SetPreferNoProxy]
//   - [NWParameters.PrivateDescription]
//   - [NWParameters.ProcessUUID]
//   - [NWParameters.SetProcessUUID]
//   - [NWParameters.ProhibitCellular]
//   - [NWParameters.ProhibitConstrainedPaths]
//   - [NWParameters.SetProhibitConstrainedPaths]
//   - [NWParameters.ProhibitExpensivePaths]
//   - [NWParameters.SetProhibitExpensivePaths]
//   - [NWParameters.ProhibitFallback]
//   - [NWParameters.SetProhibitFallback]
//   - [NWParameters.ProhibitInterface]
//   - [NWParameters.ProhibitInterfaceSubtype]
//   - [NWParameters.ProhibitInterfaceType]
//   - [NWParameters.ProhibitJoiningProtocols]
//   - [NWParameters.SetProhibitJoiningProtocols]
//   - [NWParameters.ProhibitNetworkAgentWithUUID]
//   - [NWParameters.ProhibitNetworkAgentsWithDomainType]
//   - [NWParameters.ProhibitRoaming]
//   - [NWParameters.SetProhibitRoaming]
//   - [NWParameters.ProtocolTransforms]
//   - [NWParameters.SetProtocolTransforms]
//   - [NWParameters.ProxyConfiguration]
//   - [NWParameters.SetProxyConfiguration]
//   - [NWParameters.ReduceBuffering]
//   - [NWParameters.SetReduceBuffering]
//   - [NWParameters.RequireNetworkAgentWithDomainType]
//   - [NWParameters.RequireNetworkAgentWithUUID]
//   - [NWParameters.RequiredAddressFamily]
//   - [NWParameters.SetRequiredAddressFamily]
//   - [NWParameters.RequiredCompanionProxyInterfaceType]
//   - [NWParameters.SetRequiredCompanionProxyInterfaceType]
//   - [NWParameters.RequiredInterface]
//   - [NWParameters.SetRequiredInterface]
//   - [NWParameters.RequiredInterfaceSubtype]
//   - [NWParameters.SetRequiredInterfaceSubtype]
//   - [NWParameters.RequiredInterfaceType]
//   - [NWParameters.SetRequiredInterfaceType]
//   - [NWParameters.ResolvePTR]
//   - [NWParameters.SetResolvePTR]
//   - [NWParameters.ReuseLocalAddress]
//   - [NWParameters.SetReuseLocalAddress]
//   - [NWParameters.SanitizedURL]
//   - [NWParameters.SetInitialDataPayload]
//   - [NWParameters.SetSourceApplicationWithBundleID]
//   - [NWParameters.TlsVersionWithSSLProtocol]
//   - [NWParameters.TrafficClass]
//   - [NWParameters.SetTrafficClass]
//   - [NWParameters.TransportProtocol]
//   - [NWParameters.TrustInvalidCertificates]
//   - [NWParameters.SetTrustInvalidCertificates]
//   - [NWParameters.Uid]
//   - [NWParameters.SetUid]
//   - [NWParameters.Url]
//   - [NWParameters.SetUrl]
//   - [NWParameters.UseAWDL]
//   - [NWParameters.SetUseAWDL]
//   - [NWParameters.UseLongOutstandingQueries]
//   - [NWParameters.SetUseLongOutstandingQueries]
//   - [NWParameters.UseP2P]
//   - [NWParameters.SetUseP2P]
//   - [NWParameters.InitWithCoder]
//   - [NWParameters.InitWithParameters]
//   - [NWParameters.Discretionary]
//   - [NWParameters.SetDiscretionary]
//   - [NWParameters.DryRun]
//   - [NWParameters.Valid]
type NWParameters struct {
	objectivec.Object
}

// NWParametersFromID constructs a [NWParameters] from an objc.ID.
func NWParametersFromID(id objc.ID) NWParameters {
	return NWParameters{objectivec.Object{ID: id}}
}

// Ensure NWParameters implements INWParameters.
var _ INWParameters = NWParameters{}

// An interface definition for the [NWParameters] class.
//
// # Methods
//
//   - [INWParameters.SSLCipherSuites]
//   - [INWParameters.SetSSLCipherSuites]
//   - [INWParameters.SSLCipherSuitesInternal]
//   - [INWParameters.SetSSLCipherSuitesInternal]
//   - [INWParameters.TLSSessionID]
//   - [INWParameters.SetTLSSessionID]
//   - [INWParameters.Account]
//   - [INWParameters.SetAccount]
//   - [INWParameters.AllowDuplicateStateUpdates]
//   - [INWParameters.SetAllowDuplicateStateUpdates]
//   - [INWParameters.AllowJoiningConnectedFd]
//   - [INWParameters.SetAllowJoiningConnectedFd]
//   - [INWParameters.AllowSocketAccess]
//   - [INWParameters.SetAllowSocketAccess]
//   - [INWParameters.AllowUnusableAddresses]
//   - [INWParameters.SetAllowUnusableAddresses]
//   - [INWParameters.AttachProtocolListener]
//   - [INWParameters.SetAttachProtocolListener]
//   - [INWParameters.AvoidNetworkAgentWithDomainType]
//   - [INWParameters.AvoidNetworkAgentWithUUID]
//   - [INWParameters.CopyCParameters]
//   - [INWParameters.CopyRequiredAgentsDescription]
//   - [INWParameters.CreateProtocolBufferObject]
//   - [INWParameters.DataMode]
//   - [INWParameters.SetDataMode]
//   - [INWParameters.DescriptionWithIndentShowFullContent]
//   - [INWParameters.DisableNagleAlgorithm]
//   - [INWParameters.SetDisableNagleAlgorithm]
//   - [INWParameters.EffectiveBundleID]
//   - [INWParameters.SetEffectiveBundleID]
//   - [INWParameters.EffectiveProcessUUID]
//   - [INWParameters.SetEffectiveProcessUUID]
//   - [INWParameters.EnableTFO]
//   - [INWParameters.SetEnableTFO]
//   - [INWParameters.EnableTFONoCookie]
//   - [INWParameters.SetEnableTFONoCookie]
//   - [INWParameters.EnableTLS]
//   - [INWParameters.SetEnableTLS]
//   - [INWParameters.EncodeWithCoder]
//   - [INWParameters.FastOpenForceEnable]
//   - [INWParameters.SetFastOpenForceEnable]
//   - [INWParameters.HasDelegatedPIDForOriginatingPID]
//   - [INWParameters.HasDelegatedProcessUUID]
//   - [INWParameters.HasNonEmptyProxyConfiguration]
//   - [INWParameters.HasPreferredNetworkAgents]
//   - [INWParameters.HasProhibitedNetworkAgents]
//   - [INWParameters.HasRequiredNetworkAgents]
//   - [INWParameters.HttpsProxyIsOpaque]
//   - [INWParameters.SetHttpsProxyIsOpaque]
//   - [INWParameters.HttpsProxyOverTLS]
//   - [INWParameters.SetHttpsProxyOverTLS]
//   - [INWParameters.IgnoreResolverStats]
//   - [INWParameters.SetIgnoreResolverStats]
//   - [INWParameters.Indefinite]
//   - [INWParameters.SetIndefinite]
//   - [INWParameters.InternalParameters]
//   - [INWParameters.SetInternalParameters]
//   - [INWParameters.IpProtocol]
//   - [INWParameters.IsDiscretionary]
//   - [INWParameters.IsDryRun]
//   - [INWParameters.IsValid]
//   - [INWParameters.KeepAlive]
//   - [INWParameters.SetKeepAlive]
//   - [INWParameters.KeepAliveIdleTime]
//   - [INWParameters.SetKeepAliveIdleTime]
//   - [INWParameters.KeepAliveInterval]
//   - [INWParameters.SetKeepAliveInterval]
//   - [INWParameters.KeepAliveOffload]
//   - [INWParameters.SetKeepAliveOffload]
//   - [INWParameters.LocalAddress]
//   - [INWParameters.SetLocalAddress]
//   - [INWParameters.MaximumSSLProtocolVersion]
//   - [INWParameters.SetMaximumSSLProtocolVersion]
//   - [INWParameters.MaximumSSLProtocolVersionInternal]
//   - [INWParameters.SetMaximumSSLProtocolVersionInternal]
//   - [INWParameters.Metadata]
//   - [INWParameters.SetMetadata]
//   - [INWParameters.MinimumSSLProtocolVersion]
//   - [INWParameters.SetMinimumSSLProtocolVersion]
//   - [INWParameters.MinimumSSLProtocolVersionInternal]
//   - [INWParameters.SetMinimumSSLProtocolVersionInternal]
//   - [INWParameters.Multipath]
//   - [INWParameters.SetMultipath]
//   - [INWParameters.MultipathForceEnable]
//   - [INWParameters.SetMultipathForceEnable]
//   - [INWParameters.MultipathService]
//   - [INWParameters.SetMultipathService]
//   - [INWParameters.NoProxy]
//   - [INWParameters.SetNoProxy]
//   - [INWParameters.ParentID]
//   - [INWParameters.SetParentID]
//   - [INWParameters.ParentIDs]
//   - [INWParameters.Pid]
//   - [INWParameters.SetPid]
//   - [INWParameters.PreferNetworkAgentWithDomainType]
//   - [INWParameters.PreferNetworkAgentWithUUID]
//   - [INWParameters.PreferNoProxy]
//   - [INWParameters.SetPreferNoProxy]
//   - [INWParameters.PrivateDescription]
//   - [INWParameters.ProcessUUID]
//   - [INWParameters.SetProcessUUID]
//   - [INWParameters.ProhibitCellular]
//   - [INWParameters.ProhibitConstrainedPaths]
//   - [INWParameters.SetProhibitConstrainedPaths]
//   - [INWParameters.ProhibitExpensivePaths]
//   - [INWParameters.SetProhibitExpensivePaths]
//   - [INWParameters.ProhibitFallback]
//   - [INWParameters.SetProhibitFallback]
//   - [INWParameters.ProhibitInterface]
//   - [INWParameters.ProhibitInterfaceSubtype]
//   - [INWParameters.ProhibitInterfaceType]
//   - [INWParameters.ProhibitJoiningProtocols]
//   - [INWParameters.SetProhibitJoiningProtocols]
//   - [INWParameters.ProhibitNetworkAgentWithUUID]
//   - [INWParameters.ProhibitNetworkAgentsWithDomainType]
//   - [INWParameters.ProhibitRoaming]
//   - [INWParameters.SetProhibitRoaming]
//   - [INWParameters.ProtocolTransforms]
//   - [INWParameters.SetProtocolTransforms]
//   - [INWParameters.ProxyConfiguration]
//   - [INWParameters.SetProxyConfiguration]
//   - [INWParameters.ReduceBuffering]
//   - [INWParameters.SetReduceBuffering]
//   - [INWParameters.RequireNetworkAgentWithDomainType]
//   - [INWParameters.RequireNetworkAgentWithUUID]
//   - [INWParameters.RequiredAddressFamily]
//   - [INWParameters.SetRequiredAddressFamily]
//   - [INWParameters.RequiredCompanionProxyInterfaceType]
//   - [INWParameters.SetRequiredCompanionProxyInterfaceType]
//   - [INWParameters.RequiredInterface]
//   - [INWParameters.SetRequiredInterface]
//   - [INWParameters.RequiredInterfaceSubtype]
//   - [INWParameters.SetRequiredInterfaceSubtype]
//   - [INWParameters.RequiredInterfaceType]
//   - [INWParameters.SetRequiredInterfaceType]
//   - [INWParameters.ResolvePTR]
//   - [INWParameters.SetResolvePTR]
//   - [INWParameters.ReuseLocalAddress]
//   - [INWParameters.SetReuseLocalAddress]
//   - [INWParameters.SanitizedURL]
//   - [INWParameters.SetInitialDataPayload]
//   - [INWParameters.SetSourceApplicationWithBundleID]
//   - [INWParameters.TlsVersionWithSSLProtocol]
//   - [INWParameters.TrafficClass]
//   - [INWParameters.SetTrafficClass]
//   - [INWParameters.TransportProtocol]
//   - [INWParameters.TrustInvalidCertificates]
//   - [INWParameters.SetTrustInvalidCertificates]
//   - [INWParameters.Uid]
//   - [INWParameters.SetUid]
//   - [INWParameters.Url]
//   - [INWParameters.SetUrl]
//   - [INWParameters.UseAWDL]
//   - [INWParameters.SetUseAWDL]
//   - [INWParameters.UseLongOutstandingQueries]
//   - [INWParameters.SetUseLongOutstandingQueries]
//   - [INWParameters.UseP2P]
//   - [INWParameters.SetUseP2P]
//   - [INWParameters.InitWithCoder]
//   - [INWParameters.InitWithParameters]
//   - [INWParameters.Discretionary]
//   - [INWParameters.SetDiscretionary]
//   - [INWParameters.DryRun]
//   - [INWParameters.Valid]
type INWParameters interface {
	objectivec.IObject

	// Topic: Methods

	SSLCipherSuites() foundation.INSSet
	SetSSLCipherSuites(value foundation.INSSet)
	SSLCipherSuitesInternal() foundation.INSSet
	SetSSLCipherSuitesInternal(value foundation.INSSet)
	TLSSessionID() foundation.NSData
	SetTLSSessionID(value foundation.NSData)
	Account() string
	SetAccount(value string)
	AllowDuplicateStateUpdates() bool
	SetAllowDuplicateStateUpdates(value bool)
	AllowJoiningConnectedFd() bool
	SetAllowJoiningConnectedFd(value bool)
	AllowSocketAccess() bool
	SetAllowSocketAccess(value bool)
	AllowUnusableAddresses() bool
	SetAllowUnusableAddresses(value bool)
	AttachProtocolListener() bool
	SetAttachProtocolListener(value bool)
	AvoidNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject)
	AvoidNetworkAgentWithUUID(uuid objectivec.IObject)
	CopyCParameters() objectivec.IObject
	CopyRequiredAgentsDescription() objectivec.IObject
	CreateProtocolBufferObject() objectivec.IObject
	DataMode() uint64
	SetDataMode(value uint64)
	DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject
	DisableNagleAlgorithm() bool
	SetDisableNagleAlgorithm(value bool)
	EffectiveBundleID() string
	SetEffectiveBundleID(value string)
	EffectiveProcessUUID() foundation.NSUUID
	SetEffectiveProcessUUID(value foundation.NSUUID)
	EnableTFO() bool
	SetEnableTFO(value bool)
	EnableTFONoCookie() bool
	SetEnableTFONoCookie(value bool)
	EnableTLS() bool
	SetEnableTLS(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
	FastOpenForceEnable() bool
	SetFastOpenForceEnable(value bool)
	HasDelegatedPIDForOriginatingPID(pid int) bool
	HasDelegatedProcessUUID() bool
	HasNonEmptyProxyConfiguration() bool
	HasPreferredNetworkAgents() bool
	HasProhibitedNetworkAgents() bool
	HasRequiredNetworkAgents() bool
	HttpsProxyIsOpaque() bool
	SetHttpsProxyIsOpaque(value bool)
	HttpsProxyOverTLS() bool
	SetHttpsProxyOverTLS(value bool)
	IgnoreResolverStats() bool
	SetIgnoreResolverStats(value bool)
	Indefinite() bool
	SetIndefinite(value bool)
	InternalParameters() objectivec.Object
	SetInternalParameters(value objectivec.Object)
	IpProtocol() byte
	IsDiscretionary() bool
	IsDryRun() bool
	IsValid() bool
	KeepAlive() bool
	SetKeepAlive(value bool)
	KeepAliveIdleTime() uint64
	SetKeepAliveIdleTime(value uint64)
	KeepAliveInterval() uint64
	SetKeepAliveInterval(value uint64)
	KeepAliveOffload() bool
	SetKeepAliveOffload(value bool)
	LocalAddress() unsafe.Pointer
	SetLocalAddress(value unsafe.Pointer)
	MaximumSSLProtocolVersion() uint64
	SetMaximumSSLProtocolVersion(value uint64)
	MaximumSSLProtocolVersionInternal() uint64
	SetMaximumSSLProtocolVersionInternal(value uint64)
	Metadata() foundation.NSData
	SetMetadata(value foundation.NSData)
	MinimumSSLProtocolVersion() uint64
	SetMinimumSSLProtocolVersion(value uint64)
	MinimumSSLProtocolVersionInternal() uint64
	SetMinimumSSLProtocolVersionInternal(value uint64)
	Multipath() bool
	SetMultipath(value bool)
	MultipathForceEnable() bool
	SetMultipathForceEnable(value bool)
	MultipathService() int
	SetMultipathService(value int)
	NoProxy() bool
	SetNoProxy(value bool)
	ParentID() foundation.NSUUID
	SetParentID(value foundation.NSUUID)
	ParentIDs() foundation.INSArray
	Pid() int
	SetPid(value int)
	PreferNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject)
	PreferNetworkAgentWithUUID(uuid objectivec.IObject)
	PreferNoProxy() bool
	SetPreferNoProxy(value bool)
	PrivateDescription() string
	ProcessUUID() foundation.NSUUID
	SetProcessUUID(value foundation.NSUUID)
	ProhibitCellular() bool
	ProhibitConstrainedPaths() bool
	SetProhibitConstrainedPaths(value bool)
	ProhibitExpensivePaths() bool
	SetProhibitExpensivePaths(value bool)
	ProhibitFallback() bool
	SetProhibitFallback(value bool)
	ProhibitInterface(interface_ objectivec.IObject)
	ProhibitInterfaceSubtype(subtype int64)
	ProhibitInterfaceType(type_ int64)
	ProhibitJoiningProtocols() bool
	SetProhibitJoiningProtocols(value bool)
	ProhibitNetworkAgentWithUUID(uuid objectivec.IObject)
	ProhibitNetworkAgentsWithDomainType(domain objectivec.IObject, type_ objectivec.IObject)
	ProhibitRoaming() bool
	SetProhibitRoaming(value bool)
	ProtocolTransforms() foundation.INSArray
	SetProtocolTransforms(value foundation.INSArray)
	ProxyConfiguration() foundation.INSDictionary
	SetProxyConfiguration(value foundation.INSDictionary)
	ReduceBuffering() bool
	SetReduceBuffering(value bool)
	RequireNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject)
	RequireNetworkAgentWithUUID(uuid objectivec.IObject)
	RequiredAddressFamily() byte
	SetRequiredAddressFamily(value byte)
	RequiredCompanionProxyInterfaceType() int64
	SetRequiredCompanionProxyInterfaceType(value int64)
	RequiredInterface() INWInterface
	SetRequiredInterface(value INWInterface)
	RequiredInterfaceSubtype() int64
	SetRequiredInterfaceSubtype(value int64)
	RequiredInterfaceType() int64
	SetRequiredInterfaceType(value int64)
	ResolvePTR() bool
	SetResolvePTR(value bool)
	ReuseLocalAddress() bool
	SetReuseLocalAddress(value bool)
	SanitizedURL() foundation.NSURL
	SetInitialDataPayload(payload objectivec.IObject)
	SetSourceApplicationWithBundleID(id objectivec.IObject)
	TlsVersionWithSSLProtocol(sSLProtocol int) uint16
	TrafficClass() uint64
	SetTrafficClass(value uint64)
	TransportProtocol() byte
	TrustInvalidCertificates() bool
	SetTrustInvalidCertificates(value bool)
	Uid() uint32
	SetUid(value uint32)
	Url() foundation.NSURL
	SetUrl(value foundation.NSURL)
	UseAWDL() bool
	SetUseAWDL(value bool)
	UseLongOutstandingQueries() bool
	SetUseLongOutstandingQueries(value bool)
	UseP2P() bool
	SetUseP2P(value bool)
	InitWithCoder(coder foundation.INSCoder) NWParameters
	InitWithParameters(parameters objectivec.IObject) NWParameters
	Discretionary() bool
	SetDiscretionary(value bool)
	DryRun() bool
	Valid() bool
}

// Init initializes the instance.
func (n NWParameters) Init() NWParameters {
	rv := objc.Send[NWParameters](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWParameters) Autorelease() NWParameters {
	rv := objc.Send[NWParameters](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWParameters creates a new NWParameters instance.
func NewNWParameters() NWParameters {
	class := getNWParametersClass()
	rv := objc.Send[NWParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWParametersWithCoder(coder objectivec.IObject) NWParameters {
	instance := getNWParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NWParametersFromID(rv)
}

func NewNWParametersWithParameters(parameters objectivec.IObject) NWParameters {
	instance := getNWParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return NWParametersFromID(rv)
}

func (n NWParameters) AvoidNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("avoidNetworkAgentWithDomain:type:"), domain, type_)
}
func (n NWParameters) AvoidNetworkAgentWithUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("avoidNetworkAgentWithUUID:"), uuid)
}
func (n NWParameters) CopyCParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyCParameters"))
	return objectivec.Object{ID: rv}
}
func (n NWParameters) CopyRequiredAgentsDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyRequiredAgentsDescription"))
	return objectivec.Object{ID: rv}
}
func (n NWParameters) CreateProtocolBufferObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createProtocolBufferObject"))
	return objectivec.Object{ID: rv}
}
func (n NWParameters) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}
func (n NWParameters) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (n NWParameters) HasDelegatedPIDForOriginatingPID(pid int) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasDelegatedPIDForOriginatingPID:"), pid)
	return rv
}
func (n NWParameters) HasDelegatedProcessUUID() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasDelegatedProcessUUID"))
	return rv
}
func (n NWParameters) HasNonEmptyProxyConfiguration() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasNonEmptyProxyConfiguration"))
	return rv
}
func (n NWParameters) HasPreferredNetworkAgents() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasPreferredNetworkAgents"))
	return rv
}
func (n NWParameters) HasProhibitedNetworkAgents() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasProhibitedNetworkAgents"))
	return rv
}
func (n NWParameters) HasRequiredNetworkAgents() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasRequiredNetworkAgents"))
	return rv
}
func (n NWParameters) IsDiscretionary() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isDiscretionary"))
	return rv
}
func (n NWParameters) IsDryRun() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isDryRun"))
	return rv
}
func (n NWParameters) IsValid() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isValid"))
	return rv
}
func (n NWParameters) PreferNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("preferNetworkAgentWithDomain:type:"), domain, type_)
}
func (n NWParameters) PreferNetworkAgentWithUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("preferNetworkAgentWithUUID:"), uuid)
}
func (n NWParameters) ProhibitInterface(interface_ objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("prohibitInterface:"), interface_)
}
func (n NWParameters) ProhibitInterfaceSubtype(subtype int64) {
	objc.Send[objc.ID](n.ID, objc.Sel("prohibitInterfaceSubtype:"), subtype)
}
func (n NWParameters) ProhibitInterfaceType(type_ int64) {
	objc.Send[objc.ID](n.ID, objc.Sel("prohibitInterfaceType:"), type_)
}
func (n NWParameters) ProhibitNetworkAgentWithUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("prohibitNetworkAgentWithUUID:"), uuid)
}
func (n NWParameters) ProhibitNetworkAgentsWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("prohibitNetworkAgentsWithDomain:type:"), domain, type_)
}
func (n NWParameters) RequireNetworkAgentWithDomainType(domain objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("requireNetworkAgentWithDomain:type:"), domain, type_)
}
func (n NWParameters) RequireNetworkAgentWithUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("requireNetworkAgentWithUUID:"), uuid)
}
func (n NWParameters) SetInitialDataPayload(payload objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("setInitialDataPayload:"), payload)
}
func (n NWParameters) SetSourceApplicationWithBundleID(id objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("setSourceApplicationWithBundleID:"), id)
}
func (n NWParameters) TlsVersionWithSSLProtocol(sSLProtocol int) uint16 {
	rv := objc.Send[uint16](n.ID, objc.Sel("tlsVersionWithSSLProtocol:"), sSLProtocol)
	return rv
}
func (n NWParameters) InitWithCoder(coder foundation.INSCoder) NWParameters {
	rv := objc.Send[NWParameters](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (n NWParameters) InitWithParameters(parameters objectivec.IObject) NWParameters {
	rv := objc.Send[NWParameters](n.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (_NWParametersClass NWParametersClass) ParametersWithCParameters(cParameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWParametersClass.class), objc.Sel("parametersWithCParameters:"), cParameters)
	return objectivec.Object{ID: rv}
}
func (_NWParametersClass NWParametersClass) ParametersWithProtocolBufferData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWParametersClass.class), objc.Sel("parametersWithProtocolBufferData:"), data)
	return objectivec.Object{ID: rv}
}
func (_NWParametersClass NWParametersClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_NWParametersClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (n NWParameters) SSLCipherSuites() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("SSLCipherSuites"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n NWParameters) SetSSLCipherSuites(value foundation.INSSet) {
	objc.Send[struct{}](n.ID, objc.Sel("setSSLCipherSuites:"), value)
}
func (n NWParameters) SSLCipherSuitesInternal() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("SSLCipherSuitesInternal"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n NWParameters) SetSSLCipherSuitesInternal(value foundation.INSSet) {
	objc.Send[struct{}](n.ID, objc.Sel("setSSLCipherSuitesInternal:"), value)
}
func (n NWParameters) TLSSessionID() foundation.NSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("TLSSessionID"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n NWParameters) SetTLSSessionID(value foundation.NSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setTLSSessionID:"), value)
}
func (n NWParameters) Account() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("account"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWParameters) SetAccount(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccount:"), objc.String(value))
}
func (n NWParameters) AllowDuplicateStateUpdates() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("allowDuplicateStateUpdates"))
	return rv
}
func (n NWParameters) SetAllowDuplicateStateUpdates(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAllowDuplicateStateUpdates:"), value)
}
func (n NWParameters) AllowJoiningConnectedFd() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("allowJoiningConnectedFd"))
	return rv
}
func (n NWParameters) SetAllowJoiningConnectedFd(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAllowJoiningConnectedFd:"), value)
}
func (n NWParameters) AllowSocketAccess() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("allowSocketAccess"))
	return rv
}
func (n NWParameters) SetAllowSocketAccess(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAllowSocketAccess:"), value)
}
func (n NWParameters) AllowUnusableAddresses() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("allowUnusableAddresses"))
	return rv
}
func (n NWParameters) SetAllowUnusableAddresses(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAllowUnusableAddresses:"), value)
}
func (n NWParameters) AttachProtocolListener() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("attachProtocolListener"))
	return rv
}
func (n NWParameters) SetAttachProtocolListener(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttachProtocolListener:"), value)
}
func (n NWParameters) DataMode() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("dataMode"))
	return rv
}
func (n NWParameters) SetDataMode(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setDataMode:"), value)
}
func (n NWParameters) DisableNagleAlgorithm() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("disableNagleAlgorithm"))
	return rv
}
func (n NWParameters) SetDisableNagleAlgorithm(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setDisableNagleAlgorithm:"), value)
}
func (n NWParameters) Discretionary() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("discretionary"))
	return rv
}
func (n NWParameters) SetDiscretionary(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setDiscretionary:"), value)
}
func (n NWParameters) DryRun() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("dryRun"))
	return rv
}
func (n NWParameters) EffectiveBundleID() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("effectiveBundleID"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWParameters) SetEffectiveBundleID(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setEffectiveBundleID:"), objc.String(value))
}
func (n NWParameters) EffectiveProcessUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("effectiveProcessUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (n NWParameters) SetEffectiveProcessUUID(value foundation.NSUUID) {
	objc.Send[struct{}](n.ID, objc.Sel("setEffectiveProcessUUID:"), value)
}
func (n NWParameters) EnableTFO() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("enableTFO"))
	return rv
}
func (n NWParameters) SetEnableTFO(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setEnableTFO:"), value)
}
func (n NWParameters) EnableTFONoCookie() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("enableTFONoCookie"))
	return rv
}
func (n NWParameters) SetEnableTFONoCookie(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setEnableTFONoCookie:"), value)
}
func (n NWParameters) EnableTLS() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("enableTLS"))
	return rv
}
func (n NWParameters) SetEnableTLS(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setEnableTLS:"), value)
}
func (n NWParameters) FastOpenForceEnable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("fastOpenForceEnable"))
	return rv
}
func (n NWParameters) SetFastOpenForceEnable(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setFastOpenForceEnable:"), value)
}
func (n NWParameters) HttpsProxyIsOpaque() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("httpsProxyIsOpaque"))
	return rv
}
func (n NWParameters) SetHttpsProxyIsOpaque(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHttpsProxyIsOpaque:"), value)
}
func (n NWParameters) HttpsProxyOverTLS() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("httpsProxyOverTLS"))
	return rv
}
func (n NWParameters) SetHttpsProxyOverTLS(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHttpsProxyOverTLS:"), value)
}
func (n NWParameters) IgnoreResolverStats() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("ignoreResolverStats"))
	return rv
}
func (n NWParameters) SetIgnoreResolverStats(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setIgnoreResolverStats:"), value)
}
func (n NWParameters) Indefinite() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("indefinite"))
	return rv
}
func (n NWParameters) SetIndefinite(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setIndefinite:"), value)
}
func (n NWParameters) InternalParameters() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalParameters"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n NWParameters) SetInternalParameters(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalParameters:"), value)
}
func (n NWParameters) IpProtocol() byte {
	rv := objc.Send[byte](n.ID, objc.Sel("ipProtocol"))
	return rv
}
func (n NWParameters) KeepAlive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("keepAlive"))
	return rv
}
func (n NWParameters) SetKeepAlive(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setKeepAlive:"), value)
}
func (n NWParameters) KeepAliveIdleTime() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("keepAliveIdleTime"))
	return rv
}
func (n NWParameters) SetKeepAliveIdleTime(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setKeepAliveIdleTime:"), value)
}
func (n NWParameters) KeepAliveInterval() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("keepAliveInterval"))
	return rv
}
func (n NWParameters) SetKeepAliveInterval(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setKeepAliveInterval:"), value)
}
func (n NWParameters) KeepAliveOffload() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("keepAliveOffload"))
	return rv
}
func (n NWParameters) SetKeepAliveOffload(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setKeepAliveOffload:"), value)
}
func (n NWParameters) LocalAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("localAddress"))
	return rv
}
func (n NWParameters) SetLocalAddress(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setLocalAddress:"), value)
}
func (n NWParameters) MaximumSSLProtocolVersion() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("maximumSSLProtocolVersion"))
	return rv
}
func (n NWParameters) SetMaximumSSLProtocolVersion(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumSSLProtocolVersion:"), value)
}
func (n NWParameters) MaximumSSLProtocolVersionInternal() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("maximumSSLProtocolVersionInternal"))
	return rv
}
func (n NWParameters) SetMaximumSSLProtocolVersionInternal(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumSSLProtocolVersionInternal:"), value)
}
func (n NWParameters) Metadata() foundation.NSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("metadata"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n NWParameters) SetMetadata(value foundation.NSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setMetadata:"), value)
}
func (n NWParameters) MinimumSSLProtocolVersion() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("minimumSSLProtocolVersion"))
	return rv
}
func (n NWParameters) SetMinimumSSLProtocolVersion(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumSSLProtocolVersion:"), value)
}
func (n NWParameters) MinimumSSLProtocolVersionInternal() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("minimumSSLProtocolVersionInternal"))
	return rv
}
func (n NWParameters) SetMinimumSSLProtocolVersionInternal(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumSSLProtocolVersionInternal:"), value)
}
func (n NWParameters) Multipath() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("multipath"))
	return rv
}
func (n NWParameters) SetMultipath(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setMultipath:"), value)
}
func (n NWParameters) MultipathForceEnable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("multipathForceEnable"))
	return rv
}
func (n NWParameters) SetMultipathForceEnable(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setMultipathForceEnable:"), value)
}
func (n NWParameters) MultipathService() int {
	rv := objc.Send[int](n.ID, objc.Sel("multipathService"))
	return rv
}
func (n NWParameters) SetMultipathService(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setMultipathService:"), value)
}
func (n NWParameters) NoProxy() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("noProxy"))
	return rv
}
func (n NWParameters) SetNoProxy(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setNoProxy:"), value)
}
func (n NWParameters) ParentID() foundation.NSUUID {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parentID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (n NWParameters) SetParentID(value foundation.NSUUID) {
	objc.Send[struct{}](n.ID, objc.Sel("setParentID:"), value)
}
func (n NWParameters) ParentIDs() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parentIDs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n NWParameters) Pid() int {
	rv := objc.Send[int](n.ID, objc.Sel("pid"))
	return rv
}
func (n NWParameters) SetPid(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setPid:"), value)
}
func (n NWParameters) PreferNoProxy() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("preferNoProxy"))
	return rv
}
func (n NWParameters) SetPreferNoProxy(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setPreferNoProxy:"), value)
}
func (n NWParameters) PrivateDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("privateDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (n NWParameters) ProcessUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("processUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (n NWParameters) SetProcessUUID(value foundation.NSUUID) {
	objc.Send[struct{}](n.ID, objc.Sel("setProcessUUID:"), value)
}
func (n NWParameters) ProhibitCellular() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitCellular"))
	return rv
}
func (n NWParameters) ProhibitConstrainedPaths() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitConstrainedPaths"))
	return rv
}
func (n NWParameters) SetProhibitConstrainedPaths(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setProhibitConstrainedPaths:"), value)
}
func (n NWParameters) ProhibitExpensivePaths() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitExpensivePaths"))
	return rv
}
func (n NWParameters) SetProhibitExpensivePaths(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setProhibitExpensivePaths:"), value)
}
func (n NWParameters) ProhibitFallback() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitFallback"))
	return rv
}
func (n NWParameters) SetProhibitFallback(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setProhibitFallback:"), value)
}
func (n NWParameters) ProhibitJoiningProtocols() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitJoiningProtocols"))
	return rv
}
func (n NWParameters) SetProhibitJoiningProtocols(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setProhibitJoiningProtocols:"), value)
}
func (n NWParameters) ProhibitRoaming() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("prohibitRoaming"))
	return rv
}
func (n NWParameters) SetProhibitRoaming(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setProhibitRoaming:"), value)
}
func (n NWParameters) ProtocolTransforms() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("protocolTransforms"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n NWParameters) SetProtocolTransforms(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setProtocolTransforms:"), value)
}
func (n NWParameters) ProxyConfiguration() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("proxyConfiguration"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n NWParameters) SetProxyConfiguration(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setProxyConfiguration:"), value)
}
func (n NWParameters) ReduceBuffering() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("reduceBuffering"))
	return rv
}
func (n NWParameters) SetReduceBuffering(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setReduceBuffering:"), value)
}
func (n NWParameters) RequiredAddressFamily() byte {
	rv := objc.Send[byte](n.ID, objc.Sel("requiredAddressFamily"))
	return rv
}
func (n NWParameters) SetRequiredAddressFamily(value byte) {
	objc.Send[struct{}](n.ID, objc.Sel("setRequiredAddressFamily:"), value)
}
func (n NWParameters) RequiredCompanionProxyInterfaceType() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("requiredCompanionProxyInterfaceType"))
	return rv
}
func (n NWParameters) SetRequiredCompanionProxyInterfaceType(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setRequiredCompanionProxyInterfaceType:"), value)
}
func (n NWParameters) RequiredInterface() INWInterface {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("requiredInterface"))
	return NWInterfaceFromID(objc.ID(rv))
}
func (n NWParameters) SetRequiredInterface(value INWInterface) {
	objc.Send[struct{}](n.ID, objc.Sel("setRequiredInterface:"), value)
}
func (n NWParameters) RequiredInterfaceSubtype() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("requiredInterfaceSubtype"))
	return rv
}
func (n NWParameters) SetRequiredInterfaceSubtype(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setRequiredInterfaceSubtype:"), value)
}
func (n NWParameters) RequiredInterfaceType() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("requiredInterfaceType"))
	return rv
}
func (n NWParameters) SetRequiredInterfaceType(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setRequiredInterfaceType:"), value)
}
func (n NWParameters) ResolvePTR() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("resolvePTR"))
	return rv
}
func (n NWParameters) SetResolvePTR(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setResolvePTR:"), value)
}
func (n NWParameters) ReuseLocalAddress() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("reuseLocalAddress"))
	return rv
}
func (n NWParameters) SetReuseLocalAddress(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setReuseLocalAddress:"), value)
}
func (n NWParameters) SanitizedURL() foundation.NSURL {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("sanitizedURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (n NWParameters) TrafficClass() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("trafficClass"))
	return rv
}
func (n NWParameters) SetTrafficClass(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setTrafficClass:"), value)
}
func (n NWParameters) TransportProtocol() byte {
	rv := objc.Send[byte](n.ID, objc.Sel("transportProtocol"))
	return rv
}
func (n NWParameters) TrustInvalidCertificates() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("trustInvalidCertificates"))
	return rv
}
func (n NWParameters) SetTrustInvalidCertificates(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setTrustInvalidCertificates:"), value)
}
func (n NWParameters) Uid() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("uid"))
	return rv
}
func (n NWParameters) SetUid(value uint32) {
	objc.Send[struct{}](n.ID, objc.Sel("setUid:"), value)
}
func (n NWParameters) Url() foundation.NSURL {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (n NWParameters) SetUrl(value foundation.NSURL) {
	objc.Send[struct{}](n.ID, objc.Sel("setUrl:"), value)
}
func (n NWParameters) UseAWDL() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("useAWDL"))
	return rv
}
func (n NWParameters) SetUseAWDL(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUseAWDL:"), value)
}
func (n NWParameters) UseLongOutstandingQueries() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("useLongOutstandingQueries"))
	return rv
}
func (n NWParameters) SetUseLongOutstandingQueries(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUseLongOutstandingQueries:"), value)
}
func (n NWParameters) UseP2P() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("useP2P"))
	return rv
}
func (n NWParameters) SetUseP2P(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUseP2P:"), value)
}
func (n NWParameters) Valid() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("valid"))
	return rv
}
