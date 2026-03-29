// Package framebuffer provides direct framebuffer access and capture helpers.
//
// It is intended to wrap private framebuffer, framebuffer view, and graphics
// scanout APIs so callers can capture images, inspect framebuffer geometry, and
// manage headless or off-screen presentation without relying on window server
// screenshots.
//
// The package should cover:
//
//	Framebuffer and view wrappers
//	Screenshot helpers for headed and headless VMs
//	Scanout enumeration and reconfiguration
//
// This package is the low-level display primitive for higher-level packages
// such as vnc and remote automation tooling.
package framebuffer
