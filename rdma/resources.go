package rdma

import (
	"errors"
	"fmt"
)

// Resources owns a simple ibverbs resource chain.
//
// Close tears resources down in reverse dependency order: QP, MR, CQ, PD, then
// context. Zero handles are ignored, and handles are cleared before provider
// teardown to make repeated Close calls harmless.
type Resources struct {
	Context RDMAContext
	PD      RDMAPD
	MR      RDMAMR
	CQ      RDMACQ
	QP      RDMAQP
}

// Close releases the resources in r.
func (r *Resources) Close() error {
	if r == nil {
		return nil
	}
	var errs []error
	if r.QP != 0 {
		qp := r.QP
		r.QP = 0
		rc, err := IbvDestroyQp(qp)
		if err := checkProviderReturn("ibv_destroy_qp", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.MR != 0 {
		mr := r.MR
		r.MR = 0
		rc, err := IbvDeregMr(mr)
		if err := checkProviderReturn("ibv_dereg_mr", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.CQ != 0 {
		cq := r.CQ
		r.CQ = 0
		rc, err := IbvDestroyCq(cq)
		if err := checkProviderReturn("ibv_destroy_cq", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.PD != 0 {
		pd := r.PD
		r.PD = 0
		rc, err := IbvDeallocPd(pd)
		if err := checkProviderReturn("ibv_dealloc_pd", rc, err); err != nil {
			errs = append(errs, err)
		}
	}
	if r.Context != 0 {
		context := r.Context
		r.Context = 0
		rc, err := IbvCloseDevice(context)
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
