//go:build darwin

package ane

import (
	"fmt"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
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
	return k.evalWithSharedEvents(0, 0, signalPort, signalValue, cfg)
}

// EvalBidirectional evaluates the kernel with both wait and signal events.
// The ANE waits for waitValue on waitPort before executing, then signals
// signalValue on signalPort after completion.
func (k *Kernel) EvalBidirectional(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	return k.evalWithSharedEvents(waitPort, waitValue, signalPort, signalValue, cfg)
}

const eventTypeSignal = int64(5)

func (k *Kernel) evalWithSharedEvents(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	// Build signal event.
	signalEv := iosurface.NewIOSurfaceSharedEventWithMachPort(signalPort)
	if signalEv.ID == 0 {
		return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create signal shared event")}
	}

	sigEvtClass := appleneuralengine.GetANESharedSignalEventClass()
	sigEvtObj := sigEvtClass.SignalEventWithValueSymbolIndexEventTypeSharedEvent(
		signalValue, 0, eventTypeSignal, signalEv,
	)
	if sigEvtObj == nil || sigEvtObj.GetID() == 0 {
		return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create ANESharedSignalEvent")}
	}

	signalArr := foundation.NewNSMutableArray()
	signalArr.AddObject(sigEvtObj)

	waitArr := foundation.NewNSMutableArray()
	if waitPort != 0 {
		waitEv := iosurface.NewIOSurfaceSharedEventWithMachPort(waitPort)
		if waitEv.ID == 0 {
			return &ANEError{Op: "eval", Err: fmt.Errorf("failed to create wait shared event")}
		}
		waitEvtClass := appleneuralengine.GetANESharedWaitEventClass()
		waitEvtObj := waitEvtClass.WaitEventWithValueSharedEvent(waitValue, waitEv)
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
	sharedEvents := appleneuralengine.ANESharedEventsFromID(sharedEventsObj.GetID())
	k.request.SetSharedEvents(&sharedEvents)

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
	if !k.request.Validate() {
		k.request.SetSharedEvents(nil)
		return &ANEError{Op: "eval", Err: fmt.Errorf("request validation failed")}
	}

	const qos = 21
	var evalErr error

	switch k.modelType {
	case ModelTypePackage:
		ok, err := k.client.DoEvaluateDirectWithModelOptionsRequestQosError(k.model, opts, k.request, qos)
		if err != nil || !ok {
			ok, err = k.client.EvaluateWithModelOptionsRequestQosError(k.model, opts, k.request, qos)
			if err != nil || !ok {
				evalErr = wrapErr("eval", err)
			}
		}
	case ModelTypeMIL:
		ok, err := k.inMemModel.EvaluateWithQoSOptionsRequestError(qos, opts, k.request)
		if err != nil || !ok {
			evalErr = wrapErr("eval", err)
		}
	default:
		evalErr = &ANEError{Op: "eval", Err: fmt.Errorf("unknown model type %d", k.modelType)}
	}

	k.request.SetSharedEvents(nil)
	return evalErr
}

// nsStringKey creates an NSCopying-compatible key from a string for use with NSDictionary.
func nsStringKey(s string) foundation.NSCopyingObject {
	nsStr := foundation.NewStringWithString(s)
	return foundation.NSCopyingObject{Object: nsStr.Object}
}
