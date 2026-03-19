// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [NSAdaptiveImageGlyph] class.
var (
	_NSAdaptiveImageGlyphClass     NSAdaptiveImageGlyphClass
	_NSAdaptiveImageGlyphClassOnce sync.Once
)

func getNSAdaptiveImageGlyphClass() NSAdaptiveImageGlyphClass {
	_NSAdaptiveImageGlyphClassOnce.Do(func() {
		_NSAdaptiveImageGlyphClass = NSAdaptiveImageGlyphClass{class: objc.GetClass("NSAdaptiveImageGlyph")}
	})
	return _NSAdaptiveImageGlyphClass
}

// GetNSAdaptiveImageGlyphClass returns the class object for NSAdaptiveImageGlyph.
func GetNSAdaptiveImageGlyphClass() NSAdaptiveImageGlyphClass {
	return getNSAdaptiveImageGlyphClass()
}

type NSAdaptiveImageGlyphClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAdaptiveImageGlyphClass) Alloc() NSAdaptiveImageGlyph {
	rv := objc.Send[NSAdaptiveImageGlyph](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A data object for an emoji-like image that can appear in attributed text.
//
// # Overview
// 
// An [NSAdaptiveImageGlyph] contains an image that automatically adapts to
// different sizes and resolutions. The text system creates instances of this
// type to represent custom emojis that people create using the system
// interfaces. This type manages multiple images, along with metadata
// describing how to adapt those images correctly to different fonts and font
// attributes.
// 
// Typically, you receive new [NSAdaptiveImageGlyph] objects only from the
// text-input system. When someone creates a new emoji and inserts it into
// their text, TextKit creates an instance of this type to represent it. If
// your app examines or changes the attributes of attributed strings, preserve
// the [AdaptiveImageGlyph] attribute when making any changes. For example, if
// you filter unknown attributes in a custom text-storage object, update your
// code to preserve this attribute. The value of the attribute is an
// [NSAdaptiveImageGlyph] containing the emoji data. You can save the image
// data with the rest of your content and use the data to recreate the type
// later.
//
// # Creating an adaptive image glyph
//
//   - [NSAdaptiveImageGlyph.InitWithImageContent]: Create an adaptive image glyph from the previously saved data.
//   - [NSAdaptiveImageGlyph.InitWithCoder]
//
// # Getting the image data
//
//   - [NSAdaptiveImageGlyph.ImageContent]: The raw data for the image.
//
// # Getting the content metadata
//
//   - [NSAdaptiveImageGlyph.ContentIdentifier]: A unique identifier for this image.
//   - [NSAdaptiveImageGlyph.ContentDescription]: An alternate textual description of the image contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph
type NSAdaptiveImageGlyph struct {
	objectivec.Object
}

// NSAdaptiveImageGlyphFromID constructs a [NSAdaptiveImageGlyph] from an objc.ID.
//
// A data object for an emoji-like image that can appear in attributed text.
func NSAdaptiveImageGlyphFromID(id objc.ID) NSAdaptiveImageGlyph {
	return NSAdaptiveImageGlyph{objectivec.Object{ID: id}}
}
// NOTE: NSAdaptiveImageGlyph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAdaptiveImageGlyph] class.
//
// # Creating an adaptive image glyph
//
//   - [INSAdaptiveImageGlyph.InitWithImageContent]: Create an adaptive image glyph from the previously saved data.
//   - [INSAdaptiveImageGlyph.InitWithCoder]
//
// # Getting the image data
//
//   - [INSAdaptiveImageGlyph.ImageContent]: The raw data for the image.
//
// # Getting the content metadata
//
//   - [INSAdaptiveImageGlyph.ContentIdentifier]: A unique identifier for this image.
//   - [INSAdaptiveImageGlyph.ContentDescription]: An alternate textual description of the image contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph
type INSAdaptiveImageGlyph interface {
	objectivec.IObject

	// Topic: Creating an adaptive image glyph

	// Create an adaptive image glyph from the previously saved data.
	InitWithImageContent(imageContent foundation.INSData) NSAdaptiveImageGlyph
	InitWithCoder(coder foundation.INSCoder) NSAdaptiveImageGlyph

	// Topic: Getting the image data

	// The raw data for the image.
	ImageContent() foundation.INSData

	// Topic: Getting the content metadata

	// A unique identifier for this image.
	ContentIdentifier() string
	// An alternate textual description of the image contents.
	ContentDescription() string

	// The adaptive image glyph for the text.
	AdaptiveImageGlyph() foundation.NSString
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a NSAdaptiveImageGlyph) Init() NSAdaptiveImageGlyph {
	rv := objc.Send[NSAdaptiveImageGlyph](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAdaptiveImageGlyph) Autorelease() NSAdaptiveImageGlyph {
	rv := objc.Send[NSAdaptiveImageGlyph](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAdaptiveImageGlyph creates a new NSAdaptiveImageGlyph instance.
func NewNSAdaptiveImageGlyph() NSAdaptiveImageGlyph {
	class := getNSAdaptiveImageGlyphClass()
	rv := objc.Send[NSAdaptiveImageGlyph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(coder:)
func NewAdaptiveImageGlyphWithCoder(coder foundation.INSCoder) NSAdaptiveImageGlyph {
	instance := getNSAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAdaptiveImageGlyphFromID(rv)
}

// Create an adaptive image glyph from the previously saved data.
//
// imageContent: The raw image data you obtained previously from an adaptive image glyph.
// Typically, you receive adaptive images from the text system, store their
// data with the rest of your content, and use the data to recreate the
// adaptive image later.
//
// # Return Value
// 
// A new adaptive image glyph with the identifier and details from the image
// data.
//
// # Discussion
// 
// Use this initializer to create an adaptive image glyph from data you
// previously saved.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(imageContent:)
func NewAdaptiveImageGlyphWithImageContent(imageContent foundation.INSData) NSAdaptiveImageGlyph {
	instance := getNSAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImageContent:"), imageContent)
	return NSAdaptiveImageGlyphFromID(rv)
}

// Create an adaptive image glyph from the previously saved data.
//
// imageContent: The raw image data you obtained previously from an adaptive image glyph.
// Typically, you receive adaptive images from the text system, store their
// data with the rest of your content, and use the data to recreate the
// adaptive image later.
//
// # Return Value
// 
// A new adaptive image glyph with the identifier and details from the image
// data.
//
// # Discussion
// 
// Use this initializer to create an adaptive image glyph from data you
// previously saved.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(imageContent:)
func (a NSAdaptiveImageGlyph) InitWithImageContent(imageContent foundation.INSData) NSAdaptiveImageGlyph {
	rv := objc.Send[NSAdaptiveImageGlyph](a.ID, objc.Sel("initWithImageContent:"), imageContent)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(coder:)
func (a NSAdaptiveImageGlyph) InitWithCoder(coder foundation.INSCoder) NSAdaptiveImageGlyph {
	rv := objc.Send[NSAdaptiveImageGlyph](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a NSAdaptiveImageGlyph) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The raw data for the image.
//
// # Discussion
// 
// This property contains the image data, the unique identifier for the image,
// the image description, and additional metadata. When saving your content to
// disk, save the data for any adaptive images with the rest of your content.
// If you need to specify a type for the image data, use the value in the
// [ContentType] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/imageContent
func (a NSAdaptiveImageGlyph) ImageContent() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("imageContent"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A unique identifier for this image.
//
// # Discussion
// 
// Use this property to create a persistent reference to this specific image
// in your code. The image data contains this content identifier, so the value
// persists between instantiations.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentIdentifier
func (a NSAdaptiveImageGlyph) ContentIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contentIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// An alternate textual description of the image contents.
//
// # Discussion
// 
// This string contains a brief description of the image, which is useful for
// searches or places where you need a text-based description. The adaptive
// image derives the content of this property from the underlying image data.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentDescription
func (a NSAdaptiveImageGlyph) ContentDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contentDescription"))
	return foundation.NSStringFromID(rv).String()
}
// The adaptive image glyph for the text.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/adaptiveImageGlyph
func (a NSAdaptiveImageGlyph) AdaptiveImageGlyph() foundation.NSString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("adaptiveImageGlyph"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The image data format to use for this image type.
//
// # Discussion
// 
// Use this type when you need to specify the type of the image data. Adaptive
// images are compatible with the HEIC format, but include extra metadata
// about the supported resolutions and sizes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentType
func (_NSAdaptiveImageGlyphClass NSAdaptiveImageGlyphClass) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](objc.ID(_NSAdaptiveImageGlyphClass.class), objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(objc.ID(rv))
}

