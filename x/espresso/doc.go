// Package espresso provides high-level access to Apple's Espresso neural
// network execution framework with multi-backend support (CPU, GPU, ANE).
//
// Unlike [github.com/tmc/apple/x/ane], which targets only the Apple Neural
// Engine, Espresso can dispatch individual layers to the best available
// backend. This is the same execution path that CoreML uses internally.
//
//	ctx, err := espresso.Open()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ctx.Close()
//
//	net, err := ctx.LoadNetwork(irJSON)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer net.Close()
//
//	f := espresso.NewFrame()
//	f.SetInput("input", data, espresso.Shape{Batch: 1, Channels: 3, Height: 224, Width: 224})
//	if err := net.Eval(f); err != nil {
//		log.Fatal(err)
//	}
//	output, _ := f.Output("output")
//
// # Architecture
//
// Espresso is Apple's private neural network IR and execution framework.
// It sits between CoreML/MIL and hardware backends (ANE, GPU via Metal
// Performance Shaders, CPU). Where x/ane provides direct ANE-only access,
// Espresso is the multi-backend orchestrator that CoreML uses internally.
//
// The key advantage over x/ane is per-layer backend selection: Espresso
// can run some layers on ANE, fall back to GPU for unsupported ops, and
// use CPU as a last resort. x/ane requires the entire model to be
// ANE-compatible.
//
//	CoreML (.mlmodel) → Espresso IR (JSON + weights)
//	                         │
//	              ┌──────────┼──────────┐
//	              ANE        GPU        CPU
//
// # Execution
//
// A [Context] manages execution state and selects the platform.
// [Frame] binds named input and output tensors using string keys
// matching the model's tensor names. [Network.Eval] executes the network.
//
// Note: [Context.LoadNetwork] wraps EspressoNetwork.initWithJSFile, which
// is an internal CoreML entry point that requires the full CoreML model
// loading context. It cannot be called standalone with raw Espresso IR.
// Use [CompileAndLoadEspresso] (via _ANEClient.compileModel) to load
// models from IR instead.
//
// # Profiling
//
// [ReadProfile] extracts per-layer execution data from an
// EspressoProfilingNetworkInfo, including which engine ran each layer,
// runtime samples, and platform support. This is useful for
// understanding how Espresso partitions a model across backends.
//
// [ProfileNetwork] cannot extract profiling data directly from a loaded
// network because the EspressoNetwork class does not expose profiling
// selectors in the generated bindings. Use ReadProfile with an externally
// obtained EspressoProfilingNetworkInfo instead.
//
// # Optimization Passes
//
// The [Pass] interface wraps Espresso's built-in optimization passes
// (fusion, strength reduction, quantization, normalization merging).
// Use [Optimize] to apply a standard set, or [OptimizeWith] for a
// custom pipeline.
//
// # Metal Interop
//
// For networks using the ANE backend, [Network.MetalBuffer] provides
// zero-copy GPU access to ANE IOSurfaces via Metal buffers.
package espresso
