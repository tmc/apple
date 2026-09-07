// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGPreferredSpaceAssignments] class.
var (
	_PKGPreferredSpaceAssignmentsClass     PKGPreferredSpaceAssignmentsClass
	_PKGPreferredSpaceAssignmentsClassOnce sync.Once
)

func getPKGPreferredSpaceAssignmentsClass() PKGPreferredSpaceAssignmentsClass {
	_PKGPreferredSpaceAssignmentsClassOnce.Do(func() {
		_PKGPreferredSpaceAssignmentsClass = PKGPreferredSpaceAssignmentsClass{class: objc.GetClass("PKGPreferredSpaceAssignments")}
	})
	return _PKGPreferredSpaceAssignmentsClass
}

// GetPKGPreferredSpaceAssignmentsClass returns the class object for PKGPreferredSpaceAssignments.
func GetPKGPreferredSpaceAssignmentsClass() PKGPreferredSpaceAssignmentsClass {
	return getPKGPreferredSpaceAssignmentsClass()
}

type PKGPreferredSpaceAssignmentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGPreferredSpaceAssignmentsClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGPreferredSpaceAssignmentsClass) Alloc() PKGPreferredSpaceAssignments {
	rv := objc.SendIfResponds[PKGPreferredSpaceAssignments](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGPreferredSpaceAssignments.PreferredManagedSpaceIDsForManagedDisplayID]
type PKGPreferredSpaceAssignments struct {
	objectivec.Object
}

// PKGPreferredSpaceAssignmentsFromID constructs a [PKGPreferredSpaceAssignments] from an objc.ID.
func PKGPreferredSpaceAssignmentsFromID(id objc.ID) PKGPreferredSpaceAssignments {
	return PKGPreferredSpaceAssignments{objectivec.Object{ID: id}}
}

// Ensure PKGPreferredSpaceAssignments implements IPKGPreferredSpaceAssignments.
var _ IPKGPreferredSpaceAssignments = PKGPreferredSpaceAssignments{}

// An interface definition for the [PKGPreferredSpaceAssignments] class.
//
// # Methods
//
//   - [IPKGPreferredSpaceAssignments.PreferredManagedSpaceIDsForManagedDisplayID]
type IPKGPreferredSpaceAssignments interface {
	objectivec.IObject

	// Topic: Methods

	PreferredManagedSpaceIDsForManagedDisplayID(id objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (p PKGPreferredSpaceAssignments) Init() PKGPreferredSpaceAssignments {
	rv := objc.SendIfResponds[PKGPreferredSpaceAssignments](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGPreferredSpaceAssignments) Autorelease() PKGPreferredSpaceAssignments {
	rv := objc.SendIfResponds[PKGPreferredSpaceAssignments](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGPreferredSpaceAssignments creates a new PKGPreferredSpaceAssignments instance.
func NewPKGPreferredSpaceAssignments() PKGPreferredSpaceAssignments {
	class := getPKGPreferredSpaceAssignmentsClass()
	rv := objc.SendIfResponds[PKGPreferredSpaceAssignments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGPreferredSpaceAssignments) PreferredManagedSpaceIDsForManagedDisplayID(id objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("preferredManagedSpaceIDsForManagedDisplayID:"), id)
	return objectivec.Object{ID: rv}
}
