// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSItemProvider] class.
var (
	_NSItemProviderClass     NSItemProviderClass
	_NSItemProviderClassOnce sync.Once
)

func getNSItemProviderClass() NSItemProviderClass {
	_NSItemProviderClassOnce.Do(func() {
		_NSItemProviderClass = NSItemProviderClass{class: objc.GetClass("NSItemProvider")}
	})
	return _NSItemProviderClass
}

// GetNSItemProviderClass returns the class object for NSItemProvider.
func GetNSItemProviderClass() NSItemProviderClass {
	return getNSItemProviderClass()
}

type NSItemProviderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSItemProviderClass) Alloc() NSItemProvider {
	rv := objc.Send[NSItemProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An item provider for conveying data or a file between processes during
// drag-and-drop or copy-and-paste activities, or from a host app to an app
// extension.
//
// # Overview
// 
// Starting in iOS 11, item providers play a central role in drag and drop,
// and in copy and paste. They continue to play a role with app extensions.
// 
// The system uses an internal queue when calling the completion blocks for
// the [NSItemProvider] class. When using an item provider with drag and drop,
// ensure that UI updates take place on the main queue as follows:
// 
// # App extension support
// 
// An app extension typically encounters item providers when examining the
// [NSItemProvider.Attachments] property of an [NSExtensionItem] object. During that
// examination, the extension can use the [NSItemProvider.HasItemConformingToTypeIdentifier]
// method to look for data that it recognizes. Item providers use [Uniform
// Type Identifiers] values to identify the data they contain. After finding a
// type of data that your extension can use, it calls the
// [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler] method to load the
// actual data, which is delivered to the provided completion handler.
// 
// You can create item providers to vend data to another process. An extension
// that modifies an original data item can create a new [NSItemProvider]
// object to send back to the host app. When creating data items, you specify
// your data object and the type of that object. You can optionally use the
// [NSItemProvider.PreviewImageHandler] property to generate a preview image for your data.
// 
// A single item provider may use custom blocks to provide its data in many
// different formats. When configuring an item provider, use the
// [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler] method to register your blocks
// and the formats each one supports. When a client requests data in a
// particular format, the item provider executes the corresponding block,
// which is then responsible for coercing the data to the appropriate type and
// returning it to the client.
//
// [Uniform Type Identifiers]: https://developer.apple.com/documentation/UniformTypeIdentifiers
//
// # Creating an item provider
//
//   - [NSItemProvider.InitWithContentsOfURL]: Provides data-backed content from an existing file.
//   - [NSItemProvider.InitWithItemTypeIdentifier]: Creates an item provider with an object, according to the item provider type coercion policy.
//   - [NSItemProvider.InitWithObject]: Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load.
//
// # Configuring the provider
//
//   - [NSItemProvider.PreferredPresentationSize]: The ideal presentation size of the item.
//   - [NSItemProvider.SetPreferredPresentationSize]
//   - [NSItemProvider.SuggestedName]: The filename to use when writing the provided data to a file on disk.
//   - [NSItemProvider.SetSuggestedName]
//
// # Querying the provider’s contents
//
//   - [NSItemProvider.CanLoadObjectOfClass]: Returns a Boolean value indicating whether an item provider can load objects of a specified class.
//   - [NSItemProvider.HasItemConformingToTypeIdentifier]: Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier file options parameter with a value of zero.
//   - [NSItemProvider.HasRepresentationConformingToTypeIdentifierFileOptions]: Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier and to specified open-in-place behavior.
//   - [NSItemProvider.RegisteredTypeIdentifiers]: Returns the array of type identifiers for the item provider, in the same order they were registered.
//   - [NSItemProvider.RegisteredTypeIdentifiersWithFileOptions]: Returns an array with a subset of type identifiers for the item provider, according to the specified file options, in the same order they were registered.
//
// # Loading the provider’s contents
//
//   - [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]: Loads the item’s data and coerces it to the specified type.
//   - [NSItemProvider.LoadDataRepresentationForTypeIdentifierCompletionHandler]: Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//   - [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]: Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
//   - [NSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]: Asynchronously opens a file in place, if possible, returning a progress object.
//   - [NSItemProvider.LoadObjectOfClassCompletionHandler]: Asynchronously loads an object of a specified class to an item provider, returning a progress object.
//
// # Loading a preview image
//
//   - [NSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]: Loads the preview image for the item that the item provider represents.
//   - [NSItemProvider.PreviewImageHandler]: The custom preview image handler block for the item provider.
//   - [NSItemProvider.SetPreviewImageHandler]
//
// # Registering CloudKit shares
//
//   - [NSItemProvider.RegisterCloudKitShareContainer]: Registers a CloudKit share for the user to modify.
//   - [NSItemProvider.RegisterCloudKitShareWithPreparationHandler]: Registers a handler that prepares a new CloudKit share.
//
// # Registering content types
//
//   - [NSItemProvider.RegisteredContentTypes]: Registered content types in the order the app registers each type.
//   - [NSItemProvider.RegisteredContentTypesForOpenInPlace]: Registered content types that the system can load as open-in-place files.
//   - [NSItemProvider.RegisteredContentTypesConformingToContentType]: Returns an array of registered content types that conform to a specified content type.
//
// # Registering data
//
//   - [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]: Registers a data-backed representation for an item, specifiying item visibility and a load handler.
//   - [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler]: Lazily registers an item, according to the item provider type coercion policy.
//
// # Registering files
//
//   - [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]: Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler.
//
// # Registering objects
//
//   - [NSItemProvider.RegisterObjectVisibility]: Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//   - [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]: Lazily adds representations of a specified object class to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//
// # Getting the provider’s frame
//
//   - [NSItemProvider.SourceFrame]: The rectangle that the item occupies in the host app’s source window.
//   - [NSItemProvider.ContainerFrame]: The rectangle of the item’s visible content.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider
type NSItemProvider struct {
	objectivec.Object
}

// NSItemProviderFromID constructs a [NSItemProvider] from an objc.ID.
//
// An item provider for conveying data or a file between processes during
// drag-and-drop or copy-and-paste activities, or from a host app to an app
// extension.
func NSItemProviderFromID(id objc.ID) NSItemProvider {
	return NSItemProvider{objectivec.Object{id}}
}
// NOTE: NSItemProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSItemProvider] class.
//
// # Creating an item provider
//
//   - [INSItemProvider.InitWithContentsOfURL]: Provides data-backed content from an existing file.
//   - [INSItemProvider.InitWithItemTypeIdentifier]: Creates an item provider with an object, according to the item provider type coercion policy.
//   - [INSItemProvider.InitWithObject]: Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load.
//
// # Configuring the provider
//
//   - [INSItemProvider.PreferredPresentationSize]: The ideal presentation size of the item.
//   - [INSItemProvider.SetPreferredPresentationSize]
//   - [INSItemProvider.SuggestedName]: The filename to use when writing the provided data to a file on disk.
//   - [INSItemProvider.SetSuggestedName]
//
// # Querying the provider’s contents
//
//   - [INSItemProvider.CanLoadObjectOfClass]: Returns a Boolean value indicating whether an item provider can load objects of a specified class.
//   - [INSItemProvider.HasItemConformingToTypeIdentifier]: Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier file options parameter with a value of zero.
//   - [INSItemProvider.HasRepresentationConformingToTypeIdentifierFileOptions]: Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier and to specified open-in-place behavior.
//   - [INSItemProvider.RegisteredTypeIdentifiers]: Returns the array of type identifiers for the item provider, in the same order they were registered.
//   - [INSItemProvider.RegisteredTypeIdentifiersWithFileOptions]: Returns an array with a subset of type identifiers for the item provider, according to the specified file options, in the same order they were registered.
//
// # Loading the provider’s contents
//
//   - [INSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]: Loads the item’s data and coerces it to the specified type.
//   - [INSItemProvider.LoadDataRepresentationForTypeIdentifierCompletionHandler]: Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//   - [INSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]: Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
//   - [INSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]: Asynchronously opens a file in place, if possible, returning a progress object.
//   - [INSItemProvider.LoadObjectOfClassCompletionHandler]: Asynchronously loads an object of a specified class to an item provider, returning a progress object.
//
// # Loading a preview image
//
//   - [INSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]: Loads the preview image for the item that the item provider represents.
//   - [INSItemProvider.PreviewImageHandler]: The custom preview image handler block for the item provider.
//   - [INSItemProvider.SetPreviewImageHandler]
//
// # Registering CloudKit shares
//
//   - [INSItemProvider.RegisterCloudKitShareContainer]: Registers a CloudKit share for the user to modify.
//   - [INSItemProvider.RegisterCloudKitShareWithPreparationHandler]: Registers a handler that prepares a new CloudKit share.
//
// # Registering content types
//
//   - [INSItemProvider.RegisteredContentTypes]: Registered content types in the order the app registers each type.
//   - [INSItemProvider.RegisteredContentTypesForOpenInPlace]: Registered content types that the system can load as open-in-place files.
//   - [INSItemProvider.RegisteredContentTypesConformingToContentType]: Returns an array of registered content types that conform to a specified content type.
//
// # Registering data
//
//   - [INSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]: Registers a data-backed representation for an item, specifiying item visibility and a load handler.
//   - [INSItemProvider.RegisterItemForTypeIdentifierLoadHandler]: Lazily registers an item, according to the item provider type coercion policy.
//
// # Registering files
//
//   - [INSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]: Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler.
//
// # Registering objects
//
//   - [INSItemProvider.RegisterObjectVisibility]: Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//   - [INSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]: Lazily adds representations of a specified object class to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//
// # Getting the provider’s frame
//
//   - [INSItemProvider.SourceFrame]: The rectangle that the item occupies in the host app’s source window.
//   - [INSItemProvider.ContainerFrame]: The rectangle of the item’s visible content.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider
type INSItemProvider interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating an item provider

	// Provides data-backed content from an existing file.
	InitWithContentsOfURL(fileURL INSURL) NSItemProvider
	// Creates an item provider with an object, according to the item provider type coercion policy.
	InitWithItemTypeIdentifier(item NSSecureCoding, typeIdentifier string) NSItemProvider
	// Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load.
	InitWithObject(object NSItemProviderWriting) NSItemProvider

	// Topic: Configuring the provider

	// The ideal presentation size of the item.
	PreferredPresentationSize() corefoundation.CGSize
	SetPreferredPresentationSize(value corefoundation.CGSize)
	// The filename to use when writing the provided data to a file on disk.
	SuggestedName() string
	SetSuggestedName(value string)

	// Topic: Querying the provider’s contents

	// Returns a Boolean value indicating whether an item provider can load objects of a specified class.
	CanLoadObjectOfClass(aClass objc.Class) bool
	// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier file options parameter with a value of zero.
	HasItemConformingToTypeIdentifier(typeIdentifier string) bool
	// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier and to specified open-in-place behavior.
	HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions NSItemProviderFileOptions) bool
	// Returns the array of type identifiers for the item provider, in the same order they were registered.
	RegisteredTypeIdentifiers() []string
	// Returns an array with a subset of type identifiers for the item provider, according to the specified file options, in the same order they were registered.
	RegisteredTypeIdentifiersWithFileOptions(fileOptions NSItemProviderFileOptions) []string

	// Topic: Loading the provider’s contents

	// Loads the item’s data and coerces it to the specified type.
	LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options INSDictionary, completionHandler ErrorHandler)
	// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
	LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress
	// Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
	LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler URLErrorHandler) INSProgress
	// Asynchronously opens a file in place, if possible, returning a progress object.
	LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler URLErrorHandler) INSProgress
	// Asynchronously loads an object of a specified class to an item provider, returning a progress object.
	LoadObjectOfClassCompletionHandler(aClass objc.Class, completionHandler NSItemProviderReadingErrorHandler) INSProgress

	// Topic: Loading a preview image

	// Loads the preview image for the item that the item provider represents.
	LoadPreviewImageWithOptionsCompletionHandler(options INSDictionary, completionHandler ErrorHandler)
	// The custom preview image handler block for the item provider.
	PreviewImageHandler() NSItemProviderLoadHandler
	SetPreviewImageHandler(value NSItemProviderLoadHandler)

	// Topic: Registering CloudKit shares

	// Registers a CloudKit share for the user to modify.
	RegisterCloudKitShareContainer(share objectivec.IObject, container objectivec.IObject)
	// Registers a handler that prepares a new CloudKit share.
	RegisterCloudKitShareWithPreparationHandler(preparationHandler VoidHandler)

	// Topic: Registering content types

	// Registered content types in the order the app registers each type.
	RegisteredContentTypes() []objc.ID
	// Registered content types that the system can load as open-in-place files.
	RegisteredContentTypesForOpenInPlace() []objc.ID
	// Returns an array of registered content types that conform to a specified content type.
	RegisteredContentTypesConformingToContentType(contentType objectivec.IObject) []objc.ID

	// Topic: Registering data

	// Registers a data-backed representation for an item, specifiying item visibility and a load handler.
	RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler)
	// Lazily registers an item, according to the item provider type coercion policy.
	RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler ErrorHandler)

	// Topic: Registering files

	// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler.
	RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions NSItemProviderFileOptions, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler)

	// Topic: Registering objects

	// Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
	RegisterObjectVisibility(object NSItemProviderWriting, visibility NSItemProviderRepresentationVisibility)
	// Lazily adds representations of a specified object class to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
	RegisterObjectOfClassVisibilityLoadHandler(aClass objc.Class, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler)

	// Topic: Getting the provider’s frame

	// The rectangle that the item occupies in the host app’s source window.
	SourceFrame() NSRect
	// The rectangle of the item’s visible content.
	ContainerFrame() NSRect

	// An optional array of media data associated with the extension item.
	Attachments() INSItemProvider
	SetAttachments(value INSItemProvider)
	// Provides data-backed content from an existing file with the specified parameters.
	InitWithContentsOfURLContentTypeOpenInPlaceCoordinatedVisibility(fileURL INSURL, contentType objectivec.IObject, openInPlace bool, coordinated bool, visibility NSItemProviderRepresentationVisibility) NSItemProvider
	// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
	LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler DataErrorHandler) INSProgress
	// Asynchronously copies the content type data into a generic data object with the specified parameters.
	LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler URLErrorHandler) INSProgress
	// Registers an existing collaboration object on a server.
	RegisterCKShareContainerAllowedSharingOptions(share objectivec.IObject, container objectivec.IObject, allowedOptions objectivec.IObject)
	// Creates and registers a new collaboration object using a collection of records to share.
	RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container objectivec.IObject, allowedOptions objectivec.IObject, preparationHandler ErrorHandler)
	// Lazily registers an item, according to the item provider type coercion policy.
	RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler)
	// Registers a file-backed representation for an item with item visibility, an open-in-place option, and a load handler.
	RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, openInPlace bool, loadHandler VoidHandler)
}





// Init initializes the instance.
func (i NSItemProvider) Init() NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSItemProvider) Autorelease() NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSItemProvider creates a new NSItemProvider instance.
func NewNSItemProvider() NSItemProvider {
	class := getNSItemProviderClass()
	rv := objc.Send[NSItemProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Provides data-backed content from an existing file.
//
// fileURL: The URL of the file to use for the item provider’s data. The item
// provider uses the filename extension to determine the universal type
// identifier for the associated data.
//
// # Return Value
// 
// An item provider for the specified file, or `nil` if an error occurs.
//
// # Discussion
// 
// The system uses the URL’s filename extension to select an appropriate
// universal type identifier. If the system can’t determine a specific
// universal type identifier based on the filename extension, it assigns the
// `public.Data()` universal type identifier for the file.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(contentsOf:)
func NewItemProviderWithContentsOfURL(fileURL INSURL) NSItemProvider {
	instance := getNSItemProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), fileURL)
	return NSItemProviderFromID(rv)
}


// Provides data-backed content from an existing file with the specified
// parameters.
//
// fileURL: The URL of the file to use for the item provider’s data.
//
// contentType: The content type of the specified file.
//
// openInPlace: `true` if the system opens the file in place.
//
// coordinated: `true` if the returned file must be accessed using [NSFileCoordinator].
//
// visibility: The [NSItemProviderRepresentationVisibility] setting the system uses to
// identify which processes can see this content.
// //
// [NSItemProviderRepresentationVisibility]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
//
// # Return Value
// 
// An item provider for the specified file or `nil` if an error occurred.
//
// # Discussion
// 
// If [ItemProviderFileOptionOpenInPlace] is set to `false`, the system copies
// the file provided before the load handler returns.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:
func NewItemProviderWithContentsOfURLContentTypeOpenInPlaceCoordinatedVisibility(fileURL INSURL, contentType objectivec.IObject, openInPlace bool, coordinated bool, visibility NSItemProviderRepresentationVisibility) NSItemProvider {
	instance := getNSItemProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:"), fileURL, contentType, openInPlace, coordinated, visibility)
	return NSItemProviderFromID(rv)
}


// Creates an item provider with an object, according to the item provider
// type coercion policy.
//
// item: An object containing the data you want to provide. You may specify `nil`
// for this parameter and register items and types later.
//
// typeIdentifier: A string that represents the UTI of the item. If `item` is not `nil`, this
// parameter must not be `nil`.
//
// # Return Value
// 
// An item provider for the specified item.
//
// # Discussion
// 
// Use this method to initialize an item provider for objects in your app. The
// item provider registers your object with the specified type. Subsequent
// requests for that same type return the specified `item`.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(item:typeIdentifier:)
func NewItemProviderWithItemTypeIdentifier(item NSSecureCoding, typeIdentifier string) NSItemProvider {
	instance := getNSItemProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItem:typeIdentifier:"), item, objc.String(typeIdentifier))
	return NSItemProviderFromID(rv)
}


// Creates a new item provider, employing a specified object’s type
// identifiers to specify the data representations eligible for the provider
// to load.
//
// object: An object containing the data you want to provide.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(object:)
func NewItemProviderWithObject(object NSItemProviderWriting) NSItemProvider {
	instance := getNSItemProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return NSItemProviderFromID(rv)
}







// Provides data-backed content from an existing file.
//
// fileURL: The URL of the file to use for the item provider’s data. The item
// provider uses the filename extension to determine the universal type
// identifier for the associated data.
//
// # Return Value
// 
// An item provider for the specified file, or `nil` if an error occurs.
//
// # Discussion
// 
// The system uses the URL’s filename extension to select an appropriate
// universal type identifier. If the system can’t determine a specific
// universal type identifier based on the filename extension, it assigns the
// `public.Data()` universal type identifier for the file.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(contentsOf:)
func (i NSItemProvider) InitWithContentsOfURL(fileURL INSURL) NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("initWithContentsOfURL:"), fileURL)
	return rv
}

// Creates an item provider with an object, according to the item provider
// type coercion policy.
//
// item: An object containing the data you want to provide. You may specify `nil`
// for this parameter and register items and types later.
//
// typeIdentifier: A string that represents the UTI of the item. If `item` is not `nil`, this
// parameter must not be `nil`.
//
// # Return Value
// 
// An item provider for the specified item.
//
// # Discussion
// 
// Use this method to initialize an item provider for objects in your app. The
// item provider registers your object with the specified type. Subsequent
// requests for that same type return the specified `item`.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(item:typeIdentifier:)
func (i NSItemProvider) InitWithItemTypeIdentifier(item NSSecureCoding, typeIdentifier string) NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("initWithItem:typeIdentifier:"), item, objc.String(typeIdentifier))
	return rv
}

// Creates a new item provider, employing a specified object’s type
// identifiers to specify the data representations eligible for the provider
// to load.
//
// object: An object containing the data you want to provide.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(object:)
func (i NSItemProvider) InitWithObject(object NSItemProviderWriting) NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// Returns a Boolean value indicating whether an item provider can load
// objects of a specified class.
//
// aClass: The object class for comparison.
//
// # Return Value
// 
// `true` if the item provider can load objects of the class.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/canLoadObject(ofClass:)-3eig9
func (i NSItemProvider) CanLoadObjectOfClass(aClass objc.Class) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("canLoadObjectOfClass:"), aClass)
	return rv
}

// Returns a Boolean value indicating whether an item provider contains a data
// representation conforming to a specified universal type identifier file
// options parameter with a value of zero.
//
// typeIdentifier: A string that represents the desired UTI.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/hasItemConformingToTypeIdentifier(_:)
func (i NSItemProvider) HasItemConformingToTypeIdentifier(typeIdentifier string) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hasItemConformingToTypeIdentifier:"), objc.String(typeIdentifier))
	return rv
}

// Returns a Boolean value indicating whether an item provider contains a data
// representation conforming to a specified universal type identifier and to
// specified open-in-place behavior.
//
// # Discussion
// 
// To check all registered UTIs for type conformance, pass the value `0` in
// the `fileOptions` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/hasRepresentationConforming(toTypeIdentifier:fileOptions:)
func (i NSItemProvider) HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions NSItemProviderFileOptions) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hasRepresentationConformingToTypeIdentifier:fileOptions:"), objc.String(typeIdentifier), fileOptions)
	return rv
}

// Returns an array with a subset of type identifiers for the item provider,
// according to the specified file options, in the same order they were
// registered.
//
// fileOptions: An array of [NSItemProviderFileOptions].
// //
// [NSItemProviderFileOptions]: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
//
// # Return Value
// 
// An array of type identifier strings.
//
// # Discussion
// 
// To access the array of all registered UTIs, pass the value `0` in the
// `fileOptions` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredTypeIdentifiers(fileOptions:)
func (i NSItemProvider) RegisteredTypeIdentifiersWithFileOptions(fileOptions NSItemProviderFileOptions) []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("registeredTypeIdentifiersWithFileOptions:"), fileOptions)
	return objc.ConvertSliceToStrings(rv)
}

// Loads the item’s data and coerces it to the specified type.
//
// typeIdentifier: A string that represents the desired UTI.
//
// options: A dictionary of keys and values that provide information about the item,
// such as the size of an image. (See [NSItemProviderPreferredImageSizeKey]
// for a key you can use.)
// //
// [NSItemProviderPreferredImageSizeKey]: https://developer.apple.com/documentation/Foundation/NSItemProviderPreferredImageSizeKey
//
// completionHandler: A completion handler block to execute with the results. For information
// about the format of this block, see [NSItemProviderCompletionHandler].
//
// # Discussion
// 
// Call this method when you want to retrieve the item provider’s data. If
// the item provider object is able to provide data in the requested type, it
// does so and asynchronously executes your `completionHandler` block with the
// results. The block may be executed on a background thread.
// 
// The type information for the first parameter of your `completionHandler`
// block should be set to the class of the expected type. For example, when
// requesting text data, you might set the type of the first parameter to
// [NSString] or [NSAttributedString]. An item provider can perform simple
// type conversions of the data to the class you specify, such as from [NSURL]
// to [NSData] or [NSFileWrapper], or from [NSData] to [UIImage] (in iOS) or
// [NSImage] (in macOS). If the data could not be retrieved or coerced to the
// specified class, an error is passed to the completion block’s.
//
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadItem(forTypeIdentifier:options:completionHandler:)
func (i NSItemProvider) LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options INSDictionary, completionHandler ErrorHandler) {
		_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
		objc.Send[objc.ID](i.ID, objc.Sel("loadItemForTypeIdentifier:options:completionHandler:"), objc.String(typeIdentifier), options, _block2)
}

// Asynchronously copies the provided, typed data into a generic data object,
// returning a progress object.
//
// # Discussion
// 
// If the source app provides a folder URL, the [Data] object contains a zip
// archive with the folder as its top-level entry.
//
// [Data]: https://developer.apple.com/documentation/Foundation/Data
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentation(forTypeIdentifier:completionHandler:)
func (i NSItemProvider) LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
		_block1, _cleanup1 := NewDataErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadDataRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Asynchronously writes a copy of the provided, typed data to a temporary
// file, returning a progress object.
//
// # Discussion
// 
// This method writes a copy of the file’s data to a temporary file, which
// the system deletes when the completion handler returns.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentation(forTypeIdentifier:completionHandler:)
func (i NSItemProvider) LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler URLErrorHandler) INSProgress {
		_block1, _cleanup1 := NewURLErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Asynchronously opens a file in place, if possible, returning a progress
// object.
//
// # Discussion
// 
// The system sets the `isInPlace` parameter to [true] if the system
// successfully opened the file in place, or [false] if it made a local copy.
// In either case, you must access the returned [URL] using an
// [NSFileCoordinator] object.
// 
// If the system created a local copy of a file, it will be automatically
// deleted after your file coordinator relinquishes its read access to the
// file.
//
// [URL]: https://developer.apple.com/documentation/Foundation/URL
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadInPlaceFileRepresentation(forTypeIdentifier:completionHandler:)
func (i NSItemProvider) LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler URLErrorHandler) INSProgress {
		_block1, _cleanup1 := NewURLErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadInPlaceFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Asynchronously loads an object of a specified class to an item provider,
// returning a progress object.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadObject(ofClass:completionHandler:)-8ak5d
func (i NSItemProvider) LoadObjectOfClassCompletionHandler(aClass objc.Class, completionHandler NSItemProviderReadingErrorHandler) INSProgress {
		_block1, _cleanup1 := NewNSItemProviderReadingErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadObjectOfClass:completionHandler:"), aClass, _block1)
	return NSProgressFromID(rv)
}

// Loads the preview image for the item that the item provider represents.
//
// options: A dictionary of keys and values that provide information about the item,
// such as the size of an image. For a list of possible keys, see [Options
// Dictionary Key].
// //
// [Options Dictionary Key]: https://developer.apple.com/documentation/Foundation/options-dictionary-key
//
// completionHandler: A completion handler block to execute with the results. The first parameter
// of this block must be a parameter of type [NSData], [NSURL], [UIImage] (in
// iOS), or [NSImage] (in macOS) for receiving the image data. For more
// information about implementing the block, see
// [NSItemProviderCompletionHandler].
// //
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
//
// # Discussion
// 
// To handle image preview yourself, provide a completion handler block that
// returns an [NSData] or [NSURL] object, or an instance of a
// platform-specific image class ([UIImage] or [NSImage]).
// 
// This method supports implicit type coercion for the item parameter of the
// completion block.
//
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadPreviewImage(options:completionHandler:)
func (i NSItemProvider) LoadPreviewImageWithOptionsCompletionHandler(options INSDictionary, completionHandler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
		objc.Send[objc.ID](i.ID, objc.Sel("loadPreviewImageWithOptions:completionHandler:"), options, _block1)
}

// Registers a CloudKit share for the user to modify.
//
// share: The CloudKit share to modify.
//
// container: The CloudKit container that stores the shared records.
//
// share is a [cloudkit.CKShare].
//
// container is a [cloudkit.CKContainer].
//
// # Discussion
// 
// Use this method when the CloudKit share already exists on the server and
// you want to update it. The behavior of the sharing service depends on the
// role of the current user. An owner can edit the share’s configuration,
// which includes managing participants and their permissions. A participant
// can view the share’s configuration and choose to stop participating.
// 
// If you’re unsure which container to use, fetch the share’s metadata
// using [CKFetchShareMetadataOperation]. Then initialize an instance of
// [CKContainer] using the metadata’s [containerIdentifier] property.
// 
// Use the [NSCloudSharingServiceDelegate] protocol to respond to any changes
// the sharing service makes.
// 
// The following example shows how to create an item provider with an existing
// share. It then invokes the cloud-sharing service with the provider and
// presents the share’s configuration to the user.
//
// [CKContainer]: https://developer.apple.com/documentation/CloudKit/CKContainer
// [CKFetchShareMetadataOperation]: https://developer.apple.com/documentation/CloudKit/CKFetchShareMetadataOperation
// [NSCloudSharingServiceDelegate]: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate
// [containerIdentifier]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/containerIdentifier
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCloudKitShare(_:container:)
func (i NSItemProvider) RegisterCloudKitShareContainer(share objectivec.IObject, container objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("registerCloudKitShare:container:"), share, container)
}

// Registers a handler that prepares a new CloudKit share.
//
// preparationHandler: The handler the service invokes when it requires the CloudKit share.
//
// # Discussion
// 
// Use this method to share a hierarchy of CloudKit records with other iCloud
// users. When the service invokes the handler, create an instance of
// [CKShare] with a root record. Save the share to the server using
// [CKModifyRecordsOperation]. The root record (and its hierarchy) must
// already exist on the server or be part of the same save operation. After
// the share saves, call `preparationCompletionHandler` with the saved share
// and its container. If the save fails, pass the error to the completion
// handler instead. Invoking the sharing service with a share you register
// using this method prompts the user to begin sharing.
// 
// Use the [NSCloudSharingServiceDelegate] protocol to respond to any changes
// the sharing service makes.
// 
// The following example shows how to create an item provider with a handler
// that saves a share. It then invokes the cloud-sharing service with that
// provider.
//
// [CKModifyRecordsOperation]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
// [NSCloudSharingServiceDelegate]: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCloudKitShare(preparationHandler:)
func (i NSItemProvider) RegisterCloudKitShareWithPreparationHandler(preparationHandler VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(preparationHandler)
	defer _cleanup0()
		objc.Send[objc.ID](i.ID, objc.Sel("registerCloudKitShareWithPreparationHandler:"), _block0)
}

// Returns an array of registered content types that conform to a specified
// content type.
//
// contentType: The specified content type.
//
// # Return Value
// 
// An array of registered content types.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypes(conformingTo:)
func (i NSItemProvider) RegisteredContentTypesConformingToContentType(contentType objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("registeredContentTypesConformingToContentType:"), contentType)
	return rv
}

// Registers a data-backed representation for an item, specifiying item
// visibility and a load handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentation(forTypeIdentifier:visibility:loadHandler:)
func (i NSItemProvider) RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler) {
		_block2, _cleanup2 := NewVoidBlock(loadHandler)
	defer _cleanup2()
		objc.Send[objc.ID](i.ID, objc.Sel("registerDataRepresentationForTypeIdentifier:visibility:loadHandler:"), objc.String(typeIdentifier), visibility, _block2)
}

// Lazily registers an item, according to the item provider type coercion
// policy.
//
// typeIdentifier: A string that represents the desired UTI.
//
// loadHandler: A block capable of returning the data item as the specified type. For
// information about implementing this block, see [NSItemProviderLoadHandler].
//
// # Discussion
// 
// Use this method to register blocks that can take the item provider’s file
// or data object and convert it to a specific data format. Your `loadHandler`
// block is executed when a client passes the same `typeIdentifier` string to
// the [LoadItemForTypeIdentifierOptionsCompletionHandler] method. In the
// implementation of your block, coerce the data to the specified type and
// call the provided completion handler. You must call the completion handler,
// either with the requested data or with an error.
// 
// Item providers know how to coerce known types of objects, such as images or
// strings. Use this method to register blocks to coerce your custom data
// types.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerItem(forTypeIdentifier:loadHandler:)
func (i NSItemProvider) RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(loadHandler)
	defer _cleanup1()
		objc.Send[objc.ID](i.ID, objc.Sel("registerItemForTypeIdentifier:loadHandler:"), objc.String(typeIdentifier), _block1)
}

// Registers a file-backed representation for an item, specifying file
// options, item visibility, and a load handler.
//
// # Discussion
// 
// If a destination app must access the represented file using a file
// coordinator, set the `coordinated` parameter in the load handler block to
// [true].
// 
// To offer a representation backed by a file provider, return an [NSURL]
// object that points to your app’s file provider’s container. The file
// provider extension is then invoked to retrieve the file when requested.
// 
// To offer a representation backed by a file to open in place, set the
// fileOptions parameter to a value of [ItemProviderFileOptionOpenInPlace]; in
// addition, return an [NSURL] object that points to your app’s file
// provider’s container. Open-in-place support requires that the file
// provider is visible in the Files app.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentation(forTypeIdentifier:fileOptions:visibility:loadHandler:)
func (i NSItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions NSItemProviderFileOptions, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler) {
		_block3, _cleanup3 := NewVoidBlock(loadHandler)
	defer _cleanup3()
		objc.Send[objc.ID](i.ID, objc.Sel("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:"), objc.String(typeIdentifier), fileOptions, visibility, _block3)
}

// Adds representations of a specified object to an item provider, based on
// the object’s implementation of the item provider writing protocol, and
// adhering to a visibility specification.
//
// # Discussion
// 
// If a representation for a given UTI is already registered, it is preserved
// (specifically, duplicate representations are ignored).
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerObject(_:visibility:)
func (i NSItemProvider) RegisterObjectVisibility(object NSItemProviderWriting, visibility NSItemProviderRepresentationVisibility) {
	objc.Send[objc.ID](i.ID, objc.Sel("registerObject:visibility:"), object, visibility)
}

// Lazily adds representations of a specified object class to an item
// provider, based on the object’s implementation of the item provider
// writing protocol, and adhering to a visibility specification.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerObject(ofClass:visibility:loadHandler:)-9sndn
func (i NSItemProvider) RegisterObjectOfClassVisibilityLoadHandler(aClass objc.Class, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler) {
		_block2, _cleanup2 := NewVoidBlock(loadHandler)
	defer _cleanup2()
		objc.Send[objc.ID](i.ID, objc.Sel("registerObjectOfClass:visibility:loadHandler:"), aClass, visibility, _block2)
}

// Provides data-backed content from an existing file with the specified
// parameters.
//
// fileURL: The URL of the file to use for the item provider’s data.
//
// contentType: The content type of the specified file.
//
// openInPlace: `true` if the system opens the file in place.
//
// coordinated: `true` if the returned file must be accessed using [NSFileCoordinator].
//
// visibility: The [NSItemProviderRepresentationVisibility] setting the system uses to
// identify which processes can see this content.
// //
// [NSItemProviderRepresentationVisibility]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
//
// # Return Value
// 
// An item provider for the specified file or `nil` if an error occurred.
//
// # Discussion
// 
// If [ItemProviderFileOptionOpenInPlace] is set to `false`, the system copies
// the file provided before the load handler returns.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:
func (i NSItemProvider) InitWithContentsOfURLContentTypeOpenInPlaceCoordinatedVisibility(fileURL INSURL, contentType objectivec.IObject, openInPlace bool, coordinated bool, visibility NSItemProviderRepresentationVisibility) NSItemProvider {
	rv := objc.Send[NSItemProvider](i.ID, objc.Sel("initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:"), fileURL, contentType, openInPlace, coordinated, visibility)
	return rv
}

// Asynchronously copies the provided, typed data into a generic data object,
// returning a progress object.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentationForContentType:completionHandler:
func (i NSItemProvider) LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler DataErrorHandler) INSProgress {
		_block1, _cleanup1 := NewDataErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadDataRepresentationForContentType:completionHandler:"), contentType, _block1)
	return NSProgressFromID(rv)
}

// Asynchronously copies the content type data into a generic data object with
// the specified parameters.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentationForContentType:openInPlace:completionHandler:
func (i NSItemProvider) LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler URLErrorHandler) INSProgress {
		_block2, _cleanup2 := NewURLErrorBlock(completionHandler)
	defer _cleanup2()
		rv := objc.Send[objc.ID](i.ID, objc.Sel("loadFileRepresentationForContentType:openInPlace:completionHandler:"), contentType, openInPlace, _block2)
	return NSProgressFromID(rv)
}

// Registers an existing collaboration object on a server.
//
// share: An existing [CKShare] on the server.
// //
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
//
// container: A [CKContainer] the system uses to coordinate all the interactions between
// your app and the server.
// //
// [CKContainer]: https://developer.apple.com/documentation/CloudKit/CKContainer
//
// allowedOptions: The [CKAllowedSharingOptions]. The standard option is the default.
// //
// [CKAllowedSharingOptions]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
//
// share is a [cloudkit.CKShare].
//
// container is a [cloudkit.CKContainer].
//
// allowedOptions is a [cloudkit.CKAllowedSharingOptions].
//
// # Discussion
// 
// Use this method when a [CKShare] currently exists on the server. When the
// system invokes the share sheet with a [CKShare] that you register with this
// method, it allows the owner to make modifications to the share settings,
// and allows a participant to view the share settings.
//
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCKShare:container:allowedSharingOptions:
func (i NSItemProvider) RegisterCKShareContainerAllowedSharingOptions(share objectivec.IObject, container objectivec.IObject, allowedOptions objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("registerCKShare:container:allowedSharingOptions:"), share, container, allowedOptions)
}

// Creates and registers a new collaboration object using a collection of
// records to share.
//
// container: A [CKContainer] the system uses to coordinate all the interactions between
// your app and the server.
// //
// [CKContainer]: https://developer.apple.com/documentation/CloudKit/CKContainer
//
// allowedOptions: The [CKAllowedSharingOptions]. The standard option is the default.
// //
// [CKAllowedSharingOptions]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
//
// preparationHandler: The handler the system calls in your app to create a new [CKShare].
// //
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
//
// container is a [cloudkit.CKContainer].
//
// allowedOptions is a [cloudkit.CKAllowedSharingOptions].
//
// preparationHandler is a [cloudkit.CKSharePreparationHandler].
//
// # Discussion
// 
// Use this method to share a collection of [CKRecord] objects that aren’t
// assigned to an existing [CKShare]. When the system calls the
// `preparationHandler`, your app creates a new [CKShare] with the appropriate
// root [CKRecord] or [CKRecordZone.ID].
// 
// After the server successfully saves the share, invoke the
// [CKSharePreparationCompletionHandler] with either the resulting [CKShare]
// or an `NSError,` if the save failed.
// 
// When the system invokes the share sheet with a [CKShare] registered with
// this method, it prompts the user to start sharing.
//
// [CKRecordZone.ID]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID
// [CKRecord]: https://developer.apple.com/documentation/CloudKit/CKRecord
// [CKSharePreparationCompletionHandler]: https://developer.apple.com/documentation/CloudKit/CKSharePreparationCompletionHandler
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCKShareWithContainer:allowedSharingOptions:preparationHandler:
func (i NSItemProvider) RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container objectivec.IObject, allowedOptions objectivec.IObject, preparationHandler ErrorHandler) {
		_block2, _cleanup2 := NewErrorBlock(preparationHandler)
	defer _cleanup2()
		objc.Send[objc.ID](i.ID, objc.Sel("registerCKShareWithContainer:allowedSharingOptions:preparationHandler:"), container, allowedOptions, _block2)
}

// Lazily registers an item, according to the item provider type coercion
// policy.
//
// contentType: A string that represents the desired UTI.
//
// visibility: The [NSItemProviderRepresentationVisibility] setting.
// //
// [NSItemProviderRepresentationVisibility]: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
//
// loadHandler: A block capable of returning the data item as the specified type. For
// information about implementing this block, see [NSItemProviderLoadHandler].
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentationForContentType:visibility:loadHandler:
func (i NSItemProvider) RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, loadHandler VoidHandler) {
		_block2, _cleanup2 := NewVoidBlock(loadHandler)
	defer _cleanup2()
		objc.Send[objc.ID](i.ID, objc.Sel("registerDataRepresentationForContentType:visibility:loadHandler:"), contentType, visibility, _block2)
}

// Registers a file-backed representation for an item with item visibility, an
// open-in-place option, and a load handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentationForContentType:visibility:openInPlace:loadHandler:
func (i NSItemProvider) RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, openInPlace bool, loadHandler VoidHandler) {
		_block3, _cleanup3 := NewVoidBlock(loadHandler)
	defer _cleanup3()
		objc.Send[objc.ID](i.ID, objc.Sel("registerFileRepresentationForContentType:visibility:openInPlace:loadHandler:"), contentType, visibility, openInPlace, _block3)
}












// The ideal presentation size of the item.
//
// # Discussion
// 
// When displaying the item, the value in this property represents the ideal
// size at which to display the item. The size in this property may differ
// from the size in the [SourceFrame] rectangle. For images, video, and other
// content with a natural size, the item automatically derives the size from
// that content. If the value in this property is [NSZeroSize], use the size
// specified in the [SourceFrame] rectangle.
//
// [NSZeroSize]: https://developer.apple.com/documentation/Foundation/NSZeroSize
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationSize
func (i NSItemProvider) PreferredPresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](i.ID, objc.Sel("preferredPresentationSize"))
	return corefoundation.CGSize(rv)
}
func (i NSItemProvider) SetPreferredPresentationSize(value corefoundation.CGSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setPreferredPresentationSize:"), value)
}



// The filename to use when writing the provided data to a file on disk.
//
// # Discussion
// 
// Setting this property is recommended when providing [NSData] or text data
// from an item provider.
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/suggestedName
func (i NSItemProvider) SuggestedName() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("suggestedName"))
	return NSStringFromID(rv).String()
}
func (i NSItemProvider) SetSuggestedName(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setSuggestedName:"), objc.String(value))
}



// Returns the array of type identifiers for the item provider, in the same
// order they were registered.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredTypeIdentifiers
func (i NSItemProvider) RegisteredTypeIdentifiers() []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("registeredTypeIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}



// The custom preview image handler block for the item provider.
//
// # Discussion
// 
// In your image handler block, return an [NSURL] object that specifies a
// file, or return an [NSData] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/previewImageHandler
func (i NSItemProvider) PreviewImageHandler() NSItemProviderLoadHandler {
	rv := objc.Send[NSItemProviderLoadHandler](i.ID, objc.Sel("previewImageHandler"))
	return NSItemProviderLoadHandler(rv)
}
func (i NSItemProvider) SetPreviewImageHandler(value NSItemProviderLoadHandler) {
	objc.Send[struct{}](i.ID, objc.Sel("setPreviewImageHandler:"), value)
}



// Registered content types in the order the app registers each type.
//
// # Discussion
// 
// You app should register content types in order of fidelity. The system uses
// content types that appear earlier in the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypes
func (i NSItemProvider) RegisteredContentTypes() []objc.ID {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("registeredContentTypes"))
	return rv
}



// Registered content types that the system can load as open-in-place files.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypesForOpenInPlace
func (i NSItemProvider) RegisteredContentTypesForOpenInPlace() []objc.ID {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("registeredContentTypesForOpenInPlace"))
	return rv
}



// The rectangle that the item occupies in the host app’s source window.
//
// # Discussion
// 
// This property contains the rectangle, in screen coordinates, that encloses
// the item. This rectangle includes areas that might be clipped and not
// currently visible onscreen.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/sourceFrame
func (i NSItemProvider) SourceFrame() NSRect {
	rv := objc.Send[NSRect](i.ID, objc.Sel("sourceFrame"))
	return NSRect(rv)
}



// The rectangle of the item’s visible content.
//
// # Discussion
// 
// The rectangle in this property corresponds to the onscreen frame rectangle
// of the item. This rectangle may or may not intersect the [SourceFrame]
// rectangle of the item. An intersection of the rectangles means that at
// least part of the item is visible onscreen.
// 
// The rectangle in this property may be a clipped version of the source frame
// or it might be [NSZeroRect] if the item is offscreen or the system can’t
// determine the clipping rectangle. The system treats a value of [NSZeroRect]
// as meaning the item is fully visible.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/containerFrame
func (i NSItemProvider) ContainerFrame() NSRect {
	rv := objc.Send[NSRect](i.ID, objc.Sel("containerFrame"))
	return NSRect(rv)
}



// An optional array of media data associated with the extension item.
//
// See: https://developer.apple.com/documentation/foundation/nsextensionitem/attachments
func (i NSItemProvider) Attachments() INSItemProvider {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("attachments"))
	return NSItemProviderFromID(objc.ID(rv))
}
func (i NSItemProvider) SetAttachments(value INSItemProvider) {
	objc.Send[struct{}](i.ID, objc.Sel("setAttachments:"), value)
}













			// Protocol methods for NSCopying
			









// LoadDataRepresentationForTypeIdentifier is a synchronous wrapper around [NSItemProvider.LoadDataRepresentationForTypeIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) LoadDataRepresentationForTypeIdentifier(ctx context.Context, typeIdentifier string) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	i.LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadFileRepresentationForTypeIdentifier is a synchronous wrapper around [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) LoadFileRepresentationForTypeIdentifier(ctx context.Context, typeIdentifier string) (*NSURL, error) {
	type result struct {
		val *NSURL
		err error
	}
	done := make(chan result, 1)
	i.LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier, func(val *NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadObjectOfClass is a synchronous wrapper around [NSItemProvider.LoadObjectOfClassCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) LoadObjectOfClass(ctx context.Context, aClass objc.Class) (NSItemProviderReading, error) {
	type result struct {
		val NSItemProviderReading
		err error
	}
	done := make(chan result, 1)
	i.LoadObjectOfClassCompletionHandler(aClass, func(val NSItemProviderReading, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RegisterCloudKitShareWithPreparationHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterCloudKitShareWithPreparationHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterCloudKitShareWithPreparationHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	i.RegisterCloudKitShareWithPreparationHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandlerSync(ctx context.Context, typeIdentifier string, visibility NSItemProviderRepresentationVisibility) error {
	done := make(chan struct{}, 1)
	i.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier, visibility, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandlerSync(ctx context.Context, typeIdentifier string, fileOptions NSItemProviderFileOptions, visibility NSItemProviderRepresentationVisibility) error {
	done := make(chan struct{}, 1)
	i.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier, fileOptions, visibility, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterObjectOfClassVisibilityLoadHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterObjectOfClassVisibilityLoadHandlerSync(ctx context.Context, aClass objc.Class, visibility NSItemProviderRepresentationVisibility) error {
	done := make(chan struct{}, 1)
	i.RegisterObjectOfClassVisibilityLoadHandler(aClass, visibility, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadDataRepresentationForContentType is a synchronous wrapper around [NSItemProvider.LoadDataRepresentationForContentTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) LoadDataRepresentationForContentType(ctx context.Context, contentType objectivec.IObject) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	i.LoadDataRepresentationForContentTypeCompletionHandler(contentType, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RegisterDataRepresentationForContentTypeVisibilityLoadHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterDataRepresentationForContentTypeVisibilityLoadHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterDataRepresentationForContentTypeVisibilityLoadHandlerSync(ctx context.Context, contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility) error {
	done := make(chan struct{}, 1)
	i.RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType, visibility, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandlerSync is a synchronous wrapper around [NSItemProvider.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i NSItemProvider) RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandlerSync(ctx context.Context, contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, openInPlace bool) error {
	done := make(chan struct{}, 1)
	i.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler(contentType, visibility, openInPlace, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}





