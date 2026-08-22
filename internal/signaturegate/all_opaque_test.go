// Copyright 2026 The tmc/apple Authors. All rights reserved.

package signaturegate

import (
	"go/ast"
	"os"
	"sort"
	"strings"
	"testing"
)

// A record whose every field is unsafe.Pointer has no layout.
//
// This gate takes no baseline, and that is the point of it. Every other check
// in this package is a difference over a window -- against a pinned revision,
// against HEAD -- and a difference cannot see a constant. The
// applicationservices damage proves the gap rather than illustrating it: nine
// typed fields of CMDeviceInfo became unsafe.Pointer on 2026-08-02, and it was
// invisible to the pinned baseline (77eeb3409), to the working-tree baseline,
// and to a 313-package API recensus, because all three windows opened after it
// landed. It was found only by widening a baseline for an unrelated reason.
//
// The failure is worth stating plainly: a struct that loses its layout gets
// wider (a uint32 field becomes an 8-byte pointer, so every offset after it is
// wrong across the C boundary) and gains GC roots (unsafe.Pointer is scanned,
// so the collector follows integers as if they were pointers). It still
// compiles, and every gate keyed on compilation stays green.
//
// The check is cheap because the population is small: 44 of 2,147 records with
// two or more fields. That is reviewable by hand, which a 2,419-entry list
// derived from a baseline diff is not.
//
// Each record is listed with a reason, exactly like knownFindings and for the
// same argument: never a pattern, so nothing else can hide behind an entry,
// and an entry that stops reproducing is a failure rather than a pass. Repair
// a record and this gate goes red until its line is deleted.
//
// Reasons are not decoration. Three dispositions are live here and they must
// not be allowed to blur:
//
//	opaque by design  the C declaration really is opaque
//	DEFECT            adjudicated against the SDK header; layout was lost
//	UNADJUDICATED     nobody has read the header yet
//
// An UNADJUDICATED entry is suppressed exactly as effectively as an adjudicated
// one. It is written that way so the list cannot be read as 44 reviewed
// decisions when 39 of them are.
var allOpaqueRecords = map[string]string{
	// DEFECT, adjudicated 2026-08-22 against ColorSyncDeprecated.h. The header
	// declares typed fields -- UInt32 dataVersion, CMDeviceClass deviceClass,
	// CMDeviceScope deviceScope, CFDictionaryRef *deviceName -- and the tree
	// renders all nine as unsafe.Pointer. origin/main rendered them correctly,
	// so this is lost ground, not ground never held.
	//
	// Mechanism, a LEAD and not a conclusion: all four live in
	// QD.framework, a sub-framework nested inside ApplicationServices, which
	// carries 1 top-level header and 11 nested frameworks. The regen log for
	// this package says "2 unplaced by any header". The counts do not match
	// (four records, two typedefs), so the correlation is suggestive only.
	//
	// Held for v0.7.1: the API has been deprecated since Mac OS X 10.6, so the
	// exposure does not justify holding a 313-package release.
	"applicationservices.CMDeviceInfo":         "DEFECT v0.7.1: QD.framework nested header, layout lost",
	"applicationservices.CMDeviceProfileArray": "DEFECT v0.7.1: QD.framework nested header, layout lost",
	"applicationservices.CMDeviceScope":        "DEFECT v0.7.1: QD.framework nested header, layout lost",
	"applicationservices.CMMultiFunctLutType":  "DEFECT v0.7.1: QD.framework nested header, layout lost",

	// SUSPECTED same defect, and the reason it is grouped here rather than
	// guessed at: MDImporter.h lives in Metadata.framework, nested inside
	// CoreServices.framework -- the same shape as QD.framework above. Two
	// unrelated frameworks losing layout for records that share only the
	// property of being declared in a nested sub-framework is what raised the
	// mechanism from coincidence to lead. Their true fields are a CFPlugIn
	// COM-style mix of function pointers and a refcount, so all-opaque is less
	// obviously wrong here than for CMDeviceInfo, and it has NOT been
	// adjudicated field by field.
	"coreservices.MDExporterInterfaceStruct":                 "UNADJUDICATED: Metadata.framework nested header, same shape as QD",
	"coreservices.MDImporterBundleWrapperURLInterfaceStruct": "UNADJUDICATED: Metadata.framework nested header, same shape as QD",
	"coreservices.MDImporterInterfaceStruct":                 "UNADJUDICATED: Metadata.framework nested header, same shape as QD",
	"coreservices.MDImporterURLInterfaceStruct":              "UNADJUDICATED: Metadata.framework nested header, same shape as QD",

	// Opaque by design. The IOKit and libkern surface is C++: OSObject and its
	// subclasses have vtables and private members, and no C declaration of
	// their storage exists to render. An opaque handle is the correct Go
	// rendering, not a degraded one.
	"kernel.AVCSubunitInfo":                    "opaque by design: IOKit/libkern C++ class",
	"kernel.IOATACommand":                      "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFWAsyncPHYCommand":               "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFWIsochChannel":                  "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFireWireAVCAsynchronousCommand":  "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFireWireMultiIsochReceivePacket": "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFireWireSBP2ORB":                 "opaque by design: IOKit/libkern C++ class",
	"kernel.IOFramebuffer":                     "opaque by design: IOKit/libkern C++ class",
	"kernel.IOHIDEventService":                 "opaque by design: IOKit/libkern C++ class",
	"kernel.IOMemoryDescriptor":                "opaque by design: IOKit/libkern C++ class",
	"kernel.IONotifier":                        "opaque by design: IOKit/libkern C++ class",
	"kernel.IORegistryEntry":                   "opaque by design: IOKit/libkern C++ class",
	"kernel.IORegistryIterator":                "opaque by design: IOKit/libkern C++ class",
	"kernel.IOService":                         "opaque by design: IOKit/libkern C++ class",
	"kernel.OSArray":                           "opaque by design: IOKit/libkern C++ class",
	"kernel.OSBoolean":                         "opaque by design: IOKit/libkern C++ class",
	"kernel.OSCollectionIterator":              "opaque by design: IOKit/libkern C++ class",
	"kernel.OSData":                            "opaque by design: IOKit/libkern C++ class",
	"kernel.OSDictionary":                      "opaque by design: IOKit/libkern C++ class",
	"kernel.OSMetaClassBase":                   "opaque by design: IOKit/libkern C++ class",
	"kernel.OSNumber":                          "opaque by design: IOKit/libkern C++ class",
	"kernel.OSObject":                          "opaque by design: IOKit/libkern C++ class",
	"kernel.OSOrderedSet":                      "opaque by design: IOKit/libkern C++ class",
	"kernel.OSSerialize":                       "opaque by design: IOKit/libkern C++ class",
	"kernel.OSSerializer":                      "opaque by design: IOKit/libkern C++ class",
	"kernel.OSSet":                             "opaque by design: IOKit/libkern C++ class",
	"kernel.OSString":                          "opaque by design: IOKit/libkern C++ class",
	"kernel.OSSymbol":                          "opaque by design: IOKit/libkern C++ class",

	// UNADJUDICATED. The smrq_* records are the kernel's intrusive queue links,
	// whose members are pointers in C, so all-opaque is plausibly right -- but
	// plausibly is not measured.
	"kernel.Smrq_link":        "UNADJUDICATED: intrusive queue link, members are pointers in C",
	"kernel.Smrq_stailq_head": "UNADJUDICATED: intrusive queue link, members are pointers in C",
	"kernel.Smrq_tailq_head":  "UNADJUDICATED: intrusive queue link, members are pointers in C",

	// UNADJUDICATED. Private-framework records with no public header to read,
	// so the SDK cannot settle these the way it settled CMDeviceInfo.
	//
	// Keyed by Go PACKAGE name, not by directory. These three live under
	// private/ but declare packages diskimages2, gtshaderprofiler and
	// virtualization; the first draft of this list used the directory and the
	// gate reported all three as both unlisted and stale in one run. A key that
	// names the wrong thing is suppression that protects nothing.
	"diskimages2.AAAsyncByteStreamImpl":                   "UNADJUDICATED: private framework, no public header",
	"gtshaderprofiler.GTAGX2ShaderProfilerProgramAddress": "UNADJUDICATED: private framework, no public header",
	"virtualization.PluginIdentifier":                     "UNADJUDICATED: private framework, no public header",

	// UNADJUDICATED. CSSM is a deprecated C API whose headers are no longer
	// shipped in the SDK, so there is nothing to adjudicate against on this
	// machine -- absence of the header is not evidence the rendering is right.
	"security.CE_SemanticsInformation": "UNADJUDICATED: CSSM headers absent from 25F70",
	"security.Cssm_tp_authority_id":    "UNADJUDICATED: CSSM headers absent from 25F70",
}

// allOpaqueRecordsInTree returns every generated record with at least two
// fields whose fields are all unsafe.Pointer, keyed as package.TypeName.
//
// The two-field floor is deliberate. A single-field wrapper around
// unsafe.Pointer is the ordinary rendering of an opaque handle and carries no
// layout claim to lose; including those would bury the finding in hundreds of
// correct records, which is how an allowlist stops being read.
func allOpaqueRecordsInTree(t *testing.T) (all map[string]int, records int) {
	t.Helper()
	root := repoRoot(t)
	all = map[string]int{}
	for _, path := range generatedFiles(t, root) {
		name := rel(root, path)
		src, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		file, err := parseGo(name, src)
		if err != nil {
			continue
		}
		pkg := file.Name.Name
		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok || st.Fields == nil {
				return true
			}
			var fields, opaque int
			for _, f := range st.Fields.List {
				n := len(f.Names)
				if n == 0 {
					n = 1 // embedded
				}
				fields += n
				if isUnsafePointer(f.Type) {
					opaque += n
				}
			}
			if fields < 2 {
				return true
			}
			records++
			if fields == opaque {
				all[pkg+"."+ts.Name.Name] = fields
			}
			return true
		})
	}
	return all, records
}

func isUnsafePointer(e ast.Expr) bool {
	sel, ok := e.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Pointer" {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	return ok && id.Name == "unsafe"
}

func TestNoAllOpaqueRecords(t *testing.T) {
	all, records := allOpaqueRecordsInTree(t)
	if records == 0 {
		t.Fatal("no multi-field records found; the gate would pass without inspecting anything")
	}
	t.Logf("inspected %d record(s) with >=2 fields; %d are entirely unsafe.Pointer", records, len(all))

	var unexpected []string
	for name, fields := range all {
		if _, known := allOpaqueRecords[name]; !known {
			unexpected = append(unexpected, name+" ("+itoa(fields)+" fields)")
		}
	}
	if len(unexpected) > 0 {
		sort.Strings(unexpected)
		t.Errorf("%d record(s) have no layout and are not listed in allOpaqueRecords:\n\t%s\n"+
			"a record whose every field is unsafe.Pointer is wider than its C declaration and "+
			"adds GC roots; it still compiles, so no build-based gate can see it",
			len(unexpected), strings.Join(unexpected, "\n\t"))
	}

	// A listed record that is no longer all-opaque means it was repaired, and
	// the entry has to go or it becomes a place for the next one to hide.
	var stale []string
	for name := range allOpaqueRecords {
		if _, found := all[name]; !found {
			stale = append(stale, name)
		}
	}
	if len(stale) > 0 {
		sort.Strings(stale)
		t.Errorf("%d entry(s) in allOpaqueRecords no longer reproduce; delete them:\n\t%s",
			len(stale), strings.Join(stale, "\n\t"))
	}
}

// TestAllOpaqueGateCanFail is the mutation control. Without it a gate whose
// scan silently found nothing would report the same green as a clean tree --
// the failure mode this package has hit more than once.
func TestAllOpaqueGateCanFail(t *testing.T) {
	const src = `package sample

import "unsafe"

type Damaged struct {
	A unsafe.Pointer
	B unsafe.Pointer
}

type Healthy struct {
	A uint32
	B unsafe.Pointer
}

type Handle struct {
	p unsafe.Pointer
}
`
	file, err := parseGo("sample.gen.go", []byte(src))
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	ast.Inspect(file, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok || st.Fields == nil {
			return true
		}
		var fields, opaque int
		for _, f := range st.Fields.List {
			c := len(f.Names)
			if c == 0 {
				c = 1
			}
			fields += c
			if isUnsafePointer(f.Type) {
				opaque += c
			}
		}
		if fields >= 2 && fields == opaque {
			got[ts.Name.Name] = true
		}
		return true
	})
	if !got["Damaged"] {
		t.Error("gate did not flag an all-opaque record; it cannot fail and its green means nothing")
	}
	if got["Healthy"] {
		t.Error("gate flagged a record with a typed field; it would fail on correct output")
	}
	// The single-field floor is a deliberate exclusion, so it gets a control
	// of its own rather than being trusted to stay right.
	if got["Handle"] {
		t.Error("gate flagged a single-field opaque handle; that is the correct rendering, not damage")
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
