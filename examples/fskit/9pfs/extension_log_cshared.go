//go:build darwin && cshared

package main

/*
#include <dlfcn.h>
#include <stdint.h>
#include <stdlib.h>

static void ninepfs_log_cstring(const char *message) {
	typedef void (*log_fn)(const char *);
	log_fn fn = (log_fn)dlsym(RTLD_DEFAULT, "NinePFSLogCString");
	if (fn != NULL) {
		fn(message);
	}
}
*/
import "C"

import "unsafe"

func nativeExtensionLog(msg string) {
	cstr := C.CString(msg)
	C.ninepfs_log_cstring(cstr)
	C.free(unsafe.Pointer(cstr))
}
