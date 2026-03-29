// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MLWritable protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLWritable
type MLWritable interface {
	objectivec.IObject
}

// MLWritableObject wraps an existing Objective-C object that conforms to the MLWritable protocol.
type MLWritableObject struct {
	objectivec.Object
}
func (o MLWritableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLWritableObjectFromID constructs a [MLWritableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLWritableObjectFromID(id objc.ID) MLWritableObject {
	return MLWritableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLWritable/writeToURL:error:
func (o MLWritableObject) WriteToURLError(url foundation.INSURL) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("writeToURL:error:"), url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}

