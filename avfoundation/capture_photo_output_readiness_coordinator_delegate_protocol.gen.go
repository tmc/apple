// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A delegate protocol to receive updates about a photo output’s capture readiness.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinatorDelegate
type AVCapturePhotoOutputReadinessCoordinatorDelegate interface {
	objectivec.IObject
}

// AVCapturePhotoOutputReadinessCoordinatorDelegateObject wraps an existing Objective-C object that conforms to the AVCapturePhotoOutputReadinessCoordinatorDelegate protocol.
type AVCapturePhotoOutputReadinessCoordinatorDelegateObject struct {
	objectivec.Object
}

func (o AVCapturePhotoOutputReadinessCoordinatorDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCapturePhotoOutputReadinessCoordinatorDelegateObjectFromID constructs a [AVCapturePhotoOutputReadinessCoordinatorDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCapturePhotoOutputReadinessCoordinatorDelegateObjectFromID(id objc.ID) AVCapturePhotoOutputReadinessCoordinatorDelegateObject {
	return AVCapturePhotoOutputReadinessCoordinatorDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the capture readiness state of a photo output
// changed.
//
// coordinator: The delegate’s coordinator object.
//
// captureReadiness: An updated capture readiness value.
//
// # Discussion
//
// The system always performs this call on the main queue, so you can use it
// to update your user interface’s shutter button availability and
// appearance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinatorDelegate/readinessCoordinator(_:captureReadinessDidChange:)
func (o AVCapturePhotoOutputReadinessCoordinatorDelegateObject) ReadinessCoordinatorCaptureReadinessDidChange(coordinator IAVCapturePhotoOutputReadinessCoordinator, captureReadiness AVCapturePhotoOutputCaptureReadiness) {
	objc.Send[struct{}](o.ID, objc.Sel("readinessCoordinator:captureReadinessDidChange:"), coordinator, captureReadiness)
}

// AVCapturePhotoOutputReadinessCoordinatorDelegateConfig holds optional typed callbacks for [AVCapturePhotoOutputReadinessCoordinatorDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutputreadinesscoordinatordelegate
type AVCapturePhotoOutputReadinessCoordinatorDelegateConfig struct {

	// Monitoring capture readiness
	// ReadinessCoordinatorCaptureReadinessDidChange — Tells the delegate that the capture readiness state of a photo output changed.
	ReadinessCoordinatorCaptureReadinessDidChange func(coordinator AVCapturePhotoOutputReadinessCoordinator, captureReadiness AVCapturePhotoOutputCaptureReadiness)
}

// NewAVCapturePhotoOutputReadinessCoordinatorDelegate creates an Objective-C object implementing the [AVCapturePhotoOutputReadinessCoordinatorDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCapturePhotoOutputReadinessCoordinatorDelegateObject] satisfies the [AVCapturePhotoOutputReadinessCoordinatorDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutputreadinesscoordinatordelegate
func NewAVCapturePhotoOutputReadinessCoordinatorDelegate(config AVCapturePhotoOutputReadinessCoordinatorDelegateConfig) AVCapturePhotoOutputReadinessCoordinatorDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCapturePhotoOutputReadinessCoordinatorDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ReadinessCoordinatorCaptureReadinessDidChange != nil {
		fn := config.ReadinessCoordinatorCaptureReadinessDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("readinessCoordinator:captureReadinessDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, coordinatorID objc.ID, captureReadiness AVCapturePhotoOutputCaptureReadiness) {
				coordinator := AVCapturePhotoOutputReadinessCoordinatorFromID(coordinatorID)
				fn(coordinator, captureReadiness)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCapturePhotoOutputReadinessCoordinatorDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCapturePhotoOutputReadinessCoordinatorDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCapturePhotoOutputReadinessCoordinatorDelegateObjectFromID(instance)
}
