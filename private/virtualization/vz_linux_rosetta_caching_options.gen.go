// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZLinuxRosettaCachingOptions] class.
var (
	_VZLinuxRosettaCachingOptionsClass     VZLinuxRosettaCachingOptionsClass
	_VZLinuxRosettaCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaCachingOptionsClass() VZLinuxRosettaCachingOptionsClass {
	_VZLinuxRosettaCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaCachingOptionsClass = VZLinuxRosettaCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaCachingOptions")}
	})
	return _VZLinuxRosettaCachingOptionsClass
}

// GetVZLinuxRosettaCachingOptionsClass returns the class object for VZLinuxRosettaCachingOptions.
func GetVZLinuxRosettaCachingOptionsClass() VZLinuxRosettaCachingOptionsClass {
	return getVZLinuxRosettaCachingOptionsClass()
}

type VZLinuxRosettaCachingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaCachingOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaCachingOptionsClass) Alloc() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZLinuxRosettaCachingOptions._init]
//   - [VZLinuxRosettaCachingOptions._options]
//   - [VZLinuxRosettaCachingOptions.DebugDescription]
//   - [VZLinuxRosettaCachingOptions.Description]
//   - [VZLinuxRosettaCachingOptions.Hash]
//   - [VZLinuxRosettaCachingOptions.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type VZLinuxRosettaCachingOptions struct {
	objectivec.Object
}

// VZLinuxRosettaCachingOptionsFromID constructs a [VZLinuxRosettaCachingOptions] from an objc.ID.
func VZLinuxRosettaCachingOptionsFromID(id objc.ID) VZLinuxRosettaCachingOptions {
	return VZLinuxRosettaCachingOptions{objectivec.Object{ID: id}}
}

// Ensure VZLinuxRosettaCachingOptions implements IVZLinuxRosettaCachingOptions.
var _ IVZLinuxRosettaCachingOptions = VZLinuxRosettaCachingOptions{}

// An interface definition for the [VZLinuxRosettaCachingOptions] class.
//
// # Methods
//
//   - [IVZLinuxRosettaCachingOptions._init]
//   - [IVZLinuxRosettaCachingOptions._options]
//   - [IVZLinuxRosettaCachingOptions.DebugDescription]
//   - [IVZLinuxRosettaCachingOptions.Description]
//   - [IVZLinuxRosettaCachingOptions.Hash]
//   - [IVZLinuxRosettaCachingOptions.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type IVZLinuxRosettaCachingOptions interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_options() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZLinuxRosettaCachingOptions) Init() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinuxRosettaCachingOptions) Autorelease() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaCachingOptions creates a new VZLinuxRosettaCachingOptions instance.
func NewVZLinuxRosettaCachingOptions() VZLinuxRosettaCachingOptions {
	class := getVZLinuxRosettaCachingOptionsClass()
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/_init
func (v VZLinuxRosettaCachingOptions) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/_options
func (v VZLinuxRosettaCachingOptions) _options() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_options"))
	return objectivec.Object{ID: rv}
}

// CanOptions reports whether the receiver responds to the private selector _options.
func (v VZLinuxRosettaCachingOptions) CanOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_options"))
}

// Options is an exported wrapper for the private property _options.
func (v VZLinuxRosettaCachingOptions) Options() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_options")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_options"}
	}
	return v._options(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/debugDescription
func (v VZLinuxRosettaCachingOptions) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/description
func (v VZLinuxRosettaCachingOptions) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/hash
func (v VZLinuxRosettaCachingOptions) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/superclass
func (v VZLinuxRosettaCachingOptions) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
