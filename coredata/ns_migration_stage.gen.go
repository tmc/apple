// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMigrationStage] class.
var (
	_NSMigrationStageClass     NSMigrationStageClass
	_NSMigrationStageClassOnce sync.Once
)

func getNSMigrationStageClass() NSMigrationStageClass {
	_NSMigrationStageClassOnce.Do(func() {
		_NSMigrationStageClass = NSMigrationStageClass{class: objc.GetClass("NSMigrationStage")}
	})
	return _NSMigrationStageClass
}

// GetNSMigrationStageClass returns the class object for NSMigrationStage.
func GetNSMigrationStageClass() NSMigrationStageClass {
	return getNSMigrationStageClass()
}

type NSMigrationStageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMigrationStageClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMigrationStageClass) Alloc() NSMigrationStage {
	rv := objc.Send[NSMigrationStage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for describing an individual stage of a migration.
//
// # Overview
//
// # Describing the purpose
//
//   - [NSMigrationStage.Label]: The textual description of the migration stage’s purpose.
//   - [NSMigrationStage.SetLabel]
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationStage
type NSMigrationStage struct {
	objectivec.Object
}

// NSMigrationStageFromID constructs a [NSMigrationStage] from an objc.ID.
//
// An abstract base class for describing an individual stage of a migration.
func NSMigrationStageFromID(id objc.ID) NSMigrationStage {
	return NSMigrationStage{objectivec.Object{ID: id}}
}

// NOTE: NSMigrationStage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMigrationStage] class.
//
// # Describing the purpose
//
//   - [INSMigrationStage.Label]: The textual description of the migration stage’s purpose.
//   - [INSMigrationStage.SetLabel]
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationStage
type INSMigrationStage interface {
	objectivec.IObject

	// Topic: Describing the purpose

	// The textual description of the migration stage’s purpose.
	Label() string
	SetLabel(value string)
}

// Init initializes the instance.
func (m NSMigrationStage) Init() NSMigrationStage {
	rv := objc.Send[NSMigrationStage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMigrationStage) Autorelease() NSMigrationStage {
	rv := objc.Send[NSMigrationStage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMigrationStage creates a new NSMigrationStage instance.
func NewNSMigrationStage() NSMigrationStage {
	class := getNSMigrationStageClass()
	rv := objc.Send[NSMigrationStage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The textual description of the migration stage’s purpose.
//
// # Discussion
//
// Persistent history tracking, if enabled, records the label for later use.
// The default value is an empty string.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationStage/label
func (m NSMigrationStage) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMigrationStage) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}
