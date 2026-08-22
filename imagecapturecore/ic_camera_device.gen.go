// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICCameraDevice] class.
var (
	_ICCameraDeviceClass     ICCameraDeviceClass
	_ICCameraDeviceClassOnce sync.Once
)

func getICCameraDeviceClass() ICCameraDeviceClass {
	_ICCameraDeviceClassOnce.Do(func() {
		_ICCameraDeviceClass = ICCameraDeviceClass{class: objc.GetClass("ICCameraDevice")}
	})
	return _ICCameraDeviceClass
}

// GetICCameraDeviceClass returns the class object for ICCameraDevice.
func GetICCameraDeviceClass() ICCameraDeviceClass {
	return getICCameraDeviceClass()
}

type ICCameraDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICCameraDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICCameraDeviceClass) Alloc() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a camera.
//
// # Reading Files
//
//   - [ICCameraDevice.Contents]: All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//   - [ICCameraDevice.MediaFiles]: All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//   - [ICCameraDevice.ContentCatalogPercentCompleted]: The percentage of the camera’s content that has been catalogued.
//   - [ICCameraDevice.FilesOfType]: Returns an array of files of the selected type on the camera.
//   - [ICCameraDevice.RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo]: Asynchronously reads data of a specified length from a specified offset.
//
// # Downloading Files
//
//   - [ICCameraDevice.CancelDownload]: Cancels a download from the camera.
//   - [ICCameraDevice.RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo]: Downloads a file from the camera.
//
// # Deleting Files
//
//   - [ICCameraDevice.IsLocked]: A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//   - [ICCameraDevice.RequestDeleteFiles]: Deletes files from the camera.
//   - [ICCameraDevice.CancelDelete]: Cancels the current delete operation.
//
// # Taking Pictures
//
//   - [ICCameraDevice.TetheredCaptureEnabled]: A Boolean value indicating whether tethered capture is enabled on the camera.
//   - [ICCameraDevice.PtpEventHandler]: A closure for handling PTP event packets.
//   - [ICCameraDevice.SetPtpEventHandler]
//   - [ICCameraDevice.RequestTakePicture]: Captures a new image using the camera.
//   - [ICCameraDevice.RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo]: Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//   - [ICCameraDevice.RequestSendPTPCommandOutDataCompletion]: Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// # Inspecting the Battery Charge Level
//
//   - [ICCameraDevice.BatteryLevelAvailable]: A Boolean value that indicates whether the battery charge level is available.
//   - [ICCameraDevice.BatteryLevel]: The battery charge level.
//
// # Synchronizing the Clock
//
//   - [ICCameraDevice.TimeOffset]: The time offset, in seconds, between the camera’s clock and the computer’s clock.
//   - [ICCameraDevice.RequestSyncClock]: Synchronizes the camera’s clock with the computer’s clock.
//
// # Detecting Apple Devices
//
//   - [ICCameraDevice.IsAccessRestrictedAppleDevice]: A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//   - [ICCameraDevice.ICloudPhotosEnabled]: A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// # Detecting Mass Storage Devices
//
//   - [ICCameraDevice.MountPoint]: The file system mount point for a camera using the mass storage transport type.
//
// # Removing a Device
//
//   - [ICCameraDevice.IsEjectable]: A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// # Instance Properties
//
//   - [ICCameraDevice.MediaPresentation]
//   - [ICCameraDevice.SetMediaPresentation]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice
type ICCameraDevice struct {
	ICDevice
}

// ICCameraDeviceFromID constructs a [ICCameraDevice] from an objc.ID.
//
// An object that represents a camera.
func ICCameraDeviceFromID(id objc.ID) ICCameraDevice {
	return ICCameraDevice{ICDevice: ICDeviceFromID(id)}
}

// NOTE: ICCameraDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICCameraDevice] class.
//
// # Reading Files
//
//   - [IICCameraDevice.Contents]: All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//   - [IICCameraDevice.MediaFiles]: All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//   - [IICCameraDevice.ContentCatalogPercentCompleted]: The percentage of the camera’s content that has been catalogued.
//   - [IICCameraDevice.FilesOfType]: Returns an array of files of the selected type on the camera.
//   - [IICCameraDevice.RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo]: Asynchronously reads data of a specified length from a specified offset.
//
// # Downloading Files
//
//   - [IICCameraDevice.CancelDownload]: Cancels a download from the camera.
//   - [IICCameraDevice.RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo]: Downloads a file from the camera.
//
// # Deleting Files
//
//   - [IICCameraDevice.IsLocked]: A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//   - [IICCameraDevice.RequestDeleteFiles]: Deletes files from the camera.
//   - [IICCameraDevice.CancelDelete]: Cancels the current delete operation.
//
// # Taking Pictures
//
//   - [IICCameraDevice.TetheredCaptureEnabled]: A Boolean value indicating whether tethered capture is enabled on the camera.
//   - [IICCameraDevice.PtpEventHandler]: A closure for handling PTP event packets.
//   - [IICCameraDevice.SetPtpEventHandler]
//   - [IICCameraDevice.RequestTakePicture]: Captures a new image using the camera.
//   - [IICCameraDevice.RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo]: Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//   - [IICCameraDevice.RequestSendPTPCommandOutDataCompletion]: Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// # Inspecting the Battery Charge Level
//
//   - [IICCameraDevice.BatteryLevelAvailable]: A Boolean value that indicates whether the battery charge level is available.
//   - [IICCameraDevice.BatteryLevel]: The battery charge level.
//
// # Synchronizing the Clock
//
//   - [IICCameraDevice.TimeOffset]: The time offset, in seconds, between the camera’s clock and the computer’s clock.
//   - [IICCameraDevice.RequestSyncClock]: Synchronizes the camera’s clock with the computer’s clock.
//
// # Detecting Apple Devices
//
//   - [IICCameraDevice.IsAccessRestrictedAppleDevice]: A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//   - [IICCameraDevice.ICloudPhotosEnabled]: A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// # Detecting Mass Storage Devices
//
//   - [IICCameraDevice.MountPoint]: The file system mount point for a camera using the mass storage transport type.
//
// # Removing a Device
//
//   - [IICCameraDevice.IsEjectable]: A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// # Instance Properties
//
//   - [IICCameraDevice.MediaPresentation]
//   - [IICCameraDevice.SetMediaPresentation]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice
type IICCameraDevice interface {
	IICDevice

	// Topic: Reading Files

	// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
	Contents() []ICCameraItem
	// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
	MediaFiles() []ICCameraItem
	// The percentage of the camera’s content that has been catalogued.
	ContentCatalogPercentCompleted() uint
	// Returns an array of files of the selected type on the camera.
	FilesOfType(fileUTType string) []string
	// Asynchronously reads data of a specified length from a specified offset.
	RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo(file IICCameraFile, offset int64, length int64, readDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr)

	// Topic: Downloading Files

	// Cancels a download from the camera.
	CancelDownload()
	// Downloads a file from the camera.
	RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo(file IICCameraFile, options foundation.INSDictionary, downloadDelegate ICCameraDeviceDownloadDelegate, selector objc.SEL, contextInfo uintptr)

	// Topic: Deleting Files

	// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
	IsLocked() bool
	// Deletes files from the camera.
	RequestDeleteFiles(files []ICCameraItem)
	// Cancels the current delete operation.
	CancelDelete()

	// Topic: Taking Pictures

	// A Boolean value indicating whether tethered capture is enabled on the camera.
	TetheredCaptureEnabled() bool
	// A closure for handling PTP event packets.
	PtpEventHandler() DataHandler
	SetPtpEventHandler(value DataHandler)
	// Captures a new image using the camera.
	RequestTakePicture()
	// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
	RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo(command foundation.NSData, data foundation.NSData, sendCommandDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr)
	// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
	RequestSendPTPCommandOutDataCompletion(ptpCommand foundation.NSData, ptpData foundation.NSData, completion DataDataErrorHandler)

	// Topic: Inspecting the Battery Charge Level

	// A Boolean value that indicates whether the battery charge level is available.
	BatteryLevelAvailable() bool
	// The battery charge level.
	BatteryLevel() uint

	// Topic: Synchronizing the Clock

	// The time offset, in seconds, between the camera’s clock and the computer’s clock.
	TimeOffset() foundation.NSTimeInterval
	// Synchronizes the camera’s clock with the computer’s clock.
	RequestSyncClock()

	// Topic: Detecting Apple Devices

	// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
	IsAccessRestrictedAppleDevice() bool
	// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
	ICloudPhotosEnabled() bool

	// Topic: Detecting Mass Storage Devices

	// The file system mount point for a camera using the mass storage transport type.
	MountPoint() string

	// Topic: Removing a Device

	// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
	IsEjectable() bool

	// Topic: Instance Properties

	MediaPresentation() ICMediaPresentation
	SetMediaPresentation(value ICMediaPresentation)
}

// Init initializes the instance.
func (c ICCameraDevice) Init() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ICCameraDevice) Autorelease() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraDevice creates a new ICCameraDevice instance.
func NewICCameraDevice() ICCameraDevice {
	class := getICCameraDeviceClass()
	rv := objc.Send[ICCameraDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an array of files of the selected type on the camera.
//
// # Discussion
//
// For the `fileType` parameter, pass one of the following uniform type
// identifier strings: `kUTTypeImage`, `kUTTypeMovie`, `kUTTypeAudio`, or
// `kUTTypeData`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/files(ofType:)
func (c ICCameraDevice) FilesOfType(fileUTType string) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("filesOfType:"), objc.String(fileUTType))
	return objc.ConvertSliceToStrings(rv)
}

// Asynchronously reads data of a specified length from a specified offset.
//
// # Discussion
//
// The `readDelegate` must implement a function with the signature
// `didReadData(NSData, ICCameraFile, NSError, Any)`, to be called when the
// request is completed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestReadData(from:atOffset:length:readDelegate:didReadDataSelector:contextInfo:)
func (c ICCameraDevice) RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo(file IICCameraFile, offset int64, length int64, readDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("requestReadDataFromFile:atOffset:length:readDelegate:didReadDataSelector:contextInfo:"), file, offset, length, readDelegate, selector, contextInfo)
}

// Cancels a download from the camera.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/cancelDownload()
func (c ICCameraDevice) CancelDownload() {
	objc.Send[objc.ID](c.ID, objc.Sel("cancelDownload"))
}

// Downloads a file from the camera.
//
// # Discussion
//
// Once this request completes, [DidDownloadFileErrorOptionsContextInfo] is
// called on the `downloadDelegate`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDownloadFile(_:options:downloadDelegate:didDownloadSelector:contextInfo:)
func (c ICCameraDevice) RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo(file IICCameraFile, options foundation.INSDictionary, downloadDelegate ICCameraDeviceDownloadDelegate, selector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("requestDownloadFile:options:downloadDelegate:didDownloadSelector:contextInfo:"), file, options, downloadDelegate, selector, contextInfo)
}

// Deletes files from the camera.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDeleteFiles(_:)
func (c ICCameraDevice) RequestDeleteFiles(files []ICCameraItem) {
	objc.Send[objc.ID](c.ID, objc.Sel("requestDeleteFiles:"), files)
}

// Cancels the current delete operation.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/cancelDelete()
func (c ICCameraDevice) CancelDelete() {
	objc.Send[objc.ID](c.ID, objc.Sel("cancelDelete"))
}

// Captures a new image using the camera.
//
// # Discussion
//
// Before taking a picture, you must first enable tethering by calling
// [ICCameraDevice.RequestEnableTethering].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestTakePicture()
func (c ICCameraDevice) RequestTakePicture() {
	objc.Send[objc.ID](c.ID, objc.Sel("requestTakePicture"))
}

// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// # Discussion
//
// Call this method only if the [ICDevice.Capabilities] property contains
// [cameraDeviceCanAcceptPTPCommands]. All PTP cameras have this capability.
//
// The `sendCommandDelegate` must implement a function with the signature `-
// (void)(NSData*)command (NSData*)data (NSData*)response (NSError*)error
// (void*)contextInfo`, to be called when the request is completed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSendPTPCommand(_:outData:sendCommandDelegate:didSendCommand:contextInfo:)
//
// [cameraDeviceCanAcceptPTPCommands]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanAcceptPTPCommands
func (c ICCameraDevice) RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo(command foundation.NSData, data foundation.NSData, sendCommandDelegate objectivec.IObject, selector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("requestSendPTPCommand:outData:sendCommandDelegate:didSendCommandSelector:contextInfo:"), command, data, sendCommandDelegate, selector, contextInfo)
}

// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// # Discussion
//
// The block receives the response, data, and an error message, if present.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSendPTPCommand(_:outData:completion:)
func (c ICCameraDevice) RequestSendPTPCommandOutDataCompletion(ptpCommand foundation.NSData, ptpData foundation.NSData, completion DataDataErrorHandler) {
	_block2, _ := NewDataDataErrorBlock(completion)
	objc.Send[objc.ID](c.ID, objc.Sel("requestSendPTPCommand:outData:completion:"), ptpCommand, ptpData, _block2)
}

// Synchronizes the camera’s clock with the computer’s clock.
//
// # Discussion
//
// Send this request only if the camera has the [cameraDeviceCanSyncClock]
// capability.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSyncClock()
//
// [cameraDeviceCanSyncClock]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanSyncClock
func (c ICCameraDevice) RequestSyncClock() {
	objc.Send[objc.ID](c.ID, objc.Sel("requestSyncClock"))
}

// All image, movie, and audio files stored on the camera, in an order that
// reflects the camera’s storage folder structure.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/contents
func (c ICCameraDevice) Contents() []ICCameraItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contents"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICCameraItem {
		return ICCameraItemFromID(id)
	})
}

// All image, movie and audio files stored on the camera, without regard to
// the camera’s storage folder structure.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/mediaFiles
func (c ICCameraDevice) MediaFiles() []ICCameraItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("mediaFiles"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICCameraItem {
		return ICCameraItemFromID(id)
	})
}

// The percentage of the camera’s content that has been catalogued.
//
// # Discussion
//
// The value of this property ranges from 0 to 100.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/contentCatalogPercentCompleted
func (c ICCameraDevice) ContentCatalogPercentCompleted() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("contentCatalogPercentCompleted"))
	return rv
}

// A Boolean value indicating whether the device is locked, preventing
// deletion of any asset.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isLocked
func (c ICCameraDevice) IsLocked() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLocked"))
	return rv
}

// A Boolean value indicating whether tethered capture is enabled on the
// camera.
//
// # Discussion
//
// Use [ICCameraDevice.RequestEnableTethering] and
// [ICCameraDevice.RequestDisableTethering] to enable or disable tethered
// capture.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/tetheredCaptureEnabled
func (c ICCameraDevice) TetheredCaptureEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("tetheredCaptureEnabled"))
	return rv
}

// A closure for handling PTP event packets.
//
// # Discussion
//
// Set this property as an alternative to setting up an object to handle PTP
// event packets. If the handler is set, it is called in place of the
// [ICDevice.Delegate]. If the handler is `nil`, the [ICDevice.Delegate] is
// called, if present. If both are set, only the handler is called.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/ptpEventHandler
func (c ICCameraDevice) PtpEventHandler() DataHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ptpEventHandler"))
	_ = rv
	return nil
}
func (c ICCameraDevice) SetPtpEventHandler(value DataHandler) {
	block, cleanup := NewDataBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setPtpEventHandler:"), block)
}

// A Boolean value that indicates whether the battery charge level is
// available.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/batteryLevelAvailable
func (c ICCameraDevice) BatteryLevelAvailable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("batteryLevelAvailable"))
	return rv
}

// The battery charge level.
//
// # Discussion
//
// The value of this property ranges from 0 to 100.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/batteryLevel
func (c ICCameraDevice) BatteryLevel() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("batteryLevel"))
	return rv
}

// The time offset, in seconds, between the camera’s clock and the
// computer’s clock.
//
// # Discussion
//
// The value of this property is positive if the camera’s clock is ahead of
// the computer’s clock. Ignore this property if the camera’s
// [ICDevice.Capabilities] do not include [cameraDeviceCanSyncClock].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/timeOffset
//
// [cameraDeviceCanSyncClock]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanSyncClock
func (c ICCameraDevice) TimeOffset() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](c.ID, objc.Sel("timeOffset"))
	return foundation.NSTimeInterval(rv)
}

// A Boolean value indicating whether the device is an Apple device,
// passcode-locked, and connected to an untrusted host.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isAccessRestrictedAppleDevice
func (c ICCameraDevice) IsAccessRestrictedAppleDevice() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAccessRestrictedAppleDevice"))
	return rv
}

// A Boolean value indicating whether the iCloud Photo Library is enabled on
// the device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/iCloudPhotosEnabled
func (c ICCameraDevice) ICloudPhotosEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("iCloudPhotosEnabled"))
	return rv
}

// The file system mount point for a camera using the mass storage transport
// type.
//
// # Discussion
//
// This property is set for cameras whose [ICDevice.TransportType] is
// [transportTypeMassStorage].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/mountPoint
//
// [transportTypeMassStorage]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeMassStorage
func (c ICCameraDevice) MountPoint() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mountPoint"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value indicating whether the device can be ‘soft’ removed or
// disconnected.
//
// # Discussion
//
// Soft ejecting an SD card is equivalent to unmounting it in Finder without
// physically removing it from the host.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isEjectable
func (c ICCameraDevice) IsEjectable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEjectable"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/mediaPresentation
func (c ICCameraDevice) MediaPresentation() ICMediaPresentation {
	rv := objc.Send[ICMediaPresentation](c.ID, objc.Sel("mediaPresentation"))
	return ICMediaPresentation(rv)
}
func (c ICCameraDevice) SetMediaPresentation(value ICMediaPresentation) {
	objc.Send[struct{}](c.ID, objc.Sel("setMediaPresentation:"), value)
}
