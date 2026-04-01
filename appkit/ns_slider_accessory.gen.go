// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSliderAccessory] class.
var (
	_NSSliderAccessoryClass     NSSliderAccessoryClass
	_NSSliderAccessoryClassOnce sync.Once
)

func getNSSliderAccessoryClass() NSSliderAccessoryClass {
	_NSSliderAccessoryClassOnce.Do(func() {
		_NSSliderAccessoryClass = NSSliderAccessoryClass{class: objc.GetClass("NSSliderAccessory")}
	})
	return _NSSliderAccessoryClass
}

// GetNSSliderAccessoryClass returns the class object for NSSliderAccessory.
func GetNSSliderAccessoryClass() NSSliderAccessoryClass {
	return getNSSliderAccessoryClass()
}

type NSSliderAccessoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSliderAccessoryClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderAccessoryClass) Alloc() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [NSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [NSSliderAccessory.SetBehavior]
//   - [NSSliderAccessory.Enabled]
//   - [NSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type NSSliderAccessory struct {
	objectivec.Object
}

// NSSliderAccessoryFromID constructs a [NSSliderAccessory] from an objc.ID.
func NSSliderAccessoryFromID(id objc.ID) NSSliderAccessory {
	return NSSliderAccessory{objectivec.Object{ID: id}}
}

// NOTE: NSSliderAccessory adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSliderAccessory] class.
//
// # Instance Properties
//
//   - [INSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [INSSliderAccessory.SetBehavior]
//   - [INSSliderAccessory.Enabled]
//   - [INSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type INSSliderAccessory interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The effect on interaction with the accessory.
	Behavior() INSSliderAccessoryBehavior
	SetBehavior(value INSSliderAccessoryBehavior)
	Enabled() bool
	SetEnabled(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSliderAccessory) Init() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderAccessory) Autorelease() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderAccessory creates a new NSSliderAccessory instance.
func NewNSSliderAccessory() NSSliderAccessory {
	class := getNSSliderAccessoryClass()
	rv := objc.Send[NSSliderAccessory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(image:)
func NewSliderAccessoryWithImage(image INSImage) NSSliderAccessory {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderAccessoryClass().class), objc.Sel("accessoryWithImage:"), image)
	return NSSliderAccessoryFromID(rv)
}

func (s NSSliderAccessory) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The effect on interaction with the accessory.
//
// # Discussion
//
// The default value is `automaticBehavior`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/behavior
func (s NSSliderAccessory) Behavior() INSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("behavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}
func (s NSSliderAccessory) SetBehavior(value INSSliderAccessoryBehavior) {
	objc.Send[struct{}](s.ID, objc.Sel("setBehavior:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/isEnabled
func (s NSSliderAccessory) Enabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabled"))
	return rv
}
func (s NSSliderAccessory) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}

// Protocol methods for NSAccessibilityElementProtocol

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSSliderAccessory) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
//
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSSliderAccessory) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s identity.
//
// # Return Value
//
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSSliderAccessory) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSSliderAccessory) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSAccessibilityProtocol

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSSliderAccessory) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (o NSSliderAccessory) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSSliderAccessory) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (o NSSliderAccessory) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (o NSSliderAccessory) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (o NSSliderAccessory) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (o NSSliderAccessory) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (o NSSliderAccessory) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (o NSSliderAccessory) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (o NSSliderAccessory) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (o NSSliderAccessory) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (o NSSliderAccessory) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (o NSSliderAccessory) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
//
// true, if accessibility clients can call the selector; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (o NSSliderAccessory) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (o NSSliderAccessory) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (o NSSliderAccessory) SetAccessibilityContents(accessibilityContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (o NSSliderAccessory) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (o NSSliderAccessory) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (o NSSliderAccessory) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (o NSSliderAccessory) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (o NSSliderAccessory) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (o NSSliderAccessory) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (o NSSliderAccessory) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (o NSSliderAccessory) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (o NSSliderAccessory) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSSliderAccessory) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (o NSSliderAccessory) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSSliderAccessory) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (o NSSliderAccessory) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (o NSSliderAccessory) AccessibilityURL() foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (o NSSliderAccessory) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (o NSSliderAccessory) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (o NSSliderAccessory) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (o NSSliderAccessory) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (o NSSliderAccessory) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (o NSSliderAccessory) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (o NSSliderAccessory) SetAccessibilityChildren(accessibilityChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (o NSSliderAccessory) AccessibilityChildrenInNavigationOrder() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (o NSSliderAccessory) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (o NSSliderAccessory) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (o NSSliderAccessory) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedChildren(accessibilitySelectedChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (o NSSliderAccessory) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (o NSSliderAccessory) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (o NSSliderAccessory) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (o NSSliderAccessory) SetAccessibilityVisibleChildren(accessibilityVisibleChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (o NSSliderAccessory) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (o NSSliderAccessory) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (o NSSliderAccessory) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (o NSSliderAccessory) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (o NSSliderAccessory) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (o NSSliderAccessory) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (o NSSliderAccessory) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSSliderAccessory) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (o NSSliderAccessory) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (o NSSliderAccessory) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (o NSSliderAccessory) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (o NSSliderAccessory) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (o NSSliderAccessory) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (o NSSliderAccessory) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (o NSSliderAccessory) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (o NSSliderAccessory) AccessibilityCustomActions() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (o NSSliderAccessory) SetAccessibilityCustomActions(accessibilityCustomActions objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (o NSSliderAccessory) AccessibilityCustomRotors() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (o NSSliderAccessory) SetAccessibilityCustomRotors(accessibilityCustomRotors objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (o NSSliderAccessory) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (o NSSliderAccessory) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (o NSSliderAccessory) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (o NSSliderAccessory) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (o NSSliderAccessory) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (o NSSliderAccessory) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (o NSSliderAccessory) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (o NSSliderAccessory) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (o NSSliderAccessory) AccessibilitySelectedTextRanges() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSArrayFromID(rv)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (o NSSliderAccessory) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (o NSSliderAccessory) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (o NSSliderAccessory) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (o NSSliderAccessory) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (o NSSliderAccessory) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (o NSSliderAccessory) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
//
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
func (o NSSliderAccessory) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
func (o NSSliderAccessory) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
//
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
func (o NSSliderAccessory) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// The rectangle that encloses the specified characters.
//
// # Discussion
//
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
func (o NSSliderAccessory) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
//
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
func (o NSSliderAccessory) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
//
// The range of characters for the glyph.
//
// # Discussion
//
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
func (o NSSliderAccessory) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
//
// A range of characters with the same style as the specified character.
//
// # Discussion
//
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
func (o NSSliderAccessory) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
//
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
func (o NSSliderAccessory) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
//
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
func (o NSSliderAccessory) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (o NSSliderAccessory) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (o NSSliderAccessory) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSSliderAccessory) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (o NSSliderAccessory) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (o NSSliderAccessory) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (o NSSliderAccessory) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (o NSSliderAccessory) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (o NSSliderAccessory) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (o NSSliderAccessory) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (o NSSliderAccessory) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (o NSSliderAccessory) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (o NSSliderAccessory) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (o NSSliderAccessory) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (o NSSliderAccessory) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSSliderAccessory) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (o NSSliderAccessory) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (o NSSliderAccessory) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (o NSSliderAccessory) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSSliderAccessory) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (o NSSliderAccessory) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSSliderAccessory) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (o NSSliderAccessory) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (o NSSliderAccessory) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (o NSSliderAccessory) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (o NSSliderAccessory) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (o NSSliderAccessory) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (o NSSliderAccessory) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (o NSSliderAccessory) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (o NSSliderAccessory) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (o NSSliderAccessory) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (o NSSliderAccessory) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (o NSSliderAccessory) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (o NSSliderAccessory) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (o NSSliderAccessory) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSSliderAccessory) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (o NSSliderAccessory) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSSliderAccessory) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (o NSSliderAccessory) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (o NSSliderAccessory) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (o NSSliderAccessory) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (o NSSliderAccessory) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (o NSSliderAccessory) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (o NSSliderAccessory) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (o NSSliderAccessory) SetAccessibilityWindows(accessibilityWindows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (o NSSliderAccessory) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (o NSSliderAccessory) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSSliderAccessory) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (o NSSliderAccessory) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (o NSSliderAccessory) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (o NSSliderAccessory) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (o NSSliderAccessory) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (o NSSliderAccessory) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (o NSSliderAccessory) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (o NSSliderAccessory) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (o NSSliderAccessory) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (o NSSliderAccessory) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (o NSSliderAccessory) SetAccessibilityColumns(accessibilityColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (o NSSliderAccessory) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (o NSSliderAccessory) SetAccessibilityColumnTitles(accessibilityColumnTitles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSSliderAccessory) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (o NSSliderAccessory) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (o NSSliderAccessory) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (o NSSliderAccessory) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (o NSSliderAccessory) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (o NSSliderAccessory) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (o NSSliderAccessory) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (o NSSliderAccessory) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (o NSSliderAccessory) SetAccessibilityRows(accessibilityRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (o NSSliderAccessory) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedColumns(accessibilitySelectedColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (o NSSliderAccessory) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedRows(accessibilitySelectedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (o NSSliderAccessory) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (o NSSliderAccessory) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (o NSSliderAccessory) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (o NSSliderAccessory) SetAccessibilityVisibleColumns(accessibilityVisibleColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (o NSSliderAccessory) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (o NSSliderAccessory) SetAccessibilityVisibleRows(accessibilityVisibleRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSSliderAccessory) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (o NSSliderAccessory) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (o NSSliderAccessory) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (o NSSliderAccessory) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (o NSSliderAccessory) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (o NSSliderAccessory) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (o NSSliderAccessory) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (o NSSliderAccessory) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (o NSSliderAccessory) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (o NSSliderAccessory) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (o NSSliderAccessory) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (o NSSliderAccessory) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (o NSSliderAccessory) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (o NSSliderAccessory) SetAccessibilitySelectedCells(accessibilitySelectedCells objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (o NSSliderAccessory) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (o NSSliderAccessory) SetAccessibilityVisibleCells(accessibilityVisibleCells objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
//
// The cell specified by the column and row indexes.
//
// # Discussion
//
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)
func (o NSSliderAccessory) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (o NSSliderAccessory) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (o NSSliderAccessory) SetAccessibilityHandles(accessibilityHandles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (o NSSliderAccessory) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (o NSSliderAccessory) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (o NSSliderAccessory) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (o NSSliderAccessory) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (o NSSliderAccessory) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (o NSSliderAccessory) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (o NSSliderAccessory) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (o NSSliderAccessory) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
//
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
func (o NSSliderAccessory) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
//
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
func (o NSSliderAccessory) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
//
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
func (o NSSliderAccessory) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
//
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
func (o NSSliderAccessory) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (o NSSliderAccessory) AccessibilityAllowedValues() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSArrayFromID(rv)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (o NSSliderAccessory) SetAccessibilityAllowedValues(accessibilityAllowedValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (o NSSliderAccessory) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityLabelUIElements(accessibilityLabelUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (o NSSliderAccessory) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (o NSSliderAccessory) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (o NSSliderAccessory) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (o NSSliderAccessory) SetAccessibilityNextContents(accessibilityNextContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (o NSSliderAccessory) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (o NSSliderAccessory) SetAccessibilityPreviousContents(accessibilityPreviousContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (o NSSliderAccessory) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (o NSSliderAccessory) SetAccessibilitySplitters(accessibilitySplitters objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (o NSSliderAccessory) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (o NSSliderAccessory) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (o NSSliderAccessory) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (o NSSliderAccessory) SetAccessibilityTabs(accessibilityTabs objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (o NSSliderAccessory) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (o NSSliderAccessory) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (o NSSliderAccessory) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (o NSSliderAccessory) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (o NSSliderAccessory) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (o NSSliderAccessory) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (o NSSliderAccessory) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (o NSSliderAccessory) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (o NSSliderAccessory) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (o NSSliderAccessory) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (o NSSliderAccessory) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (o NSSliderAccessory) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (o NSSliderAccessory) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (o NSSliderAccessory) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (o NSSliderAccessory) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSSliderAccessory) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (o NSSliderAccessory) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (o NSSliderAccessory) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (o NSSliderAccessory) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (o NSSliderAccessory) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (o NSSliderAccessory) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (o NSSliderAccessory) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (o NSSliderAccessory) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (o NSSliderAccessory) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (o NSSliderAccessory) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (o NSSliderAccessory) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (o NSSliderAccessory) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (o NSSliderAccessory) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (o NSSliderAccessory) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (o NSSliderAccessory) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
}

// Cancels the current operation.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (o NSSliderAccessory) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (o NSSliderAccessory) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
}

// Selects the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (o NSSliderAccessory) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
}

// Simulates clicking the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (o NSSliderAccessory) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (o NSSliderAccessory) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

// Returns to the accessibility element’s original UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (o NSSliderAccessory) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

// Displays the menu accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (o NSSliderAccessory) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Brings the window to the front.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (o NSSliderAccessory) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
}

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (o NSSliderAccessory) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (o NSSliderAccessory) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (o NSSliderAccessory) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (o NSSliderAccessory) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
}

// Increments the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSSliderAccessory) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

// Decrements the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSSliderAccessory) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

// Deletes the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (o NSSliderAccessory) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (o NSSliderAccessory) AccessibilityAttributedUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (o NSSliderAccessory) AccessibilityUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (o NSSliderAccessory) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (o NSSliderAccessory) SetAccessibilityUserInputLabels(accessibilityUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
}
