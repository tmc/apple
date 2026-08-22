//go:build darwin

package engineattest

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
	"github.com/tmc/apple/x/ane"
)

// ANE runs fn and returns an error unless m executed on Neural Engine
// hardware while fn ran. Evidence is the driver's own counter: a zeroed
// performance-stats object is attached to m's request before fn, and the
// hardware execution time it reports afterward must be nonzero.
//
// Before trusting a zero, ANE proves the counter can move by evaluating
// m once (the sensitivity canary). If that evaluation reports no
// hardware time the claim is undecidable and ANE returns
// [ErrUnattestable] — never a false "did not run". The canary is a real
// evaluation: it overwrites m's output surfaces from its current inputs.
//
// If fn returns an error, ANE returns it unchanged without judging the
// claim.
func ANE(m *ane.Model, fn func() error) error {
	if m == nil {
		return fmt.Errorf("engineattest: nil model")
	}
	canary := attachStats(m)
	if err := m.Eval(); err != nil {
		return fmt.Errorf("engineattest: sensitivity canary eval: %w", err)
	}
	if canary.hwExecutionNS() == 0 {
		return fmt.Errorf("engineattest: canary eval reported zero hardware time: %w", ErrUnattestable)
	}

	probe := attachStats(m)
	if err := fn(); err != nil {
		return err
	}
	if probe.hwExecutionNS() == 0 {
		return fmt.Errorf("engineattest: ANE hardware counter did not advance: %w", ErrDidNotRun)
	}
	return nil
}

// attachedStats is a fresh zero-valued ANEPerformanceStats attached to a
// model's request. The driver fills it (or replaces it) on each hardware
// evaluation, so a nonzero reading after a region proves the hardware ran
// during that region.
type attachedStats struct {
	req appleneuralengine.ANERequest
	ps  appleneuralengine.ANEPerformanceStats
}

func attachStats(m *ane.Model) *attachedStats {
	cls := appleneuralengine.GetANEPerformanceStatsClass()
	ps := appleneuralengine.ANEPerformanceStatsFromID(cls.StatsWithHardwareExecutionNS(0).GetID())
	req := appleneuralengine.ANERequestFromID(objc.ID(m.RawRequest()))
	req.SetPerfStats(&ps)
	return &attachedStats{req: req, ps: ps}
}

// hwExecutionNS reads the hardware execution time recorded since the
// stats object was attached. All ObjC reads are recover-guarded: on any
// firmware or class-shape surprise the answer degrades to zero, which
// the callers treat as "did not advance" (for the probe) or
// "unattestable" (for the canary) — never as a pass.
func (s *attachedStats) hwExecutionNS() uint64 {
	ps := s.ps
	// Some firmware versions report through PerfStatsArray rather than
	// the attached object; prefer the array entry when present.
	func() {
		defer func() { recover() }()
		arr := s.req.PerfStatsArray()
		if arr != nil && int(arr.Count()) > 0 {
			ps = appleneuralengine.ANEPerformanceStatsFromID(arr.ObjectAtIndex(0).GetID())
		}
	}()
	if ps.GetID() == 0 {
		func() {
			defer func() { recover() }()
			if p := s.req.PerfStats(); p != nil && p.GetID() != 0 {
				ps = appleneuralengine.ANEPerformanceStatsFromID(p.GetID())
			}
		}()
	}
	var hw uint64
	func() {
		defer func() { recover() }()
		hw = ps.HwExecutionTime()
	}()
	return hw
}
