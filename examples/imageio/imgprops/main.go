// Command imgprops prints the properties of an image file using the Image I/O
// framework. It reports the container type, the number of images in the file,
// and the property dictionary of each image, including nested dictionaries such
// as Exif, TIFF and GPS.
//
// Usage: imgprops <image-file>...
package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/imageio"
)

const utf8Encoding = uint32(corefoundation.KCFStringEncodingUTF8)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: imgprops <image-file>...\n")
		os.Exit(1)
	}
	status := 0
	for _, path := range os.Args[1:] {
		if err := dump(path); err != nil {
			fmt.Fprintf(os.Stderr, "imgprops: %v\n", err)
			status = 1
		}
	}
	os.Exit(status)
}

// dump prints the Image I/O properties of the image file at path.
func dump(path string) error {
	cfPath := corefoundation.CFStringCreateWithCString(0, path, utf8Encoding)
	if cfPath == 0 {
		return fmt.Errorf("%s: cannot encode path", path)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cfPath))

	url := corefoundation.CFURLCreateWithFileSystemPath(0, cfPath, corefoundation.KCFURLPOSIXPathStyle, false)
	if url == 0 {
		return fmt.Errorf("%s: cannot make file URL", path)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(url))

	src := imageio.CGImageSourceCreateWithURL(url, 0)
	if src == 0 {
		return fmt.Errorf("%s: not a readable image", path)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(src))

	count := int(imageio.CGImageSourceGetCount(src))
	fmt.Printf("%s\n", path)
	fmt.Printf("  type:   %s\n", goString(imageio.CGImageSourceGetType(src)))
	fmt.Printf("  images: %d\n", count)

	for i := 0; i < count; i++ {
		props := imageio.CGImageSourceCopyPropertiesAtIndex(src, uintptr(i), 0)
		if props == 0 {
			fmt.Printf("  image %d: no properties\n", i)
			continue
		}
		fmt.Printf("  image %d:\n", i)
		printDict(props, "    ")
		corefoundation.CFRelease(corefoundation.CFTypeRef(props))
	}
	return nil
}

// printDict prints the entries of dict, one per line, indented by indent.
// Values that are themselves dictionaries are printed recursively.
func printDict(dict corefoundation.CFDictionaryRef, indent string) {
	n := corefoundation.CFDictionaryGetCount(dict)
	if n == 0 {
		return
	}
	keys := make([]unsafe.Pointer, n)
	values := make([]unsafe.Pointer, n)
	corefoundation.CFDictionaryGetKeysAndValues(dict, unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]))

	type entry struct{ key, value unsafe.Pointer }
	entries := make([]entry, n)
	for i := range entries {
		entries[i] = entry{keys[i], values[i]}
	}
	sort.Slice(entries, func(i, j int) bool {
		return goString(corefoundation.CFStringRef(entries[i].key)) < goString(corefoundation.CFStringRef(entries[j].key))
	})

	dictType := corefoundation.CFDictionaryGetTypeID()
	for _, e := range entries {
		name := goString(corefoundation.CFStringRef(e.key))
		value := corefoundation.CFTypeRef(e.value)
		if value != nil && corefoundation.CFGetTypeID(value) == dictType {
			fmt.Printf("%s%s:\n", indent, name)
			printDict(corefoundation.CFDictionaryRef(e.value), indent+"  ")
			continue
		}
		fmt.Printf("%s%s: %s\n", indent, name, description(value))
	}
}

// description returns a readable rendering of the Core Foundation value v.
// Strings, numbers and booleans are unwrapped; anything else falls back to
// CFCopyDescription.
func description(v corefoundation.CFTypeRef) string {
	if v == nil {
		return "<nil>"
	}
	switch corefoundation.CFGetTypeID(v) {
	case corefoundation.CFStringGetTypeID():
		return goString(corefoundation.CFStringRef(v))
	case corefoundation.CFBooleanGetTypeID():
		return fmt.Sprint(corefoundation.CFBooleanGetValue(corefoundation.CFBooleanRef(v)))
	case corefoundation.CFNumberGetTypeID():
		num := corefoundation.CFNumberRef(v)
		if corefoundation.CFNumberIsFloatType(num) {
			var f float64
			if corefoundation.CFNumberGetValue(num, corefoundation.KCFNumberDoubleType, unsafe.Pointer(&f)) {
				return fmt.Sprintf("%g", f)
			}
		} else {
			var i int64
			if corefoundation.CFNumberGetValue(num, corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&i)) {
				return fmt.Sprint(i)
			}
		}
	}
	desc := corefoundation.CFCopyDescription(v)
	if desc == 0 {
		return "<no description>"
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(desc))
	return goString(desc)
}

// goString converts a CFString to a Go string. It returns "" if s is nil or
// cannot be converted.
func goString(s corefoundation.CFStringRef) string {
	if s == 0 {
		return ""
	}
	size := 4*corefoundation.CFStringGetLength(s) + 1
	buf := make([]byte, size)
	if !corefoundation.CFStringGetCString(s, &buf[0], size, utf8Encoding) {
		return ""
	}
	if i := bytes.IndexByte(buf, 0); i >= 0 {
		buf = buf[:i]
	}
	return string(buf)
}
