// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ArchiveReader protocol.
//
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader
type ArchiveReader interface {
	objectivec.IObject
}

// ArchiveReaderObject wraps an existing Objective-C object that conforms to the ArchiveReader protocol.
type ArchiveReaderObject struct {
	objectivec.Object
}
func (o ArchiveReaderObject) BaseObject() objectivec.Object {
	return o.Object
}

// ArchiveReaderObjectFromID constructs a [ArchiveReaderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ArchiveReaderObjectFromID(id objc.ID) ArchiveReaderObject {
	return ArchiveReaderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/copyLayerShapesToContainer:
func (o ArchiveReaderObject) CopyLayerShapesToContainer(container objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("copyLayerShapesToContainer:"), container)
	}
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/layerInfos
func (o ArchiveReaderObject) LayerInfos() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("layerInfos"))
	return objectivec.Object{ID: rv}
	}
//
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/loadUpdatableParams:
func (o ArchiveReaderObject) LoadUpdatableParams(params []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("loadUpdatableParams:"), objectivec.IObjectSliceToNSArray(params))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/modelName
func (o ArchiveReaderObject) ModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelName"))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/modelPath
func (o ArchiveReaderObject) ModelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/CoreML/_ArchiveReader/transformParams
func (o ArchiveReaderObject) TransformParams() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("transformParams"))
	return objectivec.Object{ID: rv}
	}

