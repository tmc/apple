package rdma

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// RDMACompChannel is an ibv_comp_channel handle.
type RDMACompChannel = uintptr

var _ibvCreateCompChannel func(context RDMAContext) RDMACompChannel
var _ibvCreateCompChannelErr error

var _ibvDestroyCompChannel func(channel RDMACompChannel) int
var _ibvDestroyCompChannelErr error

var _ibvGetCqEvent func(channel RDMACompChannel, cq uintptr, cqContext uintptr) int
var _ibvGetCqEventErr error

var _ibvAckCqEvents func(cq RDMACQ, nevents uint32)
var _ibvAckCqEventsErr error

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_ibvCreateCompChannel, &_ibvCreateCompChannelErr, frameworkHandle, "ibv_create_comp_channel", "")
	registerFunc(&_ibvDestroyCompChannel, &_ibvDestroyCompChannelErr, frameworkHandle, "ibv_destroy_comp_channel", "")
	registerFunc(&_ibvGetCqEvent, &_ibvGetCqEventErr, frameworkHandle, "ibv_get_cq_event", "")
	registerFunc(&_ibvAckCqEvents, &_ibvAckCqEventsErr, frameworkHandle, "ibv_ack_cq_events", "")
}

func tryIbvCreateCompChannel(context RDMAContext) (RDMACompChannel, error) {
	if _ibvCreateCompChannel == nil {
		return 0, symbolCallError("ibv_create_comp_channel", "", _ibvCreateCompChannelErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_create_comp_channel", "context")
	}
	channel, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMACompChannel {
		return _ibvCreateCompChannel(context)
	})
	rdmaKeepAlive(context)
	if channel == 0 {
		return 0, rdmaNilProviderResultError("ibv_create_comp_channel", "completion channel", 0, errno, errnoSet, context, true)
	}
	return channel, nil
}

// IbvCreateCompChannel creates a completion event channel for context.
func IbvCreateCompChannel(context RDMAContext) (RDMACompChannel, error) {
	return tryIbvCreateCompChannel(context)
}

func tryIbvDestroyCompChannel(channel RDMACompChannel) (int, error) {
	if _ibvDestroyCompChannel == nil {
		return 0, symbolCallError("ibv_destroy_comp_channel", "", _ibvDestroyCompChannelErr)
	}
	if channel == 0 {
		return 0, rdmaNilHandleError("ibv_destroy_comp_channel", "completion channel")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvDestroyCompChannel(channel)
	})
	rdmaKeepAlive(channel)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_destroy_comp_channel", rc, errno, errnoSet, rdmaContextFromCompChannel(channel), true)
	}
	return rc, nil
}

// IbvDestroyCompChannel destroys a completion event channel.
func IbvDestroyCompChannel(channel RDMACompChannel) (int, error) {
	return tryIbvDestroyCompChannel(channel)
}

func tryIbvGetCqEvent(channel RDMACompChannel, cq *RDMACQ, cqContext *uintptr) (int, error) {
	if _ibvGetCqEvent == nil {
		return 0, symbolCallError("ibv_get_cq_event", "", _ibvGetCqEventErr)
	}
	if channel == 0 {
		return 0, rdmaNilHandleError("ibv_get_cq_event", "completion channel")
	}
	if cq == nil {
		return 0, rdmaNilPointerError("ibv_get_cq_event", "completion queue")
	}
	if cqContext == nil {
		return 0, rdmaNilPointerError("ibv_get_cq_event", "completion queue context")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvGetCqEvent(channel, uintptr(unsafe.Pointer(cq)), uintptr(unsafe.Pointer(cqContext)))
	})
	rdmaKeepAlive(channel)
	rdmaKeepAlive(cq)
	rdmaKeepAlive(cqContext)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_get_cq_event", rc, errno, errnoSet, rdmaContextFromCompChannel(channel), true)
	}
	return rc, nil
}

// IbvGetCqEvent reads the next completion event from channel.
func IbvGetCqEvent(channel RDMACompChannel, cq *RDMACQ, cqContext *uintptr) (int, error) {
	return tryIbvGetCqEvent(channel, cq, cqContext)
}

func tryIbvAckCqEvents(cq RDMACQ, nevents uint32) error {
	if _ibvAckCqEvents == nil {
		return symbolCallError("ibv_ack_cq_events", "", _ibvAckCqEventsErr)
	}
	if cq == 0 {
		return rdmaNilHandleError("ibv_ack_cq_events", "completion queue")
	}
	rdmaProviderCall0(func() {
		_ibvAckCqEvents(cq, nevents)
	})
	rdmaKeepAlive(cq)
	return nil
}

// IbvAckCqEvents acknowledges completion events returned by IbvGetCqEvent.
func IbvAckCqEvents(cq RDMACQ, nevents uint32) error {
	return tryIbvAckCqEvents(cq, nevents)
}

type ibvReqNotifyCQFunc func(RDMACQ, int) int

var ibvReqNotifyCQFuncs sync.Map // map[uintptr]ibvReqNotifyCQFunc

func rdmaReqNotifyCQFunc(fnPtr uintptr) ibvReqNotifyCQFunc {
	if fn, ok := ibvReqNotifyCQFuncs.Load(fnPtr); ok {
		return fn.(ibvReqNotifyCQFunc)
	}
	ibvFuncMu.Lock()
	defer ibvFuncMu.Unlock()
	if fn, ok := ibvReqNotifyCQFuncs.Load(fnPtr); ok {
		return fn.(ibvReqNotifyCQFunc)
	}
	var fn ibvReqNotifyCQFunc
	purego.RegisterFunc(&fn, fnPtr)
	actual, _ := ibvReqNotifyCQFuncs.LoadOrStore(fnPtr, fn)
	return actual.(ibvReqNotifyCQFunc)
}

func tryIbvReqNotifyCq(cq RDMACQ, solicitedOnly int) (int, error) {
	if cq == 0 {
		return 0, rdmaNilHandleError("ibv_req_notify_cq", "completion queue")
	}
	fnPtr := rdmaContextOp(rdmaContextFromCQ(cq), 96)
	if fnPtr == 0 {
		return 0, fmt.Errorf("rdma: ibv_req_notify_cq unavailable")
	}
	fn := rdmaReqNotifyCQFunc(fnPtr)
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return fn(cq, solicitedOnly)
	})
	rdmaKeepAlive(cq)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_req_notify_cq", rc, errno, errnoSet, rdmaContextFromCQ(cq), true)
	}
	return rc, nil
}

// IbvReqNotifyCq requests a completion notification for cq.
func IbvReqNotifyCq(cq RDMACQ, solicitedOnly int) (int, error) {
	return tryIbvReqNotifyCq(cq, solicitedOnly)
}

// IbvCompChannelFD returns the file descriptor stored in channel.
func IbvCompChannelFD(channel RDMACompChannel) int {
	if channel == 0 {
		return -1
	}
	return int(*(*int32)(unsafe.Add(unsafe.Pointer(channel), 8)))
}

// IbvCQChannel returns the completion channel associated with cq, if any.
func IbvCQChannel(cq RDMACQ) RDMACompChannel {
	if cq == 0 {
		return 0
	}
	return *(*RDMACompChannel)(unsafe.Add(unsafe.Pointer(cq), 8))
}

func rdmaContextFromCompChannel(channel RDMACompChannel) RDMAContext {
	if channel == 0 {
		return 0
	}
	return *(*RDMAContext)(unsafe.Pointer(channel))
}
