// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerDevice] class.
var (
	_ICScannerDeviceClass     ICScannerDeviceClass
	_ICScannerDeviceClassOnce sync.Once
)

func getICScannerDeviceClass() ICScannerDeviceClass {
	_ICScannerDeviceClassOnce.Do(func() {
		_ICScannerDeviceClass = ICScannerDeviceClass{class: objc.GetClass("ICScannerDevice")}
	})
	return _ICScannerDeviceClass
}

// GetICScannerDeviceClass returns the class object for ICScannerDevice.
func GetICScannerDeviceClass() ICScannerDeviceClass {
	return getICScannerDeviceClass()
}

type ICScannerDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerDeviceClass) Alloc() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a scanner.
//
// # Overview
//
// An instance of ICScannerDevice class is intended to be used by the
// ICScannerDeviceView object. The ICScannerDeviceView class encapsulates the
// complexities of setting scan parameters, performing scans and saving the
// result. The developer should consider using ICScannerDeviceView instead of
// building their own views using the ICScannerDevice object.
//
// # Selecting a Functional Unit
//
//   - [ICScannerDevice.AvailableFunctionalUnitTypes]: An array of functional unit types available on this scanner.
//   - [ICScannerDevice.SelectedFunctionalUnit]: The currently selected functional unit on the scanner.
//   - [ICScannerDevice.RequestSelectFunctionalUnit]: Requests to select a functional unit on the scanner.
//
// # Performing a Scan
//
//   - [ICScannerDevice.RequestOpenSessionWithCredentialsPassword]: Opens a session on the protected device with the authorized username and passcode.
//   - [ICScannerDevice.RequestOverviewScan]: Starts an overview scan on the selected functional unit.
//   - [ICScannerDevice.RequestScan]: Starts a scan on the selected functional unit.
//   - [ICScannerDevice.CancelScan]: Cancels the current scan.
//   - [ICScannerDevice.DocumentName]: The document’s name.
//   - [ICScannerDevice.SetDocumentName]
//   - [ICScannerDevice.DocumentUTI]: The document’s uniform type identifier.
//   - [ICScannerDevice.SetDocumentUTI]
//   - [ICScannerDevice.DownloadsDirectory]: The downloads directory.
//   - [ICScannerDevice.SetDownloadsDirectory]
//   - [ICScannerDevice.TransferMode]: The transfer mode for the scanned document.
//   - [ICScannerDevice.SetTransferMode]
//   - [ICScannerDevice.MaxMemoryBandSize]: The total maximum band size requested when performing a memory-based transfer.
//   - [ICScannerDevice.SetMaxMemoryBandSize]
//
// # Logging into a Protected Device
//
//   - [ICScannerDevice.DefaultUsername]: A default username on protected scanners.
//   - [ICScannerDevice.SetDefaultUsername]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice
type ICScannerDevice struct {
	ICDevice
}

// ICScannerDeviceFromID constructs a [ICScannerDevice] from an objc.ID.
//
// An object that represents a scanner.
func ICScannerDeviceFromID(id objc.ID) ICScannerDevice {
	return ICScannerDevice{ICDevice: ICDeviceFromID(id)}
}

// NOTE: ICScannerDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerDevice] class.
//
// # Selecting a Functional Unit
//
//   - [IICScannerDevice.AvailableFunctionalUnitTypes]: An array of functional unit types available on this scanner.
//   - [IICScannerDevice.SelectedFunctionalUnit]: The currently selected functional unit on the scanner.
//   - [IICScannerDevice.RequestSelectFunctionalUnit]: Requests to select a functional unit on the scanner.
//
// # Performing a Scan
//
//   - [IICScannerDevice.RequestOpenSessionWithCredentialsPassword]: Opens a session on the protected device with the authorized username and passcode.
//   - [IICScannerDevice.RequestOverviewScan]: Starts an overview scan on the selected functional unit.
//   - [IICScannerDevice.RequestScan]: Starts a scan on the selected functional unit.
//   - [IICScannerDevice.CancelScan]: Cancels the current scan.
//   - [IICScannerDevice.DocumentName]: The document’s name.
//   - [IICScannerDevice.SetDocumentName]
//   - [IICScannerDevice.DocumentUTI]: The document’s uniform type identifier.
//   - [IICScannerDevice.SetDocumentUTI]
//   - [IICScannerDevice.DownloadsDirectory]: The downloads directory.
//   - [IICScannerDevice.SetDownloadsDirectory]
//   - [IICScannerDevice.TransferMode]: The transfer mode for the scanned document.
//   - [IICScannerDevice.SetTransferMode]
//   - [IICScannerDevice.MaxMemoryBandSize]: The total maximum band size requested when performing a memory-based transfer.
//   - [IICScannerDevice.SetMaxMemoryBandSize]
//
// # Logging into a Protected Device
//
//   - [IICScannerDevice.DefaultUsername]: A default username on protected scanners.
//   - [IICScannerDevice.SetDefaultUsername]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice
type IICScannerDevice interface {
	IICDevice

	// Topic: Selecting a Functional Unit

	// An array of functional unit types available on this scanner.
	AvailableFunctionalUnitTypes() []foundation.NSNumber
	// The currently selected functional unit on the scanner.
	SelectedFunctionalUnit() IICScannerFunctionalUnit
	// Requests to select a functional unit on the scanner.
	RequestSelectFunctionalUnit(type_ ICScannerFunctionalUnitType)

	// Topic: Performing a Scan

	// Opens a session on the protected device with the authorized username and passcode.
	RequestOpenSessionWithCredentialsPassword(username string, password string)
	// Starts an overview scan on the selected functional unit.
	RequestOverviewScan()
	// Starts a scan on the selected functional unit.
	RequestScan()
	// Cancels the current scan.
	CancelScan()
	// The document’s name.
	DocumentName() string
	SetDocumentName(value string)
	// The document’s uniform type identifier.
	DocumentUTI() string
	SetDocumentUTI(value string)
	// The downloads directory.
	DownloadsDirectory() foundation.NSURL
	SetDownloadsDirectory(value foundation.NSURL)
	// The transfer mode for the scanned document.
	TransferMode() ICScannerTransferMode
	SetTransferMode(value ICScannerTransferMode)
	// The total maximum band size requested when performing a memory-based transfer.
	MaxMemoryBandSize() uint32
	SetMaxMemoryBandSize(value uint32)

	// Topic: Logging into a Protected Device

	// A default username on protected scanners.
	DefaultUsername() string
	SetDefaultUsername(value string)
}

// Init initializes the instance.
func (s ICScannerDevice) Init() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerDevice) Autorelease() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerDevice creates a new ICScannerDevice instance.
func NewICScannerDevice() ICScannerDevice {
	class := getICScannerDeviceClass()
	rv := objc.Send[ICScannerDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests to select a functional unit on the scanner.
//
// # Discussion
//
// When the request has completed, [ScannerDeviceDidSelectFunctionalUnitError]
// is called on the delegate.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestSelect(_:)
func (s ICScannerDevice) RequestSelectFunctionalUnit(type_ ICScannerFunctionalUnitType) {
	objc.Send[objc.ID](s.ID, objc.Sel("requestSelectFunctionalUnit:"), type_)
}

// Opens a session on the protected device with the authorized username and
// passcode.
//
// # Discussion
//
// If the device reports a failure of credentials, you can provide them here
// for the launch. A client must open a session on a device in order to use
// the device.
//
// Before calling this method, set the receiver’s delegate; otherwise, the
// request is ignored.
//
// Once the request to open the session has completed,
// [DeviceDidOpenSessionWithError] is called on the delegate.
//
// No more messages are sent to the delegate if this request fails.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestOpenSession(withCredentials:password:)
func (s ICScannerDevice) RequestOpenSessionWithCredentialsPassword(username string, password string) {
	objc.Send[objc.ID](s.ID, objc.Sel("requestOpenSessionWithCredentials:password:"), objc.String(username), objc.String(password))
}

// Starts an overview scan on the selected functional unit.
//
// # Discussion
//
// Once the request to start an overview scan has completed,
// [ScannerDeviceDidCompleteOverviewScanWithError] is called on the delegate.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestOverviewScan()
func (s ICScannerDevice) RequestOverviewScan() {
	objc.Send[objc.ID](s.ID, objc.Sel("requestOverviewScan"))
}

// Starts a scan on the selected functional unit.
//
// # Discussion
//
// Once the request to start an overview scan has completed,
// [ScannerDeviceDidCompleteScanWithError] is called on the delegate.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestScan()
func (s ICScannerDevice) RequestScan() {
	objc.Send[objc.ID](s.ID, objc.Sel("requestScan"))
}

// Cancels the current scan.
//
// # Discussion
//
// Cancels the scan in progress by calling
// [ICScannerDevice.RequestOverviewScan] or [ICScannerDevice.RequestScan].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/cancelScan()
func (s ICScannerDevice) CancelScan() {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelScan"))
}

// An array of functional unit types available on this scanner.
//
// # Discussion
//
// This array contains [NSNumber] objects whose values are of type
// [ICScannerFunctionalUnitType].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/availableFunctionalUnitTypes
//
// [ICScannerFunctionalUnitType]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitType
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (s ICScannerDevice) AvailableFunctionalUnitTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("availableFunctionalUnitTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The currently selected functional unit on the scanner.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/selectedFunctionalUnit
func (s ICScannerDevice) SelectedFunctionalUnit() IICScannerFunctionalUnit {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectedFunctionalUnit"))
	return ICScannerFunctionalUnitFromID(objc.ID(rv))
}

// The document’s name.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/documentName
func (s ICScannerDevice) DocumentName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("documentName"))
	return foundation.NSStringFromID(rv).String()
}
func (s ICScannerDevice) SetDocumentName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentName:"), objc.String(value))
}

// The document’s uniform type identifier.
//
// # Discussion
//
// Supported uniform type identifiers are `kUTTypeJPEG`, `kUTTypeJPEG2000`,
// `kUTTypeTIFF`, and `kUTTypePNG`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/documentUTI
func (s ICScannerDevice) DocumentUTI() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("documentUTI"))
	return foundation.NSStringFromID(rv).String()
}
func (s ICScannerDevice) SetDocumentUTI(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentUTI:"), objc.String(value))
}

// The downloads directory.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/downloadsDirectory
func (s ICScannerDevice) DownloadsDirectory() foundation.NSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadsDirectory"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (s ICScannerDevice) SetDownloadsDirectory(value foundation.NSURL) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadsDirectory:"), value)
}

// The transfer mode for the scanned document.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/transferMode
func (s ICScannerDevice) TransferMode() ICScannerTransferMode {
	rv := objc.Send[ICScannerTransferMode](s.ID, objc.Sel("transferMode"))
	return ICScannerTransferMode(rv)
}
func (s ICScannerDevice) SetTransferMode(value ICScannerTransferMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setTransferMode:"), value)
}

// The total maximum band size requested when performing a memory-based
// transfer.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/maxMemoryBandSize
func (s ICScannerDevice) MaxMemoryBandSize() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("maxMemoryBandSize"))
	return rv
}
func (s ICScannerDevice) SetMaxMemoryBandSize(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxMemoryBandSize:"), value)
}

// A default username on protected scanners.
//
// # Discussion
//
// If the scanner is protected, you can set this property to a specific
// username as a convenience, instead of prompting the user for a username.
// The value persists until reset to `nil`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/defaultUsername
func (s ICScannerDevice) DefaultUsername() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("defaultUsername"))
	return foundation.NSStringFromID(rv).String()
}
func (s ICScannerDevice) SetDefaultUsername(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setDefaultUsername:"), objc.String(value))
}
