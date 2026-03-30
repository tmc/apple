// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods that a popover delegate can implement to provide additional or custom functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate
type NSPopoverDelegate interface {
	objectivec.IObject
}

// NSPopoverDelegateObject wraps an existing Objective-C object that conforms to the NSPopoverDelegate protocol.
type NSPopoverDelegateObject struct {
	objectivec.Object
}

func (o NSPopoverDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPopoverDelegateObjectFromID constructs a [NSPopoverDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPopoverDelegateObjectFromID(id objc.ID) NSPopoverDelegateObject {
	return NSPopoverDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Detaches the popover creating a window containing the content.
//
// popover: The popover.
//
// # Return Value
//
// Returns a window instance to which the popover should be detached.
//
// # Discussion
//
// You should not remove the popover’s content view as part of your
// implementation of this method.
//
// The popover and the detachable window may be shown at the same time and
// therefore cannot share a content view or content view controller.
//
// If the popover and the detachable window should have the same content, you
// should define the content in a separate nib file and use a view controller
// to instantiate separate copies of the content for the popover and the
// detachable window.
//
// The popover will animate to appear as though it morphs into the detachable
// window (unless the popover’s [Animates] property is set to false). The
// exact animation used is not guaranteed.
//
// If there is no delegate, the delegate does not implement this method, or
// the delegate returns nil, the popup will not be displayed detached.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/detachableWindow(for:)
func (o NSPopoverDelegateObject) DetachableWindowForPopover(popover INSPopover) INSWindow {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("detachableWindowForPopover:"), popover)
	return NSWindowFromID(rv)
}

// Allows a delegate to override a close request.
//
// popover: The popover.
//
// # Return Value
//
// true if the popover should close, false otherwise.
//
// # Discussion
//
// The popover invokes this method on its delegate whenever it is about to
// close. This gives the delegate a chance to override the close.
//
// If there is no delegate or the delegate does not implement this method the
// default behavior is that the popover will close.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverShouldClose(_:)
func (o NSPopoverDelegateObject) PopoverShouldClose(popover INSPopover) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("popoverShouldClose:"), popover)
	return rv
}

// Invoked when the popover will show.
//
// # Discussion
//
// Invoked on the delegate when the [willShowNotification] notification is
// sent.
//
// This method will also be invoked on the delegate’s popover, if the method
// has been implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverWillShow(_:)
//
// [willShowNotification]: https://developer.apple.com/documentation/AppKit/NSPopover/willShowNotification
func (o NSPopoverDelegateObject) PopoverWillShow(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("popoverWillShow:"), notification)
}

// Invoked when the popover has been shown.
//
// # Discussion
//
// Invoked on the delegate when the [didShowNotification] notification is
// sent.
//
// This method will also be invoked on the delegate’s popover, if the method
// has been implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverDidShow(_:)
//
// [didShowNotification]: https://developer.apple.com/documentation/AppKit/NSPopover/didShowNotification
func (o NSPopoverDelegateObject) PopoverDidShow(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("popoverDidShow:"), notification)
}

// Invoked when the popover is about to close.
//
// # Discussion
//
// Invoked on the delegate when the [willCloseNotification] notification is
// sent.
//
// This method will also be invoked on the delegate’s popover, if the method
// has been implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverWillClose(_:)
//
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSPopover/willCloseNotification
func (o NSPopoverDelegateObject) PopoverWillClose(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("popoverWillClose:"), notification)
}

// Invoked when the popover did close.
//
// # Discussion
//
// Invoked on the delegate when the [didCloseNotification] notification is
// sent.
//
// This method will also be invoked on the delegate’s popover, if the method
// has been implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverDidClose(_:)
//
// [didCloseNotification]: https://developer.apple.com/documentation/AppKit/NSPopover/didCloseNotification
func (o NSPopoverDelegateObject) PopoverDidClose(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("popoverDidClose:"), notification)
}

// Indicates that a popover has been released while it’s in an implicitly
// detached state.
//
// popover: The popover that detached from its anchor view and is not closing.
//
// # Discussion
//
// This method is not called when the popover’s detached window is returned
// by [DetachableWindowForPopover].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverDidDetach(_:)
func (o NSPopoverDelegateObject) PopoverDidDetach(popover INSPopover) {
	objc.Send[struct{}](o.ID, objc.Sel("popoverDidDetach:"), popover)
}

// Returns a Boolean value that indicates whether a popover should detach from
// its positioning view and become a separate window.
//
// popover: The popover that may be detached.
//
// # Discussion
//
// If you don’t implement this method, it returns false by default. If you
// return true from this method, but you don’t implement
// [DetachableWindowForPopover] or you implement it to return `nil`, a
// detachable window is created with the popover’s [ContentViewController].
//
// An automatically created window has the same appearance as the detached
// popover. For example, if the popover’s [ContentViewController] has a
// title, it will be bound to and displayed as the title of the detached
// window. When a popover is released in a detached state, it calls
// [PopoverDidDetach] on the delegate. When a detached popover is closed,
// calls to [PopoverShouldClose], [PopoverWillClose], and [PopoverDidClose],
// in addition to the related notifications, specify the reason [standard].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverDelegate/popoverShouldDetach(_:)
//
// [standard]: https://developer.apple.com/documentation/AppKit/NSPopover/CloseReason/standard
func (o NSPopoverDelegateObject) PopoverShouldDetach(popover INSPopover) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("popoverShouldDetach:"), popover)
	return rv
}

// NSPopoverDelegateConfig holds optional typed callbacks for [NSPopoverDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspopoverdelegate
type NSPopoverDelegateConfig struct {

	// Popover Visibility
	// ShouldClose — Allows a delegate to override a close request.
	ShouldClose func(popover NSPopover) bool
	// WillShow — Invoked when the popover will show.
	WillShow func(notification foundation.NSNotification)
	// DidShow — Invoked when the popover has been shown.
	DidShow func(notification foundation.NSNotification)
	// WillClose — Invoked when the popover is about to close.
	WillClose func(notification foundation.NSNotification)
	// DidClose — Invoked when the popover did close.
	DidClose func(notification foundation.NSNotification)
	// DidDetach — Indicates that a popover has been released while it’s in an implicitly detached state.
	DidDetach func(popover NSPopover)
	// ShouldDetach — Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window.
	ShouldDetach func(popover NSPopover) bool

	// Other Methods
	// DetachableWindowForPopover — Detaches the popover creating a window containing the content.
	DetachableWindowForPopover func(popover NSPopover) NSWindow
}

// NewNSPopoverDelegate creates an Objective-C object implementing the [NSPopoverDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSPopoverDelegateObject] satisfies the [NSPopoverDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspopoverdelegate
func NewNSPopoverDelegate(config NSPopoverDelegateConfig) NSPopoverDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSPopoverDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DetachableWindowForPopover != nil {
		fn := config.DetachableWindowForPopover
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("detachableWindowForPopover:"),
			Fn: func(self objc.ID, _cmd objc.SEL, popoverID objc.ID) objc.ID {
				popover := NSPopoverFromID(popoverID)
				return fn(popover).GetID()
			},
		})
	}

	if config.ShouldClose != nil {
		fn := config.ShouldClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverShouldClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, popoverID objc.ID) bool {
				popover := NSPopoverFromID(popoverID)
				return fn(popover)
			},
		})
	}

	if config.WillShow != nil {
		fn := config.WillShow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverWillShow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidShow != nil {
		fn := config.DidShow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverDidShow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillClose != nil {
		fn := config.WillClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverWillClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidClose != nil {
		fn := config.DidClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverDidClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidDetach != nil {
		fn := config.DidDetach
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverDidDetach:"),
			Fn: func(self objc.ID, _cmd objc.SEL, popoverID objc.ID) {
				popover := NSPopoverFromID(popoverID)
				fn(popover)
			},
		})
	}

	if config.ShouldDetach != nil {
		fn := config.ShouldDetach
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("popoverShouldDetach:"),
			Fn: func(self objc.ID, _cmd objc.SEL, popoverID objc.ID) bool {
				popover := NSPopoverFromID(popoverID)
				return fn(popover)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSPopoverDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSPopoverDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSPopoverDelegateObjectFromID(instance)
}
