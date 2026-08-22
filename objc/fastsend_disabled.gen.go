// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

//go:build darwin && objc_slowpath

package objc

// Built with -tags objc_slowpath: initFastSend is never called, so
// objcMsgSendAddr stays zero and Send takes the reflect-based path.
// This file exists to make that choice visible rather than implicit.
