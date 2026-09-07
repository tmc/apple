// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Environment] class.
var (
	_EnvironmentClass     EnvironmentClass
	_EnvironmentClassOnce sync.Once
)

func getEnvironmentClass() EnvironmentClass {
	_EnvironmentClassOnce.Do(func() {
		_EnvironmentClass = EnvironmentClass{class: objc.GetClass("_TtCC8SkyLight27SystemGestureEventProcessor11Environment")}
	})
	return _EnvironmentClass
}

// GetEnvironmentClass returns the class object for _TtCC8SkyLight27SystemGestureEventProcessor11Environment.
func GetEnvironmentClass() EnvironmentClass {
	return getEnvironmentClass()
}

type EnvironmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EnvironmentClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EnvironmentClass) Alloc() Environment {
	rv := objc.SendIfResponds[Environment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type Environment struct {
	objectivec.Object
}

// EnvironmentFromID constructs a [Environment] from an objc.ID.
func EnvironmentFromID(id objc.ID) Environment {
	return Environment{objectivec.Object{ID: id}}
}

// Ensure Environment implements IEnvironment.
var _ IEnvironment = Environment{}

// An interface definition for the [Environment] class.
type IEnvironment interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e Environment) Init() Environment {
	rv := objc.SendIfResponds[Environment](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e Environment) Autorelease() Environment {
	rv := objc.SendIfResponds[Environment](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironment creates a new Environment instance.
func NewEnvironment() Environment {
	class := getEnvironmentClass()
	rv := objc.SendIfResponds[Environment](objc.ID(class.class), objc.Sel("new"))
	return rv
}
