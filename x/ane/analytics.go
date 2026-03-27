//go:build darwin

package ane

import (
	"fmt"

	pe "github.com/tmc/apple/private/espresso"
)

// CompilerAnalytics wraps _ANECompilerAnalytics for inspecting the
// ANE compiler's layer-by-layer breakdown and task profiling data.
type CompilerAnalytics struct {
	raw pe.ANECompilerAnalytics
}

// AnalyticsGroup describes a group of compiled layers.
type AnalyticsGroup struct {
	GroupID int
	Layers  []AnalyticsLayer
}

// AnalyticsLayer describes a single compiled layer.
type AnalyticsLayer struct {
	Name   string
	Weight float64
}

// ParseAnalytics creates a CompilerAnalytics from a raw NSData buffer.
// The buffer is typically obtained from the ANE compiler output.
func ParseAnalytics(data []byte) (*CompilerAnalytics, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("ane analytics: empty buffer")
	}
	raw := pe.NewANECompilerAnalytics()
	if raw.GetID() == 0 {
		return nil, fmt.Errorf("ane analytics: failed to allocate _ANECompilerAnalytics")
	}
	return &CompilerAnalytics{raw: raw}, nil
}

// ParseAnalyticsFromObj creates a CompilerAnalytics from an existing
// _ANECompilerAnalytics ObjC object.
func ParseAnalyticsFromObj(raw pe.ANECompilerAnalytics) *CompilerAnalytics {
	return &CompilerAnalytics{raw: raw}
}

// Populate triggers the analytics engine to parse the buffer.
// Returns false if the buffer contains no valid analytics data.
func (a *CompilerAnalytics) Populate() bool {
	return a.raw.PopulateAnalytics()
}

// Groups returns the analytics groups extracted from the compilation.
// Each group contains layers and associated task metrics.
func (a *CompilerAnalytics) Groups() []AnalyticsGroup {
	procs := a.raw.ProcedureAnalytics()
	if procs == nil {
		return nil
	}
	count := procs.Count()
	var groups []AnalyticsGroup
	for i := uint(0); i < count; i++ {
		obj := procs.ObjectAtIndex(i)
		grp := pe.ANEAnalyticsGroupFromID(obj.GetID())
		if grp.GetID() == 0 {
			continue
		}
		ag := AnalyticsGroup{
			GroupID: grp.GroupID().IntValue(),
		}
		layerArr := grp.LayerInfo()
		if layerArr != nil {
			layerCount := layerArr.Count()
			for j := uint(0); j < layerCount; j++ {
				lobj := layerArr.ObjectAtIndex(j)
				layer := pe.ANEAnalyticsLayerFromID(lobj.GetID())
				if layer.GetID() == 0 {
					continue
				}
				ag.Layers = append(ag.Layers, AnalyticsLayer{
					Name:   layer.LayerName(),
					Weight: layer.Weight().DoubleValue(),
				})
			}
		}
		groups = append(groups, ag)
	}
	return groups
}

// Serialize returns the analytics data as a serialized ObjC object
// (typically an NSDictionary). Returns nil if serialization fails.
func (a *CompilerAnalytics) Serialize() uintptr {
	obj := a.raw.Serialize()
	if obj == nil {
		return 0
	}
	return uintptr(obj.GetID())
}

// BufferSize returns the size of the analytics buffer in bytes.
func (a *CompilerAnalytics) BufferSize() int {
	return a.raw.BufferSizeInBytes().IntValue()
}
