// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelErrorUtils] class.
var (
	_MLModelErrorUtilsClass     MLModelErrorUtilsClass
	_MLModelErrorUtilsClassOnce sync.Once
)

func getMLModelErrorUtilsClass() MLModelErrorUtilsClass {
	_MLModelErrorUtilsClassOnce.Do(func() {
		_MLModelErrorUtilsClass = MLModelErrorUtilsClass{class: objc.GetClass("MLModelErrorUtils")}
	})
	return _MLModelErrorUtilsClass
}

// GetMLModelErrorUtilsClass returns the class object for MLModelErrorUtils.
func GetMLModelErrorUtilsClass() MLModelErrorUtilsClass {
	return getMLModelErrorUtilsClass()
}

type MLModelErrorUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelErrorUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelErrorUtilsClass) Alloc() MLModelErrorUtils {
	rv := objc.SendIfResponds[MLModelErrorUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLModelErrorUtils struct {
	objectivec.Object
}

// MLModelErrorUtilsFromID constructs a [MLModelErrorUtils] from an objc.ID.
func MLModelErrorUtilsFromID(id objc.ID) MLModelErrorUtils {
	return MLModelErrorUtils{objectivec.Object{ID: id}}
}

// Ensure MLModelErrorUtils implements IMLModelErrorUtils.
var _ IMLModelErrorUtils = MLModelErrorUtils{}

// An interface definition for the [MLModelErrorUtils] class.
type IMLModelErrorUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelErrorUtils) Init() MLModelErrorUtils {
	rv := objc.SendIfResponds[MLModelErrorUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelErrorUtils) Autorelease() MLModelErrorUtils {
	rv := objc.SendIfResponds[MLModelErrorUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelErrorUtils creates a new MLModelErrorUtils instance.
func NewMLModelErrorUtils() MLModelErrorUtils {
	class := getMLModelErrorUtilsClass()
	rv := objc.SendIfResponds[MLModelErrorUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) IOErrorWithFormat(format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("IOErrorWithFormat:"), format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) CustomLayerErrorWithUnderlyingErrorWithFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("customLayerErrorWithUnderlyingError:withFormat:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ErrorWithCodeFormat(code int64, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("errorWithCode:format:"), code, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ErrorWithCodeFormatArgs(code int64, format objectivec.IObject, args string) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("errorWithCode:format:args:"), code, format, unsafe.Pointer(unsafe.StringData(args+"\x00")))
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ErrorWithCodeUnderlyingErrorFormat(code int64, error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("errorWithCode:underlyingError:format:"), code, error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ErrorWithCodeUnderlyingErrorFormatArgs(code int64, error_ objectivec.IObject, format objectivec.IObject, args string) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("errorWithCode:underlyingError:format:args:"), code, error_, format, unsafe.Pointer(unsafe.StringData(args+"\x00")))
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ErrorWithIntegerCodeUnderlyingErrorFormatArgs(code int64, error_ objectivec.IObject, format objectivec.IObject, args string) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("errorWithIntegerCode:underlyingError:format:args:"), code, error_, format, unsafe.Pointer(unsafe.StringData(args+"\x00")))
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) FeatureTypeErrorWithFormat(format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("featureTypeErrorWithFormat:"), format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) GenericErrorWithFormat(format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("genericErrorWithFormat:"), format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ModelDecryptionErrorWithUnderlyingErrorFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("modelDecryptionErrorWithUnderlyingError:format:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ModelDecryptionKeyFetchErrorWithUnderlyingErrorFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("modelDecryptionKeyFetchErrorWithUnderlyingError:format:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ModelEncryptionErrorWithUnderlyingErrorFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("modelEncryptionErrorWithUnderlyingError:format:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ParameterErrorWithUnderlyingErrorFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("parameterErrorWithUnderlyingError:format:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) PrivateErrorWithCodeUnderlyingErrorFormatArgs(code int64, error_ objectivec.IObject, format objectivec.IObject, args string) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("privateErrorWithCode:underlyingError:format:args:"), code, error_, format, unsafe.Pointer(unsafe.StringData(args+"\x00")))
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ProgramEvaluationErrorWithUnderlyingErrorFormat(error_ objectivec.IObject, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("programEvaluationErrorWithUnderlyingError:format:"), error_, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ProgramParsingAtLoadErrorWithReasonFormat(reason int, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("programParsingAtLoadErrorWithReason:format:"), reason, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) ProgramValidationAtLoadErrorWithReasonFormat(reason int, format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("programValidationAtLoadErrorWithReason:format:"), reason, format)
	return objectivec.Object{ID: rv}
}
func (_MLModelErrorUtilsClass MLModelErrorUtilsClass) UpdateErrorWithFormat(format objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelErrorUtilsClass.class), objc.Sel("updateErrorWithFormat:"), format)
	return objectivec.Object{ID: rv}
}
