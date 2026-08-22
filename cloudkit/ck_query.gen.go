// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKQuery] class.
var (
	_CKQueryClass     CKQueryClass
	_CKQueryClassOnce sync.Once
)

func getCKQueryClass() CKQueryClass {
	_CKQueryClassOnce.Do(func() {
		_CKQueryClass = CKQueryClass{class: objc.GetClass("CKQuery")}
	})
	return _CKQueryClass
}

// GetCKQueryClass returns the class object for CKQuery.
func GetCKQueryClass() CKQueryClass {
	return getCKQueryClass()
}

type CKQueryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKQueryClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKQueryClass) Alloc() CKQuery {
	rv := objc.Send[CKQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A query that describes the criteria to apply when searching for records in
// a database.
//
// # Overview
//
// You create a query as the first step in the search process. The query
// stores the search parameters, including the type of records to search, the
// match criteria (predicate) to apply, and the sort parameters to apply to
// the results. Then you use the query to initialize an instance of
// [CKQueryOperation], which you execute to generate the results.
//
// Always designate a record type and predicate when you create a query
// object. The record type narrows the scope of the search to one type of
// record, and the predicate defines the conditions for matching records of
// that type. Predicates usually compare one or more fields of a record to
// constant values, but you can create predicates that return all records of a
// specific type or perform more nuanced searches.
//
// Because you can’t change the record type and predicate after
// initialization, you can use the same query to initialize multiple instances
// of [CKQueryOperation], each of which targets a different database or record
// zone.
//
// # Building Your Predicates
//
// An [NSPredicate] object defines the logical conditions for determining
// whether a record is a match for a query. Queries support only a subset of
// the predicate behaviors that the [NSPredicate] class offers.
//
// # Predicate Rules for Query Objects
//
// The predicates you create for your query objects must follow these rules:
//
// - Predicates derive from a format string. You can’t use value or
// block-based predicates. - Predicates use only the operators in [CKQuery]. -
// Predicates operate only on fields that contain the following types of data:
// - - [NSString] - [NSData] - [NSDate] - [NSNumber] - [NSArray] -
// [CKReference] - [CLLocation] - Key names in predicates correspond to fields
// in the currently evaluated record. Key names can include the names of the
// record’s metadata properties, such as `creationDate`, or any data fields
// you add to the record. You can’t use key paths to specify fields in
// related records. - Predicates support the following variable substitution
// strings: - - Use `%@` for value objects, such as strings, numbers, and
// dates. - Use `%K` for the name of a field. This substitution variable
// indicates that the system uses the substitution string to look up a field
// name. - With one exception, the [CONTAINS] operator is only for testing
// list membership. The exception is when you use it to perform full-text
// searches in conjunction with the `self` key path. The `self` key path
// causes the server to look in searchable string-based fields for the
// specified token string. For example, a predicate string of `@"self contains
// 'blue'"` searches for the word blue in all fields that you mark for
// inclusion in full-text searches. You can’t use the `self` key path to
// search in fields with a type that isn’t a string. - You can combine the
// [ANY] and [SOME] aggregate operators with the [IN] and [CONTAINS] operators
// to perform list membership tests. - The “ operator function performs a
// radius-based location comparison and that comparison must determine whether
// the location value is inside the circular area you provide. You can’t use
// it to search for locations outside the specified circular area. Location
// indexes have a resolution of no less than 10 km. - CloudKit doesn’t
// support the [ALL] aggregate operator. - CloudKit doesn’t support the
// [NOT] compound operator in the following cases: - - You can’t use it to
// negate an [AND] compound predicate. - You can’t use it in tokenized
// queries, such as `self CONTAINS 'value'`. - You can’t use it with the “
// function. - You can’t use it in [BETWEEN] queries.
//
// # Supported Predicate Operators
//
// The following table lists the operators you can use in predicates for a
// query.
//
// [Table data omitted]
//
// Specifying an unsupported operator or data type in your query’s predicate
// results in an error when you execute the query. For more information about
// creating predicate objects, see [Predicate Programming Guide].
//
// # Sample Predicate Format Strings
//
// To match records that link to a different record with an ID you know,
// create a predicate that matches a field that contains a reference as
// Listing 1 shows. In the example, the `employee` field of the record
// contains a [CKReference] object that points to another record. When
// CloudKit executes the query, a match occurs when the ID in the
// locally-created [CKReference] object is the same ID as in the specified
// field of the record.
//
// Listing 1. Matching the ID of a record
//
// To match the contents of a field to a specific value, use a predicate
// similar to the ones in Listing 2. All of the listed predicates generate the
// same set of results, which in the example means that the `favoriteColors`
// field contains the value red. The value in the field must match the value
// you specify in the predicate exactly. String-based comparisons are
// case-insensitive, but otherwise, all comparisons must be an exact match of
// the specified value.
//
// Listing 2. Matching a field to a specific value
//
// You can match more than one value at a time by using a predicate similar to
// the ones in Listing 3. In the example, the predicates report a match if the
// value in the `favoriteColor` field of a record matches either of the values
// `red` or `green`.
//
// Listing 3. Matching a field to one or more values
//
// For fields that contain string values, you can match the beginning portion
// of the string using the [BEGINSWITH] operator as Listing 4 shows. You
// can’t use other string comparison operators, such as [CONTAINS] or
// [ENDSWITH]. When using this operator, the field must contain a string value
// and must start with the string you specify. Matches are case-sensitive. In
// the examples, the predicate matches records where the `favoriteColors`
// field contains the strings red, reddish, or red` `green` `duct` `tape.
//
// Listing 4. Matching a field that starts with a string value
//
// To perform a tokenized search of a record’s fields, use the special
// operator `self`. A tokenized search searches any fields where you enable
// full-text search, which is all string-based fields by default. CloudKit
// treats each distinct word in the tokenized string as a separate token for
// the purpose of searching. Comparisons are case- and diacritic-insensitive.
// These token strings can be in a single field or in multiple fields.
//
// Listing 5 shows an example that searches the fields of a record for the
// token strings `bob` or `smith`:
//
// Listing 5. Matching a field that contains one or more tokens
//
// To search for multiple tokens present in the fields, use the [AND]
// predicate operator, as Listing 6 shows.
//
// Listing 6. Matching a field that contains multiple tokens
//
// To test whether two locations are near each other, create a predicate using
// the “ function as Listing 7 shows. Predicates that use this function must
// have the structure in the listing. In your code, replace the `location`
// variable with a field name from one of your records. This data type for the
// field must be a [CLLocation] object. Similarly, replace the `fixedLoc` and
// `radius` values with appropriate values from your app. The `fixedLoc` value
// is the geographic coordinate that marks the center of a circle with the
// specified radius. In this example, the predicate returns a match if the
// location in the record is within 10 kilometers of the specified latitude
// and longitude.
//
// Listing 7. Matching by distance from a location
//
// To retrieve all records of a specific type, use the [TRUEPREDICATE]
// expression as Listing 8 shows. A predicate with this operator always
// evaluates to `true` and, therefore, matches every record. When using such
// an operator, use a cursor to batch the results into smaller groups for
// processing.
//
// Listing 8. Retrieving all records of a specific type
//
// # Indexes and Full-Text Search
//
// Indexes make it possible to search the contents of your records
// efficiently. During development, the server indexes all fields with data
// types it can use in the predicate of a query. This automatic indexing makes
// it easier to experiment with queries during development, but the indexes
// require space in a database, and require time to generate and maintain. So
// when migrating to a production environment, remove the indexes for any
// fields that you don’t use in queries.
//
// Full-text search is another feature that is on by default for all fields
// during development. When you move to the production environment, disable
// full-text search for fields with content you don’t need to search. As
// with removing indexes, disabling full-text search improves the performance
// of your tokenized searches. To configure the indexing and full-text search
// options for fields in your schema, use CloudKit Dashboard.
//
// In a full-text search, CloudKit ignores the following words if they appear
// in the token strings:
//
// [Table data omitted]
//
// # Executing a Search Using Your Query Object
//
// To execute a query, do one of the following:
//
// - Create an instance of [CKQueryOperation] using your query. Run the
// operation directly or add it to an operation queue to perform the query and
// deliver the results. - Call the [perform(_:inZoneWith:completionHandler:)]
// method of [CKDatabase] to execute the query. Process the results in your
// completion handler.
//
// Queries always run asynchronously and deliver results to a completion
// handler that you provide.
//
// # Creating a Query
//
//   - [CKQuery.InitWithCoder]: Creates an operation group from a serialized instance.
//
// # Accessing the Query Parameters
//
//   - [CKQuery.RecordType]: The record type to search.
//   - [CKQuery.SetRecordType]
//   - [CKQuery.Predicate]: The predicate to use for matching records.
//   - [CKQuery.SortDescriptors]: The sort descriptors for organizing the query’s results.
//   - [CKQuery.SetSortDescriptors]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery
//
// [CLLocation]: https://developer.apple.com/documentation/CoreLocation/CLLocation
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSPredicate]: https://developer.apple.com/documentation/Foundation/NSPredicate
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
// [perform(_:inZoneWith:completionHandler:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/perform(_:inZoneWith:completionHandler:)
type CKQuery struct {
	objectivec.Object
}

// CKQueryFromID constructs a [CKQuery] from an objc.ID.
//
// A query that describes the criteria to apply when searching for records in
// a database.
func CKQueryFromID(id objc.ID) CKQuery {
	return CKQuery{objectivec.Object{ID: id}}
}

// NOTE: CKQuery adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKQuery] class.
//
// # Creating a Query
//
//   - [ICKQuery.InitWithCoder]: Creates an operation group from a serialized instance.
//
// # Accessing the Query Parameters
//
//   - [ICKQuery.RecordType]: The record type to search.
//   - [ICKQuery.SetRecordType]
//   - [ICKQuery.Predicate]: The predicate to use for matching records.
//   - [ICKQuery.SortDescriptors]: The sort descriptors for organizing the query’s results.
//   - [ICKQuery.SetSortDescriptors]
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery
type ICKQuery interface {
	objectivec.IObject

	// Topic: Creating a Query

	// Creates an operation group from a serialized instance.
	InitWithCoder(aDecoder foundation.INSCoder) CKQuery

	// Topic: Accessing the Query Parameters

	// The record type to search.
	RecordType() CKRecordType
	SetRecordType(value CKRecordType)
	// The predicate to use for matching records.
	Predicate() foundation.NSPredicate
	// The sort descriptors for organizing the query’s results.
	SortDescriptors() []foundation.NSSortDescriptor
	SetSortDescriptors(value []foundation.NSSortDescriptor)

	EncodeWithCoder(coder foundation.INSCoder)
	InitWithRecordTypePredicate(recordType CKRecordType, predicate foundation.NSPredicate) CKQuery
}

// Init initializes the instance.
func (c CKQuery) Init() CKQuery {
	rv := objc.Send[CKQuery](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKQuery) Autorelease() CKQuery {
	rv := objc.Send[CKQuery](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQuery creates a new CKQuery instance.
func NewCKQuery() CKQuery {
	class := getCKQueryClass()
	rv := objc.Send[CKQuery](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation group from a serialized instance.
//
// aDecoder: The coder to use when deserializing the group.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery/init(coder:)
func NewCKQueryWithCoder(aDecoder foundation.INSCoder) CKQuery {
	instance := getCKQueryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKQueryFromID(rv)
}

// Creates an operation group from a serialized instance.
//
// aDecoder: The coder to use when deserializing the group.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery/init(coder:)
func (c CKQuery) InitWithCoder(aDecoder foundation.INSCoder) CKQuery {
	rv := objc.Send[CKQuery](c.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}
func (c CKQuery) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (c CKQuery) InitWithRecordTypePredicate(recordType CKRecordType, predicate foundation.NSPredicate) CKQuery {
	rv := objc.Send[CKQuery](c.ID, objc.Sel("initWithRecordType:predicate:"), objc.String(string(recordType)), predicate)
	return rv
}

// The record type to search.
//
// See: https://developer.apple.com/documentation/cloudkit/ckquery/recordtype-6ajii
func (c CKQuery) RecordType() CKRecordType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordType"))
	return CKRecordType(foundation.NSStringFromID(rv).String())
}
func (c CKQuery) SetRecordType(value CKRecordType) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordType:"), objc.String(string(value)))
}

// The predicate to use for matching records.
//
// # Discussion
//
// A predicate contains one or more expressions that evaluate to true or
// false. Expressions are often value-based comparisons, but predicates
// support other types of operators, including string comparisons and
// aggregate operations. For guidelines on how to construct predicates for
// your queries, see [CKQuery].
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery/predicate
func (c CKQuery) Predicate() foundation.NSPredicate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}

// The sort descriptors for organizing the query’s results.
//
// # Discussion
//
// You can add sort descriptors to a query and change them later as necessary.
// Each sort descriptor contains a field name of the intended record type and
// information about whether to sort values in that field in ascending or
// descending order. The default value of this property is `nil`, which means
// that records return in an indeterminate order.
//
// The order of the items in the array defines the order that CloudKit applies
// the sort descriptors to the results. In other words, CloudKit applies the
// first sort descriptor in the array, then the second sort descriptor, if
// necessary, then the third, and so on.
//
// See: https://developer.apple.com/documentation/CloudKit/CKQuery/sortDescriptors
func (c CKQuery) SortDescriptors() []foundation.NSSortDescriptor {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("sortDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSSortDescriptor {
		return foundation.NSSortDescriptorFromID(id)
	})
}
func (c CKQuery) SetSortDescriptors(value []foundation.NSSortDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setSortDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}
