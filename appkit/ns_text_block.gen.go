// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextBlock] class.
var (
	_NSTextBlockClass     NSTextBlockClass
	_NSTextBlockClassOnce sync.Once
)

func getNSTextBlockClass() NSTextBlockClass {
	_NSTextBlockClassOnce.Do(func() {
		_NSTextBlockClass = NSTextBlockClass{class: objc.GetClass("NSTextBlock")}
	})
	return _NSTextBlockClass
}

// GetNSTextBlockClass returns the class object for NSTextBlock.
func GetNSTextBlockClass() NSTextBlockClass {
	return getNSTextBlockClass()
}

type NSTextBlockClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextBlockClass) Alloc() NSTextBlock {
	rv := objc.Send[NSTextBlock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A block of text laid out in a subregion of the text container.
//
// # Overview
// 
// A text block appears as an attribute of a paragraph, and as part of the
// paragraph style. The most important subclass of [NSTextBlock] is
// [NSTextTableBlock], which represents a block of text that appears as a cell
// in a table. The table itself is a [NSTextTable] object. All [NSTextBlock]
// objects reference this table, which controls their sizing and positioning.
//
// # Working with dimensions of content
//
//   - [NSTextBlock.SetValueTypeForDimension]: Sets a dimension of the text block.
//   - [NSTextBlock.ValueForDimension]: Returns the value of the specified text block dimension.
//   - [NSTextBlock.ValueTypeForDimension]: Returns the value type of the specified text block dimension.
//   - [NSTextBlock.SetContentWidthType]: Sets the width of the text block.
//   - [NSTextBlock.ContentWidth]: The width of the text block.
//   - [NSTextBlock.ContentWidthValueType]: The type of value stored for the text block width.
//
// # Getting and setting margins, borders, and padding
//
//   - [NSTextBlock.SetWidthTypeForLayerEdge]: Sets the width of a specified edge of a specified layer of the text block.
//   - [NSTextBlock.SetWidthTypeForLayer]: Sets the width of all edges of a specified layer of the text block.
//   - [NSTextBlock.WidthForLayerEdge]: Returns the width of an edge of a specified layer of the text block.
//   - [NSTextBlock.WidthValueTypeForLayerEdge]: Returns the value type of an edge of a specified layer of the text block.
//
// # Getting and setting alignment
//
//   - [NSTextBlock.VerticalAlignment]: The vertical alignment of the text block.
//   - [NSTextBlock.SetVerticalAlignment]
//
// # Working with color
//
//   - [NSTextBlock.BackgroundColor]: The background color of the text block.
//   - [NSTextBlock.SetBackgroundColor]
//   - [NSTextBlock.SetBorderColorForEdge]: Sets the border color of the specified edge of the text block.
//   - [NSTextBlock.SetBorderColor]: Sets the color of all borders of the text block.
//   - [NSTextBlock.BorderColorForEdge]: Returns the border color of the specified text block edge.
//
// # Determining size and position of a text block
//
//   - [NSTextBlock.RectForLayoutAtPointInRectTextContainerCharacterRange]: Returns the rectangle within which glyphs should be laid out for the specified arguments.
//   - [NSTextBlock.BoundsRectForContentRectInRectTextContainerCharacterRange]: Returns the rectangle the text in the block actually occupies, including padding, borders, and margins.
//
// # Drawing colors and decorations
//
//   - [NSTextBlock.DrawBackgroundWithFrameInViewCharacterRangeLayoutManager]: Called by the layout manager to draw any colors and other decorations before the text is drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock
type NSTextBlock struct {
	objectivec.Object
}

// NSTextBlockFromID constructs a [NSTextBlock] from an objc.ID.
//
// A block of text laid out in a subregion of the text container.
func NSTextBlockFromID(id objc.ID) NSTextBlock {
	return NSTextBlock{objectivec.Object{ID: id}}
}
// NOTE: NSTextBlock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextBlock] class.
//
// # Working with dimensions of content
//
//   - [INSTextBlock.SetValueTypeForDimension]: Sets a dimension of the text block.
//   - [INSTextBlock.ValueForDimension]: Returns the value of the specified text block dimension.
//   - [INSTextBlock.ValueTypeForDimension]: Returns the value type of the specified text block dimension.
//   - [INSTextBlock.SetContentWidthType]: Sets the width of the text block.
//   - [INSTextBlock.ContentWidth]: The width of the text block.
//   - [INSTextBlock.ContentWidthValueType]: The type of value stored for the text block width.
//
// # Getting and setting margins, borders, and padding
//
//   - [INSTextBlock.SetWidthTypeForLayerEdge]: Sets the width of a specified edge of a specified layer of the text block.
//   - [INSTextBlock.SetWidthTypeForLayer]: Sets the width of all edges of a specified layer of the text block.
//   - [INSTextBlock.WidthForLayerEdge]: Returns the width of an edge of a specified layer of the text block.
//   - [INSTextBlock.WidthValueTypeForLayerEdge]: Returns the value type of an edge of a specified layer of the text block.
//
// # Getting and setting alignment
//
//   - [INSTextBlock.VerticalAlignment]: The vertical alignment of the text block.
//   - [INSTextBlock.SetVerticalAlignment]
//
// # Working with color
//
//   - [INSTextBlock.BackgroundColor]: The background color of the text block.
//   - [INSTextBlock.SetBackgroundColor]
//   - [INSTextBlock.SetBorderColorForEdge]: Sets the border color of the specified edge of the text block.
//   - [INSTextBlock.SetBorderColor]: Sets the color of all borders of the text block.
//   - [INSTextBlock.BorderColorForEdge]: Returns the border color of the specified text block edge.
//
// # Determining size and position of a text block
//
//   - [INSTextBlock.RectForLayoutAtPointInRectTextContainerCharacterRange]: Returns the rectangle within which glyphs should be laid out for the specified arguments.
//   - [INSTextBlock.BoundsRectForContentRectInRectTextContainerCharacterRange]: Returns the rectangle the text in the block actually occupies, including padding, borders, and margins.
//
// # Drawing colors and decorations
//
//   - [INSTextBlock.DrawBackgroundWithFrameInViewCharacterRangeLayoutManager]: Called by the layout manager to draw any colors and other decorations before the text is drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock
type INSTextBlock interface {
	objectivec.IObject

	// Topic: Working with dimensions of content

	// Sets a dimension of the text block.
	SetValueTypeForDimension(val float64, type_ NSTextBlockValueType, dimension NSTextBlockDimension)
	// Returns the value of the specified text block dimension.
	ValueForDimension(dimension NSTextBlockDimension) float64
	// Returns the value type of the specified text block dimension.
	ValueTypeForDimension(dimension NSTextBlockDimension) NSTextBlockValueType
	// Sets the width of the text block.
	SetContentWidthType(val float64, type_ NSTextBlockValueType)
	// The width of the text block.
	ContentWidth() float64
	// The type of value stored for the text block width.
	ContentWidthValueType() NSTextBlockValueType

	// Topic: Getting and setting margins, borders, and padding

	// Sets the width of a specified edge of a specified layer of the text block.
	SetWidthTypeForLayerEdge(val float64, type_ NSTextBlockValueType, layer NSTextBlockLayer, edge foundation.NSRectEdge)
	// Sets the width of all edges of a specified layer of the text block.
	SetWidthTypeForLayer(val float64, type_ NSTextBlockValueType, layer NSTextBlockLayer)
	// Returns the width of an edge of a specified layer of the text block.
	WidthForLayerEdge(layer NSTextBlockLayer, edge foundation.NSRectEdge) float64
	// Returns the value type of an edge of a specified layer of the text block.
	WidthValueTypeForLayerEdge(layer NSTextBlockLayer, edge foundation.NSRectEdge) NSTextBlockValueType

	// Topic: Getting and setting alignment

	// The vertical alignment of the text block.
	VerticalAlignment() NSTextBlockVerticalAlignment
	SetVerticalAlignment(value NSTextBlockVerticalAlignment)

	// Topic: Working with color

	// The background color of the text block.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// Sets the border color of the specified edge of the text block.
	SetBorderColorForEdge(color INSColor, edge foundation.NSRectEdge)
	// Sets the color of all borders of the text block.
	SetBorderColor(color INSColor)
	// Returns the border color of the specified text block edge.
	BorderColorForEdge(edge foundation.NSRectEdge) INSColor

	// Topic: Determining size and position of a text block

	// Returns the rectangle within which glyphs should be laid out for the specified arguments.
	RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect
	// Returns the rectangle the text in the block actually occupies, including padding, borders, and margins.
	BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect

	// Topic: Drawing colors and decorations

	// Called by the layout manager to draw any colors and other decorations before the text is drawn.
	DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect corefoundation.CGRect, controlView INSView, charRange foundation.NSRange, layoutManager INSLayoutManager)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextBlock) Init() NSTextBlock {
	rv := objc.Send[NSTextBlock](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextBlock) Autorelease() NSTextBlock {
	rv := objc.Send[NSTextBlock](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextBlock creates a new NSTextBlock instance.
func NewNSTextBlock() NSTextBlock {
	class := getNSTextBlockClass()
	rv := objc.Send[NSTextBlock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets a dimension of the text block.
//
// val: The new value for the dimension.
//
// type: The type of value being provided. This controls how `val` is interpreted.
//
// dimension: The dimension to set.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setValue(_:type:for:)
func (t NSTextBlock) SetValueTypeForDimension(val float64, type_ NSTextBlockValueType, dimension NSTextBlockDimension) {
	objc.Send[objc.ID](t.ID, objc.Sel("setValue:type:forDimension:"), val, type_, dimension)
}
// Returns the value of the specified text block dimension.
//
// # Return Value
// 
// The value for the specified dimension. This value should be interpreted
// according to the value type returned by [ValueTypeForDimension].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/value(for:)
func (t NSTextBlock) ValueForDimension(dimension NSTextBlockDimension) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("valueForDimension:"), dimension)
	return rv
}
// Returns the value type of the specified text block dimension.
//
// # Return Value
// 
// The value type for the specified text block dimension. This result
// determines how the value for the dimension should be interpreted.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/valueType(for:)
func (t NSTextBlock) ValueTypeForDimension(dimension NSTextBlockDimension) NSTextBlockValueType {
	rv := objc.Send[NSTextBlockValueType](t.ID, objc.Sel("valueTypeForDimension:"), dimension)
	return NSTextBlockValueType(rv)
}
// Sets the width of the text block.
//
// val: The new value for the width.
//
// type: The type of value being provided. This controls how `val` is interpreted.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setContentWidth(_:type:)
func (t NSTextBlock) SetContentWidthType(val float64, type_ NSTextBlockValueType) {
	objc.Send[objc.ID](t.ID, objc.Sel("setContentWidth:type:"), val, type_)
}
// Sets the width of a specified edge of a specified layer of the text block.
//
// val: The new value for the specified edge width.
//
// type: The type of value being provided. This controls how `val` is interpreted.
//
// layer: The layer of the text block to modify.
//
// edge: The edge of the layer to modify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setWidth(_:type:for:edge:)
func (t NSTextBlock) SetWidthTypeForLayerEdge(val float64, type_ NSTextBlockValueType, layer NSTextBlockLayer, edge foundation.NSRectEdge) {
	objc.Send[objc.ID](t.ID, objc.Sel("setWidth:type:forLayer:edge:"), val, type_, layer, edge)
}
// Sets the width of all edges of a specified layer of the text block.
//
// val: The new value for the specified edge width.
//
// type: The type of value being provided. This controls how `val` is interpreted.
//
// layer: The layer of the text block to modify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setWidth(_:type:for:)
func (t NSTextBlock) SetWidthTypeForLayer(val float64, type_ NSTextBlockValueType, layer NSTextBlockLayer) {
	objc.Send[objc.ID](t.ID, objc.Sel("setWidth:type:forLayer:"), val, type_, layer)
}
// Returns the width of an edge of a specified layer of the text block.
//
// layer: The layer to examine.
//
// edge: The edge of the layer to examine.
//
// # Return Value
// 
// The width of the `edge` of `layer`. This value must be interpreted
// according to the value type returned by [WidthValueTypeForLayerEdge].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/width(for:edge:)
func (t NSTextBlock) WidthForLayerEdge(layer NSTextBlockLayer, edge foundation.NSRectEdge) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("widthForLayer:edge:"), layer, edge)
	return rv
}
// Returns the value type of an edge of a specified layer of the text block.
//
// layer: The layer to examine.
//
// edge: The edge of the layer to examine.
//
// # Return Value
// 
// The value type of the `edge` of `layer`. This determines how the value for
// this `edge` of `layer` should be interpreted.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/widthValueType(for:edge:)
func (t NSTextBlock) WidthValueTypeForLayerEdge(layer NSTextBlockLayer, edge foundation.NSRectEdge) NSTextBlockValueType {
	rv := objc.Send[NSTextBlockValueType](t.ID, objc.Sel("widthValueTypeForLayer:edge:"), layer, edge)
	return NSTextBlockValueType(rv)
}
// Sets the border color of the specified edge of the text block.
//
// color: The new color.
//
// edge: The edge whose color is to be set.
//
// # Discussion
// 
// This setting has no visible effect unless the border width is larger than
// the default, which is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setBorderColor(_:for:)
func (t NSTextBlock) SetBorderColorForEdge(color INSColor, edge foundation.NSRectEdge) {
	objc.Send[objc.ID](t.ID, objc.Sel("setBorderColor:forEdge:"), color, edge)
}
// Sets the color of all borders of the text block.
//
// color: The new color.
//
// # Discussion
// 
// This setting has no visible effect unless the border width is larger than
// the default, which is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/setBorderColor(_:)
func (t NSTextBlock) SetBorderColor(color INSColor) {
	objc.Send[objc.ID](t.ID, objc.Sel("setBorderColor:"), color)
}
// Returns the border color of the specified text block edge.
//
// edge: The edge of the text block in question.
//
// # Return Value
// 
// The border color of the text block edge `edge`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/borderColor(for:)
func (t NSTextBlock) BorderColorForEdge(edge foundation.NSRectEdge) INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("borderColorForEdge:"), edge)
	return NSColorFromID(rv)
}
// Returns the rectangle within which glyphs should be laid out for the
// specified arguments.
//
// startingPoint: The location, in container coordinates, where layout begins.
//
// rect: The rectangle in which the block is constrained to lie. For top-level
// blocks, this is the container rectangle of `textContainer`; for nested
// blocks, this is the layout rectangle of the enclosing block.
//
// textContainer: The text container being used for the layout.
//
// charRange: The range of the characters in the [NSTextStorage] object whose glyphs are
// to be drawn.
//
// # Return Value
// 
// The rectangle within which glyphs should be laid out.
//
// # Discussion
// 
// This method is called by the typesetter before the text block is laid out
// to return the rectangle within which glyphs should be laid out.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/rectForLayout(at:in:textContainer:characterRange:)
func (t NSTextBlock) RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("rectForLayoutAtPoint:inRect:textContainer:characterRange:"), startingPoint, rect, textContainer, charRange)
	return corefoundation.CGRect(rv)
}
// Returns the rectangle the text in the block actually occupies, including
// padding, borders, and margins.
//
// contentRect: The actual rectangle in which the text was laid out, as determined by
// [RectForLayoutAtPointInRectTextContainerCharacterRange].
//
// rect: The initial rectangle in `textContainer` proposed by the typesetter.
//
// textContainer: The text container being used for the layout.
//
// charRange: The range of the characters in the [NSTextStorage] object whose glyphs are
// to be drawn.
//
// # Return Value
// 
// The rectangle the text in the block actually occupies, including padding,
// borders, and margins.
//
// # Discussion
// 
// This methods is called by the typesetter after the text block is laid out
// to return the rectangle the text in the block actually occupies, including
// padding, borders, and margins.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/boundsRect(forContentRect:in:textContainer:characterRange:)
func (t NSTextBlock) BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("boundsRectForContentRect:inRect:textContainer:characterRange:"), contentRect, rect, textContainer, charRange)
	return corefoundation.CGRect(rv)
}
// Called by the layout manager to draw any colors and other decorations
// before the text is drawn.
//
// frameRect: The bounds rectangle in view coordinates.
//
// controlView: The view in which drawing occurs.
//
// charRange: The range of the characters in the [NSTextStorage] object whose glyphs are
// to be drawn.
//
// layoutManager: The layout manager controlling the typesetting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/drawBackground(withFrame:in:characterRange:layoutManager:)
func (t NSTextBlock) DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect corefoundation.CGRect, controlView INSView, charRange foundation.NSRange, layoutManager INSLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawBackgroundWithFrame:inView:characterRange:layoutManager:"), frameRect, controlView, charRange, layoutManager)
}
func (t NSTextBlock) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The width of the text block.
//
// # Discussion
// 
// This property interpreted according to the value type returned by
// [ContentWidthValueType].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/contentWidth
func (t NSTextBlock) ContentWidth() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("contentWidth"))
	return rv
}
// The type of value stored for the text block width.
//
// # Discussion
// 
// This property determines how the width value should be interpreted.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/contentWidthValueType
func (t NSTextBlock) ContentWidthValueType() NSTextBlockValueType {
	rv := objc.Send[NSTextBlockValueType](t.ID, objc.Sel("contentWidthValueType"))
	return NSTextBlockValueType(rv)
}
// The vertical alignment of the text block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/verticalAlignment-swift.property
func (t NSTextBlock) VerticalAlignment() NSTextBlockVerticalAlignment {
	rv := objc.Send[NSTextBlockVerticalAlignment](t.ID, objc.Sel("verticalAlignment"))
	return NSTextBlockVerticalAlignment(rv)
}
func (t NSTextBlock) SetVerticalAlignment(value NSTextBlockVerticalAlignment) {
	objc.Send[struct{}](t.ID, objc.Sel("setVerticalAlignment:"), value)
}
// The background color of the text block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextBlock/backgroundColor
func (t NSTextBlock) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextBlock) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}

