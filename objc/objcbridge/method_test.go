package objcbridge

import (
	"testing"
	"unsafe"

	"github.com/tmc/apple/objc"
)

func TestMethodTypeEncoding(t *testing.T) {
	tests := []struct {
		name string
		fn   any
		want string
	}{
		{"void", func(objc.ID, objc.SEL) {}, "v@:"},
		{"object", func(objc.ID, objc.SEL, objc.ID) objc.ID { return 0 }, "@@:@"},
		{"values", func(objc.ID, objc.SEL, bool, int32, uint64, unsafe.Pointer) {}, "v@:BiQ^v"},
		{"integers", func(objc.ID, objc.SEL, int64, uint64, uintptr) int64 { return -1 }, "q@:qQQ"},
		{"block", func(objc.ID, objc.SEL, objc.Block) {}, "v@:@?"},
		{"nil function", (func(objc.ID, objc.SEL))(nil), ""},
		{"variadic", func(objc.ID, objc.SEL, ...int) {}, ""},
		{"not function", 1, ""},
		{"missing method arguments", func() {}, ""},
		{"multiple results", func(objc.ID, objc.SEL) (int, int) { return 0, 0 }, ""},
		{"unsupported argument", func(objc.ID, objc.SEL, []byte) {}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := methodTypeEncoding(tt.fn)
			if tt.want == "" {
				if err == nil {
					t.Fatalf("methodTypeEncoding() = %q, nil; want error", got)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("methodTypeEncoding() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAddMethodsInstallsTypeEncoding(t *testing.T) {
	class, err := objc.RegisterClass("ObjCBridgeMethodEncodingTest", objc.GetClass("NSObject"), nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	selector := objc.Sel("objcBridgeMethodEncoding:flag:")
	if err := AddMethods(class, "ObjCBridgeMethodEncodingTest", []objc.MethodDef{{
		Cmd: selector,
		Fn:  func(objc.ID, objc.SEL, objc.ID, bool) {},
	}}); err != nil {
		t.Fatal(err)
	}

	instance := objc.Send[objc.ID](objc.ID(class), objc.Sel("new"))
	defer objc.Send[struct{}](instance, objc.Sel("release"))
	signature := objc.Send[objc.ID](instance, objc.Sel("methodSignatureForSelector:"), selector)
	if signature == 0 {
		t.Fatal("methodSignatureForSelector: returned nil")
	}
	if got := objc.Send[uintptr](signature, objc.Sel("numberOfArguments")); got != 4 {
		t.Fatalf("numberOfArguments = %d, want 4", got)
	}
}
