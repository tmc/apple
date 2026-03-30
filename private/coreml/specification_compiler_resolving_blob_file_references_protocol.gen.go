// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLSpecificationCompilerResolvingBlobFileReferences protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLSpecificationCompilerResolvingBlobFileReferences
type MLSpecificationCompilerResolvingBlobFileReferences interface {
	objectivec.IObject
}

// MLSpecificationCompilerResolvingBlobFileReferencesObject wraps an existing Objective-C object that conforms to the MLSpecificationCompilerResolvingBlobFileReferences protocol.
type MLSpecificationCompilerResolvingBlobFileReferencesObject struct {
	objectivec.Object
}

func (o MLSpecificationCompilerResolvingBlobFileReferencesObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLSpecificationCompilerResolvingBlobFileReferencesObjectFromID constructs a [MLSpecificationCompilerResolvingBlobFileReferencesObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLSpecificationCompilerResolvingBlobFileReferencesObjectFromID(id objc.ID) MLSpecificationCompilerResolvingBlobFileReferencesObject {
	return MLSpecificationCompilerResolvingBlobFileReferencesObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLSpecificationCompilerResolvingBlobFileReferences/compileSpecification:blobMapping:toArchive:options:error:
func (o MLSpecificationCompilerResolvingBlobFileReferencesObject) CompileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
