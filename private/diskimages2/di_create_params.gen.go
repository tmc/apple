// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DICreateParams] class.
var (
	_DICreateParamsClass     DICreateParamsClass
	_DICreateParamsClassOnce sync.Once
)

func getDICreateParamsClass() DICreateParamsClass {
	_DICreateParamsClassOnce.Do(func() {
		_DICreateParamsClass = DICreateParamsClass{class: objc.GetClass("DICreateParams")}
	})
	return _DICreateParamsClass
}

// GetDICreateParamsClass returns the class object for DICreateParams.
func GetDICreateParamsClass() DICreateParamsClass {
	return getDICreateParamsClass()
}

type DICreateParamsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICreateParamsClass) Alloc() DICreateParams {
	rv := objc.Send[DICreateParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other diskimages2 classes. [Full Topic]
type DICreateParams struct {
	objectivec.Object
}

// DICreateParamsFromID constructs a [DICreateParams] from an objc.ID.
//
// A parent class referenced by other diskimages2 classes.
func DICreateParamsFromID(id objc.ID) DICreateParams {
	return DICreateParams{objectivec.Object{id}}
}
// Ensure DICreateParams implements IDICreateParams.
var _ IDICreateParams = DICreateParams{}





// An interface definition for the [DICreateParams] class.
type IDICreateParams interface {
	objectivec.IObject
}





// Init initializes the instance.
func (d DICreateParams) Init() DICreateParams {
	rv := objc.Send[DICreateParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICreateParams) Autorelease() DICreateParams {
	rv := objc.Send[DICreateParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICreateParams creates a new DICreateParams instance.
func NewDICreateParams() DICreateParams {
	class := getDICreateParamsClass()
	rv := objc.Send[DICreateParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































