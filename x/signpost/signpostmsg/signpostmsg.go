package signpostmsg

// Name identifies a signpost. The os_signpost macros require the name to be a
// compile-time literal, so the set of names is fixed; pass the varying part of
// a span's identity as the message.
type Name int

const (
	// Model marks a high-level pass, such as one generation step.
	Model Name = iota
	// Layer marks a block boundary within a pass.
	Layer
	// Op marks a single compute dispatch.
	Op
)

// String returns the name as it appears in trace output.
func (n Name) String() string {
	switch n {
	case Model:
		return "Model"
	case Layer:
		return "Layer"
	case Op:
		return "Op"
	}
	return "Unknown"
}

// Category names understood by Instruments and the logging system. They mirror
// the constants in [github.com/tmc/apple/x/signpost].
const (
	// PointsOfInterest is the category Instruments displays in the Points of
	// Interest track.
	PointsOfInterest = "PointsOfInterest"
	// DynamicTracing is a category whose signposts are disabled until a tool
	// such as Instruments enables them.
	DynamicTracing = "DynamicTracing"
)
