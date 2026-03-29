//go:build darwin

package ane

import (
	"fmt"
	"sync/atomic"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// maxPoolDepth is the ANE hardware queue limit.
const maxPoolDepth = 127

// RequestPool pre-allocates a ring of ANE requests for pipelined evaluation.
// All requests share the same IOSurface mappings, enabling up to depth
// requests in flight simultaneously.
type RequestPool struct {
	model    *Model
	requests []appleneuralengine.ANERequest
	unmaps   []func()
	depth    int
	next     atomic.Uint64
}

// NewRequestPool creates a pool of depth pre-validated requests for the model.
// All requests share the model's IOSurfaces. Maximum depth is 127.
func NewRequestPool(m *Model, depth int) (*RequestPool, error) {
	if depth < 1 {
		return nil, fmt.Errorf("ane: pool depth must be >= 1")
	}
	if depth > maxPoolDepth {
		return nil, fmt.Errorf("ane: pool depth %d exceeds maximum %d", depth, maxPoolDepth)
	}
	if !m.mapped {
		return nil, &ANEError{Op: "pool", Err: fmt.Errorf("kernel not mapped")}
	}

	requests := make([]appleneuralengine.ANERequest, depth)
	unmaps := make([]func(), depth)

	// First slot reuses the model's existing request.
	requests[0] = m.request

	// Create additional requests referencing the same IOSurfaces.
	inputBindings := m.inputBindings()
	outputBindings := m.outputBindings()
	for i := 1; i < depth; i++ {
		req, _, err := createRequestFromBindingsWithSharedEvents(inputBindings, outputBindings, 0, nil)
		if err != nil {
			return nil, fmt.Errorf("ane: pool request[%d]: %w", i, err)
		}

		unmap, err := m.mapRequestWithFallback(req)
		if err != nil {
			return nil, &ANEError{Op: "pool", Err: fmt.Errorf("map request[%d] failed: %w", i, err)}
		}

		requests[i] = req
		unmaps[i] = unmap
	}

	return &RequestPool{
		model:    m,
		requests: requests,
		unmaps:   unmaps,
		depth:    depth,
	}, nil
}

// createRequestFromSurfaces builds an ANERequest referencing existing IOSurfaces.
func createRequestFromSurfaces(inputs, outputs []coregraphics.IOSurfaceRef) (appleneuralengine.ANERequest, error) {
	req, _, err := createRequestFromSurfacesWithSharedEvents(inputs, outputs, nil)
	return req, err
}

func createRequestFromSurfacesWithSharedEvents(inputs, outputs []coregraphics.IOSurfaceRef, sharedEvents objectivec.IObject) (appleneuralengine.ANERequest, []objectivec.IObject, error) {
	inputBindings := make([]SurfaceBinding, len(inputs))
	for i, ref := range inputs {
		inputBindings[i] = SurfaceBinding{Surface: ref, SymbolIndex: i}
	}
	outputBindings := make([]SurfaceBinding, len(outputs))
	for i, ref := range outputs {
		outputBindings[i] = SurfaceBinding{Surface: ref, SymbolIndex: i}
	}
	return createRequestFromBindingsWithSharedEvents(inputBindings, outputBindings, 0, sharedEvents)
}

func createRequestFromBindingsWithSharedEvents(inputs, outputs []SurfaceBinding, procedureIndex int, sharedEvents objectivec.IObject) (appleneuralengine.ANERequest, []objectivec.IObject, error) {
	ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()

	inputArr := foundation.NewNSMutableArray()
	inputIdxArr := foundation.NewNSMutableArray()
	for i, binding := range inputs {
		if binding.Surface == 0 {
			return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("input[%d] has nil IOSurface", i)}
		}
		if binding.SymbolIndex < 0 {
			return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("input[%d] has invalid symbol index %d", i, binding.SymbolIndex)}
		}
		wrapped := ioClass.ObjectWithIOSurface(binding.Surface)
		inputArr.AddObject(wrapped)
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(binding.SymbolIndex))
	}

	outputArr := foundation.NewNSMutableArray()
	outputIdxArr := foundation.NewNSMutableArray()
	for i, binding := range outputs {
		if binding.Surface == 0 {
			return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("output[%d] has nil IOSurface", i)}
		}
		if binding.SymbolIndex < 0 {
			return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("output[%d] has invalid symbol index %d", i, binding.SymbolIndex)}
		}
		wrapped := ioClass.ObjectWithIOSurface(binding.Surface)
		outputArr.AddObject(wrapped)
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(binding.SymbolIndex))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(procedureIndex)
	txnHandle := foundation.GetNSNumberClass().NumberWithUnsignedLongLong(1)

	reqClass := appleneuralengine.GetANERequestClass()
	reqObj := reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(
		inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx, sharedEvents, txnHandle,
	)
	if reqObj == nil || reqObj.GetID() == 0 {
		if sharedEvents != nil {
			return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("failed to create request with shared events")}
		}
		reqObj = reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(
			inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx,
		)
	}
	if reqObj == nil || reqObj.GetID() == 0 {
		return appleneuralengine.ANERequest{}, nil, &ANEError{Op: "pool", Err: fmt.Errorf("failed to create request")}
	}
	req := appleneuralengine.ANERequestFromID(reqObj.GetID())
	if req.TransactionHandle().GetID() == 0 {
		req.SetTransactionHandle(txnHandle)
	}
	keepAlive := []objectivec.IObject{inputArr, inputIdxArr, outputArr, outputIdxArr}
	if sharedEvents != nil {
		keepAlive = append(keepAlive, sharedEvents)
	}
	return req, keepAlive, nil
}

// PooledRequest is a request checked out from a pool.
type PooledRequest struct {
	Request appleneuralengine.ANERequest
	pool    *RequestPool
	idx     int
}

// Acquire returns the next request from the pool in round-robin order.
func (p *RequestPool) Acquire() *PooledRequest {
	idx := int(p.next.Add(1)-1) % p.depth
	return &PooledRequest{
		Request: p.requests[idx],
		pool:    p,
		idx:     idx,
	}
}

// Eval evaluates this request on the ANE.
func (pr *PooledRequest) Eval() error {
	mdl := pr.pool.model
	if !pr.Request.Validate() {
		return &ANEError{Op: "eval", Err: fmt.Errorf("pooled request validation failed")}
	}

	switch mdl.modelType {
	case ModelTypeMIL:
		return mdl.evaluateRequestWithOptions(pr.Request, nil, true)
	case ModelTypePackage:
		return mdl.evaluateRequestWithOptions(pr.Request, nil, true)
	default:
		return &ANEError{Op: "eval", Err: fmt.Errorf("unknown model type %d", mdl.modelType)}
	}
}

// Release returns this request to the pool.
func (pr *PooledRequest) Release() {}

// Close unmaps all pooled requests except the model's original request.
func (p *RequestPool) Close() error {
	mdl := p.model
	for i := 1; i < p.depth; i++ {
		if i < len(p.unmaps) && p.unmaps[i] != nil {
			p.unmaps[i]()
			p.unmaps[i] = nil
			continue
		}
		switch mdl.modelType {
		case ModelTypeMIL:
			mdl.inMemModel.UnmapIOSurfacesWithRequest(p.requests[i])
		case ModelTypePackage:
			mdl.aneClient.UnmapIOSurfacesWithModelRequest(mdl.aneModel, p.requests[i])
		}
	}
	return nil
}
