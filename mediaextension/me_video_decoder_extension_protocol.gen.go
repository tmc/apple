// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a factory to create new video decoders for a codec type that the extension implements.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderExtension
type MEVideoDecoderExtension interface {
	objectivec.IObject

	// Creates a new video decoder that matches the codec type, format description, decoder specifications, and pixel buffer manager that you specify.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderExtension/makeVideoDecoder(codecType:videoFormatDescription:videoDecoderSpecifications:pixelBufferManager:)
	VideoDecoderWithCodecTypeVideoFormatDescriptionVideoDecoderSpecificationsExtensionDecoderPixelBufferManagerError(codecType coremedia.CMVideoCodecType, videoFormatDescription coremedia.CMVideoFormatDescriptionRef, videoDecoderSpecifications foundation.INSDictionary, extensionDecoderPixelBufferManager IMEVideoDecoderPixelBufferManager) (MEVideoDecoder, error)
}

// MEVideoDecoderExtensionObject wraps an existing Objective-C object that conforms to the MEVideoDecoderExtension protocol.
type MEVideoDecoderExtensionObject struct {
	objectivec.Object
}

func (o MEVideoDecoderExtensionObject) BaseObject() objectivec.Object {
	return o.Object
}

// MEVideoDecoderExtensionObjectFromID constructs a [MEVideoDecoderExtensionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MEVideoDecoderExtensionObjectFromID(id objc.ID) MEVideoDecoderExtensionObject {
	return MEVideoDecoderExtensionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates a new video decoder that matches the codec type, format
// description, decoder specifications, and pixel buffer manager that you
// specify.
//
// codecType: The codec type for the requested decoder.
//
// videoFormatDescription: An object that describes the video data.
//
// videoDecoderSpecifications: A dictionary that contains video decoder specification values, which may be
// empty. See [Decompression Properties] for a list of keys to use.
//
// extensionDecoderPixelBufferManager: A pixel buffer manager.
//
// # Return Value
//
// A new [MEVideoDecoder].
//
// # Discussion
//
// The `videoDecoderSpecifications` parameter accepts the following keys:
//
// - [kVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder] -
// [kVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder] -
// [kVTVideoDecoderSpecification_RequiredDecoderGPURegistryID] -
// [kVTVideoDecoderSpecification_PreferredDecoderGPURegistryID]
//
// If the parameter values aren’t compatible with the video decoder, this
// method fails with the error [MEError.Code.unsupportedFeature].
//
// The video decoder needs to retain the pixel buffer manager and use it to
// allocate and configure output pixel buffers.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderExtension/makeVideoDecoder(codecType:videoFormatDescription:videoDecoderSpecifications:pixelBufferManager:)
//
// [Decompression Properties]: https://developer.apple.com/documentation/VideoToolbox/decompression-properties
// [MEError.Code.unsupportedFeature]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/unsupportedFeature
// [kVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder]: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder
// [kVTVideoDecoderSpecification_PreferredDecoderGPURegistryID]: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_PreferredDecoderGPURegistryID
// [kVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder]: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder
// [kVTVideoDecoderSpecification_RequiredDecoderGPURegistryID]: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_RequiredDecoderGPURegistryID
func (o MEVideoDecoderExtensionObject) VideoDecoderWithCodecTypeVideoFormatDescriptionVideoDecoderSpecificationsExtensionDecoderPixelBufferManagerError(codecType coremedia.CMVideoCodecType, videoFormatDescription coremedia.CMVideoFormatDescriptionRef, videoDecoderSpecifications foundation.INSDictionary, extensionDecoderPixelBufferManager IMEVideoDecoderPixelBufferManager) (MEVideoDecoder, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("videoDecoderWithCodecType:videoFormatDescription:videoDecoderSpecifications:extensionDecoderPixelBufferManager:error:"), codecType, videoFormatDescription, videoDecoderSpecifications, extensionDecoderPixelBufferManager)
	if err != nil {
		return nil, err
	}
	return MEVideoDecoderObjectFromID(rv), nil
}
