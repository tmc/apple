// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that enables the delegate of a metadata query to provide substitute results or attributes.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryDelegate
type NSMetadataQueryDelegate interface {
	objectivec.IObject
}

// NSMetadataQueryDelegateObject wraps an existing Objective-C object that conforms to the NSMetadataQueryDelegate protocol.
type NSMetadataQueryDelegateObject struct {
	objectivec.Object
}
func (o NSMetadataQueryDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSMetadataQueryDelegateObjectFromID constructs a [NSMetadataQueryDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMetadataQueryDelegateObjectFromID(id objc.ID) NSMetadataQueryDelegateObject {
	return NSMetadataQueryDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a different object for a given query result object.
//
// query: The query that produced the result object to replace.
//
// result: The query result object to replace.
//
// # Return Value
// 
// Object that replaces the query result object.
//
// # Discussion
// 
// By default query result objects are instances of the [NSMetadataItem]
// class. By implementing this method, you can return an object of a different
// class type for the specified result object.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryDelegate/metadataQuery(_:replacementObjectForResultObject:)
func (o NSMetadataQueryDelegateObject) MetadataQueryReplacementObjectForResultObject(query INSMetadataQuery, result INSMetadataItem) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metadataQuery:replacementObjectForResultObject:"), query, result)
	return objectivec.Object{ID: rv}
	}
// Returns a different value for a given attribute and value.
//
// query: The query that produced the result object with `attrName`.
//
// attrName: The attribute in question.
//
// attrValue: The attribute value to replace.
//
// # Return Value
// 
// Object that replaces the value of `attrName` in the result object
//
// # Discussion
// 
// The delegate implementation of this method could convert specific query
// attribute values to other attribute values, for example, converting date
// object values to formatted strings for display.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryDelegate/metadataQuery(_:replacementValueForAttribute:value:)
func (o NSMetadataQueryDelegateObject) MetadataQueryReplacementValueForAttributeValue(query INSMetadataQuery, attrName string, attrValue objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metadataQuery:replacementValueForAttribute:value:"), query, objc.String(attrName), attrValue)
	return objectivec.Object{ID: rv}
	}

