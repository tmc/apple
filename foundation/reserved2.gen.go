// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Reserved2] class.
var (
	_Reserved2Class     Reserved2Class
	_Reserved2ClassOnce sync.Once
)

func getReserved2Class() Reserved2Class {
	_Reserved2ClassOnce.Do(func() {
		_Reserved2Class = Reserved2Class{class: objc.GetClass("reserved2")}
	})
	return _Reserved2Class
}

// GetReserved2Class returns the class object for reserved2.
func GetReserved2Class() Reserved2Class {
	return getReserved2Class()
}

type Reserved2Class struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (rc Reserved2Class) Alloc() Reserved2 {
	rv := objc.Send[Reserved2](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/reserved2
type Reserved2 struct {
	objectivec.Object
}

// Reserved2FromID constructs a [Reserved2] from an objc.ID.
func Reserved2FromID(id objc.ID) Reserved2 {
	return Reserved2{objectivec.Object{ID: id}}
}
// Ensure Reserved2 implements IReserved2.
var _ IReserved2 = Reserved2{}

// An interface definition for the [Reserved2] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/reserved2
type IReserved2 interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r Reserved2) Init() Reserved2 {
	rv := objc.Send[Reserved2](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r Reserved2) Autorelease() Reserved2 {
	rv := objc.Send[Reserved2](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewReserved2 creates a new Reserved2 instance.
func NewReserved2() Reserved2 {
	class := getReserved2Class()
	rv := objc.Send[Reserved2](objc.ID(class.class), objc.Sel("new"))
	return rv
}

