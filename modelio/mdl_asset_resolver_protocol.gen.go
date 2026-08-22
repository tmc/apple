// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MDLAssetResolver protocol.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLAssetResolver
type MDLAssetResolver interface {
	objectivec.IObject
}

// MDLAssetResolverObject wraps an existing Objective-C object that conforms to the MDLAssetResolver protocol.
type MDLAssetResolverObject struct {
	objectivec.Object
}

func (o MDLAssetResolverObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLAssetResolverObjectFromID constructs a [MDLAssetResolverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLAssetResolverObjectFromID(id objc.ID) MDLAssetResolverObject {
	return MDLAssetResolverObject{
		Object: objectivec.ObjectFromID(id),
	}
}
