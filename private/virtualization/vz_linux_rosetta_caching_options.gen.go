// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

//
// # Methods
//
//   - [VZLinuxRosettaCachingOptions._init]
//   - [VZLinuxRosettaCachingOptions._options]
//   - [VZLinuxRosettaCachingOptions.DebugDescription]
//   - [VZLinuxRosettaCachingOptions.Description]
//   - [VZLinuxRosettaCachingOptions.Hash]
//   - [VZLinuxRosettaCachingOptions.Superclass]
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
func (l VZLinuxRosettaCachingOptions) Init() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxRosettaCachingOptions) Autorelease() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaCachingOptions creates a new VZLinuxRosettaCachingOptions instance.
func NewVZLinuxRosettaCachingOptions() VZLinuxRosettaCachingOptions {
	class := getVZLinuxRosettaCachingOptionsClass()
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/_init
func (l VZLinuxRosettaCachingOptions) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/_options
func (l VZLinuxRosettaCachingOptions) _options() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("_options"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/debugDescription
func (l VZLinuxRosettaCachingOptions) DebugDescription() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/description
func (l VZLinuxRosettaCachingOptions) Description() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/hash
func (l VZLinuxRosettaCachingOptions) Hash() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions/superclass
func (l VZLinuxRosettaCachingOptions) Superclass() objc.Class {
	rv := objc.Send[objc.Class](l.ID, objc.Sel("superclass"))
	return rv
}

