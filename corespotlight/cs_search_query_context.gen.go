// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSSearchQueryContext] class.
var (
	_CSSearchQueryContextClass     CSSearchQueryContextClass
	_CSSearchQueryContextClassOnce sync.Once
)

func getCSSearchQueryContextClass() CSSearchQueryContextClass {
	_CSSearchQueryContextClassOnce.Do(func() {
		_CSSearchQueryContextClass = CSSearchQueryContextClass{class: objc.GetClass("CSSearchQueryContext")}
	})
	return _CSSearchQueryContextClass
}

// GetCSSearchQueryContextClass returns the class object for CSSearchQueryContext.
func GetCSSearchQueryContextClass() CSSearchQueryContextClass {
	return getCSSearchQueryContextClass()
}

type CSSearchQueryContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSearchQueryContextClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSearchQueryContextClass) Alloc() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The behavior configuration to use for a search query.
//
// # Configuring search behavior
//
//   - [CSSearchQueryContext.FetchAttributes]: The attributes the system fetches for the searchable items.
//   - [CSSearchQueryContext.SetFetchAttributes]
//   - [CSSearchQueryContext.KeyboardLanguage]: The language used for the query.
//   - [CSSearchQueryContext.SetKeyboardLanguage]
//   - [CSSearchQueryContext.SourceOptions]: The query source options to allow or deny Mail messages in the search.
//   - [CSSearchQueryContext.SetSourceOptions]
//
// # Filtering the results
//
//   - [CSSearchQueryContext.FilterQueries]: The query string used to filter the results.
//   - [CSSearchQueryContext.SetFilterQueries]
//
// # Initializers
//
//   - [CSSearchQueryContext.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext
type CSSearchQueryContext struct {
	objectivec.Object
}

// CSSearchQueryContextFromID constructs a [CSSearchQueryContext] from an objc.ID.
//
// The behavior configuration to use for a search query.
func CSSearchQueryContextFromID(id objc.ID) CSSearchQueryContext {
	return CSSearchQueryContext{objectivec.Object{ID: id}}
}

// NOTE: CSSearchQueryContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSearchQueryContext] class.
//
// # Configuring search behavior
//
//   - [ICSSearchQueryContext.FetchAttributes]: The attributes the system fetches for the searchable items.
//   - [ICSSearchQueryContext.SetFetchAttributes]
//   - [ICSSearchQueryContext.KeyboardLanguage]: The language used for the query.
//   - [ICSSearchQueryContext.SetKeyboardLanguage]
//   - [ICSSearchQueryContext.SourceOptions]: The query source options to allow or deny Mail messages in the search.
//   - [ICSSearchQueryContext.SetSourceOptions]
//
// # Filtering the results
//
//   - [ICSSearchQueryContext.FilterQueries]: The query string used to filter the results.
//   - [ICSSearchQueryContext.SetFilterQueries]
//
// # Initializers
//
//   - [ICSSearchQueryContext.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext
type ICSSearchQueryContext interface {
	objectivec.IObject

	// Topic: Configuring search behavior

	// The attributes the system fetches for the searchable items.
	FetchAttributes() []string
	SetFetchAttributes(value []string)
	// The language used for the query.
	KeyboardLanguage() string
	SetKeyboardLanguage(value string)
	// The query source options to allow or deny Mail messages in the search.
	SourceOptions() CSSearchQuerySourceOptions
	SetSourceOptions(value CSSearchQuerySourceOptions)

	// Topic: Filtering the results

	// The query string used to filter the results.
	FilterQueries() []string
	SetFilterQueries(value []string)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CSSearchQueryContext

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSSearchQueryContext) Init() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSearchQueryContext) Autorelease() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchQueryContext creates a new CSSearchQueryContext instance.
func NewCSSearchQueryContext() CSSearchQueryContext {
	class := getCSSearchQueryContextClass()
	rv := objc.Send[CSSearchQueryContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/init(coder:)
func NewCSSearchQueryContextWithCoder(coder foundation.INSCoder) CSSearchQueryContext {
	instance := getCSSearchQueryContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSSearchQueryContextFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/init(coder:)
func (c CSSearchQueryContext) InitWithCoder(coder foundation.INSCoder) CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CSSearchQueryContext) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The attributes the system fetches for the searchable items.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/fetchAttributes
func (c CSSearchQueryContext) FetchAttributes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("fetchAttributes"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchQueryContext) SetFetchAttributes(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchAttributes:"), objectivec.StringSliceToNSArray(value))
}

// The language used for the query.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/keyboardLanguage
func (c CSSearchQueryContext) KeyboardLanguage() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keyboardLanguage"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchQueryContext) SetKeyboardLanguage(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeyboardLanguage:"), objc.String(value))
}

// The query source options to allow or deny Mail messages in the search.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/sourceOptions-swift.property
func (c CSSearchQueryContext) SourceOptions() CSSearchQuerySourceOptions {
	rv := objc.Send[CSSearchQuerySourceOptions](c.ID, objc.Sel("sourceOptions"))
	return CSSearchQuerySourceOptions(rv)
}
func (c CSSearchQueryContext) SetSourceOptions(value CSSearchQuerySourceOptions) {
	objc.Send[struct{}](c.ID, objc.Sel("setSourceOptions:"), value)
}

// The query string used to filter the results.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/filterQueries
func (c CSSearchQueryContext) FilterQueries() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("filterQueries"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchQueryContext) SetFilterQueries(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setFilterQueries:"), objectivec.StringSliceToNSArray(value))
}
