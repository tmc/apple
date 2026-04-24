// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWindowManagementFallbackBridge] class.
var (
	_SLSWindowManagementFallbackBridgeClass     SLSWindowManagementFallbackBridgeClass
	_SLSWindowManagementFallbackBridgeClassOnce sync.Once
)

func getSLSWindowManagementFallbackBridgeClass() SLSWindowManagementFallbackBridgeClass {
	_SLSWindowManagementFallbackBridgeClassOnce.Do(func() {
		_SLSWindowManagementFallbackBridgeClass = SLSWindowManagementFallbackBridgeClass{class: objc.GetClass("SLSWindowManagementFallbackBridge")}
	})
	return _SLSWindowManagementFallbackBridgeClass
}

// GetSLSWindowManagementFallbackBridgeClass returns the class object for SLSWindowManagementFallbackBridge.
func GetSLSWindowManagementFallbackBridgeClass() SLSWindowManagementFallbackBridgeClass {
	return getSLSWindowManagementFallbackBridgeClass()
}

type SLSWindowManagementFallbackBridgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWindowManagementFallbackBridgeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWindowManagementFallbackBridgeClass) Alloc() SLSWindowManagementFallbackBridge {
	rv := objc.Send[SLSWindowManagementFallbackBridge](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWindowManagementFallbackBridge.PerformAsynchronousBridgedWindowManagementOperation]
//   - [SLSWindowManagementFallbackBridge.PerformSynchronousBridgedWindowManagementOperation]
//   - [SLSWindowManagementFallbackBridge.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [SLSWindowManagementFallbackBridge.DebugDescription]
//   - [SLSWindowManagementFallbackBridge.Description]
//   - [SLSWindowManagementFallbackBridge.Hash]
//   - [SLSWindowManagementFallbackBridge.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge
type SLSWindowManagementFallbackBridge struct {
	objectivec.Object
}

// SLSWindowManagementFallbackBridgeFromID constructs a [SLSWindowManagementFallbackBridge] from an objc.ID.
func SLSWindowManagementFallbackBridgeFromID(id objc.ID) SLSWindowManagementFallbackBridge {
	return SLSWindowManagementFallbackBridge{objectivec.Object{ID: id}}
}

// Ensure SLSWindowManagementFallbackBridge implements ISLSWindowManagementFallbackBridge.
var _ ISLSWindowManagementFallbackBridge = SLSWindowManagementFallbackBridge{}

// An interface definition for the [SLSWindowManagementFallbackBridge] class.
//
// # Methods
//
//   - [ISLSWindowManagementFallbackBridge.PerformAsynchronousBridgedWindowManagementOperation]
//   - [ISLSWindowManagementFallbackBridge.PerformSynchronousBridgedWindowManagementOperation]
//   - [ISLSWindowManagementFallbackBridge.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [ISLSWindowManagementFallbackBridge.DebugDescription]
//   - [ISLSWindowManagementFallbackBridge.Description]
//   - [ISLSWindowManagementFallbackBridge.Hash]
//   - [ISLSWindowManagementFallbackBridge.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge
type ISLSWindowManagementFallbackBridge interface {
	objectivec.IObject

	// Topic: Methods

	PerformAsynchronousBridgedWindowManagementOperation(operation objectivec.IObject)
	PerformSynchronousBridgedWindowManagementOperation(operation objectivec.IObject) objectivec.IObject
	PerformWindowManagementBridgeTransactionUsingBlock(block VoidHandler)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLSWindowManagementFallbackBridge) Init() SLSWindowManagementFallbackBridge {
	rv := objc.Send[SLSWindowManagementFallbackBridge](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWindowManagementFallbackBridge) Autorelease() SLSWindowManagementFallbackBridge {
	rv := objc.Send[SLSWindowManagementFallbackBridge](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWindowManagementFallbackBridge creates a new SLSWindowManagementFallbackBridge instance.
func NewSLSWindowManagementFallbackBridge() SLSWindowManagementFallbackBridge {
	class := getSLSWindowManagementFallbackBridgeClass()
	rv := objc.Send[SLSWindowManagementFallbackBridge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/performAsynchronousBridgedWindowManagementOperation:
func (s SLSWindowManagementFallbackBridge) PerformAsynchronousBridgedWindowManagementOperation(operation objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("performAsynchronousBridgedWindowManagementOperation:"), operation)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/performSynchronousBridgedWindowManagementOperation:
func (s SLSWindowManagementFallbackBridge) PerformSynchronousBridgedWindowManagementOperation(operation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("performSynchronousBridgedWindowManagementOperation:"), operation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/performWindowManagementBridgeTransactionUsingBlock:
func (s SLSWindowManagementFallbackBridge) PerformWindowManagementBridgeTransactionUsingBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("performWindowManagementBridgeTransactionUsingBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/debugDescription
func (s SLSWindowManagementFallbackBridge) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/description
func (s SLSWindowManagementFallbackBridge) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/hash
func (s SLSWindowManagementFallbackBridge) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementFallbackBridge/superclass
func (s SLSWindowManagementFallbackBridge) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

// PerformWindowManagementBridgeTransactionUsingBlockSync is a synchronous wrapper around [SLSWindowManagementFallbackBridge.PerformWindowManagementBridgeTransactionUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSWindowManagementFallbackBridge) PerformWindowManagementBridgeTransactionUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.PerformWindowManagementBridgeTransactionUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
