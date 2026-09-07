// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// MLTaskStateTransitionDelegate protocol.
type MLTaskStateTransitionDelegate interface {
	objectivec.IObject

	// OnCancellation protocol.
	OnCancellation()

	// OnCompletionWithTaskContext protocol.
	OnCompletionWithTaskContext(context objectivec.IObject)

	// OnFailureWithTaskContext protocol.
	OnFailureWithTaskContext(context objectivec.IObject)

	// OnResumptionWithTaskContext protocol.
	OnResumptionWithTaskContext(context objectivec.IObject)

	// OnSuspensionWithTaskContext protocol.
	OnSuspensionWithTaskContext(context objectivec.IObject)
}

// MLTaskStateTransitionDelegateObject wraps an existing Objective-C object that conforms to the MLTaskStateTransitionDelegate protocol.
type MLTaskStateTransitionDelegateObject struct {
	objectivec.Object
}

func (o MLTaskStateTransitionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLTaskStateTransitionDelegateObjectFromID constructs a [MLTaskStateTransitionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLTaskStateTransitionDelegateObjectFromID(id objc.ID) MLTaskStateTransitionDelegateObject {
	return MLTaskStateTransitionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLTaskStateTransitionDelegateObject) OnCancellation() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("onCancellation"))
}
func (o MLTaskStateTransitionDelegateObject) OnCompletionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("onCompletionWithTaskContext:"), context)
}
func (o MLTaskStateTransitionDelegateObject) OnFailureWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("onFailureWithTaskContext:"), context)
}
func (o MLTaskStateTransitionDelegateObject) OnResumptionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("onResumptionWithTaskContext:"), context)
}
func (o MLTaskStateTransitionDelegateObject) OnSuspensionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("onSuspensionWithTaskContext:"), context)
}

// MLTaskStateTransitionDelegateConfig holds optional typed callbacks for [MLTaskStateTransitionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
type MLTaskStateTransitionDelegateConfig struct {

	// Other Methods
	OnCancellation func()
}

// NewMLTaskStateTransitionDelegate creates an Objective-C object implementing the [MLTaskStateTransitionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [MLTaskStateTransitionDelegateObject] satisfies the [MLTaskStateTransitionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
func NewMLTaskStateTransitionDelegate(config MLTaskStateTransitionDelegateConfig) MLTaskStateTransitionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoMLTaskStateTransitionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OnCancellation != nil {
		fn := config.OnCancellation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("onCancellation"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("MLTaskStateTransitionDelegate", "onCancellation")
					}
				}()
				fn()
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("MLTaskStateTransitionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewMLTaskStateTransitionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return MLTaskStateTransitionDelegateObjectFromID(instance)
}
