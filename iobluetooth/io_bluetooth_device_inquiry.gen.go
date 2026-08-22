// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothDeviceInquiry] class.
var (
	_IOBluetoothDeviceInquiryClass     IOBluetoothDeviceInquiryClass
	_IOBluetoothDeviceInquiryClassOnce sync.Once
)

func getIOBluetoothDeviceInquiryClass() IOBluetoothDeviceInquiryClass {
	_IOBluetoothDeviceInquiryClassOnce.Do(func() {
		_IOBluetoothDeviceInquiryClass = IOBluetoothDeviceInquiryClass{class: objc.GetClass("IOBluetoothDeviceInquiry")}
	})
	return _IOBluetoothDeviceInquiryClass
}

// GetIOBluetoothDeviceInquiryClass returns the class object for IOBluetoothDeviceInquiry.
func GetIOBluetoothDeviceInquiryClass() IOBluetoothDeviceInquiryClass {
	return getIOBluetoothDeviceInquiryClass()
}

type IOBluetoothDeviceInquiryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothDeviceInquiryClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothDeviceInquiryClass) Alloc() IOBluetoothDeviceInquiry {
	rv := objc.Send[IOBluetoothDeviceInquiry](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// Object representing a device inquiry that finds Bluetooth devices in-range
// of the computer, and (optionally) retrieves name information for them.
//
// # Overview
//
// You should only use this object if your application needs to know about
// in-range devices and cannot use the GUI provided by the IOBluetoothUI
// framework. It will not let you perform unlimited back-to-back inquiries,
// but will instead throttle the number of attempted inquiries if too many are
// attempted within a small window of time. Important Note: DO NOT perform
// remote name requests on devices from delegate methods or while this object
// is in use. If you wish to do your own remote name requests on devices, do
// them after you have stopped this object. If you do not heed this warning,
// you could potentially deadlock your process.
//
// # Initializers
//
//   - [IOBluetoothDeviceInquiry.InitWithDelegate]: Initializes an alloc’d inquiry object, and sets the delegate object, as if -setDelegate: were called on it.
//
// # Instance Properties
//
//   - [IOBluetoothDeviceInquiry.Delegate]
//   - [IOBluetoothDeviceInquiry.SetDelegate]
//   - [IOBluetoothDeviceInquiry.InquiryLength]: Set the length of the inquiry that is performed each time -start is used on an inquiry object.
//   - [IOBluetoothDeviceInquiry.SetInquiryLength]
//   - [IOBluetoothDeviceInquiry.SearchType]: Set the devices that are found.
//   - [IOBluetoothDeviceInquiry.SetSearchType]
//   - [IOBluetoothDeviceInquiry.UpdateNewDeviceNames]: Sets whether or not the inquiry object will retrieve the names of devices found during the search.
//   - [IOBluetoothDeviceInquiry.SetUpdateNewDeviceNames]
//
// # Instance Methods
//
//   - [IOBluetoothDeviceInquiry.ClearFoundDevices]: Removes all found devices from the inquiry object.
//   - [IOBluetoothDeviceInquiry.FoundDevices]: Returns found IOBluetoothDevice objects as an array.
//   - [IOBluetoothDeviceInquiry.SetSearchCriteriaMajorDeviceClassMinorDeviceClass]: Use this method to set the criteria for the device search.
//   - [IOBluetoothDeviceInquiry.Start]: Tells inquiry object to begin the inquiry and name updating process, if specified.
//   - [IOBluetoothDeviceInquiry.Stop]: Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry
type IOBluetoothDeviceInquiry struct {
	objectivec.Object
}

// IOBluetoothDeviceInquiryFromID constructs a [IOBluetoothDeviceInquiry] from an objc.ID.
//
// Object representing a device inquiry that finds Bluetooth devices in-range
// of the computer, and (optionally) retrieves name information for them.
func IOBluetoothDeviceInquiryFromID(id objc.ID) IOBluetoothDeviceInquiry {
	return IOBluetoothDeviceInquiry{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothDeviceInquiry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothDeviceInquiry] class.
//
// # Initializers
//
//   - [IIOBluetoothDeviceInquiry.InitWithDelegate]: Initializes an alloc’d inquiry object, and sets the delegate object, as if -setDelegate: were called on it.
//
// # Instance Properties
//
//   - [IIOBluetoothDeviceInquiry.Delegate]
//   - [IIOBluetoothDeviceInquiry.SetDelegate]
//   - [IIOBluetoothDeviceInquiry.InquiryLength]: Set the length of the inquiry that is performed each time -start is used on an inquiry object.
//   - [IIOBluetoothDeviceInquiry.SetInquiryLength]
//   - [IIOBluetoothDeviceInquiry.SearchType]: Set the devices that are found.
//   - [IIOBluetoothDeviceInquiry.SetSearchType]
//   - [IIOBluetoothDeviceInquiry.UpdateNewDeviceNames]: Sets whether or not the inquiry object will retrieve the names of devices found during the search.
//   - [IIOBluetoothDeviceInquiry.SetUpdateNewDeviceNames]
//
// # Instance Methods
//
//   - [IIOBluetoothDeviceInquiry.ClearFoundDevices]: Removes all found devices from the inquiry object.
//   - [IIOBluetoothDeviceInquiry.FoundDevices]: Returns found IOBluetoothDevice objects as an array.
//   - [IIOBluetoothDeviceInquiry.SetSearchCriteriaMajorDeviceClassMinorDeviceClass]: Use this method to set the criteria for the device search.
//   - [IIOBluetoothDeviceInquiry.Start]: Tells inquiry object to begin the inquiry and name updating process, if specified.
//   - [IIOBluetoothDeviceInquiry.Stop]: Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry
type IIOBluetoothDeviceInquiry interface {
	objectivec.IObject

	// Topic: Initializers

	// Initializes an alloc’d inquiry object, and sets the delegate object, as if -setDelegate: were called on it.
	InitWithDelegate(delegate objectivec.IObject) IOBluetoothDeviceInquiry

	// Topic: Instance Properties

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	// Set the length of the inquiry that is performed each time -start is used on an inquiry object.
	InquiryLength() uint8
	SetInquiryLength(value uint8)
	// Set the devices that are found.
	SearchType() IOBluetoothDeviceSearchTypes
	SetSearchType(value IOBluetoothDeviceSearchTypes)
	// Sets whether or not the inquiry object will retrieve the names of devices found during the search.
	UpdateNewDeviceNames() bool
	SetUpdateNewDeviceNames(value bool)

	// Topic: Instance Methods

	// Removes all found devices from the inquiry object.
	ClearFoundDevices()
	// Returns found IOBluetoothDevice objects as an array.
	FoundDevices() foundation.INSArray
	// Use this method to set the criteria for the device search.
	SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor BluetoothServiceClassMajor, inMajorDeviceClass BluetoothDeviceClassMajor, inMinorDeviceClass BluetoothDeviceClassMinor)
	// Tells inquiry object to begin the inquiry and name updating process, if specified.
	Start() kernel.IOReturn
	// Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
	Stop() kernel.IOReturn
}

// Init initializes the instance.
func (b IOBluetoothDeviceInquiry) Init() IOBluetoothDeviceInquiry {
	rv := objc.Send[IOBluetoothDeviceInquiry](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothDeviceInquiry) Autorelease() IOBluetoothDeviceInquiry {
	rv := objc.Send[IOBluetoothDeviceInquiry](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothDeviceInquiry creates a new IOBluetoothDeviceInquiry instance.
func NewIOBluetoothDeviceInquiry() IOBluetoothDeviceInquiry {
	class := getIOBluetoothDeviceInquiryClass()
	rv := objc.Send[IOBluetoothDeviceInquiry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an alloc’d inquiry object, and sets the delegate object, as
// if -setDelegate: were called on it.
//
// delegate: A delegate object that wishes to receive messages from the inquiry object.
// Delegate methods are listed below, under IOBluetoothDeviceInquiryDelegate.
//
// # Return Value
//
// A pointer to the initialized IOBluetoothDeviceInquiry object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/init(delegate:)
func NewBluetoothDeviceInquiryWithDelegate(delegate objectivec.IObject) IOBluetoothDeviceInquiry {
	instance := getIOBluetoothDeviceInquiryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	return IOBluetoothDeviceInquiryFromID(rv)
}

// Initializes an alloc’d inquiry object, and sets the delegate object, as
// if -setDelegate: were called on it.
//
// delegate: A delegate object that wishes to receive messages from the inquiry object.
// Delegate methods are listed below, under IOBluetoothDeviceInquiryDelegate.
//
// # Return Value
//
// A pointer to the initialized IOBluetoothDeviceInquiry object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/init(delegate:)
func (b IOBluetoothDeviceInquiry) InitWithDelegate(delegate objectivec.IObject) IOBluetoothDeviceInquiry {
	rv := objc.Send[IOBluetoothDeviceInquiry](b.ID, objc.Sel("initWithDelegate:"), delegate)
	return rv
}

// Removes all found devices from the inquiry object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/clearFoundDevices()
func (b IOBluetoothDeviceInquiry) ClearFoundDevices() {
	objc.Send[objc.ID](b.ID, objc.Sel("clearFoundDevices"))
}

// Returns found IOBluetoothDevice objects as an array.
//
// # Return Value
//
// Returns an NSArray of IOBluetoothDevice objects.
//
// # Discussion
//
// Will not return nil. If there are no devices found, returns an array with
// length of 0.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/foundDevices()
func (b IOBluetoothDeviceInquiry) FoundDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("foundDevices"))
	return foundation.NSArrayFromID(rv)
}

// Use this method to set the criteria for the device search.
//
// inServiceClassMajor: Set the major service class for found devices. Set to
// kBluetoothServiceClassMajorAny for all devices. See
// BluetoothAssignedNumbers.h for possible values.
//
// inMajorDeviceClass: Set the major device class for found devices. Set to
// kBluetoothDeviceClassMajorAny for all devices. See
// BluetoothAssignedNumbers.h for possible values.
//
// inMinorDeviceClass: Set the minor device class for found devices. Set to
// kBluetoothDeviceClassMinorAny for all devices. See
// BluetoothAssignedNumbers.h for possible values.
//
// # Discussion
//
// The default inquiry object will search for all types of devices. If you
// wish to find only keyboards, for example, you might use this method like
// this:
//
// [myInquiryObject setSearchCriteria:kBluetoothServiceClassMajorAny
// majorDeviceClass:kBluetoothDeviceClassMajorPeripheral
// minorDeviceClass:kBluetoothDeviceClassMinorPeripheral1Keyboard];
//
// However, we recommend only using this if you are certain of the device
// class you are looking for, as some devices may report a
// different/unexpected device class, and the search may miss the device you
// are interested in.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/setSearchCriteria(_:majorDeviceClass:minorDeviceClass:)
func (b IOBluetoothDeviceInquiry) SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor BluetoothServiceClassMajor, inMajorDeviceClass BluetoothDeviceClassMajor, inMinorDeviceClass BluetoothDeviceClassMinor) {
	objc.Send[objc.ID](b.ID, objc.Sel("setSearchCriteria:majorDeviceClass:minorDeviceClass:"), inServiceClassMajor, inMajorDeviceClass, inMinorDeviceClass)
}

// Tells inquiry object to begin the inquiry and name updating process, if
// specified.
//
// # Return Value
//
// Returns kIOReturnSuccess if start was successful. Returns kIOReturnBusy if
// the object is already in process. May return other IOReturn values, as
// appropriate.
//
// # Discussion
//
// Calling start multiple times in rapid succession or back-to-back will
// probably not produce the intended results. Inquiries are throttled if they
// are called too quickly in succession.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/start()
func (b IOBluetoothDeviceInquiry) Start() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("start"))
	return kernel.IOReturn(rv)
}

// Halts the inquiry object. Could either stop the search for new devices, or
// the updating of found device names.
//
// # Return Value
//
// Returns kIOReturnSuccess if the inquiry is successfully stopped. Returns
// kIOReturnNotPermitted if the inquiry object is already stopped. May return
// other IOReturn values, as appropriate.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/stop()
func (b IOBluetoothDeviceInquiry) Stop() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("stop"))
	return kernel.IOReturn(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/delegate
func (b IOBluetoothDeviceInquiry) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (b IOBluetoothDeviceInquiry) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegate:"), value)
}

// Set the length of the inquiry that is performed each time -start is used on
// an inquiry object.
//
// # Discussion
//
// A default of 10 seconds is used, unless a different value is specified
// using this method. Note that if you have called -start again too quickly,
// your inquiry may actually take much longer than what length you specify, as
// inquiries are throttled in the system. Also note that if you have the
// inquiry object updating device names for you, the whole inquiry process
// could be much longer than the specified length, depending on the number of
// devices found and how responsive to name requests they are. If you -must-
// have a strict inquiry length, disable name updates. In other words, this
// “length” only refers to the actual device discovery portion of the
// whole inquiry process.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryLength
func (b IOBluetoothDeviceInquiry) InquiryLength() uint8 {
	rv := objc.Send[uint8](b.ID, objc.Sel("inquiryLength"))
	return rv
}
func (b IOBluetoothDeviceInquiry) SetInquiryLength(value uint8) {
	objc.Send[struct{}](b.ID, objc.Sel("setInquiryLength:"), value)
}

// Set the devices that are found.
//
// # Discussion
//
// A default of kIOBluetoothDeviceSearchClassic is used, unless a different
// value is specified using this method.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/searchType
func (b IOBluetoothDeviceInquiry) SearchType() IOBluetoothDeviceSearchTypes {
	rv := objc.Send[IOBluetoothDeviceSearchTypes](b.ID, objc.Sel("searchType"))
	return IOBluetoothDeviceSearchTypes(rv)
}
func (b IOBluetoothDeviceInquiry) SetSearchType(value IOBluetoothDeviceSearchTypes) {
	objc.Send[struct{}](b.ID, objc.Sel("setSearchType:"), value)
}

// Sets whether or not the inquiry object will retrieve the names of devices
// found during the search.
//
// # Discussion
//
// The default value for the inquiry object is TRUE, unless this method is
// used to change it.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/updateNewDeviceNames
func (b IOBluetoothDeviceInquiry) UpdateNewDeviceNames() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("updateNewDeviceNames"))
	return rv
}
func (b IOBluetoothDeviceInquiry) SetUpdateNewDeviceNames(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setUpdateNewDeviceNames:"), value)
}
