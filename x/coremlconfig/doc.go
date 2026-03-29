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
// PlanOptions.ComputeUnits controls the target compute unit selection
// (default: all available units including ANE).
//
// # Model Configuration
//
// [NewConfig] creates an MLModelConfiguration with helpers for
// compute unit selection and private ANE compiler options.
package coremlconfig
