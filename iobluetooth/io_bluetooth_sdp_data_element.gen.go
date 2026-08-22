// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothSDPDataElement] class.
var (
	_IOBluetoothSDPDataElementClass     IOBluetoothSDPDataElementClass
	_IOBluetoothSDPDataElementClassOnce sync.Once
)

func getIOBluetoothSDPDataElementClass() IOBluetoothSDPDataElementClass {
	_IOBluetoothSDPDataElementClassOnce.Do(func() {
		_IOBluetoothSDPDataElementClass = IOBluetoothSDPDataElementClass{class: objc.GetClass("IOBluetoothSDPDataElement")}
	})
	return _IOBluetoothSDPDataElementClass
}

// GetIOBluetoothSDPDataElementClass returns the class object for IOBluetoothSDPDataElement.
func GetIOBluetoothSDPDataElementClass() IOBluetoothSDPDataElementClass {
	return getIOBluetoothSDPDataElementClass()
}

type IOBluetoothSDPDataElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothSDPDataElementClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothSDPDataElementClass) Alloc() IOBluetoothSDPDataElement {
	rv := objc.Send[IOBluetoothSDPDataElement](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An instance of this class represents a single SDP data element as defined
// by the Bluetooth SDP spec.
//
// # Overview
//
// The data types described by the spec have been mapped onto the base
// Foundation classes NSNumber, NSArray, NSData as well as IOBluetoothSDPUUID.
// The number and boolean types (type descriptor 1, 2 and 5) are represented
// as NSNumber objects with the exception of 128-bit numbers which are
// represented as NSData objects in their raw format. The UUID type (type
// descriptor 3) is represented by IOBluetoothSDPUUID. The string and URL
// types (type descriptor 4 and 8) are represented by NSString. The sequence
// types (type descriptor 6 and 7) are represented by NSArray.
//
// Typically, you will not need to create an IOBluetoothSDPDataElement
// directly, the system will do that automatically for both client and server
// operations. However, the current API for adding SDP services to the system
// does allow the use of an NSDictionary based format for creating new
// services. The purpose for that is to allow a service to be built up
// completely in a text file (a plist for example) and then easily imported
// into an app and added to the system without a lot of tedious code to build
// up the entire SDP service record.
//
// The basis for that NSDictionary structure comes from the
// IOBluetoothSDPDataElement. At its simplest, a data element is made up of
// three parts: the type descriptor, the size (from which the size descriptor
// is generated) and the actual value. To provide a complete representation of
// a data element, an NSDictionary with three entries can be used. Each of the
// three entries has a key/value pair representing one of the three attributes
// of a data element. The first key/value pair has a key ‘DataElementType’
// that contains a number value with the actual type descriptor for the data
// element. The second pair has a key ‘DataElementSize’ that contains the
// actual size of the element in bytes. The size descriptor will be calculated
// based on the size and type of the element. The third pair is the value
// itself whose key is ‘DataElementValue’ and whose type corresponds to
// the type mapping above.
//
// In addition to this complete description of a data element, their are some
// shortcuts that can be used for some of the common types and sizes.
//
// If the ‘DataElementType’ value is one of the numeric types (1, 2), the
// ‘DataElementValue’ can be an NSData instead of an NSNumber. In that
// case, the numeric data is taken in network byte order (MSB first).
// Additionally, the ‘DataElementSize’ parameter may be omitted and the
// size will be taken from the length of the data object.
//
// If the ‘DataElementType’ value is the nil type (0), no
// ‘DataElementSize’ or ‘DataElementValue’ entries are needed.
//
// If the ‘DataElementType’ value is any of the other types, the
// ‘DataElementSize’ entry is not needed since the size will be taken
// directly from the value (data, array, string).
//
// In the case where the element is an unsigned, 32-bit integer (type
// descriptor 1, size descriptor 4), the value itself may simply be a number
// (instead of a dictionary as in the previous examples). In the case where
// the element is a UUID (type descriptor 3), the value itself may be a data
// object. The UUID type will be inferred and the size taken from the length
// of the data object.
//
// In the case where the element is a text string (type descriptor 4), the
// value may be a string object. The text string type will be inferred and the
// size taken from the length of the string.
//
// In the case where the element is a data element sequence, the value may be
// an array object. The type will be inferred and the size taken from the
// length of the array. Additionally, the array must contain sub-elements that
// will be parsed out individually.
//
// # Initializers
//
//   - [IOBluetoothSDPDataElement.InitWithElementValue]: Initializes a new IOBluetoothSDPDataElement with the given value.
//   - [IOBluetoothSDPDataElement.InitWithTypeSizeDescriptorSizeValue]: Initializes a new IOBluetoothSDPDataElement with the given attributes.
//
// # Instance Methods
//
//   - [IOBluetoothSDPDataElement.ContainsDataElement]: Checks to see if the target data element is the same as the dataElement parameter or if it contains the dataElement parameter (if its a sequence type).
//   - [IOBluetoothSDPDataElement.ContainsValue]: Checks to see if the target data element’s value is the same as the value parameter or if it contains the value parameter.
//   - [IOBluetoothSDPDataElement.GetArrayValue]: If the data element is represented by an array object, it returns the value as an NSArray.
//   - [IOBluetoothSDPDataElement.GetDataValue]: If the data element is represented by a data object, it returns the value as an NSData.
//   - [IOBluetoothSDPDataElement.GetNumberValue]: If the data element is represented by a number, it returns the value as an NSNumber.
//   - [IOBluetoothSDPDataElement.GetSDPDataElementRef]: Returns an IOBluetoothSDPDataElementRef representation of the target IOBluetoothSDPDataElement object.
//   - [IOBluetoothSDPDataElement.GetSize]: Returns the size in bytes of the target data element.
//   - [IOBluetoothSDPDataElement.GetSizeDescriptor]: Returns the SDP spec defined data element size descriptor for the target data element.
//   - [IOBluetoothSDPDataElement.GetStringValue]: If the data element is represented by a string object, it returns the value as an NSString.
//   - [IOBluetoothSDPDataElement.GetTypeDescriptor]: Returns the SDP spec defined data element type descriptor for the target data element.
//   - [IOBluetoothSDPDataElement.GetUUIDValue]: If the data element is a UUID (type 3), it returns the value as an IOBluetoothSDPUUID.
//   - [IOBluetoothSDPDataElement.GetValue]: Returns the object value of the data element.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement
type IOBluetoothSDPDataElement struct {
	objectivec.Object
}

// IOBluetoothSDPDataElementFromID constructs a [IOBluetoothSDPDataElement] from an objc.ID.
//
// An instance of this class represents a single SDP data element as defined
// by the Bluetooth SDP spec.
func IOBluetoothSDPDataElementFromID(id objc.ID) IOBluetoothSDPDataElement {
	return IOBluetoothSDPDataElement{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothSDPDataElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothSDPDataElement] class.
//
// # Initializers
//
//   - [IIOBluetoothSDPDataElement.InitWithElementValue]: Initializes a new IOBluetoothSDPDataElement with the given value.
//   - [IIOBluetoothSDPDataElement.InitWithTypeSizeDescriptorSizeValue]: Initializes a new IOBluetoothSDPDataElement with the given attributes.
//
// # Instance Methods
//
//   - [IIOBluetoothSDPDataElement.ContainsDataElement]: Checks to see if the target data element is the same as the dataElement parameter or if it contains the dataElement parameter (if its a sequence type).
//   - [IIOBluetoothSDPDataElement.ContainsValue]: Checks to see if the target data element’s value is the same as the value parameter or if it contains the value parameter.
//   - [IIOBluetoothSDPDataElement.GetArrayValue]: If the data element is represented by an array object, it returns the value as an NSArray.
//   - [IIOBluetoothSDPDataElement.GetDataValue]: If the data element is represented by a data object, it returns the value as an NSData.
//   - [IIOBluetoothSDPDataElement.GetNumberValue]: If the data element is represented by a number, it returns the value as an NSNumber.
//   - [IIOBluetoothSDPDataElement.GetSDPDataElementRef]: Returns an IOBluetoothSDPDataElementRef representation of the target IOBluetoothSDPDataElement object.
//   - [IIOBluetoothSDPDataElement.GetSize]: Returns the size in bytes of the target data element.
//   - [IIOBluetoothSDPDataElement.GetSizeDescriptor]: Returns the SDP spec defined data element size descriptor for the target data element.
//   - [IIOBluetoothSDPDataElement.GetStringValue]: If the data element is represented by a string object, it returns the value as an NSString.
//   - [IIOBluetoothSDPDataElement.GetTypeDescriptor]: Returns the SDP spec defined data element type descriptor for the target data element.
//   - [IIOBluetoothSDPDataElement.GetUUIDValue]: If the data element is a UUID (type 3), it returns the value as an IOBluetoothSDPUUID.
//   - [IIOBluetoothSDPDataElement.GetValue]: Returns the object value of the data element.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement
type IIOBluetoothSDPDataElement interface {
	objectivec.IObject

	// Topic: Initializers

	// Initializes a new IOBluetoothSDPDataElement with the given value.
	InitWithElementValue(element objectivec.NSObject) IOBluetoothSDPDataElement
	// Initializes a new IOBluetoothSDPDataElement with the given attributes.
	InitWithTypeSizeDescriptorSizeValue(newType BluetoothSDPDataElementTypeDescriptor, newSizeDescriptor BluetoothSDPDataElementSizeDescriptor, newSize uint32, newValue objectivec.NSObject) IOBluetoothSDPDataElement

	// Topic: Instance Methods

	// Checks to see if the target data element is the same as the dataElement parameter or if it contains the dataElement parameter (if its a sequence type).
	ContainsDataElement(dataElement IIOBluetoothSDPDataElement) bool
	// Checks to see if the target data element’s value is the same as the value parameter or if it contains the value parameter.
	ContainsValue(cmpValue objectivec.NSObject) bool
	// If the data element is represented by an array object, it returns the value as an NSArray.
	GetArrayValue() foundation.INSArray
	// If the data element is represented by a data object, it returns the value as an NSData.
	GetDataValue() foundation.NSData
	// If the data element is represented by a number, it returns the value as an NSNumber.
	GetNumberValue() foundation.NSNumber
	// Returns an IOBluetoothSDPDataElementRef representation of the target IOBluetoothSDPDataElement object.
	GetSDPDataElementRef() IOBluetoothSDPDataElementRef
	// Returns the size in bytes of the target data element.
	GetSize() uint32
	// Returns the SDP spec defined data element size descriptor for the target data element.
	GetSizeDescriptor() BluetoothSDPDataElementSizeDescriptor
	// If the data element is represented by a string object, it returns the value as an NSString.
	GetStringValue() string
	// Returns the SDP spec defined data element type descriptor for the target data element.
	GetTypeDescriptor() BluetoothSDPDataElementTypeDescriptor
	// If the data element is a UUID (type 3), it returns the value as an IOBluetoothSDPUUID.
	GetUUIDValue() IIOBluetoothSDPUUID
	// Returns the object value of the data element.
	GetValue() objectivec.Object

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b IOBluetoothSDPDataElement) Init() IOBluetoothSDPDataElement {
	rv := objc.Send[IOBluetoothSDPDataElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothSDPDataElement) Autorelease() IOBluetoothSDPDataElement {
	rv := objc.Send[IOBluetoothSDPDataElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothSDPDataElement creates a new IOBluetoothSDPDataElement instance.
func NewIOBluetoothSDPDataElement() IOBluetoothSDPDataElement {
	class := getIOBluetoothSDPDataElementClass()
	rv := objc.Send[IOBluetoothSDPDataElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new IOBluetoothSDPDataElement with the given value.
//
// element: The data element value of one of the specified types.
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error parsing the
// element value.
//
// # Discussion
//
// The value must follow the format listed above and must be an instance of
// NSData, NSString, NSNumber, NSArray, NSDictionary, IOBluetoothSDPUUID.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(elementValue:)
func NewBluetoothSDPDataElementWithElementValue(element objectivec.NSObject) IOBluetoothSDPDataElement {
	instance := getIOBluetoothSDPDataElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithElementValue:"), element)
	return IOBluetoothSDPDataElementFromID(rv)
}

// Initializes a new IOBluetoothSDPDataElement with the given attributes.
//
// newType: The type descriptor for the data element.
//
// newSizeDescriptor: The size descriptor for the data element (verify it matches the size
// parameter).
//
// newSize: The size of the data element in bytes (make sure it is a valid size for the
// given size descriptor).
//
// newValue: The raw value itself. This must be the base NSString, NSNumber, NSArray or
// NSData objects. It may not be NSDictionary. If a dictionary format is
// present, use +withElementValue:.
//
// # Return Value
//
// Returns self if successful. Returns nil if an error is encountered (not
// likely due to the limited error checking currently done).
//
// # Discussion
//
// Warning - be careful using this method. There is next to no error checking
// done on the attributes. It is entirely possible to construct an invalid
// data element. It is recommended that +withElementValue: be used instead of
// this one.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(type:sizeDescriptor:size:value:)
func NewBluetoothSDPDataElementWithTypeSizeDescriptorSizeValue(newType BluetoothSDPDataElementTypeDescriptor, newSizeDescriptor BluetoothSDPDataElementSizeDescriptor, newSize uint32, newValue objectivec.NSObject) IOBluetoothSDPDataElement {
	instance := getIOBluetoothSDPDataElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:sizeDescriptor:size:value:"), newType, newSizeDescriptor, newSize, newValue)
	return IOBluetoothSDPDataElementFromID(rv)
}

// Initializes a new IOBluetoothSDPDataElement with the given value.
//
// element: The data element value of one of the specified types.
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error parsing the
// element value.
//
// # Discussion
//
// The value must follow the format listed above and must be an instance of
// NSData, NSString, NSNumber, NSArray, NSDictionary, IOBluetoothSDPUUID.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(elementValue:)
func (b IOBluetoothSDPDataElement) InitWithElementValue(element objectivec.NSObject) IOBluetoothSDPDataElement {
	rv := objc.Send[IOBluetoothSDPDataElement](b.ID, objc.Sel("initWithElementValue:"), element)
	return rv
}

// Initializes a new IOBluetoothSDPDataElement with the given attributes.
//
// newType: The type descriptor for the data element.
//
// newSizeDescriptor: The size descriptor for the data element (verify it matches the size
// parameter).
//
// newSize: The size of the data element in bytes (make sure it is a valid size for the
// given size descriptor).
//
// newValue: The raw value itself. This must be the base NSString, NSNumber, NSArray or
// NSData objects. It may not be NSDictionary. If a dictionary format is
// present, use +withElementValue:.
//
// # Return Value
//
// Returns self if successful. Returns nil if an error is encountered (not
// likely due to the limited error checking currently done).
//
// # Discussion
//
// Warning - be careful using this method. There is next to no error checking
// done on the attributes. It is entirely possible to construct an invalid
// data element. It is recommended that +withElementValue: be used instead of
// this one.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(type:sizeDescriptor:size:value:)
func (b IOBluetoothSDPDataElement) InitWithTypeSizeDescriptorSizeValue(newType BluetoothSDPDataElementTypeDescriptor, newSizeDescriptor BluetoothSDPDataElementSizeDescriptor, newSize uint32, newValue objectivec.NSObject) IOBluetoothSDPDataElement {
	rv := objc.Send[IOBluetoothSDPDataElement](b.ID, objc.Sel("initWithType:sizeDescriptor:size:value:"), newType, newSizeDescriptor, newSize, newValue)
	return rv
}

// Checks to see if the target data element is the same as the dataElement
// parameter or if it contains the dataElement parameter (if its a sequence
// type).
//
// dataElement: The data element to compare with (and search for).
//
// # Return Value
//
// Returns TRUE if the target either matches the given data element or if it
// contains the given data element.
//
// # Discussion
//
// If the target data element is not a sequence type, this method simply
// compares the two data elements. If it is a sequence type, it will search
// through the sequence (and sub-sequences) for the dataElement parameter.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/contains(_:)
func (b IOBluetoothSDPDataElement) ContainsDataElement(dataElement IIOBluetoothSDPDataElement) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("containsDataElement:"), dataElement)
	return rv
}

// Checks to see if the target data element’s value is the same as the value
// parameter or if it contains the value parameter.
//
// cmpValue: The value to compare with (and search for).
//
// # Return Value
//
// Returns TRUE if the target’s value either matches the given value or if
// it contains the given value.
//
// # Discussion
//
// This method works just like -containsDataElement: except that it is
// comparing the value objects directly.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/containsValue(_:)
func (b IOBluetoothSDPDataElement) ContainsValue(cmpValue objectivec.NSObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("containsValue:"), cmpValue)
	return rv
}

// If the data element is represented by an array object, it returns the value
// as an NSArray.
//
// # Return Value
//
// Returns an NSArray representation of the data element if it is a sequence
// type.
//
// # Discussion
//
// The data types represented by an array object are 6 (data element sequence)
// and 7 (data element alternative).
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getArrayValue()
func (b IOBluetoothSDPDataElement) GetArrayValue() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getArrayValue"))
	return foundation.NSArrayFromID(rv)
}

// If the data element is represented by a data object, it returns the value
// as an NSData.
//
// # Return Value
//
// Returns an NSData representation of the data element if it is a 128-bit
// number.
//
// # Discussion
//
// The data types represented by a data object are 128-bit versions of 1
// (unsigned int) and 2 (signed int).
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getDataValue()
func (b IOBluetoothSDPDataElement) GetDataValue() foundation.NSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getDataValue"))
	return foundation.NSDataFromID(rv)
}

// If the data element is represented by a number, it returns the value as an
// NSNumber.
//
// # Return Value
//
// Returns an NSNumber representation of the data element if it is a numeric
// type.
//
// # Discussion
//
// The data types represented by a number are 1 (unsigned int), 2 (signed int)
// and 5 (boolean) except for 128-bit versions of 1 and 2.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getNumberValue()
func (b IOBluetoothSDPDataElement) GetNumberValue() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getNumberValue"))
	return foundation.NSNumberFromID(rv)
}

// Returns an IOBluetoothSDPDataElementRef representation of the target
// IOBluetoothSDPDataElement object.
//
// # Return Value
//
// Returns an IOBluetoothSDPDataElementRef representation of the target
// IOBluetoothSDPDataElement object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getRef()
func (b IOBluetoothSDPDataElement) GetSDPDataElementRef() IOBluetoothSDPDataElementRef {
	rv := objc.Send[IOBluetoothSDPDataElementRef](b.ID, objc.Sel("getSDPDataElementRef"))
	return IOBluetoothSDPDataElementRef(rv)
}

// Returns the size in bytes of the target data element.
//
// # Return Value
//
// Returns the size in bytes of the target data element.
//
// # Discussion
//
// The size is valid whether the data element has a fixed or variable size
// descriptor.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getSize()
func (b IOBluetoothSDPDataElement) GetSize() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("getSize"))
	return rv
}

// Returns the SDP spec defined data element size descriptor for the target
// data element.
//
// # Return Value
//
// Returns the size descriptor for the target data element.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getSizeDescriptor()
func (b IOBluetoothSDPDataElement) GetSizeDescriptor() BluetoothSDPDataElementSizeDescriptor {
	rv := objc.Send[BluetoothSDPDataElementSizeDescriptor](b.ID, objc.Sel("getSizeDescriptor"))
	return BluetoothSDPDataElementSizeDescriptor(rv)
}

// If the data element is represented by a string object, it returns the value
// as an NSString.
//
// # Return Value
//
// Returns an NSString representation of the data element if it is a text or
// URL type.
//
// # Discussion
//
// The data types represented by a string object are 4 (text string) and 8
// (URL).
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getStringValue()
func (b IOBluetoothSDPDataElement) GetStringValue() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getStringValue"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the SDP spec defined data element type descriptor for the target
// data element.
//
// # Return Value
//
// Returns the type descriptor for the target data element.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getTypeDescriptor()
func (b IOBluetoothSDPDataElement) GetTypeDescriptor() BluetoothSDPDataElementTypeDescriptor {
	rv := objc.Send[BluetoothSDPDataElementTypeDescriptor](b.ID, objc.Sel("getTypeDescriptor"))
	return BluetoothSDPDataElementTypeDescriptor(rv)
}

// If the data element is a UUID (type 3), it returns the value as an
// IOBluetoothSDPUUID.
//
// # Return Value
//
// Returns an IOBluetoothSDPUUID representation of the data element if it is a
// UUID.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getUUIDValue()
func (b IOBluetoothSDPDataElement) GetUUIDValue() IIOBluetoothSDPUUID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getUUIDValue"))
	return IOBluetoothSDPUUIDFromID(rv)
}

// Returns the object value of the data element.
//
// # Return Value
//
// Returns the object value of the target data element.
//
// # Discussion
//
// The value returned may be an NSNumber, NSString, NSData, NSArray or
// IOBluetoothSDPDataElement depending on the type of the data element.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getValue()
func (b IOBluetoothSDPDataElement) GetValue() objectivec.Object {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getValue"))
	return objectivec.ObjectFromID(rv)
}
func (b IOBluetoothSDPDataElement) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a new IOBluetoothSDPDataElement with the given value.
//
// element: The data element value of one of the specified types.
//
// # Return Value
//
// Returns the newly allocated data element object. Returns nil if there was
// an error parsing the element value. The returned IOBluetoothSDPDataElement
// object has been autoreleased, so it is not necessary for the caller to
// release it. If the object is to be referenced and kept around, retain
// should be called.
//
// # Discussion
//
// The value must follow the format listed above and must be an instance of
// NSData, NSString, NSNumber, NSArray, NSDictionary, IOBluetoothSDPUUID.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withElementValue(_:)
func (_IOBluetoothSDPDataElementClass IOBluetoothSDPDataElementClass) WithElementValue(element objectivec.NSObject) IOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPDataElementClass.class), objc.Sel("withElementValue:"), element)
	return IOBluetoothSDPDataElementFromID(rv)
}

// Method call to convert an IOBluetoothSDPDataElementRef into an
// IOBluetoothSDPDataElement *.
//
// sdpDataElementRef: IOBluetoothSDPDataElementRef for which an IOBluetoothSDPDataElement * is
// desired.
//
// # Return Value
//
// Returns the IOBluetoothSDPDataElement * for the given
// IOBluetoothSDPDataElementRef.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withSDPDataElementRef(_:)
func (_IOBluetoothSDPDataElementClass IOBluetoothSDPDataElementClass) WithSDPDataElementRef(sdpDataElementRef IOBluetoothSDPDataElementRef) IOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPDataElementClass.class), objc.Sel("withSDPDataElementRef:"), sdpDataElementRef)
	return IOBluetoothSDPDataElementFromID(rv)
}

// Creates a new IOBluetoothSDPDataElement with the given attributes.
//
// type: The type descriptor for the data element.
//
// newSizeDescriptor: The size descriptor for the data element (verify it matches the size
// parameter).
//
// newSize: The size of the data element in bytes (make sure it is a valid size for the
// given size descriptor).
//
// newValue: The raw value itself. This must be the base NSString, NSNumber, NSArray or
// NSData objects. It may not be NSDictionary. If a dictionary format is
// present, use +withElementValue:.
//
// # Return Value
//
// Returns the newly allocated data element object. Returns nil if an error is
// encountered (not likely due to the limited error checking currently done).
// The returned IOBluetoothSDPDataElement object has been autoreleased, so it
// is not necessary for the caller to release it. If the object is to be
// referenced and kept around, retain should be called.
//
// # Discussion
//
// Warning - be careful using this method. There is next to no error checking
// done on the attributes. It is entirely possible to construct an invalid
// data element. It is recommended that +withElementValue: be used instead of
// this one.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withType(_:sizeDescriptor:size:value:)
func (_IOBluetoothSDPDataElementClass IOBluetoothSDPDataElementClass) WithTypeSizeDescriptorSizeValue(type_ BluetoothSDPDataElementTypeDescriptor, newSizeDescriptor BluetoothSDPDataElementSizeDescriptor, newSize uint32, newValue objectivec.NSObject) IOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPDataElementClass.class), objc.Sel("withType:sizeDescriptor:size:value:"), type_, newSizeDescriptor, newSize, newValue)
	return IOBluetoothSDPDataElementFromID(rv)
}
