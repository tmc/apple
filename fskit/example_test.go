//go:build darwin

package fskit_test

import (
	"fmt"

	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

func ExampleFSAccessMask() {
	fmt.Println(fskit.FSAccessAddFile)
	fmt.Println(fskit.FSAccessExecute)
	fmt.Println(fskit.FSAccessDelete)

	// Output:
	// FSAccessAddFile
	// FSAccessExecute
	// FSAccessDelete
}

func ExampleFSItemAttributesClass() {
	cls := fskit.GetFSItemAttributesClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)

	// Output:
	// class name: FSItemAttributes
}
