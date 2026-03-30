// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWritingToolsCoordinatorContext] class.
var (
	_NSWritingToolsCoordinatorContextClass     NSWritingToolsCoordinatorContextClass
	_NSWritingToolsCoordinatorContextClassOnce sync.Once
)

func getNSWritingToolsCoordinatorContextClass() NSWritingToolsCoordinatorContextClass {
	_NSWritingToolsCoordinatorContextClassOnce.Do(func() {
		_NSWritingToolsCoordinatorContextClass = NSWritingToolsCoordinatorContextClass{class: objc.GetClass("NSWritingToolsCoordinatorContext")}
	})
	return _NSWritingToolsCoordinatorContextClass
}

// GetNSWritingToolsCoordinatorContextClass returns the class object for NSWritingToolsCoordinatorContext.
func GetNSWritingToolsCoordinatorContextClass() NSWritingToolsCoordinatorContextClass {
	return getNSWritingToolsCoordinatorContextClass()
}

type NSWritingToolsCoordinatorContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWritingToolsCoordinatorContextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWritingToolsCoordinatorContextClass) Alloc() NSWritingToolsCoordinatorContext {
	rv := objc.Send[NSWritingToolsCoordinatorContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A data object that you use to share your custom view’s text with Writing
// Tools.
//
// # Overview
//
// At the start of every Writing Tools operation, you create one or more
// `NSWritingToolsCoordinator.Context` objects with a copy of the text you
// want Writing Tools to evaluate. Each Writing Tools operation starts with a
// call to the [WritingToolsCoordinatorRequestsContextsForScopeCompletion]
// method of your [NSWritingToolsCoordinatorDelegate] object. Use the
// parameters of that method to determine how much of your view’s text to
// provide. For some operations, Writing Tools asks for all of your view’s
// text, but in others it asks for only a portion of the text. When Writing
// Tools finishes its evaluation, it reports changes back to your delegate
// relative to the context objects you provided.
//
// When Writing Tools asks for your view’s text, create one or more
// `NSWritingToolsCoordinator.Context` objects with the requested content. If
// your view contains only one text storage object, create only one context
// object for the request. However, if you use multiple text storage objects
// to manage different parts of your view’s content, you might need to
// create multiple context objects. The actual number depends on how much of
// your text Writing Tools asks for. For example, when Writing Tools asks for
// all of your view’s content, you return one context object for each text
// storage object in your view. However, if Writing Tools asks for the current
// selection, and one text storage object contains all of the selected text,
// you create only one context object for the content.
//
// Writing Tools uses your context objects as the starting point for its
// evaluations, and as a reference point for any changes. Because Writing
// Tools doesn’t know anything about your view or its content, it makes
// suggestions only relative to your context objects. It’s your
// responsibility to take those suggestions and incorporate them back into
// your view’s text storage. In some cases, you might need to store
// additional information to update your storage correctly. For example, you
// might need to store, and update as needed, the offset from the start of
// your document to the start of the text in your context object.
//
// When Writing Tools asks for the currently selected text in your view,
// include some of the surrounding text in your context object as well. Supply
// a string that includes the selection and any text up to the nearest
// paragraph boundary. When creating your context object, specify a range
// value that represents the portion of that string that corresponds to the
// text selection. Providing some additional text in your context object can
// help Writing Tools improve its evaluation of your content. Writing Tools
// uses the [NSWritingToolsCoordinatorContext.ResolvedRange] property of your context object to indicate what
// text it considered.
//
// If your context object includes text that you don’t want Writing Tools to
// evaluate, add the `excludeFromWritingTools` attribute to the corresponding
// characters of your [NSAttributedString] object. You might add this
// attribute if the text string includes a code listing or readonly content
// that you don’t want Writing Tools to change.
//
// # Creating a context object
//
//   - [NSWritingToolsCoordinatorContext.InitWithAttributedStringRange]: Creates a context object with the specified attributed string and range information.
//
// # Getting the source text details
//
//   - [NSWritingToolsCoordinatorContext.AttributedString]: The portion of your view’s text to evaluate.
//   - [NSWritingToolsCoordinatorContext.Range]: The unique identifier of the context object.
//
// # Getting the assessed text range
//
//   - [NSWritingToolsCoordinatorContext.ResolvedRange]: The actual range of text that Writing Tools might change, which can be different than the range of text you supplied.
//
// # Identifying the context object
//
//   - [NSWritingToolsCoordinatorContext.Identifier]: The unique identifier of the context object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
type NSWritingToolsCoordinatorContext struct {
	objectivec.Object
}

// NSWritingToolsCoordinatorContextFromID constructs a [NSWritingToolsCoordinatorContext] from an objc.ID.
//
// A data object that you use to share your custom view’s text with Writing
// Tools.
func NSWritingToolsCoordinatorContextFromID(id objc.ID) NSWritingToolsCoordinatorContext {
	return NSWritingToolsCoordinatorContext{objectivec.Object{ID: id}}
}

// NOTE: NSWritingToolsCoordinatorContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWritingToolsCoordinatorContext] class.
//
// # Creating a context object
//
//   - [INSWritingToolsCoordinatorContext.InitWithAttributedStringRange]: Creates a context object with the specified attributed string and range information.
//
// # Getting the source text details
//
//   - [INSWritingToolsCoordinatorContext.AttributedString]: The portion of your view’s text to evaluate.
//   - [INSWritingToolsCoordinatorContext.Range]: The unique identifier of the context object.
//
// # Getting the assessed text range
//
//   - [INSWritingToolsCoordinatorContext.ResolvedRange]: The actual range of text that Writing Tools might change, which can be different than the range of text you supplied.
//
// # Identifying the context object
//
//   - [INSWritingToolsCoordinatorContext.Identifier]: The unique identifier of the context object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context
type INSWritingToolsCoordinatorContext interface {
	objectivec.IObject

	// Topic: Creating a context object

	// Creates a context object with the specified attributed string and range information.
	InitWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSWritingToolsCoordinatorContext

	// Topic: Getting the source text details

	// The portion of your view’s text to evaluate.
	AttributedString() foundation.NSAttributedString
	// The unique identifier of the context object.
	Range() foundation.NSRange

	// Topic: Getting the assessed text range

	// The actual range of text that Writing Tools might change, which can be different than the range of text you supplied.
	ResolvedRange() foundation.NSRange

	// Topic: Identifying the context object

	// The unique identifier of the context object.
	Identifier() foundation.NSUUID
}

// Init initializes the instance.
func (w NSWritingToolsCoordinatorContext) Init() NSWritingToolsCoordinatorContext {
	rv := objc.Send[NSWritingToolsCoordinatorContext](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWritingToolsCoordinatorContext) Autorelease() NSWritingToolsCoordinatorContext {
	rv := objc.Send[NSWritingToolsCoordinatorContext](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWritingToolsCoordinatorContext creates a new NSWritingToolsCoordinatorContext instance.
func NewNSWritingToolsCoordinatorContext() NSWritingToolsCoordinatorContext {
	class := getNSWritingToolsCoordinatorContextClass()
	rv := objc.Send[NSWritingToolsCoordinatorContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a context object with the specified attributed string and range
// information.
//
// attributedString: A string that contains some or all of the content from your view’s text
// storage. This initializer makes a copy of the string you provide, so you
// can discard the original when you’re done.
//
// range: The portion of `attributedString` you want Writing Tools to evaluate. If
// you want Writing Tools to evaluate the entire string you provided, specify
// a range with a location of `0` and a length equal to your string’s
// length. If you want Writing Tools to evaluate only part of the string,
// provide the appropriate range in this parameter. Writing Tools suggests
// changes only to the range of text you specify, but it can consider text
// outside that range during the evaluation process.
//
// # Discussion
//
// When Writing Tools asks for your view’s current selection, it’s best to
// create a string that includes text before and after that selection. During
// the evaluation process, Writing Tools can use the additional text you
// provided to improve the results it delivers. If you do provide additional
// text, set the `range` parameter to the portion of `attributedString` with
// the current selection. Don’t use the `range` parameter to specify the
// location of the text in your view’s text storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/init(attributedString:range:)
func NewWritingToolsCoordinatorContextWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSWritingToolsCoordinatorContext {
	instance := getNSWritingToolsCoordinatorContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:range:"), attributedString, range_)
	return NSWritingToolsCoordinatorContextFromID(rv)
}

// Creates a context object with the specified attributed string and range
// information.
//
// attributedString: A string that contains some or all of the content from your view’s text
// storage. This initializer makes a copy of the string you provide, so you
// can discard the original when you’re done.
//
// range: The portion of `attributedString` you want Writing Tools to evaluate. If
// you want Writing Tools to evaluate the entire string you provided, specify
// a range with a location of `0` and a length equal to your string’s
// length. If you want Writing Tools to evaluate only part of the string,
// provide the appropriate range in this parameter. Writing Tools suggests
// changes only to the range of text you specify, but it can consider text
// outside that range during the evaluation process.
//
// # Discussion
//
// When Writing Tools asks for your view’s current selection, it’s best to
// create a string that includes text before and after that selection. During
// the evaluation process, Writing Tools can use the additional text you
// provided to improve the results it delivers. If you do provide additional
// text, set the `range` parameter to the portion of `attributedString` with
// the current selection. Don’t use the `range` parameter to specify the
// location of the text in your view’s text storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/init(attributedString:range:)
func (w NSWritingToolsCoordinatorContext) InitWithAttributedStringRange(attributedString foundation.NSAttributedString, range_ foundation.NSRange) NSWritingToolsCoordinatorContext {
	rv := objc.Send[NSWritingToolsCoordinatorContext](w.ID, objc.Sel("initWithAttributedString:range:"), attributedString, range_)
	return rv
}

// The portion of your view’s text to evaluate.
//
// # Discussion
//
// The [NSWritingToolsCoordinatorContext] object initializes the value of this
// property at creation time and doesn’t change it during the course of an
// operation. Instead, it suggests changes to the text in the indicated range
// and reports those changes to your [NSWritingToolsCoordinatorDelegate]
// object. Use the methods of your delegate object to integrate those changes
// back into your view’s text storage.
//
// It’s your responsibility to track the location of this text in your
// view’s text storage object. When Writing Tools reports changes, it
// provides range values relative to this string. If you initialize this
// property with a subset of your view’s content, you must adjust any ranges
// that Writing Tools provides to get the correct location in your text
// storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/attributedString
func (w NSWritingToolsCoordinatorContext) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

// The unique identifier of the context object.
//
// # Discussion
//
// The [NSWritingToolsCoordinatorContext] object initializes the value of this
// property at creation time. Use this value to identify the context object
// within your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/range
func (w NSWritingToolsCoordinatorContext) Range() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](w.ID, objc.Sel("range"))
	return foundation.NSRange(rv)
}

// The actual range of text that Writing Tools might change, which can be
// different than the range of text you supplied.
//
// # Discussion
//
// After analyzing the text in your context object, Writing Tools sets this
// property to the portion of [AttributedString] it might modify. Initially,
// this property has a location of [NSNotFound] and a length of `0`, but
// Writing Tools updates those values before making any changes to the text.
//
// While the Writing Tools operation is active, make sure Writing Tools has
// exclusive access to the text in this range. Your
// [NSWritingToolsCoordinatorDelegate] object can make changes to the text as
// part of incorporating Writing Tools results, but don’t allow changes to
// come from other sources. For example, don’t let someone edit the text in
// this range directly until Writing Tools finishes.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/resolvedRange
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
func (w NSWritingToolsCoordinatorContext) ResolvedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](w.ID, objc.Sel("resolvedRange"))
	return foundation.NSRange(rv)
}

// The unique identifier of the context object.
//
// # Discussion
//
// The [NSWritingToolsCoordinatorContext] object initializes the value of this
// property at creation time. Use this value to identify the context object
// within your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/identifier
func (w NSWritingToolsCoordinatorContext) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
