//go:build darwin

package objcinspect_test

import (
	"fmt"
	"reflect"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcinspect"
)

func ExampleParseSignature() {
	// Parse an Objective-C method type encoding.
	// "v@:@" represents a void method taking self (id), _cmd (SEL), and an object argument (id).
	sig, err := objcinspect.ParseSignature("v@:@")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Return kind:", sig.Return.Kind)
	fmt.Println("Total args:", len(sig.Args))
	for i, arg := range sig.Args {
		fmt.Printf("Arg %d: kind=%s encoding=%s\n", i, arg.Kind, arg.Encoding)
	}

	// Output:
	// Return kind: void
	// Total args: 3
	// Arg 0: kind=id encoding=@
	// Arg 1: kind=SEL encoding=:
	// Arg 2: kind=id encoding=@
}

func ExampleParseStruct() {
	// Parse an Objective-C struct type encoding.
	name, fields, size, err := objcinspect.ParseStruct("{CGPoint=dd}")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Struct: %s, Fields: %d, Size: %d\n", name, len(fields), size)
	for i, f := range fields {
		fmt.Printf("Field %d: kind=%s align=%d\n", i, f.Kind, f.Align())
	}

	// Output:
	// Struct: CGPoint, Fields: 2, Size: 16
	// Field 0: kind=double align=8
	// Field 1: kind=double align=8
}

func ExampleLooksLikeObject() {
	// 0 is invalid address (reject)
	invalidZero := objcinspect.LooksLikeObject(0)
	// 0x1001 is unaligned address (reject)
	invalidUnaligned := objcinspect.LooksLikeObject(0x1001)

	// Valid NSString object ID passes validation
	nsStr := objc.String("Test")
	validObj := objcinspect.LooksLikeObject(nsStr)

	fmt.Printf("Zero ID valid: %t\n", invalidZero)
	fmt.Printf("Unaligned ID valid: %t\n", invalidUnaligned)
	fmt.Printf("NSString object valid: %t\n", validObj)

	// Output:
	// Zero ID valid: false
	// Unaligned ID valid: false
	// NSString object valid: true
}

func ExampleRegistry() {
	registry := objcinspect.NewRegistry()

	// Parse and register a struct encoding for CGRect (registers CGRect, CGPoint, and CGSize).
	err := registry.RegisterEncoding("{CGRect={CGPoint=dd}{CGSize=dd}}")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Registry count:", registry.Len())
	opaque := objcinspect.Type{Kind: objcinspect.KindStruct, Name: "CGPoint"}
	resolved := registry.Resolve(opaque)
	fmt.Printf("Resolved CGPoint: name=%s fields=%d size=%d\n", resolved.Name, len(resolved.Fields), resolved.NaturalSize())

	// Output:
	// Registry count: 3
	// Resolved CGPoint: name=CGPoint fields=2 size=16
}

func ExampleCheck() {
	nsStr := objc.String("hello")
	selLength := objc.Sel("length")

	// Check calling length expecting uint return type
	err := objcinspect.Check(nsStr, selLength, reflect.TypeOf(uint(0)))
	fmt.Println("Check length (uint return):", err)

	// Check calling length expecting mismatched argument count
	errMismatch := objcinspect.Check(nsStr, selLength, reflect.TypeOf(uint(0)), "extra arg")
	fmt.Println("Check length (extra arg):", errMismatch)

	// Output:
	// Check length (uint return): <nil>
	// Check length (extra arg): argument count mismatch: got 1 args, selector expects 0
}
