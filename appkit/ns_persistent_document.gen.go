// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentDocument] class.
var (
	_NSPersistentDocumentClass     NSPersistentDocumentClass
	_NSPersistentDocumentClassOnce sync.Once
)

func getNSPersistentDocumentClass() NSPersistentDocumentClass {
	_NSPersistentDocumentClassOnce.Do(func() {
		_NSPersistentDocumentClass = NSPersistentDocumentClass{class: objc.GetClass("NSPersistentDocument")}
	})
	return _NSPersistentDocumentClass
}

// GetNSPersistentDocumentClass returns the class object for NSPersistentDocument.
func GetNSPersistentDocumentClass() NSPersistentDocumentClass {
	return getNSPersistentDocumentClass()
}

type NSPersistentDocumentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentDocumentClass) Alloc() NSPersistentDocument {
	rv := objc.Send[NSPersistentDocument](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A document object that can integrate with Core Data.
//
// # Overview
// 
// The [NSPersistentDocument] class is a subclass of [NSDocument] that is
// designed to easily integrate into the Core Data framework. It provides
// methods to access a document-wide [NSManagedObjectContext] object, and
// provides default implementations of methods to read and write files using
// the persistence framework. In a persistent document, the undo manager
// functionality is taken over by managed object context.
// 
// Standard document behavior is implemented as follows:
// 
// - Opening a document invokes
// [NSPersistentDocument.ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError]
// with the new URL, and adds a store of the default type (XML). Objects are
// loaded from the persistent store on demand through the document’s
// context. - Saving a new document adds a store of the default type with the
// chosen URL and invokes save: on the context. For an existing document, a
// save just invokes [save()] on the context. - Save As for a new document
// simply invokes save. For an opened document, it migrates the persistent
// store to the new URL and invokes [save()] on the context. - Revert resets
// the document’s managed object context. Objects are subsequently loaded
// from the persistent store on demand, as with opening a new document.
// 
// By default an [NSPersistentDocument] instance creates its own ready-to-use
// persistence stack including managed object context, persistent object store
// coordinator and persistent store. There is a one-to-one mapping between the
// document and the backing object store.
// 
// You can customize the architecture of the persistence stack by overriding
// the [NSPersistentDocument.ManagedObjectModel] property and
// [NSPersistentDocument.ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError]
// method. You might wish to do this, for example, to specify a particular
// managed object model.
// 
// # Undo Support
// 
// The persistent document uses the managed object context’s undo manager.
// 
// The [DocumentEdited] method returns [true] if the persistent document’s
// managed object context, or editors registered with the context, have
// uncommitted changes, otherwise it returns [false].
//
// [NSManagedObjectContext]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext
// [false]: https://developer.apple.com/documentation/Swift/false
// [save()]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Managing the Persistence Objects
//
//   - [NSPersistentDocument.ManagedObjectContext]: The managed object context for the document.
//   - [NSPersistentDocument.SetManagedObjectContext]
//   - [NSPersistentDocument.ManagedObjectModel]: The managed object model of the document.
//   - [NSPersistentDocument.ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError]: Configures the receiver’s persistent store coordinator with the appropriate stores for a given URL.
//   - [NSPersistentDocument.PersistentStoreTypeForFileType]: Returns the type of persistent store associated with the specified file type.
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument
type NSPersistentDocument struct {
	NSDocument
}

// NSPersistentDocumentFromID constructs a [NSPersistentDocument] from an objc.ID.
//
// A document object that can integrate with Core Data.
func NSPersistentDocumentFromID(id objc.ID) NSPersistentDocument {
	return NSPersistentDocument{NSDocument: NSDocumentFromID(id)}
}
// NOTE: NSPersistentDocument adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentDocument] class.
//
// # Managing the Persistence Objects
//
//   - [INSPersistentDocument.ManagedObjectContext]: The managed object context for the document.
//   - [INSPersistentDocument.SetManagedObjectContext]
//   - [INSPersistentDocument.ManagedObjectModel]: The managed object model of the document.
//   - [INSPersistentDocument.ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError]: Configures the receiver’s persistent store coordinator with the appropriate stores for a given URL.
//   - [INSPersistentDocument.PersistentStoreTypeForFileType]: Returns the type of persistent store associated with the specified file type.
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument
type INSPersistentDocument interface {
	INSDocument

	// Topic: Managing the Persistence Objects

	// The managed object context for the document.
	ManagedObjectContext() objectivec.IObject
	SetManagedObjectContext(value objectivec.IObject)
	// The managed object model of the document.
	ManagedObjectModel() objectivec.IObject
	// Configures the receiver’s persistent store coordinator with the appropriate stores for a given URL.
	ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError(url foundation.INSURL, fileType string, configuration string, storeOptions foundation.INSDictionary) (bool, error)
	// Returns the type of persistent store associated with the specified file type.
	PersistentStoreTypeForFileType(fileType string) string

	// A Boolean value that indicates whether the document has unsaved changes.
	IsDocumentEdited() bool
	SetIsDocumentEdited(value bool)
}

// Init initializes the instance.
func (p NSPersistentDocument) Init() NSPersistentDocument {
	rv := objc.Send[NSPersistentDocument](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentDocument) Autorelease() NSPersistentDocument {
	rv := objc.Send[NSPersistentDocument](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentDocument creates a new NSPersistentDocument instance.
func NewNSPersistentDocument() NSPersistentDocument {
	class := getNSPersistentDocumentClass()
	rv := objc.Send[NSPersistentDocument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a document with the specified contents, and places the
// resulting document’s file at the designated location.
//
// urlOrNil: The location for the document’s file. This value is `nil` for an
// autosaved document that the user never explicitly saved.
//
// contentsURL: The URL of the file that contains the document’s contents. When loading
// an autosaved document, this URL contains the location of the autosave file.
// The contents of this file replace the contents of the file in
// `absoluteDocumentURL`.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized document object, or `nil` if the document could not be
// created.
//
// # Discussion
// 
// The system calls this method to open a document that has an associated
// autosave file . You can override this method to handle any document
// initialization specific to autosave contents.
// 
// After reading the contents from the specified autosave file, this method
// updates the document’s change count using the [NSChangeReadOtherContents]
// change type.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(for:withContentsOf:ofType:)
func NewPersistentDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (NSPersistentDocument, error) {
	var errorPtr objc.ID
	instance := getNSPersistentDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentDocumentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentDocumentFromID(rv), nil
}

// Initializes a document located by a URL of a specified type.
//
// url: The URL from which the contents of the document are obtained.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// You can override this method to customize the reopening of autosaved
// documents.
// 
// This method is invoked by the [NSDocumentController] method
// [DocumentWithContentsOfURLOfTypeError]. The default implementation of this
// method calls the [Init] and [ReadFromURLOfTypeError] methods and sets
// values for the [FileURL], [FileType], and [FileModificationDate]
// properties.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes
// [initWithContentsOfFile:ofType:] if it is overridden and the URL uses the
// `` scheme. It still updates the [FileModificationDate] property in this
// situation.
//
// [initWithContentsOfFile:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/initWithContentsOfFile:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(contentsOf:ofType:)
func NewPersistentDocumentWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (NSPersistentDocument, error) {
	var errorPtr objc.ID
	instance := getNSPersistentDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentDocumentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentDocumentFromID(rv), nil
}

// Initializes a document of a specified type.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// The default implementation of this method just invokes `[self init]` and
// `[self typeName]`.
// 
// You can override this method to perform initialization that must be done
// when creating new documents but should not be done when opening existing
// documents. Your override should typically invoke `super`, or at least it
// must invoke [Init], the [NSDocument] designated initializer, to initialize
// the [NSDocument] private instance variables.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(type:)
func NewPersistentDocumentWithTypeError(typeName string) (NSPersistentDocument, error) {
	var errorPtr objc.ID
	instance := getNSPersistentDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentDocumentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentDocumentFromID(rv), nil
}

// Configures the receiver’s persistent store coordinator with the
// appropriate stores for a given URL.
//
// url: An URL that specifies the location of the document’s store.
//
// fileType: The document type.
//
// configuration: The name of the managed object model configuration to use. (The managed
// object model is associated with the persistent store coordinator.) Pass
// `nil` if you do not want to specify a configuration.
//
// storeOptions: Options for the store. See “Store Options” in
// [NSPersistentStoreCoordinator] for possible values.
// //
// [NSPersistentStoreCoordinator]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator
//
// # Discussion
// 
// This method is invoked automatically when an existing document is opened.
// You override this method to customize creation of a persistent store for a
// given document or store type. You can retrieve the persistent store
// coordinator with the following code:
// 
// You can override this method to create the store to save to or load from
// (invoked from within the other [NSDocument] methods to read/write files),
// which gives developers the ability to load/save from/to different
// persistent store types (default type is XML).
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument/configurePersistentStoreCoordinator(for:ofType:modelConfiguration:storeOptions:)
func (p NSPersistentDocument) ConfigurePersistentStoreCoordinatorForURLOfTypeModelConfigurationStoreOptionsError(url foundation.INSURL, fileType string, configuration string, storeOptions foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("configurePersistentStoreCoordinatorForURL:ofType:modelConfiguration:storeOptions:error:"), url, objc.String(fileType), objc.String(configuration), storeOptions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("configurePersistentStoreCoordinatorForURL:ofType:modelConfiguration:storeOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the type of persistent store associated with the specified file
// type.
//
// fileType: A document file type.
//
// # Return Value
// 
// The type of persistent store associated with `fileType`. For possible
// values, see [NSPersistentStoreCoordinator].
//
// [NSPersistentStoreCoordinator]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator
//
// # Discussion
// 
// You set the persistent store type in the application’s property list.
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument/persistentStoreType(forFileType:)
func (p NSPersistentDocument) PersistentStoreTypeForFileType(fileType string) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("persistentStoreTypeForFileType:"), objc.String(fileType))
	return foundation.NSStringFromID(rv).String()
}

// The managed object context for the document.
//
// # Discussion
// 
// If a managed object context for the document does not exist, one is created
// automatically. If you want to customize the creation of the persistence
// stack, reimplement this property in your custom subclass and use your
// implementation to create the appropriate objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument/managedObjectContext
func (p NSPersistentDocument) ManagedObjectContext() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectContext"))
	return objectivec.Object{ID: rv}
}
func (p NSPersistentDocument) SetManagedObjectContext(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setManagedObjectContext:"), value)
}

// The managed object model of the document.
//
// # Discussion
// 
// By default the Core Data framework creates a merged model from all models
// in the application bundle (`[NSBundle mainBundle]`). You can reimplement
// this property and return a specific model to use to create persistent
// stores. A typical implementation might include code similar to the
// following fragment:
// 
// # Special Considerations
// 
// In applications built in OS X v10.4, by default the Core Data framework
// creates a merged model from all the models found in the application bundle
// .
//
// See: https://developer.apple.com/documentation/AppKit/NSPersistentDocument/managedObjectModel
func (p NSPersistentDocument) ManagedObjectModel() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectModel"))
	return objectivec.Object{ID: rv}
}

// A Boolean value that indicates whether the document has unsaved changes.
//
// See: https://developer.apple.com/documentation/appkit/nsdocument/isdocumentedited
func (p NSPersistentDocument) IsDocumentEdited() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("documentEdited"))
	return rv
}
func (p NSPersistentDocument) SetIsDocumentEdited(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDocumentEdited:"), value)
}

