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
	HasANE        bool
	NumCores      uint32
	Architecture  string
	Product       string
	BuildVersion  string
	IsVM          bool
	NumANEs       uint32
	SubType       string
	BoardType     int64
	InternalBuild bool
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

	// QoS selects the ANE quality-of-service scheduling class.
	// Zero uses the default value of 21. Higher values may receive
	// priority on shared hardware.
	QoS uint32

	// PerfStatsMask enables hardware performance counters during evaluation.
	// Zero disables stats collection. When non-zero, telemetry.EvalWithStats
	// returns counter names and timing data. Valid masks are 4 bits wide
	// (0x1–0xF); higher values are silently converted to zero by the driver.
	PerfStatsMask uint32
}

// WeightFile describes a named MIL BLOBFILE entry.
type WeightFile struct {
	Path string
	Blob []byte
}

// SharedEventEvalOptions configures shared event evaluation behavior.
type SharedEventEvalOptions struct {
	DisableIOFencesUseSharedEvents bool // set kANEFDisableIOFencesUseSharedEventsKey
	EnableFWToFWSignal             bool // set kANEFEnableFWToFWSignal (keep false for ANE→Metal on physical hosts)
}

// CompileStats contains timing measurements from model compilation.
type CompileStats struct {
	CompileNS int64 // wall-clock compile phase
	LoadNS    int64 // wall-clock load phase
	TotalNS   int64 // wall-clock total
}

// Available reports whether any timing data was collected.
func (s CompileStats) Available() bool { return s.TotalNS > 0 }

// ReportMetrics reports compilation timing to a testing.B-compatible reporter.
func (s CompileStats) ReportMetrics(b interface{ ReportMetric(float64, string) }) {
	if !s.Available() {
		return
	}
	b.ReportMetric(float64(s.CompileNS), "compile-ns/op")
	b.ReportMetric(float64(s.LoadNS), "load-ns/op")
	b.ReportMetric(float64(s.TotalNS), "compile-total-ns/op")
}

// String returns a compact human-readable summary of compile timing.
func (s CompileStats) String() string {
	if !s.Available() {
		return "CompileStats{}"
	}
	return fmt.Sprintf("CompileStats{compile=%dns load=%dns total=%dns}", s.CompileNS, s.LoadNS, s.TotalNS)
}
