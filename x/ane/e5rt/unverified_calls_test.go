//go:build darwin

package e5rt_test

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

// The e5rt entry points below had never been called from Go. Their argument
// lists come from ANEForge, whose prose has been wrong where its call sites were
// right, so these tests replace citation with observation.
//
// Some of them abort the process. E5RT reports at least one failure by throwing
// a C++ exception rather than returning a status, and an exception thrown
// through a purego call cannot be caught in Go: the wrapper returns success and
// the process terminates later, in an unrelated frame. Each probe therefore runs
// in a subprocess, and its exit status is the result.

const probeEnv = "E5RT_PROBE_CASE"

// buildExampleModel writes a model directory holding a 1x1 identity convolution.
func buildExampleModel(t *testing.T, channels, spatial int) string {
	t.Helper()
	weights := make([]float32, channels*channels)
	for i := range channels {
		weights[i*channels+i] = 1
	}
	blob, err := mil.BuildWeightBlob(weights, channels, channels)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		t.Fatal(err)
	}
	text := mil.GenConvFP16IO(channels, channels, spatial)
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

// route holds the handles of a compiled, bound, encodable program.
type route struct {
	lib      *e5rt.Lib
	dir      string
	function uintptr
	op       uintptr
	stream   uintptr
	inPtr    uintptr
	outPtr   uintptr
}

// openRoute compiles the identity convolution and binds its ports. Handles are
// deliberately not released: every caller is a subprocess that is about to exit,
// and releasing a stream whose operation has been reset aborts.
func openRoute(t *testing.T) *route {
	t.Helper()
	lib, err := e5rt.Open()
	if err != nil {
		t.Skip(err)
	}
	const channels, spatial = 4, 4
	dir := buildExampleModel(t, channels, spatial)

	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		t.Fatal(err)
	}
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		t.Fatal(err)
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		t.Fatal(err)
	}
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		t.Fatal(err)
	}
	library, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options)
	if err != nil {
		t.Fatal(err)
	}
	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		t.Fatal(err)
	}
	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		t.Fatal(err)
	}
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		t.Fatal(err)
	}
	op, err := lib.OperationCreatePrecompiled(opOptions)
	if err != nil {
		t.Fatal(err)
	}
	_, inPtr := bindExamplePort(lib, op, "x", channels*spatial*2, true)
	_, outPtr := bindExamplePort(lib, op, "y", channels*spatial*2, false)
	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		t.Fatal(err)
	}
	return &route{lib: lib, dir: dir, function: function, op: op, stream: stream, inPtr: inPtr, outPtr: outPtr}
}

// dispatch runs the identity convolution and reports whether it round-tripped.
func (r *route) dispatch() error {
	in := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	writeExampleFP16(r.inPtr, in)
	if err := r.lib.ExecuteSync(r.stream); err != nil {
		return err
	}
	got := readExampleFP16(r.outPtr, len(in))
	for i := range in {
		if got[i] != in[i] {
			return fmt.Errorf("identity convolution returned %v, want %v", got, in)
		}
	}
	return nil
}

func findBundle(root string) string {
	var found string
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || found != "" {
			return nil
		}
		if d.IsDir() && strings.HasSuffix(d.Name(), ".bundle") {
			found = path
			return fs.SkipAll
		}
		return nil
	})
	return found
}

// probes are the individual experiments. Each runs in its own process and
// reports by returning an error or by aborting.
var probes = map[string]func(t *testing.T){
	// ANEForge resolves this symbol but never calls it, and the paper's dispatch
	// listings skip it.
	"loadForExecution": func(t *testing.T) {
		r := openRoute(t)
		err := r.lib.ProgramFunctionLoadForExecution(r.function)
		fmt.Printf("RESULT ProgramFunctionLoadForExecution: %v\n", err)
	},

	// ANEForge (ane_e5rt_dispatch.mm:460-464) says prepare_op_for_encode rejects
	// an operation that has never been encoded.
	"prepareBeforeEncode": func(t *testing.T) {
		r := openRoute(t)
		err := r.lib.PrepareOpForEncode(r.op)
		fmt.Printf("RESULT PrepareOpForEncode before any encode: %v\n", err)
	},

	// The same call once the operation has been encoded.
	"prepareAfterEncode": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatal(err)
		}
		err := r.lib.PrepareOpForEncode(r.op)
		fmt.Printf("RESULT PrepareOpForEncode after encode: %v\n", err)
	},

	// Whether an operation still dispatches after being reset and re-encoded on
	// the same stream.
	"prepareThenReencode": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatal(err)
		}
		if err := r.lib.PrepareOpForEncode(r.op); err != nil {
			t.Fatal(err)
		}
		err := r.lib.EncodeOperation(r.stream, r.op)
		fmt.Printf("RESULT re-encode on the same stream after prepare: %v\n", err)
		if err != nil {
			return
		}
		fmt.Printf("RESULT dispatch after prepare and re-encode: %v\n", r.dispatch())
	},

	// The same, with the stream reset between the prepare and the re-encode.
	"prepareResetThenReencode": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatal(err)
		}
		if err := r.lib.PrepareOpForEncode(r.op); err != nil {
			t.Fatal(err)
		}
		if err := r.lib.ExecutionStreamReset(r.stream); err != nil {
			fmt.Printf("RESULT stream reset after prepare: %v\n", err)
			return
		}
		err := r.lib.EncodeOperation(r.stream, r.op)
		fmt.Printf("RESULT re-encode after prepare and stream reset: %v\n", err)
		if err != nil {
			return
		}
		fmt.Printf("RESULT dispatch after prepare, reset and re-encode: %v\n", r.dispatch())
	},

	// Whether preparing an operation that was never encoded leaves the ordinary
	// route working.
	"prepareBeforeThenDispatch": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.PrepareOpForEncode(r.op); err != nil {
			t.Fatal(err)
		}
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatalf("encode after preparing a never-encoded op: %v", err)
		}
		fmt.Printf("RESULT dispatch after preparing a never-encoded op: %v\n", r.dispatch())
	},

	// Releasing a stream whose operation has been reset out of its encoded
	// state. This is what aborted the process when the calls above were first
	// run together in one program.
	"prepareAfterEncodeThenRelease": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatal(err)
		}
		if err := r.lib.PrepareOpForEncode(r.op); err != nil {
			t.Fatal(err)
		}
		fmt.Println("RESULT reached the stream release")
		err := r.lib.ExecutionStreamRelease(r.stream)
		fmt.Printf("RESULT stream release after prepare: %v\n", err)
	},

	// Asynchronous submission with a real Objective-C block. The callee retains
	// the block unconditionally, so a wrong second argument faults inside
	// objc_retain.
	"submitAsync": func(t *testing.T) {
		r := openRoute(t)
		if err := r.lib.EncodeOperation(r.stream, r.op); err != nil {
			t.Fatal(err)
		}
		in := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		writeExampleFP16(r.inPtr, in)

		done := make(chan struct{})
		release, err := r.lib.SubmitAsync(r.stream, func() e5rt.Status {
			close(done)
			return 0
		})
		if err != nil {
			t.Fatalf("SubmitAsync: %v", err)
		}
		select {
		case <-done:
			fmt.Println("RESULT SubmitAsync: the completion block ran")
		case <-time.After(10 * time.Second):
			t.Fatal("the completion block never ran")
		}
		release()

		got := readExampleFP16(r.outPtr, len(in))
		for i := range in {
			if got[i] != in[i] {
				t.Fatalf("async dispatch returned %v, want %v", got, in)
			}
		}
		fmt.Println("RESULT SubmitAsync: the output matches the reference")
	},

	// Opening a compiled bundle directly, without the compiler.
	"programLibraryCreate": func(t *testing.T) {
		r := openRoute(t)
		bundle := findBundle(r.dir)
		if bundle == "" {
			t.Fatal("no .bundle in the cache directory")
		}
		lib2, err := r.lib.ProgramLibraryCreate(bundle)
		if err != nil || lib2 == 0 {
			t.Fatalf("ProgramLibraryCreate: handle=%#x err=%v", lib2, err)
		}
		fn2, err := r.lib.ProgramLibraryRetainProgramFunction(lib2, "main")
		if err != nil || fn2 == 0 {
			t.Fatalf("RetainProgramFunction on the reopened library: handle=%#x err=%v", fn2, err)
		}
		fmt.Println("RESULT ProgramLibraryCreate: a compiled bundle reopens without the compiler")
	},
}

// TestE5RTProbe is the subprocess entry point. It does nothing unless the parent
// selected a case.
func TestE5RTProbe(t *testing.T) {
	name := os.Getenv(probeEnv)
	if name == "" {
		t.Skip("not a probe subprocess")
	}
	fn, ok := probes[name]
	if !ok {
		t.Fatalf("unknown probe %q", name)
	}
	fn(t)
}

// TestUnverifiedCalls runs each probe in its own process and reports what
// happened, including the cases that abort.
func TestUnverifiedCalls(t *testing.T) {
	if os.Getenv(probeEnv) != "" {
		t.Skip("probe subprocess")
	}
	if _, err := e5rt.Open(); err != nil {
		t.Skip(err)
	}

	names := make([]string, 0, len(probes))
	for name := range probes {
		names = append(names, name)
	}
	// Deterministic order.
	for i := range names {
		for j := i + 1; j < len(names); j++ {
			if names[j] < names[i] {
				names[i], names[j] = names[j], names[i]
			}
		}
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=TestE5RTProbe", "-test.v")
			cmd.Env = append(os.Environ(), probeEnv+"="+name)
			out, err := cmd.CombinedOutput()
			text := string(out)

			for line := range strings.SplitSeq(text, "\n") {
				if strings.HasPrefix(line, "RESULT ") {
					t.Log(line)
				}
			}
			if err == nil {
				return
			}
			t.Logf("subprocess exited with %v", err)
			for _, marker := range []string{"E5RTError", "no longer supported", "libc++abi"} {
				for line := range strings.SplitSeq(text, "\n") {
					if strings.Contains(line, marker) {
						t.Logf("C++ said: %s", strings.TrimSpace(line))
					}
				}
			}
		})
	}
}
