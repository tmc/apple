//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
)

// RealTimeTask brackets a sequence of ANE evaluations with real-time
// scheduling priority. Call End when the latency-critical section is done.
type RealTimeTask struct {
	client *Client
	active bool
}

// BeginRealTimeTask enters real-time scheduling mode on the ANE.
func (c *Client) BeginRealTimeTask() (*RealTimeTask, error) {
	if !c.aneClient.BeginRealTimeTask() {
		return nil, &ANEError{Op: "begin-rt", Err: fmt.Errorf("failed to begin real-time task")}
	}
	return &RealTimeTask{client: c, active: true}, nil
}

// Eval executes the model under real-time scheduling.
func (rt *RealTimeTask) Eval(m *Model) error {
	if !rt.active {
		return &ANEError{Op: "eval-rt", Err: fmt.Errorf("real-time task not active")}
	}
	emptyOpts := foundation.NewNSMutableDictionary()

	var model appleneuralengine.ANEModel
	switch m.modelType {
	case ModelTypeMIL:
		inner := m.inMemModel.Model()
		if inner == nil {
			return &ANEError{Op: "eval-rt", Err: fmt.Errorf("nil inner model")}
		}
		model = *inner
	case ModelTypePackage:
		model = m.aneModel
	default:
		return &ANEError{Op: "eval-rt", Err: fmt.Errorf("unknown model type %d", m.modelType)}
	}

	ok, err := rt.client.aneClient.EvaluateRealTimeWithModelOptionsRequestError(model, emptyOpts, m.request)
	if err != nil || !ok {
		return wrapErr("eval-rt", err)
	}
	return nil
}

// End exits real-time scheduling mode.
func (rt *RealTimeTask) End() error {
	if !rt.active {
		return nil
	}
	rt.active = false
	if !rt.client.aneClient.EndRealTimeTask() {
		return &ANEError{Op: "end-rt", Err: fmt.Errorf("failed to end real-time task")}
	}
	return nil
}
