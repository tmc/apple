// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SpringBehavior] class.
var (
	_SpringBehaviorClass     SpringBehaviorClass
	_SpringBehaviorClassOnce sync.Once
)

func getSpringBehaviorClass() SpringBehaviorClass {
	_SpringBehaviorClassOnce.Do(func() {
		_SpringBehaviorClass = SpringBehaviorClass{class: objc.GetClass("_TtCV8SkyLight31PKGSystemStatusIndicatorsSpringP33_87FE1EE3101D9959C103B7B6E6890EE214SpringBehavior")}
	})
	return _SpringBehaviorClass
}

// GetSpringBehaviorClass returns the class object for _TtCV8SkyLight31PKGSystemStatusIndicatorsSpringP33_87FE1EE3101D9959C103B7B6E6890EE214SpringBehavior.
func GetSpringBehaviorClass() SpringBehaviorClass {
	return getSpringBehaviorClass()
}

type SpringBehaviorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SpringBehaviorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SpringBehaviorClass) Alloc() SpringBehavior {
	rv := objc.SendIfResponds[SpringBehavior](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SpringBehavior struct {
	objectivec.Object
}

// SpringBehaviorFromID constructs a [SpringBehavior] from an objc.ID.
func SpringBehaviorFromID(id objc.ID) SpringBehavior {
	return SpringBehavior{objectivec.Object{ID: id}}
}

// Ensure SpringBehavior implements ISpringBehavior.
var _ ISpringBehavior = SpringBehavior{}

// An interface definition for the [SpringBehavior] class.
type ISpringBehavior interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SpringBehavior) Init() SpringBehavior {
	rv := objc.SendIfResponds[SpringBehavior](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SpringBehavior) Autorelease() SpringBehavior {
	rv := objc.SendIfResponds[SpringBehavior](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpringBehavior creates a new SpringBehavior instance.
func NewSpringBehavior() SpringBehavior {
	class := getSpringBehaviorClass()
	rv := objc.SendIfResponds[SpringBehavior](objc.ID(class.class), objc.Sel("new"))
	return rv
}
