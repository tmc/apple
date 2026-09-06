package jaccl

import (
	"context"
	"encoding/json"
	"fmt"
)

// destinationRecord carries one opaque, peer-specific link record. The native
// backend serializes its QP destination in Data; the rendezvous layer checks
// only identity and topology.
type destinationRecord struct {
	Peer      int    `json:"peer"`
	Direction int    `json:"direction"`
	Wire      int    `json:"wire"`
	Data      []byte `json:"data"`
}

func (c *coordinator) exchangeDestinations(ctx context.Context, cfg Config, local []destinationRecord) ([]destinationRecord, error) {
	if c == nil || cfg.Rank != 0 || c.incarnation == 0 {
		return nil, fmt.Errorf("coordinator destination exchange: %w", ErrProtocol)
	}
	if err := validateDestinationRecords(local, cfg, 0); err != nil {
		return nil, err
	}
	all := make([][]destinationRecord, cfg.Size)
	all[0] = local
	for rank, peer := range c.peers {
		frame, err := readControlFrameContext(ctx, peer.conn)
		if err != nil {
			return nil, fmt.Errorf("read rank %d destinations: %w", rank, err)
		}
		if err := validateDestinationFrame(frame, cfg.GroupID, c.incarnation, rank, 0); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(frame.Payload, &all[rank]); err != nil {
			return nil, fmt.Errorf("decode rank %d destinations: %w", rank, err)
		}
		if err := validateDestinationRecords(all[rank], cfg, rank); err != nil {
			return nil, fmt.Errorf("rank %d destinations: %w", rank, err)
		}
	}
	for rank, peer := range c.peers {
		inbound := inboundDestinations(all, rank)
		payload, err := json.Marshal(inbound)
		if err != nil {
			return nil, fmt.Errorf("encode rank %d destinations: %w", rank, err)
		}
		if err := peer.send(controlFrame{
			Magic:       protocolMagic,
			Version:     protocolVersion,
			Kind:        frameDestination,
			GroupID:     cfg.GroupID,
			Incarnation: c.incarnation,
			Source:      0,
			Destination: rank,
			Payload:     payload,
		}); err != nil {
			return nil, fmt.Errorf("send rank %d destinations: %w", rank, err)
		}
	}
	return inboundDestinations(all, 0), nil
}

func (p *controlPeer) exchangeDestinations(ctx context.Context, cfg Config, local []destinationRecord) ([]destinationRecord, error) {
	if p == nil || cfg.Rank == 0 || p.rank != 0 || p.incarnation == 0 {
		return nil, fmt.Errorf("peer destination exchange: %w", ErrProtocol)
	}
	if err := validateDestinationRecords(local, cfg, cfg.Rank); err != nil {
		return nil, err
	}
	payload, err := json.Marshal(local)
	if err != nil {
		return nil, fmt.Errorf("encode destinations: %w", err)
	}
	if err := p.send(controlFrame{
		Magic:       protocolMagic,
		Version:     protocolVersion,
		Kind:        frameDestination,
		GroupID:     cfg.GroupID,
		Incarnation: p.incarnation,
		Source:      cfg.Rank,
		Destination: 0,
		Payload:     payload,
	}); err != nil {
		return nil, fmt.Errorf("send destinations: %w", err)
	}
	frame, err := readControlFrameContext(ctx, p.conn)
	if err != nil {
		return nil, fmt.Errorf("read destinations: %w", err)
	}
	if err := validateDestinationFrame(frame, cfg.GroupID, p.incarnation, 0, cfg.Rank); err != nil {
		return nil, err
	}
	var inbound []destinationRecord
	if err := json.Unmarshal(frame.Payload, &inbound); err != nil {
		return nil, fmt.Errorf("decode inbound destinations: %w", err)
	}
	if err := validateDestinationRecords(inbound, cfg, cfg.Rank); err != nil {
		return nil, fmt.Errorf("validate inbound destinations: %w", err)
	}
	return inbound, nil
}

func validateDestinationFrame(frame controlFrame, groupID string, incarnation uint64, source, destination int) error {
	if frame.Kind != frameDestination || frame.GroupID != groupID || frame.Incarnation != incarnation || frame.Source != source || frame.Destination != destination {
		return fmt.Errorf("destination frame source=%d destination=%d: %w", frame.Source, frame.Destination, ErrProtocol)
	}
	return nil
}

func validateDestinationRecords(records []destinationRecord, cfg Config, rank int) error {
	plans, err := cfg.linkPlansForRank(rank)
	if err != nil {
		return err
	}
	if len(records) != len(plans) {
		return fmt.Errorf("destination count %d, want %d: %w", len(records), len(plans), ErrProtocol)
	}
	want := make(map[linkKey]bool, len(plans))
	for _, plan := range plans {
		want[linkKey{Peer: plan.Peer, Direction: plan.Direction, Wire: plan.Wire}] = true
	}
	for _, record := range records {
		key := linkKey{Peer: record.Peer, Direction: record.Direction, Wire: record.Wire}
		if record.Peer < 0 || record.Peer >= cfg.Size || record.Wire < 0 || !want[key] || len(record.Data) == 0 || len(record.Data) > maxControlFrameBytes {
			return fmt.Errorf("destination peer=%d direction=%d wire=%d: %w", record.Peer, record.Direction, record.Wire, ErrProtocol)
		}
		delete(want, key)
	}
	if len(want) != 0 {
		return fmt.Errorf("destination peers missing: %w", ErrProtocol)
	}
	return nil
}

func inboundDestinations(all [][]destinationRecord, destination int) []destinationRecord {
	var inbound []destinationRecord
	for source, records := range all {
		for _, record := range records {
			if record.Peer == destination {
				inbound = append(inbound, destinationRecord{Peer: source, Direction: -record.Direction, Wire: record.Wire, Data: record.Data})
			}
		}
	}
	return inbound
}
