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
	rv := objc.SendIfResponds[QuarantineFileHandler](objc.ID(qc.class), objc.Sel("alloc"))
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
type IQuarantineFileHandler interface {
	objectivec.IObject

	// Topic: Methods

	ApplyMountPointsWithBSDNameError(bSDName objectivec.IObject) (bool, error)
	CheckErrorWithQtnInitResultError(result int) (bool, error)
	GetFileInfoWithError() (objectivec.IObject, error)
	IsQuarantined() bool
	QtFile() QtnFileRef
	SetQtFile(value QtnFileRef)
	InitWithBackendError(backend unsafe.Pointer) (QuarantineFileHandler, error)
	InitWithFDError(fd int) (QuarantineFileHandler, error)
	InitWithFlagError(flag uint32) (QuarantineFileHandler, error)
}

// Init initializes the instance.
func (q QuarantineFileHandler) Init() QuarantineFileHandler {
	rv := objc.SendIfResponds[QuarantineFileHandler](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q QuarantineFileHandler) Autorelease() QuarantineFileHandler {
	rv := objc.SendIfResponds[QuarantineFileHandler](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuarantineFileHandler creates a new QuarantineFileHandler instance.
func NewQuarantineFileHandler() QuarantineFileHandler {
	class := getQuarantineFileHandlerClass()
	rv := objc.SendIfResponds[QuarantineFileHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewQuarantineFileHandlerWithBackendError(backend unsafe.Pointer) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackend:error:"), backend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return QuarantineFileHandler{}, objc.ErrInitFailed
	}
	return QuarantineFileHandlerFromID(rv), nil
}

func NewQuarantineFileHandlerWithFDError(fd int) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFD:error:"), fd, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return QuarantineFileHandler{}, objc.ErrInitFailed
	}
	return QuarantineFileHandlerFromID(rv), nil
}

func NewQuarantineFileHandlerWithFlagError(flag uint32) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	instance := getQuarantineFileHandlerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFlag:error:"), flag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return QuarantineFileHandler{}, objc.ErrInitFailed
	}
	return QuarantineFileHandlerFromID(rv), nil
}

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
func (q QuarantineFileHandler) GetFileInfoWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("getFileInfoWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (q QuarantineFileHandler) InitWithBackendError(backend unsafe.Pointer) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithBackend:error:"), backend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}
func (q QuarantineFileHandler) InitWithFDError(fd int) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithFD:error:"), fd, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}
func (q QuarantineFileHandler) InitWithFlagError(flag uint32) (QuarantineFileHandler, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](q.ID, objc.Sel("initWithFlag:error:"), flag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return QuarantineFileHandler{}, foundation.NSErrorFrom(errorPtr)
	}
	return QuarantineFileHandlerFromID(rv), nil

}

func (q QuarantineFileHandler) IsQuarantined() bool {
	rv := objc.SendIfResponds[bool](q.ID, objc.Sel("isQuarantined"))
	return rv
}
func (q QuarantineFileHandler) QtFile() QtnFileRef {
	rv := objc.SendIfResponds[objc.ID](q.ID, objc.Sel("qtFile"))
	return QtnFileRef(rv)
}
func (q QuarantineFileHandler) SetQtFile(value QtnFileRef) {
	objc.SendIfResponds[struct{}](q.ID, objc.Sel("setQtFile:"), value)
}
