// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that defines the interface to text attachment objects from a layout manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentContainer
type NSTextAttachmentContainer interface {
	objectivec.IObject

	// Returns the layout bounds of the text attachment to the layout manager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentContainer/attachmentBounds(for:proposedLineFragment:glyphPosition:characterIndex:)
	AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect

	// Returns the image object that the layout manager renders in the specified image bounds rectangle inside the text container.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentContainer/image(forBounds:textContainer:characterIndex:)
	ImageForBoundsTextContainerCharacterIndex(imageBounds corefoundation.CGRect, textContainer INSTextContainer, charIndex uint) INSImage
}

// NSTextAttachmentContainerObject wraps an existing Objective-C object that conforms to the NSTextAttachmentContainer protocol.
type NSTextAttachmentContainerObject struct {
	objectivec.Object
}
func (o NSTextAttachmentContainerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextAttachmentContainerObjectFromID constructs a [NSTextAttachmentContainerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextAttachmentContainerObjectFromID(id objc.ID) NSTextAttachmentContainerObject {
	return NSTextAttachmentContainerObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o NSTextAttachmentContainerObject) AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("attachmentBoundsForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), textContainer, lineFrag, position, charIndex)
	return rv
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
func (o NSTextAttachmentContainerObject) ImageForBoundsTextContainerCharacterIndex(imageBounds corefoundation.CGRect, textContainer INSTextContainer, charIndex uint) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("imageForBounds:textContainer:characterIndex:"), imageBounds, textContainer, charIndex)
	return NSImageFromID(rv)
	}

