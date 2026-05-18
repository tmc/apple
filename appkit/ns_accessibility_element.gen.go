// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAccessibilityElement] class.
var (
	_NSAccessibilityElementClass     NSAccessibilityElementClass
	_NSAccessibilityElementClassOnce sync.Once
)

func getNSAccessibilityElementClass() NSAccessibilityElementClass {
	_NSAccessibilityElementClassOnce.Do(func() {
		_NSAccessibilityElementClass = NSAccessibilityElementClass{class: objc.GetClass("NSAccessibilityElement")}
	})
	return _NSAccessibilityElementClass
}

// GetNSAccessibilityElementClass returns the class object for NSAccessibilityElement.
func GetNSAccessibilityElementClass() NSAccessibilityElementClass {
	return getNSAccessibilityElementClass()
}

type NSAccessibilityElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAccessibilityElementClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAccessibilityElementClass) Alloc() NSAccessibilityElement {
	rv := objc.Send[NSAccessibilityElement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The basic infrastructure necessary for interacting with an assistive app.
//
// # Overview
//
// Create subclasses of the [NSAccessibilityElement] class to represent any of
// your user interface elements that don’t inherit from [NSView] or from one
// of the standard AppKit controls. This class represents your user interface
// element in the accessibility hierarchy and manages the details necessary
// for working with assistive apps.
//
// To support accessibility features for a custom user interface element:
//
// - Create your [NSAccessibilityElement] subclass by using
// [NSAccessibilityElement.AccessibilityElementWithRoleFrameLabelParent]. You can also set these
// values using [SetAccessibilityRole], [SetAccessibilityLabel] and
// [SetAccessibilityParent]. - Call the parent’s
// [NSAccessibilityElement.AccessibilityAddChildElement] method to add your subclass. You can also
// add the subclass to its parent’s [accessibilityChildren] array using
// [SetAccessibilityChildren]. - In your subclass, call
// [NSAccessibilityElement.SetAccessibilityFrameInParentSpace]. This ensures that your control moves
// with its superview. - In your subclass, adopt a role-specific protocol,
// customize the role, and post notifications just as you would handle any
// other accessible control. See [Custom Controls]. - In your subclass,
// implement any additional properties and methods you may need to use to
// further customize your user interface element’s accessibility behavior.
// See [NSAccessibilityProtocol].
//
// # Supporting the Accessibility Hierarchy
//
//   - [NSAccessibilityElement.AccessibilityAddChildElement]: Adds a child to the accessibility element in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class
//
// [Custom Controls]: https://developer.apple.com/documentation/AppKit/custom-controls
// [accessibilityChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
type NSAccessibilityElement struct {
	objectivec.Object
}

// NSAccessibilityElementFromID constructs a [NSAccessibilityElement] from an objc.ID.
//
// The basic infrastructure necessary for interacting with an assistive app.
func NSAccessibilityElementFromID(id objc.ID) NSAccessibilityElement {
	return NSAccessibilityElement{objectivec.Object{ID: id}}
}

// NOTE: NSAccessibilityElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAccessibilityElement] class.
//
// # Supporting the Accessibility Hierarchy
//
//   - [INSAccessibilityElement.AccessibilityAddChildElement]: Adds a child to the accessibility element in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class
type INSAccessibilityElement interface {
	objectivec.IObject

	// Topic: Supporting the Accessibility Hierarchy

	// Adds a child to the accessibility element in the accessibility hierarchy.
	AccessibilityAddChildElement(childElement NSAccessibilityElement)

	// The accessibility element’s frame in its parent’s coordinate system.
	AccessibilityFrameInParentSpace() corefoundation.CGRect
	SetAccessibilityFrameInParentSpace(value corefoundation.CGRect)
}

// Init initializes the instance.
func (a NSAccessibilityElement) Init() NSAccessibilityElement {
	rv := objc.Send[NSAccessibilityElement](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAccessibilityElement) Autorelease() NSAccessibilityElement {
	rv := objc.Send[NSAccessibilityElement](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAccessibilityElement creates a new NSAccessibilityElement instance.
func NewNSAccessibilityElement() NSAccessibilityElement {
	class := getNSAccessibilityElementClass()
	rv := objc.Send[NSAccessibilityElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a child to the accessibility element in the accessibility hierarchy.
//
// childElement: The child element to be added.
//
// # Discussion
//
// Calling this method sets up the proper parent-child relationship between
// the current element and the provided child element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a NSAccessibilityElement) AccessibilityAddChildElement(childElement NSAccessibilityElement) {
	objc.Send[objc.ID](a.ID, objc.Sel("accessibilityAddChildElement:"), childElement)
}

// Instantiates and configures a new accessibility element.
//
// role: The new element’s intended role. For a complete list of roles, see Roles.
//
// frame: The element’s frame in screen coordinates. Additionally, you need to set
// the element’s [AccessibilityFrameInParentSpace] property.
//
// label: A short description of the new element. Do not include the element’s type
// in the label (for example, use [Play], not `Play button`). If possible, use
// a single word. To help ensure that accessibility clients such as VoiceOver
// read the label with the correct intonation, start the label with a capital
// letter. Do not put a period at the end. Always localize the label.
//
// parent: The new element’s parent in the accessibility hierarchy.
//
// # Return Value
//
// A newly instantiated and initialized accessibility element.
//
// # Discussion
//
// Alternatively, instead of calling this convenience method, you can create
// an accessibility element and set its [accessibilityRole],
// [accessibilityLabel], and [accessibilityParent] properties. Regardless of
// how you create the accessibility element, you need to set its
// [AccessibilityFrameInParentSpace] property to ensure that the element’s
// frame is updated as its parent moves.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
// [accessibilityRole]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
func (_NSAccessibilityElementClass NSAccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role NSAccessibilityRole, frame corefoundation.CGRect, label string, parent objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSAccessibilityElementClass.class), objc.Sel("accessibilityElementWithRole:frame:label:parent:"), objc.String(string(role)), frame, objc.String(label), parent)
	return objectivec.Object{ID: rv}
}

// The accessibility element’s frame in its parent’s coordinate system.
//
// # Discussion
//
// Setting this property ensures that the accessibility client receives the
// correct frame (in screen coordinates) as the element’s parent moves.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
func (a NSAccessibilityElement) AccessibilityFrameInParentSpace() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("accessibilityFrameInParentSpace"))
	return corefoundation.CGRect(rv)
}
func (a NSAccessibilityElement) SetAccessibilityFrameInParentSpace(value corefoundation.CGRect) {
	objc.Send[struct{}](a.ID, objc.Sel("setAccessibilityFrameInParentSpace:"), value)
}
