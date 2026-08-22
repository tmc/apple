// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentCloudKitContainerOptions] class.
var (
	_NSPersistentCloudKitContainerOptionsClass     NSPersistentCloudKitContainerOptionsClass
	_NSPersistentCloudKitContainerOptionsClassOnce sync.Once
)

func getNSPersistentCloudKitContainerOptionsClass() NSPersistentCloudKitContainerOptionsClass {
	_NSPersistentCloudKitContainerOptionsClassOnce.Do(func() {
		_NSPersistentCloudKitContainerOptionsClass = NSPersistentCloudKitContainerOptionsClass{class: objc.GetClass("NSPersistentCloudKitContainerOptions")}
	})
	return _NSPersistentCloudKitContainerOptionsClass
}

// GetNSPersistentCloudKitContainerOptionsClass returns the class object for NSPersistentCloudKitContainerOptions.
func GetNSPersistentCloudKitContainerOptionsClass() NSPersistentCloudKitContainerOptionsClass {
	return getNSPersistentCloudKitContainerOptionsClass()
}

type NSPersistentCloudKitContainerOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentCloudKitContainerOptionsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentCloudKitContainerOptionsClass) Alloc() NSPersistentCloudKitContainerOptions {
	rv := objc.Send[NSPersistentCloudKitContainerOptions](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that customizes how a store description aligns with a CloudKit
// database.
//
// # Overview
//
// Use [NSPersistentCloudKitContainerOptions] to customize the behavior of an
// [NSPersistentCloudKitContainer] or to create additional store descriptions
// that sync to other containers.
//
// For more information about setting up multiple stores, see [Setting Up Core
// Data with CloudKit].
//
// # Creating Container Options
//
//   - [NSPersistentCloudKitContainerOptions.InitWithContainerIdentifier]: Initializes container options using the given CloudKit container identifier.
//   - [NSPersistentCloudKitContainerOptions.ContainerIdentifier]: The identifier of the CloudKit container associated with a given store description.
//   - [NSPersistentCloudKitContainerOptions.DatabaseScope]: The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container.
//   - [NSPersistentCloudKitContainerOptions.SetDatabaseScope]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions
//
// [Setting Up Core Data with CloudKit]: https://developer.apple.com/documentation/CoreData/setting-up-core-data-with-cloudkit
type NSPersistentCloudKitContainerOptions struct {
	objectivec.Object
}

// NSPersistentCloudKitContainerOptionsFromID constructs a [NSPersistentCloudKitContainerOptions] from an objc.ID.
//
// An object that customizes how a store description aligns with a CloudKit
// database.
func NSPersistentCloudKitContainerOptionsFromID(id objc.ID) NSPersistentCloudKitContainerOptions {
	return NSPersistentCloudKitContainerOptions{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentCloudKitContainerOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentCloudKitContainerOptions] class.
//
// # Creating Container Options
//
//   - [INSPersistentCloudKitContainerOptions.InitWithContainerIdentifier]: Initializes container options using the given CloudKit container identifier.
//   - [INSPersistentCloudKitContainerOptions.ContainerIdentifier]: The identifier of the CloudKit container associated with a given store description.
//   - [INSPersistentCloudKitContainerOptions.DatabaseScope]: The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container.
//   - [INSPersistentCloudKitContainerOptions.SetDatabaseScope]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions
type INSPersistentCloudKitContainerOptions interface {
	objectivec.IObject

	// Topic: Creating Container Options

	// Initializes container options using the given CloudKit container identifier.
	InitWithContainerIdentifier(containerIdentifier string) NSPersistentCloudKitContainerOptions
	// The identifier of the CloudKit container associated with a given store description.
	ContainerIdentifier() string
	// The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container.
	DatabaseScope() uint
	SetDatabaseScope(value uint)
}

// Init initializes the instance.
func (p NSPersistentCloudKitContainerOptions) Init() NSPersistentCloudKitContainerOptions {
	rv := objc.Send[NSPersistentCloudKitContainerOptions](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentCloudKitContainerOptions) Autorelease() NSPersistentCloudKitContainerOptions {
	rv := objc.Send[NSPersistentCloudKitContainerOptions](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentCloudKitContainerOptions creates a new NSPersistentCloudKitContainerOptions instance.
func NewNSPersistentCloudKitContainerOptions() NSPersistentCloudKitContainerOptions {
	class := getNSPersistentCloudKitContainerOptionsClass()
	rv := objc.Send[NSPersistentCloudKitContainerOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes container options using the given CloudKit container
// identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/init(containerIdentifier:)
func NewPersistentCloudKitContainerOptionsWithContainerIdentifier(containerIdentifier string) NSPersistentCloudKitContainerOptions {
	instance := getNSPersistentCloudKitContainerOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerIdentifier:"), objc.String(containerIdentifier))
	return NSPersistentCloudKitContainerOptionsFromID(rv)
}

// Initializes container options using the given CloudKit container
// identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/init(containerIdentifier:)
func (p NSPersistentCloudKitContainerOptions) InitWithContainerIdentifier(containerIdentifier string) NSPersistentCloudKitContainerOptions {
	rv := objc.Send[NSPersistentCloudKitContainerOptions](p.ID, objc.Sel("initWithContainerIdentifier:"), objc.String(containerIdentifier))
	return rv
}

// The identifier of the CloudKit container associated with a given store
// description.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/containerIdentifier
func (p NSPersistentCloudKitContainerOptions) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The database scope — public, private, or shared — to use for a
// specified store in a persistent CloudKit container.
//
// See: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontaineroptions/databasescope-4c72t
func (p NSPersistentCloudKitContainerOptions) DatabaseScope() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("databaseScope"))
	return rv
}
func (p NSPersistentCloudKitContainerOptions) SetDatabaseScope(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setDatabaseScope:"), value)
}
