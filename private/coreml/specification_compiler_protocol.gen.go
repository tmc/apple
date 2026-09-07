// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLSpecificationCompiler protocol.
type MLSpecificationCompiler interface {
	objectivec.IObject

	// CompileSpecificationToArchiveOptionsError protocol.
	CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error)

	// CompiledVersionForSpecificationOptionsError protocol.
	CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error)
}

// MLSpecificationCompilerObject wraps an existing Objective-C object that conforms to the MLSpecificationCompiler protocol.
type MLSpecificationCompilerObject struct {
	objectivec.Object
}

func (o MLSpecificationCompilerObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLSpecificationCompilerObjectFromID constructs a [MLSpecificationCompilerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLSpecificationCompilerObjectFromID(id objc.ID) MLSpecificationCompilerObject {
	return MLSpecificationCompilerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLSpecificationCompilerObject) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLSpecificationCompilerObject) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("compiledVersionForSpecification:options:error:"), specification, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
