// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionReflection] class.
var (
	_MTLFunctionReflectionClass     MTLFunctionReflectionClass
	_MTLFunctionReflectionClassOnce sync.Once
)

func getMTLFunctionReflectionClass() MTLFunctionReflectionClass {
	_MTLFunctionReflectionClassOnce.Do(func() {
		_MTLFunctionReflectionClass = MTLFunctionReflectionClass{class: objc.GetClass("MTLFunctionReflection")}
	})
	return _MTLFunctionReflectionClass
}

// GetMTLFunctionReflectionClass returns the class object for MTLFunctionReflection.
func GetMTLFunctionReflectionClass() MTLFunctionReflectionClass {
	return getMTLFunctionReflectionClass()
}

type MTLFunctionReflectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionReflectionClass) Alloc() MTLFunctionReflection {
	rv := objc.Send[MTLFunctionReflection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Represents a reflection object containing information about a function in a
// Metal library.
//
// # Instance Properties
//
//   - [MTLFunctionReflection.Bindings]: Provides a list of inputs and outputs of the function.
//   - [MTLFunctionReflection.UserAnnotation]: The string passed to the user annotation attribute for this function. Null if no user annotation is present for this function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionReflection
type MTLFunctionReflection struct {
	objectivec.Object
}

// MTLFunctionReflectionFromID constructs a [MTLFunctionReflection] from an objc.ID.
//
// Represents a reflection object containing information about a function in a
// Metal library.
func MTLFunctionReflectionFromID(id objc.ID) MTLFunctionReflection {
	return MTLFunctionReflection{objectivec.Object{id}}
}
// NOTE: MTLFunctionReflection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLFunctionReflection] class.
//
// # Instance Properties
//
//   - [IMTLFunctionReflection.Bindings]: Provides a list of inputs and outputs of the function.
//   - [IMTLFunctionReflection.UserAnnotation]: The string passed to the user annotation attribute for this function. Null if no user annotation is present for this function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionReflection
type IMTLFunctionReflection interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides a list of inputs and outputs of the function.
	Bindings() []objectivec.IObject
	// The string passed to the user annotation attribute for this function. Null if no user annotation is present for this function.
	UserAnnotation() string
}





// Init initializes the instance.
func (f MTLFunctionReflection) Init() MTLFunctionReflection {
	rv := objc.Send[MTLFunctionReflection](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionReflection) Autorelease() MTLFunctionReflection {
	rv := objc.Send[MTLFunctionReflection](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionReflection creates a new MTLFunctionReflection instance.
func NewMTLFunctionReflection() MTLFunctionReflection {
	class := getMTLFunctionReflectionClass()
	rv := objc.Send[MTLFunctionReflection](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Provides a list of inputs and outputs of the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/bindings
func (f MTLFunctionReflection) Bindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("bindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// The string passed to the user annotation attribute for this function. Null
// if no user annotation is present for this function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/userAnnotation
func (f MTLFunctionReflection) UserAnnotation() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("userAnnotation"))
	return foundation.NSStringFromID(rv).String()
}

























