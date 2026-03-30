// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSShadow] class.
var (
	_NSShadowClass     NSShadowClass
	_NSShadowClassOnce sync.Once
)

func getNSShadowClass() NSShadowClass {
	_NSShadowClassOnce.Do(func() {
		_NSShadowClass = NSShadowClass{class: objc.GetClass("NSShadow")}
	})
	return _NSShadowClass
}

// GetNSShadowClass returns the class object for NSShadow.
func GetNSShadowClass() NSShadowClass {
	return getNSShadowClass()
}

type NSShadowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSShadowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSShadowClass) Alloc() NSShadow {
	rv := objc.Send[NSShadow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to specify attributes to create and style a drop shadow
// during drawing operations.
//
// # Overview
//
// When you create shadows, the system draws them in the default user
// coordinate space, where coordinates are independent from the pixel values
// of any particular device. Rotations, translations, and other
// transformations of the current transformation matrix (CTM) don’t affect
// the shadow or the apparent position of the shadow’s light source.
//
// A shadow has two positional parameters: an x-offset and a y-offset. Express
// these values with a single size data type ([CGSize] in iOS, [NSSize] in
// macOS), using the units of the default user coordinate space. Positive
// values for these offsets extend down and to the right from the user’s
// perspective.
//
// In addition to its positional parameters, a shadow also contains a blur
// radius, which specifies how much the system blurs a drawn object’s image
// mask before compositing the image onto the destination. A value of `0`
// produces no blur. Larger values produce an increasingly large blurred
// shadow.
//
// You can use an [NSShadow] object in one of two ways. First, you can set it,
// like a color or a font, where [NSShadow] attributes apply to everything you
// draw until you apply another shadow or restore a previous graphics state.
// Second, you can use an [NSShadow] instance as the value for the [Shadow]
// text attribute, so the system applies the shadow to the glyphs
// corresponding to the characters bearing this attribute.
//
// # Managing a shadow
//
//   - [NSShadow.ShadowOffset]: The shadow’s relative position, which you specify with horizontal and vertical offset values.
//   - [NSShadow.SetShadowOffset]
//   - [NSShadow.ShadowBlurRadius]: The blur radius of the shadow.
//   - [NSShadow.SetShadowBlurRadius]
//   - [NSShadow.ShadowColor]: The color of the shadow.
//   - [NSShadow.SetShadowColor]
//
// # Setting a shadow
//
//   - [NSShadow.Set]: Sets the shadow of subsequent drawing operations to the current shadow.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow
//
// [CGSize]: https://developer.apple.com/documentation/CoreFoundation/CGSize
// [NSSize]: https://developer.apple.com/documentation/Foundation/NSSize
type NSShadow struct {
	objectivec.Object
}

// NSShadowFromID constructs a [NSShadow] from an objc.ID.
//
// An object you use to specify attributes to create and style a drop shadow
// during drawing operations.
func NSShadowFromID(id objc.ID) NSShadow {
	return NSShadow{objectivec.Object{ID: id}}
}

// NOTE: NSShadow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSShadow] class.
//
// # Managing a shadow
//
//   - [INSShadow.ShadowOffset]: The shadow’s relative position, which you specify with horizontal and vertical offset values.
//   - [INSShadow.SetShadowOffset]
//   - [INSShadow.ShadowBlurRadius]: The blur radius of the shadow.
//   - [INSShadow.SetShadowBlurRadius]
//   - [INSShadow.ShadowColor]: The color of the shadow.
//   - [INSShadow.SetShadowColor]
//
// # Setting a shadow
//
//   - [INSShadow.Set]: Sets the shadow of subsequent drawing operations to the current shadow.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow
type INSShadow interface {
	objectivec.IObject

	// Topic: Managing a shadow

	// The shadow’s relative position, which you specify with horizontal and vertical offset values.
	ShadowOffset() corefoundation.CGSize
	SetShadowOffset(value corefoundation.CGSize)
	// The blur radius of the shadow.
	ShadowBlurRadius() float64
	SetShadowBlurRadius(value float64)
	// The color of the shadow.
	ShadowColor() INSColor
	SetShadowColor(value INSColor)

	// Topic: Setting a shadow

	// Sets the shadow of subsequent drawing operations to the current shadow.
	Set()

	// The shadow of the text.
	Shadow() foundation.NSString
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSShadow) Init() NSShadow {
	rv := objc.Send[NSShadow](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSShadow) Autorelease() NSShadow {
	rv := objc.Send[NSShadow](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSShadow creates a new NSShadow instance.
func NewNSShadow() NSShadow {
	class := getNSShadowClass()
	rv := objc.Send[NSShadow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the shadow of subsequent drawing operations to the current shadow.
//
// # Discussion
//
// The shadow attributes of the receiver are used until another shadow is set
// or until the graphics state is restored.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow/set()
func (s NSShadow) Set() {
	objc.Send[objc.ID](s.ID, objc.Sel("set"))
}
func (s NSShadow) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The shadow’s relative position, which you specify with horizontal and
// vertical offset values.
//
// # Discussion
//
// This property contains the horizontal and vertical offset values that you
// specify using the `width` and `height` fields of the [CGSize] or [NSSize]
// data type. These offsets use the default user coordinate space and are not
// affected by custom transformations. Positive offset values extend down and
// to the right from the user’s perspective.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow/shadowOffset
func (s NSShadow) ShadowOffset() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("shadowOffset"))
	return corefoundation.CGSize(rv)
}
func (s NSShadow) SetShadowOffset(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setShadowOffset:"), value)
}

// The blur radius of the shadow.
//
// # Discussion
//
// This property contains the shadow’s blur radius, as measured in the
// default user coordinate space. A value of `0` produces no blur, while
// larger values produce an increasingly large blurred shadow. This value must
// not be negative. The default value is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow/shadowBlurRadius
func (s NSShadow) ShadowBlurRadius() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("shadowBlurRadius"))
	return rv
}
func (s NSShadow) SetShadowBlurRadius(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setShadowBlurRadius:"), value)
}

// The color of the shadow.
//
// # Discussion
//
// The default shadow color is black with an alpha of 1/3. If you set this
// property to `nil`, the shadow is not drawn. The color you specify must be
// convertible to an RGBA color and may contain alpha information.
//
// See: https://developer.apple.com/documentation/AppKit/NSShadow/shadowColor
func (s NSShadow) ShadowColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shadowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (s NSShadow) SetShadowColor(value INSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setShadowColor:"), value)
}

// The shadow of the text.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/shadow
func (s NSShadow) Shadow() foundation.NSString {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shadow"))
	return foundation.NSStringFromID(objc.ID(rv))
}
