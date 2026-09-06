package jaccl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// operation describes one ordered collective. It intentionally contains only
// shape, not payload bytes: the TCP control plane establishes agreement but is
// never a data path.
type operation struct {
	Name        string   `json:"name"`
	Length      int      `json:"length"`
	DType       DType    `json:"dtype,omitempty"`
	Reduce      ReduceOp `json:"reduce,omitempty"`
	Source      int      `json:"source,omitempty"`
	Destination int      `json:"destination,omitempty"`
}

func (o operation) validate() error {
	switch o.Name {
	case "barrier":
		if o.Length != 0 || o.DType != 0 || o.Reduce != 0 {
			return fmt.Errorf("barrier shape: %w", ErrProtocol)
		}
	case "all gather":
		if o.Length < 0 || o.DType != 0 || o.Reduce != 0 {
			return fmt.Errorf("all gather shape: %w", ErrProtocol)
		}
	case "all reduce":
		width, err := o.DType.Size()
		if o.Length < 0 || err != nil || o.Length%width != 0 || !o.Reduce.valid() {
			return fmt.Errorf("all reduce shape: %w", ErrProtocol)
		}
	case "transfer":
		if o.Length < 0 || o.DType != 0 || o.Reduce != 0 || o.Source < 0 || o.Destination < 0 || o.Source == o.Destination {
			return fmt.Errorf("transfer shape: %w", ErrProtocol)
		}
	default:
		return fmt.Errorf("operation %q: %w", o.Name, ErrProtocol)
	}
	return nil
}

func (o operation) validateForConfig(cfg Config) error {
	if err := o.validate(); err != nil {
		return err
	}
	if o.Name == "transfer" && (o.Source >= cfg.Size || o.Destination >= cfg.Size) {
		return fmt.Errorf("transfer ranks source=%d destination=%d size=%d: %w", o.Source, o.Destination, cfg.Size, ErrProtocol)
	}
	return nil
}

func encodeOperation(o operation) ([]byte, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	payload, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("encode operation: %w", err)
	}
	if len(payload) > maxControlFrameBytes {
		return nil, fmt.Errorf("operation length %d: %w", len(payload), ErrProtocol)
	}
	return payload, nil
}

func decodeOperation(payload []byte) (operation, error) {
	if len(payload) == 0 || len(payload) > maxControlFrameBytes {
		return operation{}, fmt.Errorf("operation length %d: %w", len(payload), ErrProtocol)
	}
	var o operation
	if err := json.Unmarshal(payload, &o); err != nil {
		return operation{}, fmt.Errorf("decode operation: %w", err)
	}
	if err := o.validate(); err != nil {
		return operation{}, err
	}
	return o, nil
}

func (c *coordinator) agree(ctx context.Context, cfg Config, epoch uint32, local operation) error {
	if c == nil || cfg.Rank != 0 || c.incarnation == 0 {
		return fmt.Errorf("coordinator operation agreement: %w", ErrProtocol)
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := local.validateForConfig(cfg); err != nil {
		return err
	}
	payload, err := encodeOperation(local)
	if err != nil {
		return err
	}
	for rank, peer := range c.peers {
		frame, err := readControlFrameContext(ctx, peer.conn)
		if err != nil {
			return fmt.Errorf("read rank %d operation: %w", rank, err)
		}
		if err := validateOperationFrame(frame, cfg.GroupID, c.incarnation, rank, 0, epoch); err != nil {
			return err
		}
		remote, err := decodeOperation(frame.Payload)
		if err == nil {
			err = remote.validateForConfig(cfg)
		}
		if err != nil || remote != local {
			_ = c.abort(cfg, epoch)
			if err != nil {
				return fmt.Errorf("rank %d operation: %w", rank, err)
			}
			return fmt.Errorf("rank %d operation does not match: %w", rank, ErrProtocol)
		}
	}
	for rank, peer := range c.peers {
		if err := peer.send(controlFrame{
			Magic: protocolMagic, Version: protocolVersion, Kind: frameOperation,
			GroupID: cfg.GroupID, Incarnation: c.incarnation,
			Source: 0, Destination: rank, Epoch: epoch, Payload: payload,
		}); err != nil {
			return fmt.Errorf("acknowledge rank %d operation: %w", rank, err)
		}
	}
	return nil
}

func (c *coordinator) abort(cfg Config, epoch uint32) error {
	if c == nil {
		return nil
	}
	var errs []error
	for rank, peer := range c.peers {
		err := peer.send(controlFrame{
			Magic: protocolMagic, Version: protocolVersion, Kind: frameAbort,
			GroupID: cfg.GroupID, Incarnation: c.incarnation,
			Source: 0, Destination: rank, Epoch: epoch,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("abort rank %d: %w", rank, err))
		}
	}
	return errors.Join(errs...)
}

func (p *controlPeer) agree(ctx context.Context, cfg Config, epoch uint32, local operation) error {
	if p == nil || cfg.Rank == 0 || p.rank != 0 || p.incarnation == 0 {
		return fmt.Errorf("peer operation agreement: %w", ErrProtocol)
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := local.validateForConfig(cfg); err != nil {
		return err
	}
	payload, err := encodeOperation(local)
	if err != nil {
		return err
	}
	if err := p.send(controlFrame{
		Magic: protocolMagic, Version: protocolVersion, Kind: frameOperation,
		GroupID: cfg.GroupID, Incarnation: p.incarnation,
		Source: cfg.Rank, Destination: 0, Epoch: epoch, Payload: payload,
	}); err != nil {
		return fmt.Errorf("send operation: %w", err)
	}
	frame, err := readControlFrameContext(ctx, p.conn)
	if err != nil {
		return fmt.Errorf("read operation acknowledgement: %w", err)
	}
	if frame.Kind == frameAbort && frame.GroupID == cfg.GroupID && frame.Incarnation == p.incarnation && frame.Source == 0 && frame.Destination == cfg.Rank && frame.Epoch == epoch {
		return fmt.Errorf("operation rejected by coordinator: %w", ErrProtocol)
	}
	if err := validateOperationFrame(frame, cfg.GroupID, p.incarnation, 0, cfg.Rank, epoch); err != nil {
		return err
	}
	remote, err := decodeOperation(frame.Payload)
	if err != nil {
		return err
	}
	if err := remote.validateForConfig(cfg); err != nil {
		return err
	}
	if remote != local {
		return fmt.Errorf("operation acknowledgement does not match: %w", ErrProtocol)
	}
	return nil
}

func validateOperationFrame(frame controlFrame, groupID string, incarnation uint64, source, destination int, epoch uint32) error {
	if frame.Kind != frameOperation || frame.GroupID != groupID || frame.Incarnation != incarnation || frame.Source != source || frame.Destination != destination || frame.Epoch != epoch {
		return fmt.Errorf("operation frame source=%d destination=%d epoch=%d: %w", frame.Source, frame.Destination, frame.Epoch, ErrProtocol)
	}
	return nil
}
