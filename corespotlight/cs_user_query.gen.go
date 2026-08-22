// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSUserQuery] class.
var (
	_CSUserQueryClass     CSUserQueryClass
	_CSUserQueryClassOnce sync.Once
)

func getCSUserQueryClass() CSUserQueryClass {
	_CSUserQueryClassOnce.Do(func() {
		_CSUserQueryClass = CSUserQueryClass{class: objc.GetClass("CSUserQuery")}
	})
	return _CSUserQueryClass
}

// GetCSUserQueryClass returns the class object for CSUserQuery.
func GetCSUserQueryClass() CSUserQueryClass {
	return getCSUserQueryClass()
}

type CSUserQueryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSUserQueryClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSUserQueryClass) Alloc() CSUserQuery {
	rv := objc.Send[CSUserQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A type you use to initiate searches from your interface and offer suggested
// text completions.
//
// # Overview
//
// A [CSUserQuery] object provides the back-end support for your app’s
// search features. Combine this object with your app’s search interface to
// perform lexical and semantic searches of human-entered search terms. You
// can configure a query object to return ranked or unranked results. You can
// also use it to get a list of suggestions to display from your search
// interface.
//
// When the text in your search control changes, create a query object to
// begin searching for results based on the current text. You use a query
// object only once to perform a search. If the text changes again while a
// previous query is in progress, cancel the old query and execute the new
// one. For this reason, it’s a good idea to delay the start of each query
// until there is a sufficient gap between changes.
//
// Configure the query parameters using a [CSUserQueryContext] object, which
// you can reuse for multiple queries. The context lets you configure the
// behavior for ranking results, specify the maximum number of results and
// suggestions, and filter the results using a predicate string. When you’re
// ready to start the query, choose one of the following options:
//
// - Get the value of the [CSUserQuery.Responses] property and iterate over
// the results. - Configure the [CSSearchQuery.FoundItemsHandler] property and
// call [CSUserQuery.Start] to execute the query manually.
//
// Each query runs until Spotlight returns the requested maximum number of
// results. If you don’t specify the maximum number of results, Spotlight
// runs until it returns all results. To end a search before you receive all
// the results, call the [CSUserQuery.Cancel] method. Cancelling a query is
// especially important if you’re about to start a new query with an updated
// search string.
//
// For more information about configuring a [CSUserQuery] object, see
// [Building a search interface for your app].
//
// # Creating a user query
//
//   - [CSUserQuery.InitWithUserQueryStringUserQueryContext]: Creates a new user query that searches for the specified term.
//
// # Executing the query automatically
//
//   - [CSUserQuery.Responses]: The matching results and suggestions for the current query string.
//   - [CSUserQuery.SetResponses]
//   - [CSUserQuery.Suggestions]: An asynchronous sequence of suggested completions for the current query text.
//   - [CSUserQuery.SetSuggestions]
//
// # Executing the query with handler blocks
//
//   - [CSUserQuery.FoundSuggestionsHandler]: The block to execute when the query delivers a new batch of suggested items.
//   - [CSUserQuery.SetFoundSuggestionsHandler]
//   - [CSUserQuery.FoundSuggestionCount]: The number of suggested items the query found so far.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery
//
// [Building a search interface for your app]: https://developer.apple.com/documentation/CoreSpotlight/building-a-search-interface-for-your-app
type CSUserQuery struct {
	CSSearchQuery
}

// CSUserQueryFromID constructs a [CSUserQuery] from an objc.ID.
//
// A type you use to initiate searches from your interface and offer suggested
// text completions.
func CSUserQueryFromID(id objc.ID) CSUserQuery {
	return CSUserQuery{CSSearchQuery: CSSearchQueryFromID(id)}
}

// NOTE: CSUserQuery adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSUserQuery] class.
//
// # Creating a user query
//
//   - [ICSUserQuery.InitWithUserQueryStringUserQueryContext]: Creates a new user query that searches for the specified term.
//
// # Executing the query automatically
//
//   - [ICSUserQuery.Responses]: The matching results and suggestions for the current query string.
//   - [ICSUserQuery.SetResponses]
//   - [ICSUserQuery.Suggestions]: An asynchronous sequence of suggested completions for the current query text.
//   - [ICSUserQuery.SetSuggestions]
//
// # Executing the query with handler blocks
//
//   - [ICSUserQuery.FoundSuggestionsHandler]: The block to execute when the query delivers a new batch of suggested items.
//   - [ICSUserQuery.SetFoundSuggestionsHandler]
//   - [ICSUserQuery.FoundSuggestionCount]: The number of suggested items the query found so far.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery
type ICSUserQuery interface {
	ICSSearchQuery

	// Topic: Creating a user query

	// Creates a new user query that searches for the specified term.
	InitWithUserQueryStringUserQueryContext(userQueryString string, userQueryContext ICSUserQueryContext) CSUserQuery

	// Topic: Executing the query automatically

	// The matching results and suggestions for the current query string.
	Responses() unsafe.Pointer
	SetResponses(value unsafe.Pointer)
	// An asynchronous sequence of suggested completions for the current query text.
	Suggestions() unsafe.Pointer
	SetSuggestions(value unsafe.Pointer)

	// Topic: Executing the query with handler blocks

	// The block to execute when the query delivers a new batch of suggested items.
	FoundSuggestionsHandler() CSSuggestionArrayHandler
	SetFoundSuggestionsHandler(value CSSuggestionArrayHandler)
	// The number of suggested items the query found so far.
	FoundSuggestionCount() int
}

// Init initializes the instance.
func (c CSUserQuery) Init() CSUserQuery {
	rv := objc.Send[CSUserQuery](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSUserQuery) Autorelease() CSUserQuery {
	rv := objc.Send[CSUserQuery](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSUserQuery creates a new CSUserQuery instance.
func NewCSUserQuery() CSUserQuery {
	class := getCSUserQueryClass()
	rv := objc.Send[CSUserQuery](objc.ID(class.class), objc.Sel("new"))
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
func NewCSUserQueryWithQueryStringQueryContext(queryString string, queryContext ICSSearchQueryContext) CSUserQuery {
	instance := getCSUserQueryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithQueryString:queryContext:"), objc.String(queryString), queryContext)
	return CSUserQueryFromID(rv)
}

// Creates a new user query that searches for the specified term.
//
// userQueryString: The term to search for. You may specify an empty string for this parameter.
//
// userQueryContext: A context object with options for how to run the query and generate
// results.
//
// # Return Value
//
// An initialized query object.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/init(userQueryString:userQueryContext:)
func NewCSUserQueryWithUserQueryStringUserQueryContext(userQueryString string, userQueryContext ICSUserQueryContext) CSUserQuery {
	instance := getCSUserQueryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUserQueryString:userQueryContext:"), objc.String(userQueryString), userQueryContext)
	return CSUserQueryFromID(rv)
}

// Creates a new user query that searches for the specified term.
//
// userQueryString: The term to search for. You may specify an empty string for this parameter.
//
// userQueryContext: A context object with options for how to run the query and generate
// results.
//
// # Return Value
//
// An initialized query object.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/init(userQueryString:userQueryContext:)
func (c CSUserQuery) InitWithUserQueryStringUserQueryContext(userQueryString string, userQueryContext ICSUserQueryContext) CSUserQuery {
	rv := objc.Send[CSUserQuery](c.ID, objc.Sel("initWithUserQueryString:userQueryContext:"), objc.String(userQueryString), userQueryContext)
	return rv
}

// Performs one-time tasks that prepare Spotlight to search for content in all
// search indexes.
//
// # Discussion
//
// Call this method once during your app’s lifecycle to give Spotlight time
// to load the resources it needs for search. This preparation comes at a
// cost, so measure your app’s performance and determine an appropriate time
// to call it. For example, you might call it when you load an interface that
// includes search features.
//
// You don’t need to call this method more than once during the lifetime of
// your app, but it’s safe to call the method multiple times.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepare()
func (_CSUserQueryClass CSUserQueryClass) Prepare() {
	objc.Send[objc.ID](objc.ID(_CSUserQueryClass.class), objc.Sel("prepare"))
}

// Performs one-time tasks that prepare Spotlight to search for content in one
// or more protected search indexes.
//
// protectionClasses: The file protection types associated with the indexes you plan to search.
//
// # Discussion
//
// Call this method once during your app’s lifecycle to give Spotlight time
// to load the resources it needs for search. This method prepares only the
// indexes that have the specified protetction levels. This preparation comes
// at a cost, so measure your app’s performance and determine an appropriate
// time to call it. For example, you might call it when you load an interface
// that includes search features.
//
// You don’t need to call this method more than once during the lifetime of
// your app, but it’s safe to call the method multiple times.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepareProtectionClasses(_:)
func (_CSUserQueryClass CSUserQueryClass) PrepareProtectionClasses(protectionClasses []string) {
	objc.Send[objc.ID](objc.ID(_CSUserQueryClass.class), objc.Sel("prepareProtectionClasses:"), objectivec.StringSliceToNSArray(protectionClasses))
}

// The matching results and suggestions for the current query string.
//
// See: https://developer.apple.com/documentation/corespotlight/csuserquery/responses-swift.property
func (c CSUserQuery) Responses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("responses"))
	return rv
}
func (c CSUserQuery) SetResponses(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setResponses:"), value)
}

// An asynchronous sequence of suggested completions for the current query
// text.
//
// See: https://developer.apple.com/documentation/corespotlight/csuserquery/suggestions-swift.property
func (c CSUserQuery) Suggestions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("suggestions"))
	return rv
}
func (c CSUserQuery) SetSuggestions(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setSuggestions:"), value)
}

// The block to execute when the query delivers a new batch of suggested
// items.
//
// # Discussion
//
// Specify a value for this property only if you start your query with the
// [CSUserQuery.Start] method. While the query runs, the query object executes
// the provided closure one or more times to deliver suggested completions for
// the current search term. Use your handler to retrieve the suggested
// completions and update your app’s search interface. The query object
// stops delivering suggested items when it runs out of suggestions or reaches
// the maximum number found in the [CSUserQueryContext.MaxSuggestionCount]
// property of the query configuration parameters.
//
// If you start the query by accessing the [CSUserQuery.Responses] property,
// the query object doesn’t execute the block in this property.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionsHandler
func (c CSUserQuery) FoundSuggestionsHandler() CSSuggestionArrayHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("foundSuggestionsHandler"))
	_ = rv
	return nil
}
func (c CSUserQuery) SetFoundSuggestionsHandler(value CSSuggestionArrayHandler) {
	block, cleanup := NewCSSuggestionArrayBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setFoundSuggestionsHandler:"), block)
}

// The number of suggested items the query found so far.
//
// # Discussion
//
// As the query runs, it updates the value in this property to reflect the
// total number of suggestions.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionCount
func (c CSUserQuery) FoundSuggestionCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("foundSuggestionCount"))
	return rv
}
