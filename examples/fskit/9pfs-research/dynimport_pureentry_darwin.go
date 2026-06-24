//go:build darwin && pureentry

package main

import "os"

//go:cgo_import_dynamic nsext_main NSExtensionMain "/System/Library/Frameworks/Foundation.framework/Versions/C/Foundation"

func nsextMainEntry()

func init() {
	if os.Getenv("NINEPFS_FORCE_SWIFT_METADATA_REF") == "1" {
		nsextMainEntry()
	}
}
