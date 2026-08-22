// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

//go:build darwin && !objc_slowpath

package objc

// The fast path is armed here rather than in objc.gen.go so that the
// objc_slowpath build tag actually disables it. An unconditional init in
// objc.gen.go would arm it under every tag combination.
func init() {
	initFastSend()
}
