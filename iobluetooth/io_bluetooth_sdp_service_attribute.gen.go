// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothSDPServiceAttribute] class.
var (
	_IOBluetoothSDPServiceAttributeClass     IOBluetoothSDPServiceAttributeClass
	_IOBluetoothSDPServiceAttributeClassOnce sync.Once
)

func getIOBluetoothSDPServiceAttributeClass() IOBluetoothSDPServiceAttributeClass {
	_IOBluetoothSDPServiceAttributeClassOnce.Do(func() {
		_IOBluetoothSDPServiceAttributeClass = IOBluetoothSDPServiceAttributeClass{class: objc.GetClass("IOBluetoothSDPServiceAttribute")}
	})
	return _IOBluetoothSDPServiceAttributeClass
}

// GetIOBluetoothSDPServiceAttributeClass returns the class object for IOBluetoothSDPServiceAttribute.
func GetIOBluetoothSDPServiceAttributeClass() IOBluetoothSDPServiceAttributeClass {
	return getIOBluetoothSDPServiceAttributeClass()
}

type IOBluetoothSDPServiceAttributeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothSDPServiceAttributeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothSDPServiceAttributeClass) Alloc() IOBluetoothSDPServiceAttribute {
	rv := objc.Send[IOBluetoothSDPServiceAttribute](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
//
// # Overview
//
// A service attribute contains two components: an attribute ID and a data
// element.
//
// # Initializers
//
//   - [IOBluetoothSDPServiceAttribute.InitWithIDAttributeElement]: Initializes a new service attribute with the given ID and data element.
//   - [IOBluetoothSDPServiceAttribute.InitWithIDAttributeElementValue]: Initializes a new service attribute with the given ID and element value.
//
// # Instance Methods
//
//   - [IOBluetoothSDPServiceAttribute.GetDataElement]: Returns the data element for the target service attribute.
//   - [IOBluetoothSDPServiceAttribute.GetAttributeID]: Returns the attribute ID for the target service attribute.
//   - [IOBluetoothSDPServiceAttribute.GetIDDataElement]: Returns the data element representing the attribute ID for the target service attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute
type IOBluetoothSDPServiceAttribute struct {
	objectivec.Object
}

// IOBluetoothSDPServiceAttributeFromID constructs a [IOBluetoothSDPServiceAttribute] from an objc.ID.
//
// IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
func IOBluetoothSDPServiceAttributeFromID(id objc.ID) IOBluetoothSDPServiceAttribute {
	return IOBluetoothSDPServiceAttribute{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothSDPServiceAttribute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothSDPServiceAttribute] class.
//
// # Initializers
//
//   - [IIOBluetoothSDPServiceAttribute.InitWithIDAttributeElement]: Initializes a new service attribute with the given ID and data element.
//   - [IIOBluetoothSDPServiceAttribute.InitWithIDAttributeElementValue]: Initializes a new service attribute with the given ID and element value.
//
// # Instance Methods
//
//   - [IIOBluetoothSDPServiceAttribute.GetDataElement]: Returns the data element for the target service attribute.
//   - [IIOBluetoothSDPServiceAttribute.GetAttributeID]: Returns the attribute ID for the target service attribute.
//   - [IIOBluetoothSDPServiceAttribute.GetIDDataElement]: Returns the data element representing the attribute ID for the target service attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute
type IIOBluetoothSDPServiceAttribute interface {
	objectivec.IObject

	// Topic: Initializers

	// Initializes a new service attribute with the given ID and data element.
	InitWithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID, attributeElement IIOBluetoothSDPDataElement) IOBluetoothSDPServiceAttribute
	// Initializes a new service attribute with the given ID and element value.
	InitWithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID, attributeElementValue objectivec.NSObject) IOBluetoothSDPServiceAttribute

	// Topic: Instance Methods

	// Returns the data element for the target service attribute.
	GetDataElement() IIOBluetoothSDPDataElement
	// Returns the attribute ID for the target service attribute.
	GetAttributeID() BluetoothSDPServiceAttributeID
	// Returns the data element representing the attribute ID for the target service attribute.
	GetIDDataElement() IIOBluetoothSDPDataElement

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b IOBluetoothSDPServiceAttribute) Init() IOBluetoothSDPServiceAttribute {
	rv := objc.Send[IOBluetoothSDPServiceAttribute](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothSDPServiceAttribute) Autorelease() IOBluetoothSDPServiceAttribute {
	rv := objc.Send[IOBluetoothSDPServiceAttribute](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothSDPServiceAttribute creates a new IOBluetoothSDPServiceAttribute instance.
func NewIOBluetoothSDPServiceAttribute() IOBluetoothSDPServiceAttribute {
	class := getIOBluetoothSDPServiceAttributeClass()
	rv := objc.Send[IOBluetoothSDPServiceAttribute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new service attribute with the given ID and data element.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElement: The data element of the new service attribute.
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElement:)
func NewBluetoothSDPServiceAttributeWithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID, attributeElement IIOBluetoothSDPDataElement) IOBluetoothSDPServiceAttribute {
	instance := getIOBluetoothSDPServiceAttributeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithID:attributeElement:"), newAttributeID, attributeElement)
	return IOBluetoothSDPServiceAttributeFromID(rv)
}

// Initializes a new service attribute with the given ID and element value.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElementValue: The data element value of the new service attribute
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error parsing the
// element value.
//
// # Discussion
//
// See +[IOBluetoothSDPDataElement withElementValue:] for a description of the
// types that may be passed in as the attributeElementValue.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElementValue:)
func NewBluetoothSDPServiceAttributeWithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID, attributeElementValue objectivec.NSObject) IOBluetoothSDPServiceAttribute {
	instance := getIOBluetoothSDPServiceAttributeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithID:attributeElementValue:"), newAttributeID, attributeElementValue)
	return IOBluetoothSDPServiceAttributeFromID(rv)
}

// Initializes a new service attribute with the given ID and data element.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElement: The data element of the new service attribute.
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElement:)
func (b IOBluetoothSDPServiceAttribute) InitWithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID, attributeElement IIOBluetoothSDPDataElement) IOBluetoothSDPServiceAttribute {
	rv := objc.Send[IOBluetoothSDPServiceAttribute](b.ID, objc.Sel("initWithID:attributeElement:"), newAttributeID, attributeElement)
	return rv
}

// Initializes a new service attribute with the given ID and element value.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElementValue: The data element value of the new service attribute
//
// # Return Value
//
// Returns self if successful. Returns nil if there was an error parsing the
// element value.
//
// # Discussion
//
// See +[IOBluetoothSDPDataElement withElementValue:] for a description of the
// types that may be passed in as the attributeElementValue.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElementValue:)
func (b IOBluetoothSDPServiceAttribute) InitWithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID, attributeElementValue objectivec.NSObject) IOBluetoothSDPServiceAttribute {
	rv := objc.Send[IOBluetoothSDPServiceAttribute](b.ID, objc.Sel("initWithID:attributeElementValue:"), newAttributeID, attributeElementValue)
	return rv
}

// Returns the data element for the target service attribute.
//
// # Return Value
//
// Returns the data element for the target service attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getDataElement()
func (b IOBluetoothSDPServiceAttribute) GetDataElement() IIOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getDataElement"))
	return IOBluetoothSDPDataElementFromID(rv)
}

// Returns the attribute ID for the target service attribute.
//
// # Return Value
//
// Returns the attribute ID for the target service attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getID()
func (b IOBluetoothSDPServiceAttribute) GetAttributeID() BluetoothSDPServiceAttributeID {
	rv := objc.Send[BluetoothSDPServiceAttributeID](b.ID, objc.Sel("getAttributeID"))
	return BluetoothSDPServiceAttributeID(rv)
}

// Returns the data element representing the attribute ID for the target
// service attribute.
//
// # Return Value
//
// Returns the data element representing the attribute ID for the target
// service attribute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getIDDataElement()
func (b IOBluetoothSDPServiceAttribute) GetIDDataElement() IIOBluetoothSDPDataElement {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getIDDataElement"))
	return IOBluetoothSDPDataElementFromID(rv)
}
func (b IOBluetoothSDPServiceAttribute) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a new service attribute with the given ID and data element.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElement: The data element of the new service attribute.
//
// # Return Value
//
// Returns the newly allocated service attribute object. Returns nil if there
// was an error. The returned IOBluetoothSDPDataElement object has been
// autoreleased, so it is not necessary for the caller to release it. If the
// object is to be referenced and kept around, retain should be called.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/withID(_:attributeElement:)
func (_IOBluetoothSDPServiceAttributeClass IOBluetoothSDPServiceAttributeClass) WithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID, attributeElement IIOBluetoothSDPDataElement) IOBluetoothSDPServiceAttribute {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPServiceAttributeClass.class), objc.Sel("withID:attributeElement:"), newAttributeID, attributeElement)
	return IOBluetoothSDPServiceAttributeFromID(rv)
}

// Creates a new service attribute with the given ID and element value.
//
// newAttributeID: The attribute ID of the new service attribute.
//
// attributeElementValue: The data element value of the new service attribute
//
// # Return Value
//
// Returns the newly allocated service attribute object. Returns nil if there
// was an error parsing the element value. The returned
// IOBluetoothSDPDataElement object has been autoreleased, so it is not
// necessary for the caller to release it. If the object is to be referenced
// and kept around, retain should be called.
//
// # Discussion
//
// See +[IOBluetoothSDPDataElement withElementValue:] for a description of the
// types that may be passed in as the attributeElementValue.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/withID(_:attributeElementValue:)
func (_IOBluetoothSDPServiceAttributeClass IOBluetoothSDPServiceAttributeClass) WithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID, attributeElementValue objectivec.NSObject) IOBluetoothSDPServiceAttribute {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothSDPServiceAttributeClass.class), objc.Sel("withID:attributeElementValue:"), newAttributeID, attributeElementValue)
	return IOBluetoothSDPServiceAttributeFromID(rv)
}
