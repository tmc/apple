package jaccl

import (
	"fmt"
	"strings"
)

// Config describes one rank in a group.
type Config struct {
	Rank        int
	Size        int
	GroupID     string
	Coordinator string
	Device      string
	// Devices is the upstream JACCL directed connectivity matrix. Each cell
	// names zero or more RDMA devices for one source-to-destination link.
	Devices [][][]string
	// PreferRing asks ConfigFromEnv to choose a valid ring before a mesh.
	PreferRing bool
	Port       uint8
	Topology   Topology
}

// Topology reports the directly connected peers for a rank.
//
// Peers must return no duplicate, self, or out-of-range ranks. Open validates
// both the local report and reciprocal reports exchanged with peers.
type Topology interface {
	Peers(rank int) []int
}

func (c Config) validate() ([]int, error) {
	peers, err := c.validateTopology()
	if err != nil {
		return nil, err
	}
	if _, err := c.linkPlans(); err != nil {
		return nil, err
	}
	return peers, nil
}

func (c Config) validateTopology() ([]int, error) {
	if c.Size < 2 {
		return nil, fmt.Errorf("size %d: %w", c.Size, ErrInvalidConfig)
	}
	if c.Rank < 0 || c.Rank >= c.Size {
		return nil, fmt.Errorf("rank %d for size %d: %w", c.Rank, c.Size, ErrInvalidConfig)
	}
	if strings.TrimSpace(c.GroupID) == "" {
		return nil, fmt.Errorf("group id is empty: %w", ErrInvalidConfig)
	}
	if strings.TrimSpace(c.Coordinator) == "" {
		return nil, fmt.Errorf("coordinator is empty: %w", ErrInvalidConfig)
	}
	if c.Port == 0 {
		return nil, fmt.Errorf("port is zero: %w", ErrInvalidConfig)
	}
	if c.Topology == nil {
		return nil, fmt.Errorf("topology is nil: %w", ErrInvalidConfig)
	}
	peers := append([]int(nil), c.Topology.Peers(c.Rank)...)
	seen := make(map[int]bool, len(peers))
	for _, peer := range peers {
		if peer < 0 || peer >= c.Size {
			return nil, fmt.Errorf("peer %d out of range for size %d: %w", peer, c.Size, ErrInvalidConfig)
		}
		if peer == c.Rank {
			return nil, fmt.Errorf("peer %d is self: %w", peer, ErrInvalidConfig)
		}
		if seen[peer] {
			return nil, fmt.Errorf("peer %d is duplicated: %w", peer, ErrInvalidConfig)
		}
		seen[peer] = true
	}
	return peers, nil
}

// Mesh returns a topology in which every rank connects directly to every other rank.
func Mesh(size int) Topology {
	return meshTopology{size: size}
}

// Ring returns a topology in which each rank connects to its two neighbors.
// A two-rank ring has one peer per rank.
func Ring(size int) Topology {
	return ringTopology{size: size}
}

type meshTopology struct {
	size int
}

func (t meshTopology) Peers(rank int) []int {
	if t.size <= 0 {
		return nil
	}
	peers := make([]int, 0, t.size-1)
	for peer := 0; peer < t.size; peer++ {
		if peer != rank {
			peers = append(peers, peer)
		}
	}
	return peers
}

type ringTopology struct {
	size int
}

func (t ringTopology) Peers(rank int) []int {
	if t.size < 2 || rank < 0 || rank >= t.size {
		return nil
	}
	if t.size == 2 {
		return []int{1 - rank}
	}
	left := (rank + t.size - 1) % t.size
	right := (rank + 1) % t.size
	return []int{left, right}
}
