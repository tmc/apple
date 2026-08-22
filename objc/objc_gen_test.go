//go:build darwin

package objc

import (
	"errors"
	"testing"

	pobjc "github.com/ebitengine/purego/objc"
)

var _ func(Class, SEL, IMP, string) bool = AddMethod

func TestGetClassIsExact(t *testing.T) {
	const (
		exactName = "_GoGetClassExactTest"
		otherName = "GoGetClassExactTest"
	)
	class, err := pobjc.RegisterClass(exactName, pobjc.GetClass("NSObject"), nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := GetClass(exactName); got != class {
		t.Fatalf("GetClass(%q) = %d, want %d", exactName, got, class)
	}
	// Generated bindings pass the name the runtime reports, underscore
	// included, so a near miss must stay a miss rather than resolve to a
	// class the caller did not name.
	if got := GetClass(otherName); got != 0 {
		t.Fatalf("GetClass(%q) = %d, want 0", otherName, got)
	}
}

func TestObjCError(t *testing.T) {
	AutoreleasePool(func() {
		const (
			description = "operation failed"
			domain      = "GoNSErrorTestDomain"
			code        = 7
		)
		userInfo := Send[ID](ID(GetClass("NSDictionary")), Sel("dictionaryWithObject:forKey:"),
			String(description), String("NSLocalizedDescription"))
		errorID := Send[ID](ID(GetClass("NSError")), Sel("errorWithDomain:code:userInfo:"),
			String(domain), code, userInfo)

		sel := Sel("failWithError:")
		class, err := pobjc.RegisterClass("GoNSErrorOutTest", pobjc.GetClass("NSObject"), nil, nil, []pobjc.MethodDef{
			{
				Cmd: sel,
				Fn: func(_ ID, _ SEL, out *ID) bool {
					*out = errorID
					return false
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		object := Send[ID](ID(class), Sel("new"))
		defer Send[struct{}](object, Sel("release"))

		ok, gotErr := SendWithError[bool](object, sel)
		if ok {
			t.Fatal("SendWithError returned true")
		}
		var objcError ObjCError
		if !errors.As(gotErr, &objcError) {
			t.Fatalf("error type = %T, want ObjCError", gotErr)
		}
		if objcError.ID != errorID {
			t.Fatalf("ObjCError.ID = %d, want %d", objcError.ID, errorID)
		}
		if got, want := objcError.Error(), description+" ("+domain+" error 7)"; got != want {
			t.Fatalf("ObjCError.Error() = %q, want %q", got, want)
		}
	})
}
