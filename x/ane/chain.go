//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
)

// ShareSurface binds dst's input[dstInput] to src's output[srcOutput],
// sharing the same IOSurface for zero-copy handoff between models.
// Both models must have compatible tensor layouts at the given indices.
func ShareSurface(src *Model, srcOutput int, dst *Model, dstInput int) error {
	if srcOutput < 0 || srcOutput >= len(src.outputs) {
		return fmt.Errorf("ane: src output index %d out of range [0,%d)", srcOutput, len(src.outputs))
	}
	if dstInput < 0 || dstInput >= len(dst.inputs) {
		return fmt.Errorf("ane: dst input index %d out of range [0,%d)", dstInput, len(dst.inputs))
	}
	dst.inputs[dstInput] = src.outputs[srcOutput]
	return nil
}

// ChainLink describes one step in an intra-model procedure chain.
type ChainLink struct {
	ProcedureIndex    int
	LoopbackInputIdx  int
	LoopbackOutputIdx int
	EnqueueDelay      uint64
}

// PrepareChain sets up intra-model procedure chaining using ANEChainingRequest.
// This is for multi-procedure .mlmodelc packages where procedures feed into
// each other via loopback symbols. Package models only.
func (m *Model) PrepareChain(links []ChainLink) error {
	if m.modelType != ModelTypePackage {
		return &ANEError{Op: "chain", Err: fmt.Errorf("chaining requires package models")}
	}
	if len(links) == 0 {
		return nil
	}

	numClass := foundation.GetNSNumberClass()
	emptyOpts := foundation.NewNSMutableDictionary()

	for _, link := range links {
		inputArr := foundation.NewNSMutableArray()
		for _, surf := range m.inputs {
			ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()
			inputArr.AddObject(ioClass.ObjectWithIOSurface(surf))
		}

		outputArr := foundation.NewNSMutableArray()
		for _, surf := range m.outputs {
			ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()
			outputArr.AddObject(ioClass.ObjectWithIOSurface(surf))
		}

		lbInputIdx := foundation.NewNSMutableArray()
		lbInputIdx.AddObject(numClass.NumberWithInt(link.LoopbackInputIdx))

		lbOutputIdx := foundation.NewNSMutableArray()
		lbOutputIdx.AddObject(numClass.NumberWithInt(link.LoopbackOutputIdx))

		procIdx := numClass.NumberWithInt(link.ProcedureIndex)
		txnHandle := numClass.NumberWithUnsignedLongLong(1)
		delay := numClass.NumberWithUnsignedLongLong(link.EnqueueDelay)
		memPoolId := numClass.NumberWithInt(0)

		reqClass := appleneuralengine.GetANEChainingRequestClass()
		reqObj := reqClass.ChainingRequestWithInputsOutputSetsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId(
			inputArr, outputArr, lbInputIdx, lbOutputIdx, procIdx, nil, txnHandle, delay, memPoolId,
		)
		if reqObj == nil || reqObj.GetID() == 0 {
			return &ANEError{Op: "chain", Err: fmt.Errorf("failed to create chaining request for procedure %d", link.ProcedureIndex)}
		}

		ok, err := m.aneClient.PrepareChainingWithModelOptionsChainingReqQosError(
			m.aneModel, emptyOpts, reqObj, m.qosValue(),
		)
		if err != nil || !ok {
			return &ANEError{Op: "chain", Err: fmt.Errorf("prepare chaining failed for procedure %d: %w", link.ProcedureIndex, err)}
		}
	}
	return nil
}
