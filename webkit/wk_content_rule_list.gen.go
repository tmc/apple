// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKContentRuleList] class.
var (
	_WKContentRuleListClass     WKContentRuleListClass
	_WKContentRuleListClassOnce sync.Once
)

func getWKContentRuleListClass() WKContentRuleListClass {
	_WKContentRuleListClassOnce.Do(func() {
		_WKContentRuleListClass = WKContentRuleListClass{class: objc.GetClass("WKContentRuleList")}
	})
	return _WKContentRuleListClass
}

// GetWKContentRuleListClass returns the class object for WKContentRuleList.
func GetWKContentRuleListClass() WKContentRuleListClass {
	return getWKContentRuleListClass()
}

type WKContentRuleListClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKContentRuleListClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKContentRuleListClass) Alloc() WKContentRuleList {
	rv := objc.Send[WKContentRuleList](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A compiled list of rules to apply to web content.
//
// # Overview
//
// A [WKContentRuleList] object represents a compiled set of rules for
// modifying how a webpage loads content. You don’t create a
// [WKContentRuleList] directly. Instead, you specify your rules in JSON
// format and compile them using the
// [CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]
// method of [WKContentRuleListStore]. That method compiles your rules into an
// efficient byte format and returns them in an instance of this class.
//
// Content rule lists use the same syntax as content blocker extensions in
// Safari. For more information on how to specify the JSON for your rule
// lists, see [Creating a content blocker].
//
// # Getting the Rules List Identifier
//
//   - [WKContentRuleList.Identifier]: The identifier for the rule list.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleList
//
// [Creating a content blocker]: https://developer.apple.com/documentation/SafariServices/creating-a-content-blocker
type WKContentRuleList struct {
	objectivec.Object
}

// WKContentRuleListFromID constructs a [WKContentRuleList] from an objc.ID.
//
// A compiled list of rules to apply to web content.
func WKContentRuleListFromID(id objc.ID) WKContentRuleList {
	return WKContentRuleList{objectivec.Object{ID: id}}
}

// NOTE: WKContentRuleList adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKContentRuleList] class.
//
// # Getting the Rules List Identifier
//
//   - [IWKContentRuleList.Identifier]: The identifier for the rule list.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleList
type IWKContentRuleList interface {
	objectivec.IObject

	// Topic: Getting the Rules List Identifier

	// The identifier for the rule list.
	Identifier() string
}

// Init initializes the instance.
func (c WKContentRuleList) Init() WKContentRuleList {
	rv := objc.Send[WKContentRuleList](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c WKContentRuleList) Autorelease() WKContentRuleList {
	rv := objc.Send[WKContentRuleList](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKContentRuleList creates a new WKContentRuleList instance.
func NewWKContentRuleList() WKContentRuleList {
	class := getWKContentRuleListClass()
	rv := objc.Send[WKContentRuleList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The identifier for the rule list.
//
// # Discussion
//
// You specify the identifier for your rule lists at compile time in the
// [CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]
// method of [WKContentRuleListStore]. You also use this identifier to look up
// the rules list later.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleList/identifier
func (c WKContentRuleList) Identifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
