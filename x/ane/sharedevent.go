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

// Close releases the shared event.
func (e *SharedEvent) Close() error {
	return nil
}

// EvalWithSignalEvent evaluates the kernel and signals the given event on completion.
// The ANE hardware signals the event; no CPU wait is performed.
func (k *Kernel) EvalWithSignalEvent(signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	signal, err := SharedEventFromPort(signalPort)
	if err != nil {
		return err
	}
	return k.EvalWithSignal(signal, signalValue, cfg)
}

// EvalWithWaitEvent evaluates the kernel after waiting for the given event.
// The ANE waits for waitValue on waitPort before executing.
func (k *Kernel) EvalWithWaitEvent(waitPort uint32, waitValue uint64, cfg SharedEventEvalOptions) error {
	wait, err := SharedEventFromPort(waitPort)
	if err != nil {
		return err
	}
	return k.EvalWithWait(wait, waitValue, cfg)
}

// EvalBidirectional evaluates the kernel with both wait and signal events.
// The ANE waits for waitValue on waitPort before executing, then signals
// signalValue on signalPort after completion.
func (k *Kernel) EvalBidirectional(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	wait, err := SharedEventFromPort(waitPort)
	if err != nil {
		return err
	}
	signal, err := SharedEventFromPort(signalPort)
	if err != nil {
		return err
	}
	return k.EvalBidirectionalEvents(wait, waitValue, signal, signalValue, cfg)
}

const eventTypeSignal = int64(5)

// EvalWithSignal evaluates the kernel and signals the given shared event.
func (k *Kernel) EvalWithSignal(signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	return k.evalWithSharedEvents(nil, 0, signal, signalValue, cfg)
}

// EvalWithWait evaluates the kernel after waiting on the given shared event.
func (k *Kernel) EvalWithWait(wait *SharedEvent, waitValue uint64, cfg SharedEventEvalOptions) error {
	return k.evalWithSharedEvents(wait, waitValue, nil, 0, cfg)
}

// EvalBidirectionalEvents evaluates the kernel with wait and signal shared events.
func (k *Kernel) EvalBidirectionalEvents(wait *SharedEvent, waitValue uint64, signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	return k.evalWithSharedEvents(wait, waitValue, signal, signalValue, cfg)
}

func (k *Kernel) evalWithSharedEvents(wait *SharedEvent, waitValue uint64, signal *SharedEvent, signalValue uint64, cfg SharedEventEvalOptions) error {
	if k.modelType != ModelTypePackage {
		return &ANEError{Op: "eval", Err: fmt.Errorf("shared events require package-backed models")}
	}
	k.sharedEventUsed = true

	signalArr := foundation.NewNSMutableArray()
	if signal != nil {
		sigEvtClass := appleneuralengine.GetANESharedSignalEventClass()
		sigEvtObj := sigEvtClass.SignalEventWithValueSymbolIndexEventTypeSharedEvent(
			signalValue, 0, eventTypeSignal, signal.ev,
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
	req, reqObjects, err := createRequestFromSurfacesWithSharedEvents(k.inputs, k.outputs, sharedEventsObj)
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
	ok, err := k.client.MapIOSurfacesWithModelRequestCacheInferenceError(k.model, req, true)
	if err != nil || !ok {
		ok, err = k.client.MapIOSurfacesWithModelRequestCacheInferenceError(k.model, req, false)
		if err != nil || !ok {
			return &ANEError{Op: "map", Err: fmt.Errorf("map shared-event request failed: %w", err)}
		}
	}
	defer k.client.UnmapIOSurfacesWithModelRequest(k.model, req)
	if wait != nil {
		keepAlive = append(keepAlive, waitEvtObj, waitArr)
	}
	defer func() {
		runtime.KeepAlive(keepAlive)
		runtime.KeepAlive(req)
	}()

	// Build options dict for shared events.
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

	const qos = 21
	ok, err = k.client.EvaluateWithModelOptionsRequestQosError(k.model, opts, req, qos)
	if err == nil && ok {
		return nil
	}
	if os.Getenv("ANE_SHARED_USE_DIRECT") != "" {
		ok, err = k.client.DoEvaluateDirectWithModelOptionsRequestQosError(k.model, opts, req, qos)
		if err == nil && ok {
			return nil
		}
	}
	return wrapErr("eval", err)
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
