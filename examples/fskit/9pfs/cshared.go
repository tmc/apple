//go:build darwin && cshared && !ninepfs_stubgo

package main

/*
#include <stdint.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/x/fskitbridge"
)

func init() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT)
	go func() {
		for sig := range ch {
			extensionLog("caught signal " + sig.String())
		}
	}()
}

// replyFallback delivers failure replies when the bridge server is not
// available, so an exported entry point can still answer FSKit.
var replyFallback = fskitbridge.NewReplyBlocksWithShims(ninepShims)

// csharedInit records whether the bridge is initialized and the most recent
// failure for NinePFSLastError. A failure is not sticky: every exported entry
// point retries initialization, so an early call that races extension
// startup (for example before the Swift shim registers its class) does not
// poison the process.
var csharedInit struct {
	sync.Mutex
	ok  bool
	err error
}

//export NinePFSInit
func NinePFSInit() C.int {
	csharedInit.Lock()
	defer csharedInit.Unlock()
	if csharedInit.ok {
		return 0
	}
	cls := objc.GetClass("NinePFileSystem")
	if cls == 0 {
		csharedInit.err = errors.New("swift shim did not register NinePFileSystem")
		return -1
	}
	var err error
	objc.AutoreleasePool(func() {
		_, err = ensureServer(cls, &ninepFileSystem{config: defaultFSConfigFromEnv()})
	})
	if err != nil {
		csharedInit.err = err
		extensionLog("register bridge: " + err.Error())
		return -1
	}
	csharedInit.ok, csharedInit.err = true, nil
	return 0
}

//export NinePFSNewFileSystem
func NinePFSNewFileSystem() unsafe.Pointer {
	if NinePFSInit() != 0 {
		return nil
	}
	// An objc.ID is an object pointer; reinterpret rather than convert
	// through uintptr, which vet rejects.
	fs := currentServer().NewFileSystem()
	return *(*unsafe.Pointer)(unsafe.Pointer(&fs))
}

//export NinePFSConfigureFileSystem
func NinePFSConfigureFileSystem(raw unsafe.Pointer) C.int {
	extensionLog("configure file system begin")
	if NinePFSInit() != 0 {
		extensionLog("configure file system: init failed")
		return -1
	}
	if raw == nil {
		extensionLog("configure file system: nil object")
		return -1
	}
	extensionLog("configure file system ok")
	return 0
}

//export NinePFSProbeResource
func NinePFSProbeResource(self unsafe.Pointer, resource unsafe.Pointer, reply unsafe.Pointer) {
	defer recoverExported("probeResource")
	extensionLog("probeResource begin")
	if NinePFSInit() != 0 {
		_ = replyFallback.ObjectError(objc.ID(uintptr(reply)), 0, posixError(syscall.EINVAL))
		return
	}
	currentServer().ProbeResource(objc.ID(uintptr(self)), objc.ID(uintptr(resource)), objc.ID(uintptr(reply)))
	extensionLog("probeResource end")
}

//export NinePFSLoadResource
func NinePFSLoadResource(self unsafe.Pointer, resource unsafe.Pointer, options unsafe.Pointer, reply unsafe.Pointer) {
	defer recoverExported("loadResource")
	extensionLog("loadResource begin")
	if NinePFSInit() != 0 {
		_ = replyFallback.ObjectError(objc.ID(uintptr(reply)), 0, posixError(syscall.EINVAL))
		return
	}
	currentServer().LoadResource(objc.ID(uintptr(self)), objc.ID(uintptr(resource)), objc.ID(uintptr(options)), objc.ID(uintptr(reply)))
	extensionLog("loadResource end")
}

//export NinePFSUnloadResource
func NinePFSUnloadResource(self unsafe.Pointer, resource unsafe.Pointer, options unsafe.Pointer, reply unsafe.Pointer) {
	defer recoverExported("unloadResource")
	extensionLog("unloadResource begin")
	if NinePFSInit() != 0 {
		_ = replyFallback.Error(objc.ID(uintptr(reply)), posixError(syscall.EINVAL))
		return
	}
	currentServer().UnloadResource(objc.ID(uintptr(self)), objc.ID(uintptr(resource)), objc.ID(uintptr(options)), objc.ID(uintptr(reply)))
	extensionLog("unloadResource end")
}

//export NinePFSLastError
func NinePFSLastError() *C.char {
	csharedInit.Lock()
	defer csharedInit.Unlock()
	if csharedInit.err == nil {
		return nil
	}
	return C.CString(csharedInit.err.Error())
}

func extensionLog(msg string) {
	nativeExtensionLog(msg)
	if os.Getenv("NINEPFS_DEBUG") != "" {
		fmt.Fprintln(os.Stderr, "9pfs: "+msg)
	}
}

func recoverExported(name string) {
	if r := recover(); r != nil {
		extensionLog(fmt.Sprintf("%s panic: %v\n%s", name, r, debug.Stack()))
		os.Exit(2)
	}
}
