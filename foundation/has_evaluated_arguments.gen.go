// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasEvaluatedArguments] class.
var (
	_HasEvaluatedArgumentsClass     HasEvaluatedArgumentsClass
	_HasEvaluatedArgumentsClassOnce sync.Once
)

func getHasEvaluatedArgumentsClass() HasEvaluatedArgumentsClass {
	_HasEvaluatedArgumentsClassOnce.Do(func() {
		_HasEvaluatedArgumentsClass = HasEvaluatedArgumentsClass{class: objc.GetClass("hasEvaluatedArguments")}
	})
	return _HasEvaluatedArgumentsClass
}

// GetHasEvaluatedArgumentsClass returns the class object for hasEvaluatedArguments.
func GetHasEvaluatedArgumentsClass() HasEvaluatedArgumentsClass {
	return getHasEvaluatedArgumentsClass()
}

type HasEvaluatedArgumentsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasEvaluatedArgumentsClass) Alloc() HasEvaluatedArguments {
	rv := objc.Send[HasEvaluatedArguments](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/hasEvaluatedArguments
type HasEvaluatedArguments struct {
	objectivec.Object
}

// HasEvaluatedArgumentsFromID constructs a [HasEvaluatedArguments] from an objc.ID.
func HasEvaluatedArgumentsFromID(id objc.ID) HasEvaluatedArguments {
	return HasEvaluatedArguments{objectivec.Object{ID: id}}
}
// Ensure HasEvaluatedArguments implements IHasEvaluatedArguments.
var _ IHasEvaluatedArguments = HasEvaluatedArguments{}

// An interface definition for the [HasEvaluatedArguments] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/struct_(unnamed)/hasEvaluatedArguments
type IHasEvaluatedArguments interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasEvaluatedArguments) Init() HasEvaluatedArguments {
	rv := objc.Send[HasEvaluatedArguments](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasEvaluatedArguments) Autorelease() HasEvaluatedArguments {
	rv := objc.Send[HasEvaluatedArguments](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasEvaluatedArguments creates a new HasEvaluatedArguments instance.
func NewHasEvaluatedArguments() HasEvaluatedArguments {
	class := getHasEvaluatedArgumentsClass()
	rv := objc.Send[HasEvaluatedArguments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

