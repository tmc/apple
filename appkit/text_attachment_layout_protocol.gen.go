// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that defines the interface to attachment objects from a text layout manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout
type NSTextAttachmentLayout interface {
	objectivec.IObject

	// Returns the layout bounds of the attachment you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/attachmentBounds(for:location:textContainer:proposedLineFragment:position:)
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect

	// Returns the image object rendered at the bounds and inside the text container you specify.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/image(for:attributes:location:textContainer:)
	ImageForBoundsAttributesLocationTextContainer(bounds corefoundation.CGRect, attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer) INSImage

	// Returns the text attachment view provider corresponding to the file type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/viewProvider(for:location:textContainer:)
	ViewProviderForParentViewLocationTextContainer(parentView INSView, location NSTextLocation, textContainer INSTextContainer) INSTextAttachmentViewProvider
}

// NSTextAttachmentLayoutObject wraps an existing Objective-C object that conforms to the NSTextAttachmentLayout protocol.
type NSTextAttachmentLayoutObject struct {
	objectivec.Object
}

func (o NSTextAttachmentLayoutObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextAttachmentLayoutObjectFromID constructs a [NSTextAttachmentLayoutObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextAttachmentLayoutObjectFromID(id objc.ID) NSTextAttachmentLayoutObject {
	return NSTextAttachmentLayoutObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the layout bounds of the attachment you specify.
//
// attributes: A dictionary of [NSAttributedString.Key] attributes.
//
// location: An [NSTextLocation] that indicates that start of the string.
//
// textContainer: The [NSTextContainer] that contains the source text.
//
// proposedLineFragment: A [CGRect] that describes the boundaries of the line fragment.
//
// position: A [CGPoint] inside `proposedLineFragment`.
//
// # Return Value
//
// Returns a [CGRect] that describes the boundaries of the attachment, or
// `CGRectZero.`
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
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentLayout/attachmentBounds(for:location:textContainer:proposedLineFragment:position:)
//
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (o NSTextAttachmentLayoutObject) AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, location, textContainer, proposedLineFragment, position)
	return rv
}

// Returns the image object rendered at the bounds and inside the text
// container you specify.
//
// bounds: The [CGRect] that presents the image boundaries inside `textContainer`.
//
// attributes: A dictionary of [NSAttributedString.Key] attributes.
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
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
func (o NSTextAttachmentLayoutObject) ImageForBoundsAttributesLocationTextContainer(bounds corefoundation.CGRect, attributes foundation.INSDictionary, location NSTextLocation, textContainer INSTextContainer) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("imageForBounds:attributes:location:textContainer:"), bounds, attributes, location, textContainer)
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
func (o NSTextAttachmentLayoutObject) ViewProviderForParentViewLocationTextContainer(parentView INSView, location NSTextLocation, textContainer INSTextContainer) INSTextAttachmentViewProvider {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("viewProviderForParentView:location:textContainer:"), parentView, location, textContainer)
	return NSTextAttachmentViewProviderFromID(rv)
}
