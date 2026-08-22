// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods that, when implemented, allow any class to participate in the SpriteKit render loop callbacks.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate
type SKSceneDelegate interface {
	objectivec.IObject
}

// SKSceneDelegateObject wraps an existing Objective-C object that conforms to the SKSceneDelegate protocol.
type SKSceneDelegateObject struct {
	objectivec.Object
}

func (o SKSceneDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SKSceneDelegateObjectFromID constructs a [SKSceneDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SKSceneDelegateObjectFromID(id objc.ID) SKSceneDelegateObject {
	return SKSceneDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells you to perform any app specific logic to update your scene.
//
// currentTime: The current system time.
//
// scene: The scene that is being animated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// This is the first method called when animating the scene, before any
// actions are evaluated and before any physics are simulated.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate/update(_:for:)
func (o SKSceneDelegateObject) UpdateForScene(currentTime foundation.NSTimeInterval, scene ISKScene) {
	objc.Send[struct{}](o.ID, objc.Sel("update:forScene:"), currentTime, scene)
}

// Tells you to peform any necessary logic after scene actions are evaluated.
//
// scene: The scene that is being animated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after any actions have been evaluated by nodes in the scene
// but before any physics are simulated.
//
// Any additional actions applied are not evaluated until the next update.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate/didEvaluateActions(for:)
func (o SKSceneDelegateObject) DidEvaluateActionsForScene(scene ISKScene) {
	objc.Send[struct{}](o.ID, objc.Sel("didEvaluateActionsForScene:"), scene)
}

// Tells you to peform any necessary logic after physics simulations are
// performed.
//
// scene: The scene that is being animated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after physics has been simulated in the scene.
//
// Any additional actions applied are not evaluated until the next update.
//
// Any changes to physics bodies are not simulated until the next update.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate/didSimulatePhysics(for:)
func (o SKSceneDelegateObject) DidSimulatePhysicsForScene(scene ISKScene) {
	objc.Send[struct{}](o.ID, objc.Sel("didSimulatePhysicsForScene:"), scene)
}

// Tells you to peform any necessary logic after constraints are applied.
//
// scene: The scene that is being animated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after any enabled constraints in the scene have been applied.
//
// Any additional actions applied are not evaluated until the next update.
//
// Any changes to physics bodies is not simulated until the next update.
//
// Any changes to constraints will not be applied until the next update.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate/didApplyConstraints(for:)
func (o SKSceneDelegateObject) DidApplyConstraintsForScene(scene ISKScene) {
	objc.Send[struct{}](o.ID, objc.Sel("didApplyConstraintsForScene:"), scene)
}

// Tells you to peform any necessary logic after the scene has finished all of
// the steps required to process animations.
//
// scene: The scene that is being animated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after all update logic has been completed and before the scene
// is rendered.
//
// Any additional actions applied are not evaluated until the next update.
//
// Any changes to physics bodies are not simulated until the next update.
//
// Any changes to constraints will not be applied until the next update.
//
// No further update logic will be applied to the scene after this call. Any
// values set on nodes here will be used when the scene is rendered for the
// current frame.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSceneDelegate/didFinishUpdate(for:)
func (o SKSceneDelegateObject) DidFinishUpdateForScene(scene ISKScene) {
	objc.Send[struct{}](o.ID, objc.Sel("didFinishUpdateForScene:"), scene)
}

// SKSceneDelegateConfig holds optional typed callbacks for [SKSceneDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skscenedelegate
type SKSceneDelegateConfig struct {

	// Other Methods
	// UpdateForScene — Tells you to perform any app specific logic to update your scene.
	UpdateForScene func(currentTime foundation.NSTimeInterval, scene SKScene)
	// DidEvaluateActionsForScene — Tells you to peform any necessary logic after scene actions are evaluated.
	DidEvaluateActionsForScene func(scene SKScene)
	// DidSimulatePhysicsForScene — Tells you to peform any necessary logic after physics simulations are performed.
	DidSimulatePhysicsForScene func(scene SKScene)
	// DidApplyConstraintsForScene — Tells you to peform any necessary logic after constraints are applied.
	DidApplyConstraintsForScene func(scene SKScene)
	// DidFinishUpdateForScene — Tells you to peform any necessary logic after the scene has finished all of the steps required to process animations.
	DidFinishUpdateForScene func(scene SKScene)
}

// NewSKSceneDelegate creates an Objective-C object implementing the [SKSceneDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SKSceneDelegateObject] satisfies the [SKSceneDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skscenedelegate
func NewSKSceneDelegate(config SKSceneDelegateConfig) SKSceneDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSKSceneDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UpdateForScene != nil {
		fn := config.UpdateForScene
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("update:forScene:"),
			Fn: func(self objc.ID, _cmd objc.SEL, currentTime foundation.NSTimeInterval, sceneID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKSceneDelegate", "update:forScene:")
					}
				}()
				scene := SKSceneFromID(sceneID)
				fn(currentTime, scene)
				_delegateDone = true
			},
		})
	}

	if config.DidEvaluateActionsForScene != nil {
		fn := config.DidEvaluateActionsForScene
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didEvaluateActionsForScene:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sceneID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKSceneDelegate", "didEvaluateActionsForScene:")
					}
				}()
				scene := SKSceneFromID(sceneID)
				fn(scene)
				_delegateDone = true
			},
		})
	}

	if config.DidSimulatePhysicsForScene != nil {
		fn := config.DidSimulatePhysicsForScene
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didSimulatePhysicsForScene:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sceneID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKSceneDelegate", "didSimulatePhysicsForScene:")
					}
				}()
				scene := SKSceneFromID(sceneID)
				fn(scene)
				_delegateDone = true
			},
		})
	}

	if config.DidApplyConstraintsForScene != nil {
		fn := config.DidApplyConstraintsForScene
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didApplyConstraintsForScene:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sceneID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKSceneDelegate", "didApplyConstraintsForScene:")
					}
				}()
				scene := SKSceneFromID(sceneID)
				fn(scene)
				_delegateDone = true
			},
		})
	}

	if config.DidFinishUpdateForScene != nil {
		fn := config.DidFinishUpdateForScene
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didFinishUpdateForScene:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sceneID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKSceneDelegate", "didFinishUpdateForScene:")
					}
				}()
				scene := SKSceneFromID(sceneID)
				fn(scene)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SKSceneDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSKSceneDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SKSceneDelegateObjectFromID(instance)
}
