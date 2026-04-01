// Code generated from Apple documentation. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CMAttachmentBearerRef is an object that can carry attachments.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAttachmentBearer
type CMAttachmentBearerRef = corefoundation.CFTypeRef

// CMAttachmentMode is the mode to use when propagating attachments.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAttachmentMode
type CMAttachmentMode = uint32

// CMAudioCodecType is an audio codec type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioCodecType
type CMAudioCodecType = uint32

// CMAudioFormatDescriptionMask is a type for mask bits that represent parts of an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionMask
type CMAudioFormatDescriptionMask = uint32

// CMAudioFormatDescriptionRef is a type you use to interact with audio format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescription
type CMAudioFormatDescriptionRef uintptr

// See: https://developer.apple.com/documentation/CoreMedia/CMBaseClassVersion
type CMBaseClassVersion = uint32

// CMBlockBufferFlags is a type for flags that control behaviors and features of block buffer APIs.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferFlags
type CMBlockBufferFlags = uint32

// CMBlockBufferRef is a reference to a block buffer instance.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBuffer
type CMBlockBufferRef uintptr

// CMBufferCompareCallback is callback that compares one [CMBuffer] with another.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferCompareCallback
type CMBufferCompareCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) corefoundation.CFComparisonResult

// See: https://developer.apple.com/documentation/CoreMedia/CMBufferCompareHandler
type CMBufferCompareHandler = func(unsafe.Pointer, unsafe.Pointer) corefoundation.CFComparisonResult

// CMBufferGetBooleanCallback is callback that returns a Boolean value from a [CMBuffer].
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetBooleanCallback
type CMBufferGetBooleanCallback = func(unsafe.Pointer, unsafe.Pointer) uint8

// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetBooleanHandler
type CMBufferGetBooleanHandler = func(unsafe.Pointer) byte

// CMBufferGetSizeCallback is a client callback that returns a size.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetSizeCallback
type CMBufferGetSizeCallback = func(unsafe.Pointer, unsafe.Pointer) uint

// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetSizeHandler
type CMBufferGetSizeHandler = func(unsafe.Pointer) uint64

// CMBufferGetTimeCallback is callback that returns a [CMTime] from a [CMBuffer].
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetTimeCallback
type CMBufferGetTimeCallback = func(unsafe.Pointer, unsafe.Pointer) CMTime

// See: https://developer.apple.com/documentation/CoreMedia/CMBufferGetTimeHandler
type CMBufferGetTimeHandler = func(unsafe.Pointer) CMTime

// CMBufferQueueRef is a reference to a buffer queue instance.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueue
type CMBufferQueueRef uintptr

// CMBufferQueueTriggerCallback is a callback for the system to invoke when a trigger condition becomes true.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerCallback
type CMBufferQueueTriggerCallback = func(unsafe.Pointer, uintptr)

// CMBufferQueueTriggerCondition is a type to specify conditions to associate with a buffer queue trigger.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerCondition
type CMBufferQueueTriggerCondition = int32

// CMBufferQueueTriggerHandler is a type alias for a trigger handler.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerHandler
type CMBufferQueueTriggerHandler = func(objectivec.IObject)

// CMBufferQueueTriggerToken is a type alias for a trigger token.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerToken
type CMBufferQueueTriggerToken = uintptr

// CMBufferRef is a reference to a buffer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBuffer
type CMBufferRef = corefoundation.CFTypeRef

// CMBufferValidationCallback is a type alias for a callback that tests whether a buffer is in a valid state to add to a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferValidationCallback
type CMBufferValidationCallback = func(uintptr, unsafe.Pointer, unsafe.Pointer) int

// CMBufferValidationHandler is a type alias for a handler that tests whether a buffer is in a valid state to add to a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferValidationHandler
type CMBufferValidationHandler = func(objectivec.IObject, unsafe.Pointer) int

// CMClockOrTimebaseRef is a type you use in argument lists and function results to indicate that you can pass either a clock or timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockOrTimebase
type CMClockOrTimebaseRef = corefoundation.CFTypeRef

// CMClockRef is an object that represents a source of time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClock
type CMClockRef uintptr

// CMClosedCaptionDescriptionFlavor is types that represent closed caption format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionDescriptionFlavor
type CMClosedCaptionDescriptionFlavor = corefoundation.CFStringRef

// CMClosedCaptionFormatDescriptionRef is a type you use to interact with closed caption format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescription
type CMClosedCaptionFormatDescriptionRef uintptr

// CMClosedCaptionFormatType is a closed caption format type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatType
type CMClosedCaptionFormatType = uint32

// CMFormatDescriptionRef is an object that describes a media format descriptor.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
type CMFormatDescriptionRef uintptr

// CMImageDescriptionFlavor is types that represent image format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor
type CMImageDescriptionFlavor = corefoundation.CFStringRef

// CMItemCount is a datatype that represents an item count.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMItemCount
type CMItemCount = int

// CMItemIndex is a datatype that represents an item index.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMItemIndex
type CMItemIndex = int

// CMMediaType is constants that represent media types.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMediaType
type CMMediaType = uint32

// CMMemoryPoolRef is an instance that optimizes memory allocation when working with large blocks of memory.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPool
type CMMemoryPoolRef uintptr

// CMMetadataDescriptionFlavor is types that represent metadata format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDescriptionFlavor
type CMMetadataDescriptionFlavor = corefoundation.CFStringRef

// CMMetadataFormatDescriptionRef is a type you use to interact with metadata format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescription
type CMMetadataFormatDescriptionRef uintptr

// CMMetadataFormatType is a metadata format type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatType
type CMMetadataFormatType = uint32

// CMMutableTagCollectionRef is a mutable reference to a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMutableTagCollectionRef
type CMMutableTagCollectionRef uintptr

// CMMuxedFormatDescriptionRef is a type you use to interact with muxed format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMuxedFormatDescription
type CMMuxedFormatDescriptionRef uintptr

// CMMuxedStreamType is a datatype that represents a muxed stream of data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMuxedStreamType
type CMMuxedStreamType = uint32

// CMPersistentTrackID is a datatype that represents a persistent track identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMPersistentTrackID
type CMPersistentTrackID = int32

// CMPixelFormatType is a pixel format type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMPixelFormatType
type CMPixelFormatType = uint32

// CMSampleBufferInvalidateCallback is client callback called by [CMSampleBufferInvalidate(_:)].
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidateCallback
type CMSampleBufferInvalidateCallback = func(uintptr, uint64)

// CMSampleBufferInvalidateHandler is client callback called by [CMSampleBufferInvalidate(_:)].
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidateHandler
type CMSampleBufferInvalidateHandler = func(objectivec.IObject)

// CMSampleBufferMakeDataReadyCallback is client callback called by [CMSampleBufferMakeDataReady(_:)].
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReadyCallback
type CMSampleBufferMakeDataReadyCallback = func(uintptr, unsafe.Pointer) int

// CMSampleBufferMakeDataReadyHandler is a block the system calls to make the sample buffer ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReadyHandler
type CMSampleBufferMakeDataReadyHandler = func(objectivec.IObject) int

// CMSampleBufferRef is a reference to a buffer of media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
type CMSampleBufferRef uintptr

// CMSimpleQueueRef is a reference to an instance that provides a simple lockless queue of elements.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueue
type CMSimpleQueueRef uintptr

// CMSoundDescriptionFlavor is types that represent sound format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor
type CMSoundDescriptionFlavor = corefoundation.CFStringRef

// See: https://developer.apple.com/documentation/CoreMedia/CMStructVersion
type CMStructVersion = uint32

// CMSubtitleFormatType is a type that represents a text subtitle format.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSubtitleFormatType
type CMSubtitleFormatType = uint32

// CMTagCollectionApplierFunction is a type for function application over elements of a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApplierFunction
type CMTagCollectionApplierFunction = func(CMTag, unsafe.Pointer)

// CMTagCollectionRef is a reference to a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRef
type CMTagCollectionRef uintptr

// CMTagCollectionTagFilterFunction is a type for filtering of tag collections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionTagFilterFunction
type CMTagCollectionTagFilterFunction = func(CMTag, unsafe.Pointer) uint8

// CMTagValue is the type used to represent tag values.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagValue
type CMTagValue = uint64

// CMTaggedBufferGroupFormatDescriptionRef is a type for tagged buffer format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescription
type CMTaggedBufferGroupFormatDescriptionRef uintptr

// CMTaggedBufferGroupFormatType is a type for tagged buffer format information.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatType
type CMTaggedBufferGroupFormatType = uint32

// CMTaggedBufferGroupRef is a reference to a tagged buffer group instance.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupRef
type CMTaggedBufferGroupRef uintptr

// CMTextDescriptionFlavor is types that represent text format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextDescriptionFlavor
type CMTextDescriptionFlavor = corefoundation.CFStringRef

// CMTextDisplayFlags is an integer value that describes the display mode flags for text media.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextDisplayFlags
type CMTextDisplayFlags = uint32

// CMTextFormatDescriptionRef is a type you use to interact with text format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescription
type CMTextFormatDescriptionRef uintptr

// CMTextFormatType is a text format type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatType
type CMTextFormatType = uint32

// CMTextJustificationValue is an integer value that describes the justification modes for text media.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextJustificationValue
type CMTextJustificationValue = int8

// CMTimeCodeDescriptionFlavor is types that represent time code format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeDescriptionFlavor
type CMTimeCodeDescriptionFlavor = corefoundation.CFStringRef

// CMTimeCodeFormatDescriptionRef is a type you use to interact with time code format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescription
type CMTimeCodeFormatDescriptionRef uintptr

// CMTimeCodeFormatType is a time code format type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatType
type CMTimeCodeFormatType = uint32

// CMTimeEpoch is an epoch for a time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeEpoch
type CMTimeEpoch = int64

// CMTimeScale is an integer timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeScale
type CMTimeScale = int32

// CMTimeValue is an integer time value.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeValue
type CMTimeValue = int64

// CMTimebaseRef is a model of a timeline under application control.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebase
type CMTimebaseRef uintptr

// CMVideoCodecType is a video codec type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoCodecType
type CMVideoCodecType = uint32

// CMVideoFormatDescriptionRef is a type you use to interact with video format descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescription
type CMVideoFormatDescriptionRef uintptr
