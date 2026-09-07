//go:build darwin

package objcinspect_test

import (
	"fmt"
	"reflect"

	_ "github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcinspect"
)

func ExampleSendChecked() {
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))
	n, err := objcinspect.SendChecked[uint64](objc.String("hello"), objc.Sel("length"))
	fmt.Println(n, err)
	// Output: 5 <nil>
}

func ExampleCallChecked() {
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))
	err := objcinspect.CallChecked(objc.String("hello"), objc.Sel("hash"))
	fmt.Println(err)
	// Output: <nil>
}

func ExampleSetEnabled() {
	previous := objcinspect.SetEnabled(true)
	defer objcinspect.SetEnabled(previous)
	fmt.Println(objcinspect.Enabled())
	// Output: true
}

func ExampleEnabled() {
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(false))
	fmt.Println(objcinspect.Enabled())
	// Output: false
}

func ExampleCall() {
	call := objcinspect.Call{
		ID:     objc.String("hello"),
		Sel:    objc.Sel("length"),
		Return: reflect.TypeFor[uint64](),
	}
	fmt.Println(objcinspect.CheckAll(call))
	// Output: <nil>
}

func ExampleCheckAll() {
	s := objc.String("hello")
	err := objcinspect.CheckAll(
		objcinspect.Call{ID: s, Sel: objc.Sel("length"), Return: reflect.TypeFor[uint64]()},
		objcinspect.Call{ID: s, Sel: objc.Sel("hash")},
	)
	fmt.Println(err)
	// Output: <nil>
}
