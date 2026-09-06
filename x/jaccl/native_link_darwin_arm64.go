//go:build darwin && arm64

package jaccl

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/tmc/apple/rdma"
)

// nativeSlotBytes matches standalone JACCL's FRAME_SIZE. Apple UC accepts
// these page-sized receive frames; larger payloads are split by the backend.
const nativeSlotBytes = 4 << 10

// nativeLink owns one UC QP and its persistent registered send and receive
// buffers. Descriptors are link fields because the provider may read them for
// the whole verbs call; they must not be per-operation stack values.
type nativeLink struct {
	device *nativeDevice
	peer   int

	cq     rdma.RDMACQ
	qp     rdma.RDMAQP
	sendMR rdma.RDMAMR
	recvMR rdma.RDMAMR

	send []byte
	recv []byte

	sendLKey uint32
	recvLKey uint32
	poller   rdma.IbvCQPoller
	poster   rdma.IbvQPPoster

	sendSGE rdma.IbvSGE
	sendWR  rdma.IbvSendWR
	sendBad *rdma.IbvSendWR
	recvSGE rdma.IbvSGE
	recvWR  rdma.IbvRecvWR
	recvBad *rdma.IbvRecvWR
	pollWC  [2 * nativeSlotDepth]rdma.IbvWC

	cleanup closeStack
}

func newNativeLink(device *nativeDevice, peer int) (*nativeLink, error) {
	if device == nil || device.context == 0 || device.pd == 0 {
		return nil, fmt.Errorf("create link to rank %d: %w", peer, ErrUnsupported)
	}
	if peer < 0 || peer > workIDPeerMax {
		return nil, fmt.Errorf("create link peer %d: %w", peer, ErrProtocol)
	}
	cq, err := rdma.IbvCreateCq(device.context, 2*nativeSlotDepth, 0, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("create completion queue: %w", err)
	}
	if cq == 0 {
		return nil, fmt.Errorf("create completion queue: %w", ErrUnsupported)
	}
	init := rdma.IbvQPInitAttr{
		SendCQ: cq,
		RecvCQ: cq,
		Cap: rdma.IbvQPCap{
			MaxSendWR:  nativeSlotDepth,
			MaxRecvWR:  nativeSlotDepth,
			MaxSendSGE: 1,
			MaxRecvSGE: 1,
		},
		QPType:   rdma.IBV_QPT_UC,
		SQSigAll: 1,
	}
	qp, err := rdma.IbvCreateQpAttr(device.pd, &init)
	if err != nil {
		_ = nativeCloseLink(0, 0, 0, nil, nil, cq)
		return nil, fmt.Errorf("create queue pair: %w", err)
	}
	if qp == 0 {
		_ = nativeCloseLink(0, 0, 0, nil, nil, cq)
		return nil, fmt.Errorf("create queue pair: %w", ErrUnsupported)
	}

	send, err := newNativeMemory(nativeSlotDepth * nativeSlotBytes)
	if err != nil {
		_ = nativeCloseLink(qp, 0, 0, nil, nil, cq)
		return nil, fmt.Errorf("allocate send memory: %w", err)
	}
	sendMR, err := registerNativeMemory(device.pd, send)
	if err != nil {
		_ = nativeCloseLink(qp, 0, 0, send, nil, cq)
		return nil, fmt.Errorf("register send memory: %w", err)
	}
	recv, err := newNativeMemory(nativeSlotDepth * nativeSlotBytes)
	if err != nil {
		_ = nativeCloseLink(qp, sendMR, 0, send, nil, cq)
		return nil, fmt.Errorf("allocate receive memory: %w", err)
	}
	recvMR, err := registerNativeMemory(device.pd, recv)
	if err != nil {
		_ = nativeCloseLink(qp, sendMR, 0, send, recv, cq)
		return nil, fmt.Errorf("register receive memory: %w", err)
	}

	poller, err := rdma.NewIbvCQPoller(cq)
	if err != nil {
		_ = nativeCloseLink(qp, sendMR, recvMR, send, recv, cq)
		return nil, fmt.Errorf("create completion poller: %w", err)
	}
	poster, err := rdma.NewIbvQPPoster(qp)
	if err != nil {
		_ = nativeCloseLink(qp, sendMR, recvMR, send, recv, cq)
		return nil, fmt.Errorf("create queue-pair poster: %w", err)
	}
	var cleanup closeStack
	cleanup.add(func() error {
		rc, err := rdma.IbvDestroyCq(cq)
		return nativeProviderRC("ibv_destroy_cq", int(rc), err)
	})
	cleanup.add(releaseNativeMemory(send))
	cleanup.add(nativeDeregister(sendMR))
	cleanup.add(releaseNativeMemory(recv))
	cleanup.add(nativeDeregister(recvMR))
	cleanup.add(func() error {
		rc, err := rdma.IbvDestroyQp(qp)
		return nativeProviderRC("ibv_destroy_qp", int(rc), err)
	})
	return &nativeLink{
		device:   device,
		peer:     peer,
		cq:       cq,
		qp:       qp,
		sendMR:   sendMR,
		recvMR:   recvMR,
		send:     send,
		recv:     recv,
		sendLKey: rdma.Ibv_mr_lkey(sendMR),
		recvLKey: rdma.Ibv_mr_lkey(recvMR),
		poller:   poller,
		poster:   poster,
		cleanup:  cleanup,
	}, nil
}

func nativeCloseLink(qp rdma.RDMAQP, sendMR, recvMR rdma.RDMAMR, send, recv []byte, cq rdma.RDMACQ) error {
	var errs []error
	if qp != 0 {
		rc, err := rdma.IbvDestroyQp(qp)
		errs = append(errs, nativeProviderRC("ibv_destroy_qp", int(rc), err))
	}
	if recvMR != 0 {
		errs = append(errs, nativeDeregister(recvMR)())
	}
	if recv != nil {
		errs = append(errs, releaseNativeMemory(recv)())
	}
	if sendMR != 0 {
		errs = append(errs, nativeDeregister(sendMR)())
	}
	if send != nil {
		errs = append(errs, releaseNativeMemory(send)())
	}
	if cq != 0 {
		rc, err := rdma.IbvDestroyCq(cq)
		errs = append(errs, nativeProviderRC("ibv_destroy_cq", int(rc), err))
	}
	return errors.Join(errs...)
}

// newNativeMemory allocates a page-backed buffer acceptable to the Apple RDMA
// provider. It must be released after its registered memory region is closed.
func newNativeMemory(length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("allocate memory length %d: %w", length, ErrProtocol)
	}
	buffer, err := syscall.Mmap(-1, 0, length, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, fmt.Errorf("mmap %d bytes: %w", length, err)
	}
	return buffer, nil
}

func releaseNativeMemory(buffer []byte) func() error {
	return func() error {
		if len(buffer) == 0 {
			return nil
		}
		if err := syscall.Munmap(buffer); err != nil {
			return fmt.Errorf("munmap %d bytes: %w", len(buffer), err)
		}
		return nil
	}
}

func registerNativeMemory(pd rdma.RDMAPD, buffer []byte) (rdma.RDMAMR, error) {
	if len(buffer) == 0 {
		return 0, fmt.Errorf("register zero-length memory: %w", ErrProtocol)
	}
	mr, err := rdma.IbvRegMr(pd, uintptr(unsafe.Pointer(unsafe.SliceData(buffer))), uintptr(len(buffer)), rdma.IBV_ACCESS_LOCAL_WRITE)
	runtime.KeepAlive(buffer)
	if err != nil {
		return 0, err
	}
	if mr == 0 {
		return 0, fmt.Errorf("provider returned zero memory region: %w", ErrUnsupported)
	}
	return mr, nil
}

func nativeDeregister(mr rdma.RDMAMR) func() error {
	return func() error {
		rc, err := rdma.IbvDeregMr(mr)
		return nativeProviderRC("ibv_dereg_mr", int(rc), err)
	}
}

// Close destroys the QP before either MR, then its CQ. The shared PD and
// context belong to nativeDevice and close only after all links are closed.
func (l *nativeLink) Close() error {
	if l == nil {
		return nil
	}
	err := l.cleanup.close()
	l.cq = 0
	l.qp = 0
	l.sendMR = 0
	l.recvMR = 0
	l.send = nil
	l.recv = nil
	l.sendLKey = 0
	l.recvLKey = 0
	l.poller = rdma.IbvCQPoller{}
	l.poster = rdma.IbvQPPoster{}
	l.sendBad = nil
	l.recvBad = nil
	l.device = nil
	return err
}

func (l *nativeLink) postRecv(slot, length int, id uint64) error {
	if err := l.slotRange(slot, length); err != nil {
		return fmt.Errorf("post receive: %w", err)
	}
	offset := slot * nativeSlotBytes
	l.recvSGE = rdma.IbvSGE{
		Addr:   uint64(uintptr(unsafe.Pointer(unsafe.SliceData(l.recv[offset:])))),
		Length: uint32(length),
		LKey:   l.recvLKey,
	}
	l.recvWR = rdma.IbvRecvWR{WRID: id, SGList: &l.recvSGE, NumSGE: 1}
	l.recvBad = nil
	rc := l.poster.PostRecv(&l.recvWR, &l.recvBad)
	runtime.KeepAlive(l)
	if rc != 0 {
		return fmt.Errorf("ibv_post_recv: rc=%d", rc)
	}
	return nil
}

func (l *nativeLink) postSend(slot int, payload []byte, id uint64) error {
	if err := l.slotRange(slot, len(payload)); err != nil {
		return fmt.Errorf("post send: %w", err)
	}
	offset := slot * nativeSlotBytes
	copy(l.send[offset:offset+len(payload)], payload)
	l.sendSGE = rdma.IbvSGE{
		Addr:   uint64(uintptr(unsafe.Pointer(unsafe.SliceData(l.send[offset:])))),
		Length: uint32(len(payload)),
		LKey:   l.sendLKey,
	}
	l.sendWR = rdma.IbvSendWR{
		WRID:      id,
		SGList:    &l.sendSGE,
		NumSGE:    1,
		Opcode:    rdma.IBV_WR_SEND,
		SendFlags: rdma.IBV_SEND_SIGNALED,
	}
	l.sendBad = nil
	rc := l.poster.PostSend(&l.sendWR, &l.sendBad)
	runtime.KeepAlive(l)
	if rc != 0 {
		return fmt.Errorf("ibv_post_send: rc=%d", rc)
	}
	return nil
}

func (l *nativeLink) waitCompletion(ctx context.Context, closed <-chan struct{}, want uint64, opcode int32, length int) error {
	if l == nil || l.poller == (rdma.IbvCQPoller{}) {
		return fmt.Errorf("poll completion: %w", ErrProtocol)
	}
	for {
		n := l.poller.Poll(len(l.pollWC), &l.pollWC[0])
		if n < 0 {
			return fmt.Errorf("ibv_poll_cq: rc=%d", n)
		}
		if n > len(l.pollWC) {
			return fmt.Errorf("ibv_poll_cq: returned %d for capacity %d: %w", n, len(l.pollWC), ErrProtocol)
		}
		for _, completion := range l.pollWC[:n] {
			if err := validateCompletion(completion, want, opcode, length); err != nil {
				return err
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-closed:
			return ErrClosed
		default:
			runtime.Gosched()
		}
	}
}

// waitPair waits for the one send and one receive completion posted for a
// one-slot UC transfer. Polling them together is important: either completion
// may arrive first, and treating the other as unexpected would turn normal CQ
// ordering into a protocol failure.
func (l *nativeLink) waitPair(ctx context.Context, closed <-chan struct{}, sendID, recvID uint64, recvLength int) error {
	if l == nil || l.poller == (rdma.IbvCQPoller{}) {
		return fmt.Errorf("poll completion pair: %w", ErrProtocol)
	}
	seenSend := false
	seenRecv := false
	for !seenSend || !seenRecv {
		n := l.poller.Poll(len(l.pollWC), &l.pollWC[0])
		if n < 0 {
			return fmt.Errorf("ibv_poll_cq: rc=%d", n)
		}
		if n > len(l.pollWC) {
			return fmt.Errorf("ibv_poll_cq: returned %d for capacity %d: %w", n, len(l.pollWC), ErrProtocol)
		}
		for _, completion := range l.pollWC[:n] {
			switch completion.WRID {
			case sendID:
				if seenSend {
					return fmt.Errorf("duplicate send completion id=%#x: %w", completion.WRID, ErrProtocol)
				}
				if err := validateCompletion(completion, sendID, rdma.IBV_WC_SEND, 0); err != nil {
					return err
				}
				seenSend = true
			case recvID:
				if seenRecv {
					return fmt.Errorf("duplicate receive completion id=%#x: %w", completion.WRID, ErrProtocol)
				}
				if err := validateCompletion(completion, recvID, rdma.IBV_WC_RECV, recvLength); err != nil {
					return err
				}
				seenRecv = true
			default:
				return fmt.Errorf("unexpected completion id=%#x opcode=%d status=%d: %w", completion.WRID, completion.Opcode, completion.Status, ErrProtocol)
			}
		}
		if seenSend && seenRecv {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-closed:
			return ErrClosed
		default:
			runtime.Gosched()
		}
	}
	return nil
}

func validateCompletion(completion rdma.IbvWC, want uint64, opcode int32, length int) error {
	if completion.WRID != want || completion.Opcode != opcode || completion.Status != rdma.IBV_WC_SUCCESS {
		return fmt.Errorf("completion id=%#x opcode=%d status=%d, want id=%#x opcode=%d success: %w", completion.WRID, completion.Opcode, completion.Status, want, opcode, ErrProtocol)
	}
	if opcode == rdma.IBV_WC_RECV && int(completion.ByteLen) != length {
		return fmt.Errorf("receive completion length %d, want %d: %w", completion.ByteLen, length, ErrProtocol)
	}
	return nil
}

func (l *nativeLink) slotRange(slot, length int) error {
	if l == nil || slot < 0 || slot >= nativeSlotDepth || length < 0 || length > nativeSlotBytes {
		return fmt.Errorf("slot=%d length=%d: %w", slot, length, ErrProtocol)
	}
	return nil
}
