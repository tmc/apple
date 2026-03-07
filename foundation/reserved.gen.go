// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RESERVED] class.
var (
	_RESERVEDClass     RESERVEDClass
	_RESERVEDClassOnce sync.Once
)

func getRESERVEDClass() RESERVEDClass {
	_RESERVEDClassOnce.Do(func() {
		_RESERVEDClass = RESERVEDClass{class: objc.GetClass("RESERVED")}
	})
	return _RESERVEDClass
}

// GetRESERVEDClass returns the class object for RESERVED.
func GetRESERVEDClass() RESERVEDClass {
	return getRESERVEDClass()
}

type RESERVEDClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (rc RESERVEDClass) Alloc() RESERVED {
	rv := objc.Send[RESERVED](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/RESERVED
type RESERVED struct {
	objectivec.Object
}

// RESERVEDFromID constructs a [RESERVED] from an objc.ID.
func RESERVEDFromID(id objc.ID) RESERVED {
	return RESERVED{objectivec.Object{id}}
}
// Ensure RESERVED implements IRESERVED.
var _ IRESERVED = RESERVED{}





// An interface definition for the [RESERVED] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/RESERVED
type IRESERVED interface {
	objectivec.IObject
}





// Init initializes the instance.
func (r RESERVED) Init() RESERVED {
	rv := objc.Send[RESERVED](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RESERVED) Autorelease() RESERVED {
	rv := objc.Send[RESERVED](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRESERVED creates a new RESERVED instance.
func NewRESERVED() RESERVED {
	class := getRESERVEDClass()
	rv := objc.Send[RESERVED](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































