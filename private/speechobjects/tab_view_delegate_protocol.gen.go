// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSTabViewDelegate protocol.
type NSTabViewDelegate interface {
	objectivec.IObject

	// TabViewDidSelectTabViewItem protocol.
	TabViewDidSelectTabViewItem(view objectivec.IObject, item objectivec.IObject)

	// TabViewShouldSelectTabViewItem protocol.
	TabViewShouldSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) bool

	// TabViewWillSelectTabViewItem protocol.
	TabViewWillSelectTabViewItem(view objectivec.IObject, item objectivec.IObject)

	// TabViewDidChangeNumberOfTabViewItems protocol.
	TabViewDidChangeNumberOfTabViewItems(items objectivec.IObject)
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

func (o NSTabViewDelegateObject) TabViewDidSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("tabView:didSelectTabViewItem:"), view, item)
}
func (o NSTabViewDelegateObject) TabViewShouldSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), view, item)
	return rv
}
func (o NSTabViewDelegateObject) TabViewWillSelectTabViewItem(view objectivec.IObject, item objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("tabView:willSelectTabViewItem:"), view, item)
}
func (o NSTabViewDelegateObject) TabViewDidChangeNumberOfTabViewItems(items objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), items)
}
