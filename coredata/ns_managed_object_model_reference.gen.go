// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSManagedObjectModelReference] class.
var (
	_NSManagedObjectModelReferenceClass     NSManagedObjectModelReferenceClass
	_NSManagedObjectModelReferenceClassOnce sync.Once
)

func getNSManagedObjectModelReferenceClass() NSManagedObjectModelReferenceClass {
	_NSManagedObjectModelReferenceClassOnce.Do(func() {
		_NSManagedObjectModelReferenceClass = NSManagedObjectModelReferenceClass{class: objc.GetClass("NSManagedObjectModelReference")}
	})
	return _NSManagedObjectModelReferenceClass
}

// GetNSManagedObjectModelReferenceClass returns the class object for NSManagedObjectModelReference.
func GetNSManagedObjectModelReferenceClass() NSManagedObjectModelReferenceClass {
	return getNSManagedObjectModelReferenceClass()
}

type NSManagedObjectModelReferenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSManagedObjectModelReferenceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSManagedObjectModelReferenceClass) Alloc() NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a specific version of an object model.
//
// # Creating a reference
//
//   - [NSManagedObjectModelReference.InitWithModelVersionChecksum]: Creates an object model reference for the specified model.
//   - [NSManagedObjectModelReference.InitWithFileURLVersionChecksum]: Creates an object model reference for the model at the specified file URL.
//   - [NSManagedObjectModelReference.InitWithNameInBundleVersionChecksum]: Creates an object model reference for the named model in the specified bundle.
//   - [NSManagedObjectModelReference.InitWithEntityVersionHashesInBundleVersionChecksum]: Creates an object model reference with the entities corresponding to the specified entity version hashes.
//
// # Resolving the model object
//
//   - [NSManagedObjectModelReference.ResolvedModel]: The resolved object model.
//   - [NSManagedObjectModelReference.VersionChecksum]: The version checksum of the resolved model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference
type NSManagedObjectModelReference struct {
	objectivec.Object
}

// NSManagedObjectModelReferenceFromID constructs a [NSManagedObjectModelReference] from an objc.ID.
//
// An object that describes a specific version of an object model.
func NSManagedObjectModelReferenceFromID(id objc.ID) NSManagedObjectModelReference {
	return NSManagedObjectModelReference{objectivec.Object{ID: id}}
}

// NOTE: NSManagedObjectModelReference adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSManagedObjectModelReference] class.
//
// # Creating a reference
//
//   - [INSManagedObjectModelReference.InitWithModelVersionChecksum]: Creates an object model reference for the specified model.
//   - [INSManagedObjectModelReference.InitWithFileURLVersionChecksum]: Creates an object model reference for the model at the specified file URL.
//   - [INSManagedObjectModelReference.InitWithNameInBundleVersionChecksum]: Creates an object model reference for the named model in the specified bundle.
//   - [INSManagedObjectModelReference.InitWithEntityVersionHashesInBundleVersionChecksum]: Creates an object model reference with the entities corresponding to the specified entity version hashes.
//
// # Resolving the model object
//
//   - [INSManagedObjectModelReference.ResolvedModel]: The resolved object model.
//   - [INSManagedObjectModelReference.VersionChecksum]: The version checksum of the resolved model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference
type INSManagedObjectModelReference interface {
	objectivec.IObject

	// Topic: Creating a reference

	// Creates an object model reference for the specified model.
	InitWithModelVersionChecksum(model INSManagedObjectModel, versionChecksum string) NSManagedObjectModelReference
	// Creates an object model reference for the model at the specified file URL.
	InitWithFileURLVersionChecksum(fileURL foundation.NSURL, versionChecksum string) NSManagedObjectModelReference
	// Creates an object model reference for the named model in the specified bundle.
	InitWithNameInBundleVersionChecksum(modelName string, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference
	// Creates an object model reference with the entities corresponding to the specified entity version hashes.
	InitWithEntityVersionHashesInBundleVersionChecksum(versionHash foundation.INSDictionary, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference

	// Topic: Resolving the model object

	// The resolved object model.
	ResolvedModel() INSManagedObjectModel
	// The version checksum of the resolved model.
	VersionChecksum() string
}

// Init initializes the instance.
func (m NSManagedObjectModelReference) Init() NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSManagedObjectModelReference) Autorelease() NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSManagedObjectModelReference creates a new NSManagedObjectModelReference instance.
func NewNSManagedObjectModelReference() NSManagedObjectModelReference {
	class := getNSManagedObjectModelReferenceClass()
	rv := objc.Send[NSManagedObjectModelReference](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object model reference with the entities corresponding to the
// specified entity version hashes.
//
// versionHash: The dictionary of entity names and their corresponding version hashes.
//
// bundle: The bundle to search.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(entityVersionHashes:in:versionChecksum:)
func NewManagedObjectModelReferenceWithEntityVersionHashesInBundleVersionChecksum(versionHash foundation.INSDictionary, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference {
	instance := getNSManagedObjectModelReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEntityVersionHashes:inBundle:versionChecksum:"), versionHash, bundle, objc.String(versionChecksum))
	return NSManagedObjectModelReferenceFromID(rv)
}

// Creates an object model reference for the model at the specified file URL.
//
// fileURL: The on-disk location of the managed object model.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(fileURL:versionChecksum:)
func NewManagedObjectModelReferenceWithFileURLVersionChecksum(fileURL foundation.NSURL, versionChecksum string) NSManagedObjectModelReference {
	instance := getNSManagedObjectModelReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileURL:versionChecksum:"), fileURL, objc.String(versionChecksum))
	return NSManagedObjectModelReferenceFromID(rv)
}

// Creates an object model reference for the specified model.
//
// model: The managed object model.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(model:versionChecksum:)
func NewManagedObjectModelReferenceWithModelVersionChecksum(model INSManagedObjectModel, versionChecksum string) NSManagedObjectModelReference {
	instance := getNSManagedObjectModelReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModel:versionChecksum:"), model, objc.String(versionChecksum))
	return NSManagedObjectModelReferenceFromID(rv)
}

// Creates an object model reference for the named model in the specified
// bundle.
//
// modelName: The name of the managed object model in the specified bundle.
//
// bundle: The bundle to search.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(name:in:versionChecksum:)
func NewManagedObjectModelReferenceWithNameInBundleVersionChecksum(modelName string, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference {
	instance := getNSManagedObjectModelReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inBundle:versionChecksum:"), objc.String(modelName), bundle, objc.String(versionChecksum))
	return NSManagedObjectModelReferenceFromID(rv)
}

// Creates an object model reference for the specified model.
//
// model: The managed object model.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(model:versionChecksum:)
func (m NSManagedObjectModelReference) InitWithModelVersionChecksum(model INSManagedObjectModel, versionChecksum string) NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("initWithModel:versionChecksum:"), model, objc.String(versionChecksum))
	return rv
}

// Creates an object model reference for the model at the specified file URL.
//
// fileURL: The on-disk location of the managed object model.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(fileURL:versionChecksum:)
func (m NSManagedObjectModelReference) InitWithFileURLVersionChecksum(fileURL foundation.NSURL, versionChecksum string) NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("initWithFileURL:versionChecksum:"), fileURL, objc.String(versionChecksum))
	return rv
}

// Creates an object model reference for the named model in the specified
// bundle.
//
// modelName: The name of the managed object model in the specified bundle.
//
// bundle: The bundle to search.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(name:in:versionChecksum:)
func (m NSManagedObjectModelReference) InitWithNameInBundleVersionChecksum(modelName string, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("initWithName:inBundle:versionChecksum:"), objc.String(modelName), bundle, objc.String(versionChecksum))
	return rv
}

// Creates an object model reference with the entities corresponding to the
// specified entity version hashes.
//
// versionHash: The dictionary of entity names and their corresponding version hashes.
//
// bundle: The bundle to search.
//
// versionChecksum: The checksum of the object model’s version.
//
// # Discussion
//
// To determine an object model’s version checksum, use its
// [NSManagedObjectModel.VersionChecksum] property. Alternatively, you can
// find the checksum in the versioned model’s `VersionInfo.Plist()` file or
// in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/init(entityVersionHashes:in:versionChecksum:)
func (m NSManagedObjectModelReference) InitWithEntityVersionHashesInBundleVersionChecksum(versionHash foundation.INSDictionary, bundle foundation.NSBundle, versionChecksum string) NSManagedObjectModelReference {
	rv := objc.Send[NSManagedObjectModelReference](m.ID, objc.Sel("initWithEntityVersionHashes:inBundle:versionChecksum:"), versionHash, bundle, objc.String(versionChecksum))
	return rv
}

// The resolved object model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/resolvedModel
func (m NSManagedObjectModelReference) ResolvedModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resolvedModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The version checksum of the resolved model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference/versionChecksum
func (m NSManagedObjectModelReference) VersionChecksum() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("versionChecksum"))
	return foundation.NSStringFromID(rv).String()
}
