// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [QuarantineFileHandler] class.
var (
	_QuarantineFileHandlerClass     QuarantineFileHandlerClass
	_QuarantineFileHandlerClassOnce sync.Once
)

func getQuarantineFileHandlerClass() QuarantineFileHandlerClass {
	_QuarantineFileHandlerClassOnce.Do(func() {
		_QuarantineFileHandlerClass = QuarantineFileHandlerClass{class: objc.GetClass("QuarantineFileHandler")}
	})
	return _QuarantineFileHandlerClass
}

// GetQuarantineFileHandlerClass returns the class object for QuarantineFileHandler.
func GetQuarantineFileHandlerClass() QuarantineFileHandlerClass {
	return getQuarantineFileHandlerClass()
}

type QuarantineFileHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (qc QuarantineFileHandlerClass) Class() objc.Class {
	return qc.class
}

// Alloc allocates memory for a new instance of the class.
func (qc QuarantineFileHandlerClass) Alloc() QuarantineFileHandler {
	rv := objc.Send[QuarantineFileHandler](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [QuarantineFileHandler.ApplyMountPointsWithBSDNameError]
//   - [QuarantineFileHandler.CheckErrorWithQtnInitResultError]
//   - [QuarantineFileHandler.GetFileInfoWithError]
//   - [QuarantineFileHandler.IsQuarantined]
//   - [QuarantineFileHandler.QtFile]
//   - [QuarantineFileHandler.SetQtFile]
//   - [QuarantineFileHandler.InitWithBackendError]
//   - [QuarantineFileHandler.InitWithFDError]
//   - [QuarantineFileHandler.InitWithFlagError]
//
// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler
type QuarantineFileHandler struct {
	objectivec.Object
}

// QuarantineFileHandlerFromID constructs a [QuarantineFileHandler] from an objc.ID.
func QuarantineFileHandlerFromID(id objc.ID) QuarantineFileHandler {
	return QuarantineFileHandler{objectivec.Object{ID: id}}
}

// Ensure QuarantineFileHandler implements IQuarantineFileHandler.
var _ IQuarantineFileHandler = QuarantineFileHandler{}

// An interface definition for the [QuarantineFileHandler] class.
//
// # Methods
//
//   - [IQuarantineFileHandler.ApplyMountPointsWithBSDNameError]
//   - [IQuarantineFileHandler.CheckErrorWithQtnInitResultError]
//   - [IQuarantineFileHandler.GetFileInfoWithError]
//   - [IQuarantineFileHandler.IsQuarantined]
//   - [IQuarantineFileHandler.QtFile]
//   - [IQuarantineFileHandler.SetQtFile]
//   - [IQuarantineFileHandler.InitWithBackendError]
//   - [IQuarantineFileHandler.InitWithFDError]
//   - [IQuarantineFileHandler.InitWithFlagError]
//
// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler
type IQuarantineFileHandler interface {
	objectivec.IObject

	// Topic: Methods

	ApplyMountPointsWithBSDNameError(bSDName objectivec.IObject) (bool, error)
	CheckErrorWithQtnInitResultError(result int) (bool, error)
	GetFileInfoWithError() (objectivec.IObject, error)
	IsQuarantined() bool
	QtFile() objectivec.IObject
	SetQtFile(value objectivec.IObject)
	InitWithBackendError(backend objectivec.IObject) (QuarantineFileHandler, error)
	InitWithFDError(fd int) (QuarantineFileHandler, error)
	InitWithFlagError(flag uint32) (QuarantineFileHandler, error)
}

// Init initializes the instance.
func (q QuarantineFileHandler) Init() QuarantineFileHandler {
	rv := objc.Send[QuarantineFileHandler](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q QuarantineFileHandler) Autorelease() QuarantineFileHandler {
	rv := objc.Send[QuarantineFileHandler](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuarantineFileHandler creates a new QuarantineFileHandler instance.
func NewQuarantineFileHandler() QuarantineFileHandler {
	class := getQuarantineFileHandlerClass()
	rv := objc.Send[QuarantineFileHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithBackend:error:
func NewQuarantineFileHandlerWithBackendError(backend objectivec.IObject) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackend:error:"), backend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithFD:error:
func NewQuarantineFileHandlerWithFDError(fd int) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFD:error:"), fd, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithFlag:error:
func NewQuarantineFileHandlerWithFlagError(flag uint32) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFlag:error:"), flag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/applyMountPointsWithBSDName:error:
func (q QuarantineFileHandler) ApplyMountPointsWithBSDNameError(bSDName objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](q.ID, objc.Sel("applyMountPointsWithBSDName:error:"), bSDName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("applyMountPointsWithBSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/checkErrorWithQtnInitResult:error:
func (q QuarantineFileHandler) CheckErrorWithQtnInitResultError(result int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](q.ID, objc.Sel("checkErrorWithQtnInitResult:error:"), result, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkErrorWithQtnInitResult:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/getFileInfoWithError:
func (q QuarantineFileHandler) GetFileInfoWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("getFileInfoWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithBackend:error:
func (q QuarantineFileHandler) InitWithBackendError(backend objectivec.IObject) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithBackend:error:"), backend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithFD:error:
func (q QuarantineFileHandler) InitWithFDError(fd int) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithFD:error:"), fd, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/initWithFlag:error:
func (q QuarantineFileHandler) InitWithFlagError(flag uint32) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithFlag:error:"), flag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/isQuarantined
func (q QuarantineFileHandler) IsQuarantined() bool {
	rv := objc.Send[bool](q.ID, objc.Sel("isQuarantined"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/QuarantineFileHandler/qtFile
func (q QuarantineFileHandler) QtFile() objectivec.IObject {
	rv := objc.Send[objc.ID](q.ID, objc.Sel("qtFile"))
	return objectivec.Object{ID: rv}
}
func (q QuarantineFileHandler) SetQtFile(value objectivec.IObject) {
	objc.Send[struct{}](q.ID, objc.Sel("setQtFile:"), value)
}
