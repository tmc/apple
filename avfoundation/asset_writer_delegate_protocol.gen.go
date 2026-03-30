// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A delegate protocol that defines the methods to implement to respond to asset-writing events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterDelegate
type AVAssetWriterDelegate interface {
	objectivec.IObject
}

// AVAssetWriterDelegateObject wraps an existing Objective-C object that conforms to the AVAssetWriterDelegate protocol.
type AVAssetWriterDelegateObject struct {
	objectivec.Object
}

func (o AVAssetWriterDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAssetWriterDelegateObjectFromID constructs a [AVAssetWriterDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAssetWriterDelegateObjectFromID(id objc.ID) AVAssetWriterDelegateObject {
	return AVAssetWriterDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the asset writer output segment data.
//
// writer: The asset writer that output segment data.
//
// segmentData: The data for the segment.
//
// segmentType: The type of segment data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterDelegate/assetWriter(_:didOutputSegmentData:segmentType:)
func (o AVAssetWriterDelegateObject) AssetWriterDidOutputSegmentDataSegmentType(writer IAVAssetWriter, segmentData foundation.INSData, segmentType AVAssetSegmentType) {
	objc.Send[struct{}](o.ID, objc.Sel("assetWriter:didOutputSegmentData:segmentType:"), writer, segmentData, segmentType)
}

// Tells the delegate that the asset writer output segment data and a report.
//
// writer: The asset writer that output segment data.
//
// segmentData: The data for the segment.
//
// segmentType: The type of segment data.
//
// segmentReport: A report for the segment data.
//
// # Discussion
//
// The asset writer stops normal file writing when you implement this method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterDelegate/assetWriter(_:didOutputSegmentData:segmentType:segmentReport:)
func (o AVAssetWriterDelegateObject) AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport(writer IAVAssetWriter, segmentData foundation.INSData, segmentType AVAssetSegmentType, segmentReport IAVAssetSegmentReport) {
	objc.Send[struct{}](o.ID, objc.Sel("assetWriter:didOutputSegmentData:segmentType:segmentReport:"), writer, segmentData, segmentType, segmentReport)
}

// AVAssetWriterDelegateConfig holds optional typed callbacks for [AVAssetWriterDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate
type AVAssetWriterDelegateConfig struct {

	// Responding to segment output
	// AssetWriterDidOutputSegmentDataSegmentType — Tells the delegate that the asset writer output segment data.
	AssetWriterDidOutputSegmentDataSegmentType func(writer AVAssetWriter, segmentData foundation.NSData, segmentType AVAssetSegmentType)
	// AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport — Tells the delegate that the asset writer output segment data and a report.
	AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport func(writer AVAssetWriter, segmentData foundation.NSData, segmentType AVAssetSegmentType, segmentReport AVAssetSegmentReport)
}

// NewAVAssetWriterDelegate creates an Objective-C object implementing the [AVAssetWriterDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVAssetWriterDelegateObject] satisfies the [AVAssetWriterDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate
func NewAVAssetWriterDelegate(config AVAssetWriterDelegateConfig) AVAssetWriterDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVAssetWriterDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AssetWriterDidOutputSegmentDataSegmentType != nil {
		fn := config.AssetWriterDidOutputSegmentDataSegmentType
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("assetWriter:didOutputSegmentData:segmentType:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writerID objc.ID, segmentDataID objc.ID, segmentType AVAssetSegmentType) {
				writer := AVAssetWriterFromID(writerID)
				segmentData := foundation.NSDataFromID(segmentDataID)
				fn(writer, segmentData, segmentType)
			},
		})
	}

	if config.AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport != nil {
		fn := config.AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("assetWriter:didOutputSegmentData:segmentType:segmentReport:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writerID objc.ID, segmentDataID objc.ID, segmentType AVAssetSegmentType, segmentReportID objc.ID) {
				writer := AVAssetWriterFromID(writerID)
				segmentData := foundation.NSDataFromID(segmentDataID)
				segmentReport := AVAssetSegmentReportFromID(segmentReportID)
				fn(writer, segmentData, segmentType, segmentReport)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVAssetWriterDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVAssetWriterDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVAssetWriterDelegateObjectFromID(instance)
}
