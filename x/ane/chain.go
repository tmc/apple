//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// ShareSurface binds dst's input[dstInput] to src's output[srcOutput],
// sharing the same IOSurface for zero-copy handoff between models.
// Both models must have compatible tensor layouts at the given indices.
//
// Rebinding the surface is not enough on its own. A Model's ANERequest holds
// ANEIOSurfaceObject wrappers built when the model was compiled, and Eval runs
// that request, so a surface swapped into the Go slice afterwards is visible to
// WriteInput and ReadOutput but not to the engine. This rebuilds dst's request
// so that the next Eval reads the surface named here.
func ShareSurface(src *Model, srcOutput int, dst *Model, dstInput int) error {
	if srcOutput < 0 || srcOutput >= len(src.outputs) {
		return fmt.Errorf("ane: src output index %d out of range [0,%d)", srcOutput, len(src.outputs))
	}
	if dstInput < 0 || dstInput >= len(dst.inputs) {
		return fmt.Errorf("ane: dst input index %d out of range [0,%d)", dstInput, len(dst.inputs))
	}

	previous := dst.inputs[dstInput]
	dst.inputs[dstInput] = src.outputs[srcOutput]
	request, err := buildRequest(dst.inputs, dst.outputs, dst.inputLayouts, dst.outputLayouts)
	if err != nil {
		dst.inputs[dstInput] = previous
		return fmt.Errorf("ane: rebuild request after sharing surface: %w", err)
	}
	// Hand the retain over to the new request rather than adding one.
	//
	// While objsRetained is set the Model owns exactly one retain on its
	// request, and Close releases exactly one — so retaining the new request
	// without releasing the old leaks every request but the last, which is what
	// an earlier version of this did. The order matters only in that the new
	// request is retained before the old one is released.
	//
	// Calling this while an evaluation is in flight is not supported, here or
	// anywhere else in the type: Eval reads m.request without holding m.mu.
	previousRequest := dst.request
	dst.request = request
	if dst.objsRetained {
		objectivec.ObjectFromID(request.ID).Retain()
		if previousRequest.ID != 0 {
			objectivec.ObjectFromID(previousRequest.ID).Release()
		}
	}
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
			inputArr.AddObject(ioClass.ObjectWithIOSurface(iosurface.IOSurfaceRef(surf)))
		}

		outputArr := foundation.NewNSMutableArray()
		for _, surf := range m.outputs {
			ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()
			outputArr.AddObject(ioClass.ObjectWithIOSurface(iosurface.IOSurfaceRef(surf)))
		}

		lbInputIdx := foundation.NewNSMutableArray()
		lbInputIdx.AddObject(numClass.NumberWithInt(int32(link.LoopbackInputIdx)))

		lbOutputIdx := foundation.NewNSMutableArray()
		lbOutputIdx.AddObject(numClass.NumberWithInt(int32(link.LoopbackOutputIdx)))

		procIdx := numClass.NumberWithInt(int32(link.ProcedureIndex))
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
