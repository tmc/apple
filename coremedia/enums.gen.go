// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreMedia/CMPackingType
type CMPackingType uint64

const (
	// KCMPackingType_None: Each frame contains only a single image, and isn’t frame-packed.
	KCMPackingType_None CMPackingType = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'none'
	// KCMPackingType_OverUnder: The video contains packed frames that have a left eye image on the top and right eye image on the bottom.
	KCMPackingType_OverUnder CMPackingType = 'o'<<24 | 'v'<<16 | 'e'<<8 | 'r' // 'over'
	// KCMPackingType_SideBySide: The video contains packed frames that have a left eye image on the left and right eye image on the right.
	KCMPackingType_SideBySide CMPackingType = 's'<<24 | 'i'<<16 | 'd'<<8 | 'e' // 'side'
)

func (e CMPackingType) String() string {
	switch e {
	case KCMPackingType_None:
		return "KCMPackingType_None"
	case KCMPackingType_OverUnder:
		return "KCMPackingType_OverUnder"
	case KCMPackingType_SideBySide:
		return "KCMPackingType_SideBySide"
	default:
		return fmt.Sprintf("CMPackingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMProjectionType
type CMProjectionType uint64

const (
	// KCMProjectionType_Equirectangular: Video content displays as a 360 degree equirectangular projection.
	KCMProjectionType_Equirectangular CMProjectionType = 'e'<<24 | 'q'<<16 | 'u'<<8 | 'i' // 'equi'
	// KCMProjectionType_Fisheye: Video content displays as a fisheye projection.
	KCMProjectionType_Fisheye CMProjectionType = 'f'<<24 | 'i'<<16 | 's'<<8 | 'h' // 'fish'
	// KCMProjectionType_HalfEquirectangular: Video content displays as a 180 degree equirectangular projection.
	KCMProjectionType_HalfEquirectangular CMProjectionType = 'h'<<24 | 'e'<<16 | 'q'<<8 | 'u' // 'hequ'
	KCMProjectionType_ParametricImmersive CMProjectionType = 'p'<<24 | 'r'<<16 | 'i'<<8 | 'm' // 'prim'
	// KCMProjectionType_Rectangular: Video content displays on a flat, rectangular 2D surface.
	KCMProjectionType_Rectangular CMProjectionType = 'r'<<24 | 'e'<<16 | 'c'<<8 | 't' // 'rect'
)

func (e CMProjectionType) String() string {
	switch e {
	case KCMProjectionType_Equirectangular:
		return "KCMProjectionType_Equirectangular"
	case KCMProjectionType_Fisheye:
		return "KCMProjectionType_Fisheye"
	case KCMProjectionType_HalfEquirectangular:
		return "KCMProjectionType_HalfEquirectangular"
	case KCMProjectionType_ParametricImmersive:
		return "KCMProjectionType_ParametricImmersive"
	case KCMProjectionType_Rectangular:
		return "KCMProjectionType_Rectangular"
	default:
		return fmt.Sprintf("CMProjectionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents
type CMStereoViewComponents uint64

const (
	// KCMStereoView_LeftEye: The stereo video track includes a left eye layer.
	KCMStereoView_LeftEye CMStereoViewComponents = 1
	// KCMStereoView_None: A constant for video metadata to have no available stereo frames.
	KCMStereoView_None CMStereoViewComponents = 0
	// KCMStereoView_RightEye: The stereo video track includes a right eye layer.
	KCMStereoView_RightEye CMStereoViewComponents = 2
)

func (e CMStereoViewComponents) String() string {
	switch e {
	case KCMStereoView_LeftEye:
		return "KCMStereoView_LeftEye"
	case KCMStereoView_None:
		return "KCMStereoView_None"
	case KCMStereoView_RightEye:
		return "KCMStereoView_RightEye"
	default:
		return fmt.Sprintf("CMStereoViewComponents(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions
type CMStereoViewInterpretationOptions uint64

const (
	// KCMStereoViewInterpretation_AdditionalViews: A flag indicating that the video content contains additional views beyond the left or right eye.
	KCMStereoViewInterpretation_AdditionalViews CMStereoViewInterpretationOptions = 2
	// KCMStereoViewInterpretation_Default: The default options for stereo video views.
	KCMStereoViewInterpretation_Default CMStereoViewInterpretationOptions = 0
	// KCMStereoViewInterpretation_StereoOrderReversed: Changes the default ordering of eye data, switching it from left-to-right to right-to-left.
	KCMStereoViewInterpretation_StereoOrderReversed CMStereoViewInterpretationOptions = 1
)

func (e CMStereoViewInterpretationOptions) String() string {
	switch e {
	case KCMStereoViewInterpretation_AdditionalViews:
		return "KCMStereoViewInterpretation_AdditionalViews"
	case KCMStereoViewInterpretation_Default:
		return "KCMStereoViewInterpretation_Default"
	case KCMStereoViewInterpretation_StereoOrderReversed:
		return "KCMStereoViewInterpretation_StereoOrderReversed"
	default:
		return fmt.Sprintf("CMStereoViewInterpretationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTagCategory
type CMTagCategory int

const (
	// KCMTagCategory_ChannelID: A category used for tagging a channel ID.
	KCMTagCategory_ChannelID CMTagCategory = 'v'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'vchn'
	// KCMTagCategory_MediaSubType: A category used for tagging media subtype metadata.
	KCMTagCategory_MediaSubType CMTagCategory = 'm'<<24 | 's'<<16 | 'u'<<8 | 'b' // 'msub'
	// KCMTagCategory_MediaType: A category used for tagging media type metadata.
	KCMTagCategory_MediaType CMTagCategory = 'm'<<24 | 'd'<<16 | 'i'<<8 | 'a' // 'mdia'
	// KCMTagCategory_PackingType: A category used for tagging frame-packing information.
	KCMTagCategory_PackingType CMTagCategory = 'p'<<24 | 'a'<<16 | 'c'<<8 | 'k' // 'pack'
	// KCMTagCategory_PixelFormat: A category used for tagging pixel format information.
	KCMTagCategory_PixelFormat CMTagCategory = 'p'<<24 | 'i'<<16 | 'x'<<8 | 'f' // 'pixf'
	// KCMTagCategory_ProjectionType: A category used for tagging projection surface information.
	KCMTagCategory_ProjectionType CMTagCategory = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'j' // 'proj'
	// KCMTagCategory_StereoView: A category used for tagging eye information for 3D video.
	KCMTagCategory_StereoView CMTagCategory = 'e'<<24 | 'y'<<16 | 'e'<<8 | 's' // 'eyes'
	// KCMTagCategory_StereoViewInterpretation: A category used for tagging how to interpret stereo view metadata.
	KCMTagCategory_StereoViewInterpretation CMTagCategory = 'e'<<24 | 'y'<<16 | 'i'<<8 | 'p' // 'eyip'
	// KCMTagCategory_TrackID: A category used for tagging a track ID.
	KCMTagCategory_TrackID CMTagCategory = 't'<<24 | 'r'<<16 | 'a'<<8 | 'k' // 'trak'
	// KCMTagCategory_Undefined: An unknown or undefined tag category.
	KCMTagCategory_Undefined CMTagCategory = 0
	// KCMTagCategory_VideoLayerID: A category used for tagging a video layer ID.
	KCMTagCategory_VideoLayerID CMTagCategory = 'v'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'vlay'
)

func (e CMTagCategory) String() string {
	switch e {
	case KCMTagCategory_ChannelID:
		return "KCMTagCategory_ChannelID"
	case KCMTagCategory_MediaSubType:
		return "KCMTagCategory_MediaSubType"
	case KCMTagCategory_MediaType:
		return "KCMTagCategory_MediaType"
	case KCMTagCategory_PackingType:
		return "KCMTagCategory_PackingType"
	case KCMTagCategory_PixelFormat:
		return "KCMTagCategory_PixelFormat"
	case KCMTagCategory_ProjectionType:
		return "KCMTagCategory_ProjectionType"
	case KCMTagCategory_StereoView:
		return "KCMTagCategory_StereoView"
	case KCMTagCategory_StereoViewInterpretation:
		return "KCMTagCategory_StereoViewInterpretation"
	case KCMTagCategory_TrackID:
		return "KCMTagCategory_TrackID"
	case KCMTagCategory_Undefined:
		return "KCMTagCategory_Undefined"
	case KCMTagCategory_VideoLayerID:
		return "KCMTagCategory_VideoLayerID"
	default:
		return fmt.Sprintf("CMTagCategory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError
type CMTagCollectionError int

const (
	// KCMTagCollectionError_AllocationFailed: Indicates an internal allocation failed.
	KCMTagCollectionError_AllocationFailed CMTagCollectionError = -15741
	// KCMTagCollectionError_ExhaustedBufferSize: Indicates that a buffer was smaller than the number of requested tags.
	KCMTagCollectionError_ExhaustedBufferSize CMTagCollectionError = -15748
	// KCMTagCollectionError_InternalError: Indicates an error occurred inside of the Core Media framework.
	KCMTagCollectionError_InternalError CMTagCollectionError = -15742
	// KCMTagCollectionError_InvalidTag: Indicates that the collection contains an invalid tag.
	KCMTagCollectionError_InvalidTag CMTagCollectionError = -15743
	// KCMTagCollectionError_InvalidTagCollectionData: Indicates that a Core Foundation data instance failed to initialize a new tag collection.
	KCMTagCollectionError_InvalidTagCollectionData CMTagCollectionError = -15745
	// KCMTagCollectionError_InvalidTagCollectionDataVersion: Indicates that a Core Foundation data instance failed to initialize a new tag collection due to a versioning problem.
	KCMTagCollectionError_InvalidTagCollectionDataVersion CMTagCollectionError = -15747
	// KCMTagCollectionError_InvalidTagCollectionDictionary: Indicates that a Core Foundation dictionary instance failed to initialize a new tag collection.
	KCMTagCollectionError_InvalidTagCollectionDictionary CMTagCollectionError = -15744
	// KCMTagCollectionError_NotYetImplemented: Indicates a function lacks a necessary backing implementation in Core Media.
	KCMTagCollectionError_NotYetImplemented CMTagCollectionError = -15749
	// KCMTagCollectionError_ParamErr: Indicates a parameter to a function was of the wrong type or didn’t meet a necessary condition.
	KCMTagCollectionError_ParamErr CMTagCollectionError = -15740
	// KCMTagCollectionError_TagNotFound: Indicates that there was no match in a collection for a tag.
	KCMTagCollectionError_TagNotFound CMTagCollectionError = -15746
)

func (e CMTagCollectionError) String() string {
	switch e {
	case KCMTagCollectionError_AllocationFailed:
		return "KCMTagCollectionError_AllocationFailed"
	case KCMTagCollectionError_ExhaustedBufferSize:
		return "KCMTagCollectionError_ExhaustedBufferSize"
	case KCMTagCollectionError_InternalError:
		return "KCMTagCollectionError_InternalError"
	case KCMTagCollectionError_InvalidTag:
		return "KCMTagCollectionError_InvalidTag"
	case KCMTagCollectionError_InvalidTagCollectionData:
		return "KCMTagCollectionError_InvalidTagCollectionData"
	case KCMTagCollectionError_InvalidTagCollectionDataVersion:
		return "KCMTagCollectionError_InvalidTagCollectionDataVersion"
	case KCMTagCollectionError_InvalidTagCollectionDictionary:
		return "KCMTagCollectionError_InvalidTagCollectionDictionary"
	case KCMTagCollectionError_NotYetImplemented:
		return "KCMTagCollectionError_NotYetImplemented"
	case KCMTagCollectionError_ParamErr:
		return "KCMTagCollectionError_ParamErr"
	case KCMTagCollectionError_TagNotFound:
		return "KCMTagCollectionError_TagNotFound"
	default:
		return fmt.Sprintf("CMTagCollectionError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTagDataType
type CMTagDataType uint32

const (
	// KCMTagDataType_Flags: The tag value is a 64-bit wide bitflag field.
	KCMTagDataType_Flags CMTagDataType = 7
	// KCMTagDataType_Float64: The tag value is a 64-bit floating point number.
	KCMTagDataType_Float64 CMTagDataType = 3
	// KCMTagDataType_Invalid: The tag value isn’t associated with any known data type.
	KCMTagDataType_Invalid CMTagDataType = 0
	// KCMTagDataType_OSType: The tag value is a 64-bit identifier used by the operating system.
	KCMTagDataType_OSType CMTagDataType = 5
	// KCMTagDataType_SInt64: The tag value is a signed 64-bit integer.
	KCMTagDataType_SInt64 CMTagDataType = 2
)

func (e CMTagDataType) String() string {
	switch e {
	case KCMTagDataType_Flags:
		return "KCMTagDataType_Flags"
	case KCMTagDataType_Float64:
		return "KCMTagDataType_Float64"
	case KCMTagDataType_Invalid:
		return "KCMTagDataType_Invalid"
	case KCMTagDataType_OSType:
		return "KCMTagDataType_OSType"
	case KCMTagDataType_SInt64:
		return "KCMTagDataType_SInt64"
	default:
		return fmt.Sprintf("CMTagDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTagError
type CMTagError int

const (
	// KCMTagError_AllocationFailed: An error where the system can’t allocate enough memory for the tag.
	KCMTagError_AllocationFailed CMTagError = -15731
	// KCMTagError_ParamErr: An error where input or output parameters didn’t match the requirements of Core Media.
	KCMTagError_ParamErr CMTagError = -15730
)

func (e CMTagError) String() string {
	switch e {
	case KCMTagError_AllocationFailed:
		return "KCMTagError_AllocationFailed"
	case KCMTagError_ParamErr:
		return "KCMTagError_ParamErr"
	default:
		return fmt.Sprintf("CMTagError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupError
type CMTaggedBufferGroupError int

const (
	// KCMTaggedBufferGroupError_AllocationFailed: Indicates an internal allocation failed.
	KCMTaggedBufferGroupError_AllocationFailed CMTaggedBufferGroupError = -15781
	// KCMTaggedBufferGroupError_InternalError: Indicates an error occurred inside of the Core Media framework.
	KCMTaggedBufferGroupError_InternalError CMTaggedBufferGroupError = -15782
	// KCMTaggedBufferGroupError_ParamErr: Indicates a parameter to a function was of the wrong type or didn’t meet a necessary condition.
	KCMTaggedBufferGroupError_ParamErr CMTaggedBufferGroupError = -15780
)

func (e CMTaggedBufferGroupError) String() string {
	switch e {
	case KCMTaggedBufferGroupError_AllocationFailed:
		return "KCMTaggedBufferGroupError_AllocationFailed"
	case KCMTaggedBufferGroupError_InternalError:
		return "KCMTaggedBufferGroupError_InternalError"
	case KCMTaggedBufferGroupError_ParamErr:
		return "KCMTaggedBufferGroupError_ParamErr"
	default:
		return fmt.Sprintf("CMTaggedBufferGroupError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags
type CMTimeFlags uint32

const (
	// KCMTimeFlags_HasBeenRounded: A flag that indicates a previous time calculation rounded the result.
	KCMTimeFlags_HasBeenRounded CMTimeFlags = 2
	// KCMTimeFlags_ImpliedValueFlagsMask: A flag that indicates the time is positive or negative infinity, or indefinite.
	KCMTimeFlags_ImpliedValueFlagsMask CMTimeFlags = 4
	// KCMTimeFlags_Indefinite: A flag that indicates the time is indefinite.
	KCMTimeFlags_Indefinite CMTimeFlags = 16
	// KCMTimeFlags_NegativeInfinity: A flag that indicates the time is negative infinity.
	KCMTimeFlags_NegativeInfinity CMTimeFlags = 8
	// KCMTimeFlags_PositiveInfinity: A flag that indicates the time is positive infinity.
	KCMTimeFlags_PositiveInfinity CMTimeFlags = 4
	// KCMTimeFlags_Valid: A flag that indicates a time is valid.
	KCMTimeFlags_Valid CMTimeFlags = 1
)

func (e CMTimeFlags) String() string {
	switch e {
	case KCMTimeFlags_HasBeenRounded:
		return "KCMTimeFlags_HasBeenRounded"
	case KCMTimeFlags_ImpliedValueFlagsMask:
		return "KCMTimeFlags_ImpliedValueFlagsMask"
	case KCMTimeFlags_Indefinite:
		return "KCMTimeFlags_Indefinite"
	case KCMTimeFlags_NegativeInfinity:
		return "KCMTimeFlags_NegativeInfinity"
	case KCMTimeFlags_Valid:
		return "KCMTimeFlags_Valid"
	default:
		return fmt.Sprintf("CMTimeFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod
type CMTimeRoundingMethod uint32

const (
	// KCMTimeRoundingMethod_Default: The default rounding method.
	KCMTimeRoundingMethod_Default CMTimeRoundingMethod = 1
	// KCMTimeRoundingMethod_QuickTime: Rounds using the QuickTime method.
	KCMTimeRoundingMethod_QuickTime CMTimeRoundingMethod = 4
	// KCMTimeRoundingMethod_RoundAwayFromZero: Rounds away from zero.
	KCMTimeRoundingMethod_RoundAwayFromZero CMTimeRoundingMethod = 3
	// KCMTimeRoundingMethod_RoundHalfAwayFromZero: Rounds half away from zero.
	KCMTimeRoundingMethod_RoundHalfAwayFromZero CMTimeRoundingMethod = 1
	// KCMTimeRoundingMethod_RoundTowardNegativeInfinity: Rounds toward negative infinity.
	KCMTimeRoundingMethod_RoundTowardNegativeInfinity CMTimeRoundingMethod = 6
	// KCMTimeRoundingMethod_RoundTowardPositiveInfinity: Rounds toward positive infinity.
	KCMTimeRoundingMethod_RoundTowardPositiveInfinity CMTimeRoundingMethod = 5
	// KCMTimeRoundingMethod_RoundTowardZero: Rounds toward zero.
	KCMTimeRoundingMethod_RoundTowardZero CMTimeRoundingMethod = 2
)

func (e CMTimeRoundingMethod) String() string {
	switch e {
	case KCMTimeRoundingMethod_Default:
		return "KCMTimeRoundingMethod_Default"
	case KCMTimeRoundingMethod_QuickTime:
		return "KCMTimeRoundingMethod_QuickTime"
	case KCMTimeRoundingMethod_RoundAwayFromZero:
		return "KCMTimeRoundingMethod_RoundAwayFromZero"
	case KCMTimeRoundingMethod_RoundTowardNegativeInfinity:
		return "KCMTimeRoundingMethod_RoundTowardNegativeInfinity"
	case KCMTimeRoundingMethod_RoundTowardPositiveInfinity:
		return "KCMTimeRoundingMethod_RoundTowardPositiveInfinity"
	case KCMTimeRoundingMethod_RoundTowardZero:
		return "KCMTimeRoundingMethod_RoundTowardZero"
	default:
		return fmt.Sprintf("CMTimeRoundingMethod(%d)", e)
	}
}

type KCMAudioFormatDescriptionMask uint

const (
	// KCMAudioFormatDescriptionMask_All: A mask that represents all parts of an audio format description.
	KCMAudioFormatDescriptionMask_All KCMAudioFormatDescriptionMask = 0
	// KCMAudioFormatDescriptionMask_ChannelLayout: A mask that represents the audio channel layout.
	KCMAudioFormatDescriptionMask_ChannelLayout KCMAudioFormatDescriptionMask = 0
	// KCMAudioFormatDescriptionMask_Extensions: A mask that represents the format description extensions.
	KCMAudioFormatDescriptionMask_Extensions KCMAudioFormatDescriptionMask = 0
	// KCMAudioFormatDescriptionMask_MagicCookie: A mask that represents the magic cookie.
	KCMAudioFormatDescriptionMask_MagicCookie KCMAudioFormatDescriptionMask = 0
	// KCMAudioFormatDescriptionMask_StreamBasicDescription: A mask that represents the audio stream description.
	KCMAudioFormatDescriptionMask_StreamBasicDescription KCMAudioFormatDescriptionMask = 0
)

func (e KCMAudioFormatDescriptionMask) String() string {
	switch e {
	case KCMAudioFormatDescriptionMask_All:
		return "KCMAudioFormatDescriptionMask_All"
	default:
		return fmt.Sprintf("KCMAudioFormatDescriptionMask(%d)", e)
	}
}

type KCMBlockBuffer uint

const (
	// KCMBlockBufferAlwaysCopyDataFlag: Used with CMBlockBuffer to cause it to always produce an allocated copy of the desired data.
	KCMBlockBufferAlwaysCopyDataFlag KCMBlockBuffer = 0
	// KCMBlockBufferAssureMemoryNowFlag: When passed to routines that accept block allocators, causes the memory block to be allocated immediately.
	KCMBlockBufferAssureMemoryNowFlag KCMBlockBuffer = 0
	// KCMBlockBufferBadCustomBlockSourceErr: An error code that indicates the custom block source is invalid.
	KCMBlockBufferBadCustomBlockSourceErr KCMBlockBuffer = 0
	// KCMBlockBufferBadLengthParameterErr: An error code that indicates the block length is zero or doesn’t equal the size of the memory block.
	KCMBlockBufferBadLengthParameterErr KCMBlockBuffer = 0
	// KCMBlockBufferBadOffsetParameterErr: An error code that indicates the offset doesn’t point to the location of data in the memory block.
	KCMBlockBufferBadOffsetParameterErr KCMBlockBuffer = 0
	// KCMBlockBufferBadPointerParameterErr: An error code that indicates the block buffer reference is invalid.
	KCMBlockBufferBadPointerParameterErr KCMBlockBuffer = 0
	// KCMBlockBufferBlockAllocationFailedErr: An error code that indicates the block allocator failed to allocate a memory block.
	KCMBlockBufferBlockAllocationFailedErr KCMBlockBuffer = 0
	// KCMBlockBufferDontOptimizeDepthFlag: Passed to block buffers to suppress reference depth optimization.
	KCMBlockBufferDontOptimizeDepthFlag KCMBlockBuffer = 0
	// KCMBlockBufferEmptyBBufErr: An error code that indicates the block buffer is empty.
	KCMBlockBufferEmptyBBufErr KCMBlockBuffer = 0
	// KCMBlockBufferInsufficientSpaceErr: An error code that indicates the system failed to create a new buffer because of insufficient space at the buffer out location.
	KCMBlockBufferInsufficientSpaceErr KCMBlockBuffer = 0
	// KCMBlockBufferNoErr: A code that indicates the system successfully created a block buffer with no errors.
	KCMBlockBufferNoErr KCMBlockBuffer = 0
	// KCMBlockBufferPermitEmptyReferenceFlag: Passed to CMBlockBuffer and CMBlockBuffer to allow references into a [CMBlockBuffer] that may not yet be populated.
	KCMBlockBufferPermitEmptyReferenceFlag KCMBlockBuffer = 0
	// KCMBlockBufferStructureAllocationFailedErr: An error code that indicates the structure allocator failed to allocate a block buffer.
	KCMBlockBufferStructureAllocationFailedErr KCMBlockBuffer = 0
	// KCMBlockBufferUnallocatedBlockErr: An error code that indicates the system encountered an unallocated memory block.
	KCMBlockBufferUnallocatedBlockErr KCMBlockBuffer = 0
)

func (e KCMBlockBuffer) String() string {
	switch e {
	case KCMBlockBufferAlwaysCopyDataFlag:
		return "KCMBlockBufferAlwaysCopyDataFlag"
	default:
		return fmt.Sprintf("KCMBlockBuffer(%d)", e)
	}
}

type KCMBlockBufferCustomBlockSource uint

const (
	// KCMBlockBufferCustomBlockSourceVersion: The value is the block source version.
	KCMBlockBufferCustomBlockSourceVersion KCMBlockBufferCustomBlockSource = 0
)

func (e KCMBlockBufferCustomBlockSource) String() string {
	switch e {
	case KCMBlockBufferCustomBlockSourceVersion:
		return "KCMBlockBufferCustomBlockSourceVersion"
	default:
		return fmt.Sprintf("KCMBlockBufferCustomBlockSource(%d)", e)
	}
}

type KCMBufferQueueError int

const (
	// KCMBufferQueueError_AllocationFailed: The system failed to allocate memory.
	KCMBufferQueueError_AllocationFailed KCMBufferQueueError = 0
	// KCMBufferQueueError_BadTriggerDuration: You specified an invalid trigger duration.
	KCMBufferQueueError_BadTriggerDuration KCMBufferQueueError = 0
	// KCMBufferQueueError_CannotModifyQueueFromTriggerCallback: A trigger callback attempted to modify a queue.
	KCMBufferQueueError_CannotModifyQueueFromTriggerCallback KCMBufferQueueError = 0
	// KCMBufferQueueError_EnqueueAfterEndOfData: You attempted to enqueue a buffer on a queue that disallows it.
	KCMBufferQueueError_EnqueueAfterEndOfData KCMBufferQueueError = 0
	// KCMBufferQueueError_InvalidBuffer: A buffer validation callback rejected the buffer.
	KCMBufferQueueError_InvalidBuffer KCMBufferQueueError = 0
	// KCMBufferQueueError_InvalidCMBufferCallbacksStruct: The format of a callbacks structure isn’t correct.
	KCMBufferQueueError_InvalidCMBufferCallbacksStruct KCMBufferQueueError = 0
	// KCMBufferQueueError_InvalidTriggerCondition: You specified an invalid trigger condition.
	KCMBufferQueueError_InvalidTriggerCondition KCMBufferQueueError = 0
	// KCMBufferQueueError_InvalidTriggerToken: You specified a trigger token that isn’t a trigger currently associated with this queue.
	KCMBufferQueueError_InvalidTriggerToken KCMBufferQueueError = 0
	// KCMBufferQueueError_QueueIsFull: You attempted to enqueue a buffer on a queue that’s full.
	KCMBufferQueueError_QueueIsFull KCMBufferQueueError = 0
	// KCMBufferQueueError_RequiredParameterMissing: You failed to provide a valid value for a required parameter.
	KCMBufferQueueError_RequiredParameterMissing KCMBufferQueueError = 0
)

func (e KCMBufferQueueError) String() string {
	switch e {
	case KCMBufferQueueError_AllocationFailed:
		return "KCMBufferQueueError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMBufferQueueError(%d)", e)
	}
}

type KCMClockError int

const (
	// KCMClockError_AllocationFailed: A clock error that indicates the memory allocation fails.
	KCMClockError_AllocationFailed KCMClockError = 0
	// KCMClockError_InvalidParameter: A clock error that indicates a parameter isn’t valid.
	KCMClockError_InvalidParameter KCMClockError = 0
	// KCMClockError_MissingRequiredParameter: A clock error that indicates a parameter is missing.
	KCMClockError_MissingRequiredParameter KCMClockError = 0
	// KCMClockError_UnsupportedOperation: A clock error that indicates the operation isn’t supported.
	KCMClockError_UnsupportedOperation KCMClockError = 0
)

func (e KCMClockError) String() string {
	switch e {
	case KCMClockError_AllocationFailed:
		return "KCMClockError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMClockError(%d)", e)
	}
}

type KCMClosedCaptionFormatType uint

const (
	// KCMClosedCaptionFormatType_ATSC: A type that describes ATSC-compliant samples.
	KCMClosedCaptionFormatType_ATSC KCMClosedCaptionFormatType = 0
	// KCMClosedCaptionFormatType_CEA608: A type that describes CEA 608-compliant samples.
	KCMClosedCaptionFormatType_CEA608 KCMClosedCaptionFormatType = 0
	// KCMClosedCaptionFormatType_CEA708: A type that describes CEA 708-compliant samples.
	KCMClosedCaptionFormatType_CEA708 KCMClosedCaptionFormatType = 0
)

func (e KCMClosedCaptionFormatType) String() string {
	switch e {
	case KCMClosedCaptionFormatType_ATSC:
		return "KCMClosedCaptionFormatType_ATSC"
	default:
		return fmt.Sprintf("KCMClosedCaptionFormatType(%d)", e)
	}
}

type KCMFormatDescriptionBridgeError int

const (
	// KCMFormatDescriptionBridgeError_AllocationFailed: A bridge error that indicates when an allocation fails.
	KCMFormatDescriptionBridgeError_AllocationFailed KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_IncompatibleFormatDescription: A bridge error that indicates the format description has an unknown format.
	KCMFormatDescriptionBridgeError_IncompatibleFormatDescription KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_InvalidFormatDescription: A bridge error that indicates the format description isn’t valid.
	KCMFormatDescriptionBridgeError_InvalidFormatDescription KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_InvalidParameter: A bridge error that indicates that the function recieves an empty value for a parameter it requires.
	KCMFormatDescriptionBridgeError_InvalidParameter KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_InvalidSerializedSampleDescription: A bridge error that indicates that the sample isn’t valid.
	KCMFormatDescriptionBridgeError_InvalidSerializedSampleDescription KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_InvalidSlice: A bridge error that indicates the slice isn’t valid.
	KCMFormatDescriptionBridgeError_InvalidSlice KCMFormatDescriptionBridgeError = 0
	// KCMFormatDescriptionBridgeError_UnsupportedSampleDescriptionFlavor: A bridge error that indicates the sample isn’t supported for the format flavor you specify.
	KCMFormatDescriptionBridgeError_UnsupportedSampleDescriptionFlavor KCMFormatDescriptionBridgeError = 0
)

func (e KCMFormatDescriptionBridgeError) String() string {
	switch e {
	case KCMFormatDescriptionBridgeError_AllocationFailed:
		return "KCMFormatDescriptionBridgeError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMFormatDescriptionBridgeError(%d)", e)
	}
}

type KCMFormatDescriptionError int

const (
	// KCMFormatDescriptionError_AllocationFailed: An error that indicates when an allocation fails.
	KCMFormatDescriptionError_AllocationFailed KCMFormatDescriptionError = 0
	// KCMFormatDescriptionError_InvalidParameter: An error that indicates that the function recieves an empty value for a parameter it requires.
	KCMFormatDescriptionError_InvalidParameter KCMFormatDescriptionError = 0
	// KCMFormatDescriptionError_ValueNotAvailable: An error that indicates the format description doesn’t contain the value you request.
	KCMFormatDescriptionError_ValueNotAvailable KCMFormatDescriptionError = 0
)

func (e KCMFormatDescriptionError) String() string {
	switch e {
	case KCMFormatDescriptionError_AllocationFailed:
		return "KCMFormatDescriptionError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMFormatDescriptionError(%d)", e)
	}
}

type KCMMPEG2VideoProfile uint

const (
	// KCMMPEG2VideoProfile_HDV_1080i50: A profile that represents the Apple Intermediate Codec HDV 1080i50 format.
	KCMMPEG2VideoProfile_HDV_1080i50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_1080i60: A profile that represents the Apple Intermediate Codec HDV 1080i60 format.
	KCMMPEG2VideoProfile_HDV_1080i60 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_1080p24: A profile that represents the Apple ProRes 422 codec HDV 1080p24 format.
	KCMMPEG2VideoProfile_HDV_1080p24 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_1080p25: A profile that represents the HDV 1080p25 format.
	KCMMPEG2VideoProfile_HDV_1080p25 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_1080p30: A profile that represents the HDV 1080p30 format.
	KCMMPEG2VideoProfile_HDV_1080p30 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_720p24: A profile that represents the HDV 720p24 format.
	KCMMPEG2VideoProfile_HDV_720p24 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_720p25: A profile that represents the HDV_720p25 format.
	KCMMPEG2VideoProfile_HDV_720p25 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_720p30: A profile that represents the Apple Intermediate Codec HDV 720p30 format.
	KCMMPEG2VideoProfile_HDV_720p30 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_720p50: A profile that represents the HDV 720p50 format.
	KCMMPEG2VideoProfile_HDV_720p50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_HDV_720p60: A profile that represents the HDV 720p60 format.
	KCMMPEG2VideoProfile_HDV_720p60 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_1080i50_VBR35: A profile that represents the XDCAM EX 1080i50 HQ video format with HQ 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_1080i50_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_1080i60_VBR35: A profile that represents the XDCAM EX 1080i60 HQ video format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_1080i60_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_1080p24_VBR35: A profile that represents the XDCAM EX 1080p24 HQ video format with HQ 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_1080p24_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_1080p25_VBR35: A profile that represents the XDCAM EX 1080p25 HQ video format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_1080p25_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_1080p30_VBR35: A profile that represents the XDCAM EX 1080p30 HQ video format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_1080p30_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_720p24_VBR35: A profile that represents the XDCAM EX 720p24 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_720p24_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_720p25_VBR35: A profile that represents the XDCAM EX 720p25 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_720p25_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_720p30_VBR35: A profile that represents the XDCAM EX 720p30 video HQ format with 35 Mbps bit rate .
	KCMMPEG2VideoProfile_XDCAM_EX_720p30_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_720p50_VBR35: A profile that represents the XDCAM EX 720p50 HQ video format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_720p50_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_EX_720p60_VBR35: A profile that represents the XDCAM EX 720p60 HQ video format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_EX_720p60_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_1080i50_CBR50: A profile that represents the XDCAM HD422 1080i50 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_1080i50_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_1080i60_CBR50: A profile that represents the XDCAM HD422 1080i60 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_1080i60_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_1080p24_CBR50: A profile that represents the XDCAM HD422 1080p24 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_1080p24_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_1080p25_CBR50: A profile that represents the XDCAM HD422 1080p25 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_1080p25_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_1080p30_CBR50: A profile that represents the XDCAM HD422 1080p30 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_1080p30_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_540p: A profile that represents the XDCAM HD422 540 video format.
	KCMMPEG2VideoProfile_XDCAM_HD422_540p KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_720p24_CBR50: A profile that represents the XDCAM HD 422 720p24 video format.
	KCMMPEG2VideoProfile_XDCAM_HD422_720p24_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_720p25_CBR50: A profile that represents the XDCAM HD 422 720p25 video format.
	KCMMPEG2VideoProfile_XDCAM_HD422_720p25_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_720p30_CBR50: A profile that represents the XDCAM HD 422 720p30 video format.
	KCMMPEG2VideoProfile_XDCAM_HD422_720p30_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_720p50_CBR50: A profile that represents the XDCAM HD422 720p50 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_720p50_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD422_720p60_CBR50: A profile that represents the XDCAM HD422 720p60 video format with 50 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD422_720p60_CBR50 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_1080i50_VBR35: A profile that represents the XDCAM HD 1080i50 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD_1080i50_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_1080i60_VBR35: A profile that represents the XDCAM HD 1080i60 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD_1080i60_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_1080p24_VBR35: A profile that represents the XDCAM HD 1080p24 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD_1080p24_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_1080p25_VBR35: A profile that represents the XDCAM HD 1080p25 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD_1080p25_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_1080p30_VBR35: A profile that represents the DCAM HD 1080p30 video HQ format with 35 Mbps bit rate.
	KCMMPEG2VideoProfile_XDCAM_HD_1080p30_VBR35 KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XDCAM_HD_540p: A profile that represents the XDCAM HD 540p video format.
	KCMMPEG2VideoProfile_XDCAM_HD_540p KCMMPEG2VideoProfile = 0
	// KCMMPEG2VideoProfile_XF: A profile that represents the XF video format.
	KCMMPEG2VideoProfile_XF KCMMPEG2VideoProfile = 0
)

func (e KCMMPEG2VideoProfile) String() string {
	switch e {
	case KCMMPEG2VideoProfile_HDV_1080i50:
		return "KCMMPEG2VideoProfile_HDV_1080i50"
	default:
		return fmt.Sprintf("KCMMPEG2VideoProfile(%d)", e)
	}
}

type KCMMediaType uint

const (
	// KCMMediaType_Audio: Audio media.
	KCMMediaType_Audio            KCMMediaType = 0
	KCMMediaType_AuxiliaryPicture KCMMediaType = 0
	// KCMMediaType_ClosedCaption: Closed-caption media.
	KCMMediaType_ClosedCaption KCMMediaType = 0
	// KCMMediaType_Metadata: Meta data.
	KCMMediaType_Metadata KCMMediaType = 0
	// KCMMediaType_Muxed: Muxed media.
	KCMMediaType_Muxed KCMMediaType = 0
	// KCMMediaType_Subtitle: Subtitle media.
	KCMMediaType_Subtitle KCMMediaType = 0
	// KCMMediaType_TaggedBufferGroup: Media containing a tagged buffer group.
	KCMMediaType_TaggedBufferGroup KCMMediaType = 0
	// KCMMediaType_Text: Text media.
	KCMMediaType_Text KCMMediaType = 0
	// KCMMediaType_TimeCode: Time code media.
	KCMMediaType_TimeCode KCMMediaType = 0
	// KCMMediaType_Video: Video media.
	KCMMediaType_Video KCMMediaType = 0
)

func (e KCMMediaType) String() string {
	switch e {
	case KCMMediaType_Audio:
		return "KCMMediaType_Audio"
	default:
		return fmt.Sprintf("KCMMediaType(%d)", e)
	}
}

type KCMMemoryPoolError int

const (
	// KCMMemoryPoolError_AllocationFailed: An error that indicates the system failed to allocate an internal data structure.
	KCMMemoryPoolError_AllocationFailed KCMMemoryPoolError = 0
	// KCMMemoryPoolError_InvalidParameter: An error that indicates you called an API with an invalid parameter.
	KCMMemoryPoolError_InvalidParameter KCMMemoryPoolError = 0
)

func (e KCMMemoryPoolError) String() string {
	switch e {
	case KCMMemoryPoolError_AllocationFailed:
		return "KCMMemoryPoolError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMMemoryPoolError(%d)", e)
	}
}

type KCMMetadataDataTypeRegistryError int

const (
	// KCMMetadataDataTypeRegistryError_AllocationFailed: An error that indicates the allocation fails.
	KCMMetadataDataTypeRegistryError_AllocationFailed KCMMetadataDataTypeRegistryError = 0
	// KCMMetadataDataTypeRegistryError_BadDataTypeIdentifier: An error that indicates a bad datatype identifier.
	KCMMetadataDataTypeRegistryError_BadDataTypeIdentifier KCMMetadataDataTypeRegistryError = 0
	// KCMMetadataDataTypeRegistryError_DataTypeAlreadyRegistered: An error that indicates the datatype is in a registered state.
	KCMMetadataDataTypeRegistryError_DataTypeAlreadyRegistered KCMMetadataDataTypeRegistryError = 0
	// KCMMetadataDataTypeRegistryError_MultipleConformingBaseTypes: An error that indicates that type you provide has more than one base data type.
	KCMMetadataDataTypeRegistryError_MultipleConformingBaseTypes KCMMetadataDataTypeRegistryError = 0
	// KCMMetadataDataTypeRegistryError_RequiredParameterMissing: An error that indicates a parameter the function requires is empty.
	KCMMetadataDataTypeRegistryError_RequiredParameterMissing KCMMetadataDataTypeRegistryError = 0
	// KCMMetadataDataTypeRegistryError_RequiresConformingBaseType: An error that indicates the data type you specify requires a conforming data type that resolves to a base data type.
	KCMMetadataDataTypeRegistryError_RequiresConformingBaseType KCMMetadataDataTypeRegistryError = 0
)

func (e KCMMetadataDataTypeRegistryError) String() string {
	switch e {
	case KCMMetadataDataTypeRegistryError_AllocationFailed:
		return "KCMMetadataDataTypeRegistryError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMMetadataDataTypeRegistryError(%d)", e)
	}
}

type KCMMetadataFormatType uint

const (
	// KCMMetadataFormatType_Boxed: CoreMedia boxed format.
	KCMMetadataFormatType_Boxed KCMMetadataFormatType = 0
	KCMMetadataFormatType_EMSG  KCMMetadataFormatType = 0
	// KCMMetadataFormatType_ICY: SHOUTCast format.
	KCMMetadataFormatType_ICY KCMMetadataFormatType = 0
	// KCMMetadataFormatType_ID3: ID3 format.
	KCMMetadataFormatType_ID3 KCMMetadataFormatType = 0
)

func (e KCMMetadataFormatType) String() string {
	switch e {
	case KCMMetadataFormatType_Boxed:
		return "KCMMetadataFormatType_Boxed"
	default:
		return fmt.Sprintf("KCMMetadataFormatType(%d)", e)
	}
}

type KCMMetadataIdentifierError int

const (
	// KCMMetadataIdentifierError_AllocationFailed: An error that indicates an allocation fails.
	KCMMetadataIdentifierError_AllocationFailed KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadIdentifier: An error that indicates the identifier isn’t valid.
	KCMMetadataIdentifierError_BadIdentifier KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadKey: An error that indicates a key isn’t valid.
	KCMMetadataIdentifierError_BadKey KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadKeyLength: An error that indicates that a key doesn’t have the correct length.
	KCMMetadataIdentifierError_BadKeyLength KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadKeySpace: An error that indicates the keyspace isn’t valid.
	KCMMetadataIdentifierError_BadKeySpace KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadKeyType: An error that indicates a key has a bad type.
	KCMMetadataIdentifierError_BadKeyType KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_BadNumberKey: An error that indicates the number for a key isn’t valid.
	KCMMetadataIdentifierError_BadNumberKey KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_NoKeyValueAvailable: An error that indicates a request for a key value in the anonymous keyspace.
	KCMMetadataIdentifierError_NoKeyValueAvailable KCMMetadataIdentifierError = 0
	// KCMMetadataIdentifierError_RequiredParameterMissing: An error that indicates a parameter the function requires is empty.
	KCMMetadataIdentifierError_RequiredParameterMissing KCMMetadataIdentifierError = 0
)

func (e KCMMetadataIdentifierError) String() string {
	switch e {
	case KCMMetadataIdentifierError_AllocationFailed:
		return "KCMMetadataIdentifierError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMMetadataIdentifierError(%d)", e)
	}
}

type KCMMuxedStreamType uint

const (
	// KCMMuxedStreamType_DV: DV stream.
	KCMMuxedStreamType_DV                            KCMMuxedStreamType = 0
	KCMMuxedStreamType_EmbeddedDeviceScreenRecording KCMMuxedStreamType = 0
	// KCMMuxedStreamType_MPEG1System: MPEG-1 System stream.
	KCMMuxedStreamType_MPEG1System KCMMuxedStreamType = 0
	// KCMMuxedStreamType_MPEG2Program: MPEG-2 Program stream.
	KCMMuxedStreamType_MPEG2Program KCMMuxedStreamType = 0
	// KCMMuxedStreamType_MPEG2Transport: MPEG-2 Transport stream.
	KCMMuxedStreamType_MPEG2Transport KCMMuxedStreamType = 0
)

func (e KCMMuxedStreamType) String() string {
	switch e {
	case KCMMuxedStreamType_DV:
		return "KCMMuxedStreamType_DV"
	default:
		return fmt.Sprintf("KCMMuxedStreamType(%d)", e)
	}
}

type KCMPersistentTrackID uint

const (
	// KCMPersistentTrackID_Invalid: Indicates an invalid track ID.
	KCMPersistentTrackID_Invalid KCMPersistentTrackID = 0
)

func (e KCMPersistentTrackID) String() string {
	switch e {
	case KCMPersistentTrackID_Invalid:
		return "KCMPersistentTrackID_Invalid"
	default:
		return fmt.Sprintf("KCMPersistentTrackID(%d)", e)
	}
}

type KCMPixelFormat uint

const (
	// KCMPixelFormat_16BE555: A type that describes 16-bit big-endian 5-5-5.
	KCMPixelFormat_16BE555 KCMPixelFormat = 0
	// KCMPixelFormat_16BE565: A type that describes 16-bit big-endian 5-6-5.
	KCMPixelFormat_16BE565 KCMPixelFormat = 0
	// KCMPixelFormat_16LE555: A type that describes 16-bit little-endian 5-5-5.
	KCMPixelFormat_16LE555 KCMPixelFormat = 0
	// KCMPixelFormat_16LE5551: A type that describes 16-bit little-endian 5-5-5-1.
	KCMPixelFormat_16LE5551 KCMPixelFormat = 0
	// KCMPixelFormat_16LE565: A type that describes 16-bit little-endian 5-6-5.
	KCMPixelFormat_16LE565 KCMPixelFormat = 0
	// KCMPixelFormat_24RGB: A type that describes 24-bit RGB.
	KCMPixelFormat_24RGB KCMPixelFormat = 0
	// KCMPixelFormat_32ARGB: A type that describes 32-bit ARGB.
	KCMPixelFormat_32ARGB KCMPixelFormat = 0
	// KCMPixelFormat_32BGRA: A type that describes 32-bit BGRA.
	KCMPixelFormat_32BGRA KCMPixelFormat = 0
	// KCMPixelFormat_422YpCbCr10: A type that describes component Y’CbCr 10-bit 4:2:2.
	KCMPixelFormat_422YpCbCr10 KCMPixelFormat = 0
	// KCMPixelFormat_422YpCbCr16: A type that describes component Y’CbCr 10,12,14,16-bit 4:2:2.
	KCMPixelFormat_422YpCbCr16 KCMPixelFormat = 0
	// KCMPixelFormat_422YpCbCr8: A type that describes component Y’CbCr 8-bit 4:2:2 ordered Cb Y’0 Cr Y’1.
	KCMPixelFormat_422YpCbCr8 KCMPixelFormat = 0
	// KCMPixelFormat_422YpCbCr8_yuvs: A type that describes component Y’CbCr 8-bit 4:2:2 ordered Y’0 Cb Y’1 Cr.
	KCMPixelFormat_422YpCbCr8_yuvs KCMPixelFormat = 0
	// KCMPixelFormat_4444YpCbCrA8: A type that describes component Y’CbCrA 8-bit 4:4:4:4.
	KCMPixelFormat_4444YpCbCrA8 KCMPixelFormat = 0
	// KCMPixelFormat_444YpCbCr10: A type that describes component Y’CbCr 10-bit 4:4:4
	KCMPixelFormat_444YpCbCr10 KCMPixelFormat = 0
	// KCMPixelFormat_444YpCbCr8: A type that describes component Y’CbCr 8-bit 4:4:4.
	KCMPixelFormat_444YpCbCr8 KCMPixelFormat = 0
	// KCMPixelFormat_8IndexedGray_WhiteIsZero: A type that describes 8-bit indexed gray, white is zero.
	KCMPixelFormat_8IndexedGray_WhiteIsZero KCMPixelFormat = 0
)

func (e KCMPixelFormat) String() string {
	switch e {
	case KCMPixelFormat_16BE555:
		return "KCMPixelFormat_16BE555"
	default:
		return fmt.Sprintf("KCMPixelFormat(%d)", e)
	}
}

type KCMSampleBufferError int

const (
	// KCMSampleBufferError_AllocationFailed: An error code that indicates the system failed to allocate memory.
	KCMSampleBufferError_AllocationFailed KCMSampleBufferError = 0
	// KCMSampleBufferError_AlreadyHasDataBuffer: An error code that indicates an attempt to set data on a sample buffer failed because that buffer already contains media data.
	KCMSampleBufferError_AlreadyHasDataBuffer KCMSampleBufferError = 0
	// KCMSampleBufferError_ArrayTooSmall: An error code that indicates the output array isn’t large enough to hold the requested array.
	KCMSampleBufferError_ArrayTooSmall KCMSampleBufferError = 0
	// KCMSampleBufferError_BufferHasNoSampleSizes: An error code that indicates a request for sample sizes on a buffer failed because the buffer doesn’t provide that information.
	KCMSampleBufferError_BufferHasNoSampleSizes KCMSampleBufferError = 0
	// KCMSampleBufferError_BufferHasNoSampleTimingInfo: An error code that indicates a request for sample timing on a buffer failed because the buffer doesn’t contain that information.
	KCMSampleBufferError_BufferHasNoSampleTimingInfo KCMSampleBufferError = 0
	// KCMSampleBufferError_BufferNotReady: An error code that indicates the system can’t make the buffer’s data ready for use.
	KCMSampleBufferError_BufferNotReady KCMSampleBufferError = 0
	// KCMSampleBufferError_CannotSubdivide: An error code that indicates a sample buffer doesn’t contain sample sizes.
	KCMSampleBufferError_CannotSubdivide KCMSampleBufferError = 0
	// KCMSampleBufferError_DataCanceled: An error code that indicates a sample buffer canceled its data-loading operation.
	KCMSampleBufferError_DataCanceled KCMSampleBufferError = 0
	// KCMSampleBufferError_DataFailed: An error code that indicates a sample buffer failed to load its data.
	KCMSampleBufferError_DataFailed KCMSampleBufferError = 0
	// KCMSampleBufferError_InvalidEntryCount: An error code that indicates a timing or size value isn’t within the allowed range.
	KCMSampleBufferError_InvalidEntryCount KCMSampleBufferError = 0
	// KCMSampleBufferError_InvalidMediaFormat: An error code that indicates the media format doesn’t match the sample buffer’s format description.
	KCMSampleBufferError_InvalidMediaFormat KCMSampleBufferError = 0
	// KCMSampleBufferError_InvalidMediaTypeForOperation: An error code that indicates the media type that the format description defines isn’t a value for the requested operation.
	KCMSampleBufferError_InvalidMediaTypeForOperation KCMSampleBufferError = 0
	// KCMSampleBufferError_InvalidSampleData: An error code that indicates the sample buffer contains bad data.
	KCMSampleBufferError_InvalidSampleData KCMSampleBufferError = 0
	// KCMSampleBufferError_Invalidated: An error code that indicates a sample buffer invalidated its data.
	KCMSampleBufferError_Invalidated KCMSampleBufferError = 0
	// KCMSampleBufferError_RequiredParameterMissing: An error code that indicates a required parameter’s value is invalid.
	KCMSampleBufferError_RequiredParameterMissing KCMSampleBufferError = 0
	// KCMSampleBufferError_SampleIndexOutOfRange: An error code that indicates the sample index is outside the range of samples that the buffer contains.
	KCMSampleBufferError_SampleIndexOutOfRange KCMSampleBufferError = 0
	// KCMSampleBufferError_SampleTimingInfoInvalid: An error code that indicates the sample buffer unexpectedly contains nonnumeric sample-timing information.
	KCMSampleBufferError_SampleTimingInfoInvalid KCMSampleBufferError = 0
)

func (e KCMSampleBufferError) String() string {
	switch e {
	case KCMSampleBufferError_AllocationFailed:
		return "KCMSampleBufferError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMSampleBufferError(%d)", e)
	}
}

type KCMSimpleQueueError int

const (
	// KCMSimpleQueueError_AllocationFailed: The system failed to allocate memory.
	KCMSimpleQueueError_AllocationFailed KCMSimpleQueueError = 0
	// KCMSimpleQueueError_ParameterOutOfRange: You passed a parameter to a function that’s outside the range of allowed values.
	KCMSimpleQueueError_ParameterOutOfRange KCMSimpleQueueError = 0
	// KCMSimpleQueueError_QueueIsFull: An operation failed because the queue is full.
	KCMSimpleQueueError_QueueIsFull KCMSimpleQueueError = 0
	// KCMSimpleQueueError_RequiredParameterMissing: You failed to pass a required parameter to a function.
	KCMSimpleQueueError_RequiredParameterMissing KCMSimpleQueueError = 0
)

func (e KCMSimpleQueueError) String() string {
	switch e {
	case KCMSimpleQueueError_AllocationFailed:
		return "KCMSimpleQueueError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMSimpleQueueError(%d)", e)
	}
}

type KCMSubtitleFormatType uint

const (
	KCMSubtitleFormatType_3GText KCMSubtitleFormatType = 0
	KCMSubtitleFormatType_WebVTT KCMSubtitleFormatType = 0
)

func (e KCMSubtitleFormatType) String() string {
	switch e {
	case KCMSubtitleFormatType_3GText:
		return "KCMSubtitleFormatType_3GText"
	default:
		return fmt.Sprintf("KCMSubtitleFormatType(%d)", e)
	}
}

type KCMSyncError int

const (
	// KCMSyncError_AllocationFailed: A sync error that indicates the memory allocation fails.
	KCMSyncError_AllocationFailed KCMSyncError = 0
	// KCMSyncError_InvalidParameter: A sync error that indicates a parameter isn’t valid.
	KCMSyncError_InvalidParameter KCMSyncError = 0
	// KCMSyncError_MissingRequiredParameter: A sync error that indicates a parameter is missing.
	KCMSyncError_MissingRequiredParameter KCMSyncError = 0
	// KCMSyncError_RateMustBeNonZero: A sync error that indicates the rate is nonzero.
	KCMSyncError_RateMustBeNonZero KCMSyncError = 0
)

func (e KCMSyncError) String() string {
	switch e {
	case KCMSyncError_AllocationFailed:
		return "KCMSyncError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMSyncError(%d)", e)
	}
}

type KCMTaggedBufferGroupFormatType uint

const (
	KCMTaggedBufferGroupFormatType_TaggedBufferGroup KCMTaggedBufferGroupFormatType = 0
)

func (e KCMTaggedBufferGroupFormatType) String() string {
	switch e {
	case KCMTaggedBufferGroupFormatType_TaggedBufferGroup:
		return "KCMTaggedBufferGroupFormatType_TaggedBufferGroup"
	default:
		return fmt.Sprintf("KCMTaggedBufferGroupFormatType(%d)", e)
	}
}

type KCMTextDisplayFlag uint

const (
	// KCMTextDisplayFlag_allSubtitlesForced: A flag that describes treating all subtitle samples as if they contain forced subtitles.
	KCMTextDisplayFlag_allSubtitlesForced KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_continuousKaraoke: A flag that describes enabling the continuous karaoke mode where the range of karaoke highlighting extends to include additional ranges rather than the highlighting moves onto the next range.
	KCMTextDisplayFlag_continuousKaraoke KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_fillTextRegion: A flag that describes the subtitle display bounds are to be filled with the color specified by `kCMTextFormatDescriptionExtension_BackgroundColor`.
	KCMTextDisplayFlag_fillTextRegion KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_forcedSubtitlesPresent: A flag that describes forcing subtitles are present, for example, a subtitle which only displays during foreign language sections of the video.
	KCMTextDisplayFlag_forcedSubtitlesPresent KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_obeySubtitleFormatting: A flag that describes using the subtitle display bounds to determine if the system places the subtitltes near the top or bottom of the video.
	KCMTextDisplayFlag_obeySubtitleFormatting KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollDirectionMask: A flag that describes the scrolling direction is set by a two-bit field, obtained from displayFlags using kCMTextDisplayFlag_scrollDirectionMask.
	KCMTextDisplayFlag_scrollDirectionMask KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollDirection_bottomToTop: A flag that describes the text is vertically scrolled up (“credits style”), entering from the bottom and leaving towards the top.
	KCMTextDisplayFlag_scrollDirection_bottomToTop KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollDirection_leftToRight: A flag that describes the text is horizontally scrolled, entering from the left and leaving towards the right.
	KCMTextDisplayFlag_scrollDirection_leftToRight KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollDirection_rightToLeft: A flag that describes the text is horizontally scrolled (“marquee style”), entering from the right and leaving towards the left.
	KCMTextDisplayFlag_scrollDirection_rightToLeft KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollDirection_topToBottom: A flag that describes the text is vertically scrolled down, entering from the top and leaving towards the bottom.
	KCMTextDisplayFlag_scrollDirection_topToBottom KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollIn: A flag that describes the text scrolls into the display region.
	KCMTextDisplayFlag_scrollIn KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_scrollOut: A flag that describes the text scrolls out of the display region.
	KCMTextDisplayFlag_scrollOut KCMTextDisplayFlag = 0
	// KCMTextDisplayFlag_writeTextVertically: A flag that describes the text renders vertically.
	KCMTextDisplayFlag_writeTextVertically KCMTextDisplayFlag = 0
)

func (e KCMTextDisplayFlag) String() string {
	switch e {
	case KCMTextDisplayFlag_allSubtitlesForced:
		return "KCMTextDisplayFlag_allSubtitlesForced"
	default:
		return fmt.Sprintf("KCMTextDisplayFlag(%d)", e)
	}
}

type KCMTextFormatType uint

const (
	// KCMTextFormatType_3GText: A type that describes 3GPP text media.
	KCMTextFormatType_3GText KCMTextFormatType = 0
	// KCMTextFormatType_QTText: A type that describes QuickTime text media.
	KCMTextFormatType_QTText KCMTextFormatType = 0
)

func (e KCMTextFormatType) String() string {
	switch e {
	case KCMTextFormatType_3GText:
		return "KCMTextFormatType_3GText"
	default:
		return fmt.Sprintf("KCMTextFormatType(%d)", e)
	}
}

type KCMTextJustification uint

const (
	// KCMTextJustification_bottom_right: A type that describes bottom justification when specified for vertical justification, right justification for horizontal justification.
	KCMTextJustification_bottom_right KCMTextJustification = 0
	// KCMTextJustification_centered: A type that describes center justification (both horizontal and vertical justification).
	KCMTextJustification_centered KCMTextJustification = 0
	// KCMTextJustification_left_top: A type that describes left justification when specified for horizontal justification, top justification for vertical justification.
	KCMTextJustification_left_top KCMTextJustification = 0
)

func (e KCMTextJustification) String() string {
	switch e {
	case KCMTextJustification_bottom_right:
		return "KCMTextJustification_bottom_right"
	default:
		return fmt.Sprintf("KCMTextJustification(%d)", e)
	}
}

type KCMTimeCodeFlag uint

const (
	// KCMTimeCodeFlag_24HourMax: A type that describes timecode rolls over every 24 hours.
	KCMTimeCodeFlag_24HourMax KCMTimeCodeFlag = 0
	// KCMTimeCodeFlag_DropFrame: A type that describes timecodes are to be rendered in drop-frame format.
	KCMTimeCodeFlag_DropFrame KCMTimeCodeFlag = 0
	// KCMTimeCodeFlag_NegTimesOK: A type that describes that the track may contain negative timecodes.
	KCMTimeCodeFlag_NegTimesOK KCMTimeCodeFlag = 0
)

func (e KCMTimeCodeFlag) String() string {
	switch e {
	case KCMTimeCodeFlag_24HourMax:
		return "KCMTimeCodeFlag_24HourMax"
	default:
		return fmt.Sprintf("KCMTimeCodeFlag(%d)", e)
	}
}

type KCMTimeCodeFormatType uint

const (
	// KCMTimeCodeFormatType_Counter32: 32-bit counter-mode sample.
	KCMTimeCodeFormatType_Counter32 KCMTimeCodeFormatType = 0
	// KCMTimeCodeFormatType_Counter64: 64-bit counter-mode sample.
	KCMTimeCodeFormatType_Counter64 KCMTimeCodeFormatType = 0
	// KCMTimeCodeFormatType_TimeCode32: 32-bit time code sample.
	KCMTimeCodeFormatType_TimeCode32 KCMTimeCodeFormatType = 0
	// KCMTimeCodeFormatType_TimeCode64: 64-bit time code sample.
	KCMTimeCodeFormatType_TimeCode64 KCMTimeCodeFormatType = 0
)

func (e KCMTimeCodeFormatType) String() string {
	switch e {
	case KCMTimeCodeFormatType_Counter32:
		return "KCMTimeCodeFormatType_Counter32"
	default:
		return fmt.Sprintf("KCMTimeCodeFormatType(%d)", e)
	}
}

type KCMTimebaseError int

const (
	// KCMTimebaseError_AllocationFailed: A timebase error that indicates the memory allocation fails.
	KCMTimebaseError_AllocationFailed KCMTimebaseError = 0
	// KCMTimebaseError_InvalidParameter: A timebase error that indicates a parameter isn’t valid.
	KCMTimebaseError_InvalidParameter KCMTimebaseError = 0
	// KCMTimebaseError_MissingRequiredParameter: A timebase error that indicates a parameter is missing.
	KCMTimebaseError_MissingRequiredParameter KCMTimebaseError = 0
	// KCMTimebaseError_ReadOnly: A timebase error that indicates the system attempts to modify a read-only timebase.
	KCMTimebaseError_ReadOnly KCMTimebaseError = 0
	// KCMTimebaseError_TimerIntervalTooShort: A timebase error that indicates the time interval is too short.
	KCMTimebaseError_TimerIntervalTooShort KCMTimebaseError = 0
)

func (e KCMTimebaseError) String() string {
	switch e {
	case KCMTimebaseError_AllocationFailed:
		return "KCMTimebaseError_AllocationFailed"
	default:
		return fmt.Sprintf("KCMTimebaseError(%d)", e)
	}
}

type KCMVideoCodecType uint

const (
	// KCMVideoCodecType_422YpCbCr8: A type that identifies a component with the format of Y’CbCr 8-bit 4:2:2 ordered Cb Y’0 Cr Y’1.
	KCMVideoCodecType_422YpCbCr8 KCMVideoCodecType = 0
	KCMVideoCodecType_AV1        KCMVideoCodecType = 0
	// KCMVideoCodecType_Animation: A type that identifies the apple animation format.
	KCMVideoCodecType_Animation KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes422: A type that identifies the Apple ProRes 422 format.
	KCMVideoCodecType_AppleProRes422 KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes422HQ: A type that identifies the Apple ProRes 422 HQ format.
	KCMVideoCodecType_AppleProRes422HQ KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes422LT: A type that identifies the Apple ProRes 422 LT format.
	KCMVideoCodecType_AppleProRes422LT KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes422Proxy: A type that identifies the Apple ProRes 422 proxy format.
	KCMVideoCodecType_AppleProRes422Proxy KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes4444: A type that identifies the Apple ProRes 4444 format.
	KCMVideoCodecType_AppleProRes4444 KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProRes4444XQ: A type that identifies the Apple ProRes 4444 XQ format.
	KCMVideoCodecType_AppleProRes4444XQ KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProResRAW: A type that identifies the Apple ProRes RAW format.
	KCMVideoCodecType_AppleProResRAW KCMVideoCodecType = 0
	// KCMVideoCodecType_AppleProResRAWHQ: A type that identifies the Apple ProRes RAW HQ format.
	KCMVideoCodecType_AppleProResRAWHQ KCMVideoCodecType = 0
	// KCMVideoCodecType_Cinepak: A type that identifies the cinepak format.
	KCMVideoCodecType_Cinepak KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCNTSC: A type that identifies the DV NTSC format.
	KCMVideoCodecType_DVCNTSC KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPAL: A type that identifies the DV PAL format.
	KCMVideoCodecType_DVCPAL KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD1080i50: A type that identifies the Panasonic DVCPro-HD 1080i50 format.
	KCMVideoCodecType_DVCPROHD1080i50 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD1080i60: A type that identifies the Panasonic DVCPro-HD 1080i60 format.
	KCMVideoCodecType_DVCPROHD1080i60 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD1080p25: A type that identifies the Panasonic DVCPro-HD 1080p25 format.
	KCMVideoCodecType_DVCPROHD1080p25 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD1080p30: A type that identifies the Panasonic DVCPro-HD 1080p30 format.
	KCMVideoCodecType_DVCPROHD1080p30 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD720p50: A type that identifies the Panasonic DVCPro-HD 720p50 format.
	KCMVideoCodecType_DVCPROHD720p50 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPROHD720p60: A type that identifies the Panasonic DVCPro-HD 720p60 format.
	KCMVideoCodecType_DVCPROHD720p60 KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPro50NTSC: A type that identifies the Panasonic DVCPro-50 NTSC format.
	KCMVideoCodecType_DVCPro50NTSC KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCPro50PAL: A type that identifies the Panasonic DVCPro-50 PAL format.
	KCMVideoCodecType_DVCPro50PAL KCMVideoCodecType = 0
	// KCMVideoCodecType_DVCProPAL: A type that identifies the Panasonic DVCPro PAL format.
	KCMVideoCodecType_DVCProPAL KCMVideoCodecType = 0
	// KCMVideoCodecType_DepthHEVC: A type that identifies the depth HEVC format.
	KCMVideoCodecType_DepthHEVC KCMVideoCodecType = 0
	// KCMVideoCodecType_DisparityHEVC: A type that identifies the disparity HEVC format.
	KCMVideoCodecType_DisparityHEVC KCMVideoCodecType = 0
	// KCMVideoCodecType_DolbyVisionHEVC: A type that identifies the Dolby Vision HEVC format.
	KCMVideoCodecType_DolbyVisionHEVC KCMVideoCodecType = 0
	// KCMVideoCodecType_H263: A type that identifies the ITU-T H.263 format.
	KCMVideoCodecType_H263 KCMVideoCodecType = 0
	// KCMVideoCodecType_H264: A type that identifies the ITU-T H.264 format.
	KCMVideoCodecType_H264 KCMVideoCodecType = 0
	// KCMVideoCodecType_HEVC: A type that identifies the ITU-T HEVC format.
	KCMVideoCodecType_HEVC KCMVideoCodecType = 0
	// KCMVideoCodecType_HEVCWithAlpha: A type that identifies the HEVC format with alpha support.
	KCMVideoCodecType_HEVCWithAlpha KCMVideoCodecType = 0
	// KCMVideoCodecType_JPEG: A type that identifies the Joint Photographic Experts Group (JPEG) format.
	KCMVideoCodecType_JPEG KCMVideoCodecType = 0
	// KCMVideoCodecType_JPEG_OpenDML: A type that identifies the JPEG format with Open-DML extensions.
	KCMVideoCodecType_JPEG_OpenDML KCMVideoCodecType = 0
	KCMVideoCodecType_JPEG_XL      KCMVideoCodecType = 0
	// KCMVideoCodecType_MPEG1Video: A type that identifies the MPEG-1 video format.
	KCMVideoCodecType_MPEG1Video KCMVideoCodecType = 0
	// KCMVideoCodecType_MPEG2Video: A type that identifies the MPEG-2 video format.
	KCMVideoCodecType_MPEG2Video KCMVideoCodecType = 0
	// KCMVideoCodecType_MPEG4Video: A type that identifies the Moving Picture Experts Group (MPEG) MPEG-4 Part 2 video format.
	KCMVideoCodecType_MPEG4Video KCMVideoCodecType = 0
	// KCMVideoCodecType_SorensonVideo: A type that identifies the sorenson video format.
	KCMVideoCodecType_SorensonVideo KCMVideoCodecType = 0
	// KCMVideoCodecType_SorensonVideo3: A type that identifies the sorenson 3 video format.
	KCMVideoCodecType_SorensonVideo3 KCMVideoCodecType = 0
	// KCMVideoCodecType_VP9: A type that identifies the VP9 format.
	KCMVideoCodecType_VP9 KCMVideoCodecType = 0
)

func (e KCMVideoCodecType) String() string {
	switch e {
	case KCMVideoCodecType_422YpCbCr8:
		return "KCMVideoCodecType_422YpCbCr8"
	default:
		return fmt.Sprintf("KCMVideoCodecType(%d)", e)
	}
}

type KcmattachmentmodeShould uint

const (
	// KCMAttachmentMode_ShouldNotPropagate: A mode that doesn’t propagate attachments to another object.
	KCMAttachmentMode_ShouldNotPropagate KcmattachmentmodeShould = 0
	// KCMAttachmentMode_ShouldPropagate: A mode that propagates attachments to another object.
	KCMAttachmentMode_ShouldPropagate KcmattachmentmodeShould = 0
)

func (e KcmattachmentmodeShould) String() string {
	switch e {
	case KCMAttachmentMode_ShouldNotPropagate:
		return "KCMAttachmentMode_ShouldNotPropagate"
	default:
		return fmt.Sprintf("KcmattachmentmodeShould(%d)", e)
	}
}

type KcmaudiocodectypeAac uint

const (
	// KCMAudioCodecType_AAC_AudibleProtected: Audible’s protected AAC.
	KCMAudioCodecType_AAC_AudibleProtected KcmaudiocodectypeAac = 0
	// KCMAudioCodecType_AAC_LCProtected: iTMS protected low-complexity AAC.
	KCMAudioCodecType_AAC_LCProtected KcmaudiocodectypeAac = 0
)

func (e KcmaudiocodectypeAac) String() string {
	switch e {
	case KCMAudioCodecType_AAC_AudibleProtected:
		return "KCMAudioCodecType_AAC_AudibleProtected"
	default:
		return fmt.Sprintf("KcmaudiocodectypeAac(%d)", e)
	}
}

type KcmbufferqueuetriggerWhen uint

const (
	// KCMBufferQueueTrigger_WhenBufferCountBecomesGreaterThan: Trigger fires when buffer count becomes > the specified threshold number.
	KCMBufferQueueTrigger_WhenBufferCountBecomesGreaterThan KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenBufferCountBecomesLessThan: Trigger fires when buffer count becomes less than the specified threshold number.
	KCMBufferQueueTrigger_WhenBufferCountBecomesLessThan KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenDataBecomesReady: Trigger fires when next dequeueable buffer becomes ready (that is, CMBufferQueueDequeueIfDataReady(_:) will now succeed).
	KCMBufferQueueTrigger_WhenDataBecomesReady KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenDurationBecomesGreaterThan: Trigger fires when queue duration becomes greater than the specified duration.
	KCMBufferQueueTrigger_WhenDurationBecomesGreaterThan KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenDurationBecomesGreaterThanOrEqualTo: Trigger fires when queue duration becomes greater than or equal to the specified duration.
	KCMBufferQueueTrigger_WhenDurationBecomesGreaterThanOrEqualTo                                 KcmbufferqueuetriggerWhen = 0
	KCMBufferQueueTrigger_WhenDurationBecomesGreaterThanOrEqualToAndBufferCountBecomesGreaterThan KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenDurationBecomesLessThan: Trigger fires when queue duration becomes less than the specified duration.
	KCMBufferQueueTrigger_WhenDurationBecomesLessThan KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenDurationBecomesLessThanOrEqualTo: Trigger fires when queue duration becomes less than or equal to the specified duration.
	KCMBufferQueueTrigger_WhenDurationBecomesLessThanOrEqualTo KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenEndOfDataReached: Trigger fires when CMBufferQueueIsAtEndOfData’s condition becomes true.
	KCMBufferQueueTrigger_WhenEndOfDataReached KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenMaxPresentationTimeStampChanges: Trigger fires when the maximum presentation timestamp changes (triggerDuration is ignored).
	KCMBufferQueueTrigger_WhenMaxPresentationTimeStampChanges KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenMinPresentationTimeStampChanges: Trigger fires when the minimum presentation timestamp changes (triggerDuration is ignored).
	KCMBufferQueueTrigger_WhenMinPresentationTimeStampChanges KcmbufferqueuetriggerWhen = 0
	// KCMBufferQueueTrigger_WhenReset: Trigger fires when CMBufferQueueReset called.
	KCMBufferQueueTrigger_WhenReset KcmbufferqueuetriggerWhen = 0
)

func (e KcmbufferqueuetriggerWhen) String() string {
	switch e {
	case KCMBufferQueueTrigger_WhenBufferCountBecomesGreaterThan:
		return "KCMBufferQueueTrigger_WhenBufferCountBecomesGreaterThan"
	default:
		return fmt.Sprintf("KcmbufferqueuetriggerWhen(%d)", e)
	}
}

type KcmsamplebufferflagAudiobufferlist uint

const (
	// KCMSampleBufferFlag_AudioBufferList_Assure16ByteAlignment: Indicates that memory involved in audio buffer lists is 16-byte aligned.
	KCMSampleBufferFlag_AudioBufferList_Assure16ByteAlignment KcmsamplebufferflagAudiobufferlist = 0
)

func (e KcmsamplebufferflagAudiobufferlist) String() string {
	switch e {
	case KCMSampleBufferFlag_AudioBufferList_Assure16ByteAlignment:
		return "KCMSampleBufferFlag_AudioBufferList_Assure16ByteAlignment"
	default:
		return fmt.Sprintf("KcmsamplebufferflagAudiobufferlist(%d)", e)
	}
}
