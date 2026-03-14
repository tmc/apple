//go:build darwin

package telemetry

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/x/ane"
)

// ProbeDiagnostics returns best-effort diagnostic information about the kernel.
// Fields marked "Known" are true only if the corresponding selector was available.
func ProbeDiagnostics(m *ane.Model) Diagnostics {
	var d Diagnostics

	switch m.CompileModelType() {
	case ane.ModelTypeMIL:
		d.ProgramClass = "ANEInMemoryModel"
		d.ProgramClassKnown = true
		modelID := objc.ID(m.InMemModelObjcID())
		probeModelDiags(modelID, &d)

	case ane.ModelTypePackage:
		d.ProgramClass = "ANEModel"
		d.ProgramClassKnown = true
		modelID := objc.ID(m.ModelObjcID())
		probeModelDiags(modelID, &d)
	}

	return d
}

func probeModelDiags(id objc.ID, d *Diagnostics) {
	if depth, ok := probeQueueDepth(id); ok {
		d.ModelQueueDepth = depth
		d.ModelQueueDepthKnown = true
	}
	if prog := probeProgramForEval(id); prog != 0 {
		probeProgramDiags(prog, d)
	}
	if handle, ok := probeProgramHandle(id); ok {
		d.ProgramHandle = handle
		d.ProgramHandleKnown = true
	}
	if state, ok := probeModelState(id); ok {
		d.ModelState = state
		d.ModelStateKnown = true
	}
}

func probeQueueDepth(id objc.ID) (depth int, ok bool) {
	defer func() { recover() }()
	rv := objc.Send[int8](id, objc.Sel("queueDepth"))
	return int(rv), true
}

func probeProgramForEval(id objc.ID) objc.ID {
	defer func() { recover() }()
	return objc.Send[objc.ID](id, objc.Sel("program"))
}

func probeProgramDiags(prog objc.ID, d *Diagnostics) {
	func() {
		defer func() { recover() }()
		rv := objc.Send[int64](prog, objc.Sel("currentAsyncRequestsInFlight"))
		d.AsyncRequestsInFlight = rv
		d.AsyncRequestsInFlightKnown = true
	}()
	func() {
		defer func() { recover() }()
		rv := objc.Send[int8](prog, objc.Sel("queueDepth"))
		d.ProgramQueueDepth = int(rv)
		d.ProgramQueueDepthKnown = true
	}()
	func() {
		defer func() { recover() }()
		rv := objc.Send[uint64](prog, objc.Sel("programHandle"))
		if !d.ProgramHandleKnown {
			d.ProgramHandle = rv
			d.ProgramHandleKnown = true
		}
	}()
}

func probeProgramHandle(id objc.ID) (uint64, bool) {
	defer func() { recover() }()
	rv := objc.Send[uint64](id, objc.Sel("programHandle"))
	return rv, true
}

func probeModelState(id objc.ID) (uint64, bool) {
	defer func() { recover() }()
	rv := objc.Send[uint64](id, objc.Sel("state"))
	return rv, true
}
