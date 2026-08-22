# gpucompute

Capture the first display with ScreenCaptureKit and run a Metal luminance
reduction directly over each frame's IOSurface:

```sh
go run ./examples/screencapturekit/gpucompute -duration 10s -fps 30
```

Grant the process Screen Recording permission when macOS asks. The output
includes the number of captured and computed frames, GPU time, and the frame
budget. The shader reads the IOSurface-backed BGRA texture; there is no Go
pixel-buffer copy in the capture-to-compute path.
