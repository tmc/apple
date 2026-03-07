
// Code generated from Apple documentation for Vision. DO NOT EDIT.

// Package vision provides Go bindings for the Vision framework.
//
// Apply computer vision algorithms to perform a variety of tasks on input images and videos.
//
// The Vision framework combines machine learning technologies and Swift’s concurrency features to perform computer vision tasks in your app. Use the Vision framework to analyze images for a variety of purposes:
//
// # Text and document analysis
//
//   - Locating and displaying recognized text: Perform text recognition on a photo using the Vision framework’s text-recognition request.
//   - Recognizing tables within a document: Scan a document that contains a table and extract its content in a formatted way.
//   - DetectBarcodesRequest: A request that detects barcodes in an image. ([BarcodeObservation], [BarcodeSymbology])
//   - DetectDocumentSegmentationRequest: A request that detects rectangular regions that contain text in the input image.
//   - DetectTextRectanglesRequest: An image-analysis request that finds regions of visible text in an image. ([TextObservation])
//   - RecognizeDocumentsRequest: An image-analysis request to scan an image of a document and provide information about its structure.
//   - RecognizeTextRequest: An image-analysis request that recognizes text in an image. ([RecognizedTextObservation])
//
// # Facial analysis
//
//   - Analyzing a selfie and visualizing its content: Calculate face-capture quality and visualize facial features for a collection of images using the Vision framework.
//   - DetectFaceCaptureQualityRequest: A request that produces a floating-point number that represents the capture quality of a face in a photo. ([FaceObservation])
//   - DetectFaceLandmarksRequest: An image-analysis request that finds facial features like eyes and mouth in an image. ([FaceObservation])
//   - DetectFaceRectanglesRequest: A request that finds faces within an image. ([FaceObservation])
//
// # Image segmentation and subject lifting
//
//   - GenerateForegroundInstanceMaskRequest: A request that generates an instance mask of noticeable objects to separate from the background. ([InstanceMaskObservation])
//   - GeneratePersonInstanceMaskRequest: A request that produces a mask of individual people it finds in the input image. ([InstanceMaskObservation])
//   - GeneratePersonSegmentationRequest: A request that produces a matte image for a person it finds in the input image. ([PixelBufferObservation])
//
// # Pose analysis
//
//   - DetectAnimalBodyPoseRequest: A request that detects an animal body pose. ([AnimalBodyPoseObservation])
//   - DetectHumanBodyPose3DRequest: A request that detects points on human bodies in 3D space, relative to the camera. ([HumanBodyPose3DObservation])
//   - DetectHumanBodyPoseRequest: A request that detects a human body pose. ([HumanBodyPoseObservation])
//   - DetectHumanHandPoseRequest: A request that detects a human hand pose. ([HumanHandPoseObservation])
//   - Supporting Pose Types: Types you use when working with pose analysis. ([Chirality])
//
// # Image classification and recognition
//
//   - Classifying images for categorization and search: Analyze and label images using a Vision classification request.
//   - ClassifyImageRequest: A request to classify an image. ([ClassificationObservation])
//   - DetectHumanRectanglesRequest: A request that finds rectangular regions that contain people in an image. ([HumanObservation])
//   - RecognizeAnimalsRequest: A request that recognizes animals in an image. ([RecognizedObjectObservation])
//
// # Shape and edge detection
//
//   - DetectContoursRequest: A request that detects the contours of the edges of an image. ([ContoursObservation])
//   - DetectHorizonRequest: An image-analysis request that determines the horizon angle in an image. ([HorizonObservation])
//   - DetectRectanglesRequest: An image-analysis request that finds projected rectangular regions in an image. ([RectangleObservation])
//
// # Image quality and saliency analysis
//
//   - Generating high-quality thumbnails from videos: Identify the most visually pleasing frames in a video by using the image-aesthetics scores request.
//   - CalculateImageAestheticsScoresRequest: A request that analyzes an image for aesthetically pleasing attributes. ([ImageAestheticsScoresObservation])
//   - DetectLensSmudgeRequest: A request that detects a smudge on a lens from an image or video frame capture.
//   - GenerateAttentionBasedSaliencyImageRequest: An object that produces a heat map that identifies the parts of an image most likely to draw attention. ([SaliencyImageObservation])
//   - GenerateObjectnessBasedSaliencyImageRequest: A request that generates a heat map that identifies the parts of an image most likely to represent objects. ([SaliencyImageObservation])
//
// # Motion and object tracking
//
//   - DetectTrajectoriesRequest: A request that detects the trajectories of shapes moving along a parabolic path. ([TrajectoryObservation])
//   - TrackObjectRequest: An image-analysis request that tracks the movement of a previously identified object across multiple images or video frames. ([DetectedObjectObservation])
//   - TrackOpticalFlowRequest: A request that determines the direction change of vectors for each pixel from a previous to current image.
//   - TrackRectangleRequest: An image-analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames. ([RectangleObservation])
//
// # Image registration and comparison
//
//   - GenerateImageFeaturePrintRequest: An image-based request to generate feature prints from an image. ([FeaturePrintObservation])
//   - TrackHomographicImageRegistrationRequest: An image-analysis request that you track over time to determine the perspective warp matrix necessary to align the content of two images. ([ImageHomographicAlignmentObservation])
//   - TrackTranslationalImageRegistrationRequest: An image-analysis request that you track over time to determine the affine transform necessary to align the content of two images. ([ImageTranslationAlignmentObservation], [ComputeStage])
//
// # Custom Core ML integration
//
//   - CoreMLRequest: An image-analysis request that uses a Core ML model to process images. ([PixelBufferObservation], [ClassificationObservation], [CoreMLFeatureValueObservation], [ComputeStage])
//
// # Protocols
//
//   - ImageProcessingRequest: A type for image-analysis requests that focus on a specific part of an image.
//   - PoseProviding: An observation that provides a collection of joints that make up a pose.
//   - StatefulRequest: The protocol for a type that builds evidence of a condition over time.
//   - TargetedRequest: A type for analyzing two images together.
//   - VisionObservation: A type for objects produced by image-analysis requests.
//   - VisionRequest: A type for image-analysis requests. ([ComputeStage])
//
// # Request handlers
//
//   - ImageRequestHandler: An object that processes one or more image-analysis requests pertaining to a single image.
//   - TargetedImageRequestHandler: An object that performs image-analysis requests on two images.
//   - VideoProcessor: An object that performs offline analysis of video content.
//
// # Image locations and regions
//
//   - NormalizedPoint: A point in a 2D coordinate system.
//   - NormalizedRect: The location and dimensions of a rectangle.
//   - NormalizedRegion: A polygon composed of normalized points.
//   - NormalizedCircle: The center point and radius of a 2D circle.
//   - BoundingBoxProviding: A protocol for objects that have a bounding box.
//   - BoundingRegionProviding: A protocol for objects that have a defined boundary in an image.
//   - QuadrilateralProviding: A protocol for objects that have a bounding quadrilateral.
//   - CoordinateOrigin: The origin of a coordinate system relative to an image.
//
// # Errors
//
//   - VisionError: The errors that the framework produces.
//
// # Legacy API
//
//   - Original Objective-C and Swift API ([VNRequest], [VNImageBasedRequest], [VNClassifyImageRequest], [VNGenerateImageFeaturePrintRequest], [VNFeaturePrintObservation])
//
// # Key Types
//
//   - [VNGenerateOpticalFlowRequest] - An object that generates directional change vectors for each pixel in the targeted image.
//   - [VNHomographicImageRegistrationRequest] - An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//   - [VNTranslationalImageRegistrationRequest] - An image-analysis request that determines the affine transform necessary to align the content of two images.
//   - [VNImageRegistrationRequest] - The abstract superclass for image-analysis requests that align images according to their content.
//   - [VNTargetedImageRequest] - The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//   - [VNFaceLandmarks2D] - A collection of facial features that a request detects.
//   - [VNRecognizeTextRequest] - An image-analysis request that finds and recognizes text in an image.
//   - [VNBarcodeObservation] - An object that represents barcode information that an image analysis request detects.
//   - [VNImageRequestHandler] - An object that processes one or more image-analysis request pertaining to a single image.
//   - [VNRequest] - The abstract superclass for analysis requests.
//
// [Vision Documentation]: https://developer.apple.com/documentation/Vision
package vision

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Vision.framework/Vision"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Vision: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

