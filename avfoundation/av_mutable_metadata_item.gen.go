// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableMetadataItem] class.
var (
	_AVMutableMetadataItemClass     AVMutableMetadataItemClass
	_AVMutableMetadataItemClassOnce sync.Once
)

func getAVMutableMetadataItemClass() AVMutableMetadataItemClass {
	_AVMutableMetadataItemClassOnce.Do(func() {
		_AVMutableMetadataItemClass = AVMutableMetadataItemClass{class: objc.GetClass("AVMutableMetadataItem")}
	})
	return _AVMutableMetadataItemClass
}

// GetAVMutableMetadataItemClass returns the class object for AVMutableMetadataItem.
func GetAVMutableMetadataItemClass() AVMutableMetadataItemClass {
	return getAVMutableMetadataItemClass()
}

type AVMutableMetadataItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableMetadataItemClass) Alloc() AVMutableMetadataItem {
	rv := objc.Send[AVMutableMetadataItem](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable metadata item for an audiovisual asset or for one of its tracks.
//
// # Overview
// 
// You can initialize a mutable metadata item from an existing
// [AVMetadataItem] object or with a one or more of the basic properties of a
// metadata item: a key, a key space, a locale, and a value.
//
// # Accessing values
//
//   - [AVMutableMetadataItem.Value]: The value for the mutable metadata item.
//   - [AVMutableMetadataItem.SetValue]
//   - [AVMutableMetadataItem.ExtraAttributes]: A dictionary of additional attributes for a metadata item.
//   - [AVMutableMetadataItem.SetExtraAttributes]
//   - [AVMutableMetadataItem.StringValue]: The value of the metadata item as a string.
//   - [AVMutableMetadataItem.SetStringValue]
//   - [AVMutableMetadataItem.NumberValue]: The value of the metadata item as a number.
//   - [AVMutableMetadataItem.SetNumberValue]
//   - [AVMutableMetadataItem.DateValue]: The value of the metadata item as a date.
//   - [AVMutableMetadataItem.SetDateValue]
//   - [AVMutableMetadataItem.DataValue]: The value of the metadata item as a data value.
//   - [AVMutableMetadataItem.SetDataValue]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem
type AVMutableMetadataItem struct {
	AVMetadataItem
}

// AVMutableMetadataItemFromID constructs a [AVMutableMetadataItem] from an objc.ID.
//
// A mutable metadata item for an audiovisual asset or for one of its tracks.
func AVMutableMetadataItemFromID(id objc.ID) AVMutableMetadataItem {
	return AVMutableMetadataItem{AVMetadataItem: AVMetadataItemFromID(id)}
}
// NOTE: AVMutableMetadataItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableMetadataItem] class.
//
// # Accessing values
//
//   - [IAVMutableMetadataItem.Value]: The value for the mutable metadata item.
//   - [IAVMutableMetadataItem.SetValue]
//   - [IAVMutableMetadataItem.ExtraAttributes]: A dictionary of additional attributes for a metadata item.
//   - [IAVMutableMetadataItem.SetExtraAttributes]
//   - [IAVMutableMetadataItem.StringValue]: The value of the metadata item as a string.
//   - [IAVMutableMetadataItem.SetStringValue]
//   - [IAVMutableMetadataItem.NumberValue]: The value of the metadata item as a number.
//   - [IAVMutableMetadataItem.SetNumberValue]
//   - [IAVMutableMetadataItem.DateValue]: The value of the metadata item as a date.
//   - [IAVMutableMetadataItem.SetDateValue]
//   - [IAVMutableMetadataItem.DataValue]: The value of the metadata item as a data value.
//   - [IAVMutableMetadataItem.SetDataValue]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem
type IAVMutableMetadataItem interface {
	IAVMetadataItem

	// Topic: Accessing values

	// The value for the mutable metadata item.
	Value() objectivec.IObject
	SetValue(value objectivec.IObject)
	// A dictionary of additional attributes for a metadata item.
	ExtraAttributes() foundation.INSDictionary
	SetExtraAttributes(value foundation.INSDictionary)
	// The value of the metadata item as a string.
	StringValue() string
	SetStringValue(value string)
	// The value of the metadata item as a number.
	NumberValue() foundation.NSNumber
	SetNumberValue(value foundation.NSNumber)
	// The value of the metadata item as a date.
	DateValue() foundation.INSDate
	SetDateValue(value foundation.INSDate)
	// The value of the metadata item as a data value.
	DataValue() foundation.INSData
	SetDataValue(value foundation.INSData)
}

// Init initializes the instance.
func (m AVMutableMetadataItem) Init() AVMutableMetadataItem {
	rv := objc.Send[AVMutableMetadataItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableMetadataItem) Autorelease() AVMutableMetadataItem {
	rv := objc.Send[AVMutableMetadataItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableMetadataItem creates a new AVMutableMetadataItem instance.
func NewAVMutableMetadataItem() AVMutableMetadataItem {
	class := getAVMutableMetadataItemClass()
	rv := objc.Send[AVMutableMetadataItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a new mutable metadata item.
//
// # Return Value
// 
// A new mutable metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/metadataItem
func (_AVMutableMetadataItemClass AVMutableMetadataItemClass) MetadataItem() AVMutableMetadataItem {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableMetadataItemClass.class), objc.Sel("metadataItem"))
	return AVMutableMetadataItemFromID(rv)
}

// The value for the mutable metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/value
func (m AVMutableMetadataItem) Value() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("value"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMetadataItem) SetValue(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setValue:"), value)
}
// A dictionary of additional attributes for a metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extraAttributes
func (m AVMutableMetadataItem) ExtraAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extraAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m AVMutableMetadataItem) SetExtraAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setExtraAttributes:"), value)
}
// The value of the metadata item as a string.
//
// # Discussion
// 
// This value is `nil` if the system can’t represent the value as a string.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/stringValue
func (m AVMutableMetadataItem) StringValue() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (m AVMutableMetadataItem) SetStringValue(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setStringValue:"), objc.String(value))
}
// The value of the metadata item as a number.
//
// # Discussion
// 
// This value is `nil` if the system can’t represent the value as a number.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/numberValue
func (m AVMutableMetadataItem) NumberValue() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numberValue"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m AVMutableMetadataItem) SetNumberValue(value foundation.NSNumber) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberValue:"), value)
}
// The value of the metadata item as a date.
//
// # Discussion
// 
// This value is `nil` if the system can’t represent the value as a date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dateValue
func (m AVMutableMetadataItem) DateValue() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dateValue"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (m AVMutableMetadataItem) SetDateValue(value foundation.INSDate) {
	objc.Send[struct{}](m.ID, objc.Sel("setDateValue:"), value)
}
// The value of the metadata item as a data value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataValue
func (m AVMutableMetadataItem) DataValue() foundation.INSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataValue"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m AVMutableMetadataItem) SetDataValue(value foundation.INSData) {
	objc.Send[struct{}](m.ID, objc.Sel("setDataValue:"), value)
}

