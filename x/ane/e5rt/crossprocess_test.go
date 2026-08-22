//go:build darwin

package e5rt_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

// ANEForge reports (docs/e5rt-dispatch-reference.md:305-309) that compile and
// dispatch must happen in the same process: the signed HWX lives only in aned's
// per-PID cache, so loading a previously compiled bundle in a fresh process
// fails with "Must re-compile the E5 bundle".
//
// This package's own note on [e5rt.Lib.ProgramLibraryCreate] said the opposite —
// that a compiled bundle is reusable across processes. That was not measured.
// The probe behind it compiled and reopened the bundle inside ONE process, which
// shows only that the call needs no compiler object, not that it needs no
// compiling process. A same-process test cannot speak to a per-PID cache.
//
// These tests settle it with two processes that never overlap: the first
// compiles and exits, the second opens the bundle it left behind.

const (
	crossEnvRole = "E5RT_CROSS_ROLE"
	crossEnvDir  = "E5RT_CROSS_DIR"
)

// TestCrossProcessChild is the subprocess entry point for the cross-process
// tests. It does nothing unless the parent assigned it a role.
func TestCrossProcessChild(t *testing.T) {
	role := os.Getenv(crossEnvRole)
	dir := os.Getenv(crossEnvDir)
	if role == "" {
		t.Skip("not a cross-process child")
	}
	lib, err := e5rt.Open()
	if err != nil {
		t.Skip(err)
	}
	switch role {
	case "compile":
		crossCompile(t, lib, dir)
	case "open":
		crossOpen(t, lib, dir)
	case "compileAndOpen":
		crossCompile(t, lib, dir)
		crossOpen(t, lib, dir)
	default:
		t.Fatalf("unknown role %q", role)
	}
}

// crossCompile writes the model into dir and compiles it there.
func crossCompile(t *testing.T, lib *e5rt.Lib, dir string) {
	t.Helper()
	const channels, spatial = 4, 4
	weights := make([]float32, channels*channels)
	for i := range channels {
		weights[i*channels+i] = 1
	}
	blob, err := mil.BuildWeightBlob(weights, channels, channels)
	if err != nil {
		t.Fatal(err)
	}
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
	if _, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options); err != nil {
		t.Fatal(err)
	}
	if findBundle(dir) == "" {
		t.Fatal("the compile left no .bundle in the cache directory")
	}
	fmt.Printf("RESULT compiled in pid %d\n", os.Getpid())
}

// crossOpen opens the bundle in dir and drives it as far as it will go, naming
// the step that fails. Each step is reported separately: a per-PID cache could
// stop this at the library, at the function, at the operation or at dispatch,
// and which one it is says where the signed program is checked.
func crossOpen(t *testing.T, lib *e5rt.Lib, dir string) {
	t.Helper()
	const channels, spatial = 4, 4
	bundle := findBundle(dir)
	if bundle == "" {
		t.Fatal("no .bundle to open")
	}
	fmt.Printf("RESULT opening in pid %d\n", os.Getpid())

	library, err := lib.ProgramLibraryCreate(bundle)
	fmt.Printf("RESULT ProgramLibraryCreate: handle=%#x err=%v\n", library, err)
	if err != nil || library == 0 {
		return
	}
	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	fmt.Printf("RESULT RetainProgramFunction: handle=%#x err=%v\n", function, err)
	if err != nil || function == 0 {
		return
	}
	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		fmt.Printf("RESULT PrecompiledComputeOpOptionsCreate: %v\n", err)
		return
	}
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		fmt.Printf("RESULT SetOperationName: %v\n", err)
		return
	}
	op, err := lib.OperationCreatePrecompiled(opOptions)
	fmt.Printf("RESULT OperationCreatePrecompiled: handle=%#x err=%v\n", op, err)
	if err != nil || op == 0 {
		return
	}
	_, inPtr := bindExamplePort(lib, op, "x", channels*spatial*2, true)
	_, outPtr := bindExamplePort(lib, op, "y", channels*spatial*2, false)
	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		fmt.Printf("RESULT ExecutionStreamCreate: %v\n", err)
		return
	}
	if err := lib.EncodeOperation(stream, op); err != nil {
		fmt.Printf("RESULT EncodeOperation: %v\n", err)
		return
	}
	in := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	writeExampleFP16(inPtr, in)
	if err := lib.ExecuteSync(stream); err != nil {
		fmt.Printf("RESULT ExecuteSync: %v\n", err)
		return
	}
	got := readExampleFP16(outPtr, len(in))
	for i := range in {
		if got[i] != in[i] {
			fmt.Printf("RESULT dispatched but the output is wrong: got %v want %v\n", got, in)
			return
		}
	}
	fmt.Printf("RESULT dispatched and the output matches the reference\n")
}

// bundleBackends reports the backends the compiler emitted for fn, read from
// the directories it leaves under the cache location: each function gets one
// <fn>_<backend> directory per selected backend. The device mask is a
// permission rather than a placement, so this is the only local evidence of
// where the work was put.
func bundleBackends(t *testing.T, cacheDir, fn string) []string {
	t.Helper()
	var seen []string
	filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() || filepath.Base(filepath.Dir(path)) != fn {
			return nil
		}
		if suffix, ok := strings.CutPrefix(d.Name(), fn+"_"); ok {
			seen = append(seen, suffix)
		}
		return nil
	})
	return seen
}

// runCrossChild runs one child with the given role and returns its output.
func runCrossChild(t *testing.T, role, dir string) (string, error) {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=TestCrossProcessChild", "-test.v")
	cmd.Env = append(os.Environ(), crossEnvRole+"="+role, crossEnvDir+"="+dir)
	out, err := cmd.CombinedOutput()
	for line := range strings.SplitSeq(string(out), "\n") {
		if strings.Contains(line, "RESULT ") || strings.Contains(line, "re-compile") ||
			strings.Contains(line, "E5RTError") || strings.Contains(line, "load has failed") {
			t.Log(strings.TrimSpace(line))
		}
	}
	return string(out), err
}

// TestCompiledBundleAcrossProcesses is the measurement. The compiling process
// exits before the opening process starts, so nothing about the first is live
// when the second runs.
func TestCompiledBundleAcrossProcesses(t *testing.T) {
	if os.Getenv(crossEnvRole) != "" {
		t.Skip("cross-process child")
	}
	if _, err := e5rt.Open(); err != nil {
		t.Skip(err)
	}
	// Not t.TempDir inside the child: both children must see the same directory,
	// and it must outlive the first.
	dir := t.TempDir()

	if _, err := runCrossChild(t, "compile", dir); err != nil {
		t.Fatalf("the compiling child failed: %v", err)
	}

	// The compiler must have chosen the Neural Engine, or a later dispatch says
	// nothing about a signed program on the engine: the identity convolution
	// returns the same correct answer from BNNS on the CPU.
	backends := bundleBackends(t, dir, "main")
	t.Logf("the compiler emitted %v for \"main\"", backends)
	if len(backends) != 1 || backends[0] != "ane" {
		t.Fatalf("compiled to %v, want exactly [ane]; the cross-process question is about the engine", backends)
	}

	// Remove the sources the compiler read. Whatever the second process does, it
	// cannot be recompiling this model — there is nothing left to compile from.
	// Without this a success is equally consistent with a silent recompile,
	// which would leave the per-PID claim untouched.
	if err := os.Remove(filepath.Join(dir, "model.mil")); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(filepath.Join(dir, "weights")); err != nil {
		t.Fatal(err)
	}

	out, err := runCrossChild(t, "open", dir)
	t.Logf("the opening child exited with %v", err)

	// The control. If the same sequence also fails inside one process, the
	// failure is not about process boundaries and this test has measured
	// nothing. A fresh directory keeps the two runs independent.
	same, sameErr := runCrossChild(t, "compileAndOpen", t.TempDir())
	t.Logf("the same-process control exited with %v", sameErr)

	sameWorks := strings.Contains(same, "dispatched and the output matches")
	crossWorks := strings.Contains(out, "dispatched and the output matches")
	if !sameWorks {
		t.Fatalf("the same-process control did not dispatch; nothing can be concluded about the cross-process case")
	}
	switch {
	case crossWorks:
		t.Log("VERDICT: a compiled bundle IS reusable in a fresh process")
	default:
		t.Log("VERDICT: a compiled bundle is NOT reusable in a fresh process; see the step that failed above")
	}
}
