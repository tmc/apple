// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Maintenance operations for a file system’s resources.
//
// See: https://developer.apple.com/documentation/FSKit/FSManageableResourceMaintenanceOperations
type FSManageableResourceMaintenanceOperations interface {
	objectivec.IObject

	// Starts checking the file system with the given options.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSManageableResourceMaintenanceOperations/startCheck(task:options:)
	StartCheckWithTaskOptionsError(task IFSTask, options IFSTaskOptions) (foundation.NSProgress, error)

	// Starts formatting the file system with the given options.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSManageableResourceMaintenanceOperations/startFormat(task:options:)
	StartFormatWithTaskOptionsError(task IFSTask, options IFSTaskOptions) (foundation.NSProgress, error)
}

// FSManageableResourceMaintenanceOperationsObject wraps an existing Objective-C object that conforms to the FSManageableResourceMaintenanceOperations protocol.
type FSManageableResourceMaintenanceOperationsObject struct {
	objectivec.Object
}

func (o FSManageableResourceMaintenanceOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSManageableResourceMaintenanceOperationsObjectFromID constructs a [FSManageableResourceMaintenanceOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSManageableResourceMaintenanceOperationsObjectFromID(id objc.ID) FSManageableResourceMaintenanceOperationsObject {
	return FSManageableResourceMaintenanceOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Starts checking the file system with the given options.
//
// task: A task object you use to communicate back to the client.
//
// options: Options for performing the check.
//
// # Return Value
//
// An [Progress] object that you use to update progress as the check operation
// progresses. Return `nil` if starting the file system check encountered an
// error.
//
// See: https://developer.apple.com/documentation/FSKit/FSManageableResourceMaintenanceOperations/startCheck(task:options:)
//
// [Progress]: https://developer.apple.com/documentation/Foundation/Progress
func (o FSManageableResourceMaintenanceOperationsObject) StartCheckWithTaskOptionsError(task IFSTask, options IFSTaskOptions) (foundation.NSProgress, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("startCheckWithTask:options:error:"), task, options)
	if err != nil {
		return *new(foundation.NSProgress), err
	}
	return foundation.NSProgressFromID(rv), nil
}

// Starts formatting the file system with the given options.
//
// task: A task object you use to communicate back to the client.
//
// options: Options for performing the format.
//
// # Return Value
//
// An [Progress] object that you use to update progress as the format
// operation progresses. Return `nil` if starting to format the file system
// encountered an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSManageableResourceMaintenanceOperations/startFormat(task:options:)
//
// [Progress]: https://developer.apple.com/documentation/Foundation/Progress
func (o FSManageableResourceMaintenanceOperationsObject) StartFormatWithTaskOptionsError(task IFSTask, options IFSTaskOptions) (foundation.NSProgress, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("startFormatWithTask:options:error:"), task, options)
	if err != nil {
		return *new(foundation.NSProgress), err
	}
	return foundation.NSProgressFromID(rv), nil
}
