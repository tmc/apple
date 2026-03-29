// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATextLayer] class.
var (
	_CATextLayerClass     CATextLayerClass
	_CATextLayerClassOnce sync.Once
)

func getCATextLayerClass() CATextLayerClass {
	_CATextLayerClassOnce.Do(func() {
		_CATextLayerClass = CATextLayerClass{class: objc.GetClass("CATextLayer")}
	})
	return _CATextLayerClass
}

// GetCATextLayerClass returns the class object for CATextLayer.
func GetCATextLayerClass() CATextLayerClass {
	return getCATextLayerClass()
}

type CATextLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CATextLayerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATextLayerClass) Alloc() CATextLayer {
	rv := objc.Send[CATextLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that provides simple text layout and rendering of plain or
// attributed strings.
//
// # Overview
// 
// The first line is aligned to the top of the layer.
//
// # Text Visual Properties
//
//   - [CATextLayer.Font]: The font used to render the receiver’s text.
//   - [CATextLayer.SetFont]
//   - [CATextLayer.FontSize]: The font size used to render the receiver’s text. Animatable.
//   - [CATextLayer.SetFontSize]
//   - [CATextLayer.ForegroundColor]: The color used to render the receiver’s text. Animatable.
//   - [CATextLayer.SetForegroundColor]
//   - [CATextLayer.AllowsFontSubpixelQuantization]: Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//   - [CATextLayer.SetAllowsFontSubpixelQuantization]
//
// # Text Alignment and Truncation
//
//   - [CATextLayer.Wrapped]: Determines whether the text is wrapped to fit within the receiver’s bounds.
//   - [CATextLayer.SetWrapped]
//   - [CATextLayer.AlignmentMode]: Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//   - [CATextLayer.SetAlignmentMode]
//   - [CATextLayer.TruncationMode]: Determines how the text is truncated to fit within the receiver’s bounds.
//   - [CATextLayer.SetTruncationMode]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer
type CATextLayer struct {
	CALayer
}

// CATextLayerFromID constructs a [CATextLayer] from an objc.ID.
//
// A layer that provides simple text layout and rendering of plain or
// attributed strings.
func CATextLayerFromID(id objc.ID) CATextLayer {
	return CATextLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CATextLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATextLayer] class.
//
// # Text Visual Properties
//
//   - [ICATextLayer.Font]: The font used to render the receiver’s text.
//   - [ICATextLayer.SetFont]
//   - [ICATextLayer.FontSize]: The font size used to render the receiver’s text. Animatable.
//   - [ICATextLayer.SetFontSize]
//   - [ICATextLayer.ForegroundColor]: The color used to render the receiver’s text. Animatable.
//   - [ICATextLayer.SetForegroundColor]
//   - [ICATextLayer.AllowsFontSubpixelQuantization]: Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//   - [ICATextLayer.SetAllowsFontSubpixelQuantization]
//
// # Text Alignment and Truncation
//
//   - [ICATextLayer.Wrapped]: Determines whether the text is wrapped to fit within the receiver’s bounds.
//   - [ICATextLayer.SetWrapped]
//   - [ICATextLayer.AlignmentMode]: Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//   - [ICATextLayer.SetAlignmentMode]
//   - [ICATextLayer.TruncationMode]: Determines how the text is truncated to fit within the receiver’s bounds.
//   - [ICATextLayer.SetTruncationMode]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer
type ICATextLayer interface {
	ICALayer

	// Topic: Text Visual Properties

	// The font used to render the receiver’s text.
	Font() corefoundation.CFTypeRef
	SetFont(value corefoundation.CFTypeRef)
	// The font size used to render the receiver’s text. Animatable.
	FontSize() float64
	SetFontSize(value float64)
	// The color used to render the receiver’s text. Animatable.
	ForegroundColor() coregraphics.CGColorRef
	SetForegroundColor(value coregraphics.CGColorRef)
	// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)

	// Topic: Text Alignment and Truncation

	// Determines whether the text is wrapped to fit within the receiver’s bounds.
	Wrapped() bool
	SetWrapped(value bool)
	// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
	AlignmentMode() CATextLayerAlignmentMode
	SetAlignmentMode(value CATextLayerAlignmentMode)
	// Determines how the text is truncated to fit within the receiver’s bounds.
	TruncationMode() CATextLayerTruncationMode
	SetTruncationMode(value CATextLayerTruncationMode)
}

// Init initializes the instance.
func (t CATextLayer) Init() CATextLayer {
	rv := objc.Send[CATextLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATextLayer) Autorelease() CATextLayer {
	rv := objc.Send[CATextLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATextLayer creates a new CATextLayer instance.
func NewCATextLayer() CATextLayer {
	class := getCATextLayerClass()
	rv := objc.Send[CATextLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewTextLayerWithLayer(layer objectivec.IObject) CATextLayer {
	instance := getCATextLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CATextLayerFromID(rv)
}

// The font used to render the receiver’s text.
//
// # Discussion
// 
// May be either a [CTFont], a [CGFont], an instance of [NSFont] (macOS only),
// or a string naming the font. In iOS, you cannot assign a [UIFont] object to
// this property. Defaults to Helvetica.
// 
// The `font` property is only used when the [String] property is not an
// [NSAttributedString].
//
// [CGFont]: https://developer.apple.com/documentation/CoreGraphics/CGFont
// [CTFont]: https://developer.apple.com/documentation/CoreText/CTFont
// [UIFont]: https://developer.apple.com/documentation/UIKit/UIFont
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/font
func (t CATextLayer) Font() corefoundation.CFTypeRef {
	rv := objc.Send[corefoundation.CFTypeRef](t.ID, objc.Sel("font"))
	return corefoundation.CFTypeRef(rv)
}
func (t CATextLayer) SetFont(value corefoundation.CFTypeRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setFont:"), value)
}
// The font size used to render the receiver’s text. Animatable.
//
// # Discussion
// 
// Defaults to 36.0.
// 
// The `fontSize` property is only used when the [String] property is not an
// [NSAttributedString].
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/fontSize
func (t CATextLayer) FontSize() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("fontSize"))
	return rv
}
func (t CATextLayer) SetFontSize(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setFontSize:"), value)
}
// The color used to render the receiver’s text. Animatable.
//
// # Discussion
// 
// Defaults to opaque white.
// 
// The `foregroundColor` property is only used when the [String] property is
// not an [NSAttributedString].
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/foregroundColor
func (t CATextLayer) ForegroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](t.ID, objc.Sel("foregroundColor"))
	return coregraphics.CGColorRef(rv)
}
func (t CATextLayer) SetForegroundColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setForegroundColor:"), value)
}
// Determines whether to allow subpixel quantization for the graphics context
// used for text rendering.
//
// # Discussion
// 
// When enabled, the graphics context used for text rendering may quantize the
// subpixel positions of glyphs.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/allowsFontSubpixelQuantization
func (t CATextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsFontSubpixelQuantization"))
	return rv
}
func (t CATextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsFontSubpixelQuantization:"), value)
}
// Determines whether the text is wrapped to fit within the receiver’s
// bounds.
//
// # Discussion
// 
// Defaults to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/isWrapped
func (t CATextLayer) Wrapped() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isWrapped"))
	return rv
}
func (t CATextLayer) SetWrapped(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setWrapped:"), value)
}
// Determines how individual lines of text are horizontally aligned within the
// receiver’s bounds.
//
// # Discussion
// 
// The possible values are described in [Horizontal alignment modes]. Defaults
// to [natural].
//
// [Horizontal alignment modes]: https://developer.apple.com/documentation/QuartzCore/horizontal-alignment-modes
// [natural]: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/natural
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/alignmentMode
func (t CATextLayer) AlignmentMode() CATextLayerAlignmentMode {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("alignmentMode"))
	return CATextLayerAlignmentMode(foundation.NSStringFromID(rv).String())
}
func (t CATextLayer) SetAlignmentMode(value CATextLayerAlignmentMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setAlignmentMode:"), objc.String(string(value)))
}
// Determines how the text is truncated to fit within the receiver’s bounds.
//
// # Discussion
// 
// The possible values are described in [Truncation modes]. Defaults to
// [none].
//
// [Truncation modes]: https://developer.apple.com/documentation/QuartzCore/truncation-modes
// [none]: https://developer.apple.com/documentation/QuartzCore/CATextLayerTruncationMode/none
//
// See: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t CATextLayer) TruncationMode() CATextLayerTruncationMode {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("truncationMode"))
	return CATextLayerTruncationMode(foundation.NSStringFromID(rv).String())
}
func (t CATextLayer) SetTruncationMode(value CATextLayerTruncationMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setTruncationMode:"), objc.String(string(value)))
}

