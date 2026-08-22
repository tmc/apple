// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of methods that a scrubber delegate implements to respond to user interactions.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate
type NSScrubberDelegate interface {
	objectivec.IObject
}

// NSScrubberDelegateObject wraps an existing Objective-C object that conforms to the NSScrubberDelegate protocol.
type NSScrubberDelegateObject struct {
	objectivec.Object
}

func (o NSScrubberDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSScrubberDelegateObjectFromID constructs a [NSScrubberDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSScrubberDelegateObjectFromID(id objc.ID) NSScrubberDelegateObject {
	return NSScrubberDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the item at the specified index was selected.
//
// scrubber: The scrubber object that is notifying you of the selection change.
//
// selectedIndex: The index of the item that was selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didSelectItemAt:)
func (o NSScrubberDelegateObject) ScrubberDidSelectItemAtIndex(scrubber INSScrubber, selectedIndex int) {
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didSelectItemAtIndex:"), scrubber, selectedIndex)
}

// Tells the delegate that the item at the specified index was highlighted.
//
// scrubber: The scrubber object that is notifying you of the highlight change.
//
// highlightedIndex: The index of the item that is now highlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didHighlightItemAt:)
func (o NSScrubberDelegateObject) ScrubberDidHighlightItemAtIndex(scrubber INSScrubber, highlightedIndex int) {
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didHighlightItemAtIndex:"), scrubber, highlightedIndex)
}

// Tells the delegate that the range of items currently visible in the
// scrubber has changed.
//
// scrubber: The scrubber object that is notifying you of the change in the range of
// items that are currently visible.
//
// visibleRange: The range of items that are now visible in the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/scrubber(_:didChangeVisibleRange:)
func (o NSScrubberDelegateObject) ScrubberDidChangeVisibleRange(scrubber INSScrubber, visibleRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("scrubber:didChangeVisibleRange:"), scrubber, visibleRange)
}

// Tells the delegate that the user is panning or scrolling the scrubber.
//
// scrubber: The scrubber the user is interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didBeginInteracting(with:)
func (o NSScrubberDelegateObject) DidBeginInteractingWithScrubber(scrubber INSScrubber) {
	objc.Send[struct{}](o.ID, objc.Sel("didBeginInteractingWithScrubber:"), scrubber)
}

// Tells the delegate that a pan or scroll interaction with the scrubber has
// ended.
//
// scrubber: The scrubber the user was interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didFinishInteracting(with:)
func (o NSScrubberDelegateObject) DidFinishInteractingWithScrubber(scrubber INSScrubber) {
	objc.Send[struct{}](o.ID, objc.Sel("didFinishInteractingWithScrubber:"), scrubber)
}

// Tells the delegate that a user interaction with the scrubber has been
// canceled.
//
// scrubber: The scrubber the user is interacting with.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDelegate/didCancelInteracting(with:)
func (o NSScrubberDelegateObject) DidCancelInteractingWithScrubber(scrubber INSScrubber) {
	objc.Send[struct{}](o.ID, objc.Sel("didCancelInteractingWithScrubber:"), scrubber)
}

// NSScrubberDelegateConfig holds optional typed callbacks for [NSScrubberDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate
type NSScrubberDelegateConfig struct {

	// Handling scrubber scrolling
	// ScrubberDidChangeVisibleRange — Tells the delegate that the range of items currently visible in the scrubber has changed.
	ScrubberDidChangeVisibleRange func(scrubber NSScrubber, visibleRange foundation.NSRange)

	// Other Methods
	// ScrubberDidSelectItemAtIndex — Tells the delegate that the item at the specified index was selected.
	ScrubberDidSelectItemAtIndex func(scrubber NSScrubber, selectedIndex int)
	// ScrubberDidHighlightItemAtIndex — Tells the delegate that the item at the specified index was highlighted.
	ScrubberDidHighlightItemAtIndex func(scrubber NSScrubber, highlightedIndex int)
	// DidBeginInteractingWithScrubber — Tells the delegate that the user is panning or scrolling the scrubber.
	DidBeginInteractingWithScrubber func(scrubber NSScrubber)
	// DidFinishInteractingWithScrubber — Tells the delegate that a pan or scroll interaction with the scrubber has ended.
	DidFinishInteractingWithScrubber func(scrubber NSScrubber)
	// DidCancelInteractingWithScrubber — Tells the delegate that a user interaction with the scrubber has been canceled.
	DidCancelInteractingWithScrubber func(scrubber NSScrubber)
}

// NewNSScrubberDelegate creates an Objective-C object implementing the [NSScrubberDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSScrubberDelegateObject] satisfies the [NSScrubberDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate
func NewNSScrubberDelegate(config NSScrubberDelegateConfig) NSScrubberDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSScrubberDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ScrubberDidSelectItemAtIndex != nil {
		fn := config.ScrubberDidSelectItemAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scrubber:didSelectItemAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID, selectedIndex int) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "scrubber:didSelectItemAtIndex:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber, selectedIndex)
				_delegateDone = true
			},
		})
	}

	if config.ScrubberDidHighlightItemAtIndex != nil {
		fn := config.ScrubberDidHighlightItemAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scrubber:didHighlightItemAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID, highlightedIndex int) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "scrubber:didHighlightItemAtIndex:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber, highlightedIndex)
				_delegateDone = true
			},
		})
	}

	if config.ScrubberDidChangeVisibleRange != nil {
		fn := config.ScrubberDidChangeVisibleRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scrubber:didChangeVisibleRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID, visibleRange foundation.NSRange) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "scrubber:didChangeVisibleRange:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber, visibleRange)
				_delegateDone = true
			},
		})
	}

	if config.DidBeginInteractingWithScrubber != nil {
		fn := config.DidBeginInteractingWithScrubber
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didBeginInteractingWithScrubber:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "didBeginInteractingWithScrubber:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber)
				_delegateDone = true
			},
		})
	}

	if config.DidFinishInteractingWithScrubber != nil {
		fn := config.DidFinishInteractingWithScrubber
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didFinishInteractingWithScrubber:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "didFinishInteractingWithScrubber:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber)
				_delegateDone = true
			},
		})
	}

	if config.DidCancelInteractingWithScrubber != nil {
		fn := config.DidCancelInteractingWithScrubber
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didCancelInteractingWithScrubber:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scrubberID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSScrubberDelegate", "didCancelInteractingWithScrubber:")
					}
				}()
				scrubber := NSScrubberFromID(scrubberID)
				fn(scrubber)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSScrubberDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSScrubberDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSScrubberDelegateObjectFromID(instance)
}
