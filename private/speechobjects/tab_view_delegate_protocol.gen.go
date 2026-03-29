// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSTabViewDelegate protocol.
//
// See: https://developer.apple.com/documentation/SpeechObjects/NSTabViewDelegate
type NSTabViewDelegate interface {
	objectivec.IObject
}

// NSTabViewDelegateObject wraps an existing Objective-C object that conforms to the NSTabViewDelegate protocol.
type NSTabViewDelegateObject struct {
	objectivec.Object
}
func (o NSTabViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTabViewDelegateObjectFromID constructs a [NSTabViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTabViewDelegateObjectFromID(id objc.ID) NSTabViewDelegateObject {
	return NSTabViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/NSTabViewDelegate/tabView:didSelectTabViewItem:
func (o NSTabViewDelegateObject) TabViewDidSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("tabView:didSelectTabViewItem:"), view, item)
	}
//
// See: https://developer.apple.com/documentation/SpeechObjects/NSTabViewDelegate/tabView:shouldSelectTabViewItem:
func (o NSTabViewDelegateObject) TabViewShouldSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), view, item)
	return rv
	}
//
// See: https://developer.apple.com/documentation/SpeechObjects/NSTabViewDelegate/tabView:willSelectTabViewItem:
func (o NSTabViewDelegateObject) TabViewWillSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("tabView:willSelectTabViewItem:"), view, item)
	}
//
// See: https://developer.apple.com/documentation/SpeechObjects/NSTabViewDelegate/tabViewDidChangeNumberOfTabViewItems:
func (o NSTabViewDelegateObject) TabViewDidChangeNumberOfTabViewItems(items objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), items)
	}

