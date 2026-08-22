// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSManagedObjectModel] class.
var (
	_NSManagedObjectModelClass     NSManagedObjectModelClass
	_NSManagedObjectModelClassOnce sync.Once
)

func getNSManagedObjectModelClass() NSManagedObjectModelClass {
	_NSManagedObjectModelClassOnce.Do(func() {
		_NSManagedObjectModelClass = NSManagedObjectModelClass{class: objc.GetClass("NSManagedObjectModel")}
	})
	return _NSManagedObjectModelClass
}

// GetNSManagedObjectModelClass returns the class object for NSManagedObjectModel.
func GetNSManagedObjectModelClass() NSManagedObjectModelClass {
	return getNSManagedObjectModelClass()
}

type NSManagedObjectModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSManagedObjectModelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSManagedObjectModelClass) Alloc() NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A programmatic representation of the `XCUIElementTypeXcdatamodeld` file
// describing your objects.
//
// # Overview
//
// The model contains one or more [NSEntityDescription] objects representing
// the entities in the schema. Each [NSEntityDescription] object has property
// description objects (instances of subclasses of [NSPropertyDescription])
// that represent the properties (or fields) of the entity in the schema. The
// Core Data framework uses this description in several ways:
//
// - Constraining UI creation in Interface Builder - Validating attribute and
// relationship values at runtime - Mapping between your managed objects and a
// database or file-based schema for object persistence
//
// A managed object model maintains a mapping between each of its entity
// objects and a corresponding managed object class for use with the
// persistent storage mechanisms in the Core Data framework. You can determine
// the entity for a particular managed object with the `entity` method.
//
// You typically create managed object models using the data modeling tool in
// Xcode, but it’s possible to build a model programmatically if needed.
//
// # Loading a model file
//
// Managed object model files are typically stored in a project or a
// framework. To load a model, you provide an URL to the constructor. Note
// that loading a model doesn’t have the effect of loading all of its
// entities.
//
// # Storing fetch requests
//
// Frequently, you need a collection of objects that share features in common.
// Sometimes you can define those features (property values) in advance;
// sometimes you need to be able to supply values at runtime. For example,
// suppose you want to retrieve all movies owned by Pixar, or retrieve all
// movies that earned more than an amount specified by the user at runtime.
//
// Fetch requests are often predefined in a managed object model as templates.
// They allow you to predefine named queries and their parameters in the
// model. Typically they contain variables that need to be substituted at
// runtime. [NSManagedObjectModel] provides an API to retrieve a stored fetch
// request by name, and to perform variable substitution—see
// [NSManagedObjectModel.FetchRequestTemplateForName] and
// [NSManagedObjectModel.FetchRequestFromTemplateWithNameSubstitutionVariables].
//
// You typically define fetch request templates using the Data Model editor in
// Xcode. You can also create fetch request templates programmatically, and
// associate them with a model using
// [NSManagedObjectModel.SetFetchRequestTemplateForName].
//
// # Supporting multiple configurations for the same model
//
// You may want to specify different sets of entities for the same model to be
// used in different situations. For example, suppose certain entities should
// only be available if a user has administrative privileges. To support this
// requirement, a model may have more than one configuration. Each
// configuration is named, and has an associated set of entities. The sets may
// overlap. You establish configurations programmatically using
// [NSManagedObjectModel.SetEntitiesForConfiguration] or using the Xcode
// design tool, and retrieve the entities for a given configuration name using
// [NSManagedObjectModel.EntitiesForConfiguration].
//
// # Changing models
//
// Because a model describes the structure of the data in a persistent store,
// changing any parts of a model that alters the schema renders it
// incompatible with (and so unable to open) the stores it previously created.
// If you change your schema, you therefore need to migrate the data in
// existing stores to new version (see [Core Data Model Versioning and Data
// Migration Programming Guide]). For example, if you add a new entity or a
// new attribute to an existing entity, you can’t open old stores; if you
// add a validation constraint or set a new default value for an attribute,
// you can open old stores.
//
// # Editing models at runtime
//
// Managed object models are editable until they are used by an object graph
// manager (a managed object context or a persistent store coordinator). This
// allows you to create or modify them dynamically until their first use.
// However, once a model is being used, it must not be changed. This is
// enforced at runtime—when the object manager first fetches data using a
// model, the whole of that model becomes uneditable. Any attempt to mutate a
// model or any of its sub-objects after that point throws an exception. If
// you need to modify a model that’s in use, create a copy, modify the copy,
// and then discard the objects with the old model.
//
// # Enumerating entities with fast enumeration
//
// In macOS 10.5 and later and on iOS, [NSManagedObjectModel] supports the
// [NSFastEnumeration] protocol. You can use this to enumerate over a
// model’s entities, as illustrated in the following example:
//
// # Creating a managed object model
//
//   - [NSManagedObjectModel.InitWithContentsOfURL]: Initializes the managed object model using the model file at the specified URL.
//
// # Managing entities and configurations
//
//   - [NSManagedObjectModel.Entities]: The entities in the model.
//   - [NSManagedObjectModel.SetEntities]
//   - [NSManagedObjectModel.EntitiesByName]: The entities of the model, keyed by name.
//   - [NSManagedObjectModel.Configurations]: All the available configuration names of the model.
//   - [NSManagedObjectModel.EntitiesForConfiguration]: Returns the entities of the model for a specified configuration.
//   - [NSManagedObjectModel.SetEntitiesForConfiguration]: Associates the specified entities with the model using the given configuration name.
//
// # Manipulating fetch request templates
//
//   - [NSManagedObjectModel.FetchRequestTemplatesByName]: A dictionary of the receiver’s fetch request templates, keyed by name.
//   - [NSManagedObjectModel.FetchRequestTemplateForName]: Returns the fetch request with a specified name.
//   - [NSManagedObjectModel.FetchRequestFromTemplateWithNameSubstitutionVariables]: Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary.
//   - [NSManagedObjectModel.SetFetchRequestTemplateForName]: Associates the specified fetch request with the receiver using the given name.
//
// # Handling localization
//
//   - [NSManagedObjectModel.LocalizationDictionary]: The localization dictionary of the model.
//   - [NSManagedObjectModel.SetLocalizationDictionary]
//
// # Versioning and migrating entities
//
//   - [NSManagedObjectModel.VersionChecksum]: The Base64-encoded 128-bit model version hash.
//   - [NSManagedObjectModel.VersionIdentifiers]: The set of developer-defined version identifiers for the object model.
//   - [NSManagedObjectModel.SetVersionIdentifiers]
//   - [NSManagedObjectModel.EntityVersionHashesByName]: The dictionary of the model’s entity names and their corresponding version hashes.
//   - [NSManagedObjectModel.IsConfigurationCompatibleWithStoreMetadata]: Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel
//
// [Core Data Model Versioning and Data Migration Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreDataVersioning/Articles/Introduction.html#//apple_ref/doc/uid/TP40004399
// [NSFastEnumeration]: https://developer.apple.com/documentation/Foundation/NSFastEnumeration
type NSManagedObjectModel struct {
	objectivec.Object
}

// NSManagedObjectModelFromID constructs a [NSManagedObjectModel] from an objc.ID.
//
// A programmatic representation of the `XCUIElementTypeXcdatamodeld` file
// describing your objects.
func NSManagedObjectModelFromID(id objc.ID) NSManagedObjectModel {
	return NSManagedObjectModel{objectivec.Object{ID: id}}
}

// NOTE: NSManagedObjectModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSManagedObjectModel] class.
//
// # Creating a managed object model
//
//   - [INSManagedObjectModel.InitWithContentsOfURL]: Initializes the managed object model using the model file at the specified URL.
//
// # Managing entities and configurations
//
//   - [INSManagedObjectModel.Entities]: The entities in the model.
//   - [INSManagedObjectModel.SetEntities]
//   - [INSManagedObjectModel.EntitiesByName]: The entities of the model, keyed by name.
//   - [INSManagedObjectModel.Configurations]: All the available configuration names of the model.
//   - [INSManagedObjectModel.EntitiesForConfiguration]: Returns the entities of the model for a specified configuration.
//   - [INSManagedObjectModel.SetEntitiesForConfiguration]: Associates the specified entities with the model using the given configuration name.
//
// # Manipulating fetch request templates
//
//   - [INSManagedObjectModel.FetchRequestTemplatesByName]: A dictionary of the receiver’s fetch request templates, keyed by name.
//   - [INSManagedObjectModel.FetchRequestTemplateForName]: Returns the fetch request with a specified name.
//   - [INSManagedObjectModel.FetchRequestFromTemplateWithNameSubstitutionVariables]: Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary.
//   - [INSManagedObjectModel.SetFetchRequestTemplateForName]: Associates the specified fetch request with the receiver using the given name.
//
// # Handling localization
//
//   - [INSManagedObjectModel.LocalizationDictionary]: The localization dictionary of the model.
//   - [INSManagedObjectModel.SetLocalizationDictionary]
//
// # Versioning and migrating entities
//
//   - [INSManagedObjectModel.VersionChecksum]: The Base64-encoded 128-bit model version hash.
//   - [INSManagedObjectModel.VersionIdentifiers]: The set of developer-defined version identifiers for the object model.
//   - [INSManagedObjectModel.SetVersionIdentifiers]
//   - [INSManagedObjectModel.EntityVersionHashesByName]: The dictionary of the model’s entity names and their corresponding version hashes.
//   - [INSManagedObjectModel.IsConfigurationCompatibleWithStoreMetadata]: Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel
type INSManagedObjectModel interface {
	objectivec.IObject

	// Topic: Creating a managed object model

	// Initializes the managed object model using the model file at the specified URL.
	InitWithContentsOfURL(url foundation.NSURL) NSManagedObjectModel

	// Topic: Managing entities and configurations

	// The entities in the model.
	Entities() []NSEntityDescription
	SetEntities(value []NSEntityDescription)
	// The entities of the model, keyed by name.
	EntitiesByName() foundation.INSDictionary
	// All the available configuration names of the model.
	Configurations() []string
	// Returns the entities of the model for a specified configuration.
	EntitiesForConfiguration(configuration string) []NSEntityDescription
	// Associates the specified entities with the model using the given configuration name.
	SetEntitiesForConfiguration(entities []NSEntityDescription, configuration string)

	// Topic: Manipulating fetch request templates

	// A dictionary of the receiver’s fetch request templates, keyed by name.
	FetchRequestTemplatesByName() foundation.INSDictionary
	// Returns the fetch request with a specified name.
	FetchRequestTemplateForName(name string) INSFetchRequest
	// Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary.
	FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables foundation.INSDictionary) INSFetchRequest
	// Associates the specified fetch request with the receiver using the given name.
	SetFetchRequestTemplateForName(fetchRequestTemplate INSFetchRequest, name string)

	// Topic: Handling localization

	// The localization dictionary of the model.
	LocalizationDictionary() foundation.INSDictionary
	SetLocalizationDictionary(value foundation.INSDictionary)

	// Topic: Versioning and migrating entities

	// The Base64-encoded 128-bit model version hash.
	VersionChecksum() string
	// The set of developer-defined version identifiers for the object model.
	VersionIdentifiers() foundation.INSSet
	SetVersionIdentifiers(value foundation.INSSet)
	// The dictionary of the model’s entity names and their corresponding version hashes.
	EntityVersionHashesByName() foundation.INSDictionary
	// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
	IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata foundation.INSDictionary) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m NSManagedObjectModel) Init() NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSManagedObjectModel) Autorelease() NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSManagedObjectModel creates a new NSManagedObjectModel instance.
func NewNSManagedObjectModel() NSManagedObjectModel {
	class := getNSManagedObjectModelClass()
	rv := objc.Send[NSManagedObjectModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a single model from an array of existing models.
//
// models: An array of instances of [NSManagedObjectModel].
//
// # Return Value
//
// A single model made by combining the models in `models`.
//
// # Discussion
//
// You use this method to combine multiple models (typically from different
// frameworks) into one.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/init(byMerging:)
func NewManagedObjectModelByMergingModels(models []NSManagedObjectModel) NSManagedObjectModel {
	rv := objc.Send[objc.ID](objc.ID(getNSManagedObjectModelClass().class), objc.Sel("modelByMergingModels:"), objectivec.IObjectSliceToNSArray(models))
	return NSManagedObjectModelFromID(rv)
}

// Returns, for the version information in given metadata, a model merged from
// a given array of models.
//
// models: An array of instances of [NSManagedObjectModel].
//
// metadata: A dictionary containing version information from the metadata for a
// persistent store.
//
// # Return Value
//
// A merged model from `models` for the version information in `metadata`. If
// a model cannot be created to match the version information in `metadata`,
// returns `nil`.
//
// # Discussion
//
// This is the companion method to
// [NSManagedObjectModelClass.MergedModelFromBundlesForStoreMetadata].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/init(byMerging:forStoreMetadata:)
func NewManagedObjectModelByMergingModelsForStoreMetadata(models []NSManagedObjectModel, metadata foundation.INSDictionary) NSManagedObjectModel {
	rv := objc.Send[objc.ID](objc.ID(getNSManagedObjectModelClass().class), objc.Sel("modelByMergingModels:forStoreMetadata:"), objectivec.IObjectSliceToNSArray(models), metadata)
	return NSManagedObjectModelFromID(rv)
}

// Initializes the managed object model using the model file at the specified
// URL.
//
// url: An URL object specifying the location of a model file.
//
// # Return Value
//
// A managed object model initialized using the file at `url`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/init(contentsOf:)
func NewManagedObjectModelWithContentsOfURL(url foundation.NSURL) NSManagedObjectModel {
	instance := getNSManagedObjectModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSManagedObjectModelFromID(rv)
}

// Initializes the managed object model using the model file at the specified
// URL.
//
// url: An URL object specifying the location of a model file.
//
// # Return Value
//
// A managed object model initialized using the file at `url`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/init(contentsOf:)
func (m NSManagedObjectModel) InitWithContentsOfURL(url foundation.NSURL) NSManagedObjectModel {
	rv := objc.Send[NSManagedObjectModel](m.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}

// Returns the entities of the model for a specified configuration.
//
// configuration: The name of a configuration in the receiver.
//
// # Return Value
//
// An array containing the entities of the receiver for the configuration
// specified by `configuration`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/entities(forConfigurationName:)
func (m NSManagedObjectModel) EntitiesForConfiguration(configuration string) []NSEntityDescription {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("entitiesForConfiguration:"), objc.String(configuration))
	return objc.ConvertSlice(rv, func(id objc.ID) NSEntityDescription {
		return NSEntityDescriptionFromID(id)
	})
}

// Associates the specified entities with the model using the given
// configuration name.
//
// entities: An array of instances of [NSEntityDescription].
//
// configuration: A name for the configuration.
//
// # Discussion
//
// This method raises an exception if the receiver has been used by an object
// graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setEntities(_:forConfigurationName:)
func (m NSManagedObjectModel) SetEntitiesForConfiguration(entities []NSEntityDescription, configuration string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setEntities:forConfiguration:"), objectivec.IObjectSliceToNSArray(entities), objc.String(configuration))
}

// Returns the fetch request with a specified name.
//
// name: A string containing the name of a fetch request template.
//
// # Return Value
//
// The fetch request named `name`.
//
// # Discussion
//
// If the template contains substitution variables, you should instead use
// [NSManagedObjectModel.FetchRequestFromTemplateWithNameSubstitutionVariables]
// to create a new fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestTemplate(forName:)
func (m NSManagedObjectModel) FetchRequestTemplateForName(name string) INSFetchRequest {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fetchRequestTemplateForName:"), objc.String(name))
	return NSFetchRequestFromID(rv)
}

// Returns a copy of the fetch request template with the variables substituted
// by values from the substitutions dictionary.
//
// name: A string containing the name of a fetch request template.
//
// variables: A dictionary containing key-value pairs where the keys are the names of
// variables specified in the template; the corresponding values are
// substituted before the fetch request is returned. The dictionary must
// provide values for all the variables in the template.
//
// # Return Value
//
// A copy of the fetch request template with the variables substituted by
// values from `variables`.
//
// # Discussion
//
// The `variables` dictionary must provide values for all the variables. If
// you want to test for a nil value, use `[NSNull null]`.
//
// This method provides the usual way to bind an “abstractly” defined
// fetch request template to a concrete fetch. For more details on using this
// method, see [Creating Predicates].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestFromTemplate(withName:substitutionVariables:)
//
// [Creating Predicates]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pCreating.html#//apple_ref/doc/uid/TP40001793
func (m NSManagedObjectModel) FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables foundation.INSDictionary) INSFetchRequest {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fetchRequestFromTemplateWithName:substitutionVariables:"), objc.String(name), variables)
	return NSFetchRequestFromID(rv)
}

// Associates the specified fetch request with the receiver using the given
// name.
//
// fetchRequestTemplate: A fetch request, typically containing predicates with variables for
// substitution.
//
// name: A string that specifies the name of the fetch request template.
//
// # Discussion
//
// For more details on using this method, see [Creating Predicates].
//
// # Special Considerations
//
// This method raises an exception if the receiver has been used by an object
// graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setFetchRequestTemplate(_:forName:)
//
// [Creating Predicates]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/Articles/pCreating.html#//apple_ref/doc/uid/TP40001793
func (m NSManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate INSFetchRequest, name string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setFetchRequestTemplate:forName:"), fetchRequestTemplate, objc.String(name))
}

// Returns a Boolean value that indicates whether a given configuration in the
// model is compatible with given metadata from a persistent store.
//
// configuration: The name of a configuration in the receiver. Pass `nil` to specify no
// configuration.
//
// metadata: Metadata for a persistent store.
//
// # Return Value
//
// true if the configuration in the receiver specified by `configuration` is
// compatible with the store metadata given by `metadata`, otherwise false.
//
// # Discussion
//
// This method compares the version information in the store metadata with the
// entity versions of a given configuration. For information on specific
// differences, use [NSManagedObjectModel.EntityVersionHashesByName] and
// perform an entity-by-entity comparison.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/isConfiguration(withName:compatibleWithStoreMetadata:)
func (m NSManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata foundation.INSDictionary) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), objc.String(configuration), metadata)
	return rv
}
func (m NSManagedObjectModel) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a model created by merging all the models found in given bundles.
//
// bundles: An array of instances of [NSBundle] to search. If you specify `nil`, then
// the main bundle is searched.
//
// # Return Value
//
// A model created by merging all the models found in `bundles`.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/mergedModel(from:)
func (_NSManagedObjectModelClass NSManagedObjectModelClass) MergedModelFromBundles(bundles []foundation.NSBundle) NSManagedObjectModel {
	rv := objc.Send[objc.ID](objc.ID(_NSManagedObjectModelClass.class), objc.Sel("mergedModelFromBundles:"), objectivec.IObjectSliceToNSArray(bundles))
	return NSManagedObjectModelFromID(rv)
}

// Returns a merged model from a specified array for the version information
// in provided metadata.
//
// bundles: An array of bundles.
//
// metadata: A dictionary containing version information from the metadata for a
// persistent store.
//
// # Return Value
//
// The managed object model used to create the store for the metadata. If a
// model cannot be created to match the version information specified by
// `metadata`, returns `nil`.
//
// # Discussion
//
// This method is a companion to
// [NSManagedObjectModelClass.MergedModelFromBundles].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/mergedModel(from:forStoreMetadata:)
func (_NSManagedObjectModelClass NSManagedObjectModelClass) MergedModelFromBundlesForStoreMetadata(bundles []foundation.NSBundle, metadata foundation.INSDictionary) NSManagedObjectModel {
	rv := objc.Send[objc.ID](objc.ID(_NSManagedObjectModelClass.class), objc.Sel("mergedModelFromBundles:forStoreMetadata:"), objectivec.IObjectSliceToNSArray(bundles), metadata)
	return NSManagedObjectModelFromID(rv)
}

// The entities in the model.
//
// # Discussion
//
// Entities are instances of [NSEntityDescription].
//
// # Special Considerations
//
// Setting the entities for an object model raises an exception if the object
// model has been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/entities
func (m NSManagedObjectModel) Entities() []NSEntityDescription {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("entities"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSEntityDescription {
		return NSEntityDescriptionFromID(id)
	})
}
func (m NSManagedObjectModel) SetEntities(value []NSEntityDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setEntities:"), objectivec.IObjectSliceToNSArray(value))
}

// The entities of the model, keyed by name.
//
// # Discussion
//
// Entities are instances of [NSEntityDescription].
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/entitiesByName
func (m NSManagedObjectModel) EntitiesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entitiesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// All the available configuration names of the model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/configurations
func (m NSManagedObjectModel) Configurations() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("configurations"))
	return objc.ConvertSliceToStrings(rv)
}

// A dictionary of the receiver’s fetch request templates, keyed by name.
//
// # Discussion
//
// If the template contains a predicate with substitution variables, you
// should instead use
// [NSManagedObjectModel.FetchRequestFromTemplateWithNameSubstitutionVariables]
// to create a new fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestTemplatesByName
func (m NSManagedObjectModel) FetchRequestTemplatesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fetchRequestTemplatesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The localization dictionary of the model.
//
// # Discussion
//
// The following table describes the key and value pattern for the
// localization dictionary.
//
// [Table data omitted]
//
// (1) For properties in different entities with the same non-localized name
// but that should have different localized names.
//
// # Special Considerations
//
// In OS X v10.4, `localizationDictionary` may return `nil` until Core Data
// lazily loads the dictionary for its own purposes (for example, reporting a
// localized error).
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/localizationDictionary
func (m NSManagedObjectModel) LocalizationDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("localizationDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m NSManagedObjectModel) SetLocalizationDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setLocalizationDictionary:"), value)
}

// The Base64-encoded 128-bit model version hash.
//
// # Discussion
//
// This value is also available in the versioned model’s
// `VersionInfo.Plist()` file and in Xcode’s build log.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/versionChecksum
func (m NSManagedObjectModel) VersionChecksum() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("versionChecksum"))
	return foundation.NSStringFromID(rv).String()
}

// The set of developer-defined version identifiers for the object model.
//
// # Discussion
//
// Merged models return the combined collection of identifiers. The Core Data
// framework does not assign a default identifier to object models, nor does
// it depend on this value at runtime. For models you create in Xcode, set
// this value in the model inspector.
//
// Use this value when debugging to help determine the models that Core Data
// merges to create the merged model.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/versionIdentifiers
func (m NSManagedObjectModel) VersionIdentifiers() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("versionIdentifiers"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (m NSManagedObjectModel) SetVersionIdentifiers(value foundation.INSSet) {
	objc.Send[struct{}](m.ID, objc.Sel("setVersionIdentifiers:"), value)
}

// The dictionary of the model’s entity names and their corresponding
// version hashes.
//
// # Discussion
//
// Core Data use the dictionary of version hash information is to determine
// schema compatibility.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/entityVersionHashesByName
func (m NSManagedObjectModel) EntityVersionHashesByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entityVersionHashesByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
