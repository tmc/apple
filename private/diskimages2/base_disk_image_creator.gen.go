// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BaseDiskImageCreator] class.
var (
	_BaseDiskImageCreatorClass     BaseDiskImageCreatorClass
	_BaseDiskImageCreatorClassOnce sync.Once
)

func getBaseDiskImageCreatorClass() BaseDiskImageCreatorClass {
	_BaseDiskImageCreatorClassOnce.Do(func() {
		_BaseDiskImageCreatorClass = BaseDiskImageCreatorClass{class: objc.GetClass("BaseDiskImageCreator")}
	})
	return _BaseDiskImageCreatorClass
}

// GetBaseDiskImageCreatorClass returns the class object for BaseDiskImageCreator.
func GetBaseDiskImageCreatorClass() BaseDiskImageCreatorClass {
	return getBaseDiskImageCreatorClass()
}

type BaseDiskImageCreatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BaseDiskImageCreatorClass) Alloc() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other diskimages2 classes. [Full Topic]
type BaseDiskImageCreator struct {
	objectivec.Object
}

// BaseDiskImageCreatorFromID constructs a [BaseDiskImageCreator] from an objc.ID.
//
// A parent class referenced by other diskimages2 classes.
func BaseDiskImageCreatorFromID(id objc.ID) BaseDiskImageCreator {
	return BaseDiskImageCreator{objectivec.Object{id}}
}
// Ensure BaseDiskImageCreator implements IBaseDiskImageCreator.
var _ IBaseDiskImageCreator = BaseDiskImageCreator{}

// An interface definition for the [BaseDiskImageCreator] class.
type IBaseDiskImageCreator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BaseDiskImageCreator) Init() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BaseDiskImageCreator) Autorelease() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBaseDiskImageCreator creates a new BaseDiskImageCreator instance.
func NewBaseDiskImageCreator() BaseDiskImageCreator {
	class := getBaseDiskImageCreatorClass()
	rv := objc.Send[BaseDiskImageCreator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

