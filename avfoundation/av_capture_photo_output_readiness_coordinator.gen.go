// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCapturePhotoOutputReadinessCoordinator] class.
var (
	_AVCapturePhotoOutputReadinessCoordinatorClass     AVCapturePhotoOutputReadinessCoordinatorClass
	_AVCapturePhotoOutputReadinessCoordinatorClassOnce sync.Once
)

func getAVCapturePhotoOutputReadinessCoordinatorClass() AVCapturePhotoOutputReadinessCoordinatorClass {
	_AVCapturePhotoOutputReadinessCoordinatorClassOnce.Do(func() {
		_AVCapturePhotoOutputReadinessCoordinatorClass = AVCapturePhotoOutputReadinessCoordinatorClass{class: objc.GetClass("AVCapturePhotoOutputReadinessCoordinator")}
	})
	return _AVCapturePhotoOutputReadinessCoordinatorClass
}

// GetAVCapturePhotoOutputReadinessCoordinatorClass returns the class object for AVCapturePhotoOutputReadinessCoordinator.
func GetAVCapturePhotoOutputReadinessCoordinatorClass() AVCapturePhotoOutputReadinessCoordinatorClass {
	return getAVCapturePhotoOutputReadinessCoordinatorClass()
}

type AVCapturePhotoOutputReadinessCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCapturePhotoOutputReadinessCoordinatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCapturePhotoOutputReadinessCoordinatorClass) Alloc() AVCapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[AVCapturePhotoOutputReadinessCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that monitors changes to a photo output’s capture readiness.
//
// # Overview
// 
// Use this object to coordinate user interface updates on the main queue with
// a [AVCapturePhotoOutput] that runs on a background queue. Adopt the
// [AVCapturePhotoOutputReadinessCoordinatorDelegate] protocol in your app and
// set its implementation as the coordinator’s delegate object to receive
// callbacks as the associated photo output’s [AVCapturePhotoOutputReadinessCoordinator.CaptureReadiness] state
// changes.
// 
// You can track additional capture requests with this object by calling its
// [AVCapturePhotoOutputReadinessCoordinator.StartTrackingCaptureRequestUsingPhotoSettings] method. You can use it to
// synchronously update shutter button availability and appearance and on the
// main thread while calling the photo output’s
// [CapturePhotoWithSettingsDelegate] method asynchronously on a background
// queue.
//
// # Creating a coordinator
//
//   - [AVCapturePhotoOutputReadinessCoordinator.InitWithPhotoOutput]: Creates an object that helps coordinate user interface changes with a photo output that runs on a background queue.
//
// # Setting the delegate object
//
//   - [AVCapturePhotoOutputReadinessCoordinator.Delegate]: The coordinator’s delegate object.
//   - [AVCapturePhotoOutputReadinessCoordinator.SetDelegate]
//
// # Performing tracking requests
//
//   - [AVCapturePhotoOutputReadinessCoordinator.StartTrackingCaptureRequestUsingPhotoSettings]: Tracks a capture request that uses the specified photo settings.
//   - [AVCapturePhotoOutputReadinessCoordinator.StopTrackingCaptureRequestUsingPhotoSettingsUniqueID]: Stop tracking the capture request represented by the specified photo setting’s unique identifier.
//
// # Determining readiness for capture
//
//   - [AVCapturePhotoOutputReadinessCoordinator.CaptureReadiness]: A value that indicates whether the coordinator’s photo output is ready to respond to new capture requests in a timely manner.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator
type AVCapturePhotoOutputReadinessCoordinator struct {
	objectivec.Object
}

// AVCapturePhotoOutputReadinessCoordinatorFromID constructs a [AVCapturePhotoOutputReadinessCoordinator] from an objc.ID.
//
// An object that monitors changes to a photo output’s capture readiness.
func AVCapturePhotoOutputReadinessCoordinatorFromID(id objc.ID) AVCapturePhotoOutputReadinessCoordinator {
	return AVCapturePhotoOutputReadinessCoordinator{objectivec.Object{ID: id}}
}
// NOTE: AVCapturePhotoOutputReadinessCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCapturePhotoOutputReadinessCoordinator] class.
//
// # Creating a coordinator
//
//   - [IAVCapturePhotoOutputReadinessCoordinator.InitWithPhotoOutput]: Creates an object that helps coordinate user interface changes with a photo output that runs on a background queue.
//
// # Setting the delegate object
//
//   - [IAVCapturePhotoOutputReadinessCoordinator.Delegate]: The coordinator’s delegate object.
//   - [IAVCapturePhotoOutputReadinessCoordinator.SetDelegate]
//
// # Performing tracking requests
//
//   - [IAVCapturePhotoOutputReadinessCoordinator.StartTrackingCaptureRequestUsingPhotoSettings]: Tracks a capture request that uses the specified photo settings.
//   - [IAVCapturePhotoOutputReadinessCoordinator.StopTrackingCaptureRequestUsingPhotoSettingsUniqueID]: Stop tracking the capture request represented by the specified photo setting’s unique identifier.
//
// # Determining readiness for capture
//
//   - [IAVCapturePhotoOutputReadinessCoordinator.CaptureReadiness]: A value that indicates whether the coordinator’s photo output is ready to respond to new capture requests in a timely manner.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator
type IAVCapturePhotoOutputReadinessCoordinator interface {
	objectivec.IObject

	// Topic: Creating a coordinator

	// Creates an object that helps coordinate user interface changes with a photo output that runs on a background queue.
	InitWithPhotoOutput(photoOutput IAVCapturePhotoOutput) AVCapturePhotoOutputReadinessCoordinator

	// Topic: Setting the delegate object

	// The coordinator’s delegate object.
	Delegate() AVCapturePhotoOutputReadinessCoordinatorDelegate
	SetDelegate(value AVCapturePhotoOutputReadinessCoordinatorDelegate)

	// Topic: Performing tracking requests

	// Tracks a capture request that uses the specified photo settings.
	StartTrackingCaptureRequestUsingPhotoSettings(settings IAVCapturePhotoSettings)
	// Stop tracking the capture request represented by the specified photo setting’s unique identifier.
	StopTrackingCaptureRequestUsingPhotoSettingsUniqueID(settingsUniqueID int64)

	// Topic: Determining readiness for capture

	// A value that indicates whether the coordinator’s photo output is ready to respond to new capture requests in a timely manner.
	CaptureReadiness() AVCapturePhotoOutputCaptureReadiness
}

// Init initializes the instance.
func (c AVCapturePhotoOutputReadinessCoordinator) Init() AVCapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[AVCapturePhotoOutputReadinessCoordinator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCapturePhotoOutputReadinessCoordinator) Autorelease() AVCapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[AVCapturePhotoOutputReadinessCoordinator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCapturePhotoOutputReadinessCoordinator creates a new AVCapturePhotoOutputReadinessCoordinator instance.
func NewAVCapturePhotoOutputReadinessCoordinator() AVCapturePhotoOutputReadinessCoordinator {
	class := getAVCapturePhotoOutputReadinessCoordinatorClass()
	rv := objc.Send[AVCapturePhotoOutputReadinessCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that helps coordinate user interface changes with a photo
// output that runs on a background queue.
//
// photoOutput: The photo output to monitor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/init(photoOutput:)
func NewCapturePhotoOutputReadinessCoordinatorWithPhotoOutput(photoOutput IAVCapturePhotoOutput) AVCapturePhotoOutputReadinessCoordinator {
	instance := getAVCapturePhotoOutputReadinessCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPhotoOutput:"), photoOutput)
	return AVCapturePhotoOutputReadinessCoordinatorFromID(rv)
}

// Creates an object that helps coordinate user interface changes with a photo
// output that runs on a background queue.
//
// photoOutput: The photo output to monitor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/init(photoOutput:)
func (c AVCapturePhotoOutputReadinessCoordinator) InitWithPhotoOutput(photoOutput IAVCapturePhotoOutput) AVCapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[AVCapturePhotoOutputReadinessCoordinator](c.ID, objc.Sel("initWithPhotoOutput:"), photoOutput)
	return rv
}
// Tracks a capture request that uses the specified photo settings.
//
// settings: A settings object that the system passes [CapturePhotoWithSettingsDelegate]
// for this capture request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/startTrackingCaptureRequest(using:)
func (c AVCapturePhotoOutputReadinessCoordinator) StartTrackingCaptureRequestUsingPhotoSettings(settings IAVCapturePhotoSettings) {
	objc.Send[objc.ID](c.ID, objc.Sel("startTrackingCaptureRequestUsingPhotoSettings:"), settings)
}
// Stop tracking the capture request represented by the specified photo
// setting’s unique identifier.
//
// settingsUniqueID: The [UniqueID] value of the related photo settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/stopTrackingCaptureRequest(using:)
func (c AVCapturePhotoOutputReadinessCoordinator) StopTrackingCaptureRequestUsingPhotoSettingsUniqueID(settingsUniqueID int64) {
	objc.Send[objc.ID](c.ID, objc.Sel("stopTrackingCaptureRequestUsingPhotoSettingsUniqueID:"), settingsUniqueID)
}

// The coordinator’s delegate object.
//
// # Discussion
// 
// The capture delegate receives callbacks when the photo output’s
// captureReadiness changes. It calls its delegate on the main queue, which
// allows you to perform user interface updates directly from the delegate’s
// [ReadinessCoordinatorCaptureReadinessDidChange] method.
// 
// The coordinator provides an initial value to the delegate when you first
// set it on this object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/delegate
func (c AVCapturePhotoOutputReadinessCoordinator) Delegate() AVCapturePhotoOutputReadinessCoordinatorDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return AVCapturePhotoOutputReadinessCoordinatorDelegateObjectFromID(rv)
}
func (c AVCapturePhotoOutputReadinessCoordinator) SetDelegate(value AVCapturePhotoOutputReadinessCoordinatorDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}
// A value that indicates whether the coordinator’s photo output is ready to
// respond to new capture requests in a timely manner.
//
// # Discussion
// 
// The value incorporates the photo output’s [CaptureReadiness] property
// value and any requests registered by calling the
// [StartTrackingCaptureRequestUsingPhotoSettings] method. The system updates
// this value before calling the
// [ReadinessCoordinatorCaptureReadinessDidChange] method.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/captureReadiness
func (c AVCapturePhotoOutputReadinessCoordinator) CaptureReadiness() AVCapturePhotoOutputCaptureReadiness {
	rv := objc.Send[AVCapturePhotoOutputCaptureReadiness](c.ID, objc.Sel("captureReadiness"))
	return AVCapturePhotoOutputCaptureReadiness(rv)
}

