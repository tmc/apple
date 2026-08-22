// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/coremedia"
)

// C struct types

// AVCaptionDimension - A structure that defines a caption dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionDimension
type AVCaptionDimension struct {
	Value float64            // The value of the coordinate or length.
	Units AVCaptionUnitsType // The units of the coordinate, such as cells or points.

}

// AVCaptionPoint - A structure that defines the origin point for a caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionPoint
type AVCaptionPoint struct {
	X AVCaptionDimension // The point’s x coordinate.
	Y AVCaptionDimension // The point’s y coordinate.

}

// AVCaptionSize - A structure that defines the height and width of a caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSize
type AVCaptionSize struct {
	Width  AVCaptionDimension // The width of the caption.
	Height AVCaptionDimension // The height of the caption.

}

// AVCaptureTimecode - This structure represents a timecode, adhering to SMPTE standards, which define precise time information and associated timestamps for video or audio synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode
type AVCaptureTimecode struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint64
	storage [40]byte
}

// Hours returns the Hours field from the record's packed storage.
func (s *AVCaptureTimecode) Hours() uint8 {
	return uint8(s.storage[0])
}

// SetHours updates the Hours field in the record's packed storage.
func (s *AVCaptureTimecode) SetHours(v uint8) {
	s.storage[0] = uint8(v)
}

// Minutes returns the Minutes field from the record's packed storage.
func (s *AVCaptureTimecode) Minutes() uint8 {
	return uint8(s.storage[1])
}

// SetMinutes updates the Minutes field in the record's packed storage.
func (s *AVCaptureTimecode) SetMinutes(v uint8) {
	s.storage[1] = uint8(v)
}

// Seconds returns the Seconds field from the record's packed storage.
func (s *AVCaptureTimecode) Seconds() uint8 {
	return uint8(s.storage[2])
}

// SetSeconds updates the Seconds field in the record's packed storage.
func (s *AVCaptureTimecode) SetSeconds(v uint8) {
	s.storage[2] = uint8(v)
}

// Frames returns the Frames field from the record's packed storage.
func (s *AVCaptureTimecode) Frames() uint8 {
	return uint8(s.storage[3])
}

// SetFrames updates the Frames field in the record's packed storage.
func (s *AVCaptureTimecode) SetFrames(v uint8) {
	s.storage[3] = uint8(v)
}

// UserBits returns the UserBits field from the record's packed storage.
func (s *AVCaptureTimecode) UserBits() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetUserBits updates the UserBits field in the record's packed storage.
func (s *AVCaptureTimecode) SetUserBits(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FrameDuration returns the FrameDuration field from the record's packed storage.
func (s *AVCaptureTimecode) FrameDuration() coremedia.CMTime {
	return *(*coremedia.CMTime)(unsafe.Pointer(&s.storage[8]))
}

// SetFrameDuration updates the FrameDuration field in the record's packed storage.
func (s *AVCaptureTimecode) SetFrameDuration(v coremedia.CMTime) {
	*(*coremedia.CMTime)(unsafe.Pointer(&s.storage[8])) = v
}

// SourceType returns the SourceType field from the record's packed storage.
func (s *AVCaptureTimecode) SourceType() AVCaptureTimecodeSourceType {
	return *(*AVCaptureTimecodeSourceType)(unsafe.Pointer(&s.storage[32]))
}

// SetSourceType updates the SourceType field in the record's packed storage.
func (s *AVCaptureTimecode) SetSourceType(v AVCaptureTimecodeSourceType) {
	*(*AVCaptureTimecodeSourceType)(unsafe.Pointer(&s.storage[32])) = v
}

// AVCaptureWhiteBalanceChromaticityValues - A structure that defines CIE 1931 xy chromaticity values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceChromaticityValues
type AVCaptureWhiteBalanceChromaticityValues struct {
	X float32 // The x component of the CIE 1931 chromaticity value.
	Y float32 // The y component of the CIE 1931 chromaticity value.

}

// AVCaptureWhiteBalanceGains - A structure that defines RGB white balance gain values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceGains
type AVCaptureWhiteBalanceGains struct {
	RedGain   float32 // The red gain component of the white balance value.
	GreenGain float32 // The green gain component of the white balance value.
	BlueGain  float32 // The blue gain component of the white balance value.

}

// AVCaptureWhiteBalanceTemperatureAndTintValues - A structure that defines temperature and tint values correlated to a white-balance color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceTemperatureAndTintValues
type AVCaptureWhiteBalanceTemperatureAndTintValues struct {
	Temperature float32 // The white balance color correlated temperature in kelvin.
	Tint        float32 // The white balance tint value in the range of `-150.0` through `+150.0`.

}

// AVEdgeWidths - A structure that defines edge processing region widths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVEdgeWidths
type AVEdgeWidths struct {
	Left   float64 // The left-edge width.
	Top    float64 // The top-edge width.
	Right  float64 // The right-edge width.
	Bottom float64 // The bottom-edge width.

}

// AVPixelAspectRatio - A structure that defines a pixel aspect ratio for a rendering context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPixelAspectRatio
type AVPixelAspectRatio struct {
	HorizontalSpacing int // The pixel aspect ratio’s horizontal spacing value.
	VerticalSpacing   int // The pixel aspect ratio’s vertical spacing value.

}

// AVSampleCursorAudioDependencyInfo - A structure that describes the independent decodability of audio samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorAudioDependencyInfo
type AVSampleCursorAudioDependencyInfo struct {
	AudioSampleIsIndependentlyDecodable bool // A Boolean value indicating whether the sample is independently decodable.
	AudioSamplePacketRefreshCount       int  // The number of samples, starting at the current sample, that must be fed to the decoder to achieve full decoder refresh.

}

// AVSampleCursorChunkInfo - A value that provides information about a chunk of media samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorChunkInfo
type AVSampleCursorChunkInfo struct {
	ChunkSampleCount                  int64 // The count of media samples in the chunk.
	ChunkHasUniformSampleSizes        bool  // The samples in the chunk occupy the same number of bytes in storage.
	ChunkHasUniformSampleDurations    bool  // The samples in the chunk have the same duration.
	ChunkHasUniformFormatDescriptions bool  // The samples in the chunk have the same format description.

}

// AVSampleCursorDependencyInfo - A value for describing dependencies between a media sample and other media samples in the same sample sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorDependencyInfo
type AVSampleCursorDependencyInfo struct {
	SampleIndicatesWhetherItHasDependentSamples bool // A Boolean value that determines whether the sample indicates if other samples depend on it.
	SampleHasDependentSamples                   bool // A Boolean value that determines whether the sample has dependent samples.
	SampleIndicatesWhetherItDependsOnOthers     bool // A Boolean value that determines whether the sample indicates that it depends on other samples.
	SampleDependsOnOthers                       bool // A Boolean value that determines whether the sample depends on other samples.
	SampleIndicatesWhetherItHasRedundantCoding  bool // A Boolean value that determines whether the sample indicates that it has redundant coding.
	SampleHasRedundantCoding                    bool // A Boolean value that determines whether the sample has redundant coding.

}

// AVSampleCursorStorageRange - A structure that indicates the offset and length of storage for a media sample or its chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorStorageRange
type AVSampleCursorStorageRange struct {
	Offset int64 // The offset of the first byte of storage that a media sample or its chunk occupies.
	Length int64 // The count of bytes of storage that a media sample or its chunk occupies.

}

// AVSampleCursorSyncInfo - A structure that describes the attributes of media samples to consider when resynchronizing a decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorSyncInfo
type AVSampleCursorSyncInfo struct {
	SampleIsFullSync    bool // A Boolean value that indicates whether a sample is a full sync sample.
	SampleIsPartialSync bool // A Boolean value that indicates whether a sample is a partial sync sample.
	SampleIsDroppable   bool // A Boolean value that indicates whether a sample is droppable.

}
