// Command audiomix demonstrates editing audio with AVAudioMix.
//
// It loads an audio or video file, applies volume ramp effects using
// AVMutableAudioMix, and either previews the result with AVPlayer or
// exports a new file with the mix baked in via AVAssetReader/Writer.
//
// This example is a Go port of Apple's "Editing Spatial Audio with an
// Audio Mix" sample, adapted to use the general-purpose AVAudioMix APIs
// available through the avfoundation bindings.
//
// Usage:
//
//	audiomix preview [--duration <seconds>] <file>
//	audiomix bake [--fade-in <seconds>] [--fade-out <seconds>] --output <outfile> <file>
//
// Examples:
//
//	audiomix preview --duration 10 recording.mov
//	audiomix bake --fade-in 2 --fade-out 3 --output mixed.mov recording.mov
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// cmTime mirrors CoreMedia's CMTime struct layout.
type cmTime struct {
	Value     int64
	Timescale int32
	Flags     uint32
	Epoch     int64
}

// cmTimeRange mirrors CoreMedia's CMTimeRange struct layout.
type cmTimeRange struct {
	Start    cmTime
	Duration cmTime
}

const kCMTimeFlagsValid uint32 = 1

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if len(os.Args) < 3 {
		usage()
		os.Exit(2)
	}

	action := os.Args[1]
	switch action {
	case "preview":
		if err := runPreview(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "audiomix preview: %v\n", err)
			os.Exit(1)
		}
	case "bake":
		if err := runBake(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "audiomix bake: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown action %q\n", action)
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage:
  audiomix preview [--duration <seconds>] <file>
  audiomix bake [--fade-in <seconds>] [--fade-out <seconds>] --output <outfile> <file>
`)
}

func runPreview(args []string) error {
	fs := flag.NewFlagSet("preview", flag.ExitOnError)
	duration := fs.Float64("duration", 5.0, "playback duration in seconds")
	fs.Parse(args)
	if fs.NArg() < 1 {
		return fmt.Errorf("usage: audiomix preview <file>")
	}
	return preview(fs.Arg(0), *duration)
}

func runBake(args []string) error {
	fs := flag.NewFlagSet("bake", flag.ExitOnError)
	output := fs.String("output", "", "output file path")
	fadeIn := fs.Float64("fade-in", 0, "fade-in duration in seconds")
	fadeOut := fs.Float64("fade-out", 0, "fade-out duration in seconds")
	fs.Parse(args)
	if fs.NArg() < 1 || *output == "" {
		return fmt.Errorf("usage: audiomix bake <file> --output <outfile>")
	}
	return bake(fs.Arg(0), *output, *fadeIn, *fadeOut)
}

// preview plays the input file with an audio mix that fades in the first two
// seconds. It uses AVPlayer with an AVAudioMix set on the player item.
func preview(path string, duration float64) error {
	url := foundation.NewURLFileURLWithPath(path)
	asset := avfoundation.NewURLAssetWithURL(url)

	track, err := loadFirstAudioTrack(asset)
	if err != nil {
		return err
	}

	mix := buildFadeMix(track, 2.0, 0, 0)

	item := avfoundation.NewPlayerItemWithAsset(asset)
	item.SetAudioMix(mix)

	player := avfoundation.NewPlayerWithPlayerItem(item)

	fmt.Printf("previewing %s for %.1fs (2s fade-in)\n", path, duration)
	player.Play()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	timer := time.NewTimer(time.Duration(duration * float64(time.Second)))
	defer timer.Stop()

	select {
	case <-timer.C:
	case <-sig:
		fmt.Println()
	}

	player.Pause()
	runtime.KeepAlive(player)
	runtime.KeepAlive(item)
	runtime.KeepAlive(mix)
	return nil
}

// bake reads the input file, applies volume ramp effects (fade-in and/or
// fade-out), and writes the result to a new file using AVAssetReader and
// AVAssetWriter.
func bake(inputPath, outputPath string, fadeIn, fadeOut float64) error {
	inputURL := foundation.NewURLFileURLWithPath(inputPath)
	outputURL := foundation.NewURLFileURLWithPath(outputPath)
	asset := avfoundation.NewURLAssetWithURL(inputURL)

	track, err := loadFirstAudioTrack(asset)
	if err != nil {
		return err
	}

	dur := loadTrackDuration(track)
	fmt.Printf("baking %s -> %s (fade-in=%.1fs, fade-out=%.1fs, duration=%.1fs)\n",
		inputPath, outputPath, fadeIn, fadeOut, dur)

	mix := buildFadeMix(track, fadeIn, fadeOut, dur)

	reader, readerErr := avfoundation.NewAssetReaderWithAssetError(asset)
	if readerErr != nil {
		return fmt.Errorf("create asset reader: %w", readerErr)
	}

	readerOutput := avfoundation.NewAssetReaderAudioMixOutputWithAudioTracksAudioSettings(
		[]avfoundation.AVAssetTrack{track}, nil)
	readerOutput.SetAudioMix(mix)

	// addOutput: is not in generated bindings.
	objc.Send[objc.ID](reader.ID, objc.Sel("addOutput:"), readerOutput.ID)

	writer, writerErr := avfoundation.NewAssetWriterWithURLFileTypeError(outputURL, avfoundation.AVFileTypes.QuickTimeMovie)
	if writerErr != nil {
		return fmt.Errorf("create asset writer: %w", writerErr)
	}

	writerInput := avfoundation.NewAssetWriterInputWithMediaTypeOutputSettings(
		avfoundation.AVMediaTypes.Audio, nil)

	// addInput: is not in generated bindings.
	objc.Send[objc.ID](writer.ID, objc.Sel("addInput:"), writerInput.ID)

	objc.Send[bool](reader.ID, objc.Sel("startReading"))
	objc.Send[bool](writer.ID, objc.Sel("startWriting"))

	// startSessionAtSourceTime: expects a raw CMTime struct.
	zero := cmTime{Value: 0, Timescale: 1, Flags: kCMTimeFlagsValid}
	objc.Send[objc.ID](writer.ID, objc.Sel("startSessionAtSourceTime:"), zero)

	samples := 0
	for {
		// Wait until the writer input is ready for more data.
		for !objc.Send[bool](writerInput.ID, objc.Sel("isReadyForMoreMediaData")) {
			time.Sleep(10 * time.Millisecond)
		}
		buf := objc.Send[objc.ID](readerOutput.ID, objc.Sel("copyNextSampleBuffer"))
		if buf == 0 {
			break
		}
		objc.Send[bool](writerInput.ID, objc.Sel("appendSampleBuffer:"), buf)
		cfRelease(buf)
		samples++
	}

	writerInput.MarkAsFinished()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := writer.FinishWritingSync(ctx); err != nil {
		return fmt.Errorf("finish writing: %w", err)
	}

	fmt.Printf("done: wrote %d sample buffers to %s\n", samples, outputPath)
	return nil
}

// loadFirstAudioTrack loads the first audio track from the asset.
func loadFirstAudioTrack(asset avfoundation.AVURLAsset) (avfoundation.AVAssetTrack, error) {
	// tracksWithMediaType: is not in generated bindings.
	tracks := objc.Send[[]objc.ID](asset.ID, objc.Sel("tracksWithMediaType:"),
		objc.String(string(avfoundation.AVMediaTypes.Audio)))
	if len(tracks) == 0 {
		return avfoundation.AVAssetTrack{}, fmt.Errorf("no audio tracks found")
	}
	return avfoundation.AVAssetTrackFromID(tracks[0]), nil
}

// loadTrackDuration returns the track duration in seconds.
func loadTrackDuration(track avfoundation.AVAssetTrack) float64 {
	tr := objc.Send[cmTimeRange](track.ID, objc.Sel("timeRange"))
	if tr.Duration.Timescale == 0 {
		return 0
	}
	return float64(tr.Duration.Value) / float64(tr.Duration.Timescale)
}

// buildFadeMix creates an AVMutableAudioMix with optional fade-in and fade-out
// volume ramps.
func buildFadeMix(track avfoundation.AVAssetTrack, fadeIn, fadeOut, totalDuration float64) avfoundation.AVMutableAudioMix {
	params := avfoundation.NewMutableAudioMixInputParametersWithTrack(track)
	if params.ID == 0 {
		// Fall back to creating params without a track and setting trackID manually.
		params = avfoundation.GetAVMutableAudioMixInputParametersClass().AudioMixInputParameters()
		objc.Send[objc.ID](params.ID, objc.Sel("setTrackID:"), track.TrackID())
	}

	if fadeIn > 0 {
		r := newCMTimeRange(0, fadeIn)
		// Pass float32 values as float32 (not uintptr) so purego places
		// them in ARM64 float registers via the reflect slow path.
		objc.Send[objc.ID](params.ID,
			objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"),
			float32(0.0), float32(1.0), r)
	}

	if fadeOut > 0 && totalDuration > fadeOut {
		r := newCMTimeRange(totalDuration-fadeOut, fadeOut)
		objc.Send[objc.ID](params.ID,
			objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"),
			float32(1.0), float32(0.0), r)
	}

	mix := avfoundation.GetAVMutableAudioMixClass().AudioMix()
	// setInputParameters: is not in generated bindings.
	objc.Send[objc.ID](mix.ID, objc.Sel("setInputParameters:"),
		foundation.NewArrayWithObject(params))
	return mix
}

func newCMTimeRange(startSeconds, durationSeconds float64) cmTimeRange {
	const timescale int32 = 600
	return cmTimeRange{
		Start: cmTime{
			Value:     int64(startSeconds * float64(timescale)),
			Timescale: timescale,
			Flags:     kCMTimeFlagsValid,
		},
		Duration: cmTime{
			Value:     int64(durationSeconds * float64(timescale)),
			Timescale: timescale,
			Flags:     kCMTimeFlagsValid,
		},
	}
}

// cfRelease calls CFRelease on a Core Foundation object.
func cfRelease(id objc.ID) {
	objc.Send[objc.ID](id, objc.Sel("release"))
}
