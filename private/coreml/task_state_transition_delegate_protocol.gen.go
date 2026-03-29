// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// MLTaskStateTransitionDelegate protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate
type MLTaskStateTransitionDelegate interface {
	objectivec.IObject

	// OnCancellation protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onCancellation
	OnCancellation()
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

// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onCancellation
func (o MLTaskStateTransitionDelegateObject) OnCancellation() {
	objc.Send[struct{}](o.ID, objc.Sel("onCancellation"))
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onCompletionWithTaskContext:
func (o MLTaskStateTransitionDelegateObject) OnCompletionWithTaskContext(context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("onCompletionWithTaskContext:"), context)
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onFailureWithTaskContext:
func (o MLTaskStateTransitionDelegateObject) OnFailureWithTaskContext(context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("onFailureWithTaskContext:"), context)
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onResumptionWithTaskContext:
func (o MLTaskStateTransitionDelegateObject) OnResumptionWithTaskContext(context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("onResumptionWithTaskContext:"), context)
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLTaskStateTransitionDelegate/onSuspensionWithTaskContext:
func (o MLTaskStateTransitionDelegateObject) OnSuspensionWithTaskContext(context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("onSuspensionWithTaskContext:"), context)
	}

// MLTaskStateTransitionDelegateConfig holds optional typed callbacks for [MLTaskStateTransitionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coreml/mltaskstatetransitiondelegate
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
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coreml/mltaskstatetransitiondelegate
func NewMLTaskStateTransitionDelegate(config MLTaskStateTransitionDelegateConfig) MLTaskStateTransitionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoMLTaskStateTransitionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OnCancellation != nil {
		fn := config.OnCancellation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("onCancellation"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				fn()
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

