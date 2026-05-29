// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

// Package coremedia provides Go bindings for the CoreMedia framework.
//
// Represent time-based audio-visual assets with essential data types.
//
// The Core Media framework defines the media pipeline used by AVFoundation
// and other high-level media frameworks found on Apple platforms. Use Core
// Media’s low-level data types and interfaces to efficiently process media
// samples and manage queues of media data.
//
// # Sample Processing
//
//   - CMSampleBuffer: An object that contains zero or more media samples of a uniform media type. ([CMSampleTimingInfo], [CMBufferGetSizeCallback], [CMItemIndex], [CMItemCount], [CMPersistentTrackID])
//   - CMBlockBuffer: An object the system uses to move blocks of memory through a processing system. ([CMBlockBufferFlags], [CMBlockBufferCustomBlockSource])
//   - CMTaggedBufferGroup: Objective-C types and interfaces for working with Core Media tagged buffer groups. ([CMTaggedBufferGroupFormatType])
//   - CMFormatDescription: A media format descriptor that describes the samples in a sample buffer. ([CMSoundDescriptionFlavor], [CMImageDescriptionFlavor], [CMMetadataDescriptionFlavor], [CMTextDescriptionFlavor], [CMTimeCodeDescriptionFlavor])
//   - CMAttachment: Add supporting metadata to sample buffers. ([CMAttachmentMode])
//
// # Time Representation
//
//   - [CMTime]: A structure that represents time. ([CMTimeRoundingMethod], [CMTime], [CMTimeValue], [CMTimeScale], [CMTimeEpoch])
//   - [CMTimeRange]: A structure that represents a range of time. ([CMTimeRange])
//   - [CMTimeMapping]: A structure that maps a segment of a source time range to a target time range. ([CMTimeMapping])
//
// # Media Synchronization
//
//   - CMClock: A reference clock you use to synchronize applications and devices.
//   - CMAudioClock: A specialized reference clock that synchronizes with audio sources.
//   - CMTimebase: A model of a timeline under application control.
//
// # Text Markup
//
//   - CMTextMarkup: Attributes that specify text markup in legible media.
//
// # Metadata
//
//   - CMMetadata: The APIs for working with the framework’s Metadata Identifier Services and Metadata Data Type Registry.
//   - [CMTag]: Types and interfaces for working with Core Media tags. ([CMTag])
//   - CMTagCollection: Objective-C types and interfaces for working with Core Media tag collections.
//   - [CMProjectionType]: Constants describing the projection surface information in a 3D video buffer or channel.
//   - [CMStereoViewComponents]: Constants describing the stereo views contained within a buffer or channel.
//   - [CMStereoViewInterpretationOptions]: Create a set of stereo view interpretation options from a constant.
//   - [CMPackingType]: The type of packing within each video frame, if any.
//
// # Queues
//
//   - CMSimpleQueue: A simple, lockless FIFO queue of elements.
//   - CMBufferQueue: A queue of timed buffers. ([CMBufferCallbacks], [CMBufferQueueTriggerHandler], [CMBufferQueueTriggerToken], [CMBufferQueueTriggerCallback], [CMBufferQueueTriggerCondition])
//   - CMMemoryPool: An object that optimizes memory allocation when working with large blocks of memory.
//
// # Variables
//
//   - [KCMTagProjectionTypeParametricImmersive]
//
// # Functions
//
//   - [CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions]
//
// # Macros
//
//   - CM_BRIDGED_MUTABLE_TYPE
package coremedia

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreMedia library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreMedia.framework/CoreMedia",
	"/usr/lib/libCoreMedia.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreMedia: failed to load framework from any known path\n")
	}
}
