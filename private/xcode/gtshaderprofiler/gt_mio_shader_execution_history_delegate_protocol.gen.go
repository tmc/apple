// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioShaderExecutionHistoryDelegate protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDelegate
type GTMioShaderExecutionHistoryDelegate interface {
	objectivec.IObject
}

// GTMioShaderExecutionHistoryDelegateObject wraps an existing Objective-C object that conforms to the GTMioShaderExecutionHistoryDelegate protocol.
type GTMioShaderExecutionHistoryDelegateObject struct {
	objectivec.Object
}
func (o GTMioShaderExecutionHistoryDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioShaderExecutionHistoryDelegateObjectFromID constructs a [GTMioShaderExecutionHistoryDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioShaderExecutionHistoryDelegateObjectFromID(id objc.ID) GTMioShaderExecutionHistoryDelegateObject {
	return GTMioShaderExecutionHistoryDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDelegate/uniqueIdentifierForFile:debugFunctionName:line:column:
func (o GTMioShaderExecutionHistoryDelegateObject) UniqueIdentifierForFileDebugFunctionNameLineColumn(file objectivec.IObject, name objectivec.IObject, line uint32, column uint32) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("uniqueIdentifierForFile:debugFunctionName:line:column:"), file, name, line, column)
	return rv
	}

