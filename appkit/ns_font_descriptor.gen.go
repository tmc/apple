// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFontDescriptor] class.
var (
	_NSFontDescriptorClass     NSFontDescriptorClass
	_NSFontDescriptorClassOnce sync.Once
)

func getNSFontDescriptorClass() NSFontDescriptorClass {
	_NSFontDescriptorClassOnce.Do(func() {
		_NSFontDescriptorClass = NSFontDescriptorClass{class: objc.GetClass("NSFontDescriptor")}
	})
	return _NSFontDescriptorClass
}

// GetNSFontDescriptorClass returns the class object for NSFontDescriptor.
func GetNSFontDescriptorClass() NSFontDescriptorClass {
	return getNSFontDescriptorClass()
}

type NSFontDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontDescriptorClass) Alloc() NSFontDescriptor {
	rv := objc.Send[NSFontDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dictionary of attributes that describe a font.
//
// # Overview
// 
// A font descriptor can be used to create or modify an [NSFont] object. The
// system provides a font matching capability, so that you can partially
// describe a font by creating a font descriptor with, for example, just a
// family name. You can then find all the available fonts on the system with a
// matching family name using [NSFontDescriptor.MatchingFontDescriptorsWithMandatoryKeys].
// 
// There are several ways to create a new [NSFontDescriptor] object. You can
// use `alloc` and [NSFontDescriptor.InitWithFontAttributes],
// [fontDescriptorWithFontAttributes:], [NSFontDescriptor.FontDescriptorWithNameMatrix], or
// [NSFontDescriptor.FontDescriptorWithNameSize]. to create a font descriptor based on either
// your custom attributes dictionary or on a specific font’s name and size.
// Alternatively you can use one of the `fontDescriptor…` instance methods
// (such as [NSFontDescriptor.FontDescriptorWithFace]) to create a modified version of an
// existing descriptor. The latter methods are useful if you have an existing
// descriptor and simply want to change one aspect.
// 
// All attributes in the attributes dictionary are optional.
//
// [fontDescriptorWithFontAttributes:]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontDescriptorWithFontAttributes:
//
// # Creating a Font Descriptor
//
//   - [NSFontDescriptor.InitWithFontAttributes]: Initializes and returns a new font descriptor with the specified attributes.
//
// # Modifying an Existing Font Descriptor
//
//   - [NSFontDescriptor.FontDescriptorByAddingAttributes]: Returns a new font descriptor based on the current object, but with the specified attributes taking precedence over the existing ones.
//   - [NSFontDescriptor.FontDescriptorWithFace]: Returns a new font descriptor based on the current object, but with the specified face.
//   - [NSFontDescriptor.FontDescriptorWithFamily]: Returns a new font descriptor based on the current object, but with the specified font family.
//   - [NSFontDescriptor.FontDescriptorWithMatrix]: Returns a new font descriptor based on the current object, but with the specified font matrix.
//   - [NSFontDescriptor.FontDescriptorWithSize]: Returns a new font descriptor based on the current object, but with the specified point size.
//   - [NSFontDescriptor.FontDescriptorWithSymbolicTraits]: Returns a new font descriptor based on the current object, but with the specified symbolic traits taking precedence over the existing ones.
//   - [NSFontDescriptor.FontDescriptorWithDesign]: Returns a new font descriptor based on the current object, but with the specified design style.
//
// # Finding Fonts
//
//   - [NSFontDescriptor.MatchingFontDescriptorsWithMandatoryKeys]: Returns all the fonts available on the system whose specified attributes match those of the receiver.
//   - [NSFontDescriptor.MatchingFontDescriptorWithMandatoryKeys]: Returns a normalized font descriptor whose specified attributes match those of the receiver.
//
// # Getting the Font Attributes
//
//   - [NSFontDescriptor.FontAttributes]: The receiver’s dictionary of attributes.
//   - [NSFontDescriptor.ObjectForKey]: Returns the font attribute specified by the given key.
//   - [NSFontDescriptor.Matrix]: The current transform matrix of the receiver.
//   - [NSFontDescriptor.PointSize]: The point size of the receiver.
//   - [NSFontDescriptor.PostscriptName]: The PostScript name of the receiver.
//   - [NSFontDescriptor.NSFontFamilyClassMask]: Constant you use to access 
//   - [NSFontDescriptor.SetNSFontFamilyClassMask]
//
// # Getting the Font Traits
//
//   - [NSFontDescriptor.SymbolicTraits]: A bit mask that describes the traits of the receiver.
//
// # Requiring Font Assets
//
//   - [NSFontDescriptor.RequiresFontAssetRequest]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor
type NSFontDescriptor struct {
	objectivec.Object
}

// NSFontDescriptorFromID constructs a [NSFontDescriptor] from an objc.ID.
//
// A dictionary of attributes that describe a font.
func NSFontDescriptorFromID(id objc.ID) NSFontDescriptor {
	return NSFontDescriptor{objectivec.Object{ID: id}}
}
// NOTE: NSFontDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFontDescriptor] class.
//
// # Creating a Font Descriptor
//
//   - [INSFontDescriptor.InitWithFontAttributes]: Initializes and returns a new font descriptor with the specified attributes.
//
// # Modifying an Existing Font Descriptor
//
//   - [INSFontDescriptor.FontDescriptorByAddingAttributes]: Returns a new font descriptor based on the current object, but with the specified attributes taking precedence over the existing ones.
//   - [INSFontDescriptor.FontDescriptorWithFace]: Returns a new font descriptor based on the current object, but with the specified face.
//   - [INSFontDescriptor.FontDescriptorWithFamily]: Returns a new font descriptor based on the current object, but with the specified font family.
//   - [INSFontDescriptor.FontDescriptorWithMatrix]: Returns a new font descriptor based on the current object, but with the specified font matrix.
//   - [INSFontDescriptor.FontDescriptorWithSize]: Returns a new font descriptor based on the current object, but with the specified point size.
//   - [INSFontDescriptor.FontDescriptorWithSymbolicTraits]: Returns a new font descriptor based on the current object, but with the specified symbolic traits taking precedence over the existing ones.
//   - [INSFontDescriptor.FontDescriptorWithDesign]: Returns a new font descriptor based on the current object, but with the specified design style.
//
// # Finding Fonts
//
//   - [INSFontDescriptor.MatchingFontDescriptorsWithMandatoryKeys]: Returns all the fonts available on the system whose specified attributes match those of the receiver.
//   - [INSFontDescriptor.MatchingFontDescriptorWithMandatoryKeys]: Returns a normalized font descriptor whose specified attributes match those of the receiver.
//
// # Getting the Font Attributes
//
//   - [INSFontDescriptor.FontAttributes]: The receiver’s dictionary of attributes.
//   - [INSFontDescriptor.ObjectForKey]: Returns the font attribute specified by the given key.
//   - [INSFontDescriptor.Matrix]: The current transform matrix of the receiver.
//   - [INSFontDescriptor.PointSize]: The point size of the receiver.
//   - [INSFontDescriptor.PostscriptName]: The PostScript name of the receiver.
//   - [INSFontDescriptor.NSFontFamilyClassMask]: Constant you use to access 
//   - [INSFontDescriptor.SetNSFontFamilyClassMask]
//
// # Getting the Font Traits
//
//   - [INSFontDescriptor.SymbolicTraits]: A bit mask that describes the traits of the receiver.
//
// # Requiring Font Assets
//
//   - [INSFontDescriptor.RequiresFontAssetRequest]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor
type INSFontDescriptor interface {
	objectivec.IObject

	// Topic: Creating a Font Descriptor

	// Initializes and returns a new font descriptor with the specified attributes.
	InitWithFontAttributes(attributes foundation.INSDictionary) NSFontDescriptor

	// Topic: Modifying an Existing Font Descriptor

	// Returns a new font descriptor based on the current object, but with the specified attributes taking precedence over the existing ones.
	FontDescriptorByAddingAttributes(attributes foundation.INSDictionary) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified face.
	FontDescriptorWithFace(newFace string) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified font family.
	FontDescriptorWithFamily(newFamily string) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified font matrix.
	FontDescriptorWithMatrix(matrix foundation.NSAffineTransform) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified point size.
	FontDescriptorWithSize(newPointSize float64) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified symbolic traits taking precedence over the existing ones.
	FontDescriptorWithSymbolicTraits(symbolicTraits NSFontDescriptorSymbolicTraits) INSFontDescriptor
	// Returns a new font descriptor based on the current object, but with the specified design style.
	FontDescriptorWithDesign(design NSFontDescriptorSystemDesign) INSFontDescriptor

	// Topic: Finding Fonts

	// Returns all the fonts available on the system whose specified attributes match those of the receiver.
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.INSSet) []NSFontDescriptor
	// Returns a normalized font descriptor whose specified attributes match those of the receiver.
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.INSSet) INSFontDescriptor

	// Topic: Getting the Font Attributes

	// The receiver’s dictionary of attributes.
	FontAttributes() foundation.INSDictionary
	// Returns the font attribute specified by the given key.
	ObjectForKey(attribute NSFontDescriptorAttributeName) objectivec.IObject
	// The current transform matrix of the receiver.
	Matrix() foundation.NSAffineTransform
	// The point size of the receiver.
	PointSize() float64
	// The PostScript name of the receiver.
	PostscriptName() string
	// Constant you use to access 
	NSFontFamilyClassMask() uint32
	SetNSFontFamilyClassMask(value uint32)

	// Topic: Getting the Font Traits

	// A bit mask that describes the traits of the receiver.
	SymbolicTraits() NSFontDescriptorSymbolicTraits

	// Topic: Requiring Font Assets

	RequiresFontAssetRequest() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFontDescriptor) Init() NSFontDescriptor {
	rv := objc.Send[NSFontDescriptor](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFontDescriptor) Autorelease() NSFontDescriptor {
	rv := objc.Send[NSFontDescriptor](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFontDescriptor creates a new NSFontDescriptor instance.
func NewNSFontDescriptor() NSFontDescriptor {
	class := getNSFontDescriptorClass()
	rv := objc.Send[NSFontDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new font descriptor with the specified
// attributes.
//
// attributes: The attributes for the new font descriptor. If `nil`, the font
// descriptor’s attribute dictionary will be empty.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(fontAttributes:)
func NewFontDescriptorWithFontAttributes(attributes foundation.INSDictionary) NSFontDescriptor {
	instance := getNSFontDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFontAttributes:"), attributes)
	return NSFontDescriptorFromID(rv)
}

// Returns a font descriptor with the name and matrix attributes set to the
// given values.
//
// fontName: The value for [NSFontNameAttribute].
//
// matrix: The value for [NSFontMatrixAttribute].
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:matrix:)
func NewFontDescriptorWithNameMatrix(fontName string, matrix foundation.NSAffineTransform) NSFontDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSFontDescriptorClass().class), objc.Sel("fontDescriptorWithName:matrix:"), objc.String(fontName), matrix)
	return NSFontDescriptorFromID(rv)
}

// Returns a font descriptor with the name and size attributes set to the
// given values.
//
// fontName: The value for [NSFontNameAttribute].
//
// size: The value for [NSFontSizeAttribute].
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:size:)
func NewFontDescriptorWithNameSize(fontName string, size float64) NSFontDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSFontDescriptorClass().class), objc.Sel("fontDescriptorWithName:size:"), objc.String(fontName), size)
	return NSFontDescriptorFromID(rv)
}

// Initializes and returns a new font descriptor with the specified
// attributes.
//
// attributes: The attributes for the new font descriptor. If `nil`, the font
// descriptor’s attribute dictionary will be empty.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(fontAttributes:)
func (f NSFontDescriptor) InitWithFontAttributes(attributes foundation.INSDictionary) NSFontDescriptor {
	rv := objc.Send[NSFontDescriptor](f.ID, objc.Sel("initWithFontAttributes:"), attributes)
	return rv
}
// Returns a new font descriptor based on the current object, but with the
// specified attributes taking precedence over the existing ones.
//
// attributes: The replacement attributes for the new font descriptor.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/addingAttributes(_:)
func (f NSFontDescriptor) FontDescriptorByAddingAttributes(attributes foundation.INSDictionary) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorByAddingAttributes:"), attributes)
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified face.
//
// newFace: The replacement font face.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withFace(_:)
func (f NSFontDescriptor) FontDescriptorWithFace(newFace string) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithFace:"), objc.String(newFace))
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified font family.
//
// newFamily: The replacement font family.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withFamily(_:)
func (f NSFontDescriptor) FontDescriptorWithFamily(newFamily string) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithFamily:"), objc.String(newFamily))
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified font matrix.
//
// matrix: The replacement font matrix.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withMatrix(_:)
func (f NSFontDescriptor) FontDescriptorWithMatrix(matrix foundation.NSAffineTransform) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithMatrix:"), matrix)
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified point size.
//
// newPointSize: The replacement point size.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withSize(_:)
func (f NSFontDescriptor) FontDescriptorWithSize(newPointSize float64) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithSize:"), newPointSize)
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified symbolic traits taking precedence over the existing ones.
//
// symbolicTraits: The replacement symbolic traits.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withSymbolicTraits(_:)
func (f NSFontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits NSFontDescriptorSymbolicTraits) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithSymbolicTraits:"), symbolicTraits)
	return NSFontDescriptorFromID(rv)
}
// Returns a new font descriptor based on the current object, but with the
// specified design style.
//
// design: The replacement design style for the font. For a list of possible values,
// see `UIFontDescriptor.SystemDesign`.
//
// # Return Value
// 
// The new font descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withDesign(_:)
func (f NSFontDescriptor) FontDescriptorWithDesign(design NSFontDescriptorSystemDesign) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontDescriptorWithDesign:"), objc.String(string(design)))
	return NSFontDescriptorFromID(rv)
}
// Returns all the fonts available on the system whose specified attributes
// match those of the receiver.
//
// mandatoryKeys: Keys that must be identical to be matched. Can be `nil`.
//
// # Return Value
// 
// The matching font descriptors. If the attribute value specified does not
// exist in the input dictionary or if there is no font that matches the given
// mandatory key values, an empty array is returned.
//
// # Discussion
// 
// For example, suppose there are two versions of a given font installed that
// differ in the number of glyphs covered (the new version has more glyphs).
// If you explicitly specify [name] as the only mandatory key, then a font
// descriptor that specifies a font name and character set by default matches
// both versions, since the character set attribute is not used for matching.
// If you specify that font name and character set keys are mandatory, the
// returned array contains only the font that matches both keys.
//
// [name]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName/name
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matchingFontDescriptors(withMandatoryKeys:)
func (f NSFontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.INSSet) []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("matchingFontDescriptorsWithMandatoryKeys:"), mandatoryKeys)
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
// Returns a normalized font descriptor whose specified attributes match those
// of the receiver.
//
// mandatoryKeys: Keys that must be identical to be matched. Can be `nil`.
//
// # Return Value
// 
// The matching font descriptor. If there is no font that matches the given
// mandatory key values, returns `nil`.
//
// # Discussion
// 
// If more than one font matches the [[NSFontNameAttribute],
// [NSFontFamilyAttribute], [NSFontVisibleNameAttribute],
// [NSFontFaceAttribute]] attributes, the list of font descriptors is filtered
// by the other mandatory keys, if any, and the top result that is returned is
// the same as the first element returned from
// [MatchingFontDescriptorsWithMandatoryKeys].
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matchingFontDescriptor(withMandatoryKeys:)
func (f NSFontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.INSSet) INSFontDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("matchingFontDescriptorWithMandatoryKeys:"), mandatoryKeys)
	return NSFontDescriptorFromID(rv)
}
// Returns the font attribute specified by the given key.
//
// attribute: The font attribute key.
//
// # Return Value
// 
// The font attribute corresponding to `anAttribute`. For valid values of
// `anAttribute`, see `Font Attributes`.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/object(forKey:)
func (f NSFontDescriptor) ObjectForKey(attribute NSFontDescriptorAttributeName) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("objectForKey:"), objc.String(string(attribute)))
	return objectivec.Object{ID: rv}
}
func (f NSFontDescriptor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a font descriptor that contains the text style.
//
// style: The text style for which to return a font descriptor. See [NSFontTextStyle]
// for available values.
//
// options: A dictionary you use to further configure the returned font descriptor. See
// [NSFontTextStyleOptionKey] for a list of valid keys. Pass an empty
// dictionary to use the default configuration.
//
// # Return Value
// 
// The font descriptor that contains the text style.
//
// # Discussion
// 
// The font descriptor contains a dictionary of attributes that you use to
// create an [NSFont] object. See [NSFontDescriptor] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/preferredFontDescriptor(forTextStyle:options:)
func (_NSFontDescriptorClass NSFontDescriptorClass) PreferredFontDescriptorForTextStyleOptions(style NSFontTextStyle, options foundation.INSDictionary) NSFontDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSFontDescriptorClass.class), objc.Sel("preferredFontDescriptorForTextStyle:options:"), objc.String(string(style)), options)
	return NSFontDescriptorFromID(rv)
}

// The receiver’s dictionary of attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontAttributes
func (f NSFontDescriptor) FontAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The current transform matrix of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matrix
func (f NSFontDescriptor) Matrix() foundation.NSAffineTransform {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("matrix"))
	return foundation.NSAffineTransformFromID(objc.ID(rv))
}
// The point size of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/pointSize
func (f NSFontDescriptor) PointSize() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("pointSize"))
	return rv
}
// The PostScript name of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/postscriptName
func (f NSFontDescriptor) PostscriptName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("postscriptName"))
	return foundation.NSStringFromID(rv).String()
}
// Constant you use to access
//
// See: https://developer.apple.com/documentation/appkit/nsfontfamilyclassmask
func (f NSFontDescriptor) NSFontFamilyClassMask() uint32 {
	rv := objc.Send[uint32](f.ID, objc.Sel("NSFontFamilyClassMask"))
	return rv
}
func (f NSFontDescriptor) SetNSFontFamilyClassMask(value uint32) {
	objc.Send[struct{}](f.ID, objc.Sel("setNSFontFamilyClassMask:"), value)
}
// A bit mask that describes the traits of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/symbolicTraits-swift.property
func (f NSFontDescriptor) SymbolicTraits() NSFontDescriptorSymbolicTraits {
	rv := objc.Send[NSFontDescriptorSymbolicTraits](f.ID, objc.Sel("symbolicTraits"))
	return NSFontDescriptorSymbolicTraits(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/requiresFontAssetRequest
func (f NSFontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("requiresFontAssetRequest"))
	return rv
}

