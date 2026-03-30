// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFBorder] class.
var (
	_PDFBorderClass     PDFBorderClass
	_PDFBorderClassOnce sync.Once
)

func getPDFBorderClass() PDFBorderClass {
	_PDFBorderClassOnce.Do(func() {
		_PDFBorderClass = PDFBorderClass{class: objc.GetClass("PDFBorder")}
	})
	return _PDFBorderClass
}

// GetPDFBorderClass returns the class object for PDFBorder.
func GetPDFBorderClass() PDFBorderClass {
	return getPDFBorderClass()
}

type PDFBorderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFBorderClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFBorderClass) Alloc() PDFBorder {
	rv := objc.Send[PDFBorder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An optional border for an annotation that lies completely within the
// annotation rectangle.
//
// # Working with Border Styles and Characteristics
//
//   - [PDFBorder.Style]: Sets the border style.
//   - [PDFBorder.SetStyle]
//   - [PDFBorder.LineWidth]: Sets the line width (in points) for the border.
//   - [PDFBorder.SetLineWidth]
//   - [PDFBorder.DashPattern]: Gets the dash pattern for the border as an array of NSNumber objects.
//   - [PDFBorder.SetDashPattern]
//   - [PDFBorder.BorderKeyValues]: A dictionary that contains a deep copy of all border properties.
//
// # Drawing Borders
//
//   - [PDFBorder.DrawInRect]: Draws the border.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder
type PDFBorder struct {
	objectivec.Object
}

// PDFBorderFromID constructs a [PDFBorder] from an objc.ID.
//
// An optional border for an annotation that lies completely within the
// annotation rectangle.
func PDFBorderFromID(id objc.ID) PDFBorder {
	return PDFBorder{objectivec.Object{ID: id}}
}

// NOTE: PDFBorder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFBorder] class.
//
// # Working with Border Styles and Characteristics
//
//   - [IPDFBorder.Style]: Sets the border style.
//   - [IPDFBorder.SetStyle]
//   - [IPDFBorder.LineWidth]: Sets the line width (in points) for the border.
//   - [IPDFBorder.SetLineWidth]
//   - [IPDFBorder.DashPattern]: Gets the dash pattern for the border as an array of NSNumber objects.
//   - [IPDFBorder.SetDashPattern]
//   - [IPDFBorder.BorderKeyValues]: A dictionary that contains a deep copy of all border properties.
//
// # Drawing Borders
//
//   - [IPDFBorder.DrawInRect]: Draws the border.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder
type IPDFBorder interface {
	objectivec.IObject

	// Topic: Working with Border Styles and Characteristics

	// Sets the border style.
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
	// Sets the line width (in points) for the border.
	LineWidth() float64
	SetLineWidth(value float64)
	// Gets the dash pattern for the border as an array of NSNumber objects.
	DashPattern() foundation.INSArray
	SetDashPattern(value foundation.INSArray)
	// A dictionary that contains a deep copy of all border properties.
	BorderKeyValues() foundation.INSDictionary

	// Topic: Drawing Borders

	// Draws the border.
	DrawInRect(rect corefoundation.CGRect)

	// The alignment of the free text and text widget annotation’s text content.
	Alignment() appkit.NSTextAlignment
	SetAlignment(value appkit.NSTextAlignment)
	// Sets the border style for the annotation.
	Border() IPDFBorder
	SetBorder(value IPDFBorder)
	// Returns the bounding box for the annotation in page space.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	// Sets the stroke color for the annotation.
	Color() objc.ID
	SetColor(value objc.ID)
	// Returns the textual content (if any) associated with the annotation.
	Contents() string
	SetContents(value string)
	// The font the annotation uses to display text.
	Font() objectivec.IObject
	SetFont(value objectivec.IObject)
	// The font color the annotation uses to display text.
	FontColor() objc.ID
	SetFontColor(value objc.ID)
	// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
	HasAppearanceStream() bool
	SetHasAppearanceStream(value bool)
	// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p PDFBorder) Init() PDFBorder {
	rv := objc.Send[PDFBorder](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFBorder) Autorelease() PDFBorder {
	rv := objc.Send[PDFBorder](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFBorder creates a new PDFBorder instance.
func NewPDFBorder() PDFBorder {
	class := getPDFBorderClass()
	rv := objc.Send[PDFBorder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Draws the border.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/draw(in:)
func (p PDFBorder) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawInRect:"), rect)
}
func (p PDFBorder) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Sets the border style.
//
// # Discussion
//
// Refer to [Constants] for the available border styles.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p PDFBorder) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p.ID, objc.Sel("style"))
	return PDFBorderStyle(rv)
}
func (p PDFBorder) SetStyle(value PDFBorderStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setStyle:"), value)
}

// Sets the line width (in points) for the border.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/lineWidth
func (p PDFBorder) LineWidth() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("lineWidth"))
	return rv
}
func (p PDFBorder) SetLineWidth(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLineWidth:"), value)
}

// Gets the dash pattern for the border as an array of NSNumber objects.
//
// # Discussion
//
// Refer to the description for [NSBezierPath] for more information.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p PDFBorder) DashPattern() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dashPattern"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (p PDFBorder) SetDashPattern(value foundation.INSArray) {
	objc.Send[struct{}](p.ID, objc.Sel("setDashPattern:"), value)
}

// A dictionary that contains a deep copy of all border properties.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/borderKeyValues
func (p PDFBorder) BorderKeyValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("borderKeyValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The alignment of the free text and text widget annotation’s text content.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p PDFBorder) Alignment() appkit.NSTextAlignment {
	rv := objc.Send[appkit.NSTextAlignment](p.ID, objc.Sel("alignment"))
	return appkit.NSTextAlignment(rv)
}
func (p PDFBorder) SetAlignment(value appkit.NSTextAlignment) {
	objc.Send[struct{}](p.ID, objc.Sel("setAlignment:"), value)
}

// Sets the border style for the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p PDFBorder) Border() IPDFBorder {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("border"))
	return PDFBorderFromID(objc.ID(rv))
}
func (p PDFBorder) SetBorder(value IPDFBorder) {
	objc.Send[struct{}](p.ID, objc.Sel("setBorder:"), value)
}

// Returns the bounding box for the annotation in page space.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p PDFBorder) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (p PDFBorder) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](p.ID, objc.Sel("setBounds:"), value)
}

// Sets the stroke color for the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p PDFBorder) Color() objc.ID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("color"))
	return rv
}
func (p PDFBorder) SetColor(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setColor:"), value)
}

// Returns the textual content (if any) associated with the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p PDFBorder) Contents() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contents"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFBorder) SetContents(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setContents:"), objc.String(value))
}

// The font the annotation uses to display text.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p PDFBorder) Font() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("font"))
	return objectivec.Object{ID: rv}
}
func (p PDFBorder) SetFont(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setFont:"), value)
}

// The font color the annotation uses to display text.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p PDFBorder) FontColor() objc.ID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fontColor"))
	return rv
}
func (p PDFBorder) SetFontColor(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setFontColor:"), value)
}

// Returns a Boolean value that indicates whether the annotation has an
// appearance stream associated with it.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p PDFBorder) HasAppearanceStream() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasAppearanceStream"))
	return rv
}
func (p PDFBorder) SetHasAppearanceStream(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHasAppearanceStream:"), value)
}

// A Boolean value that indicates whether the annotation is in a highlighted
// state, such as when the mouse is down on a link annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p PDFBorder) IsHighlighted() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("highlighted"))
	return rv
}
func (p PDFBorder) SetIsHighlighted(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHighlighted:"), value)
}
