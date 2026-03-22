// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A command buffer that contains input/output commands that work with files in the file systems and Metal resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer
type MTLIOCommandBuffer interface {
	objectivec.IObject

	// Encodes a command that loads data from a file handle into a GPU buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/load(_:offset:size:sourceHandle:sourceHandleOffset:)
	LoadBufferOffsetSizeSourceHandleSourceHandleOffset(buffer MTLBuffer, offset uint, size uint, sourceHandle MTLIOFileHandle, sourceHandleOffset uint)

	// Encodes a command that loads data from a file handle into a GPU texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/load(_:slice:level:size:sourceBytesPerRow:sourceBytesPerImage:destinationOrigin:sourceHandle:sourceHandleOffset:)
	LoadTextureSliceLevelSizeSourceBytesPerRowSourceBytesPerImageDestinationOriginSourceHandleSourceHandleOffset(texture MTLTexture, slice uint, level uint, size MTLSize, sourceBytesPerRow uint, sourceBytesPerImage uint, destinationOrigin MTLOrigin, sourceHandle MTLIOFileHandle, sourceHandleOffset uint)

	// Encodes a command that loads data from a file handle into CPU-accessible memory buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/loadBytes(_:size:sourceHandle:sourceHandleOffset:)
	LoadBytesSizeSourceHandleSourceHandleOffset(pointer unsafe.Pointer, size uint, sourceHandle MTLIOFileHandle, sourceHandleOffset uint)

	// Encodes a barrier into the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/addBarrier()
	AddBarrier()

	// Encodes a command that signals a shared event to other parts of your app.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/signalEvent(_:value:)
	SignalEventValue(event MTLSharedEvent, value uint64)

	// Encodes a command that pauses the command buffer’s execution until another part of your app signals a shared event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/waitForEvent(_:value:)
	WaitForEventValue(event MTLSharedEvent, value uint64)

	// Encodes a command that writes the input/output command buffer’s status to a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/copyStatus(buffer:offset:)
	CopyStatusToBufferOffset(buffer MTLBuffer, offset uint)

	// Adds a closure that Metal calls immediately after the GPU finishes executing the commands in the input/output command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/addCompletedHandler(_:)
	AddCompletedHandler(block MTLIOCommandBufferHandler)

	// Submits the command buffer to the queue for execution on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/commit()
	Commit()

	// Reserves a place for the input/output command buffer in the input/output command queue without committing the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/enqueue()
	Enqueue()

	// Submits a request to abandon a command buffer the queue is currently running.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/tryCancel()
	TryCancel()

	// Blocks the current thread until the GPU finishes executing the input/output command buffer and all of its completion handlers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/waitUntilCompleted()
	WaitUntilCompleted()

	// Represents the state of the input/output command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/status
	Status() MTLIOStatus

	// Stores the details of an error when the GPU experienced a problem with the input/output command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/error
	Error() foundation.INSError

	// An optional name for the input/output command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/label
	Label() string

	// Sets the current name for this input/output command encoder by adding it to the top of the debug name stack.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/pushDebugGroup(_:)
	PushDebugGroup(string_ string)

	// Restores the previous name for this input/output command encoder by removing the top item of the debug name stack.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/popDebugGroup()
	PopDebugGroup()

	// An optional name for the input/output command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/label
	SetLabel(value string)
}

// MTLIOCommandBufferObject wraps an existing Objective-C object that conforms to the MTLIOCommandBuffer protocol.
type MTLIOCommandBufferObject struct {
	objectivec.Object
}
func (o MTLIOCommandBufferObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIOCommandBufferObjectFromID constructs a [MTLIOCommandBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIOCommandBufferObjectFromID(id objc.ID) MTLIOCommandBufferObject {
	return MTLIOCommandBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes a command that loads data from a file handle into a GPU buffer.
//
// buffer: A buffer instance the method loads data into.
//
// offset: A starting location relative to the beginning of the buffer, in bytes, the
// method copies data to.
//
// size: The number of bytes the method loads from the file into the buffer.
//
// sourceHandle: A handle to a source file.
//
// sourceHandleOffset: A starting location relative to the beginning of the file, in bytes, the
// method copies data from.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/load(_:offset:size:sourceHandle:sourceHandleOffset:)
func (o MTLIOCommandBufferObject) LoadBufferOffsetSizeSourceHandleSourceHandleOffset(buffer MTLBuffer, offset uint, size uint, sourceHandle MTLIOFileHandle, sourceHandleOffset uint) {
	objc.Send[struct{}](o.ID, objc.Sel("loadBuffer:offset:size:sourceHandle:sourceHandleOffset:"), buffer, offset, size, sourceHandle, sourceHandleOffset)
	}
// Encodes a command that loads data from a file handle into a GPU texture.
//
// texture: A texture instance the method loads data into.
//
// slice: A slice within the texture.
//
// level: A level within the texture.
//
// size: The region of the texture the method copies to.
//
// sourceBytesPerRow: The number of bytes in a row of data from the source file.
//
// sourceBytesPerImage: The number of bytes in an image from the source file.
//
// destinationOrigin: A starting location within the texture the method copies data to.
//
// sourceHandle: A handle to a source file.
//
// sourceHandleOffset: A starting location relative to the beginning of the file, in bytes, the
// method copies data from.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/load(_:slice:level:size:sourceBytesPerRow:sourceBytesPerImage:destinationOrigin:sourceHandle:sourceHandleOffset:)
func (o MTLIOCommandBufferObject) LoadTextureSliceLevelSizeSourceBytesPerRowSourceBytesPerImageDestinationOriginSourceHandleSourceHandleOffset(texture MTLTexture, slice uint, level uint, size MTLSize, sourceBytesPerRow uint, sourceBytesPerImage uint, destinationOrigin MTLOrigin, sourceHandle MTLIOFileHandle, sourceHandleOffset uint) {
	objc.Send[struct{}](o.ID, objc.Sel("loadTexture:slice:level:size:sourceBytesPerRow:sourceBytesPerImage:destinationOrigin:sourceHandle:sourceHandleOffset:"), texture, slice, level, size, sourceBytesPerRow, sourceBytesPerImage, destinationOrigin, sourceHandle, sourceHandleOffset)
	}
// Encodes a command that loads data from a file handle into CPU-accessible
// memory buffer.
//
// pointer: A pointer to memory the method loads data into.
//
// size: The number of bytes the method loads from the file.
//
// sourceHandle: A handle to a source file.
//
// sourceHandleOffset: A starting location relative to the beginning of the file, in bytes, the
// method copies data from.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/loadBytes(_:size:sourceHandle:sourceHandleOffset:)
func (o MTLIOCommandBufferObject) LoadBytesSizeSourceHandleSourceHandleOffset(pointer unsafe.Pointer, size uint, sourceHandle MTLIOFileHandle, sourceHandleOffset uint) {
	objc.Send[struct{}](o.ID, objc.Sel("loadBytes:size:sourceHandle:sourceHandleOffset:"), pointer, size, sourceHandle, sourceHandleOffset)
	}
// Encodes a barrier into the command buffer.
//
// # Discussion
// 
// The method encodes a barrier that starts any subsequent commands only after
// all the previously encoded commands have completed.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/addBarrier()
func (o MTLIOCommandBufferObject) AddBarrier() {
	objc.Send[struct{}](o.ID, objc.Sel("addBarrier"))
	}
// Encodes a command that signals a shared event to other parts of your app.
//
// event: A shared event instance the method waits for.
//
// value: A value the command uses to signal for the event to other parts of your
// app.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/signalEvent(_:value:)
func (o MTLIOCommandBufferObject) SignalEventValue(event MTLSharedEvent, value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("signalEvent:value:"), event, value)
	}
// Encodes a command that pauses the command buffer’s execution until
// another part of your app signals a shared event.
//
// event: A shared event instance the method waits for.
//
// value: A value the method compares to the event’s value. The method returns when
// the event’s value is greater than or equal to `value`.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/waitForEvent(_:value:)
func (o MTLIOCommandBufferObject) WaitForEventValue(event MTLSharedEvent, value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("waitForEvent:value:"), event, value)
	}
// Encodes a command that writes the input/output command buffer’s status to
// a buffer.
//
// buffer: A buffer instance the method copies the status into.
//
// offset: A starting location relative to the beginning of the buffer, in bytes, the
// method copies data to.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/copyStatus(buffer:offset:)
func (o MTLIOCommandBufferObject) CopyStatusToBufferOffset(buffer MTLBuffer, offset uint) {
	objc.Send[struct{}](o.ID, objc.Sel("copyStatusToBuffer:offset:"), buffer, offset)
	}
// Adds a closure that Metal calls immediately after the GPU finishes
// executing the commands in the input/output command buffer.
//
// block: A Swift closure or an Objective-C block with your code.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/addCompletedHandler(_:)
func (o MTLIOCommandBufferObject) AddCompletedHandler(block MTLIOCommandBufferHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("addCompletedHandler:"), block)
	}
// Submits the command buffer to the queue for execution on the GPU.
//
// # Discussion
// 
// If you haven’t already called [Enqueue] for the command buffer, the
// [Commit] method enqueues it at the next position in the input/output
// command queue.
// 
// You can only commit an input/output command buffer once, after which you
// can’t encode any additional commands or add more completion handlers to
// it.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/commit()
func (o MTLIOCommandBufferObject) Commit() {
	objc.Send[struct{}](o.ID, objc.Sel("commit"))
	}
// Reserves a place for the input/output command buffer in the input/output
// command queue without committing the command buffer.
//
// # Discussion
// 
// The method saves the next position for the command buffer in the
// input/output command queue. You can call [Enqueue] at any time relative to
// encoding commands, but you can only enqueue a command buffer once. To
// submit a command buffer to GPU for execution, call its [Commit] method.
// 
// For example, to fill multiple command buffers asynchronously that execute
// in a specific order:
// 
// - Call each command buffer’s [Enqueue] method in order. - Encode commands
// into each command buffer on its own, separate thread. - Call each command
// buffer’s [Commit] in any order.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/enqueue()
func (o MTLIOCommandBufferObject) Enqueue() {
	objc.Send[struct{}](o.ID, objc.Sel("enqueue"))
	}
// Submits a request to abandon a command buffer the queue is currently
// running.
//
// # Discussion
// 
// Check the command buffer’s [Status] property after it completes, either
// after [WaitUntilCompleted] or in one of your completion handlers (see
// [AddCompletedHandler]).
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/tryCancel()
func (o MTLIOCommandBufferObject) TryCancel() {
	objc.Send[struct{}](o.ID, objc.Sel("tryCancel"))
	}
// Blocks the current thread until the GPU finishes executing the input/output
// command buffer and all of its completion handlers.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/waitUntilCompleted()
func (o MTLIOCommandBufferObject) WaitUntilCompleted() {
	objc.Send[struct{}](o.ID, objc.Sel("waitUntilCompleted"))
	}
// Represents the state of the input/output command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/status
func (o MTLIOCommandBufferObject) Status() MTLIOStatus {
	rv := objc.Send[MTLIOStatus](o.ID, objc.Sel("status"))
	return rv
	}
// Stores the details of an error when the GPU experienced a problem with the
// input/output command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/error
func (o MTLIOCommandBufferObject) Error() foundation.INSError {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(rv)
	}
// An optional name for the input/output command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/label
func (o MTLIOCommandBufferObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the current name for this input/output command encoder by adding it to
// the top of the debug name stack.
//
// string: A new debugging name.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/pushDebugGroup(_:)
func (o MTLIOCommandBufferObject) PushDebugGroup(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
	}
// Restores the previous name for this input/output command encoder by
// removing the top item of the debug name stack.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBuffer/popDebugGroup()
func (o MTLIOCommandBufferObject) PopDebugGroup() {
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}

func (o MTLIOCommandBufferObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

