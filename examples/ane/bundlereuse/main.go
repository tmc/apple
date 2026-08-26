// Command bundlereuse compiles a MIL program to a Neural Engine bundle once,
// then reopens that bundle — in this process and in a second process, with the
// source MIL and its weights deleted — and checks that every arm computes the
// same answer as a float64 CPU reference.
//
// Compilation is the expensive part of the private e5rt route: the aned daemon
// lowers MIL to a signed hardware program, which takes hundreds of milliseconds
// to seconds, while binding and dispatch are microseconds. A deployed program
// wants to pay that once. [e5rt.OpenBundle] is the interface for it, and until
// this example nothing outside the package's own tests called it.
//
// # What "reuse" has to mean to be worth anything
//
// A program that reopened a bundle and silently recompiled from the source
// would look exactly like one that reused it — same answer, same code path from
// the caller's side. So this example deletes the model directory, weights and
// all, between compiling and reopening. What that establishes is narrower than
// "no recompilation happened": it establishes that neither OpenBundle nor the
// second process fell back to the caller-side MIL and weights, because those no
// longer exist. What aned does internally when it materializes a compiled
// bundle is not observed here.
//
// The second process is a genuinely separate one, spawned after the source is
// gone. It shares this program's code-signing identity and has it as a parent.
// That is the case E5RT bundle reuse has been observed in; reuse after an aned
// reboot or from an unrelated process is not measured here
// and this program does not claim it.
//
// # What each check would catch
//
//   - Every arm is compared against a float64 CPU evaluation of the
//     convolution. An arm that loaded a different or stale program fails here.
//
//   - The reopened arms are compared against the freshly compiled arm as well,
//     which is a tighter bar than the reference: both run the same fp16 program
//     on the same hardware, so they should agree exactly, not merely closely.
//
//   - A mutation control runs two inputs through one opened program and
//     requires the second execution both to move and to match its own
//     recomputed reference. Two fresh opens would show only that two opens
//     react to two inputs; running twice through one program is what a stale
//     output or rebinding bug would fail, and it is what a decode loop needs.
//
//   - The second process declares how many values it returns, and the parent
//     re-checks that count. Six malformed child results are fed to the parser
//     on every run to prove it still refuses them. Without this, a truncated
//     result compared over its prefix was reported identical.
//
//   - An unrelated-process arm opens the bundle from a process that started
//     before the model existed, runs in its own session and process group,
//     shares no file descriptors with the compiler, and touches nothing until
//     after the source MIL and weights are gone. It is compared against the
//     freshly compiled arm exactly like the second-process arm. What remains
//     outside the claim: the opener still runs as the same user as whoever ran
//     this program, so reuse across a uid change stays unmeasured.
//
//   - A wrong function name and a damaged bundle file are both required to be
//     refused, and they fail at different native entry points. Without them,
//     "the bundle opened" would be evidence only that OpenBundle returns
//     non-nil.
//
// Compile and open are timed separately, so the number that motivates the whole
// interface is visible rather than asserted. Compile time varies substantially
// between runs, so the printed ratio is a diagnostic showing the order of
// magnitude, not a benchmark figure.
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

var (
	inCh    = flag.Int("inch", 32, "input channels")
	outCh   = flag.Int("outch", 32, "output channels")
	spatial = flag.Int("spatial", 8, "spatial extent")
	tol     = flag.Float64("tol", 5e-2, "maximum allowed difference from the CPU reference")

	child      = flag.Bool("child", false, "internal: run as the second process")
	childBundl = flag.String("child-bundle", "", "internal: bundle for the second process to open")
	// The unrelated child is spawned before any model exists and blocks on a
	// request file, so at open time it has inherited nothing from a compiler
	// that had not even started when the process came up.
	unrelatedChild = flag.Bool("unrelated-child", false, "internal: run as the unrelated opener process")
	unrelatedReq   = flag.String("unrelated-request", "", "internal: request file carrying the bundle path")
	unrelatedRes   = flag.String("unrelated-result", "", "internal: file the result line is written to")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if *unrelatedChild {
		if err := runUnrelatedChild(); err != nil {
			log.Fatal(err)
		}
		return
	}
	if *child {
		if err := runChild(); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if *inCh <= 0 || *outCh <= 0 || *spatial <= 0 {
		return fmt.Errorf("inch, outch and spatial must all be positive")
	}
	if !(*tol > 0) || math.IsInf(*tol, 0) {
		return fmt.Errorf("tol must be a positive finite number, have %v", *tol)
	}
	for _, p := range []struct {
		name string
		a, b int
	}{{"inch*spatial", *inCh, *spatial}, {"outch*spatial", *outCh, *spatial}, {"outch*inch", *outCh, *inCh}} {
		if p.a > math.MaxInt32/2/p.b {
			return fmt.Errorf("%s overflows a sensible buffer size", p.name)
		}
	}
	fmt.Printf("bundle reuse: inch=%d outch=%d spatial=%d\n", *inCh, *outCh, *spatial)

	dir, err := os.MkdirTemp("", "bundlereuse-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	modelDir := filepath.Join(dir, "model")
	cacheDir := filepath.Join(dir, "cache")
	weights := sampleWeights(*outCh * *inCh)
	if err := writeModel(modelDir, weights); err != nil {
		return fmt.Errorf("write model: %w", err)
	}

	x := sampleInput(*inCh * *spatial)
	want := cpuReference(x, weights)

	// Arm 1: compile from source, run, and note where the bundle landed.
	compiled, bundlePath, compileTime, err := compileArm(modelDir, cacheDir, x)
	if err != nil {
		return fmt.Errorf("compile arm: %w", err)
	}
	fmt.Printf("\n  compiled in %v\n", compileTime.Round(time.Millisecond))
	reportBackend(cacheDir)

	baseDiff, err := compare(compiled, want)
	if err != nil {
		return fmt.Errorf("freshly compiled vs CPU reference: %w", err)
	}
	if baseDiff > *tol {
		return fmt.Errorf("the freshly compiled program disagrees with the CPU reference by %.3g", baseDiff)
	}
	fmt.Printf("  freshly compiled vs float64 CPU reference: max diff %.3g (tolerance %.3g)\n", baseDiff, *tol)

	// Delete the source. Nothing after this point can fall back to the
	// caller-side MIL and weights, because they no longer exist. That is the
	// claim; what aned does internally is not observed here.
	if err := os.RemoveAll(modelDir); err != nil {
		return fmt.Errorf("remove model source: %w", err)
	}
	if _, err := os.Stat(filepath.Join(modelDir, "model.mil")); !os.IsNotExist(err) {
		return fmt.Errorf("model source still present after removal; the reuse arms would be unattributable")
	}
	fmt.Println("\n  source MIL and weights deleted")

	// Arm 2: reopen in this process.
	reopened, openTime, err := openArm(bundlePath, x)
	if err != nil {
		return fmt.Errorf("reopen arm: %w", err)
	}
	fmt.Printf("  reopened in %v, %.0fx faster than compiling\n",
		openTime.Round(time.Microsecond), float64(compileTime)/float64(openTime))
	reopenDiff, err := compare(reopened, compiled)
	if err != nil {
		return fmt.Errorf("reopened vs freshly compiled: %w", err)
	}
	// Exact equality is required because both arms are the same fp16 program on
	// the same binary and host. That is an observation about this run, not a
	// portability contract: a compiler or runtime revision could legitimately
	// change rounding, and the float64 reference above is the arithmetic oracle.
	if reopenDiff != 0 {
		return fmt.Errorf("the reopened program differs from the compiled one by %.3g", reopenDiff)
	}
	fmt.Println("  reopened vs freshly compiled: identical")

	// Mutation control, both inputs through ONE opened program: a bundle that
	// replayed a buffer would pass every comparison above.
	perturbed := make([]float32, len(x))
	for i := range x {
		perturbed[i] = -x[i]
	}
	firstOut, secondOut, err := reopenTwice(bundlePath, x, perturbed)
	if err != nil {
		return fmt.Errorf("mutation control: %w", err)
	}
	// The second execution must be right, not merely different: negating the
	// input negates the output of a linear map.
	wantPerturbed := cpuReference(perturbed, weights)
	if d, err := compare(secondOut, wantPerturbed); err != nil || d > *tol {
		return fmt.Errorf("the second execution of one reopened program disagrees with its own reference (diff %.3g, err %v)", d, err)
	}
	movedBy, err := compare(secondOut, firstOut)
	if err != nil {
		return fmt.Errorf("mutation control: %w", err)
	}
	fmt.Printf("  mutation control: one opened program, second execution with a negated input moved the output by %.3g and still matches its own reference\n", movedBy)
	if movedBy <= *tol {
		return fmt.Errorf("the second execution moved the output by only %.3g; the reopened program is not reading its input on re-execution", movedBy)
	}

	// Arm 3: a second process, started after the source was deleted.
	crossed, err := childArm(bundlePath)
	if err != nil {
		return fmt.Errorf("second-process arm: %w", err)
	}
	crossDiff, err := compare(crossed, compiled)
	if err != nil {
		return fmt.Errorf("second process vs freshly compiled: %w", err)
	}
	if crossDiff != 0 {
		return fmt.Errorf("the second process differs from the compiled arm by %.3g", crossDiff)
	}
	fmt.Printf("  second process vs freshly compiled: identical over all %d values\n", len(crossed))

	// Arm 4: an unrelated process — spawned before anything was compiled,
	// in its own session, waiting on a request file that names a bundle which
	// does not exist yet.
	unrel, err := unrelatedArm(bundlePath)
	if err != nil {
		return fmt.Errorf("unrelated-process arm: %w", err)
	}
	unrelDiff, err := compare(unrel, compiled)
	if err != nil {
		return fmt.Errorf("unrelated process vs freshly compiled: %w", err)
	}
	if unrelDiff != 0 {
		return fmt.Errorf("the unrelated process differs from the compiled arm by %.3g", unrelDiff)
	}
	fmt.Printf("  unrelated process vs freshly compiled: identical over all %d values\n", len(unrel))

	if err := parserControls(); err != nil {
		return err
	}

	if err := refusalControls(dir, bundlePath); err != nil {
		return err
	}

	fmt.Println("\nOK")
	fmt.Println("\nMeasured here: reuse in a second process that shares this one's code-signing")
	fmt.Println("identity and has it as a parent; reuse from an unrelated process spawned before")
	fmt.Println("the model existed, in its own session, opening only after the source was gone")
	fmt.Println("(same user — cross-uid reuse stays unmeasured); and reuse across an aned restart,")
	fmt.Println("measured separately in examples/ane/internal/anedrestart, which survives. Reuse")
	fmt.Println("after a reboot is UNMEASURED. Bundles are not durable: two roughly day-old bundles")
	fmt.Println("stopped opening; the expiry boundary is UNMEASURED.")
	return nil
}

// compileArm compiles the model, runs it, and reports where the compiler left
// its bundle.
func compileArm(modelDir, cacheDir string, x []float32) ([]float32, string, time.Duration, error) {
	start := time.Now()
	p, err := e5rt.Compile(e5rt.ProgramOptions{
		ModelPath: filepath.Join(modelDir, "model.mil"),
		CacheDir:  cacheDir,
		Inputs:    []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:   []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err != nil {
		return nil, "", 0, err
	}
	elapsed := time.Since(start)
	defer p.Close()

	out, err := evalProgram(p, x)
	if err != nil {
		return nil, "", 0, err
	}
	bundle, err := findBundle(cacheDir)
	if err != nil {
		return nil, "", 0, err
	}
	return out, bundle, elapsed, nil
}

// reopenTwice opens the bundle once and runs two different inputs through that
// single program, returning both results.
//
// Running each input through its own fresh open would show only that two opens
// react to two inputs. It would not show that a reopened program reacts on its
// second execution, which is the case a stale-output or rebinding bug actually
// lives in — and it is the case a decode loop depends on.
func reopenTwice(bundlePath string, first, second []float32) ([]float32, []float32, error) {
	p, err := e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: bundlePath,
		Inputs:     []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err != nil {
		return nil, nil, err
	}
	defer p.Close()

	a, err := evalProgram(p, first)
	if err != nil {
		return nil, nil, err
	}
	b, err := evalProgram(p, second)
	if err != nil {
		return nil, nil, err
	}
	return a, b, nil
}

// openArm opens an existing bundle and runs it.
func openArm(bundlePath string, x []float32) ([]float32, time.Duration, error) {
	start := time.Now()
	p, err := e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: bundlePath,
		Inputs:     []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err != nil {
		return nil, 0, err
	}
	elapsed := time.Since(start)
	defer p.Close()

	out, err := evalProgram(p, x)
	if err != nil {
		return nil, 0, err
	}
	return out, elapsed, nil
}

// childArm re-executes this program with -child, so the bundle is opened by a
// process that did not compile it and that started after the source was gone.
func childArm(bundlePath string) ([]float32, error) {
	self, err := os.Executable()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(self,
		"-child",
		"-child-bundle", bundlePath,
		"-inch", strconv.Itoa(*inCh),
		"-outch", strconv.Itoa(*outCh),
		"-spatial", strconv.Itoa(*spatial),
	)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("second process: %w", err)
	}
	return parseResult(string(out), *outCh**spatial)
}

// unrelatedArm opens the bundle from a process that cannot have inherited any
// compiler state: it is spawned BEFORE this function runs — before the model
// directory even exists — it starts its own session and process group via
// setsid, exec.Command passes it no extra file descriptors, and it blocks on a
// request file until after the source MIL and weights are deleted.
//
// The claim this arm establishes is deliberately narrower than the word
// "unrelated" suggests. The opener still runs as the same user; reuse across a
// uid change or a different code-signing identity remains unmeasured here.
func unrelatedArm(bundlePath string) ([]float32, error) {
	self, err := os.Executable()
	if err != nil {
		return nil, err
	}
	dir, err := os.MkdirTemp("", "bundlereuse-unrelated-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	req := filepath.Join(dir, "request")
	res := filepath.Join(dir, "result")

	cmd := exec.Command(self,
		"-unrelated-child",
		"-unrelated-request", req,
		"-unrelated-result", res,
		"-inch", strconv.Itoa(*inCh),
		"-outch", strconv.Itoa(*outCh),
		"-spatial", strconv.Itoa(*spatial),
	)
	// A new session detaches the opener from this one's process group and
	// controlling terminal. It does not change the uid, which is exactly why
	// the doc comment bounds the claim the way it does.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("spawn unrelated opener: %w", err)
	}
	defer func() {
		// If anything above failed before the result appeared, stop the waiter.
		os.Remove(req)
		_, statErr := os.Stat(res)
		if statErr != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	}()

	// The request file is written only after compileArm has run and the source
	// has been deleted, so from the opener's point of view the bundle simply
	// appears, already final, with nothing about its provenance observable.
	if err := os.WriteFile(req, []byte(bundlePath+"\n"), 0o644); err != nil {
		return nil, fmt.Errorf("write request file: %w", err)
	}
	deadline := time.Now().Add(5 * time.Minute)
	for {
		data, err := os.ReadFile(res)
		if err == nil {
			return parseResult(string(data), *outCh**spatial)
		}
		if time.Now().After(deadline) {
			return nil, errors.New("the unrelated opener produced no result within 5 minutes")
		}
		time.Sleep(20 * time.Millisecond)
	}
}

// runUnrelatedChild waits for a request naming a bundle that did not exist
// when this process started, then opens it and writes the result line.
func runUnrelatedChild() error {
	if *unrelatedReq == "" || *unrelatedRes == "" {
		return fmt.Errorf("unrelated-child mode needs -unrelated-request and -unrelated-result")
	}
	deadline := time.Now().Add(10 * time.Minute)
	var data []byte
	for {
		d, err := os.ReadFile(*unrelatedReq)
		if err == nil && len(d) > 0 {
			data = d
			break
		}
		if !os.IsNotExist(err) && err != nil {
			return fmt.Errorf("read request: %w", err)
		}
		if time.Now().After(deadline) {
			return errors.New("no request arrived within 10 minutes")
		}
		time.Sleep(20 * time.Millisecond)
	}
	bundlePath := strings.TrimSpace(string(data))
	if bundlePath == "" {
		return errors.New("empty request")
	}
	x := sampleInput(*inCh * *spatial)
	out, _, err := openArm(bundlePath, x)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var b strings.Builder
	fmt.Fprintf(&b, "RESULT %d", len(out))
	for _, v := range out {
		fmt.Fprintf(&b, " %v", v)
	}
	fmt.Fprintln(w, b.String())
	return os.WriteFile(*unrelatedRes, []byte(b.String()+"\n"), 0o644)
}

// parserControls prove the result parser refuses the shapes that used to slip
// through it. Without these, the parser's strictness is asserted rather than
// demonstrated, and a later edit could quietly restore the false accept.
func parserControls() error {
	n := *outCh * *spatial
	full := "RESULT " + strconv.Itoa(n) + strings.Repeat(" 1", n)
	if _, err := parseResult(full, n); err != nil {
		return fmt.Errorf("parser control: a well-formed result was refused: %w", err)
	}
	bad := []struct{ name, text string }{
		{"no result line", "nothing here\n"},
		{"empty result line", "RESULT \n"},
		{"zero values declared and none printed", "RESULT 0\n"},
		{"one value where the program returns many", "RESULT 1 1\n"},
		{"a declared count that does not match", "RESULT 99 1 2 3\n"},
		{"two result lines", full + "\n" + full + "\n"},
	}
	for _, c := range bad {
		if _, err := parseResult(c.text, n); err == nil {
			return fmt.Errorf("parser control: %s was accepted", c.name)
		}
	}
	fmt.Printf("  parser controls: a well-formed result parses, and %d malformed ones are refused\n", len(bad))
	return nil
}

// runChild is the second process: it opens the bundle it is given, runs the
// same input, and prints the result for the parent to compare.
func runChild() error {
	if *childBundl == "" {
		return fmt.Errorf("no bundle path")
	}
	x := sampleInput(*inCh * *spatial)
	out, _, err := openArm(*childBundl, x)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	// The count leads the line so the parent can tell a truncated result from
	// a short one, rather than comparing whatever prefix arrived.
	fmt.Fprintf(w, "RESULT %d", len(out))
	for _, v := range out {
		fmt.Fprintf(w, " %v", v)
	}
	fmt.Fprintln(w)
	return nil
}

// parseResult reads the second process's output, requiring exactly one result
// line carrying exactly the expected number of values.
//
// The count is declared by the child and re-checked here against what this
// process expects. Without that, a child that printed a truncated line — or a
// bare "RESULT" with no values at all — parsed cleanly, and the comparison then
// ran over the short prefix and called it identical. That was a false accept on
// the cross-process claim, which is the whole point of the arm.
func parseResult(s string, wantCount int) ([]float32, error) {
	var found []string
	for _, line := range strings.Split(s, "\n") {
		if rest, ok := strings.CutPrefix(strings.TrimSpace(line), "RESULT "); ok {
			found = append(found, rest)
		}
	}
	if len(found) == 0 {
		return nil, errors.New("the second process printed no result line")
	}
	if len(found) > 1 {
		return nil, fmt.Errorf("the second process printed %d result lines, expected exactly one", len(found))
	}
	fields := strings.Fields(found[0])
	if len(fields) < 1 {
		return nil, errors.New("the second process printed an empty result line")
	}
	declared, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("the second process did not declare a value count: %w", err)
	}
	values := fields[1:]
	if declared != len(values) {
		return nil, fmt.Errorf("the second process declared %d values and printed %d", declared, len(values))
	}
	if declared != wantCount {
		return nil, fmt.Errorf("the second process returned %d values, expected %d", declared, wantCount)
	}
	out := make([]float32, len(values))
	for i, f := range values {
		v, err := strconv.ParseFloat(f, 32)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", f, err)
		}
		out[i] = float32(v)
	}
	return out, nil
}

// refusalControls require a wrong function name and a corrupted bundle to be
// rejected. Without them, a successful open would be evidence only that
// OpenBundle can return a non-nil program.
func refusalControls(dir, bundlePath string) error {
	fmt.Println()

	p, err := e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath:   bundlePath,
		FunctionName: "no_such_function",
		Inputs:       []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:      []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err == nil {
		p.Close()
		return fmt.Errorf("refusal control: a bundle opened under a function name it does not contain")
	}
	fmt.Printf("  refused, as it must be: a function name the bundle does not contain\n    %v\n", err)

	corrupt := filepath.Join(dir, "corrupt.bundle")
	if err := copyTree(bundlePath, corrupt); err != nil {
		return fmt.Errorf("copy bundle: %w", err)
	}
	damaged, wasSize, err := truncateOneFile(corrupt)
	if err != nil {
		return fmt.Errorf("corrupt bundle: %w", err)
	}
	p, err = e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: corrupt,
		Inputs:     []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err == nil {
		p.Close()
		return fmt.Errorf("refusal control: a bundle with %s truncated from %d bytes was accepted", damaged, wasSize)
	}
	fmt.Printf("  refused, as it must be: a bundle with %s truncated from %d bytes to 0\n    %v\n", damaged, wasSize, err)
	return nil
}

func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(target)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, in)
		return err
	})
}

// truncateOneFile empties the largest regular file in the tree and reports
// which file it was and how large.
//
// It reports rather than assumes: an earlier comment here called the largest
// file "the compiled program rather than a manifest", which is an inference
// about bundle layout that can change with format or model size. What the
// control establishes is that damaging a selected retained bundle file causes a
// refusal — not that every file in the bundle is validated.
func truncateOneFile(dir string) (string, int64, error) {
	var biggest string
	var size int64
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.Size() > size {
			size, biggest = info.Size(), path
		}
		return nil
	})
	if err != nil {
		return "", 0, err
	}
	if biggest == "" {
		return "", 0, fmt.Errorf("no file found in %s", dir)
	}
	if err := os.Truncate(biggest, 0); err != nil {
		return "", 0, err
	}
	rel, relErr := filepath.Rel(dir, biggest)
	if relErr != nil {
		rel = biggest
	}
	return rel, size, nil
}

// findBundle returns the outermost compiled bundle, failing closed unless there
// is exactly one.
//
// The compiler nests: it emits <digest>.bundle containing a per-hardware
// <generation>.bundle, so a naive search finds two. Taking the first one
// encountered gets the right answer only because WalkDir happens to visit a
// parent before its children — an ordering accident, not a selection rule.
// Selecting the bundle that has no .bundle ancestor states the intent, and
// requiring exactly one such bundle means a cache holding two compiled models
// fails rather than silently picking one.
func findBundle(cacheDir string) (string, error) {
	var top []string
	err := filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() || !strings.HasSuffix(d.Name(), ".bundle") {
			return nil
		}
		// Nested bundles belong to the enclosing one; skip them.
		if strings.Contains(filepath.Dir(path), ".bundle") {
			return nil
		}
		top = append(top, path)
		return nil
	})
	if err != nil {
		return "", err
	}
	switch len(top) {
	case 0:
		return "", fmt.Errorf("no top-level .bundle directory under %s", cacheDir)
	case 1:
		return top[0], nil
	default:
		return "", fmt.Errorf("expected exactly one top-level bundle under %s, found %d: %v", cacheDir, len(top), top)
	}
}

func evalProgram(p *e5rt.Program, x []float32) ([]float32, error) {
	in, err := p.Input("x")
	if err != nil {
		return nil, err
	}
	out, err := p.Output("y")
	if err != nil {
		return nil, err
	}
	if err := in.WriteFP16(x); err != nil {
		return nil, err
	}
	if err := p.Execute(); err != nil {
		return nil, err
	}
	values := make([]float32, *outCh**spatial)
	if err := out.ReadFP16(values); err != nil {
		return nil, err
	}
	return values, nil
}

func writeModel(dir string, weights []float32) error {
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return err
	}
	text := mil.GenConvFP16(*inCh, *outCh, *spatial)
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(text), 0o644); err != nil {
		return err
	}
	blob, err := mil.BuildWeightBlob(weights, *outCh, *inCh)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644)
}

// cpuReference evaluates the 1x1 convolution in float64. Tensors are
// [1, C, 1, S] row-major and the weight is [outCh, inCh, 1, 1], so output
// (o, s) is the dot product of weight row o with input column s.
func cpuReference(x, weights []float32) []float32 {
	out := make([]float32, *outCh**spatial)
	for o := 0; o < *outCh; o++ {
		for s := 0; s < *spatial; s++ {
			var acc float64
			for i := 0; i < *inCh; i++ {
				acc += float64(weights[o**inCh+i]) * float64(x[i**spatial+s])
			}
			out[o**spatial+s] = float32(acc)
		}
	}
	return out
}

// compare returns the largest absolute difference between got and want, and
// refuses the two ways this check can pass while measuring nothing.
//
// The earlier version iterated over got and stopped at want's length, so a
// second process that printed one of its values compared one element and was
// reported identical. A NaN is the other hole: it compares false against every
// threshold, so `d > max` never fires and an all-NaN result reports a maximum
// difference of zero. An infinity at least survives and fails loudly; NaN is the
// silent one, and NaN is what an fp16 overflow produces.
func compare(got, want []float32) (float64, error) {
	if len(got) != len(want) {
		return 0, fmt.Errorf("comparing %d values against %d", len(got), len(want))
	}
	if len(got) == 0 {
		return 0, errors.New("nothing to compare")
	}
	var max float64
	for i := range got {
		g, w := float64(got[i]), float64(want[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			return 0, fmt.Errorf("result element %d is %v", i, g)
		}
		if math.IsNaN(w) || math.IsInf(w, 0) {
			return 0, fmt.Errorf("reference element %d is %v", i, w)
		}
		if d := math.Abs(g - w); d > max {
			max = d
		}
	}
	return max, nil
}

// sampleInput and sampleWeights are deterministic, so the parent and the second
// process feed the engine exactly the same numbers and the CPU reference stays
// comparable across runs.
func sampleInput(n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(math.Sin(float64(i)*0.61 + 0.2))
	}
	return out
}

func sampleWeights(n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(math.Cos(float64(i)*0.29)) * 0.2
	}
	return out
}

// reportBackend prints which backend the compiler chose, read from the
// main_<backend> directory it leaves in the cache location. The device mask is
// a permission and not a placement, so this is the only local evidence that the
// work was put on the engine.
func reportBackend(cacheDir string) {
	seen := map[string]bool{}
	filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() || filepath.Base(filepath.Dir(path)) != "main" {
			return nil
		}
		if suffix, ok := strings.CutPrefix(d.Name(), "main_"); ok {
			seen[suffix] = true
		}
		return nil
	})
	if len(seen) == 0 {
		fmt.Println("  backend: UNKNOWN (no main_* directory in the compiled bundle)")
		return
	}
	names := make([]string, 0, len(seen))
	for n := range seen {
		names = append(names, n)
	}
	fmt.Printf("  backend: the compiler emitted %v\n", names)
}
