// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

// Package avfoundation provides Go bindings for the AVFoundation framework.
//
// Work with audiovisual assets, control device cameras, process audio, and
// configure system audio interactions.
//
// AVFoundation combines several major technology areas that together
// encompass a wide range of tasks for inspecting, playing, capturing, and
// processing audiovisual media on Apple platforms.
//
// # Essentials
//
//   - [AVFoundation updates]: Learn about important changes to AVFoundation.
//
// # Common
//
//   - [Media assets]: Load media assets from files and streams to inspect their attributes, tracks, and embedded metadata. ([AVAsset], [AVURLAsset], [AVAssetTrack], [AVAssetTrackSegment], [AVAssetTrackGroup])
//   - [Media reading and writing]: Read images from video, export to alternative formats, and perform sample-level reading and writing of media data. ([AVAssetExportSession], [AVAssetImageGenerator], [AVAssetReader], [AVAssetReaderOutput], [AVAssetReaderTrackOutput])
//   - [Media types and utilities]: Identify the types of content and file formats that AVFoundation supports. ([AVMediaType], [AVMediaCharacteristic], [AVFileType], [AVFileTypeProfile])
//   - [Video settings]: Configure video processing settings using standard key and value constants. ([AVVideoCodecType])
//   - [Audio settings]: Configure audio processing settings using standard key and value constants.
//
// # Playback
//
//   - [Media playback]: Manage the playback of media assets and interstitial content, independent of how you present that content in your interface. ([AVPlayer], [AVPlayerItem], [AVPlayerItemTrack], [AVQueuePlayer], [AVPlayerLooper])
//   - [Offline playback and storage]: Download streamed content to disk to allow offline playback, and define policies to automatically remove downloaded assets. ([AVAssetDownloadURLSession], [AVAssetDownloadTask], [AVAssetDownloadStorageManager], [AVAssetDownloadStorageManagementPolicy], [AVMutableAssetDownloadStorageManagementPolicy])
//   - [Streaming and AirPlay]: Stream content wirelessly to other devices using AirPlay, and handle requests involving FairPlay-protected assets. ([AVRouteDetector], [AVAssetResourceLoader], [AVAssetResourceLoaderDelegate], [AVAssetResourceLoadingRequest], [AVAssetResourceRenewalRequest])
//   - [Sample buffer playback]: Create custom controllers to play and synchronize the timing of sample buffer streams. ([AVSampleBufferRequest], [AVSampleBufferGenerator], [AVSampleBufferGeneratorBatch], [AVQueuedSampleBufferRendering], [AVSampleBufferRenderSynchronizer])
//
// # Capture
//
//   - [Capture setup]: Configure built-in cameras and microphones, and external capture devices, for media capture. ([AVCaptureSession], [AVCaptureInput], [AVCaptureOutput], [AVCaptureConnection], [AVCaptureDevice])
//   - [Photo capture]: Capture high-quality still images, Live Photos, and supporting photo data. ([AVCapturePhoto], [AVCapturePhotoOutput], [AVCapturePhotoCaptureDelegate], [AVCapturePhotoOutputReadinessCoordinator], [AVCapturePhotoOutputReadinessCoordinatorDelegate])
//   - [Audio and video capture]: Capture audio and video directly to media files, or capture streams of media for direct access to media sample buffers. ([AVCaptureMovieFileOutput], [AVCaptureAudioFileOutput], [AVCaptureFileOutput], [AVCaptureFileOutputDelegate], [AVCaptureFileOutputRecordingDelegate])
//   - [Additional data capture]: Capture additional data including depth and metadata, and synchronize capture from multiple outputs. ([AVDepthData], [AVCameraCalibrationData], [AVCaptureMetadataOutput], [AVMetadataObject])
//
// # Editing
//
//   - [Composite assets]: Combine tracks and segments of tracks from multiple assets into a composite asset that you can play or process. ([AVComposition], [AVCompositionTrack], [AVCompositionTrackSegment], [AVMutableComposition], [AVMutableCompositionTrack])
//   - [QuickTime movies]: Access the contents of a QuickTime movie file, and perform sample-level edits of its media tracks. ([AVMovie], [AVMovieTrack], [AVMutableMovie], [AVMutableMovieTrack], [AVFragmentedMovie])
//   - [Video effects]: Define standard video transition effects, synchronize layer animations with media timing, and create custom video compositors. ([AVVideoCompositionCoreAnimationTool], [AVVideoComposition], [AVVideoCompositionInstruction], [AVVideoCompositionLayerInstruction], [AVMutableVideoComposition])
//   - [Audio mixing]: Define how to mix the audio levels from multiple audio tracks over an asset’s duration. ([AVAudioMix], [AVMutableAudioMix], [AVAudioMixInputParameters], [AVMutableAudioMixInputParameters])
//
// # Audio
//
//   - [Audio playback, recording, and processing]: Play, record, and process audio; configure your app’s system audio behavior.
//   - [Speech synthesis]: Configure voices to speak strings of text.
//
// # Errors
//
//   - [AVFoundationErrorDomain]: The error domain of AVFoundation errors.
//   - [AVError]: An enumeration that defines the errors that framework operations can generate.//
//
// # Key Types
//
//   - [AVCaptureDevice] - An object that represents a hardware or virtual capture device like a camera or microphone.
//   - [AVPlayerItem] - An object that models the timing and presentation state of an asset during playback.
//   - [AVCapturePhotoOutput] - A capture output for still image, Live Photos, and other photography workflows.
//   - [AVCaptureDeviceFormat] - A class that defines media formats and capture settings that capture devices support.
//   - [AVPlayer] - An object that provides the interface to control the player’s transport behavior.
//   - [AVMutableMovie] - A mutable object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//   - [AVMutableMovieTrack] - A mutable track that conforms to the QuickTime or ISO base media file format.
//   - [AVCaptureSession] - An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
//   - [AVCapturePhotoSettings] - A specification of the features and settings to use for a single photo capture request.
//   - [AVComposition] - An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
//
// [AVFoundation updates]: https://developer.apple.com/documentation/Updates/AVFoundation
// [Additional data capture]: https://developer.apple.com/documentation/avfoundation/additional-data-capture
// [Audio and video capture]: https://developer.apple.com/documentation/avfoundation/audio-and-video-capture
// [Audio mixing]: https://developer.apple.com/documentation/avfoundation/audio-mixing
// [Audio playback, recording, and processing]: https://developer.apple.com/documentation/avfoundation/audio-playback-recording-and-processing
// [Audio settings]: https://developer.apple.com/documentation/avfoundation/audio-settings
// [Capture setup]: https://developer.apple.com/documentation/avfoundation/capture-setup
// [Composite assets]: https://developer.apple.com/documentation/avfoundation/composite-assets
// [Media assets]: https://developer.apple.com/documentation/avfoundation/media-assets
// [Media playback]: https://developer.apple.com/documentation/avfoundation/media-playback
// [Media reading and writing]: https://developer.apple.com/documentation/avfoundation/media-reading-and-writing
// [Media types and utilities]: https://developer.apple.com/documentation/avfoundation/media-types-and-utilities
// [Offline playback and storage]: https://developer.apple.com/documentation/avfoundation/offline-playback-and-storage
// [Photo capture]: https://developer.apple.com/documentation/avfoundation/photo-capture
// [QuickTime movies]: https://developer.apple.com/documentation/avfoundation/quicktime-movies
// [Sample buffer playback]: https://developer.apple.com/documentation/avfoundation/sample-buffer-playback
// [Speech synthesis]: https://developer.apple.com/documentation/avfoundation/speech-synthesis
// [Streaming and AirPlay]: https://developer.apple.com/documentation/avfoundation/streaming-and-airplay
// [Video effects]: https://developer.apple.com/documentation/avfoundation/video-effects
// [Video settings]: https://developer.apple.com/documentation/avfoundation/video-settings
package avfoundation

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the AVFoundation library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/AVFoundation.framework/AVFoundation",
	"/usr/lib/libAVFoundation.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: AVFoundation: failed to load framework from any known path\n")
	}
}
