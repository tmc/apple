// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXCallbackScheduling protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling
type CPXCallbackScheduling interface {
	objectivec.IObject

	// DescheduleForceLogoutCallback protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/descheduleForceLogoutCallback
	DescheduleForceLogoutCallback()

	// DescheduleKillProcessCallback protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/descheduleKillProcessCallback
	DescheduleKillProcessCallback()

	// ScheduleFixBadForegroundCallbackForProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleFixBadForegroundCallbackForProcess:
	ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRecRef)

	// ScheduleForceLogoutCallbackForTime protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleForceLogoutCallbackForTime:
	ScheduleForceLogoutCallbackForTime(time float64)

	// ScheduleKillProcessCallbackForTime protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleKillProcessCallbackForTime:
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

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/descheduleForceLogoutCallback
func (o CPXCallbackSchedulingObject) DescheduleForceLogoutCallback() {
	objc.Send[struct{}](o.ID, objc.Sel("descheduleForceLogoutCallback"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/descheduleKillProcessCallback
func (o CPXCallbackSchedulingObject) DescheduleKillProcessCallback() {
	objc.Send[struct{}](o.ID, objc.Sel("descheduleKillProcessCallback"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleFixBadForegroundCallbackForProcess:
func (o CPXCallbackSchedulingObject) ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleFixBadForegroundCallbackForProcess:"), process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleForceLogoutCallbackForTime:
func (o CPXCallbackSchedulingObject) ScheduleForceLogoutCallbackForTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleForceLogoutCallbackForTime:"), time)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduling/scheduleKillProcessCallbackForTime:
func (o CPXCallbackSchedulingObject) ScheduleKillProcessCallbackForTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("scheduleKillProcessCallbackForTime:"), time)
}
