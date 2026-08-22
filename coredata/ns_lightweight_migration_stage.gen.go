// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSLightweightMigrationStage] class.
var (
	_NSLightweightMigrationStageClass     NSLightweightMigrationStageClass
	_NSLightweightMigrationStageClassOnce sync.Once
)

func getNSLightweightMigrationStageClass() NSLightweightMigrationStageClass {
	_NSLightweightMigrationStageClassOnce.Do(func() {
		_NSLightweightMigrationStageClass = NSLightweightMigrationStageClass{class: objc.GetClass("NSLightweightMigrationStage")}
	})
	return _NSLightweightMigrationStageClass
}

// GetNSLightweightMigrationStageClass returns the class object for NSLightweightMigrationStage.
func GetNSLightweightMigrationStageClass() NSLightweightMigrationStageClass {
	return getNSLightweightMigrationStageClass()
}

type NSLightweightMigrationStageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLightweightMigrationStageClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLightweightMigrationStageClass) Alloc() NSLightweightMigrationStage {
	rv := objc.Send[NSLightweightMigrationStage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a series of models suitable for lightweight
// migration.
//
// # Overview
//
// Use [NSLightweightMigrationStage] when you have a series of models to
// migrate and those models are compatible with lightweight migrations.
// Instances of this class supplement your custom migration stages and help
// maintain a consistent stage order for the entire migration.
//
// # Accessing the checksums
//
//   - [NSLightweightMigrationStage.VersionChecksums]: The array of version checksums.
//
// See: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage
type NSLightweightMigrationStage struct {
	NSMigrationStage
}

// NSLightweightMigrationStageFromID constructs a [NSLightweightMigrationStage] from an objc.ID.
//
// An object that describes a series of models suitable for lightweight
// migration.
func NSLightweightMigrationStageFromID(id objc.ID) NSLightweightMigrationStage {
	return NSLightweightMigrationStage{NSMigrationStage: NSMigrationStageFromID(id)}
}

// NOTE: NSLightweightMigrationStage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLightweightMigrationStage] class.
//
// # Accessing the checksums
//
//   - [INSLightweightMigrationStage.VersionChecksums]: The array of version checksums.
//
// See: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage
type INSLightweightMigrationStage interface {
	INSMigrationStage

	// Topic: Accessing the checksums

	// The array of version checksums.
	VersionChecksums() []string
}

// Init initializes the instance.
func (l NSLightweightMigrationStage) Init() NSLightweightMigrationStage {
	rv := objc.Send[NSLightweightMigrationStage](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLightweightMigrationStage) Autorelease() NSLightweightMigrationStage {
	rv := objc.Send[NSLightweightMigrationStage](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLightweightMigrationStage creates a new NSLightweightMigrationStage instance.
func NewNSLightweightMigrationStage() NSLightweightMigrationStage {
	class := getNSLightweightMigrationStageClass()
	rv := objc.Send[NSLightweightMigrationStage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of version checksums.
//
// # Discussion
//
// Core Data sets this property to the checksum you specify when creating the
// lightweight migration stage.
//
// See: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage/versionChecksums
func (l NSLightweightMigrationStage) VersionChecksums() []string {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("versionChecksums"))
	return objc.ConvertSliceToStrings(rv)
}
