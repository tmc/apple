// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"context"
	"sync"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICDevice] class.
var (
	_ICDeviceClass     ICDeviceClass
	_ICDeviceClassOnce sync.Once
)

func getICDeviceClass() ICDeviceClass {
	_ICDeviceClassOnce.Do(func() {
		_ICDeviceClass = ICDeviceClass{class: objc.GetClass("ICDevice")}
	})
	return _ICDeviceClass
}

// GetICDeviceClass returns the class object for ICDevice.
func GetICDeviceClass() ICDeviceClass {
	return getICDeviceClass()
}

type ICDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICDeviceClass) Alloc() ICDevice {
	rv := objc.Send[ICDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An abstract object that represents a device.
//
// # Overview
//
// The device browser uses the concrete subclasses [ICCameraDevice] and
// [ICScannerDevice] to represent the cameras and scanners it finds.
//
// # Identifying a Device
//
//   - [ICDevice.Name]: The device’s name as reported by the device module, or if no device module is in control of this device, by the device transport.
//   - [ICDevice.ProductKind]: The device’s type.
//   - [ICDevice.Icon]: The device’s icon image.
//   - [ICDevice.UUIDString]: A string representation of the device’s universally unique identifier (UUID).
//   - [ICDevice.PersistentIDString]: A string representation of the device’s persistent ID.
//   - [ICDevice.SerialNumberString]: The device’s serial number.
//
// # Inspecting a Device’s Type and Location
//
//   - [ICDevice.Type]: A combination of the device’s type and its location type.
//   - [ICDevice.LocationDescription]: A nonlocalized location description for the device.
//   - [ICDevice.ModulePath]: The file system path of the device module associated with this device.
//   - [ICDevice.ModuleVersion]: The bundle version of the device module associated with this device.
//   - [ICDevice.UsbLocationID]: The USB location that the device is occupying.
//   - [ICDevice.UsbProductID]: The USB Product ID (PID) associated with the device.
//   - [ICDevice.UsbVendorID]: The USB Vendor ID (VID) associated with the device.
//
// # Inspecting a Device’s Transport Type
//
//   - [ICDevice.TransportType]: The hardware connection type the device is using.
//
// # Inspecting a Device’s Capabilities
//
//   - [ICDevice.Capabilities]: The capabilities of the device as reported by the device module.
//
// # Managing a Device
//
//   - [ICDevice.Delegate]: The delegate to receive messages once a session is opened on the device.
//   - [ICDevice.SetDelegate]
//   - [ICDevice.HasOpenSession]: A Boolean value that indicates whether the device has an open session.
//   - [ICDevice.RequestOpenSession]: Requests to open a session on the device.
//   - [ICDevice.RequestOpenSessionWithOptionsCompletion]: Requests to open a session on the device, then executes the completion handler.
//   - [ICDevice.RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo]: Asynchronously sends an arbitrary message with optional data to a device.
//   - [ICDevice.RequestCloseSession]: Requests to close an open session on the device.
//   - [ICDevice.RequestCloseSessionWithOptionsCompletion]: Requests to close an open session on the device, then executes the completion handler.
//   - [ICDevice.RequestEject]: Requests to eject the media if permitted by the device, or to disconnect from a remote device.
//   - [ICDevice.RequestEjectWithCompletion]: Requests to eject the media if permitted by the device, or to disconnect from a remote device, then executes the completion handler.
//
// # Configuring a Device’s Characteristics
//
//   - [ICDevice.UserData]: A bookkeeping object for client convenience.
//   - [ICDevice.AutolaunchApplicationPath]: The file system path of an application to launch automatically when this device is added.
//   - [ICDevice.SetAutolaunchApplicationPath]
//   - [ICDevice.IsRemote]: A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
//
// # Instance Properties
//
//   - [ICDevice.SystemSymbolName]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice
type ICDevice struct {
	objectivec.Object
}

// ICDeviceFromID constructs a [ICDevice] from an objc.ID.
//
// An abstract object that represents a device.
func ICDeviceFromID(id objc.ID) ICDevice {
	return ICDevice{objectivec.Object{ID: id}}
}

// NOTE: ICDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICDevice] class.
//
// # Identifying a Device
//
//   - [IICDevice.Name]: The device’s name as reported by the device module, or if no device module is in control of this device, by the device transport.
//   - [IICDevice.ProductKind]: The device’s type.
//   - [IICDevice.Icon]: The device’s icon image.
//   - [IICDevice.UUIDString]: A string representation of the device’s universally unique identifier (UUID).
//   - [IICDevice.PersistentIDString]: A string representation of the device’s persistent ID.
//   - [IICDevice.SerialNumberString]: The device’s serial number.
//
// # Inspecting a Device’s Type and Location
//
//   - [IICDevice.Type]: A combination of the device’s type and its location type.
//   - [IICDevice.LocationDescription]: A nonlocalized location description for the device.
//   - [IICDevice.ModulePath]: The file system path of the device module associated with this device.
//   - [IICDevice.ModuleVersion]: The bundle version of the device module associated with this device.
//   - [IICDevice.UsbLocationID]: The USB location that the device is occupying.
//   - [IICDevice.UsbProductID]: The USB Product ID (PID) associated with the device.
//   - [IICDevice.UsbVendorID]: The USB Vendor ID (VID) associated with the device.
//
// # Inspecting a Device’s Transport Type
//
//   - [IICDevice.TransportType]: The hardware connection type the device is using.
//
// # Inspecting a Device’s Capabilities
//
//   - [IICDevice.Capabilities]: The capabilities of the device as reported by the device module.
//
// # Managing a Device
//
//   - [IICDevice.Delegate]: The delegate to receive messages once a session is opened on the device.
//   - [IICDevice.SetDelegate]
//   - [IICDevice.HasOpenSession]: A Boolean value that indicates whether the device has an open session.
//   - [IICDevice.RequestOpenSession]: Requests to open a session on the device.
//   - [IICDevice.RequestOpenSessionWithOptionsCompletion]: Requests to open a session on the device, then executes the completion handler.
//   - [IICDevice.RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo]: Asynchronously sends an arbitrary message with optional data to a device.
//   - [IICDevice.RequestCloseSession]: Requests to close an open session on the device.
//   - [IICDevice.RequestCloseSessionWithOptionsCompletion]: Requests to close an open session on the device, then executes the completion handler.
//   - [IICDevice.RequestEject]: Requests to eject the media if permitted by the device, or to disconnect from a remote device.
//   - [IICDevice.RequestEjectWithCompletion]: Requests to eject the media if permitted by the device, or to disconnect from a remote device, then executes the completion handler.
//
// # Configuring a Device’s Characteristics
//
//   - [IICDevice.UserData]: A bookkeeping object for client convenience.
//   - [IICDevice.AutolaunchApplicationPath]: The file system path of an application to launch automatically when this device is added.
//   - [IICDevice.SetAutolaunchApplicationPath]
//   - [IICDevice.IsRemote]: A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
//
// # Instance Properties
//
//   - [IICDevice.SystemSymbolName]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice
type IICDevice interface {
	objectivec.IObject

	// Topic: Identifying a Device

	// The device’s name as reported by the device module, or if no device module is in control of this device, by the device transport.
	Name() string
	// The device’s type.
	ProductKind() string
	// The device’s icon image.
	Icon() coregraphics.CGImageRef
	// A string representation of the device’s universally unique identifier (UUID).
	UUIDString() string
	// A string representation of the device’s persistent ID.
	PersistentIDString() string
	// The device’s serial number.
	SerialNumberString() string

	// Topic: Inspecting a Device’s Type and Location

	// A combination of the device’s type and its location type.
	Type() ICDeviceType
	// A nonlocalized location description for the device.
	LocationDescription() string
	// The file system path of the device module associated with this device.
	ModulePath() string
	// The bundle version of the device module associated with this device.
	ModuleVersion() string
	// The USB location that the device is occupying.
	UsbLocationID() int32
	// The USB Product ID (PID) associated with the device.
	UsbProductID() int32
	// The USB Vendor ID (VID) associated with the device.
	UsbVendorID() int32

	// Topic: Inspecting a Device’s Transport Type

	// The hardware connection type the device is using.
	TransportType() string

	// Topic: Inspecting a Device’s Capabilities

	// The capabilities of the device as reported by the device module.
	Capabilities() []string

	// Topic: Managing a Device

	// The delegate to receive messages once a session is opened on the device.
	Delegate() ICDeviceDelegate
	SetDelegate(value ICDeviceDelegate)
	// A Boolean value that indicates whether the device has an open session.
	HasOpenSession() bool
	// Requests to open a session on the device.
	RequestOpenSession()
	// Requests to open a session on the device, then executes the completion handler.
	RequestOpenSessionWithOptionsCompletion(options foundation.INSDictionary, completion ErrorHandler)
	// Asynchronously sends an arbitrary message with optional data to a device.
	RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo(messageCode uint32, data foundation.NSData, maxReturnedDataSize uint32, sendMessageDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr)
	// Requests to close an open session on the device.
	RequestCloseSession()
	// Requests to close an open session on the device, then executes the completion handler.
	RequestCloseSessionWithOptionsCompletion(options foundation.INSDictionary, completion ErrorHandler)
	// Requests to eject the media if permitted by the device, or to disconnect from a remote device.
	RequestEject()
	// Requests to eject the media if permitted by the device, or to disconnect from a remote device, then executes the completion handler.
	RequestEjectWithCompletion(completion ErrorHandler)

	// Topic: Configuring a Device’s Characteristics

	// A bookkeeping object for client convenience.
	UserData() foundation.INSDictionary
	// The file system path of an application to launch automatically when this device is added.
	AutolaunchApplicationPath() string
	SetAutolaunchApplicationPath(value string)
	// A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
	IsRemote() bool

	// Topic: Instance Properties

	SystemSymbolName() string
}

// Init initializes the instance.
func (d ICDevice) Init() ICDevice {
	rv := objc.Send[ICDevice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d ICDevice) Autorelease() ICDevice {
	rv := objc.Send[ICDevice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewICDevice creates a new ICDevice instance.
func NewICDevice() ICDevice {
	class := getICDeviceClass()
	rv := objc.Send[ICDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests to open a session on the device.
//
// # Discussion
//
// Set the receiver’s delegate before calling this method; otherwise, the
// request is ignored.
//
// Once the request to open the session has completed,
// [DeviceDidOpenSessionWithError] is called on the delegate.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestOpenSession()
func (d ICDevice) RequestOpenSession() {
	objc.Send[objc.ID](d.ID, objc.Sel("requestOpenSession"))
}

// Requests to open a session on the device, then executes the completion
// handler.
//
// # Discussion
//
// Execution of the completion block occurs on the calling thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestOpenSession(options:completion:)
func (d ICDevice) RequestOpenSessionWithOptionsCompletion(options foundation.INSDictionary, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("requestOpenSessionWithOptions:completion:"), options, _block1)
}

// Asynchronously sends an arbitrary message with optional data to a device.
//
// # Discussion
//
// Use this method to send a private message from a client application to a
// device module.
//
// The `sendMessageDelegate` must implement a function with the signature
// `(void)(UInt32)messageCode (NSData*)data (NSError*)error
// (void*)contextInfo`, to be called when the request is completed.
//
// Do not use this method to send PTP pass-through commands to a PTP camera.
// Use
// [ICCameraDevice.RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo]
// instead.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestSendMessage(_:outData:maxReturnedDataSize:sendMessageDelegate:didSendMessageSelector:contextInfo:)
func (d ICDevice) RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo(messageCode uint32, data foundation.NSData, maxReturnedDataSize uint32, sendMessageDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("requestSendMessage:outData:maxReturnedDataSize:sendMessageDelegate:didSendMessageSelector:contextInfo:"), messageCode, data, maxReturnedDataSize, sendMessageDelegate, selector, contextInfo)
}

// Requests to close an open session on the device.
//
// # Discussion
//
// Once the request to close the session has completed,
// [DeviceDidCloseSessionWithError] is called on the delegate.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestCloseSession()
func (d ICDevice) RequestCloseSession() {
	objc.Send[objc.ID](d.ID, objc.Sel("requestCloseSession"))
}

// Requests to close an open session on the device, then executes the
// completion handler.
//
// # Discussion
//
// Execution of the completion block occurs on the calling thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestCloseSession(options:completion:)
func (d ICDevice) RequestCloseSessionWithOptionsCompletion(options foundation.INSDictionary, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("requestCloseSessionWithOptions:completion:"), options, _block1)
}

// Requests to eject the media if permitted by the device, or to disconnect
// from a remote device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestEject()
func (d ICDevice) RequestEject() {
	objc.Send[objc.ID](d.ID, objc.Sel("requestEject"))
}

// Requests to eject the media if permitted by the device, or to disconnect
// from a remote device, then executes the completion handler.
//
// # Discussion
//
// Execution of the completion block occurs on the calling thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestEject(completion:)
func (d ICDevice) RequestEjectWithCompletion(completion ErrorHandler) {
	_block0, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("requestEjectWithCompletion:"), _block0)
}

// The device’s name as reported by the device module, or if no device
// module is in control of this device, by the device transport.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/name
func (d ICDevice) Name() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The device’s type.
//
// # Discussion
//
// Possible values are `"iPhone"`, `"iPod"`, and `"Camera"`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/productKind
func (d ICDevice) ProductKind() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("productKind"))
	return foundation.NSStringFromID(rv).String()
}

// The device’s icon image.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/icon
func (d ICDevice) Icon() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](d.ID, objc.Sel("icon"))
	return coregraphics.CGImageRef(rv)
}

// A string representation of the device’s universally unique identifier
// (UUID).
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/uuidString
func (d ICDevice) UUIDString() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("UUIDString"))
	return foundation.NSStringFromID(rv).String()
}

// A string representation of the device’s persistent ID.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/persistentIDString
func (d ICDevice) PersistentIDString() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("persistentIDString"))
	return foundation.NSStringFromID(rv).String()
}

// The device’s serial number.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/serialNumberString
func (d ICDevice) SerialNumberString() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("serialNumberString"))
	return foundation.NSStringFromID(rv).String()
}

// A combination of the device’s type and its location type.
//
// # Discussion
//
// This property combines the device’s type, for example
// [ICDeviceType.camera], with its location type, for example
// [ICDeviceLocationType.bluetooth].
//
// To isolate the device’s type, perform bitwise [AND] on this property with
// an [ICDeviceTypeMask].
//
// To isolate the device’s location type, perform bitwise [AND] on this
// property with an [ICDeviceLocationTypeMask].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/type
//
// [ICDeviceLocationType.bluetooth]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType/bluetooth
// [ICDeviceLocationTypeMask]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask
// [ICDeviceType.camera]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceType/camera
// [ICDeviceTypeMask]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTypeMask
func (d ICDevice) Type() ICDeviceType {
	rv := objc.Send[ICDeviceType](d.ID, objc.Sel("type"))
	return ICDeviceType(rv)
}

// A nonlocalized location description for the device.
//
// # Discussion
//
// This property returns either the description of an [ICDeviceLocationType],
// or a description obtained from the Bonjour [TXT] record of a network
// device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/locationDescription
//
// [ICDeviceLocationType]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType
func (d ICDevice) LocationDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("locationDescription"))
	return foundation.NSStringFromID(rv).String()
}

// The file system path of the device module associated with this device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/modulePath
func (d ICDevice) ModulePath() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("modulePath"))
	return foundation.NSStringFromID(rv).String()
}

// The bundle version of the device module associated with this device.
//
// # Discussion
//
// The bundle version may change if an existing device module associated with
// this device is updated or a new device module for this device is installed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/moduleVersion
func (d ICDevice) ModuleVersion() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("moduleVersion"))
	return foundation.NSStringFromID(rv).String()
}

// The USB location that the device is occupying.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/usbLocationID
func (d ICDevice) UsbLocationID() int32 {
	rv := objc.Send[int32](d.ID, objc.Sel("usbLocationID"))
	return rv
}

// The USB Product ID (PID) associated with the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/usbProductID
func (d ICDevice) UsbProductID() int32 {
	rv := objc.Send[int32](d.ID, objc.Sel("usbProductID"))
	return rv
}

// The USB Vendor ID (VID) associated with the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/usbVendorID
func (d ICDevice) UsbVendorID() int32 {
	rv := objc.Send[int32](d.ID, objc.Sel("usbVendorID"))
	return rv
}

// The hardware connection type the device is using.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/transportType
func (d ICDevice) TransportType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("transportType"))
	return foundation.NSStringFromID(rv).String()
}

// The capabilities of the device as reported by the device module.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/capabilities
func (d ICDevice) Capabilities() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("capabilities"))
	return objc.ConvertSliceToStrings(rv)
}

// The delegate to receive messages once a session is opened on the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/delegate
func (d ICDevice) Delegate() ICDeviceDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return ICDeviceDelegateObjectFromID(rv)
}
func (d ICDevice) SetDelegate(value ICDeviceDelegate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the device has an open session.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/hasOpenSession
func (d ICDevice) HasOpenSession() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasOpenSession"))
	return rv
}

// A bookkeeping object for client convenience.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/userData
func (d ICDevice) UserData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("userData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The file system path of an application to launch automatically when this
// device is added.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/autolaunchApplicationPath
func (d ICDevice) AutolaunchApplicationPath() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("autolaunchApplicationPath"))
	return foundation.NSStringFromID(rv).String()
}
func (d ICDevice) SetAutolaunchApplicationPath(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setAutolaunchApplicationPath:"), objc.String(value))
}

// A Boolean value indicating whether the device is published by the Image
// Capture device-sharing facility.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/isRemote
func (d ICDevice) IsRemote() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isRemote"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/systemSymbolName
func (d ICDevice) SystemSymbolName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("systemSymbolName"))
	return foundation.NSStringFromID(rv).String()
}

// RequestOpenSessionWithOptionsCompletionSync is a synchronous wrapper around [ICDevice.RequestOpenSessionWithOptionsCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (d ICDevice) RequestOpenSessionWithOptionsCompletionSync(ctx context.Context, options foundation.INSDictionary) error {
	done := make(chan error, 1)
	d.RequestOpenSessionWithOptionsCompletion(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestCloseSessionWithOptionsCompletionSync is a synchronous wrapper around [ICDevice.RequestCloseSessionWithOptionsCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (d ICDevice) RequestCloseSessionWithOptionsCompletionSync(ctx context.Context, options foundation.INSDictionary) error {
	done := make(chan error, 1)
	d.RequestCloseSessionWithOptionsCompletion(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestEjectSync is a synchronous wrapper around [ICDevice.RequestEjectWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (d ICDevice) RequestEjectSync(ctx context.Context) error {
	done := make(chan error, 1)
	d.RequestEjectWithCompletion(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
