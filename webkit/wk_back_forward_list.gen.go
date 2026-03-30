// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKBackForwardList] class.
var (
	_WKBackForwardListClass     WKBackForwardListClass
	_WKBackForwardListClassOnce sync.Once
)

func getWKBackForwardListClass() WKBackForwardListClass {
	_WKBackForwardListClassOnce.Do(func() {
		_WKBackForwardListClass = WKBackForwardListClass{class: objc.GetClass("WKBackForwardList")}
	})
	return _WKBackForwardListClass
}

// GetWKBackForwardListClass returns the class object for WKBackForwardList.
func GetWKBackForwardListClass() WKBackForwardListClass {
	return getWKBackForwardListClass()
}

type WKBackForwardListClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKBackForwardListClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKBackForwardListClass) Alloc() WKBackForwardList {
	rv := objc.Send[WKBackForwardList](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the list of previously loaded webpages, which the
// web view uses for forward and backward navigation.
//
// # Overview
//
// Use a [WKBackForwardList] object to retrieve a web view’s previously
// loaded pages. Typically, you don’t create [WKBackForwardList] objects
// directly. Each web view creates one automatically and uses it to store the
// history of all loaded pages. Fetch this object from your web view’s
// [BackForwardList] property and use its contents to facilitate programmatic
// navigation.
//
// # Getting the Most Recent Items
//
//   - [WKBackForwardList.BackItem]: The item immediately preceding the current item, if any.
//   - [WKBackForwardList.CurrentItem]: The current item.
//   - [WKBackForwardList.ForwardItem]: The item immediately following the current item, if any.
//
// # Getting Specific Items in the List
//
//   - [WKBackForwardList.ItemAtIndex]: Returns the item at the relative offset from the current item.
//
// # Getting Sublists
//
//   - [WKBackForwardList.BackList]: The array of items that precede the current item.
//   - [WKBackForwardList.ForwardList]: The array of items that follow the current item.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList
type WKBackForwardList struct {
	objectivec.Object
}

// WKBackForwardListFromID constructs a [WKBackForwardList] from an objc.ID.
//
// An object that manages the list of previously loaded webpages, which the
// web view uses for forward and backward navigation.
func WKBackForwardListFromID(id objc.ID) WKBackForwardList {
	return WKBackForwardList{objectivec.Object{ID: id}}
}

// NOTE: WKBackForwardList adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKBackForwardList] class.
//
// # Getting the Most Recent Items
//
//   - [IWKBackForwardList.BackItem]: The item immediately preceding the current item, if any.
//   - [IWKBackForwardList.CurrentItem]: The current item.
//   - [IWKBackForwardList.ForwardItem]: The item immediately following the current item, if any.
//
// # Getting Specific Items in the List
//
//   - [IWKBackForwardList.ItemAtIndex]: Returns the item at the relative offset from the current item.
//
// # Getting Sublists
//
//   - [IWKBackForwardList.BackList]: The array of items that precede the current item.
//   - [IWKBackForwardList.ForwardList]: The array of items that follow the current item.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList
type IWKBackForwardList interface {
	objectivec.IObject

	// Topic: Getting the Most Recent Items

	// The item immediately preceding the current item, if any.
	BackItem() IWKBackForwardListItem
	// The current item.
	CurrentItem() IWKBackForwardListItem
	// The item immediately following the current item, if any.
	ForwardItem() IWKBackForwardListItem

	// Topic: Getting Specific Items in the List

	// Returns the item at the relative offset from the current item.
	ItemAtIndex(index int) IWKBackForwardListItem

	// Topic: Getting Sublists

	// The array of items that precede the current item.
	BackList() []WKBackForwardListItem
	// The array of items that follow the current item.
	ForwardList() []WKBackForwardListItem

	// The web view’s back-forward list.
	BackForwardList() IWKBackForwardList
	SetBackForwardList(value IWKBackForwardList)
}

// Init initializes the instance.
func (b WKBackForwardList) Init() WKBackForwardList {
	rv := objc.Send[WKBackForwardList](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b WKBackForwardList) Autorelease() WKBackForwardList {
	rv := objc.Send[WKBackForwardList](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKBackForwardList creates a new WKBackForwardList instance.
func NewWKBackForwardList() WKBackForwardList {
	class := getWKBackForwardListClass()
	rv := objc.Send[WKBackForwardList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the item at the relative offset from the current item.
//
// index: The offset of the desired item from the current item. Specify `0` for the
// current item, `-1` for the immediately preceding item, `1` for the
// immediately following item, and so on.
//
// # Return Value
//
// The item at the specified offset from the current item, or `nil` if the
// `index` exceeds the limits of the list.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/item(at:)
func (b WKBackForwardList) ItemAtIndex(index int) IWKBackForwardListItem {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("itemAtIndex:"), index)
	return WKBackForwardListItemFromID(rv)
}

// The item immediately preceding the current item, if any.
//
// # Discussion
//
// If the current item is the first item in the list, the value in this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/backItem
func (b WKBackForwardList) BackItem() IWKBackForwardListItem {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("backItem"))
	return WKBackForwardListItemFromID(objc.ID(rv))
}

// The current item.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/currentItem
func (b WKBackForwardList) CurrentItem() IWKBackForwardListItem {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("currentItem"))
	return WKBackForwardListItemFromID(objc.ID(rv))
}

// The item immediately following the current item, if any.
//
// # Discussion
//
// If the current item is the last item in the list, this value in this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/forwardItem
func (b WKBackForwardList) ForwardItem() IWKBackForwardListItem {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("forwardItem"))
	return WKBackForwardListItemFromID(objc.ID(rv))
}

// The array of items that precede the current item.
//
// # Discussion
//
// The items are in the order in which the web view originally visited them.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/backList
func (b WKBackForwardList) BackList() []WKBackForwardListItem {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("backList"))
	return objc.ConvertSlice(rv, func(id objc.ID) WKBackForwardListItem {
		return WKBackForwardListItemFromID(id)
	})
}

// The array of items that follow the current item.
//
// # Discussion
//
// The items are in the order in which the web view originally visited them.
//
// See: https://developer.apple.com/documentation/WebKit/WKBackForwardList/forwardList
func (b WKBackForwardList) ForwardList() []WKBackForwardListItem {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("forwardList"))
	return objc.ConvertSlice(rv, func(id objc.ID) WKBackForwardListItem {
		return WKBackForwardListItemFromID(id)
	})
}

// The web view’s back-forward list.
//
// See: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (b WKBackForwardList) BackForwardList() IWKBackForwardList {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("backForwardList"))
	return WKBackForwardListFromID(objc.ID(rv))
}
func (b WKBackForwardList) SetBackForwardList(value IWKBackForwardList) {
	objc.Send[struct{}](b.ID, objc.Sel("setBackForwardList:"), value)
}
