//go:build darwin

package ane

import (
	"fmt"
	"sync/atomic"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// maxPoolDepth bounds the number of requests a pool may pre-allocate. It is a
// limit on ring capacity, not on execution concurrency: it says nothing about
// how many evaluations the driver or firmware will run at once.
//
// TODO(unverified): the value 127 is inherited from earlier code in this
// package and has not been traced to a documented or measured driver limit.
const maxPoolDepth = 127

// RequestPool pre-allocates a fixed ring of ANERequest objects for a model, so
// that a hot loop can reuse validated requests instead of building one per
// evaluation. It saves the per-call cost of constructing the request object
// graph (the NSArrays of ANEIOSurfaceObject wrappers and index NSNumbers) and
// of mapping each new request's surfaces into the model.
//
// The pool does not provide concurrency. Every request in the pool references
// the same IOSurfaces as the model (see NewRequestPool, which builds each slot
// from the model's own input and output bindings), so two evaluations running
// at the same time would read and write the same operand memory and clobber
// each other. Callers must therefore keep evaluation serialized: acquire, eval,
// and consume the outputs before the next Eval. Round-robin Acquire only
// spreads reuse across slots; it does not make overlapping use safe.
//
// Whether concurrent submissions would serialize below us on the _ANEClient
// path is not determined by this code: Eval makes a blocking evaluate call and
// this package does not observe the driver's queueing. Bryngelson (arXiv
// 2606.22283, section 2.4) reports, for an M1 reached through the e5rt/IOKit
// route rather than _ANEClient, that the driver keeps at most one firmware
// command in flight and that two concurrent submission threads serialized at
// 1.04x. That is the paper's measurement on a different route, not ours.
//
// A related hazard for future work, from the same paper (section 6.4) and again
// on the e5rt route: pipelining several operations within a single stream is
// sound, but overlapping two or more streams is not, as the completion event of
// the second and later streams never notifies and the waiter blocks. This
// package does not use that stream API today.
type RequestPool struct {
	model    *Model
	requests []appleneuralengine.ANERequest
	unmaps   []func()
	depth    int
	next     atomic.Uint64
}

// NewRequestPool creates a pool of depth requests for the model. Slot 0 reuses
// the model's existing request; the remaining slots are new ANERequest objects
// built from the model's input and output bindings, so all slots reference the
// same IOSurfaces and must not be evaluated concurrently. depth must be between
// 1 and maxPoolDepth.
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
		wrapped := ioClass.ObjectWithIOSurface(iosurface.IOSurfaceRef(binding.Surface))
		inputArr.AddObject(wrapped)
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(int32(binding.SymbolIndex)))
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
		wrapped := ioClass.ObjectWithIOSurface(iosurface.IOSurfaceRef(binding.Surface))
		outputArr.AddObject(wrapped)
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(int32(binding.SymbolIndex)))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(int32(procedureIndex))
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

// Acquire returns the next request from the pool in round-robin order. The
// index advance is atomic, but the returned request shares IOSurfaces with
// every other slot, so the caller must still evaluate one request at a time.
func (p *RequestPool) Acquire() *PooledRequest {
	idx := int(p.next.Add(1)-1) % p.depth
	return &PooledRequest{
		Request: p.requests[idx],
		pool:    p,
		idx:     idx,
	}
}

// Eval evaluates this request on the ANE, blocking until the underlying
// evaluate call returns. The request's operands are the model's IOSurfaces, so
// the caller must read the outputs before evaluating another pooled request.
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

// Release is a no-op. Slots are handed out round-robin by Acquire and are never
// checked back in; it exists so callers can defer a symmetric call.
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
