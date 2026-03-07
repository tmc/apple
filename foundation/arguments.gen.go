// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Arguments] class.
var (
	_ArgumentsClass     ArgumentsClass
	_ArgumentsClassOnce sync.Once
)

func getArgumentsClass() ArgumentsClass {
	_ArgumentsClassOnce.Do(func() {
		_ArgumentsClass = ArgumentsClass{class: objc.GetClass("arguments")}
	})
	return _ArgumentsClass
}

// GetArgumentsClass returns the class object for arguments.
func GetArgumentsClass() ArgumentsClass {
	return getArgumentsClass()
}

type ArgumentsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ArgumentsClass) Alloc() Arguments {
	rv := objc.Send[Arguments](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/arguments-c.ivar
type Arguments struct {
	objectivec.Object
}

// ArgumentsFromID constructs a [Arguments] from an objc.ID.
func ArgumentsFromID(id objc.ID) Arguments {
	return Arguments{objectivec.Object{id}}
}
// Ensure Arguments implements IArguments.
var _ IArguments = Arguments{}





// An interface definition for the [Arguments] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/arguments-c.ivar
type IArguments interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a Arguments) Init() Arguments {
	rv := objc.Send[Arguments](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a Arguments) Autorelease() Arguments {
	rv := objc.Send[Arguments](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewArguments creates a new Arguments instance.
func NewArguments() Arguments {
	class := getArgumentsClass()
	rv := objc.Send[Arguments](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































