# rtsynth

Renders audio from a Go thread promoted to the Mach real-time scheduling
class (`x/mach.Thread.SetTimeConstraint`) and joined to the audio device's
workgroup (`kAudioDevicePropertyIOThreadOSWorkgroup` + `os_workgroup_join`).

A dedicated Go OS thread synthesizes an additive tone into a ~21 ms
lock-free ring; the AudioUnit render callback drains the ring on Core
Audio's IO thread. Underruns are zero-filled, counted, and audible.

## The A/B

Load every core, then run both scheduling classes:

    for i in $(seq $(sysctl -n hw.ncpu)); do yes > /dev/null & done
    go run . -dur 4s -harmonics 1024                # promoted
    go run . -dur 4s -harmonics 1024 -no-promote    # default class

`-harmonics` is the CPU-cost dial: one `math.Sin` per harmonic per sample.

## Measured, 2026-08-11

16-core arm64, macOS 26.x, go1.26.3, all cores loaded with `yes`:

| harmonics | promoted | default class |
|---|---|---|
| 1024 | **0** underruns / 375 pulls | **46** underruns / 375 pulls |
| 4096 | 148 / 375 | 316 / 375 |

At 1024 harmonics the promotion is the whole difference: silent vs
audibly glitching. At 4096 the synthesis exceeds the physical CPU budget
and both classes fail — promotion buys scheduling, not cycles.

Exit status is 1 if any underrun occurred, so the run is scriptable as a
gate.
