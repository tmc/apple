package espresso

import (
	"errors"
	"fmt"
)

// Platform selects the execution backend.
type Platform int

const (
	PlatformAuto Platform = iota // Espresso selects per-layer
	PlatformCPU
	PlatformGPU
	PlatformANE
)

func (p Platform) String() string {
	switch p {
	case PlatformAuto:
		return "auto"
	case PlatformCPU:
		return "cpu"
	case PlatformGPU:
		return "gpu"
	case PlatformANE:
		return "ane"
	default:
		return fmt.Sprintf("platform(%d)", int(p))
	}
}

// Engine identifies which hardware executed a layer.
type Engine int

const (
	EngineCPU Engine = iota
	EngineGPU
	EngineANE
)

func (e Engine) String() string {
	switch e {
	case EngineCPU:
		return "cpu"
	case EngineGPU:
		return "gpu"
	case EngineANE:
		return "ane"
	default:
		return fmt.Sprintf("engine(%d)", int(e))
	}
}

// Shape describes tensor dimensions.
type Shape struct {
	Batch    int
	Channels int
	Height   int
	Width    int
	Sequence int // for RNN/transformer models
}

// Elements returns the total number of elements.
func (s Shape) Elements() int {
	n := s.Channels * s.Height * s.Width
	if s.Batch > 0 {
		n *= s.Batch
	}
	if s.Sequence > 0 {
		n *= s.Sequence
	}
	return n
}

// Pass is a network optimization pass that transforms an Espresso network.
type Pass interface {
	// Name returns the pass identifier.
	Name() string
	// Run applies the optimization to the network. Returns true if modified.
	Run(n *Network) bool
}

var (
	ErrFramework    = errors.New("espresso: framework not available")
	ErrLoad         = errors.New("espresso: network load failed")
	ErrNilNetwork   = errors.New("espresso: nil network returned")
	ErrEval         = errors.New("espresso: evaluation failed")
	ErrBindInput    = errors.New("espresso: input binding failed")
	ErrBindOutput   = errors.New("espresso: output binding failed")
	ErrNoSuchOutput = errors.New("espresso: output not found")
	ErrClosed       = errors.New("espresso: context closed")
)

// Error wraps an error from the Espresso subsystem.
type Error struct {
	Op  string
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("espresso %s: %v", e.Op, e.Err)
}

func (e *Error) Unwrap() error { return e.Err }

