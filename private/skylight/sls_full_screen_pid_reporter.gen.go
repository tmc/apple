// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSFullScreenPidReporter] class.
var (
	_SLSFullScreenPidReporterClass     SLSFullScreenPidReporterClass
	_SLSFullScreenPidReporterClassOnce sync.Once
)

func getSLSFullScreenPidReporterClass() SLSFullScreenPidReporterClass {
	_SLSFullScreenPidReporterClassOnce.Do(func() {
		_SLSFullScreenPidReporterClass = SLSFullScreenPidReporterClass{class: objc.GetClass("SLSFullScreenPidReporter")}
	})
	return _SLSFullScreenPidReporterClass
}

// GetSLSFullScreenPidReporterClass returns the class object for SLSFullScreenPidReporter.
func GetSLSFullScreenPidReporterClass() SLSFullScreenPidReporterClass {
	return getSLSFullScreenPidReporterClass()
}

type SLSFullScreenPidReporterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSFullScreenPidReporterClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSFullScreenPidReporterClass) Alloc() SLSFullScreenPidReporter {
	rv := objc.Send[SLSFullScreenPidReporter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSFullScreenPidReporter.EqualCurrentSeed]
//   - [SLSFullScreenPidReporter.HandleConnectionInterrupt]
//   - [SLSFullScreenPidReporter.IncrementSeed]
//   - [SLSFullScreenPidReporter.ReceiveMessages]
//   - [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [SLSFullScreenPidReporter.SetDisconnectHandler]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter
type SLSFullScreenPidReporter struct {
	objectivec.Object
}

// SLSFullScreenPidReporterFromID constructs a [SLSFullScreenPidReporter] from an objc.ID.
func SLSFullScreenPidReporterFromID(id objc.ID) SLSFullScreenPidReporter {
	return SLSFullScreenPidReporter{objectivec.Object{ID: id}}
}

// Ensure SLSFullScreenPidReporter implements ISLSFullScreenPidReporter.
var _ ISLSFullScreenPidReporter = SLSFullScreenPidReporter{}

// An interface definition for the [SLSFullScreenPidReporter] class.
//
// # Methods
//
//   - [ISLSFullScreenPidReporter.EqualCurrentSeed]
//   - [ISLSFullScreenPidReporter.HandleConnectionInterrupt]
//   - [ISLSFullScreenPidReporter.IncrementSeed]
//   - [ISLSFullScreenPidReporter.ReceiveMessages]
//   - [ISLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [ISLSFullScreenPidReporter.SetDisconnectHandler]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter
type ISLSFullScreenPidReporter interface {
	objectivec.IObject

	// Topic: Methods

	EqualCurrentSeed(seed uint64) bool
	HandleConnectionInterrupt()
	IncrementSeed() uint64
	ReceiveMessages(messages objectivec.IObject)
	ReportFullScreenStatusWithFilterAndHandler(filter objectivec.IObject, handler VoidHandler)
	SetDisconnectHandler(handler VoidHandler)
}

// Init initializes the instance.
func (s SLSFullScreenPidReporter) Init() SLSFullScreenPidReporter {
	rv := objc.Send[SLSFullScreenPidReporter](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSFullScreenPidReporter) Autorelease() SLSFullScreenPidReporter {
	rv := objc.Send[SLSFullScreenPidReporter](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSFullScreenPidReporter creates a new SLSFullScreenPidReporter instance.
func NewSLSFullScreenPidReporter() SLSFullScreenPidReporter {
	class := getSLSFullScreenPidReporterClass()
	rv := objc.Send[SLSFullScreenPidReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/equalCurrentSeed:
func (s SLSFullScreenPidReporter) EqualCurrentSeed(seed uint64) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("equalCurrentSeed:"), seed)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/handleConnectionInterrupt
func (s SLSFullScreenPidReporter) HandleConnectionInterrupt() {
	objc.Send[objc.ID](s.ID, objc.Sel("handleConnectionInterrupt"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/incrementSeed
func (s SLSFullScreenPidReporter) IncrementSeed() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("incrementSeed"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/receiveMessages:
func (s SLSFullScreenPidReporter) ReceiveMessages(messages objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("receiveMessages:"), messages)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/reportFullScreenStatusWithFilter:andHandler:
func (s SLSFullScreenPidReporter) ReportFullScreenStatusWithFilterAndHandler(filter objectivec.IObject, handler VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("reportFullScreenStatusWithFilter:andHandler:"), filter, _block1)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/setDisconnectHandler:
func (s SLSFullScreenPidReporter) SetDisconnectHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("setDisconnectHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSFullScreenPidReporter/sharedReporter
func (_SLSFullScreenPidReporterClass SLSFullScreenPidReporterClass) SharedReporter() SLSFullScreenPidReporter {
	rv := objc.Send[objc.ID](objc.ID(_SLSFullScreenPidReporterClass.class), objc.Sel("sharedReporter"))
	return SLSFullScreenPidReporterFromID(rv)
}

// ReportFullScreenStatusWithFilterAndHandlerSync is a synchronous wrapper around [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSFullScreenPidReporter) ReportFullScreenStatusWithFilterAndHandlerSync(ctx context.Context, filter objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.ReportFullScreenStatusWithFilterAndHandler(filter, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetDisconnectHandlerSync is a synchronous wrapper around [SLSFullScreenPidReporter.SetDisconnectHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSFullScreenPidReporter) SetDisconnectHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetDisconnectHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
