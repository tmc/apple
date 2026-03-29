//go:build darwin

package ane

import (
	"fmt"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// MultiSurfaceEvalPlanConfig controls reusable mapped-request construction.
type MultiSurfaceEvalPlanConfig struct {
	ProcedureIndex      int
	QoS                 uint32
	DisableCacheMapping bool
	PreferDirect        bool
	EnableMetalWait     bool
	EnableMetalSignal   bool
	WaitValue           uint64
	SignalValue         uint64
	EnableFWToFWSignal  bool
	SignalOutputBinding int
}

// DefaultMultiSurfaceEvalPlanConfig returns conservative defaults.
func DefaultMultiSurfaceEvalPlanConfig() MultiSurfaceEvalPlanConfig {
	return MultiSurfaceEvalPlanConfig{
		ProcedureIndex:      0,
		PreferDirect:        true,
		WaitValue:           1,
		SignalValue:         1,
		SignalOutputBinding: 0,
	}
}

// MultiSurfaceEvalPlan keeps a mapped ANE request alive across evals.
type MultiSurfaceEvalPlan struct {
	model        *Model
	request      appleneuralengine.ANERequest
	options      foundation.NSMutableDictionary
	qos          uint32
	preferDirect bool
	waitEvent    *SharedEvent
	waitValue    uint64
	signalEvent  *SharedEvent
	signalValue  uint64
	retained     []objectivec.IObject
	inputs       []SurfaceBinding
	outputs      []SurfaceBinding
	unmap        func()
	closeOnce    sync.Once
}

// NewMultiSurfaceEvalPlan creates a reusable mapped request for caller-owned surfaces.
func NewMultiSurfaceEvalPlan(model *Model, inputs, outputs []SurfaceBinding, cfg MultiSurfaceEvalPlanConfig) (*MultiSurfaceEvalPlan, error) {
	if model == nil {
		return nil, fmt.Errorf("ane: multi-surface eval plan model is nil")
	}
	if len(inputs) == 0 || len(outputs) == 0 {
		return nil, fmt.Errorf("ane: multi-surface eval plan requires at least one input and output")
	}
	def := DefaultMultiSurfaceEvalPlanConfig()
	if cfg == (MultiSurfaceEvalPlanConfig{}) {
		cfg = def
	}
	if cfg.QoS == 0 {
		cfg.QoS = model.qosValue()
	}
	if cfg.WaitValue == 0 {
		cfg.WaitValue = def.WaitValue
	}
	if cfg.SignalValue == 0 {
		cfg.SignalValue = def.SignalValue
	}
	if cfg.EnableMetalWait {
		cfg.EnableFWToFWSignal = true
	}
	if cfg.EnableMetalSignal {
		if cfg.SignalOutputBinding < 0 || cfg.SignalOutputBinding >= len(outputs) {
			return nil, fmt.Errorf("ane: signal output binding %d out of range [0,%d)", cfg.SignalOutputBinding, len(outputs))
		}
		if outputs[cfg.SignalOutputBinding].SymbolIndex < 0 {
			return nil, fmt.Errorf("ane: signal output binding %d has invalid symbol index %d", cfg.SignalOutputBinding, outputs[cfg.SignalOutputBinding].SymbolIndex)
		}
	}

	options := foundation.NewNSMutableDictionary()
	var waitEvent *SharedEvent
	var signalEvent *SharedEvent
	var retained []objectivec.IObject
	var sharedEvents objectivec.IObject

	if cfg.EnableMetalWait || cfg.EnableMetalSignal {
		if cfg.DisableCacheMapping {
			options.SetObjectForKey(
				foundation.GetNSNumberClass().NumberWithBool(true),
				nsStringKey("kANEFDisableIOFencesUseSharedEventsKey"),
			)
		} else {
			options.SetObjectForKey(
				foundation.GetNSNumberClass().NumberWithBool(true),
				nsStringKey("kANEFDisableIOFencesUseSharedEventsKey"),
			)
		}
		if cfg.EnableFWToFWSignal {
			options.SetObjectForKey(
				foundation.GetNSNumberClass().NumberWithBool(true),
				nsStringKey("kANEFEnableFWToFWSignal"),
			)
		}

		signalArr := foundation.NewNSMutableArray()
		waitArr := foundation.NewNSMutableArray()
		retained = append(retained, signalArr, waitArr)

		if cfg.EnableMetalSignal {
			var err error
			signalEvent, err = NewSharedEvent()
			if err != nil {
				return nil, err
			}
			sig := appleneuralengine.GetANESharedSignalEventClass().SignalEventWithValueSymbolIndexEventTypeSharedEvent(
				cfg.SignalValue,
				uint32(outputs[cfg.SignalOutputBinding].SymbolIndex),
				eventTypeSignal,
				signalEvent.ev,
			)
			if sig == nil || sig.GetID() == 0 {
				signalEvent.Close()
				return nil, fmt.Errorf("ane: failed to create ANESharedSignalEvent")
			}
			signalArr.AddObject(sig)
			retained = append(retained, sig)
		}
		if cfg.EnableMetalWait {
			var err error
			waitEvent, err = NewSharedEvent()
			if err != nil {
				if signalEvent != nil {
					signalEvent.Close()
				}
				return nil, err
			}
			waitObj := appleneuralengine.GetANESharedWaitEventClass().WaitEventWithValueSharedEvent(cfg.WaitValue, waitEvent.ev)
			if waitObj == nil || waitObj.GetID() == 0 {
				waitEvent.Close()
				if signalEvent != nil {
					signalEvent.Close()
				}
				return nil, fmt.Errorf("ane: failed to create ANESharedWaitEvent")
			}
			waitArr.AddObject(waitObj)
			retained = append(retained, waitObj)
		}

		sharedEvents = appleneuralengine.GetANESharedEventsClass().SharedEventsWithSignalEventsWaitEvents(signalArr, waitArr)
		if sharedEvents == nil || sharedEvents.GetID() == 0 {
			if waitEvent != nil {
				waitEvent.Close()
			}
			if signalEvent != nil {
				signalEvent.Close()
			}
			return nil, fmt.Errorf("ane: failed to create ANESharedEvents")
		}
	}

	request, keepAlive, err := createRequestFromBindingsWithSharedEvents(inputs, outputs, cfg.ProcedureIndex, sharedEvents)
	if err != nil {
		if waitEvent != nil {
			waitEvent.Close()
		}
		if signalEvent != nil {
			signalEvent.Close()
		}
		return nil, err
	}
	if !request.Validate() {
		if waitEvent != nil {
			waitEvent.Close()
		}
		if signalEvent != nil {
			signalEvent.Close()
		}
		return nil, fmt.Errorf("ane: request validation failed")
	}
	retained = append(retained, keepAlive...)

	unmap, err := model.mapRequest(request, !cfg.DisableCacheMapping)
	if err != nil {
		if waitEvent != nil {
			waitEvent.Close()
		}
		if signalEvent != nil {
			signalEvent.Close()
		}
		return nil, wrapErr("map", err)
	}
	model.mu.Lock()
	model.sharedEventUsed = model.sharedEventUsed || waitEvent != nil || signalEvent != nil
	model.mu.Unlock()

	return &MultiSurfaceEvalPlan{
		model:        model,
		request:      request,
		options:      options,
		qos:          cfg.QoS,
		preferDirect: cfg.PreferDirect,
		waitEvent:    waitEvent,
		waitValue:    cfg.WaitValue,
		signalEvent:  signalEvent,
		signalValue:  cfg.SignalValue,
		retained:     retained,
		inputs:       append([]SurfaceBinding(nil), inputs...),
		outputs:      append([]SurfaceBinding(nil), outputs...),
		unmap:        unmap,
	}, nil
}

// Eval evaluates the mapped request.
func (p *MultiSurfaceEvalPlan) Eval() error {
	if p == nil {
		return fmt.Errorf("ane: multi-surface eval plan is nil")
	}
	return wrapErr("eval", p.model.evaluateRequestWithQoS(p.request, p.options, p.qos, p.preferDirect))
}

// WaitEventPort returns the shared-event Mach port used for Metal->ANE wait sync.
func (p *MultiSurfaceEvalPlan) WaitEventPort() uint32 {
	if p == nil || p.waitEvent == nil {
		return 0
	}
	return p.waitEvent.Port()
}

// SignalEventPort returns the shared-event Mach port used for ANE->Metal signal sync.
func (p *MultiSurfaceEvalPlan) SignalEventPort() uint32 {
	if p == nil || p.signalEvent == nil {
		return 0
	}
	return p.signalEvent.Port()
}

// WaitValue returns the configured wait-event value.
func (p *MultiSurfaceEvalPlan) WaitValue() uint64 {
	if p == nil {
		return 0
	}
	return p.waitValue
}

// SignalValue returns the configured signal-event value.
func (p *MultiSurfaceEvalPlan) SignalValue() uint64 {
	if p == nil {
		return 0
	}
	return p.signalValue
}

// WaitEvent returns a wrapper for the wait-event Mach port, if present.
func (p *MultiSurfaceEvalPlan) WaitEvent() *SharedEvent {
	if p == nil || p.waitEvent == nil {
		return nil
	}
	ev, err := SharedEventFromPort(p.waitEvent.Port())
	if err != nil {
		return nil
	}
	return ev
}

// SignalEvent returns a wrapper for the signal-event Mach port, if present.
func (p *MultiSurfaceEvalPlan) SignalEvent() *SharedEvent {
	if p == nil || p.signalEvent == nil {
		return nil
	}
	ev, err := SharedEventFromPort(p.signalEvent.Port())
	if err != nil {
		return nil
	}
	return ev
}

// Close unmaps the request and releases owned shared events.
func (p *MultiSurfaceEvalPlan) Close() error {
	if p == nil {
		return nil
	}
	var err error
	p.closeOnce.Do(func() {
		if p.unmap != nil {
			p.unmap()
			p.unmap = nil
		}
		if p.waitEvent != nil {
			err = p.waitEvent.Close()
			p.waitEvent = nil
		}
		if p.signalEvent != nil {
			if closeErr := p.signalEvent.Close(); err == nil {
				err = closeErr
			}
			p.signalEvent = nil
		}
		for _, obj := range p.retained {
			if obj != nil && obj.GetID() != 0 {
				objectivec.ObjectFromID(obj.GetID()).Release()
			}
		}
		p.retained = nil
	})
	return err
}
