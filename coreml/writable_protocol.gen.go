// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that saves a machine learning type to the file system.
//
// See: https://developer.apple.com/documentation/CoreML/MLWritable
type MLWritable interface {
	objectivec.IObject

	// Exports a machine learning file to the file system.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLWritable/write(to:)
	WriteToURLError(url foundation.INSURL) (bool, error)
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




// Exports a machine learning file to the file system.
//
// url: The location in the file system where the file should be written.
//
// See: https://developer.apple.com/documentation/CoreML/MLWritable/write(to:)

func (o MLWritableObject) WriteToURLError(url foundation.INSURL) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("writeToURL:error:"), url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}








