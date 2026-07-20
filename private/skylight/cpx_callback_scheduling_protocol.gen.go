// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXCallbackScheduling protocol.
type CPXCallbackScheduling interface {
	objectivec.IObject

	// DescheduleForceLogoutCallback protocol.
	DescheduleForceLogoutCallback()

	// DescheduleKillProcessCallback protocol.
	DescheduleKillProcessCallback()

	// ScheduleFixBadForegroundCallbackForProcess protocol.
	ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRec)

	// ScheduleForceLogoutCallbackForTime protocol.
	ScheduleForceLogoutCallbackForTime(time float64)

	// ScheduleKillProcessCallbackForTime protocol.
	ScheduleKillProcessCallbackForTime(time float64)
}

// CPXCallbackSchedulingObject wraps an existing Objective-C object that conforms to the CPXCallbackScheduling protocol.
type CPXCallbackSchedulingObject struct {
	objectivec.Object
}

func (o CPXCallbackSchedulingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXCallbackSchedulingObjectFromID constructs a [CPXCallbackSchedulingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXCallbackSchedulingObjectFromID(id objc.ID) CPXCallbackSchedulingObject {
	return CPXCallbackSchedulingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXCallbackSchedulingObject) DescheduleForceLogoutCallback() {
	objc.Send[struct{}](o.ID, objc.Sel("descheduleForceLogoutCallback"))
}
func (o CPXCallbackSchedulingObject) DescheduleKillProcessCallback() {
	objc.Send[struct{}](o.ID, objc.Sel("descheduleKillProcessCallback"))
}
func (o CPXCallbackSchedulingObject) ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRec) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleFixBadForegroundCallbackForProcess:"), process)
}
func (o CPXCallbackSchedulingObject) ScheduleForceLogoutCallbackForTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleForceLogoutCallbackForTime:"), time)
}
func (o CPXCallbackSchedulingObject) ScheduleKillProcessCallbackForTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleKillProcessCallbackForTime:"), time)
}
