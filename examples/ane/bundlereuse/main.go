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
// all, between compiling and reopening. Anything that still runs afterwards
// cannot have recompiled, because there is nothing left to compile.
//
// The second process is a genuinely separate one, spawned after the source is
// gone. It shares this program's code-signing identity and has it as a parent.
// That is the case E5RT bundle reuse has been observed in; reuse after an aned
// restart, after a reboot, or from an unrelated process is not measured here
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
//   - A mutation control feeds the reopened program a different input and
//     requires the output to move. Without it, a bundle that returned a
//     constant, or replayed the previous run's output buffer, would pass every
//     comparison above.
//
//   - A wrong function name and a corrupted bundle are both required to be
//     refused. Without them, "the bundle opened" would be evidence only that
//     OpenBundle returns non-nil.
//
// Compile and open are timed separately, so the number that motivates the whole
// interface is visible rather than asserted.
package main

import (
	"bufio"
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
)

func main() {
	log.SetFlags(0)
	flag.Parse()
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

	if d := compare(compiled, want); d > *tol {
		return fmt.Errorf("the freshly compiled program disagrees with the CPU reference by %.3g", d)
	}
	fmt.Printf("  freshly compiled vs float64 CPU reference: max diff %.3g (tolerance %.3g)\n",
		compare(compiled, want), *tol)

	// Delete the source. Everything after this point cannot be a recompile,
	// because there is nothing left to compile from.
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
	if d := compare(reopened, compiled); d != 0 {
		return fmt.Errorf("the reopened program differs from the compiled one by %.3g; the same fp16 program on the same hardware should agree exactly", d)
	}
	fmt.Println("  reopened vs freshly compiled: identical")

	// Mutation control on the reopened program: a bundle that replayed a
	// buffer would pass every comparison above.
	perturbed := make([]float32, len(x))
	copy(perturbed, x)
	for i := range perturbed {
		perturbed[i] = -perturbed[i]
	}
	moved, _, err := openArm(bundlePath, perturbed)
	if err != nil {
		return fmt.Errorf("mutation control: %w", err)
	}
	movedBy := compare(moved, reopened)
	fmt.Printf("  mutation control: negating the input moved the output by %.3g\n", movedBy)
	if movedBy <= *tol {
		return fmt.Errorf("negating the input moved the output by only %.3g; the reopened bundle is not reading its input", movedBy)
	}

	// Arm 3: a second process, started after the source was deleted.
	crossed, err := childArm(bundlePath)
	if err != nil {
		return fmt.Errorf("second-process arm: %w", err)
	}
	if d := compare(crossed, compiled); d != 0 {
		return fmt.Errorf("the second process differs from the compiled arm by %.3g", d)
	}
	fmt.Println("  second process vs freshly compiled: identical")

	if err := refusalControls(dir, bundlePath); err != nil {
		return err
	}

	fmt.Println("\nOK")
	fmt.Println("\nMeasured here: reuse in a second process that shares this one's code-signing")
	fmt.Println("identity and has it as a parent. Reuse after an aned restart, after a reboot,")
	fmt.Println("or from an unrelated process is UNMEASURED.")
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
	return parseResult(string(out))
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
	fmt.Fprint(w, "RESULT")
	for _, v := range out {
		fmt.Fprintf(w, " %v", v)
	}
	fmt.Fprintln(w)
	return nil
}

func parseResult(s string) ([]float32, error) {
	for _, line := range strings.Split(s, "\n") {
		rest, ok := strings.CutPrefix(strings.TrimSpace(line), "RESULT ")
		if !ok {
			continue
		}
		fields := strings.Fields(rest)
		out := make([]float32, len(fields))
		for i, f := range fields {
			v, err := strconv.ParseFloat(f, 32)
			if err != nil {
				return nil, fmt.Errorf("parse %q: %w", f, err)
			}
			out[i] = float32(v)
		}
		return out, nil
	}
	return nil, fmt.Errorf("the second process printed no result")
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
	if err := truncateOneFile(corrupt); err != nil {
		return fmt.Errorf("corrupt bundle: %w", err)
	}
	p, err = e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: corrupt,
		Inputs:     []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err == nil {
		p.Close()
		return fmt.Errorf("refusal control: a bundle with a truncated file was accepted")
	}
	fmt.Printf("  refused, as it must be: a bundle with a truncated file\n    %v\n", err)
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

// truncateOneFile empties the largest regular file in the tree, which is the
// compiled program rather than a manifest.
func truncateOneFile(dir string) error {
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
		return err
	}
	if biggest == "" {
		return fmt.Errorf("no file found in %s", dir)
	}
	return os.Truncate(biggest, 0)
}

func findBundle(cacheDir string) (string, error) {
	var bundle string
	err := filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.HasSuffix(d.Name(), ".bundle") && bundle == "" {
			bundle = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if bundle == "" {
		return "", fmt.Errorf("no .bundle directory under %s", cacheDir)
	}
	return bundle, nil
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

func compare(got, want []float32) float64 {
	var max float64
	for i := range got {
		if i >= len(want) {
			break
		}
		if d := math.Abs(float64(got[i]) - float64(want[i])); d > max {
			max = d
		}
	}
	return max
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
