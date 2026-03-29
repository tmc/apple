// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSHelpManager] class.
var (
	_NSHelpManagerClass     NSHelpManagerClass
	_NSHelpManagerClassOnce sync.Once
)

func getNSHelpManagerClass() NSHelpManagerClass {
	_NSHelpManagerClassOnce.Do(func() {
		_NSHelpManagerClass = NSHelpManagerClass{class: objc.GetClass("NSHelpManager")}
	})
	return _NSHelpManagerClass
}

// GetNSHelpManagerClass returns the class object for NSHelpManager.
func GetNSHelpManagerClass() NSHelpManagerClass {
	return getNSHelpManagerClass()
}

type NSHelpManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSHelpManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSHelpManagerClass) Alloc() NSHelpManager {
	rv := objc.Send[NSHelpManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object for displaying online help for an app.
//
// # Overview
// 
// The [NSHelpManager] class provides an approach to displaying online help.
// An app contains one [NSHelpManager] object.
//
// # Displaying Help
//
//   - [NSHelpManager.FindStringInBook]: Performs a search for the specified string in the specified book.
//   - [NSHelpManager.OpenHelpAnchorInBook]: Finds and displays the text at the given anchor location in the given book.
//
// # Dynamically Adding Help Books
//
//   - [NSHelpManager.RegisterBooksInBundle]: Registers one or more help books in the given bundle.
//
// # Configuring Context-Sensitive Help
//
//   - [NSHelpManager.SetContextHelpForObject]: Associates help content with an object.
//   - [NSHelpManager.RemoveContextHelpForObject]: Removes the association between an object and its context-sensitive help.
//
// # Displaying Context-Sensitive Help
//
//   - [NSHelpManager.ContextHelpForObject]: Returns context-sensitive help for an object.
//   - [NSHelpManager.ShowContextHelpForObjectLocationHint]: Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager
type NSHelpManager struct {
	objectivec.Object
}

// NSHelpManagerFromID constructs a [NSHelpManager] from an objc.ID.
//
// An object for displaying online help for an app.
func NSHelpManagerFromID(id objc.ID) NSHelpManager {
	return NSHelpManager{objectivec.Object{ID: id}}
}
// NOTE: NSHelpManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSHelpManager] class.
//
// # Displaying Help
//
//   - [INSHelpManager.FindStringInBook]: Performs a search for the specified string in the specified book.
//   - [INSHelpManager.OpenHelpAnchorInBook]: Finds and displays the text at the given anchor location in the given book.
//
// # Dynamically Adding Help Books
//
//   - [INSHelpManager.RegisterBooksInBundle]: Registers one or more help books in the given bundle.
//
// # Configuring Context-Sensitive Help
//
//   - [INSHelpManager.SetContextHelpForObject]: Associates help content with an object.
//   - [INSHelpManager.RemoveContextHelpForObject]: Removes the association between an object and its context-sensitive help.
//
// # Displaying Context-Sensitive Help
//
//   - [INSHelpManager.ContextHelpForObject]: Returns context-sensitive help for an object.
//   - [INSHelpManager.ShowContextHelpForObjectLocationHint]: Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager
type INSHelpManager interface {
	objectivec.IObject

	// Topic: Displaying Help

	// Performs a search for the specified string in the specified book.
	FindStringInBook(query string, book NSHelpBookName)
	// Finds and displays the text at the given anchor location in the given book.
	OpenHelpAnchorInBook(anchor NSHelpAnchorName, book NSHelpBookName)

	// Topic: Dynamically Adding Help Books

	// Registers one or more help books in the given bundle.
	RegisterBooksInBundle(bundle foundation.NSBundle) bool

	// Topic: Configuring Context-Sensitive Help

	// Associates help content with an object.
	SetContextHelpForObject(attrString foundation.NSAttributedString, object objectivec.IObject)
	// Removes the association between an object and its context-sensitive help.
	RemoveContextHelpForObject(object objectivec.IObject)

	// Topic: Displaying Context-Sensitive Help

	// Returns context-sensitive help for an object.
	ContextHelpForObject(object objectivec.IObject) foundation.NSAttributedString
	// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point.
	ShowContextHelpForObjectLocationHint(object objectivec.IObject, pt corefoundation.CGPoint) bool
}

// Init initializes the instance.
func (h NSHelpManager) Init() NSHelpManager {
	rv := objc.Send[NSHelpManager](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h NSHelpManager) Autorelease() NSHelpManager {
	rv := objc.Send[NSHelpManager](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSHelpManager creates a new NSHelpManager instance.
func NewNSHelpManager() NSHelpManager {
	class := getNSHelpManagerClass()
	rv := objc.Send[NSHelpManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Performs a search for the specified string in the specified book.
//
// query: String to search for.
//
// book: Localized help book to search. When `nil`, all installed help books are
// searched.
//
// # Discussion
// 
// To search for a string in your bundle’s localized help book, you could
// use code similar to the following:
// 
// This is a wrapper for [AHRegisterHelpBook] (which is called only once to
// register the help book specified in the application’s main bundle) and
// [AHSearch].
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h NSHelpManager) FindStringInBook(query string, book NSHelpBookName) {
	objc.Send[objc.ID](h.ID, objc.Sel("findString:inBook:"), objc.String(query), objc.String(string(book)))
}
// Finds and displays the text at the given anchor location in the given book.
//
// anchor: Location of the desired text.
//
// book: Help book containing the anchor. When `nil`, all installed help books are
// searched.
//
// # Discussion
// 
// To open an anchor in your bundle’s localized help book, you could use
// code similar to the following:
// 
// This method is a wrapper for [AHRegisterHelpBook] (which is called only
// once to register the help book specified in the application’s main
// bundle) and [AHLookupAnchor].
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h NSHelpManager) OpenHelpAnchorInBook(anchor NSHelpAnchorName, book NSHelpBookName) {
	objc.Send[objc.ID](h.ID, objc.Sel("openHelpAnchor:inBook:"), objc.String(string(anchor)), objc.String(string(book)))
}
// Registers one or more help books in the given bundle.
//
// bundle: The bundle for additional help books. Books in the main bundle are
// automatically registered.
//
// # Return Value
// 
// [true] if registration is successful, [false] if the bundle doesn’t
// contain any help books or if registration fails.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You use `` to register help books in, for example, a plug-in bundle. The
// `Info.Plist()` in the bundle should contain a help book directory path,
// which specifies one or more folders containing help books.
// 
// The main bundle is automatically registered by [OpenHelpAnchorInBook] and
// [FindStringInBook].
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h NSHelpManager) RegisterBooksInBundle(bundle foundation.NSBundle) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("registerBooksInBundle:"), bundle)
	return rv
}
// Associates help content with an object.
//
// attrString: Help content to associate with `object`.
//
// object: Object to associate with `help`.
//
// # Discussion
// 
// When the application enters context-sensitive help mode, if `object` is
// clicked, `help` appears in the context-sensitive help window.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h NSHelpManager) SetContextHelpForObject(attrString foundation.NSAttributedString, object objectivec.IObject) {
	objc.Send[objc.ID](h.ID, objc.Sel("setContextHelp:forObject:"), attrString, object)
}
// Removes the association between an object and its context-sensitive help.
//
// object: Object to disassociate from its help content.
//
// # Discussion
// 
// If `object` does not have context-sensitive help associated with it, this
// method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h NSHelpManager) RemoveContextHelpForObject(object objectivec.IObject) {
	objc.Send[objc.ID](h.ID, objc.Sel("removeContextHelpForObject:"), object)
}
// Returns context-sensitive help for an object.
//
// object: Object for which context-sensitive help is sought.
//
// # Return Value
// 
// Context-sensitive help content.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h NSHelpManager) ContextHelpForObject(object objectivec.IObject) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("contextHelpForObject:"), object)
	return foundation.NSAttributedStringFromID(rv)
}
// Displays the context-sensitive help for a given object at or near the point
// on the screen specified by a given point.
//
// object: Object for which context-sensitive help is sought.
//
// pt: Screen location at which to display the help content; it’s usually under
// the cursor.
//
// # Return Value
// 
// [true] when help content is successfully displayed. [false] if help content
// is not displayed (for example, when there is no context-sensitive help
// associated with `object`).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h NSHelpManager) ShowContextHelpForObjectLocationHint(object objectivec.IObject, pt corefoundation.CGPoint) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}

// Returns the shared [NSHelpManager] instance, creating it if it does not
// already exist.
//
// # Return Value
// 
// Shared help manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/shared
func (_NSHelpManagerClass NSHelpManagerClass) SharedHelpManager() NSHelpManager {
	rv := objc.Send[objc.ID](objc.ID(_NSHelpManagerClass.class), objc.Sel("sharedHelpManager"))
	return NSHelpManagerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (_NSHelpManagerClass NSHelpManagerClass) ContextHelpModeActive() bool {
	rv := objc.Send[bool](objc.ID(_NSHelpManagerClass.class), objc.Sel("isContextHelpModeActive"))
	return rv
}
func (_NSHelpManagerClass NSHelpManagerClass) SetContextHelpModeActive(value bool) {
	objc.Send[struct{}](objc.ID(_NSHelpManagerClass.class), objc.Sel("setContextHelpModeActive:"), value)
}
// Posted when the application enters context-sensitive help mode. This
// typically happens when the user holds down the Help key.
//
// See: https://developer.apple.com/documentation/appkit/nshelpmanager/contexthelpmodedidactivatenotification
func (_NSHelpManagerClass NSHelpManagerClass) ContextHelpModeDidActivateNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSHelpManagerClass.class), objc.Sel("NSContextHelpModeDidActivateNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// Posted when the application exits context-sensitive help mode. This happens
// when the user clicks the mouse button while the cursor is anywhere on the
// screen after displaying a context-sensitive help topic.
//
// See: https://developer.apple.com/documentation/appkit/nshelpmanager/contexthelpmodediddeactivatenotification
func (_NSHelpManagerClass NSHelpManagerClass) ContextHelpModeDidDeactivateNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSHelpManagerClass.class), objc.Sel("NSContextHelpModeDidDeactivateNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

