// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRemoteCoreMLErrors] class.
var (
	_MLRemoteCoreMLErrorsClass     MLRemoteCoreMLErrorsClass
	_MLRemoteCoreMLErrorsClassOnce sync.Once
)

func getMLRemoteCoreMLErrorsClass() MLRemoteCoreMLErrorsClass {
	_MLRemoteCoreMLErrorsClassOnce.Do(func() {
		_MLRemoteCoreMLErrorsClass = MLRemoteCoreMLErrorsClass{class: objc.GetClass("_MLRemoteCoreMLErrors")}
	})
	return _MLRemoteCoreMLErrorsClass
}

// GetMLRemoteCoreMLErrorsClass returns the class object for _MLRemoteCoreMLErrors.
func GetMLRemoteCoreMLErrorsClass() MLRemoteCoreMLErrorsClass {
	return getMLRemoteCoreMLErrorsClass()
}

type MLRemoteCoreMLErrorsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLRemoteCoreMLErrorsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRemoteCoreMLErrorsClass) Alloc() MLRemoteCoreMLErrors {
	rv := objc.SendIfResponds[MLRemoteCoreMLErrors](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLRemoteCoreMLErrors struct {
	objectivec.Object
}

// MLRemoteCoreMLErrorsFromID constructs a [MLRemoteCoreMLErrors] from an objc.ID.
func MLRemoteCoreMLErrorsFromID(id objc.ID) MLRemoteCoreMLErrors {
	return MLRemoteCoreMLErrors{objectivec.Object{ID: id}}
}

// Ensure MLRemoteCoreMLErrors implements IMLRemoteCoreMLErrors.
var _ IMLRemoteCoreMLErrors = MLRemoteCoreMLErrors{}

// An interface definition for the [MLRemoteCoreMLErrors] class.
type IMLRemoteCoreMLErrors interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLRemoteCoreMLErrors) Init() MLRemoteCoreMLErrors {
	rv := objc.SendIfResponds[MLRemoteCoreMLErrors](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLRemoteCoreMLErrors) Autorelease() MLRemoteCoreMLErrors {
	rv := objc.SendIfResponds[MLRemoteCoreMLErrors](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRemoteCoreMLErrors creates a new MLRemoteCoreMLErrors instance.
func NewMLRemoteCoreMLErrors() MLRemoteCoreMLErrors {
	class := getMLRemoteCoreMLErrorsClass()
	rv := objc.SendIfResponds[MLRemoteCoreMLErrors](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLRemoteCoreMLErrorsClass MLRemoteCoreMLErrorsClass) ClientTimeoutErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLRemoteCoreMLErrorsClass.class), objc.Sel("clientTimeoutErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}
func (_MLRemoteCoreMLErrorsClass MLRemoteCoreMLErrorsClass) CreateErrorWithCodeDescription(code int64, description objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLRemoteCoreMLErrorsClass.class), objc.Sel("createErrorWithCode:description:"), code, description)
	return objectivec.Object{ID: rv}
}
