//go:build darwin

package ane

import (
	"github.com/tmc/apple/objc"
)

// Diagnostics returns best-effort diagnostic information about the kernel.
// Fields marked "Known" are true only if the corresponding selector was available.
func (k *Kernel) Diagnostics() Diagnostics {
	var d Diagnostics

	switch k.modelType {
	case ModelTypeMIL:
		d.ProgramClass = "ANEInMemoryModel"
		d.ProgramClassKnown = true
	case ModelTypePackage:
		d.ProgramClass = "ANEModel"
		d.ProgramClassKnown = true

		if depth, ok := probeQueueDepth(k.model.ID); ok {
			d.ModelQueueDepth = depth
			d.ModelQueueDepthKnown = true
		}
	}

	return d
}

// probeQueueDepth attempts to read the queueDepth property from an ANEModel.
func probeQueueDepth(id objc.ID) (depth int, ok bool) {
	defer func() { recover() }()
	rv := objc.Send[int8](id, objc.Sel("queueDepth"))
	return int(rv), true
}
