// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ProgramLibrary] class.
var (
	_MLE5ProgramLibraryClass     MLE5ProgramLibraryClass
	_MLE5ProgramLibraryClassOnce sync.Once
)

func getMLE5ProgramLibraryClass() MLE5ProgramLibraryClass {
	_MLE5ProgramLibraryClassOnce.Do(func() {
		_MLE5ProgramLibraryClass = MLE5ProgramLibraryClass{class: objc.GetClass("MLE5ProgramLibrary")}
	})
	return _MLE5ProgramLibraryClass
}

// GetMLE5ProgramLibraryClass returns the class object for MLE5ProgramLibrary.
func GetMLE5ProgramLibraryClass() MLE5ProgramLibraryClass {
	return getMLE5ProgramLibraryClass()
}

type MLE5ProgramLibraryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ProgramLibraryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ProgramLibraryClass) Alloc() MLE5ProgramLibrary {
	rv := objc.Send[MLE5ProgramLibrary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5ProgramLibrary._allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError]
//   - [MLE5ProgramLibrary._programLibraryHandleWithForceRespecializationError]
//   - [MLE5ProgramLibrary.Container]
//   - [MLE5ProgramLibrary.CreateOperationForFunctionNameForceRespecializationHasRangeShapeInputsError]
//   - [MLE5ProgramLibrary.FunctionNames]
//   - [MLE5ProgramLibrary.Impl]
//   - [MLE5ProgramLibrary.LazyInitQueue]
//   - [MLE5ProgramLibrary.ModelConfiguration]
//   - [MLE5ProgramLibrary.ModelDisplayName]
//   - [MLE5ProgramLibrary.NewStateForFunctionNamedStateNamesClientBuffersError]
//   - [MLE5ProgramLibrary.PrepareAndReturnError]
//   - [MLE5ProgramLibrary.SegmentationAnalyticsAndReturnError]
//   - [MLE5ProgramLibrary.SerializedMILText]
//   - [MLE5ProgramLibrary.InitWithContainerConfigurationError]
//   - [MLE5ProgramLibrary.InitWithImplContainerConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary
type MLE5ProgramLibrary struct {
	objectivec.Object
}

// MLE5ProgramLibraryFromID constructs a [MLE5ProgramLibrary] from an objc.ID.
func MLE5ProgramLibraryFromID(id objc.ID) MLE5ProgramLibrary {
	return MLE5ProgramLibrary{objectivec.Object{ID: id}}
}

// Ensure MLE5ProgramLibrary implements IMLE5ProgramLibrary.
var _ IMLE5ProgramLibrary = MLE5ProgramLibrary{}

// An interface definition for the [MLE5ProgramLibrary] class.
//
// # Methods
//
//   - [IMLE5ProgramLibrary._allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError]
//   - [IMLE5ProgramLibrary._programLibraryHandleWithForceRespecializationError]
//   - [IMLE5ProgramLibrary.Container]
//   - [IMLE5ProgramLibrary.CreateOperationForFunctionNameForceRespecializationHasRangeShapeInputsError]
//   - [IMLE5ProgramLibrary.FunctionNames]
//   - [IMLE5ProgramLibrary.Impl]
//   - [IMLE5ProgramLibrary.LazyInitQueue]
//   - [IMLE5ProgramLibrary.ModelConfiguration]
//   - [IMLE5ProgramLibrary.ModelDisplayName]
//   - [IMLE5ProgramLibrary.NewStateForFunctionNamedStateNamesClientBuffersError]
//   - [IMLE5ProgramLibrary.PrepareAndReturnError]
//   - [IMLE5ProgramLibrary.SegmentationAnalyticsAndReturnError]
//   - [IMLE5ProgramLibrary.SerializedMILText]
//   - [IMLE5ProgramLibrary.InitWithContainerConfigurationError]
//   - [IMLE5ProgramLibrary.InitWithImplContainerConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary
type IMLE5ProgramLibrary interface {
	objectivec.IObject

	// Topic: Methods

	_allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError(named objectivec.IObject, name objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error)
	_programLibraryHandleWithForceRespecializationError(respecialization bool) (objectivec.IObject, error)
	Container() IMLProgramE5Container
	CreateOperationForFunctionNameForceRespecializationHasRangeShapeInputsError(name objectivec.IObject, respecialization bool, inputs bool) (objectivec.IObject, error)
	FunctionNames() foundation.INSArray
	Impl() objectivec.IObject
	LazyInitQueue() objectivec.Object
	ModelConfiguration() IMLModelConfiguration
	ModelDisplayName() string
	NewStateForFunctionNamedStateNamesClientBuffersError(named objectivec.IObject, names objectivec.IObject, buffers objectivec.IObject) (objectivec.IObject, error)
	PrepareAndReturnError() (bool, error)
	SegmentationAnalyticsAndReturnError() (objectivec.IObject, error)
	SerializedMILText() string
	InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5ProgramLibrary, error)
	InitWithImplContainerConfiguration(impl objectivec.IObject, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibrary
}

// Init initializes the instance.
func (e MLE5ProgramLibrary) Init() MLE5ProgramLibrary {
	rv := objc.Send[MLE5ProgramLibrary](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5ProgramLibrary) Autorelease() MLE5ProgramLibrary {
	rv := objc.Send[MLE5ProgramLibrary](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ProgramLibrary creates a new MLE5ProgramLibrary instance.
func NewMLE5ProgramLibrary() MLE5ProgramLibrary {
	class := getMLE5ProgramLibraryClass()
	rv := objc.Send[MLE5ProgramLibrary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/initWithContainer:configuration:error:
func NewE5ProgramLibraryWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5ProgramLibrary, error) {
	var errorPtr objc.ID
	instance := getMLE5ProgramLibraryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLE5ProgramLibrary{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLE5ProgramLibraryFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/initWithImpl:container:configuration:
func NewE5ProgramLibraryWithImplContainerConfiguration(impl objectivec.IObject, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibrary {
	instance := getMLE5ProgramLibraryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:container:configuration:"), impl, container, configuration)
	return MLE5ProgramLibraryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/_allocateStateBufferForFeatureNamed:entryFunctionName:programFunctionNames:error:
func (e MLE5ProgramLibrary) _allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError(named objectivec.IObject, name objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_allocateStateBufferForFeatureNamed:entryFunctionName:programFunctionNames:error:"), named, name, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError is an exported wrapper for the private method _allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError.
func (e MLE5ProgramLibrary) AllocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError(named objectivec.IObject, name objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	return e._allocateStateBufferForFeatureNamedEntryFunctionNameProgramFunctionNamesError(named, name, names)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/_programLibraryHandleWithForceRespecialization:error:
func (e MLE5ProgramLibrary) _programLibraryHandleWithForceRespecializationError(respecialization bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_programLibraryHandleWithForceRespecialization:error:"), respecialization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ProgramLibraryHandleWithForceRespecializationError is an exported wrapper for the private method _programLibraryHandleWithForceRespecializationError.
func (e MLE5ProgramLibrary) ProgramLibraryHandleWithForceRespecializationError(respecialization bool) (objectivec.IObject, error) {
	return e._programLibraryHandleWithForceRespecializationError(respecialization)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/createOperationForFunctionName:forceRespecialization:hasRangeShapeInputs:error:
func (e MLE5ProgramLibrary) CreateOperationForFunctionNameForceRespecializationHasRangeShapeInputsError(name objectivec.IObject, respecialization bool, inputs bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("createOperationForFunctionName:forceRespecialization:hasRangeShapeInputs:error:"), name, respecialization, inputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/newStateForFunctionNamed:stateNames:clientBuffers:error:
func (e MLE5ProgramLibrary) NewStateForFunctionNamedStateNamesClientBuffersError(named objectivec.IObject, names objectivec.IObject, buffers objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("newStateForFunctionNamed:stateNames:clientBuffers:error:"), named, names, buffers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/prepareAndReturnError:
func (e MLE5ProgramLibrary) PrepareAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/segmentationAnalyticsAndReturnError:
func (e MLE5ProgramLibrary) SegmentationAnalyticsAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("segmentationAnalyticsAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/initWithContainer:configuration:error:
func (e MLE5ProgramLibrary) InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5ProgramLibrary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLE5ProgramLibrary{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLE5ProgramLibraryFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/initWithImpl:container:configuration:
func (e MLE5ProgramLibrary) InitWithImplContainerConfiguration(impl objectivec.IObject, container objectivec.IObject, configuration objectivec.IObject) MLE5ProgramLibrary {
	rv := objc.Send[MLE5ProgramLibrary](e.ID, objc.Sel("initWithImpl:container:configuration:"), impl, container, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/container
func (e MLE5ProgramLibrary) Container() IMLProgramE5Container {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("container"))
	return MLProgramE5ContainerFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/functionNames
func (e MLE5ProgramLibrary) FunctionNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("functionNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/impl
func (e MLE5ProgramLibrary) Impl() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("impl"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/lazyInitQueue
func (e MLE5ProgramLibrary) LazyInitQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("lazyInitQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/modelConfiguration
func (e MLE5ProgramLibrary) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/modelDisplayName
func (e MLE5ProgramLibrary) ModelDisplayName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDisplayName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ProgramLibrary/serializedMILText
func (e MLE5ProgramLibrary) SerializedMILText() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}
