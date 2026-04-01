// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFAppearanceCharacteristics] class.
var (
	_PDFAppearanceCharacteristicsClass     PDFAppearanceCharacteristicsClass
	_PDFAppearanceCharacteristicsClassOnce sync.Once
)

func getPDFAppearanceCharacteristicsClass() PDFAppearanceCharacteristicsClass {
	_PDFAppearanceCharacteristicsClassOnce.Do(func() {
		_PDFAppearanceCharacteristicsClass = PDFAppearanceCharacteristicsClass{class: objc.GetClass("PDFAppearanceCharacteristics")}
	})
	return _PDFAppearanceCharacteristicsClass
}

// GetPDFAppearanceCharacteristicsClass returns the class object for PDFAppearanceCharacteristics.
func GetPDFAppearanceCharacteristicsClass() PDFAppearanceCharacteristicsClass {
	return getPDFAppearanceCharacteristicsClass()
}

type PDFAppearanceCharacteristicsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFAppearanceCharacteristicsClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFAppearanceCharacteristicsClass) Alloc() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents appearance characteristics of a widget
// annotation.
//
// # Configuring a Widget’s Appearance
//
//   - [PDFAppearanceCharacteristics.BackgroundColor]: The background color of the widget annotation.
//   - [PDFAppearanceCharacteristics.SetBackgroundColor]
//   - [PDFAppearanceCharacteristics.BorderColor]: The border color of the widget annotation.
//   - [PDFAppearanceCharacteristics.SetBorderColor]
//   - [PDFAppearanceCharacteristics.Caption]: The text that the button widget annotation displays when the user isn’t interacting with it.
//   - [PDFAppearanceCharacteristics.SetCaption]
//   - [PDFAppearanceCharacteristics.ControlType]: The type of button widget annotation.
//   - [PDFAppearanceCharacteristics.SetControlType]
//   - [PDFAppearanceCharacteristics.DownCaption]: The text that the button widget annotation displays when the user holds down on it.
//   - [PDFAppearanceCharacteristics.SetDownCaption]
//   - [PDFAppearanceCharacteristics.RolloverCaption]: The text that the widget annotation displays when the user hovers the pointer over it.
//   - [PDFAppearanceCharacteristics.SetRolloverCaption]
//   - [PDFAppearanceCharacteristics.Rotation]: The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//   - [PDFAppearanceCharacteristics.SetRotation]
//   - [PDFAppearanceCharacteristics.AppearanceCharacteristicsKeyValues]: A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics
type PDFAppearanceCharacteristics struct {
	objectivec.Object
}

// PDFAppearanceCharacteristicsFromID constructs a [PDFAppearanceCharacteristics] from an objc.ID.
//
// An object that represents appearance characteristics of a widget
// annotation.
func PDFAppearanceCharacteristicsFromID(id objc.ID) PDFAppearanceCharacteristics {
	return PDFAppearanceCharacteristics{objectivec.Object{ID: id}}
}

// NOTE: PDFAppearanceCharacteristics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFAppearanceCharacteristics] class.
//
// # Configuring a Widget’s Appearance
//
//   - [IPDFAppearanceCharacteristics.BackgroundColor]: The background color of the widget annotation.
//   - [IPDFAppearanceCharacteristics.SetBackgroundColor]
//   - [IPDFAppearanceCharacteristics.BorderColor]: The border color of the widget annotation.
//   - [IPDFAppearanceCharacteristics.SetBorderColor]
//   - [IPDFAppearanceCharacteristics.Caption]: The text that the button widget annotation displays when the user isn’t interacting with it.
//   - [IPDFAppearanceCharacteristics.SetCaption]
//   - [IPDFAppearanceCharacteristics.ControlType]: The type of button widget annotation.
//   - [IPDFAppearanceCharacteristics.SetControlType]
//   - [IPDFAppearanceCharacteristics.DownCaption]: The text that the button widget annotation displays when the user holds down on it.
//   - [IPDFAppearanceCharacteristics.SetDownCaption]
//   - [IPDFAppearanceCharacteristics.RolloverCaption]: The text that the widget annotation displays when the user hovers the pointer over it.
//   - [IPDFAppearanceCharacteristics.SetRolloverCaption]
//   - [IPDFAppearanceCharacteristics.Rotation]: The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
//   - [IPDFAppearanceCharacteristics.SetRotation]
//   - [IPDFAppearanceCharacteristics.AppearanceCharacteristicsKeyValues]: A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics
type IPDFAppearanceCharacteristics interface {
	objectivec.IObject

	// Topic: Configuring a Widget’s Appearance

	// The background color of the widget annotation.
	BackgroundColor() appkit.NSColor
	SetBackgroundColor(value appkit.NSColor)
	// The border color of the widget annotation.
	BorderColor() appkit.NSColor
	SetBorderColor(value appkit.NSColor)
	// The text that the button widget annotation displays when the user isn’t interacting with it.
	Caption() string
	SetCaption(value string)
	// The type of button widget annotation.
	ControlType() PDFWidgetControlType
	SetControlType(value PDFWidgetControlType)
	// The text that the button widget annotation displays when the user holds down on it.
	DownCaption() string
	SetDownCaption(value string)
	// The text that the widget annotation displays when the user hovers the pointer over it.
	RolloverCaption() string
	SetRolloverCaption(value string)
	// The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
	Rotation() int
	SetRotation(value int)
	// A dictionary that contains a deep copy of the appearance characteristic key-value pairs.
	AppearanceCharacteristicsKeyValues() foundation.INSDictionary

	// The widget identifier for form annotation actions and behaviors.
	FieldName() string
	SetFieldName(value string)
	// A Boolean value that determines whether the widget is editable.
	IsReadOnly() bool
	SetIsReadOnly(value bool)
	// The string value that the widget reverts to when performing a reset form action.
	WidgetDefaultStringValue() string
	SetWidgetDefaultStringValue(value string)
	// The type of widget annotation, such as button, choice, or text.
	WidgetFieldType() PDFAnnotationWidgetSubtype
	SetWidgetFieldType(value PDFAnnotationWidgetSubtype)
	// The string value of the widget annotation.
	WidgetStringValue() string
	SetWidgetStringValue(value string)
}

// Init initializes the instance.
func (p PDFAppearanceCharacteristics) Init() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFAppearanceCharacteristics) Autorelease() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAppearanceCharacteristics creates a new PDFAppearanceCharacteristics instance.
func NewPDFAppearanceCharacteristics() PDFAppearanceCharacteristics {
	class := getPDFAppearanceCharacteristicsClass()
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The background color of the widget annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/backgroundColor
func (p PDFAppearanceCharacteristics) BackgroundColor() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAppearanceCharacteristics) SetBackgroundColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}

// The border color of the widget annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/borderColor
func (p PDFAppearanceCharacteristics) BorderColor() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("borderColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAppearanceCharacteristics) SetBorderColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setBorderColor:"), value)
}

// The text that the button widget annotation displays when the user isn’t
// interacting with it.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/caption
func (p PDFAppearanceCharacteristics) Caption() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("caption"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetCaption(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setCaption:"), objc.String(value))
}

// The type of button widget annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/controlType
func (p PDFAppearanceCharacteristics) ControlType() PDFWidgetControlType {
	rv := objc.Send[PDFWidgetControlType](p.ID, objc.Sel("controlType"))
	return PDFWidgetControlType(rv)
}
func (p PDFAppearanceCharacteristics) SetControlType(value PDFWidgetControlType) {
	objc.Send[struct{}](p.ID, objc.Sel("setControlType:"), value)
}

// The text that the button widget annotation displays when the user holds
// down on it.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/downCaption
func (p PDFAppearanceCharacteristics) DownCaption() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("downCaption"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetDownCaption(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setDownCaption:"), objc.String(value))
}

// The text that the widget annotation displays when the user hovers the
// pointer over it.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rolloverCaption
func (p PDFAppearanceCharacteristics) RolloverCaption() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("rolloverCaption"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetRolloverCaption(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setRolloverCaption:"), objc.String(value))
}

// The number of degrees, in multiples of 90, that the widget annotation
// rotates counterclockwise relative to the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/rotation
func (p PDFAppearanceCharacteristics) Rotation() int {
	rv := objc.Send[int](p.ID, objc.Sel("rotation"))
	return rv
}
func (p PDFAppearanceCharacteristics) SetRotation(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setRotation:"), value)
}

// A dictionary that contains a deep copy of the appearance characteristic
// key-value pairs.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics/appearanceCharacteristicsKeyValues
func (p PDFAppearanceCharacteristics) AppearanceCharacteristicsKeyValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("appearanceCharacteristicsKeyValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The widget identifier for form annotation actions and behaviors.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/fieldname
func (p PDFAppearanceCharacteristics) FieldName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fieldName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetFieldName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setFieldName:"), objc.String(value))
}

// A Boolean value that determines whether the widget is editable.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p PDFAppearanceCharacteristics) IsReadOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("readOnly"))
	return rv
}
func (p PDFAppearanceCharacteristics) SetIsReadOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setReadOnly:"), value)
}

// The string value that the widget reverts to when performing a reset form
// action.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetdefaultstringvalue
func (p PDFAppearanceCharacteristics) WidgetDefaultStringValue() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("widgetDefaultStringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetWidgetDefaultStringValue(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetDefaultStringValue:"), objc.String(value))
}

// The type of widget annotation, such as button, choice, or text.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetfieldtype
func (p PDFAppearanceCharacteristics) WidgetFieldType() PDFAnnotationWidgetSubtype {
	rv := objc.Send[PDFAnnotationWidgetSubtype](p.ID, objc.Sel("widgetFieldType"))
	return PDFAnnotationWidgetSubtype(rv)
}
func (p PDFAppearanceCharacteristics) SetWidgetFieldType(value PDFAnnotationWidgetSubtype) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetFieldType:"), value)
}

// The string value of the widget annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/widgetstringvalue
func (p PDFAppearanceCharacteristics) WidgetStringValue() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("widgetStringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAppearanceCharacteristics) SetWidgetStringValue(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetStringValue:"), objc.String(value))
}
