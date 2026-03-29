// Package vnc provides headless remote display serving for virtual machines.
//
// It is intended to wrap the private VZVNCServer APIs and their security
// configurations behind a small Go surface for attaching a virtual machine or
// graphics display to a TCP listener or Bonjour-published VNC service.
//
// The package should focus on a small set of concepts:
//
//	Config for port, vm.Queue, and security settings
//	Server for lifecycle management
//	Helpers to attach a VM or specific display
//
// This package relies on private Virtualization.framework APIs and should be
// kept separate from the public VM construction helpers in the parent package.
package vnc
