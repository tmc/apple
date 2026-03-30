// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNetworkOptions] class.
var (
	_MLNetworkOptionsClass     MLNetworkOptionsClass
	_MLNetworkOptionsClassOnce sync.Once
)

func getMLNetworkOptionsClass() MLNetworkOptionsClass {
	_MLNetworkOptionsClassOnce.Do(func() {
		_MLNetworkOptionsClass = MLNetworkOptionsClass{class: objc.GetClass("_MLNetworkOptions")}
	})
	return _MLNetworkOptionsClass
}

// GetMLNetworkOptionsClass returns the class object for _MLNetworkOptions.
func GetMLNetworkOptionsClass() MLNetworkOptionsClass {
	return getMLNetworkOptionsClass()
}

type MLNetworkOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNetworkOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNetworkOptionsClass) Alloc() MLNetworkOptions {
	rv := objc.Send[MLNetworkOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNetworkOptions.Family]
//   - [MLNetworkOptions.LocalAddr]
//   - [MLNetworkOptions.LocalPort]
//   - [MLNetworkOptions.NetworkNameIdentifier]
//   - [MLNetworkOptions.NetworkOptions]
//   - [MLNetworkOptions.Port]
//   - [MLNetworkOptions.Psk]
//   - [MLNetworkOptions.ReceiveTimeout]
//   - [MLNetworkOptions.ReceiveTimeoutValue]
//   - [MLNetworkOptions.UseAWDL]
//   - [MLNetworkOptions.UseBonjour]
//   - [MLNetworkOptions.UseTLS]
//   - [MLNetworkOptions.UseUDP]
//   - [MLNetworkOptions.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions
type MLNetworkOptions struct {
	objectivec.Object
}

// MLNetworkOptionsFromID constructs a [MLNetworkOptions] from an objc.ID.
func MLNetworkOptionsFromID(id objc.ID) MLNetworkOptions {
	return MLNetworkOptions{objectivec.Object{ID: id}}
}

// Ensure MLNetworkOptions implements IMLNetworkOptions.
var _ IMLNetworkOptions = MLNetworkOptions{}

// An interface definition for the [MLNetworkOptions] class.
//
// # Methods
//
//   - [IMLNetworkOptions.Family]
//   - [IMLNetworkOptions.LocalAddr]
//   - [IMLNetworkOptions.LocalPort]
//   - [IMLNetworkOptions.NetworkNameIdentifier]
//   - [IMLNetworkOptions.NetworkOptions]
//   - [IMLNetworkOptions.Port]
//   - [IMLNetworkOptions.Psk]
//   - [IMLNetworkOptions.ReceiveTimeout]
//   - [IMLNetworkOptions.ReceiveTimeoutValue]
//   - [IMLNetworkOptions.UseAWDL]
//   - [IMLNetworkOptions.UseBonjour]
//   - [IMLNetworkOptions.UseTLS]
//   - [IMLNetworkOptions.UseUDP]
//   - [IMLNetworkOptions.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions
type IMLNetworkOptions interface {
	objectivec.IObject

	// Topic: Methods

	Family() uint64
	LocalAddr() string
	LocalPort() string
	NetworkNameIdentifier() string
	NetworkOptions() foundation.INSDictionary
	Port() string
	Psk() string
	ReceiveTimeout() int64
	ReceiveTimeoutValue() int64
	UseAWDL() bool
	UseBonjour() bool
	UseTLS() bool
	UseUDP() bool
	InitWithOptions(options objectivec.IObject) MLNetworkOptions
}

// Init initializes the instance.
func (m MLNetworkOptions) Init() MLNetworkOptions {
	rv := objc.Send[MLNetworkOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworkOptions) Autorelease() MLNetworkOptions {
	rv := objc.Send[MLNetworkOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworkOptions creates a new MLNetworkOptions instance.
func NewMLNetworkOptions() MLNetworkOptions {
	class := getMLNetworkOptionsClass()
	rv := objc.Send[MLNetworkOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/initWithOptions:
func NewMLNetworkOptionsWithOptions(options objectivec.IObject) MLNetworkOptions {
	instance := getMLNetworkOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return MLNetworkOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/family
func (m MLNetworkOptions) Family() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("family"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/localAddr
func (m MLNetworkOptions) LocalAddr() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("localAddr"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/localPort
func (m MLNetworkOptions) LocalPort() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("localPort"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/networkNameIdentifier
func (m MLNetworkOptions) NetworkNameIdentifier() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("networkNameIdentifier"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/port
func (m MLNetworkOptions) Port() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("port"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/psk
func (m MLNetworkOptions) Psk() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("psk"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/receiveTimeoutValue
func (m MLNetworkOptions) ReceiveTimeoutValue() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("receiveTimeoutValue"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/useAWDL
func (m MLNetworkOptions) UseAWDL() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useAWDL"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/useBonjour
func (m MLNetworkOptions) UseBonjour() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useBonjour"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/useTLS
func (m MLNetworkOptions) UseTLS() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useTLS"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/useUDP
func (m MLNetworkOptions) UseUDP() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useUDP"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/initWithOptions:
func (m MLNetworkOptions) InitWithOptions(options objectivec.IObject) MLNetworkOptions {
	rv := objc.Send[MLNetworkOptions](m.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/networkOptions
func (m MLNetworkOptions) NetworkOptions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("networkOptions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkOptions/receiveTimeout
func (m MLNetworkOptions) ReceiveTimeout() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("receiveTimeout"))
	return rv
}
