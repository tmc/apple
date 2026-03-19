// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A fast-addition container for collecting data during pipeline state creation.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializer
type MTL4PipelineDataSetSerializer interface {
	objectivec.IObject

	// Serializes a pipeline data set to an archive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializer/serializeAsArchiveAndFlush(url:)
	SerializeAsArchiveAndFlushToURLError(url foundation.INSURL) (bool, error)

	// Serializes a serializer data set to a pipeline script as raw data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializer/serializeAsPipelinesScript()
	SerializeAsPipelinesScriptWithError() (foundation.INSData, error)
}

// MTL4PipelineDataSetSerializerObject wraps an existing Objective-C object that conforms to the MTL4PipelineDataSetSerializer protocol.
type MTL4PipelineDataSetSerializerObject struct {
	objectivec.Object
}
func (o MTL4PipelineDataSetSerializerObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4PipelineDataSetSerializerObjectFromID constructs a [MTL4PipelineDataSetSerializerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4PipelineDataSetSerializerObjectFromID(id objc.ID) MTL4PipelineDataSetSerializerObject {
	return MTL4PipelineDataSetSerializerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Serializes a pipeline data set to an archive.
//
// url: The URL used to serialize the serializer data set as an archive to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializer/serializeAsArchiveAndFlush(url:)
func (o MTL4PipelineDataSetSerializerObject) SerializeAsArchiveAndFlushToURLError(url foundation.INSURL) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("serializeAsArchiveAndFlushToURL:error:"), url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// Serializes a serializer data set to a pipeline script as raw data.
//
// # Return Value
// 
// An [NSData] instance containing the pipeline script.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializer/serializeAsPipelinesScript()
func (o MTL4PipelineDataSetSerializerObject) SerializeAsPipelinesScriptWithError() (foundation.INSData, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("serializeAsPipelinesScriptWithError:"))
	if err != nil {
		return nil, err
	}
	return foundation.NSDataFromID(rv), nil
	}

