//go:build !darwin || !cshared

package main

func nativeExtensionLog(msg string) {}
