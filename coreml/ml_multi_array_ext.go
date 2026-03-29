package coreml

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// DataPointer returns a pointer to the contiguous data backing this multiarray.
//
// Deprecated by Apple in favor of GetBytesWithHandler/GetMutableBytesWithHandler.
// Still widely used because the scoped alternatives cannot support zero-copy
// handoff to Metal or MLX where the pointer must outlive the callback.
//
// The pointer is valid for the lifetime of the MLMultiArray (and any parent
// prediction result that owns it). Callers must retain the MLMultiArray or its
// owner to keep the pointer valid.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataPointer
func (m MLMultiArray) DataPointer() unsafe.Pointer {
	return objc.Send[unsafe.Pointer](m.ID, objc.Sel("dataPointer"))
}

// GetBytesWithHandler calls the handler with a read-only pointer to the
// multiarray's contiguous data and its size in bytes. The pointer is valid
// only for the duration of the handler call.
//
// This is Apple's non-deprecated replacement for DataPointer (macOS 12.3+,
// iOS 15.4+). Use this when you can complete all work within the callback.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m MLMultiArray) GetBytesWithHandler(handler func(bytes unsafe.Pointer, size int)) {
	block := objc.NewBlock(func(_ objc.Block, bytes unsafe.Pointer, size int) {
		handler(bytes, size)
	})
	defer block.Release()
	objc.Send[objc.ID](m.ID, objc.Sel("getBytesWithHandler:"), objc.ID(block))
}

// GetMutableBytesWithHandler calls the handler with a mutable pointer to
// the multiarray's data, its size in bytes, and the current strides. The
// pointer is valid only for the duration of the handler call.
//
// The strides may differ from the array's Strides property if CoreML has
// relaid out the data internally. Use the provided strides when interpreting
// the buffer.
//
// This is Apple's non-deprecated replacement for DataPointer (macOS 12.3+,
// iOS 15.4+).
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getMutableBytesWithHandler:
func (m MLMultiArray) GetMutableBytesWithHandler(handler func(mutableBytes unsafe.Pointer, size int, strides []foundation.NSNumber)) {
	block := objc.NewBlock(func(_ objc.Block, mutableBytes unsafe.Pointer, size int, stridesID objc.ID) {
		stridesArr := foundation.NSArrayFromID(stridesID)
		count := stridesArr.Count()
		strides := make([]foundation.NSNumber, count)
		for i := uint(0); i < count; i++ {
			obj := stridesArr.ObjectAtIndex(i)
			strides[i] = foundation.NSNumberFromID(obj.GetID())
		}
		handler(mutableBytes, size, strides)
	})
	defer block.Release()
	objc.Send[objc.ID](m.ID, objc.Sel("getMutableBytesWithHandler:"), objc.ID(block))
}
