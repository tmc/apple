// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Methods your app can implement to respond when animations start and stop.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationDelegate
type CAAnimationDelegate interface {
	objectivec.IObject
}

// CAAnimationDelegateObject wraps an existing Objective-C object that conforms to the CAAnimationDelegate protocol.
type CAAnimationDelegateObject struct {
	objectivec.Object
}
func (o CAAnimationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CAAnimationDelegateObjectFromID constructs a [CAAnimationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CAAnimationDelegateObjectFromID(id objc.ID) CAAnimationDelegateObject {
	return CAAnimationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the animation has started.
//
// anim: The [CAAnimation] object that has started.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationDelegate/animationDidStart(_:)
func (o CAAnimationDelegateObject) AnimationDidStart(anim ICAAnimation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("animationDidStart:"), anim)
	}
// Tells the delegate the animation has ended.
//
// anim: The [CAAnimation] object that has ended.
//
// flag: A flag indicating whether the animation has completed by reaching the end
// of its duration.
//
// # Discussion
// 
// The animation may have ended because it has completed its active duration
// or because it has been removed from the layer it is attached to. `flag` is
// true if the animation reached the end of its duration without being
// removed.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationDelegate/animationDidStop(_:finished:)
func (o CAAnimationDelegateObject) AnimationDidStopFinished(anim ICAAnimation, flag bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("animationDidStop:finished:"), anim, flag)
	}

// CAAnimationDelegateConfig holds optional typed callbacks for [CAAnimationDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate
type CAAnimationDelegateConfig struct {

	// Customizing Start and Stop Times
	// AnimationDidStart — Tells the delegate the animation has started.
	AnimationDidStart func(anim CAAnimation)
	// AnimationDidStopFinished — Tells the delegate the animation has ended.
	AnimationDidStopFinished func(anim CAAnimation, flag bool)
}

// NewCAAnimationDelegate creates an Objective-C object implementing the [CAAnimationDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CAAnimationDelegateObject] satisfies the [CAAnimationDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate
func NewCAAnimationDelegate(config CAAnimationDelegateConfig) CAAnimationDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCAAnimationDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AnimationDidStart != nil {
		fn := config.AnimationDidStart
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animationDidStart:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animID objc.ID) {
				anim := CAAnimationFromID(animID)
				fn(anim)
			},
		})
	}

	if config.AnimationDidStopFinished != nil {
		fn := config.AnimationDidStopFinished
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animationDidStop:finished:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animID objc.ID, flag bool) {
				anim := CAAnimationFromID(animID)
				fn(anim, flag)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CAAnimationDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCAAnimationDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CAAnimationDelegateObjectFromID(instance)
}

