// Code generated from Apple documentation by applegen. DO NOT EDIT.

package objc

import (
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// objc_setAssociatedObject policy constants. Only
// OBJC_ASSOCIATION_RETAIN_NONATOMIC is used here: blocks must be
// retained for the framework to keep them alive past return, and the
// associated-object table is already serialized by the Objective-C
// runtime, so the nonatomic variant is safe and cheaper.
const (
	associationRetainNonatomic = 1
)

var (
	associationOnce sync.Once

	// objcSetAssociatedObjectFn is the package-wide pointer to
	// objc_setAssociatedObject. action.gen.go calls ensureAssociation
	// before using it so the registration happens exactly once.
	objcSetAssociatedObjectFn func(object uintptr, key unsafe.Pointer, value uintptr, policy uintptr)
)

func ensureAssociation() {
	associationOnce.Do(func() {
		ensureLibObjC()
		purego.RegisterLibFunc(&objcSetAssociatedObjectFn, libobjc, "objc_setAssociatedObject")
	})
}

// AssociateBlockWithReceiver ties the lifetime of block to receiver via
// objc_setAssociatedObject with OBJC_ASSOCIATION_RETAIN_NONATOMIC.
//
// When the receiver deallocates, the Objective-C runtime releases the
// associated block, which invokes purego's dispose callback and clears
// the Go closure from the block cache.
//
// Passing the same key on a subsequent call replaces the prior
// association: the runtime releases the previous block before retaining
// the new one. This is the mechanism used by setter-style escaping
// methods (set*Block:, set*Handler:, set*Callback:) to free the prior
// block when overwritten.
//
// The function is unexported: keys are package-level vars emitted by
// applegen at each call site, so external callers have no need to
// construct them. The block must be a Go-owned block returned by
// NewBlock; the caller transfers ownership.
func AssociateBlockWithReceiver(receiver ID, key *byte, block Block) {
	if receiver == 0 || key == nil || block == 0 {
		return
	}
	ensureAssociation()
	objcSetAssociatedObjectFn(uintptr(receiver), unsafe.Pointer(key),
		uintptr(block), associationRetainNonatomic)
	// The association retained the block; drop our local refcount so
	// the only remaining reference belongs to the receiver. When the
	// receiver deallocs (or a new value is associated to this key) the
	// runtime issues the final release, which invokes purego's dispose
	// callback and clears the Go closure cache.
	block.Release()
}
