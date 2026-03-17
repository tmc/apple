//go:generate -command applegen-bootstrap applegen bootstrap --module github.com/tmc/apple/private/xcode --force --verbose

//go:generate applegen-bootstrap --private --extra-flags=--framework-path=/Applications/Xcode.app/Contents/PlugIns/GPUDebugger.ideplugin/Contents/Frameworks/GTShaderProfiler.framework GTShaderProfiler

// Package xcode contains generated bindings for private frameworks shipped
// inside Xcode.app and its plug-ins.
package xcode
