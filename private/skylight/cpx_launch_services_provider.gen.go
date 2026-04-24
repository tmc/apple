// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXLaunchServicesProvider] class.
var (
	_CPXLaunchServicesProviderClass     CPXLaunchServicesProviderClass
	_CPXLaunchServicesProviderClassOnce sync.Once
)

func getCPXLaunchServicesProviderClass() CPXLaunchServicesProviderClass {
	_CPXLaunchServicesProviderClassOnce.Do(func() {
		_CPXLaunchServicesProviderClass = CPXLaunchServicesProviderClass{class: objc.GetClass("CPXLaunchServicesProvider")}
	})
	return _CPXLaunchServicesProviderClass
}

// GetCPXLaunchServicesProviderClass returns the class object for CPXLaunchServicesProvider.
func GetCPXLaunchServicesProviderClass() CPXLaunchServicesProviderClass {
	return getCPXLaunchServicesProviderClass()
}

type CPXLaunchServicesProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXLaunchServicesProviderClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXLaunchServicesProviderClass) Alloc() CPXLaunchServicesProvider {
	rv := objc.Send[CPXLaunchServicesProvider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXLaunchServicesProvider._init]
//   - [CPXLaunchServicesProvider.LaunchServicesInterface]
//   - [CPXLaunchServicesProvider.SetLaunchServicesInterface]
//   - [CPXLaunchServicesProvider.DebugDescription]
//   - [CPXLaunchServicesProvider.Description]
//   - [CPXLaunchServicesProvider.Hash]
//   - [CPXLaunchServicesProvider.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider
type CPXLaunchServicesProvider struct {
	objectivec.Object
}

// CPXLaunchServicesProviderFromID constructs a [CPXLaunchServicesProvider] from an objc.ID.
func CPXLaunchServicesProviderFromID(id objc.ID) CPXLaunchServicesProvider {
	return CPXLaunchServicesProvider{objectivec.Object{ID: id}}
}

// Ensure CPXLaunchServicesProvider implements ICPXLaunchServicesProvider.
var _ ICPXLaunchServicesProvider = CPXLaunchServicesProvider{}

// An interface definition for the [CPXLaunchServicesProvider] class.
//
// # Methods
//
//   - [ICPXLaunchServicesProvider._init]
//   - [ICPXLaunchServicesProvider.LaunchServicesInterface]
//   - [ICPXLaunchServicesProvider.SetLaunchServicesInterface]
//   - [ICPXLaunchServicesProvider.DebugDescription]
//   - [ICPXLaunchServicesProvider.Description]
//   - [ICPXLaunchServicesProvider.Hash]
//   - [ICPXLaunchServicesProvider.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider
type ICPXLaunchServicesProvider interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	LaunchServicesInterface() ICPXLaunchServicesInterface
	SetLaunchServicesInterface(value ICPXLaunchServicesInterface)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXLaunchServicesProvider) Init() CPXLaunchServicesProvider {
	rv := objc.Send[CPXLaunchServicesProvider](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXLaunchServicesProvider) Autorelease() CPXLaunchServicesProvider {
	rv := objc.Send[CPXLaunchServicesProvider](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXLaunchServicesProvider creates a new CPXLaunchServicesProvider instance.
func NewCPXLaunchServicesProvider() CPXLaunchServicesProvider {
	class := getCPXLaunchServicesProviderClass()
	rv := objc.Send[CPXLaunchServicesProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/_init
func (c CPXLaunchServicesProvider) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/debugDescription
func (c CPXLaunchServicesProvider) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/description
func (c CPXLaunchServicesProvider) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/hash
func (c CPXLaunchServicesProvider) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/launchServicesInterface
func (c CPXLaunchServicesProvider) LaunchServicesInterface() ICPXLaunchServicesInterface {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("launchServicesInterface"))
	return CPXLaunchServicesInterfaceFromID(objc.ID(rv))
}
func (c CPXLaunchServicesProvider) SetLaunchServicesInterface(value ICPXLaunchServicesInterface) {
	objc.Send[struct{}](c.ID, objc.Sel("setLaunchServicesInterface:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLaunchServicesProvider/superclass
func (c CPXLaunchServicesProvider) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
