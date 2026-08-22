//go:build cgo && darwin

package signpost

// This file contributes no runtime cgo calls. It exists to plant the strings
// the package emits as signpost formats into the binary's
// __TEXT,__oslogstring section, where the log tools can resolve them by
// offset at decode time (see pool.go). Binaries built without cgo can plant
// strings with a .syso instead; pool.go shows how.

/*
__attribute__((used, section("__TEXT,__oslogstring")))
static const char signpost_pool_message_format[] = "%{public}s";
__attribute__((used, section("__TEXT,__oslogstring")))
static const char signpost_pool_empty_format[] = "";
*/
import "C"
