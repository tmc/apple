// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingNetworkInfo] class.
var (
	_EspressoProfilingNetworkInfoClass     EspressoProfilingNetworkInfoClass
	_EspressoProfilingNetworkInfoClassOnce sync.Once
)

func getEspressoProfilingNetworkInfoClass() EspressoProfilingNetworkInfoClass {
	_EspressoProfilingNetworkInfoClassOnce.Do(func() {
		_EspressoProfilingNetworkInfoClass = EspressoProfilingNetworkInfoClass{class: objc.GetClass("EspressoProfilingNetworkInfo")}
	})
	return _EspressoProfilingNetworkInfoClass
}

// GetEspressoProfilingNetworkInfoClass returns the class object for EspressoProfilingNetworkInfo.
func GetEspressoProfilingNetworkInfoClass() EspressoProfilingNetworkInfoClass {
	return getEspressoProfilingNetworkInfoClass()
}

type EspressoProfilingNetworkInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoProfilingNetworkInfoClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingNetworkInfoClass) Alloc() EspressoProfilingNetworkInfo {
	rv := objc.Send[EspressoProfilingNetworkInfo](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoProfilingNetworkInfo.Ane_compiler_analytics]
//   - [EspressoProfilingNetworkInfo.SetAne_compiler_analytics]
//   - [EspressoProfilingNetworkInfo.Ane_performance_info]
//   - [EspressoProfilingNetworkInfo.SetAne_performance_info]
//   - [EspressoProfilingNetworkInfo.Layers]
//   - [EspressoProfilingNetworkInfo.SetLayers]
//   - [EspressoProfilingNetworkInfo.Network_at_path]
//   - [EspressoProfilingNetworkInfo.SetNetwork_at_path]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo
type EspressoProfilingNetworkInfo struct {
	objectivec.Object
}

// EspressoProfilingNetworkInfoFromID constructs a [EspressoProfilingNetworkInfo] from an objc.ID.
func EspressoProfilingNetworkInfoFromID(id objc.ID) EspressoProfilingNetworkInfo {
	return EspressoProfilingNetworkInfo{objectivec.Object{ID: id}}
}
// Ensure EspressoProfilingNetworkInfo implements IEspressoProfilingNetworkInfo.
var _ IEspressoProfilingNetworkInfo = EspressoProfilingNetworkInfo{}

// An interface definition for the [EspressoProfilingNetworkInfo] class.
//
// # Methods
//
//   - [IEspressoProfilingNetworkInfo.Ane_compiler_analytics]
//   - [IEspressoProfilingNetworkInfo.SetAne_compiler_analytics]
//   - [IEspressoProfilingNetworkInfo.Ane_performance_info]
//   - [IEspressoProfilingNetworkInfo.SetAne_performance_info]
//   - [IEspressoProfilingNetworkInfo.Layers]
//   - [IEspressoProfilingNetworkInfo.SetLayers]
//   - [IEspressoProfilingNetworkInfo.Network_at_path]
//   - [IEspressoProfilingNetworkInfo.SetNetwork_at_path]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo
type IEspressoProfilingNetworkInfo interface {
	objectivec.IObject

	// Topic: Methods

	Ane_compiler_analytics() IEspressoProfilingANEcompilerAnalytics
	SetAne_compiler_analytics(value IEspressoProfilingANEcompilerAnalytics)
	Ane_performance_info() IEspressoProfilingNetworkANEInfo
	SetAne_performance_info(value IEspressoProfilingNetworkANEInfo)
	Layers() foundation.INSArray
	SetLayers(value foundation.INSArray)
	Network_at_path() string
	SetNetwork_at_path(value string)
}

// Init initializes the instance.
func (e EspressoProfilingNetworkInfo) Init() EspressoProfilingNetworkInfo {
	rv := objc.Send[EspressoProfilingNetworkInfo](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingNetworkInfo) Autorelease() EspressoProfilingNetworkInfo {
	rv := objc.Send[EspressoProfilingNetworkInfo](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingNetworkInfo creates a new EspressoProfilingNetworkInfo instance.
func NewEspressoProfilingNetworkInfo() EspressoProfilingNetworkInfo {
	class := getEspressoProfilingNetworkInfoClass()
	rv := objc.Send[EspressoProfilingNetworkInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo/ane_compiler_analytics
func (e EspressoProfilingNetworkInfo) Ane_compiler_analytics() IEspressoProfilingANEcompilerAnalytics {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ane_compiler_analytics"))
	return EspressoProfilingANEcompilerAnalyticsFromID(objc.ID(rv))
}
func (e EspressoProfilingNetworkInfo) SetAne_compiler_analytics(value IEspressoProfilingANEcompilerAnalytics) {
	objc.Send[struct{}](e.ID, objc.Sel("setAne_compiler_analytics:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo/ane_performance_info
func (e EspressoProfilingNetworkInfo) Ane_performance_info() IEspressoProfilingNetworkANEInfo {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ane_performance_info"))
	return EspressoProfilingNetworkANEInfoFromID(objc.ID(rv))
}
func (e EspressoProfilingNetworkInfo) SetAne_performance_info(value IEspressoProfilingNetworkANEInfo) {
	objc.Send[struct{}](e.ID, objc.Sel("setAne_performance_info:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo/layers
func (e EspressoProfilingNetworkInfo) Layers() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoProfilingNetworkInfo) SetLayers(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setLayers:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkInfo/network_at_path
func (e EspressoProfilingNetworkInfo) Network_at_path() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("network_at_path"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoProfilingNetworkInfo) SetNetwork_at_path(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setNetwork_at_path:"), objc.String(value))
}

