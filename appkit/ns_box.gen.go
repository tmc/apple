// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBox] class.
var (
	_NSBoxClass     NSBoxClass
	_NSBoxClassOnce sync.Once
)

func getNSBoxClass() NSBoxClass {
	_NSBoxClassOnce.Do(func() {
		_NSBoxClass = NSBoxClass{class: objc.GetClass("NSBox")}
	})
	return _NSBoxClass
}

// GetNSBoxClass returns the class object for NSBox.
func GetNSBoxClass() NSBoxClass {
	return getNSBoxClass()
}

type NSBoxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBoxClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBoxClass) Alloc() NSBox {
	rv := objc.Send[NSBox](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A stylized rectangular box with an optional title.
//
// # Overview
// 
// Use box objects to visually group the contents of your window. For example,
// you might use boxes to group related views. Use an [NSBox] object to
// configure the appearance of the box.
// 
// # Subclassing Notes
// 
// An [NSBox] object is a view that draws a line around its rectangular bounds
// and that displays a title on or near the line (or might display neither
// line nor title). You can adjust the style of the line (bezel, grooved, or
// plain) as well as the placement and font of the title. An [NSBox] also has
// a content view to which other views can be added; it thus offers a way for
// an application to group related views. You could create a custom subclass
// of [NSBox] that alters or augments its appearance or that modifies its
// grouping behavior. For example, you might add color to the lines or
// background, add a new line style, or have the views in the group
// automatically snap to an invisible grid when added.
// 
// # Methods to Override
// 
// You must override the [DrawRect] method (inherited from [NSView]) if you
// want to customize the appearance of your [NSBox] objects. Depending on the
// visual effect you’re trying to achieve, you may have to invoke
// `super`‘s implementation first. For example, if you are compositing a
// small image in a corner of the box, you would invoke the superclass
// implementation first. If you’re adding a new style of line, you would
// provide a way to store a request for this line type (such as a boolean
// instance variable and related accessor methods). Then, in [DrawRect], if a
// request for this line type exists, you would draw the entire view yourself
// (that is, without calling `super`). Otherwise, you would invoke the
// superclass implementation.
// 
// If you wish to change grouping behavior or other behavioral characteristics
// of the [NSBox] class, consider overriding [NSBox.ContentView], [NSBox.SizeToFit], or
// [AddSubview] (inherited from [NSView]).
// 
// # Special Considerations
// 
// If you are drawing the custom [NSBox] entirely by yourself, and you want it
// to look exactly like the superclass object (except for your changes), it
// may take some effort and time to get the details right.
//
// # Configuring Boxes
//
//   - [NSBox.BorderRect]: The rectangle in which the receiver’s border is drawn.
//   - [NSBox.BoxType]: The receiver’s box type.
//   - [NSBox.SetBoxType]
//   - [NSBox.Transparent]: A Boolean value that indicates whether the receiver is transparent.
//   - [NSBox.SetTransparent]
//   - [NSBox.Title]: The receiver’s title.
//   - [NSBox.SetTitle]
//   - [NSBox.TitleFont]: The font object used to draw the receiver’s title.
//   - [NSBox.SetTitleFont]
//   - [NSBox.TitlePosition]: A constant representing the title position.
//   - [NSBox.SetTitlePosition]
//   - [NSBox.TitleCell]: The cell used to display the receiver’s title.
//   - [NSBox.TitleRect]: The rectangle in which the receiver’s title is drawn.
//
// # Customizing
//
//   - [NSBox.BorderColor]: The color of the receiver’s border when the receiver is a custom box with a simple line border.
//   - [NSBox.SetBorderColor]
//   - [NSBox.BorderWidth]: The width of the receiver’s border when the receiver is a custom box with a simple line border.
//   - [NSBox.SetBorderWidth]
//   - [NSBox.CornerRadius]: The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//   - [NSBox.SetCornerRadius]
//   - [NSBox.FillColor]: The color of the receiver’s background when the receiver is a custom box with a simple line border.
//   - [NSBox.SetFillColor]
//
// # Managing Content
//
//   - [NSBox.ContentView]: The receiver’s content view.
//   - [NSBox.SetContentView]
//   - [NSBox.ContentViewMargins]: The distances between the border and the content view.
//   - [NSBox.SetContentViewMargins]
//
// # Sizing
//
//   - [NSBox.SetFrameFromContentFrame]: Places the receiver so its content view lies on the specified frame.
//   - [NSBox.SizeToFit]: Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox
type NSBox struct {
	NSView
}

// NSBoxFromID constructs a [NSBox] from an objc.ID.
//
// A stylized rectangular box with an optional title.
func NSBoxFromID(id objc.ID) NSBox {
	return NSBox{NSView: NSViewFromID(id)}
}
// NOTE: NSBox adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBox] class.
//
// # Configuring Boxes
//
//   - [INSBox.BorderRect]: The rectangle in which the receiver’s border is drawn.
//   - [INSBox.BoxType]: The receiver’s box type.
//   - [INSBox.SetBoxType]
//   - [INSBox.Transparent]: A Boolean value that indicates whether the receiver is transparent.
//   - [INSBox.SetTransparent]
//   - [INSBox.Title]: The receiver’s title.
//   - [INSBox.SetTitle]
//   - [INSBox.TitleFont]: The font object used to draw the receiver’s title.
//   - [INSBox.SetTitleFont]
//   - [INSBox.TitlePosition]: A constant representing the title position.
//   - [INSBox.SetTitlePosition]
//   - [INSBox.TitleCell]: The cell used to display the receiver’s title.
//   - [INSBox.TitleRect]: The rectangle in which the receiver’s title is drawn.
//
// # Customizing
//
//   - [INSBox.BorderColor]: The color of the receiver’s border when the receiver is a custom box with a simple line border.
//   - [INSBox.SetBorderColor]
//   - [INSBox.BorderWidth]: The width of the receiver’s border when the receiver is a custom box with a simple line border.
//   - [INSBox.SetBorderWidth]
//   - [INSBox.CornerRadius]: The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//   - [INSBox.SetCornerRadius]
//   - [INSBox.FillColor]: The color of the receiver’s background when the receiver is a custom box with a simple line border.
//   - [INSBox.SetFillColor]
//
// # Managing Content
//
//   - [INSBox.ContentView]: The receiver’s content view.
//   - [INSBox.SetContentView]
//   - [INSBox.ContentViewMargins]: The distances between the border and the content view.
//   - [INSBox.SetContentViewMargins]
//
// # Sizing
//
//   - [INSBox.SetFrameFromContentFrame]: Places the receiver so its content view lies on the specified frame.
//   - [INSBox.SizeToFit]: Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox
type INSBox interface {
	INSView

	// Topic: Configuring Boxes

	// The rectangle in which the receiver’s border is drawn.
	BorderRect() corefoundation.CGRect
	// The receiver’s box type.
	BoxType() NSBoxType
	SetBoxType(value NSBoxType)
	// A Boolean value that indicates whether the receiver is transparent.
	Transparent() bool
	SetTransparent(value bool)
	// The receiver’s title.
	Title() string
	SetTitle(value string)
	// The font object used to draw the receiver’s title.
	TitleFont() NSFont
	SetTitleFont(value NSFont)
	// A constant representing the title position.
	TitlePosition() NSTitlePosition
	SetTitlePosition(value NSTitlePosition)
	// The cell used to display the receiver’s title.
	TitleCell() objectivec.IObject
	// The rectangle in which the receiver’s title is drawn.
	TitleRect() corefoundation.CGRect

	// Topic: Customizing

	// The color of the receiver’s border when the receiver is a custom box with a simple line border.
	BorderColor() INSColor
	SetBorderColor(value INSColor)
	// The width of the receiver’s border when the receiver is a custom box with a simple line border.
	BorderWidth() float64
	SetBorderWidth(value float64)
	// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
	CornerRadius() float64
	SetCornerRadius(value float64)
	// The color of the receiver’s background when the receiver is a custom box with a simple line border.
	FillColor() INSColor
	SetFillColor(value INSColor)

	// Topic: Managing Content

	// The receiver’s content view.
	ContentView() INSView
	SetContentView(value INSView)
	// The distances between the border and the content view.
	ContentViewMargins() corefoundation.CGSize
	SetContentViewMargins(value corefoundation.CGSize)

	// Topic: Sizing

	// Places the receiver so its content view lies on the specified frame.
	SetFrameFromContentFrame(contentFrame corefoundation.CGRect)
	// Resizes and moves the receiver’s content view so it just encloses its subviews.
	SizeToFit()
}

// Init initializes the instance.
func (b NSBox) Init() NSBox {
	rv := objc.Send[NSBox](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBox) Autorelease() NSBox {
	rv := objc.Send[NSBox](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBox creates a new NSBox instance.
func NewNSBox() NSBox {
	class := getNSBoxClass()
	rv := objc.Send[NSBox](objc.ID(class.class), objc.Sel("new"))
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
func NewBoxWithCoder(coder foundation.INSCoder) NSBox {
	instance := getNSBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSBoxFromID(rv)
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
func NewBoxWithFrame(frameRect corefoundation.CGRect) NSBox {
	instance := getNSBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSBoxFromID(rv)
}

// Places the receiver so its content view lies on the specified frame.
//
// contentFrame: The rectangle specifying the frame of the box’s content view, reckoned in
// the coordinate system of the box’s superview. The box is marked for
// redisplay.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b NSBox) SetFrameFromContentFrame(contentFrame corefoundation.CGRect) {
	objc.Send[objc.ID](b.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}
// Resizes and moves the receiver’s content view so it just encloses its
// subviews.
//
// # Discussion
// 
// The receiver is then moved and resized to wrap around the content view. The
// receiver’s width is constrained so its title will be fully displayed.
// 
// You should invoke this method after:
// 
// - Adding a subview (to the content view) - Altering the size or location of
// such a subview - Setting the margins around the content view
// 
// The mechanism by which the content view is moved and resized depends on
// whether the object responds to its own `sizeToFit` message: If it does
// respond, then that message is sent, and the content view is expected to be
// so modified. If the content view doesn’t respond, the box moves and
// resizes the content view itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b NSBox) SizeToFit() {
	objc.Send[objc.ID](b.ID, objc.Sel("sizeToFit"))
}

// The rectangle in which the receiver’s border is drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/borderRect
func (b NSBox) BorderRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("borderRect"))
	return corefoundation.CGRect(rv)
}
// The receiver’s box type.
//
// # Discussion
// 
// A constant describing the type of box. These constants are described in
// [NSBox.BoxType]. By default, the box type of an [NSBox] is [NSBoxPrimary].
//
// [NSBox.BoxType]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b NSBox) BoxType() NSBoxType {
	rv := objc.Send[NSBoxType](b.ID, objc.Sel("boxType"))
	return NSBoxType(rv)
}
func (b NSBox) SetBoxType(value NSBoxType) {
	objc.Send[struct{}](b.ID, objc.Sel("setBoxType:"), value)
}
// A Boolean value that indicates whether the receiver is transparent.
//
// # Discussion
// 
// [true] when the receiver is transparent, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b NSBox) Transparent() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isTransparent"))
	return rv
}
func (b NSBox) SetTransparent(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setTransparent:"), value)
}
// The receiver’s title.
//
// # Discussion
// 
// The title of the [NSBox]. By default, a box’s title is “Title.” If
// the size of the new title is different from that of the old title, the
// content view is resized to absorb the difference.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b NSBox) Title() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSBox) SetTitle(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The font object used to draw the receiver’s title.
//
// # Discussion
// 
// By default, the title is drawn using the small system font (obtained using
// ([SmallSystemFontSize] as the parameter of [SystemFontOfSize], both
// [NSFont] class methods). If the size of the new font is different from that
// of the old font, the content view is resized to absorb the difference.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b NSBox) TitleFont() NSFont {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("titleFont"))
	return NSFontFromID(objc.ID(rv))
}
func (b NSBox) SetTitleFont(value NSFont) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitleFont:"), value)
}
// A constant representing the title position.
//
// # Discussion
// 
// A constant representing the position of the receiver’s title. See
// [NSBox.TitlePosition] for a list of these constants. If the new title
// position changes the size of the box’s border area, the content view is
// resized to absorb the difference, and the box is marked as needing
// redisplay.
//
// [NSBox.TitlePosition]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b NSBox) TitlePosition() NSTitlePosition {
	rv := objc.Send[NSTitlePosition](b.ID, objc.Sel("titlePosition"))
	return NSTitlePosition(rv)
}
func (b NSBox) SetTitlePosition(value NSTitlePosition) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitlePosition:"), value)
}
// The cell used to display the receiver’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/titleCell
func (b NSBox) TitleCell() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("titleCell"))
	return objectivec.Object{ID: rv}
}
// The rectangle in which the receiver’s title is drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/titleRect
func (b NSBox) TitleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("titleRect"))
	return corefoundation.CGRect(rv)
}
// The color of the receiver’s border when the receiver is a custom box with
// a simple line border.
//
// # Discussion
// 
// The receiver’s border color. It must be a custom box—that is, it has a
// type of [NSBox.BoxType.custom]—and it must have a border style of
// [NSBorderType.lineBorder].
// 
// # Special Considerations
// 
// Functional only when the receiver’s box type ([BoxType]) is [NSBoxCustom]
// and its border type ([borderType]) is [NSLineBorder].
//
// [NSBorderType.lineBorder]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
// [NSBox.BoxType.custom]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/custom
// [borderType]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b NSBox) BorderColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("borderColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSBox) SetBorderColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setBorderColor:"), value)
}
// The width of the receiver’s border when the receiver is a custom box with
// a simple line border.
//
// # Discussion
// 
// The receiver’s border width. It must be a custom box—that is, it has a
// type of [NSBox.BoxType.custom]—and it must have a border style of
// [NSBorderType.lineBorder].
// 
// # Special Considerations
// 
// Functional only when the receiver’s box type ([BoxType]) is [NSBoxCustom]
// and its border type ([borderType]) is [NSLineBorder].
//
// [NSBorderType.lineBorder]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
// [NSBox.BoxType.custom]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/custom
// [borderType]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b NSBox) BorderWidth() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("borderWidth"))
	return rv
}
func (b NSBox) SetBorderWidth(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setBorderWidth:"), value)
}
// The radius of the receiver’s corners when the receiver is a custom box
// with a simple line border.
//
// # Discussion
// 
// The receiver’s corner radius. It must be a custom box—that is, it has a
// type of [NSBox.BoxType.custom]—and it must have a border style of
// [NSBorderType.lineBorder].
// 
// # Special Considerations
// 
// Functional only when the receiver’s box type ([BoxType]) is [NSBoxCustom]
// and its border type ([borderType]) is [NSLineBorder].
//
// [NSBorderType.lineBorder]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
// [NSBox.BoxType.custom]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/custom
// [borderType]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b NSBox) CornerRadius() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("cornerRadius"))
	return rv
}
func (b NSBox) SetCornerRadius(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setCornerRadius:"), value)
}
// The color of the receiver’s background when the receiver is a custom box
// with a simple line border.
//
// # Discussion
// 
// The receiver’s fill color. It must be a custom box—that is, it has a
// type of [NSBox.BoxType.custom]—and it must have a border style of
// [NSBorderType.lineBorder].
// 
// # Special Considerations
// 
// Functional only when the receiver’s box type ([BoxType]) is [NSBoxCustom]
// and its border type ([borderType]) is [NSLineBorder].
//
// [NSBorderType.lineBorder]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
// [NSBox.BoxType.custom]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/custom
// [borderType]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b NSBox) FillColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("fillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSBox) SetFillColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setFillColor:"), value)
}
// The receiver’s content view.
//
// # Discussion
// 
// The content view of the [NSBox] object. The content view is created
// automatically when the box is created and resized as the box is resized
// (you should never send frame-altering messages directly to a box’s
// content view). You can replace it with an [NSView] of your own.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b NSBox) ContentView() INSView {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (b NSBox) SetContentView(value INSView) {
	objc.Send[struct{}](b.ID, objc.Sel("setContentView:"), value)
}
// The distances between the border and the content view.
//
// # Discussion
// 
// The width (the horizontal distance between the innermost edge of the border
// and the content view) and height (the vertical distance between the
// innermost edge of the border and the content view) describing the distance
// between the border and the content view. By default, these are both 5.0 in
// the box’s coordinate system.
// 
// Unlike changing a box’s other attributes, such as its title position or
// border type, changing the offsets doesn’t automatically resize the
// content view. In general, you should send a [SizeToFit] message to the box
// after changing the size of its offsets. This message causes the content
// view to remain unchanged while the box is sized to fit around it.
//
// See: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b NSBox) ContentViewMargins() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](b.ID, objc.Sel("contentViewMargins"))
	return corefoundation.CGSize(rv)
}
func (b NSBox) SetContentViewMargins(value corefoundation.CGSize) {
	objc.Send[struct{}](b.ID, objc.Sel("setContentViewMargins:"), value)
}

