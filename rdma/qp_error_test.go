package rdma

import (
	"errors"
	"strings"
	"testing"
	"unsafe"
)

func TestModifyQPErrorIncludesRTRDiagnostics(t *testing.T) {
	var raw [64]byte
	*(*RDMAContext)(unsafe.Pointer(&raw[0])) = 123
	*(*uint32)(unsafe.Pointer(&raw[52])) = 777
	qp := RDMAQP(uintptr(unsafe.Pointer(&raw[0])))
	rdmaRememberContext(123, "rdma_en_test")
	t.Cleanup(func() { rdmaForgetContext(123) })

	attr := IbvQPAttr{
		QPState:   IBV_QPS_RTR,
		PathMTU:   5,
		RQPSN:     7,
		DestQPNum: 42,
		AHAttr: IbvAHAttr{
			DLID:     2,
			PortNum:  1,
			IsGlobal: 1,
			GRH: IbvGlobalRoute{
				SGIDIndex: 3,
				DGID:      IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
	}
	mask := IBV_QP_STATE | IBV_QP_AV | IBV_QP_PATH_MTU | IBV_QP_DEST_QPN | IBV_QP_RQ_PSN
	err := NewModifyQPError(qp, &attr, mask, 60, nil)
	if !errors.Is(err, ErrProviderStatus) {
		t.Fatalf("NewModifyQPError error = %v, want ErrProviderStatus", err)
	}
	msg := err.Error()
	for _, want := range []string{
		"INIT->RTR",
		"rdma_en_test",
		"qpn=777",
		"errno 60 (ETIMEDOUT)",
		"IBV_QP_DEST_QPN",
		"dest_qpn=42",
		"rq_psn=7",
		"gid_index=3",
		"gid=fe800000000000000102030405060708",
		"path_mtu=4096",
		"hint:",
	} {
		if !strings.Contains(msg, want) {
			t.Fatalf("ModifyQPError missing %q in %q", want, msg)
		}
	}
}

func TestQPAttrMaskNames(t *testing.T) {
	mask := IBV_QP_STATE | IBV_QP_AV | IBV_QP_DEST_QPN | 0x40000000
	got := strings.Join(QPAttrMaskNames(mask), ",")
	want := "IBV_QP_STATE,IBV_QP_AV,IBV_QP_DEST_QPN,unknown(0x40000000)"
	if got != want {
		t.Fatalf("QPAttrMaskNames = %q, want %q", got, want)
	}
}
