// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods your app can implement to respond when physics bodies come into contact.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContactDelegate
type SKPhysicsContactDelegate interface {
	objectivec.IObject
}

// SKPhysicsContactDelegateObject wraps an existing Objective-C object that conforms to the SKPhysicsContactDelegate protocol.
type SKPhysicsContactDelegateObject struct {
	objectivec.Object
}

func (o SKPhysicsContactDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SKPhysicsContactDelegateObjectFromID constructs a [SKPhysicsContactDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SKPhysicsContactDelegateObjectFromID(id objc.ID) SKPhysicsContactDelegateObject {
	return SKPhysicsContactDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when two bodies first contact each other.
//
// contact: An object that describes the contact.
//
// # Discussion
//
// The two physics bodies described in the contact parameter are not passed in
// a guaranteed order. The following code shows how you might respond to the
// beginning of a contact event to execute code if either physics body is
// owned by a node with the name `ground`.
//
// Listing 1. Responding to a contact event.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContactDelegate/didBegin(_:)
func (o SKPhysicsContactDelegateObject) DidBeginContact(contact ISKPhysicsContact) {
	objc.Send[struct{}](o.ID, objc.Sel("didBeginContact:"), contact)
}

// Called when the contact ends between two physics bodies.
//
// contact: An object that describes the contact.
//
// # Discussion
//
// The two physics bodies described in the contact parameter are not passed in
// a guaranteed order. The following code shows how you might respond to the
// end of a contact event to execute code if either physics body is owned by a
// node with the name `ground`.
//
// Listing 1. Responding to a contact event
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContactDelegate/didEnd(_:)
func (o SKPhysicsContactDelegateObject) DidEndContact(contact ISKPhysicsContact) {
	objc.Send[struct{}](o.ID, objc.Sel("didEndContact:"), contact)
}

// SKPhysicsContactDelegateConfig holds optional typed callbacks for [SKPhysicsContactDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skphysicscontactdelegate
type SKPhysicsContactDelegateConfig struct {

	// Other Methods
	// DidBeginContact — Called when two bodies first contact each other.
	DidBeginContact func(contact SKPhysicsContact)
	// DidEndContact — Called when the contact ends between two physics bodies.
	DidEndContact func(contact SKPhysicsContact)
}

// NewSKPhysicsContactDelegate creates an Objective-C object implementing the [SKPhysicsContactDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SKPhysicsContactDelegateObject] satisfies the [SKPhysicsContactDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/spritekit/skphysicscontactdelegate
func NewSKPhysicsContactDelegate(config SKPhysicsContactDelegateConfig) SKPhysicsContactDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSKPhysicsContactDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DidBeginContact != nil {
		fn := config.DidBeginContact
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didBeginContact:"),
			Fn: func(self objc.ID, _cmd objc.SEL, contactID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKPhysicsContactDelegate", "didBeginContact:")
					}
				}()
				contact := SKPhysicsContactFromID(contactID)
				fn(contact)
				_delegateDone = true
			},
		})
	}

	if config.DidEndContact != nil {
		fn := config.DidEndContact
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didEndContact:"),
			Fn: func(self objc.ID, _cmd objc.SEL, contactID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SKPhysicsContactDelegate", "didEndContact:")
					}
				}()
				contact := SKPhysicsContactFromID(contactID)
				fn(contact)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SKPhysicsContactDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSKPhysicsContactDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SKPhysicsContactDelegateObjectFromID(instance)
}
