// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFAnnotation] class.
var (
	_PDFAnnotationClass     PDFAnnotationClass
	_PDFAnnotationClassOnce sync.Once
)

func getPDFAnnotationClass() PDFAnnotationClass {
	_PDFAnnotationClassOnce.Do(func() {
		_PDFAnnotationClass = PDFAnnotationClass{class: objc.GetClass("PDFAnnotation")}
	})
	return _PDFAnnotationClass
}

// GetPDFAnnotationClass returns the class object for PDFAnnotation.
func GetPDFAnnotationClass() PDFAnnotationClass {
	return getPDFAnnotationClass()
}

type PDFAnnotationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFAnnotationClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFAnnotationClass) Alloc() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An annotation in a PDF document.
//
// # Overview
//
// In addition to its primary textual content, a PDF file can contain
// annotations that represent links, form elements, highlighting circles,
// textual notes, and so on. Each annotation has a specific location on a page
// and may offer interactivity with the user.
//
// # Creating an Annotation
//
//   - [PDFAnnotation.InitWithBoundsForTypeWithProperties]: Creates a PDF annotation with the specified bounds, type, and optional properties.
//
// # Accessing Information About an Annotation
//
//   - [PDFAnnotation.Page]: Returns the page that the annotation is associated with.
//   - [PDFAnnotation.SetPage]
//   - [PDFAnnotation.ModificationDate]: Returns the modification date of the annotation.
//   - [PDFAnnotation.SetModificationDate]
//   - [PDFAnnotation.UserName]: Returns the name of the user who created the annotation.
//   - [PDFAnnotation.SetUserName]
//   - [PDFAnnotation.Type]: Returns the type of the annotation.
//   - [PDFAnnotation.SetType]
//   - [PDFAnnotation.Action]: An object that represents an action for a PDF element, such as a link annotation.
//   - [PDFAnnotation.SetAction]
//
// # Managing Annotation Drawing and Output
//
//   - [PDFAnnotation.DrawWithBoxInContext]: Draws the annotation in a graphics context using page-space coordinates relative to the origin of the specified box.
//   - [PDFAnnotation.ShouldDisplay]: Returns a Boolean value indicating whether the annotation should be displayed.
//   - [PDFAnnotation.SetShouldDisplay]
//   - [PDFAnnotation.ShouldPrint]: Returns a Boolean value indicating whether the annotation should appear when the document is printed.
//   - [PDFAnnotation.SetShouldPrint]
//
// # Modifying Annotation Attributes
//
//   - [PDFAnnotation.AnnotationKeyValues]: A dictionary that contains a deep copy of the widget’s properties.
//   - [PDFAnnotation.ValueForAnnotationKey]: Returns a deep copy of the key-value pairs of properties for the specified key.
//   - [PDFAnnotation.SetValueForAnnotationKey]: Sets a value in the annotation’s dictionary.
//   - [PDFAnnotation.SetBooleanForAnnotationKey]: Sets a Boolean value in the annotation’s dictionary.
//   - [PDFAnnotation.SetRectForAnnotationKey]: Sets a rectangle value in the annotation’s dictionary.
//   - [PDFAnnotation.RemoveValueForAnnotationKey]: Removes a value from the annotation’s dictionary.
//
// # Managing Annotation Display Characteristics
//
//   - [PDFAnnotation.Alignment]: The alignment of the free text and text widget annotation’s text content.
//   - [PDFAnnotation.SetAlignment]
//   - [PDFAnnotation.Bounds]: Returns the bounding box for the annotation in page space.
//   - [PDFAnnotation.SetBounds]
//   - [PDFAnnotation.Contents]: Returns the textual content (if any) associated with the annotation.
//   - [PDFAnnotation.SetContents]
//   - [PDFAnnotation.Font]: The font the annotation uses to display text.
//   - [PDFAnnotation.SetFont]
//   - [PDFAnnotation.FontColor]: The font color the annotation uses to display text.
//   - [PDFAnnotation.SetFontColor]
//   - [PDFAnnotation.Border]: Sets the border style for the annotation.
//   - [PDFAnnotation.SetBorder]
//   - [PDFAnnotation.Highlighted]: A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//   - [PDFAnnotation.SetHighlighted]
//   - [PDFAnnotation.Color]: Sets the stroke color for the annotation.
//   - [PDFAnnotation.SetColor]
//   - [PDFAnnotation.HasAppearanceStream]: Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// # Configuring Shape Annotations
//
//   - [PDFAnnotation.InteriorColor]: The fill color for drawing a circle, line, or square annotation.
//   - [PDFAnnotation.SetInteriorColor]
//
// # Configuring Line Annotations
//
//   - [PDFAnnotation.StartPoint]: The point where a line begins, in annotation-space coordinates.
//   - [PDFAnnotation.SetStartPoint]
//   - [PDFAnnotation.EndPoint]: The point where a line ends, in annotation-space coordinates.
//   - [PDFAnnotation.SetEndPoint]
//   - [PDFAnnotation.StartLineStyle]: The style of the line annotation’s starting point, such as square or filled arrowhead.
//   - [PDFAnnotation.SetStartLineStyle]
//   - [PDFAnnotation.EndLineStyle]: The style of the line annotation’s ending point, such as square or filled arrowhead.
//   - [PDFAnnotation.SetEndLineStyle]
//
// # Configuring Link Annotations
//
//   - [PDFAnnotation.Destination]: The destination for a link annotation.
//   - [PDFAnnotation.SetDestination]
//   - [PDFAnnotation.URL]: A URL for a link annotation.
//   - [PDFAnnotation.SetURL]
//
// # Configuring Text Annotations
//
//   - [PDFAnnotation.IconType]: The type of icon to display for a pop-up text annotation.
//   - [PDFAnnotation.SetIconType]
//
// # Configuring Pop-Up Annotations
//
//   - [PDFAnnotation.Popup]: Returns the pop-up annotation associated with an annotation.
//   - [PDFAnnotation.SetPopup]
//   - [PDFAnnotation.Open]: A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//   - [PDFAnnotation.SetOpen]
//
// # Configuring Text Markup Annotations
//
//   - [PDFAnnotation.MarkupType]: The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
//   - [PDFAnnotation.SetMarkupType]
//   - [PDFAnnotation.QuadrilateralPoints]: An array of values that represents the points bounding the marked-up text.
//   - [PDFAnnotation.SetQuadrilateralPoints]
//
// # Configuring Widget Annotations
//
//   - [PDFAnnotation.WidgetFieldType]: The type of widget annotation, such as button, choice, or text.
//   - [PDFAnnotation.SetWidgetFieldType]
//   - [PDFAnnotation.WidgetStringValue]: The string value of the widget annotation.
//   - [PDFAnnotation.SetWidgetStringValue]
//   - [PDFAnnotation.WidgetDefaultStringValue]: The string value that the widget reverts to when performing a reset form action.
//   - [PDFAnnotation.SetWidgetDefaultStringValue]
//   - [PDFAnnotation.FieldName]: The widget identifier for form annotation actions and behaviors.
//   - [PDFAnnotation.SetFieldName]
//   - [PDFAnnotation.BackgroundColor]: The color of the widget’s background.
//   - [PDFAnnotation.SetBackgroundColor]
//   - [PDFAnnotation.ReadOnly]: A Boolean value that determines whether the widget is editable.
//   - [PDFAnnotation.SetReadOnly]
//
// # Configuring Text Widget Annotations
//
//   - [PDFAnnotation.Multiline]: A Boolean value that indicates whether the text widget annotation displays multiple lines.
//   - [PDFAnnotation.SetMultiline]
//   - [PDFAnnotation.IsPasswordField]: A Boolean value that indicates whether the text widget annotation displays a password field using bullet characters.
//   - [PDFAnnotation.MaximumLength]: The maximum number of characters the text widget annotation allows.
//   - [PDFAnnotation.SetMaximumLength]
//   - [PDFAnnotation.Comb]: A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//   - [PDFAnnotation.SetComb]
//
// # Configuring Button Widget Annotations
//
//   - [PDFAnnotation.WidgetControlType]: The type of button widget control, either radio button, push button, or checkbox.
//   - [PDFAnnotation.SetWidgetControlType]
//   - [PDFAnnotation.ButtonWidgetState]: The current state of the button widget annotation.
//   - [PDFAnnotation.SetButtonWidgetState]
//   - [PDFAnnotation.ButtonWidgetStateString]: A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
//   - [PDFAnnotation.SetButtonWidgetStateString]
//   - [PDFAnnotation.Caption]: The title of push button widget annotations.
//   - [PDFAnnotation.SetCaption]
//   - [PDFAnnotation.AllowsToggleToOff]: A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
//   - [PDFAnnotation.SetAllowsToggleToOff]
//   - [PDFAnnotation.RadiosInUnison]: A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
//   - [PDFAnnotation.SetRadiosInUnison]
//
// # Configuring Choice Widget Annotations
//
//   - [PDFAnnotation.Choices]: An array of strings that specifies the options in either a list or a pop-up menu.
//   - [PDFAnnotation.SetChoices]
//   - [PDFAnnotation.ListChoice]: A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//   - [PDFAnnotation.SetListChoice]
//   - [PDFAnnotation.Values]: An array of strings that specifies the export values for items in a list or a pop-up menu.
//   - [PDFAnnotation.SetValues]
//
// # Configuring Ink Annotations
//
//   - [PDFAnnotation.Paths]: An array of bezier paths, in annotation-space coordinates, that compose the annotation.
//   - [PDFAnnotation.AddBezierPath]: Adds a bezier path to the ink annotation.
//   - [PDFAnnotation.RemoveBezierPath]: Removes a bezier path from an ink annotation.
//
// # Configuring Stamp Annotations
//
//   - [PDFAnnotation.StampName]: The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
//   - [PDFAnnotation.SetStampName]
//
// # Instance Properties
//
//   - [PDFAnnotation.ActivatableTextField]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation
type PDFAnnotation struct {
	objectivec.Object
}

// PDFAnnotationFromID constructs a [PDFAnnotation] from an objc.ID.
//
// An annotation in a PDF document.
func PDFAnnotationFromID(id objc.ID) PDFAnnotation {
	return PDFAnnotation{objectivec.Object{ID: id}}
}

// NOTE: PDFAnnotation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFAnnotation] class.
//
// # Creating an Annotation
//
//   - [IPDFAnnotation.InitWithBoundsForTypeWithProperties]: Creates a PDF annotation with the specified bounds, type, and optional properties.
//
// # Accessing Information About an Annotation
//
//   - [IPDFAnnotation.Page]: Returns the page that the annotation is associated with.
//   - [IPDFAnnotation.SetPage]
//   - [IPDFAnnotation.ModificationDate]: Returns the modification date of the annotation.
//   - [IPDFAnnotation.SetModificationDate]
//   - [IPDFAnnotation.UserName]: Returns the name of the user who created the annotation.
//   - [IPDFAnnotation.SetUserName]
//   - [IPDFAnnotation.Type]: Returns the type of the annotation.
//   - [IPDFAnnotation.SetType]
//   - [IPDFAnnotation.Action]: An object that represents an action for a PDF element, such as a link annotation.
//   - [IPDFAnnotation.SetAction]
//
// # Managing Annotation Drawing and Output
//
//   - [IPDFAnnotation.DrawWithBoxInContext]: Draws the annotation in a graphics context using page-space coordinates relative to the origin of the specified box.
//   - [IPDFAnnotation.ShouldDisplay]: Returns a Boolean value indicating whether the annotation should be displayed.
//   - [IPDFAnnotation.SetShouldDisplay]
//   - [IPDFAnnotation.ShouldPrint]: Returns a Boolean value indicating whether the annotation should appear when the document is printed.
//   - [IPDFAnnotation.SetShouldPrint]
//
// # Modifying Annotation Attributes
//
//   - [IPDFAnnotation.AnnotationKeyValues]: A dictionary that contains a deep copy of the widget’s properties.
//   - [IPDFAnnotation.ValueForAnnotationKey]: Returns a deep copy of the key-value pairs of properties for the specified key.
//   - [IPDFAnnotation.SetValueForAnnotationKey]: Sets a value in the annotation’s dictionary.
//   - [IPDFAnnotation.SetBooleanForAnnotationKey]: Sets a Boolean value in the annotation’s dictionary.
//   - [IPDFAnnotation.SetRectForAnnotationKey]: Sets a rectangle value in the annotation’s dictionary.
//   - [IPDFAnnotation.RemoveValueForAnnotationKey]: Removes a value from the annotation’s dictionary.
//
// # Managing Annotation Display Characteristics
//
//   - [IPDFAnnotation.Alignment]: The alignment of the free text and text widget annotation’s text content.
//   - [IPDFAnnotation.SetAlignment]
//   - [IPDFAnnotation.Bounds]: Returns the bounding box for the annotation in page space.
//   - [IPDFAnnotation.SetBounds]
//   - [IPDFAnnotation.Contents]: Returns the textual content (if any) associated with the annotation.
//   - [IPDFAnnotation.SetContents]
//   - [IPDFAnnotation.Font]: The font the annotation uses to display text.
//   - [IPDFAnnotation.SetFont]
//   - [IPDFAnnotation.FontColor]: The font color the annotation uses to display text.
//   - [IPDFAnnotation.SetFontColor]
//   - [IPDFAnnotation.Border]: Sets the border style for the annotation.
//   - [IPDFAnnotation.SetBorder]
//   - [IPDFAnnotation.Highlighted]: A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//   - [IPDFAnnotation.SetHighlighted]
//   - [IPDFAnnotation.Color]: Sets the stroke color for the annotation.
//   - [IPDFAnnotation.SetColor]
//   - [IPDFAnnotation.HasAppearanceStream]: Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// # Configuring Shape Annotations
//
//   - [IPDFAnnotation.InteriorColor]: The fill color for drawing a circle, line, or square annotation.
//   - [IPDFAnnotation.SetInteriorColor]
//
// # Configuring Line Annotations
//
//   - [IPDFAnnotation.StartPoint]: The point where a line begins, in annotation-space coordinates.
//   - [IPDFAnnotation.SetStartPoint]
//   - [IPDFAnnotation.EndPoint]: The point where a line ends, in annotation-space coordinates.
//   - [IPDFAnnotation.SetEndPoint]
//   - [IPDFAnnotation.StartLineStyle]: The style of the line annotation’s starting point, such as square or filled arrowhead.
//   - [IPDFAnnotation.SetStartLineStyle]
//   - [IPDFAnnotation.EndLineStyle]: The style of the line annotation’s ending point, such as square or filled arrowhead.
//   - [IPDFAnnotation.SetEndLineStyle]
//
// # Configuring Link Annotations
//
//   - [IPDFAnnotation.Destination]: The destination for a link annotation.
//   - [IPDFAnnotation.SetDestination]
//   - [IPDFAnnotation.URL]: A URL for a link annotation.
//   - [IPDFAnnotation.SetURL]
//
// # Configuring Text Annotations
//
//   - [IPDFAnnotation.IconType]: The type of icon to display for a pop-up text annotation.
//   - [IPDFAnnotation.SetIconType]
//
// # Configuring Pop-Up Annotations
//
//   - [IPDFAnnotation.Popup]: Returns the pop-up annotation associated with an annotation.
//   - [IPDFAnnotation.SetPopup]
//   - [IPDFAnnotation.Open]: A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//   - [IPDFAnnotation.SetOpen]
//
// # Configuring Text Markup Annotations
//
//   - [IPDFAnnotation.MarkupType]: The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
//   - [IPDFAnnotation.SetMarkupType]
//   - [IPDFAnnotation.QuadrilateralPoints]: An array of values that represents the points bounding the marked-up text.
//   - [IPDFAnnotation.SetQuadrilateralPoints]
//
// # Configuring Widget Annotations
//
//   - [IPDFAnnotation.WidgetFieldType]: The type of widget annotation, such as button, choice, or text.
//   - [IPDFAnnotation.SetWidgetFieldType]
//   - [IPDFAnnotation.WidgetStringValue]: The string value of the widget annotation.
//   - [IPDFAnnotation.SetWidgetStringValue]
//   - [IPDFAnnotation.WidgetDefaultStringValue]: The string value that the widget reverts to when performing a reset form action.
//   - [IPDFAnnotation.SetWidgetDefaultStringValue]
//   - [IPDFAnnotation.FieldName]: The widget identifier for form annotation actions and behaviors.
//   - [IPDFAnnotation.SetFieldName]
//   - [IPDFAnnotation.BackgroundColor]: The color of the widget’s background.
//   - [IPDFAnnotation.SetBackgroundColor]
//   - [IPDFAnnotation.ReadOnly]: A Boolean value that determines whether the widget is editable.
//   - [IPDFAnnotation.SetReadOnly]
//
// # Configuring Text Widget Annotations
//
//   - [IPDFAnnotation.Multiline]: A Boolean value that indicates whether the text widget annotation displays multiple lines.
//   - [IPDFAnnotation.SetMultiline]
//   - [IPDFAnnotation.IsPasswordField]: A Boolean value that indicates whether the text widget annotation displays a password field using bullet characters.
//   - [IPDFAnnotation.MaximumLength]: The maximum number of characters the text widget annotation allows.
//   - [IPDFAnnotation.SetMaximumLength]
//   - [IPDFAnnotation.Comb]: A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//   - [IPDFAnnotation.SetComb]
//
// # Configuring Button Widget Annotations
//
//   - [IPDFAnnotation.WidgetControlType]: The type of button widget control, either radio button, push button, or checkbox.
//   - [IPDFAnnotation.SetWidgetControlType]
//   - [IPDFAnnotation.ButtonWidgetState]: The current state of the button widget annotation.
//   - [IPDFAnnotation.SetButtonWidgetState]
//   - [IPDFAnnotation.ButtonWidgetStateString]: A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
//   - [IPDFAnnotation.SetButtonWidgetStateString]
//   - [IPDFAnnotation.Caption]: The title of push button widget annotations.
//   - [IPDFAnnotation.SetCaption]
//   - [IPDFAnnotation.AllowsToggleToOff]: A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
//   - [IPDFAnnotation.SetAllowsToggleToOff]
//   - [IPDFAnnotation.RadiosInUnison]: A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
//   - [IPDFAnnotation.SetRadiosInUnison]
//
// # Configuring Choice Widget Annotations
//
//   - [IPDFAnnotation.Choices]: An array of strings that specifies the options in either a list or a pop-up menu.
//   - [IPDFAnnotation.SetChoices]
//   - [IPDFAnnotation.ListChoice]: A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//   - [IPDFAnnotation.SetListChoice]
//   - [IPDFAnnotation.Values]: An array of strings that specifies the export values for items in a list or a pop-up menu.
//   - [IPDFAnnotation.SetValues]
//
// # Configuring Ink Annotations
//
//   - [IPDFAnnotation.Paths]: An array of bezier paths, in annotation-space coordinates, that compose the annotation.
//   - [IPDFAnnotation.AddBezierPath]: Adds a bezier path to the ink annotation.
//   - [IPDFAnnotation.RemoveBezierPath]: Removes a bezier path from an ink annotation.
//
// # Configuring Stamp Annotations
//
//   - [IPDFAnnotation.StampName]: The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
//   - [IPDFAnnotation.SetStampName]
//
// # Instance Properties
//
//   - [IPDFAnnotation.ActivatableTextField]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation
type IPDFAnnotation interface {
	objectivec.IObject

	// Topic: Creating an Annotation

	// Creates a PDF annotation with the specified bounds, type, and optional properties.
	InitWithBoundsForTypeWithProperties(bounds corefoundation.CGRect, annotationType PDFAnnotationSubtype, properties foundation.INSDictionary) PDFAnnotation

	// Topic: Accessing Information About an Annotation

	// Returns the page that the annotation is associated with.
	Page() IPDFPage
	SetPage(value IPDFPage)
	// Returns the modification date of the annotation.
	ModificationDate() foundation.INSDate
	SetModificationDate(value foundation.INSDate)
	// Returns the name of the user who created the annotation.
	UserName() string
	SetUserName(value string)
	// Returns the type of the annotation.
	Type() string
	SetType(value string)
	// An object that represents an action for a PDF element, such as a link annotation.
	Action() IPDFAction
	SetAction(value IPDFAction)

	// Topic: Managing Annotation Drawing and Output

	// Draws the annotation in a graphics context using page-space coordinates relative to the origin of the specified box.
	DrawWithBoxInContext(box PDFDisplayBox, context coregraphics.CGContextRef)
	// Returns a Boolean value indicating whether the annotation should be displayed.
	ShouldDisplay() bool
	SetShouldDisplay(value bool)
	// Returns a Boolean value indicating whether the annotation should appear when the document is printed.
	ShouldPrint() bool
	SetShouldPrint(value bool)

	// Topic: Modifying Annotation Attributes

	// A dictionary that contains a deep copy of the widget’s properties.
	AnnotationKeyValues() foundation.INSDictionary
	// Returns a deep copy of the key-value pairs of properties for the specified key.
	ValueForAnnotationKey(key PDFAnnotationKey) objectivec.IObject
	// Sets a value in the annotation’s dictionary.
	SetValueForAnnotationKey(value objectivec.IObject, key PDFAnnotationKey) bool
	// Sets a Boolean value in the annotation’s dictionary.
	SetBooleanForAnnotationKey(value bool, key PDFAnnotationKey) bool
	// Sets a rectangle value in the annotation’s dictionary.
	SetRectForAnnotationKey(value corefoundation.CGRect, key PDFAnnotationKey) bool
	// Removes a value from the annotation’s dictionary.
	RemoveValueForAnnotationKey(key PDFAnnotationKey)

	// Topic: Managing Annotation Display Characteristics

	// The alignment of the free text and text widget annotation’s text content.
	Alignment() appkit.NSTextAlignment
	SetAlignment(value appkit.NSTextAlignment)
	// Returns the bounding box for the annotation in page space.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	// Returns the textual content (if any) associated with the annotation.
	Contents() string
	SetContents(value string)
	// The font the annotation uses to display text.
	Font() appkit.NSFont
	SetFont(value appkit.NSFont)
	// The font color the annotation uses to display text.
	FontColor() appkit.NSColor
	SetFontColor(value appkit.NSColor)
	// Sets the border style for the annotation.
	Border() IPDFBorder
	SetBorder(value IPDFBorder)
	// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
	Highlighted() bool
	SetHighlighted(value bool)
	// Sets the stroke color for the annotation.
	Color() appkit.NSColor
	SetColor(value appkit.NSColor)
	// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
	HasAppearanceStream() bool

	// Topic: Configuring Shape Annotations

	// The fill color for drawing a circle, line, or square annotation.
	InteriorColor() appkit.NSColor
	SetInteriorColor(value appkit.NSColor)

	// Topic: Configuring Line Annotations

	// The point where a line begins, in annotation-space coordinates.
	StartPoint() corefoundation.CGPoint
	SetStartPoint(value corefoundation.CGPoint)
	// The point where a line ends, in annotation-space coordinates.
	EndPoint() corefoundation.CGPoint
	SetEndPoint(value corefoundation.CGPoint)
	// The style of the line annotation’s starting point, such as square or filled arrowhead.
	StartLineStyle() PDFLineStyle
	SetStartLineStyle(value PDFLineStyle)
	// The style of the line annotation’s ending point, such as square or filled arrowhead.
	EndLineStyle() PDFLineStyle
	SetEndLineStyle(value PDFLineStyle)

	// Topic: Configuring Link Annotations

	// The destination for a link annotation.
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)
	// A URL for a link annotation.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)

	// Topic: Configuring Text Annotations

	// The type of icon to display for a pop-up text annotation.
	IconType() PDFTextAnnotationIconType
	SetIconType(value PDFTextAnnotationIconType)

	// Topic: Configuring Pop-Up Annotations

	// Returns the pop-up annotation associated with an annotation.
	Popup() IPDFAnnotation
	SetPopup(value IPDFAnnotation)
	// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
	Open() bool
	SetOpen(value bool)

	// Topic: Configuring Text Markup Annotations

	// The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
	MarkupType() PDFMarkupType
	SetMarkupType(value PDFMarkupType)
	// An array of values that represents the points bounding the marked-up text.
	QuadrilateralPoints() []foundation.NSValue
	SetQuadrilateralPoints(value []foundation.NSValue)

	// Topic: Configuring Widget Annotations

	// The type of widget annotation, such as button, choice, or text.
	WidgetFieldType() PDFAnnotationWidgetSubtype
	SetWidgetFieldType(value PDFAnnotationWidgetSubtype)
	// The string value of the widget annotation.
	WidgetStringValue() string
	SetWidgetStringValue(value string)
	// The string value that the widget reverts to when performing a reset form action.
	WidgetDefaultStringValue() string
	SetWidgetDefaultStringValue(value string)
	// The widget identifier for form annotation actions and behaviors.
	FieldName() string
	SetFieldName(value string)
	// The color of the widget’s background.
	BackgroundColor() appkit.NSColor
	SetBackgroundColor(value appkit.NSColor)
	// A Boolean value that determines whether the widget is editable.
	ReadOnly() bool
	SetReadOnly(value bool)

	// Topic: Configuring Text Widget Annotations

	// A Boolean value that indicates whether the text widget annotation displays multiple lines.
	Multiline() bool
	SetMultiline(value bool)
	// A Boolean value that indicates whether the text widget annotation displays a password field using bullet characters.
	IsPasswordField() bool
	// The maximum number of characters the text widget annotation allows.
	MaximumLength() int
	SetMaximumLength(value int)
	// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
	Comb() bool
	SetComb(value bool)

	// Topic: Configuring Button Widget Annotations

	// The type of button widget control, either radio button, push button, or checkbox.
	WidgetControlType() PDFWidgetControlType
	SetWidgetControlType(value PDFWidgetControlType)
	// The current state of the button widget annotation.
	ButtonWidgetState() PDFWidgetCellState
	SetButtonWidgetState(value PDFWidgetCellState)
	// A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
	ButtonWidgetStateString() string
	SetButtonWidgetStateString(value string)
	// The title of push button widget annotations.
	Caption() string
	SetCaption(value string)
	// A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
	AllowsToggleToOff() bool
	SetAllowsToggleToOff(value bool)
	// A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
	RadiosInUnison() bool
	SetRadiosInUnison(value bool)

	// Topic: Configuring Choice Widget Annotations

	// An array of strings that specifies the options in either a list or a pop-up menu.
	Choices() []string
	SetChoices(value []string)
	// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
	ListChoice() bool
	SetListChoice(value bool)
	// An array of strings that specifies the export values for items in a list or a pop-up menu.
	Values() []string
	SetValues(value []string)

	// Topic: Configuring Ink Annotations

	// An array of bezier paths, in annotation-space coordinates, that compose the annotation.
	Paths() []appkit.NSBezierPath
	// Adds a bezier path to the ink annotation.
	AddBezierPath(path appkit.NSBezierPath)
	// Removes a bezier path from an ink annotation.
	RemoveBezierPath(path appkit.NSBezierPath)

	// Topic: Configuring Stamp Annotations

	// The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
	StampName() string
	SetStampName(value string)

	// Topic: Instance Properties

	ActivatableTextField() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p PDFAnnotation) Init() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFAnnotation) Autorelease() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotation creates a new PDFAnnotation instance.
func NewPDFAnnotation() PDFAnnotation {
	class := getPDFAnnotationClass()
	rv := objc.Send[PDFAnnotation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a PDF annotation with the specified bounds, type, and optional
// properties.
//
// bounds: The bounding box of the annotation, in page-space coordinates.
//
// annotationType: The subtype of the annotation, such as text, link, or line.
//
// properties: A dictionary that contains properties of the annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(bounds:forType:withProperties:)
func NewPDFAnnotationWithBoundsForTypeWithProperties(bounds corefoundation.CGRect, annotationType PDFAnnotationSubtype, properties foundation.INSDictionary) PDFAnnotation {
	instance := getPDFAnnotationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBounds:forType:withProperties:"), bounds, annotationType, properties)
	return PDFAnnotationFromID(rv)
}

// Creates a PDF annotation with the specified bounds, type, and optional
// properties.
//
// bounds: The bounding box of the annotation, in page-space coordinates.
//
// annotationType: The subtype of the annotation, such as text, link, or line.
//
// properties: A dictionary that contains properties of the annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(bounds:forType:withProperties:)
func (p PDFAnnotation) InitWithBoundsForTypeWithProperties(bounds corefoundation.CGRect, annotationType PDFAnnotationSubtype, properties foundation.INSDictionary) PDFAnnotation {
	rv := objc.Send[PDFAnnotation](p.ID, objc.Sel("initWithBounds:forType:withProperties:"), bounds, annotationType, properties)
	return rv
}

// Draws the annotation in a graphics context using page-space coordinates
// relative to the origin of the specified box.
//
// box: The display box that represents the rectangle to draw the annotation in, in
// page-space coordinates.
//
// context: The graphics context to draw the annotation in.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/draw(with:in:)
func (p PDFAnnotation) DrawWithBoxInContext(box PDFDisplayBox, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawWithBox:inContext:"), box, context)
}

// Returns a deep copy of the key-value pairs of properties for the specified
// key.
//
// key: A [PDFAnnotationKey] or appropriate string from the Adobe PDF
// Specification.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/value(forAnnotationKey:)
func (p PDFAnnotation) ValueForAnnotationKey(key PDFAnnotationKey) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("valueForAnnotationKey:"), key)
	return objectivec.Object{ID: rv}
}

// Sets a value in the annotation’s dictionary.
//
// value: The value to set in the attribute’s dictionary.
//
// key: A [PDFAnnotationKey] or appropriate string from the Adobe PDF
// Specification.
//
// # Return Value
//
// true if the value sets successfully; otherwise, false.
//
// # Discussion
//
// Some keys expect a complex type. For example, the [color] key expects an
// array of zero, one, two, three, or four elements, where each element is a
// floating-point number from `0.0` to `1.0`. As a convenience, this key
// accepts an [NSColor] or [UIColor] value. For details about other
// conveniences, see the individual [PDFAnnotationKey] properties or the
// `PDFAnnotationUtilities.H()` header file.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setValue(_:forAnnotationKey:)
//
// [NSColor]: https://developer.apple.com/documentation/AppKit/NSColor
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
// [color]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationKey/color
func (p PDFAnnotation) SetValueForAnnotationKey(value objectivec.IObject, key PDFAnnotationKey) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setValue:forAnnotationKey:"), value, key)
	return rv
}

// Sets a Boolean value in the annotation’s dictionary.
//
// value: The Boolean value to set in the annotation’s dictionary.
//
// key: A [PDFAnnotationKey] or appropriate string from the Adobe PDF
// Specification.
//
// # Return Value
//
// true if the value sets successfully; otherwise, false.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setBoolean(_:forAnnotationKey:)
func (p PDFAnnotation) SetBooleanForAnnotationKey(value bool, key PDFAnnotationKey) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setBoolean:forAnnotationKey:"), value, key)
	return rv
}

// Sets a rectangle value in the annotation’s dictionary.
//
// value: The rectangle value to set in the annotation’s dictionary.
//
// key: A [PDFAnnotationKey] or appropriate string from the Adobe PDF
// Specification.
//
// # Return Value
//
// true if the value sets successfully; otherwise, false.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setRect(_:forAnnotationKey:)
func (p PDFAnnotation) SetRectForAnnotationKey(value corefoundation.CGRect, key PDFAnnotationKey) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setRect:forAnnotationKey:"), value, key)
	return rv
}

// Removes a value from the annotation’s dictionary.
//
// key: A [PDFAnnotationKey] or appropriate string from the Adobe PDF
// Specification.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/removeValue(forAnnotationKey:)
func (p PDFAnnotation) RemoveValueForAnnotationKey(key PDFAnnotationKey) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeValueForAnnotationKey:"), key)
}

// Adds a bezier path to the ink annotation.
//
// path: The bezier path to add, in annotation-space coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/add(_:)
func (p PDFAnnotation) AddBezierPath(path appkit.NSBezierPath) {
	objc.Send[objc.ID](p.ID, objc.Sel("addBezierPath:"), path)
}

// Removes a bezier path from an ink annotation.
//
// path: The bezier path to remove, in annotation-space coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/remove(_:)
func (p PDFAnnotation) RemoveBezierPath(path appkit.NSBezierPath) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeBezierPath:"), path)
}
func (p PDFAnnotation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a line style that corresponds to the specified name.
//
// name: The name of a line style, which matches the definition in the Adobe PDF
// Specification.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/lineStyle(fromName:)
func (_PDFAnnotationClass PDFAnnotationClass) LineStyleFromName(name string) PDFLineStyle {
	rv := objc.Send[PDFLineStyle](objc.ID(_PDFAnnotationClass.class), objc.Sel("lineStyleFromName:"), objc.String(name))
	return PDFLineStyle(rv)
}

// Returns the name of the line style, which matches the definition in the
// Adobe PDF Specification.
//
// style: A line style to get a name for.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/name(for:)
func (_PDFAnnotationClass PDFAnnotationClass) NameForLineStyle(style PDFLineStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_PDFAnnotationClass.class), objc.Sel("nameForLineStyle:"), style)
	return foundation.NSStringFromID(rv).String()
}

// Returns the page that the annotation is associated with.
//
// # Return Value
//
// The PDF page associated with the annotation.
//
// # Discussion
//
// The [AddAnnotation] method in the [PDFPage] class lets you associate an
// annotation with a page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/page
func (p PDFAnnotation) Page() IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("page"))
	return PDFPageFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetPage(value IPDFPage) {
	objc.Send[struct{}](p.ID, objc.Sel("setPage:"), value)
}

// Returns the modification date of the annotation.
//
// # Return Value
//
// The modification date of the annotation, or [NULL] if there is no
// modification date.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/modificationDate
func (p PDFAnnotation) ModificationDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetModificationDate(value foundation.INSDate) {
	objc.Send[struct{}](p.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the name of the user who created the annotation.
//
// # Return Value
//
// The name of the user who created the annotation, or NULL if no user name is
// set.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/userName
func (p PDFAnnotation) UserName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetUserName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setUserName:"), objc.String(value))
}

// Returns the type of the annotation.
//
// # Return Value
//
// The type of the annotation. Types include [Line], [Link], [Text], and so
// on, referring to the [PDFAnnotation] subclasses. In the Adobe PDF
// Specification, this attribute is called [Subtype], and the common
// “type” for all annotations in the PDF Specification is [Annot].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/type
func (p PDFAnnotation) Type() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetType(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setType:"), objc.String(value))
}

// An object that represents an action for a PDF element, such as a link
// annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/action
func (p PDFAnnotation) Action() IPDFAction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("action"))
	return PDFActionFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetAction(value IPDFAction) {
	objc.Send[struct{}](p.ID, objc.Sel("setAction:"), value)
}

// Returns a Boolean value indicating whether the annotation should be
// displayed.
//
// # Return Value
//
// true if the annotation should be displayed; otherwise false.
//
// # Discussion
//
// [PDFPage] respects this flag when drawing.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldDisplay
func (p PDFAnnotation) ShouldDisplay() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldDisplay"))
	return rv
}
func (p PDFAnnotation) SetShouldDisplay(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldDisplay:"), value)
}

// Returns a Boolean value indicating whether the annotation should appear
// when the document is printed.
//
// # Return Value
//
// true if the annotation should appear when the PDF document is printed;
// otherwise false.
//
// # Discussion
//
// [PDFPage] respects this flag when printing.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldPrint
func (p PDFAnnotation) ShouldPrint() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldPrint"))
	return rv
}
func (p PDFAnnotation) SetShouldPrint(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldPrint:"), value)
}

// A dictionary that contains a deep copy of the widget’s properties.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/annotationKeyValues
func (p PDFAnnotation) AnnotationKeyValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("annotationKeyValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The alignment of the free text and text widget annotation’s text content.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/alignment
func (p PDFAnnotation) Alignment() appkit.NSTextAlignment {
	rv := objc.Send[appkit.NSTextAlignment](p.ID, objc.Sel("alignment"))
	return appkit.NSTextAlignment(rv)
}
func (p PDFAnnotation) SetAlignment(value appkit.NSTextAlignment) {
	objc.Send[struct{}](p.ID, objc.Sel("setAlignment:"), value)
}

// Returns the bounding box for the annotation in page space.
//
// # Return Value
//
// The bounding box for the annotation in page space.
//
// # Discussion
//
// Page space is a 72-dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/bounds
func (p PDFAnnotation) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (p PDFAnnotation) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](p.ID, objc.Sel("setBounds:"), value)
}

// Returns the textual content (if any) associated with the annotation.
//
// # Return Value
//
// A string representing the textual content associated with the annotation.
//
// # Discussion
//
// Textual content is typically associated with [PDFAnnotationText] and
// [PDFAnnotationFreeText] annotations.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/contents
func (p PDFAnnotation) Contents() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contents"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetContents(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setContents:"), objc.String(value))
}

// The font the annotation uses to display text.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/font
func (p PDFAnnotation) Font() appkit.NSFont {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("font"))
	return appkit.NSFontFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetFont(value appkit.NSFont) {
	objc.Send[struct{}](p.ID, objc.Sel("setFont:"), value)
}

// The font color the annotation uses to display text.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fontColor
func (p PDFAnnotation) FontColor() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fontColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetFontColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setFontColor:"), value)
}

// Sets the border style for the annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/border
func (p PDFAnnotation) Border() IPDFBorder {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("border"))
	return PDFBorderFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetBorder(value IPDFBorder) {
	objc.Send[struct{}](p.ID, objc.Sel("setBorder:"), value)
}

// A Boolean value that indicates whether the annotation is in a highlighted
// state, such as when the mouse is down on a link annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isHighlighted
func (p PDFAnnotation) Highlighted() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isHighlighted"))
	return rv
}
func (p PDFAnnotation) SetHighlighted(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHighlighted:"), value)
}

// Sets the stroke color for the annotation.
//
// # Discussion
//
// Where this color is used depends on the annotation type.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/color
func (p PDFAnnotation) Color() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("color"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setColor:"), value)
}

// Returns a Boolean value that indicates whether the annotation has an
// appearance stream associated with it.
//
// # Return Value
//
// true if the annotation has an appearance stream; otherwise false.
//
// # Discussion
//
// An appearance stream is a sequence of draw instructions used to render a
// PDF item. If an appearance stream exists, PDF Kit draws the annotation
// using the stream, which may override existing set parameters (such as the
// stroke color set with `setColor`).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasAppearanceStream
func (p PDFAnnotation) HasAppearanceStream() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasAppearanceStream"))
	return rv
}

// The fill color for drawing a circle, line, or square annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/interiorColor
func (p PDFAnnotation) InteriorColor() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("interiorColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetInteriorColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setInteriorColor:"), value)
}

// The point where a line begins, in annotation-space coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startPoint
func (p PDFAnnotation) StartPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("startPoint"))
	return corefoundation.CGPoint(rv)
}
func (p PDFAnnotation) SetStartPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](p.ID, objc.Sel("setStartPoint:"), value)
}

// The point where a line ends, in annotation-space coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endPoint
func (p PDFAnnotation) EndPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("endPoint"))
	return corefoundation.CGPoint(rv)
}
func (p PDFAnnotation) SetEndPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](p.ID, objc.Sel("setEndPoint:"), value)
}

// The style of the line annotation’s starting point, such as square or
// filled arrowhead.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startLineStyle
func (p PDFAnnotation) StartLineStyle() PDFLineStyle {
	rv := objc.Send[PDFLineStyle](p.ID, objc.Sel("startLineStyle"))
	return PDFLineStyle(rv)
}
func (p PDFAnnotation) SetStartLineStyle(value PDFLineStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setStartLineStyle:"), value)
}

// The style of the line annotation’s ending point, such as square or filled
// arrowhead.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endLineStyle
func (p PDFAnnotation) EndLineStyle() PDFLineStyle {
	rv := objc.Send[PDFLineStyle](p.ID, objc.Sel("endLineStyle"))
	return PDFLineStyle(rv)
}
func (p PDFAnnotation) SetEndLineStyle(value PDFLineStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setEndLineStyle:"), value)
}

// The destination for a link annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/destination
func (p PDFAnnotation) Destination() IPDFDestination {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destination"))
	return PDFDestinationFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetDestination(value IPDFDestination) {
	objc.Send[struct{}](p.ID, objc.Sel("setDestination:"), value)
}

// A URL for a link annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/url
func (p PDFAnnotation) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

// The type of icon to display for a pop-up text annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/iconType
func (p PDFAnnotation) IconType() PDFTextAnnotationIconType {
	rv := objc.Send[PDFTextAnnotationIconType](p.ID, objc.Sel("iconType"))
	return PDFTextAnnotationIconType(rv)
}
func (p PDFAnnotation) SetIconType(value PDFTextAnnotationIconType) {
	objc.Send[struct{}](p.ID, objc.Sel("setIconType:"), value)
}

// Returns the pop-up annotation associated with an annotation.
//
// # Return Value
//
// The pop-up annotation associated with the annotation, or [NULL] if no
// pop-up exists.
//
// # Discussion
//
// Pop-up annotations are not used with links or widgets. The bounds and open
// state of the pop-up annotation indicate the placement and open state of the
// pop-up window.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/popup
func (p PDFAnnotation) Popup() IPDFAnnotation {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("popup"))
	return PDFAnnotationFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetPopup(value IPDFAnnotation) {
	objc.Send[struct{}](p.ID, objc.Sel("setPopup:"), value)
}

// A Boolean value that indicates whether the pop-up annotation is in an
// opened state, displaying its text content, or in a closed state, displaying
// an icon.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isOpen
func (p PDFAnnotation) Open() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isOpen"))
	return rv
}
func (p PDFAnnotation) SetOpen(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setOpen:"), value)
}

// The markup type that the annotation displays, either highlight,
// strikethrough, underline, or redact.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/markupType
func (p PDFAnnotation) MarkupType() PDFMarkupType {
	rv := objc.Send[PDFMarkupType](p.ID, objc.Sel("markupType"))
	return PDFMarkupType(rv)
}
func (p PDFAnnotation) SetMarkupType(value PDFMarkupType) {
	objc.Send[struct{}](p.ID, objc.Sel("setMarkupType:"), value)
}

// An array of values that represents the points bounding the marked-up text.
//
// # Discussion
//
// The array contains `N * 4` [NSValue] objects that use [pointValue] or
// [cgPointValue] to define [N] quadrilaterals in page-space coordinates,
// where [N] is the number of quad points. The order of the points is a Z
// pattern as follows:
//
// - Upper-left point - Upper-right point - Lower-left point - Lower-right
// point
//
// The coordinates of each point are relative to the bound’s origin of the
// annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/quadrilateralPoints
//
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
// [cgPointValue]: https://developer.apple.com/documentation/Foundation/NSValue/cgPointValue
// [pointValue]: https://developer.apple.com/documentation/Foundation/NSValue/pointValue
func (p PDFAnnotation) QuadrilateralPoints() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("quadrilateralPoints"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
func (p PDFAnnotation) SetQuadrilateralPoints(value []foundation.NSValue) {
	objc.Send[struct{}](p.ID, objc.Sel("setQuadrilateralPoints:"), objectivec.IObjectSliceToNSArray(value))
}

// The type of widget annotation, such as button, choice, or text.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetFieldType
func (p PDFAnnotation) WidgetFieldType() PDFAnnotationWidgetSubtype {
	rv := objc.Send[PDFAnnotationWidgetSubtype](p.ID, objc.Sel("widgetFieldType"))
	return PDFAnnotationWidgetSubtype(rv)
}
func (p PDFAnnotation) SetWidgetFieldType(value PDFAnnotationWidgetSubtype) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetFieldType:"), value)
}

// The string value of the widget annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetStringValue
func (p PDFAnnotation) WidgetStringValue() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("widgetStringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetWidgetStringValue(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetStringValue:"), objc.String(value))
}

// The string value that the widget reverts to when performing a reset form
// action.
//
// # Discussion
//
// For text and choice widgets, this property specifies the string value that
// a reset form action sets on a widget.
//
// For radio buttons and checkboxes, set this property to [Off] if the desired
// default is for the button to be in an unselected state; otherwise, set it
// to the [ButtonWidgetStateString].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetDefaultStringValue
func (p PDFAnnotation) WidgetDefaultStringValue() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("widgetDefaultStringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetWidgetDefaultStringValue(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetDefaultStringValue:"), objc.String(value))
}

// The widget identifier for form annotation actions and behaviors.
//
// # Discussion
//
// Use this identifier to refer to the widget when using a
// [PDFActionResetForm] action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fieldName
func (p PDFAnnotation) FieldName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fieldName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetFieldName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setFieldName:"), objc.String(value))
}

// The color of the widget’s background.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/backgroundColor
func (p PDFAnnotation) BackgroundColor() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFAnnotation) SetBackgroundColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value that determines whether the widget is editable.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isReadOnly
func (p PDFAnnotation) ReadOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isReadOnly"))
	return rv
}
func (p PDFAnnotation) SetReadOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setReadOnly:"), value)
}

// A Boolean value that indicates whether the text widget annotation displays
// multiple lines.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isMultiline
func (p PDFAnnotation) Multiline() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isMultiline"))
	return rv
}
func (p PDFAnnotation) SetMultiline(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setMultiline:"), value)
}

// A Boolean value that indicates whether the text widget annotation displays
// a password field using bullet characters.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isPasswordField
func (p PDFAnnotation) IsPasswordField() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPasswordField"))
	return rv
}

// The maximum number of characters the text widget annotation allows.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/maximumLength
func (p PDFAnnotation) MaximumLength() int {
	rv := objc.Send[int](p.ID, objc.Sel("maximumLength"))
	return rv
}
func (p PDFAnnotation) SetMaximumLength(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaximumLength:"), value)
}

// A Boolean value that indicates whether the annotation divides the text
// widget’s bounds into equally spaced segments, such as in a form entry
// field.
//
// # Discussion
//
// The [MaximumLength] property specifies the number of spaces the text widget
// divides the bounds into.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasComb
func (p PDFAnnotation) Comb() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasComb"))
	return rv
}
func (p PDFAnnotation) SetComb(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setComb:"), value)
}

// The type of button widget control, either radio button, push button, or
// checkbox.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetControlType
func (p PDFAnnotation) WidgetControlType() PDFWidgetControlType {
	rv := objc.Send[PDFWidgetControlType](p.ID, objc.Sel("widgetControlType"))
	return PDFWidgetControlType(rv)
}
func (p PDFAnnotation) SetWidgetControlType(value PDFWidgetControlType) {
	objc.Send[struct{}](p.ID, objc.Sel("setWidgetControlType:"), value)
}

// The current state of the button widget annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetState
func (p PDFAnnotation) ButtonWidgetState() PDFWidgetCellState {
	rv := objc.Send[PDFWidgetCellState](p.ID, objc.Sel("buttonWidgetState"))
	return PDFWidgetCellState(rv)
}
func (p PDFAnnotation) SetButtonWidgetState(value PDFWidgetCellState) {
	objc.Send[struct{}](p.ID, objc.Sel("setButtonWidgetState:"), value)
}

// A string value that differentiates button widgets in the same group, such
// as to identify mutually exclusive radio buttons from each other.
//
// # Discussion
//
// The default value is [Yes].
//
// To group button widgets, set the same [FieldName] on the button widgets.
// The [ButtonWidgetStateString] property allows you to identify individual
// button widgets in that group.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetStateString
func (p PDFAnnotation) ButtonWidgetStateString() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("buttonWidgetStateString"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetButtonWidgetStateString(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setButtonWidgetStateString:"), objc.String(value))
}

// The title of push button widget annotations.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/caption
func (p PDFAnnotation) Caption() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("caption"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetCaption(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setCaption:"), objc.String(value))
}

// A Boolean value that indicates whether clicking or tapping a selected radio
// button toggles it to an unselected state.
//
// # Discussion
//
// To implement a group of radio buttons where at least one option must remain
// in a selected state, set [AllowsToggleToOff] to false on each button in the
// group.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/allowsToggleToOff
func (p PDFAnnotation) AllowsToggleToOff() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsToggleToOff"))
	return rv
}
func (p PDFAnnotation) SetAllowsToggleToOff(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsToggleToOff:"), value)
}

// A Boolean value that indicates whether radio buttons in a group turn on and
// off in unison.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/radiosInUnison
func (p PDFAnnotation) RadiosInUnison() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("radiosInUnison"))
	return rv
}
func (p PDFAnnotation) SetRadiosInUnison(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setRadiosInUnison:"), value)
}

// An array of strings that specifies the options in either a list or a pop-up
// menu.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/choices
func (p PDFAnnotation) Choices() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("choices"))
	return objc.ConvertSliceToStrings(rv)
}
func (p PDFAnnotation) SetChoices(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setChoices:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean value that indicates whether the choice widget annotation is a
// list or a pop-up menu.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isListChoice
func (p PDFAnnotation) ListChoice() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isListChoice"))
	return rv
}
func (p PDFAnnotation) SetListChoice(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setListChoice:"), value)
}

// An array of strings that specifies the export values for items in a list or
// a pop-up menu.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/values
func (p PDFAnnotation) Values() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("values"))
	return objc.ConvertSliceToStrings(rv)
}
func (p PDFAnnotation) SetValues(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setValues:"), objectivec.StringSliceToNSArray(value))
}

// An array of bezier paths, in annotation-space coordinates, that compose the
// annotation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/paths
func (p PDFAnnotation) Paths() []appkit.NSBezierPath {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("paths"))
	return objc.ConvertSlice(rv, func(id objc.ID) appkit.NSBezierPath {
		return appkit.NSBezierPathFromID(id)
	})
}

// The name of the stamp, a text or graphics annotation that emulates a rubber
// stamp effect.
//
// # Discussion
//
// The default value for this property is [Default]. Additional possible
// values for `stampName` are: [Approved], [Asis], [Confidential],
// [Departmental], [Experimental], [Expired], [Final], [ForComment],
// [ForPublicRelease], [NotApproved], [NotForPublicRelease], [Sold], and
// [TopSecret]. Custom values are also supported.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/stampName
func (p PDFAnnotation) StampName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("stampName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAnnotation) SetStampName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setStampName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isActivatableTextField
func (p PDFAnnotation) ActivatableTextField() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isActivatableTextField"))
	return rv
}
