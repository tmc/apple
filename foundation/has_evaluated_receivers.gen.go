// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasEvaluatedReceivers] class.
var (
	_HasEvaluatedReceiversClass     HasEvaluatedReceiversClass
	_HasEvaluatedReceiversClassOnce sync.Once
)

func getHasEvaluatedReceiversClass() HasEvaluatedReceiversClass {
	_HasEvaluatedReceiversClassOnce.Do(func() {
		_HasEvaluatedReceiversClass = HasEvaluatedReceiversClass{class: objc.GetClass("hasEvaluatedReceivers")}
	})
	return _HasEvaluatedReceiversClass
}

// GetHasEvaluatedReceiversClass returns the class object for hasEvaluatedReceivers.
func GetHasEvaluatedReceiversClass() HasEvaluatedReceiversClass {
	return getHasEvaluatedReceiversClass()
}

type HasEvaluatedReceiversClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasEvaluatedReceiversClass) Alloc() HasEvaluatedReceivers {
	rv := objc.Send[HasEvaluatedReceivers](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/hasEvaluatedReceivers
type HasEvaluatedReceivers struct {
	objectivec.Object
}

// HasEvaluatedReceiversFromID constructs a [HasEvaluatedReceivers] from an objc.ID.
func HasEvaluatedReceiversFromID(id objc.ID) HasEvaluatedReceivers {
	return HasEvaluatedReceivers{objectivec.Object{ID: id}}
}
// Ensure HasEvaluatedReceivers implements IHasEvaluatedReceivers.
var _ IHasEvaluatedReceivers = HasEvaluatedReceivers{}

// An interface definition for the [HasEvaluatedReceivers] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/hasEvaluatedReceivers
type IHasEvaluatedReceivers interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasEvaluatedReceivers) Init() HasEvaluatedReceivers {
	rv := objc.Send[HasEvaluatedReceivers](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasEvaluatedReceivers) Autorelease() HasEvaluatedReceivers {
	rv := objc.Send[HasEvaluatedReceivers](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasEvaluatedReceivers creates a new HasEvaluatedReceivers instance.
func NewHasEvaluatedReceivers() HasEvaluatedReceivers {
	class := getHasEvaluatedReceiversClass()
	rv := objc.Send[HasEvaluatedReceivers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

