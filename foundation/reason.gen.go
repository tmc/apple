// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Reason] class.
var (
	_ReasonClass     ReasonClass
	_ReasonClassOnce sync.Once
)

func getReasonClass() ReasonClass {
	_ReasonClassOnce.Do(func() {
		_ReasonClass = ReasonClass{class: objc.GetClass("reason")}
	})
	return _ReasonClass
}

// GetReasonClass returns the class object for reason.
func GetReasonClass() ReasonClass {
	return getReasonClass()
}

type ReasonClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (rc ReasonClass) Alloc() Reason {
	rv := objc.Send[Reason](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Foundation/NSException/reason-c.ivar
type Reason struct {
	objectivec.Object
}

// ReasonFromID constructs a [Reason] from an objc.ID.
func ReasonFromID(id objc.ID) Reason {
	return Reason{objectivec.Object{id}}
}
// Ensure Reason implements IReason.
var _ IReason = Reason{}





// An interface definition for the [Reason] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/reason-c.ivar
type IReason interface {
	objectivec.IObject
}





// Init initializes the instance.
func (r Reason) Init() Reason {
	rv := objc.Send[Reason](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r Reason) Autorelease() Reason {
	rv := objc.Send[Reason](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewReason creates a new Reason instance.
func NewReason() Reason {
	class := getReasonClass()
	rv := objc.Send[Reason](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































