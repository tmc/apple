// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNetworkUtilities] class.
var (
	_MLNetworkUtilitiesClass     MLNetworkUtilitiesClass
	_MLNetworkUtilitiesClassOnce sync.Once
)

func getMLNetworkUtilitiesClass() MLNetworkUtilitiesClass {
	_MLNetworkUtilitiesClassOnce.Do(func() {
		_MLNetworkUtilitiesClass = MLNetworkUtilitiesClass{class: objc.GetClass("_MLNetworkUtilities")}
	})
	return _MLNetworkUtilitiesClass
}

// GetMLNetworkUtilitiesClass returns the class object for _MLNetworkUtilities.
func GetMLNetworkUtilitiesClass() MLNetworkUtilitiesClass {
	return getMLNetworkUtilitiesClass()
}

type MLNetworkUtilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNetworkUtilitiesClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNetworkUtilitiesClass) Alloc() MLNetworkUtilities {
	rv := objc.Send[MLNetworkUtilities](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities
type MLNetworkUtilities struct {
	objectivec.Object
}

// MLNetworkUtilitiesFromID constructs a [MLNetworkUtilities] from an objc.ID.
func MLNetworkUtilitiesFromID(id objc.ID) MLNetworkUtilities {
	return MLNetworkUtilities{objectivec.Object{ID: id}}
}
// Ensure MLNetworkUtilities implements IMLNetworkUtilities.
var _ IMLNetworkUtilities = MLNetworkUtilities{}

// An interface definition for the [MLNetworkUtilities] class.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities
type IMLNetworkUtilities interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLNetworkUtilities) Init() MLNetworkUtilities {
	rv := objc.Send[MLNetworkUtilities](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworkUtilities) Autorelease() MLNetworkUtilities {
	rv := objc.Send[MLNetworkUtilities](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworkUtilities creates a new MLNetworkUtilities instance.
func NewMLNetworkUtilities() MLNetworkUtilities {
	class := getMLNetworkUtilitiesClass()
	rv := objc.Send[MLNetworkUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/bindEndPoints:localAddr:localPort:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) BindEndPointsLocalAddrLocalPort(points objectivec.IObject, addr string, port string) {
	objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("bindEndPoints:localAddr:localPort:"), points, unsafe.Pointer(unsafe.StringData(addr + "\x00")), unsafe.Pointer(unsafe.StringData(port + "\x00")))
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/configureTLS:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) ConfigureTLS(tls objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("configureTLS:"), tls)
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/createSecureConnectionParameter:useUDP:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) CreateSecureConnectionParameterUseUDP(parameter VoidHandler, udp bool) objectivec.IObject {
_block0, _ := NewVoidBlock(parameter)
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("createSecureConnectionParameter:useUDP:"), _block0, udp)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/doInitNetwork:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) DoInitNetwork(network objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("doInitNetwork:"), network)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/setAWDL:useAWDL:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) SetAWDLUseAWDL(awdl objectivec.IObject, awdl2 bool) {
	objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("setAWDL:useAWDL:"), awdl, awdl2)
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/setProtocolStack:family:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) SetProtocolStackFamily(stack objectivec.IObject, family uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("setProtocolStack:family:"), stack, family)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/setupBonjour:name:useBonjour:useUDP:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) SetupBonjourNameUseBonjourUseUDP(bonjour objectivec.IObject, name string, bonjour2 bool, udp bool) {
	objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("setupBonjour:name:useBonjour:useUDP:"), bonjour, unsafe.Pointer(unsafe.StringData(name + "\x00")), bonjour2, udp)
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkUtilities/setupListenerStateChangeHandler:useUDP:
func (_MLNetworkUtilitiesClass MLNetworkUtilitiesClass) SetupListenerStateChangeHandlerUseUDP(handler ErrorHandler, udp bool) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLNetworkUtilitiesClass.class), objc.Sel("setupListenerStateChangeHandler:useUDP:"), _block0, udp)
}

