// Copyright 2026 The tmc/apple Authors. All rights reserved.

package appkit

import (
	"reflect"
	"strings"
	"testing"
)

// Apple's role-based accessibility protocols are generated as Go interfaces,
// but nothing in the package proves a class satisfies one. As of 2026-08-03 ten
// of the twelve pairs below do not, for two reasons: AccessibilityLabel is
// never emitted on a class, and the protocol emitter and the class emitter
// disagree about how to map a property's type (interface says string, class
// says objectivec.IObject). Both are generator defects, so these cases are
// expected to fail until the emitters share one type-mapping path.
//
// The assertions are deliberately runtime rather than compile-time. Writing
// them as `var _ NSAccessibilityButton = NSButton{}` would make the whole
// package fail to build, which stops anyone from running any other appkit test
// while the generator is being fixed. Reflection lets the tree stay buildable
// and still report the gap.

func iface[T any]() reflect.Type { return reflect.TypeFor[T]() }

// unrelatedType is a generated wrapper with no connection to accessibility. It
// must be a real wrapper rather than struct{}: every generated protocol embeds
// objectivec.IObject, so a bare struct fails to satisfy even a protocol with no
// required members, and the guard below would pass for the wrong reason.
func unrelatedType() reflect.Type { return reflect.TypeFor[NSColor]() }

// TestProtocolsAreFalsifiable rejects protocols that cannot fail. 98 of the 192
// generated protocol interfaces in appkit and foundation have no required
// members, and a conformance result against one of those is unfalsifiable
// rather than merely weak. Guarding the suite is cheaper than auditing results
// after the fact.
func TestProtocolsAreFalsifiable(t *testing.T) {
	for _, c := range accessibilityConformances() {
		if unrelatedType().Implements(c.proto) {
			t.Errorf("%s: %s has no required members; an unrelated type "+
				"satisfies it, so this conformance result is unfalsifiable",
				c.name, c.proto)
		}
	}
}

// TestProtocolsAreFalsifiableCanary proves the guard above can fire, by feeding
// it a protocol measured to be empty. Without this, a green
// TestProtocolsAreFalsifiable would be indistinguishable from a guard that
// never evaluates its condition.
func TestProtocolsAreFalsifiableCanary(t *testing.T) {
	empty := iface[NSStandardKeyBindingResponding]()
	if !unrelatedType().Implements(empty) {
		t.Fatal("canary: NSStandardKeyBindingResponding was expected to have " +
			"no required members, but an unrelated type does not satisfy it; " +
			"the falsifiability guard cannot fire and must be re-measured")
	}
}

type conformance struct {
	class reflect.Type
	proto reflect.Type
	name  string
}

func accessibilityConformances() []conformance {
	return []conformance{
		{reflect.TypeFor[NSView](), iface[NSAccessibilityElementProtocol](), "NSView/NSAccessibilityElementProtocol"},
		{reflect.TypeFor[NSView](), iface[NSAccessibilityGroup](), "NSView/NSAccessibilityGroup"},
		{reflect.TypeFor[NSView](), iface[NSAccessibilityProtocol](), "NSView/NSAccessibilityProtocol"},
		{reflect.TypeFor[NSButton](), iface[NSAccessibilityButton](), "NSButton/NSAccessibilityButton"},
		{reflect.TypeFor[NSTextField](), iface[NSAccessibilityStaticText](), "NSTextField/NSAccessibilityStaticText"},
		{reflect.TypeFor[NSSlider](), iface[NSAccessibilitySlider](), "NSSlider/NSAccessibilitySlider"},
		{reflect.TypeFor[NSSwitch](), iface[NSAccessibilitySwitch](), "NSSwitch/NSAccessibilitySwitch"},
		{reflect.TypeFor[NSTableView](), iface[NSAccessibilityTable](), "NSTableView/NSAccessibilityTable"},
		{reflect.TypeFor[NSOutlineView](), iface[NSAccessibilityOutline](), "NSOutlineView/NSAccessibilityOutline"},
		{reflect.TypeFor[NSProgressIndicator](), iface[NSAccessibilityProgressIndicator](), "NSProgressIndicator/NSAccessibilityProgressIndicator"},
		{reflect.TypeFor[NSStepper](), iface[NSAccessibilityStepper](), "NSStepper/NSAccessibilityStepper"},
		{reflect.TypeFor[NSImageView](), iface[NSAccessibilityImage](), "NSImageView/NSAccessibilityImage"},
	}
}

// missing reports the protocol methods the class does not satisfy, and why.
func missing(class, proto reflect.Type) []string {
	var out []string
	for i := 0; i < proto.NumMethod(); i++ {
		pm := proto.Method(i)
		cm, ok := class.MethodByName(pm.Name)
		if !ok {
			out = append(out, "missing "+pm.Name+pm.Type.String()[4:])
			continue
		}
		// A method value's type carries the receiver; the interface method's
		// does not. Compare the signature without it.
		if cm.Type.NumIn()-1 != pm.Type.NumIn() || cm.Type.NumOut() != pm.Type.NumOut() {
			out = append(out, "arity "+pm.Name)
			continue
		}
		for j := 0; j < pm.Type.NumIn(); j++ {
			if cm.Type.In(j+1) != pm.Type.In(j) {
				out = append(out, "wrong arg "+pm.Name)
				break
			}
		}
		for j := 0; j < pm.Type.NumOut(); j++ {
			if cm.Type.Out(j) != pm.Type.Out(j) {
				out = append(out, "wrong result "+pm.Name+
					": have "+cm.Type.Out(j).String()+", want "+pm.Type.Out(j).String())
				break
			}
		}
	}
	return out
}

// TestAccessibilityConformanceCanary proves the check can fail. Without it a
// harness that silently reports nothing looks exactly like a tree with no
// defects -- which is how a vet run aborting on a broken dependency once
// certified this package as clean.
func TestAccessibilityConformanceCanary(t *testing.T) {
	// NSView cannot possibly satisfy a delegate protocol it does not
	// implement; if this reports conformance the harness is broken.
	impossible := iface[NSAccessibilityCustomRotorItemSearchDelegate]()
	if reflect.TypeFor[NSView]().Implements(impossible) {
		t.Fatal("canary: NSView reported as implementing " +
			"NSAccessibilityCustomRotorItemSearchDelegate; the harness cannot fail")
	}
	if len(missing(reflect.TypeFor[NSView](), impossible)) == 0 {
		t.Fatal("canary: missing() found no gap in a known-unsatisfied protocol")
	}
}

// Causes are grouped so that a regeneration can be compared against a
// prediction group by group rather than by eye. The three known groups each
// have a fix in flight; a cause landing in groupOther is the interesting
// result, because it is the one nobody predicted.
const (
	groupWrongType  = "wrong-type"      // emitters disagree about a property's type
	groupBoolNaming = "boolean-naming"  // class says IsFoo, protocol says Foo
	groupMixedIn    = "mixedin-missing" // property absent entirely (mixed-in filter)
	groupOther      = "OTHER"           // unpredicted: report this first
)

func group(cause string) string {
	switch {
	case strings.HasPrefix(cause, "wrong result"), strings.HasPrefix(cause, "wrong arg"):
		return groupWrongType
	case strings.HasPrefix(cause, "missing ") && strings.HasSuffix(cause, "() bool"):
		// class emits IsFoo, protocol declares Foo
		return groupBoolNaming
	case strings.HasPrefix(cause, "missing "):
		// getter and setter both absent: excluded by the mixed-in filter
		return groupMixedIn
	default:
		return groupOther
	}
}

func TestAccessibilityConformance(t *testing.T) {
	counts := map[string]int{}
	for _, c := range accessibilityConformances() {
		t.Run(c.name, func(t *testing.T) {
			if c.class.Implements(c.proto) {
				return
			}
			for _, m := range missing(c.class, c.proto) {
				g := group(m)
				counts[g]++
				t.Errorf("[%s] %s: %s", g, c.name, m)
			}
			if !t.Failed() {
				t.Errorf("[%s] %s: does not implement, but no per-method cause found",
					groupOther, c.name)
			}
		})
	}
	t.Cleanup(func() {
		for _, g := range []string{groupWrongType, groupBoolNaming, groupMixedIn, groupOther} {
			t.Logf("GROUP %-15s %d", g, counts[g])
		}
	})
}
