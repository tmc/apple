// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The [NSTabViewDelegate] protocol defines the optional methods implemented by delegates of [NSTabView] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate
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

// Informs the delegate that the number of tab view items in `tabView` has
// changed.
//
// tabView: The tab view that added or removed tabview items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate/tabViewDidChangeNumberOfTabViewItems(_:)
func (o NSTabViewDelegateObject) TabViewDidChangeNumberOfTabViewItems(tabView INSTabView) {
	objc.Send[struct{}](o.ID, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), tabView)
}

// Invoked just before `tabViewItem` in `tabView` is selected.
//
// tabView: The tab view that sent the request.
//
// tabViewItem: The tab view item to select.
//
// # Return Value
//
// true if the tab view item should be selected, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate/tabView(_:shouldSelect:)
func (o NSTabViewDelegateObject) TabViewShouldSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), tabView, tabViewItem)
	return rv
}

// Informs the delegate that `tabView` is about to select `tabViewItem`.
//
// tabView: The tab view that sent the request.
//
// tabViewItem: The tab view item that is about to be selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate/tabView(_:willSelect:)
func (o NSTabViewDelegateObject) TabViewWillSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) {
	objc.Send[struct{}](o.ID, objc.Sel("tabView:willSelectTabViewItem:"), tabView, tabViewItem)
}

// Informs the delegate that `tabView` has selected `tabViewItem`.
//
// tabView: The tab view that sent the request.
//
// tabViewItem: The tab view item that was selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate/tabView(_:didSelect:)
func (o NSTabViewDelegateObject) TabViewDidSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) {
	objc.Send[struct{}](o.ID, objc.Sel("tabView:didSelectTabViewItem:"), tabView, tabViewItem)
}

// NSTabViewDelegateConfig holds optional typed callbacks for [NSTabViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstabviewdelegate
type NSTabViewDelegateConfig struct {

	// Adding and Removing Tabs
	// DidChangeNumberOfTabViewItems — Informs the delegate that the number of tab view items in `tabView` has changed.
	DidChangeNumberOfTabViewItems func(tabView NSTabView)

	// Other Methods
	// ShouldSelectTabViewItem — Invoked just before `tabViewItem` in `tabView` is selected.
	ShouldSelectTabViewItem func(tabView NSTabView, tabViewItem NSTabViewItem) bool
	// WillSelectTabViewItem — Informs the delegate that `tabView` is about to select `tabViewItem`.
	WillSelectTabViewItem func(tabView NSTabView, tabViewItem NSTabViewItem)
	// DidSelectTabViewItem — Informs the delegate that `tabView` has selected `tabViewItem`.
	DidSelectTabViewItem func(tabView NSTabView, tabViewItem NSTabViewItem)
}

// NewNSTabViewDelegate creates an Objective-C object implementing the [NSTabViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTabViewDelegateObject] satisfies the [NSTabViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstabviewdelegate
func NewNSTabViewDelegate(config NSTabViewDelegateConfig) NSTabViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTabViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DidChangeNumberOfTabViewItems != nil {
		fn := config.DidChangeNumberOfTabViewItems
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tabViewDidChangeNumberOfTabViewItems:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tabViewID objc.ID) {
				tabView := NSTabViewFromID(tabViewID)
				fn(tabView)
			},
		})
	}

	if config.ShouldSelectTabViewItem != nil {
		fn := config.ShouldSelectTabViewItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tabView:shouldSelectTabViewItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tabViewID objc.ID, tabViewItemID objc.ID) bool {
				tabView := NSTabViewFromID(tabViewID)
				tabViewItem := NSTabViewItemFromID(tabViewItemID)
				return fn(tabView, tabViewItem)
			},
		})
	}

	if config.WillSelectTabViewItem != nil {
		fn := config.WillSelectTabViewItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tabView:willSelectTabViewItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tabViewID objc.ID, tabViewItemID objc.ID) {
				tabView := NSTabViewFromID(tabViewID)
				tabViewItem := NSTabViewItemFromID(tabViewItemID)
				fn(tabView, tabViewItem)
			},
		})
	}

	if config.DidSelectTabViewItem != nil {
		fn := config.DidSelectTabViewItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tabView:didSelectTabViewItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tabViewID objc.ID, tabViewItemID objc.ID) {
				tabView := NSTabViewFromID(tabViewID)
				tabViewItem := NSTabViewItemFromID(tabViewItemID)
				fn(tabView, tabViewItem)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTabViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTabViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTabViewDelegateObjectFromID(instance)
}
