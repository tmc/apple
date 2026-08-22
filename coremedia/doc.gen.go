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
// # Metadata
//
//   - [CMTag]: Types and interfaces for working with Core Media tags. ([CMTag])
//   - [CMProjectionType]: Constants describing the projection surface information in a 3D video buffer or channel.
//   - [CMStereoViewComponents]: Constants describing the stereo views contained within a buffer or channel.
//   - [CMStereoViewInterpretationOptions]: Create a set of stereo view interpretation options from a constant.
//   - [CMPackingType]: The type of packing within each video frame, if any.
//
// # Queues
//
//   - CMBufferQueue: A queue of timed buffers. ([CMBufferCallbacks], [CMBufferQueueTriggerHandler], [CMBufferQueueTriggerToken], [CMBufferQueueTriggerCallback], [CMBufferQueueTriggerCondition])
//
// # Variables
//
//   - [KCMTagProjectionTypeParametricImmersive]
//
// # Functions
//
//   - [CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions]
package coremedia

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreMedia library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
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
