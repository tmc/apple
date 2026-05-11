// Command inputtap captures audio from the default microphone via
// AVAudioEngine + tap, resamples it to PCM16 mono at 24 kHz with
// AVAudioConverter, and prints frame counts and peak amplitude every second.
//
// Run with:
//
//	go run .
//
// macOS prompts for microphone permission on first launch. To get past TCC
// reliably, build into a code-signed .app bundle (e.g. with tmc/macgo) or
// codesign the binary with a Microphone usage entitlement.
//
// This example exercises the input side of AVAudioEngine, which works
// reliably under purego because tap callbacks fire on a normal dispatch
// queue. The streaming output side (AVAudioPlayerNode + ScheduleBuffer for
// a continuous stream, or AVAudioSourceNode render block) does not currently
// work under purego — see the avfaudio package documentation for details.
package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; cancel() }()

	const outRate = 24000

	engine := avfaudio.NewAVAudioEngine()
	bus := avfaudio.AVAudioNodeBus(0)
	input := engine.InputNode()
	inFmt := input.InputFormatForBus(bus)

	outFmt := avfaudio.NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved(
		avfaudio.AVAudioPCMFormatInt16, float64(outRate), 1, true,
	)
	conv := avfaudio.NewAudioConverterFromFormatToFormat(inFmt, outFmt)

	var totalFrames atomic.Int64
	var peakBits atomic.Uint32

	tap := func(in avfaudio.AVAudioPCMBuffer, _ avfaudio.AVAudioTime) {
		inFrames := uint32(in.FrameLength())
		ratio := float64(outRate) / inFmt.SampleRate()
		out := avfaudio.NewAudioPCMBufferWithPCMFormatFrameCapacity(
			outFmt, avfaudio.AVAudioFrameCount(float64(inFrames)*ratio)+1024,
		)

		consumed := false
		inputBlock := func(_ uint32, status *avfaudio.AVAudioConverterInputStatus) avfaudio.AVAudioBuffer {
			if consumed {
				*status = avfaudio.AVAudioConverterInputStatus_NoDataNow
				return avfaudio.AVAudioBuffer{}
			}
			consumed = true
			*status = avfaudio.AVAudioConverterInputStatus_HaveData
			return in.AVAudioBuffer
		}
		var nsErr foundation.NSError
		conv.ConvertToBufferErrorWithInputFromBlock(out, nsErr, inputBlock)

		nFrames := int(out.FrameLength())
		if nFrames == 0 {
			return
		}
		channels := (**int16)(out.Int16ChannelData())
		src := unsafe.Slice(*channels, nFrames)

		var peak int16
		for _, s := range src {
			if s < 0 {
				s = -s
			}
			if s > peak {
				peak = s
			}
		}
		totalFrames.Add(int64(nFrames))
		peakBits.Store(math.Float32bits(float32(peak) / 32768.0))
	}

	const tapFrames avfaudio.AVAudioFrameCount = 4096
	input.RemoveTapOnBus(bus)
	releaseTap := avfaudio.InstallTapOnBus(input, bus, tapFrames, inFmt, tap)
	defer func() {
		input.RemoveTapOnBus(bus)
		releaseTap()
	}()

	if ok, err := engine.StartAndReturnError(); !ok || err != nil {
		fmt.Fprintf(os.Stderr, "AVAudioEngine start: %v\n", err)
		os.Exit(1)
	}
	defer engine.Stop()

	fmt.Fprintf(os.Stderr, "capturing: hwRate=%.0f → outRate=%d (Ctrl-C to stop)\n",
		inFmt.SampleRate(), outRate)

	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	last := totalFrames.Load()
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stderr, "stopping")
			return
		case <-tick.C:
			cur := totalFrames.Load()
			peak := math.Float32frombits(peakBits.Swap(0))
			fmt.Fprintf(os.Stderr, "frames=%d (+%d) peak=%.3f\n", cur, cur-last, peak)
			last = cur
		default:
			corefoundation.CFRunLoopRunInMode(
				corefoundation.KCFRunLoopDefaultMode, 0.05, false,
			)
		}
	}
}
