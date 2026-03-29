// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIError] class.
var (
	_DIErrorClass     DIErrorClass
	_DIErrorClassOnce sync.Once
)

func getDIErrorClass() DIErrorClass {
	_DIErrorClassOnce.Do(func() {
		_DIErrorClass = DIErrorClass{class: objc.GetClass("DIError")}
	})
	return _DIErrorClass
}

// GetDIErrorClass returns the class object for DIError.
func GetDIErrorClass() DIErrorClass {
	return getDIErrorClass()
}

type DIErrorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIErrorClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIErrorClass) Alloc() DIError {
	rv := objc.Send[DIError](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIError
type DIError struct {
	objectivec.Object
}

// DIErrorFromID constructs a [DIError] from an objc.ID.
func DIErrorFromID(id objc.ID) DIError {
	return DIError{objectivec.Object{ID: id}}
}
// Ensure DIError implements IDIError.
var _ IDIError = DIError{}

// An interface definition for the [DIError] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError
type IDIError interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DIError) Init() DIError {
	rv := objc.Send[DIError](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIError) Autorelease() DIError {
	rv := objc.Send[DIError](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIError creates a new DIError instance.
func NewDIError() DIError {
	class := getDIErrorClass()
	rv := objc.Send[DIError](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/copyDefaultLocalizedStringForDIErrorCode:
func (_DIErrorClass DIErrorClass) CopyDefaultLocalizedStringForDIErrorCode(code int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("copyDefaultLocalizedStringForDIErrorCode:"), code)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/errorWithDIException:description:prefix:error:
func (_DIErrorClass DIErrorClass) ErrorWithDIExceptionDescriptionPrefixError(dIException unsafe.Pointer, description objectivec.IObject, prefix objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("errorWithDIException:description:prefix:error:"), dIException, description, prefix, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/errorWithDomain:code:description:verboseInfo:error:
func (_DIErrorClass DIErrorClass) ErrorWithDomainCodeDescriptionVerboseInfoError(domain objectivec.IObject, code int64, description objectivec.IObject, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("errorWithDomain:code:description:verboseInfo:error:"), domain, code, description, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/errorWithEnumValue:verboseInfo:
func (_DIErrorClass DIErrorClass) ErrorWithEnumValueVerboseInfo(value int64, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("errorWithEnumValue:verboseInfo:"), value, info)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/errorWithPOSIXCode:verboseInfo:
func (_DIErrorClass DIErrorClass) ErrorWithPOSIXCodeVerboseInfo(pOSIXCode int, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("errorWithPOSIXCode:verboseInfo:"), pOSIXCode, info)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/errorWithUnexpected:verboseInfo:error:
func (_DIErrorClass DIErrorClass) ErrorWithUnexpectedVerboseInfoError(unexpected objectivec.IObject, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("errorWithUnexpected:verboseInfo:error:"), unexpected, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithDIException:description:error:
func (_DIErrorClass DIErrorClass) FailWithDIExceptionDescriptionError(dIException unsafe.Pointer, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithDIException:description:error:"), dIException, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithDIException:description:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithDIException:prefix:error:
func (_DIErrorClass DIErrorClass) FailWithDIExceptionPrefixError(dIException unsafe.Pointer, prefix objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithDIException:prefix:error:"), dIException, prefix, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithDIException:prefix:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithEnumValue:description:error:
func (_DIErrorClass DIErrorClass) FailWithEnumValueDescriptionError(value int64, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithEnumValue:description:error:"), value, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithEnumValue:description:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithEnumValue:verboseInfo:error:
func (_DIErrorClass DIErrorClass) FailWithEnumValueVerboseInfoError(value int64, info objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithEnumValue:verboseInfo:error:"), value, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithEnumValue:verboseInfo:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithInError:outError:
func (_DIErrorClass DIErrorClass) FailWithInErrorOutError(error_ objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithInError:outError:"), error_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithInError:outError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithOSStatus:description:error:
func (_DIErrorClass DIErrorClass) FailWithOSStatusDescriptionError(oSStatus int, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithOSStatus:description:error:"), oSStatus, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithOSStatus:description:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithOSStatus:verboseInfo:error:
func (_DIErrorClass DIErrorClass) FailWithOSStatusVerboseInfoError(oSStatus int, info objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithOSStatus:verboseInfo:error:"), oSStatus, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithOSStatus:verboseInfo:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithPOSIXCode:description:error:
func (_DIErrorClass DIErrorClass) FailWithPOSIXCodeDescriptionError(pOSIXCode int, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithPOSIXCode:description:error:"), pOSIXCode, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithPOSIXCode:description:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithPOSIXCode:error:
func (_DIErrorClass DIErrorClass) FailWithPOSIXCodeError(pOSIXCode int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithPOSIXCode:error:"), pOSIXCode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithPOSIXCode:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithPOSIXCode:verboseInfo:error:
func (_DIErrorClass DIErrorClass) FailWithPOSIXCodeVerboseInfoError(pOSIXCode int, info objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithPOSIXCode:verboseInfo:error:"), pOSIXCode, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithPOSIXCode:verboseInfo:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithUnexpected:error:
func (_DIErrorClass DIErrorClass) FailWithUnexpectedError(unexpected objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithUnexpected:error:"), unexpected, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithUnexpected:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/failWithUnexpected:verboseInfo:error:
func (_DIErrorClass DIErrorClass) FailWithUnexpectedVerboseInfoError(unexpected objectivec.IObject, info objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("failWithUnexpected:verboseInfo:error:"), unexpected, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithUnexpected:verboseInfo:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DIError/frameworkBundle
func (_DIErrorClass DIErrorClass) FrameworkBundle() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("frameworkBundle"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/mandatoryArgumentFailWithError:
func (_DIErrorClass DIErrorClass) MandatoryArgumentFailWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIErrorClass.class), objc.Sel("mandatoryArgumentFailWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("mandatoryArgumentFailWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithDIException:description:error:
func (_DIErrorClass DIErrorClass) NilWithDIExceptionDescriptionError(dIException unsafe.Pointer, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithDIException:description:error:"), dIException, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithDIException:prefix:error:
func (_DIErrorClass DIErrorClass) NilWithDIExceptionPrefixError(dIException unsafe.Pointer, prefix objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithDIException:prefix:error:"), dIException, prefix, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithEnumValue:description:error:
func (_DIErrorClass DIErrorClass) NilWithEnumValueDescriptionError(value int64, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithEnumValue:description:error:"), value, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithEnumValue:verboseInfo:error:
func (_DIErrorClass DIErrorClass) NilWithEnumValueVerboseInfoError(value int64, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithEnumValue:verboseInfo:error:"), value, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithOSStatus:verboseInfo:error:
func (_DIErrorClass DIErrorClass) NilWithOSStatusVerboseInfoError(oSStatus int, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithOSStatus:verboseInfo:error:"), oSStatus, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithPOSIXCode:description:error:
func (_DIErrorClass DIErrorClass) NilWithPOSIXCodeDescriptionError(pOSIXCode int, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithPOSIXCode:description:error:"), pOSIXCode, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithPOSIXCode:verboseInfo:error:
func (_DIErrorClass DIErrorClass) NilWithPOSIXCodeVerboseInfoError(pOSIXCode int, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithPOSIXCode:verboseInfo:error:"), pOSIXCode, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithUnexpected:error:
func (_DIErrorClass DIErrorClass) NilWithUnexpectedError(unexpected objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithUnexpected:error:"), unexpected, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIError/nilWithUnexpected:verboseInfo:error:
func (_DIErrorClass DIErrorClass) NilWithUnexpectedVerboseInfoError(unexpected objectivec.IObject, info objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIErrorClass.class), objc.Sel("nilWithUnexpected:verboseInfo:error:"), unexpected, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

