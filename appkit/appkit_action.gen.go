// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Actionable is the interface for controls that support a target/action pair.
type Actionable interface {
	objectivec.IObject
	SetTarget(objectivec.IObject)
	SetAction(objc.SEL)
}

// SetActionHandler sets a Go function as the action handler for the control.
// The handler replaces any previous handler (last wins), matching ObjC's
// single target/action semantics. The handler's lifetime is tied to the
// control via associated objects — no manual cleanup is required.
func (n NSControl) SetActionHandler(fn func()) {
	target, sel := objc.NewActionTarget(n.ID, func(_ objc.ID) { fn() })
	n.SetTarget(objectivec.Object{ID: target})
	n.SetAction(sel)
}

// SetActionHandler sets a Go function as the action handler for the menu item.
func (n NSMenuItem) SetActionHandler(fn func()) {
	target, sel := objc.NewActionTarget(n.ID, func(_ objc.ID) { fn() })
	n.SetTarget(objectivec.Object{ID: target})
	n.SetAction(sel)
}

// SetActionHandler sets a Go function as the action handler for the toolbar item.
func (n NSToolbarItem) SetActionHandler(fn func()) {
	target, sel := objc.NewActionTarget(n.ID, func(_ objc.ID) { fn() })
	n.SetTarget(objectivec.Object{ID: target})
	n.SetAction(sel)
}

// SetActionHandlerFor sets a Go function as the action handler for any
// Actionable control. The handler receives the sender typed as T.
func SetActionHandlerFor[T Actionable](ctrl T, fn func(T)) {
	id := ctrl.GetID()
	target, sel := objc.NewActionTarget(id, func(_ objc.ID) {
		fn(ctrl)
	})
	ctrl.SetTarget(objectivec.Object{ID: target})
	ctrl.SetAction(sel)
}


