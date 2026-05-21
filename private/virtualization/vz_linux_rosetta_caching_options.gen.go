// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
type IVZLinuxRosettaCachingOptions interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_options() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func (v VZLinuxRosettaCachingOptions) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZLinuxRosettaCachingOptions) _options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_options"))
	return rv
}

// CanOptions reports whether the receiver responds to the private selector _options.
func (v VZLinuxRosettaCachingOptions) CanOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_options"))
}

// Options is an exported wrapper for the private property _options.
func (v VZLinuxRosettaCachingOptions) Options() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_options")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_options"}
	}
	return v._options(), nil
}
func (v VZLinuxRosettaCachingOptions) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZLinuxRosettaCachingOptions) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZLinuxRosettaCachingOptions) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZLinuxRosettaCachingOptions) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
