// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGSpaceAssignmentController] class.
var (
	_PKGSpaceAssignmentControllerClass     PKGSpaceAssignmentControllerClass
	_PKGSpaceAssignmentControllerClassOnce sync.Once
)

func getPKGSpaceAssignmentControllerClass() PKGSpaceAssignmentControllerClass {
	_PKGSpaceAssignmentControllerClassOnce.Do(func() {
		_PKGSpaceAssignmentControllerClass = PKGSpaceAssignmentControllerClass{class: objc.GetClass("PKGSpaceAssignmentController")}
	})
	return _PKGSpaceAssignmentControllerClass
}

// GetPKGSpaceAssignmentControllerClass returns the class object for PKGSpaceAssignmentController.
func GetPKGSpaceAssignmentControllerClass() PKGSpaceAssignmentControllerClass {
	return getPKGSpaceAssignmentControllerClass()
}

type PKGSpaceAssignmentControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSpaceAssignmentControllerClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSpaceAssignmentControllerClass) Alloc() PKGSpaceAssignmentController {
	rv := objc.SendIfResponds[PKGSpaceAssignmentController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGSpaceAssignmentController.CopyApplicationPersistencePropertyListForWindowID]
//   - [PKGSpaceAssignmentController.CopyDebugPropertyList]
//   - [PKGSpaceAssignmentController.CopyPersistencePropertyList]
//   - [PKGSpaceAssignmentController.LoadApplicationPersistencePropertyListForWindowID]
//   - [PKGSpaceAssignmentController.LoadPersistencePropertyList]
//   - [PKGSpaceAssignmentController.OrderedManagedSpaceIDsForManagedDisplayID]
//   - [PKGSpaceAssignmentController.PreferredSpaceAssignmentsForActiveDisplayIDs]
//   - [PKGSpaceAssignmentController.RecordManagedSpaceIDPrefersManagedDisplayID]
//   - [PKGSpaceAssignmentController.RecordManagedSpaceOrderForManagedDisplayID]
//   - [PKGSpaceAssignmentController.RecordWindowIDWillCollapseFromManagedSpaceID]
//   - [PKGSpaceAssignmentController.RemoveManagedDisplayID]
//   - [PKGSpaceAssignmentController.RemoveManagedSpaceID]
//   - [PKGSpaceAssignmentController.RemoveWindowID]
//   - [PKGSpaceAssignmentController.TakeCollapsedWindowAssignmentsForActiveManagedSpaceIDs]
type PKGSpaceAssignmentController struct {
	objectivec.Object
}

// PKGSpaceAssignmentControllerFromID constructs a [PKGSpaceAssignmentController] from an objc.ID.
func PKGSpaceAssignmentControllerFromID(id objc.ID) PKGSpaceAssignmentController {
	return PKGSpaceAssignmentController{objectivec.Object{ID: id}}
}

// Ensure PKGSpaceAssignmentController implements IPKGSpaceAssignmentController.
var _ IPKGSpaceAssignmentController = PKGSpaceAssignmentController{}

// An interface definition for the [PKGSpaceAssignmentController] class.
//
// # Methods
//
//   - [IPKGSpaceAssignmentController.CopyApplicationPersistencePropertyListForWindowID]
//   - [IPKGSpaceAssignmentController.CopyDebugPropertyList]
//   - [IPKGSpaceAssignmentController.CopyPersistencePropertyList]
//   - [IPKGSpaceAssignmentController.LoadApplicationPersistencePropertyListForWindowID]
//   - [IPKGSpaceAssignmentController.LoadPersistencePropertyList]
//   - [IPKGSpaceAssignmentController.OrderedManagedSpaceIDsForManagedDisplayID]
//   - [IPKGSpaceAssignmentController.PreferredSpaceAssignmentsForActiveDisplayIDs]
//   - [IPKGSpaceAssignmentController.RecordManagedSpaceIDPrefersManagedDisplayID]
//   - [IPKGSpaceAssignmentController.RecordManagedSpaceOrderForManagedDisplayID]
//   - [IPKGSpaceAssignmentController.RecordWindowIDWillCollapseFromManagedSpaceID]
//   - [IPKGSpaceAssignmentController.RemoveManagedDisplayID]
//   - [IPKGSpaceAssignmentController.RemoveManagedSpaceID]
//   - [IPKGSpaceAssignmentController.RemoveWindowID]
//   - [IPKGSpaceAssignmentController.TakeCollapsedWindowAssignmentsForActiveManagedSpaceIDs]
type IPKGSpaceAssignmentController interface {
	objectivec.IObject

	// Topic: Methods

	CopyApplicationPersistencePropertyListForWindowID(id uint32) objectivec.IObject
	CopyDebugPropertyList() objectivec.IObject
	CopyPersistencePropertyList() objectivec.IObject
	LoadApplicationPersistencePropertyListForWindowID(list objectivec.IObject, id uint32)
	LoadPersistencePropertyList(list objectivec.IObject)
	OrderedManagedSpaceIDsForManagedDisplayID(id objectivec.IObject) objectivec.IObject
	PreferredSpaceAssignmentsForActiveDisplayIDs(iDs objectivec.IObject) objectivec.IObject
	RecordManagedSpaceIDPrefersManagedDisplayID(id objectivec.IObject, id2 objectivec.IObject)
	RecordManagedSpaceOrderForManagedDisplayID(order objectivec.IObject, id objectivec.IObject)
	RecordWindowIDWillCollapseFromManagedSpaceID(id uint32, id2 objectivec.IObject)
	RemoveManagedDisplayID(id objectivec.IObject)
	RemoveManagedSpaceID(id objectivec.IObject)
	RemoveWindowID(id uint32)
	TakeCollapsedWindowAssignmentsForActiveManagedSpaceIDs(iDs objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (p PKGSpaceAssignmentController) Init() PKGSpaceAssignmentController {
	rv := objc.SendIfResponds[PKGSpaceAssignmentController](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGSpaceAssignmentController) Autorelease() PKGSpaceAssignmentController {
	rv := objc.SendIfResponds[PKGSpaceAssignmentController](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSpaceAssignmentController creates a new PKGSpaceAssignmentController instance.
func NewPKGSpaceAssignmentController() PKGSpaceAssignmentController {
	class := getPKGSpaceAssignmentControllerClass()
	rv := objc.SendIfResponds[PKGSpaceAssignmentController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGSpaceAssignmentController) CopyApplicationPersistencePropertyListForWindowID(id uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("copyApplicationPersistencePropertyListForWindowID:"), id)
	return objectivec.Object{ID: rv}
}
func (p PKGSpaceAssignmentController) CopyDebugPropertyList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("copyDebugPropertyList"))
	return objectivec.Object{ID: rv}
}
func (p PKGSpaceAssignmentController) CopyPersistencePropertyList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("copyPersistencePropertyList"))
	return objectivec.Object{ID: rv}
}
func (p PKGSpaceAssignmentController) LoadApplicationPersistencePropertyListForWindowID(list objectivec.IObject, id uint32) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("loadApplicationPersistencePropertyList:forWindowID:"), list, id)
}
func (p PKGSpaceAssignmentController) LoadPersistencePropertyList(list objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("loadPersistencePropertyList:"), list)
}
func (p PKGSpaceAssignmentController) OrderedManagedSpaceIDsForManagedDisplayID(id objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("orderedManagedSpaceIDsForManagedDisplayID:"), id)
	return objectivec.Object{ID: rv}
}
func (p PKGSpaceAssignmentController) PreferredSpaceAssignmentsForActiveDisplayIDs(iDs objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("preferredSpaceAssignmentsForActiveDisplayIDs:"), iDs)
	return objectivec.Object{ID: rv}
}
func (p PKGSpaceAssignmentController) RecordManagedSpaceIDPrefersManagedDisplayID(id objectivec.IObject, id2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("recordManagedSpaceID:prefersManagedDisplayID:"), id, id2)
}
func (p PKGSpaceAssignmentController) RecordManagedSpaceOrderForManagedDisplayID(order objectivec.IObject, id objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("recordManagedSpaceOrder:forManagedDisplayID:"), order, id)
}
func (p PKGSpaceAssignmentController) RecordWindowIDWillCollapseFromManagedSpaceID(id uint32, id2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("recordWindowID:willCollapseFromManagedSpaceID:"), id, id2)
}
func (p PKGSpaceAssignmentController) RemoveManagedDisplayID(id objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("removeManagedDisplayID:"), id)
}
func (p PKGSpaceAssignmentController) RemoveManagedSpaceID(id objectivec.IObject) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("removeManagedSpaceID:"), id)
}
func (p PKGSpaceAssignmentController) RemoveWindowID(id uint32) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("removeWindowID:"), id)
}
func (p PKGSpaceAssignmentController) TakeCollapsedWindowAssignmentsForActiveManagedSpaceIDs(iDs objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("takeCollapsedWindowAssignmentsForActiveManagedSpaceIDs:"), iDs)
	return objectivec.Object{ID: rv}
}
