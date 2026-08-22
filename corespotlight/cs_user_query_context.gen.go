// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CSUserQueryContext] class.
var (
	_CSUserQueryContextClass     CSUserQueryContextClass
	_CSUserQueryContextClassOnce sync.Once
)

func getCSUserQueryContextClass() CSUserQueryContextClass {
	_CSUserQueryContextClassOnce.Do(func() {
		_CSUserQueryContextClass = CSUserQueryContextClass{class: objc.GetClass("CSUserQueryContext")}
	})
	return _CSUserQueryContextClass
}

// GetCSUserQueryContextClass returns the class object for CSUserQueryContext.
func GetCSUserQueryContextClass() CSUserQueryContextClass {
	return getCSUserQueryContextClass()
}

type CSUserQueryContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSUserQueryContextClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSUserQueryContextClass) Alloc() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The configuration details to apply to a user query.
//
// # Overview
//
// Use an instance of [CSUserQueryContext] to configure the search parameters
// for a [CSUserQuery] object. This object stores configuration details that
// the query uses to modify the search results it delivers. For example, use
// this object to specify the maximum number of results or suggestions you
// want the query to return. You can also use it to enable or disable the
// ranking of results by Spotlight.
//
// For information about search filters and other configurable query
// parameters, see the parent class [CSSearchQueryContext].
//
// # Configuring search options
//
//   - [CSUserQueryContext.MaxResultCount]: The maximum number of search results for the query to return.
//   - [CSUserQueryContext.SetMaxResultCount]
//   - [CSUserQueryContext.MaxSuggestionCount]: The maximum number of suggested text completions for the query to return.
//   - [CSUserQueryContext.SetMaxSuggestionCount]
//   - [CSUserQueryContext.DisableSemanticSearch]: A Boolean value that indicates whether to exclude semantic-based search results from the output.
//   - [CSUserQueryContext.SetDisableSemanticSearch]
//
// # Configuring the ranked results behavior
//
//   - [CSUserQueryContext.EnableRankedResults]: A Boolean value that indicates whether the query sorts results by their relevance.
//   - [CSUserQueryContext.SetEnableRankedResults]
//   - [CSUserQueryContext.MaxRankedResultCount]: The maximum number of ranked results to return during the query.
//   - [CSUserQueryContext.SetMaxRankedResultCount]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext
type CSUserQueryContext struct {
	CSSearchQueryContext
}

// CSUserQueryContextFromID constructs a [CSUserQueryContext] from an objc.ID.
//
// The configuration details to apply to a user query.
func CSUserQueryContextFromID(id objc.ID) CSUserQueryContext {
	return CSUserQueryContext{CSSearchQueryContext: CSSearchQueryContextFromID(id)}
}

// NOTE: CSUserQueryContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSUserQueryContext] class.
//
// # Configuring search options
//
//   - [ICSUserQueryContext.MaxResultCount]: The maximum number of search results for the query to return.
//   - [ICSUserQueryContext.SetMaxResultCount]
//   - [ICSUserQueryContext.MaxSuggestionCount]: The maximum number of suggested text completions for the query to return.
//   - [ICSUserQueryContext.SetMaxSuggestionCount]
//   - [ICSUserQueryContext.DisableSemanticSearch]: A Boolean value that indicates whether to exclude semantic-based search results from the output.
//   - [ICSUserQueryContext.SetDisableSemanticSearch]
//
// # Configuring the ranked results behavior
//
//   - [ICSUserQueryContext.EnableRankedResults]: A Boolean value that indicates whether the query sorts results by their relevance.
//   - [ICSUserQueryContext.SetEnableRankedResults]
//   - [ICSUserQueryContext.MaxRankedResultCount]: The maximum number of ranked results to return during the query.
//   - [ICSUserQueryContext.SetMaxRankedResultCount]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext
type ICSUserQueryContext interface {
	ICSSearchQueryContext

	// Topic: Configuring search options

	// The maximum number of search results for the query to return.
	MaxResultCount() int
	SetMaxResultCount(value int)
	// The maximum number of suggested text completions for the query to return.
	MaxSuggestionCount() int
	SetMaxSuggestionCount(value int)
	// A Boolean value that indicates whether to exclude semantic-based search results from the output.
	DisableSemanticSearch() bool
	SetDisableSemanticSearch(value bool)

	// Topic: Configuring the ranked results behavior

	// A Boolean value that indicates whether the query sorts results by their relevance.
	EnableRankedResults() bool
	SetEnableRankedResults(value bool)
	// The maximum number of ranked results to return during the query.
	MaxRankedResultCount() int
	SetMaxRankedResultCount(value int)
}

// Init initializes the instance.
func (c CSUserQueryContext) Init() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSUserQueryContext) Autorelease() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSUserQueryContext creates a new CSUserQueryContext instance.
func NewCSUserQueryContext() CSUserQueryContext {
	class := getCSUserQueryContextClass()
	rv := objc.Send[CSUserQueryContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/init(coder:)
func NewCSUserQueryContextWithCoder(coder foundation.INSCoder) CSUserQueryContext {
	instance := getCSUserQueryContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSUserQueryContextFromID(rv)
}

// Creates a new query context object with an optional suggested search
// string.
//
// currentSuggestion: The suggested text completion that the person selected in your interface.
// Specify `nil` if the person hasn’t chosen a suggestion.
//
// # Return Value
//
// An initialized user query context object. Configure the properties of the
// returned object and use it to construct a [CSUserQuery] object.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/init(currentSuggestion:)
func NewCSUserQueryContextWithCurrentSuggestion(currentSuggestion ICSSuggestion) CSUserQueryContext {
	rv := objc.Send[objc.ID](objc.ID(getCSUserQueryContextClass().class), objc.Sel("userQueryContextWithCurrentSuggestion:"), currentSuggestion)
	return CSUserQueryContextFromID(rv)
}

// The maximum number of search results for the query to return.
//
// # Discussion
//
// Spotlight returns all results that it finds by default. Set a maximum limit
// to terminate the search early, which can improve performance.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxResultCount
func (c CSUserQueryContext) MaxResultCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maxResultCount"))
	return rv
}
func (c CSUserQueryContext) SetMaxResultCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxResultCount:"), value)
}

// The maximum number of suggested text completions for the query to return.
//
// # Discussion
//
// You might specify different limits to account for the amount of available
// space.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxSuggestionCount
func (c CSUserQueryContext) MaxSuggestionCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maxSuggestionCount"))
	return rv
}
func (c CSUserQueryContext) SetMaxSuggestionCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxSuggestionCount:"), value)
}

// A Boolean value that indicates whether to exclude semantic-based search
// results from the output.
//
// # Discussion
//
// Semantic searching finds matches that are related to the original term, but
// not necessarily a lexical match. For example, a search for the string
// “Sun and Moon” might also return a result with a title like “Sol and
// Luna”. The default value of this property is `true`, which enables the
// delivery of semantic search results.
//
// You might set this property to `false` when you want to perform only a
// lexical match. For example, you might disable semantic search when looking
// for a proper name.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/disableSemanticSearch
func (c CSUserQueryContext) DisableSemanticSearch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("disableSemanticSearch"))
	return rv
}
func (c CSUserQueryContext) SetDisableSemanticSearch(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDisableSemanticSearch:"), value)
}

// A Boolean value that indicates whether the query sorts results by their
// relevance.
//
// # Discussion
//
// The default value of this property is `true`. Setting the property to
// `false` causes Spotlight to return results as it finds them instead of
// ordering them by their relevance or saliency to the search term.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/enableRankedResults
func (c CSUserQueryContext) EnableRankedResults() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("enableRankedResults"))
	return rv
}
func (c CSUserQueryContext) SetEnableRankedResults(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnableRankedResults:"), value)
}

// The maximum number of ranked results to return during the query.
//
// # Discussion
//
// Spotlight ranks a limited number of results by default, but you can change
// the default value to improve performance or better suit your app’s
// interface. For example, you might want to return only the five most
// relevant results due to space constraints in your UI.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxRankedResultCount
func (c CSUserQueryContext) MaxRankedResultCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maxRankedResultCount"))
	return rv
}
func (c CSUserQueryContext) SetMaxRankedResultCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxRankedResultCount:"), value)
}
