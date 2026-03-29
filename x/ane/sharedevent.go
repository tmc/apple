//go:build darwin

package ane

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// SharedEvent wraps an IOSurfaceSharedEvent for ANE↔GPU/CPU synchronization.
type SharedEvent struct {
	ev iosurface.IOSurfaceSharedEvent
}

// NewSharedEvent creates a new shared event.
func NewSharedEvent() (*SharedEvent, error) {
	ev := iosurface.NewIOSurfaceSharedEvent()
	if ev.ID == 0 {
		return nil, fmt.Errorf("ane: failed to create shared event")
	}
	return &SharedEvent{ev: ev}, nil
}

// SharedEventFromPort creates a shared event bound to an existing Mach port.
func SharedEventFromPort(port uint32) (*SharedEvent, error) {
	ev := iosurface.NewIOSurfaceSharedEventWithMachPort(port)
	if ev.ID == 0 {
		return nil, fmt.Errorf("ane: failed to create shared event from port %d", port)
	}
	return &SharedEvent{ev: ev}, nil
}

// Port returns the Mach port name for this event.
func (e *SharedEvent) Port() uint32 { return e.ev.EventPort() }

// SignaledValue returns the current signaled value.
func (e *SharedEvent) SignaledValue() uint64 { return e.ev.SignaledValue() }

// Signal sets the event to the given value from the CPU.
func (e *SharedEvent) Signal(value uint64) { e.ev.SetSignaledValue(value) }

// Wait blocks until the event reaches the given value or timeout expires.
// Returns true if the event was signaled, false on timeout.
func (e *SharedEvent) Wait(value uint64, timeout time.Duration) bool {
	return e.ev.WaitUntilSignaledValueTimeoutMS(value, uint64(timeout.Milliseconds()))
}

// Close releases the shared event's underlying ObjC object.
func (e *SharedEvent) Close() error {
	if e.ev.ID != 0 {
		e.ev.Release()
		e.ev.ID = 0
	}
	return nil
}

// EvalWithSignalEvent evaluates the kernel and signals the given event on completion.
// The ANE hardware signals the event; no CPU wait is performed.
func (m *Model) EvalWithSignalEvent(signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	signal, err := SharedEventFromPort(signalPort)
	if err != nil {
		return err
	}
	defer signal.Close()
	return m.EvalWithSignal(signal, signalValue, cfg)
}

// EvalWithWaitEvent evaluates the kernel after waiting for the given event.
// The ANE waits for waitValue on waitPort before executing.
func (m *Model) EvalWithWaitEvent(waitPort uint32, waitValue uint64, cfg SharedEventEvalOptions) error {
	wait, err := SharedEventFromPort(waitPort)
	if err != nil {
		return err
	}
	defer wait.Close()
	return m.EvalWithWait(wait, waitValue, cfg)
}

// EvalBidirectional evaluates the kernel with both wait and signal events.
// The ANE waits for waitValue on waitPort before executing, then signals
// signalValue on signalPort after completion.
func (m *Model) EvalBidirectional(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	wait, err := SharedEventFromPort(waitPort)
	if err != nil {
		return err
	}
	defer wait.Close()
	signal, err := SharedEventFromPort(signalPort)
	if err != nil {
		return err
	}
	defer signal.Close()
	return m.EvalBidirectionalEvents(wait, waitValue, signal, signalValue, cfg)
}

const eventTypeSignal = int64(5)

// EvalWithSignal evaluates the kernel and signals the given shared event.
func (m *Model) EvalWithSignal(signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	return m.evalWithSharedEvents(nil, 0, signal, signalValue, cfg)
}

// EvalWithWait evaluates the kernel after waiting on the given shared event.
func (m *Model) EvalWithWait(wait *SharedEvent, waitValue uint64, cfg SharedEventEvalOptions) error {
	return m.evalWithSharedEvents(wait, waitValue, nil, 0, cfg)
}

// EvalBidirectionalEvents evaluates the kernel with wait and signal shared events.
func (m *Model) EvalBidirectionalEvents(wait *SharedEvent, waitValue uint64, signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	return m.evalWithSharedEvents(wait, waitValue, signal, signalValue, cfg)
}

func (m *Model) evalWithSharedEvents(wait *SharedEvent, waitValue uint64, signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	m.mu.Lock()
	m.sharedEventUsed = true
	m.mu.Unlock()

	signalArr := foundation.NewNSMutableArray()
	if signal != nil {
		symbolIndex := 0
		if len(m.outputLayouts) > 0 && m.outputLayouts[0].SymbolIndex >= 0 {
			symbolIndex = m.outputLayouts[0].SymbolIndex
		}
		sigEvtClass := appleneuralengine.GetANESharedSignalEventClass()
		sigEvtObj := sigEvtClass.SignalEventWithValueSymbolIndexEventTypeSharedEvent(
			signalValue, uint32(symbolIndex), eventTypeSignal, signal.ev,
		)
		if sigEvtObj == nil || sigEvtObj.GetID() == 0 {
			return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create ANESharedSignalEvent")}
		}
		signalArr.AddObject(sigEvtObj)
	}

	waitArr := foundation.NewNSMutableArray()
	var waitEvtObj objectivec.IObject
	if wait != nil {
		waitEvtClass := appleneuralengine.GetANESharedWaitEventClass()
		waitEvtObj = waitEvtClass.WaitEventWithValueSharedEvent(waitValue, wait.ev)
		if waitEvtObj == nil || waitEvtObj.GetID() == 0 {
			return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create ANESharedWaitEvent")}
		}
		waitArr.AddObject(waitEvtObj)
	}

	// Build shared events container.
	sharedEventsClass := appleneuralengine.GetANESharedEventsClass()
	sharedEventsObj := sharedEventsClass.SharedEventsWithSignalEventsWaitEvents(signalArr, waitArr)
	if sharedEventsObj == nil || sharedEventsObj.GetID() == 0 {
		return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create ANESharedEvents")}
	}
	req, reqObjects, err := createRequestFromBindingsWithSharedEvents(m.inputBindings(), m.outputBindings(), 0, sharedEventsObj)
	if err != nil {
		return err
	}
	keepAlive := []objectivec.IObject{signalArr, sharedEventsObj}
	keepAlive = append(keepAlive, reqObjects...)
	if signal != nil {
		keepAlive = append(keepAlive, signal.ev)
	}
	if wait != nil {
		keepAlive = append(keepAlive, wait.ev)
	}
	if wait != nil {
		keepAlive = append(keepAlive, waitEvtObj, waitArr)
	}
	defer func() {
		runtime.KeepAlive(keepAlive)
		runtime.KeepAlive(req)
	}()

	opts := foundation.NewNSMutableDictionary()
	if cfg.DisableIOFencesUseSharedEvents {
		opts.SetObjectForKey(
			foundation.GetNSNumberClass().NumberWithBool(true),
			nsStringKey("kANEFDisableIOFencesUseSharedEventsKey"),
		)
	}
	if cfg.EnableFWToFWSignal {
		opts.SetObjectForKey(
			foundation.GetNSNumberClass().NumberWithBool(true),
			nsStringKey("kANEFEnableFWToFWSignal"),
		)
	}

	// Validate and evaluate.
	if !req.Validate() {
		return &ANEError{Op: "eval", Err: fmt.Errorf("request validation failed")}
	}

	unmap, err := m.mapRequestWithFallback(req)
	if err != nil {
		return wrapErr("map", err)
	}
	defer unmap()

	preferDirect := os.Getenv("ANE_SHARED_USE_DIRECT") != ""
	return wrapErr("eval", m.evaluateRequestWithOptions(req, opts, preferDirect))
}

// TimeWait blocks until the event reaches the given value or timeout expires,
// returning the signaled status and the wall-clock duration of the wait.
func (e *SharedEvent) TimeWait(value uint64, timeout time.Duration) (bool, time.Duration) {
	start := time.Now()
	ok := e.Wait(value, timeout)
	return ok, time.Since(start)
}

// nsStringKey creates an NSCopying-compatible key from a string for use with NSDictionary.
func nsStringKey(s string) foundation.NSCopyingObject {
	nsStr := foundation.NewStringWithString(s)
	return foundation.NSCopyingObject{Object: nsStr.Object}
}
