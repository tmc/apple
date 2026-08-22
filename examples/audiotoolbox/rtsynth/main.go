// Command rtsynth renders audio from a Go thread promoted to the real-time
// scheduling class and joined to the audio device's workgroup.
//
// A dedicated Go OS thread synthesizes samples into a short lock-free ring;
// the AudioUnit render callback drains the ring on Core Audio's IO thread.
// The ring is deliberately shallow, so if the synth thread misses its
// cadence the callback underruns and the gap is audible as a dropout.
//
// Run it twice to hear what promotion buys:
//
//	go run . -harmonics 64                # promoted: survives CPU load
//	go run . -harmonics 64 -no-promote    # default class: glitches under load
//
// with something like "yes > /dev/null &" per CPU in another terminal.
// Underruns are counted and reported either way.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/audiotoolbox"
	"github.com/tmc/apple/coreaudio"
	"github.com/tmc/apple/coreaudiotypes"
	appleos "github.com/tmc/apple/os"
	"github.com/tmc/apple/x/mach"
)

const (
	sampleRate = 48000
	channels   = 2

	// ringFrames is the ring capacity. ~21 ms at 48 kHz: deep enough to
	// absorb scheduling jitter on a promoted thread, shallow enough that a
	// preempted default-class producer underruns audibly.
	ringFrames = 1024
	lowWater   = ringFrames / 4
)

// ring is a single-producer single-consumer ring of interleaved float32
// frames. The consumer side runs on Core Audio's IO thread and must not
// allocate, lock, or call into Go's scheduler.
type ring struct {
	buf  [ringFrames * channels]float32
	head atomic.Uint64 // written by consumer
	tail atomic.Uint64 // written by producer
}

func (r *ring) filled() int { return int(r.tail.Load() - r.head.Load()) }

var (
	rng       ring
	underruns atomic.Uint64
	pulls     atomic.Uint64
)

// render drains the ring into the AudioUnit's buffer list. Short reads are
// zero-filled and counted: that is the audible glitch.
func render(refcon, flags, ts unsafe.Pointer, bus, frames uint32, iodata unsafe.Pointer) int32 {
	type audioBuffer struct {
		channels uint32
		byteSize uint32
		data     unsafe.Pointer
	}
	type bufferList struct {
		count   uint32
		buffers [1]audioBuffer
	}
	list := (*bufferList)(iodata)
	pulls.Add(1)
	for i := uint32(0); i < list.count; i++ {
		b := &list.buffers[i]
		n := int(b.byteSize) / 4
		out := unsafe.Slice((*float32)(b.data), n)
		head, tail := rng.head.Load(), rng.tail.Load()
		avail := int(tail - head)
		if avail > n {
			avail = n
		}
		for j := 0; j < avail; j++ {
			out[j] = rng.buf[(head+uint64(j))%uint64(len(rng.buf))]
		}
		for j := avail; j < n; j++ {
			out[j] = 0
		}
		if avail < n {
			underruns.Add(1)
		}
		rng.head.Store(head + uint64(avail))
	}
	return 0
}

// synthesize appends one frame of an additive-synthesis tone. The harmonic
// count is the demo's CPU-cost dial: math.Sin per harmonic per sample.
func synthesize(phase float64, freq float64, harmonics int) float32 {
	var s float64
	for h := 1; h <= harmonics; h++ {
		s += math.Sin(2*math.Pi*freq*float64(h)*phase) / float64(h)
	}
	return float32(s * 0.3)
}

func producer(freq float64, harmonics int, promote bool, wg appleos.OSWorkgroup, done <-chan struct{}) {
	runtime.LockOSThread()

	if promote {
		t, err := mach.ThreadSelf()
		if err != nil {
			log.Fatalf("thread self: %v", err)
		}
		// The producer wakes roughly every quarter-ring (~5 ms). Ask for
		// a matching period with headroom for the synthesis loop.
		err = t.SetTimeConstraint(mach.TimeConstraint{
			Period:      5 * time.Millisecond,
			Computation: 2 * time.Millisecond,
			Constraint:  4 * time.Millisecond,
			Preemptible: true,
		})
		if err != nil {
			log.Fatalf("promote: %v", err)
		}
		if wg != nil {
			var token appleos.OSWorkgroupJoinTokenOpaqueS
			if rc := appleos.OSWorkgroupJoin(wg, appleos.OSWorkgroupJoinToken(unsafe.Pointer(&token))); rc != 0 {
				log.Printf("workgroup join failed (%d); continuing promoted but unjoined", rc)
			} else {
				defer appleos.OSWorkgroupLeave(wg, appleos.OSWorkgroupJoinToken(unsafe.Pointer(&token)))
				log.Printf("joined audio device workgroup")
			}
		}
	}

	var n uint64 // frames synthesized, drives phase
	for {
		select {
		case <-done:
			return
		default:
		}
		for rng.filled() < (ringFrames-lowWater)*channels {
			phase := float64(n) / sampleRate
			v := synthesize(phase, freq, harmonics)
			tail := rng.tail.Load()
			for c := 0; c < channels; c++ {
				rng.buf[(tail+uint64(c))%uint64(len(rng.buf))] = v
			}
			rng.tail.Store(tail + channels)
			n++
		}
		time.Sleep(2 * time.Millisecond)
	}
}

// defaultOutputWorkgroup returns the default output device's IO thread
// workgroup, or nil if the property is unavailable.
func defaultOutputWorkgroup() appleos.OSWorkgroup {
	var device uint32
	size := uint32(unsafe.Sizeof(device))
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: uint32(coreaudio.KAudioHardwarePropertyDefaultOutputDevice),
		MScope:    uint32(coreaudio.KAudioObjectPropertyScopeGlobalValue),
		MElement:  uint32(coreaudio.KAudioObjectPropertyElementMain),
	}
	if rc := coreaudio.AudioObjectGetPropertyData(uint32(coreaudio.KAudioObjectSystemObject), &addr, 0, nil, &size, unsafe.Pointer(&device)); rc != 0 {
		return nil
	}
	var wg appleos.OSWorkgroup
	size = uint32(unsafe.Sizeof(wg))
	addr.MSelector = uint32(coreaudio.KAudioDevicePropertyIOThreadOSWorkgroup)
	if rc := coreaudio.AudioObjectGetPropertyData(device, &addr, 0, nil, &size, unsafe.Pointer(&wg)); rc != 0 {
		return nil
	}
	return wg
}

func main() {
	log.SetFlags(0)
	freq := flag.Float64("freq", 220, "fundamental frequency in Hz")
	harmonics := flag.Int("harmonics", 32, "harmonics per sample (CPU cost dial)")
	dur := flag.Duration("dur", 10*time.Second, "how long to play")
	noPromote := flag.Bool("no-promote", false, "skip real-time promotion and workgroup join")
	flag.Parse()

	// Default output AudioUnit.
	desc := audiotoolbox.AudioComponentDescription{
		ComponentType:         uint32(audiotoolbox.KAudioUnitType_Output),
		ComponentSubType:      uint32(audiotoolbox.KAudioUnitSubType_DefaultOutput),
		ComponentManufacturer: uint32(audiotoolbox.KAudioUnitManufacturer_Apple),
	}
	comp := audiotoolbox.AudioComponentFindNext(0, &desc)
	if comp == 0 {
		log.Fatal("no default output audio component")
	}
	var unit unsafe.Pointer
	if rc := audiotoolbox.AudioComponentInstanceNew(comp, (*audiotoolbox.AudioComponentInstance)(unsafe.Pointer(&unit))); rc != 0 {
		log.Fatalf("AudioComponentInstanceNew: %d", rc)
	}

	// Interleaved native float32 stereo. 'lpcm', IsFloat|IsPacked.
	const lpcm = 'l'<<24 | 'p'<<16 | 'c'<<8 | 'm'
	format := struct {
		sampleRate                                                float64
		formatID, formatFlags, bytesPerPacket, framesPerPacket    uint32
		bytesPerFrame, channelsPerFrame, bitsPerChannel, reserved uint32
	}{
		sampleRate:       sampleRate,
		formatID:         lpcm,
		formatFlags:      uint32(coreaudiotypes.KAudioFormatFlagIsFloatValue) | uint32(coreaudiotypes.KAudioFormatFlagIsPacked),
		bytesPerPacket:   4 * channels,
		framesPerPacket:  1,
		bytesPerFrame:    4 * channels,
		channelsPerFrame: channels,
		bitsPerChannel:   32,
	}
	if rc := audiotoolbox.AudioUnitSetProperty(unit, uint32(audiotoolbox.KAudioUnitProperty_StreamFormat),
		uint32(audiotoolbox.KAudioUnitScope_Input), 0, unsafe.Pointer(&format), uint32(unsafe.Sizeof(format))); rc != 0 {
		log.Fatalf("set stream format: %d", rc)
	}

	// The render callback must be a C function pointer, so the generated
	// AURenderCallbackStruct (whose field is a Go func type) cannot be
	// used here; build the C layout with a purego trampoline instead.
	cb := struct {
		proc   uintptr
		refcon unsafe.Pointer
	}{proc: purego.NewCallback(render)}
	if rc := audiotoolbox.AudioUnitSetProperty(unit, uint32(audiotoolbox.KAudioUnitProperty_SetRenderCallback),
		uint32(audiotoolbox.KAudioUnitScope_Input), 0, unsafe.Pointer(&cb), uint32(unsafe.Sizeof(cb))); rc != 0 {
		log.Fatalf("set render callback: %d", rc)
	}

	var wg appleos.OSWorkgroup
	if !*noPromote {
		if wg = defaultOutputWorkgroup(); wg == nil {
			log.Printf("device workgroup unavailable; promoting without joining")
		}
	}

	done := make(chan struct{})
	go producer(*freq, *harmonics, !*noPromote, wg, done)

	if rc := audiotoolbox.AudioUnitInitialize(unit); rc != 0 {
		log.Fatalf("AudioUnitInitialize: %d", rc)
	}
	if rc := audiotoolbox.AudioOutputUnitStart(unit); rc != 0 {
		log.Fatalf("AudioOutputUnitStart: %d", rc)
	}
	mode := "promoted"
	if *noPromote {
		mode = "default-class"
	}
	log.Printf("playing %.0f Hz, %d harmonics, %s, for %v", *freq, *harmonics, mode, *dur)

	deadline := time.After(*dur)
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	var lastUnder, lastPulls uint64
	for {
		select {
		case <-tick.C:
			u, p := underruns.Load(), pulls.Load()
			log.Printf("pulls %4d  underruns %4d  ring %4d/%d", p-lastPulls, u-lastUnder, rng.filled()/channels, ringFrames)
			lastUnder, lastPulls = u, p
		case <-deadline:
			close(done)
			audiotoolbox.AudioOutputUnitStop(unit)
			audiotoolbox.AudioUnitUninitialize(unit)
			audiotoolbox.AudioComponentInstanceDispose(audiotoolbox.AudioComponentInstance(uintptr(unit)))
			u, p := underruns.Load(), pulls.Load()
			fmt.Printf("\n%s: %d underruns in %d pulls over %v\n", mode, u, p, *dur)
			if u > 0 {
				os.Exit(1)
			}
			return
		}
	}
}
