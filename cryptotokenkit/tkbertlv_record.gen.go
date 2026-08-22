// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKBERTLVRecord] class.
var (
	_TKBERTLVRecordClass     TKBERTLVRecordClass
	_TKBERTLVRecordClassOnce sync.Once
)

func getTKBERTLVRecordClass() TKBERTLVRecordClass {
	_TKBERTLVRecordClassOnce.Do(func() {
		_TKBERTLVRecordClass = TKBERTLVRecordClass{class: objc.GetClass("TKBERTLVRecord")}
	})
	return _TKBERTLVRecordClass
}

// GetTKBERTLVRecordClass returns the class object for TKBERTLVRecord.
func GetTKBERTLVRecordClass() TKBERTLVRecordClass {
	return getTKBERTLVRecordClass()
}

type TKBERTLVRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKBERTLVRecordClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKBERTLVRecordClass) Alloc() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An object that parses BER-encoded data and produces DER-encoded data for
// TLV records.
//
// # Creating TLV Records
//
//   - [TKBERTLVRecord.InitWithTagValue]: Initializes a BER-TLV record with the specified tag and value.
//   - [TKBERTLVRecord.InitWithTagRecords]: Initializes a BER-TLV record with the specified tag and an array of TLV subrecords.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord
type TKBERTLVRecord struct {
	TKTLVRecord
}

// TKBERTLVRecordFromID constructs a [TKBERTLVRecord] from an objc.ID.
//
// An object that parses BER-encoded data and produces DER-encoded data for
// TLV records.
func TKBERTLVRecordFromID(id objc.ID) TKBERTLVRecord {
	return TKBERTLVRecord{TKTLVRecord: TKTLVRecordFromID(id)}
}

// NOTE: TKBERTLVRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKBERTLVRecord] class.
//
// # Creating TLV Records
//
//   - [ITKBERTLVRecord.InitWithTagValue]: Initializes a BER-TLV record with the specified tag and value.
//   - [ITKBERTLVRecord.InitWithTagRecords]: Initializes a BER-TLV record with the specified tag and an array of TLV subrecords.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord
type ITKBERTLVRecord interface {
	ITKTLVRecord

	// Topic: Creating TLV Records

	// Initializes a BER-TLV record with the specified tag and value.
	InitWithTagValue(tag TKTLVTag, value foundation.NSData) TKBERTLVRecord
	// Initializes a BER-TLV record with the specified tag and an array of TLV subrecords.
	InitWithTagRecords(tag TKTLVTag, records []TKTLVRecord) TKBERTLVRecord
}

// Init initializes the instance.
func (t TKBERTLVRecord) Init() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKBERTLVRecord) Autorelease() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKBERTLVRecord creates a new TKBERTLVRecord instance.
func NewTKBERTLVRecord() TKBERTLVRecord {
	class := getTKBERTLVRecordClass()
	rv := objc.Send[TKBERTLVRecord](objc.ID(class.class), objc.Sel("new"))
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
func NewTKBERTLVRecordFromData(data foundation.NSData) TKBERTLVRecord {
	rv := objc.Send[objc.ID](objc.ID(getTKBERTLVRecordClass().class), objc.Sel("recordFromData:"), data)
	return TKBERTLVRecordFromID(rv)
}

// Initializes a BER-TLV record with the specified tag and an array of TLV
// subrecords.
//
// tag: The tag field of the record.
//
// records: The TLV subrecords of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag field and subrecords.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:records:)
func NewTKBERTLVRecordWithTagRecords(tag TKTLVTag, records []TKTLVRecord) TKBERTLVRecord {
	instance := getTKBERTLVRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:records:"), tag, objectivec.IObjectSliceToNSArray(records))
	return TKBERTLVRecordFromID(rv)
}

// Initializes a BER-TLV record with the specified tag and value.
//
// tag: The tag field of the record.
//
// value: The value field of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:value:)
func NewTKBERTLVRecordWithTagValue(tag TKTLVTag, value foundation.NSData) TKBERTLVRecord {
	instance := getTKBERTLVRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	return TKBERTLVRecordFromID(rv)
}

// Initializes a BER-TLV record with the specified tag and value.
//
// tag: The tag field of the record.
//
// value: The value field of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:value:)
func (t TKBERTLVRecord) InitWithTagValue(tag TKTLVTag, value foundation.NSData) TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t.ID, objc.Sel("initWithTag:value:"), tag, value)
	return rv
}

// Initializes a BER-TLV record with the specified tag and an array of TLV
// subrecords.
//
// tag: The tag field of the record.
//
// records: The TLV subrecords of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag field and subrecords.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:records:)
func (t TKBERTLVRecord) InitWithTagRecords(tag TKTLVTag, records []TKTLVRecord) TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t.ID, objc.Sel("initWithTag:records:"), tag, objectivec.IObjectSliceToNSArray(records))
	return rv
}

// Encodes a specified tag using BER-TLV tag encoding rules.
//
// tag: The tag value to encode.
//
// # Return Value
//
// A data object that encodes a tag value using BER-TLV encoding.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/data(forTag:)
func (_TKBERTLVRecordClass TKBERTLVRecordClass) DataForTag(tag TKTLVTag) foundation.NSData {
	rv := objc.Send[objc.ID](objc.ID(_TKBERTLVRecordClass.class), objc.Sel("dataForTag:"), tag)
	return foundation.NSDataFromID(rv)
}
