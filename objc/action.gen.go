// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

package objc

import (
	"strconv"
	"sync"
	"unsafe"

	pobjc "github.com/ebitengine/purego/objc"
)

// actionKey is the associated object key used to tie the trampoline's
// lifetime to its owner control.
var actionKey byte

var (
	actionTargetOnce  sync.Once
	actionTargetClass Class
	actionInvokeSel   SEL
	actionDeallocSel  SEL
	actionHandlers    sync.Map // ID → func(ID)
)

func ensureActionTarget() {
	actionTargetOnce.Do(func() {
		ensureAssociation()
		actionInvokeSel = Sel("invoke:")
		actionDeallocSel = Sel("dealloc")
		actionTargetClass = registerActionTargetClass("GoActionTarget")
	})
}

func registerActionTargetClass(baseName string) Class {
	methods := []pobjc.MethodDef{
		{Cmd: actionInvokeSel, Fn: actionTargetInvoke},
		{Cmd: actionDeallocSel, Fn: actionTargetDealloc},
	}
	for suffix := 1; ; suffix++ {
		name := baseName
		if suffix > 1 {
			name += strconv.Itoa(suffix)
		}
		class, err := pobjc.RegisterClass(name, pobjc.GetClass("NSObject"), nil, nil, methods)
		if err == nil {
			return class
		}
		if pobjc.GetClass(name) == 0 {
			panic(err)
		}
	}
}

func actionTargetInvoke(self ID, _cmd SEL, sender ID) {
	if fn, ok := actionHandlers.Load(self); ok {
		fn.(func(ID))(sender)
	}
}

func actionTargetDealloc(self ID, _cmd SEL) {
	actionHandlers.Delete(self)
	SendSuper[struct{}](self, actionDeallocSel)
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

	// OBJC_ASSOCIATION_RETAIN_NONATOMIC = associationRetainNonatomic
	objcSetAssociatedObjectFn(uintptr(owner), unsafe.Pointer(&actionKey), uintptr(target), associationRetainNonatomic)

	// Balance the alloc — the associated object already retained it.
	Send[struct{}](target, Sel("release"))

	return target, actionInvokeSel
}

// ClearActionTarget removes the action trampoline associated with owner,
// releasing it immediately. It is safe to call when no trampoline is set.
func ClearActionTarget(owner ID) {
	ensureAssociation()
	objcSetAssociatedObjectFn(uintptr(owner), unsafe.Pointer(&actionKey), 0, associationRetainNonatomic)
}
