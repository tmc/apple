// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSGraphViewerNodePropertySPI protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphViewerNodePropertySPI
type MPSGraphViewerNodePropertySPI interface {
	objectivec.IObject
}

// MPSGraphViewerNodePropertySPIObject wraps an existing Objective-C object that conforms to the MPSGraphViewerNodePropertySPI protocol.
type MPSGraphViewerNodePropertySPIObject struct {
	objectivec.Object
}

func (o MPSGraphViewerNodePropertySPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSGraphViewerNodePropertySPIObjectFromID constructs a [MPSGraphViewerNodePropertySPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSGraphViewerNodePropertySPIObjectFromID(id objc.ID) MPSGraphViewerNodePropertySPIObject {
	return MPSGraphViewerNodePropertySPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphViewerNodePropertySPI/jsonDictionary
func (o MPSGraphViewerNodePropertySPIObject) JsonDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("jsonDictionary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphViewerNodePropertySPI/name
func (o MPSGraphViewerNodePropertySPIObject) Name() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSGraphViewerNodePropertySPI/type
func (o MPSGraphViewerNodePropertySPIObject) Type() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("type"))
	return objectivec.Object{ID: rv}
}
