// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLScreenTelemetryConnection] class.
var (
	_SLScreenTelemetryConnectionClass     SLScreenTelemetryConnectionClass
	_SLScreenTelemetryConnectionClassOnce sync.Once
)

func getSLScreenTelemetryConnectionClass() SLScreenTelemetryConnectionClass {
	_SLScreenTelemetryConnectionClassOnce.Do(func() {
		_SLScreenTelemetryConnectionClass = SLScreenTelemetryConnectionClass{class: objc.GetClass("SLScreenTelemetryConnection")}
	})
	return _SLScreenTelemetryConnectionClass
}

// GetSLScreenTelemetryConnectionClass returns the class object for SLScreenTelemetryConnection.
func GetSLScreenTelemetryConnectionClass() SLScreenTelemetryConnectionClass {
	return getSLScreenTelemetryConnectionClass()
}

type SLScreenTelemetryConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLScreenTelemetryConnectionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLScreenTelemetryConnectionClass) Alloc() SLScreenTelemetryConnection {
	rv := objc.SendIfResponds[SLScreenTelemetryConnection](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLScreenTelemetryConnection.Close]
//   - [SLScreenTelemetryConnection.Closed]
//   - [SLScreenTelemetryConnection.SetClosed]
//   - [SLScreenTelemetryConnection.ClosedWithError]
//   - [SLScreenTelemetryConnection.ConnectUsingXPCConnectionAndConfigMessage]
//   - [SLScreenTelemetryConnection.Connection]
//   - [SLScreenTelemetryConnection.SetConnection]
//   - [SLScreenTelemetryConnection.HandleMessage]
//   - [SLScreenTelemetryConnection.Queue]
//   - [SLScreenTelemetryConnection.UpdateBlock]
//   - [SLScreenTelemetryConnection.ZeroingContainer]
//   - [SLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
type SLScreenTelemetryConnection struct {
	objectivec.Object
}

// SLScreenTelemetryConnectionFromID constructs a [SLScreenTelemetryConnection] from an objc.ID.
func SLScreenTelemetryConnectionFromID(id objc.ID) SLScreenTelemetryConnection {
	return SLScreenTelemetryConnection{objectivec.Object{ID: id}}
}

// Ensure SLScreenTelemetryConnection implements ISLScreenTelemetryConnection.
var _ ISLScreenTelemetryConnection = SLScreenTelemetryConnection{}

// An interface definition for the [SLScreenTelemetryConnection] class.
//
// # Methods
//
//   - [ISLScreenTelemetryConnection.Close]
//   - [ISLScreenTelemetryConnection.Closed]
//   - [ISLScreenTelemetryConnection.SetClosed]
//   - [ISLScreenTelemetryConnection.ClosedWithError]
//   - [ISLScreenTelemetryConnection.ConnectUsingXPCConnectionAndConfigMessage]
//   - [ISLScreenTelemetryConnection.Connection]
//   - [ISLScreenTelemetryConnection.SetConnection]
//   - [ISLScreenTelemetryConnection.HandleMessage]
//   - [ISLScreenTelemetryConnection.Queue]
//   - [ISLScreenTelemetryConnection.UpdateBlock]
//   - [ISLScreenTelemetryConnection.ZeroingContainer]
//   - [ISLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
type ISLScreenTelemetryConnection interface {
	objectivec.IObject

	// Topic: Methods

	Close()
	Closed() bool
	SetClosed(value bool)
	ClosedWithError(error_ objectivec.IObject)
	ConnectUsingXPCConnectionAndConfigMessage(xPCConnection objectivec.IObject, message objectivec.IObject)
	Connection() objectivec.Object
	SetConnection(value objectivec.Object)
	HandleMessage(message objectivec.IObject)
	Queue() objectivec.Object
	UpdateBlock() VoidHandler
	ZeroingContainer() ISLSZeroingWeakContainer
	InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock(width uint32, height uint32, rows uint32, columns uint32, interval float64, queue objectivec.IObject, block VoidHandler) SLScreenTelemetryConnection
}

// Init initializes the instance.
func (s SLScreenTelemetryConnection) Init() SLScreenTelemetryConnection {
	rv := objc.SendIfResponds[SLScreenTelemetryConnection](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLScreenTelemetryConnection) Autorelease() SLScreenTelemetryConnection {
	rv := objc.SendIfResponds[SLScreenTelemetryConnection](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLScreenTelemetryConnection creates a new SLScreenTelemetryConnection instance.
func NewSLScreenTelemetryConnection() SLScreenTelemetryConnection {
	class := getSLScreenTelemetryConnectionClass()
	rv := objc.SendIfResponds[SLScreenTelemetryConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SLScreenTelemetryConnection) Close() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("close"))
}
func (s SLScreenTelemetryConnection) ClosedWithError(error_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("closedWithError:"), error_)
}
func (s SLScreenTelemetryConnection) ConnectUsingXPCConnectionAndConfigMessage(xPCConnection objectivec.IObject, message objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("connectUsingXPCConnection:andConfigMessage:"), xPCConnection, message)
}
func (s SLScreenTelemetryConnection) HandleMessage(message objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("handleMessage:"), message)
}
func (s SLScreenTelemetryConnection) InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock(width uint32, height uint32, rows uint32, columns uint32, interval float64, queue objectivec.IObject, block VoidHandler) SLScreenTelemetryConnection {
	_block6, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[SLScreenTelemetryConnection](s.ID, objc.Sel("initWithZoneWidth:zoneHeight:zoneRows:zoneColumns:samplingInterval:queue:andUpdateBlock:"), width, height, rows, columns, interval, queue, _block6)
	return rv
}

func (_SLScreenTelemetryConnectionClass SLScreenTelemetryConnectionClass) ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock(width uint32, height uint32, rows uint32, columns uint32, interval float64, queue objectivec.IObject, block VoidHandler) objectivec.IObject {
	_block6, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLScreenTelemetryConnectionClass.class), objc.Sel("connectionWithZoneWidth:zoneHeight:zoneRows:zoneColumns:samplingInterval:queue:andUpdateBlock:"), width, height, rows, columns, interval, queue, _block6)
	return objectivec.Object{ID: rv}
}

func (s SLScreenTelemetryConnection) Closed() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("closed"))
	return rv
}
func (s SLScreenTelemetryConnection) SetClosed(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setClosed:"), value)
}
func (s SLScreenTelemetryConnection) Connection() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("connection"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLScreenTelemetryConnection) SetConnection(value objectivec.Object) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setConnection:"), value)
}
func (s SLScreenTelemetryConnection) Queue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLScreenTelemetryConnection) UpdateBlock() VoidHandler {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("updateBlock"))
	_ = rv
	return nil
}
func (s SLScreenTelemetryConnection) ZeroingContainer() ISLSZeroingWeakContainer {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("zeroingContainer"))
	return SLSZeroingWeakContainerFromID(objc.ID(rv))
}

// InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlockSync is a synchronous wrapper around [SLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLScreenTelemetryConnection) InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlockSync(ctx context.Context, width uint32, height uint32, rows uint32, columns uint32, interval float64, queue objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock(width, height, rows, columns, interval, queue, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlockSync is a synchronous wrapper around [SLScreenTelemetryConnection.ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SLScreenTelemetryConnectionClass) ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlockSync(ctx context.Context, width uint32, height uint32, rows uint32, columns uint32, interval float64, queue objectivec.IObject) error {
	done := make(chan struct{}, 1)
	sc.ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock(width, height, rows, columns, interval, queue, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
