// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ProgramLibraryE5BundleImpl] class.
var (
	_MLE5ProgramLibraryE5BundleImplClass     MLE5ProgramLibraryE5BundleImplClass
	_MLE5ProgramLibraryE5BundleImplClassOnce sync.Once
)

func getMLE5ProgramLibraryE5BundleImplClass() MLE5ProgramLibraryE5BundleImplClass {
	_MLE5ProgramLibraryE5BundleImplClassOnce.Do(func() {
		_MLE5ProgramLibraryE5BundleImplClass = MLE5ProgramLibraryE5BundleImplClass{class: objc.GetClass("MLE5ProgramLibraryE5BundleImpl")}
	})
	return _MLE5ProgramLibraryE5BundleImplClass
}

// GetMLE5ProgramLibraryE5BundleImplClass returns the class object for MLE5ProgramLibraryE5BundleImpl.
func GetMLE5ProgramLibraryE5BundleImplClass() MLE5ProgramLibraryE5BundleImplClass {
	return getMLE5ProgramLibraryE5BundleImplClass()
}

type MLE5ProgramLibraryE5BundleImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ProgramLibraryE5BundleImplClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ProgramLibraryE5BundleImplClass) Alloc() MLE5ProgramLibraryE5BundleImpl {
	rv := objc.Send[MLE5ProgramLibraryE5BundleImpl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5ProgramLibraryE5BundleImpl.Configuration]
//   - [MLE5ProgramLibraryE5BundleImpl.CreateProgramLibraryHandleWithRespecializationError]
//   - [MLE5ProgramLibraryE5BundleImpl.E5BundleURL]
//   - [MLE5ProgramLibraryE5BundleImpl.ModelDisplayName]
//   - [MLE5ProgramLibraryE5BundleImpl.SerializedMILText]
//   - [MLE5ProgramLibraryE5BundleImpl.InitWithE5BundleAtURLConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl
type MLE5ProgramLibraryE5BundleImpl struct {
	objectivec.Object
}

// MLE5ProgramLibraryE5BundleImplFromID constructs a [MLE5ProgramLibraryE5BundleImpl] from an objc.ID.
func MLE5ProgramLibraryE5BundleImplFromID(id objc.ID) MLE5ProgramLibraryE5BundleImpl {
	return MLE5ProgramLibraryE5BundleImpl{objectivec.Object{ID: id}}
}

// Ensure MLE5ProgramLibraryE5BundleImpl implements IMLE5ProgramLibraryE5BundleImpl.
var _ IMLE5ProgramLibraryE5BundleImpl = MLE5ProgramLibraryE5BundleImpl{}

// An interface definition for the [MLE5ProgramLibraryE5BundleImpl] class.
//
// # Methods
//
//   - [IMLE5ProgramLibraryE5BundleImpl.Configuration]
//   - [IMLE5ProgramLibraryE5BundleImpl.CreateProgramLibraryHandleWithRespecializationError]
//   - [IMLE5ProgramLibraryE5BundleImpl.E5BundleURL]
//   - [IMLE5ProgramLibraryE5BundleImpl.ModelDisplayName]
//   - [IMLE5ProgramLibraryE5BundleImpl.SerializedMILText]
//   - [IMLE5ProgramLibraryE5BundleImpl.InitWithE5BundleAtURLConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl
type IMLE5ProgramLibraryE5BundleImpl interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error)
	E5BundleURL() foundation.INSURL
	ModelDisplayName() string
	SerializedMILText() string
	InitWithE5BundleAtURLConfiguration(url foundation.INSURL, configuration objectivec.IObject) MLE5ProgramLibraryE5BundleImpl
}

// Init initializes the instance.
func (m MLE5ProgramLibraryE5BundleImpl) Init() MLE5ProgramLibraryE5BundleImpl {
	rv := objc.Send[MLE5ProgramLibraryE5BundleImpl](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5ProgramLibraryE5BundleImpl) Autorelease() MLE5ProgramLibraryE5BundleImpl {
	rv := objc.Send[MLE5ProgramLibraryE5BundleImpl](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ProgramLibraryE5BundleImpl creates a new MLE5ProgramLibraryE5BundleImpl instance.
func NewMLE5ProgramLibraryE5BundleImpl() MLE5ProgramLibraryE5BundleImpl {
	class := getMLE5ProgramLibraryE5BundleImplClass()
	rv := objc.Send[MLE5ProgramLibraryE5BundleImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/initWithE5BundleAtURL:configuration:
func NewE5ProgramLibraryE5BundleImplWithE5BundleAtURLConfiguration(url foundation.INSURL, configuration objectivec.IObject) MLE5ProgramLibraryE5BundleImpl {
	instance := getMLE5ProgramLibraryE5BundleImplClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithE5BundleAtURL:configuration:"), url, configuration)
	return MLE5ProgramLibraryE5BundleImplFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/createProgramLibraryHandleWithRespecialization:error:
func (m MLE5ProgramLibraryE5BundleImpl) CreateProgramLibraryHandleWithRespecializationError(respecialization bool) (E5rtProgramLibraryRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("createProgramLibraryHandleWithRespecialization:error:"), respecialization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(E5rtProgramLibraryRef), foundation.NSErrorFrom(errorPtr)
	}
	return E5rtProgramLibraryRef(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/initWithE5BundleAtURL:configuration:
func (m MLE5ProgramLibraryE5BundleImpl) InitWithE5BundleAtURLConfiguration(url foundation.INSURL, configuration objectivec.IObject) MLE5ProgramLibraryE5BundleImpl {
	rv := objc.Send[MLE5ProgramLibraryE5BundleImpl](m.ID, objc.Sel("initWithE5BundleAtURL:configuration:"), url, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/configuration
func (m MLE5ProgramLibraryE5BundleImpl) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/e5BundleURL
func (m MLE5ProgramLibraryE5BundleImpl) E5BundleURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("e5BundleURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/modelDisplayName
func (m MLE5ProgramLibraryE5BundleImpl) ModelDisplayName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDisplayName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibraryE5BundleImpl/serializedMILText
func (m MLE5ProgramLibraryE5BundleImpl) SerializedMILText() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}
