// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextAttachmentViewProvider] class.
var (
	_NSTextAttachmentViewProviderClass     NSTextAttachmentViewProviderClass
	_NSTextAttachmentViewProviderClassOnce sync.Once
)

func getNSTextAttachmentViewProviderClass() NSTextAttachmentViewProviderClass {
	_NSTextAttachmentViewProviderClassOnce.Do(func() {
		_NSTextAttachmentViewProviderClass = NSTextAttachmentViewProviderClass{class: objc.GetClass("NSTextAttachmentViewProvider")}
	})
	return _NSTextAttachmentViewProviderClass
}

// GetNSTextAttachmentViewProviderClass returns the class object for NSTextAttachmentViewProvider.
func GetNSTextAttachmentViewProviderClass() NSTextAttachmentViewProviderClass {
	return getNSTextAttachmentViewProviderClass()
}

type NSTextAttachmentViewProviderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextAttachmentViewProviderClass) Alloc() NSTextAttachmentViewProvider {
	rv := objc.Send[NSTextAttachmentViewProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container object that associates a text attachment at a particular
// document location with a view object.
//
// # Overview
// 
// Use [NSTextAttachmentViewProvider] when you need to represent document
// locations in terms of an [NSTextLocation] or an [NSTextRange] or you want
// to support view-based text attachments. The view provider controls the view
// placement and layout without requiring view classes to be aware of the text
// attachment coordination using a [NSTextLayoutManager] in macOS 12 or iOS 15
// and later.
//
// # Initializing a text attachment view
//
//   - [NSTextAttachmentViewProvider.InitWithTextAttachmentParentViewTextLayoutManagerLocation]: Creates a new text attachment view whose content starts at the location you provide.
//
// # Defining the contents
//
//   - [NSTextAttachmentViewProvider.Location]: The location that indicates the start of the text attachment.
//   - [NSTextAttachmentViewProvider.TextAttachment]: The text attachment for this view.
//   - [NSTextAttachmentViewProvider.TextLayoutManager]: The text layout manager for this view.
//   - [NSTextAttachmentViewProvider.TracksTextAttachmentViewBounds]: A Boolean value that determines the text attachment’s bounds policy.
//   - [NSTextAttachmentViewProvider.SetTracksTextAttachmentViewBounds]
//   - [NSTextAttachmentViewProvider.View]: The text attachment’s view.
//   - [NSTextAttachmentViewProvider.SetView]
//
// # Defining a custom view hierarchy
//
//   - [NSTextAttachmentViewProvider.LoadView]: Draws the custom view hierarchy that text attachment view subclasses implement.
//
// # Determining the Attachment’s Bounds
//
//   - [NSTextAttachmentViewProvider.AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition]: Returns the layout bounds for an attachment at a specific text location that contains the text attributes you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider
type NSTextAttachmentViewProvider struct {
	objectivec.Object
}

// NSTextAttachmentViewProviderFromID constructs a [NSTextAttachmentViewProvider] from an objc.ID.
//
// A container object that associates a text attachment at a particular
// document location with a view object.
func NSTextAttachmentViewProviderFromID(id objc.ID) NSTextAttachmentViewProvider {
	return NSTextAttachmentViewProvider{objectivec.Object{ID: id}}
}
// NOTE: NSTextAttachmentViewProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextAttachmentViewProvider] class.
//
// # Initializing a text attachment view
//
//   - [INSTextAttachmentViewProvider.InitWithTextAttachmentParentViewTextLayoutManagerLocation]: Creates a new text attachment view whose content starts at the location you provide.
//
// # Defining the contents
//
//   - [INSTextAttachmentViewProvider.Location]: The location that indicates the start of the text attachment.
//   - [INSTextAttachmentViewProvider.TextAttachment]: The text attachment for this view.
//   - [INSTextAttachmentViewProvider.TextLayoutManager]: The text layout manager for this view.
//   - [INSTextAttachmentViewProvider.TracksTextAttachmentViewBounds]: A Boolean value that determines the text attachment’s bounds policy.
//   - [INSTextAttachmentViewProvider.SetTracksTextAttachmentViewBounds]
//   - [INSTextAttachmentViewProvider.View]: The text attachment’s view.
//   - [INSTextAttachmentViewProvider.SetView]
//
// # Defining a custom view hierarchy
//
//   - [INSTextAttachmentViewProvider.LoadView]: Draws the custom view hierarchy that text attachment view subclasses implement.
//
// # Determining the Attachment’s Bounds
//
//   - [INSTextAttachmentViewProvider.AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition]: Returns the layout bounds for an attachment at a specific text location that contains the text attributes you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider
type INSTextAttachmentViewProvider interface {
	objectivec.IObject

	// Topic: Initializing a text attachment view

	// Creates a new text attachment view whose content starts at the location you provide.
	InitWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment INSTextAttachment, parentView INSView, textLayoutManager INSTextLayoutManager, location NSTextLocation) NSTextAttachmentViewProvider

	// Topic: Defining the contents

	// The location that indicates the start of the text attachment.
	Location() NSTextLocation
	// The text attachment for this view.
	TextAttachment() INSTextAttachment
	// The text layout manager for this view.
	TextLayoutManager() INSTextLayoutManager
	// A Boolean value that determines the text attachment’s bounds policy.
	TracksTextAttachmentViewBounds() bool
	SetTracksTextAttachmentViewBounds(value bool)
	// The text attachment’s view.
	View() INSView
	SetView(value INSView)

	// Topic: Defining a custom view hierarchy

	// Draws the custom view hierarchy that text attachment view subclasses implement.
	LoadView()

	// Topic: Determining the Attachment’s Bounds

	// Returns the layout bounds for an attachment at a specific text location that contains the text attributes you specify.
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect
}

// Init initializes the instance.
func (t NSTextAttachmentViewProvider) Init() NSTextAttachmentViewProvider {
	rv := objc.Send[NSTextAttachmentViewProvider](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextAttachmentViewProvider) Autorelease() NSTextAttachmentViewProvider {
	rv := objc.Send[NSTextAttachmentViewProvider](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextAttachmentViewProvider creates a new NSTextAttachmentViewProvider instance.
func NewNSTextAttachmentViewProvider() NSTextAttachmentViewProvider {
	class := getNSTextAttachmentViewProviderClass()
	rv := objc.Send[NSTextAttachmentViewProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new text attachment view whose content starts at the location you
// provide.
//
// textAttachment: The [NSTextAttachment] for this view.
//
// parentView: The parent view of this attachment.
//
// textLayoutManager: The [NSTextLayoutManager] for this view.
//
// location: The [NSTextLocation] that identifies the start of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/init(textAttachment:parentView:textLayoutManager:location:)
func NewTextAttachmentViewProviderWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment INSTextAttachment, parentView INSView, textLayoutManager INSTextLayoutManager, location NSTextLocation) NSTextAttachmentViewProvider {
	instance := getNSTextAttachmentViewProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextAttachment:parentView:textLayoutManager:location:"), textAttachment, parentView, textLayoutManager, location)
	return NSTextAttachmentViewProviderFromID(rv)
}

// Creates a new text attachment view whose content starts at the location you
// provide.
//
// textAttachment: The [NSTextAttachment] for this view.
//
// parentView: The parent view of this attachment.
//
// textLayoutManager: The [NSTextLayoutManager] for this view.
//
// location: The [NSTextLocation] that identifies the start of the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/init(textAttachment:parentView:textLayoutManager:location:)
func (t NSTextAttachmentViewProvider) InitWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment INSTextAttachment, parentView INSView, textLayoutManager INSTextLayoutManager, location NSTextLocation) NSTextAttachmentViewProvider {
	rv := objc.Send[NSTextAttachmentViewProvider](t.ID, objc.Sel("initWithTextAttachment:parentView:textLayoutManager:location:"), textAttachment, parentView, textLayoutManager, location)
	return rv
}

// Draws the custom view hierarchy that text attachment view subclasses
// implement.
//
// # Discussion
// 
// Use this method to create a custom view hierarchy. Don’t call this method
// directly, the framework calls it at the appropriate time.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/loadView()
func (t NSTextAttachmentViewProvider) LoadView() {
	objc.Send[objc.ID](t.ID, objc.Sel("loadView"))
}

// Returns the layout bounds for an attachment at a specific text location
// that contains the text attributes you specify.
//
// attributes: A dictionary that contains a list of key and attribute pairs that describe
// the customization of the string.
//
// location: An [NSTextLocation] that indicates that start of the string.
//
// textContainer: The [NSTextContainer] that contains the source string.
//
// proposedLineFragment: A [CGRect] that describes the boundaries of the line fragment.
// //
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// position: A [CGPoint] inside `proposedLineFragment`.
// //
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// # Return Value
// 
// Returns a [CGRect] that describes the bounds of the attachment.
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/attachmentBounds(for:location:textContainer:proposedLineFragment:position:)
func (t NSTextAttachmentViewProvider) AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, location, textContainer, proposedLineFragment, position)
	return corefoundation.CGRect(rv)
}

// The location that indicates the start of the text attachment.
//
// # Discussion
// 
// Specify the value of this property at initialization time using the
// [InitWithTextAttachmentParentViewTextLayoutManagerLocation] initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/location
func (t NSTextAttachmentViewProvider) Location() NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("location"))
	return NSTextLocationObjectFromID(rv)
}

// The text attachment for this view.
//
// # Discussion
// 
// Specify the value of this property at initialization time using the
// [InitWithTextAttachmentParentViewTextLayoutManagerLocation] initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/textAttachment
func (t NSTextAttachmentViewProvider) TextAttachment() INSTextAttachment {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textAttachment"))
	return NSTextAttachmentFromID(objc.ID(rv))
}

// The text layout manager for this view.
//
// # Discussion
// 
// Specify the value of this property at initialization time using the
// [InitWithTextAttachmentParentViewTextLayoutManagerLocation] initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/textLayoutManager
func (t NSTextAttachmentViewProvider) TextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}

// A Boolean value that determines the text attachment’s bounds policy.
//
// # Discussion
// 
// If `true`, the framework calls the `textAttachment` property’s
// [AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition]
// method and examines the text attachment view provider to determine the
// bounds instead of using the `bounds` property of this instance. Defaults to
// `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/tracksTextAttachmentViewBounds
func (t NSTextAttachmentViewProvider) TracksTextAttachmentViewBounds() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("tracksTextAttachmentViewBounds"))
	return rv
}
func (t NSTextAttachmentViewProvider) SetTracksTextAttachmentViewBounds(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setTracksTextAttachmentViewBounds:"), value)
}

// The text attachment’s view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider/view
func (t NSTextAttachmentViewProvider) View() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTextAttachmentViewProvider) SetView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setView:"), value)
}

