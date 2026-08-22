// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for managing camera file downloads.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDownloadDelegate
type ICCameraDeviceDownloadDelegate interface {
	objectivec.IObject
}

// ICCameraDeviceDownloadDelegateObject wraps an existing Objective-C object that conforms to the ICCameraDeviceDownloadDelegate protocol.
type ICCameraDeviceDownloadDelegateObject struct {
	objectivec.Object
}

func (o ICCameraDeviceDownloadDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ICCameraDeviceDownloadDelegateObjectFromID constructs a [ICCameraDeviceDownloadDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ICCameraDeviceDownloadDelegateObjectFromID(id objc.ID) ICCameraDeviceDownloadDelegateObject {
	return ICCameraDeviceDownloadDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the requested download has completed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDownloadDelegate/didDownloadFile(_:error:options:contextInfo:)
func (o ICCameraDeviceDownloadDelegateObject) DidDownloadFileErrorOptionsContextInfo(file IICCameraFile, error_ foundation.NSError, options foundation.INSDictionary, contextInfo uintptr) {
	objc.Send[struct{}](o.ID, objc.Sel("didDownloadFile:error:options:contextInfo:"), file, error_, options, contextInfo)
}

// Updates the delegate about the status of the download.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDownloadDelegate/didReceiveDownloadProgress(for:downloadedBytes:maxBytes:)
func (o ICCameraDeviceDownloadDelegateObject) DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes(file IICCameraFile, downloadedBytes int64, maxBytes int64) {
	objc.Send[struct{}](o.ID, objc.Sel("didReceiveDownloadProgressForFile:downloadedBytes:maxBytes:"), file, downloadedBytes, maxBytes)
}

// ICCameraDeviceDownloadDelegateConfig holds optional typed callbacks for [ICCameraDeviceDownloadDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/iccameradevicedownloaddelegate
type ICCameraDeviceDownloadDelegateConfig struct {

	// Other Methods
	// DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes — Updates the delegate about the status of the download.
	DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes func(file ICCameraFile, downloadedBytes int64, maxBytes int64)
}

// NewICCameraDeviceDownloadDelegate creates an Objective-C object implementing the [ICCameraDeviceDownloadDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ICCameraDeviceDownloadDelegateObject] satisfies the [ICCameraDeviceDownloadDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/iccameradevicedownloaddelegate
func NewICCameraDeviceDownloadDelegate(config ICCameraDeviceDownloadDelegateConfig) ICCameraDeviceDownloadDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoICCameraDeviceDownloadDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes != nil {
		fn := config.DidReceiveDownloadProgressForFileDownloadedBytesMaxBytes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didReceiveDownloadProgressForFile:downloadedBytes:maxBytes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileID objc.ID, downloadedBytes int64, maxBytes int64) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDownloadDelegate", "didReceiveDownloadProgressForFile:downloadedBytes:maxBytes:")
					}
				}()
				file := ICCameraFileFromID(fileID)
				fn(file, downloadedBytes, maxBytes)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ICCameraDeviceDownloadDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewICCameraDeviceDownloadDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ICCameraDeviceDownloadDelegateObjectFromID(instance)
}
