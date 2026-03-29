// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioConverter] class.
var (
	_AVAudioConverterClass     AVAudioConverterClass
	_AVAudioConverterClassOnce sync.Once
)

func getAVAudioConverterClass() AVAudioConverterClass {
	_AVAudioConverterClassOnce.Do(func() {
		_AVAudioConverterClass = AVAudioConverterClass{class: objc.GetClass("AVAudioConverter")}
	})
	return _AVAudioConverterClass
}

// GetAVAudioConverterClass returns the class object for AVAudioConverter.
func GetAVAudioConverterClass() AVAudioConverterClass {
	return getAVAudioConverterClass()
}

type AVAudioConverterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioConverterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioConverterClass) Alloc() AVAudioConverter {
	rv := objc.Send[AVAudioConverter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that converts streams of audio between formats.
//
// # Overview
// 
// The audio converter class transforms audio between file formats and audio
// encodings.
// 
// Supported transformations include:
// 
// - PCM float, integer, or bit depth conversions - PCM sample rate conversion
// - PCM interleaving and deinterleaving - Encoding PCM to compressed formats
// - Decoding compressed formats to PCM
// 
// A single audio converter instance may perform more than one of the above
// transformations.
//
// # Creating an Audio Converter
//
//   - [AVAudioConverter.InitFromFormatToFormat]: Creates an audio converter object from the specified input and output formats.
//
// # Converting Audio Formats
//
//   - [AVAudioConverter.ConvertToBufferErrorWithInputFromBlock]: Performs a conversion between audio formats, if the system supports it.
//   - [AVAudioConverter.ConvertToBufferFromBufferError]: Performs a basic conversion between audio formats that doesn’t involve converting codecs or sample rates.
//
// # Resetting an Audio Converter
//
//   - [AVAudioConverter.Reset]: Resets the converter so you can convert a new audio stream.
//
// # Getting Audio Converter Properties
//
//   - [AVAudioConverter.ChannelMap]: An array of integers that indicates which input to derive each output from.
//   - [AVAudioConverter.SetChannelMap]
//   - [AVAudioConverter.Dither]: A Boolean value that indicates whether dither is on.
//   - [AVAudioConverter.SetDither]
//   - [AVAudioConverter.Downmix]: A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//   - [AVAudioConverter.SetDownmix]
//   - [AVAudioConverter.InputFormat]: The format of the input audio stream.
//   - [AVAudioConverter.OutputFormat]: The format of the output audio stream.
//   - [AVAudioConverter.MagicCookie]: An object that contains metadata for encoders and decoders.
//   - [AVAudioConverter.SetMagicCookie]
//   - [AVAudioConverter.MaximumOutputPacketSize]: The maximum size of an output packet, in bytes.
//
// # Getting Bit Rate Properties
//
//   - [AVAudioConverter.ApplicableEncodeBitRates]: An array of bit rates the framework applies during encoding according to the current formats and settings.
//   - [AVAudioConverter.AvailableEncodeBitRates]: An array of all bit rates the codec provides when encoding.
//   - [AVAudioConverter.AvailableEncodeChannelLayoutTags]: An array of all output channel layout tags the codec provides when encoding.
//   - [AVAudioConverter.BitRate]: The bit rate, in bits per second.
//   - [AVAudioConverter.SetBitRate]
//   - [AVAudioConverter.BitRateStrategy]: A key value constant the framework uses during encoding.
//   - [AVAudioConverter.SetBitRateStrategy]
//
// # Getting Sample Rate Properties
//
//   - [AVAudioConverter.SampleRateConverterQuality]: A sample rate converter algorithm key value.
//   - [AVAudioConverter.SetSampleRateConverterQuality]
//   - [AVAudioConverter.SampleRateConverterAlgorithm]: The priming method the sample rate converter or decoder uses.
//   - [AVAudioConverter.SetSampleRateConverterAlgorithm]
//   - [AVAudioConverter.ApplicableEncodeSampleRates]: An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//   - [AVAudioConverter.AvailableEncodeSampleRates]: An array of all output sample rates the codec provides when encoding.
//
// # Getting Priming Information
//
//   - [AVAudioConverter.PrimeInfo]: The number of priming frames the converter uses.
//   - [AVAudioConverter.SetPrimeInfo]
//   - [AVAudioConverter.PrimeMethod]: The priming method the sample rate converter or decoder uses.
//   - [AVAudioConverter.SetPrimeMethod]
//
// # Managing packet dependencies
//
//   - [AVAudioConverter.AudioSyncPacketFrequency]
//   - [AVAudioConverter.SetAudioSyncPacketFrequency]
//   - [AVAudioConverter.ContentSource]
//   - [AVAudioConverter.SetContentSource]
//   - [AVAudioConverter.DynamicRangeControlConfiguration]
//   - [AVAudioConverter.SetDynamicRangeControlConfiguration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter
type AVAudioConverter struct {
	objectivec.Object
}

// AVAudioConverterFromID constructs a [AVAudioConverter] from an objc.ID.
//
// An object that converts streams of audio between formats.
func AVAudioConverterFromID(id objc.ID) AVAudioConverter {
	return AVAudioConverter{objectivec.Object{ID: id}}
}
// NOTE: AVAudioConverter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioConverter] class.
//
// # Creating an Audio Converter
//
//   - [IAVAudioConverter.InitFromFormatToFormat]: Creates an audio converter object from the specified input and output formats.
//
// # Converting Audio Formats
//
//   - [IAVAudioConverter.ConvertToBufferErrorWithInputFromBlock]: Performs a conversion between audio formats, if the system supports it.
//   - [IAVAudioConverter.ConvertToBufferFromBufferError]: Performs a basic conversion between audio formats that doesn’t involve converting codecs or sample rates.
//
// # Resetting an Audio Converter
//
//   - [IAVAudioConverter.Reset]: Resets the converter so you can convert a new audio stream.
//
// # Getting Audio Converter Properties
//
//   - [IAVAudioConverter.ChannelMap]: An array of integers that indicates which input to derive each output from.
//   - [IAVAudioConverter.SetChannelMap]
//   - [IAVAudioConverter.Dither]: A Boolean value that indicates whether dither is on.
//   - [IAVAudioConverter.SetDither]
//   - [IAVAudioConverter.Downmix]: A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//   - [IAVAudioConverter.SetDownmix]
//   - [IAVAudioConverter.InputFormat]: The format of the input audio stream.
//   - [IAVAudioConverter.OutputFormat]: The format of the output audio stream.
//   - [IAVAudioConverter.MagicCookie]: An object that contains metadata for encoders and decoders.
//   - [IAVAudioConverter.SetMagicCookie]
//   - [IAVAudioConverter.MaximumOutputPacketSize]: The maximum size of an output packet, in bytes.
//
// # Getting Bit Rate Properties
//
//   - [IAVAudioConverter.ApplicableEncodeBitRates]: An array of bit rates the framework applies during encoding according to the current formats and settings.
//   - [IAVAudioConverter.AvailableEncodeBitRates]: An array of all bit rates the codec provides when encoding.
//   - [IAVAudioConverter.AvailableEncodeChannelLayoutTags]: An array of all output channel layout tags the codec provides when encoding.
//   - [IAVAudioConverter.BitRate]: The bit rate, in bits per second.
//   - [IAVAudioConverter.SetBitRate]
//   - [IAVAudioConverter.BitRateStrategy]: A key value constant the framework uses during encoding.
//   - [IAVAudioConverter.SetBitRateStrategy]
//
// # Getting Sample Rate Properties
//
//   - [IAVAudioConverter.SampleRateConverterQuality]: A sample rate converter algorithm key value.
//   - [IAVAudioConverter.SetSampleRateConverterQuality]
//   - [IAVAudioConverter.SampleRateConverterAlgorithm]: The priming method the sample rate converter or decoder uses.
//   - [IAVAudioConverter.SetSampleRateConverterAlgorithm]
//   - [IAVAudioConverter.ApplicableEncodeSampleRates]: An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//   - [IAVAudioConverter.AvailableEncodeSampleRates]: An array of all output sample rates the codec provides when encoding.
//
// # Getting Priming Information
//
//   - [IAVAudioConverter.PrimeInfo]: The number of priming frames the converter uses.
//   - [IAVAudioConverter.SetPrimeInfo]
//   - [IAVAudioConverter.PrimeMethod]: The priming method the sample rate converter or decoder uses.
//   - [IAVAudioConverter.SetPrimeMethod]
//
// # Managing packet dependencies
//
//   - [IAVAudioConverter.AudioSyncPacketFrequency]
//   - [IAVAudioConverter.SetAudioSyncPacketFrequency]
//   - [IAVAudioConverter.ContentSource]
//   - [IAVAudioConverter.SetContentSource]
//   - [IAVAudioConverter.DynamicRangeControlConfiguration]
//   - [IAVAudioConverter.SetDynamicRangeControlConfiguration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter
type IAVAudioConverter interface {
	objectivec.IObject

	// Topic: Creating an Audio Converter

	// Creates an audio converter object from the specified input and output formats.
	InitFromFormatToFormat(fromFormat IAVAudioFormat, toFormat IAVAudioFormat) AVAudioConverter

	// Topic: Converting Audio Formats

	// Performs a conversion between audio formats, if the system supports it.
	ConvertToBufferErrorWithInputFromBlock(outputBuffer IAVAudioBuffer, outError foundation.INSError, inputBlock AVAudioConverterInputBlock) AVAudioConverterOutputStatus
	// Performs a basic conversion between audio formats that doesn’t involve converting codecs or sample rates.
	ConvertToBufferFromBufferError(outputBuffer IAVAudioPCMBuffer, inputBuffer IAVAudioPCMBuffer) (bool, error)

	// Topic: Resetting an Audio Converter

	// Resets the converter so you can convert a new audio stream.
	Reset()

	// Topic: Getting Audio Converter Properties

	// An array of integers that indicates which input to derive each output from.
	ChannelMap() []foundation.NSNumber
	SetChannelMap(value []foundation.NSNumber)
	// A Boolean value that indicates whether dither is on.
	Dither() bool
	SetDither(value bool)
	// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
	Downmix() bool
	SetDownmix(value bool)
	// The format of the input audio stream.
	InputFormat() IAVAudioFormat
	// The format of the output audio stream.
	OutputFormat() IAVAudioFormat
	// An object that contains metadata for encoders and decoders.
	MagicCookie() foundation.INSData
	SetMagicCookie(value foundation.INSData)
	// The maximum size of an output packet, in bytes.
	MaximumOutputPacketSize() int

	// Topic: Getting Bit Rate Properties

	// An array of bit rates the framework applies during encoding according to the current formats and settings.
	ApplicableEncodeBitRates() []foundation.NSNumber
	// An array of all bit rates the codec provides when encoding.
	AvailableEncodeBitRates() []foundation.NSNumber
	// An array of all output channel layout tags the codec provides when encoding.
	AvailableEncodeChannelLayoutTags() []foundation.NSNumber
	// The bit rate, in bits per second.
	BitRate() int
	SetBitRate(value int)
	// A key value constant the framework uses during encoding.
	BitRateStrategy() string
	SetBitRateStrategy(value string)

	// Topic: Getting Sample Rate Properties

	// A sample rate converter algorithm key value.
	SampleRateConverterQuality() int
	SetSampleRateConverterQuality(value int)
	// The priming method the sample rate converter or decoder uses.
	SampleRateConverterAlgorithm() string
	SetSampleRateConverterAlgorithm(value string)
	// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
	ApplicableEncodeSampleRates() []foundation.NSNumber
	// An array of all output sample rates the codec provides when encoding.
	AvailableEncodeSampleRates() []foundation.NSNumber

	// Topic: Getting Priming Information

	// The number of priming frames the converter uses.
	PrimeInfo() AVAudioConverterPrimeInfo
	SetPrimeInfo(value AVAudioConverterPrimeInfo)
	// The priming method the sample rate converter or decoder uses.
	PrimeMethod() AVAudioConverterPrimeMethod
	SetPrimeMethod(value AVAudioConverterPrimeMethod)

	// Topic: Managing packet dependencies

	AudioSyncPacketFrequency() int
	SetAudioSyncPacketFrequency(value int)
	ContentSource() AVAudioContentSource
	SetContentSource(value AVAudioContentSource)
	DynamicRangeControlConfiguration() AVAudioDynamicRangeControlConfiguration
	SetDynamicRangeControlConfiguration(value AVAudioDynamicRangeControlConfiguration)
}

// Init initializes the instance.
func (a AVAudioConverter) Init() AVAudioConverter {
	rv := objc.Send[AVAudioConverter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioConverter) Autorelease() AVAudioConverter {
	rv := objc.Send[AVAudioConverter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioConverter creates a new AVAudioConverter instance.
func NewAVAudioConverter() AVAudioConverter {
	class := getAVAudioConverterClass()
	rv := objc.Send[AVAudioConverter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio converter object from the specified input and output
// formats.
//
// fromFormat: The input audio format.
//
// toFormat: The audio format to convert to.
//
// # Return Value
// 
// An [AVAudioConverter] instance, or `nil` if the format conversion isn’t
// possible.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/init(from:to:)
func NewAudioConverterFromFormatToFormat(fromFormat IAVAudioFormat, toFormat IAVAudioFormat) AVAudioConverter {
	instance := getAVAudioConverterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromFormat:toFormat:"), fromFormat, toFormat)
	return AVAudioConverterFromID(rv)
}

// Creates an audio converter object from the specified input and output
// formats.
//
// fromFormat: The input audio format.
//
// toFormat: The audio format to convert to.
//
// # Return Value
// 
// An [AVAudioConverter] instance, or `nil` if the format conversion isn’t
// possible.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/init(from:to:)
func (a AVAudioConverter) InitFromFormatToFormat(fromFormat IAVAudioFormat, toFormat IAVAudioFormat) AVAudioConverter {
	rv := objc.Send[AVAudioConverter](a.ID, objc.Sel("initFromFormat:toFormat:"), fromFormat, toFormat)
	return rv
}
// Performs a conversion between audio formats, if the system supports it.
//
// outputBuffer: The output audio buffer.
//
// outError: The error if the conversion fails.
//
// inputBlock: A block the framework calls to get input data.
//
// # Return Value
// 
// An [AVAudioConverterOutputStatus] type that indicates the conversion
// status.
//
// [AVAudioConverterOutputStatus]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus
//
// # Discussion
// 
// The method attempts to fill the buffer to its capacity. On return, the
// buffer’s length indicates the number of sample frames the framework
// successfully converts.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:error:withInputFrom:)
func (a AVAudioConverter) ConvertToBufferErrorWithInputFromBlock(outputBuffer IAVAudioBuffer, outError foundation.INSError, inputBlock AVAudioConverterInputBlock) AVAudioConverterOutputStatus {
	rv := objc.Send[AVAudioConverterOutputStatus](a.ID, objc.Sel("convertToBuffer:error:withInputFromBlock:"), outputBuffer, outError, inputBlock)
	return AVAudioConverterOutputStatus(rv)
}
// Performs a basic conversion between audio formats that doesn’t involve
// converting codecs or sample rates.
//
// outputBuffer: The output audio buffer.
//
// inputBuffer: The input audio buffer.
//
// # Discussion
// 
// The output buffer’s [FrameCapacity] value needs to be at least at large
// as the [FrameLength] value of the `inputBuffer`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:from:)
func (a AVAudioConverter) ConvertToBufferFromBufferError(outputBuffer IAVAudioPCMBuffer, inputBuffer IAVAudioPCMBuffer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("convertToBuffer:fromBuffer:error:"), outputBuffer, inputBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("convertToBuffer:fromBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Resets the converter so you can convert a new audio stream.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/reset()
func (a AVAudioConverter) Reset() {
	objc.Send[objc.ID](a.ID, objc.Sel("reset"))
}

// An array of integers that indicates which input to derive each output from.
//
// # Discussion
// 
// The array size equals the number of output channels. Each element’s value
// is the input channel number, starting with zero, that the framework copies
// to that output.
// 
// A negative value means that the output channel doesn’t have a source and
// is silent.
// 
// Setting a channel map overrides channel mapping due to any channel layouts
// in the input and output formats that you supply.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/channelMap
func (a AVAudioConverter) ChannelMap() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelMap"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (a AVAudioConverter) SetChannelMap(value []foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setChannelMap:"), objectivec.IObjectSliceToNSArray(value))
}
// A Boolean value that indicates whether dither is on.
//
// # Discussion
// 
// This property defaults to `false`. When `true`, the framework determines
// whether dithering makes sense for the formats and settings.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dither
func (a AVAudioConverter) Dither() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("dither"))
	return rv
}
func (a AVAudioConverter) SetDither(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setDither:"), value)
}
// A Boolean value that indicates whether the framework mixes the channels
// instead of remapping.
//
// # Discussion
// 
// This property defaults to `false`, indicating that the framework remaps the
// channels. When `true`, and channel remapping is necessary, the framework
// mixes the channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/downmix
func (a AVAudioConverter) Downmix() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("downmix"))
	return rv
}
func (a AVAudioConverter) SetDownmix(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setDownmix:"), value)
}
// The format of the input audio stream.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/inputFormat
func (a AVAudioConverter) InputFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}
// The format of the output audio stream.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/outputFormat
func (a AVAudioConverter) OutputFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}
// An object that contains metadata for encoders and decoders.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/magicCookie
func (a AVAudioConverter) MagicCookie() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("magicCookie"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a AVAudioConverter) SetMagicCookie(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setMagicCookie:"), value)
}
// The maximum size of an output packet, in bytes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/maximumOutputPacketSize
func (a AVAudioConverter) MaximumOutputPacketSize() int {
	rv := objc.Send[int](a.ID, objc.Sel("maximumOutputPacketSize"))
	return rv
}
// An array of bit rates the framework applies during encoding according to
// the current formats and settings.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/applicableEncodeBitRates
func (a AVAudioConverter) ApplicableEncodeBitRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("applicableEncodeBitRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// An array of all bit rates the codec provides when encoding.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeBitRates
func (a AVAudioConverter) AvailableEncodeBitRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableEncodeBitRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// An array of all output channel layout tags the codec provides when
// encoding.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeChannelLayoutTags
func (a AVAudioConverter) AvailableEncodeChannelLayoutTags() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableEncodeChannelLayoutTags"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The bit rate, in bits per second.
//
// # Discussion
// 
// This value only applies when encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRate
func (a AVAudioConverter) BitRate() int {
	rv := objc.Send[int](a.ID, objc.Sel("bitRate"))
	return rv
}
func (a AVAudioConverter) SetBitRate(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setBitRate:"), value)
}
// A key value constant the framework uses during encoding.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding. For information about
// possible values, see [AVEncoderBitRateStrategyKey].
//
// [AVEncoderBitRateStrategyKey]: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitRateStrategyKey
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRateStrategy
func (a AVAudioConverter) BitRateStrategy() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bitRateStrategy"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioConverter) SetBitRateStrategy(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setBitRateStrategy:"), objc.String(value))
}
// A sample rate converter algorithm key value.
//
// # Discussion
// 
// For information about possible key values, see
// [AVSampleRateConverterAlgorithmKey].
//
// [AVSampleRateConverterAlgorithmKey]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAlgorithmKey
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterQuality
func (a AVAudioConverter) SampleRateConverterQuality() int {
	rv := objc.Send[int](a.ID, objc.Sel("sampleRateConverterQuality"))
	return rv
}
func (a AVAudioConverter) SetSampleRateConverterQuality(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setSampleRateConverterQuality:"), value)
}
// The priming method the sample rate converter or decoder uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterAlgorithm
func (a AVAudioConverter) SampleRateConverterAlgorithm() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sampleRateConverterAlgorithm"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioConverter) SetSampleRateConverterAlgorithm(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSampleRateConverterAlgorithm:"), objc.String(value))
}
// An array of output sample rates that the converter applies according to the
// current formats and settings, when encoding.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/applicableEncodeSampleRates
func (a AVAudioConverter) ApplicableEncodeSampleRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("applicableEncodeSampleRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// An array of all output sample rates the codec provides when encoding.
//
// # Discussion
// 
// This property returns `nil` if you’re not encoding.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeSampleRates
func (a AVAudioConverter) AvailableEncodeSampleRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableEncodeSampleRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The number of priming frames the converter uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeInfo
func (a AVAudioConverter) PrimeInfo() AVAudioConverterPrimeInfo {
	rv := objc.Send[AVAudioConverterPrimeInfo](a.ID, objc.Sel("primeInfo"))
	return AVAudioConverterPrimeInfo(rv)
}
func (a AVAudioConverter) SetPrimeInfo(value AVAudioConverterPrimeInfo) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimeInfo:"), value)
}
// The priming method the sample rate converter or decoder uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeMethod
func (a AVAudioConverter) PrimeMethod() AVAudioConverterPrimeMethod {
	rv := objc.Send[AVAudioConverterPrimeMethod](a.ID, objc.Sel("primeMethod"))
	return AVAudioConverterPrimeMethod(rv)
}
func (a AVAudioConverter) SetPrimeMethod(value AVAudioConverterPrimeMethod) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimeMethod:"), value)
}
//
// # Discussion
// 
// Number of packets between consecutive sync packets.
// 
// A sync packet is an independently-decodable packet that completely
// refreshes the decoder without needing to decode other packets. When
// compressing to a format which supports it (such as APAC), the audio sync
// packet frequency indicates the distance in packets between two sync
// packets, with non-sync packets between. This is useful to set when saving
// compressed packets to a file and efficient random access is desired. Note:
// Separating sync packets by at least one second of encoded audio (e.g. 75
// packets) is recommended.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/audioSyncPacketFrequency
func (a AVAudioConverter) AudioSyncPacketFrequency() int {
	rv := objc.Send[int](a.ID, objc.Sel("audioSyncPacketFrequency"))
	return rv
}
func (a AVAudioConverter) SetAudioSyncPacketFrequency(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioSyncPacketFrequency:"), value)
}
//
// # Discussion
// 
// Index to select a pre-defined content source type that describes the
// content type and how it was generated. Note: This is only supported when
// compressing audio to formats which support it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/contentSource
func (a AVAudioConverter) ContentSource() AVAudioContentSource {
	rv := objc.Send[AVAudioContentSource](a.ID, objc.Sel("contentSource"))
	return AVAudioContentSource(rv)
}
func (a AVAudioConverter) SetContentSource(value AVAudioContentSource) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentSource:"), value)
}
//
// # Discussion
// 
// Encoder Dynamic Range Control (DRC) configuration.
// 
// When supported by the encoder, this property controls which configuration
// is applied when a bitstream is generated. Note: This is only supported when
// compressing audio to formats which support it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dynamicRangeControlConfiguration
func (a AVAudioConverter) DynamicRangeControlConfiguration() AVAudioDynamicRangeControlConfiguration {
	rv := objc.Send[AVAudioDynamicRangeControlConfiguration](a.ID, objc.Sel("dynamicRangeControlConfiguration"))
	return AVAudioDynamicRangeControlConfiguration(rv)
}
func (a AVAudioConverter) SetDynamicRangeControlConfiguration(value AVAudioDynamicRangeControlConfiguration) {
	objc.Send[struct{}](a.ID, objc.Sel("setDynamicRangeControlConfiguration:"), value)
}

