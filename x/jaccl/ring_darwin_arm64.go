//go:build darwin && arm64

package jaccl

import (
	"context"
	"fmt"
)

// ringGather circulates one rank-sized block per step. It deliberately uses
// one UC transfer per direction and stage: parity first, then pipelining.
func (b *nativeBackend) ringGather(ctx context.Context, epoch uint32, dst, src []byte) error {
	if err := b.requireRing(); err != nil {
		return err
	}
	left := (b.cfg.Rank + b.cfg.Size - 1) % b.cfg.Size
	right := (b.cfg.Rank + 1) % b.cfg.Size
	return ringGatherBlocks(b.cfg.Rank, b.cfg.Size, dst, src, len(b.ringLeft), func(step, wire int, forward, received []byte) error {
		if epoch > ^uint32(0)>>8 || step > 255 {
			return fmt.Errorf("ring stage epoch %d step %d: %w", epoch, step, ErrProtocol)
		}
		return b.ringStage(ctx, epoch<<8|uint32(step), right, left, b.ringRight[wire], b.ringLeft[wire], forward, received)
	})
}

func (b *nativeBackend) ringStage(ctx context.Context, epoch uint32, sendPeer, recvPeer int, sendLink, recvLink *nativeLink, src, dst []byte) error {
	type sessionResult struct {
		session *p2pSession
		err     error
	}
	sent := make(chan sessionResult, 1)
	received := make(chan sessionResult, 1)
	go func() {
		s, err := b.openP2P(ctx, operation{Name: "transfer", Length: len(src), Source: b.cfg.Rank, Destination: sendPeer}, epoch)
		sent <- sessionResult{session: s, err: err}
	}()
	go func() {
		s, err := b.openP2P(ctx, operation{Name: "transfer", Length: len(dst), Source: recvPeer, Destination: b.cfg.Rank}, epoch)
		received <- sessionResult{session: s, err: err}
	}()
	sendSession := <-sent
	recvSession := <-received
	if sendSession.err != nil || recvSession.err != nil {
		if sendSession.session != nil {
			_ = sendSession.session.close()
		}
		if recvSession.session != nil {
			_ = recvSession.session.close()
		}
		return fmt.Errorf("open ring sessions: %w", joinRingErrors(sendSession.err, recvSession.err))
	}
	defer sendSession.session.close()
	defer recvSession.session.close()
	errs := make(chan error, 2)
	go func() {
		errs <- b.transferLink(ctx, epoch, sendPeer, sendLink, src, nil, true, sendSession.session.exchangeCredits)
	}()
	go func() {
		errs <- b.transferLink(ctx, epoch, recvPeer, recvLink, nil, dst, false, recvSession.session.exchangeCredits)
	}()
	return joinRingErrors(<-errs, <-errs)
}

func (b *nativeBackend) requireRing() error {
	if b.cfg.Size < 2 {
		return nil
	}
	if !b.ring || len(b.ringLeft) == 0 || len(b.ringLeft) != len(b.ringRight) {
		return fmt.Errorf("ring has %d left and %d right wires: %w", len(b.ringLeft), len(b.ringRight), ErrUnsupported)
	}
	for wire := range b.ringLeft {
		if b.ringLeft[wire] == nil || b.ringRight[wire] == nil {
			return fmt.Errorf("ring wire %d is missing: %w", wire, ErrUnsupported)
		}
	}
	return nil
}

func joinRingErrors(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
