// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for managing the addition and removal of devices and responding to device changes.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate
type ICDeviceBrowserDelegate interface {
	objectivec.IObject

	// Tells the delegate that a device has been added.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:didAdd:moreComing:)
	DeviceBrowserDidAddDeviceMoreComing(browser IICDeviceBrowser, device IICDevice, moreComing bool)

	// Tells the delegate that a device has been removed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:didRemove:moreGoing:)
	DeviceBrowserDidRemoveDeviceMoreGoing(browser IICDeviceBrowser, device IICDevice, moreGoing bool)
}

// ICDeviceBrowserDelegateObject wraps an existing Objective-C object that conforms to the ICDeviceBrowserDelegate protocol.
type ICDeviceBrowserDelegateObject struct {
	objectivec.Object
}

func (o ICDeviceBrowserDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ICDeviceBrowserDelegateObjectFromID constructs a [ICDeviceBrowserDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ICDeviceBrowserDelegateObjectFromID(id objc.ID) ICDeviceBrowserDelegateObject {
	return ICDeviceBrowserDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that a device has been added.
//
// # Discussion
//
// If several devices are found during the initial search, then this message
// is sent once for each device with the value of `moreComing` set to `true`
// in each message except the last one.
//
// Not all devices are reported using this method. Devices that fail to
// communicate successfully are silently ignored.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:didAdd:moreComing:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidAddDeviceMoreComing(browser IICDeviceBrowser, device IICDevice, moreComing bool) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowser:didAddDevice:moreComing:"), browser, device, moreComing)
}

// Tells the delegate that a device has been removed.
//
// # Discussion
//
// If several devices are removed at the same time, then this message is sent
// once for each device with the value of `moreGoing` set to `true` in each
// message except the last one.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:didRemove:moreGoing:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidRemoveDeviceMoreGoing(browser IICDeviceBrowser, device IICDevice, moreGoing bool) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowser:didRemoveDevice:moreGoing:"), browser, device, moreGoing)
}

// Tells the delegate that the device browser has completed sending
// [DeviceBrowserDidAddDeviceMoreComing] for all local devices.
//
// # Discussion
//
// Detecting locally connected devices (USB and FireWire devices) is faster
// than detecting devices connected using a network protocol. An Image Capture
// client application may use this message to update its user interface to let
// the user know that it has completed looking for locally connected devices
// and then started looking for network devices.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowserDidEnumerateLocalDevices(_:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidEnumerateLocalDevices(browser IICDeviceBrowser) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowserDidEnumerateLocalDevices:"), browser)
}

// Tells the delegate when an event occurs on the device that may be of
// interest to the client application.
//
// # Discussion
//
// This message is sent when a button is pressed on a device and the current
// application is the target for that button press. When this happens, if a
// session is open on the device, this message is not sent to the browser
// delegate; instead the message `device(_:)` is sent to the device delegate.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:requestsSelect:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserRequestsSelectDevice(browser IICDeviceBrowser, device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowser:requestsSelectDevice:"), browser, device)
}

// Tells the delegate when the name of a device changes.
//
// # Discussion
//
// A device’s name may change if a device module overrides the default name
// reported by the device’s transport layer, or if a user changes the name
// of the file system volume mounted by the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:deviceDidChangeName:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDeviceDidChangeName(browser IICDeviceBrowser, device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowser:deviceDidChangeName:"), browser, device)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowserDidCancelSuspendOperations(_:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidCancelSuspendOperations(browser IICDeviceBrowser) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowserDidCancelSuspendOperations:"), browser)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowserDidResumeOperations(_:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidResumeOperations(browser IICDeviceBrowser) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowserDidResumeOperations:"), browser)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowserDidSuspendOperations(_:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserDidSuspendOperations(browser IICDeviceBrowser) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowserDidSuspendOperations:"), browser)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowserWillSuspendOperations(_:)
func (o ICDeviceBrowserDelegateObject) DeviceBrowserWillSuspendOperations(browser IICDeviceBrowser) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceBrowserWillSuspendOperations:"), browser)
}

// ICDeviceBrowserDelegateConfig holds optional typed callbacks for [ICDeviceBrowserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowserdelegate
type ICDeviceBrowserDelegateConfig struct {

	// Adding and Removing Devices
	// DeviceBrowserDidEnumerateLocalDevices — Tells the delegate that the device browser has completed sending [deviceBrowser(_:didAdd:moreComing:)](<https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowserDelegate/deviceBrowser(_:didAdd:moreComing:)>) for all local devices.
	DeviceBrowserDidEnumerateLocalDevices func(browser ICDeviceBrowser)

	// Responding to Device Changes
	// DeviceBrowserDeviceDidChangeName — Tells the delegate when the name of a device changes.
	DeviceBrowserDeviceDidChangeName func(browser ICDeviceBrowser, device ICDevice)

	// Other Methods
	// DeviceBrowserDidAddDeviceMoreComing — Tells the delegate that a device has been added.
	DeviceBrowserDidAddDeviceMoreComing func(browser ICDeviceBrowser, device ICDevice, moreComing bool)
	// DeviceBrowserDidRemoveDeviceMoreGoing — Tells the delegate that a device has been removed.
	DeviceBrowserDidRemoveDeviceMoreGoing func(browser ICDeviceBrowser, device ICDevice, moreGoing bool)
	// DeviceBrowserRequestsSelectDevice — Tells the delegate when an event occurs on the device that may be of interest to the client application.
	DeviceBrowserRequestsSelectDevice func(browser ICDeviceBrowser, device ICDevice)
}

// NewICDeviceBrowserDelegate creates an Objective-C object implementing the [ICDeviceBrowserDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ICDeviceBrowserDelegateObject] satisfies the [ICDeviceBrowserDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowserdelegate
func NewICDeviceBrowserDelegate(config ICDeviceBrowserDelegateConfig) ICDeviceBrowserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoICDeviceBrowserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DeviceBrowserDidAddDeviceMoreComing != nil {
		fn := config.DeviceBrowserDidAddDeviceMoreComing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceBrowser:didAddDevice:moreComing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, deviceID objc.ID, moreComing bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceBrowserDelegate", "deviceBrowser:didAddDevice:moreComing:")
					}
				}()
				browser := ICDeviceBrowserFromID(browserID)
				device := ICDeviceFromID(deviceID)
				fn(browser, device, moreComing)
				_delegateDone = true
			},
		})
	}

	if config.DeviceBrowserDidRemoveDeviceMoreGoing != nil {
		fn := config.DeviceBrowserDidRemoveDeviceMoreGoing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceBrowser:didRemoveDevice:moreGoing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, deviceID objc.ID, moreGoing bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceBrowserDelegate", "deviceBrowser:didRemoveDevice:moreGoing:")
					}
				}()
				browser := ICDeviceBrowserFromID(browserID)
				device := ICDeviceFromID(deviceID)
				fn(browser, device, moreGoing)
				_delegateDone = true
			},
		})
	}

	if config.DeviceBrowserDidEnumerateLocalDevices != nil {
		fn := config.DeviceBrowserDidEnumerateLocalDevices
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceBrowserDidEnumerateLocalDevices:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceBrowserDelegate", "deviceBrowserDidEnumerateLocalDevices:")
					}
				}()
				browser := ICDeviceBrowserFromID(browserID)
				fn(browser)
				_delegateDone = true
			},
		})
	}

	if config.DeviceBrowserRequestsSelectDevice != nil {
		fn := config.DeviceBrowserRequestsSelectDevice
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceBrowser:requestsSelectDevice:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceBrowserDelegate", "deviceBrowser:requestsSelectDevice:")
					}
				}()
				browser := ICDeviceBrowserFromID(browserID)
				device := ICDeviceFromID(deviceID)
				fn(browser, device)
				_delegateDone = true
			},
		})
	}

	if config.DeviceBrowserDeviceDidChangeName != nil {
		fn := config.DeviceBrowserDeviceDidChangeName
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceBrowser:deviceDidChangeName:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceBrowserDelegate", "deviceBrowser:deviceDidChangeName:")
					}
				}()
				browser := ICDeviceBrowserFromID(browserID)
				device := ICDeviceFromID(deviceID)
				fn(browser, device)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ICDeviceBrowserDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewICDeviceBrowserDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ICDeviceBrowserDelegateObjectFromID(instance)
}
