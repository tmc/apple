// Command fsevents watches a directory with Apple's journal-backed FSEvents
// service and prints each file-level event delivered to the Go callback.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
)

const utf8 = 0x08000100

// FSEvents is a subframework of CoreServices. On current macOS releases its
// dylib is not loadable through the umbrella CoreServices path, so resolve
// this small C surface from the owning subframework.
type fseventsAPI struct {
	create     func(corefoundation.CFAllocatorRef, unsafe.Pointer, unsafe.Pointer, corefoundation.CFArrayRef, uint64, float64, uint32) coreservices.FSEventStreamRef
	setQueue   func(coreservices.FSEventStreamRef, uintptr)
	start      func(coreservices.FSEventStreamRef) bool
	stop       func(coreservices.FSEventStreamRef)
	invalidate func(coreservices.FSEventStreamRef)
	release    func(coreservices.FSEventStreamRef)
}

func loadFSEvents() (fseventsAPI, error) {
	for _, path := range []string{
		"/System/Library/Frameworks/CoreServices.framework/Frameworks/FSEvents.framework/FSEvents",
		"/System/Library/Frameworks/FSEvents.framework/FSEvents",
	} {
		lib, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			continue
		}
		var api fseventsAPI
		purego.RegisterLibFunc(&api.create, lib, "FSEventStreamCreate")
		purego.RegisterLibFunc(&api.setQueue, lib, "FSEventStreamSetDispatchQueue")
		purego.RegisterLibFunc(&api.start, lib, "FSEventStreamStart")
		purego.RegisterLibFunc(&api.stop, lib, "FSEventStreamStop")
		purego.RegisterLibFunc(&api.invalidate, lib, "FSEventStreamInvalidate")
		purego.RegisterLibFunc(&api.release, lib, "FSEventStreamRelease")
		return api, nil
	}
	return fseventsAPI{}, fmt.Errorf("FSEvents framework unavailable")
}

func main() {
	path := flag.String("path", ".", "directory to watch")
	duration := flag.Duration("duration", 10*time.Second, "watch duration")
	flag.Parse()
	if *duration <= 0 {
		fmt.Fprintln(os.Stderr, "fsevents: duration must be positive")
		os.Exit(2)
	}
	abs, err := filepath.Abs(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fsevents: %v\n", err)
		os.Exit(1)
	}
	api, err := loadFSEvents()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fsevents: %v\n", err)
		os.Exit(1)
	}

	paths := corefoundation.CFArrayCreateMutable(0, 1, &corefoundation.KCFTypeArrayCallBacks)
	pathString := corefoundation.CFStringCreateWithCString(0, abs, utf8)
	corefoundation.CFArrayAppendValue(paths, unsafe.Pointer(pathString))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(paths))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(pathString))

	queue := dispatch.QueueCreate("github.com/tmc/apple.examples.coreservices.fsevents")
	callback := purego.NewCallback(func(_ coreservices.ConstFSEventStreamRef, _ unsafe.Pointer, n int32, pathData unsafe.Pointer, ids unsafe.Pointer, flags unsafe.Pointer) {
		pathPtrs := unsafe.Slice((**byte)(pathData), n)
		eventIDs := unsafe.Slice((*uint64)(ids), n)
		eventFlags := unsafe.Slice((*uint32)(flags), n)
		for i := range pathPtrs {
			name := "<unknown>"
			if pathPtrs[i] != nil {
				name = objc.GoString(pathPtrs[i])
			}
			fmt.Printf("id=%d flags=0x%x path=%s\n", eventIDs[i], eventFlags[i], name)
		}
	})
	stream := api.create(0, unsafe.Pointer(callback), nil,
		corefoundation.CFArrayRef(paths), uint64(coreservices.KFSEventStreamEventIdSinceNow),
		0.25, uint32(coreservices.KFSEventStreamCreateFlagFileEvents|coreservices.KFSEventStreamCreateFlagNoDefer))
	if stream == 0 {
		fmt.Fprintln(os.Stderr, "fsevents: create stream failed")
		os.Exit(1)
	}
	defer api.invalidate(stream)
	defer api.release(stream)
	api.setQueue(stream, uintptr(queue.Handle()))
	if !api.start(stream) {
		fmt.Fprintln(os.Stderr, "fsevents: start stream failed")
		os.Exit(1)
	}
	// Keep the callback trampoline live until after the stream is stopped.
	timer := time.NewTimer(*duration)
	<-timer.C
	api.stop(stream)
	_ = callback
}
