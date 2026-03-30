// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKBackForwardListItem] class.
var (
	_WKBackForwardListItemClass     WKBackForwardListItemClass
	_WKBackForwardListItemClassOnce sync.Once
)

func getWKBackForwardListItemClass() WKBackForwardListItemClass {
	_WKBackForwardListItemClassOnce.Do(func() {
		_WKBackForwardListItemClass = WKBackForwardListItemClass{class: objc.GetClass("WKBackForwardListItem")}
	})
	return _WKBackForwardListItemClass
}

// GetWKBackForwardListItemClass returns the class object for WKBackForwardListItem.
func GetWKBackForwardListItemClass() WKBackForwardListItemClass {
	return getWKBackForwardListItemClass()
}

type WKBackForwardListItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKBackForwardListItemClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKBackForwardListItemClass) Alloc() WKBackForwardListItem {
	rv := objc.Send[WKBackForwardListItem](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a webpage that the web view previously visited.
//
// # Overview
//
// Use a [WKBackForwardListItem] object to get information about previously
// visited webpages. This object identifies the page’s title and URL. It
// also identifes the URL that requested the webpage.
//
// You don’t create [WKBackForwardListItem] objects directly. Instead, a
// [WKBackForwardList] object creates them in conjunction with its associated
// web view when the web view loads new pages.
//
// # Getting the Page-Specific Information
//
//   - [WKBackForwardListItem.Title]: The title of the webpage this item represents.
//   - [WKBackForwardListItem.URL]: The URL of the webpage this item represents.
//
// # Getting the Requesting Page
//
//   - [WKBackForwardListItem.InitialURL]: The source URL that originally asked the web view to load this page.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem
type WKBackForwardListItem struct {
	objectivec.Object
}

// WKBackForwardListItemFromID constructs a [WKBackForwardListItem] from an objc.ID.
//
// A representation of a webpage that the web view previously visited.
func WKBackForwardListItemFromID(id objc.ID) WKBackForwardListItem {
	return WKBackForwardListItem{objectivec.Object{ID: id}}
}

// NOTE: WKBackForwardListItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKBackForwardListItem] class.
//
// # Getting the Page-Specific Information
//
//   - [IWKBackForwardListItem.Title]: The title of the webpage this item represents.
//   - [IWKBackForwardListItem.URL]: The URL of the webpage this item represents.
//
// # Getting the Requesting Page
//
//   - [IWKBackForwardListItem.InitialURL]: The source URL that originally asked the web view to load this page.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem
type IWKBackForwardListItem interface {
	objectivec.IObject

	// Topic: Getting the Page-Specific Information

	// The title of the webpage this item represents.
	Title() string
	// The URL of the webpage this item represents.
	URL() foundation.INSURL

	// Topic: Getting the Requesting Page

	// The source URL that originally asked the web view to load this page.
	InitialURL() foundation.INSURL
}

// Init initializes the instance.
func (b WKBackForwardListItem) Init() WKBackForwardListItem {
	rv := objc.Send[WKBackForwardListItem](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b WKBackForwardListItem) Autorelease() WKBackForwardListItem {
	rv := objc.Send[WKBackForwardListItem](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKBackForwardListItem creates a new WKBackForwardListItem instance.
func NewWKBackForwardListItem() WKBackForwardListItem {
	class := getWKBackForwardListItemClass()
	rv := objc.Send[WKBackForwardListItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The title of the webpage this item represents.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem/title
func (b WKBackForwardListItem) Title() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// The URL of the webpage this item represents.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem/url
func (b WKBackForwardListItem) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The source URL that originally asked the web view to load this page.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem/initialURL
func (b WKBackForwardListItem) InitialURL() foundation.INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("initialURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
