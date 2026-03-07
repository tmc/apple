// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextAttachment] class.
var (
	_NSTextAttachmentClass     NSTextAttachmentClass
	_NSTextAttachmentClassOnce sync.Once
)

func getNSTextAttachmentClass() NSTextAttachmentClass {
	_NSTextAttachmentClassOnce.Do(func() {
		_NSTextAttachmentClass = NSTextAttachmentClass{class: objc.GetClass("NSTextAttachment")}
	})
	return _NSTextAttachmentClass
}

// GetNSTextAttachmentClass returns the class object for NSTextAttachment.
func GetNSTextAttachmentClass() NSTextAttachmentClass {
	return getNSTextAttachmentClass()
}

type NSTextAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextAttachmentClass) Alloc() NSTextAttachment {
	rv := objc.Send[NSTextAttachment](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The values for the attachment characteristics of attributed strings and
// related objects.
//
// # Overview
// 
// The [NSAttributedString] class uses text attachment objects as the values
// for attachment attributes (stored in the attributed string under the
// [NSTextAttachment.Attachment] key).
// 
// A text attachment object contains either an [NSData] object or an
// [NSTextAttachment.FileWrapper] object, which in turn holds the contents of the attached
// file. The properties of this class configure the appearance of the text
// attachment in your interface. In macOS, the text attachment also uses a
// cell object that conforms to the [NSTextAttachmentCellProtocol] protocol to
// draw the image that represents the text and handles mouse events. For more
// information about text attachments, see the [NSAttributedString] and
// [NSTextView].
// 
// In macOS 12 and iOS 15 and later, [NSTextAttachmentViewProvider] and
// [NSTextAttachmentLayout] provide additional capabilities to represent
// document locations in terms of an [NSTextLocation] or an [NSTextRange], and
// provide support for view-based text attachments.
//
// [NSTextAttachment.FileWrapper]: https://developer.apple.com/documentation/Foundation/FileWrapper
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// # Initializing a text attachment
//
//   - [NSTextAttachment.InitWithFileWrapper]: Creates a text attachment object to contain the specified file wrapper.
//   - [NSTextAttachment.InitWithDataOfType]: Creates a text attachment object with the specified data.
//
// # Defining the attachment’s contents
//
//   - [NSTextAttachment.Bounds]: The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//   - [NSTextAttachment.SetBounds]
//   - [NSTextAttachment.Contents]: The contents for the text attachment.
//   - [NSTextAttachment.SetContents]
//   - [NSTextAttachment.FileType]: The file type of the contents for the text attachment.
//   - [NSTextAttachment.SetFileType]
//   - [NSTextAttachment.Image]: An instance of the relevant image class that represents the contents of the text attachment object.
//   - [NSTextAttachment.SetImage]
//   - [NSTextAttachment.FileWrapper]: The text attachment’s file wrapper.
//   - [NSTextAttachment.SetFileWrapper]
//   - [NSTextAttachment.AllowsTextAttachmentView]: A Boolean value that determines whether the text attachment uses text attachment views.
//   - [NSTextAttachment.SetAllowsTextAttachmentView]
//   - [NSTextAttachment.UsesTextAttachmentView]: A Boolean value that indicates whether the text attachment uses text attachment views.
//   - [NSTextAttachment.LineLayoutPadding]: The layout padding before and after the text attachment bounds.
//   - [NSTextAttachment.SetLineLayoutPadding]
//
// # Setting the attachment cell
//
//   - [NSTextAttachment.AttachmentCell]: The object that draws the icon for the text attachment and handles mouse events.
//   - [NSTextAttachment.SetAttachmentCell]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment
type NSTextAttachment struct {
	objectivec.Object
}

// NSTextAttachmentFromID constructs a [NSTextAttachment] from an objc.ID.
//
// The values for the attachment characteristics of attributed strings and
// related objects.
func NSTextAttachmentFromID(id objc.ID) NSTextAttachment {
	return NSTextAttachment{objectivec.Object{id}}
}
// NOTE: NSTextAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextAttachment] class.
//
// # Initializing a text attachment
//
//   - [INSTextAttachment.InitWithFileWrapper]: Creates a text attachment object to contain the specified file wrapper.
//   - [INSTextAttachment.InitWithDataOfType]: Creates a text attachment object with the specified data.
//
// # Defining the attachment’s contents
//
//   - [INSTextAttachment.Bounds]: The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//   - [INSTextAttachment.SetBounds]
//   - [INSTextAttachment.Contents]: The contents for the text attachment.
//   - [INSTextAttachment.SetContents]
//   - [INSTextAttachment.FileType]: The file type of the contents for the text attachment.
//   - [INSTextAttachment.SetFileType]
//   - [INSTextAttachment.Image]: An instance of the relevant image class that represents the contents of the text attachment object.
//   - [INSTextAttachment.SetImage]
//   - [INSTextAttachment.FileWrapper]: The text attachment’s file wrapper.
//   - [INSTextAttachment.SetFileWrapper]
//   - [INSTextAttachment.AllowsTextAttachmentView]: A Boolean value that determines whether the text attachment uses text attachment views.
//   - [INSTextAttachment.SetAllowsTextAttachmentView]
//   - [INSTextAttachment.UsesTextAttachmentView]: A Boolean value that indicates whether the text attachment uses text attachment views.
//   - [INSTextAttachment.LineLayoutPadding]: The layout padding before and after the text attachment bounds.
//   - [INSTextAttachment.SetLineLayoutPadding]
//
// # Setting the attachment cell
//
//   - [INSTextAttachment.AttachmentCell]: The object that draws the icon for the text attachment and handles mouse events.
//   - [INSTextAttachment.SetAttachmentCell]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment
type INSTextAttachment interface {
	objectivec.IObject
	NSTextAttachmentContainer
	NSTextAttachmentLayout

	// Topic: Initializing a text attachment

	// Creates a text attachment object to contain the specified file wrapper.
	InitWithFileWrapper(fileWrapper *foundation.NSFileWrapper) NSTextAttachment
	// Creates a text attachment object with the specified data.
	InitWithDataOfType(contentData foundation.INSData, uti string) NSTextAttachment

	// Topic: Defining the attachment’s contents

	// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	// The contents for the text attachment.
	Contents() foundation.INSData
	SetContents(value foundation.INSData)
	// The file type of the contents for the text attachment.
	FileType() string
	SetFileType(value string)
	// An instance of the relevant image class that represents the contents of the text attachment object.
	Image() INSImage
	SetImage(value INSImage)
	// The text attachment’s file wrapper.
	FileWrapper() *foundation.NSFileWrapper
	SetFileWrapper(value *foundation.NSFileWrapper)
	// A Boolean value that determines whether the text attachment uses text attachment views.
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	// A Boolean value that indicates whether the text attachment uses text attachment views.
	UsesTextAttachmentView() bool
	// The layout padding before and after the text attachment bounds.
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)

	// Topic: Setting the attachment cell

	// The object that draws the icon for the text attachment and handles mouse events.
	AttachmentCell() NSTextAttachmentCell
	SetAttachmentCell(value NSTextAttachmentCell)

	// The attachment for the text.
	Attachment() foundation.NSString
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTextAttachment) Init() NSTextAttachment {
	rv := objc.Send[NSTextAttachment](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextAttachment) Autorelease() NSTextAttachment {
	rv := objc.Send[NSTextAttachment](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextAttachment creates a new NSTextAttachment instance.
func NewNSTextAttachment() NSTextAttachment {
	class := getNSTextAttachmentClass()
	rv := objc.Send[NSTextAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a text attachment object with the specified data.
//
// contentData: Data to use for the text attachment contents. Can be `nil`.
//
// uti: A uniform type identifier specifying the data type of the attachment
// contents. Can be `nil`.
//
// # Return Value
// 
// A new [NSTextAttachment] object.
//
// # Discussion
// 
// This method is the designated initializer for the [NSTextAttachment] class
// on iOS.
// 
// When either `contentData` or `uti` is `nil`, TextKit considers the receiver
// to be an attachment without document contents. In this case, the
// [NSAttributedString] external file writing methods try to save the value of
// the [Image] property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(data:ofType:)
func NewTextAttachmentWithDataOfType(contentData foundation.INSData, uti string) NSTextAttachment {
	instance := getNSTextAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:ofType:"), contentData, objc.String(uti))
	return NSTextAttachmentFromID(rv)
}


// Creates a text attachment object to contain the specified file wrapper.
//
// fileWrapper: The file wrapper for the attachment.
//
// # Return Value
// 
// A new text attachment object initialized with the file wrapper.
//
// # Discussion
// 
// This method is the designated initializer for the [NSTextAttachment] class
// in macOS.
// 
// If `aWrapper` contains an image file that the receiver can interpret as an
// [NSImage] object, this method sets the attachment cell’s image to that
// image rather than to the icon of `aWrapper`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(fileWrapper:)
func NewTextAttachmentWithFileWrapper(fileWrapper *foundation.NSFileWrapper) NSTextAttachment {
	instance := getNSTextAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileWrapper:"), fileWrapper)
	return NSTextAttachmentFromID(rv)
}







// Creates a text attachment object to contain the specified file wrapper.
//
// fileWrapper: The file wrapper for the attachment.
//
// # Return Value
// 
// A new text attachment object initialized with the file wrapper.
//
// # Discussion
// 
// This method is the designated initializer for the [NSTextAttachment] class
// in macOS.
// 
// If `aWrapper` contains an image file that the receiver can interpret as an
// [NSImage] object, this method sets the attachment cell’s image to that
// image rather than to the icon of `aWrapper`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(fileWrapper:)
func (t NSTextAttachment) InitWithFileWrapper(fileWrapper *foundation.NSFileWrapper) NSTextAttachment {
	rv := objc.Send[NSTextAttachment](t.ID, objc.Sel("initWithFileWrapper:"), fileWrapper)
	return rv
}

// Creates a text attachment object with the specified data.
//
// contentData: Data to use for the text attachment contents. Can be `nil`.
//
// uti: A uniform type identifier specifying the data type of the attachment
// contents. Can be `nil`.
//
// # Return Value
// 
// A new [NSTextAttachment] object.
//
// # Discussion
// 
// This method is the designated initializer for the [NSTextAttachment] class
// on iOS.
// 
// When either `contentData` or `uti` is `nil`, TextKit considers the receiver
// to be an attachment without document contents. In this case, the
// [NSAttributedString] external file writing methods try to save the value of
// the [Image] property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(data:ofType:)
func (t NSTextAttachment) InitWithDataOfType(contentData foundation.INSData, uti string) NSTextAttachment {
	rv := objc.Send[NSTextAttachment](t.ID, objc.Sel("initWithData:ofType:"), contentData, objc.String(uti))
	return rv
}

// Returns the layout bounds of the text attachment to the layout manager.
//
// textContainer: The text container for the text being laid out.
//
// lineFrag: The line fragment containing the text attachment.
//
// position: The glyph location inside `lineFrag` which is the origin of the returned
// bounds rectangle.
//
// charIndex: The character location inside the text storage for the attachment
// character.
//
// # Return Value
// 
// The [Bounds] rectangle of the text attachment if not [CGRectZero];
// otherwise, the rectangle of the [size] property of the attachment’s
// [Image] property.
//
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
// [size]: https://developer.apple.com/documentation/UIKit/UIImage/size
//
// # Discussion
// 
// Conforming objects can implement more sophisticated logic for negotiating
// the attachment bounds based on the available container space and proposed
// line fragment rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentContainer/attachmentBounds(for:proposedLineFragment:glyphPosition:characterIndex:)
func (t NSTextAttachment) AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("attachmentBoundsForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), textContainer, lineFrag, position, charIndex)
	return corefoundation.CGRect(rv)
}

// Returns the image object that the layout manager renders in the specified
// image bounds rectangle inside the text container.
//
// imageBounds: The rectangle in which the image is laid out.
//
// textContainer: The text container in which the image is laid out.
//
// charIndex: The character location inside the text storage for the attachment
// character.
//
// # Return Value
// 
// The image rendered in the bounds rectangle.
//
// # Discussion
// 
// The method should return an image appropriate for the target rendering
// context derived by arguments passed into this method. The
// [NSTextAttachment] implementation returns the text attachment’s [Image]
// when non-`nil`. If the image is `nil`, it returns an image based on the
// text attachment’s [Contents] and [FileType] properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentContainer/image(forBounds:textContainer:characterIndex:)
func (t NSTextAttachment) ImageForBoundsTextContainerCharacterIndex(imageBounds corefoundation.CGRect, textContainer INSTextContainer, charIndex uint) INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("imageForBounds:textContainer:characterIndex:"), imageBounds, textContainer, charIndex)
	return NSImageFromID(rv)
}

// Returns the text attachment view provider corresponding to the file type.
//
// parentView: The parent view.
//
// location: An <[NSTextLocation] that indicates that start of the string.
//
// textContainer: The[NSTextContainer] that contains the source text.
//
// # Return Value
// 
// An [NSTextAttachmentViewProvider].
//
// # Discussion
// 
// The default implementation queries the text attachment view provider class
// using the [TextAttachmentViewProviderClassForFileType] method of
// [NSTextAttachment]. When non-`nil`, it instantiates a view, then, fills
// properties declared in [NSTextAttachmentViewProvider] if implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/viewProvider(for:location:textContainer:)
func (t NSTextAttachment) ViewProviderForParentViewLocationTextContainer(parentView INSView, location NSTextLocation, textContainer INSTextContainer) INSTextAttachmentViewProvider {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewProviderForParentView:location:textContainer:"), parentView, location, textContainer)
	return NSTextAttachmentViewProviderFromID(rv)
}

// Returns the image object rendered at the bounds and inside the text
// container you specify.
//
// bounds: The [CGRect] that presents the image boundaries inside `textContainer`.
// //
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// attributes: A dictionary of [NSAttributedString.Key] attributes.
// //
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
//
// location: An [NSTextLocation] that indicates that start of the string.
//
// textContainer: The [NSTextContainer] that contains the source text.
//
// # Return Value
// 
// An optional image object.
//
// # Discussion
// 
// A custom implementation should return an image appropriate for the target
// rendering context that you derive by arguments to this method. The default
// [NSTextAttachment] implementation returns the contents of the `image`
// property when non-`nil`. If the `image` property is `nil`, it returns an
// image based on the `contents` and `fileType` properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/image(for:attributes:location:textContainer:)
func (t NSTextAttachment) ImageForBoundsAttributesLocationTextContainer(bounds corefoundation.CGRect, attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer) INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("imageForBounds:attributes:location:textContainer:"), bounds, attributes, location, textContainer)
	return NSImageFromID(rv)
}

// Returns the layout bounds of the attachment you specify.
//
// attributes: A dictionary of [NSAttributedString.Key] attributes.
// //
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
//
// location: An [NSTextLocation] that indicates that start of the string.
//
// textContainer: The [NSTextContainer] that contains the source text.
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
// Returns a [CGRect] that describes the boundaries of the attachment, or
// `CGRectZero.`
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// # Discussion
// 
// The framework interprets the bounds origin to match `position` inside
// `proposedLineFragment`. The default [NSTextAttachment] implementation
// returns bounds if the value isn’t equivalent to [CGRectZero]; otherwise,
// it derives the bounds value from `image.Size()`. Conforming objects can
// implement more sophisticated logic for negotiating the frame size based on
// the available container space and proposed line fragment rectangle.
//
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/attachmentBounds(for:location:textContainer:proposedLineFragment:position:)
func (t NSTextAttachment) AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, location, textContainer, proposedLineFragment, position)
	return corefoundation.CGRect(rv)
}
func (t NSTextAttachment) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Registers a specific file type with the attachment view provider.
//
// textAttachmentViewProviderClass: The text attachment view provider class.
//
// fileType: A [String] that represents the file type.
// //
// [String]: https://developer.apple.com/documentation/Swift/String
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/registerViewProviderClass(_:forFileType:)
func (_NSTextAttachmentClass NSTextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.Class, fileType string) {
	objc.Send[objc.ID](objc.ID(_NSTextAttachmentClass.class), objc.Sel("registerTextAttachmentViewProviderClass:forFileType:"), textAttachmentViewProviderClass, objc.String(fileType))
}

// Returns the text attachment view provider class, if any, for the file type
// you specify.
//
// fileType: A [String] that represents the file type.
// //
// [String]: https://developer.apple.com/documentation/Swift/String
//
// # Return Value
// 
// The text attachment view provider class, or `nil` if the there is no class
// for the specified file type.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/textAttachmentViewProviderClass(forFileType:)
func (_NSTextAttachmentClass NSTextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSTextAttachmentClass.class), objc.Sel("textAttachmentViewProviderClassForFileType:"), objc.String(fileType))
	return rv
}








// The layout bounds of the text attachment’s graphical representation in
// the text coordinate system.
//
// # Discussion
// 
// The bounds rectangle origin is at the current glyph location on the text
// baseline. The default value is [CGRectZero].
//
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t NSTextAttachment) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (t NSTextAttachment) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](t.ID, objc.Sel("setBounds:"), value)
}



// The contents for the text attachment.
//
// # Discussion
// 
// Modifying this property has the side effect of invalidating the [Image]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t NSTextAttachment) Contents() foundation.INSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("contents"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t NSTextAttachment) SetContents(value foundation.INSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setContents:"), value)
}



// The file type of the contents for the text attachment.
//
// # Discussion
// 
// Modifying this property has the side effect of invalidating the [Image]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t NSTextAttachment) FileType() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("fileType"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTextAttachment) SetFileType(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setFileType:"), objc.String(value))
}



// An instance of the relevant image class that represents the contents of the
// text attachment object.
//
// # Discussion
// 
// For details about using the [UIImage] class to create text attachments that
// automatically adjust to surrounding font and color attributes, see the
// [init(image:)] initializer.
//
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
// [init(image:)]: https://developer.apple.com/documentation/UIKit/NSTextAttachment/init(image:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/image
func (t NSTextAttachment) Image() INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (t NSTextAttachment) SetImage(value INSImage) {
	objc.Send[struct{}](t.ID, objc.Sel("setImage:"), value)
}



// The text attachment’s file wrapper.
//
// # Discussion
// 
// The file wrapper holds the contents of the attached file. In iOS, modifying
// this property has a side effect of invalidating the [Image], [Contents],
// and [FileType] properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t NSTextAttachment) FileWrapper() *foundation.NSFileWrapper {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("fileWrapper"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSFileWrapperFromID(objc.ID(rv))
	return &val
}
func (t NSTextAttachment) SetFileWrapper(value *foundation.NSFileWrapper) {
	objc.Send[struct{}](t.ID, objc.Sel("setFileWrapper:"), value)
}



// A Boolean value that determines whether the text attachment uses text
// attachment views.
//
// # Discussion
// 
// When `true`, the text attachment tries to use a text attachment view
// returned by [ViewProviderForParentViewLocationTextContainer]. Default is
// `true`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t NSTextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsTextAttachmentView"))
	return rv
}
func (t NSTextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsTextAttachmentView:"), value)
}



// A Boolean value that indicates whether the text attachment uses text
// attachment views.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/usesTextAttachmentView
func (t NSTextAttachment) UsesTextAttachmentView() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesTextAttachmentView"))
	return rv
}



// The layout padding before and after the text attachment bounds.
//
// # Discussion
// 
// The layout and rendering bounds X origin is inset by the padding value.
// This affects the relationship between the text attachment bounds and
// [NSLayoutManager] glyph metrics methods [LocationForGlyphAtIndex] and
// [AttachmentSizeForGlyphAtIndex]. The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t NSTextAttachment) LineLayoutPadding() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("lineLayoutPadding"))
	return rv
}
func (t NSTextAttachment) SetLineLayoutPadding(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setLineLayoutPadding:"), value)
}



// The object that draws the icon for the text attachment and handles mouse
// events.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t NSTextAttachment) AttachmentCell() NSTextAttachmentCell {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attachmentCell"))
	return NSTextAttachmentCellFromID(objc.ID(rv))
}
func (t NSTextAttachment) SetAttachmentCell(value NSTextAttachmentCell) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttachmentCell:"), value)
}



// The attachment for the text.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/attachment
func (t NSTextAttachment) Attachment() foundation.NSString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attachment"))
	return foundation.NSStringFromID(objc.ID(rv))
}







// Specifies a character that denotes an attachment.
//
// See: https://developer.apple.com/documentation/appkit/nstextattachment/character
func (_NSTextAttachmentClass NSTextAttachmentClass) Character() int {
	rv := objc.Send[int](objc.ID(_NSTextAttachmentClass.class), objc.Sel("NSAttachmentCharacter"))
	return rv
}
func (_NSTextAttachmentClass NSTextAttachmentClass) SetCharacter(value int) {
	objc.Send[struct{}](objc.ID(_NSTextAttachmentClass.class), objc.Sel("setNSAttachmentCharacter:"), value)
}













			// Protocol methods for NSTextAttachmentContainer
			




			// Protocol methods for NSTextAttachmentLayout
			













