package jaccl

import (
	"context"
	"fmt"
)

func (c *coordinator) ready(ctx context.Context, cfg Config) error {
	if c == nil || cfg.Rank != 0 || c.incarnation == 0 {
		return fmt.Errorf("coordinator ready: %w", ErrProtocol)
	}
	for rank, peer := range c.peers {
		frame, err := readControlFrameContext(ctx, peer.conn)
		if err != nil {
			return fmt.Errorf("read rank %d ready: %w", rank, err)
		}
		if frame.Kind != frameReady || frame.GroupID != cfg.GroupID || frame.Incarnation != c.incarnation || frame.Source != rank || frame.Destination != 0 {
			return fmt.Errorf("rank %d ready frame: %w", rank, ErrProtocol)
		}
	}
	for rank, peer := range c.peers {
		if err := peer.send(controlFrame{
			Magic:       protocolMagic,
			Version:     protocolVersion,
			Kind:        frameReady,
			GroupID:     cfg.GroupID,
			Incarnation: c.incarnation,
			Source:      0,
			Destination: rank,
		}); err != nil {
			return fmt.Errorf("acknowledge rank %d ready: %w", rank, err)
		}
	}
	return nil
}

func (p *controlPeer) ready(ctx context.Context, cfg Config) error {
	if p == nil || cfg.Rank == 0 || p.rank != 0 || p.incarnation == 0 {
		return fmt.Errorf("peer ready: %w", ErrProtocol)
	}
	if err := p.send(controlFrame{
		Magic:       protocolMagic,
		Version:     protocolVersion,
		Kind:        frameReady,
		GroupID:     cfg.GroupID,
		Incarnation: p.incarnation,
		Source:      cfg.Rank,
		Destination: 0,
	}); err != nil {
		return fmt.Errorf("send ready: %w", err)
	}
	frame, err := readControlFrameContext(ctx, p.conn)
	if err != nil {
		return fmt.Errorf("read ready acknowledgement: %w", err)
	}
	if frame.Kind != frameReady || frame.GroupID != cfg.GroupID || frame.Incarnation != p.incarnation || frame.Source != 0 || frame.Destination != cfg.Rank {
		return fmt.Errorf("ready acknowledgement: %w", ErrProtocol)
	}
	return nil
}
