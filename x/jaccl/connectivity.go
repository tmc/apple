package jaccl

import "fmt"

const (
	jacclMeshMaxPeers = 8
	jacclRingMaxWires = 4
)

// linkPlan identifies one local UC connection selected from Config.Devices.
// Direction is -1 for the left side of a ring, +1 for the right side, and 0
// for a mesh connection. Wire is zero based within one directed link.
type linkPlan struct {
	Peer      int
	Direction int
	Wire      int
	Device    string
}

type linkKey struct {
	Peer      int
	Direction int
	Wire      int
}

func (c Config) linkPlans() ([]linkPlan, error) {
	return c.linkPlansForRank(c.Rank)
}

func (c Config) linkPlansForRank(rank int) ([]linkPlan, error) {
	if rank < 0 || rank >= c.Size {
		return nil, fmt.Errorf("rank %d for size %d: %w", rank, c.Size, ErrInvalidConfig)
	}
	c.Rank = rank
	peers, err := c.validateTopology()
	if err != nil {
		return nil, err
	}
	if len(c.Devices) == 0 {
		if c.usesRing() {
			if !topologyHasPeers(c.Topology, c.Size, Ring(c.Size)) {
				return nil, fmt.Errorf("ring devices and topology disagree: %w", ErrInvalidConfig)
			}
			return c.legacyRingLinkPlans()
		}
		if !topologyHasPeers(c.Topology, c.Size, Mesh(c.Size)) {
			return nil, fmt.Errorf("topology is neither a JACCL mesh nor ring: %w", ErrInvalidConfig)
		}
		if c.Size > jacclMeshMaxPeers {
			return nil, fmt.Errorf("mesh size %d exceeds JACCL limit %d: %w", c.Size, jacclMeshMaxPeers, ErrInvalidConfig)
		}
		plans := make([]linkPlan, 0, len(peers))
		for _, peer := range peers {
			plans = append(plans, linkPlan{Peer: peer, Device: c.Device})
		}
		return plans, nil
	}
	if err := validateDeviceMatrix(c.Devices, c.Size); err != nil {
		return nil, err
	}
	// The three-rank mesh and ring have the same peer sets. PreferRing carries
	// the upstream Config preference that disambiguates that case. For other
	// sizes a ring-only topology also selects the ring path.
	if c.usesRing() {
		if !topologyHasPeers(c.Topology, c.Size, Ring(c.Size)) {
			return nil, fmt.Errorf("ring devices and topology disagree: %w", ErrInvalidConfig)
		}
		return c.ringLinkPlans()
	}
	if validDeviceMesh(c.Devices) {
		if !topologyHasPeers(c.Topology, c.Size, Mesh(c.Size)) {
			return nil, fmt.Errorf("mesh devices and topology disagree: %w", ErrInvalidConfig)
		}
		return c.meshLinkPlans()
	}
	if validDeviceRing(c.Devices) {
		return c.ringLinkPlans()
	}
	return nil, fmt.Errorf("device connectivity needs a JACCL mesh or ring topology: %w", ErrInvalidConfig)
}

func topologyHasPeers(got Topology, size int, want Topology) bool {
	if got == nil || want == nil {
		return false
	}
	for rank := 0; rank < size; rank++ {
		if !sameRanks(got.Peers(rank), want.Peers(rank)) {
			return false
		}
	}
	return true
}

func sameRanks(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[int]bool, len(a))
	for _, rank := range a {
		seen[rank] = true
	}
	for _, rank := range b {
		if !seen[rank] {
			return false
		}
	}
	return true
}

func (c Config) usesRing() bool {
	if len(c.Devices) == 0 {
		_, ok := c.Topology.(ringTopology)
		return ok
	}
	return validDeviceRing(c.Devices) && (c.PreferRing || !validDeviceMesh(c.Devices))
}

func (c Config) legacyRingLinkPlans() ([]linkPlan, error) {
	if c.Size < 2 {
		return nil, fmt.Errorf("ring size %d: %w", c.Size, ErrInvalidConfig)
	}
	left := (c.Rank + c.Size - 1) % c.Size
	right := (c.Rank + 1) % c.Size
	return []linkPlan{
		{Peer: left, Direction: -1, Device: c.Device},
		{Peer: right, Direction: 1, Device: c.Device},
	}, nil
}

func (c Config) meshLinkPlans() ([]linkPlan, error) {
	if c.Size > jacclMeshMaxPeers {
		return nil, fmt.Errorf("mesh size %d exceeds JACCL limit %d: %w", c.Size, jacclMeshMaxPeers, ErrInvalidConfig)
	}
	plans := make([]linkPlan, 0, c.Size-1)
	for peer := 0; peer < c.Size; peer++ {
		if peer == c.Rank {
			continue
		}
		names := c.Devices[c.Rank][peer]
		if len(names) == 0 {
			return nil, fmt.Errorf("mesh device %d->%d is empty: %w", c.Rank, peer, ErrInvalidConfig)
		}
		plans = append(plans, linkPlan{Peer: peer, Device: names[0]})
	}
	return plans, nil
}

func (c Config) ringLinkPlans() ([]linkPlan, error) {
	if c.Size < 2 {
		return nil, fmt.Errorf("ring size %d: %w", c.Size, ErrInvalidConfig)
	}
	left := (c.Rank + c.Size - 1) % c.Size
	right := (c.Rank + 1) % c.Size
	leftNames := c.Devices[c.Rank][left]
	rightNames := c.Devices[c.Rank][right]
	if len(leftNames) == 0 || len(leftNames) != len(rightNames) {
		return nil, fmt.Errorf("ring devices for rank %d have %d left and %d right wires: %w", c.Rank, len(leftNames), len(rightNames), ErrInvalidConfig)
	}
	if len(leftNames) > jacclRingMaxWires {
		return nil, fmt.Errorf("ring has %d wires, exceeds JACCL limit %d: %w", len(leftNames), jacclRingMaxWires, ErrInvalidConfig)
	}
	plans := make([]linkPlan, 0, 2*len(leftNames))
	for wire, name := range leftNames {
		plans = append(plans, linkPlan{Peer: left, Direction: -1, Wire: wire, Device: name})
	}
	for wire, name := range rightNames {
		plans = append(plans, linkPlan{Peer: right, Direction: 1, Wire: wire, Device: name})
	}
	return plans, nil
}

func validateDeviceMatrix(devices [][][]string, size int) error {
	if len(devices) != size {
		return fmt.Errorf("device matrix has %d rows, want %d: %w", len(devices), size, ErrInvalidConfig)
	}
	for rank, row := range devices {
		if len(row) != size {
			return fmt.Errorf("device matrix row %d has %d columns, want %d: %w", rank, len(row), size, ErrInvalidConfig)
		}
		for peer, names := range row {
			for wire, name := range names {
				if name == "" {
					return fmt.Errorf("device matrix %d->%d wire %d is empty: %w", rank, peer, wire, ErrInvalidConfig)
				}
			}
		}
	}
	return nil
}
