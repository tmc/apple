// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageCreatorFromFolder] class.
var (
	_DiskImageCreatorFromFolderClass     DiskImageCreatorFromFolderClass
	_DiskImageCreatorFromFolderClassOnce sync.Once
)

func getDiskImageCreatorFromFolderClass() DiskImageCreatorFromFolderClass {
	_DiskImageCreatorFromFolderClassOnce.Do(func() {
		_DiskImageCreatorFromFolderClass = DiskImageCreatorFromFolderClass{class: objc.GetClass("DiskImageCreatorFromFolder")}
	})
	return _DiskImageCreatorFromFolderClass
}

// GetDiskImageCreatorFromFolderClass returns the class object for DiskImageCreatorFromFolder.
func GetDiskImageCreatorFromFolderClass() DiskImageCreatorFromFolderClass {
	return getDiskImageCreatorFromFolderClass()
}

type DiskImageCreatorFromFolderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageCreatorFromFolderClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageCreatorFromFolderClass) Alloc() DiskImageCreatorFromFolder {
	rv := objc.Send[DiskImageCreatorFromFolder](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DiskImageCreatorFromFolder.CompactAndEjectWithCreateParamsError]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderProgressCreateParamsConvertParamsError]
//   - [DiskImageCreatorFromFolder.ResizeDataPartitionWithError]
//   - [DiskImageCreatorFromFolder.UpdateNumBlocksWithFolderSizeNumFiles]
//   - [DiskImageCreatorFromFolder.UpdatePartitionMapWithError]
//   - [DiskImageCreatorFromFolder.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder
type DiskImageCreatorFromFolder struct {
	BaseDiskImageCreator
}

// DiskImageCreatorFromFolderFromID constructs a [DiskImageCreatorFromFolder] from an objc.ID.
func DiskImageCreatorFromFolderFromID(id objc.ID) DiskImageCreatorFromFolder {
	return DiskImageCreatorFromFolder{BaseDiskImageCreator: BaseDiskImageCreatorFromID(id)}
}

// Ensure DiskImageCreatorFromFolder implements IDiskImageCreatorFromFolder.
var _ IDiskImageCreatorFromFolder = DiskImageCreatorFromFolder{}

// An interface definition for the [DiskImageCreatorFromFolder] class.
//
// # Methods
//
//   - [IDiskImageCreatorFromFolder.CompactAndEjectWithCreateParamsError]
//   - [IDiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
//   - [IDiskImageCreatorFromFolder.CreateImageWithSrcFolderProgressCreateParamsConvertParamsError]
//   - [IDiskImageCreatorFromFolder.ResizeDataPartitionWithError]
//   - [IDiskImageCreatorFromFolder.UpdateNumBlocksWithFolderSizeNumFiles]
//   - [IDiskImageCreatorFromFolder.UpdatePartitionMapWithError]
//   - [IDiskImageCreatorFromFolder.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder
type IDiskImageCreatorFromFolder interface {
	IBaseDiskImageCreator

	// Topic: Methods

	CompactAndEjectWithCreateParamsError(params objectivec.IObject) (bool, error)
	CreateImageWithSrcFolderCompletionBlock(folder objectivec.IObject, block VoidHandler) objectivec.IObject
	CreateImageWithSrcFolderProgressCreateParamsConvertParamsError(folder objectivec.IObject, progress objectivec.IObject, params objectivec.IObject, params2 []objectivec.IObject) (bool, error)
	ResizeDataPartitionWithError() (bool, error)
	UpdateNumBlocksWithFolderSizeNumFiles(size uint64, files uint64)
	UpdatePartitionMapWithError() (bool, error)
	InitWithURLError(url foundation.INSURL) (DiskImageCreatorFromFolder, error)
}

// Init initializes the instance.
func (d DiskImageCreatorFromFolder) Init() DiskImageCreatorFromFolder {
	rv := objc.Send[DiskImageCreatorFromFolder](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageCreatorFromFolder) Autorelease() DiskImageCreatorFromFolder {
	rv := objc.Send[DiskImageCreatorFromFolder](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageCreatorFromFolder creates a new DiskImageCreatorFromFolder instance.
func NewDiskImageCreatorFromFolder() DiskImageCreatorFromFolder {
	class := getDiskImageCreatorFromFolderClass()
	rv := objc.Send[DiskImageCreatorFromFolder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/initWithURL:defaultFormat:error:
func NewDiskImageCreatorFromFolderWithURLDefaultFormatError(url foundation.INSURL, format int64) (DiskImageCreatorFromFolder, error) {
	var errorPtr objc.ID
	instance := getDiskImageCreatorFromFolderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromFolder{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageCreatorFromFolderFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/initWithURL:error:
func NewDiskImageCreatorFromFolderWithURLError(url foundation.INSURL) (DiskImageCreatorFromFolder, error) {
	var errorPtr objc.ID
	instance := getDiskImageCreatorFromFolderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromFolder{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageCreatorFromFolderFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/compactAndEjectWithCreateParams:error:
func (d DiskImageCreatorFromFolder) CompactAndEjectWithCreateParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("compactAndEjectWithCreateParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("compactAndEjectWithCreateParams:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/createImageWithSrcFolder:completionBlock:
func (d DiskImageCreatorFromFolder) CreateImageWithSrcFolderCompletionBlock(folder objectivec.IObject, block VoidHandler) objectivec.IObject {
	_block1, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createImageWithSrcFolder:completionBlock:"), folder, _block1)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/createImageWithSrcFolder:progress:createParams:convertParams:error:
func (d DiskImageCreatorFromFolder) CreateImageWithSrcFolderProgressCreateParamsConvertParamsError(folder objectivec.IObject, progress objectivec.IObject, params objectivec.IObject, params2 []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createImageWithSrcFolder:progress:createParams:convertParams:error:"), folder, progress, params, objectivec.IObjectSliceToNSArray(params2), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createImageWithSrcFolder:progress:createParams:convertParams:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/resizeDataPartitionWithError:
func (d DiskImageCreatorFromFolder) ResizeDataPartitionWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeDataPartitionWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeDataPartitionWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/updateNumBlocksWithFolderSize:numFiles:
func (d DiskImageCreatorFromFolder) UpdateNumBlocksWithFolderSizeNumFiles(size uint64, files uint64) {
	objc.Send[objc.ID](d.ID, objc.Sel("updateNumBlocksWithFolderSize:numFiles:"), size, files)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/updatePartitionMapWithError:
func (d DiskImageCreatorFromFolder) UpdatePartitionMapWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("updatePartitionMapWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updatePartitionMapWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/initWithURL:error:
func (d DiskImageCreatorFromFolder) InitWithURLError(url foundation.INSURL) (DiskImageCreatorFromFolder, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageCreatorFromFolder{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageCreatorFromFolderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageCreatorFromFolder/allowParallelModeWithURL:outMode:error:
func (_DiskImageCreatorFromFolderClass DiskImageCreatorFromFolderClass) AllowParallelModeWithURLOutModeError(url foundation.NSURL) (bool, error) {
	var mode bool
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageCreatorFromFolderClass.class), objc.Sel("allowParallelModeWithURL:outMode:error:"), url, unsafe.Pointer(&mode), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("allowParallelModeWithURL:outMode:error: returned NO with nil NSError")
	}
	return mode, nil
}

// CreateImageWithSrcFolderCompletionBlockSync is a synchronous wrapper around [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d DiskImageCreatorFromFolder) CreateImageWithSrcFolderCompletionBlockSync(ctx context.Context, folder objectivec.IObject) error {
	done := make(chan struct{}, 1)
	d.CreateImageWithSrcFolderCompletionBlock(folder, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
