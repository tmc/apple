// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIFilterGenerator] class.
var (
	_CIFilterGeneratorClass     CIFilterGeneratorClass
	_CIFilterGeneratorClassOnce sync.Once
)

func getCIFilterGeneratorClass() CIFilterGeneratorClass {
	_CIFilterGeneratorClassOnce.Do(func() {
		_CIFilterGeneratorClass = CIFilterGeneratorClass{class: objc.GetClass("CIFilterGenerator")}
	})
	return _CIFilterGeneratorClass
}

// GetCIFilterGeneratorClass returns the class object for CIFilterGenerator.
func GetCIFilterGeneratorClass() CIFilterGeneratorClass {
	return getCIFilterGeneratorClass()
}

type CIFilterGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIFilterGeneratorClass) Alloc() CIFilterGenerator {
	rv := objc.Send[CIFilterGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that creates and configures chains of individual image filters.
//
// # Overview
// 
// The [CIFilterGenerator] class provides methods for creating a [CIFilter]
// object by chaining together existing [CIFilter] objects to create complex
// effects. (A refers to the [CIFilter] objects that are connected in the
// [CIFilterGenerator] object.) The complex effect can be encapsulated as a
// [CIFilterGenerator] object and saved as a file so that it can be used
// again. The contains an archived instance of all the [CIFilter] objects that
// are chained together.
// 
// Any filter generator files that you copy to `/Library/Graphics/Image
// Units/` are loaded when any of the loading methods provided by the
// [CIPlugIn] class are invoked. A [CIFilterGenerator] object is registered by
// its filename or, if present, by a class attribute that you supply in its
// description.
// 
// You can create a [CIFilterGenerator] object programmatically, using the
// methods provided by the [CIFilterGenerator] class, or by using the editor
// view provided by Core Image.
//
// # Initializing a Filter Generator Object
//
//   - [CIFilterGenerator.InitWithContentsOfURL]: Initializes a filter generator object with the contents of a filter generator file.
//
// # Connecting and Disconnecting Objects
//
//   - [CIFilterGenerator.ConnectObjectWithKeyToObjectWithKey]: Adds an object to the filter chain.
//   - [CIFilterGenerator.DisconnectObjectWithKeyToObjectWithKey]: Removes the connection between two objects in the filter chain.
//
// # Managing Exported Keys
//
//   - [CIFilterGenerator.ExportedKeys]: Returns an array of the exported keys.
//   - [CIFilterGenerator.ExportKeyFromObjectWithName]: Exports an input or output key of an object in the filter chain.
//   - [CIFilterGenerator.RemoveExportedKey]: Removes a key that was previously exported.
//   - [CIFilterGenerator.SetAttributesForExportedKey]: Sets a dictionary of attributes for an exported key.
//
// # Setting and Getting Class Attributes
//
//   - [CIFilterGenerator.ClassAttributes]: The class attributes associated with the filter.
//   - [CIFilterGenerator.SetClassAttributes]
//
// # Archiving a Filter Generator Object
//
//   - [CIFilterGenerator.WriteToURLAtomically]: Archives a filter generator object to a filter generator file.
//
// # Registering a Filter Chain
//
//   - [CIFilterGenerator.RegisterFilterName]: Registers the name associated with a filter chain.
//
// # Creating a Filter from a Filter Chain
//
//   - [CIFilterGenerator.Filter]: Creates a filter object based on the filter chain.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator
type CIFilterGenerator struct {
	objectivec.Object
}

// CIFilterGeneratorFromID constructs a [CIFilterGenerator] from an objc.ID.
//
// An object that creates and configures chains of individual image filters.
func CIFilterGeneratorFromID(id objc.ID) CIFilterGenerator {
	return CIFilterGenerator{objectivec.Object{ID: id}}
}
// NOTE: CIFilterGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIFilterGenerator] class.
//
// # Initializing a Filter Generator Object
//
//   - [ICIFilterGenerator.InitWithContentsOfURL]: Initializes a filter generator object with the contents of a filter generator file.
//
// # Connecting and Disconnecting Objects
//
//   - [ICIFilterGenerator.ConnectObjectWithKeyToObjectWithKey]: Adds an object to the filter chain.
//   - [ICIFilterGenerator.DisconnectObjectWithKeyToObjectWithKey]: Removes the connection between two objects in the filter chain.
//
// # Managing Exported Keys
//
//   - [ICIFilterGenerator.ExportedKeys]: Returns an array of the exported keys.
//   - [ICIFilterGenerator.ExportKeyFromObjectWithName]: Exports an input or output key of an object in the filter chain.
//   - [ICIFilterGenerator.RemoveExportedKey]: Removes a key that was previously exported.
//   - [ICIFilterGenerator.SetAttributesForExportedKey]: Sets a dictionary of attributes for an exported key.
//
// # Setting and Getting Class Attributes
//
//   - [ICIFilterGenerator.ClassAttributes]: The class attributes associated with the filter.
//   - [ICIFilterGenerator.SetClassAttributes]
//
// # Archiving a Filter Generator Object
//
//   - [ICIFilterGenerator.WriteToURLAtomically]: Archives a filter generator object to a filter generator file.
//
// # Registering a Filter Chain
//
//   - [ICIFilterGenerator.RegisterFilterName]: Registers the name associated with a filter chain.
//
// # Creating a Filter from a Filter Chain
//
//   - [ICIFilterGenerator.Filter]: Creates a filter object based on the filter chain.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator
type ICIFilterGenerator interface {
	objectivec.IObject
	CIFilterConstructor

	// Topic: Initializing a Filter Generator Object

	// Initializes a filter generator object with the contents of a filter generator file.
	InitWithContentsOfURL(aURL foundation.INSURL) CIFilterGenerator

	// Topic: Connecting and Disconnecting Objects

	// Adds an object to the filter chain.
	ConnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey string, targetObject objectivec.IObject, targetKey string)
	// Removes the connection between two objects in the filter chain.
	DisconnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey string, targetObject objectivec.IObject, targetKey string)

	// Topic: Managing Exported Keys

	// Returns an array of the exported keys.
	ExportedKeys() foundation.INSDictionary
	// Exports an input or output key of an object in the filter chain.
	ExportKeyFromObjectWithName(key string, targetObject objectivec.IObject, exportedKeyName string)
	// Removes a key that was previously exported.
	RemoveExportedKey(exportedKeyName string)
	// Sets a dictionary of attributes for an exported key.
	SetAttributesForExportedKey(attributes foundation.INSDictionary, key string)

	// Topic: Setting and Getting Class Attributes

	// The class attributes associated with the filter.
	ClassAttributes() foundation.INSDictionary
	SetClassAttributes(value foundation.INSDictionary)

	// Topic: Archiving a Filter Generator Object

	// Archives a filter generator object to a filter generator file.
	WriteToURLAtomically(aURL foundation.INSURL, flag bool) bool

	// Topic: Registering a Filter Chain

	// Registers the name associated with a filter chain.
	RegisterFilterName(name string)

	// Topic: Creating a Filter from a Filter Chain

	// Creates a filter object based on the filter chain.
	Filter() CIFilter

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f CIFilterGenerator) Init() CIFilterGenerator {
	rv := objc.Send[CIFilterGenerator](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CIFilterGenerator) Autorelease() CIFilterGenerator {
	rv := objc.Send[CIFilterGenerator](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIFilterGenerator creates a new CIFilterGenerator instance.
func NewCIFilterGenerator() CIFilterGenerator {
	class := getCIFilterGeneratorClass()
	rv := objc.Send[CIFilterGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a filter generator object with the contents of a filter
// generator file.
//
// aURL: The location of a filter generator file.
//
// # Return Value
// 
// The initialized [CIFilterGenerator] object. Returns `nil` if the file
// can’t be read.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/init(contentsOf:)
func NewFilterGeneratorWithContentsOfURL(aURL foundation.INSURL) CIFilterGenerator {
	instance := getCIFilterGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), aURL)
	return CIFilterGeneratorFromID(rv)
}

// Initializes a filter generator object with the contents of a filter
// generator file.
//
// aURL: The location of a filter generator file.
//
// # Return Value
// 
// The initialized [CIFilterGenerator] object. Returns `nil` if the file
// can’t be read.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/init(contentsOf:)
func (f CIFilterGenerator) InitWithContentsOfURL(aURL foundation.INSURL) CIFilterGenerator {
	rv := objc.Send[CIFilterGenerator](f.ID, objc.Sel("initWithContentsOfURL:"), aURL)
	return rv
}
// Adds an object to the filter chain.
//
// sourceObject: A [CIFilter] object, a [CIImage] object, or the path (an [NSString] or
// [NSURL] object) to an image.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// sourceKey: The key that specifies the source object. For example, if the source is the
// output image of a filter, pass the `outputImage` key. Pass `nil` if the
// source object is used directly.
//
// targetObject: The object to which the source object links.
//
// targetKey: The key that specifies the target for the source. For example, if you are
// connecting the source to the input image of a [CIFilter] object, you would
// pass the `inputImage` key.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/connect(_:withKey:to:withKey:)
func (f CIFilterGenerator) ConnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey string, targetObject objectivec.IObject, targetKey string) {
	objc.Send[objc.ID](f.ID, objc.Sel("connectObject:withKey:toObject:withKey:"), sourceObject, objc.String(sourceKey), targetObject, objc.String(targetKey))
}
// Removes the connection between two objects in the filter chain.
//
// sourceObject: A [CIFilter] object, a [CIImage] object, or the path (an [NSString] or
// [NSURL] object) to an image.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// sourceKey: The key that specifies the source object. Pass `nil` if the source object
// is used directly.
//
// targetObject: The object from which you want to disconnect the source object.
//
// targetKey: The key that specifies the target that the source object is currently
// connected to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/disconnectObject(_:withKey:to:withKey:)
func (f CIFilterGenerator) DisconnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey string, targetObject objectivec.IObject, targetKey string) {
	objc.Send[objc.ID](f.ID, objc.Sel("disconnectObject:withKey:toObject:withKey:"), sourceObject, objc.String(sourceKey), targetObject, objc.String(targetKey))
}
// Exports an input or output key of an object in the filter chain.
//
// key: The key to export from the target object (for example, `inputImage`).
//
// targetObject: The object associated with the key (for example, the filter).
//
// exportedKeyName: A unique name to use for the exported key. Pass `nil` to use the original
// key name.
//
// # Discussion
// 
// When you create a [CIFilter] object from a [CIFilterGenerator] object, you
// might want the filter client to be able to set some of the parameters
// associated with the filter chain. You can make a parameter settable by
// exporting the key associated with the parameter. If the exported key
// represents an input parameter of the filter, the key is exported as an
// input key. If the key represents an output parameter, it is exported as an
// output key.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/exportKey(_:from:withName:)
func (f CIFilterGenerator) ExportKeyFromObjectWithName(key string, targetObject objectivec.IObject, exportedKeyName string) {
	objc.Send[objc.ID](f.ID, objc.Sel("exportKey:fromObject:withName:"), objc.String(key), targetObject, objc.String(exportedKeyName))
}
// Removes a key that was previously exported.
//
// exportedKeyName: The name of the key you want to remove.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/removeExportedKey(_:)
func (f CIFilterGenerator) RemoveExportedKey(exportedKeyName string) {
	objc.Send[objc.ID](f.ID, objc.Sel("removeExportedKey:"), objc.String(exportedKeyName))
}
// Sets a dictionary of attributes for an exported key.
//
// attributes: A dictionary that describes the attributes associated with the specified
// key.
//
// key: The exported key whose attributes you want to set.
//
// # Discussion
// 
// By default, the exported key inherits the attributes from its original key
// and target object. You can use this method to change one or more of the
// existing attributes for the key, such as the default value or maximum
// value. For more information on attributes, see [CIFilter] and [Core Image
// Programming Guide].
//
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/setAttributes(_:forExportedKey:)
func (f CIFilterGenerator) SetAttributesForExportedKey(attributes foundation.INSDictionary, key string) {
	objc.Send[objc.ID](f.ID, objc.Sel("setAttributes:forExportedKey:"), attributes, objc.String(key))
}
// Archives a filter generator object to a filter generator file.
//
// aURL: A location for the file generator file.
//
// flag: Pass `true` to specify that Core Image should create an interim file to
// avoid overwriting an existing file.
//
// # Return Value
// 
// Returns `true` if the object is successfully archived to the file.
//
// # Discussion
// 
// Use this method to save your filter chain to a file for later use.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/write(to:atomically:)
func (f CIFilterGenerator) WriteToURLAtomically(aURL foundation.INSURL, flag bool) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("writeToURL:atomically:"), aURL, flag)
	return rv
}
// Registers the name associated with a filter chain.
//
// name: A unique name for the filter chain you want to register.
//
// # Discussion
// 
// This method allows you to register the filter chain as a named filter in
// the Core Image filter repository. You can then create a [CIFilter] object
// from it using the [FilterWithName] method of the [CIFilter] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/registerFilterName(_:)
func (f CIFilterGenerator) RegisterFilterName(name string) {
	objc.Send[objc.ID](f.ID, objc.Sel("registerFilterName:"), objc.String(name))
}
// Creates a filter object based on the filter chain.
//
// # Return Value
// 
// A [CIFilter] object.
//
// # Discussion
// 
// The topology of the filter chain is immutable, meaning that any changes you
// make to the filter chain are not reflected in the filter. The returned
// filter holds the export input and output keys.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filter()
func (f CIFilterGenerator) Filter() CIFilter {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filter"))
	return CIFilterFromID(rv)
}
// Returns a filter object specified by name.
//
// name: The name of the requested custom filter.
//
// # Return Value
// 
// A [CIFilter] object implementing the custom filter.
//
// # Discussion
// 
// Core Image calls this method when a filter is requested by name using the
// [CIFilter] class method [FilterWithName] method (or related methods). Your
// implementation of this method should provide a new instance of the
// [CIFilter] subclass for your custom filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterConstructor/filter(withName:)
func (f CIFilterGenerator) FilterWithName(name string) CIFilter {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filterWithName:"), objc.String(name))
	return CIFilterFromID(rv)
}
func (f CIFilterGenerator) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an empty filter generator object.
//
// # Return Value
// 
// A [CIFilterGenerator] object.
//
// # Discussion
// 
// You use the returned object to connect two or more [CIFilter] objects and
// input images. It is also valid to have only one [CIFilter] object in a
// filter generator.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGenerator
func (_CIFilterGeneratorClass CIFilterGeneratorClass) FilterGenerator() CIFilterGenerator {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterGeneratorClass.class), objc.Sel("filterGenerator"))
	return CIFilterGeneratorFromID(rv)
}
// Creates and returns a filter generator object and initializes it with the
// contents of a filter generator file.
//
// aURL: The location of a filter generator file.
//
// # Return Value
// 
// A [CIFilterGenerator] object; returns `nil` if the file can’t be read.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGeneratorWithContentsOfURL:
func (_CIFilterGeneratorClass CIFilterGeneratorClass) FilterGeneratorWithContentsOfURL(aURL foundation.INSURL) CIFilterGenerator {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterGeneratorClass.class), objc.Sel("filterGeneratorWithContentsOfURL:"), aURL)
	return CIFilterGeneratorFromID(rv)
}

// Returns an array of the exported keys.
//
// # Return Value
// 
// An array of dictionaries that describe the exported key and target object.
// See [kCIFilterGeneratorExportedKey],
// [kCIFilterGeneratorExportedKeyTargetObject], and
// [kCIFilterGeneratorExportedKey] for keys used in the dictionary.
// 
// # Discussion
// 
// This method returns the keys that you exported using the
// [ExportKeyFromObjectWithName] method or that were exported before being
// written to the file from which you read the filter chain.
//
// [kCIFilterGeneratorExportedKeyTargetObject]: https://developer.apple.com/documentation/CoreImage/kCIFilterGeneratorExportedKeyTargetObject
// [kCIFilterGeneratorExportedKey]: https://developer.apple.com/documentation/CoreImage/kCIFilterGeneratorExportedKey
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/exportedKeys
func (f CIFilterGenerator) ExportedKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("exportedKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The class attributes associated with the filter.
//
// # Discussion
// 
// For more information about class attributes for a filter, see [Core Image
// Programming Guide] and the filter attributes key constants defined in
// [CIFilter].
//
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/classAttributes
func (f CIFilterGenerator) ClassAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("classAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (f CIFilterGenerator) SetClassAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](f.ID, objc.Sel("setClassAttributes:"), value)
}

			// Protocol methods for CIFilterConstructor
			

