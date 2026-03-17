// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingLayerSupportInfo] class.
var (
	_EspressoProfilingLayerSupportInfoClass     EspressoProfilingLayerSupportInfoClass
	_EspressoProfilingLayerSupportInfoClassOnce sync.Once
)

func getEspressoProfilingLayerSupportInfoClass() EspressoProfilingLayerSupportInfoClass {
	_EspressoProfilingLayerSupportInfoClassOnce.Do(func() {
		_EspressoProfilingLayerSupportInfoClass = EspressoProfilingLayerSupportInfoClass{class: objc.GetClass("EspressoProfilingLayerSupportInfo")}
	})
	return _EspressoProfilingLayerSupportInfoClass
}

// GetEspressoProfilingLayerSupportInfoClass returns the class object for EspressoProfilingLayerSupportInfo.
func GetEspressoProfilingLayerSupportInfoClass() EspressoProfilingLayerSupportInfoClass {
	return getEspressoProfilingLayerSupportInfoClass()
}

type EspressoProfilingLayerSupportInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingLayerSupportInfoClass) Alloc() EspressoProfilingLayerSupportInfo {
	rv := objc.Send[EspressoProfilingLayerSupportInfo](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoProfilingLayerSupportInfo.Error_private]
//   - [EspressoProfilingLayerSupportInfo.SetError_private]
//   - [EspressoProfilingLayerSupportInfo.Error_public]
//   - [EspressoProfilingLayerSupportInfo.SetError_public]
//   - [EspressoProfilingLayerSupportInfo.Exists]
//   - [EspressoProfilingLayerSupportInfo.SetExists]
//   - [EspressoProfilingLayerSupportInfo.Has_perf_warning]
//   - [EspressoProfilingLayerSupportInfo.SetHas_perf_warning]
//   - [EspressoProfilingLayerSupportInfo.Internal_layer]
//   - [EspressoProfilingLayerSupportInfo.SetInternal_layer]
//   - [EspressoProfilingLayerSupportInfo.Supported]
//   - [EspressoProfilingLayerSupportInfo.SetSupported]
//   - [EspressoProfilingLayerSupportInfo.Type]
//   - [EspressoProfilingLayerSupportInfo.SetType]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo
type EspressoProfilingLayerSupportInfo struct {
	objectivec.Object
}

// EspressoProfilingLayerSupportInfoFromID constructs a [EspressoProfilingLayerSupportInfo] from an objc.ID.
func EspressoProfilingLayerSupportInfoFromID(id objc.ID) EspressoProfilingLayerSupportInfo {
	return EspressoProfilingLayerSupportInfo{objectivec.Object{ID: id}}
}
// Ensure EspressoProfilingLayerSupportInfo implements IEspressoProfilingLayerSupportInfo.
var _ IEspressoProfilingLayerSupportInfo = EspressoProfilingLayerSupportInfo{}

// An interface definition for the [EspressoProfilingLayerSupportInfo] class.
//
// # Methods
//
//   - [IEspressoProfilingLayerSupportInfo.Error_private]
//   - [IEspressoProfilingLayerSupportInfo.SetError_private]
//   - [IEspressoProfilingLayerSupportInfo.Error_public]
//   - [IEspressoProfilingLayerSupportInfo.SetError_public]
//   - [IEspressoProfilingLayerSupportInfo.Exists]
//   - [IEspressoProfilingLayerSupportInfo.SetExists]
//   - [IEspressoProfilingLayerSupportInfo.Has_perf_warning]
//   - [IEspressoProfilingLayerSupportInfo.SetHas_perf_warning]
//   - [IEspressoProfilingLayerSupportInfo.Internal_layer]
//   - [IEspressoProfilingLayerSupportInfo.SetInternal_layer]
//   - [IEspressoProfilingLayerSupportInfo.Supported]
//   - [IEspressoProfilingLayerSupportInfo.SetSupported]
//   - [IEspressoProfilingLayerSupportInfo.Type]
//   - [IEspressoProfilingLayerSupportInfo.SetType]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo
type IEspressoProfilingLayerSupportInfo interface {
	objectivec.IObject

	// Topic: Methods

	Error_private() foundation.INSArray
	SetError_private(value foundation.INSArray)
	Error_public() foundation.INSArray
	SetError_public(value foundation.INSArray)
	Exists() bool
	SetExists(value bool)
	Has_perf_warning() bool
	SetHas_perf_warning(value bool)
	Internal_layer() bool
	SetInternal_layer(value bool)
	Supported() bool
	SetSupported(value bool)
	Type() string
	SetType(value string)
}

// Init initializes the instance.
func (e EspressoProfilingLayerSupportInfo) Init() EspressoProfilingLayerSupportInfo {
	rv := objc.Send[EspressoProfilingLayerSupportInfo](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingLayerSupportInfo) Autorelease() EspressoProfilingLayerSupportInfo {
	rv := objc.Send[EspressoProfilingLayerSupportInfo](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingLayerSupportInfo creates a new EspressoProfilingLayerSupportInfo instance.
func NewEspressoProfilingLayerSupportInfo() EspressoProfilingLayerSupportInfo {
	class := getEspressoProfilingLayerSupportInfoClass()
	rv := objc.Send[EspressoProfilingLayerSupportInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/error_private
func (e EspressoProfilingLayerSupportInfo) Error_private() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("error_private"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoProfilingLayerSupportInfo) SetError_private(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setError_private:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/error_public
func (e EspressoProfilingLayerSupportInfo) Error_public() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("error_public"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoProfilingLayerSupportInfo) SetError_public(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setError_public:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/exists
func (e EspressoProfilingLayerSupportInfo) Exists() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("exists"))
	return rv
}
func (e EspressoProfilingLayerSupportInfo) SetExists(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setExists:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/has_perf_warning
func (e EspressoProfilingLayerSupportInfo) Has_perf_warning() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("has_perf_warning"))
	return rv
}
func (e EspressoProfilingLayerSupportInfo) SetHas_perf_warning(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setHas_perf_warning:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/internal_layer
func (e EspressoProfilingLayerSupportInfo) Internal_layer() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("internal_layer"))
	return rv
}
func (e EspressoProfilingLayerSupportInfo) SetInternal_layer(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setInternal_layer:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/supported
func (e EspressoProfilingLayerSupportInfo) Supported() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("supported"))
	return rv
}
func (e EspressoProfilingLayerSupportInfo) SetSupported(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setSupported:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerSupportInfo/type
func (e EspressoProfilingLayerSupportInfo) Type() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoProfilingLayerSupportInfo) SetType(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setType:"), objc.String(value))
}

