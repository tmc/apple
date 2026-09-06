//go:build darwin && arm64

package jaccl

import (
	"context"
	"errors"
	"fmt"
)

// open returns only after every peer link has passed destination exchange, QP
// transitions, and the readiness acknowledgement. Every failure closes the
// control plane first so another rank cannot remain blocked in setup.
func open(ctx context.Context, cfg Config) (*Group, error) {
	var coordinator *coordinator
	var peer *controlPeer
	var err error
	if cfg.Rank == 0 {
		coordinator, err = listenRendezvous(ctx, cfg)
	} else {
		peer, err = dialRendezvous(ctx, cfg)
	}
	if err != nil {
		return nil, err
	}
	devices := make(map[string]*nativeDevice)
	allLinks := make(map[linkKey]*nativeLink)
	fail := func(err error) (*Group, error) {
		var errs []error
		if coordinator != nil {
			errs = append(errs, coordinator.abort(cfg, 0), coordinator.close())
		} else if peer != nil {
			errs = append(errs, peer.close())
		}
		for _, key := range sortedLinkKeys(allLinks) {
			errs = append(errs, allLinks[key].Close())
		}
		for _, name := range sortedDeviceNames(devices) {
			errs = append(errs, devices[name].Close())
		}
		return nil, errors.Join(append([]error{err}, errs...)...)
	}
	plans, err := cfg.linkPlans()
	if err != nil {
		return fail(err)
	}
	local := make([]destinationRecord, 0, len(plans))
	localDestinations := make(map[linkKey]nativeDestination, len(plans))
	for _, plan := range plans {
		device := devices[plan.Device]
		if device == nil {
			device, err = openNativeDevice(Config{Device: plan.Device, Port: cfg.Port})
			if err != nil {
				return fail(fmt.Errorf("open native device %q: %w", plan.Device, err))
			}
			devices[plan.Device] = device
		}
		link, err := newNativeLink(device, plan.Peer)
		if err != nil {
			return fail(fmt.Errorf("create link to rank %d wire %d: %w", plan.Peer, plan.Wire, err))
		}
		key := linkKey{Peer: plan.Peer, Direction: plan.Direction, Wire: plan.Wire}
		allLinks[key] = link
		destination, err := link.localDestination()
		if err != nil {
			return fail(fmt.Errorf("build destination for rank %d wire %d: %w", plan.Peer, plan.Wire, err))
		}
		data, err := encodeNativeDestination(destination)
		if err != nil {
			return fail(fmt.Errorf("encode destination for rank %d wire %d: %w", plan.Peer, plan.Wire, err))
		}
		localDestinations[key] = destination
		local = append(local, destinationRecord{Peer: plan.Peer, Direction: plan.Direction, Wire: plan.Wire, Data: data})
	}
	var remote []destinationRecord
	if coordinator != nil {
		remote, err = coordinator.exchangeDestinations(ctx, cfg, local)
	} else {
		remote, err = peer.exchangeDestinations(ctx, cfg, local)
	}
	if err != nil {
		return fail(fmt.Errorf("exchange destinations: %w", err))
	}
	if err := connectLinks(ctx, allLinks, localDestinations, remote); err != nil {
		return fail(err)
	}
	if coordinator != nil {
		err = coordinator.ready(ctx, cfg)
	} else {
		err = peer.ready(ctx, cfg)
	}
	if err != nil {
		return fail(fmt.Errorf("ready: %w", err))
	}
	links, ringLeft, ringRight, err := primaryLinks(cfg, allLinks)
	if err != nil {
		return fail(err)
	}
	backend := &nativeBackend{
		cfg:          cfg,
		devices:      devices,
		links:        links,
		allLinks:     allLinks,
		destinations: localDestinations,
		ring:         cfg.usesRing(),
		ringLeft:     ringLeft,
		ringRight:    ringRight,
		coordinator:  coordinator,
		peer:         peer,
		closed:       make(chan struct{}),
		p2pEpoch:     make(map[int]uint32),
	}
	if coordinator != nil {
		backend.p2p = newP2PBroker(cfg, coordinator, backend.closed)
	}
	return &Group{rank: cfg.Rank, size: cfg.Size, backend: backend}, nil
}

func connectLinks(ctx context.Context, links map[linkKey]*nativeLink, local map[linkKey]nativeDestination, remote []destinationRecord) error {
	if len(remote) != len(links) {
		return fmt.Errorf("received %d destinations for %d links: %w", len(remote), len(links), ErrProtocol)
	}
	seen := make(map[linkKey]bool, len(remote))
	for _, record := range remote {
		key := linkKey{Peer: record.Peer, Direction: record.Direction, Wire: record.Wire}
		link := links[key]
		if link == nil || seen[key] {
			return fmt.Errorf("destination from rank %d direction %d wire %d: %w", record.Peer, record.Direction, record.Wire, ErrProtocol)
		}
		seen[key] = true
		remoteDestination, err := decodeNativeDestination(record.Data)
		if err != nil {
			return fmt.Errorf("decode destination from rank %d: %w", record.Peer, err)
		}
		if err := link.initialize(); err != nil {
			return fmt.Errorf("initialize link to rank %d wire %d: %w", record.Peer, record.Wire, err)
		}
		if err := link.connect(ctx, local[key], remoteDestination); err != nil {
			return fmt.Errorf("connect link to rank %d wire %d: %w", record.Peer, record.Wire, err)
		}
	}
	if len(seen) != len(links) {
		return fmt.Errorf("missing link destinations: %w", ErrProtocol)
	}
	return nil
}

func primaryLinks(cfg Config, all map[linkKey]*nativeLink) (map[int]*nativeLink, []*nativeLink, []*nativeLink, error) {
	links := make(map[int]*nativeLink)
	if !cfg.usesRing() {
		for key, link := range all {
			if key.Direction != 0 || key.Wire != 0 || links[key.Peer] != nil {
				return nil, nil, nil, fmt.Errorf("mesh link peer=%d direction=%d wire=%d: %w", key.Peer, key.Direction, key.Wire, ErrProtocol)
			}
			links[key.Peer] = link
		}
		return links, nil, nil, nil
	}
	plans, err := cfg.linkPlans()
	if err != nil {
		return nil, nil, nil, err
	}
	wires := 0
	for _, plan := range plans {
		if plan.Direction == -1 && plan.Wire+1 > wires {
			wires = plan.Wire + 1
		}
	}
	if wires == 0 {
		return nil, nil, nil, fmt.Errorf("ring has no left links: %w", ErrProtocol)
	}
	left := make([]*nativeLink, wires)
	right := make([]*nativeLink, wires)
	for _, plan := range plans {
		key := linkKey{Peer: plan.Peer, Direction: plan.Direction, Wire: plan.Wire}
		link := all[key]
		if link == nil || plan.Wire >= wires {
			return nil, nil, nil, fmt.Errorf("ring link peer=%d direction=%d wire=%d: %w", plan.Peer, plan.Direction, plan.Wire, ErrProtocol)
		}
		switch plan.Direction {
		case -1:
			if left[plan.Wire] != nil {
				return nil, nil, nil, fmt.Errorf("duplicate ring left wire %d: %w", plan.Wire, ErrProtocol)
			}
			left[plan.Wire] = link
			links[plan.Peer] = link
		case 1:
			if right[plan.Wire] != nil {
				return nil, nil, nil, fmt.Errorf("duplicate ring right wire %d: %w", plan.Wire, ErrProtocol)
			}
			right[plan.Wire] = link
			// For a two-rank ring, this intentional overwrite selects the
			// right link for the otherwise ambiguous public Send and Recv API.
			links[plan.Peer] = link
		default:
			return nil, nil, nil, fmt.Errorf("ring link direction %d: %w", plan.Direction, ErrProtocol)
		}
	}
	for wire := range left {
		if left[wire] == nil || right[wire] == nil {
			return nil, nil, nil, fmt.Errorf("ring wire %d is incomplete: %w", wire, ErrProtocol)
		}
	}
	return links, left, right, nil
}
