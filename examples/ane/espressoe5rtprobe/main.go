// Command espressoe5rtprobe reports what the generated private/espresso
// binding of Espresso's e5rt_* direct-dispatch route can be shown to do, and
// refuses to report anything else.
//
// The normal compile, load, bind, and synchronous-dispatch route has been
// called through generated signatures under a CPU reference. Those 33
// signatures carry the manifest's strongest grade, but it was earned from the
// handwritten e5rt adapter's own call sites: a record that the route is
// exercised, not independent confirmation of it. The independent evidence is
// the end-to-end CPU-reference and mutation-control examples,
// espressoe5rtdispatch and espressoe5rtblock. Resolving the remaining names
// with dlsym does not establish their calling conventions.
//
// The tool keeps its shape anyway, because the hazard it was built around
// has not changed. Calling a wrong signature through purego faults in C,
// where the fault is fatal and unrecoverable — a Go recover cannot see it —
// and E5RT reports at least one failure by throwing a C++ exception, which a
// purego caller cannot catch either, so the wrapper returns success and the
// process dies later in an unrelated frame. Anything that reaches a private
// entry point belongs in a child process whatever the current confidence.
//
// # Commands
//
// The symbols command is the default and is safe: it loads the framework,
// reports which symbols resolved, and compares that set against the export
// list in the SDK's Espresso.tbd stub. It calls no e5rt_* function.
//
// The canary command proves the crash harness works, by running a child
// that jumps to an invalid address and checking that the parent reports
// CRASHED. A crash report from a harness that has not been shown to be able
// to report a crash is worth nothing.
//
// The stages command calls the unverified wrappers. Each sequence of calls
// runs in its own child process and streams one record per step it reaches,
// so a fault localizes to a named step rather than to "the example", and
// the parent survives to report it. It is gated behind both
// -allow-unverified-calls and CONFIRM_E5RT_UNVERIFIED_CALLS=call-guessed-signatures.
//
// # Outcomes
//
// No step is ever reported as OK. A step is RESOLVED or MISSING (dlsym, the
// one verified operation), RETURNED (the call returned, with its status
// code — evidence about control flow only, not that the call did what its
// name says), ERROR (the wrapper failed before calling), CRASHED, or
// UNREACHED. A step that was not run is UNREACHED, never a success.
//
//	go run ./examples/ane/espressoe5rtprobe
//	go run ./examples/ane/espressoe5rtprobe symbols -json
//	go run ./examples/ane/espressoe5rtprobe canary
//	CONFIRM_E5RT_UNVERIFIED_CALLS=call-guessed-signatures \
//	  go run ./examples/ane/e5rtprobe stages -allow-unverified-calls
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	e5rt "github.com/tmc/apple/examples/ane/internal/espressoe5rt"
)

// The stages command calls signatures nobody has confirmed. Both
// acknowledgements are required, matching examples/rdma/rdmainfo.
const (
	confirmEnv   = "CONFIRM_E5RT_UNVERIFIED_CALLS"
	confirmValue = "call-guessed-signatures"
)

// Outcomes. There is deliberately no "OK": nothing this tool can observe
// distinguishes a call that worked from a call that returned zero after
// scribbling on memory.
const (
	outcomeResolved  = "RESOLVED"  // dlsym found the name
	outcomeMissing   = "MISSING"   // dlsym did not find the name
	outcomeReturned  = "RETURNED"  // the call returned; see Status
	outcomeError     = "ERROR"     // the wrapper reported an error
	outcomeCrashed   = "CRASHED"   // the child died before reaching this step
	outcomeUnreached = "UNREACHED" // not attempted
)

func main() {
	args := os.Args[1:]
	cmd := "symbols"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		cmd, args = args[0], args[1:]
	}
	switch cmd {
	case "symbols":
		symbols(args)
	case "canary":
		canary(args)
	case "stages":
		stages(args)
	case "__sequence":
		sequenceChild(args)
	case "help", "-h", "--help":
		usage(os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "e5rtprobe: unknown command %q\n\n", cmd)
		usage(os.Stderr)
		os.Exit(2)
	}
}

func usage(w *os.File) {
	fmt.Fprintf(w, `usage: e5rtprobe [command] [options]

Commands:
  symbols   Resolve every name in e5rt.Symbols and compare against the SDK
            stub's export list. Calls no e5rt_* function. (default)
  canary    Run a child that jumps to an invalid address, to show the crash
            harness can report a crash.
  stages    Call the unverified wrappers, one child process per sequence.

Options:
  -json                      Print JSON instead of text.
  -model path                Model path for the compile sequence.
  -allow-unverified-calls    Required by stages, together with %s=%s.
`, confirmEnv, confirmValue)
}

// A record is one step of one sequence. The child emits these as JSON lines
// as it reaches them; every declared step with no record is UNREACHED.
type record struct {
	Sequence string `json:"sequence"`
	Step     string `json:"step"`
	Outcome  string `json:"outcome"`
	Status   *int64 `json:"status,omitempty"`
	Handle   string `json:"handle,omitempty"`
	Error    string `json:"error,omitempty"`
	Note     string `json:"note,omitempty"`
}

// symbols

type symbolReport struct {
	GOOS          string   `json:"goos"`
	Framework     string   `json:"framework"`
	OpenError     string   `json:"open_error,omitempty"`
	Listed        int      `json:"listed"`
	Resolved      []string `json:"resolved"`
	Unresolved    []string `json:"unresolved"`
	Stub          string   `json:"stub,omitempty"`
	StubError     string   `json:"stub_error,omitempty"`
	StubE5RTNames int      `json:"stub_e5rt_names,omitempty"`
	NotInStub     []string `json:"listed_not_in_stub,omitempty"`
	Wrapped       []string `json:"wrapped,omitempty"`
	Unwrapped     []string `json:"unwrapped,omitempty"`
}

// wrapped lists the names for which espressoe5rt has a typed method using a
// signature with the manifest's strongest grade. That grade is not independent:
// it is derived from the handwritten e5rt adapter's call sites. The remainder
// of e5rt.Symbols is resolved only; a raw caller through Lib.Sym must supply
// its own convention and run in an isolated child process. This list records
// the adapter's boundary, not a claim about every generated Espresso
// declaration.
var wrapped = map[string]bool{
	"e5rt_e5_compiler_config_options_create":                                            true,
	"e5rt_e5_compiler_config_options_set_cache_bundle_location":                         true,
	"e5rt_e5_compiler_config_options_release":                                           true,
	"e5rt_e5_compiler_create_with_config":                                               true,
	"e5rt_e5_compiler_release":                                                          true,
	"e5rt_e5_compiler_options_create":                                                   true,
	"e5rt_e5_compiler_options_set_compute_device_types_mask":                            true,
	"e5rt_e5_compiler_options_set_force_recompilation":                                  true,
	"e5rt_e5_compiler_options_set_segmenter":                                            true,
	"e5rt_e5_compiler_options_release":                                                  true,
	"e5rt_e5_compiler_compile":                                                          true,
	"e5rt_program_library_release":                                                      true,
	"e5rt_program_library_retain_program_function":                                      true,
	"e5rt_program_function_load_for_execution":                                          true,
	"e5rt_program_function_release":                                                     true,
	"e5rt_precompiled_compute_op_create_options_create_with_program_function":           true,
	"e5rt_precompiled_compute_op_create_options_set_operation_name":                     true,
	"e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers":      true,
	"e5rt_precompiled_compute_op_create_options_release":                                true,
	"e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options": true,
	"e5rt_execution_stream_operation_release":                                           true,
	"e5rt_buffer_object_alloc":                                                          true,
	"e5rt_buffer_object_get_data_ptr":                                                   true,
	"e5rt_buffer_object_release":                                                        true,
	"e5rt_io_port_bind_buffer_object":                                                   true,
	"e5rt_io_port_release":                                                              true,
	"e5rt_execution_stream_operation_retain_input_port":                                 true,
	"e5rt_execution_stream_operation_retain_output_port":                                true,
	"e5rt_execution_stream_create":                                                      true,
	"e5rt_execution_stream_encode_operation":                                            true,
	"e5rt_execution_stream_execute_sync":                                                true,
	"e5rt_execution_stream_reset":                                                       true,
	"e5rt_execution_stream_release":                                                     true,
}

func symbols(args []string) {
	fs := flag.NewFlagSet("symbols", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	rep := symbolReport{
		GOOS:      runtime.GOOS,
		Framework: e5rt.FrameworkPath,
		Listed:    len(e5rt.Symbols),
	}
	for _, name := range e5rt.Symbols {
		if wrapped[name] {
			rep.Wrapped = append(rep.Wrapped, name)
		} else {
			rep.Unwrapped = append(rep.Unwrapped, name)
		}
	}

	lib, err := e5rt.Open()
	if err != nil {
		rep.OpenError = err.Error()
	}
	got := make(map[string]bool)
	for _, name := range lib.Resolved() {
		got[name] = true
	}
	for _, name := range e5rt.Symbols {
		if got[name] {
			rep.Resolved = append(rep.Resolved, name)
		} else {
			rep.Unresolved = append(rep.Unresolved, name)
		}
	}

	stubPath, stubNames, stubErr := stubE5RTNames()
	rep.Stub = stubPath
	if stubErr != nil {
		rep.StubError = stubErr.Error()
	} else {
		rep.StubE5RTNames = len(stubNames)
		for _, name := range e5rt.Symbols {
			if !stubNames[name] {
				rep.NotInStub = append(rep.NotInStub, name)
			}
		}
	}

	if *jsonOut {
		writeJSON(rep)
		return
	}
	printSymbolReport(rep)
}

func printSymbolReport(rep symbolReport) {
	fmt.Printf("framework: %s\n", rep.Framework)
	fmt.Printf("goos:      %s\n\n", rep.GOOS)

	if rep.OpenError != "" {
		fmt.Printf("open: %s\n\n", rep.OpenError)
	}
	fmt.Printf("dlsym over the %d names in e5rt.Symbols:\n", rep.Listed)
	for _, name := range rep.Resolved {
		fmt.Printf("  %-9s %s\n", outcomeResolved, name)
	}
	for _, name := range rep.Unresolved {
		fmt.Printf("  %-9s %s\n", outcomeMissing, name)
	}
	fmt.Printf("\n%d RESOLVED, %d MISSING. Resolution says an address exists, not\n", len(rep.Resolved), len(rep.Unresolved))
	fmt.Printf("that the argument list the wrapper passes is right. What establishes\n")
	fmt.Printf("that is the route running: see espressoe5rtdispatch,\n")
	fmt.Printf("espressoe5rtblock, and espressoe5rtstack, not this command.\n\n")

	fmt.Printf("called through a typed wrapper by this example: %d of %d listed names.\n", len(rep.Wrapped), rep.Listed)
	fmt.Printf("listed but not called here at all. Some the package wraps and this\n")
	fmt.Printf("example does not exercise; the rest it leaves to Lib.Sym:\n")
	for _, name := range rep.Unwrapped {
		fmt.Printf("  %s\n", name)
	}
	fmt.Println()

	fmt.Printf("coverage against the SDK stub\n")
	fmt.Printf("  stub: %s\n", rep.Stub)
	if rep.StubError != "" {
		fmt.Printf("  denominator: UNAVAILABLE (%s)\n", rep.StubError)
		fmt.Printf("  No coverage fraction is reported, because none was measured.\n\n")
		return
	}
	fmt.Printf("  %d of %d distinct _e5rt_* spellings in the stub are listed here.\n",
		rep.Listed-len(rep.NotInStub), rep.StubE5RTNames)
	fmt.Printf("  CAVEAT: the stub's symbols: list mixes functions and data with no\n")
	fmt.Printf("  marker, so %d is an upper bound on e5rt_* functions, not a count of\n", rep.StubE5RTNames)
	fmt.Printf("  them, and the fraction above is not a function-coverage figure. The\n")
	fmt.Printf("  precedent: AppleNeuralEngine's 65 stub symbols are 6 functions and 59\n")
	fmt.Printf("  constants; counting the list as functions reports 52 phantom gaps.\n")
	fmt.Printf("  See appledocs internal/privateheaders/tbd.go.\n")
	if len(rep.NotInStub) > 0 {
		fmt.Printf("\n  listed here but absent from the stub (the binding names something\n")
		fmt.Printf("  the framework does not export):\n")
		for _, name := range rep.NotInStub {
			fmt.Printf("    %s\n", name)
		}
	}
	fmt.Println()
}

var e5rtName = regexp.MustCompile(`_e5rt_[A-Za-z0-9_]+`)

// stubE5RTNames returns the distinct _e5rt_* spellings in the SDK's Espresso
// stub, with the leading underscore stripped. It scans for the spelling
// rather than parsing the stub's sections: the names are wanted as a
// denominator, and the sections do not say which entries are functions.
func stubE5RTNames() (path string, names map[string]bool, err error) {
	out, err := exec.Command("xcrun", "--sdk", "macosx", "--show-sdk-path").Output()
	if err != nil {
		return "", nil, fmt.Errorf("xcrun --show-sdk-path: %w", err)
	}
	path = strings.TrimSpace(string(out)) +
		"/System/Library/PrivateFrameworks/Espresso.framework/Versions/A/Espresso.tbd"
	b, err := os.ReadFile(path)
	if err != nil {
		return path, nil, err
	}
	names = make(map[string]bool)
	for _, m := range e5rtName.FindAllString(string(b), -1) {
		names[strings.TrimPrefix(m, "_")] = true
	}
	return path, names, nil
}

// canary

func canary(args []string) {
	fs := flag.NewFlagSet("canary", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	results := []sequenceResult{
		runSequence("crash-canary", nil),
		runSequence("readback-canary", nil),
	}
	if *jsonOut {
		writeJSON(results)
		return
	}
	held := true
	for _, res := range results {
		printSequence(res)
		if res.Steps[len(res.Steps)-1].Outcome != outcomeCrashed {
			held = false
		}
		fmt.Println()
	}
	if held {
		fmt.Println("controls hold: the harness reports a child fault as CRASHED, and")
		fmt.Println("the readback probe faults on an address nothing backs. So in a")
		fmt.Println("stages report a CRASHED means a fault was seen, and a passing")
		fmt.Println("readback means the pointer was writable.")
		return
	}
	fmt.Println("CONTROL FAILED: a canary did not fault, or the parent did not see it.")
	fmt.Println("Until this passes, read the stages report as UNMEASURED rather than")
	fmt.Println("as evidence that nothing faulted.")
	os.Exit(1)
}

// stages

func stages(args []string) {
	fs := flag.NewFlagSet("stages", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	model := fs.String("model", "", "model path for the compile sequence")
	allow := fs.Bool("allow-unverified-calls", false, "acknowledge that these calls use guessed signatures")
	fs.Parse(args)

	if !*allow || os.Getenv(confirmEnv) != confirmValue {
		fmt.Fprintf(os.Stderr, `stages calls C functions whose argument lists are guesses from a paper.
A wrong guess corrupts memory or kills the process, and the fault is not
recoverable from Go. Both acknowledgements are required:

  %s=%s \
    go run ./examples/ane/e5rtprobe stages -allow-unverified-calls

`, confirmEnv, confirmValue)
		os.Exit(2)
	}

	var results []sequenceResult
	// The canary runs first so every stages report carries its own proof
	// that the harness can report a crash at all.
	results = append(results, runSequence("crash-canary", nil))
	results = append(results, runSequence("readback-canary", nil))
	var childArgs []string
	if *model != "" {
		childArgs = []string{"-model", *model}
	}
	for _, s := range sequences {
		if strings.HasSuffix(s.name, "-canary") {
			continue
		}
		if s.needsModel && *model == "" {
			results = append(results, skippedSequence(s, "no -model given"))
			continue
		}
		results = append(results, runSequence(s.name, childArgs))
	}

	if *jsonOut {
		writeJSON(struct {
			Sequences []sequenceResult `json:"sequences"`
			Blocked   []blockedStep    `json:"blocked"`
		}{results, blocked})
		return
	}
	for _, res := range results {
		printSequence(res)
		fmt.Println()
	}
	fmt.Println("route steps with no typed wrapper, so not attemptable from Go:")
	for _, b := range blocked {
		fmt.Printf("  %-9s %s\n            %s\n", outcomeUnreached, b.Symbol, b.Reason)
	}
	fmt.Println()
	fmt.Println("A RETURNED above reports only that control came back with that")
	fmt.Println("status. It is not evidence that the argument list was right, that")
	fmt.Println("a handle is valid, or that the Neural Engine did anything.")
}

// A blockedStep is a route step this tool cannot attempt because e5rt
// exposes no typed wrapper for it.
type blockedStep struct {
	Symbol string `json:"symbol"`
	Reason string `json:"reason"`
}

// Entries this list used to carry have left it one at a time. First
// program_library_create and the two operation-creation symbols, once ANEForge
// settled the argument order the paper left ambiguous; the route they blocked is
// now driven end to end by the e5rtdispatch example. Then submit_async, which
// was listed as unreachable on the grounds that the package had no way to build
// an Objective-C block. It did: objc.NewBlock builds a real __NSMallocBlock__,
// and Lib.SubmitAsync now uses it. The reason given was about this repository
// rather than about e5rt, and it was not true.
//
// Only one name is left, and it is missing an argument list rather than a
// calling convention.
var blocked = []blockedStep{
	{"e5rt_e5_compiler_is_new_compile_required",
		"named only in the paper's phase table; no call site, no argument list, and ANEForge does not use it either"},
}

// A sequence is a run of calls made in one child process. Steps are declared
// up front so the parent can name the ones a fault prevented.
type sequence struct {
	name       string
	steps      []string
	needsModel bool
	run        func(l *e5rt.Lib, model string, emit func(record))
}

var sequences = []sequence{
	{
		name:  "crash-canary",
		steps: []string{"before", "call-invalid-address"},
		run:   runCrashCanary,
	},
	{
		name:  "readback-canary",
		steps: []string{"before", "readback-invalid-address"},
		run:   runReadbackCanary,
	},
	{
		name:  "stream",
		steps: []string{"execution-stream-create", "execution-stream-reset"},
		run:   runStream,
	},
	{
		name:  "buffer",
		steps: []string{"buffer-object-alloc", "buffer-object-get-data-ptr", "data-ptr-readback"},
		run:   runBuffer,
	},
	{
		name:  "compiler",
		steps: []string{"compiler-create-with-config"},
		run:   runCompiler,
	},
	{
		name:  "compiler-config",
		steps: []string{"config-options-create", "compiler-create-with-config"},
		run:   runCompilerConfig,
	},
	{
		name:       "compile",
		steps:      []string{"compiler-create-with-config", "compiler-compile", "retain-program-function", "load-for-execution"},
		needsModel: true,
		run:        runCompile,
	},
}

func lookupSequence(name string) (sequence, bool) {
	for _, s := range sequences {
		if s.name == name {
			return s, true
		}
	}
	return sequence{}, false
}

func runCrashCanary(l *e5rt.Lib, _ string, emit func(record)) {
	emit(record{Step: "before", Outcome: outcomeReturned,
		Note: "the child is alive and its records reach the parent"})
	// Jump to an address that cannot hold code. This is the failure mode a
	// wrong e5rt_* signature produces: a fault in C that no recover sees.
	purego.SyscallN(1)
	emit(record{Step: "call-invalid-address", Outcome: outcomeReturned,
		Note: "did not fault; the control did not establish anything"})
}

// runReadbackCanary is the control for the buffer sequence's readback step: it
// runs that probe against an address nothing backs. Without it, a readback
// that passes is not evidence, because a probe that cannot fail proves
// nothing about the pointer it was handed.
func runReadbackCanary(l *e5rt.Lib, _ string, emit func(record)) {
	emit(record{Step: "before", Outcome: outcomeReturned,
		Note: "about to run the readback probe against an unbacked address"})
	r := readbackRecord(1, bufferBytes)
	r.Step = "readback-invalid-address"
	r.Note = "did not fault; the readback probe does not discriminate, so a " +
		"passing readback in the buffer sequence establishes nothing"
	emit(r)
}

func runStream(l *e5rt.Lib, _ string, emit func(record)) {
	stream, err := l.ExecutionStreamCreate()
	emit(callRecord("execution-stream-create", stream, err))
	if err != nil {
		return
	}
	err = l.ExecutionStreamReset(stream)
	emit(callRecord("execution-stream-reset", 0, err))
}

const bufferBytes = 4096

func runBuffer(l *e5rt.Lib, _ string, emit func(record)) {
	buf, err := l.BufferObjectAlloc(bufferBytes, 0)
	emit(callRecord("buffer-object-alloc", buf, err))
	if err != nil {
		return
	}
	ptr, err := l.BufferObjectGetDataPtr(buf)
	emit(callRecord("buffer-object-get-data-ptr", ptr, err))
	if err != nil || ptr == 0 {
		return
	}
	// Everything above is a returned value. This step is the one thing in
	// the sequence that can be observed: write a pattern through the
	// returned pointer and read it back. If the pointer is not backed by
	// writable memory the child faults here and the parent names this step.
	emit(readbackRecord(ptr, bufferBytes))
}

// pointerAt reinterprets a foreign address as a pointer. The address is
// returned by C and never refers to Go memory, so the unsafe.Pointer rule
// this looks like it breaks does not apply; the conversion goes through the
// address of a local so that go vet can see a real pointer to convert.
func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

// readbackRecord writes a pattern to the first and last byte of the region
// the returned pointer is claimed to cover and reads it back.
func readbackRecord(ptr uintptr, n uintptr) record {
	r := record{Step: "data-ptr-readback"}
	const pattern = 0xa5
	first := (*byte)(pointerAt(ptr))
	last := (*byte)(pointerAt(ptr + n - 1))
	*first, *last = pattern, pattern
	if *first != pattern || *last != pattern {
		r.Outcome = outcomeError
		r.Error = "wrote the pattern but did not read it back"
		return r
	}
	r.Outcome = outcomeReturned
	r.Note = fmt.Sprintf("wrote and read back %#x at offsets 0 and %d: the returned "+
		"pointer is backed by at least %d writable bytes. That is a property of "+
		"the address, not proof it is this buffer object's storage.", pattern, n-1, n)
	return r
}

func runCompiler(l *e5rt.Lib, _ string, emit func(record)) {
	compiler, err := l.CompilerCreateWithConfig(0)
	emit(callRecord("compiler-create-with-config", compiler, err))
}

// configOptionsCreate is reached through Lib.Sym with a convention supplied
// here — the escape hatch the package documents — falling back to a live dlsym
// if the name is not in e5rt.Symbols. The signature is
// e5rt_error_code_t(void **out), matching the out-first shape every e5rt_*
// constructor uses.
const configOptionsCreate = "e5rt_e5_compiler_config_options_create"

// newOut allocates an out-parameter cell on the heap. Passing the address of a
// stack local through a uintptr would let a stack growth during the call
// relocate it; returning the pointer from a non-inlinable function forces the
// escape. The caller must KeepAlive it across the call.
//
//go:noinline
func newOut() *uintptr { return new(uintptr) }

// dlsym resolves a name in Espresso directly. Dlopen is reference-counted and
// the framework is already loaded, so this adds no second copy.
func dlsym(name string) (uintptr, error) {
	h, err := purego.Dlopen(e5rt.FrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, err
	}
	return purego.Dlsym(h, name)
}

// runCompilerConfig tests whether the status the compiler constructor returns
// for a null config is explained by a config being required. It builds one
// first, then retries the constructor that failed without it.
func runCompilerConfig(l *e5rt.Lib, _ string, emit func(record)) {
	route := "Lib.Sym"
	sym, err := l.Sym(configOptionsCreate)
	if err != nil {
		// Lib.Sym is a lookup over what Open resolved, not a live dlsym, so
		// it reaches e5rt.Symbols and nothing else. This symbol has since
		// been added to that list, so the lookup now succeeds and this arm is
		// not normally taken; it stays as the fallback for a build of the
		// package that does not list the name, and the report says which
		// route was used either way rather than hiding the difference.
		sym, err = dlsym(configOptionsCreate)
		if err != nil {
			emit(record{Step: "config-options-create", Outcome: outcomeError,
				Error: err.Error()})
			return
		}
		route = "dlsym"
	}
	out := newOut()
	st, _, _ := purego.SyscallN(sym, uintptr(unsafe.Pointer(out)))
	config := *out
	runtime.KeepAlive(out)
	r := record{Step: "config-options-create", Outcome: outcomeReturned}
	status := int64(st)
	r.Status = &status
	if config != 0 {
		r.Handle = fmt.Sprintf("%#x", config)
	}
	r.Note = fmt.Sprintf("reached by %s: Lib.Sym resolves e5rt.Symbols and nothing "+
		"else, so a name outside that list falls back to a live dlsym", route)
	emit(r)
	if status != 0 || config == 0 {
		return
	}
	compiler, err := l.CompilerCreateWithConfig(config)
	rec := callRecord("compiler-create-with-config", compiler, err)
	if err == nil {
		rec.Note = "with a config built above, where a null config returned status 1"
	}
	emit(rec)
}

func runCompile(l *e5rt.Lib, model string, emit func(record)) {
	compiler, err := l.CompilerCreateWithConfig(0)
	emit(callRecord("compiler-create-with-config", compiler, err))
	if err != nil {
		return
	}
	library, err := l.CompilerCompile(compiler, model, 0)
	emit(callRecord("compiler-compile", library, err))
	if err != nil {
		return
	}
	fn, err := l.ProgramLibraryRetainProgramFunction(library, "main")
	emit(callRecord("retain-program-function", fn, err))
	if err != nil {
		return
	}
	err = l.ProgramFunctionLoadForExecution(fn)
	emit(callRecord("load-for-execution", 0, err))
}

// callRecord describes one returned call. A returned call is RETURNED, never
// OK: the wrapper reports status zero as success, and a guessed signature can
// return zero without having done anything the name implies.
func callRecord(step string, handle uintptr, err error) record {
	r := record{Step: step}
	if err != nil {
		r.Outcome = outcomeError
		r.Error = err.Error()
		return r
	}
	r.Outcome = outcomeReturned
	zero := int64(0)
	r.Status = &zero
	if handle != 0 {
		r.Handle = fmt.Sprintf("%#x", handle)
		r.Note = "handle value only; not shown to point at a valid object"
	}
	return r
}

// sequenceChild runs one sequence and streams its records to stdout. It is
// the crash domain: if a guessed signature faults, this process dies and the
// parent reports which step it died on.
func sequenceChild(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "e5rtprobe: __sequence needs a name")
		os.Exit(2)
	}
	name, args := args[0], args[1:]
	fs := flag.NewFlagSet("__sequence", flag.ExitOnError)
	model := fs.String("model", "", "model path")
	fs.Parse(args)

	s, ok := lookupSequence(name)
	if !ok {
		fmt.Fprintf(os.Stderr, "e5rtprobe: unknown sequence %q\n", name)
		os.Exit(2)
	}
	enc := json.NewEncoder(os.Stdout)
	emit := func(r record) {
		r.Sequence = name
		enc.Encode(r) // os.Stdout is unbuffered, so each record is out before the next call
	}

	lib, err := e5rt.Open()
	if err != nil && lib == nil {
		emit(record{Step: s.steps[0], Outcome: outcomeError, Error: err.Error()})
		return
	}
	s.run(lib, *model, emit)
}

// A sequenceResult is what the parent observed of one child.
type sequenceResult struct {
	Name        string   `json:"name"`
	Steps       []record `json:"steps"`
	Termination string   `json:"termination,omitempty"`
	Stderr      string   `json:"stderr,omitempty"`
	StderrLines int      `json:"stderr_lines,omitempty"` // total lines, before truncation
}

// stderrHeadLines is how much of a dying child's stderr the report keeps. A
// Go runtime traceback runs to hundreds of lines and buries the fault; the
// first few carry the signal, the faulting PC, and the address.
const stderrHeadLines = 6

// headLines returns the first n lines of s and the total line count, so a
// truncated report can say how much it dropped instead of implying that was
// all of it.
func headLines(s string, n int) (string, int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", 0
	}
	lines := strings.Split(s, "\n")
	if len(lines) <= n {
		return s, len(lines)
	}
	return strings.Join(lines[:n], "\n"), len(lines)
}

// faultSignal reports the signal named at the head of a Go runtime traceback,
// or the empty string if the child's stderr does not start with one.
func faultSignal(stderr string) string {
	line, _, _ := strings.Cut(stderr, "\n")
	if !strings.HasPrefix(line, "SIG") {
		return ""
	}
	return line
}

func skippedSequence(s sequence, reason string) sequenceResult {
	res := sequenceResult{Name: s.name, Termination: "not started: " + reason}
	for _, step := range s.steps {
		res.Steps = append(res.Steps, record{Sequence: s.name, Step: step,
			Outcome: outcomeUnreached, Note: reason})
	}
	return res
}

// runSequence runs one sequence in a child process and reconciles the records
// it streamed against the steps it declared. Steps with no record are
// UNREACHED, or CRASHED for the first one after the last record if the child
// died — that step is where the fault happened.
func runSequence(name string, extra []string) sequenceResult {
	s, ok := lookupSequence(name)
	if !ok {
		return sequenceResult{Name: name, Termination: "unknown sequence"}
	}
	res := sequenceResult{Name: name}

	args := append([]string{"__sequence", name}, extra...)
	cmd := exec.Command(os.Args[0], args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		res.Termination = err.Error()
		return res
	}
	var stderr strings.Builder
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		res.Termination = err.Error()
		return res
	}

	got := make(map[string]record)
	sc := bufio.NewScanner(stdout)
	for sc.Scan() {
		var r record
		if err := json.Unmarshal(sc.Bytes(), &r); err != nil {
			continue
		}
		got[r.Step] = r
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	var waitErr error
	select {
	case waitErr = <-done:
	case <-time.After(60 * time.Second):
		cmd.Process.Kill()
		<-done
		waitErr = errors.New("timed out after 60s")
	}
	if waitErr != nil {
		res.Termination = waitErr.Error()
	}
	res.Stderr, res.StderrLines = headLines(stderr.String(), stderrHeadLines)
	if sig := faultSignal(res.Stderr); sig != "" && res.Termination != "" {
		// A fault in C reaches the parent as an ordinary nonzero exit,
		// because the Go runtime catches the signal and prints the
		// traceback itself. Surface the signal so the report names the
		// fault rather than an exit code that could mean anything.
		res.Termination += " (" + sig + ")"
	}

	died := waitErr != nil
	crashAssigned := false
	for _, step := range s.steps {
		if r, ok := got[step]; ok {
			res.Steps = append(res.Steps, r)
			continue
		}
		r := record{Sequence: name, Step: step, Outcome: outcomeUnreached}
		if died && !crashAssigned {
			r.Outcome = outcomeCrashed
			r.Note = "child terminated here: " + res.Termination
			crashAssigned = true
		} else if died {
			r.Note = "not attempted; the child had already died"
		}
		res.Steps = append(res.Steps, r)
	}
	return res
}

func printSequence(res sequenceResult) {
	fmt.Printf("sequence %s\n", res.Name)
	for _, r := range res.Steps {
		fmt.Printf("  %-9s %s", r.Outcome, r.Step)
		if r.Status != nil {
			fmt.Printf(" status=%d", *r.Status)
		}
		if r.Handle != "" {
			fmt.Printf(" handle=%s", r.Handle)
		}
		fmt.Println()
		if r.Error != "" {
			fmt.Printf("            %s\n", r.Error)
		}
		if r.Note != "" {
			fmt.Printf("            %s\n", r.Note)
		}
	}
	if res.Termination != "" {
		fmt.Printf("  child: %s\n", res.Termination)
	}
	if res.Stderr != "" {
		shown := 0
		for line := range strings.SplitSeq(res.Stderr, "\n") {
			fmt.Printf("  stderr: %s\n", line)
			shown++
		}
		if res.StderrLines > shown {
			fmt.Printf("  stderr: ... %d further lines not shown\n", res.StderrLines-shown)
		}
	}
}

func writeJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "e5rtprobe: %v\n", err)
		os.Exit(1)
	}
}
