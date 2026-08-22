// Copyright 2026 The tmc/apple Authors. All rights reserved.

// Package smoketest calls one real entry point per framework, in a subprocess,
// so that a binding which aborts the process is a failing test rather than a
// dead test run.
//
// Every other gate in this repository is compile-time. A binding that passes an
// ObjC object as the address of its Go wrapper struct compiles, vets and type
// checks; it only misbehaves when ObjC dereferences the pointer, at which point
// the process takes SIGABRT. `go build`, `go vet` and `go test` were all clean
// on the vision package while exactly that defect was live in 19 of its call
// sites, because no test called an entry point that passed such an argument.
//
// The subprocess is not a stylistic choice. The failure mode is process death,
// so an in-process t.Run cannot observe it: the harness dies with the callee and
// reports nothing. Each target is therefore re-executed as a child of the test
// binary and judged by its exit status, which turns SIGABRT into a red test.
//
// A target that cannot run headlessly is SKIPPED WITH A REASON and counted
// apart from the passes. A skipped framework is unmeasured, never clean.
//
// The crash is not always deterministic, so each target runs [repeats] times
// and any single failure fails the target. Objective-C reads the first word of
// the address it is handed as an isa pointer; that word is the wrapper's ID
// field, which holds a valid object address, so whether the process dies
// depends on what that address happens to point at.
//
// Measured over 20 runs each against an unfixed package: vision (19 defective
// sites) crashed 20 of 20, but objectivec (1 site) crashed only 9 of 20. One
// run of a defective call is therefore not evidence of anything.
//
// The consequence for a reader: a green board here is corroboration, not proof.
// A defective call can decline to crash, so only a static census of the call
// sites can show the defect is absent. This gate exists to catch what a census
// misses, not to replace it.
package smoketest

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/vision"
)

// targetEnv names the target to run when the test binary re-executes itself.
const targetEnv = "APPLE_SMOKE_TARGET"

// repeats is how many times each target runs. The crash is data-dependent, so
// one run of a defective call can pass. At the least reliable rate measured
// here (objectivec, 11 of 20 runs passing) ten consecutive passes happen about
// once in 400, which is the odds of this gate missing a live single-site defect.
const repeats = 10

// A target exercises one framework through a real entry point.
//
// Prefer an entry point that takes a CROSS-FRAMEWORK object parameter: that is
// the shape which fails, because the defect is in how such an argument is
// marshalled at the call site. An entry point taking only scalars proves the
// framework loads and nothing more.
type target struct {
	name string
	// skip, when non-empty, is why this framework cannot be exercised here.
	skip string
	// crossFramework records whether run passes an object owned by another
	// framework. A target without one is weak evidence and says so.
	crossFramework bool
	run            func() error
}

var targets = []target{{
	name: "foundation",
	// No cross-framework object: this is the control. It proves the harness
	// reports success when the call is sound, which is what makes a red
	// result elsewhere meaningful.
	run: func() error {
		if s := foundation.NSHomeDirectory().UTF8String(); s == "" {
			return errors.New("NSHomeDirectory returned an empty string")
		}
		return nil
	},
}, {
	name: "coreimage",
	run: func() error {
		img := coreimage.NewCIImage()
		if img.ID == 0 {
			return errors.New("NewCIImage returned a nil object")
		}
		_ = img.Extent()
		return nil
	},
}, {
	name:           "foundation-protocol",
	crossFramework: true,
	run: func() error {
		// Objc_getProtocol is declared as returning **Protocol, but the C
		// function returns a bare Protocol*, so the value it hands back IS the
		// protocol pointer -- do not dereference it. Rebuild the wrapper from
		// that raw pointer instead. (The declared return type is a separate
		// defect from the one this target exercises; see the census.)
		raw := objectivec.Objc_getProtocol("NSCopying")
		if raw == nil {
			return errors.New("objc_getProtocol(NSCopying) returned nil")
		}
		proto := objectivec.ProtocolFromID(objc.ID(uintptr(unsafe.Pointer(raw))))
		// Passes *objectivec.Protocol, a cross-framework object parameter.
		iface := foundation.NewXPCInterfaceWithProtocol(&proto)
		_ = iface
		return nil
	},
}, {
	name:           "objectivec",
	crossFramework: true,
	run: func() error {
		raw := objectivec.Objc_getProtocol("NSCopying")
		if raw == nil {
			return errors.New("objc_getProtocol(NSCopying) returned nil")
		}
		proto := objectivec.ProtocolFromID(objc.ID(uintptr(unsafe.Pointer(raw))))
		// A class object is itself an NSObject and answers conformsToProtocol:,
		// so this needs no instance and no other framework.
		obj := objectivec.NSObjectObjectFromID(objc.ID(objc.GetClass("NSString")))
		// conformsToProtocol: takes a single Protocol*, so passing the address
		// of the Go wrapper is the defect and not an out-parameter. Runtime
		// conformance checks go through here, so this one is load-bearing.
		_ = obj.ConformsToProtocol(&proto)
		return nil
	},
}, {
	name:           "vision",
	crossFramework: true,
	run: func() error {
		img := coreimage.NewCIImage()
		req := vision.NewVNDetectFaceRectanglesRequest()
		seq := vision.NewVNSequenceRequestHandler()
		// Passes *coreimage.CIImage, a cross-framework object parameter.
		_, err := seq.PerformRequestsOnCIImageError([]vision.VNRequest{req.VNRequest}, &img)
		// The error is irrelevant: Vision legitimately rejects a
		// zero-dimensioned image. Surviving the call is the assertion.
		_ = err
		return nil
	},
}, {
	name:           "appkit",
	crossFramework: true,
	run: func() error {
		img := coreimage.NewCIImage()
		// NSCIImageRep is an image representation, not a view, so it needs no
		// window server. It takes *coreimage.CIImage.
		rep := appkit.NewCIImageRepWithCIImage(&img)
		_ = rep
		return nil
	},
}, {
	name: "naturallanguage",
	// The site to reach is NewModelWithMLModelError, which takes *coreml.MLModel.
	skip: "naturallanguage/nl_model.gen.go:173 needs an MLModel, so a compiled .mlmodelc on disk; no fixture in-tree yet",
}, {
	name: "iosurface",
	// The site to reach is NewRenderDestinationWithIOSurface, *iosurface.IOSurface.
	skip: "coreimage/ci_render_destination.gen.go:323 needs a live IOSurface, which needs an IOSurfaceWidth/Height/BytesPerElement properties dictionary; not built here yet",
}}

// TestFrameworkSmoke runs each target in a subprocess and reports its exit
// status. Passes, failures and skips are counted separately and summarised, so
// that "everything green" cannot be confused with "everything measured".
func TestFrameworkSmoke(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Apple frameworks are only available on darwin")
	}
	var passed, failed, skipped int
	for _, tg := range targets {
		t.Run(tg.name, func(t *testing.T) {
			if tg.skip != "" {
				skipped++
				t.Skipf("SKIPPED (unmeasured, not clean): %s", tg.skip)
			}
			for i := range repeats {
				out, err := runTarget(tg.name)
				if err != nil {
					failed++
					t.Errorf("%s aborted or failed in a subprocess on run %d of %d: %v\n%s",
						tg.name, i+1, repeats, err, indent(out))
					return
				}
			}
			passed++
			if !tg.crossFramework {
				t.Logf("%s ok (control: no cross-framework object argument, "+
					"so this proves the framework loads, not that object "+
					"marshalling is correct)", tg.name)
			}
		})
	}
	t.Logf("smoke summary: %d passed, %d failed, %d SKIPPED of %d frameworks",
		passed, failed, skipped, len(targets))
}

// TestSmokeHelper is the subprocess entry point. It is a no-op unless the
// parent selected a target through the environment.
func TestSmokeHelper(t *testing.T) {
	name := os.Getenv(targetEnv)
	if name == "" {
		t.Skip("not invoked as a smoke subprocess")
	}
	for _, tg := range targets {
		if tg.name != name {
			continue
		}
		if tg.run == nil {
			t.Fatalf("target %s has no run function", name)
		}
		// AppKit and other UI frameworks assume the main thread.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		if err := tg.run(); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		return
	}
	t.Fatalf("unknown smoke target %q", name)
}

func runTarget(name string) (string, error) {
	cmd := exec.Command(os.Args[0], "-test.run=^TestSmokeHelper$", "-test.v")
	cmd.Env = append(os.Environ(), targetEnv+"="+name)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func indent(s string) string {
	const max = 24
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > max {
		lines = append(lines[:max], fmt.Sprintf("... (%d more lines)", len(lines)-max))
	}
	return "\t" + strings.Join(lines, "\n\t")
}
