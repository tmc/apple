// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMappingModel] class.
var (
	_NSMappingModelClass     NSMappingModelClass
	_NSMappingModelClassOnce sync.Once
)

func getNSMappingModelClass() NSMappingModelClass {
	_NSMappingModelClassOnce.Do(func() {
		_NSMappingModelClass = NSMappingModelClass{class: objc.GetClass("NSMappingModel")}
	})
	return _NSMappingModelClass
}

// GetNSMappingModelClass returns the class object for NSMappingModel.
func GetNSMappingModelClass() NSMappingModelClass {
	return getNSMappingModelClass()
}

type NSMappingModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMappingModelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMappingModelClass) Alloc() NSMappingModel {
	rv := objc.Send[NSMappingModel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A model instance that specifies how to map a model from a source to a
// destination managed object model.
//
// # Creating a Mapping
//
//   - [NSMappingModel.InitWithContentsOfURL]: Returns a mapping model initialized from a given URL.
//
// # Managing Entity Mappings
//
//   - [NSMappingModel.EntityMappings]: The entity mappings for the mapping model.
//   - [NSMappingModel.SetEntityMappings]
//   - [NSMappingModel.EntityMappingsByName]: The entity mappings for the mapping model, keyed by name.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel
type NSMappingModel struct {
	objectivec.Object
}

// NSMappingModelFromID constructs a [NSMappingModel] from an objc.ID.
//
// A model instance that specifies how to map a model from a source to a
// destination managed object model.
func NSMappingModelFromID(id objc.ID) NSMappingModel {
	return NSMappingModel{objectivec.Object{ID: id}}
}

// NOTE: NSMappingModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMappingModel] class.
//
// # Creating a Mapping
//
//   - [INSMappingModel.InitWithContentsOfURL]: Returns a mapping model initialized from a given URL.
//
// # Managing Entity Mappings
//
//   - [INSMappingModel.EntityMappings]: The entity mappings for the mapping model.
//   - [INSMappingModel.SetEntityMappings]
//   - [INSMappingModel.EntityMappingsByName]: The entity mappings for the mapping model, keyed by name.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel
type INSMappingModel interface {
	objectivec.IObject

	// Topic: Creating a Mapping

	// Returns a mapping model initialized from a given URL.
	InitWithContentsOfURL(url foundation.NSURL) NSMappingModel

	// Topic: Managing Entity Mappings

	// The entity mappings for the mapping model.
	EntityMappings() []NSEntityMapping
	SetEntityMappings(value []NSEntityMapping)
	// The entity mappings for the mapping model, keyed by name.
	EntityMappingsByName() foundation.INSDictionary
}

// Init initializes the instance.
func (m NSMappingModel) Init() NSMappingModel {
	rv := objc.Send[NSMappingModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMappingModel) Autorelease() NSMappingModel {
	rv := objc.Send[NSMappingModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMappingModel creates a new NSMappingModel instance.
func NewNSMappingModel() NSMappingModel {
	class := getNSMappingModelClass()
	rv := objc.Send[NSMappingModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the mapping model that will translate data from the source to the
// destination model.
//
// bundles: An array of bundles in which to search for mapping models.
//
// sourceModel: The managed object model for the source store.
//
// destinationModel: The managed object model for the destination store.
//
// # Return Value
//
// Returns the mapping model to translate data from `sourceModel` to
// `destinationModel`. If a suitable mapping model cannot be found, returns
// `nil`.
//
// # Discussion
//
// This method is a companion to the
// [NSManagedObjectModelClass.MergedModelFromBundles] and
// [NSManagedObjectModelClass.MergedModelFromBundlesForStoreMetadata] methods.
// In this case, the framework uses the version information from the models to
// locate the appropriate mapping model in the available bundles.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(from:forSourceModel:destinationModel:)
func NewMappingModelFromBundlesForSourceModelDestinationModel(bundles []foundation.NSBundle, sourceModel INSManagedObjectModel, destinationModel INSManagedObjectModel) NSMappingModel {
	rv := objc.Send[objc.ID](objc.ID(getNSMappingModelClass().class), objc.Sel("mappingModelFromBundles:forSourceModel:destinationModel:"), objectivec.IObjectSliceToNSArray(bundles), sourceModel, destinationModel)
	return NSMappingModelFromID(rv)
}

// Returns a mapping model initialized from a given URL.
//
// url: The location of an archived mapping model.
//
// # Return Value
//
// A mapping model initialized from `url`.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(contentsOf:)
func NewMappingModelWithContentsOfURL(url foundation.NSURL) NSMappingModel {
	instance := getNSMappingModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSMappingModelFromID(rv)
}

// Returns a mapping model initialized from a given URL.
//
// url: The location of an archived mapping model.
//
// # Return Value
//
// A mapping model initialized from `url`.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/init(contentsOf:)
func (m NSMappingModel) InitWithContentsOfURL(url foundation.NSURL) NSMappingModel {
	rv := objc.Send[NSMappingModel](m.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}

// Returns a newly created mapping model that will migrate data from the
// source to the destination model.
//
// sourceModel: The source managed object model.
//
// destinationModel: The destination managed object model.
//
// # Return Value
//
// A newly-created mapping model to migrate data from the source to the
// destination model. If the mapping model can not be created, returns `nil`.
//
// # Discussion
//
// A model will be created only if all changes are simple enough to be able to
// reasonably infer a mapping (for example, removing or renaming an attribute,
// adding an optional attribute or relationship, or adding renaming or
// deleting an entity). Element IDs are used to track renamed properties and
// entities.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/inferredMappingModel(forSourceModel:destinationModel:)
func (_NSMappingModelClass NSMappingModelClass) InferredMappingModelForSourceModelDestinationModelError(sourceModel INSManagedObjectModel, destinationModel INSManagedObjectModel) (NSMappingModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSMappingModelClass.class), objc.Sel("inferredMappingModelForSourceModel:destinationModel:error:"), sourceModel, destinationModel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMappingModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSMappingModelFromID(rv), nil

}

// The entity mappings for the mapping model.
//
// # Discussion
//
// The order of the mappings in the array determines the order in which they
// will be processed during migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/entityMappings
func (m NSMappingModel) EntityMappings() []NSEntityMapping {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("entityMappings"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSEntityMapping {
		return NSEntityMappingFromID(id)
	})
}
func (m NSMappingModel) SetEntityMappings(value []NSEntityMapping) {
	objc.Send[struct{}](m.ID, objc.Sel("setEntityMappings:"), objectivec.IObjectSliceToNSArray(value))
}

// The entity mappings for the mapping model, keyed by name.
//
// See: https://developer.apple.com/documentation/CoreData/NSMappingModel/entityMappingsByName
func (m NSMappingModel) EntityMappingsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entityMappingsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
