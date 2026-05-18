// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKQueryCursor] class.
var (
	_CKQueryCursorClass     CKQueryCursorClass
	_CKQueryCursorClassOnce sync.Once
)

func getCKQueryCursorClass() CKQueryCursorClass {
	_CKQueryCursorClassOnce.Do(func() {
		_CKQueryCursorClass = CKQueryCursorClass{class: objc.GetClass("CKQueryCursor")}
	})
	return _CKQueryCursorClass
}

// GetCKQueryCursorClass returns the class object for CKQueryCursor.
func GetCKQueryCursorClass() CKQueryCursorClass {
	return getCKQueryCursorClass()
}

type CKQueryCursorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKQueryCursorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKQueryCursorClass) Alloc() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that marks the stopping point for a query and the starting point
// for retrieving the remaining results.
//
// # Overview
//
// You don’t create instances of this class yourself. When fetching records
// using a query operation, if the number of results exceeds the limit for the
// query, CloudKit provides a cursor. Use that cursor to create a new instance
// of [CKQueryOperation] and retrieve the next batch of results for the same
// query.
//
// For information about how to use a [CKQueryCursor] object, see
// [CKQueryOperation].
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/Cursor-swift.class
type CKQueryCursor struct {
	objectivec.Object
}

// CKQueryCursorFromID constructs a [CKQueryCursor] from an objc.ID.
//
// An object that marks the stopping point for a query and the starting point
// for retrieving the remaining results.
func CKQueryCursorFromID(id objc.ID) CKQueryCursor {
	return CKQueryCursor{objectivec.Object{ID: id}}
}

// NOTE: CKQueryCursor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKQueryCursor] class.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/Cursor-swift.class
type ICKQueryCursor interface {
	objectivec.IObject

	// The  cursor for continuing the search.
	Cursor() ICKQueryCursor
	SetCursor(value ICKQueryCursor)
	// The fields of the records to fetch.
	DesiredKeys() string
	SetDesiredKeys(value string)
	// The query for the search.
	Query() ICKQuery
	SetQuery(value ICKQuery)
	// The maximum number of records to return at one time.
	ResultsLimit() int
	SetResultsLimit(value int)
	// The ID of the record zone that contains the records to search.
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKQueryCursor) Init() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKQueryCursor) Autorelease() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryCursor creates a new CKQueryCursor instance.
func NewCKQueryCursor() CKQueryCursor {
	class := getCKQueryCursorClass()
	rv := objc.Send[CKQueryCursor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKQueryCursor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The cursor for continuing the search.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/cursor-swift.property
func (c CKQueryCursor) Cursor() ICKQueryCursor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cursor"))
	return CKQueryCursorFromID(objc.ID(rv))
}
func (c CKQueryCursor) SetCursor(value ICKQueryCursor) {
	objc.Send[struct{}](c.ID, objc.Sel("setCursor:"), value)
}

// The fields of the records to fetch.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c CKQueryCursor) DesiredKeys() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("desiredKeys"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKQueryCursor) SetDesiredKeys(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDesiredKeys:"), objc.String(value))
}

// The query for the search.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/query
func (c CKQueryCursor) Query() ICKQuery {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("query"))
	return CKQueryFromID(objc.ID(rv))
}
func (c CKQueryCursor) SetQuery(value ICKQuery) {
	objc.Send[struct{}](c.ID, objc.Sel("setQuery:"), value)
}

// The maximum number of records to return at one time.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/resultslimit
func (c CKQueryCursor) ResultsLimit() int {
	rv := objc.Send[int](c.ID, objc.Sel("resultsLimit"))
	return rv
}
func (c CKQueryCursor) SetResultsLimit(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setResultsLimit:"), value)
}

// The ID of the record zone that contains the records to search.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c CKQueryCursor) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
func (c CKQueryCursor) SetZoneID(value ICKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setZoneID:"), value)
}

// A constant value that represents the maximum number of results CloudKit
// retrieves.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/maximumresults
func (_CKQueryCursorClass CKQueryCursorClass) MaximumResults() int {
	rv := objc.Send[int](objc.ID(_CKQueryCursorClass.class), objc.Sel("CKQueryOperationMaximumResults"))
	return rv
}
