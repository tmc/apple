// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [MEByteSource] class.
var (
	_MEByteSourceClass     MEByteSourceClass
	_MEByteSourceClassOnce sync.Once
)

func getMEByteSourceClass() MEByteSourceClass {
	_MEByteSourceClassOnce.Do(func() {
		_MEByteSourceClass = MEByteSourceClass{class: objc.GetClass("MEByteSource")}
	})
	return _MEByteSourceClass
}

// GetMEByteSourceClass returns the class object for MEByteSource.
func GetMEByteSourceClass() MEByteSourceClass {
	return getMEByteSourceClass()
}

type MEByteSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEByteSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEByteSourceClass) Alloc() MEByteSource {
	rv := objc.Send[MEByteSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Provides read access to the data in a media asset file.
//
// # Overview
//
// Media Toolbox passes an [MEByteSource] instance for the media asset’s
// primary file when it initializes an [MEFormatReader] object. The format
// reader may call [MEByteSource.ByteSourceForRelatedFileNameError] to request
// additional byte sources for related files in the same directory as the
// primary file.
//
// # Inspecting a byte source
//
//   - [MEByteSource.FileName]: The name of the file for the byte source.
//   - [MEByteSource.FileLength]: The length of the byte source file.
//   - [MEByteSource.ContentType]: The format of the byte source file.
//   - [MEByteSource.RelatedFileNamesInSameDirectory]: An array of related file names in the parent directory of the byte source file.
//
// # Performing operations on a byte source
//
//   - [MEByteSource.AvailableLengthAtOffset]: Gets the number of available bytes from the offset within the byte source.
//   - [MEByteSource.ByteSourceForRelatedFileNameError]: Creates a new byte source for a related file.
//   - [MEByteSource.ReadDataOfLengthFromOffsetCompletionHandler]: Reads bytes from a byte source into a data object.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource
type MEByteSource struct {
	objectivec.Object
}

// MEByteSourceFromID constructs a [MEByteSource] from an objc.ID.
//
// Provides read access to the data in a media asset file.
func MEByteSourceFromID(id objc.ID) MEByteSource {
	return MEByteSource{objectivec.Object{ID: id}}
}

// NOTE: MEByteSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEByteSource] class.
//
// # Inspecting a byte source
//
//   - [IMEByteSource.FileName]: The name of the file for the byte source.
//   - [IMEByteSource.FileLength]: The length of the byte source file.
//   - [IMEByteSource.ContentType]: The format of the byte source file.
//   - [IMEByteSource.RelatedFileNamesInSameDirectory]: An array of related file names in the parent directory of the byte source file.
//
// # Performing operations on a byte source
//
//   - [IMEByteSource.AvailableLengthAtOffset]: Gets the number of available bytes from the offset within the byte source.
//   - [IMEByteSource.ByteSourceForRelatedFileNameError]: Creates a new byte source for a related file.
//   - [IMEByteSource.ReadDataOfLengthFromOffsetCompletionHandler]: Reads bytes from a byte source into a data object.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource
type IMEByteSource interface {
	objectivec.IObject

	// Topic: Inspecting a byte source

	// The name of the file for the byte source.
	FileName() string
	// The length of the byte source file.
	FileLength() int64
	// The format of the byte source file.
	ContentType() uniformtypeidentifiers.UTType
	// An array of related file names in the parent directory of the byte source file.
	RelatedFileNamesInSameDirectory() []string

	// Topic: Performing operations on a byte source

	// Gets the number of available bytes from the offset within the byte source.
	AvailableLengthAtOffset(offset int64) int64
	// Creates a new byte source for a related file.
	ByteSourceForRelatedFileNameError(fileName string) (IMEByteSource, error)
	// Reads bytes from a byte source into a data object.
	ReadDataOfLengthFromOffsetCompletionHandler(length uintptr, offset int64, completionHandler DataErrorHandler)
}

// Init initializes the instance.
func (m MEByteSource) Init() MEByteSource {
	rv := objc.Send[MEByteSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEByteSource) Autorelease() MEByteSource {
	rv := objc.Send[MEByteSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEByteSource creates a new MEByteSource instance.
func NewMEByteSource() MEByteSource {
	class := getMEByteSourceClass()
	rv := objc.Send[MEByteSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets the number of available bytes from the offset within the byte source.
//
// offset: The offset in bytes from the beginning of the byte source.
//
// # Return Value
//
// An integer that specifies the number of available bytes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/availableLength(at:)
func (m MEByteSource) AvailableLengthAtOffset(offset int64) int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("availableLengthAtOffset:"), offset)
	return rv
}

// Creates a new byte source for a related file.
//
// fileName: The related file name that exists in the byte source’s parent directory.
//
// # Return Value
//
// A byte source.
//
// # Discussion
//
// Requests creation of a new [MEByteSource] for a file related to the
// receiving [MEByteSource]. Only file names returned by the
// [MEByteSource.RelatedFileNamesInSameDirectory] method may be accessed.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/byteSourceForRelatedFileName(_:)
func (m MEByteSource) ByteSourceForRelatedFileNameError(fileName string) (IMEByteSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("byteSourceForRelatedFileName:error:"), objc.String(fileName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MEByteSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MEByteSourceFromID(rv), nil

}

// Reads bytes from a byte source into a data object.
//
// length: The number of bytes to read.
//
// offset: The relative offset in bytes from the beginning of the file from which to
// start reading.
//
// completionHandler: The completion block to execute when the read operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/read(length:from:completionHandler:)
func (m MEByteSource) ReadDataOfLengthFromOffsetCompletionHandler(length uintptr, offset int64, completionHandler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](m.ID, objc.Sel("readDataOfLength:fromOffset:completionHandler:"), length, offset, _block2)
}

// The name of the file for the byte source.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/fileName
func (m MEByteSource) FileName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fileName"))
	return foundation.NSStringFromID(rv).String()
}

// The length of the byte source file.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/fileLength
func (m MEByteSource) FileLength() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("fileLength"))
	return rv
}

// The format of the byte source file.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/contentType
func (m MEByteSource) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(objc.ID(rv))
}

// An array of related file names in the parent directory of the byte source
// file.
//
// # Discussion
//
// The array of related files within the [MEByteSource]’s parent directory
// that are accessible to the [MEByteSource]. Only the relative file names are
// returned, not the paths. Only files with file extensions listed in the
// [kMEFormatReaderFileNameExtensionArrayKey] array in the format reader
// property list will be returned. If no related files are available, returns
// an empty array.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEByteSource/relatedFileNamesInSameDirectory
//
// [kMEFormatReaderFileNameExtensionArrayKey]: https://developer.apple.com/documentation/MediaExtension/kMEFormatReaderFileNameExtensionArrayKey
func (m MEByteSource) RelatedFileNamesInSameDirectory() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("relatedFileNamesInSameDirectory"))
	return objc.ConvertSliceToStrings(rv)
}

// ReadDataOfLengthFromOffset is a synchronous wrapper around [MEByteSource.ReadDataOfLengthFromOffsetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MEByteSource) ReadDataOfLengthFromOffset(ctx context.Context, length uintptr, offset int64) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	m.ReadDataOfLengthFromOffsetCompletionHandler(length, offset, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
