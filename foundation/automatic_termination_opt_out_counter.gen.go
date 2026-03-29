// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AutomaticTerminationOptOutCounter] class.
var (
	_AutomaticTerminationOptOutCounterClass     AutomaticTerminationOptOutCounterClass
	_AutomaticTerminationOptOutCounterClassOnce sync.Once
)

func getAutomaticTerminationOptOutCounterClass() AutomaticTerminationOptOutCounterClass {
	_AutomaticTerminationOptOutCounterClassOnce.Do(func() {
		_AutomaticTerminationOptOutCounterClass = AutomaticTerminationOptOutCounterClass{class: objc.GetClass("automaticTerminationOptOutCounter")}
	})
	return _AutomaticTerminationOptOutCounterClass
}

// GetAutomaticTerminationOptOutCounterClass returns the class object for automaticTerminationOptOutCounter.
func GetAutomaticTerminationOptOutCounterClass() AutomaticTerminationOptOutCounterClass {
	return getAutomaticTerminationOptOutCounterClass()
}

type AutomaticTerminationOptOutCounterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AutomaticTerminationOptOutCounterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AutomaticTerminationOptOutCounterClass) Alloc() AutomaticTerminationOptOutCounter {
	rv := objc.Send[AutomaticTerminationOptOutCounter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/automaticTerminationOptOutCounter
type AutomaticTerminationOptOutCounter struct {
	objectivec.Object
}

// AutomaticTerminationOptOutCounterFromID constructs a [AutomaticTerminationOptOutCounter] from an objc.ID.
func AutomaticTerminationOptOutCounterFromID(id objc.ID) AutomaticTerminationOptOutCounter {
	return AutomaticTerminationOptOutCounter{objectivec.Object{ID: id}}
}
// Ensure AutomaticTerminationOptOutCounter implements IAutomaticTerminationOptOutCounter.
var _ IAutomaticTerminationOptOutCounter = AutomaticTerminationOptOutCounter{}

// An interface definition for the [AutomaticTerminationOptOutCounter] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/automaticTerminationOptOutCounter
type IAutomaticTerminationOptOutCounter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AutomaticTerminationOptOutCounter) Init() AutomaticTerminationOptOutCounter {
	rv := objc.Send[AutomaticTerminationOptOutCounter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AutomaticTerminationOptOutCounter) Autorelease() AutomaticTerminationOptOutCounter {
	rv := objc.Send[AutomaticTerminationOptOutCounter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutomaticTerminationOptOutCounter creates a new AutomaticTerminationOptOutCounter instance.
func NewAutomaticTerminationOptOutCounter() AutomaticTerminationOptOutCounter {
	class := getAutomaticTerminationOptOutCounterClass()
	rv := objc.Send[AutomaticTerminationOptOutCounter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

