// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSController] class.
var (
	_NSControllerClass     NSControllerClass
	_NSControllerClassOnce sync.Once
)

func getNSControllerClass() NSControllerClass {
	_NSControllerClassOnce.Do(func() {
		_NSControllerClass = NSControllerClass{class: objc.GetClass("NSController")}
	})
	return _NSControllerClass
}

// GetNSControllerClass returns the class object for NSController.
func GetNSControllerClass() NSControllerClass {
	return getNSControllerClass()
}

type NSControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSControllerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSControllerClass) Alloc() NSController {
	rv := objc.Send[NSController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that implements the [NSEditor] and [NSEditorRegistration]
// informal protocols required for controller classes.
//
// # Managing editing
//
//   - [NSController.CommitEditing]: Attempts to commit any pending edits.
//   - [NSController.CommitEditingWithDelegateDidCommitSelectorContextInfo]: Attempts to commit any pending changes in known editors of the receiver.
//   - [NSController.DiscardEditing]: Discards any pending changes by registered editors.
//   - [NSController.Editing]: A Boolean value indicating if any editors are registered with the controller.
//
// # Initializers
//
//   - [NSController.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSController
type NSController struct {
	objectivec.Object
}

// NSControllerFromID constructs a [NSController] from an objc.ID.
//
// An abstract class that implements the [NSEditor] and [NSEditorRegistration]
// informal protocols required for controller classes.
func NSControllerFromID(id objc.ID) NSController {
	return NSController{objectivec.Object{ID: id}}
}

// NOTE: NSController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSController] class.
//
// # Managing editing
//
//   - [INSController.CommitEditing]: Attempts to commit any pending edits.
//   - [INSController.CommitEditingWithDelegateDidCommitSelectorContextInfo]: Attempts to commit any pending changes in known editors of the receiver.
//   - [INSController.DiscardEditing]: Discards any pending changes by registered editors.
//   - [INSController.Editing]: A Boolean value indicating if any editors are registered with the controller.
//
// # Initializers
//
//   - [INSController.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSController
type INSController interface {
	objectivec.IObject
	NSEditorRegistration

	// Topic: Managing editing

	// Attempts to commit any pending edits.
	CommitEditing() bool
	// Attempts to commit any pending changes in known editors of the receiver.
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo uintptr)
	// Discards any pending changes by registered editors.
	DiscardEditing()
	// A Boolean value indicating if any editors are registered with the controller.
	Editing() bool

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSController

	CommitEditingAndReturnError() (bool, error)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSController) Init() NSController {
	rv := objc.Send[NSController](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSController) Autorelease() NSController {
	rv := objc.Send[NSController](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSController creates a new NSController instance.
func NewNSController() NSController {
	class := getNSControllerClass()
	rv := objc.Send[NSController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSController/init(coder:)
func NewControllerWithCoder(coder foundation.INSCoder) NSController {
	instance := getNSControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSControllerFromID(rv)
}

// Invoked to inform the receiver that `editor` has uncommitted changes that
// can affect the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/objectDidBeginEditing(_:)
func (c NSController) ObjectDidBeginEditing(editor NSEditor) {
	objc.Send[objc.ID](c.ID, objc.Sel("objectDidBeginEditing:"), editor)
}

// Invoked to inform the receiver that `editor` has committed or discarded its
// changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/objectDidEndEditing(_:)
func (c NSController) ObjectDidEndEditing(editor NSEditor) {
	objc.Send[objc.ID](c.ID, objc.Sel("objectDidEndEditing:"), editor)
}

// Attempts to commit any pending edits.
//
// # Return Value
//
// true if successful or no edits were pending.
//
// # Discussion
//
// The receiver invokes [commitEditing] on any current editors, returning
// their response. A commit is denied if the receiver fails to apply the
// changes to the model object, perhaps due to a validation error.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/commitEditing()
//
// [commitEditing]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitEditing
func (c NSController) CommitEditing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("commitEditing"))
	return rv
}

// Attempts to commit any pending changes in known editors of the receiver.
//
// delegate: An object that can serve as the receiver’s delegate. It should implement
// the method specified by `didCommitSelector`.
//
// didCommitSelector: A selector that is invoked on delegate. The method specified by the
// selector must have the same signature as the following method:
//
// contextInfo: Contextual information that is sent as the `contextInfo` argument to
// delegate when `didCommitSelector` is invoked.
//
// # Discussion
//
// Provides support for the NSEditor informal protocol. This method attempts
// to commit pending changes in known editors. Known editors are either
// instances of a subclass of [NSController] or (more rarely) user interface
// controls that may contain pending edits—such as text fields—that
// registered with the context using [ObjectDidBeginEditing] and have not yet
// unregistered using a subsequent invocation of [ObjectDidEndEditing].
//
// The receiver iterates through the array of its known editors and invokes
// `commitEditing` on each. The receiver then sends the message specified by
// the `didCommitSelector` selector to the specified delegate.
//
// The `didCommit` argument is the value returned by the editor specified by
// `editor` from the `commitEditing` message. The `contextInfo` argument is
// the same value specified as the `contextInfo` parameter—you may use this
// value however you wish.
//
// If an error occurs while attempting to commit, for example if key-value
// coding validation fails, your implementation of this method should
// typically send the view in which editing is being performed a “ message,
// specifying the view’s containing window.
//
// You may find this method useful in some situations (typically if you are
// using Cocoa Bindings) when you want to ensure that pending changes are
// applied before a change in user interface state. For example, you may need
// to ensure that changes pending in a text field are applied before a window
// is closed. See also [CommitEditing] which performs a similar function but
// which allows you to handle any errors directly, although it provides no
// information beyond simple success/failure.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/commitEditing(withDelegate:didCommit:contextInfo:)
func (c NSController) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

// Discards any pending changes by registered editors.
//
// # Discussion
//
// The receiver invokes [discardEditing] on any current editors.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/discardEditing()
//
// [discardEditing]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/discardEditing
func (c NSController) DiscardEditing() {
	objc.Send[objc.ID](c.ID, objc.Sel("discardEditing"))
}

// See: https://developer.apple.com/documentation/AppKit/NSController/init(coder:)
func (c NSController) InitWithCoder(coder foundation.INSCoder) NSController {
	rv := objc.Send[NSController](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditingWithoutPresentingError()
func (c NSController) CommitEditingAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("commitEditingAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("commitEditingAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
func (c NSController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value indicating if any editors are registered with the
// controller.
//
// # Discussion
//
// The value of this property is true when an editor is registered with the
// controller object or false when no editor is registered.
//
// See: https://developer.apple.com/documentation/AppKit/NSController/isEditing
func (c NSController) Editing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEditing"))
	return rv
}

// Protocol methods for NSEditorRegistration
