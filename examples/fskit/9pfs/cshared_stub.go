//go:build darwin && cshared && ninepfs_stubgo

package main

/*
#include <stdint.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tmc/apple/objc"
)

//export NinePFSInit
func NinePFSInit() C.int {
	return 0
}

//export NinePFSNewFileSystem
func NinePFSNewFileSystem() unsafe.Pointer {
	return nil
}

//export NinePFSConfigureFileSystem
func NinePFSConfigureFileSystem(raw unsafe.Pointer) C.int {
	return 0
}

//export NinePFSProbeResource
func NinePFSProbeResource(self unsafe.Pointer, resource unsafe.Pointer, reply unsafe.Pointer) {
}

//export NinePFSLoadResource
func NinePFSLoadResource(self unsafe.Pointer, resource unsafe.Pointer, options unsafe.Pointer, reply unsafe.Pointer) {
	server, err := ensureServer(objc.GetClass("NinePFileSystem"), &ninepFileSystem{config: defaultFSConfigFromEnv()})
	if err != nil {
		extensionLog("register bridge: " + err.Error())
		return
	}
	server.LoadResource(objcID(self), objcID(resource), objcID(options), objcID(reply))
}

//export NinePFSUnloadResource
func NinePFSUnloadResource(self unsafe.Pointer, resource unsafe.Pointer, options unsafe.Pointer, reply unsafe.Pointer) {
}

//export NinePFSLastError
func NinePFSLastError() *C.char {
	return nil
}

func extensionLog(msg string) {
	line := "9pfs: " + msg
	fmt.Fprintln(os.Stderr, line)
}

func objcID(p unsafe.Pointer) objc.ID {
	return objc.ID(uintptr(p))
}
