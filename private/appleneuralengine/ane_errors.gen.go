// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEErrors] class.
var (
	_ANEErrorsClass     ANEErrorsClass
	_ANEErrorsClassOnce sync.Once
)

func getANEErrorsClass() ANEErrorsClass {
	_ANEErrorsClassOnce.Do(func() {
		_ANEErrorsClass = ANEErrorsClass{class: objc.GetClass("_ANEErrors")}
	})
	return _ANEErrorsClass
}

// GetANEErrorsClass returns the class object for _ANEErrors.
func GetANEErrorsClass() ANEErrorsClass {
	return getANEErrorsClass()
}

type ANEErrorsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEErrorsClass) Alloc() ANEErrors {
	rv := objc.Send[ANEErrors](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors
type ANEErrors struct {
	objectivec.Object
}

// ANEErrorsFromID constructs a [ANEErrors] from an objc.ID.
func ANEErrorsFromID(id objc.ID) ANEErrors {
	return ANEErrors{objectivec.Object{id}}
}
// Ensure ANEErrors implements IANEErrors.
var _ IANEErrors = ANEErrors{}





// An interface definition for the [ANEErrors] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors
type IANEErrors interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANEErrors) Init() ANEErrors {
	rv := objc.Send[ANEErrors](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEErrors) Autorelease() ANEErrors {
	rv := objc.Send[ANEErrors](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEErrors creates a new ANEErrors instance.
func NewANEErrors() ANEErrors {
	class := getANEErrorsClass()
	rv := objc.Send[ANEErrors](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/badArgumentForMethod:
func (_ANEErrorsClass ANEErrorsClass) BadArgumentForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("badArgumentForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/baseModelIdentifierNotFoundForNewInstanceMethod:
func (_ANEErrorsClass ANEErrorsClass) BaseModelIdentifierNotFoundForNewInstanceMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("baseModelIdentifierNotFoundForNewInstanceMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/createErrorWithCode:description:
func (_ANEErrorsClass ANEErrorsClass) CreateErrorWithCodeDescription(code int64, description objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("createErrorWithCode:description:"), code, description)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/dataNotFoundForMethod:
func (_ANEErrorsClass ANEErrorsClass) DataNotFoundForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("dataNotFoundForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/entitlementErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) EntitlementErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("entitlementErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/fileAccessErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) FileAccessErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("fileAccessErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/fileNotFoundErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) FileNotFoundErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("fileNotFoundErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/guestToHostInterfaceTooOld:
func (_ANEErrorsClass ANEErrorsClass) GuestToHostInterfaceTooOld(old objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("guestToHostInterfaceTooOld:"), old)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/hostTooOld:
func (_ANEErrorsClass ANEErrorsClass) HostTooOld(old objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("hostTooOld:"), old)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/invalidModelErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) InvalidModelErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("invalidModelErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/invalidModelInstanceErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) InvalidModelInstanceErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("invalidModelInstanceErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/invalidModelKeyErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) InvalidModelKeyErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("invalidModelKeyErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/missingCodeSigningErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) MissingCodeSigningErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("missingCodeSigningErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/modelIdentifierNotFoundForMethod:
func (_ANEErrorsClass ANEErrorsClass) ModelIdentifierNotFoundForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("modelIdentifierNotFoundForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/modelNewInstanceCacheIdentifierNotNilMethod:
func (_ANEErrorsClass ANEErrorsClass) ModelNewInstanceCacheIdentifierNotNilMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("modelNewInstanceCacheIdentifierNotNilMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/notSupportedErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) NotSupportedErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("notSupportedErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/priorityErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) PriorityErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("priorityErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programChainingPrepareErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramChainingPrepareErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programChainingPrepareErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programCreationErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramCreationErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programCreationErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programIOSurfacesMapErrorForMethod:code:
func (_ANEErrorsClass ANEErrorsClass) ProgramIOSurfacesMapErrorForMethodCode(method objectivec.IObject, code int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programIOSurfacesMapErrorForMethod:code:"), method, code)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programIOSurfacesUnmapErrorForMethod:code:
func (_ANEErrorsClass ANEErrorsClass) ProgramIOSurfacesUnmapErrorForMethodCode(method objectivec.IObject, code int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programIOSurfacesUnmapErrorForMethod:code:"), method, code)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programInferenceOtherErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramInferenceOtherErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programInferenceOtherErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programInferenceOverflowErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramInferenceOverflowErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programInferenceOverflowErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programInferenceProgramNotFoundForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramInferenceProgramNotFoundForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programInferenceProgramNotFoundForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programLoadErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramLoadErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programLoadErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programLoadErrorForMethod:code:
func (_ANEErrorsClass ANEErrorsClass) ProgramLoadErrorForMethodCode(method objectivec.IObject, code int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programLoadErrorForMethod:code:"), method, code)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programLoadNewInstanceErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) ProgramLoadNewInstanceErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programLoadNewInstanceErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/programLoadNewInstanceErrorForMethod:code:
func (_ANEErrorsClass ANEErrorsClass) ProgramLoadNewInstanceErrorForMethodCode(method objectivec.IObject, code int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("programLoadNewInstanceErrorForMethod:code:"), method, code)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/systemModelPurgeNotAllowedForMethod:
func (_ANEErrorsClass ANEErrorsClass) SystemModelPurgeNotAllowedForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("systemModelPurgeNotAllowedForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/timeoutErrorForMethod:
func (_ANEErrorsClass ANEErrorsClass) TimeoutErrorForMethod(method objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("timeoutErrorForMethod:"), method)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/virtualizationDataError:
func (_ANEErrorsClass ANEErrorsClass) VirtualizationDataError(error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("virtualizationDataError:"), error_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/virtualizationHostError:
func (_ANEErrorsClass ANEErrorsClass) VirtualizationHostError(error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("virtualizationHostError:"), error_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/virtualizationHostError:error:
func (_ANEErrorsClass ANEErrorsClass) VirtualizationHostErrorError(hostError foundation.INSError, underlyingError foundation.INSError) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("virtualizationHostError:error:"), hostError, underlyingError)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEErrors/virtualizationKernelError:kernelErrorCode:
func (_ANEErrorsClass ANEErrorsClass) VirtualizationKernelErrorKernelErrorCode(error_ objectivec.IObject, code int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEErrorsClass.class), objc.Sel("virtualizationKernelError:kernelErrorCode:"), error_, code)
	return objectivec.Object{ID: rv}
}






















