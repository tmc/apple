package ane

import (
	"errors"
	"fmt"
)

var (
	ErrNoANE                    = errors.New("ane: no ANE hardware available")
	ErrCompileBudgetExhausted   = errors.New("ane: compile budget exhausted")
	ErrMapFailed                = errors.New("ane: IOSurface mapping failed")
	ErrUnsupportedSelector      = errors.New("ane: unsupported selector")
	ErrVirtualClientUnavailable = errors.New("ane: virtual client unavailable")
	ErrModelLoad                = errors.New("ane: model load failed")
	ErrEval                     = errors.New("ane: evaluation failed")
	ErrUnsupportedLayout        = errors.New("ane: unsupported tensor layout")
)

// ANEError wraps an error from the ANE subsystem with context.
type ANEError struct {
	Op    string // operation that failed
	Class string // error domain or class
	Code  int    // error code
	Err   error  // underlying error
}

func (e *ANEError) Error() string {
	if e.Class != "" {
		return fmt.Sprintf("ane %s: %s (code %d): %v", e.Op, e.Class, e.Code, e.Err)
	}
	return fmt.Sprintf("ane %s: %v", e.Op, e.Err)
}

func (e *ANEError) Unwrap() error { return e.Err }

func (e *ANEError) Is(target error) bool {
	t, ok := target.(*ANEError)
	if !ok {
		return false
	}
	return e.Op == t.Op && e.Class == t.Class && e.Code == t.Code
}

func wrapErr(op string, err error) error {
	if err == nil {
		return nil
	}
	return &ANEError{Op: op, Err: err}
}

// DeviceInfo describes the ANE hardware present on the system.
type DeviceInfo struct {
	HasANE       bool
	NumCores     uint32
	Architecture string
	Product      string
	BuildVersion string
	IsVM         bool
}

// ModelType selects the compilation path.
type ModelType int

const (
	ModelTypeMIL     ModelType = iota // In-memory MIL text + weights
	ModelTypePackage                  // On-disk .mlmodelc package
)

// CompileOptions configures model compilation.
type CompileOptions struct {
	ModelType ModelType

	// For ModelTypeMIL:
	MILText []byte // MIL program text
	// Legacy single-weight MIL fields.
	WeightBlob []byte // weight binary blob (may be nil)
	WeightPath string // model path key for weight dict (default: "@model_path/weights/weight.bin")
	// WeightFiles allows MIL graphs with multiple named BLOBFILE inputs.
	// If both WeightBlob and WeightFiles are set, they are merged.
	WeightFiles []WeightFile

	// For ModelTypePackage:
	PackagePath string // path to .mlmodelc directory
	ModelKey    string // key for model dictionary (default: "s")

	QoS           uint32 // quality-of-service class (0 = default 21)
	PerfStatsMask uint32 // perf stats bitmask (0 = disabled)
}

// WeightFile describes a named MIL BLOBFILE entry.
type WeightFile struct {
	Path string
	Blob []byte
}

// EvalStats contains hardware performance statistics from an evaluation.
type EvalStats struct {
	HWExecutionNS uint64 // hardware execution time in nanoseconds
}

// Diagnostics contains best-effort runtime diagnostic information about a kernel.
type Diagnostics struct {
	ModelQueueDepth      int
	ModelQueueDepthKnown bool
	ProgramClass         string
	ProgramClassKnown    bool
}

// SharedEventEvalOptions configures shared event evaluation behavior.
type SharedEventEvalOptions struct {
	DisableIOFencesUseSharedEvents bool // set kANEFDisableIOFencesUseSharedEventsKey
	EnableFWToFWSignal             bool // set kANEFEnableFWToFWSignal (keep false for ANE→Metal on physical hosts)
}
