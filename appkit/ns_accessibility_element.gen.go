// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
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
// values using [setAccessibilityRole(_:)], [setAccessibilityLabel(_:)] and
// [setAccessibilityParent(_:)]. - Call the parent’s
// [NSAccessibilityElement.AccessibilityAddChildElement] method to add your subclass. You can also
// add the subclass to its parent’s [accessibilityChildren] array using
// [setAccessibilityChildren(_:)]. - In your subclass, call
// [NSAccessibilityElement.SetAccessibilityFrameInParentSpace]. This ensures that your control moves
// with its superview. - In your subclass, adopt a role-specific protocol,
// customize the role, and post notifications just as you would handle any
// other accessible control. See [Custom Controls]. - In your subclass,
// implement any additional properties and methods you may need to use to
// further customize your user interface element’s accessibility behavior.
// See [NSAccessibilityProtocol].
//
// [Custom Controls]: https://developer.apple.com/documentation/AppKit/custom-controls
// [accessibilityChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
// [setAccessibilityChildren(_:)]: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
// [setAccessibilityLabel(_:)]: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
// [setAccessibilityParent(_:)]: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
// [setAccessibilityRole(_:)]: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
//
// # Supporting the Accessibility Hierarchy
//
//   - [NSAccessibilityElement.AccessibilityAddChildElement]: Adds a child to the accessibility element in the accessibility hierarchy.
//   - [NSAccessibilityElement.AccessibilityFrameInParentSpace]: Returns the accessibility element’s frame in its parent’s coordinate system.
//   - [NSAccessibilityElement.SetAccessibilityFrameInParentSpace]: Sets the accessibility element’s frame in its parent’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class
type NSAccessibilityElement struct {
	objectivec.Object
}

// NSAccessibilityElementFromID constructs a [NSAccessibilityElement] from an objc.ID.
//
// The basic infrastructure necessary for interacting with an assistive app.
func NSAccessibilityElementFromID(id objc.ID) NSAccessibilityElement {
	return NSAccessibilityElement{objectivec.Object{id}}
}
// NOTE: NSAccessibilityElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSAccessibilityElement] class.
//
// # Supporting the Accessibility Hierarchy
//
//   - [INSAccessibilityElement.AccessibilityAddChildElement]: Adds a child to the accessibility element in the accessibility hierarchy.
//   - [INSAccessibilityElement.AccessibilityFrameInParentSpace]: Returns the accessibility element’s frame in its parent’s coordinate system.
//   - [INSAccessibilityElement.SetAccessibilityFrameInParentSpace]: Sets the accessibility element’s frame in its parent’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class
type INSAccessibilityElement interface {
	objectivec.IObject
	

	// Topic: Supporting the Accessibility Hierarchy

	// Adds a child to the accessibility element in the accessibility hierarchy.
	AccessibilityAddChildElement(childElement NSAccessibilityElement)
	// Returns the accessibility element’s frame in its parent’s coordinate system.
	AccessibilityFrameInParentSpace() corefoundation.CGRect
	// Sets the accessibility element’s frame in its parent’s coordinate system.
	SetAccessibilityFrameInParentSpace(accessibilityFrameInParentSpace corefoundation.CGRect)
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

// Returns the accessibility element’s frame in its parent’s coordinate
// system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace()
func (a NSAccessibilityElement) AccessibilityFrameInParentSpace() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("accessibilityFrameInParentSpace"))
	return corefoundation.CGRect(rv)
}

// Sets the accessibility element’s frame in its parent’s coordinate
// system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/setAccessibilityFrameInParentSpace(_:)
func (a NSAccessibilityElement) SetAccessibilityFrameInParentSpace(accessibilityFrameInParentSpace corefoundation.CGRect) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFrameInParentSpace:"), accessibilityFrameInParentSpace)
}





// Instantiates and configures a new accessibility element.
//
// role: The new element’s intended role. For a complete list of roles, see Roles.
//
// frame: The element’s frame in screen coordinates. Additionally, you need to set
// the element’s [accessibilityFrameInParentSpace] property.
// //
// [accessibilityFrameInParentSpace]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
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
// [accessibilityFrameInParentSpace] property to ensure that the element’s
// frame is updated as its parent moves.
//
// [accessibilityFrameInParentSpace]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
// [accessibilityRole]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (_NSAccessibilityElementClass NSAccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role NSAccessibilityRole, frame corefoundation.CGRect, label string, parent objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSAccessibilityElementClass.class), objc.Sel("accessibilityElementWithRole:frame:label:parent:"), objc.String(string(role)), frame, objc.String(label), parent)
	return objectivec.Object{ID: rv}
}


















			// Protocol methods for NSAccessibilityProtocol
			

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()

func (o NSAccessibilityElement) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)

func (o NSAccessibilityElement) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()

func (o NSAccessibilityElement) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)

func (o NSAccessibilityElement) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}

// Returns the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityframe()

func (o NSAccessibilityElement) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)

func (o NSAccessibilityElement) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}

// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()

func (o NSAccessibilityElement) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)

func (o NSAccessibilityElement) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}

// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()

func (o NSAccessibilityElement) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)

func (o NSAccessibilityElement) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()

func (o NSAccessibilityElement) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)

func (o NSAccessibilityElement) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}

// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()

func (o NSAccessibilityElement) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)

func (o NSAccessibilityElement) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}

// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()

func (o NSAccessibilityElement) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)

func (o NSAccessibilityElement) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}

// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()

func (o NSAccessibilityElement) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}

// Returns the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityidentifier()

func (o NSAccessibilityElement) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)

func (o NSAccessibilityElement) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}

// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()

func (o NSAccessibilityElement) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}

// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()

func (o NSAccessibilityElement) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}

// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()

func (o NSAccessibilityElement) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}

// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)

func (o NSAccessibilityElement) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()

func (o NSAccessibilityElement) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)

func (o NSAccessibilityElement) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()

func (o NSAccessibilityElement) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)

func (o NSAccessibilityElement) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}

// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()

func (o NSAccessibilityElement) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}

// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)

func (o NSAccessibilityElement) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()

func (o NSAccessibilityElement) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)

func (o NSAccessibilityElement) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}

// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()

func (o NSAccessibilityElement) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()

func (o NSAccessibilityElement) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)

func (o NSAccessibilityElement) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()

func (o NSAccessibilityElement) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)

func (o NSAccessibilityElement) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityparent()

func (o NSAccessibilityElement) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)

func (o NSAccessibilityElement) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}

// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()

func (o NSAccessibilityElement) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}

// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()

func (o NSAccessibilityElement) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)

func (o NSAccessibilityElement) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()

func (o NSAccessibilityElement) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)

func (o NSAccessibilityElement) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}

// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()

func (o NSAccessibilityElement) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)

func (o NSAccessibilityElement) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// has the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfocused()

func (o NSAccessibilityElement) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)

func (o NSAccessibilityElement) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}

// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()

func (o NSAccessibilityElement) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)

func (o NSAccessibilityElement) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()

func (o NSAccessibilityElement) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)

func (o NSAccessibilityElement) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()

func (o NSAccessibilityElement) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)

func (o NSAccessibilityElement) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()

func (o NSAccessibilityElement) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)

func (o NSAccessibilityElement) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()

func (o NSAccessibilityElement) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)

func (o NSAccessibilityElement) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()

func (o NSAccessibilityElement) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)

func (o NSAccessibilityElement) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}

// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()

func (o NSAccessibilityElement) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}

// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)

func (o NSAccessibilityElement) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}

// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()

func (o NSAccessibilityElement) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}

// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)

func (o NSAccessibilityElement) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}

// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()

func (o NSAccessibilityElement) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}

// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)

func (o NSAccessibilityElement) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}

// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()

func (o NSAccessibilityElement) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}

// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)

func (o NSAccessibilityElement) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}

// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()

func (o NSAccessibilityElement) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)

func (o NSAccessibilityElement) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}

// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()

func (o NSAccessibilityElement) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}

// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()

func (o NSAccessibilityElement) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}

// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}

// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()

func (o NSAccessibilityElement) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}

// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}

// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()

func (o NSAccessibilityElement) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}

// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)

func (o NSAccessibilityElement) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}

// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()

func (o NSAccessibilityElement) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)

func (o NSAccessibilityElement) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}

// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()

func (o NSAccessibilityElement) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}

// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)

func (o NSAccessibilityElement) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
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

func (o NSAccessibilityElement) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
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

func (o NSAccessibilityElement) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
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

func (o NSAccessibilityElement) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
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

func (o NSAccessibilityElement) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
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

func (o NSAccessibilityElement) AccessibilityLineForIndex(index int) int {
	
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

func (o NSAccessibilityElement) AccessibilityRangeForIndex(index int) foundation.NSRange {
	
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

func (o NSAccessibilityElement) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	
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

func (o NSAccessibilityElement) AccessibilityRangeForLine(line int) foundation.NSRange {
	
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

func (o NSAccessibilityElement) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
	}

// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()

func (o NSAccessibilityElement) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}

// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)

func (o NSAccessibilityElement) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()

func (o NSAccessibilityElement) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)

func (o NSAccessibilityElement) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()

func (o NSAccessibilityElement) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()

func (o NSAccessibilityElement) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)

func (o NSAccessibilityElement) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()

func (o NSAccessibilityElement) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()

func (o NSAccessibilityElement) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()

func (o NSAccessibilityElement) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)

func (o NSAccessibilityElement) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()

func (o NSAccessibilityElement) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)

func (o NSAccessibilityElement) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()

func (o NSAccessibilityElement) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)

func (o NSAccessibilityElement) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()

func (o NSAccessibilityElement) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)

func (o NSAccessibilityElement) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}

// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()

func (o NSAccessibilityElement) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}

// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)

func (o NSAccessibilityElement) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()

func (o NSAccessibilityElement) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)

func (o NSAccessibilityElement) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}

// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()

func (o NSAccessibilityElement) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)

func (o NSAccessibilityElement) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()

func (o NSAccessibilityElement) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}

// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()

func (o NSAccessibilityElement) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)

func (o NSAccessibilityElement) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()

func (o NSAccessibilityElement) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)

func (o NSAccessibilityElement) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}

// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()

func (o NSAccessibilityElement) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)

func (o NSAccessibilityElement) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()

func (o NSAccessibilityElement) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)

func (o NSAccessibilityElement) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()

func (o NSAccessibilityElement) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)

func (o NSAccessibilityElement) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}

// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()

func (o NSAccessibilityElement) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)

func (o NSAccessibilityElement) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}

// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()

func (o NSAccessibilityElement) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)

func (o NSAccessibilityElement) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}

// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()

func (o NSAccessibilityElement) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)

func (o NSAccessibilityElement) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}

// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()

func (o NSAccessibilityElement) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}

// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)

func (o NSAccessibilityElement) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()

func (o NSAccessibilityElement) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)

func (o NSAccessibilityElement) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}

// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()

func (o NSAccessibilityElement) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}

// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)

func (o NSAccessibilityElement) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}

// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()

func (o NSAccessibilityElement) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)

func (o NSAccessibilityElement) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}

// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()

func (o NSAccessibilityElement) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)

func (o NSAccessibilityElement) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}

// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()

func (o NSAccessibilityElement) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)

func (o NSAccessibilityElement) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}

// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()

func (o NSAccessibilityElement) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)

func (o NSAccessibilityElement) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}

// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()

func (o NSAccessibilityElement) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}

// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)

func (o NSAccessibilityElement) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()

func (o NSAccessibilityElement) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)

func (o NSAccessibilityElement) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}

// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()

func (o NSAccessibilityElement) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}

// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)

func (o NSAccessibilityElement) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()

func (o NSAccessibilityElement) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)

func (o NSAccessibilityElement) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}

// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()

func (o NSAccessibilityElement) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)

func (o NSAccessibilityElement) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}

// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()

func (o NSAccessibilityElement) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)

func (o NSAccessibilityElement) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}

// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()

func (o NSAccessibilityElement) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}

// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()

func (o NSAccessibilityElement) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}

// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()

func (o NSAccessibilityElement) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}

// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)

func (o NSAccessibilityElement) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}

// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()

func (o NSAccessibilityElement) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)

func (o NSAccessibilityElement) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}

// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()

func (o NSAccessibilityElement) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)

func (o NSAccessibilityElement) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()

func (o NSAccessibilityElement) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)

func (o NSAccessibilityElement) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}

// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()

func (o NSAccessibilityElement) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}

// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)

func (o NSAccessibilityElement) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}

// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()

func (o NSAccessibilityElement) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)

func (o NSAccessibilityElement) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}

// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()

func (o NSAccessibilityElement) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}

// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)

func (o NSAccessibilityElement) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}

// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()

func (o NSAccessibilityElement) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}

// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)

func (o NSAccessibilityElement) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}

// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()

func (o NSAccessibilityElement) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}

// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)

func (o NSAccessibilityElement) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}

// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()

func (o NSAccessibilityElement) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)

func (o NSAccessibilityElement) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}

// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()

func (o NSAccessibilityElement) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)

func (o NSAccessibilityElement) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
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

func (o NSAccessibilityElement) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}

// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()

func (o NSAccessibilityElement) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)

func (o NSAccessibilityElement) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}

// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()

func (o NSAccessibilityElement) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}

// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)

func (o NSAccessibilityElement) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}

// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()

func (o NSAccessibilityElement) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)

func (o NSAccessibilityElement) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}

// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()

func (o NSAccessibilityElement) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}

// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)

func (o NSAccessibilityElement) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}

// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()

func (o NSAccessibilityElement) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)

func (o NSAccessibilityElement) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
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

func (o NSAccessibilityElement) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
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

func (o NSAccessibilityElement) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	
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

func (o NSAccessibilityElement) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
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

func (o NSAccessibilityElement) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
	}

// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()

func (o NSAccessibilityElement) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}

// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)

func (o NSAccessibilityElement) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}

// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()

func (o NSAccessibilityElement) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)

func (o NSAccessibilityElement) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}

// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()

func (o NSAccessibilityElement) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}

// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)

func (o NSAccessibilityElement) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}

// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()

func (o NSAccessibilityElement) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)

func (o NSAccessibilityElement) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}

// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()

func (o NSAccessibilityElement) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)

func (o NSAccessibilityElement) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()

func (o NSAccessibilityElement) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)

func (o NSAccessibilityElement) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}

// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()

func (o NSAccessibilityElement) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}

// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()

func (o NSAccessibilityElement) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}

// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)

func (o NSAccessibilityElement) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()

func (o NSAccessibilityElement) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)

func (o NSAccessibilityElement) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}

// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()

func (o NSAccessibilityElement) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)

func (o NSAccessibilityElement) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}

// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()

func (o NSAccessibilityElement) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)

func (o NSAccessibilityElement) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}

// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()

func (o NSAccessibilityElement) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}

// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)

func (o NSAccessibilityElement) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}

// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()

func (o NSAccessibilityElement) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}

// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)

func (o NSAccessibilityElement) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}

// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()

func (o NSAccessibilityElement) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}

// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)

func (o NSAccessibilityElement) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}

// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()

func (o NSAccessibilityElement) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)

func (o NSAccessibilityElement) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}

// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()

func (o NSAccessibilityElement) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)

func (o NSAccessibilityElement) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()

func (o NSAccessibilityElement) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)

func (o NSAccessibilityElement) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()

func (o NSAccessibilityElement) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)

func (o NSAccessibilityElement) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}

// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()

func (o NSAccessibilityElement) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)

func (o NSAccessibilityElement) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}

// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()

func (o NSAccessibilityElement) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)

func (o NSAccessibilityElement) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()

func (o NSAccessibilityElement) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)

func (o NSAccessibilityElement) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}

// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()

func (o NSAccessibilityElement) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}

// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()

func (o NSAccessibilityElement) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)

func (o NSAccessibilityElement) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}

// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()

func (o NSAccessibilityElement) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)

func (o NSAccessibilityElement) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}

// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()

func (o NSAccessibilityElement) AccessibilityPerformCancel() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
	}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()

func (o NSAccessibilityElement) AccessibilityPerformConfirm() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
	}

// Selects the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()

func (o NSAccessibilityElement) AccessibilityPerformPick() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
	}

// Simulates clicking the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()

func (o NSAccessibilityElement) AccessibilityPerformPress() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
	}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()

func (o NSAccessibilityElement) AccessibilityPerformShowAlternateUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
	}

// Returns to the accessibility element’s original UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()

func (o NSAccessibilityElement) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}

// Displays the menu accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()

func (o NSAccessibilityElement) AccessibilityPerformShowMenu() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
	}

// Brings the window to the front.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()

func (o NSAccessibilityElement) AccessibilityPerformRaise() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
	}

// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()

func (o NSAccessibilityElement) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}

// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()

func (o NSAccessibilityElement) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)

func (o NSAccessibilityElement) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}

// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()

func (o NSAccessibilityElement) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}

// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()

func (o NSAccessibilityElement) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}

// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()

func (o NSAccessibilityElement) AccessibilityPerformDelete() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()

func (o NSAccessibilityElement) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()

func (o NSAccessibilityElement) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)

func (o NSAccessibilityElement) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)

func (o NSAccessibilityElement) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
	}














