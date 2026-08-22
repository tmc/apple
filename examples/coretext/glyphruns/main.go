// Command glyphruns lays out a string with Core Text and reports its glyph runs.
//
// The text is attributed with a single font, but Core Text splits it into one
// run per font it actually uses, substituting fonts for characters the base
// font does not cover. Each run is reported with its character range, glyph
// count, typographic metrics, and the PostScript name of the font Core Text
// chose for it.
//
// Usage: glyphruns [-font name] [-size points] [text]
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coretext"
)

const defaultText = "Core Text 字 🙂"

func main() {
	fontName := flag.String("font", "Helvetica", "base font name")
	fontSize := flag.Float64("size", 24, "font size in points")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: glyphruns [-font name] [-size points] [text]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	text := defaultText
	if flag.NArg() > 0 {
		text = strings.Join(flag.Args(), " ")
	}

	name := cfString(*fontName)
	font := coretext.CTFontCreateWithName(name, *fontSize, nil)
	if font == 0 {
		fmt.Fprintf(os.Stderr, "glyphruns: no font named %q\n", *fontName)
		os.Exit(1)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(font))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(name))

	fontKey := cfString(coretext.KCTFontAttributeName)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(fontKey))

	attrs := cfDict(unsafe.Pointer(fontKey), unsafe.Pointer(font))
	defer corefoundation.CFRelease(attrs)

	str := cfString(text)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(str))

	attrString := corefoundation.CFAttributedStringCreate(0, str, corefoundation.CFDictionaryRef(attrs))
	if attrString == 0 {
		fmt.Fprintf(os.Stderr, "glyphruns: cannot build attributed string\n")
		os.Exit(1)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(attrString))

	line := coretext.CTLineCreateWithAttributedString(attrString)
	if line == 0 {
		fmt.Fprintf(os.Stderr, "glyphruns: cannot lay out line\n")
		os.Exit(1)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(line))

	var ascent, descent, leading float64
	width := coretext.CTLineGetTypographicBounds(line, &ascent, &descent, &leading)
	fmt.Printf("text:   %q\n", text)
	fmt.Printf("font:   %s %gpt\n", goString(coretext.CTFontCopyPostScriptName(font)), coretext.CTFontGetSize(font))
	fmt.Printf("line:   %d glyphs, width %.2f, ascent %.2f, descent %.2f, leading %.2f\n",
		coretext.CTLineGetGlyphCount(line), width, ascent, descent, leading)

	runs := coretext.CTLineGetGlyphRuns(line)
	n := corefoundation.CFArrayGetCount(runs)
	fmt.Printf("runs:   %d\n", n)
	for i := 0; i < n; i++ {
		run := coretext.CTRunRef(uintptr(corefoundation.CFArrayGetValueAtIndex(runs, i)))
		r := coretext.CTRunGetStringRange(run)
		var a, d, l float64
		w := coretext.CTRunGetTypographicBounds(run, corefoundation.CFRange{}, &a, &d, &l)

		runFont := "?"
		if attrs := coretext.CTRunGetAttributes(run); attrs != 0 {
			if p := corefoundation.CFDictionaryGetValue(attrs, unsafe.Pointer(fontKey)); p != nil {
				runFont = goString(coretext.CTFontCopyPostScriptName(coretext.CTFontRef(uintptr(p))))
			}
		}
		fmt.Printf("  [%d] chars %d..%d  glyphs %d  width %.2f  ascent %.2f  descent %.2f  status %v  font %s\n",
			i, r.Location, r.Location+r.Length, coretext.CTRunGetGlyphCount(run),
			w, a, d, coretext.CTRunGetStatus(run), runFont)
	}
}

// cfString creates a CFString from a Go string. The caller owns the result.
func cfString(s string) corefoundation.CFStringRef {
	return corefoundation.CFStringCreateWithCString(0, s, uint32(corefoundation.KCFStringEncodingUTF8))
}

// goString converts a CFString to a Go string.
func goString(s corefoundation.CFStringRef) string {
	if s == 0 {
		return ""
	}
	size := corefoundation.CFStringGetMaximumSizeForEncoding(corefoundation.CFStringGetLength(s), uint32(corefoundation.KCFStringEncodingUTF8)) + 1
	buf := make([]byte, size)
	if !corefoundation.CFStringGetCString(s, &buf[0], size, uint32(corefoundation.KCFStringEncodingUTF8)) {
		return ""
	}
	if i := bytes.IndexByte(buf, 0); i >= 0 {
		buf = buf[:i]
	}
	return string(buf)
}

// cfDict builds a CFDictionary from alternating key/value pairs.
func cfDict(pairs ...unsafe.Pointer) unsafe.Pointer {
	n := len(pairs) / 2
	keys := make([]unsafe.Pointer, n)
	vals := make([]unsafe.Pointer, n)
	for i := 0; i < n; i++ {
		keys[i] = pairs[i*2]
		vals[i] = pairs[i*2+1]
	}
	kcbs := corefoundation.KCFTypeDictionaryKeyCallBacks
	vcbs := corefoundation.KCFTypeDictionaryValueCallBacks
	return unsafe.Pointer(corefoundation.CFDictionaryCreate(0,
		unsafe.Pointer(&keys[0]),
		unsafe.Pointer(&vals[0]),
		n,
		&kcbs,
		&vcbs,
	))
}
