// Command systap records macOS system or per-process audio output to a WAV file using Core Audio process taps.
//
// Process taps (macOS 14.2+) are the modern, sanctioned in-process replacement for
// virtual audio loopback drivers (such as BlackHole or Audio Hijack).
//
// # Usage
//
//	systap -o out.wav [-pid N | -bundle com.brave.Browser] [-d 10s]
//
// # Options
//
//	-o string
//		Output WAV file path ('-' for stdout). Default is "out.wav".
//	-pid int
//		Target process ID to capture. Default 0 captures global system audio output.
//	-bundle string
//		Target application bundle ID to capture (e.g. "com.brave.Browser").
//	-d duration
//		Recording duration (e.g. 10s, 1m). Default 0 records until SIGINT/SIGTERM.
//	-rate int
//		Sample rate in Hz. Default is 48000.
//	-channels int
//		Number of audio channels. Default is 2.
//	-v
//		Enable verbose progress output.
//
// # Permissions (TCC)
//
// Capturing system or process audio requires the macOS System Audio Capture permission
// (NSAudioCaptureUsageDescription). When run for the first time, macOS will prompt the user
// to grant permission.
package main
