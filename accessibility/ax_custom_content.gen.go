// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXCustomContent] class.
var (
	_AXCustomContentClass     AXCustomContentClass
	_AXCustomContentClassOnce sync.Once
)

func getAXCustomContentClass() AXCustomContentClass {
	_AXCustomContentClassOnce.Do(func() {
		_AXCustomContentClass = AXCustomContentClass{class: objc.GetClass("AXCustomContent")}
	})
	return _AXCustomContentClass
}

// GetAXCustomContentClass returns the class object for AXCustomContent.
func GetAXCustomContentClass() AXCustomContentClass {
	return getAXCustomContentClass()
}

type AXCustomContentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXCustomContentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXCustomContentClass) Alloc() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Objects that define custom content and the timing of its output.
//
// # Overview
//
// An [AXCustomContent] object contains the accessibility strings for the
// labels you apply to your accessibility content. Combine them with the
// [AXCustomContentProvider] protocol to allow your users to experience the
// content in a more appropriate manner for each assistive technology.
//
// # Creating custom content
//
//   - [AXCustomContent.InitWithCoder]
//
// # Defining custom content
//
//   - [AXCustomContent.Label]: A localized string that identifies the label for this content.
//   - [AXCustomContent.AttributedLabel]: A localized attributed string that identifies the label for this content.
//   - [AXCustomContent.Value]: A localized string that provides a value for the label.
//   - [AXCustomContent.AttributedValue]: A localized attributed string that provides a value for the label.
//   - [AXCustomContent.Importance]: An object that determines when to output custom accessibility content.
//   - [AXCustomContent.SetImportance]
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent
type AXCustomContent struct {
	objectivec.Object
}

// AXCustomContentFromID constructs a [AXCustomContent] from an objc.ID.
//
// Objects that define custom content and the timing of its output.
func AXCustomContentFromID(id objc.ID) AXCustomContent {
	return AXCustomContent{objectivec.Object{ID: id}}
}

// NOTE: AXCustomContent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXCustomContent] class.
//
// # Creating custom content
//
//   - [IAXCustomContent.InitWithCoder]
//
// # Defining custom content
//
//   - [IAXCustomContent.Label]: A localized string that identifies the label for this content.
//   - [IAXCustomContent.AttributedLabel]: A localized attributed string that identifies the label for this content.
//   - [IAXCustomContent.Value]: A localized string that provides a value for the label.
//   - [IAXCustomContent.AttributedValue]: A localized attributed string that provides a value for the label.
//   - [IAXCustomContent.Importance]: An object that determines when to output custom accessibility content.
//   - [IAXCustomContent.SetImportance]
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent
type IAXCustomContent interface {
	objectivec.IObject

	// Topic: Creating custom content

	InitWithCoder(coder foundation.INSCoder) AXCustomContent

	// Topic: Defining custom content

	// A localized string that identifies the label for this content.
	Label() string
	// A localized attributed string that identifies the label for this content.
	AttributedLabel() foundation.NSAttributedString
	// A localized string that provides a value for the label.
	Value() string
	// A localized attributed string that provides a value for the label.
	AttributedValue() foundation.NSAttributedString
	// An object that determines when to output custom accessibility content.
	Importance() AXCustomContentImportance
	SetImportance(value AXCustomContentImportance)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AXCustomContent) Init() AXCustomContent {
	rv := objc.Send[AXCustomContent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXCustomContent) Autorelease() AXCustomContent {
	rv := objc.Send[AXCustomContent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXCustomContent creates a new AXCustomContent instance.
func NewAXCustomContent() AXCustomContent {
	class := getAXCustomContentClass()
	rv := objc.Send[AXCustomContent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates new custom content with an attributed string and attributed value.
//
// label: A localized attributed string that identifies the label for this content.
//
// value: A localized attributed string that provides a value for the label.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(attributedLabel:attributedValue:)
func NewAXCustomContentWithAttributedLabelAttributedValue(label foundation.NSAttributedString, value foundation.NSAttributedString) AXCustomContent {
	rv := objc.Send[objc.ID](objc.ID(getAXCustomContentClass().class), objc.Sel("customContentWithAttributedLabel:attributedValue:"), label, value)
	return AXCustomContentFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(coder:)
func NewAXCustomContentWithCoder(coder foundation.INSCoder) AXCustomContent {
	instance := getAXCustomContentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AXCustomContentFromID(rv)
}

// Creates new custom content with a label and value.
//
// label: A localized string that identifies the label for this content.
//
// value: A localized string that provides a value for the label.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(label:value:)
func NewAXCustomContentWithLabelValue(label string, value string) AXCustomContent {
	rv := objc.Send[objc.ID](objc.ID(getAXCustomContentClass().class), objc.Sel("customContentWithLabel:value:"), objc.String(label), objc.String(value))
	return AXCustomContentFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(coder:)
func (a AXCustomContent) InitWithCoder(coder foundation.INSCoder) AXCustomContent {
	rv := objc.Send[AXCustomContent](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a AXCustomContent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A localized string that identifies the label for this content.
//
// # Discussion
//
// Make the label succinct to work well with assistive technology. For
// example, [Orientation] is an appropriate name for photo information.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/label
func (a AXCustomContent) Label() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// A localized attributed string that identifies the label for this content.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/attributedLabel
func (a AXCustomContent) AttributedLabel() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedLabel"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

// A localized string that provides a value for the label.
//
// # Discussion
//
// Make the value succinct to work well with assistive technology. For
// example, either [Portrait] or [Landscape] is an appropriate content value
// for an [Orientation] label on a photo.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/value
func (a AXCustomContent) Value() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("value"))
	return foundation.NSStringFromID(rv).String()
}

// A localized attributed string that provides a value for the label.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/attributedValue
func (a AXCustomContent) AttributedValue() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedValue"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

// An object that determines when to output custom accessibility content.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/importance-swift.property
func (a AXCustomContent) Importance() AXCustomContentImportance {
	rv := objc.Send[AXCustomContentImportance](a.ID, objc.Sel("importance"))
	return AXCustomContentImportance(rv)
}
func (a AXCustomContent) SetImportance(value AXCustomContentImportance) {
	objc.Send[struct{}](a.ID, objc.Sel("setImportance:"), value)
}
