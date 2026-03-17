// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIBaseParams] class.
var (
	_DIBaseParamsClass     DIBaseParamsClass
	_DIBaseParamsClassOnce sync.Once
)

func getDIBaseParamsClass() DIBaseParamsClass {
	_DIBaseParamsClassOnce.Do(func() {
		_DIBaseParamsClass = DIBaseParamsClass{class: objc.GetClass("DIBaseParams")}
	})
	return _DIBaseParamsClass
}

// GetDIBaseParamsClass returns the class object for DIBaseParams.
func GetDIBaseParamsClass() DIBaseParamsClass {
	return getDIBaseParamsClass()
}

type DIBaseParamsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseParamsClass) Alloc() DIBaseParams {
	rv := objc.Send[DIBaseParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other diskimages2 classes. [Full Topic]
type DIBaseParams struct {
	objectivec.Object
}

// DIBaseParamsFromID constructs a [DIBaseParams] from an objc.ID.
//
// A parent class referenced by other diskimages2 classes.
func DIBaseParamsFromID(id objc.ID) DIBaseParams {
	return DIBaseParams{objectivec.Object{ID: id}}
}
// Ensure DIBaseParams implements IDIBaseParams.
var _ IDIBaseParams = DIBaseParams{}

// An interface definition for the [DIBaseParams] class.
type IDIBaseParams interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DIBaseParams) Init() DIBaseParams {
	rv := objc.Send[DIBaseParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseParams) Autorelease() DIBaseParams {
	rv := objc.Send[DIBaseParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseParams creates a new DIBaseParams instance.
func NewDIBaseParams() DIBaseParams {
	class := getDIBaseParamsClass()
	rv := objc.Send[DIBaseParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

