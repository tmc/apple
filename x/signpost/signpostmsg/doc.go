// Package signpostmsg emits os_signpost intervals that carry a formatted
// message under a small fixed set of names.
//
// It is a constrained layer over [github.com/tmc/apple/x/signpost]:
// Instruments groups intervals by name and shows the message as detail, so a
// small fixed set of names with a descriptive message is usually the better
// shape for profiling: it yields one track per name rather than one per span.
//
// The names are [Model], [Layer] and [Op], matching the conventional split of
// a compute pipeline into inference passes, block boundaries and individual
// dispatches. Pass the specific identity in the message:
//
//	log := signpostmsg.New("com.example.app", signpostmsg.PointsOfInterest)
//	id := log.NewID()
//	log.IntervalBegin(id, signpostmsg.Layer, "TransformerBlock_0")
//	// ... work ...
//	log.IntervalEnd(id, signpostmsg.Layer, "TransformerBlock_0")
//
// Messages are emitted as public, so they are not redacted in trace output.
//
// When built without cgo, this package delivers the message through
// [github.com/tmc/apple/x/signpost], which constructs the "%{public}s"
// argument buffer directly. The message decodes in trace output when the
// format string is pooled in the binary's __TEXT,__oslogstring section —
// automatic under cgo, via a .syso otherwise (see x/signpost's pool.go).
// On non-Darwin platforms all operations are no-ops.
package signpostmsg
