// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	_ "github.com/ebitengine/purego"
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
	IBV_MTU_4096 = 5

	// enum ibv_port_state. A port carries traffic only in ACTIVE.
	IBV_PORT_NOP          = 0
	IBV_PORT_DOWN         = 1
	IBV_PORT_INIT         = 2
	IBV_PORT_ARMED        = 3
	IBV_PORT_ACTIVE       = 4
	IBV_PORT_ACTIVE_DEFER = 5

	// enum ibv_link_layer. Thunderbolt is an Apple addition and is why the
	// value is 100 rather than the next one in sequence.
	IBV_LINK_LAYER_UNSPECIFIED = 0
	IBV_LINK_LAYER_INFINIBAND  = 1
	IBV_LINK_LAYER_ETHERNET    = 2
	IBV_LINK_LAYER_THUNDERBOLT = 100

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
	_                [9]byte
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
//
// Every field is named. The struct has no interior padding: it is 52 bytes of
// 4-byte-aligned members, so anonymous filler would only hide fields callers
// need. State and LinkLayer in particular decide whether a port carries traffic
// and how a route GID must be selected, and a binding that cannot express those
// checks fails silently rather than loudly.
type IbvPortAttr struct {
	State         int32
	MaxMTU        int32
	ActiveMTU     int32
	GIDTblLen     uint32
	PortCapFlags  uint32
	MaxMsgSz      uint32
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
}

const (
	ibvQPNumOffset  = 52
	ibvMRLKeyOffset = 36
	ibvMRRKeyOffset = 40
)

func Ibv_qp_num(qp RDMAQP) uint32 {
	if qp == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(qp + ibvQPNumOffset))
}

func Ibv_mr_lkey(mr RDMAMR) uint32 {
	if mr == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(mr + ibvMRLKeyOffset))
}

func Ibv_mr_rkey(mr RDMAMR) uint32 {
	if mr == 0 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(mr + ibvMRRKeyOffset))
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

var ibvFuncMu sync.Mutex

// The datapath verbs reach the provider through rdmaCall3 rather than a purego
// closure. Registering a closure per function pointer costs an allocation on
// every call in the hottest path in the package, which is what
// TestDatapathWrappersCall3Allocs guards against.
//
// rdmaCall3Args mirrors purego's syscallArgs. The purego trampoline reads it by
// offset, so TestRDMACall3ABI guards its layout.
type rdmaCall3Args struct {
	fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15                uintptr
	a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31, a32 uintptr
	f1, f2, f3, f4, f5, f6, f7, f8                                                      uintptr
	arm64R8                                                                             uintptr
}

var rdmaCall3ArgsPool = sync.Pool{New: func() any { return new(rdmaCall3Args) }}

//go:linkname rdmaRuntimeCgocall runtime.cgocall
func rdmaRuntimeCgocall(fn uintptr, arg unsafe.Pointer) int32

//go:linkname rdmaSyscallXABI0 github.com/ebitengine/purego.syscallXABI0
var rdmaSyscallXABI0 uintptr

func init() {
	if rdmaSyscallXABI0 == 0 {
		panic("rdma: purego syscall trampoline is unavailable")
	}
}

// rdmaCall3 invokes an RDMA context operation with its three machine-word arguments.
//
//go:uintptrescapes
func rdmaCall3(fn, a1, a2, a3 uintptr) uintptr {
	s := rdmaCall3ArgsPool.Get().(*rdmaCall3Args)
	*s = rdmaCall3Args{fn: fn, a1: a1, a2: a2, a3: a3, f1: a1, f2: a2, f3: a3}
	rdmaRuntimeCgocall(rdmaSyscallXABI0, unsafe.Pointer(s))
	r1 := s.a1
	rdmaCall3ArgsPool.Put(s)
	return r1
}

type IbvCQPoller struct {
	cq    RDMACQ
	fnPtr uintptr
}

func NewIbvCQPoller(cq RDMACQ) (IbvCQPoller, error) {
	fnPtr := rdmaContextOp(rdmaContextFromCQ(cq), 88)
	if fnPtr == 0 {
		return IbvCQPoller{}, fmt.Errorf("rdma: ibv_poll_cq unavailable")
	}
	return IbvCQPoller{cq: cq, fnPtr: fnPtr}, nil
}

// Poll polls the completion queue. It reports -1 without calling the provider
// when the poller is the zero value, which happens when NewIbvCQPoller returned
// an error the caller ignored, or when the work-completion pointer is nil.
//
// The pointer checks are not defensive style: the provider dereferences what it
// is given, so a nil here faults inside C, where the SIGSEGV arrives during cgo
// execution and no Go recover can reach it.
func (p IbvCQPoller) Poll(numEntries int, wc *IbvWC) int {
	if p.fnPtr == 0 || p.cq == 0 || wc == nil {
		return -1
	}
	r1 := rdmaCall3(p.fnPtr, uintptr(p.cq), uintptr(numEntries), uintptr(unsafe.Pointer(wc)))
	runtime.KeepAlive(wc)
	return int(r1)
}

type IbvQPPoster struct {
	qp      RDMAQP
	sendPtr uintptr
	recvPtr uintptr
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
		qp:      qp,
		sendPtr: sendPtr,
		recvPtr: recvPtr,
	}, nil
}

// PostSend posts a send work request. It reports -1 without calling the
// provider when the poster is the zero value, or when either work-request
// pointer is nil.
//
// badWR is checked along with the rest. The provider writes the offending work
// request through it on the failure path, so it is dereferenced exactly when a
// post has already gone wrong; a nil there would turn a recoverable failure
// into a killed process.
func (p IbvQPPoster) PostSend(wr *IbvSendWR, badWR **IbvSendWR) int {
	if p.sendPtr == 0 || p.qp == 0 || wr == nil || badWR == nil {
		return -1
	}
	r1 := rdmaCall3(p.sendPtr, uintptr(p.qp), uintptr(unsafe.Pointer(wr)), uintptr(unsafe.Pointer(badWR)))
	runtime.KeepAlive(wr)
	runtime.KeepAlive(badWR)
	return int(r1)
}

// PostRecv posts a receive work request. It reports -1 without calling the
// provider when the poster is the zero value, or when either work-request
// pointer is nil, as in PostSend.
func (p IbvQPPoster) PostRecv(wr *IbvRecvWR, badWR **IbvRecvWR) int {
	if p.recvPtr == 0 || p.qp == 0 || wr == nil || badWR == nil {
		return -1
	}
	r1 := rdmaCall3(p.recvPtr, uintptr(p.qp), uintptr(unsafe.Pointer(wr)), uintptr(unsafe.Pointer(badWR)))
	runtime.KeepAlive(wr)
	runtime.KeepAlive(badWR)
	return int(r1)
}

// Ibv_poll_cq calls the SDK inline ibv_poll_cq wrapper through ibv_context_ops.
func Ibv_poll_cq(cq RDMACQ, numEntries int, wc *IbvWC) (int, error) {
	fnPtr := rdmaContextOp(rdmaContextFromCQ(cq), 88)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_poll_cq unavailable")
	}
	r1 := rdmaCall3(fnPtr, uintptr(cq), uintptr(numEntries), uintptr(unsafe.Pointer(wc)))
	runtime.KeepAlive(wc)
	return int(r1), nil
}

// Ibv_post_send calls the SDK inline ibv_post_send wrapper through ibv_context_ops.
func Ibv_post_send(qp RDMAQP, wr *IbvSendWR, badWR **IbvSendWR) (int, error) {
	fnPtr := rdmaContextOp(rdmaContextFromQP(qp), 200)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_post_send unavailable")
	}
	r1 := rdmaCall3(fnPtr, uintptr(qp), uintptr(unsafe.Pointer(wr)), uintptr(unsafe.Pointer(badWR)))
	runtime.KeepAlive(wr)
	runtime.KeepAlive(badWR)
	return int(r1), nil
}

// Ibv_post_recv calls the SDK inline ibv_post_recv wrapper through ibv_context_ops.
func Ibv_post_recv(qp RDMAQP, wr *IbvRecvWR, badWR **IbvRecvWR) (int, error) {
	fnPtr := rdmaContextOp(rdmaContextFromQP(qp), 208)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_post_recv unavailable")
	}
	r1 := rdmaCall3(fnPtr, uintptr(qp), uintptr(unsafe.Pointer(wr)), uintptr(unsafe.Pointer(badWR)))
	runtime.KeepAlive(wr)
	runtime.KeepAlive(badWR)
	return int(r1), nil
}
