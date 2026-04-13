//go:build darwin

package objc_test

import (
	"runtime"
	"sync"
	"testing"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	privatevirtualization "github.com/tmc/apple/private/virtualization"
)

var (
	abiProbeOnce  sync.Once
	abiProbeClass objc.Class
	abiProbeErr   string

	lastPointerArg unsafe.Pointer
	lastPointerIdx uint32

	lastObjectArg objc.ID
	lastObjectIdx uint32

	lastStackArgs struct {
		a1 uint64
		a2 uint64
		a3 uint64
		a4 uint64
		a5 uint64
		a6 uint64
		a7 uint64
		a8 uint64
		a9 uint16
	}
)

func ensureABIProbeClass() (objc.Class, string) {
	abiProbeOnce.Do(func() {
		const name = "GoObjCSendABIProbe"
		if cls := objc.GetClass(name); cls != 0 {
			abiProbeClass = cls
			return
		}

		cls := objectivec.Objc_allocateClassPair(objc.GetClass("NSObject"), name, 0)
		if cls == 0 {
			abiProbeErr = "objc_allocateClassPair returned nil"
			return
		}

		if !objc.AddMethod(cls, objc.Sel("takePointer:index:"), purego.NewCallback(abiTakePointer), "v@:^vI") {
			abiProbeErr = "failed adding takePointer:index:"
			return
		}
		if !objc.AddMethod(cls, objc.Sel("takeObject:index:"), purego.NewCallback(abiTakeObject), "v@:@I") {
			abiProbeErr = "failed adding takeObject:index:"
			return
		}
		if !objc.AddMethod(cls, objc.Sel("takeStackA:b:c:d:e:f:g:h:i:"), purego.NewCallback(abiTakeStack), "v@:QQQQQQQQS") {
			abiProbeErr = "failed adding takeStackA:b:c:d:e:f:g:h:i:"
			return
		}

		objc.RegisterClassPair(cls)
		abiProbeClass = cls
	})
	return abiProbeClass, abiProbeErr
}

func newABIProbe(t *testing.T) objc.ID {
	t.Helper()

	cls, err := ensureABIProbeClass()
	if err != "" {
		t.Fatal(err)
	}
	if cls == 0 {
		t.Fatal("probe class is nil")
	}

	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("alloc"))
	if id == 0 {
		t.Fatal("alloc returned nil")
	}
	id = objc.Send[objc.ID](id, objc.Sel("init"))
	if id == 0 {
		t.Fatal("init returned nil")
	}
	return id
}

func abiTakePointer(_ objc.ID, _ objc.SEL, ptr unsafe.Pointer, idx uint32) {
	lastPointerArg = ptr
	lastPointerIdx = idx
}

func abiTakeObject(_ objc.ID, _ objc.SEL, id objc.ID, idx uint32) {
	lastObjectArg = id
	lastObjectIdx = idx
}

func abiTakeStack(_ objc.ID, _ objc.SEL, a1, a2, a3, a4, a5, a6, a7, a8 uint64, a9 uint16) {
	lastStackArgs = struct {
		a1 uint64
		a2 uint64
		a3 uint64
		a4 uint64
		a5 uint64
		a6 uint64
		a7 uint64
		a8 uint64
		a9 uint16
	}{
		a1: a1,
		a2: a2,
		a3: a3,
		a4: a4,
		a5: a5,
		a6: a6,
		a7: a7,
		a8: a8,
		a9: a9,
	}
}

func TestSendABIPointerAndUint32(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip("arm64-specific ABI probe")
	}

	objc.AutoreleasePool(func() {
		probe := newABIProbe(t)
		var report [8]byte
		wantPtr := unsafe.Pointer(&report[0])
		lastPointerArg = nil
		lastPointerIdx = 0

		objc.Send[struct{}](probe, objc.Sel("takePointer:index:"), wantPtr, uint32(7))

		if lastPointerArg != wantPtr {
			t.Fatalf("pointer arg = %p, want %p", lastPointerArg, wantPtr)
		}
		if lastPointerIdx != 7 {
			t.Fatalf("pointer index = %d, want 7", lastPointerIdx)
		}
	})
}

func TestSendABIObjectAndUint32(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip("arm64-specific ABI probe")
	}

	objc.AutoreleasePool(func() {
		probe := newABIProbe(t)
		want := objc.String("tab")
		lastObjectArg = 0
		lastObjectIdx = 0

		objc.Send[struct{}](probe, objc.Sel("takeObject:index:"), want, uint32(3))

		if lastObjectArg != want {
			t.Fatalf("object arg = %#x, want %#x", lastObjectArg, want)
		}
		if lastObjectIdx != 3 {
			t.Fatalf("object index = %d, want 3", lastObjectIdx)
		}
	})
}

func TestSendABIStackPassedScalars(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip("arm64-specific ABI probe")
	}

	objc.AutoreleasePool(func() {
		probe := newABIProbe(t)
		lastStackArgs = struct {
			a1 uint64
			a2 uint64
			a3 uint64
			a4 uint64
			a5 uint64
			a6 uint64
			a7 uint64
			a8 uint64
			a9 uint16
		}{}

		objc.Send[struct{}](probe, objc.Sel("takeStackA:b:c:d:e:f:g:h:i:"),
			uint64(1), uint64(2), uint64(3), uint64(4), uint64(5),
			uint64(6), uint64(7), uint64(8), uint16(0x1234))

		if lastStackArgs.a1 != 1 || lastStackArgs.a2 != 2 || lastStackArgs.a3 != 3 || lastStackArgs.a4 != 4 ||
			lastStackArgs.a5 != 5 || lastStackArgs.a6 != 6 || lastStackArgs.a7 != 7 || lastStackArgs.a8 != 8 ||
			lastStackArgs.a9 != 0x1234 {
			t.Fatalf("stack args = %+v", lastStackArgs)
		}
	})
}

func TestVirtualMachineSelectorContracts(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip("arm64-specific selector probe")
	}

	_ = privatevirtualization.GetVZVirtualMachineClass()
	cls := objc.GetClass("VZVirtualMachine")
	if cls == 0 {
		t.Skip("VZVirtualMachine class unavailable")
	}

	cases := []struct {
		selector string
		wantArgs []string
	}{
		{selector: "sendKeyboardEvents:keyboardID:", wantArgs: []string{"@", ":", "^v", "I"}},
		{selector: "sendPointerNSEvent:pointingDeviceIndex:", wantArgs: []string{"@", ":", "@", "I"}},
		{selector: "sendIOHIDEvents:hidDeviceIndex:", wantArgs: []string{"@", ":", "^v", "I"}},
		{selector: "_processHIDReports:forDevice:deviceType:", wantArgs: []string{"@", ":", "^v", "I", "i"}},
	}

	for _, tc := range cases {
		t.Run(tc.selector, func(t *testing.T) {
			method := objectivec.Class_getInstanceMethod(cls, objectivec.SEL(objc.Sel(tc.selector)))
			if method == 0 {
				t.Fatalf("method %q not found", tc.selector)
			}

			fullEncoding := objc.GoString(objectivec.Method_getTypeEncoding(method))
			returnType := objc.GoString(objectivec.Method_copyReturnType(method))
			argCount := objectivec.Method_getNumberOfArguments(method)
			argTypes := make([]string, 0, argCount)
			for i := uint(0); i < argCount; i++ {
				argTypes = append(argTypes, objc.GoString(objectivec.Method_copyArgumentType(method, i)))
			}

			t.Logf("selector=%s encoding=%q return=%q args=%v", tc.selector, fullEncoding, returnType, argTypes)
			if returnType != "v" {
				t.Fatalf("return type = %q, want %q", returnType, "v")
			}
			if len(argTypes) != len(tc.wantArgs) {
				t.Fatalf("arg count = %d, want %d (%v)", len(argTypes), len(tc.wantArgs), tc.wantArgs)
			}
			for i, want := range tc.wantArgs {
				if argTypes[i] != want {
					t.Fatalf("arg %d type = %q, want %q", i, argTypes[i], want)
				}
			}
		})
	}
}

func TestPrivateObjectInputSelectorContracts(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip("arm64-specific selector probe")
	}

	cases := []struct {
		class    string
		selector string
		wantArgs []string
	}{
		{class: "_VZKeyboard", selector: "sendKeyEvents:", wantArgs: []string{"@", ":", "@"}},
		{class: "_VZHIDDevice", selector: "sendIOHIDEvents:", wantArgs: []string{"@", ":", "@"}},
		{class: "_VZKeyEvent", selector: "initWithEvent:", wantArgs: []string{"@", ":", "@"}},
		{class: "_VZIOHIDEvent", selector: "initWithIOHIDEvent:", wantArgs: []string{"@", ":", "^{__IOHIDEvent=}"}},
		{class: "_VZIOHIDEvent", selector: "event", wantArgs: []string{"@", ":"}},
	}

	for _, tc := range cases {
		t.Run(tc.class+"."+tc.selector, func(t *testing.T) {
			cls := objc.GetClass(tc.class)
			if cls == 0 {
				t.Skipf("%s unavailable", tc.class)
			}

			method := objectivec.Class_getInstanceMethod(cls, objectivec.SEL(objc.Sel(tc.selector)))
			if method == 0 {
				t.Fatalf("method %q not found on %s", tc.selector, tc.class)
			}

			returnType := objc.GoString(objectivec.Method_copyReturnType(method))
			argCount := objectivec.Method_getNumberOfArguments(method)
			argTypes := make([]string, 0, argCount)
			for i := uint(0); i < argCount; i++ {
				argTypes = append(argTypes, objc.GoString(objectivec.Method_copyArgumentType(method, i)))
			}

			if tc.selector == "event" {
				if returnType != "^{__IOHIDEvent=}" {
					t.Fatalf("return type = %q, want ^{__IOHIDEvent=}", returnType)
				}
			} else if returnType != "v" && tc.selector[:4] == "init" {
				if returnType != "@" {
					t.Fatalf("return type = %q, want @ for init", returnType)
				}
			} else if returnType != "v" && tc.selector[:4] != "init" {
				t.Fatalf("return type = %q, want v", returnType)
			}

			if len(argTypes) != len(tc.wantArgs) {
				t.Fatalf("arg count = %d, want %d (%v)", len(argTypes), len(tc.wantArgs), tc.wantArgs)
			}
			for i, want := range tc.wantArgs {
				if argTypes[i] != want {
					t.Fatalf("arg %d type = %q, want %q", i, argTypes[i], want)
				}
			}
		})
	}
}
