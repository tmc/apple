// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the requirements for a format reader, which represents a single media asset.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader
type MEFormatReader interface {
	objectivec.IObject

	// Loads the file info object with the properties of the media asset.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadFileInfo(completionHandler:)
	LoadFileInfoWithCompletionHandler(completionHandler MEFileInfoErrorHandler)

	// Loads the array of metadata items from the media asset.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadMetadata(completionHandler:)
	LoadMetadataWithCompletionHandler(completionHandler AVMetadataItemArrayErrorHandler)

	// Loads the array of track readers that represent the tracks in the media asset.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadTrackReaders(completionHandler:)
	LoadTrackReadersWithCompletionHandler(completionHandler ErrorHandler)
}

// MEFormatReaderObject wraps an existing Objective-C object that conforms to the MEFormatReader protocol.
type MEFormatReaderObject struct {
	objectivec.Object
}

func (o MEFormatReaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MEFormatReaderObjectFromID constructs a [MEFormatReaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MEFormatReaderObjectFromID(id objc.ID) MEFormatReaderObject {
	return MEFormatReaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Loads the file info object with the properties of the media asset.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadFileInfo(completionHandler:)
func (o MEFormatReaderObject) LoadFileInfoWithCompletionHandler(completionHandler MEFileInfoErrorHandler) {
	_block0, _ := NewMEFileInfoErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadFileInfoWithCompletionHandler:"), _block0)
}

// Loads the array of metadata items from the media asset.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadMetadata(completionHandler:)
func (o MEFormatReaderObject) LoadMetadataWithCompletionHandler(completionHandler AVMetadataItemArrayErrorHandler) {
	_block0, _ := NewAVMetadataItemArrayErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadMetadataWithCompletionHandler:"), _block0)
}

// Loads the array of track readers that represent the tracks in the media
// asset.
//
// completionHandler: The completion block to execute when the load operation finishes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/loadTrackReaders(completionHandler:)
func (o MEFormatReaderObject) LoadTrackReadersWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("loadTrackReadersWithCompletionHandler:"), _block0)
}

// Incorporates additional fragments that the file received after the last
// time the format reader parsed it.
//
// completionHandler: The completion block to execute when the parse operation finishes.
//
// # Discussion
//
// This method additional fragments of the media asset if they exist. Media
// asset formats that don’t support incremental fragments don’t need
// implement this method. Create the [MEFormatReader] object with the
// [MEFormatReaderInstantiationOptions] property
// [MEFormatReaderInstantiationOptions.AllowIncrementalFragmentParsing] set to
// true. This method does nothing if the value for
// [MEFileInfo.FragmentsStatus] is
// [MEFileInfo.FragmentsStatus.containsFragments]. Once this method returns an
// error, additional calls fail.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReader/parseAdditionalFragments(completionHandler:)
//
// [MEFileInfo.FragmentsStatus.containsFragments]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum/containsFragments
// [MEFileInfo.FragmentsStatus]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum
func (o MEFormatReaderObject) ParseAdditionalFragmentsWithCompletionHandler(completionHandler MEFormatReaderParseAdditionalFragmentsStatusErrorHandler) {
	_block0, _ := NewMEFormatReaderParseAdditionalFragmentsStatusErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("parseAdditionalFragmentsWithCompletionHandler:"), _block0)
}
