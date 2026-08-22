// Package main implements system and process audio recording on macOS using Core Audio process taps.
package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"sync/atomic"
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

	// 1. Prepare output destination
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

	// Write initial WAV header (data size unknown yet)
	if err := writeWAVHeader(r.outWriter, r.opts.SampleRate, r.opts.Channels, 0); err != nil {
		return fmt.Errorf("systap: write WAV header: %w", err)
	}

	// 2. Build CATapDescription
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
	r.tapDesc.SetMuteBehavior(0) // unmuted

	uuidStr := r.tapDesc.UUID().UUIDString()
	if r.opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: tap UUID: %s\n", uuidStr)
	}

	// 3. Create Process Tap
	var tapID uint32
	status := coreaudio.AudioHardwareCreateProcessTap(&r.tapDesc, &tapID)
	if status != 0 {
		return fmt.Errorf("systap: AudioHardwareCreateProcessTap failed with status 0x%x (%d)", status, status)
	}
	r.tapID = tapID
	if r.opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: created process tap ID %d\n", tapID)
	}

	// 4. Create Aggregate Device wrapping the tap
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
	if r.opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: created aggregate device ID %d\n", aggDeviceID)
	}

	// 5. Install IOProc on aggregate device
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

			// Process float32 PCM samples and write 16-bit LE PCM
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

	// 6. Start IO proc
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

	// Update WAV header sizes if writing to a file
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

// writeWAVHeader writes a standard 44-byte 16-bit PCM WAV header.
func writeWAVHeader(w io.Writer, sampleRate, channels int, dataSize uint32) error {
	bitsPerSample := 16
	byteRate := sampleRate * channels * (bitsPerSample / 8)
	blockAlign := channels * (bitsPerSample / 8)

	header := make([]byte, 44)
	copy(header[0:4], []byte("RIFF"))
	binary.LittleEndian.PutUint32(header[4:8], 36+dataSize)
	copy(header[8:12], []byte("WAVE"))
	copy(header[12:16], []byte("fmt "))
	binary.LittleEndian.PutUint32(header[16:20], 16) // Subchunk1Size (16 for PCM)
	binary.LittleEndian.PutUint16(header[20:22], 1)  // AudioFormat (1 for PCM)
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

// CalculateRMS computes the root mean square (RMS) amplitude of 16-bit PCM audio data.
func CalculateRMS(pcmData []byte) float64 {
	if len(pcmData) < 2 {
		return 0
	}
	numSamples := len(pcmData) / 2
	var sumSquares float64
	for i := 0; i < numSamples; i++ {
		sample := int16(binary.LittleEndian.Uint16(pcmData[i*2:]))
		normalized := float64(sample) / 32767.0
		sumSquares += normalized * normalized
	}
	return math.Sqrt(sumSquares / float64(numSamples))
}
