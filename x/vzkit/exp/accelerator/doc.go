// Package accelerator provides experimental accelerator device helpers.
//
// It is intended to gather private accelerator-style device configurations such
// as Neural Engine, Bifrost, VideoToolbox, SEP, and related coprocessor-backed
// devices behind narrowly scoped builders for macOS guest configurations.
//
// This package should remain conservative and prefer small wrappers around the
// underlying configuration types until the feature set stabilizes.
package accelerator
