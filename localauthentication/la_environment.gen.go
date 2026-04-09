// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LAEnvironment] class.
var (
	_LAEnvironmentClass     LAEnvironmentClass
	_LAEnvironmentClassOnce sync.Once
)

func getLAEnvironmentClass() LAEnvironmentClass {
	_LAEnvironmentClassOnce.Do(func() {
		_LAEnvironmentClass = LAEnvironmentClass{class: objc.GetClass("LAEnvironment")}
	})
	return _LAEnvironmentClass
}

// GetLAEnvironmentClass returns the class object for LAEnvironment.
func GetLAEnvironmentClass() LAEnvironmentClass {
	return getLAEnvironmentClass()
}

type LAEnvironmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentClass) Alloc() LAEnvironment {
	rv := objc.Send[LAEnvironment](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironment.State]: The environment state information.
//
// # Instance Methods
//
//   - [LAEnvironment.AddObserver]
//   - [LAEnvironment.RemoveObserver]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment
type LAEnvironment struct {
	objectivec.Object
}

// LAEnvironmentFromID constructs a [LAEnvironment] from an objc.ID.
func LAEnvironmentFromID(id objc.ID) LAEnvironment {
	return LAEnvironment{objectivec.Object{ID: id}}
}

// NOTE: LAEnvironment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironment] class.
//
// # Instance Properties
//
//   - [ILAEnvironment.State]: The environment state information.
//
// # Instance Methods
//
//   - [ILAEnvironment.AddObserver]
//   - [ILAEnvironment.RemoveObserver]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment
type ILAEnvironment interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The environment state information.
	State() objc.ID

	// Topic: Instance Methods

	AddObserver(observer LAEnvironmentObserver)
	RemoveObserver(observer LAEnvironmentObserver)
}

// Init initializes the instance.
func (e LAEnvironment) Init() LAEnvironment {
	rv := objc.Send[LAEnvironment](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironment) Autorelease() LAEnvironment {
	rv := objc.Send[LAEnvironment](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironment creates a new LAEnvironment instance.
func NewLAEnvironment() LAEnvironment {
	class := getLAEnvironmentClass()
	rv := objc.Send[LAEnvironment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// Adds observer to monitor changes of the environment.
//
// The observer will be held weakly so its instance should be kept alive by
// the caller.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/addObserver(_:)
func (e LAEnvironment) AddObserver(observer LAEnvironmentObserver) {
	objc.Send[objc.ID](e.ID, objc.Sel("addObserver:"), observer)
}

// # Discussion
//
// Removes the previously registered observer.
//
// If the observer is deallocated, it will be removed automatically.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/removeObserver(_:)
func (e LAEnvironment) RemoveObserver(observer LAEnvironmentObserver) {
	objc.Send[objc.ID](e.ID, objc.Sel("removeObserver:"), observer)
}

// The environment state information.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/state-swift.property
func (e LAEnvironment) State() objc.ID {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("state"))
	return rv
}

// Environment of the current user.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/currentUser
func (_LAEnvironmentClass LAEnvironmentClass) CurrentUser() LAEnvironment {
	rv := objc.Send[objc.ID](objc.ID(_LAEnvironmentClass.class), objc.Sel("currentUser"))
	return LAEnvironmentFromID(objc.ID(rv))
}
