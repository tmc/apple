// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSVisualEffectView] class.
var (
	_NSVisualEffectViewClass     NSVisualEffectViewClass
	_NSVisualEffectViewClassOnce sync.Once
)

func getNSVisualEffectViewClass() NSVisualEffectViewClass {
	_NSVisualEffectViewClassOnce.Do(func() {
		_NSVisualEffectViewClass = NSVisualEffectViewClass{class: objc.GetClass("NSVisualEffectView")}
	})
	return _NSVisualEffectViewClass
}

// GetNSVisualEffectViewClass returns the class object for NSVisualEffectView.
func GetNSVisualEffectViewClass() NSVisualEffectViewClass {
	return getNSVisualEffectViewClass()
}

type NSVisualEffectViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSVisualEffectViewClass) Alloc() NSVisualEffectView {
	rv := objc.Send[NSVisualEffectView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A view that adds translucency and vibrancy effects to the views in your
// interface.
//
// # Overview
// 
// Use visual effect views primarily as background views for your app’s
// content. A visual effect view makes your foreground content more prominent
// by employing the following effects:
// 
// - and the blurring of background content adds depth to your interface. - is
// a subtle blending of foreground and background colors to increase the
// contrast and make the foreground content stand out visually.
// 
// The material and blending mode you assign determines the exact appearance
// of the visual effect. Not all materials support transparency, and materials
// apply vibrancy in different ways. The appearance and behavior of materials
// can also change based on system settings, so always pick a material based
// on its intended use. For example, use the
// [NSVisualEffectView.Material.sidebar] material when your view serves as the
// background of your window’s sidebar. Don’t select materials based on
// the apparent colors they impart on your interface.
// 
// AppKit creates visual effect views automatically for window titlebars,
// popovers, and source list table views. You don’t need to add visual
// effect views to those elements of your interface.
// 
// # Choosing a Translucency Effect for Your View
// 
// For visual effect views you create yourself, use the [NSVisualEffectView.BlendingMode]
// property to specify how and where you want the translucency applied.
// 
// - uses the content behind the window as the background for your visual
// effect view. Behind-window blending makes your entire window stand out
// above other windows and apps on the desktop. Sheets and popovers use
// behind-window blending. - uses the window’s content as the background for
// your visual effect view. Typically, you use in-window blending with
// scrolling content, so that the scrolled content remains partially visible
// under other parts of your window chrome. Toolbars always use in-window
// blending.
// 
// [media-3198506]
// 
// # Enabling Vibrancy for Foreground Content
// 
// The presence of a visual effect view in your view hierarchy does not
// automatically add vibrancy to your content. For custom views, you must
// explicitly enable vibrancy by overriding the [allowsVibrancy] property and
// returning [true].
// 
// It is recommended that you enable vibrancy only in the leaf views of your
// view hierarchy. Subviews inherit the vibrancy of their parent. Once enabled
// in a parent view, a subview cannot turn off vibrancy. As a result, enabling
// vibrancy in a parent view can lead to subviews that look incorrect if they
// are not designed to take advantage of the vibrancy effect.
// 
// Vibrancy works best when your custom views contain grayscale content.
// Combining a grayscale foreground with a color background works well,
// because AppKit improves the contrast while only subtly changing the
// foreground hue. The same isn’t always true when blending two different
// color values. Dramatically different foreground and background hues can
// cancel each other out or result in colors that don’t match your original
// designs.
// 
// Instead of defining custom grayscale color assets, consider using the
// built-in colors [labelColor], [secondaryLabelColor], [tertiaryLabelColor],
// and [quaternaryLabelColor]. While typically used with text, these colors
// are applicable with any app content. The built-in colors represent varying
// levels of contrast for your content, with [labelColor] offering the most
// contrast, and [quaternaryLabelColor] offering the least contrast.
// 
// # Subclassing Notes
// 
// If you subclass [NSVisualEffectView]:
// 
// - Always call `super` if you override [NSVisualEffectView.ViewDidMoveToWindow] or
// [NSVisualEffectView.ViewWillMoveToWindow]. - Do not override [DrawRect] or [UpdateLayer].
//
// [NSVisualEffectView.Material.sidebar]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/sidebar
// [allowsVibrancy]: https://developer.apple.com/documentation/AppKit/NSView/allowsVibrancy
// [labelColor]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
// [quaternaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
// [secondaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
// [tertiaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Specifying the Background Material
//
//   - [NSVisualEffectView.Material]: The material shown by the visual effect view.
//   - [NSVisualEffectView.SetMaterial]
//
// # Specifying the Effect Appearance
//
//   - [NSVisualEffectView.BlendingMode]: A value indicating how the view’s contents blend with the surrounding content.
//   - [NSVisualEffectView.SetBlendingMode]
//   - [NSVisualEffectView.Emphasized]: A Boolean value indicating whether to emphasize the look of the material.
//   - [NSVisualEffectView.SetEmphasized]
//   - [NSVisualEffectView.InteriorBackgroundStyle]: The view’s interior background style.
//
// # Masking the Visual Effect
//
//   - [NSVisualEffectView.MaskImage]: An image whose alpha channel masks the visual effect view’s material.
//   - [NSVisualEffectView.SetMaskImage]
//
// # Enabling or Disabling the Effect
//
//   - [NSVisualEffectView.State]: A value that indicates whether a view has a visual effect applied.
//   - [NSVisualEffectView.SetState]
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView
type NSVisualEffectView struct {
	NSView
}

// NSVisualEffectViewFromID constructs a [NSVisualEffectView] from an objc.ID.
//
// A view that adds translucency and vibrancy effects to the views in your
// interface.
func NSVisualEffectViewFromID(id objc.ID) NSVisualEffectView {
	return NSVisualEffectView{NSView: NSViewFromID(id)}
}
// NOTE: NSVisualEffectView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSVisualEffectView] class.
//
// # Specifying the Background Material
//
//   - [INSVisualEffectView.Material]: The material shown by the visual effect view.
//   - [INSVisualEffectView.SetMaterial]
//
// # Specifying the Effect Appearance
//
//   - [INSVisualEffectView.BlendingMode]: A value indicating how the view’s contents blend with the surrounding content.
//   - [INSVisualEffectView.SetBlendingMode]
//   - [INSVisualEffectView.Emphasized]: A Boolean value indicating whether to emphasize the look of the material.
//   - [INSVisualEffectView.SetEmphasized]
//   - [INSVisualEffectView.InteriorBackgroundStyle]: The view’s interior background style.
//
// # Masking the Visual Effect
//
//   - [INSVisualEffectView.MaskImage]: An image whose alpha channel masks the visual effect view’s material.
//   - [INSVisualEffectView.SetMaskImage]
//
// # Enabling or Disabling the Effect
//
//   - [INSVisualEffectView.State]: A value that indicates whether a view has a visual effect applied.
//   - [INSVisualEffectView.SetState]
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView
type INSVisualEffectView interface {
	INSView

	// Topic: Specifying the Background Material

	// The material shown by the visual effect view.
	Material() NSVisualEffectMaterial
	SetMaterial(value NSVisualEffectMaterial)

	// Topic: Specifying the Effect Appearance

	// A value indicating how the view’s contents blend with the surrounding content.
	BlendingMode() NSVisualEffectBlendingMode
	SetBlendingMode(value NSVisualEffectBlendingMode)
	// A Boolean value indicating whether to emphasize the look of the material.
	Emphasized() bool
	SetEmphasized(value bool)
	// The view’s interior background style.
	InteriorBackgroundStyle() NSBackgroundStyle

	// Topic: Masking the Visual Effect

	// An image whose alpha channel masks the visual effect view’s material.
	MaskImage() INSImage
	SetMaskImage(value INSImage)

	// Topic: Enabling or Disabling the Effect

	// A value that indicates whether a view has a visual effect applied.
	State() NSVisualEffectState
	SetState(value NSVisualEffectState)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (v NSVisualEffectView) Init() NSVisualEffectView {
	rv := objc.Send[NSVisualEffectView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSVisualEffectView) Autorelease() NSVisualEffectView {
	rv := objc.Send[NSVisualEffectView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSVisualEffectView creates a new NSVisualEffectView instance.
func NewNSVisualEffectView() NSVisualEffectView {
	class := getNSVisualEffectViewClass()
	rv := objc.Send[NSVisualEffectView](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewVisualEffectViewWithCoder(coder foundation.INSCoder) NSVisualEffectView {
	instance := getNSVisualEffectViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSVisualEffectViewFromID(rv)
}


// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewVisualEffectViewWithFrame(frameRect corefoundation.CGRect) NSVisualEffectView {
	instance := getNSVisualEffectViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSVisualEffectViewFromID(rv)
}






func (v NSVisualEffectView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The material shown by the visual effect view.
//
// # Discussion
// 
// The default value is [NSVisualEffectView.Material.appearanceBased]; the
// material updates to the correct material based on the appearance set on
// this view.
//
// [NSVisualEffectView.Material.appearanceBased]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/appearanceBased
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v NSVisualEffectView) Material() NSVisualEffectMaterial {
	rv := objc.Send[NSVisualEffectMaterial](v.ID, objc.Sel("material"))
	return NSVisualEffectMaterial(rv)
}
func (v NSVisualEffectView) SetMaterial(value NSVisualEffectMaterial) {
	objc.Send[struct{}](v.ID, objc.Sel("setMaterial:"), value)
}



// A value indicating how the view’s contents blend with the surrounding
// content.
//
// # Discussion
// 
// When the value of this property is
// [NSVisualEffectView.BlendingMode.behindWindow] (the default), the visual
// effect view blurs the content behind the window. When the value is
// [NSVisualEffectView.BlendingMode.withinWindow], it blurs the content behind
// the view of the current window.
// 
// If the visual effect view’s material is
// [NSVisualEffectView.Material.titlebar], set the blending mode to
// [NSVisualEffectView.BlendingMode.withinWindow].
//
// [NSVisualEffectView.BlendingMode.behindWindow]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum/behindWindow
// [NSVisualEffectView.BlendingMode.withinWindow]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum/withinWindow
// [NSVisualEffectView.Material.titlebar]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/titlebar
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v NSVisualEffectView) BlendingMode() NSVisualEffectBlendingMode {
	rv := objc.Send[NSVisualEffectBlendingMode](v.ID, objc.Sel("blendingMode"))
	return NSVisualEffectBlendingMode(rv)
}
func (v NSVisualEffectView) SetBlendingMode(value NSVisualEffectBlendingMode) {
	objc.Send[struct{}](v.ID, objc.Sel("setBlendingMode:"), value)
}



// A Boolean value indicating whether to emphasize the look of the material.
//
// # Discussion
// 
// Some materials change their appearance when they are emphasized. For
// example, the first responder view conveys its status.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v NSVisualEffectView) Emphasized() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isEmphasized"))
	return rv
}
func (v NSVisualEffectView) SetEmphasized(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEmphasized:"), value)
}



// The view’s interior background style.
//
// # Discussion
// 
// The background style may be light or dark, depending on the selected
// material.
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/interiorBackgroundStyle
func (v NSVisualEffectView) InteriorBackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](v.ID, objc.Sel("interiorBackgroundStyle"))
	return NSBackgroundStyle(rv)
}



// An image whose alpha channel masks the visual effect view’s material.
//
// # Discussion
// 
// The default value of this property is `nil`, which is the equivalent of
// allowing all of the visual effect view’s content to show through.
// Assigning an image to this property masks the portions of the visual effect
// view using the image’s alpha channel.
// 
// If the visual effect view is the content view of a window, the mask is
// applied in an appropriate way to the window’s shadow.
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/maskImage
func (v NSVisualEffectView) MaskImage() INSImage {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("maskImage"))
	return NSImageFromID(objc.ID(rv))
}
func (v NSVisualEffectView) SetMaskImage(value INSImage) {
	objc.Send[struct{}](v.ID, objc.Sel("setMaskImage:"), value)
}



// A value that indicates whether a view has a visual effect applied.
//
// # Discussion
// 
// The default value of this property is
// [NSVisualEffectView.State.followsWindowActiveState].
//
// [NSVisualEffectView.State.followsWindowActiveState]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum/followsWindowActiveState
//
// See: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v NSVisualEffectView) State() NSVisualEffectState {
	rv := objc.Send[NSVisualEffectState](v.ID, objc.Sel("state"))
	return NSVisualEffectState(rv)
}
func (v NSVisualEffectView) SetState(value NSVisualEffectState) {
	objc.Send[struct{}](v.ID, objc.Sel("setState:"), value)
}







// The primary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/labelcolor
func (_NSVisualEffectViewClass NSVisualEffectViewClass) LabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("labelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSVisualEffectViewClass NSVisualEffectViewClass) SetLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("setLabelColor:"), value)
}



// The quaternary color to use for text labels and separators.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/quaternarylabelcolor
func (_NSVisualEffectViewClass NSVisualEffectViewClass) QuaternaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("quaternaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSVisualEffectViewClass NSVisualEffectViewClass) SetQuaternaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("setQuaternaryLabelColor:"), value)
}



// The secondary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/secondarylabelcolor
func (_NSVisualEffectViewClass NSVisualEffectViewClass) SecondaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("secondaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSVisualEffectViewClass NSVisualEffectViewClass) SetSecondaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("setSecondaryLabelColor:"), value)
}



// The tertiary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/tertiarylabelcolor
func (_NSVisualEffectViewClass NSVisualEffectViewClass) TertiaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("tertiaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSVisualEffectViewClass NSVisualEffectViewClass) SetTertiaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSVisualEffectViewClass.class), objc.Sel("setTertiaryLabelColor:"), value)
}































