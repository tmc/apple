// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods to take custom control over the view’s render rate.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKViewDelegate
type SKViewDelegate interface {
	objectivec.IObject
}

// SKViewDelegateObject wraps an existing Objective-C object that conforms to the SKViewDelegate protocol.
type SKViewDelegateObject struct {
	objectivec.Object
}

func (o SKViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SKViewDelegateObjectFromID constructs a [SKViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SKViewDelegateObjectFromID(id objc.ID) SKViewDelegateObject {
	return SKViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Specifies whether the view should render at the given time.
//
// view: The SKView.
//
// time: The target time.
//
// # Return Value
//
// Return `true` to initiate an update and render for the target time. Return
// `false` to skip the update and render for the target time.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKViewDelegate/view(_:shouldRenderAtTime:)
func (o SKViewDelegateObject) ViewShouldRenderAtTime(view ISKView, time foundation.NSTimeInterval) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("view:shouldRenderAtTime:"), view, time)
	return rv
}

// SKViewDelegateConfig holds optional typed callbacks for [SKViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skviewdelegate
type SKViewDelegateConfig struct {

	// Instance Methods
	// ViewShouldRenderAtTime — Specifies whether the view should render at the given time.
	ViewShouldRenderAtTime func(view SKView, time foundation.NSTimeInterval) bool
}

// NewSKViewDelegate creates an Objective-C object implementing the [SKViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SKViewDelegateObject] satisfies the [SKViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skviewdelegate
func NewSKViewDelegate(config SKViewDelegateConfig) SKViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSKViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ViewShouldRenderAtTime != nil {
		fn := config.ViewShouldRenderAtTime
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("view:shouldRenderAtTime:"),
			Fn: func(self objc.ID, _cmd objc.SEL, viewID objc.ID, time foundation.NSTimeInterval) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKViewDelegate", "view:shouldRenderAtTime:")
					}
				}()
				view := SKViewFromID(viewID)
				_delegateResult := fn(view, time)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SKViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSKViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SKViewDelegateObjectFromID(instance)
}
