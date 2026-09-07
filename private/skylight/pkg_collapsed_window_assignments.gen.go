// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGCollapsedWindowAssignments] class.
var (
	_PKGCollapsedWindowAssignmentsClass     PKGCollapsedWindowAssignmentsClass
	_PKGCollapsedWindowAssignmentsClassOnce sync.Once
)

func getPKGCollapsedWindowAssignmentsClass() PKGCollapsedWindowAssignmentsClass {
	_PKGCollapsedWindowAssignmentsClassOnce.Do(func() {
		_PKGCollapsedWindowAssignmentsClass = PKGCollapsedWindowAssignmentsClass{class: objc.GetClass("PKGCollapsedWindowAssignments")}
	})
	return _PKGCollapsedWindowAssignmentsClass
}

// GetPKGCollapsedWindowAssignmentsClass returns the class object for PKGCollapsedWindowAssignments.
func GetPKGCollapsedWindowAssignmentsClass() PKGCollapsedWindowAssignmentsClass {
	return getPKGCollapsedWindowAssignmentsClass()
}

type PKGCollapsedWindowAssignmentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCollapsedWindowAssignmentsClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCollapsedWindowAssignmentsClass) Alloc() PKGCollapsedWindowAssignments {
	rv := objc.SendIfResponds[PKGCollapsedWindowAssignments](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGCollapsedWindowAssignments.CollapsedWindowIDsForManagedSpaceID]
type PKGCollapsedWindowAssignments struct {
	objectivec.Object
}

// PKGCollapsedWindowAssignmentsFromID constructs a [PKGCollapsedWindowAssignments] from an objc.ID.
func PKGCollapsedWindowAssignmentsFromID(id objc.ID) PKGCollapsedWindowAssignments {
	return PKGCollapsedWindowAssignments{objectivec.Object{ID: id}}
}

// Ensure PKGCollapsedWindowAssignments implements IPKGCollapsedWindowAssignments.
var _ IPKGCollapsedWindowAssignments = PKGCollapsedWindowAssignments{}

// An interface definition for the [PKGCollapsedWindowAssignments] class.
//
// # Methods
//
//   - [IPKGCollapsedWindowAssignments.CollapsedWindowIDsForManagedSpaceID]
type IPKGCollapsedWindowAssignments interface {
	objectivec.IObject

	// Topic: Methods

	CollapsedWindowIDsForManagedSpaceID(id objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (p PKGCollapsedWindowAssignments) Init() PKGCollapsedWindowAssignments {
	rv := objc.SendIfResponds[PKGCollapsedWindowAssignments](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGCollapsedWindowAssignments) Autorelease() PKGCollapsedWindowAssignments {
	rv := objc.SendIfResponds[PKGCollapsedWindowAssignments](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCollapsedWindowAssignments creates a new PKGCollapsedWindowAssignments instance.
func NewPKGCollapsedWindowAssignments() PKGCollapsedWindowAssignments {
	class := getPKGCollapsedWindowAssignmentsClass()
	rv := objc.SendIfResponds[PKGCollapsedWindowAssignments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGCollapsedWindowAssignments) CollapsedWindowIDsForManagedSpaceID(id objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("collapsedWindowIDsForManagedSpaceID:"), id)
	return objectivec.Object{ID: rv}
}
