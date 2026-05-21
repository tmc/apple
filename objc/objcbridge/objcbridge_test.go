package objcbridge

import (
	"testing"

	"github.com/tmc/apple/objc"
)

func TestBlockInvokerObject(t *testing.T) {
	var got objc.ID
	block := objc.NewBlock(func(_ objc.Block, object objc.ID) {
		got = object
	})
	defer block.Release()

	const want objc.ID = 42
	if err := NewBlockInvoker().Object(objc.ID(block), want); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Object block got %#x, want %#x", uintptr(got), uintptr(want))
	}
}

func TestBlockInvokerObjectObject(t *testing.T) {
	var gotA, gotB objc.ID
	block := objc.NewBlock(func(_ objc.Block, a objc.ID, b objc.ID) {
		gotA = a
		gotB = b
	})
	defer block.Release()

	const wantA objc.ID = 42
	const wantB objc.ID = 99
	if err := NewBlockInvoker().ObjectObject(objc.ID(block), wantA, wantB); err != nil {
		t.Fatal(err)
	}
	if gotA != wantA || gotB != wantB {
		t.Fatalf("ObjectObject block got %#x/%#x, want %#x/%#x", uintptr(gotA), uintptr(gotB), uintptr(wantA), uintptr(wantB))
	}
}

func TestBlockInvokerUint(t *testing.T) {
	var got uint
	block := objc.NewBlock(func(_ objc.Block, value uint) {
		got = value
	})
	defer block.Release()

	const want uint = 42
	if err := NewBlockInvoker().Uint(objc.ID(block), want); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Uint block got %d, want %d", got, want)
	}
}

func TestBlockInvokerNilBlock(t *testing.T) {
	if err := NewBlockInvoker().Object(0, 0); err == nil {
		t.Fatal("Object nil block succeeded")
	}
}
