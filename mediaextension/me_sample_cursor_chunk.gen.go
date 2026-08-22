// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MESampleCursorChunk] class.
var (
	_MESampleCursorChunkClass     MESampleCursorChunkClass
	_MESampleCursorChunkClassOnce sync.Once
)

func getMESampleCursorChunkClass() MESampleCursorChunkClass {
	_MESampleCursorChunkClassOnce.Do(func() {
		_MESampleCursorChunkClass = MESampleCursorChunkClass{class: objc.GetClass("MESampleCursorChunk")}
	})
	return _MESampleCursorChunkClass
}

// GetMESampleCursorChunkClass returns the class object for MESampleCursorChunk.
func GetMESampleCursorChunkClass() MESampleCursorChunkClass {
	return getMESampleCursorChunkClass()
}

type MESampleCursorChunkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MESampleCursorChunkClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MESampleCursorChunkClass) Alloc() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the chunk of media at the
// location of a sample.
//
// # Overview
//
// The [ChunkDetailsReturningError] method returns an instance of this class.
//
// # Creating a sample cursor chunk
//
//   - [MESampleCursorChunk.InitWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk]: Creates a new sample cursor chunk with byte source and chunk data that you provide.
//
// # Inspecting a chunk
//
//   - [MESampleCursorChunk.ByteSource]: The byte source to use to read the data for the sample.
//   - [MESampleCursorChunk.ChunkStorageRange]: The offset location and length of the sample’s chunk within the byte source.
//   - [MESampleCursorChunk.ChunkInfo]: An object that provides details about the chunk in the media.
//   - [MESampleCursorChunk.SampleIndexWithinChunk]: The offset index of the sample within the chunk, in samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk
type MESampleCursorChunk struct {
	objectivec.Object
}

// MESampleCursorChunkFromID constructs a [MESampleCursorChunk] from an objc.ID.
//
// An object that provides information about the chunk of media at the
// location of a sample.
func MESampleCursorChunkFromID(id objc.ID) MESampleCursorChunk {
	return MESampleCursorChunk{objectivec.Object{ID: id}}
}

// NOTE: MESampleCursorChunk adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MESampleCursorChunk] class.
//
// # Creating a sample cursor chunk
//
//   - [IMESampleCursorChunk.InitWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk]: Creates a new sample cursor chunk with byte source and chunk data that you provide.
//
// # Inspecting a chunk
//
//   - [IMESampleCursorChunk.ByteSource]: The byte source to use to read the data for the sample.
//   - [IMESampleCursorChunk.ChunkStorageRange]: The offset location and length of the sample’s chunk within the byte source.
//   - [IMESampleCursorChunk.ChunkInfo]: An object that provides details about the chunk in the media.
//   - [IMESampleCursorChunk.SampleIndexWithinChunk]: The offset index of the sample within the chunk, in samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk
type IMESampleCursorChunk interface {
	objectivec.IObject

	// Topic: Creating a sample cursor chunk

	// Creates a new sample cursor chunk with byte source and chunk data that you provide.
	InitWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk(byteSource IMEByteSource, chunkStorageRange avfoundation.AVSampleCursorStorageRange, chunkInfo avfoundation.AVSampleCursorChunkInfo, sampleIndexWithinChunk corefoundation.CFIndex) MESampleCursorChunk

	// Topic: Inspecting a chunk

	// The byte source to use to read the data for the sample.
	ByteSource() IMEByteSource
	// The offset location and length of the sample’s chunk within the byte source.
	ChunkStorageRange() avfoundation.AVSampleCursorStorageRange
	// An object that provides details about the chunk in the media.
	ChunkInfo() avfoundation.AVSampleCursorChunkInfo
	// The offset index of the sample within the chunk, in samples.
	SampleIndexWithinChunk() corefoundation.CFIndex
}

// Init initializes the instance.
func (m MESampleCursorChunk) Init() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MESampleCursorChunk) Autorelease() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMESampleCursorChunk creates a new MESampleCursorChunk instance.
func NewMESampleCursorChunk() MESampleCursorChunk {
	class := getMESampleCursorChunkClass()
	rv := objc.Send[MESampleCursorChunk](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new sample cursor chunk with byte source and chunk data that you
// provide.
//
// byteSource: The byte source to use to read the data for the sample.
//
// chunkStorageRange: The offset location and length of the sample’s chunk within the byte
// source.
//
// chunkInfo: An object that provides details about the chunk in the media.
//
// sampleIndexWithinChunk: The offset index of the sample within the chunk, in samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/init(byteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:)
func NewMESampleCursorChunkWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk(byteSource IMEByteSource, chunkStorageRange avfoundation.AVSampleCursorStorageRange, chunkInfo avfoundation.AVSampleCursorChunkInfo, sampleIndexWithinChunk corefoundation.CFIndex) MESampleCursorChunk {
	instance := getMESampleCursorChunkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithByteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:"), byteSource, chunkStorageRange, chunkInfo, sampleIndexWithinChunk)
	return MESampleCursorChunkFromID(rv)
}

// Creates a new sample cursor chunk with byte source and chunk data that you
// provide.
//
// byteSource: The byte source to use to read the data for the sample.
//
// chunkStorageRange: The offset location and length of the sample’s chunk within the byte
// source.
//
// chunkInfo: An object that provides details about the chunk in the media.
//
// sampleIndexWithinChunk: The offset index of the sample within the chunk, in samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/init(byteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:)
func (m MESampleCursorChunk) InitWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk(byteSource IMEByteSource, chunkStorageRange avfoundation.AVSampleCursorStorageRange, chunkInfo avfoundation.AVSampleCursorChunkInfo, sampleIndexWithinChunk corefoundation.CFIndex) MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](m.ID, objc.Sel("initWithByteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:"), byteSource, chunkStorageRange, chunkInfo, sampleIndexWithinChunk)
	return rv
}

// The byte source to use to read the data for the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/byteSource
func (m MESampleCursorChunk) ByteSource() IMEByteSource {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("byteSource"))
	return MEByteSourceFromID(objc.ID(rv))
}

// The offset location and length of the sample’s chunk within the byte
// source.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/chunkStorageRange
func (m MESampleCursorChunk) ChunkStorageRange() avfoundation.AVSampleCursorStorageRange {
	rv := objc.Send[avfoundation.AVSampleCursorStorageRange](m.ID, objc.Sel("chunkStorageRange"))
	return avfoundation.AVSampleCursorStorageRange(rv)
}

// An object that provides details about the chunk in the media.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/chunkInfo
func (m MESampleCursorChunk) ChunkInfo() avfoundation.AVSampleCursorChunkInfo {
	rv := objc.Send[avfoundation.AVSampleCursorChunkInfo](m.ID, objc.Sel("chunkInfo"))
	return avfoundation.AVSampleCursorChunkInfo(rv)
}

// The offset index of the sample within the chunk, in samples.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/sampleIndexWithinChunk
func (m MESampleCursorChunk) SampleIndexWithinChunk() corefoundation.CFIndex {
	rv := objc.Send[corefoundation.CFIndex](m.ID, objc.Sel("sampleIndexWithinChunk"))
	return corefoundation.CFIndex(rv)
}
