package jaccl

import (
	"context"
	"encoding/json"
	"fmt"
)

// creditRecord grants one remote send against a posted receive slot. Peer is
// the receiver for an outbound record and the sender for an inbound record.
type creditRecord struct {
	Peer       int    `json:"peer"`
	Slot       int    `json:"slot"`
	Generation uint16 `json:"generation"`
}

func (c *coordinator) exchangeCredits(ctx context.Context, cfg Config, epoch uint32, local []creditRecord) ([]creditRecord, error) {
	if c == nil || cfg.Rank != 0 || c.incarnation == 0 {
		return nil, fmt.Errorf("coordinator credit exchange: %w", ErrProtocol)
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if err := validateCreditRecords(local, c.topology[0], cfg.Size); err != nil {
		return nil, err
	}
	all := make([][]creditRecord, cfg.Size)
	all[0] = local
	for rank, peer := range c.peers {
		frame, err := readControlFrameContext(ctx, peer.conn)
		if err != nil {
			return nil, fmt.Errorf("read rank %d credits: %w", rank, err)
		}
		if err := validateCreditFrame(frame, cfg.GroupID, c.incarnation, rank, 0, epoch); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(frame.Payload, &all[rank]); err != nil {
			return nil, fmt.Errorf("decode rank %d credits: %w", rank, err)
		}
		if err := validateCreditRecords(all[rank], c.topology[rank], cfg.Size); err != nil {
			return nil, fmt.Errorf("rank %d credits: %w", rank, err)
		}
	}
	for rank, peer := range c.peers {
		inbound := inboundCredits(all, rank)
		payload, err := json.Marshal(inbound)
		if err != nil {
			return nil, fmt.Errorf("encode rank %d credits: %w", rank, err)
		}
		if err := peer.send(controlFrame{
			Magic: protocolMagic, Version: protocolVersion, Kind: frameCredit,
			GroupID: cfg.GroupID, Incarnation: c.incarnation,
			Source: 0, Destination: rank, Epoch: epoch, Payload: payload,
		}); err != nil {
			return nil, fmt.Errorf("send rank %d credits: %w", rank, err)
		}
	}
	return inboundCredits(all, 0), nil
}

func (p *controlPeer) exchangeCredits(ctx context.Context, cfg Config, epoch uint32, local []creditRecord) ([]creditRecord, error) {
	if p == nil || cfg.Rank == 0 || p.rank != 0 || p.incarnation == 0 {
		return nil, fmt.Errorf("peer credit exchange: %w", ErrProtocol)
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	peers, _ := cfg.validate()
	if err := validateCreditRecords(local, peers, cfg.Size); err != nil {
		return nil, err
	}
	payload, err := json.Marshal(local)
	if err != nil {
		return nil, fmt.Errorf("encode credits: %w", err)
	}
	if err := p.send(controlFrame{
		Magic: protocolMagic, Version: protocolVersion, Kind: frameCredit,
		GroupID: cfg.GroupID, Incarnation: p.incarnation,
		Source: cfg.Rank, Destination: 0, Epoch: epoch, Payload: payload,
	}); err != nil {
		return nil, fmt.Errorf("send credits: %w", err)
	}
	frame, err := readControlFrameContext(ctx, p.conn)
	if err != nil {
		return nil, fmt.Errorf("read credits: %w", err)
	}
	if err := validateCreditFrame(frame, cfg.GroupID, p.incarnation, 0, cfg.Rank, epoch); err != nil {
		return nil, err
	}
	var inbound []creditRecord
	if err := json.Unmarshal(frame.Payload, &inbound); err != nil {
		return nil, fmt.Errorf("decode inbound credits: %w", err)
	}
	if err := validateCreditRecords(inbound, peers, cfg.Size); err != nil {
		return nil, fmt.Errorf("validate inbound credits: %w", err)
	}
	return inbound, nil
}

func validateCreditFrame(frame controlFrame, groupID string, incarnation uint64, source, destination int, epoch uint32) error {
	if frame.Kind != frameCredit || frame.GroupID != groupID || frame.Incarnation != incarnation || frame.Source != source || frame.Destination != destination || frame.Epoch != epoch {
		return fmt.Errorf("credit frame source=%d destination=%d epoch=%d: %w", frame.Source, frame.Destination, frame.Epoch, ErrProtocol)
	}
	return nil
}

func validateCreditRecords(records []creditRecord, peers []int, size int) error {
	want := make(map[int]bool, len(peers))
	for _, peer := range peers {
		want[peer] = true
	}
	seen := make(map[creditRecord]bool, len(records))
	for _, record := range records {
		if record.Peer < 0 || record.Peer >= size || !want[record.Peer] || record.Slot < 0 || record.Slot >= nativeSlotDepth || record.Generation > workIDGenerationMax || seen[record] {
			return fmt.Errorf("credit peer=%d slot=%d generation=%d: %w", record.Peer, record.Slot, record.Generation, ErrProtocol)
		}
		seen[record] = true
	}
	return nil
}

func inboundCredits(all [][]creditRecord, destination int) []creditRecord {
	var inbound []creditRecord
	for source, records := range all {
		for _, record := range records {
			if record.Peer == destination {
				inbound = append(inbound, creditRecord{Peer: source, Slot: record.Slot, Generation: record.Generation})
			}
		}
	}
	return inbound
}
