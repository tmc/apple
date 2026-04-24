// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXEventDeferringPolicySanitizer] class.
var (
	_CPXEventDeferringPolicySanitizerClass     CPXEventDeferringPolicySanitizerClass
	_CPXEventDeferringPolicySanitizerClassOnce sync.Once
)

func getCPXEventDeferringPolicySanitizerClass() CPXEventDeferringPolicySanitizerClass {
	_CPXEventDeferringPolicySanitizerClassOnce.Do(func() {
		_CPXEventDeferringPolicySanitizerClass = CPXEventDeferringPolicySanitizerClass{class: objc.GetClass("CPXEventDeferringPolicySanitizer")}
	})
	return _CPXEventDeferringPolicySanitizerClass
}

// GetCPXEventDeferringPolicySanitizerClass returns the class object for CPXEventDeferringPolicySanitizer.
func GetCPXEventDeferringPolicySanitizerClass() CPXEventDeferringPolicySanitizerClass {
	return getCPXEventDeferringPolicySanitizerClass()
}

type CPXEventDeferringPolicySanitizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXEventDeferringPolicySanitizerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXEventDeferringPolicySanitizerClass) Alloc() CPXEventDeferringPolicySanitizer {
	rv := objc.Send[CPXEventDeferringPolicySanitizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXEventDeferringPolicySanitizer._isValidProcessAuditHistoryDebugProcessType]
//   - [CPXEventDeferringPolicySanitizer._sanitizeFrontmost]
//   - [CPXEventDeferringPolicySanitizer._sanitizeKeyThief]
//   - [CPXEventDeferringPolicySanitizer.Sanitize]
//   - [CPXEventDeferringPolicySanitizer.InitWithFocusManagerDataSourceProcessManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer
type CPXEventDeferringPolicySanitizer struct {
	objectivec.Object
}

// CPXEventDeferringPolicySanitizerFromID constructs a [CPXEventDeferringPolicySanitizer] from an objc.ID.
func CPXEventDeferringPolicySanitizerFromID(id objc.ID) CPXEventDeferringPolicySanitizer {
	return CPXEventDeferringPolicySanitizer{objectivec.Object{ID: id}}
}

// Ensure CPXEventDeferringPolicySanitizer implements ICPXEventDeferringPolicySanitizer.
var _ ICPXEventDeferringPolicySanitizer = CPXEventDeferringPolicySanitizer{}

// An interface definition for the [CPXEventDeferringPolicySanitizer] class.
//
// # Methods
//
//   - [ICPXEventDeferringPolicySanitizer._isValidProcessAuditHistoryDebugProcessType]
//   - [ICPXEventDeferringPolicySanitizer._sanitizeFrontmost]
//   - [ICPXEventDeferringPolicySanitizer._sanitizeKeyThief]
//   - [ICPXEventDeferringPolicySanitizer.Sanitize]
//   - [ICPXEventDeferringPolicySanitizer.InitWithFocusManagerDataSourceProcessManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer
type ICPXEventDeferringPolicySanitizer interface {
	objectivec.IObject

	// Topic: Methods

	_isValidProcessAuditHistoryDebugProcessType(process *CPSProcessRecRef, history objectivec.IObject, type_ objectivec.IObject) bool
	_sanitizeFrontmost(frontmost objectivec.IObject)
	_sanitizeKeyThief(thief objectivec.IObject)
	Sanitize(sanitize objectivec.IObject) objectivec.IObject
	InitWithFocusManagerDataSourceProcessManager(source objectivec.IObject, manager objectivec.IObject) CPXEventDeferringPolicySanitizer
}

// Init initializes the instance.
func (c CPXEventDeferringPolicySanitizer) Init() CPXEventDeferringPolicySanitizer {
	rv := objc.Send[CPXEventDeferringPolicySanitizer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXEventDeferringPolicySanitizer) Autorelease() CPXEventDeferringPolicySanitizer {
	rv := objc.Send[CPXEventDeferringPolicySanitizer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXEventDeferringPolicySanitizer creates a new CPXEventDeferringPolicySanitizer instance.
func NewCPXEventDeferringPolicySanitizer() CPXEventDeferringPolicySanitizer {
	class := getCPXEventDeferringPolicySanitizerClass()
	rv := objc.Send[CPXEventDeferringPolicySanitizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/initWithFocusManagerDataSource:processManager:
func NewCPXEventDeferringPolicySanitizerWithFocusManagerDataSourceProcessManager(source objectivec.IObject, manager objectivec.IObject) CPXEventDeferringPolicySanitizer {
	instance := getCPXEventDeferringPolicySanitizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFocusManagerDataSource:processManager:"), source, manager)
	return CPXEventDeferringPolicySanitizerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/_isValidProcess:auditHistory:debugProcessType:
func (c CPXEventDeferringPolicySanitizer) _isValidProcessAuditHistoryDebugProcessType(process *CPSProcessRecRef, history objectivec.IObject, type_ objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("_isValidProcess:auditHistory:debugProcessType:"), process, history, type_)
	return rv
}

// IsValidProcessAuditHistoryDebugProcessType is an exported wrapper for the private method _isValidProcessAuditHistoryDebugProcessType.
func (c CPXEventDeferringPolicySanitizer) IsValidProcessAuditHistoryDebugProcessType(process *CPSProcessRecRef, history objectivec.IObject, type_ objectivec.IObject) bool {
	return c._isValidProcessAuditHistoryDebugProcessType(process, history, type_)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/_sanitizeFrontmost:
func (c CPXEventDeferringPolicySanitizer) _sanitizeFrontmost(frontmost objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_sanitizeFrontmost:"), frontmost)
}

// SanitizeFrontmost is an exported wrapper for the private method _sanitizeFrontmost.
func (c CPXEventDeferringPolicySanitizer) SanitizeFrontmost(frontmost objectivec.IObject) {
	c._sanitizeFrontmost(frontmost)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/_sanitizeKeyThief:
func (c CPXEventDeferringPolicySanitizer) _sanitizeKeyThief(thief objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_sanitizeKeyThief:"), thief)
}

// SanitizeKeyThief is an exported wrapper for the private method _sanitizeKeyThief.
func (c CPXEventDeferringPolicySanitizer) SanitizeKeyThief(thief objectivec.IObject) {
	c._sanitizeKeyThief(thief)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/sanitize:
func (c CPXEventDeferringPolicySanitizer) Sanitize(sanitize objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sanitize:"), sanitize)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicySanitizer/initWithFocusManagerDataSource:processManager:
func (c CPXEventDeferringPolicySanitizer) InitWithFocusManagerDataSourceProcessManager(source objectivec.IObject, manager objectivec.IObject) CPXEventDeferringPolicySanitizer {
	rv := objc.Send[CPXEventDeferringPolicySanitizer](c.ID, objc.Sel("initWithFocusManagerDataSource:processManager:"), source, manager)
	return rv
}
