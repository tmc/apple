// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSNib] class.
var (
	_NSNibClass     NSNibClass
	_NSNibClassOnce sync.Once
)

func getNSNibClass() NSNibClass {
	_NSNibClassOnce.Do(func() {
		_NSNibClass = NSNibClass{class: objc.GetClass("NSNib")}
	})
	return _NSNibClass
}

// GetNSNibClass returns the class object for NSNib.
func GetNSNibClass() NSNibClass {
	return getNSNibClass()
}

type NSNibClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSNibClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNibClass) Alloc() NSNib {
	rv := objc.Send[NSNib](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object wrapper, or container, for an Interface Builder nib file.
//
// # Overview
//
// An [NSNib] object keeps the contents of a nib file resident in memory,
// ready for unarchiving and instantiation. When you create a nib object using
// the contents of a nib file, the object loads the contents of the referenced
// nib bundle—the object graph as well as any images and sounds—into
// memory; but it does not yet unarchive it. To unarchive all of the nib data
// and thus truly instantiate the nib you must call one of the
// `instantiate...` methods of [NSNib].
//
// During the instantiation process, each object in the archive is unarchived
// and then initialized using the method befitting its type. View classes are
// initialized using their [InitWithFrame] method. Custom objects are
// initialized using their `init` method. In the case of Cocoa views (and
// custom views that have options on an associated Interface Builder palette)
// the initialization process also reads in any values set by the user in
// Interface Builder.
//
// Once all objects have been instantiated and initialized from the archive,
// the nib loading code attempts to reestablish the connections between each
// object’s outlets and the corresponding target objects. If your custom
// objects have outlets, the [NSNib] object attempts to reestablish any
// connections you created in Interface Builder. It starts by trying to
// establish the connections using your object’s own methods first. For each
// outlet that needs a connection, the [NSNib] object looks for a method of
// the form “ in your object. If that method exists, the nib object calls it,
// passing the target object as a parameter. If you did not define a setter
// method with that exact name, the [NSNib] object searches the object for an
// instance variable (of type `IBOutlet id`) with the corresponding outlet
// name and tries to set its value directly. If an instance variable with the
// correct name cannot be found, initialization of that connection does not
// occur.
//
// After all objects have been initialized and their connections
// reestablished, each object receives an [awakeFromNib()] message. You can
// override this method in your custom objects to perform any additional
// initialization.
//
// # Subclassing Notes
//
// You can subclass [NSNib] if you want to extend or specialize nib-loading
// behavior. For example, you could create a custom [NSNib] subclass that
// performs some post-processing on the top-level objects returned from the
// `instantiateNib...` methods. If you want to modify how nib instantiations
// are performed, it is recommended that you override the primitive method
// [NSNib.InstantiateWithOwnerTopLevelObjects]. Note that the instance variables of
// [NSNib] are private and thus are not available to subclasses. Any override
// of [NSNib.InitWithNibDataBundle] or [NSNib.InitWithNibNamedBundle] should first invoke
// the superclass implementation.
//
// # Initializing a Nib
//
//   - [NSNib.InitWithNibNamedBundle]: Returns an [NSNib] object initialized to the nib file in the specified bundle.
//   - [NSNib.InitWithNibDataBundle]: Initializes an instance with nib data and specified bundle for locating resources.
//
// # Instantiating a Nib
//
//   - [NSNib.InstantiateWithOwnerTopLevelObjects]: Instantiates objects in the nib file with the specified owner.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib
//
// [awakeFromNib()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeFromNib()
type NSNib struct {
	objectivec.Object
}

// NSNibFromID constructs a [NSNib] from an objc.ID.
//
// An object wrapper, or container, for an Interface Builder nib file.
func NSNibFromID(id objc.ID) NSNib {
	return NSNib{objectivec.Object{ID: id}}
}

// NOTE: NSNib adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSNib] class.
//
// # Initializing a Nib
//
//   - [INSNib.InitWithNibNamedBundle]: Returns an [NSNib] object initialized to the nib file in the specified bundle.
//   - [INSNib.InitWithNibDataBundle]: Initializes an instance with nib data and specified bundle for locating resources.
//
// # Instantiating a Nib
//
//   - [INSNib.InstantiateWithOwnerTopLevelObjects]: Instantiates objects in the nib file with the specified owner.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib
type INSNib interface {
	objectivec.IObject

	// Topic: Initializing a Nib

	// Returns an [NSNib] object initialized to the nib file in the specified bundle.
	InitWithNibNamedBundle(nibName NSNibName, bundle foundation.NSBundle) NSNib
	// Initializes an instance with nib data and specified bundle for locating resources.
	InitWithNibDataBundle(nibData foundation.INSData, bundle foundation.NSBundle) NSNib

	// Topic: Instantiating a Nib

	// Instantiates objects in the nib file with the specified owner.
	InstantiateWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects foundation.INSArray) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n NSNib) Init() NSNib {
	rv := objc.Send[NSNib](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNib) Autorelease() NSNib {
	rv := objc.Send[NSNib](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNib creates a new NSNib instance.
func NewNSNib() NSNib {
	class := getNSNibClass()
	rv := objc.Send[NSNib](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an instance with nib data and specified bundle for locating
// resources.
//
// nibData: The nib data.
//
// bundle: The bundle for locating resources. If `nil`, the main application bundle is
// used.
//
// # Return Value
//
// The initialized [NSNib] object or `nil` if there were errors during
// initialization.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib/init(nibData:bundle:)
func NewNibWithNibDataBundle(nibData foundation.INSData, bundle foundation.NSBundle) NSNib {
	instance := getNSNibClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibData:bundle:"), nibData, bundle)
	return NSNibFromID(rv)
}

// Returns an [NSNib] object initialized to the nib file in the specified
// bundle.
//
// nibName: The name of the nib file, without any leading path information. Inclusion
// of the `XCUIElementTypeNib` extension on the nib file name is optional.
//
// bundle: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
//
// The initialized [NSNib] object or `nil` if there were errors during
// initialization or the nib file could not be located.
//
// # Discussion
//
// The [NSNib] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the [Resources]
// directory.
//
// After the nib file has been loaded, the [NSNib] object uses the bundle’s
// resource map to locate additional resources referenced by the nib. If you
// specified `nil` for the `bundle` parameter, the [NSNib] object looks for
// those resources in the bundle associated with the class of the nib file’s
// owner instead. If the nib file does not have an owner, the [NSNib] object
// looks for additional resources in the application’s main bundle.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib/init(nibNamed:bundle:)
func NewNibWithNibNamedBundle(nibName NSNibName, bundle foundation.NSBundle) NSNib {
	instance := getNSNibClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibNamed:bundle:"), objc.String(string(nibName)), bundle)
	return NSNibFromID(rv)
}

// Returns an [NSNib] object initialized to the nib file in the specified
// bundle.
//
// nibName: The name of the nib file, without any leading path information. Inclusion
// of the `XCUIElementTypeNib` extension on the nib file name is optional.
//
// bundle: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
//
// The initialized [NSNib] object or `nil` if there were errors during
// initialization or the nib file could not be located.
//
// # Discussion
//
// The [NSNib] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the [Resources]
// directory.
//
// After the nib file has been loaded, the [NSNib] object uses the bundle’s
// resource map to locate additional resources referenced by the nib. If you
// specified `nil` for the `bundle` parameter, the [NSNib] object looks for
// those resources in the bundle associated with the class of the nib file’s
// owner instead. If the nib file does not have an owner, the [NSNib] object
// looks for additional resources in the application’s main bundle.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib/init(nibNamed:bundle:)
func (n NSNib) InitWithNibNamedBundle(nibName NSNibName, bundle foundation.NSBundle) NSNib {
	rv := objc.Send[NSNib](n.ID, objc.Sel("initWithNibNamed:bundle:"), objc.String(string(nibName)), bundle)
	return rv
}

// Initializes an instance with nib data and specified bundle for locating
// resources.
//
// nibData: The nib data.
//
// bundle: The bundle for locating resources. If `nil`, the main application bundle is
// used.
//
// # Return Value
//
// The initialized [NSNib] object or `nil` if there were errors during
// initialization.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib/init(nibData:bundle:)
func (n NSNib) InitWithNibDataBundle(nibData foundation.INSData, bundle foundation.NSBundle) NSNib {
	rv := objc.Send[NSNib](n.ID, objc.Sel("initWithNibData:bundle:"), nibData, bundle)
	return rv
}

// Instantiates objects in the nib file with the specified owner.
//
// owner: The object to set as the Nib’s owner (File’s Owner).
//
// topLevelObjects: On return, an array containing the top-level objects of the nib.
//
// # Return Value
//
// true if the nib is instantiated; otherwise false.
//
// # Discussion
//
// Unlike legacy methods, the objects adhere to standard Cocoa memory
// management rules; it is necessary to keep a strong reference to the objects
// or the array to prevent the nib contents from being deallocated.
//
// Outlets to top level objects should be strong references to demonstrate
// ownership and prevent deallocation.
//
// See: https://developer.apple.com/documentation/AppKit/NSNib/instantiate(withOwner:topLevelObjects:)
func (n NSNib) InstantiateWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects foundation.INSArray) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("instantiateWithOwner:topLevelObjects:"), owner, topLevelObjects)
	return rv
}
func (n NSNib) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
