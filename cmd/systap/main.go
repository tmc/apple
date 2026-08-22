// Command systap records macOS system audio output to a 48kHz WAV file using pure-Go Core Audio process taps.
//
// Usage:
//
//	systap -o out.wav [-pid N | -bundle com.brave.Browser] [-d 10s]
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/coreaudio"
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Options configures system audio recording.
type Options struct {
	OutputFile string    // Output WAV filepath or "-" for stdout.
	PID        int       // Target process ID to capture (0 for all system audio output).
	BundleID   string    // Target application bundle ID (e.g. "com.brave.Browser").
	SampleRate int       // Audio sample rate in Hz (default 48000).
	Channels   int       // Number of audio channels (default 2).
	Verbose    bool      // Enable verbose progress messages.
	Writer     io.Writer // Optional output writer (if set, overrides OutputFile).
}

// Recorder manages system audio capture using a Core Audio process tap and aggregate device.
type Recorder struct {
	opts        Options
	tapDesc     coreaudio.CATapDescription
	tapID       uint32
	aggDeviceID uint32
	procID      coreaudio.AudioDeviceIOProcID

	mu                sync.Mutex
	file              *os.File
	outWriter         io.Writer
	dataBytesRecorded uint32
	callbacksCount    uint64
	bytesCaptured     uint64
	running           bool
}

// NewRecorder creates and initializes an audio tap recorder with the specified options.
func NewRecorder(opts Options) (*Recorder, error) {
	if opts.SampleRate <= 0 {
		opts.SampleRate = 48000
	}
	if opts.Channels <= 0 {
		opts.Channels = 2
	}
	return &Recorder{opts: opts}, nil
}

// Start begins capturing system or process audio output.
func (r *Recorder) Start() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.running {
		return fmt.Errorf("systap: recorder already running")
	}

	if r.opts.Writer != nil {
		r.outWriter = r.opts.Writer
	} else if r.opts.OutputFile == "-" || r.opts.OutputFile == "" {
		r.outWriter = os.Stdout
	} else {
		f, err := os.Create(r.opts.OutputFile)
		if err != nil {
			return fmt.Errorf("systap: create output file: %w", err)
		}
		r.file = f
		r.outWriter = f
	}

	if err := writeWAVHeader(r.outWriter, r.opts.SampleRate, r.opts.Channels, 0); err != nil {
		return fmt.Errorf("systap: write WAV header: %w", err)
	}

	if r.opts.PID > 0 {
		procNum := foundation.NewNumberWithInt(int32(r.opts.PID))
		r.tapDesc = coreaudio.NewTapDescriptionStereoMixdownOfProcesses([]foundation.NSNumber{procNum})
	} else if r.opts.BundleID != "" {
		r.tapDesc = coreaudio.NewTapDescriptionStereoGlobalTapButExcludeProcesses(nil)
		r.tapDesc.SetBundleIDs([]string{r.opts.BundleID})
	} else {
		r.tapDesc = coreaudio.NewTapDescriptionStereoGlobalTapButExcludeProcesses(nil)
	}

	r.tapDesc.SetPrivate(false)
	r.tapDesc.SetMuteBehavior(0)

	uuidStr := r.tapDesc.UUID().UUIDString()
	if r.opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: tap UUID: %s\n", uuidStr)
	}

	var tapID uint32
	status := coreaudio.AudioHardwareCreateProcessTap(&r.tapDesc, &tapID)
	if status != 0 {
		return fmt.Errorf("systap: AudioHardwareCreateProcessTap failed with status 0x%x (%d)", status, status)
	}
	r.tapID = tapID

	keyTapUID := foundation.NewStringWithString("uid")
	valTapUID := foundation.NewStringWithString(uuidStr)
	keyTapDrift := foundation.NewStringWithString("drift")
	valTapDrift := foundation.NewNumberWithInt(1)
	tapDict := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{valTapUID, valTapDrift},
		[]objectivec.IObject{keyTapUID, keyTapDrift},
	)
	tapsArray := foundation.NewArrayWithObject(tapDict)

	keyName := foundation.NewStringWithString("name")
	valName := foundation.NewStringWithString("SystapAggregate")
	keyUID := foundation.NewStringWithString("uid")
	valUID := foundation.NewStringWithString("com.tmc.systap.aggregate." + uuidStr)
	keyPrivate := foundation.NewStringWithString("private")
	valPrivate := foundation.NewNumberWithInt(1)
	keyStacked := foundation.NewStringWithString("stacked")
	valStacked := foundation.NewNumberWithInt(0)
	keyTaps := foundation.NewStringWithString("taps")

	aggDict := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{valName, valUID, valPrivate, valStacked, tapsArray},
		[]objectivec.IObject{keyName, keyUID, keyPrivate, keyStacked, keyTaps},
	)

	var aggDeviceID uint32
	status = coreaudio.AudioHardwareCreateAggregateDevice(corefoundation.CFDictionaryRef(aggDict.ID), &aggDeviceID)
	if status != 0 {
		coreaudio.AudioHardwareDestroyProcessTap(r.tapID)
		r.tapID = 0
		return fmt.Errorf("systap: AudioHardwareCreateAggregateDevice failed with status 0x%x (%d)", status, status)
	}
	r.aggDeviceID = aggDeviceID

	ioProc := func(inDevice uint32, inNow uintptr, inInputData uintptr, inInputTime uintptr, outOutputData uintptr, inOutputTime uintptr, inClientData unsafe.Pointer) int32 {
		atomic.AddUint64(&r.callbacksCount, 1)
		if inInputData == 0 {
			return 0
		}
		bufList := (*coreaudiotypes.AudioBufferList)(unsafe.Pointer(inInputData))
		for i := uint32(0); i < bufList.MNumberBuffers; i++ {
			buf := (*coreaudiotypes.AudioBuffer)(unsafe.Pointer(uintptr(unsafe.Pointer(&bufList.MBuffers[0])) + uintptr(i)*unsafe.Sizeof(coreaudiotypes.AudioBuffer{})))
			if buf.MDataByteSize == 0 || buf.MData == nil {
				continue
			}

			floatSamples := unsafe.Slice((*float32)(buf.MData), buf.MDataByteSize/4)
			pcmBytes := make([]byte, len(floatSamples)*2)
			for idx, fVal := range floatSamples {
				sample := float64(fVal)
				if sample > 1.0 {
					sample = 1.0
				} else if sample < -1.0 {
					sample = -1.0
				}
				i16 := int16(sample * 32767.0)
				binary.LittleEndian.PutUint16(pcmBytes[idx*2:], uint16(i16))
			}

			r.mu.Lock()
			if r.outWriter != nil {
				n, _ := r.outWriter.Write(pcmBytes)
				r.dataBytesRecorded += uint32(n)
				atomic.AddUint64(&r.bytesCaptured, uint64(n))
			}
			r.mu.Unlock()
		}
		return 0
	}

	var procID coreaudio.AudioDeviceIOProcID
	status = coreaudio.AudioDeviceCreateIOProcID(r.aggDeviceID, ioProc, nil, &procID)
	if status != 0 {
		coreaudio.AudioHardwareDestroyAggregateDevice(r.aggDeviceID)
		coreaudio.AudioHardwareDestroyProcessTap(r.tapID)
		r.aggDeviceID = 0
		r.tapID = 0
		return fmt.Errorf("systap: AudioDeviceCreateIOProcID failed with status 0x%x (%d)", status, status)
	}
	r.procID = procID

	status = coreaudio.AudioDeviceStart(r.aggDeviceID, r.procID)
	if status != 0 {
		coreaudio.AudioDeviceDestroyIOProcID(r.aggDeviceID, r.procID)
		coreaudio.AudioHardwareDestroyAggregateDevice(r.aggDeviceID)
		coreaudio.AudioHardwareDestroyProcessTap(r.tapID)
		r.procID = nil
		r.aggDeviceID = 0
		r.tapID = 0
		return fmt.Errorf("systap: AudioDeviceStart failed with status 0x%x (%d)", status, status)
	}

	r.running = true
	return nil
}

// Stop halts audio capture and finalizes the output file header.
func (r *Recorder) Stop() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.running {
		return nil
	}

	r.running = false

	if r.aggDeviceID != 0 && r.procID != nil {
		coreaudio.AudioDeviceStop(r.aggDeviceID, r.procID)
		coreaudio.AudioDeviceDestroyIOProcID(r.aggDeviceID, r.procID)
		r.procID = nil
	}
	if r.aggDeviceID != 0 {
		coreaudio.AudioHardwareDestroyAggregateDevice(r.aggDeviceID)
		r.aggDeviceID = 0
	}
	if r.tapID != 0 {
		coreaudio.AudioHardwareDestroyProcessTap(r.tapID)
		r.tapID = 0
	}

	if r.file != nil {
		dataLen := r.dataBytesRecorded
		r.file.Seek(4, io.SeekStart)
		binary.Write(r.file, binary.LittleEndian, uint32(36+dataLen))
		r.file.Seek(40, io.SeekStart)
		binary.Write(r.file, binary.LittleEndian, dataLen)
		r.file.Close()
		r.file = nil
	}

	return nil
}

// CallbacksCount returns the total number of audio IOProc callbacks processed.
func (r *Recorder) CallbacksCount() uint64 {
	return atomic.LoadUint64(&r.callbacksCount)
}

// BytesRecorded returns the total number of PCM data bytes written.
func (r *Recorder) BytesRecorded() uint64 {
	return atomic.LoadUint64(&r.bytesCaptured)
}

func writeWAVHeader(w io.Writer, sampleRate, channels int, dataSize uint32) error {
	bitsPerSample := 16
	byteRate := sampleRate * channels * (bitsPerSample / 8)
	blockAlign := channels * (bitsPerSample / 8)

	header := make([]byte, 44)
	copy(header[0:4], []byte("RIFF"))
	binary.LittleEndian.PutUint32(header[4:8], 36+dataSize)
	copy(header[8:12], []byte("WAVE"))
	copy(header[12:16], []byte("fmt "))
	binary.LittleEndian.PutUint32(header[16:20], 16)
	binary.LittleEndian.PutUint16(header[20:22], 1)
	binary.LittleEndian.PutUint16(header[22:24], uint16(channels))
	binary.LittleEndian.PutUint32(header[24:28], uint32(sampleRate))
	binary.LittleEndian.PutUint32(header[28:32], uint32(byteRate))
	binary.LittleEndian.PutUint16(header[32:34], uint16(blockAlign))
	binary.LittleEndian.PutUint16(header[34:36], uint16(bitsPerSample))
	copy(header[36:40], []byte("data"))
	binary.LittleEndian.PutUint32(header[40:44], dataSize)

	_, err := w.Write(header)
	return err
}

func main() {
	opts := Options{}
	flag.StringVar(&opts.OutputFile, "o", "out.wav", "Output WAV file path ('-' for stdout)")
	flag.IntVar(&opts.PID, "pid", 0, "Target process ID to capture (0 for system audio output)")
	flag.StringVar(&opts.BundleID, "bundle", "", "Target bundle ID to capture (e.g. 'com.brave.Browser')")
	flag.IntVar(&opts.SampleRate, "rate", 48000, "Audio sample rate in Hz")
	flag.IntVar(&opts.Channels, "channels", 2, "Number of audio channels")
	flag.BoolVar(&opts.Verbose, "v", false, "Enable verbose output")
	duration := flag.Duration("d", 0, "Recording duration (e.g. 10s, 1m). Default 0 records until SIGINT/SIGTERM")
	flag.Parse()

	if err := run(opts, *duration); err != nil {
		fmt.Fprintf(os.Stderr, "systap: %v\n", err)
		os.Exit(1)
	}
}

func run(opts Options, duration time.Duration) error {
	rec, err := NewRecorder(opts)
	if err != nil {
		return fmt.Errorf("initialize recorder: %w", err)
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: starting audio capture to %s (rate=%d, channels=%d)...\n",
			opts.OutputFile, opts.SampleRate, opts.Channels)
	}

	if err := rec.Start(); err != nil {
		return fmt.Errorf("start capture: %w", err)
	}
	defer rec.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	if duration > 0 {
		timer := time.NewTimer(duration)
		select {
		case <-timer.C:
			if opts.Verbose {
				fmt.Fprintf(os.Stderr, "systap: duration %v elapsed, stopping...\n", duration)
			}
		case sig := <-sigCh:
			timer.Stop()
			if opts.Verbose {
				fmt.Fprintf(os.Stderr, "systap: received signal %v, stopping...\n", sig)
			}
		}
	} else {
		sig := <-sigCh
		if opts.Verbose {
			fmt.Fprintf(os.Stderr, "systap: received signal %v, stopping...\n", sig)
		}
	}

	if err := rec.Stop(); err != nil {
		return fmt.Errorf("stop capture: %w", err)
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: capture complete. Callbacks: %d, Bytes written: %d\n",
			rec.CallbacksCount(), rec.BytesRecorded())
	}

	return nil
}
