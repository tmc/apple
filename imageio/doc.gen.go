// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

// Package imageio provides Go bindings for the ImageIO framework.
//
// Read and write most image file formats, and access an image’s metadata.
//
// The Image I/O framework allows applications to read and write most image
// file formats. This framework offers high efficiency, color management, and
// access to image metadata.
//
// # Image Management
//
//   - [CGImageSourceRef]: An opaque type that you use to read image data from a URL, data object, or data consumer. ([CGImageSourceStatus])
//   - [CGImageDestinationRef]: An opaque type that you use to write image data to a URL, data object, or data consumer.
//
// # XMP Metadata
//
//   - [CGImageMetadataRef]: An immutable object that contains the XMP metadata associated with an image. ([CGImageMetadataTagBlock])
//   - [CGMutableImageMetadataRef]: An opaque type for adding or modifying image metadata.
//   - [CGImageMetadataTagRef]: An immutable type that contains information about a single piece of image metadata. ([CGImageMetadataType])
//   - [XMP Namespaces and Prefixes]: Discover the public namespaces and prefixes that exist in XMP metadata tags.
//   - [KCFErrorDomainCGImageMetadata]: The domain for metadata-related errors that originate in the Image I/O framework.
//   - [CGImageMetadataErrors]: Constants for errors that occur when getting or setting metadata information.
//
// # Common Image Properties
//
//   - [Image Properties]: Properties that apply to the container in general, and not necessarily to an individual image in the container. ([CGImagePropertyOrientation])
//   - [EXIF Dictionary Keys]: Metadata keys for Exchangeable Image File Format (EXIF) data.
//   - [IPTC Dictionary Keys]: Metadata keys for International Press Telecommunications Council (IPTC) data.
//   - [GPS Dictionary Keys]: Keys for Global Positioning System (GPS) information.
//   - [WebP Data]: Metadata keys for WebP metadata.
//
// # Format-Specific Properties
//
//   - [CIFF Image Properties]: Metadata keys for the Camera Image File Format (CIFF) image format.
//   - [DNG Image Properties]: Metadata keys for the Digital Negative (DNG) archival format.
//   - [GIF Image Properties]: Metadata keys for the Graphics Interchange Format (GIF).
//   - [HEIC Image Properties]: Metadata keys for the High Efficiency Image Container (HEIC) format.
//   - [JFIF Image Properties]: Metadata keys for the JPEG File Interchange Format (JFIF).
//   - [PNG Image Properties]: Metadata keys for the Portable Network Graphics (PNG) format.
//   - [TGA Image Properties]: Metadata keys for the Truevision Graphics Adapter (TGA) format. ([CGImagePropertyTGACompression])
//   - [TIFF Image Properties]: Metadata keys for the Tagged Image File Format (TIFF).
//   - [8BIM Image Properties]: Metadata keys for the Adobe Photoshop image format.
//
// # Manufacturer-Specific Properties
//
//   - [Nikon Camera Dictionary Keys]: Metadata keys for an image from a Nikon camera.
//   - [Canon Camera Dictionary Keys]: Metadata keys for an image from a Canon camera.
//   - [KCGImagePropertyMakerAppleDictionary]: A dictionary of key-value pairs for an image from an Apple camera.
//   - [KCGImagePropertyMakerMinoltaDictionary]: A dictionary of key-value pairs for an image from a Minolta camera.
//   - [KCGImagePropertyMakerFujiDictionary]: A dictionary of key-value pairs for an image from a Fuji camera.
//   - [KCGImagePropertyMakerOlympusDictionary]: A dictionary of key-value pairs for an image from a Olympus camera.
//   - [KCGImagePropertyMakerPentaxDictionary]: A dictionary of key-value pairs for an image from a Pentax camera.
//   - [KCGImagePropertyRawDictionary]: A dictionary of key-value pairs for an image that contains minimally processed, or raw, data.
//
// # Spatial Photos
//
//   - [Writing spatial photos]: Create spatial photos for visionOS by packaging a pair of left- and right-eye images as a stereo HEIC file with related spatial metadata.
//   - [Creating spatial photos and videos with spatial metadata]: Add spatial metadata to stereo photos and videos to create spatial media for viewing on Apple Vision Pro.
//
// # Animations
//
//   - [CGAnimateImageAtURLWithBlock]: Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//   - [CGAnimateImageDataWithBlock]: Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//   - [CGImageSourceAnimationBlock]: The block to execute for each frame of an image animation.
//   - [KCGImageAnimationStartIndex]: A property that specifies the index of the first frame of an animation.
//   - [KCGImageAnimationDelayTime]: The number of seconds to wait before displaying the next image in an animated sequence.
//   - [KCGImageAnimationLoopCount]: The number of times to repeat the animated sequence.
//   - [CGImageAnimationStatus]: Constants that indicate the result of animating an image sequence.
//
// # Variables
//
//   - [KCGComputeHDRStats]
//   - [KCGImageDestinationEncodeAlternateColorSpace]
//   - [KCGImageDestinationEncodeBaseColorSpace]
//   - [KCGImageDestinationEncodeBaseIsSDR]
//   - [KCGImageDestinationEncodeBasePixelFormatRequest]
//   - [KCGImageDestinationEncodeGainMapPixelFormatRequest]
//   - [KCGImageDestinationEncodeGainMapSubsampleFactor]
//   - [KCGImageDestinationEncodeGenerateGainMapWithBaseImage]
//   - [KCGImageDestinationEncodeIsBaseImage]
//   - [KCGImageDestinationEncodeRequest]
//   - [KCGImageDestinationEncodeRequestOptions]
//   - [KCGImageDestinationEncodeToISOGainmap]
//   - [KCGImageDestinationEncodeToISOHDR]
//   - [KCGImageDestinationEncodeToSDR]
//   - [KCGImageDestinationEncodeTonemapMode]
//   - [KCGImagePropertyASTCBlockSize]
//   - [KCGImagePropertyASTCBlockSize4x4]
//   - [KCGImagePropertyASTCBlockSize8x8]
//   - [KCGImagePropertyASTCEncoder]
//   - [KCGImagePropertyBCEncoder]
//   - [KCGImagePropertyBCFormat]
//   - [KCGImagePropertyEncoder]
//   - [KCGImagePropertyOpenEXRCompression]
//   - [KCGImagePropertyPVREncoder]
//   - [KCGImageProviderPreferredTileHeight]
//   - [KCGImageProviderPreferredTileWidth]
//   - [KCGImageSourceGenerateImageSpecificLumaScaling]//
//
// [8BIM Image Properties]: https://developer.apple.com/documentation/imageio/8bim-image-properties
// [CIFF Image Properties]: https://developer.apple.com/documentation/imageio/ciff-image-properties
// [Canon Camera Dictionary Keys]: https://developer.apple.com/documentation/imageio/canon-camera-dictionary-keys
// [Creating spatial photos and videos with spatial metadata]: https://developer.apple.com/documentation/imageio/creating-spatial-photos-and-videos-with-spatial-metadata
// [DNG Image Properties]: https://developer.apple.com/documentation/imageio/dng-image-properties
// [EXIF Dictionary Keys]: https://developer.apple.com/documentation/imageio/exif-dictionary-keys
// [GIF Image Properties]: https://developer.apple.com/documentation/imageio/gif-image-properties
// [GPS Dictionary Keys]: https://developer.apple.com/documentation/imageio/gps-dictionary-keys
// [HEIC Image Properties]: https://developer.apple.com/documentation/imageio/heic-image-properties
// [IPTC Dictionary Keys]: https://developer.apple.com/documentation/imageio/iptc-dictionary-keys
// [Image Properties]: https://developer.apple.com/documentation/imageio/image-properties
// [JFIF Image Properties]: https://developer.apple.com/documentation/imageio/jfif-image-properties
// [Nikon Camera Dictionary Keys]: https://developer.apple.com/documentation/imageio/nikon-camera-dictionary-keys
// [PNG Image Properties]: https://developer.apple.com/documentation/imageio/png-image-properties
// [TGA Image Properties]: https://developer.apple.com/documentation/imageio/tga-image-properties
// [TIFF Image Properties]: https://developer.apple.com/documentation/imageio/tiff-image-properties
// [WebP Data]: https://developer.apple.com/documentation/imageio/webp-data
// [Writing spatial photos]: https://developer.apple.com/documentation/imageio/writing-spatial-photos
// [XMP Namespaces and Prefixes]: https://developer.apple.com/documentation/imageio/xmp-namespaces-and-prefixes
package imageio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ImageIO library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ImageIO.framework/ImageIO",
	"/usr/lib/libImageIO.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: ImageIO: failed to load framework from any known path\n")
	}
}
