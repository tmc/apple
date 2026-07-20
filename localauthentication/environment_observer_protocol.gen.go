// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// LAEnvironmentObserver protocol.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Observer
type LAEnvironmentObserver interface {
	objectivec.IObject
}

// LAEnvironmentObserverObject wraps an existing Objective-C object that conforms to the LAEnvironmentObserver protocol.
type LAEnvironmentObserverObject struct {
	objectivec.Object
}

func (o LAEnvironmentObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// LAEnvironmentObserverObjectFromID constructs a [LAEnvironmentObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func LAEnvironmentObserverObjectFromID(id objc.ID) LAEnvironmentObserverObject {
	return LAEnvironmentObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when there has been a change in the environment.
//
// oldState: The old environment state (before update)
//
// # Discussion
//
// Invoked on a queue private to LocalAuthentication framework. At the moment
// of invocation of this method, @c LAEnvironment.state already contains the
// new updated state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Observer/environment(_:stateDidChangeFromOldState:)
func (o LAEnvironmentObserverObject) EnvironmentStateDidChangeFromOldState(environment ILAEnvironment, oldState objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("environment:stateDidChangeFromOldState:"), environment, oldState)
}
