//go:build darwin

package espresso

import (
	"errors"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/espresso"
)

// ANEProfile contains ANE-specific profiling data.
type ANEProfile struct {
	ANETimePerEvalNS  uint64   // nanoseconds per ANE evaluation
	TotalANETimeNS    uint64   // total ANE execution time in nanoseconds
	CompilerFileNames []string // ANE compiler analytics file names
}

// Profile contains execution profiling data for a network.
type Profile struct {
	Path   string         // model file path
	Layers []LayerProfile // per-layer profiling
	ANE    *ANEProfile    // ANE-specific profiling (nil if unavailable)
}

// LayerProfile contains profiling data for a single layer.
type LayerProfile struct {
	Name           string
	DebugName      string
	AvgRuntimeMS   float64   // average execution time in milliseconds
	RuntimeSamples []float64 // individual runtime measurements
	Engine         Engine    // selected execution engine
	Supported      bool      // whether the main engine supports this layer
	HasPerfWarning bool      // performance warning flag
	LayerType      string    // layer type identifier
}

// ReadProfile extracts profiling data from an EspressoProfilingNetworkInfo.
// This is useful when you have profiling info from the Espresso framework.
func ReadProfile(info espresso.EspressoProfilingNetworkInfo) *Profile {
	p := &Profile{
		Path: info.Network_at_path(),
	}

	// Extract ANE performance info.
	aneInfo := info.Ane_performance_info()
	if aneInfo != nil && aneInfo.GetID() != 0 {
		p.ANE = &ANEProfile{
			ANETimePerEvalNS: aneInfo.Ane_time_per_eval_ns(),
			TotalANETimeNS:   aneInfo.Total_ane_time_ns(),
		}
	}

	// Extract ANE compiler analytics.
	analytics := info.Ane_compiler_analytics()
	if analytics != nil && analytics.GetID() != 0 {
		names := analytics.Compiler_analytics_file_names()
		if names != nil && names.GetID() != 0 {
			if p.ANE == nil {
				p.ANE = &ANEProfile{}
			}
			p.ANE.CompilerFileNames = nsArrayToStrings(names)
		}
	}

	layers := info.Layers()
	if layers == nil || layers.GetID() == 0 {
		return p
	}

	arr := foundation.NSArrayFromID(layers.GetID())
	n := arr.Count()
	p.Layers = make([]LayerProfile, 0, n)

	for i := range n {
		obj := arr.ObjectAtIndex(i)
		li := espresso.EspressoProfilingLayerInfoFromID(obj.GetID())

		lp := LayerProfile{
			Name:         li.Name(),
			DebugName:    li.Debug_name(),
			AvgRuntimeMS: li.Average_runtime(),
			Engine:       Engine(li.Selected_runtime_engine()),
		}

		// Extract runtime samples.
		runtimes := li.Runtimes()
		if runtimes != nil && runtimes.GetID() != 0 {
			ra := foundation.NSArrayFromID(runtimes.GetID())
			rn := ra.Count()
			lp.RuntimeSamples = make([]float64, 0, rn)
			for j := range rn {
				robj := ra.ObjectAtIndex(j)
				// NSNumber → double via objc message send.
				v := objc.Send[float64](robj.GetID(), objc.Sel("doubleValue"))
				lp.RuntimeSamples = append(lp.RuntimeSamples, v)
			}
		}

		// Extract main engine support info.
		support := li.Main_engine_support()
		if support != nil && support.GetID() != 0 {
			si := espresso.EspressoProfilingLayerSupportInfoFromID(support.GetID())
			lp.Supported = si.Supported()
			lp.HasPerfWarning = si.Has_perf_warning()
			lp.LayerType = si.Type()
		}

		p.Layers = append(p.Layers, lp)
	}

	return p
}

// ProfileNetwork attempts to extract profiling data from a loaded network.
//
// Status: not yet implemented. The EspressoNetwork class exposes ctx(),
// layers_size(), net(), and wipe_layers_blobs() — but no profiling selector.
// Attempted selectors: profilingInfo, profileWithOptions:, _profilingNetworkInfo,
// profilingNetworkInfo. None are present in the manifest or respond at runtime.
//
// Use ReadProfile with an externally obtained EspressoProfilingNetworkInfo instead.
func ProfileNetwork(_ *Network) (*Profile, error) {
	return nil, &Error{Op: "profile", Err: errors.New("espresso: profiling not available via generated bindings; use ReadProfile with external EspressoProfilingNetworkInfo")}
}

// TotalRuntimeMS returns the sum of average runtimes across all layers.
func (p *Profile) TotalRuntimeMS() float64 {
	var total float64
	for _, l := range p.Layers {
		total += l.AvgRuntimeMS
	}
	return total
}

// LayersByEngine returns layers grouped by their execution engine.
func (p *Profile) LayersByEngine() map[Engine][]LayerProfile {
	m := make(map[Engine][]LayerProfile)
	for _, l := range p.Layers {
		m[l.Engine] = append(m[l.Engine], l)
	}
	return m
}

// UnsupportedLayers returns layers that are not supported by the main engine.
func (p *Profile) UnsupportedLayers() []LayerProfile {
	var out []LayerProfile
	for _, l := range p.Layers {
		if !l.Supported {
			out = append(out, l)
		}
	}
	return out
}

// IsANEResident reports whether this layer runs on the ANE.
func (l *LayerProfile) IsANEResident() bool {
	return l.Engine == EngineANE && l.Supported
}

// ANEResidentLayers returns layers that are resident on the ANE.
func (p *Profile) ANEResidentLayers() []LayerProfile {
	var out []LayerProfile
	for _, l := range p.Layers {
		if l.IsANEResident() {
			out = append(out, l)
		}
	}
	return out
}

// IsFullyANEResident reports whether all supported layers run on the ANE.
func (p *Profile) IsFullyANEResident() bool {
	for _, l := range p.Layers {
		if l.Supported && l.Engine != EngineANE {
			return false
		}
	}
	return len(p.Layers) > 0
}

// CPUFallbackLayers returns supported layers that fell back to CPU instead of ANE.
func (p *Profile) CPUFallbackLayers() []LayerProfile {
	var out []LayerProfile
	for _, l := range p.Layers {
		if l.Supported && l.Engine == EngineCPU {
			out = append(out, l)
		}
	}
	return out
}
