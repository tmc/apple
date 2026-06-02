package rdma

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/rdma"
)

// QPChain owns a simple ibverbs resource chain.
//
// Close tears resources down in reverse dependency order: QP, MR, CQ, PD,
// then context. Zero handles are ignored, and handles are cleared before
// provider teardown to make repeated Close calls harmless.
type QPChain struct {
	Context rdma.RDMAContext
	PD      rdma.RDMAPD
	MR      rdma.RDMAMR
	CQ      rdma.RDMACQ
	QP      rdma.RDMAQP
}

// Teardown verbs, indirected so tests can record call order without a provider.
var (
	destroyQP    = rdma.IbvDestroyQp
	deregMR      = rdma.IbvDeregMr
	destroyCQ    = rdma.IbvDestroyCq
	deallocPD    = rdma.IbvDeallocPd
	closeContext = rdma.IbvCloseDevice
)

// Close releases the resources in r.
func (r *QPChain) Close() error {
	if r == nil {
		return nil
	}
	var errs []error
	if r.QP != 0 {
		qp := r.QP
		r.QP = 0
		rc, err := destroyQP(qp)
		if err := checkProviderReturn("ibv_destroy_qp", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.MR != 0 {
		mr := r.MR
		r.MR = 0
		rc, err := deregMR(mr)
		if err := checkProviderReturn("ibv_dereg_mr", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.CQ != 0 {
		cq := r.CQ
		r.CQ = 0
		rc, err := destroyCQ(cq)
		if err := checkProviderReturn("ibv_destroy_cq", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.PD != 0 {
		pd := r.PD
		r.PD = 0
		rc, err := deallocPD(pd)
		if err := checkProviderReturn("ibv_dealloc_pd", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.Context != 0 {
		context := r.Context
		r.Context = 0
		rc, err := closeContext(context)
		if err := checkProviderReturn("ibv_close_device", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func checkProviderReturn(name string, rc int, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if rc != 0 {
		return fmt.Errorf("%s: rc=%d", name, rc)
	}
	return nil
}
