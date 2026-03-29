// Package vm provides the core runtime helpers for Apple Virtualization VMs.
//
// It is the smallest useful set of abstractions for working with a running VM:
// queue management, VM creation and lifecycle, state queries, validation, and
// Objective-C retain helpers.
//
// Higher-level packages in x/vzkit build on this package so callers can start
// with a narrow surface and opt into advanced runtime features only when needed.
package vm
