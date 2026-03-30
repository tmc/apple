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

// The class instance for the [FastFolderCopierWrapper] class.
var (
	_FastFolderCopierWrapperClass     FastFolderCopierWrapperClass
	_FastFolderCopierWrapperClassOnce sync.Once
)

func getFastFolderCopierWrapperClass() FastFolderCopierWrapperClass {
	_FastFolderCopierWrapperClassOnce.Do(func() {
		_FastFolderCopierWrapperClass = FastFolderCopierWrapperClass{class: objc.GetClass("FastFolderCopierWrapper")}
	})
	return _FastFolderCopierWrapperClass
}

// GetFastFolderCopierWrapperClass returns the class object for FastFolderCopierWrapper.
func GetFastFolderCopierWrapperClass() FastFolderCopierWrapperClass {
	return getFastFolderCopierWrapperClass()
}

type FastFolderCopierWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FastFolderCopierWrapperClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FastFolderCopierWrapperClass) Alloc() FastFolderCopierWrapper {
	rv := objc.Send[FastFolderCopierWrapper](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [FastFolderCopierWrapper.Copier]
//   - [FastFolderCopierWrapper.SetCopier]
//   - [FastFolderCopierWrapper.CopyWithDstFolderProgressError]
//   - [FastFolderCopierWrapper.FolderSize]
//   - [FastFolderCopierWrapper.NumFiles]
//   - [FastFolderCopierWrapper.Progress]
//   - [FastFolderCopierWrapper.SetProgress]
//   - [FastFolderCopierWrapper.TraverseSrcFolderWithProgressError]
//   - [FastFolderCopierWrapper.InitWithSrcFolderParallelModeAuditToken]
//
// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper
type FastFolderCopierWrapper struct {
	objectivec.Object
}

// FastFolderCopierWrapperFromID constructs a [FastFolderCopierWrapper] from an objc.ID.
func FastFolderCopierWrapperFromID(id objc.ID) FastFolderCopierWrapper {
	return FastFolderCopierWrapper{objectivec.Object{ID: id}}
}

// Ensure FastFolderCopierWrapper implements IFastFolderCopierWrapper.
var _ IFastFolderCopierWrapper = FastFolderCopierWrapper{}

// An interface definition for the [FastFolderCopierWrapper] class.
//
// # Methods
//
//   - [IFastFolderCopierWrapper.Copier]
//   - [IFastFolderCopierWrapper.SetCopier]
//   - [IFastFolderCopierWrapper.CopyWithDstFolderProgressError]
//   - [IFastFolderCopierWrapper.FolderSize]
//   - [IFastFolderCopierWrapper.NumFiles]
//   - [IFastFolderCopierWrapper.Progress]
//   - [IFastFolderCopierWrapper.SetProgress]
//   - [IFastFolderCopierWrapper.TraverseSrcFolderWithProgressError]
//   - [IFastFolderCopierWrapper.InitWithSrcFolderParallelModeAuditToken]
//
// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper
type IFastFolderCopierWrapper interface {
	objectivec.IObject

	// Topic: Methods

	Copier() objectivec.IObject
	SetCopier(value objectivec.IObject)
	CopyWithDstFolderProgressError(folder objectivec.IObject, progress objectivec.IObject) (bool, error)
	FolderSize() uint64
	NumFiles() uint64
	Progress() *foundation.NSProgress
	SetProgress(value *foundation.NSProgress)
	TraverseSrcFolderWithProgressError(progress objectivec.IObject) (bool, error)
	InitWithSrcFolderParallelModeAuditToken(folder objectivec.IObject, mode bool, token objectivec.IObject) FastFolderCopierWrapper
}

// Init initializes the instance.
func (f FastFolderCopierWrapper) Init() FastFolderCopierWrapper {
	rv := objc.Send[FastFolderCopierWrapper](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FastFolderCopierWrapper) Autorelease() FastFolderCopierWrapper {
	rv := objc.Send[FastFolderCopierWrapper](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFastFolderCopierWrapper creates a new FastFolderCopierWrapper instance.
func NewFastFolderCopierWrapper() FastFolderCopierWrapper {
	class := getFastFolderCopierWrapperClass()
	rv := objc.Send[FastFolderCopierWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/initWithSrcFolder:parallelMode:auditToken:
func NewFastFolderCopierWrapperWithSrcFolderParallelModeAuditToken(folder objectivec.IObject, mode bool, token objectivec.IObject) FastFolderCopierWrapper {
	instance := getFastFolderCopierWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSrcFolder:parallelMode:auditToken:"), folder, mode, token)
	return FastFolderCopierWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/copyWithDstFolder:progress:error:
func (f FastFolderCopierWrapper) CopyWithDstFolderProgressError(folder objectivec.IObject, progress objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("copyWithDstFolder:progress:error:"), folder, progress, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyWithDstFolder:progress:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/traverseSrcFolderWithProgress:error:
func (f FastFolderCopierWrapper) TraverseSrcFolderWithProgressError(progress objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("traverseSrcFolderWithProgress:error:"), progress, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("traverseSrcFolderWithProgress:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/initWithSrcFolder:parallelMode:auditToken:
func (f FastFolderCopierWrapper) InitWithSrcFolderParallelModeAuditToken(folder objectivec.IObject, mode bool, token objectivec.IObject) FastFolderCopierWrapper {
	rv := objc.Send[FastFolderCopierWrapper](f.ID, objc.Sel("initWithSrcFolder:parallelMode:auditToken:"), folder, mode, token)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/copier
func (f FastFolderCopierWrapper) Copier() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("copier"))
	return objectivec.Object{ID: rv}
}
func (f FastFolderCopierWrapper) SetCopier(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setCopier:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/folderSize
func (f FastFolderCopierWrapper) FolderSize() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("folderSize"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/numFiles
func (f FastFolderCopierWrapper) NumFiles() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("numFiles"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/FastFolderCopierWrapper/progress
func (f FastFolderCopierWrapper) Progress() *foundation.NSProgress {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("progress"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSProgressFromID(objc.ID(rv))
	return &val
}
func (f FastFolderCopierWrapper) SetProgress(value *foundation.NSProgress) {
	if value == nil {
		objc.Send[struct{}](f.ID, objc.Sel("setProgress:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](f.ID, objc.Sel("setProgress:"), value)
}
