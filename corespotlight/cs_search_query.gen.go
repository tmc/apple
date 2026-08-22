// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSSearchQuery] class.
var (
	_CSSearchQueryClass     CSSearchQueryClass
	_CSSearchQueryClassOnce sync.Once
)

func getCSSearchQueryClass() CSSearchQueryClass {
	_CSSearchQueryClassOnce.Do(func() {
		_CSSearchQueryClass = CSSearchQueryClass{class: objc.GetClass("CSSearchQuery")}
	})
	return _CSSearchQueryClass
}

// GetCSSearchQueryClass returns the class object for CSSearchQuery.
func GetCSSearchQueryClass() CSSearchQueryClass {
	return getCSSearchQueryClass()
}

type CSSearchQueryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSearchQueryClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSearchQueryClass) Alloc() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A type you use to programmatically search the indexed app content.
//
// # Overview
//
// Use a [CSSearchQuery] object to search your app’s indexed content using a
// formatted search string. To perform a search, build a predicate string to
// specify the indexed attributes you want to search and the value you want
// them to match. After you start the query, you receive batches of results in
// the handlers you provide.
//
// Each [CSSearchQuery] object you create performs a single search operation
// and delivers the results back to your code. Build each predicate with an
// attribute name, one or more values, and either a comparison operator or the
// [InRange] operator. Your predicate string takes one of the following forms:
//
// - `attributeName operator value[modifiers]` - `InRange(attributeName,
// minValue, maxValue)`
//
// Queries search all of your app’s indexes by default. If your app encrypts
// some of its indexed data, you can limit your search to one or more of the
// encrypted indexes by updating the query’s
// [CSSearchQuery.ProtectionClasses] property. The query must have access to
// the protected index to search it.
//
// For more information about how to construct predicate strings for your
// query, see [Searching for information in your app].
//
// # Creating a query object
//
//   - [CSSearchQuery.InitWithQueryStringQueryContext]: Initializes and returns a query object with the specified query string and query context.
//
// # Specifying the indexes to search
//
//   - [CSSearchQuery.ProtectionClasses]: The protection types of the indexes you want to search.
//   - [CSSearchQuery.SetProtectionClasses]
//
// # Executing the query automatically
//
//   - [CSSearchQuery.Results]: The results that match the current query string.
//   - [CSSearchQuery.SetResults]
//
// # Executing the query with handler blocks
//
//   - [CSSearchQuery.Start]: Starts searching the index for items that match the current query string and parameters.
//   - [CSSearchQuery.Cancel]: Cancels the current query operation.
//   - [CSSearchQuery.IsCancelled]: A Boolean value that indicates whether the current query is no longer running.
//   - [CSSearchQuery.FoundItemCount]: The number of matching items found for the given query string.
//   - [CSSearchQuery.FoundItemsHandler]: The block to execute when the query delivers a new batch of matching items.
//   - [CSSearchQuery.SetFoundItemsHandler]
//   - [CSSearchQuery.CompletionHandler]: The block to execute when the query finishes delivering all results.
//   - [CSSearchQuery.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery
//
// [Searching for information in your app]: https://developer.apple.com/documentation/CoreSpotlight/searching-for-information-in-your-app
type CSSearchQuery struct {
	objectivec.Object
}

// CSSearchQueryFromID constructs a [CSSearchQuery] from an objc.ID.
//
// A type you use to programmatically search the indexed app content.
func CSSearchQueryFromID(id objc.ID) CSSearchQuery {
	return CSSearchQuery{objectivec.Object{ID: id}}
}

// NOTE: CSSearchQuery adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSearchQuery] class.
//
// # Creating a query object
//
//   - [ICSSearchQuery.InitWithQueryStringQueryContext]: Initializes and returns a query object with the specified query string and query context.
//
// # Specifying the indexes to search
//
//   - [ICSSearchQuery.ProtectionClasses]: The protection types of the indexes you want to search.
//   - [ICSSearchQuery.SetProtectionClasses]
//
// # Executing the query automatically
//
//   - [ICSSearchQuery.Results]: The results that match the current query string.
//   - [ICSSearchQuery.SetResults]
//
// # Executing the query with handler blocks
//
//   - [ICSSearchQuery.Start]: Starts searching the index for items that match the current query string and parameters.
//   - [ICSSearchQuery.Cancel]: Cancels the current query operation.
//   - [ICSSearchQuery.IsCancelled]: A Boolean value that indicates whether the current query is no longer running.
//   - [ICSSearchQuery.FoundItemCount]: The number of matching items found for the given query string.
//   - [ICSSearchQuery.FoundItemsHandler]: The block to execute when the query delivers a new batch of matching items.
//   - [ICSSearchQuery.SetFoundItemsHandler]
//   - [ICSSearchQuery.CompletionHandler]: The block to execute when the query finishes delivering all results.
//   - [ICSSearchQuery.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery
type ICSSearchQuery interface {
	objectivec.IObject

	// Topic: Creating a query object

	// Initializes and returns a query object with the specified query string and query context.
	InitWithQueryStringQueryContext(queryString string, queryContext ICSSearchQueryContext) CSSearchQuery

	// Topic: Specifying the indexes to search

	// The protection types of the indexes you want to search.
	ProtectionClasses() []string
	SetProtectionClasses(value []string)

	// Topic: Executing the query automatically

	// The results that match the current query string.
	Results() unsafe.Pointer
	SetResults(value unsafe.Pointer)

	// Topic: Executing the query with handler blocks

	// Starts searching the index for items that match the current query string and parameters.
	Start()
	// Cancels the current query operation.
	Cancel()
	// A Boolean value that indicates whether the current query is no longer running.
	IsCancelled() bool
	// The number of matching items found for the given query string.
	FoundItemCount() uint
	// The block to execute when the query delivers a new batch of matching items.
	FoundItemsHandler() CSSearchableItemArrayHandler
	SetFoundItemsHandler(value CSSearchableItemArrayHandler)
	// The block to execute when the query finishes delivering all results.
	CompletionHandler() ErrorHandler
	SetCompletionHandler(value ErrorHandler)
}

// Init initializes the instance.
func (c CSSearchQuery) Init() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSearchQuery) Autorelease() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchQuery creates a new CSSearchQuery instance.
func NewCSSearchQuery() CSSearchQuery {
	class := getCSSearchQueryClass()
	rv := objc.Send[CSSearchQuery](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a query object with the specified query string and
// query context.
//
// # Return Value
//
// An initialized query object.
//
// # Discussion
//
// - queryString: A formatted string that defines the matching criteria to
// apply to indexed items. To learn how to construct a query string, see
// [Create a query string for your search]. - This parameter must not be
// `nil`. - - queryContext: A [CSSearchQueryContext] object that focuses the
// query results.
//
// # Discussion
//
// After you create and initialize a query object, and call
// [CSSearchQuery.Start] to begin the query, you can’t update or reuse the
// query object for a new query.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/init(queryString:queryContext:)
//
// [Create a query string for your search]: https://developer.apple.com/documentation/CoreSpotlight/searching-for-information-in-your-app#Create-a-query-string-for-your-search
func NewCSSearchQueryWithQueryStringQueryContext(queryString string, queryContext ICSSearchQueryContext) CSSearchQuery {
	instance := getCSSearchQueryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithQueryString:queryContext:"), objc.String(queryString), queryContext)
	return CSSearchQueryFromID(rv)
}

// Initializes and returns a query object with the specified query string and
// query context.
//
// # Return Value
//
// An initialized query object.
//
// # Discussion
//
// - queryString: A formatted string that defines the matching criteria to
// apply to indexed items. To learn how to construct a query string, see
// [Create a query string for your search]. - This parameter must not be
// `nil`. - - queryContext: A [CSSearchQueryContext] object that focuses the
// query results.
//
// # Discussion
//
// After you create and initialize a query object, and call
// [CSSearchQuery.Start] to begin the query, you can’t update or reuse the
// query object for a new query.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/init(queryString:queryContext:)
//
// [Create a query string for your search]: https://developer.apple.com/documentation/CoreSpotlight/searching-for-information-in-your-app#Create-a-query-string-for-your-search
func (c CSSearchQuery) InitWithQueryStringQueryContext(queryString string, queryContext ICSSearchQueryContext) CSSearchQuery {
	rv := objc.Send[CSSearchQuery](c.ID, objc.Sel("initWithQueryString:queryContext:"), objc.String(queryString), queryContext)
	return rv
}

// Starts searching the index for items that match the current query string
// and parameters.
//
// # Discussion
//
// This method uses the configured search parameters and query string to begin
// a search of the index. It then delivers the results of that search to the
// closures in the [CSSearchQuery.FoundItemsHandler] and
// [CSSearchQuery.CompletionHandler] properties of the query object.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/start()
func (c CSSearchQuery) Start() {
	objc.Send[objc.ID](c.ID, objc.Sel("start"))
}

// Cancels the current query operation.
//
// # Discussion
//
// Cancel a query operation when you no longer need the results. You can use
// this method to cancel a query you started, by using either the
// [CSSearchQuery.Start] method or by accessing the [CSSearchQuery.Results]
// property. After you cancel a query operation, you can’t restart it. To
// perform a new query, create a new query object. For example, you might
// cancel one query and start a new one when the text in your app’s search
// control changes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/cancel()
func (c CSSearchQuery) Cancel() {
	objc.Send[objc.ID](c.ID, objc.Sel("cancel"))
}

// The protection types of the indexes you want to search.
//
// # Discussion
//
// When creating an index to store your app’s data, you can associate a file
// protection type with that index to encrypt the data it contains. File
// protections prevent access to the data unless specific conditions occur.
// For example, you can make the data available only when the device is
// unlocked, or make it available after the first successful unlocking of the
// device. This property indicates which of your app’s protected indexes you
// want to search.
//
// Possible values for this property are [complete], [completeUnlessOpen], and
// [completeUntilFirstUserAuthentication]. By default, the system retrives the
// data protection class from the
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeDefault()-data-protection`
// entitlement, if it exists; otherwise, the default value is
// [completeUntilFirstUserAuthentication].
//
// For more information on how to configure file protection types with an
// index, see [Adding your app’s content to Spotlight indexes].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/protectionClasses
//
// [Adding your app’s content to Spotlight indexes]: https://developer.apple.com/documentation/CoreSpotlight/adding-your-app-s-content-to-spotlight-indexes
// [completeUnlessOpen]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUnlessOpen
// [completeUntilFirstUserAuthentication]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUntilFirstUserAuthentication
// [complete]: https://developer.apple.com/documentation/Foundation/FileProtectionType/complete
func (c CSSearchQuery) ProtectionClasses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("protectionClasses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchQuery) SetProtectionClasses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProtectionClasses:"), objectivec.StringSliceToNSArray(value))
}

// The results that match the current query string.
//
// See: https://developer.apple.com/documentation/corespotlight/cssearchquery/results-swift.property
func (c CSSearchQuery) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("results"))
	return rv
}
func (c CSSearchQuery) SetResults(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setResults:"), value)
}

// A Boolean value that indicates whether the current query is no longer
// running.
//
// # Discussion
//
// The value of this property is `true` if you canceled it, and `false` if
// it’s still running or able to run.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/isCancelled
func (c CSSearchQuery) IsCancelled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCancelled"))
	return rv
}

// The number of matching items found for the given query string.
//
// # Discussion
//
// As Spotlight finds matches to the query string, it updates the value of
// this property and delivers the new results to the handler in the
// [CSSearchQuery.FoundItemsHandler] property. This value reflects the total
// number of matching items, including the number of items the query delivered
// to your handler block previously.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/foundItemCount
func (c CSSearchQuery) FoundItemCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("foundItemCount"))
	return rv
}

// The block to execute when the query delivers a new batch of matching items.
//
// # Discussion
//
// Assign a block to this property that has no return value and takes the
// following parameter:
//
// items: An array of items that match the specified query.
//
// Specify a value for this property only if you start your query with the
// [CSSearchQuery.Start] method. While the query runs, the query object
// executes the provided closure one or more times to deliver suggested search
// results for the current search term. Use your handler to retrieve the those
// results and display them. The query object stops delivering matching items
// when it reaches the maximum number found in the
// [CSUserQueryContext.MaxResultCount] property of the query configuration
// parameters.
//
// Each time the query object calls your handler, it delivers only the new
// items that it found since the last time it called your handler. However,
// the query object updates the value of the [CSSearchQuery.FoundItemCount]
// property with the total number of items before it calls your handler.
//
// If you start the query by accessing the [CSSearchQuery.Results] property of
// [CSSearchQuery] or the [CSUserQuery.Responses] property of [CSUserQuery],
// the query object doesn’t execute the block in this property. property of
// [CSUserQuery], the query object doesn’t execute the block in this
// property.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/foundItemsHandler
func (c CSSearchQuery) FoundItemsHandler() CSSearchableItemArrayHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("foundItemsHandler"))
	_ = rv
	return nil
}
func (c CSSearchQuery) SetFoundItemsHandler(value CSSearchableItemArrayHandler) {
	block, cleanup := NewCSSearchableItemArrayBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setFoundItemsHandler:"), block)
}

// The block to execute when the query finishes delivering all results.
//
// # Discussion
//
// Specify a value for this property only if you start your query with the
// [CSSearchQuery.Start] method. When the query finishes, the query object
// executes the provided closure once to let you know the search is complete.
// Use your handler to perform any related cleanup. The block you assign to
// this property returns no parameters and takes the following parameter:
//
// error: An error object with details about a problem that occurred, or `nil`
// if the query completed successfully.
//
// If you start the query by accessing the [CSSearchQuery.Results] property of
// [CSSearchQuery] or the [CSUserQuery.Responses] property of [CSUserQuery],
// the query object doesn’t execute the block in this property.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/completionHandler
func (c CSSearchQuery) CompletionHandler() ErrorHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("completionHandler"))
	_ = rv
	return nil
}
func (c CSSearchQuery) SetCompletionHandler(value ErrorHandler) {
	block, cleanup := NewErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setCompletionHandler:"), block)
}
