//go:build darwin

package objc_test

import (
	"testing"

	"github.com/tmc/apple/objc"
)

// TestSendUintptrKindedReturns guards the Send fast path against a regression
// where uintptr-kinded return types other than ID panicked with
// "interface conversion: interface {} is objc.ID, not <T>". The fast path boxed
// the result as ID and used a type assertion, which fails for distinct named
// types (uintptr, SEL, Class). The fix reinterprets the raw uintptr bits, so
// every uintptr-kinded T must return without panicking.
func TestSendUintptrKindedReturns(t *testing.T) {
	cls := objc.GetClass("NSObject")
	if cls == 0 {
		t.Skip("NSObject class unavailable")
	}
	obj := objc.Send[objc.ID](objc.ID(cls), objc.Sel("alloc"))
	obj = objc.Send[objc.ID](obj, objc.Sel("init"))
	if obj == 0 {
		t.Fatal("alloc/init returned nil")
	}
	defer objc.Send[objc.ID](obj, objc.Sel("release"))

	// hash returns an NSUInteger, exercised here as each uintptr-kinded T. Before
	// the fix, Send[uintptr], Send[SEL], and Send[Class] all panicked; only
	// Send[ID] (the boxed type) and Send[struct{}] (void) worked.
	t.Run("uintptr", func(t *testing.T) {
		_ = objc.Send[uintptr](obj, objc.Sel("hash"))
	})
	t.Run("ID", func(t *testing.T) {
		_ = objc.Send[objc.ID](obj, objc.Sel("hash"))
	})
	t.Run("SEL", func(t *testing.T) {
		_ = objc.Send[objc.SEL](obj, objc.Sel("hash"))
	})
	t.Run("Class", func(t *testing.T) {
		_ = objc.Send[objc.Class](obj, objc.Sel("hash"))
	})
}
