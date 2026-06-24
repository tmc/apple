//go:build darwin

package main

// These directives keep the binary cgo-free while asking the Go linker to
// record direct load commands for the Swift/FSKit frameworks used by the
// ExtensionFoundation entrypoint shim.
//
//go:cgo_import_dynamic _EFAppExtensionMain _$s19ExtensionFoundation03AppA0PAAE4mainyyKFZ "/System/Library/Frameworks/ExtensionFoundation.framework/Versions/A/ExtensionFoundation"
//go:cgo_import_dynamic _FSUnaryFileSystemExtensionConfiguration _$s5FSKit24UnaryFileSystemExtensionPAAE13configurationQrvg "/System/Library/Frameworks/FSKit.framework/Versions/A/FSKit"
//go:cgo_import_dynamic _SwiftGetWitnessTable _swift_getWitnessTable "/usr/lib/swift/libswiftCore.dylib"
