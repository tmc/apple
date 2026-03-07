// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Represents a raw or compressed file, such as a resource asset file in your app’s bundle.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOFileHandle
type MTLIOFileHandle interface {
	objectivec.IObject

	// An optional name for the file that the handle represents.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOFileHandle/label
	Label() string

	// An optional name for the file that the handle represents.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOFileHandle/label
	SetLabel(value string)
}



// MTLIOFileHandleObject wraps an existing Objective-C object that conforms to the MTLIOFileHandle protocol.
type MTLIOFileHandleObject struct {
	objectivec.Object
}
func (o MTLIOFileHandleObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLIOFileHandleObjectFromID constructs a [MTLIOFileHandleObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIOFileHandleObjectFromID(id objc.ID) MTLIOFileHandleObject {
	return MTLIOFileHandleObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// An optional name for the file that the handle represents.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOFileHandle/label

func (o MTLIOFileHandleObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}




func (o MTLIOFileHandleObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}





