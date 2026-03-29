// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia
import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
)

// C struct types
// CMBlockBufferCustomBlockSource - A structure to support custom memory allocation and deallocation for a block used in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCustomBlockSource
type CMBlockBufferCustomBlockSource struct {
	AllocateBlock func(unsafe.Pointer, uint) unsafe.Pointer // The function to allocate memory.
	FreeBlock func(unsafe.Pointer, unsafe.Pointer, uint) // A function to call once when the [CMBlockBuffer] is disposed.
	RefCon unsafe.Pointer // Contextual information passed to both the allocate and free function calls.
	Version uint32

}

// CMBufferCallbacks - A structure that stores the callbacks that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferCallbacks
type CMBufferCallbacks struct {
	Compare CMBufferCompareCallback // This callback is called multiple times from [CMBufferQueueEnqueue(_:buffer:)](<doc://com.apple.coremedia/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)>), to perform an insertion sort.
	DataBecameReadyNotification corefoundation.CFStringRef // If triggers of type `kCMBufferQueueTrigger_WhenDataBecomesReady` are installed, the queue will listen for this notification on the head buffer.
	GetDecodeTimeStamp CMBufferGetTimeCallback // Client callback that returns a [CMTime] from a [CMBuffer].
	GetDuration CMBufferGetTimeCallback // This callback is called (once) during enqueue and dequeue operations to update the total duration of the queue.
	GetPresentationTimeStamp CMBufferGetTimeCallback // Client callback that returns a [CMTime] from a [CMBuffer].
	GetSize CMBufferGetSizeCallback
	IsDataReady CMBufferGetBooleanCallback // This callback is called from [CMBufferQueueDequeueIfDataReady(_:)](<doc://com.apple.coremedia/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)>), to ask if the buffer that is about to be dequeued is ready.
	Refcon unsafe.Pointer // Contextual data to be passed to all callbacks.
	Version uint32 // The callback version.

}

// CMBufferHandlers - A structure that stores the handlers that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferHandlers
type CMBufferHandlers struct {
	Compare CMBufferCompareHandler // A handler callback the queue uses to perform an insertion sort of the queue.
	GetDuration CMBufferGetTimeHandler
	GetDecodeTimeStamp CMBufferGetTimeHandler
	GetPresentationTimeStamp CMBufferGetTimeHandler
	GetSize CMBufferGetSizeHandler
	IsDataReady CMBufferGetBooleanHandler
	DataBecameReadyNotification corefoundation.CFStringRef
	Version uintptr // The version number.

}

// CMSampleTimingInfo - A collection of timing information for a sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleTimingInfo
type CMSampleTimingInfo struct {
	DecodeTimeStamp CMTime // The time at which the sample will be decoded.
	Duration CMTime // The duration of the sample.
	PresentationTimeStamp CMTime // The time at which the sample will be presented.

}

// CMTag - A tag representing additional metadata on tagged media buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTag-c.struct
type CMTag struct {
	Category CMTagCategory // The category assigned to a tag.
	DataType CMTagDataType // The data type for the value stored in the tag.
	Value CMTagValue // The value of the tag.

}

// CMTime - A structure that represents time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTime
type CMTime struct {
	Value int64 // A time value that represents the numerator of a rational time.
	Timescale int32 // A timescale that represents the denominator of a rational time.
	Flags uint32 // The flags associated with a time.
	Epoch int64 // The epoch of the time.

}

// CMTimeMapping - A structure that maps a segment of a source time range to a target time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapping
type CMTimeMapping struct {
	Source CMTimeRange // A time range on the source timeline.
	Target CMTimeRange // A time range on the target timeline.

}

// CMTimeRange - A structure that represents a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
type CMTimeRange struct {
	Start CMTime // The start time of the time range.
	Duration CMTime // The duration of the time range.

}

// CMVideoDimensions - A structure that represents video dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoDimensions
type CMVideoDimensions struct {
	Height int32 // The height of the video.
	Width int32 // The width of the video.

}

