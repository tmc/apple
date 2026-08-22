//go:build darwin

package objc

import (
	"testing"

	pobjc "github.com/ebitengine/purego/objc"
)

func TestActionTargetReplacementReleasesOldTargets(t *testing.T) {
	owner := Send[ID](ID(GetClass("NSObject")), Sel("new"))
	defer Send[struct{}](owner, Sel("release"))

	weak := Send[ID](ID(GetClass("NSHashTable")), Sel("weakObjectsHashTable"))
	const replacements = 32
	for range replacements {
		target, _ := NewActionTarget(owner, func(ID) {})
		Send[struct{}](weak, Sel("addObject:"), target)
	}

	if got := actionHandlerCount(); got != 1 {
		t.Errorf("handler count = %d, want 1", got)
	}
	objects := Send[ID](weak, Sel("allObjects"))
	if got := Send[uint](objects, Sel("count")); got != 1 {
		t.Errorf("live target count = %d, want 1", got)
	}
}

func TestActionTargetInvoke(t *testing.T) {
	owner := Send[ID](ID(GetClass("NSObject")), Sel("new"))
	defer Send[struct{}](owner, Sel("release"))

	want := ID(42)
	var got ID
	target, sel := NewActionTarget(owner, func(sender ID) { got = sender })
	Send[struct{}](target, sel, want)
	if got != want {
		t.Fatalf("handler sender = %d, want %d", got, want)
	}
}

func TestActionTargetClear(t *testing.T) {
	owner := Send[ID](ID(GetClass("NSObject")), Sel("new"))
	defer Send[struct{}](owner, Sel("release"))

	ClearActionTarget(owner)

	weak := Send[ID](ID(GetClass("NSHashTable")), Sel("weakObjectsHashTable"))
	called := false
	target, _ := NewActionTarget(owner, func(ID) { called = true })
	Send[struct{}](weak, Sel("addObject:"), target)
	ClearActionTarget(owner)
	actionTargetInvoke(target, actionInvokeSel, owner)
	if called {
		t.Fatal("cleared action invoked its old handler")
	}
	if _, ok := actionHandlers.Load(target); ok {
		t.Fatal("cleared action retained its old handler")
	}
	objects := Send[ID](weak, Sel("allObjects"))
	if got := Send[uint](objects, Sel("count")); got != 0 {
		t.Errorf("live target count after clear = %d, want 0", got)
	}
}

func TestActionTargetClassCollision(t *testing.T) {
	ensureActionTarget()

	const name = "GoActionTargetCollision"
	foreignCalled := false
	foreign, err := pobjc.RegisterClass(name, pobjc.GetClass("NSObject"), nil, nil, []pobjc.MethodDef{
		{
			Cmd: actionInvokeSel,
			Fn:  func(ID, SEL, ID) { foreignCalled = true },
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	class := registerActionTargetClass(name)
	if class == foreign {
		t.Fatal("registration reused a foreign class")
	}

	foreignObject := Send[ID](ID(foreign), Sel("new"))
	defer Send[struct{}](foreignObject, Sel("release"))
	Send[struct{}](foreignObject, actionInvokeSel, ID(0))
	if !foreignCalled {
		t.Fatal("foreign class did not use its own callback")
	}

	object := Send[ID](ID(class), Sel("new"))
	called := false
	actionHandlers.Store(object, func(ID) { called = true })
	Send[struct{}](object, actionInvokeSel, ID(0))
	Send[struct{}](object, Sel("release"))
	if !called {
		t.Fatal("uniquified class did not use this package's callback")
	}
}

func actionHandlerCount() int {
	count := 0
	actionHandlers.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}
