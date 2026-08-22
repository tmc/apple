// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKQueryOperation] class.
var (
	_CKQueryOperationClass     CKQueryOperationClass
	_CKQueryOperationClassOnce sync.Once
)

func getCKQueryOperationClass() CKQueryOperationClass {
	_CKQueryOperationClassOnce.Do(func() {
		_CKQueryOperationClass = CKQueryOperationClass{class: objc.GetClass("CKQueryOperation")}
	})
	return _CKQueryOperationClass
}

// GetCKQueryOperationClass returns the class object for CKQueryOperation.
func GetCKQueryOperationClass() CKQueryOperationClass {
	return getCKQueryOperationClass()
}

type CKQueryOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKQueryOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKQueryOperationClass) Alloc() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation for executing queries in a database.
//
// # Overview
//
// A [CKQueryOperation] object is a concrete operation that you can use to
// execute queries. A query operation applies query parameters to the
// specified database and record zone, delivering any matching records
// asynchronously to the handlers that you provide.
//
// To perform a new search:
//
// - Initialize a [CKQueryOperation] object with a [CKQuery] object that
// contains the search criteria and sorting information for the records you
// want. - Assign a handler to the [queryCompletionBlock] property so that you
// can process the results and execute the operation.
//
// If the search yields many records, the operation object may deliver a
// portion of the total results to your blocks immediately, along with a
// cursor for obtaining the remaining records. Use the cursor to initialize
// and execute a separate [CKQueryOperation] instance when you’re ready to
// process the next batch of results. 3. Optionally, configure the results by
// specifying values for the [CKQueryOperation.ResultsLimit] and [desiredKeys]
// properties. 4. Pass the query operation object to the
// [CKDatabase.AddOperation] method of the target database to execute the
// operation.
//
// CloudKit restricts queries to the records in a single record zone. For new
// queries, you specify the zone when you initialize the query operation
// object. For cursor-based queries, the cursor contains the zone information.
// To search for records in multiple zones, you must create a separate
// [CKQueryOperation] object for each zone you want to search, although you
// can initialize each of them with the same [CKQuery] object.
//
// If you assign a handler to the operation’s [completionBlock] property,
// the operation calls it after it executes and returns any results. Use a
// handler to perform housekeeping tasks for the operation, but don’t use it
// to process the results of the operation. The handler you provide should
// manage any failures, whether due to an error or an explicit cancellation.
//
// # Creating a Query Operation
//
//   - [CKQueryOperation.InitWithQuery]: Creates an operation that searches for records in the specified record zone.
//   - [CKQueryOperation.InitWithCursor]: Creates an operation with additional results from a previous search.
//
// # Configuring the Query Operation
//
//   - [CKQueryOperation.Query]: The query for the search.
//   - [CKQueryOperation.SetQuery]
//   - [CKQueryOperation.Cursor]: The  cursor for continuing the search.
//   - [CKQueryOperation.SetCursor]
//   - [CKQueryOperation.ZoneID]: The ID of the record zone that contains the records to search.
//   - [CKQueryOperation.SetZoneID]
//   - [CKQueryOperation.ResultsLimit]: The maximum number of records to return at one time.
//   - [CKQueryOperation.SetResultsLimit]
//   - [CKQueryOperation.DesiredKeys]: The fields of the records to fetch.
//   - [CKQueryOperation.SetDesiredKeys]
//
// # Instance Properties
//
//   - [CKQueryOperation.RecordMatchedBlock]: The closure to execute when a record match is available.
//   - [CKQueryOperation.SetRecordMatchedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation
//
// [completionBlock]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
// [desiredKeys]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/desiredKeys-4a6vy
// [queryCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
type CKQueryOperation struct {
	CKDatabaseOperation
}

// CKQueryOperationFromID constructs a [CKQueryOperation] from an objc.ID.
//
// An operation for executing queries in a database.
func CKQueryOperationFromID(id objc.ID) CKQueryOperation {
	return CKQueryOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKQueryOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKQueryOperation] class.
//
// # Creating a Query Operation
//
//   - [ICKQueryOperation.InitWithQuery]: Creates an operation that searches for records in the specified record zone.
//   - [ICKQueryOperation.InitWithCursor]: Creates an operation with additional results from a previous search.
//
// # Configuring the Query Operation
//
//   - [ICKQueryOperation.Query]: The query for the search.
//   - [ICKQueryOperation.SetQuery]
//   - [ICKQueryOperation.Cursor]: The  cursor for continuing the search.
//   - [ICKQueryOperation.SetCursor]
//   - [ICKQueryOperation.ZoneID]: The ID of the record zone that contains the records to search.
//   - [ICKQueryOperation.SetZoneID]
//   - [ICKQueryOperation.ResultsLimit]: The maximum number of records to return at one time.
//   - [ICKQueryOperation.SetResultsLimit]
//   - [ICKQueryOperation.DesiredKeys]: The fields of the records to fetch.
//   - [ICKQueryOperation.SetDesiredKeys]
//
// # Instance Properties
//
//   - [ICKQueryOperation.RecordMatchedBlock]: The closure to execute when a record match is available.
//   - [ICKQueryOperation.SetRecordMatchedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation
type ICKQueryOperation interface {
	ICKDatabaseOperation

	// Topic: Creating a Query Operation

	// Creates an operation that searches for records in the specified record zone.
	InitWithQuery(query ICKQuery) CKQueryOperation
	// Creates an operation with additional results from a previous search.
	InitWithCursor(cursor ICKQueryCursor) CKQueryOperation

	// Topic: Configuring the Query Operation

	// The query for the search.
	Query() ICKQuery
	SetQuery(value ICKQuery)
	// The  cursor for continuing the search.
	Cursor() ICKQueryCursor
	SetCursor(value ICKQueryCursor)
	// The ID of the record zone that contains the records to search.
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
	// The maximum number of records to return at one time.
	ResultsLimit() uint
	SetResultsLimit(value uint)
	// The fields of the records to fetch.
	DesiredKeys() []CKRecordFieldKey
	SetDesiredKeys(value []CKRecordFieldKey)

	// Topic: Instance Properties

	// The closure to execute when a record match is available.
	RecordMatchedBlock() unsafe.Pointer
	SetRecordMatchedBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKQueryOperation) Init() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKQueryOperation) Autorelease() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryOperation creates a new CKQueryOperation instance.
func NewCKQueryOperation() CKQueryOperation {
	class := getCKQueryOperationClass()
	rv := objc.Send[CKQueryOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation with additional results from a previous search.
//
// cursor: The cursor that identifies the previous search. CloudKit passes this value
// to the completion handler of the previous search. For more information, see
// the [queryCompletionBlock] property.
//
// # Discussion
//
// Use this method to create an operation that retrieves the next batch of
// results from a previous search. When executing searches for a cursor,
// don’t cache cursors for a long time before using them. A cursor isn’t a
// snapshot of the previous search results; it stores a relative offset into
// the results list. An operation that you create from a cursor performs a new
// search, sorts the new set of results, and uses the previous offset value to
// determine where the next batch of results starts.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(cursor:)
//
// [queryCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
func NewCKQueryOperationWithCursor(cursor ICKQueryCursor) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCursor:"), cursor)
	return CKQueryOperationFromID(rv)
}

// Creates an operation that searches for records in the specified record
// zone.
//
// query: The query for the search.
//
// # Discussion
//
// You can use the operation that this method returns only once to perform a
// search, but you can reuse the query that you provide. During execution, the
// operation performs a new search and returns the first batch of results. If
// there are more results available, you must create a separate query object
// using the provided cursor object.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(query:)
func NewCKQueryOperationWithQuery(query ICKQuery) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithQuery:"), query)
	return CKQueryOperationFromID(rv)
}

// Creates an operation that searches for records in the specified record
// zone.
//
// query: The query for the search.
//
// # Discussion
//
// You can use the operation that this method returns only once to perform a
// search, but you can reuse the query that you provide. During execution, the
// operation performs a new search and returns the first batch of results. If
// there are more results available, you must create a separate query object
// using the provided cursor object.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(query:)
func (c CKQueryOperation) InitWithQuery(query ICKQuery) CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c.ID, objc.Sel("initWithQuery:"), query)
	return rv
}

// Creates an operation with additional results from a previous search.
//
// cursor: The cursor that identifies the previous search. CloudKit passes this value
// to the completion handler of the previous search. For more information, see
// the [queryCompletionBlock] property.
//
// # Discussion
//
// Use this method to create an operation that retrieves the next batch of
// results from a previous search. When executing searches for a cursor,
// don’t cache cursors for a long time before using them. A cursor isn’t a
// snapshot of the previous search results; it stores a relative offset into
// the results list. An operation that you create from a cursor performs a new
// search, sorts the new set of results, and uses the previous offset value to
// determine where the next batch of results starts.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(cursor:)
//
// [queryCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
func (c CKQueryOperation) InitWithCursor(cursor ICKQueryCursor) CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c.ID, objc.Sel("initWithCursor:"), cursor)
	return rv
}

// The query for the search.
//
// # Discussion
//
// The initial value of this property is the query that you provide to the
// [CKQueryOperation.InitWithQuery] method. When the value in the
// [CKQueryOperation.Cursor] property is `nil`, the operation uses this
// property’s value to execute a new search and return its results to your
// completion handler. If [CKQueryOperation.Cursor] isn’t `nil`, the
// operation uses the cursor instead.
//
// If you intend to specify or change the value of this property, do so before
// you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/query
func (c CKQueryOperation) Query() ICKQuery {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("query"))
	return CKQueryFromID(objc.ID(rv))
}
func (c CKQueryOperation) SetQuery(value ICKQuery) {
	objc.Send[struct{}](c.ID, objc.Sel("setQuery:"), value)
}

// The cursor for continuing the search.
//
// # Discussion
//
// The initial value of this property is the cursor that you provide to the
// [CKQueryOperation.InitWithCursor] method. When you use a cursor, the
// operation ignores the contents of the [CKQueryOperation.Query] property.
// This property’s value is an opaque value that CloudKit provides. For more
// information, see the [queryCompletionBlock] property.
//
// If you intend to specify or change the value in this property, do so before
// you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/cursor-swift.property
//
// [queryCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
func (c CKQueryOperation) Cursor() ICKQueryCursor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cursor"))
	return CKQueryCursorFromID(objc.ID(rv))
}
func (c CKQueryOperation) SetCursor(value ICKQueryCursor) {
	objc.Send[struct{}](c.ID, objc.Sel("setCursor:"), value)
}

// The ID of the record zone that contains the records to search.
//
// # Discussion
//
// The value of this property limits the scope of the search to only the
// records in the specified record zone. If you don’t specify a record zone,
// the search includes all record zones.
//
// When you create an operation using the [CKQueryOperation.InitWithCursor]
// method, this property’s value is `nil` and CloudKit ignores any changes
// that you make to it. When the operation executes, the cursor provides the
// record zone information from the original search that provides the cursor.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/zoneID
func (c CKQueryOperation) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
func (c CKQueryOperation) SetZoneID(value ICKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setZoneID:"), value)
}

// The maximum number of records to return at one time.
//
// # Discussion
//
// For most queries, leave the value of this property as the default value,
// which is the [maximumResults] constant. When using that value, CloudKit
// returns as many records as possible while minimizing delays in receiving
// those records. If you want to process a fixed number of results, change the
// value of this property accordingly.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/resultsLimit
//
// [maximumResults]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/maximumResults
func (c CKQueryOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("resultsLimit"))
	return rv
}
func (c CKQueryOperation) SetResultsLimit(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setResultsLimit:"), value)
}

// The fields of the records to fetch.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c CKQueryOperation) DesiredKeys() []CKRecordFieldKey {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("desiredKeys"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordFieldKey {
		return CKRecordFieldKey(foundation.NSStringFromID(id).String())
	})
}
func (c CKQueryOperation) SetDesiredKeys(value []CKRecordFieldKey) {
	objc.Send[struct{}](c.ID, objc.Sel("setDesiredKeys:"), objectivec.StringSliceToNSArray(value))
}

// The closure to execute when a record match is available.
//
// See: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/recordmatchedblock-2qze7
func (c CKQueryOperation) RecordMatchedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordMatchedBlock"))
	return rv
}
func (c CKQueryOperation) SetRecordMatchedBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordMatchedBlock:"), value)
}
