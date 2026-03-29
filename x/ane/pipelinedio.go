//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/private/appleneuralengine"
)

// PipelinedIO decouples input notification from evaluation and output
// enqueuing from readback, enabling overlap of GPU texture reads with
// ANE input preparation. Package models only.
type PipelinedIO struct {
	model    *Model
	openLoop bool
}

// NewPipelinedIO creates a PipelinedIO for the given model.
func NewPipelinedIO(m *Model) (*PipelinedIO, error) {
	if m.modelType != ModelTypePackage {
		return nil, &ANEError{Op: "pipelined-io", Err: fmt.Errorf("pipelined I/O requires package models")}
	}
	return &PipelinedIO{model: m}, nil
}

// NotifyInputsReady tells the ANE that input buffers for the given procedure
// are ready to be consumed.
func (p *PipelinedIO) NotifyInputsReady(procedureIdx int) error {
	cls := appleneuralengine.GetANEInputBuffersReadyClass()
	obj := cls.InputBuffersWithProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay(
		uint32(procedureIdx), nil, nil, 0,
	)
	if obj == nil || obj.GetID() == 0 {
		return &ANEError{Op: "pipelined-io", Err: fmt.Errorf("failed to create input buffers ready")}
	}
	buffers := appleneuralengine.ANEInputBuffersReadyFromID(obj.GetID())

	const qos = 21
	ok, err := p.model.aneClient.BuffersReadyWithModelInputBuffersOptionsQosError(
		p.model.aneModel, buffers, nil, qos,
	)
	if err != nil || !ok {
		return wrapErr("pipelined-io", err)
	}
	return nil
}

// EnqueueOutputSet enqueues an output set for the ANE to fill.
func (p *PipelinedIO) EnqueueOutputSet(procedureIdx, setIdx int, signalValue uint64) error {
	set := appleneuralengine.NewANEOutputSetEnqueueOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop(
		uint32(procedureIdx), uint32(setIdx), signalValue, false, p.openLoop,
	)
	if set.ID == 0 {
		return &ANEError{Op: "pipelined-io", Err: fmt.Errorf("failed to create output set enqueue")}
	}

	const qos = 21
	ok, err := p.model.aneClient.EnqueueSetsWithModelOutputSetOptionsQosError(
		p.model.aneModel, &set, nil, qos,
	)
	if err != nil || !ok {
		return wrapErr("pipelined-io", err)
	}
	return nil
}

// SetOpenLoop controls whether output reads proceed without waiting
// for the ANE to signal completion.
func (p *PipelinedIO) SetOpenLoop(open bool) {
	p.openLoop = open
}

// Close is a no-op; PipelinedIO holds no resources beyond its model reference.
func (p *PipelinedIO) Close() error {
	return nil
}
