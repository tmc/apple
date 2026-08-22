// Command surfaceffn runs an Apple Neural Engine matmul over IOSurface
// pages that another Go process wrote — the ANE arm of the surfacecompute
// family (examples/iosurface/tensorshare, examples/metal/surfacecompute).
//
// The direction of the handoff is inverted from surfacecompute, because the
// ANE compiler owns its tensors: compiling a model allocates IOSurfaces with
// the exact strided layout the NPU DMAs from. So the consumer compiles the
// kernel (x/ane/mil dynamic matmul: y = xᵀ·W in fp16, fp32 at the edges)
// and passes its input and output surfaces' mach ports BACK to the producer
// (IOSurfaceCreateMachPort → x/mach Send with MoveSend →
// IOSurfaceLookupFromMachPort). The producer writes activations and weights
// straight into the NPU's input pages and reads results from its output
// pages; the consumer only pulls the trigger. Neither process copies a
// tensor, and the consumer never sees one element of the data it computes on.
//
// Two proofs, mirroring surfacecompute:
//
//  1. The producer's float64 CPU reference matches what it reads out of the
//     ANE's output surface after the consumer evaluates (fp16-scale
//     tolerance — the ANE accumulates in reduced precision and the demo
//     says so rather than pretending fp32 accuracy).
//  2. The producer mutates one activation element in place; the consumer
//     re-evaluates without any new writes on its side, and the producer
//     sees the recomputed result. A kernel bound to a snapshot copy would
//     return the stale product.
//
// It then reports per-eval wall time from the producer's side and the
// hardware-reported execution time from the ANE's own telemetry.
//
// The ANE compile path is private-framework territory (_ANEClient) and is
// presence-checked: where unavailable the demo reports itself unavailable
// and exits cleanly.
//
//	go run ./examples/ane/surfaceffn
//	go run ./examples/ane/surfaceffn -in 1024 -out 1024 -batch 128
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/model"
	"github.com/tmc/apple/x/mach"
)

func main() {
	log.SetFlags(0)
	inDim := flag.Int("in", 512, "input dimension (matmul K)")
	outDim := flag.Int("out", 512, "output dimension (matmul N)")
	batch := flag.Int("batch", 128, "batch size (matmul M)")
	seed := flag.Int64("seed", 1, "RNG seed for tensor contents")
	consume := flag.Bool("consume", false, "run as consumer child (internal)")
	service := flag.String("service", "", "bootstrap name for the port handoff (consumer)")
	flag.Parse()

	if *consume {
		runConsumer(*service, *inDim, *outDim, *batch)
		return
	}
	runProducer(*inDim, *outDim, *batch, *seed)
}

// runConsumer compiles the ANE kernel, hands its input and output surfaces
// to the producer, and evaluates on command. It never reads or writes a
// tensor element: the data plane belongs entirely to the other process.
func runConsumer(service string, inDim, outDim, batch int) {
	k, err := model.Compile(model.CompileOptions{
		MILText:     mil.GenDynamicMatmul(inDim, outDim, batch),
		SharedModel: true,
	})
	if err != nil {
		// Presence check: private frameworks missing or no ANE. The
		// producer treats this as a clean unavailable, not a failure.
		fmt.Printf("unavail %v\n", err)
		return
	}
	defer k.Close()

	for _, l := range []struct {
		name   string
		layout ane.TensorLayout
	}{{"in", k.InputLayout(0)}, {"out", k.OutputLayout(0)}} {
		fmt.Printf("layout %s %d %d %d %d %d %d\n", l.name,
			l.layout.Channels, l.layout.Width, l.layout.Height,
			l.layout.ElemSize, l.layout.RowStride, l.layout.PlaneStride)
	}

	// Hand both surfaces to the producer in one message.
	inPort := mach.Port(iosurface.IOSurfaceCreateMachPort(iosurface.IOSurfaceRef(k.InputSurface(0))))
	outPort := mach.Port(iosurface.IOSurfaceCreateMachPort(iosurface.IOSurfaceRef(k.OutputSurface(0))))
	if inPort == mach.PortNull || outPort == mach.PortNull {
		log.Fatal("surfaceffn: IOSurfaceCreateMachPort failed")
	}
	var svc mach.Port
	deadline := time.Now().Add(5 * time.Second)
	for {
		svc, err = mach.BootstrapLookUp(service)
		if err == nil {
			break
		}
		if time.Now().After(deadline) {
			log.Fatalf("surfaceffn: bootstrap rendezvous: %v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	rights := []mach.PortRight{
		{Port: inPort, Disposition: mach.MoveSend},
		{Port: outPort, Disposition: mach.MoveSend},
	}
	if err := mach.Send(svc, mach.CopySend, 1, rights, nil, 5*time.Second); err != nil {
		log.Fatalf("surfaceffn: send surface ports: %v", err)
	}
	svc.Deallocate()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		switch in.Text() {
		case "eval":
			hwNS, err := k.EvalHWExecutionNS()
			if err != nil {
				log.Fatalf("surfaceffn: eval: %v", err)
			}
			fmt.Printf("done %d\n", hwNS)
		case "quit":
			return
		default:
			log.Fatalf("surfaceffn: unknown command %q", in.Text())
		}
	}
}

func runProducer(inDim, outDim, batch int, seed int64) {
	// Own the receive right; the consumer sends the surface ports here.
	recv, err := mach.NewPort()
	if err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	if err := recv.MakeSendRight(); err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	service := fmt.Sprintf("com.tmc.surfaceffn.%d", os.Getpid())
	if err := mach.BootstrapRegister(service, recv); err != nil {
		log.Fatalf("surfaceffn: bootstrap_register: %v", err)
	}

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	child := exec.Command(exe, "-consume",
		"-service", service,
		"-in", fmt.Sprint(inDim), "-out", fmt.Sprint(outDim),
		"-batch", fmt.Sprint(batch))
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	stdout, err := child.StdoutPipe()
	if err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	if err := child.Start(); err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}

	sc := bufio.NewScanner(stdout)
	readLine := func() string {
		if !sc.Scan() {
			log.Fatal("surfaceffn: consumer exited early")
		}
		return sc.Text()
	}

	// The consumer reports the compiled layouts (or that the ANE is absent).
	first := readLine()
	if reason, ok := strings.CutPrefix(first, "unavail "); ok {
		log.Printf("surfaceffn: ANE unavailable on this machine: %s", reason)
		child.Wait()
		return
	}
	inLayout := parseLayout(first, "in")
	outLayout := parseLayout(readLine(), "out")

	// Receive the kernel's own surfaces: the pages the NPU DMAs from.
	m, err := mach.Receive(recv, 10*time.Second)
	if err != nil {
		log.Fatalf("surfaceffn: receive surface ports: %v", err)
	}
	if len(m.Ports) != 2 {
		log.Fatalf("surfaceffn: message carried %d ports, want 2", len(m.Ports))
	}
	inRef := iosurface.IOSurfaceLookupFromMachPort(uint32(m.Ports[0]))
	outRef := iosurface.IOSurfaceLookupFromMachPort(uint32(m.Ports[1]))
	if inRef == 0 || outRef == 0 {
		log.Fatal("surfaceffn: IOSurfaceLookupFromMachPort failed")
	}
	m.Ports[0].Deallocate()
	m.Ports[1].Deallocate()
	recv.DestroyReceive()

	inSurf, err := ane.WrapIOSurfaceFloat32WithLayout(coregraphics.IOSurfaceRef(inRef), inLayout)
	if err != nil {
		log.Fatalf("surfaceffn: wrap input surface: %v", err)
	}
	outSurf, err := ane.WrapIOSurfaceFloat32WithLayout(coregraphics.IOSurfaceRef(outRef), outLayout)
	if err != nil {
		log.Fatalf("surfaceffn: wrap output surface: %v", err)
	}

	log.Printf("producer: received the ANE kernel's input and output surfaces over mach ports")
	log.Printf("producer: input [1,%d,1,%d] fp32 (%d KiB), output [1,%d,1,%d] fp32 (%d KiB)",
		inLayout.Channels, inLayout.Width, inLayout.AllocSize()>>10,
		outLayout.Channels, outLayout.Width, outLayout.AllocSize()>>10)

	// Build x (channel-first [inDim, batch]) and w (row-major [inDim,
	// outDim]) and pack them the way the kernel's slice_by_size expects:
	// each input channel's row is [batch activations][outDim weight row].
	rng := rand.New(rand.NewSource(seed))
	xCF := randTensor(rng, inDim*batch)
	w := randTensor(rng, inDim*outDim)
	spatial := batch + outDim
	packed := make([]float32, inDim*spatial)
	for in := range inDim {
		row := packed[in*spatial : (in+1)*spatial]
		copy(row[:batch], xCF[in*batch:(in+1)*batch])
		copy(row[batch:], w[in*outDim:(in+1)*outDim])
	}
	if err := inSurf.Write(packed); err != nil {
		log.Fatalf("surfaceffn: write into ANE input pages: %v", err)
	}

	eval := func() uint64 {
		fmt.Fprintln(stdin, "eval")
		line := readLine()
		var hwNS uint64
		if _, err := fmt.Sscanf(line, "done %d", &hwNS); err != nil {
			log.Fatalf("surfaceffn: consumer said %q, want done", line)
		}
		return hwNS
	}

	// Proof 1: the NPU computed over the pages this process wrote.
	hwNS := eval()
	got, err := outSurf.Read()
	if err != nil {
		log.Fatalf("surfaceffn: read ANE output pages: %v", err)
	}
	want := reference(xCF, w, inDim, outDim, batch)
	if err := verify(got, want); err != nil {
		log.Fatalf("surfaceffn: %v", err)
	}
	log.Printf("producer: ANE result read from the kernel's output surface matches the CPU reference")
	log.Printf("producer: the consumer never touched a tensor element — data plane fully cross-process")

	// Proof 2: mutate one activation in place; the next eval must see it.
	const mutIn, mutB = 5, 0
	newVal := xCF[mutIn*batch+mutB] + 2
	xCF[mutIn*batch+mutB] = newVal
	if err := inSurf.WriteAt(mutIn*spatial+mutB, []float32{newVal}); err != nil {
		log.Fatalf("surfaceffn: mutate input page: %v", err)
	}
	eval()
	got, err = outSurf.Read()
	if err != nil {
		log.Fatalf("surfaceffn: read ANE output pages: %v", err)
	}
	if err := verify(got, reference(xCF, w, inDim, outDim, batch)); err != nil {
		log.Fatalf("surfaceffn: ANE did not observe the in-place write: %v", err)
	}
	log.Printf("producer: ANE observed the in-place CPU write on re-eval — live pages, not a snapshot")

	// Bench: wall time per eval (round trip over the pipe included, and
	// said so) next to the ANE's own hardware-reported execution time.
	iters := 1
	for {
		start := time.Now()
		for range iters {
			hwNS = eval()
		}
		if elapsed := time.Since(start); elapsed >= 200*time.Millisecond {
			wall := elapsed / time.Duration(iters)
			flops := 2 * float64(inDim) * float64(outDim) * float64(batch)
			log.Printf("bench: %v wall per eval (incl. pipe round trip), %.1f GFLOPS", wall.Round(time.Microsecond), flops/wall.Seconds()/1e9)
			if hwNS > 0 {
				hw := time.Duration(hwNS)
				log.Printf("bench: %v hardware-reported on-device, %.1f GFLOPS", hw.Round(time.Microsecond), flops/hw.Seconds()/1e9)
			}
			break
		}
		iters *= 2
	}

	fmt.Fprintln(stdin, "quit")
	if err := child.Wait(); err != nil {
		log.Fatalf("surfaceffn: consumer: %v", err)
	}
}

// parseLayout decodes one "layout <name> C W H E R P" line from the consumer.
func parseLayout(line, name string) ane.TensorLayout {
	var got string
	var l ane.TensorLayout
	if _, err := fmt.Sscanf(line, "layout %s %d %d %d %d %d %d", &got,
		&l.Channels, &l.Width, &l.Height, &l.ElemSize, &l.RowStride, &l.PlaneStride); err != nil || got != name {
		log.Fatalf("surfaceffn: consumer said %q, want %s layout", line, name)
	}
	return l
}

// reference computes y[o*batch+b] = Σ_in x[in*batch+b]·w[in*outDim+o] in
// float64: the channel-first product the kernel emits.
func reference(xCF, w []float32, inDim, outDim, batch int) []float64 {
	y := make([]float64, outDim*batch)
	for in := range inDim {
		x := xCF[in*batch:]
		wr := w[in*outDim:]
		for o := range outDim {
			wv := float64(wr[o])
			out := y[o*batch:]
			for b := range batch {
				out[b] += float64(x[b]) * wv
			}
		}
	}
	return y
}

// verify compares under an fp16-scale relative tolerance: the ANE computes
// this product in reduced precision regardless of the fp32 surface edges.
func verify(got []float32, want []float64) error {
	var maxAbs float64
	for _, v := range want {
		if a := math.Abs(v); a > maxAbs {
			maxAbs = a
		}
	}
	if maxAbs == 0 {
		maxAbs = 1
	}
	for i, v := range want {
		if d := math.Abs(float64(got[i])-v) / maxAbs; d > 2e-2 {
			return fmt.Errorf("verification: relative error %.2e at index %d exceeds 2e-2", d, i)
		}
	}
	return nil
}

func randTensor(rng *rand.Rand, size int) []float32 {
	t := make([]float32, size)
	for i := range t {
		t[i] = rng.Float32()*2 - 1
	}
	return t
}
