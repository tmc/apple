// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStatusBarButton] class.
var (
	_NSStatusBarButtonClass     NSStatusBarButtonClass
	_NSStatusBarButtonClassOnce sync.Once
)

func getNSStatusBarButtonClass() NSStatusBarButtonClass {
	_NSStatusBarButtonClassOnce.Do(func() {
		_NSStatusBarButtonClass = NSStatusBarButtonClass{class: objc.GetClass("NSStatusBarButton")}
	})
	return _NSStatusBarButtonClass
}

// GetNSStatusBarButtonClass returns the class object for NSStatusBarButton.
func GetNSStatusBarButtonClass() NSStatusBarButtonClass {
	return getNSStatusBarButtonClass()
}

type NSStatusBarButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStatusBarButtonClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStatusBarButtonClass) Alloc() NSStatusBarButton {
	rv := objc.Send[NSStatusBarButton](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The appearance and behavior of an item in the systemwide menu bar.
//
// # Instance Properties
//
//   - [NSStatusBarButton.AppearsDisabled]
//   - [NSStatusBarButton.SetAppearsDisabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBarButton
type NSStatusBarButton struct {
	NSButton
}

// NSStatusBarButtonFromID constructs a [NSStatusBarButton] from an objc.ID.
//
// The appearance and behavior of an item in the systemwide menu bar.
func NSStatusBarButtonFromID(id objc.ID) NSStatusBarButton {
	return NSStatusBarButton{NSButton: NSButtonFromID(id)}
}
// NOTE: NSStatusBarButton adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStatusBarButton] class.
//
// # Instance Properties
//
//   - [INSStatusBarButton.AppearsDisabled]
//   - [INSStatusBarButton.SetAppearsDisabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBarButton
type INSStatusBarButton interface {
	INSButton

	// Topic: Instance Properties

	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

// Init initializes the instance.
func (s NSStatusBarButton) Init() NSStatusBarButton {
	rv := objc.Send[NSStatusBarButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStatusBarButton) Autorelease() NSStatusBarButton {
	rv := objc.Send[NSStatusBarButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStatusBarButton creates a new NSStatusBarButton instance.
func NewNSStatusBarButton() NSStatusBarButton {
	class := getNSStatusBarButtonClass()
	rv := objc.Send[NSStatusBarButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a standard checkbox with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewStatusBarButtonCheckboxWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSStatusBarButton {
	rv := objc.Send[objc.ID](objc.ID(getNSStatusBarButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return NSStatusBarButtonFromID(rv)
}

// Creates a standard radio button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func NewStatusBarButtonRadioButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSStatusBarButton {
	rv := objc.Send[objc.ID](objc.ID(getNSStatusBarButtonClass().class), objc.Sel("radioButtonWithTitle:target:action:"), objc.String(title), target, action)
	return NSStatusBarButtonFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewStatusBarButtonWithCoder(coder foundation.INSCoder) NSStatusBarButton {
	instance := getNSStatusBarButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStatusBarButtonFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewStatusBarButtonWithFrame(frameRect corefoundation.CGRect) NSStatusBarButton {
	instance := getNSStatusBarButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSStatusBarButtonFromID(rv)
}

// Creates a standard push button with the image you specify.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// # Discussion
// 
// Set the image’s [AccessibilityDescription] property to ensure
// accessibility for this control.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func NewStatusBarButtonWithImageTargetAction(image INSImage, target objectivec.IObject, action objc.SEL) NSStatusBarButton {
	rv := objc.Send[objc.ID](objc.ID(getNSStatusBarButtonClass().class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return NSStatusBarButtonFromID(rv)
}

// Creates a standard push button with a title and image.
//
// title: The localized title string to display on the button.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewStatusBarButtonWithTitleImageTargetAction(title string, image INSImage, target objectivec.IObject, action objc.SEL) NSStatusBarButton {
	rv := objc.Send[objc.ID](objc.ID(getNSStatusBarButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return NSStatusBarButtonFromID(rv)
}

// Creates a standard push button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the control.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewStatusBarButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSStatusBarButton {
	rv := objc.Send[objc.ID](objc.ID(getNSStatusBarButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return NSStatusBarButtonFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSStatusBarButton/appearsDisabled
func (s NSStatusBarButton) AppearsDisabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("appearsDisabled"))
	return rv
}
func (s NSStatusBarButton) SetAppearsDisabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAppearsDisabled:"), value)
}

