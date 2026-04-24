// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXEventDeferringPolicy] class.
var (
	_CPXEventDeferringPolicyClass     CPXEventDeferringPolicyClass
	_CPXEventDeferringPolicyClassOnce sync.Once
)

func getCPXEventDeferringPolicyClass() CPXEventDeferringPolicyClass {
	_CPXEventDeferringPolicyClassOnce.Do(func() {
		_CPXEventDeferringPolicyClass = CPXEventDeferringPolicyClass{class: objc.GetClass("CPXEventDeferringPolicy")}
	})
	return _CPXEventDeferringPolicyClass
}

// GetCPXEventDeferringPolicyClass returns the class object for CPXEventDeferringPolicy.
func GetCPXEventDeferringPolicyClass() CPXEventDeferringPolicyClass {
	return getCPXEventDeferringPolicyClass()
}

type CPXEventDeferringPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXEventDeferringPolicyClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXEventDeferringPolicyClass) Alloc() CPXEventDeferringPolicy {
	rv := objc.Send[CPXEventDeferringPolicy](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXEventDeferringPolicy._init]
//   - [CPXEventDeferringPolicy._initWithCopyOf]
//   - [CPXEventDeferringPolicy.AdvicePolicy]
//   - [CPXEventDeferringPolicy.AppendDescriptionToStream]
//   - [CPXEventDeferringPolicy.AuditHistory]
//   - [CPXEventDeferringPolicy.FrontmostProcess]
//   - [CPXEventDeferringPolicy.KeyThiefConnectionID]
//   - [CPXEventDeferringPolicy.SetKeyThiefConnectionID]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy
type CPXEventDeferringPolicy struct {
	objectivec.Object
}

// CPXEventDeferringPolicyFromID constructs a [CPXEventDeferringPolicy] from an objc.ID.
func CPXEventDeferringPolicyFromID(id objc.ID) CPXEventDeferringPolicy {
	return CPXEventDeferringPolicy{objectivec.Object{ID: id}}
}

// Ensure CPXEventDeferringPolicy implements ICPXEventDeferringPolicy.
var _ ICPXEventDeferringPolicy = CPXEventDeferringPolicy{}

// An interface definition for the [CPXEventDeferringPolicy] class.
//
// # Methods
//
//   - [ICPXEventDeferringPolicy._init]
//   - [ICPXEventDeferringPolicy._initWithCopyOf]
//   - [ICPXEventDeferringPolicy.AdvicePolicy]
//   - [ICPXEventDeferringPolicy.AppendDescriptionToStream]
//   - [ICPXEventDeferringPolicy.AuditHistory]
//   - [ICPXEventDeferringPolicy.FrontmostProcess]
//   - [ICPXEventDeferringPolicy.KeyThiefConnectionID]
//   - [ICPXEventDeferringPolicy.SetKeyThiefConnectionID]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy
type ICPXEventDeferringPolicy interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_initWithCopyOf(of objectivec.IObject) objectivec.IObject
	AdvicePolicy() int64
	AppendDescriptionToStream(stream objectivec.IObject)
	AuditHistory() unsafe.Pointer
	FrontmostProcess() *CPSProcessRecRef
	KeyThiefConnectionID() uint32
	SetKeyThiefConnectionID(value uint32)
}

// Init initializes the instance.
func (c CPXEventDeferringPolicy) Init() CPXEventDeferringPolicy {
	rv := objc.Send[CPXEventDeferringPolicy](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXEventDeferringPolicy) Autorelease() CPXEventDeferringPolicy {
	rv := objc.Send[CPXEventDeferringPolicy](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXEventDeferringPolicy creates a new CPXEventDeferringPolicy instance.
func NewCPXEventDeferringPolicy() CPXEventDeferringPolicy {
	class := getCPXEventDeferringPolicyClass()
	rv := objc.Send[CPXEventDeferringPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/_init
func (c CPXEventDeferringPolicy) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/_initWithCopyOf:
func (c CPXEventDeferringPolicy) _initWithCopyOf(of objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_initWithCopyOf:"), of)
	return objectivec.Object{ID: rv}
}

// InitWithCopyOf is an exported wrapper for the private method _initWithCopyOf.
func (c CPXEventDeferringPolicy) InitWithCopyOf(of objectivec.IObject) objectivec.IObject {
	return c._initWithCopyOf(of)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/appendDescriptionToStream:
func (c CPXEventDeferringPolicy) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/build:
func (_CPXEventDeferringPolicyClass CPXEventDeferringPolicyClass) Build(build VoidHandler) objectivec.IObject {
	_block0, _ := NewVoidBlock(build)
	rv := objc.Send[objc.ID](objc.ID(_CPXEventDeferringPolicyClass.class), objc.Sel("build:"), _block0)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/advicePolicy
func (c CPXEventDeferringPolicy) AdvicePolicy() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("advicePolicy"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/auditHistory
func (c CPXEventDeferringPolicy) AuditHistory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("auditHistory"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/frontmostProcess
func (c CPXEventDeferringPolicy) FrontmostProcess() *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontmostProcess"))
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringPolicy/keyThiefConnectionID
func (c CPXEventDeferringPolicy) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (c CPXEventDeferringPolicy) SetKeyThiefConnectionID(value uint32) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeyThiefConnectionID:"), value)
}

// BuildSync is a synchronous wrapper around [CPXEventDeferringPolicy.Build].
// It blocks until the completion handler fires or the context is cancelled.
func (cc CPXEventDeferringPolicyClass) BuildSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	cc.Build(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
