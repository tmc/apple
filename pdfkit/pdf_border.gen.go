// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

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

	InitWithCoder(coder foundation.INSCoder) PDFBorder
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

// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/init(coder:)
func NewPDFBorderWithCoder(coder foundation.INSCoder) PDFBorder {
	instance := getPDFBorderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return PDFBorderFromID(rv)
}

// Draws the border.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/draw(in:)
func (p PDFBorder) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawInRect:"), rect)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFBorder/init(coder:)
func (p PDFBorder) InitWithCoder(coder foundation.INSCoder) PDFBorder {
	rv := objc.Send[PDFBorder](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
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
