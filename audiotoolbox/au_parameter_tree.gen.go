// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AUParameterTree] class.
var (
	_AUParameterTreeClass     AUParameterTreeClass
	_AUParameterTreeClassOnce sync.Once
)

func getAUParameterTreeClass() AUParameterTreeClass {
	_AUParameterTreeClassOnce.Do(func() {
		_AUParameterTreeClass = AUParameterTreeClass{class: objc.GetClass("AUParameterTree")}
	})
	return _AUParameterTreeClass
}

// GetAUParameterTreeClass returns the class object for AUParameterTree.
func GetAUParameterTreeClass() AUParameterTreeClass {
	return getAUParameterTreeClass()
}

type AUParameterTreeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUParameterTreeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUParameterTreeClass) Alloc() AUParameterTree {
	rv := objc.Send[AUParameterTree](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a top-level group node that contains all of an
// audio unit’s parameters.
//
// # Overview
//
// An audio unit’s parameters are organized into a tree containing groups
// and parameters (groups may be nested).
//
// The parameter tree is KVO-compliant. An audio unit may choose to
// dynamically rearrange the tree; when doing so, it must issue a KVO
// notification on the audio unit’s [AUAudioUnit.ParameterTree] property.
//
// # Obtaining Tree Parameters
//
//   - [AUParameterTree.ParameterWithAddress]: Searches the tree for a parameter with a specific address.
//   - [AUParameterTree.ParameterWithIDScopeElement]: Searches the tree for a specific version 2 audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree
type AUParameterTree struct {
	AUParameterGroup
}

// AUParameterTreeFromID constructs a [AUParameterTree] from an objc.ID.
//
// An object that represents a top-level group node that contains all of an
// audio unit’s parameters.
func AUParameterTreeFromID(id objc.ID) AUParameterTree {
	return AUParameterTree{AUParameterGroup: AUParameterGroupFromID(id)}
}

// NOTE: AUParameterTree adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUParameterTree] class.
//
// # Obtaining Tree Parameters
//
//   - [IAUParameterTree.ParameterWithAddress]: Searches the tree for a parameter with a specific address.
//   - [IAUParameterTree.ParameterWithIDScopeElement]: Searches the tree for a specific version 2 audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree
type IAUParameterTree interface {
	IAUParameterGroup

	// Topic: Obtaining Tree Parameters

	// Searches the tree for a parameter with a specific address.
	ParameterWithAddress(address AUParameterAddress) IAUParameter
	// Searches the tree for a specific version 2 audio unit parameter.
	ParameterWithIDScopeElement(paramID AudioUnitParameterID, scope AudioUnitScope, element AudioUnitElement) IAUParameter
}

// Init initializes the instance.
func (p AUParameterTree) Init() AUParameterTree {
	rv := objc.Send[AUParameterTree](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AUParameterTree) Autorelease() AUParameterTree {
	rv := objc.Send[AUParameterTree](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUParameterTree creates a new AUParameterTree instance.
func NewAUParameterTree() AUParameterTree {
	class := getAUParameterTreeClass()
	rv := objc.Send[AUParameterTree](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterGroup/init(coder:)
func NewParameterTreeWithCoder(coder foundation.INSCoder) AUParameterTree {
	instance := getAUParameterTreeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AUParameterTreeFromID(rv)
}

// Searches the tree for a parameter with a specific address.
//
// address: The address with which to search the tree.
//
// # Return Value
//
// The parameter corresponding to the supplied address, or `nil` if no such
// parameter exists.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withAddress:)
func (p AUParameterTree) ParameterWithAddress(address AUParameterAddress) IAUParameter {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parameterWithAddress:"), address)
	return AUParameterFromID(rv)
}

// Searches the tree for a specific version 2 audio unit parameter.
//
// paramID: The parameter ID with which to search the tree.
//
// scope: The scope with which to search the tree.
//
// element: The element with which to search the tree.
//
// # Return Value
//
// The parameter corresponding to the supplied ID, scope, and element. Returns
// `nil` if the parameter is nonexistent or if it is not associated with a
// version 2 audio unit.
//
// # Discussion
//
// Version 2 audio units publish parameters identified by a parameter ID,
// scope, and element. A host that knows that it is dealing with a version 2
// audio unit can locate parameters using this method—for example, for the
// Apple-supplied system audio units.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/parameter(withID:scope:element:)
func (p AUParameterTree) ParameterWithIDScopeElement(paramID AudioUnitParameterID, scope AudioUnitScope, element AudioUnitElement) IAUParameter {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parameterWithID:scope:element:"), paramID, scope, element)
	return AUParameterFromID(rv)
}

// Creates a single parameter object.
//
// identifier: The parameter’s non-localized, permanent name.
//
// name: The parameter’s localized name for display.
//
// address: The parameter’s address.
//
// min: The parameter’s minimum value.
//
// max: The parameter’s maximum value.
//
// unit: The parameter’s unit of measurement.
//
// unitName: The parameter’s localized unit name.
//
// flags: The parameter’s characteristic details.
//
// valueStrings: The parameter’s localized value strings.
//
// dependentParameters: Any other parameter’s whose values may change as a side effect of this
// parameter’s value changing.
//
// # Return Value
//
// A newly-initialized parameter object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createParameter(withIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:)
func (_AUParameterTreeClass AUParameterTreeClass) CreateParameterWithIdentifierNameAddressMinMaxUnitUnitNameFlagsValueStringsDependentParameters(identifier string, name string, address AUParameterAddress, min AUValue, max AUValue, unit AudioUnitParameterUnit, unitName string, flags AudioUnitParameterOptions, valueStrings []string, dependentParameters []foundation.NSNumber) AUParameter {
	rv := objc.Send[objc.ID](objc.ID(_AUParameterTreeClass.class), objc.Sel("createParameterWithIdentifier:name:address:min:max:unit:unitName:flags:valueStrings:dependentParameters:"), objc.String(identifier), objc.String(name), address, min, max, unit, objc.String(unitName), flags, objectivec.StringSliceToNSArray(valueStrings), objectivec.IObjectSliceToNSArray(dependentParameters))
	return AUParameterFromID(rv)
}

// Creates a parameter group object.
//
// identifier: A non-localized, persistent identifier for the group.
//
// name: A localized display name for the group.
//
// children: The group’s child nodes.
//
// # Return Value
//
// A newly-initialized parameter group object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(withIdentifier:name:children:)
func (_AUParameterTreeClass AUParameterTreeClass) CreateGroupWithIdentifierNameChildren(identifier string, name string, children []AUParameterNode) AUParameterGroup {
	rv := objc.Send[objc.ID](objc.ID(_AUParameterTreeClass.class), objc.Sel("createGroupWithIdentifier:name:children:"), objc.String(identifier), objc.String(name), objectivec.IObjectSliceToNSArray(children))
	return AUParameterGroupFromID(rv)
}

// Creates a template group which may be used as a prototype for further group
// instances.
//
// children: The template group’s child nodes.
//
// # Return Value
//
// A newly-initialized parameter group template.
//
// # Discussion
//
// Template groups provide a way to construct multiple instances of identical
// parameter groups, sharing certain immutable state between the instances.
//
// Template groups may not appear in trees except at the root.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroupTemplate(_:)
func (_AUParameterTreeClass AUParameterTreeClass) CreateGroupTemplate(children []AUParameterNode) AUParameterGroup {
	rv := objc.Send[objc.ID](objc.ID(_AUParameterTreeClass.class), objc.Sel("createGroupTemplate:"), objectivec.IObjectSliceToNSArray(children))
	return AUParameterGroupFromID(rv)
}

// Initializes a group as a copied instance of a template group.
//
// templateGroup: A group to be used as a template and largely copied from.
//
// identifier: A non-localized, persistent identifier for the new group.
//
// name: A localized display name for the new group.
//
// addressOffset: The address offset for the new group’s parameters, with respect to the
// template group.
//
// # Return Value
//
// A newly-initialized parameter group object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createGroup(fromTemplate:identifier:name:addressOffset:)
func (_AUParameterTreeClass AUParameterTreeClass) CreateGroupFromTemplateIdentifierNameAddressOffset(templateGroup IAUParameterGroup, identifier string, name string, addressOffset AUParameterAddress) AUParameterGroup {
	rv := objc.Send[objc.ID](objc.ID(_AUParameterTreeClass.class), objc.Sel("createGroupFromTemplate:identifier:name:addressOffset:"), templateGroup, objc.String(identifier), objc.String(name), addressOffset)
	return AUParameterGroupFromID(rv)
}

// Creates a parameter tree object.
//
// children: The tree’s top-level children nodes.
//
// # Return Value
//
// A newly-initialized parameter tree object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterTree/createTree(withChildren:)
func (_AUParameterTreeClass AUParameterTreeClass) CreateTreeWithChildren(children []AUParameterNode) AUParameterTree {
	rv := objc.Send[objc.ID](objc.ID(_AUParameterTreeClass.class), objc.Sel("createTreeWithChildren:"), objectivec.IObjectSliceToNSArray(children))
	return AUParameterTreeFromID(rv)
}
