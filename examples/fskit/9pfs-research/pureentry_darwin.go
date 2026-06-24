//go:build darwin && pureentry

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	classReplaceMethod func(objc.Class, objc.SEL, uintptr, string) uintptr
	syslogPureEntry    func(int32, *byte, *byte)
)

func main() {
	runtime.LockOSThread()
	logPureEntry("start")
	initPureEntry()
	if err := installPureEntryMethods(); err != nil {
		logPureEntry("install methods: %v", err)
		os.Exit(1)
	}
	logPureEntry("installed Go file system methods")
	select {}
}

func initPureEntry() {
	libobjc, err := purego.Dlopen("libobjc.A.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		logPureEntry("dlopen libobjc: %v", err)
		os.Exit(1)
	}
	purego.RegisterLibFunc(&classReplaceMethod, libobjc, "class_replaceMethod")
	if libc, err := purego.Dlopen("/usr/lib/libSystem.B.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL); err == nil {
		purego.RegisterLibFunc(&syslogPureEntry, libc, "syslog")
	}
}

func installPureEntryMethods() error {
	cls := objc.GetClass("NinePFileSystem")
	if cls == 0 {
		return fmt.Errorf("missing NinePFileSystem class")
	}
	if err := registerFSKitBridgeWithFileSystemClass(cls); err != nil {
		return err
	}
	for _, method := range bridgeFileSystemMethods() {
		imp := purego.NewCallback(method.Fn)
		old := classReplaceMethod(cls, method.Cmd, imp, "")
		logPureEntry("replace %v old=%#x", method.Cmd, old)
	}
	return nil
}

func logPureEntry(format string, args ...any) {
	msg := fmt.Sprintf("9pfs-nocgo-entry: "+format, args...)
	fmt.Fprintln(os.Stderr, msg)
	if syslogPureEntry != nil {
		formatBytes := append([]byte("%s"), 0)
		msgBytes := append([]byte(msg), 0)
		syslogPureEntry(5, &formatBytes[0], &msgBytes[0])
		runtime.KeepAlive(formatBytes)
		runtime.KeepAlive(msgBytes)
	}
}
