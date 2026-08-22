// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCustomMigrationStage] class.
var (
	_NSCustomMigrationStageClass     NSCustomMigrationStageClass
	_NSCustomMigrationStageClassOnce sync.Once
)

func getNSCustomMigrationStageClass() NSCustomMigrationStageClass {
	_NSCustomMigrationStageClassOnce.Do(func() {
		_NSCustomMigrationStageClass = NSCustomMigrationStageClass{class: objc.GetClass("NSCustomMigrationStage")}
	})
	return _NSCustomMigrationStageClass
}

// GetNSCustomMigrationStageClass returns the class object for NSCustomMigrationStage.
func GetNSCustomMigrationStageClass() NSCustomMigrationStageClass {
	return getNSCustomMigrationStageClass()
}

type NSCustomMigrationStageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCustomMigrationStageClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCustomMigrationStageClass) Alloc() NSCustomMigrationStage {
	rv := objc.Send[NSCustomMigrationStage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that enables you to participate in the migration between two
// versions of the same model.
//
// # Overview
//
// Use [NSCustomMigrationStage] when you have two versions of a model that
// Core Data can’t automatically migrate. Custom migration stages enable you
// to participate in the migration process by assigning handlers that the
// stage invokes before and after it runs. The handlers provide an opportunity
// to prepare the persistent store’s data for the upcoming changes before
// the stage runs, and perform any cleanup tasks afterward.
//
// For example, to support a migration that changes an optional attribute to
// be nonoptional, you might assign a handler to the stage’s
// [NSCustomMigrationStage.WillMigrateHandler] property that sets any `nil`
// instances of that attribute to a default value, thereby ensuring the
// migration succeeds. To access the store you’re migrating, use the
// [NSStagedMigrationManager.Container] property of the migration manager that
// Core Data provides to every handler.
//
// # Accessing model references
//
//   - [NSCustomMigrationStage.CurrentModel]: The reference that represents the migration’s source model.
//   - [NSCustomMigrationStage.NextModel]: The reference that represents the migration’s destination model.
//
// # Assigning event handlers
//
//   - [NSCustomMigrationStage.WillMigrateHandler]: The handler to execute before the stage runs.
//   - [NSCustomMigrationStage.SetWillMigrateHandler]
//   - [NSCustomMigrationStage.DidMigrateHandler]: The handler to execute after the stage runs.
//   - [NSCustomMigrationStage.SetDidMigrateHandler]
//
// See: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage
type NSCustomMigrationStage struct {
	NSMigrationStage
}

// NSCustomMigrationStageFromID constructs a [NSCustomMigrationStage] from an objc.ID.
//
// An object that enables you to participate in the migration between two
// versions of the same model.
func NSCustomMigrationStageFromID(id objc.ID) NSCustomMigrationStage {
	return NSCustomMigrationStage{NSMigrationStage: NSMigrationStageFromID(id)}
}

// NOTE: NSCustomMigrationStage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCustomMigrationStage] class.
//
// # Accessing model references
//
//   - [INSCustomMigrationStage.CurrentModel]: The reference that represents the migration’s source model.
//   - [INSCustomMigrationStage.NextModel]: The reference that represents the migration’s destination model.
//
// # Assigning event handlers
//
//   - [INSCustomMigrationStage.WillMigrateHandler]: The handler to execute before the stage runs.
//   - [INSCustomMigrationStage.SetWillMigrateHandler]
//   - [INSCustomMigrationStage.DidMigrateHandler]: The handler to execute after the stage runs.
//   - [INSCustomMigrationStage.SetDidMigrateHandler]
//
// See: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage
type INSCustomMigrationStage interface {
	INSMigrationStage

	// Topic: Accessing model references

	// The reference that represents the migration’s source model.
	CurrentModel() INSManagedObjectModelReference
	// The reference that represents the migration’s destination model.
	NextModel() INSManagedObjectModelReference

	// Topic: Assigning event handlers

	// The handler to execute before the stage runs.
	WillMigrateHandler() objectivec.IObject
	SetWillMigrateHandler(value objectivec.IObject)
	// The handler to execute after the stage runs.
	DidMigrateHandler() objectivec.IObject
	SetDidMigrateHandler(value objectivec.IObject)
}

// Init initializes the instance.
func (c NSCustomMigrationStage) Init() NSCustomMigrationStage {
	rv := objc.Send[NSCustomMigrationStage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCustomMigrationStage) Autorelease() NSCustomMigrationStage {
	rv := objc.Send[NSCustomMigrationStage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCustomMigrationStage creates a new NSCustomMigrationStage instance.
func NewNSCustomMigrationStage() NSCustomMigrationStage {
	class := getNSCustomMigrationStageClass()
	rv := objc.Send[NSCustomMigrationStage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The reference that represents the migration’s source model.
//
// # Discussion
//
// Core Data sets this property to the `currentModel` parameter you specify
// when creating the migration stage.
//
// See: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/currentModel
func (c NSCustomMigrationStage) CurrentModel() INSManagedObjectModelReference {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentModel"))
	return NSManagedObjectModelReferenceFromID(objc.ID(rv))
}

// The reference that represents the migration’s destination model.
//
// # Discussion
//
// Core Data sets this property to the `nextModel` parameter you specify when
// creating the migration stage.
//
// See: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/nextModel
func (c NSCustomMigrationStage) NextModel() INSManagedObjectModelReference {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nextModel"))
	return NSManagedObjectModelReferenceFromID(objc.ID(rv))
}

// The handler to execute before the stage runs.
//
// See: https://developer.apple.com/documentation/coredata/nscustommigrationstage/willmigratehandler-5wead
func (c NSCustomMigrationStage) WillMigrateHandler() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("willMigrateHandler"))
	return objectivec.Object{ID: rv}
}
func (c NSCustomMigrationStage) SetWillMigrateHandler(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setWillMigrateHandler:"), value)
}

// The handler to execute after the stage runs.
//
// See: https://developer.apple.com/documentation/coredata/nscustommigrationstage/didmigratehandler-2zbss
func (c NSCustomMigrationStage) DidMigrateHandler() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("didMigrateHandler"))
	return objectivec.Object{ID: rv}
}
func (c NSCustomMigrationStage) SetDidMigrateHandler(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setDidMigrateHandler:"), value)
}
