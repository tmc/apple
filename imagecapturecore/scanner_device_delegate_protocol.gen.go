// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for determining availability, selecting a functional unit, and performing scans on connected scanners.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate
type ICScannerDeviceDelegate interface {
	objectivec.IObject
	ICDeviceDelegate
}

// ICScannerDeviceDelegateObject wraps an existing Objective-C object that conforms to the ICScannerDeviceDelegate protocol.
type ICScannerDeviceDelegateObject struct {
	objectivec.Object
}

func (o ICScannerDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ICScannerDeviceDelegateObjectFromID constructs a [ICScannerDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ICScannerDeviceDelegateObjectFromID(id objc.ID) ICScannerDeviceDelegateObject {
	return ICScannerDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the client when another client closes the current open session on the
// scanner.
//
// # Discussion
//
// Scanners require exclusive access. Only one client can open a session on a
// scanner at a time. The scanner is available if it does not have a session
// opened by another client. Attempting to open a session on a scanner that
// already has an open session for another client will result in an error.
//
// To open a session on a scanner as soon as it is available, implement this
// method and call [ICDevice.RequestOpenSession] in the method body.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDeviceDidBecomeAvailable(_:)
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidBecomeAvailable(scanner IICScannerDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDeviceDidBecomeAvailable:"), scanner)
}

// Tells the client when a functional unit is selected on the scanner.
//
// # Discussion
//
// A functional unit is selected immediately after the scanner instantiates
// and in response to calling [ICScannerDevice.RequestSelectFunctionalUnit].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDevice(_:didSelect:error:)
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidSelectFunctionalUnitError(scanner IICScannerDevice, functionalUnit IICScannerFunctionalUnit, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDevice:didSelectFunctionalUnit:error:"), scanner, functionalUnit, error_)
}

// Tells the client when the scanner completes an overview scan.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDevice(_:didCompleteOverviewScanWithError:)
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidCompleteOverviewScanWithError(scanner IICScannerDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDevice:didCompleteOverviewScanWithError:"), scanner, error_)
}

// Tells the client when the scanner completes a scan.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDevice(_:didCompleteScanWithError:)
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidCompleteScanWithError(scanner IICScannerDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDevice:didCompleteScanWithError:"), scanner, error_)
}

// Tells the client when the scanner receives the requested scan progress
// notification and a band of data is sent for each notification received.
//
// # Discussion
//
// In memory transfer mode, this method sends a band of the size selected by
// the client using the [ICScannerDevice.MaxMemoryBandSize] property.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDevice(_:didScanTo:)-6tht3
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidScanToBandData(scanner IICScannerDevice, data IICScannerBandData) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDevice:didScanToBandData:"), scanner, data)
}

// Tells the client when the scanner receives the requested scan.
//
// # Discussion
//
// If the [ICScannerDevice.SelectedFunctionalUnit] is a document feeder, then
// this message is sent once for each scanned page.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDeviceDelegate/scannerDevice(_:didScanTo:)-10whl
func (o ICScannerDeviceDelegateObject) ScannerDeviceDidScanToURL(scanner IICScannerDevice, url foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("scannerDevice:didScanToURL:"), scanner, url)
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
func (o ICScannerDeviceDelegateObject) DeviceDidOpenSessionWithError(device IICDevice, error_ foundation.NSError) {
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
func (o ICScannerDeviceDelegateObject) DeviceDidCloseSessionWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didCloseSessionWithError:"), device, error_)
}

// Tells the delegate that a device has been removed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/didRemove(_:)
func (o ICScannerDeviceDelegateObject) DidRemoveDevice(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("didRemoveDevice:"), device)
}

// Tells the delegate when the device is ready to receive requests.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/deviceDidBecomeReady(_:)
func (o ICScannerDeviceDelegateObject) DeviceDidBecomeReady(device IICDevice) {
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
func (o ICScannerDeviceDelegateObject) DeviceDidReceiveStatusInformation(device IICDevice, status foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didReceiveStatusInformation:"), device, status)
}

// Tells the delegate when a device encounters an error.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEncounterError:)
func (o ICScannerDeviceDelegateObject) DeviceDidEncounterError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didEncounterError:"), device, error_)
}

// Tells the delegate when the ejection is complete.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEjectWithError:)
func (o ICScannerDeviceDelegateObject) DeviceDidEjectWithError(device IICDevice, error_ foundation.NSError) {
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
func (o ICScannerDeviceDelegateObject) DeviceDidChangeName(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidChangeName:"), device)
}

// ICScannerDeviceDelegateConfig holds optional typed callbacks for [ICScannerDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevicedelegate
type ICScannerDeviceDelegateConfig struct {

	// Determining Scanner Availability
	// ScannerDeviceDidBecomeAvailable — Tells the client when another client closes the current open session on the scanner.
	ScannerDeviceDidBecomeAvailable func(scanner ICScannerDevice)

	// Performing a Scan
	// ScannerDeviceDidCompleteOverviewScanWithError — Tells the client when the scanner completes an overview scan.
	ScannerDeviceDidCompleteOverviewScanWithError func(scanner ICScannerDevice, error_ foundation.NSError)
	// ScannerDeviceDidCompleteScanWithError — Tells the client when the scanner completes a scan.
	ScannerDeviceDidCompleteScanWithError func(scanner ICScannerDevice, error_ foundation.NSError)

	// Other Methods
	// ScannerDeviceDidSelectFunctionalUnitError — Tells the client when a functional unit is selected on the scanner.
	ScannerDeviceDidSelectFunctionalUnitError func(scanner ICScannerDevice, functionalUnit ICScannerFunctionalUnit, error_ foundation.NSError)
	// ScannerDeviceDidScanToBandData — Tells the client when the scanner receives the requested scan progress notification and a band of data is sent for each notification received.
	ScannerDeviceDidScanToBandData func(scanner ICScannerDevice, data ICScannerBandData)
	// ScannerDeviceDidScanToURL — Tells the client when the scanner receives the requested scan.
	ScannerDeviceDidScanToURL func(scanner ICScannerDevice, url foundation.NSURL)
}

// NewICScannerDeviceDelegate creates an Objective-C object implementing the [ICScannerDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ICScannerDeviceDelegateObject] satisfies the [ICScannerDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevicedelegate
func NewICScannerDeviceDelegate(config ICScannerDeviceDelegateConfig) ICScannerDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoICScannerDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ScannerDeviceDidBecomeAvailable != nil {
		fn := config.ScannerDeviceDidBecomeAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDeviceDidBecomeAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDeviceDidBecomeAvailable:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				fn(scanner)
				_delegateDone = true
			},
		})
	}

	if config.ScannerDeviceDidSelectFunctionalUnitError != nil {
		fn := config.ScannerDeviceDidSelectFunctionalUnitError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDevice:didSelectFunctionalUnit:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID, functionalUnitID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDevice:didSelectFunctionalUnit:error:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				functionalUnit := ICScannerFunctionalUnitFromID(functionalUnitID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(scanner, functionalUnit, error_)
				_delegateDone = true
			},
		})
	}

	if config.ScannerDeviceDidCompleteOverviewScanWithError != nil {
		fn := config.ScannerDeviceDidCompleteOverviewScanWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDevice:didCompleteOverviewScanWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDevice:didCompleteOverviewScanWithError:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(scanner, error_)
				_delegateDone = true
			},
		})
	}

	if config.ScannerDeviceDidCompleteScanWithError != nil {
		fn := config.ScannerDeviceDidCompleteScanWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDevice:didCompleteScanWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDevice:didCompleteScanWithError:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(scanner, error_)
				_delegateDone = true
			},
		})
	}

	if config.ScannerDeviceDidScanToBandData != nil {
		fn := config.ScannerDeviceDidScanToBandData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDevice:didScanToBandData:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID, dataID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDevice:didScanToBandData:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				data := ICScannerBandDataFromID(dataID)
				fn(scanner, data)
				_delegateDone = true
			},
		})
	}

	if config.ScannerDeviceDidScanToURL != nil {
		fn := config.ScannerDeviceDidScanToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("scannerDevice:didScanToURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, scannerID objc.ID, urlID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICScannerDeviceDelegate", "scannerDevice:didScanToURL:")
					}
				}()
				scanner := ICScannerDeviceFromID(scannerID)
				url := foundation.NSURLFromID(urlID)
				fn(scanner, url)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ICScannerDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewICScannerDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ICScannerDeviceDelegateObjectFromID(instance)
}
