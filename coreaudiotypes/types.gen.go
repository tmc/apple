// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes

import (
	"unsafe"
)

// C struct types

// AudioBuffer - A structure that holds a buffer of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBuffer
type AudioBuffer struct {
	MNumberChannels uint32         // The number of interleaved channels in the buffer.
	MDataByteSize   uint32         // The number of bytes in the buffer.
	MData           unsafe.Pointer // A pointer to a buffer of audio data.

}

// AudioBufferList - A structure that stores a variable-length array of audio buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBufferList
type AudioBufferList struct {
	MNumberBuffers uint32         // The number of audio buffers in the list.
	MBuffers       [1]AudioBuffer // A variable-length array of audio buffers.

}

// AudioChannelDescription - A structure that describes a channel of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelDescription
type AudioChannelDescription struct {
	MChannelLabel AudioChannelLabel // A label that describes the audio channel.
	MChannelFlags AudioChannelFlags // The audio channel flags that indicate how to interpret the channel coordinates.
	MCoordinates  [3]float32        // The coordinates that specify a precise speaker location.

}

// AudioChannelLayout - A structure that specifies a channel layout in a file or in hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelLayout
type AudioChannelLayout struct {
	MChannelLayoutTag          AudioChannelLayoutTag      // The [AudioChannelLayoutTag] value that indicates the layout. See [Audio Channel Layout Tags](<https://developer.apple.com/documentation/CoreAudioTypes/audio-channel-layout-tags>) for possible values.
	MChannelBitmap             AudioChannelBitmap         // If `mChannelLayoutTag` is set to `kAudioChannelLayoutTag_UseChannelBitmap`, this field is the channel-use bitmap.
	MNumberChannelDescriptions uint32                     // The number of items in the `mChannelDescriptions` array.
	MChannelDescriptions       [1]AudioChannelDescription // A variable length array of `mNumberChannelDescription` elements that describes a layout. If the `mChannelLayoutTag` field is set to `kAudioChannelLayoutTag_UseChannelDescriptions`, use this field to describe the layout.

}

// AudioClassDescription - A structure that describes an audio codec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioClassDescription
type AudioClassDescription struct {
	MType         uint32 // A four character code that a manufacturer defines for a codec type.
	MSubType      uint32 // A four character code that a manufacturer defines for a codec subtype.
	MManufacturer uint32 // A four character code that identifies a codec manufacturer.

}

// AudioStreamBasicDescription - A format specification for an audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
type AudioStreamBasicDescription struct {
	MSampleRate       float64          // The number of frames per second of the data in the stream, when playing the stream at normal speed.
	MFormatID         AudioFormatID    // An identifier specifying the general audio data format in the stream.
	MFormatFlags      AudioFormatFlags // Format-specific flags to specify details of the format.
	MBytesPerPacket   uint32           // The number of bytes in a packet of audio data.
	MFramesPerPacket  uint32           // The number of frames in a packet of audio data.
	MBytesPerFrame    uint32           // The number of bytes from the start of one frame to the start of the next frame in an audio buffer.
	MChannelsPerFrame uint32           // The number of channels in each frame of audio data.
	MBitsPerChannel   uint32           // The number of bits for one audio sample.
	MReserved         uint32           // The amount to pad the structure to force an even 8-byte alignment.

}

// AudioStreamPacketDependencyDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDependencyDescription
type AudioStreamPacketDependencyDescription struct {
	MIsIndependentlyDecodable uint32
	MPreRollCount             uint32
	MFlags                    uint32
	MReserved                 uint32
}

// AudioStreamPacketDescription - A value that describes a packet in a buffer of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDescription
type AudioStreamPacketDescription struct {
	MStartOffset            int64  // The number of bytes from the start of the buffer to the beginning of the packet.
	MVariableFramesInPacket uint32 // The number of sample frames of data in the packet.
	MDataByteSize           uint32 // The number of bytes in the packet.

}

// AudioTimeStamp - A structure that represents a timestamp value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStamp
type AudioTimeStamp struct {
	MSampleTime    float64             // The absolute sample frame time.
	MHostTime      uint64              // The host machine’s time base (see `CoreAudio/HostTime.h`).
	MRateScalar    float64             // The ratio of actual host ticks per sample frame to the nominal host ticks per sample frame.
	MWordClockTime uint64              // The word clock time.
	MSMPTETime     SMPTETime           // The SMPTE time (see [SMPTETime](<https://developer.apple.com/documentation/CoreAudioTypes/SMPTETime>)).
	MFlags         AudioTimeStampFlags // A set of flags indicating which representations of the time are valid; see `Audio Time Stamp Flags` and `Audio Time Stamp Flag Combination Constant`.
	MReserved      uint32              // Pads the structure out to force an even 8-byte alignment.

}

// AudioValueRange - A structure that represents a continuous range of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueRange
type AudioValueRange struct {
	MMinimum float64 // The minimum value.
	MMaximum float64 // The maximum value.

}

// AudioValueTranslation - A structure that stores buffers to use in translation operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueTranslation
type AudioValueTranslation struct {
	MInputData      unsafe.Pointer // The buffer containing the data to be translated.
	MInputDataSize  uint32         // The number of bytes in the buffer pointed at by `mInputData`.
	MOutputData     unsafe.Pointer // The buffer to hold the result of the translation.
	MOutputDataSize uint32         // The number of bytes in the buffer pointed at by `mOutputData`.

}

// SMPTETime - A structure that defines an SMPTE time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETime
type SMPTETime struct {
	MSubframes       int16          // A subframe offset to the HH:MM:SS:FF time. You can use this field to position a time marker somewhere within the time span represented by a video frame, if necessary.
	MSubframeDivisor int16          // The number of subframes per video frame (typically 80).
	MCounter         uint32         // The total number of messages received. It takes 8 messages to carry a full SMPTE time code.
	MType            SMPTETimeType  // A SMPTE time type constant indicating the kind of SMPTE time used (see `SMPTE Timecode Types`).
	MFlags           SMPTETimeFlags // A set of flags that indicate the SMPTE state (see `SMPTE Time Flags`).
	MHours           int16          // The value of the hours portion of the SMPTE time.
	MMinutes         int16          // The value of the minutes portion of the SMPTE time.
	MSeconds         int16          // The value of the seconds portion of the SMPTE time.
	MFrames          int16          // The value of the frames portion of the SMPTE time.

}
