//go:build darwin

package coremlconfig

import (
	"context"
	"fmt"
	"time"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	privatecoreml "github.com/tmc/apple/private/coreml"
)

// ComputeDevice represents a target compute device for an operation.
type ComputeDevice int

const (
	DeviceUnknown      ComputeDevice = iota
	DeviceCPU                        // CPU execution
	DeviceGPU                        // GPU execution
	DeviceNeuralEngine               // Apple Neural Engine
)

func (d ComputeDevice) String() string {
	switch d {
	case DeviceCPU:
		return "CPU"
	case DeviceGPU:
		return "GPU"
	case DeviceNeuralEngine:
		return "ANE"
	default:
		return "unknown"
	}
}

// OperationPlan describes the planned execution for a single ML Program operation.
type OperationPlan struct {
	Name   string        // operator name (e.g. "conv", "linear", "softmax")
	Device ComputeDevice // preferred compute device
	Cost   float64       // estimated workload fraction [0.0, 1.0]

	// Supported lists every device that can execute this operation.
	// It is usually wider than Device: CoreML routinely reports an
	// operation as supported on the Neural Engine while preferring the
	// CPU for it. A zero [Plan.ANEFraction] therefore means "CoreML did
	// not choose the Neural Engine", not "the Neural Engine cannot run
	// this"; only Supported answers the latter.
	Supported []ComputeDevice
	Path      []string // hierarchical path components from MLModelStructurePath
	MilID     int64    // MIL operation ID for source mapping
}

// SupportState classifies a CoreML validation/support message.
type SupportState int64

// SupportStatePattern describes a matcher regex and the support state it maps to.
type SupportStatePattern struct {
	Regex string
	State SupportState
}

// Plan holds a loaded compute plan and its analyzed operations.
type Plan struct {
	raw        coreml.MLComputePlan
	operations []OperationPlan
	units      coreml.MLComputeUnits
	unreadable string
}

// Analyzable reports whether the plan's operations could be read.
//
// CoreML exposes per-operation devices only for models in the mlprogram
// form; for any other structure, such as a NeuralNetwork model, the
// operations cannot be read at all. When Analyzable is false,
// [Plan.Operations] is empty and every cost fraction is zero, which is
// indistinguishable from a model that genuinely places no work on a
// device. Check it before reading a zero fraction as a measurement.
func (p *Plan) Analyzable() bool { return p.unreadable == "" }

// Unreadable returns why the plan's operations could not be read, or the
// empty string when they were read.
func (p *Plan) Unreadable() string { return p.unreadable }

// ComputeUnits selects which compute units a plan load targets.
//
// The zero value is [ComputeUnitsDefault], which lets CoreML use every
// available unit. It is deliberately distinct from [ComputeUnitsCPUOnly]:
// coreml.MLComputeUnitsCPUOnly is itself zero, so a field of that type
// cannot distinguish an explicit CPU-only request from an unset field.
type ComputeUnits int

const (
	// ComputeUnitsDefault leaves the choice to CoreML, which uses every
	// available unit including the Neural Engine.
	ComputeUnitsDefault ComputeUnits = iota
	// ComputeUnitsCPUOnly limits the model to the CPU.
	ComputeUnitsCPUOnly
	// ComputeUnitsCPUAndGPU allows the CPU and GPU, but not the Neural Engine.
	ComputeUnitsCPUAndGPU
	// ComputeUnitsCPUAndNeuralEngine allows the CPU and Neural Engine, but not the GPU.
	ComputeUnitsCPUAndNeuralEngine
	// ComputeUnitsAll allows every available unit, including the Neural Engine.
	ComputeUnitsAll
)

func (u ComputeUnits) String() string {
	switch u {
	case ComputeUnitsDefault:
		return "default"
	case ComputeUnitsCPUOnly:
		return "CPUOnly"
	case ComputeUnitsCPUAndGPU:
		return "CPUAndGPU"
	case ComputeUnitsCPUAndNeuralEngine:
		return "CPUAndNeuralEngine"
	case ComputeUnitsAll:
		return "All"
	default:
		return fmt.Sprintf("ComputeUnits(%d)", int(u))
	}
}

// resolve maps u to the CoreML value actually applied to a configuration.
// ComputeUnitsDefault resolves to MLComputeUnitsAll; every other value maps
// to its exact CoreML counterpart. An unrecognized value is an error rather
// than a silent substitution.
func (u ComputeUnits) resolve() (coreml.MLComputeUnits, error) {
	switch u {
	case ComputeUnitsDefault, ComputeUnitsAll:
		return coreml.MLComputeUnitsAll, nil
	case ComputeUnitsCPUOnly:
		return coreml.MLComputeUnitsCPUOnly, nil
	case ComputeUnitsCPUAndGPU:
		return coreml.MLComputeUnitsCPUAndGPU, nil
	case ComputeUnitsCPUAndNeuralEngine:
		return coreml.MLComputeUnitsCPUAndNeuralEngine, nil
	default:
		return 0, fmt.Errorf("coremlconfig: unknown compute units %d", int(u))
	}
}

// PlanOptions configures compute plan loading.
type PlanOptions struct {
	// ComputeUnits selects which compute units to target.
	// The zero value is [ComputeUnitsDefault]; to restrict a plan to the
	// CPU, set [ComputeUnitsCPUOnly] explicitly.
	ComputeUnits ComputeUnits

	// Timeout for the async plan load. Zero means 30 seconds.
	Timeout time.Duration
}

// LoadPlan loads and analyzes a compute plan from a compiled .mlmodelc bundle.
// The path should point to the .mlmodelc directory on disk.
func LoadPlan(path string, opts PlanOptions) (*Plan, error) {
	url := foundation.NewURLFileURLWithPathIsDirectory(path, true)
	if url.GetID() == 0 {
		return nil, fmt.Errorf("coremlconfig: invalid path %s", path)
	}

	units, err := opts.ComputeUnits.resolve()
	if err != nil {
		return nil, err
	}

	cfg := coreml.NewMLModelConfiguration()
	cfg.SetComputeUnits(units)

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	raw, err := coreml.GetMLComputePlanClass().LoadContentsOfURLConfiguration(ctx, url, cfg)
	if err != nil {
		return nil, fmt.Errorf("coremlconfig: load plan: %w", err)
	}
	if raw == nil || raw.GetID() == 0 {
		return nil, fmt.Errorf("coremlconfig: nil plan for %s", path)
	}

	p := &Plan{raw: *raw, units: units}
	p.operations, p.unreadable = p.analyzeOperations()
	return p, nil
}

// RequestedComputeUnits reports the CoreML compute-unit selection that was
// applied to this plan's configuration.
//
// It records what was requested, not where CoreML ultimately placed each
// operation: compute units are a constraint on the planner, and the per
// operation devices reported by [Plan.Operations] are CoreML's own plan
// values. Neither is a measurement of execution.
func (p *Plan) RequestedComputeUnits() coreml.MLComputeUnits {
	return p.units
}

// Operations returns the analyzed operation plans.
func (p *Plan) Operations() []OperationPlan {
	return p.operations
}

// ANEOperations returns only operations targeting the Neural Engine.
func (p *Plan) ANEOperations() []OperationPlan {
	return p.filterByDevice(DeviceNeuralEngine)
}

// GPUOperations returns only operations targeting the GPU.
func (p *Plan) GPUOperations() []OperationPlan {
	return p.filterByDevice(DeviceGPU)
}

// CPUOperations returns only operations targeting the CPU.
func (p *Plan) CPUOperations() []OperationPlan {
	return p.filterByDevice(DeviceCPU)
}

// ANEFraction returns the fraction of total estimated cost on the Neural Engine.
func (p *Plan) ANEFraction() float64 {
	return p.costFraction(DeviceNeuralEngine)
}

// GPUFraction returns the fraction of total estimated cost on the GPU.
func (p *Plan) GPUFraction() float64 {
	return p.costFraction(DeviceGPU)
}

// CPUFraction returns the fraction of total estimated cost on the CPU.
func (p *Plan) CPUFraction() float64 {
	return p.costFraction(DeviceCPU)
}

func (p *Plan) filterByDevice(d ComputeDevice) []OperationPlan {
	var out []OperationPlan
	for _, op := range p.operations {
		if op.Device == d {
			out = append(out, op)
		}
	}
	return out
}

func (p *Plan) costFraction(d ComputeDevice) float64 {
	var total, match float64
	for _, op := range p.operations {
		total += op.Cost
		if op.Device == d {
			match += op.Cost
		}
	}
	if total == 0 {
		return 0
	}
	return match / total
}

// analyzeOperations walks the ML Program structure and queries device usage
// and cost for each operation.
// analyzeOperations reads the per-operation plan. The second result
// explains why nothing could be read, and is empty when the walk
// succeeded; an empty operation slice with an empty reason means the
// model really has no operations.
func (p *Plan) analyzeOperations() ([]OperationPlan, string) {
	structure := p.raw.ModelStructure()
	if structure == nil || structure.GetID() == 0 {
		return nil, "plan exposes no model structure"
	}

	program := structure.Program()
	if program == nil || program.GetID() == 0 {
		return nil, "model structure is not an ML program; per-operation devices are unavailable"
	}

	funcsDict := program.Functions()
	if funcsDict == nil || funcsDict.GetID() == 0 {
		return nil, "ML program exposes no functions"
	}

	var ops []OperationPlan

	allValues := objc.Send[[]objc.ID](funcsDict.GetID(), objc.Sel("allValues"))
	for _, fnID := range allValues {
		fn := coreml.MLModelStructureProgramFunctionFromID(fnID)
		block := fn.Block()
		if block == nil || block.GetID() == 0 {
			continue
		}
		ops = p.walkBlock(coreml.MLModelStructureProgramBlockFromID(block.GetID()), ops)
	}

	return ops, ""
}

// walkBlock recursively collects operations from a program block.
func (p *Plan) walkBlock(block coreml.MLModelStructureProgramBlock, ops []OperationPlan) []OperationPlan {
	for _, op := range block.Operations() {
		plan := OperationPlan{
			Name:      op.OperatorName(),
			Device:    p.deviceForOp(op),
			Cost:      p.costForOp(op),
			Supported: p.supportedForOp(op),
		}

		// Read private fields via private/coreml bindings.
		privOp := privatecoreml.MLModelStructureProgramOperationFromID(op.GetID())

		// MilId: NSNumber -> int64
		if milNum := privOp.MilId(); milNum.GetID() != 0 {
			plan.MilID = milNum.LongLongValue()
		}

		// Path: MLModelStructurePath -> []string (via Components NSArray)
		if pathObj := privOp.Path(); pathObj.GetID() != 0 {
			comps := privatecoreml.MLModelStructurePathFromID(pathObj.GetID()).Components()
			count := comps.Count()
			if count > 0 {
				plan.Path = make([]string, count)
				for i := uint(0); i < count; i++ {
					plan.Path[i] = foundation.NSStringFromID(comps.ObjectAtIndex(i).GetID()).String()
				}
			}
		}

		ops = append(ops, plan)

		for _, nested := range op.Blocks() {
			ops = p.walkBlock(nested, ops)
		}
	}
	return ops
}

func (p *Plan) deviceForOp(op coreml.MLModelStructureProgramOperation) ComputeDevice {
	usage := p.raw.ComputeDeviceUsageForMLProgramOperation(op)
	if usage == nil || usage.GetID() == 0 {
		return DeviceUnknown
	}
	dev := usage.PreferredComputeDevice()
	if dev == nil || dev.GetID() == 0 {
		return DeviceUnknown
	}
	return classifyDevice(dev.GetID())
}

// supportedForOp lists the devices that can execute op.
func (p *Plan) supportedForOp(op coreml.MLModelStructureProgramOperation) []ComputeDevice {
	usage := p.raw.ComputeDeviceUsageForMLProgramOperation(op)
	if usage == nil || usage.GetID() == 0 {
		return nil
	}
	var out []ComputeDevice
	for _, d := range usage.SupportedComputeDevices() {
		if dev := classifyDevice(d.GetID()); dev != DeviceUnknown {
			out = append(out, dev)
		}
	}
	return out
}

// SupportsANE reports whether any operation lists the Neural Engine among
// its supported devices. It answers a different question from
// [Plan.ANEFraction], which reports only what CoreML preferred.
func (p *Plan) SupportsANE() bool {
	for _, op := range p.operations {
		for _, d := range op.Supported {
			if d == DeviceNeuralEngine {
				return true
			}
		}
	}
	return false
}

func (p *Plan) costForOp(op coreml.MLModelStructureProgramOperation) float64 {
	cost := p.raw.EstimatedCostOfMLProgramOperation(op)
	if cost == nil || cost.GetID() == 0 {
		return 0
	}
	return coreml.MLComputePlanCostFromID(cost.GetID()).Weight()
}

// classifyDevice determines the ComputeDevice from an ObjC class type.
func classifyDevice(id objc.ID) ComputeDevice {
	if id == 0 {
		return DeviceUnknown
	}
	cls := objc.Send[objc.ID](id, objc.Sel("class"))
	name := objc.Send[objc.ID](cls, objc.Sel("className"))
	s := foundation.NSStringFromID(name).String()
	switch s {
	case "MLNeuralEngineComputeDevice":
		return DeviceNeuralEngine
	case "MLGPUComputeDevice":
		return DeviceGPU
	case "MLCPUComputeDevice":
		return DeviceCPU
	default:
		return DeviceUnknown
	}
}

// SupportStateForValidationMessage classifies a CoreML validation message using
// the private support-state matcher shipped with the framework.
func SupportStateForValidationMessage(message string) SupportState {
	matcher := privatecoreml.GetMLComputePlanDeviceUsageSupportStateMatcherClass().SharedInstance()
	if matcher.GetID() == 0 {
		return 0
	}
	return SupportState(matcher.MatchingSupportStateForValidationMessage(foundation.NewStringWithString(message)))
}

// SupportStatePatterns returns the framework's current support-state regex patterns.
func SupportStatePatterns() []SupportStatePattern {
	matcher := privatecoreml.GetMLComputePlanDeviceUsageSupportStateMatcherClass().SharedInstance()
	if matcher.GetID() == 0 {
		return nil
	}
	patterns := matcher.SupportStatePatterns()
	out := make([]SupportStatePattern, 0, patterns.Count())
	for i := uint(0); i < patterns.Count(); i++ {
		pattern := privatecoreml.MLComputePlanDeviceUsageSupportStatePatternFromID(patterns.ObjectAtIndex(i).GetID())
		regex := pattern.Regex()
		out = append(out, SupportStatePattern{
			Regex: regex.Pattern(),
			State: SupportState(pattern.SupportState()),
		})
	}
	return out
}
