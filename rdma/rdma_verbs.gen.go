// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	IBV_ACCESS_LOCAL_WRITE  = 1
	IBV_ACCESS_REMOTE_WRITE = 2
	IBV_ACCESS_REMOTE_READ  = 4

	IBV_QPT_UC = 3

	IBV_QPS_RESET = 0
	IBV_QPS_INIT  = 1
	IBV_QPS_RTR   = 2
	IBV_QPS_RTS   = 3
	IBV_QPS_ERR   = 6

	IBV_QP_STATE        = 1
	IBV_QP_ACCESS_FLAGS = 8
	IBV_QP_PKEY_INDEX   = 16
	IBV_QP_PORT         = 32
	IBV_QP_AV           = 128
	IBV_QP_PATH_MTU     = 256
	IBV_QP_RQ_PSN       = 4096
	IBV_QP_SQ_PSN       = 65536
	IBV_QP_DEST_QPN     = 1048576

	IBV_MTU_1024 = 3

	IBV_WR_RDMA_WRITE = 0
	IBV_WR_SEND       = 2
	IBV_SEND_SIGNALED = 2

	IBV_WC_SUCCESS = 0
)

// IbvGID is the raw 16-byte representation of union ibv_gid.
type IbvGID [16]byte

// IbvQPCap matches struct ibv_qp_cap.
type IbvQPCap struct {
	MaxSendWR     uint32
	MaxRecvWR     uint32
	MaxSendSGE    uint32
	MaxRecvSGE    uint32
	MaxInlineData uint32
}

// IbvQPInitAttr matches struct ibv_qp_init_attr.
type IbvQPInitAttr struct {
	QPContext uintptr
	SendCQ    RDMACQ
	RecvCQ    RDMACQ
	SRQ       uintptr
	Cap       IbvQPCap
	QPType    int32
	SQSigAll  int32
}

// IbvGlobalRoute matches struct ibv_global_route.
type IbvGlobalRoute struct {
	DGID         IbvGID
	FlowLabel    uint32
	SGIDIndex    uint8
	HopLimit     uint8
	TrafficClass uint8
	_            uint8
}

// IbvAHAttr matches struct ibv_ah_attr.
type IbvAHAttr struct {
	GRH         IbvGlobalRoute
	DLID        uint16
	SL          uint8
	SrcPathBits uint8
	StaticRate  uint8
	IsGlobal    uint8
	PortNum     uint8
	_           uint8
}

// IbvQPAttr matches the fields commonly used with ibv_modify_qp.
type IbvQPAttr struct {
	QPState          int32
	CurQPState       int32
	PathMTU          int32
	PathMigState     int32
	QKey             uint32
	RQPSN            uint32
	SQPSN            uint32
	DestQPNum        uint32
	QPAccessFlags    int32
	Cap              IbvQPCap
	AHAttr           IbvAHAttr
	AltAHAttr        IbvAHAttr
	PKeyIndex        uint16
	AltPKeyIndex     uint16
	EnSQDAsyncNotify uint8
	SQDraining       uint8
	MaxRDAtomic      uint8
	MaxDestRDAtomic  uint8
	MinRNRTimer      uint8
	PortNum          uint8
	Timeout          uint8
	RetryCnt         uint8
	RNRetry          uint8
	AltPortNum       uint8
	AltTimeout       uint8
	_                uint8
	RateLimit        uint32
	_                [4]byte
}

// IbvSGE matches struct ibv_sge.
type IbvSGE struct {
	Addr   uint64
	Length uint32
	LKey   uint32
}

// IbvSendWR matches struct ibv_send_wr for send operations.
type IbvSendWR struct {
	WRID      uint64
	Next      *IbvSendWR
	SGList    *IbvSGE
	NumSGE    int32
	Opcode    int32
	SendFlags int32
	ImmData   uint32
	WR        [32]byte
	QPType    [8]byte
	BindMW    [48]byte
}

// IbvRecvWR matches struct ibv_recv_wr.
type IbvRecvWR struct {
	WRID   uint64
	Next   *IbvRecvWR
	SGList *IbvSGE
	NumSGE int32
	_      int32
}

// IbvWC matches struct ibv_wc.
type IbvWC struct {
	WRID         uint64
	Status       int32
	Opcode       int32
	VendorErr    uint32
	ByteLen      uint32
	ImmData      uint32
	QPNum        uint32
	SrcQP        uint32
	WCFlags      int32
	PKeyIndex    uint16
	SLID         uint16
	SL           uint8
	DLIDPathBits uint8
	PortNum      uint8
	_            uint8
}

// IbvPortAttr matches struct ibv_port_attr.
type IbvPortAttr struct {
	State         int32
	MaxMTU        int32
	ActiveMTU     int32
	GIDTblLen     int32
	PortCapFlags  uint32
	MaxMsgSZ      uint32
	BadPKeyCntr   uint32
	QKeyViolCntr  uint32
	PKeyTblLen    uint16
	LID           uint16
	SMLID         uint16
	LMC           uint8
	MaxVLNum      uint8
	SMSL          uint8
	SubnetTimeout uint8
	InitTypeReply uint8
	ActiveWidth   uint8
	ActiveSpeed   uint8
	PhysState     uint8
	LinkLayer     uint8
	Flags         uint8
	PortCapFlags2 uint16
	_             uint16
}

func Ibv_qp_num(qp RDMAQP) uint32 {
	if qp == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(qp + 52))
}

func Ibv_mr_lkey(mr RDMAMR) uint32 {
	if mr == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(mr + 36))
}

func Ibv_mr_rkey(mr RDMAMR) uint32 {
	if mr == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(mr + 40))
}

func rdmaContextFromCQ(cq RDMACQ) RDMAContext {
	if cq == 0 {
		return 0
	}
	return *(*RDMAContext)(unsafe.Pointer(cq))
}

func rdmaContextFromQP(qp RDMAQP) RDMAContext {
	if qp == 0 {
		return 0
	}
	return *(*RDMAContext)(unsafe.Pointer(qp))
}

func rdmaContextOp(context RDMAContext, off uintptr) uintptr {
	if context == 0 {
		return 0
	}
	return *(*uintptr)(unsafe.Pointer(context + 8 + off))
}

type ibvPollCQFunc func(RDMACQ, int, *IbvWC) int
type ibvPostSendFunc func(RDMAQP, *IbvSendWR, **IbvSendWR) int
type ibvPostRecvFunc func(RDMAQP, *IbvRecvWR, **IbvRecvWR) int

var (
	ibvPollCQFuncs   sync.Map // map[uintptr]ibvPollCQFunc
	ibvPostSendFuncs sync.Map // map[uintptr]ibvPostSendFunc
	ibvPostRecvFuncs sync.Map // map[uintptr]ibvPostRecvFunc
	ibvFuncMu        sync.Mutex
)

func rdmaPollCQFunc(fnPtr uintptr) ibvPollCQFunc {
	if fn, ok := ibvPollCQFuncs.Load(fnPtr); ok {
		return fn.(ibvPollCQFunc)
	}
	ibvFuncMu.Lock()
	defer ibvFuncMu.Unlock()
	if fn, ok := ibvPollCQFuncs.Load(fnPtr); ok {
		return fn.(ibvPollCQFunc)
	}
	var fn ibvPollCQFunc
	purego.RegisterFunc(&fn, fnPtr)
	actual, _ := ibvPollCQFuncs.LoadOrStore(fnPtr, fn)
	return actual.(ibvPollCQFunc)
}

func rdmaPostSendFunc(fnPtr uintptr) ibvPostSendFunc {
	if fn, ok := ibvPostSendFuncs.Load(fnPtr); ok {
		return fn.(ibvPostSendFunc)
	}
	ibvFuncMu.Lock()
	defer ibvFuncMu.Unlock()
	if fn, ok := ibvPostSendFuncs.Load(fnPtr); ok {
		return fn.(ibvPostSendFunc)
	}
	var fn ibvPostSendFunc
	purego.RegisterFunc(&fn, fnPtr)
	actual, _ := ibvPostSendFuncs.LoadOrStore(fnPtr, fn)
	return actual.(ibvPostSendFunc)
}

func rdmaPostRecvFunc(fnPtr uintptr) ibvPostRecvFunc {
	if fn, ok := ibvPostRecvFuncs.Load(fnPtr); ok {
		return fn.(ibvPostRecvFunc)
	}
	ibvFuncMu.Lock()
	defer ibvFuncMu.Unlock()
	if fn, ok := ibvPostRecvFuncs.Load(fnPtr); ok {
		return fn.(ibvPostRecvFunc)
	}
	var fn ibvPostRecvFunc
	purego.RegisterFunc(&fn, fnPtr)
	actual, _ := ibvPostRecvFuncs.LoadOrStore(fnPtr, fn)
	return actual.(ibvPostRecvFunc)
}

type IbvCQPoller struct {
	cq RDMACQ
	fn ibvPollCQFunc
}

func NewIbvCQPoller(cq RDMACQ) (IbvCQPoller, error) {
	fnPtr := rdmaContextOp(rdmaContextFromCQ(cq), 88)
	if fnPtr == 0 {
		return IbvCQPoller{}, fmt.Errorf("rdma: ibv_poll_cq unavailable")
	}
	return IbvCQPoller{cq: cq, fn: rdmaPollCQFunc(fnPtr)}, nil
}

func (p IbvCQPoller) Poll(numEntries int, wc *IbvWC) int {
	if p.cq == 0 || p.fn == nil || wc == nil {
		return -1
	}
	rc := rdmaProviderCall(func() int {
		return p.fn(p.cq, numEntries, wc)
	})
	rdmaKeepAlive(p.cq)
	rdmaKeepAlive(wc)
	return rc
}

type IbvQPPoster struct {
	qp   RDMAQP
	send ibvPostSendFunc
	recv ibvPostRecvFunc
}

func NewIbvQPPoster(qp RDMAQP) (IbvQPPoster, error) {
	context := rdmaContextFromQP(qp)
	sendPtr := rdmaContextOp(context, 200)
	if sendPtr == 0 {
		return IbvQPPoster{}, fmt.Errorf("rdma: ibv_post_send unavailable")
	}
	recvPtr := rdmaContextOp(context, 208)
	if recvPtr == 0 {
		return IbvQPPoster{}, fmt.Errorf("rdma: ibv_post_recv unavailable")
	}
	return IbvQPPoster{
		qp:   qp,
		send: rdmaPostSendFunc(sendPtr),
		recv: rdmaPostRecvFunc(recvPtr),
	}, nil
}

func (p IbvQPPoster) PostSend(wr *IbvSendWR, badWR **IbvSendWR) int {
	if p.qp == 0 || p.send == nil || wr == nil || badWR == nil {
		return -1
	}
	rc := rdmaProviderCall(func() int {
		return p.send(p.qp, wr, badWR)
	})
	rdmaKeepAlive(p.qp)
	rdmaKeepAlive(wr)
	rdmaKeepAlive(badWR)
	return rc
}

func (p IbvQPPoster) PostRecv(wr *IbvRecvWR, badWR **IbvRecvWR) int {
	if p.qp == 0 || p.recv == nil || wr == nil || badWR == nil {
		return -1
	}
	rc := rdmaProviderCall(func() int {
		return p.recv(p.qp, wr, badWR)
	})
	rdmaKeepAlive(p.qp)
	rdmaKeepAlive(wr)
	rdmaKeepAlive(badWR)
	return rc
}

// Ibv_poll_cq calls the SDK inline ibv_poll_cq wrapper through ibv_context_ops.
func Ibv_poll_cq(cq RDMACQ, numEntries int, wc *IbvWC) (int, error) {
	if wc == nil {
		return 0, rdmaNilPointerError("ibv_poll_cq", "work completion")
	}
	fnPtr := rdmaContextOp(rdmaContextFromCQ(cq), 88)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_poll_cq unavailable")
	}
	fn := rdmaPollCQFunc(fnPtr)
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return fn(cq, numEntries, wc)
	})
	rdmaKeepAlive(cq)
	rdmaKeepAlive(wc)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_poll_cq", rc, errno, errnoSet, rdmaContextFromCQ(cq), true)
	}
	return rc, nil
}

// Ibv_post_send calls the SDK inline ibv_post_send wrapper through ibv_context_ops.
func Ibv_post_send(qp RDMAQP, wr *IbvSendWR, badWR **IbvSendWR) (int, error) {
	if wr == nil {
		return 0, rdmaNilPointerError("ibv_post_send", "send work request")
	}
	if badWR == nil {
		return 0, rdmaNilPointerError("ibv_post_send", "bad send work request")
	}
	fnPtr := rdmaContextOp(rdmaContextFromQP(qp), 200)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_post_send unavailable")
	}
	fn := rdmaPostSendFunc(fnPtr)
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return fn(qp, wr, badWR)
	})
	rdmaKeepAlive(qp)
	rdmaKeepAlive(wr)
	rdmaKeepAlive(badWR)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_post_send", rc, errno, errnoSet, rdmaContextFromQP(qp), true)
	}
	return rc, nil
}

// Ibv_post_recv calls the SDK inline ibv_post_recv wrapper through ibv_context_ops.
func Ibv_post_recv(qp RDMAQP, wr *IbvRecvWR, badWR **IbvRecvWR) (int, error) {
	if wr == nil {
		return 0, rdmaNilPointerError("ibv_post_recv", "receive work request")
	}
	if badWR == nil {
		return 0, rdmaNilPointerError("ibv_post_recv", "bad receive work request")
	}
	fnPtr := rdmaContextOp(rdmaContextFromQP(qp), 208)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_post_recv unavailable")
	}
	fn := rdmaPostRecvFunc(fnPtr)
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return fn(qp, wr, badWR)
	})
	rdmaKeepAlive(qp)
	rdmaKeepAlive(wr)
	rdmaKeepAlive(badWR)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_post_recv", rc, errno, errnoSet, rdmaContextFromQP(qp), true)
	}
	return rc, nil
}
