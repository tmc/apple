// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo
import (
	"unsafe"
)

// C struct types
// CVFillExtendedPixelsCallBackData - A structure for holding information that describes a custom extended pixel fill algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVFillExtendedPixelsCallBackData
type CVFillExtendedPixelsCallBackData struct {
	Version int // The version of this fill algorithm.
	FillCallBack CVFillExtendedPixelsCallBack
	RefCon unsafe.Pointer // A pointer to application-defined data that is passed to your custom pixel fill function.

}

// CVPlanarComponentInfo - A structure for describing planar components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarComponentInfo
type CVPlanarComponentInfo struct {
	Offset int32 // The offset from the main base address to the base address of this plane. (big-endian)
	RowBytes uint32 // The number of bytes per row of this plane. (big-endian)

}

// CVPlanarPixelBufferInfo - A structure for describing planar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo
type CVPlanarPixelBufferInfo struct {
	ComponentInfo CVPlanarComponentInfo // An array containing a [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure for each plane of the buffer.

}

// CVPlanarPixelBufferInfo_YCbCrBiPlanar - A structure for describing YCbCr biplanar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo_YCbCrBiPlanar
type CVPlanarPixelBufferInfo_YCbCrBiPlanar struct {
	ComponentInfoY CVPlanarComponentInfo // A [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure containing information on the Y component of the buffer.
	ComponentInfoCbCr CVPlanarComponentInfo // A [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure containing information on the Cb/Cr component of the buffer.

}

// CVPlanarPixelBufferInfo_YCbCrPlanar - A structure for describing YCbCr planar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo_YCbCrPlanar
type CVPlanarPixelBufferInfo_YCbCrPlanar struct {
	ComponentInfoY CVPlanarComponentInfo // A [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure containing information on the Y component of the buffer.
	ComponentInfoCb CVPlanarComponentInfo // A [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure containing information on the Cb component of the buffer.
	ComponentInfoCr CVPlanarComponentInfo // A [CVPlanarComponentInfo](<doc://com.apple.corevideo/documentation/CoreVideo/CVPlanarComponentInfo>) structure containing information on the Cr component of the buffer.

}

// CVSMPTETime - A structure for holding an SMPTE time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETime
type CVSMPTETime struct {
	Subframes int16 // The number of subframes in the full message.
	SubframeDivisor int16 // The number of subframes per frame (typically, 80).
	Counter uint32 // The total number of messages received.
	Type uint32 // The kind of SMPTE time type.
	Flags uint32 // A set of flags that indicate the SMPTE state.
	Hours int16 // The number of hours in the full message.
	Minutes int16 // The number of minutes in the full message.
	Seconds int16 // The number of seconds in the full message.
	Frames int16 // The number of frames in the full message.

}

// CVTime - A structure for reporting Core Video time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTime
type CVTime struct {
	TimeValue int64 // The time value.
	TimeScale int32 // The time scale for this value.
	Flags int32 // The flags associated with the [CVTime] value. See [CVTime Values](<doc://com.apple.corevideo/documentation/CoreVideo/cvtime-values>) for possible values. If `kCVTimeIsIndefinite` is set, you should not use any of the other fields in this structure.

}

// CVTimeStamp - A structure for defining a display timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStamp
type CVTimeStamp struct {
	Version uint32 // The current [CVTimeStamp] structure is version 0. Some functions require you to specify a version when passing in a timestamp structure to be filled.
	VideoTimeScale int32 // The scale (in units per second) of the `videoTimeScale` and `videoRefreshPeriod` fields.
	VideoTime int64 // The start of a frame (or field for interlaced video).
	HostTime uint64 // The system time measured by the timestamp.
	RateScalar float64 // The current rate of the device as measured by the timestamps, divided by the nominal rate.
	VideoRefreshPeriod int64
	SmpteTime CVSMPTETime // The SMPTE time representation of the timestamp.
	Flags uint64 // A bit field containing additional information about the timestamp.
	Reserved uint64 // Reserved. Do not use.

}








