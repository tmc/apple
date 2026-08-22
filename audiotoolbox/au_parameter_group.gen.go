// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AUParameterGroup] class.
var (
	_AUParameterGroupClass     AUParameterGroupClass
	_AUParameterGroupClassOnce sync.Once
)

func getAUParameterGroupClass() AUParameterGroupClass {
	_AUParameterGroupClassOnce.Do(func() {
		_AUParameterGroupClass = AUParameterGroupClass{class: objc.GetClass("AUParameterGroup")}
	})
	return _AUParameterGroupClass
}

// GetAUParameterGroupClass returns the class object for AUParameterGroup.
func GetAUParameterGroupClass() AUParameterGroupClass {
	return getAUParameterGroupClass()
}

type AUParameterGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUParameterGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUParameterGroupClass) Alloc() AUParameterGroup {
	rv := objc.Send[AUParameterGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A parameter group object represents a group of related audio unit
// parameters.
//
// # Overview
//
// A parameter group is KVC-compliant for its children. For example, calling
// the parameter group’s [value(forKey:)] method, with a key value of
// volume, returns a child whose [AUParameterNode.Identifier] value matches
// that key.
//
// # Obtaining Group Parameters
//
//   - [AUParameterGroup.AllParameters]: Returns a flat array of all parameters in the group, including those in child groups.
//   - [AUParameterGroup.Children]: The group’s child nodes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup
//
// [value(forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKey:)
type AUParameterGroup struct {
	AUParameterNode
}

// AUParameterGroupFromID constructs a [AUParameterGroup] from an objc.ID.
//
// A parameter group object represents a group of related audio unit
// parameters.
func AUParameterGroupFromID(id objc.ID) AUParameterGroup {
	return AUParameterGroup{AUParameterNode: AUParameterNodeFromID(id)}
}

// NOTE: AUParameterGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUParameterGroup] class.
//
// # Obtaining Group Parameters
//
//   - [IAUParameterGroup.AllParameters]: Returns a flat array of all parameters in the group, including those in child groups.
//   - [IAUParameterGroup.Children]: The group’s child nodes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup
type IAUParameterGroup interface {
	IAUParameterNode

	// Topic: Obtaining Group Parameters

	// Returns a flat array of all parameters in the group, including those in child groups.
	AllParameters() []AUParameter
	// The group’s child nodes.
	Children() []AUParameterNode
}

// Init initializes the instance.
func (p AUParameterGroup) Init() AUParameterGroup {
	rv := objc.Send[AUParameterGroup](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AUParameterGroup) Autorelease() AUParameterGroup {
	rv := objc.Send[AUParameterGroup](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUParameterGroup creates a new AUParameterGroup instance.
func NewAUParameterGroup() AUParameterGroup {
	class := getAUParameterGroupClass()
	rv := objc.Send[AUParameterGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/init(coder:)
func NewParameterGroupWithCoder(coder foundation.INSCoder) AUParameterGroup {
	instance := getAUParameterGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AUParameterGroupFromID(rv)
}

// Returns a flat array of all parameters in the group, including those in
// child groups.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/allParameters
func (p AUParameterGroup) AllParameters() []AUParameter {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("allParameters"))
	return objc.ConvertSlice(rv, func(id objc.ID) AUParameter {
		return AUParameterFromID(id)
	})
}

// The group’s child nodes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/children
func (p AUParameterGroup) Children() []AUParameterNode {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("children"))
	return objc.ConvertSlice(rv, func(id objc.ID) AUParameterNode {
		return AUParameterNodeFromID(id)
	})
}
