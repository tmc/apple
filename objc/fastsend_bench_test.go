//go:build darwin

package objc_test

import (
	"testing"

	"github.com/tmc/apple/objc"
)

// BenchmarkSendFastPath benchmarks the fast path (all args uintptr-castable, return type ID).
func BenchmarkSendFastPath(b *testing.B) {
	cls := objc.GetClass("NSObject")
	if cls == 0 {
		b.Fatal("NSObject unavailable")
	}
	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("alloc"))
	id = objc.Send[objc.ID](id, objc.Sel("init"))
	sel := objc.Sel("hash")

	b.ResetTimer()
	for b.Loop() {
		objc.Send[objc.ID](id, sel)
	}
}

// BenchmarkSendFastPathWithArgs benchmarks the fast path with a uintptr-castable arg.
func BenchmarkSendFastPathWithArgs(b *testing.B) {
	cls := objc.GetClass("NSObject")
	if cls == 0 {
		b.Fatal("NSObject unavailable")
	}
	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("alloc"))
	id = objc.Send[objc.ID](id, objc.Sel("init"))
	sel := objc.Sel("isEqual:")

	b.ResetTimer()
	for b.Loop() {
		objc.Send[objc.ID](id, sel, objc.ID(id))
	}
}

// BenchmarkSendVoidReturn benchmarks the fast path for void (struct{}) return.
func BenchmarkSendVoidReturn(b *testing.B) {
	cls := objc.GetClass("NSObject")
	if cls == 0 {
		b.Fatal("NSObject unavailable")
	}
	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("alloc"))
	id = objc.Send[objc.ID](id, objc.Sel("init"))
	sel := objc.Sel("self")

	b.ResetTimer()
	for b.Loop() {
		objc.Send[struct{}](id, sel)
	}
}
