// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStagedMigrationManager] class.
var (
	_NSStagedMigrationManagerClass     NSStagedMigrationManagerClass
	_NSStagedMigrationManagerClassOnce sync.Once
)

func getNSStagedMigrationManagerClass() NSStagedMigrationManagerClass {
	_NSStagedMigrationManagerClassOnce.Do(func() {
		_NSStagedMigrationManagerClass = NSStagedMigrationManagerClass{class: objc.GetClass("NSStagedMigrationManager")}
	})
	return _NSStagedMigrationManagerClass
}

// GetNSStagedMigrationManagerClass returns the class object for NSStagedMigrationManager.
func GetNSStagedMigrationManagerClass() NSStagedMigrationManagerClass {
	return getNSStagedMigrationManagerClass()
}

type NSStagedMigrationManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStagedMigrationManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStagedMigrationManagerClass) Alloc() NSStagedMigrationManager {
	rv := objc.Send[NSStagedMigrationManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that handles the migration event loop and provides access to the
// migrating persistent store.
//
// # Overview
//
// A staged migration manager contains the individual stages of a migration
// and applies those stages, in the order you specify, when that migration
// runs. The manager handles the migration’s event loop, and provides access
// to the migrating store through its [NSStagedMigrationManager.Container]
// property. Stages can be custom, which enables you to perform tasks
// immediately before and after a stage runs, or lightweight, which
// supplements custom stages with those that Core Data can invoke
// automatically because they’re already compatible with lightweight
// migrations.
//
// Use [NSPersistentStoreStagedMigrationManagerOptionKey] to include an
// instance of [NSStagedMigrationManager] in your persistent store’s options
// dictionary, as the following example shows:
//
// # Accessing the persistent container
//
//   - [NSStagedMigrationManager.Container]: The container that provides access to the migrating persistent store.
//
// # Accessing the stages
//
//   - [NSStagedMigrationManager.Stages]: The migration stages.
//
// See: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager
//
// [NSPersistentStoreStagedMigrationManagerOptionKey]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreStagedMigrationManagerOptionKey
type NSStagedMigrationManager struct {
	objectivec.Object
}

// NSStagedMigrationManagerFromID constructs a [NSStagedMigrationManager] from an objc.ID.
//
// An object that handles the migration event loop and provides access to the
// migrating persistent store.
func NSStagedMigrationManagerFromID(id objc.ID) NSStagedMigrationManager {
	return NSStagedMigrationManager{objectivec.Object{ID: id}}
}

// NOTE: NSStagedMigrationManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStagedMigrationManager] class.
//
// # Accessing the persistent container
//
//   - [INSStagedMigrationManager.Container]: The container that provides access to the migrating persistent store.
//
// # Accessing the stages
//
//   - [INSStagedMigrationManager.Stages]: The migration stages.
//
// See: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager
type INSStagedMigrationManager interface {
	objectivec.IObject

	// Topic: Accessing the persistent container

	// The container that provides access to the migrating persistent store.
	Container() INSPersistentContainer

	// Topic: Accessing the stages

	// The migration stages.
	Stages() []NSMigrationStage
}

// Init initializes the instance.
func (s NSStagedMigrationManager) Init() NSStagedMigrationManager {
	rv := objc.Send[NSStagedMigrationManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStagedMigrationManager) Autorelease() NSStagedMigrationManager {
	rv := objc.Send[NSStagedMigrationManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStagedMigrationManager creates a new NSStagedMigrationManager instance.
func NewNSStagedMigrationManager() NSStagedMigrationManager {
	class := getNSStagedMigrationManagerClass()
	rv := objc.Send[NSStagedMigrationManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The container that provides access to the migrating persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager/container
func (s NSStagedMigrationManager) Container() INSPersistentContainer {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("container"))
	return NSPersistentContainerFromID(objc.ID(rv))
}

// The migration stages.
//
// # Discussion
//
// Core Data sets this property to the `stages` parameter you specify when
// creating the migration manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager/stages
func (s NSStagedMigrationManager) Stages() []NSMigrationStage {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("stages"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMigrationStage {
		return NSMigrationStageFromID(id)
	})
}
