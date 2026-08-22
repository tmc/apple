// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKLabelNode] class.
var (
	_SKLabelNodeClass     SKLabelNodeClass
	_SKLabelNodeClassOnce sync.Once
)

func getSKLabelNodeClass() SKLabelNodeClass {
	_SKLabelNodeClassOnce.Do(func() {
		_SKLabelNodeClass = SKLabelNodeClass{class: objc.GetClass("SKLabelNode")}
	})
	return _SKLabelNodeClass
}

// GetSKLabelNodeClass returns the class object for SKLabelNode.
func GetSKLabelNodeClass() SKLabelNodeClass {
	return getSKLabelNodeClass()
}

type SKLabelNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKLabelNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKLabelNodeClass) Alloc() SKLabelNode {
	rv := objc.Send[SKLabelNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A graphical element that draws text.
//
// # Overview
//
// [SKLabelNode] allows you to render text in your scene. You can define a
// custom style using properties such as [SKLabelNode.FontName] and
// [SKLabelNode.FontColor], or configure the look of your text with an
// [NSAttributedString].
//
// # Creating a Label
//
//   - [SKLabelNode.InitWithFontNamed]: Initializes a new label object with a specified font.
//
// # Setting a Label’s Text
//
//   - [SKLabelNode.Text]: The string that the label node displays.
//   - [SKLabelNode.SetText]
//   - [SKLabelNode.AttributedText]: The attributed string displayed by the label.
//   - [SKLabelNode.SetAttributedText]
//
// # Specifying a Label’s Font
//
//   - [SKLabelNode.FontColor]: The color of the label.
//   - [SKLabelNode.SetFontColor]
//   - [SKLabelNode.FontName]: The font used for the text in the label.
//   - [SKLabelNode.SetFontName]
//   - [SKLabelNode.FontSize]: The size of the font used in the label.
//   - [SKLabelNode.SetFontSize]
//
// # Controlling a Label’s Alignment
//
//   - [SKLabelNode.VerticalAlignmentMode]: The vertical position of the text within the node.
//   - [SKLabelNode.SetVerticalAlignmentMode]
//   - [SKLabelNode.HorizontalAlignmentMode]: The horizontal position of the text within the node.
//   - [SKLabelNode.SetHorizontalAlignmentMode]
//
// # Defining a Label’s Line-Breaking Behavior
//
//   - [SKLabelNode.PreferredMaxLayoutWidth]: The width, in screen points, after which line-break mode should be applied.
//   - [SKLabelNode.SetPreferredMaxLayoutWidth]
//   - [SKLabelNode.LineBreakMode]: Determines the line-break mode for multiple lines.
//   - [SKLabelNode.SetLineBreakMode]
//   - [SKLabelNode.NumberOfLines]: Determines the number of lines to draw.
//   - [SKLabelNode.SetNumberOfLines]
//
// # Colorizing a Label
//
//   - [SKLabelNode.Color]: An alternative to the font color that can be used for animations.
//   - [SKLabelNode.SetColor]
//   - [SKLabelNode.ColorBlendFactor]: A floating-point value that describes how the color is blended with the font color.
//   - [SKLabelNode.SetColorBlendFactor]
//
// # Configuring Alpha Blending
//
//   - [SKLabelNode.BlendMode]: The blend mode used to draw the label into the parent’s framebuffer.
//   - [SKLabelNode.SetBlendMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
type SKLabelNode struct {
	SKNode
}

// SKLabelNodeFromID constructs a [SKLabelNode] from an objc.ID.
//
// A graphical element that draws text.
func SKLabelNodeFromID(id objc.ID) SKLabelNode {
	return SKLabelNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKLabelNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKLabelNode] class.
//
// # Creating a Label
//
//   - [ISKLabelNode.InitWithFontNamed]: Initializes a new label object with a specified font.
//
// # Setting a Label’s Text
//
//   - [ISKLabelNode.Text]: The string that the label node displays.
//   - [ISKLabelNode.SetText]
//   - [ISKLabelNode.AttributedText]: The attributed string displayed by the label.
//   - [ISKLabelNode.SetAttributedText]
//
// # Specifying a Label’s Font
//
//   - [ISKLabelNode.FontColor]: The color of the label.
//   - [ISKLabelNode.SetFontColor]
//   - [ISKLabelNode.FontName]: The font used for the text in the label.
//   - [ISKLabelNode.SetFontName]
//   - [ISKLabelNode.FontSize]: The size of the font used in the label.
//   - [ISKLabelNode.SetFontSize]
//
// # Controlling a Label’s Alignment
//
//   - [ISKLabelNode.VerticalAlignmentMode]: The vertical position of the text within the node.
//   - [ISKLabelNode.SetVerticalAlignmentMode]
//   - [ISKLabelNode.HorizontalAlignmentMode]: The horizontal position of the text within the node.
//   - [ISKLabelNode.SetHorizontalAlignmentMode]
//
// # Defining a Label’s Line-Breaking Behavior
//
//   - [ISKLabelNode.PreferredMaxLayoutWidth]: The width, in screen points, after which line-break mode should be applied.
//   - [ISKLabelNode.SetPreferredMaxLayoutWidth]
//   - [ISKLabelNode.LineBreakMode]: Determines the line-break mode for multiple lines.
//   - [ISKLabelNode.SetLineBreakMode]
//   - [ISKLabelNode.NumberOfLines]: Determines the number of lines to draw.
//   - [ISKLabelNode.SetNumberOfLines]
//
// # Colorizing a Label
//
//   - [ISKLabelNode.Color]: An alternative to the font color that can be used for animations.
//   - [ISKLabelNode.SetColor]
//   - [ISKLabelNode.ColorBlendFactor]: A floating-point value that describes how the color is blended with the font color.
//   - [ISKLabelNode.SetColorBlendFactor]
//
// # Configuring Alpha Blending
//
//   - [ISKLabelNode.BlendMode]: The blend mode used to draw the label into the parent’s framebuffer.
//   - [ISKLabelNode.SetBlendMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode
type ISKLabelNode interface {
	ISKNode

	// Topic: Creating a Label

	// Initializes a new label object with a specified font.
	InitWithFontNamed(fontName string) SKLabelNode

	// Topic: Setting a Label’s Text

	// The string that the label node displays.
	Text() string
	SetText(value string)
	// The attributed string displayed by the label.
	AttributedText() foundation.NSAttributedString
	SetAttributedText(value foundation.NSAttributedString)

	// Topic: Specifying a Label’s Font

	// The color of the label.
	FontColor() appkit.NSColor
	SetFontColor(value appkit.NSColor)
	// The font used for the text in the label.
	FontName() string
	SetFontName(value string)
	// The size of the font used in the label.
	FontSize() float64
	SetFontSize(value float64)

	// Topic: Controlling a Label’s Alignment

	// The vertical position of the text within the node.
	VerticalAlignmentMode() SKLabelVerticalAlignmentMode
	SetVerticalAlignmentMode(value SKLabelVerticalAlignmentMode)
	// The horizontal position of the text within the node.
	HorizontalAlignmentMode() SKLabelHorizontalAlignmentMode
	SetHorizontalAlignmentMode(value SKLabelHorizontalAlignmentMode)

	// Topic: Defining a Label’s Line-Breaking Behavior

	// The width, in screen points, after which line-break mode should be applied.
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
	// Determines the line-break mode for multiple lines.
	LineBreakMode() appkit.NSLineBreakMode
	SetLineBreakMode(value appkit.NSLineBreakMode)
	// Determines the number of lines to draw.
	NumberOfLines() int
	SetNumberOfLines(value int)

	// Topic: Colorizing a Label

	// An alternative to the font color that can be used for animations.
	Color() appkit.NSColor
	SetColor(value appkit.NSColor)
	// A floating-point value that describes how the color is blended with the font color.
	ColorBlendFactor() float64
	SetColorBlendFactor(value float64)

	// Topic: Configuring Alpha Blending

	// The blend mode used to draw the label into the parent’s framebuffer.
	BlendMode() SKBlendMode
	SetBlendMode(value SKBlendMode)
}

// Init initializes the instance.
func (l SKLabelNode) Init() SKLabelNode {
	rv := objc.Send[SKLabelNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l SKLabelNode) Autorelease() SKLabelNode {
	rv := objc.Send[SKLabelNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKLabelNode creates a new SKLabelNode instance.
func NewSKLabelNode() SKLabelNode {
	class := getSKLabelNodeClass()
	rv := objc.Send[SKLabelNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new label object with an attributed text string.
//
// attributedText: The attributed string from which to initialize the label.
//
// # Return Value
//
// A lable initialized from attributed text.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/init(attributedText:)
func NewLabelNodeWithAttributedText(attributedText foundation.NSAttributedString) SKLabelNode {
	rv := objc.Send[objc.ID](objc.ID(getSKLabelNodeClass().class), objc.Sel("labelNodeWithAttributedText:"), attributedText)
	return SKLabelNodeFromID(rv)
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewLabelNodeWithCoder(aDecoder foundation.INSCoder) SKLabelNode {
	instance := getSKLabelNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKLabelNodeFromID(rv)
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewLabelNodeWithFileNamed(filename string) SKLabelNode {
	rv := objc.Send[objc.ID](objc.ID(getSKLabelNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKLabelNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewLabelNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKLabelNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKLabelNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKLabelNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKLabelNode{}, objc.ErrInitFailed
	}
	return SKLabelNodeFromID(rv), nil
}

// Initializes a new label object with a specified font.
//
// fontName: The name of the font used by the label.
//
// # Return Value
//
// An initialized label object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/init(fontNamed:)
func NewLabelNodeWithFontNamed(fontName string) SKLabelNode {
	instance := getSKLabelNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFontNamed:"), objc.String(fontName))
	return SKLabelNodeFromID(rv)
}

// Initializes a new label object with a text string.
//
// text: The text to use to initialize the label node.
//
// # Return Value
//
// An initialized label object.
//
// # Discussion
//
// The label node’s font is set to Helvetica Neue Ultra Light, 32 point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/init(text:)
func NewLabelNodeWithText(text string) SKLabelNode {
	rv := objc.Send[objc.ID](objc.ID(getSKLabelNodeClass().class), objc.Sel("labelNodeWithText:"), objc.String(text))
	return SKLabelNodeFromID(rv)
}

// Initializes a new label object with a specified font.
//
// fontName: The name of the font used by the label.
//
// # Return Value
//
// An initialized label object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/init(fontNamed:)
func (l SKLabelNode) InitWithFontNamed(fontName string) SKLabelNode {
	rv := objc.Send[SKLabelNode](l.ID, objc.Sel("initWithFontNamed:"), objc.String(fontName))
	return rv
}

// The string that the label node displays.
//
// # Discussion
//
// This property is ignored if attributedText is defined.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/text
func (l SKLabelNode) Text() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
func (l SKLabelNode) SetText(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setText:"), objc.String(value))
}

// The attributed string displayed by the label.
//
// # Discussion
//
// The following properties are ignored if attributedText is defined:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/attributedText
func (l SKLabelNode) AttributedText() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("attributedText"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (l SKLabelNode) SetAttributedText(value foundation.NSAttributedString) {
	objc.Send[struct{}](l.ID, objc.Sel("setAttributedText:"), value)
}

// The color of the label.
//
// # Discussion
//
// The default color is white.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/fontColor
func (l SKLabelNode) FontColor() appkit.NSColor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("fontColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (l SKLabelNode) SetFontColor(value appkit.NSColor) {
	objc.Send[struct{}](l.ID, objc.Sel("setFontColor:"), value)
}

// The font used for the text in the label.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/fontName
func (l SKLabelNode) FontName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("fontName"))
	return foundation.NSStringFromID(rv).String()
}
func (l SKLabelNode) SetFontName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setFontName:"), objc.String(value))
}

// The size of the font used in the label.
//
// # Discussion
//
// The default font size is 32 points.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/fontSize
func (l SKLabelNode) FontSize() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("fontSize"))
	return rv
}
func (l SKLabelNode) SetFontSize(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setFontSize:"), value)
}

// The vertical position of the text within the node.
//
// # Discussion
//
// The possible values for this property are listed in
// [SKLabelVerticalAlignmentMode]. The default value of this property is
// [SKLabelVerticalAlignmentMode.baseline].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/verticalAlignmentMode
//
// [SKLabelVerticalAlignmentMode.baseline]: https://developer.apple.com/documentation/SpriteKit/SKLabelVerticalAlignmentMode/baseline
// [SKLabelVerticalAlignmentMode]: https://developer.apple.com/documentation/SpriteKit/SKLabelVerticalAlignmentMode
func (l SKLabelNode) VerticalAlignmentMode() SKLabelVerticalAlignmentMode {
	rv := objc.Send[SKLabelVerticalAlignmentMode](l.ID, objc.Sel("verticalAlignmentMode"))
	return SKLabelVerticalAlignmentMode(rv)
}
func (l SKLabelNode) SetVerticalAlignmentMode(value SKLabelVerticalAlignmentMode) {
	objc.Send[struct{}](l.ID, objc.Sel("setVerticalAlignmentMode:"), value)
}

// The horizontal position of the text within the node.
//
// # Discussion
//
// The possible values for this property are listed in
// [SKLabelHorizontalAlignmentMode]. The default value of this property is
// [SKLabelHorizontalAlignmentMode.center].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/horizontalAlignmentMode
//
// [SKLabelHorizontalAlignmentMode.center]: https://developer.apple.com/documentation/SpriteKit/SKLabelHorizontalAlignmentMode/center
// [SKLabelHorizontalAlignmentMode]: https://developer.apple.com/documentation/SpriteKit/SKLabelHorizontalAlignmentMode
func (l SKLabelNode) HorizontalAlignmentMode() SKLabelHorizontalAlignmentMode {
	rv := objc.Send[SKLabelHorizontalAlignmentMode](l.ID, objc.Sel("horizontalAlignmentMode"))
	return SKLabelHorizontalAlignmentMode(rv)
}
func (l SKLabelNode) SetHorizontalAlignmentMode(value SKLabelHorizontalAlignmentMode) {
	objc.Send[struct{}](l.ID, objc.Sel("setHorizontalAlignmentMode:"), value)
}

// The width, in screen points, after which line-break mode should be applied.
//
// # Discussion
//
// The default value is 0, which means that line break mode will not be
// applied.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/preferredMaxLayoutWidth
func (l SKLabelNode) PreferredMaxLayoutWidth() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}
func (l SKLabelNode) SetPreferredMaxLayoutWidth(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}

// Determines the line-break mode for multiple lines.
//
// # Discussion
//
// The default value is NSLineBreakByTruncatingTail.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/lineBreakMode
func (l SKLabelNode) LineBreakMode() appkit.NSLineBreakMode {
	rv := objc.Send[appkit.NSLineBreakMode](l.ID, objc.Sel("lineBreakMode"))
	return appkit.NSLineBreakMode(rv)
}
func (l SKLabelNode) SetLineBreakMode(value appkit.NSLineBreakMode) {
	objc.Send[struct{}](l.ID, objc.Sel("setLineBreakMode:"), value)
}

// Determines the number of lines to draw.
//
// # Discussion
//
// The default value is 1 (a single line). A value of 0 in interpreted as an
// unlimited number of lines. If the height of the text reaches the number of
// lines, the text will be truncated using the line break mode.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/numberOfLines
func (l SKLabelNode) NumberOfLines() int {
	rv := objc.Send[int](l.ID, objc.Sel("numberOfLines"))
	return rv
}
func (l SKLabelNode) SetNumberOfLines(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberOfLines:"), value)
}

// An alternative to the font color that can be used for animations.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/color
func (l SKLabelNode) Color() appkit.NSColor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("color"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (l SKLabelNode) SetColor(value appkit.NSColor) {
	objc.Send[struct{}](l.ID, objc.Sel("setColor:"), value)
}

// A floating-point value that describes how the color is blended with the
// font color.
//
// # Discussion
//
// The value must be a number between `0.0` and `1.0`, inclusive. The default
// value (`0.0`) indicates that the color property is ignored and that the
// label’s font color should be used unmodified. For values greater than
// `0.0`, the font color is blended with the blend color, with the maximum
// value of 1.0 determining that the font color is 100% of the blend color.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/colorBlendFactor
func (l SKLabelNode) ColorBlendFactor() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("colorBlendFactor"))
	return rv
}
func (l SKLabelNode) SetColorBlendFactor(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setColorBlendFactor:"), value)
}

// The blend mode used to draw the label into the parent’s framebuffer.
//
// # Discussion
//
// The default value is [SKBlendMode.alpha].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKLabelNode/blendMode
//
// [SKBlendMode.alpha]: https://developer.apple.com/documentation/SpriteKit/SKBlendMode/alpha
func (l SKLabelNode) BlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](l.ID, objc.Sel("blendMode"))
	return SKBlendMode(rv)
}
func (l SKLabelNode) SetBlendMode(value SKBlendMode) {
	objc.Send[struct{}](l.ID, objc.Sel("setBlendMode:"), value)
}
