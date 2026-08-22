// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothSDPServiceRecord] class.
var (
	_IOBluetoothSDPServiceRecordClass     IOBluetoothSDPServiceRecordClass
	_IOBluetoothSDPServiceRecordClassOnce sync.Once
)

func getIOBluetoothSDPServiceRecordClass() IOBluetoothSDPServiceRecordClass {
	_IOBluetoothSDPServiceRecordClassOnce.Do(func() {
		_IOBluetoothSDPServiceRecordClass = IOBluetoothSDPServiceRecordClass{class: objc.GetClass("IOBluetoothSDPServiceRecord")}
	})
	return _IOBluetoothSDPServiceRecordClass
}

// GetIOBluetoothSDPServiceRecordClass returns the class object for IOBluetoothSDPServiceRecord.
func GetIOBluetoothSDPServiceRecordClass() IOBluetoothSDPServiceRecordClass {
	return getIOBluetoothSDPServiceRecordClass()
}

type IOBluetoothSDPServiceRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothSDPServiceRecordClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothSDPServiceRecordClass) Alloc() IOBluetoothSDPServiceRecord {
	rv := objc.Send[IOBluetoothSDPServiceRecord](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An instance of this class represents a single SDP service record.
//
// # Overview
//
// As a service record, an instance of this class has an NSDictionary of
// service attributes. It also has a link to the IOBluetoothDevice that the
// service belongs to. The service dictionary is keyed off of the attribute ID
// of each attribute represented as an NSNumber.
//
// # Initializers
//
//   - [IOBluetoothSDPServiceRecord.InitWithServiceDictionaryDevice]: Returns an initialized IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// # Instance Properties
//
//   - [IOBluetoothSDPServiceRecord.Attributes]: Returns an NSDictionary containing the attributes for the service.
//   - [IOBluetoothSDPServiceRecord.Device]: Returns the IOBluetoothDevice that the target service belongs to.
//   - [IOBluetoothSDPServiceRecord.SortedAttributes]: Returns a sorted array of SDP attributes
//
// # Instance Methods
//
//   - [IOBluetoothSDPServiceRecord.GetAttributeDataElement]: Returns the data element for the given attribute ID in the target service.
//   - [IOBluetoothSDPServiceRecord.GetServiceRecordHandle]: Allows the discovery of the service record handle assigned to the service.
//   - [IOBluetoothSDPServiceRecord.GetL2CAPPSM]: Allows the discovery of the L2CAP PSM assigned to the service.
//   - [IOBluetoothSDPServiceRecord.GetRFCOMMChannelID]: Allows the discovery of the RFCOMM channel ID assigned to the service.
//   - [IOBluetoothSDPServiceRecord.GetSDPServiceRecordRef]: Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
//   - [IOBluetoothSDPServiceRecord.GetServiceName]: Returns the name of the service.
//   - [IOBluetoothSDPServiceRecord.HandsFreeSupportedFeatures]
//   - [IOBluetoothSDPServiceRecord.HasServiceFromArray]: Returns TRUE if any one of the UUIDs in the given array is found in the target service.
//   - [IOBluetoothSDPServiceRecord.MatchesSearchArray]: Returns TRUE any of the UUID arrays in the search array match the target service.
//   - [IOBluetoothSDPServiceRecord.MatchesUUID16]: Returns TRUE the UUID16 is found in the target service.
//   - [IOBluetoothSDPServiceRecord.MatchesUUIDArray]: Returns TRUE if ALL of the UUIDs in the given array is found in the target service.
//   - [IOBluetoothSDPServiceRecord.RemoveServiceRecord]: Removes the service from the local SDP server.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord
type IOBluetoothSDPServiceRecord struct {
	objectivec.Object
}

// IOBluetoothSDPServiceRecordFromID constructs a [IOBluetoothSDPServiceRecord] from an objc.ID.
//
// An instance of this class represents a single SDP service record.
func IOBluetoothSDPServiceRecordFromID(id objc.ID) IOBluetoothSDPServiceRecord {
	return IOBluetoothSDPServiceRecord{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothSDPServiceRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothSDPServiceRecord] class.
//
// # Initializers
//
//   - [IIOBluetoothSDPServiceRecord.InitWithServiceDictionaryDevice]: Returns an initialized IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// # Instance Properties
//
//   - [IIOBluetoothSDPServiceRecord.Attributes]: Returns an NSDictionary containing the attributes for the service.
//   - [IIOBluetoothSDPServiceRecord.Device]: Returns the IOBluetoothDevice that the target service belongs to.
//   - [IIOBluetoothSDPServiceRecord.SortedAttributes]: Returns a sorted array of SDP attributes
//
// # Instance Methods
//
//   - [IIOBluetoothSDPServiceRecord.GetAttributeDataElement]: Returns the data element for the given attribute ID in the target service.
//   - [IIOBluetoothSDPServiceRecord.GetServiceRecordHandle]: Allows the discovery of the service record handle assigned to the service.
//   - [IIOBluetoothSDPServiceRecord.GetL2CAPPSM]: Allows the discovery of the L2CAP PSM assigned to the service.
//   - [IIOBluetoothSDPServiceRecord.GetRFCOMMChannelID]: Allows the discovery of the RFCOMM channel ID assigned to the service.
//   - [IIOBluetoothSDPServiceRecord.GetSDPServiceRecordRef]: Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
//   - [IIOBluetoothSDPServiceRecord.GetServiceName]: Returns the name of the service.
//   - [IIOBluetoothSDPServiceRecord.HandsFreeSupportedFeatures]
//   - [IIOBluetoothSDPServiceRecord.HasServiceFromArray]: Returns TRUE if any one of the UUIDs in the given array is found in the target service.
//   - [IIOBluetoothSDPServiceRecord.MatchesSearchArray]: Returns TRUE any of the UUID arrays in the search array match the target service.
//   - [IIOBluetoothSDPServiceRecord.MatchesUUID16]: Returns TRUE the UUID16 is found in the target service.
//   - [IIOBluetoothSDPServiceRecord.MatchesUUIDArray]: Returns TRUE if ALL of the UUIDs in the given array is found in the target service.
//   - [IIOBluetoothSDPServiceRecord.RemoveServiceRecord]: Removes the service from the local SDP server.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord
type IIOBluetoothSDPServiceRecord interface {
	objectivec.IObject

	// Topic: Initializers

	// Returns an initialized IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
	InitWithServiceDictionaryDevice(serviceDict foundation.INSDictionary, device IIOBluetoothDevice) IOBluetoothSDPServiceRecord

	// Topic: Instance Properties

	// Returns an NSDictionary containing the attributes for the service.
	Attributes() foundation.INSDictionary
	// Returns the IOBluetoothDevice that the target service belongs to.
	Device() IIOBluetoothDevice
	// Returns a sorted array of SDP attributes
	SortedAttributes() foundation.INSArray

	// Topic: Instance Methods

	// Returns the data element for the given attribute ID in the target service.
	GetAttributeDataElement(attributeID BluetoothSDPServiceAttributeID) IIOBluetoothSDPDataElement
	// Allows the discovery of the service record handle assigned to the service.
	GetServiceRecordHandle(outServiceRecordHandle *BluetoothSDPServiceRecordHandle) kernel.IOReturn
	// Allows the discovery of the L2CAP PSM assigned to the service.
	GetL2CAPPSM(outPSM *BluetoothL2CAPPSM) kernel.IOReturn
	// Allows the discovery of the RFCOMM channel ID assigned to the service.
	GetRFCOMMChannelID(rfcommChannelID *BluetoothRFCOMMChannelID) kernel.IOReturn
	// Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
	GetSDPServiceRecordRef() IOBluetoothSDPServiceRecordRef
	// Returns the name of the service.
	GetServiceName() string
	HandsFreeSupportedFeatures() uint16
	// Returns TRUE if any one of the UUIDs in the given array is found in the target service.
	HasServiceFromArray(array foundation.INSArray) bool
	// Returns TRUE any of the UUID arrays in the search array match the target service.
	MatchesSearchArray(searchArray foundation.INSArray) bool
	// Returns TRUE the UUID16 is found in the target service.
	MatchesUUID16(uuid16 BluetoothSDPUUID16) bool
	// Returns TRUE if ALL of the UUIDs in the given array is found in the target service.
	MatchesUUIDArray(uuidArray foundation.INSArray) bool
	// Removes the service from the local SDP server.
	RemoveServiceRecord() kernel.IOReturn

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b IOBluetoothSDPServiceRecord) Init() IOBluetoothSDPServiceRecord {
	rv := objc.Send[IOBluetoothSDPServiceRecord](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothSDPServiceRecord) Autorelease() IOBluetoothSDPServiceRecord {
	rv := objc.Send[IOBluetoothSDPServiceRecord](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothSDPServiceRecord creates a new IOBluetoothSDPServiceRecord instance.
func NewIOBluetoothSDPServiceRecord() IOBluetoothSDPServiceRecord {
	class := getIOBluetoothSDPServiceRecordClass()
	rv := objc.Send[IOBluetoothSDPServiceRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an initialized IOBluetoothSDPServiceRecord * with the attributes
// specified in the provided service dictionary. Provide a pointer to an
// IOBlueotothDevice if you wish to associate the record to a specific
// IOBluetoothDevice.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/init(serviceDictionary:device:)
func NewBluetoothSDPServiceRecordWithServiceDictionaryDevice(serviceDict foundation.INSDictionary, device IIOBluetoothDevice) IOBluetoothSDPServiceRecord {
	instance := getIOBluetoothSDPServiceRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServiceDictionary:device:"), serviceDict, device)
	return IOBluetoothSDPServiceRecordFromID(rv)
}

// Returns an initialized IOBluetoothSDPServiceRecord * with the attributes
// specified in the provided service dictionary. Provide a pointer to an
// IOBlueotothDevice if you wish to associate the record to a specific
// IOBluetoothDevice.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/init(serviceDictionary:device:)
func (b IOBluetoothSDPServiceRecord) InitWithServiceDictionaryDevice(serviceDict foundation.INSDictionary, device IIOBluetoothDevice) IOBluetoothSDPServiceRecord {
	rv := objc.Send[IOBluetoothSDPServiceRecord](b.ID, objc.Sel("initWithServiceDictionary:device:"), serviceDict, device)
	return rv
}

// Returns the data element for the given attribute ID in the target service.
//
// attributeID: The attribute ID of the desired attribute.
//
// # Return Value
//
// Returns the data element for the given attribute ID in the target service.
// If the service does not contain an attribute with the given ID, then nil is
// returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getAttributeDataElement(_:)
func (b IOBluetoothSDPServiceRecord) GetAttributeDataElement(attributeID BluetoothSDPServiceAttributeID) IIOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getAttributeDataElement:"), attributeID)
	return IOBluetoothSDPDataElementFromID(rv)
}

// Allows the discovery of the service record handle assigned to the service.
//
// outServiceRecordHandle: A pointer to the location that will get the found service record handle.
//
// # Return Value
//
// Returns kIOReturnSuccess if the service record handle is found.
//
// # Discussion
//
// This method will search through the attributes to find the one representing
// the service record handle. If one is found the outServiceRecordHandle param
// is set with the value. The outServiceRecordHandle value only gets set when
// kIOReturnSuccess is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getHandle(_:)
func (b IOBluetoothSDPServiceRecord) GetServiceRecordHandle(outServiceRecordHandle *BluetoothSDPServiceRecordHandle) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("getServiceRecordHandle:"), unsafe.Pointer(outServiceRecordHandle))
	return kernel.IOReturn(rv)
}

// Allows the discovery of the L2CAP PSM assigned to the service.
//
// outPSM: A pointer to the location that will get the found L2CAP PSM.
//
// # Return Value
//
// Returns kIOReturnSuccess if the PSM is found.
//
// # Discussion
//
// This method will search through the ProtoclDescriptorList attribute to find
// an entry with the L2CAP UUID (UUID16: 0x0100). If one is found, it gets the
// second element of the data element sequence and sets the outPSM pointer to
// it. The PSM value only gets set when kIOReturnSuccess is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getL2CAPPSM(_:)
func (b IOBluetoothSDPServiceRecord) GetL2CAPPSM(outPSM *BluetoothL2CAPPSM) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("getL2CAPPSM:"), unsafe.Pointer(outPSM))
	return kernel.IOReturn(rv)
}

// Allows the discovery of the RFCOMM channel ID assigned to the service.
//
// rfcommChannelID: A pointer to the location that will get the found RFCOMM channel ID.
//
// # Return Value
//
// Returns kIOReturnSuccess if the channel ID is found.
//
// # Discussion
//
// This method will search through the ProtoclDescriptorList attribute to find
// an entry with the RFCOMM UUID (UUID16: 0x0003). If one is found, it gets
// the second element of the data element sequence and sets the
// rfcommChannelID pointer to it. The channel ID only gets set when
// kIOReturnSuccess is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRFCOMMChannelID(_:)
func (b IOBluetoothSDPServiceRecord) GetRFCOMMChannelID(rfcommChannelID *BluetoothRFCOMMChannelID) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("getRFCOMMChannelID:"), unsafe.Pointer(rfcommChannelID))
	return kernel.IOReturn(rv)
}

// Returns an IOBluetoothSDPServiceRecordRef representation of the target
// IOBluetoothSDPServiceRecord object.
//
// # Return Value
//
// Returns an IOBluetoothSDPServiceRecordRef representation of the target
// IOBluetoothSDPServiceRecord object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRef()
func (b IOBluetoothSDPServiceRecord) GetSDPServiceRecordRef() IOBluetoothSDPServiceRecordRef {
	rv := objc.Send[IOBluetoothSDPServiceRecordRef](b.ID, objc.Sel("getSDPServiceRecordRef"))
	return IOBluetoothSDPServiceRecordRef(rv)
}

// Returns the name of the service.
//
// # Return Value
//
// Returns the name of the target service.
//
// # Discussion
//
// This is currently implemented to simply return the attribute with an id of
// 0x0100. In the future, it will be extended to allow name localization based
// on the user’s chosen language or other languages.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getServiceName()
func (b IOBluetoothSDPServiceRecord) GetServiceName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getServiceName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/handsFreeSupportedFeatures()
func (b IOBluetoothSDPServiceRecord) HandsFreeSupportedFeatures() uint16 {
	rv := objc.Send[uint16](b.ID, objc.Sel("handsFreeSupportedFeatures"))
	return rv
}

// Returns TRUE if any one of the UUIDs in the given array is found in the
// target service.
//
// array: An NSArray of IOBluetoothSDPUUID objects to search for in the target
// service.
//
// # Return Value
//
// Returns TRUE if any of the given UUIDs are present in the service.
//
// # Discussion
//
// The given array should contain IOBluetoothSDPUUID objects. It is currently
// implemented such that it returns TRUE if any of the UUIDs are found.
// However in the future, it is likely that this will change to more closely
// match the functionality in the SDP spec so that it only returns TRUE if all
// of the given UUIDs are present. That way, both AND and OR comparisons can
// be implemented. Please make a note of this potential change.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/hasService(from:)
func (b IOBluetoothSDPServiceRecord) HasServiceFromArray(array foundation.INSArray) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("hasServiceFromArray:"), array)
	return rv
}

// Returns TRUE any of the UUID arrays in the search array match the target
// service.
//
// searchArray: An NSArray of NSArrays of IOBluetoothSDPUUID objects.
//
// # Return Value
//
// Returns [TRUE] if any of the UUID arrays match.
//
// # Discussion
//
// The given array should contain [NSArray] objects. Each sub-[NSArray] should
// contain [IOBluetoothSDPUUID] objects. In turn, each sub-NSArray gets passed
// to -matchesUUIDArray: If any of those returns [TRUE], then the search stops
// and [TRUE] is returned. Essentially the primary NSArray contains the OR
// operations and each sub-array contains the AND operations.
//
// NOTE: This method is only available in macOS 10.2.4 (Bluetooth v1.1) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesSearch(_:)
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
func (b IOBluetoothSDPServiceRecord) MatchesSearchArray(searchArray foundation.INSArray) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("matchesSearchArray:"), searchArray)
	return rv
}

// Returns TRUE the UUID16 is found in the target service.
//
// uuid16: A BluetoothSDPUUID16 to search for in the target service.
//
// # Return Value
//
// Returns TRUE if the UUID16 is present in the service.
//
// # Discussion
//
// NOTE: This method is only available in macOS 10.7 or later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUID16(_:)
func (b IOBluetoothSDPServiceRecord) MatchesUUID16(uuid16 BluetoothSDPUUID16) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("matchesUUID16:"), uuid16)
	return rv
}

// Returns TRUE if ALL of the UUIDs in the given array is found in the target
// service.
//
// uuidArray: An NSArray of IOBluetoothSDPUUID objects to search for in the target
// service.
//
// # Return Value
//
// Returns TRUE if all of the given UUIDs are present in the service.
//
// # Discussion
//
// The given array should contain IOBluetoothSDPUUID objects. It only returns
// TRUE if all of the UUIDs are found. This method is like
// hasServiceFromArray: except that it requires that all UUIDs match instead
// of any of them matching.
//
// NOTE: This method is only available in macOS 10.2.4 (Bluetooth v1.1) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUIDArray(_:)
func (b IOBluetoothSDPServiceRecord) MatchesUUIDArray(uuidArray foundation.INSArray) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("matchesUUIDArray:"), uuidArray)
	return rv
}

// Removes the service from the local SDP server.
//
// # Return Value
//
// Returns kIOReturnSuccess if successful.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/remove()
func (b IOBluetoothSDPServiceRecord) RemoveServiceRecord() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("removeServiceRecord"))
	return kernel.IOReturn(rv)
}
func (b IOBluetoothSDPServiceRecord) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Adds a service to the local SDP server.
//
// serviceDict: A dictionary containing the attributes for the new service
//
// # Return Value
//
// Returns an IOBluetoothSDPServiceRecord * with the attributes specified in
// the provided dictionary.
//
// # Discussion
//
// Each entry in the dictionary representing the service contains the
// individual attributes. Each attribute in the dict is keyed by a string that
// must begin with a hex number representing the attribute ID. The key string
// may contain additional characters if desired as long as they follow a space
// after the ID hex string. The attribute value must follow the dictionary
// format described by IOBluetoothSDPDataElement. This dictionary format
// allows a service dict to be created as a plist file and then loaded into
// the system rather than built up in code. See the example code for an
// example of how can be done.
//
// If the service record handle, L2CAP PSM or RFCOMM channel ID specified in
// the dictionary are in use, an alternate one will be assigned.
//
// In addition to attributes that represent the service itself, additional
// attributes may be specified that control the local behavior of the service.
// To specify these local attributes, an additional property titled
// “LocalAttributes” may be added to the root of the service dict. The
// value of this property must be a dictionary that contains the individual
// local attributes.
//
// Currently, only two local attributes are supported: “Persistent” and
// “TargetApplication”.
//
// The “Persistent” local attribute must be either a boolean or number
// representing whether the service should be persistent. A persistent service
// will be saved off and restored any time the Bluetooth hardware is present.
// It will persist through reboots and can only be removed by calling
// IOBluetoothRemoveServiceWithRecordHandle(). This attribute is optional. By
// default, if no “Persistent” local property is present, the service will
// only exist temporarily. It will be removed either when
// IOBluetoothRemoveServiceWithRecordHandle() is called or when the client
// application exits.
//
// The “TargetApplication” local attribute is used to specify an
// application to be launched when a remote device attempts to connect to the
// service (by opening either an L2CAP or RFCOMM channel of the type specified
// in the service). This value must be a string representing the absolute path
// to the target executable (not just the .app wrapper - i.e.
// /System/Library/CoreServices/OBEXAgent.app/Contents/MacOS/OBEXAgent). This
// attribute is optional. If no “TargetApplication” local attribute is
// specified, no special action will take place when an incoming connection to
// the service is created. It is up to the client to be monitoring for the
// connection and to do the right thing when one appears.
//
// The “LocalAttributes” property is optional. If it is not specified, by
// default the created service is transient and will be removed when the
// client exits.
//
// Additional local attributes to further control incoming services will be
// added in the future.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/publishedServiceRecord(with:)
func (_IOBluetoothSDPServiceRecordClass IOBluetoothSDPServiceRecordClass) PublishedServiceRecordWithDictionary(serviceDict foundation.INSDictionary) IOBluetoothSDPServiceRecord {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPServiceRecordClass.class), objc.Sel("publishedServiceRecordWithDictionary:"), serviceDict)
	return IOBluetoothSDPServiceRecordFromID(rv)
}

// Method call to convert an IOBluetoothSDPServiceRecordRef into an
// IOBluetoothSDPServiceRecord *.
//
// sdpServiceRecordRef: IOBluetoothSDPServiceRecordRef for which an IOBluetoothSDPServiceRecord *
// is desired.
//
// # Return Value
//
// Returns the IOBluetoothSDPServiceRecord * for the given
// IOBluetoothSDPServiceRecordRef.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withSDPServiceRecordRef(_:)
func (_IOBluetoothSDPServiceRecordClass IOBluetoothSDPServiceRecordClass) WithSDPServiceRecordRef(sdpServiceRecordRef IOBluetoothSDPServiceRecordRef) IOBluetoothSDPServiceRecord {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPServiceRecordClass.class), objc.Sel("withSDPServiceRecordRef:"), sdpServiceRecordRef)
	return IOBluetoothSDPServiceRecordFromID(rv)
}

// Returns an IOBluetoothSDPServiceRecord * with the attributes specified in
// the provided service dictionary. Provide a pointer to an IOBlueotothDevice
// if you wish to associate the record to a specific IOBluetoothDevice.
//
// # Return Value
//
// Returns an IOBluetoothSDPServiceRecord * with the attributes specified in
// the provided dictionary.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withServiceDictionary(_:device:)
func (_IOBluetoothSDPServiceRecordClass IOBluetoothSDPServiceRecordClass) WithServiceDictionaryDevice(serviceDict foundation.INSDictionary, device IIOBluetoothDevice) IOBluetoothSDPServiceRecord {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPServiceRecordClass.class), objc.Sel("withServiceDictionary:device:"), serviceDict, device)
	return IOBluetoothSDPServiceRecordFromID(rv)
}

// Returns an NSDictionary containing the attributes for the service.
//
// # Return Value
//
// Returns an NSDictionary containing the attributes for the target service.
//
// # Discussion
//
// The attribute dictionary is keyed off of the attribute id represented as an
// NSNumber. The values in the NSDictionary are IOBluetoothSDPDataElement
// objects representing the data element for the given attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/attributes
func (b IOBluetoothSDPServiceRecord) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Returns the IOBluetoothDevice that the target service belongs to.
//
// # Return Value
//
// Returns the IOBluetoothDevice that the target service belongs to. If the
// service is one the local host is vending, then nil is returned.
//
// # Discussion
//
// If the service is a local service (i.e. one the current host is vending
// out), then nil is returned.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/device
func (b IOBluetoothSDPServiceRecord) Device() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("device"))
	return IOBluetoothDeviceFromID(objc.ID(rv))
}

// Returns a sorted array of SDP attributes
//
// # Return Value
//
// # Returns a sorted array of SDP attributes
//
// # Discussion
//
// This method will walk all the elements of the service record and return an
// array of IOBluetoothSDPServiceAttribute objects sorted by attributeID
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/sortedAttributes-swift.property
func (b IOBluetoothSDPServiceRecord) SortedAttributes() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sortedAttributes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
