// escaping-tap-regression demonstrates the original use-after-free that
// motivated applegen's escaping-block classifier (appledocs c09228c541).
//
// AVAudioNode.InstallTapOnBusBufferSizeFormatBlock is the raw generated
// wrapper for installTapOnBus:bufferSize:format:block:. The audio engine
// holds the tap block past the call's return and invokes it on the audio
// I/O thread, so the pre-fix generator's `defer _block.Release()` was a
// use-after-free. AVAudioEngine does not synchronously [block copy] —
// release-before-copy meant the audio thread later saw a freed Go closure
// wrapper and crashed with EXC_BAD_ACCESS / SIGTRAP.
//
// The safe public API is the avfaudio.InstallTapOnBus helper, which
// manages the block lifetime explicitly via a release token. This example
// uses the raw method on purpose, to exercise the previously-buggy
// generator path.
//
// Run:
//
//	go run .
//
// Requires microphone permission. Build into a code-signed .app bundle
// (e.g. with tmc/macgo) or grant Terminal mic access in System Settings
// > Privacy & Security > Microphone. Exits 0 once at least one tap
// callback fires.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; cancel() }()

	engine := avfaudio.NewAVAudioEngine()
	bus := avfaudio.AVAudioNodeBus(0)
	input := engine.InputNode()
	format := input.InputFormatForBus(bus)

	var callbackCount atomic.Int64
	tap := func(buf avfaudio.AVAudioPCMBuffer, _ avfaudio.AVAudioTime) {
		callbackCount.Add(int64(buf.FrameLength()))
	}

	input.RemoveTapOnBus(bus)
	// Call the raw generated method — NOT the avfaudio.InstallTapOnBus
	// helper. Pre-fix this emitted `defer _block3.Release()` and crashed
	// when the audio thread invoked the freed block.
	input.InstallTapOnBusBufferSizeFormatBlock(bus, 4096, format, tap)
	defer input.RemoveTapOnBus(bus)

	if ok, err := engine.StartAndReturnError(); !ok || err != nil {
		fmt.Fprintf(os.Stderr, "AVAudioEngine start: %v (grant mic permission?)\n", err)
		os.Exit(1)
	}
	defer engine.Stop()

	for ctx.Err() == nil {
		if callbackCount.Load() > 0 {
			fmt.Printf("ok: tap callback fired (%d frames captured) without UAF\n", callbackCount.Load())
			return
		}
		corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, 0.05, false)
	}
	fmt.Fprintln(os.Stderr, "no tap callbacks within timeout (mic permission denied?)")
	os.Exit(1)
}
