//go:build darwin

package objectivec_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

func ExampleObject() {
	// Wrap a raw Objective-C NSString object ID in objectivec.Object.
	rawID := objc.String("Hello from objectivec package")
	obj := objectivec.ObjectFromID(rawID)

	fmt.Println("Description:", obj.Description())
	if obj.RespondsToSelector(objc.Sel("length")) {
		length := objc.Send[uint](obj.ID, objc.Sel("length"))
		fmt.Println("Length:", length)
	}

	// Output:
	// Description: Hello from objectivec package
	// Length: 29
}

func ExampleNSObjectObject() {
	// Wrap an object conforming to NSObject protocol.
	rawID := objc.String("NSObject Test")
	nsObj := objectivec.NSObjectObjectFromID(rawID)

	nsStringClass := objc.GetClass("NSString")
	if nsObj.IsKindOfClass(nsStringClass) {
		superDesc := objectivec.ObjectFromID(objc.Send[objc.ID](objc.ID(nsObj.Superclass()), objc.Sel("description"))).Description()
		fmt.Println("Superclass description:", superDesc)
	}

	// Output:
	// Superclass description: NSMutableString
}

func ExampleStringSliceToNSArray() {
	// Convert a Go []string into an Objective-C NSArray object ID.
	words := []string{"apple", "banana", "cherry"}
	arrayID := objectivec.StringSliceToNSArray(words)

	count := objc.Send[uint](arrayID, objc.Sel("count"))
	firstObj := objc.Send[objc.ID](arrayID, objc.Sel("objectAtIndex:"), uintptr(0))
	lastObj := objc.Send[objc.ID](arrayID, objc.Sel("objectAtIndex:"), uintptr(2))

	fmt.Println("NSArray count:", count)
	fmt.Println("First element:", objc.GoString(objc.Send[*byte](firstObj, objc.Sel("UTF8String"))))
	fmt.Println("Last element:", objc.GoString(objc.Send[*byte](lastObj, objc.Sel("UTF8String"))))

	// Output:
	// NSArray count: 3
	// First element: apple
	// Last element: cherry
}
