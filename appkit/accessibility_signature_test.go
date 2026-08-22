// Copyright 2026 The tmc/apple Authors. All rights reserved.

package appkit

import (
	"reflect"
	"strings"
	"testing"
)

// TestAccessibilitySignaturesAgreeWithProtocol compares a class's methods
// against the generated protocol wrapper's methods of the same name.
//
// TestAccessibilityConformance cannot reach most of this. A Go interface can
// only carry a protocol's @required members, so NSAccessibilityTable declares
// AccessibilityLabel and AccessibilityRows and nothing else -- correctly, since
// the SDK marks accessibilitySelectedRows, accessibilityVisibleRows and
// setAccessibilitySelectedRows: @optional (NSAccessibilityProtocols.h:147-155).
// The interface is not narrow by mistake; conformance is simply the wrong
// instrument for six of the eight known-bad members.
//
// The protocol wrapper type is the right yardstick, and it is not a hardcoded
// belief about what the types should be: NSAccessibilityTableObject is
// generated from the same protocol declaration and already renders these as
// []objectivec.IObject. Where the class disagrees with it, one of the two
// emitters read a different declaration of the same selector -- which is the
// defect, stated without this test having to name a single expected type.
func TestAccessibilitySignaturesAgreeWithProtocol(t *testing.T) {
	for _, c := range []struct {
		class, wrapper reflect.Type
		name           string
	}{
		{reflect.TypeFor[NSTableView](), reflect.TypeFor[NSAccessibilityTableObject](), "NSTableView/NSAccessibilityTable"},
		{reflect.TypeFor[NSOutlineView](), reflect.TypeFor[NSAccessibilityOutlineObject](), "NSOutlineView/NSAccessibilityOutline"},
		{reflect.TypeFor[NSView](), reflect.TypeFor[NSAccessibilityProtocolObject](), "NSView/NSAccessibilityProtocol"},
		{reflect.TypeFor[NSWindow](), reflect.TypeFor[NSAccessibilityProtocolObject](), "NSWindow/NSAccessibilityProtocol"},
	} {
		t.Run(c.name, func(t *testing.T) {
			var compared int
			for i := 0; i < c.wrapper.NumMethod(); i++ {
				wm := c.wrapper.Method(i)
				if !strings.HasPrefix(wm.Name, "Accessibility") && !strings.HasPrefix(wm.Name, "SetAccessibility") {
					// Autorelease is the reason this filter exists. The class
					// returns its own concrete type and the protocol wrapper
					// returns nothing, which is a real difference and a correct
					// one: it is memory-management boilerplate, not a rendering
					// of any accessibility declaration. Excluded by name rather
					// than by tolerance, so the exclusion is visible.
					continue
				}
				cm, ok := c.class.MethodByName(wm.Name)
				if !ok {
					continue
				}
				compared++
				// Both are method values carrying their own receiver, so the
				// receiver drops out of the comparison on both sides.
				if !sameSignature(cm.Type, wm.Type) {
					t.Errorf("%s.%s: class %s, protocol %s",
						c.class.Name(), wm.Name,
						withoutReceiver(cm.Type), withoutReceiver(wm.Type))
				}
			}
			// A comparison that compares nothing passes. The two classes share
			// dozens of accessibility members with their protocol wrappers, so
			// a low count means the lookup broke, not that the tree is clean.
			if compared < 10 {
				t.Fatalf("only %d method(s) compared; the yardstick is not being read", compared)
			}
		})
	}
}

func sameSignature(a, b reflect.Type) bool {
	if a.NumIn() != b.NumIn() || a.NumOut() != b.NumOut() {
		return false
	}
	// Skip the receiver at index 0 on both.
	for i := 1; i < a.NumIn(); i++ {
		if a.In(i) != b.In(i) {
			return false
		}
	}
	for i := 0; i < a.NumOut(); i++ {
		if a.Out(i) != b.Out(i) {
			return false
		}
	}
	return true
}

func withoutReceiver(t reflect.Type) string {
	in := make([]reflect.Type, 0, t.NumIn())
	for i := 1; i < t.NumIn(); i++ {
		in = append(in, t.In(i))
	}
	out := make([]reflect.Type, 0, t.NumOut())
	for i := 0; i < t.NumOut(); i++ {
		out = append(out, t.Out(i))
	}
	return reflect.FuncOf(in, out, t.IsVariadic()).String()
}

// TestAccessibilitySignatureCanary proves the comparison can fail, by feeding
// it a pair whose signatures are known to differ. Without it, a green run of
// the test above is indistinguishable from one whose lookup silently found
// nothing to compare.
func TestAccessibilitySignatureCanary(t *testing.T) {
	table := reflect.TypeFor[NSAccessibilityTableObject]()
	rows, ok := table.MethodByName("AccessibilityRows")
	if !ok {
		t.Fatal("canary: NSAccessibilityTableObject has no AccessibilityRows")
	}
	label, ok := table.MethodByName("AccessibilityLabel")
	if !ok {
		t.Fatal("canary: NSAccessibilityTableObject has no AccessibilityLabel")
	}
	if sameSignature(rows.Type, label.Type) {
		t.Fatal("canary: AccessibilityRows and AccessibilityLabel compared equal; " +
			"sameSignature cannot distinguish return types")
	}
}
