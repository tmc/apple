// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTLVRecord] class.
var (
	_TKTLVRecordClass     TKTLVRecordClass
	_TKTLVRecordClassOnce sync.Once
)

func getTKTLVRecordClass() TKTLVRecordClass {
	_TKTLVRecordClassOnce.Do(func() {
		_TKTLVRecordClass = TKTLVRecordClass{class: objc.GetClass("TKTLVRecord")}
	})
	return _TKTLVRecordClass
}

// GetTKTLVRecordClass returns the class object for TKTLVRecord.
func GetTKTLVRecordClass() TKTLVRecordClass {
	return getTKTLVRecordClass()
}

type TKTLVRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTLVRecordClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTLVRecordClass) Alloc() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// The base class encapsulating a Tag-Length-Value record.
//
// # Overview
//
// The CryptoTokenKit framework provides the following concrete subclasses for
// various TLV record encodings:
//
// - [TKBERTLVRecord] for BER-TLV encoding rules - [TKSimpleTLVRecord] for
// Simple-TLV encoding according to ISO 7816-4 - [TKCompactTLVRecord] for
// Compact-TLV encoding according to ISO 7816-4
//
// # Accessing the Tag Field
//
//   - [TKTLVRecord.Tag]: The tag field of the record.
//
// # Accessing the Value Field
//
//   - [TKTLVRecord.Value]: The value field of the record.
//
// # Accessing Record Data
//
//   - [TKTLVRecord.Data]: The record data, including the tag, length, and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord
type TKTLVRecord struct {
	objectivec.Object
}

// TKTLVRecordFromID constructs a [TKTLVRecord] from an objc.ID.
//
// The base class encapsulating a Tag-Length-Value record.
func TKTLVRecordFromID(id objc.ID) TKTLVRecord {
	return TKTLVRecord{objectivec.Object{ID: id}}
}

// NOTE: TKTLVRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTLVRecord] class.
//
// # Accessing the Tag Field
//
//   - [ITKTLVRecord.Tag]: The tag field of the record.
//
// # Accessing the Value Field
//
//   - [ITKTLVRecord.Value]: The value field of the record.
//
// # Accessing Record Data
//
//   - [ITKTLVRecord.Data]: The record data, including the tag, length, and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord
type ITKTLVRecord interface {
	objectivec.IObject

	// Topic: Accessing the Tag Field

	// The tag field of the record.
	Tag() TKTLVTag

	// Topic: Accessing the Value Field

	// The value field of the record.
	Value() foundation.NSData

	// Topic: Accessing Record Data

	// The record data, including the tag, length, and value fields.
	Data() foundation.NSData
}

// Init initializes the instance.
func (t TKTLVRecord) Init() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTLVRecord) Autorelease() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTLVRecord creates a new TKTLVRecord instance.
func NewTKTLVRecord() TKTLVRecord {
	class := getTKTLVRecordClass()
	rv := objc.Send[TKTLVRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a TLV record from by parsing the specified data.
//
// data: A data object containing the serialized representation of a TLV record.
//
// # Return Value
//
// A TLV record, or `nil` if `data` does not specify a valid record.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/init(from:)
func NewTKTLVRecordFromData(data foundation.NSData) TKTLVRecord {
	rv := objc.Send[objc.ID](objc.ID(getTKTLVRecordClass().class), objc.Sel("recordFromData:"), data)
	return TKTLVRecordFromID(rv)
}

// Creates and returns an array of TLV records from the specified data.
//
// data: A data object containing the serialized representation of zero or more TLV
// records.
//
// # Return Value
//
// A sequence of TLV records, or `nil` if `data` does not specify a sequence
// of valid records.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/sequenceOfRecords(from:)
func (_TKTLVRecordClass TKTLVRecordClass) SequenceOfRecordsFromData(data foundation.NSData) []TKTLVRecord {
	rv := objc.Send[[]objc.ID](objc.ID(_TKTLVRecordClass.class), objc.Sel("sequenceOfRecordsFromData:"), data)
	return objc.ConvertSlice(rv, func(id objc.ID) TKTLVRecord {
		return TKTLVRecordFromID(id)
	})
}

// The tag field of the record.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/tag
func (t TKTLVRecord) Tag() TKTLVTag {
	rv := objc.Send[TKTLVTag](t.ID, objc.Sel("tag"))
	return TKTLVTag(rv)
}

// The value field of the record.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/value
func (t TKTLVRecord) Value() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("value"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The record data, including the tag, length, and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/data
func (t TKTLVRecord) Data() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
