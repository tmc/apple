# inputtap

Capture audio from the default microphone via AVAudioEngine + tap, resample
to PCM16 mono at 24 kHz with AVAudioConverter, and print frame counts and
peak amplitude every second.

## Run

	go run .

macOS prompts for microphone permission on first launch. Speak into the mic;
expect output like:

	capturing: hwRate=48000 → outRate=24000 (Ctrl-C to stop)
	frames=24000 (+24000) peak=0.123
	frames=48000 (+24000) peak=0.456

Higher peak values mean louder input.

## Notes

- This example exercises the input side of AVAudioEngine, which works
  reliably under purego because tap callbacks fire on a normal dispatch
  queue.
- Streaming PCM to the speakers (AVAudioPlayerNode + ScheduleBuffer or
  AVAudioSourceNode + render block) does not currently work under purego
  for continuous streams under current purego bridging. Pipe PCM to an
  external process such as ffplay, use a small cgo shim for render callbacks,
  or use AudioQueueServices for realtime output.
- macOS TCC requires the binary to be code-signed for microphone access.
  For development, codesign the binary with a Microphone usage entitlement
  or wrap it with [tmc/macgo](https://github.com/tmc/macgo).

## Key APIs

- [`AVAudioEngine.InputNode`](https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputNode)
- [`AVAudioNode.InstallTapOnBusBufferSizeFormatBlock`](https://developer.apple.com/documentation/AVFAudio/AVAudioNode/installTap(onBus:bufferSize:format:block:))
- [`AVAudioConverter.ConvertToBufferErrorWithInputFromBlock`](https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:error:withInputFrom:))

## Implementation gotchas

- Call `runtime.LockOSThread()` at the top of `main`. AVFoundation's
  notification delivery and runloop pumping want a stable OS thread.
- The tap format must equal `InputFormatForBus(bus)` — Apple requires the
  tap and connection formats to match when both are non-nil.
- Microphone input is float32. Convert via `AVAudioConverter` to whatever
  target format you want.
- Use the input-block converter variant
  (`ConvertToBufferErrorWithInputFromBlock`), not the simpler
  `ConvertToBufferFromBufferError` — the conversion is both sample-rate
  AND format, and the basic variant only handles the latter.
- Use `avfaudio.InstallTapOnBus`, which keeps the Objective-C block alive
  until the returned cleanup function runs after `RemoveTapOnBus`.
- Set a `consumed` flag inside the input block so the converter sees
  `AVAudioConverterInputStatus_NoDataNow` after handing over the tap
  buffer once. Without it the converter spins asking for more input.
