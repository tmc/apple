// Code generated from Apple documentation by applegen. DO NOT EDIT.

package objc

import (
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// actionKey is the associated object key used to tie the trampoline's
// lifetime to its owner control.
var actionKey byte

var (
	actionTargetOnce  sync.Once
	actionTargetClass Class
	actionHandlers    sync.Map // ID → func(ID)

	objcSetAssociatedObject func(object uintptr, key unsafe.Pointer, value uintptr, policy uintptr)
	objcAllocateClassPair   func(superclass Class, name string, extraBytes uint) Class
)

func ensureActionTarget() {
	actionTargetOnce.Do(func() {
		ensureLibObjC()

		purego.RegisterLibFunc(&objcSetAssociatedObject, libobjc, "objc_setAssociatedObject")
		purego.RegisterLibFunc(&objcAllocateClassPair, libobjc, "objc_allocateClassPair")

		actionTargetClass = objcAllocateClassPair(GetClass("NSObject"), "GoActionTarget", 0)
		AddMethod(actionTargetClass, RegisterName("invoke:"),
			purego.NewCallback(actionTargetInvoke), "v@:@")
		AddMethod(actionTargetClass, RegisterName("dealloc"),
			purego.NewCallback(actionTargetDealloc), "v@:")
		RegisterClassPair(actionTargetClass)
	})
}

func actionTargetInvoke(self ID, _cmd SEL, sender ID) {
	if fn, ok := actionHandlers.Load(self); ok {
		fn.(func(ID))(sender)
	}
}

func actionTargetDealloc(self ID, _cmd SEL) {
	actionHandlers.Delete(self)
}

// NewActionTarget creates an Objective-C trampoline object that calls fn
// when it receives the "invoke:" selector. The trampoline is associated
// with owner via objc_setAssociatedObject so it is retained for the
// owner's lifetime and cleaned up automatically when the owner is
// deallocated or a new action target replaces it.
//
// Returns the trampoline ID and the selector to wire as the action.
func NewActionTarget(owner ID, fn func(sender ID)) (target ID, sel SEL) {
	ensureActionTarget()

	target = Send[ID](ID(actionTargetClass), Sel("alloc"))
	target = Send[ID](target, Sel("init"))

	actionHandlers.Store(target, fn)

	// OBJC_ASSOCIATION_RETAIN_NONATOMIC = 1
	objcSetAssociatedObject(uintptr(owner), unsafe.Pointer(&actionKey), uintptr(target), 1)

	// Balance the alloc — the associated object already retained it.
	Send[struct{}](target, Sel("release"))

	return target, RegisterName("invoke:")
}

