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
	SetAction(objectivec.SEL)
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (b NSButtonTouchBarItem) SetActionHandler(fn func()) {
	if fn == nil {
		b.SetTarget(nil)
		b.SetAction(0)
		objc.ClearActionTarget(b.ID)
		return
	}
	target, sel := objc.NewActionTarget(b.ID, func(_ objc.ID) { fn() })
	b.SetTarget(objectivec.Object{ID: target})
	b.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (c NSCell) SetActionHandler(fn func()) {
	if fn == nil {
		c.SetTarget(nil)
		c.SetAction(0)
		objc.ClearActionTarget(c.ID)
		return
	}
	target, sel := objc.NewActionTarget(c.ID, func(_ objc.ID) { fn() })
	c.SetTarget(objectivec.Object{ID: target})
	c.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (c NSColorPanel) SetActionHandler(fn func()) {
	if fn == nil {
		c.SetTarget(nil)
		c.SetAction(0)
		objc.ClearActionTarget(c.ID)
		return
	}
	target, sel := objc.NewActionTarget(c.ID, func(_ objc.ID) { fn() })
	c.SetTarget(objectivec.Object{ID: target})
	c.SetAction(objc.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (c NSColorPickerTouchBarItem) SetActionHandler(fn func()) {
	if fn == nil {
		c.SetTarget(nil)
		c.SetAction(0)
		objc.ClearActionTarget(c.ID)
		return
	}
	target, sel := objc.NewActionTarget(c.ID, func(_ objc.ID) { fn() })
	c.SetTarget(objectivec.Object{ID: target})
	c.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (c NSControl) SetActionHandler(fn func()) {
	if fn == nil {
		c.SetTarget(nil)
		c.SetAction(0)
		objc.ClearActionTarget(c.ID)
		return
	}
	target, sel := objc.NewActionTarget(c.ID, func(_ objc.ID) { fn() })
	c.SetTarget(objectivec.Object{ID: target})
	c.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (f NSFontManager) SetActionHandler(fn func()) {
	if fn == nil {
		f.SetTarget(nil)
		f.SetAction(0)
		objc.ClearActionTarget(f.ID)
		return
	}
	target, sel := objc.NewActionTarget(f.ID, func(_ objc.ID) { fn() })
	f.SetTarget(objectivec.Object{ID: target})
	f.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (g NSGestureRecognizer) SetActionHandler(fn func()) {
	if fn == nil {
		g.SetTarget(nil)
		g.SetAction(0)
		objc.ClearActionTarget(g.ID)
		return
	}
	target, sel := objc.NewActionTarget(g.ID, func(_ objc.ID) { fn() })
	g.SetTarget(objectivec.Object{ID: target})
	g.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (m NSMenuItem) SetActionHandler(fn func()) {
	if fn == nil {
		m.SetTarget(nil)
		m.SetAction(0)
		objc.ClearActionTarget(m.ID)
		return
	}
	target, sel := objc.NewActionTarget(m.ID, func(_ objc.ID) { fn() })
	m.SetTarget(objectivec.Object{ID: target})
	m.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (p NSPickerTouchBarItem) SetActionHandler(fn func()) {
	if fn == nil {
		p.SetTarget(nil)
		p.SetAction(0)
		objc.ClearActionTarget(p.ID)
		return
	}
	target, sel := objc.NewActionTarget(p.ID, func(_ objc.ID) { fn() })
	p.SetTarget(objectivec.Object{ID: target})
	p.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (s NSSliderTouchBarItem) SetActionHandler(fn func()) {
	if fn == nil {
		s.SetTarget(nil)
		s.SetAction(0)
		objc.ClearActionTarget(s.ID)
		return
	}
	target, sel := objc.NewActionTarget(s.ID, func(_ objc.ID) { fn() })
	s.SetTarget(objectivec.Object{ID: target})
	s.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (s NSStepperTouchBarItem) SetActionHandler(fn func()) {
	if fn == nil {
		s.SetTarget(nil)
		s.SetAction(0)
		objc.ClearActionTarget(s.ID)
		return
	}
	target, sel := objc.NewActionTarget(s.ID, func(_ objc.ID) { fn() })
	s.SetTarget(objectivec.Object{ID: target})
	s.SetAction(objectivec.SEL(sel))
}

// SetActionHandler sets fn as the control's action handler, replacing
// any previous handler or target/action pair (last wins, matching
// ObjC's single target/action semantics). The handler always runs on
// the main thread. Its lifetime is tied to the control; no manual
// cleanup is required. SetActionHandler(nil) clears the action.
func (t NSToolbarItem) SetActionHandler(fn func()) {
	if fn == nil {
		t.SetTarget(nil)
		t.SetAction(0)
		objc.ClearActionTarget(t.ID)
		return
	}
	target, sel := objc.NewActionTarget(t.ID, func(_ objc.ID) { fn() })
	t.SetTarget(objectivec.Object{ID: target})
	t.SetAction(objectivec.SEL(sel))
}
