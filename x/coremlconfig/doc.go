// Package coremlconfig provides compute plan analysis and model
// configuration for CoreML models.
//
// Load a compute plan from a compiled model bundle to inspect which
// operations target the ANE, GPU, or CPU:
//
//	plan, err := coremlconfig.LoadPlan("model.mlmodelc", coremlconfig.PlanOptions{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, op := range plan.Operations() {
//		fmt.Printf("%-20s device=%-4s cost=%.3f\n",
//			op.Name, op.Device, op.Cost)
//	}
//
// PlanOptions.ComputeUnits controls the target compute unit selection. Its
// zero value is [ComputeUnitsDefault], which allows every available unit
// including the ANE; restrict a plan with an explicit value such as
// [ComputeUnitsCPUOnly].
//
// The devices reported by a plan are CoreML's own planning values, not a
// measurement of execution. A compute-unit selection constrains the planner;
// it does not establish where an operation ran.
//
// # Model Configuration
//
// [NewConfig] creates an MLModelConfiguration with helpers for
// compute unit selection and private ANE compiler options.
package coremlconfig
