//go:build darwin

package model

import (
	"fmt"
	"sync/atomic"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
	xane "github.com/tmc/apple/x/ane"
)

const maxKernelPoolDepth = 127

// KernelRequestPool pre-allocates a ring of ANE requests for pipelined
// evaluation of a single kernel. For shared MIL kernels, requests use
// the clone's IOSurfaces rather than the owner model's surfaces.
//
// For non-shared kernels, delegates to xane.RequestPool internally.
type KernelRequestPool struct {
	kernel   *Kernel
	shared   bool // true if using shared MIL path
	requests []appleneuralengine.ANERequest
	xanePool *xane.RequestPool // non-nil for non-shared models
	depth    int
	next     atomic.Uint64
}

// KernelPooledRequest is a request checked out from a KernelRequestPool.
type KernelPooledRequest struct {
	pool    *KernelRequestPool
	xanePR  *xane.PooledRequest // non-nil for non-shared models
	request appleneuralengine.ANERequest
}

// NewRequestPool creates a pool of pre-mapped ANE requests for this kernel.
// All requests reference this kernel's IOSurfaces, enabling pipelined
// evaluation without per-eval request setup. Maximum depth is 127.
func (k *Kernel) NewRequestPool(depth int) (*KernelRequestPool, error) {
	if k == nil || k.closed() {
		return nil, fmt.Errorf("request pool: kernel is closed")
	}
	if depth < 1 {
		return nil, fmt.Errorf("request pool: depth must be >= 1")
	}
	if depth > maxKernelPoolDepth {
		return nil, fmt.Errorf("request pool: depth %d exceeds maximum %d", depth, maxKernelPoolDepth)
	}

	if k.shared != nil {
		return k.newSharedMILPool(depth)
	}
	return k.newModelPool(depth)
}

func (k *Kernel) newSharedMILPool(depth int) (*KernelRequestPool, error) {
	h := k.shared
	requests := make([]appleneuralengine.ANERequest, depth)

	// First slot reuses the kernel's existing request.
	requests[0] = h.request

	for i := 1; i < depth; i++ {
		req, err := requestFromExistingSurfaces(h.inputs, h.outputs)
		if err != nil {
			return nil, fmt.Errorf("request pool: request[%d]: %w", i, err)
		}
		if err := h.program.mapRequest(req); err != nil {
			return nil, fmt.Errorf("request pool: map request[%d]: %w", i, err)
		}
		requests[i] = req
	}

	return &KernelRequestPool{
		kernel:   k,
		shared:   true,
		requests: requests,
		depth:    depth,
	}, nil
}

func (k *Kernel) newModelPool(depth int) (*KernelRequestPool, error) {
	m := k.k
	if m == nil {
		return nil, fmt.Errorf("request pool: no underlying model")
	}
	pool, err := xane.NewRequestPool(m, depth)
	if err != nil {
		return nil, fmt.Errorf("request pool: %w", err)
	}
	return &KernelRequestPool{
		kernel:   k,
		xanePool: pool,
		depth:    depth,
	}, nil
}

// Acquire returns the next request from the pool in round-robin order.
func (p *KernelRequestPool) Acquire() *KernelPooledRequest {
	if p.shared {
		idx := int(p.next.Add(1)-1) % p.depth
		return &KernelPooledRequest{
			pool:    p,
			request: p.requests[idx],
		}
	}
	pr := p.xanePool.Acquire()
	return &KernelPooledRequest{
		pool:   p,
		xanePR: pr,
	}
}

// Eval evaluates this request on the ANE.
func (r *KernelPooledRequest) Eval() error {
	if r.xanePR != nil {
		return r.xanePR.Eval()
	}
	prog := r.pool.kernel.shared.program
	if !r.request.Validate() {
		return fmt.Errorf("eval: pooled request validation failed")
	}
	ok, err := prog.inMemModel.EvaluateWithQoSOptionsRequestError(prog.qos, nil, r.request)
	if err == nil && ok {
		return nil
	}
	return fmt.Errorf("eval: %w", err)
}

// Release returns this request to the pool. Currently a no-op.
func (r *KernelPooledRequest) Release() {}

// Close unmaps all pooled requests except the kernel's original request.
func (p *KernelRequestPool) Close() error {
	if p.xanePool != nil {
		return p.xanePool.Close()
	}
	k := p.kernel
	for i := 1; i < p.depth; i++ {
		k.shared.program.mapMu.Lock()
		k.shared.program.inMemModel.UnmapIOSurfacesWithRequest(p.requests[i])
		k.shared.program.mapMu.Unlock()
	}
	return nil
}

// requestFromExistingSurfaces builds an ANERequest referencing existing
// IOSurfaces. This is the model-package equivalent of the ane package's
// createRequestFromSurfaces.
func requestFromExistingSurfaces(inputs, outputs []coregraphics.IOSurfaceRef) (appleneuralengine.ANERequest, error) {
	ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()

	inputArr := foundation.NewNSMutableArray()
	inputIdxArr := foundation.NewNSMutableArray()
	for i, ref := range inputs {
		inputArr.AddObject(ioClass.ObjectWithIOSurface(ref))
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	outputArr := foundation.NewNSMutableArray()
	outputIdxArr := foundation.NewNSMutableArray()
	for i, ref := range outputs {
		outputArr.AddObject(ioClass.ObjectWithIOSurface(ref))
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(0)
	txnHandle := foundation.GetNSNumberClass().NumberWithUnsignedLongLong(1)

	reqClass := appleneuralengine.GetANERequestClass()
	reqObj := reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(
		inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx, nil, txnHandle,
	)
	if reqObj == nil || reqObj.GetID() == 0 {
		reqObj = reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(
			inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx,
		)
	}
	if reqObj == nil || reqObj.GetID() == 0 {
		return appleneuralengine.ANERequest{}, fmt.Errorf("failed to create request")
	}
	req := appleneuralengine.ANERequestFromID(reqObj.GetID())
	if req.TransactionHandle().GetID() == 0 {
		req.SetTransactionHandle(txnHandle)
	}
	return req, nil
}
