// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXCallbackScheduler] class.
var (
	_CPXCallbackSchedulerClass     CPXCallbackSchedulerClass
	_CPXCallbackSchedulerClassOnce sync.Once
)

func getCPXCallbackSchedulerClass() CPXCallbackSchedulerClass {
	_CPXCallbackSchedulerClassOnce.Do(func() {
		_CPXCallbackSchedulerClass = CPXCallbackSchedulerClass{class: objc.GetClass("CPXCallbackScheduler")}
	})
	return _CPXCallbackSchedulerClass
}

// GetCPXCallbackSchedulerClass returns the class object for CPXCallbackScheduler.
func GetCPXCallbackSchedulerClass() CPXCallbackSchedulerClass {
	return getCPXCallbackSchedulerClass()
}

type CPXCallbackSchedulerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXCallbackSchedulerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXCallbackSchedulerClass) Alloc() CPXCallbackScheduler {
	rv := objc.Send[CPXCallbackScheduler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXCallbackScheduler.DescheduleForceLogoutCallback]
//   - [CPXCallbackScheduler.DescheduleKillProcessCallback]
//   - [CPXCallbackScheduler.ScheduleFixBadForegroundCallbackForProcess]
//   - [CPXCallbackScheduler.ScheduleForceLogoutCallbackForTime]
//   - [CPXCallbackScheduler.ScheduleKillProcessCallbackForTime]
//   - [CPXCallbackScheduler.InitWithSession]
//   - [CPXCallbackScheduler.DebugDescription]
//   - [CPXCallbackScheduler.Description]
//   - [CPXCallbackScheduler.Hash]
//   - [CPXCallbackScheduler.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler
type CPXCallbackScheduler struct {
	objectivec.Object
}

// CPXCallbackSchedulerFromID constructs a [CPXCallbackScheduler] from an objc.ID.
func CPXCallbackSchedulerFromID(id objc.ID) CPXCallbackScheduler {
	return CPXCallbackScheduler{objectivec.Object{ID: id}}
}

// Ensure CPXCallbackScheduler implements ICPXCallbackScheduler.
var _ ICPXCallbackScheduler = CPXCallbackScheduler{}

// An interface definition for the [CPXCallbackScheduler] class.
//
// # Methods
//
//   - [ICPXCallbackScheduler.DescheduleForceLogoutCallback]
//   - [ICPXCallbackScheduler.DescheduleKillProcessCallback]
//   - [ICPXCallbackScheduler.ScheduleFixBadForegroundCallbackForProcess]
//   - [ICPXCallbackScheduler.ScheduleForceLogoutCallbackForTime]
//   - [ICPXCallbackScheduler.ScheduleKillProcessCallbackForTime]
//   - [ICPXCallbackScheduler.InitWithSession]
//   - [ICPXCallbackScheduler.DebugDescription]
//   - [ICPXCallbackScheduler.Description]
//   - [ICPXCallbackScheduler.Hash]
//   - [ICPXCallbackScheduler.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler
type ICPXCallbackScheduler interface {
	objectivec.IObject

	// Topic: Methods

	DescheduleForceLogoutCallback()
	DescheduleKillProcessCallback()
	ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRecRef)
	ScheduleForceLogoutCallbackForTime(time float64)
	ScheduleKillProcessCallbackForTime(time float64)
	InitWithSession(session unsafe.Pointer) CPXCallbackScheduler
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXCallbackScheduler) Init() CPXCallbackScheduler {
	rv := objc.Send[CPXCallbackScheduler](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXCallbackScheduler) Autorelease() CPXCallbackScheduler {
	rv := objc.Send[CPXCallbackScheduler](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXCallbackScheduler creates a new CPXCallbackScheduler instance.
func NewCPXCallbackScheduler() CPXCallbackScheduler {
	class := getCPXCallbackSchedulerClass()
	rv := objc.Send[CPXCallbackScheduler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/initWithSession:
func NewCPXCallbackSchedulerWithSession(session unsafe.Pointer) CPXCallbackScheduler {
	instance := getCPXCallbackSchedulerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return CPXCallbackSchedulerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/descheduleForceLogoutCallback
func (c CPXCallbackScheduler) DescheduleForceLogoutCallback() {
	objc.Send[objc.ID](c.ID, objc.Sel("descheduleForceLogoutCallback"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/descheduleKillProcessCallback
func (c CPXCallbackScheduler) DescheduleKillProcessCallback() {
	objc.Send[objc.ID](c.ID, objc.Sel("descheduleKillProcessCallback"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/scheduleFixBadForegroundCallbackForProcess:
func (c CPXCallbackScheduler) ScheduleFixBadForegroundCallbackForProcess(process *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("scheduleFixBadForegroundCallbackForProcess:"), process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/scheduleForceLogoutCallbackForTime:
func (c CPXCallbackScheduler) ScheduleForceLogoutCallbackForTime(time float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("scheduleForceLogoutCallbackForTime:"), time)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/scheduleKillProcessCallbackForTime:
func (c CPXCallbackScheduler) ScheduleKillProcessCallbackForTime(time float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("scheduleKillProcessCallbackForTime:"), time)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/initWithSession:
func (c CPXCallbackScheduler) InitWithSession(session unsafe.Pointer) CPXCallbackScheduler {
	rv := objc.Send[CPXCallbackScheduler](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/debugDescription
func (c CPXCallbackScheduler) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/description
func (c CPXCallbackScheduler) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/hash
func (c CPXCallbackScheduler) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXCallbackScheduler/superclass
func (c CPXCallbackScheduler) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
