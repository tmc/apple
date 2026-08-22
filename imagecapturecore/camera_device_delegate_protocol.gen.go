// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for detecting cameras, getting metadata and thumbnails, handling access and capability changes, and performing other actions on connected cameras.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate
type ICCameraDeviceDelegate interface {
	objectivec.IObject
	ICDeviceDelegate

	// Tells the client that the camera device is done enumerating its content and is ready to receive requests.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/deviceDidBecomeReady(withCompleteContentCatalog:)
	DeviceDidBecomeReadyWithCompleteContentCatalog(device IICCameraDevice)

	// Tells the client when objects are added to the device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didAdd:)-8oukd
	CameraDeviceDidAddItems(camera IICCameraDevice, items []ICCameraItem)

	// Tells the client when objects are removed from the device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didRemove:)-4m5al
	CameraDeviceDidRemoveItems(camera IICCameraDevice, items []ICCameraItem)

	// Tells the client when one or more objects are renamed on the device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didRenameItems:)
	CameraDeviceDidRenameItems(camera IICCameraDevice, items []ICCameraItem)

	// Tells the client when the metadata requested for an item on a camera is available.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceiveMetadata:for:error:)
	CameraDeviceDidReceiveMetadataForItemError(camera IICCameraDevice, metadata foundation.INSDictionary, item IICCameraItem, error_ foundation.NSError)

	// Tells the client when the requested thumbnail is available.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceiveThumbnail:for:error:)
	CameraDeviceDidReceiveThumbnailForItemError(camera IICCameraDevice, thumbnail coregraphics.CGImageRef, item IICCameraItem, error_ foundation.NSError)

	// Tells the client when a capability of a camera changes.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidChangeCapability(_:)
	CameraDeviceDidChangeCapability(camera IICCameraDevice)

	// Tells the client when an Apple device has been locked, and media is unavailable until the restriction has been removed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidEnableAccessRestriction(_:)
	CameraDeviceDidEnableAccessRestriction(device IICDevice)

	// Tells the client when an Apple device has been unlocked, paired to the host, and media is available.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidRemoveAccessRestriction(_:)
	CameraDeviceDidRemoveAccessRestriction(device IICDevice)

	// Tells the client about a PTP event.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceivePTPEvent:)
	CameraDeviceDidReceivePTPEvent(camera IICCameraDevice, eventData foundation.NSData)
}

// ICCameraDeviceDelegateObject wraps an existing Objective-C object that conforms to the ICCameraDeviceDelegate protocol.
type ICCameraDeviceDelegateObject struct {
	objectivec.Object
}

func (o ICCameraDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ICCameraDeviceDelegateObjectFromID constructs a [ICCameraDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ICCameraDeviceDelegateObjectFromID(id objc.ID) ICCameraDeviceDelegateObject {
	return ICCameraDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the client that the camera device is done enumerating its content and
// is ready to receive requests.
//
// # Discussion
//
// You must open a session on the device before you can enumerate its content
// and make it ready to receive requests.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/deviceDidBecomeReady(withCompleteContentCatalog:)
func (o ICCameraDeviceDelegateObject) DeviceDidBecomeReadyWithCompleteContentCatalog(device IICCameraDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidBecomeReadyWithCompleteContentCatalog:"), device)
}

// Tells the client when objects are added to the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didAdd:)-8oukd
func (o ICCameraDeviceDelegateObject) CameraDeviceDidAddItems(camera IICCameraDevice, items []ICCameraItem) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didAddItems:"), camera, items)
}

// Tells the client when objects are removed from the device.
//
// # Discussion
//
// The objects in items are instances of the ICCameraFile class.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didRemove:)-4m5al
func (o ICCameraDeviceDelegateObject) CameraDeviceDidRemoveItems(camera IICCameraDevice, items []ICCameraItem) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didRemoveItems:"), camera, items)
}

// Tells the client when one or more objects are renamed on the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didRenameItems:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidRenameItems(camera IICCameraDevice, items []ICCameraItem) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didRenameItems:"), camera, items)
}

// Tells the client when the metadata requested for an item on a camera is
// available.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceiveMetadata:for:error:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidReceiveMetadataForItemError(camera IICCameraDevice, metadata foundation.INSDictionary, item IICCameraItem, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didReceiveMetadata:forItem:error:"), camera, metadata, item, error_)
}

// Tells the client when the requested thumbnail is available.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceiveThumbnail:for:error:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidReceiveThumbnailForItemError(camera IICCameraDevice, thumbnail coregraphics.CGImageRef, item IICCameraItem, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didReceiveThumbnail:forItem:error:"), camera, thumbnail, item, error_)
}

// Tells the client when a capability of a camera changes.
//
// # Discussion
//
// For more information about device capabilities, see [ICDeviceCapability].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidChangeCapability(_:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidChangeCapability(camera IICCameraDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDeviceDidChangeCapability:"), camera)
}

// Tells the client when an Apple device has been locked, and media is
// unavailable until the restriction has been removed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidEnableAccessRestriction(_:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidEnableAccessRestriction(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDeviceDidEnableAccessRestriction:"), device)
}

// Tells the client when an Apple device has been unlocked, paired to the
// host, and media is available.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDeviceDidRemoveAccessRestriction(_:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidRemoveAccessRestriction(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDeviceDidRemoveAccessRestriction:"), device)
}

// Tells the client about a PTP event.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didReceivePTPEvent:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidReceivePTPEvent(camera IICCameraDevice, eventData foundation.NSData) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didReceivePTPEvent:"), camera, eventData)
}

// Tells the client when the camera completes a delete operation.
//
// # Discussion
//
// Initiate a delete operation using [ICCameraDevice.RequestDeleteFiles].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:didCompleteDeleteFilesWithError:)
func (o ICCameraDeviceDelegateObject) CameraDeviceDidCompleteDeleteFilesWithError(camera IICCameraDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("cameraDevice:didCompleteDeleteFilesWithError:"), camera, error_)
}

// Tells the client when the camera is about to execute queued requests for
// the metadata of a specific item.
//
// # Discussion
//
// If the request is no longer needed—for example, if the item is no longer
// displayed on the screen—the client can cancel sending a request to the
// camera, speeding up the execution queue.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:shouldGetMetadataOf:)
func (o ICCameraDeviceDelegateObject) CameraDeviceShouldGetMetadataOfItem(cameraDevice IICCameraDevice, item IICCameraItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("cameraDevice:shouldGetMetadataOfItem:"), cameraDevice, item)
	return rv
}

// Tells the client when the camera is about to execute queued requests for
// the thumbnail of a specific item.
//
// # Discussion
//
// If the request is no longer needed—for example, if the item is no longer
// displayed on the screen—the client can cancel sending a request to the
// camera, speeding up the execution queue.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDeviceDelegate/cameraDevice(_:shouldGetThumbnailOf:)
func (o ICCameraDeviceDelegateObject) CameraDeviceShouldGetThumbnailOfItem(cameraDevice IICCameraDevice, item IICCameraItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("cameraDevice:shouldGetThumbnailOfItem:"), cameraDevice, item)
	return rv
}

// Tells the delegate when a session is opened on a device.
//
// # Discussion
//
// This message completes the process initiated by the message
// “requestOpenSession” sent to the device object.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didOpenSessionWithError:)
func (o ICCameraDeviceDelegateObject) DeviceDidOpenSessionWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didOpenSessionWithError:"), device, error_)
}

// Tells the delegate when a session is closed on a device.
//
// # Discussion
//
// This message completes the process initiated by the message
// “requestCloseSession” sent to the device object. This message is also
// sent if the device module in control of the device ceases to control the
// device.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didCloseSessionWithError:)
func (o ICCameraDeviceDelegateObject) DeviceDidCloseSessionWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didCloseSessionWithError:"), device, error_)
}

// Tells the delegate that a device has been removed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/didRemove(_:)
func (o ICCameraDeviceDelegateObject) DidRemoveDevice(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("didRemoveDevice:"), device)
}

// Tells the delegate when the device is ready to receive requests.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/deviceDidBecomeReady(_:)
func (o ICCameraDeviceDelegateObject) DeviceDidBecomeReady(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidBecomeReady:"), device)
}

// Tells the delegate when status information is received from a device.
//
// # Discussion
//
// The ‘status’ dictionary contains two keys, ICStatusNotificationKey and
// ICLocalizedStatusNotificationKey, which are defined above. Status
// information keys are located in their respective ICDevice type class
// header.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didReceiveStatusInformation:)
func (o ICCameraDeviceDelegateObject) DeviceDidReceiveStatusInformation(device IICDevice, status foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didReceiveStatusInformation:"), device, status)
}

// Tells the delegate when a device encounters an error.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEncounterError:)
func (o ICCameraDeviceDelegateObject) DeviceDidEncounterError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didEncounterError:"), device, error_)
}

// Tells the delegate when the ejection is complete.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEjectWithError:)
func (o ICCameraDeviceDelegateObject) DeviceDidEjectWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didEjectWithError:"), device, error_)
}

// Tells the delegate when the name of a device changes.
//
// # Discussion
//
// This happens if the device module overrides the default name of the device
// reported by the device’s transport layer, or if the name of the
// filesystem volume mounted by the device is changed by the user.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/deviceDidChangeName(_:)
func (o ICCameraDeviceDelegateObject) DeviceDidChangeName(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidChangeName:"), device)
}

// ICCameraDeviceDelegateConfig holds optional typed callbacks for [ICCameraDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/iccameradevicedelegate
type ICCameraDeviceDelegateConfig struct {

	// Removing Objects
	// CameraDeviceDidCompleteDeleteFilesWithError — Tells the client when the camera completes a delete operation.
	CameraDeviceDidCompleteDeleteFilesWithError func(camera ICCameraDevice, error_ foundation.NSError)

	// Responding to Capability Changes
	// CameraDeviceDidChangeCapability — Tells the client when a capability of a camera changes.
	CameraDeviceDidChangeCapability func(camera ICCameraDevice)

	// Responding to Access Restrictions
	// CameraDeviceDidEnableAccessRestriction — Tells the client when an Apple device has been locked, and media is unavailable until the restriction has been removed.
	CameraDeviceDidEnableAccessRestriction func(device ICDevice)
	// CameraDeviceDidRemoveAccessRestriction — Tells the client when an Apple device has been unlocked, paired to the host, and media is available.
	CameraDeviceDidRemoveAccessRestriction func(device ICDevice)

	// Responding to PTP Events
	// CameraDeviceDidReceivePTPEvent — Tells the client about a PTP event.
	CameraDeviceDidReceivePTPEvent func(camera ICCameraDevice, eventData foundation.NSData)

	// Other Methods
	// DeviceDidBecomeReadyWithCompleteContentCatalog — Tells the client that the camera device is done enumerating its content and is ready to receive requests.
	DeviceDidBecomeReadyWithCompleteContentCatalog func(device ICCameraDevice)
	// CameraDeviceDidReceiveMetadataForItemError — Tells the client when the metadata requested for an item on a camera is available.
	CameraDeviceDidReceiveMetadataForItemError func(camera ICCameraDevice, metadata foundation.INSDictionary, item ICCameraItem, error_ foundation.NSError)
	// CameraDeviceShouldGetMetadataOfItem — Tells the client when the camera is about to execute queued requests for the metadata of a specific item.
	CameraDeviceShouldGetMetadataOfItem func(cameraDevice ICCameraDevice, item ICCameraItem) bool
	// CameraDeviceShouldGetThumbnailOfItem — Tells the client when the camera is about to execute queued requests for the thumbnail of a specific item.
	CameraDeviceShouldGetThumbnailOfItem func(cameraDevice ICCameraDevice, item ICCameraItem) bool
}

// NewICCameraDeviceDelegate creates an Objective-C object implementing the [ICCameraDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ICCameraDeviceDelegateObject] satisfies the [ICCameraDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/iccameradevicedelegate
func NewICCameraDeviceDelegate(config ICCameraDeviceDelegateConfig) ICCameraDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoICCameraDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DeviceDidBecomeReadyWithCompleteContentCatalog != nil {
		fn := config.DeviceDidBecomeReadyWithCompleteContentCatalog
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceDidBecomeReadyWithCompleteContentCatalog:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "deviceDidBecomeReadyWithCompleteContentCatalog:")
					}
				}()
				device := ICCameraDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidReceiveMetadataForItemError != nil {
		fn := config.CameraDeviceDidReceiveMetadataForItemError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDevice:didReceiveMetadata:forItem:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraID objc.ID, metadataID objc.ID, itemID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDevice:didReceiveMetadata:forItem:error:")
					}
				}()
				camera := ICCameraDeviceFromID(cameraID)
				metadata := foundation.NSDictionaryFromID(metadataID)
				item := ICCameraItemFromID(itemID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(camera, metadata, item, error_)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidChangeCapability != nil {
		fn := config.CameraDeviceDidChangeCapability
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDeviceDidChangeCapability:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDeviceDidChangeCapability:")
					}
				}()
				camera := ICCameraDeviceFromID(cameraID)
				fn(camera)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidEnableAccessRestriction != nil {
		fn := config.CameraDeviceDidEnableAccessRestriction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDeviceDidEnableAccessRestriction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDeviceDidEnableAccessRestriction:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidRemoveAccessRestriction != nil {
		fn := config.CameraDeviceDidRemoveAccessRestriction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDeviceDidRemoveAccessRestriction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDeviceDidRemoveAccessRestriction:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidReceivePTPEvent != nil {
		fn := config.CameraDeviceDidReceivePTPEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDevice:didReceivePTPEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraID objc.ID, eventDataID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDevice:didReceivePTPEvent:")
					}
				}()
				camera := ICCameraDeviceFromID(cameraID)
				eventData := foundation.NSDataFromID(eventDataID)
				fn(camera, eventData)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceDidCompleteDeleteFilesWithError != nil {
		fn := config.CameraDeviceDidCompleteDeleteFilesWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDevice:didCompleteDeleteFilesWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDevice:didCompleteDeleteFilesWithError:")
					}
				}()
				camera := ICCameraDeviceFromID(cameraID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(camera, error_)
				_delegateDone = true
			},
		})
	}

	if config.CameraDeviceShouldGetMetadataOfItem != nil {
		fn := config.CameraDeviceShouldGetMetadataOfItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDevice:shouldGetMetadataOfItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraDeviceID objc.ID, itemID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDevice:shouldGetMetadataOfItem:")
					}
				}()
				cameraDevice := ICCameraDeviceFromID(cameraDeviceID)
				item := ICCameraItemFromID(itemID)
				_delegateResult := fn(cameraDevice, item)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.CameraDeviceShouldGetThumbnailOfItem != nil {
		fn := config.CameraDeviceShouldGetThumbnailOfItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cameraDevice:shouldGetThumbnailOfItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cameraDeviceID objc.ID, itemID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICCameraDeviceDelegate", "cameraDevice:shouldGetThumbnailOfItem:")
					}
				}()
				cameraDevice := ICCameraDeviceFromID(cameraDeviceID)
				item := ICCameraItemFromID(itemID)
				_delegateResult := fn(cameraDevice, item)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ICCameraDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewICCameraDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ICCameraDeviceDelegateObjectFromID(instance)
}
