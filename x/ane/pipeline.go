//go:build darwin

package ane

import (
	"fmt"
	"sync/atomic"

	"github.com/tmc/apple/metal"
)

// Pipeline formalizes the SharedEvent + MetalDevice pattern for
// synchronizing ANE and Metal GPU execution.
//
// Each call to ANEToMetal or Bidirectional auto-increments the event
// counter to sequence operations. WaitOnANE only waits without signaling.
type Pipeline struct {
	metal    *MetalDevice
	aneEvent *SharedEvent
	mtlEvent metal.MTLSharedEventObject
	counter  atomic.Uint64
}

// NewPipeline creates a Pipeline with a shared event bridging ANE and Metal.
func NewPipeline(md *MetalDevice) (*Pipeline, error) {
	mtlEv, aneEv, err := md.NewMetalSharedEvent()
	if err != nil {
		return nil, fmt.Errorf("ane pipeline: %w", err)
	}
	p := &Pipeline{
		metal:    md,
		aneEvent: aneEv,
		mtlEvent: mtlEv,
	}
	return p, nil
}

// Counter returns the current event counter value.
func (p *Pipeline) Counter() uint64 {
	return p.counter.Load()
}

// ANEToMetal evaluates the model on ANE and signals Metal on completion.
// The event counter is incremented and used as the signal value.
func (p *Pipeline) ANEToMetal(m *Model) error {
	val := p.counter.Add(1)
	return m.EvalWithSignalEvent(p.aneEvent.Port(), val, SharedEventEvalOptions{
		DisableIOFencesUseSharedEvents: true,
	})
}

// WaitOnANE evaluates the model on ANE after waiting for Metal to signal.
// Uses the current counter as the wait value. Does not signal back or
// increment the counter.
func (p *Pipeline) WaitOnANE(m *Model) error {
	waitVal := p.counter.Load()
	return m.EvalWithWaitEvent(p.aneEvent.Port(), waitVal, SharedEventEvalOptions{
		DisableIOFencesUseSharedEvents: true,
	})
}

// Bidirectional evaluates the model with both wait and signal events.
// Waits for the current counter value, then signals the next.
func (p *Pipeline) Bidirectional(m *Model) error {
	waitVal := p.counter.Load()
	signalVal := p.counter.Add(1)
	return m.EvalBidirectional(
		p.aneEvent.Port(), waitVal,
		p.aneEvent.Port(), signalVal,
		SharedEventEvalOptions{
			DisableIOFencesUseSharedEvents: true,
		},
	)
}

// Metal returns the MetalDevice for issuing GPU commands.
func (p *Pipeline) Metal() *MetalDevice { return p.metal }

// MetalEvent returns the Metal-side shared event for encoding
// GPU command buffer waits and signals.
func (p *Pipeline) MetalEvent() metal.MTLSharedEventObject { return p.mtlEvent }

// ANEEvent returns the ANE-side shared event.
func (p *Pipeline) ANEEvent() *SharedEvent { return p.aneEvent }

// Close releases the shared event resources.
func (p *Pipeline) Close() error {
	return p.aneEvent.Close()
}
