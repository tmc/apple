//go:build darwin

package ane

import (
	"fmt"
	"sync/atomic"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
)

// maxPoolDepth is the ANE hardware queue limit.
const maxPoolDepth = 127

// RequestPool pre-allocates a ring of ANE requests for pipelined evaluation.
// All requests share the same IOSurface mappings, enabling up to depth
// requests in flight simultaneously.
type RequestPool struct {
	kernel   *Kernel
	requests []appleneuralengine.ANERequest
	depth    int
	next     atomic.Uint64
}

// NewRequestPool creates a pool of depth pre-validated requests for the kernel.
// All requests share the kernel's IOSurfaces. Maximum depth is 127.
func NewRequestPool(k *Kernel, depth int) (*RequestPool, error) {
	if depth < 1 {
		return nil, fmt.Errorf("ane: pool depth must be >= 1")
	}
	if depth > maxPoolDepth {
		return nil, fmt.Errorf("ane: pool depth %d exceeds maximum %d", depth, maxPoolDepth)
	}
	if !k.mapped {
		return nil, &ANEError{Op: "pool", Err: fmt.Errorf("kernel not mapped")}
	}

	requests := make([]appleneuralengine.ANERequest, depth)

	// First slot reuses the kernel's existing request.
	requests[0] = k.request

	// Create additional requests referencing the same IOSurfaces.
	for i := 1; i < depth; i++ {
		req, err := createRequestFromSurfaces(k.inputs, k.outputs)
		if err != nil {
			return nil, fmt.Errorf("ane: pool request[%d]: %w", i, err)
		}

		// Map the new request's surfaces.
		var ok bool
		switch k.modelType {
		case ModelTypeMIL:
			ok, err = k.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(req, true)
			if err != nil || !ok {
				ok, err = k.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(req, false)
			}
		case ModelTypePackage:
			ok, err = k.client.MapIOSurfacesWithModelRequestCacheInferenceError(k.model, req, true)
			if err != nil || !ok {
				ok, err = k.client.MapIOSurfacesWithModelRequestCacheInferenceError(k.model, req, false)
			}
		}
		if err != nil || !ok {
			return nil, &ANEError{Op: "pool", Err: fmt.Errorf("map request[%d] failed: %w", i, err)}
		}

		requests[i] = req
	}

	return &RequestPool{
		kernel:   k,
		requests: requests,
		depth:    depth,
	}, nil
}

// createRequestFromSurfaces builds an ANERequest referencing existing IOSurfaces.
func createRequestFromSurfaces(inputs, outputs []coregraphics.IOSurfaceRef) (appleneuralengine.ANERequest, error) {
	ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()

	inputArr := foundation.NewNSMutableArray()
	inputIdxArr := foundation.NewNSMutableArray()
	for i, ref := range inputs {
		wrapped := ioClass.ObjectWithIOSurface(ref)
		inputArr.AddObject(wrapped)
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	outputArr := foundation.NewNSMutableArray()
	outputIdxArr := foundation.NewNSMutableArray()
	for i, ref := range outputs {
		wrapped := ioClass.ObjectWithIOSurface(ref)
		outputArr.AddObject(wrapped)
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(0)

	reqClass := appleneuralengine.GetANERequestClass()
	reqObj := reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(
		inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx,
	)
	if reqObj == nil || reqObj.GetID() == 0 {
		return appleneuralengine.ANERequest{}, &ANEError{Op: "pool", Err: fmt.Errorf("failed to create request")}
	}
	return appleneuralengine.ANERequestFromID(reqObj.GetID()), nil
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
	k := pr.pool.kernel
	if !pr.Request.Validate() {
		return &ANEError{Op: "eval", Err: fmt.Errorf("pooled request validation failed")}
	}

	switch k.modelType {
	case ModelTypeMIL:
		const qos = 21
		ok, err := k.inMemModel.EvaluateWithQoSOptionsRequestError(qos, nil, pr.Request)
		if err == nil && ok {
			return nil
		}
		return wrapErr("eval", err)
	case ModelTypePackage:
		const qos = 21
		ok, err := k.client.DoEvaluateDirectWithModelOptionsRequestQosError(k.model, nil, pr.Request, qos)
		if err == nil && ok {
			return nil
		}
		ok, err = k.client.EvaluateWithModelOptionsRequestQosError(k.model, nil, pr.Request, qos)
		if err == nil && ok {
			return nil
		}
		return wrapErr("eval", err)
	default:
		return &ANEError{Op: "eval", Err: fmt.Errorf("unknown model type %d", k.modelType)}
	}
}

// Release returns this request to the pool.
func (pr *PooledRequest) Release() {}

// Close unmaps all pooled requests except the kernel's original request.
func (p *RequestPool) Close() error {
	k := p.kernel
	for i := 1; i < p.depth; i++ {
		switch k.modelType {
		case ModelTypeMIL:
			k.inMemModel.UnmapIOSurfacesWithRequest(p.requests[i])
		case ModelTypePackage:
			k.client.UnmapIOSurfacesWithModelRequest(k.model, p.requests[i])
		}
	}
	return nil
}
