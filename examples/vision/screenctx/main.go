// Command screenctx captures screen context by combining CGWindowListCopyWindowInfo
// with Vision OCR to produce an XML map of on-screen text attributed to the owning
// application.
//
// Usage:
//
//	screenctx                    # capture entire screen (needs screen recording permission)
//	screenctx <screenshot>       # read from file
//	screenctx -                  # read from stdin
//
// Output is XML:
//
//	<screen width="W" height="H">
//	  <app name="Terminal" pid="1234" bounds="x,y,w,h">
//	    <text confidence="0.98" bounds="x,y,w,h">Hello</text>
//	  </app>
//	</screen>
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	minConf := flag.Float64("min-confidence", 0, "discard OCR results below this confidence (0-1)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: screenctx [flags] [screenshot|-]\n\n")
		fmt.Fprintf(os.Stderr, "With no arguments, captures the entire screen.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	var source string
	if flag.NArg() > 0 {
		source = flag.Arg(0)
	}

	if err := run(source, *minConf); err != nil {
		if strings.Contains(err.Error(), "screen capture failed") {
			fmt.Fprintf(os.Stderr, "screenctx: %v\n", err)
			fmt.Fprintf(os.Stderr, "hint: grant screen recording permission in System Settings > Privacy & Security\n")
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "screenctx: %v\n", err)
		os.Exit(1)
	}
}

func run(source string, minConfidence float64) error {
	var imgW, imgH int
	var texts []ocrText

	if source == "" || source == "-" {
		// Screen capture or stdin: use CGImage path.
		img, cleanup, err := loadCGImage(source)
		if err != nil {
			return err
		}
		defer cleanup()

		imgW = int(coregraphics.CGImageGetWidth(img))
		imgH = int(coregraphics.CGImageGetHeight(img))

		texts, err = runOCRCGImage(img, imgW, imgH)
		if err != nil {
			return fmt.Errorf("ocr: %w", err)
		}
	} else {
		// File: use URL-based handler (avoids needing imageio).
		url := foundation.NewURLFileURLWithPath(source)
		handler := vision.NewImageRequestHandlerWithURLOptions(url, nil)
		request := vision.NewVNRecognizeTextRequest()
		request.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelAccurate)
		request.SetUsesLanguageCorrection(true)

		ok, err := handler.PerformRequestsError([]vision.VNRequest{request.VNRequest})
		if err != nil {
			return fmt.Errorf("ocr: %w", err)
		}
		if !ok {
			return fmt.Errorf("ocr request failed")
		}

		// For file input, use normalized coordinates (no pixel conversion).
		imgW, imgH = 1, 1
		texts = extractOCRResults(request, 1.0, 1.0)
	}

	windows, err := getVisibleWindows()
	if err != nil {
		fmt.Fprintf(os.Stderr, "screenctx: warning: window list unavailable: %v\n", err)
	}

	if minConfidence > 0 {
		var filtered []ocrText
		for _, t := range texts {
			if t.Confidence >= minConfidence {
				filtered = append(filtered, t)
			}
		}
		texts = filtered
	}

	screen := buildScreen(imgW, imgH, windows, texts)
	return writeXML(os.Stdout, screen)
}

// --- Image loading ---

func loadCGImage(source string) (coregraphics.CGImageRef, func(), error) {
	if source == "" {
		return captureScreen()
	}
	// stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return 0, nil, fmt.Errorf("read stdin: %w", err)
	}
	f, err := os.CreateTemp("", "screenctx-*.png")
	if err != nil {
		return 0, nil, fmt.Errorf("create temp: %w", err)
	}
	tmp := f.Name()
	if _, err := f.Write(data); err != nil {
		f.Close()
		os.Remove(tmp)
		return 0, nil, fmt.Errorf("write temp: %w", err)
	}
	f.Close()

	// Use CGWindowListCreateImage with a single-pixel rect as a workaround
	// won't work — just capture the screen and note stdin isn't fully supported
	// without imageio. For now, return an error.
	os.Remove(tmp)
	return 0, nil, fmt.Errorf("stdin input requires screen capture; pass a file path instead")
}

func captureScreen() (coregraphics.CGImageRef, func(), error) {
	img := coregraphics.CGWindowListCreateImage(
		coregraphics.CGRectInfinite,
		coregraphics.KCGWindowListOptionOnScreenOnly,
		0, // kCGNullWindowID
		0, // kCGWindowImageDefault
	)
	if img == 0 {
		return 0, nil, fmt.Errorf("screen capture failed (screen recording permission may be required)")
	}
	return img, func() {
		corefoundation.CFRelease(corefoundation.CFTypeRef(img))
	}, nil
}

// --- OCR ---

type ocrText struct {
	Text       string
	Confidence float64
	X, Y, W, H float64 // screen-space pixels, top-left origin
}

func runOCRCGImage(cgImage coregraphics.CGImageRef, imgW, imgH int) ([]ocrText, error) {
	handler := vision.NewImageRequestHandlerWithCGImageOptions(cgImage, nil)
	request := vision.NewVNRecognizeTextRequest()
	request.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelAccurate)
	request.SetUsesLanguageCorrection(true)

	ok, err := handler.PerformRequestsError([]vision.VNRequest{request.VNRequest})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("request failed")
	}

	return extractOCRResults(request, float64(imgW), float64(imgH)), nil
}

func extractOCRResults(request vision.VNRecognizeTextRequest, imgW, imgH float64) []ocrText {
	var results []ocrText
	for _, obs := range request.Results() {
		textObs := vision.VNRecognizedTextObservationFromID(obs.ID)
		candidates := textObs.TopCandidates(1)
		if len(candidates) == 0 {
			continue
		}
		c := candidates[0]
		bb := textObs.BoundingBox()

		// Vision bbox: normalized (0-1), bottom-left origin.
		// Convert to pixels with top-left origin.
		results = append(results, ocrText{
			Text:       c.String(),
			Confidence: float64(c.Confidence()),
			X:          bb.Origin.X * imgW,
			Y:          (1.0 - bb.Origin.Y - bb.Size.Height) * imgH,
			W:          bb.Size.Width * imgW,
			H:          bb.Size.Height * imgH,
		})
	}
	return results
}

// --- Window enumeration ---

type windowInfo struct {
	Name      string
	OwnerName string
	OwnerPID  int
	X, Y, W, H float64
}

const (
	kCFNumberSInt32Type  = 3
	kCFNumberFloat64Type = 12
	kCFStringEncodingUTF8 = uint32(0x08000100)
)

var (
	cfDictionaryGetValue func(dict uintptr, key unsafe.Pointer) unsafe.Pointer
	cfInitOnce           sync.Once
)

func ensureCFBindings() {
	cfInitOnce.Do(func() {
		lib, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			return
		}
		purego.RegisterLibFunc(&cfDictionaryGetValue, lib, "CFDictionaryGetValue")
	})
}

func getVisibleWindows() ([]windowInfo, error) {
	ensureCFBindings()
	if cfDictionaryGetValue == nil {
		return nil, fmt.Errorf("CFDictionaryGetValue not available")
	}

	listRef := coregraphics.CGWindowListCopyWindowInfo(
		coregraphics.KCGWindowListOptionOnScreenOnly|coregraphics.KCGWindowListExcludeDesktopElements,
		0, // kCGNullWindowID
	)
	if listRef == 0 {
		return nil, fmt.Errorf("CGWindowListCopyWindowInfo returned nil")
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(listRef))

	count := corefoundation.CFArrayGetCount(corefoundation.CFArrayRef(listRef))
	var windows []windowInfo
	for i := 0; i < count; i++ {
		dict := uintptr(corefoundation.CFArrayGetValueAtIndex(corefoundation.CFArrayRef(listRef), i))

		ownerName := cfDictGetString(dict, "kCGWindowOwnerName")
		name := cfDictGetString(dict, "kCGWindowName")
		pid := cfDictGetInt(dict, "kCGWindowOwnerPID")

		boundsRef := cfDictGet(dict, "kCGWindowBounds")
		x, y, w, h := parseBoundsDict(uintptr(boundsRef))

		if w == 0 || h == 0 {
			continue
		}

		windows = append(windows, windowInfo{
			Name: name, OwnerName: ownerName, OwnerPID: pid,
			X: x, Y: y, W: w, H: h,
		})
	}
	return windows, nil
}

func parseBoundsDict(dict uintptr) (x, y, w, h float64) {
	if dict == 0 {
		return
	}
	x = cfDictGetFloat(dict, "X")
	y = cfDictGetFloat(dict, "Y")
	w = cfDictGetFloat(dict, "Width")
	h = cfDictGetFloat(dict, "Height")
	return
}

func makeCFString(s string) unsafe.Pointer {
	return unsafe.Pointer(corefoundation.CFStringCreateWithCString(0, s, kCFStringEncodingUTF8))
}

func cfDictGet(dict uintptr, key string) unsafe.Pointer {
	cfKey := makeCFString(key)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cfKey))
	return cfDictionaryGetValue(dict, cfKey)
}

func cfDictGetString(dict uintptr, key string) string {
	val := cfDictGet(dict, key)
	if val == nil {
		return ""
	}
	ref := corefoundation.CFStringRef(uintptr(val))
	length := corefoundation.CFStringGetLength(ref)
	if length == 0 {
		return ""
	}
	ptr := corefoundation.CFStringGetCStringPtr(ref, kCFStringEncodingUTF8)
	if ptr != nil {
		return cstrToGo(uintptr(unsafe.Pointer(ptr)))
	}
	bufSize := length*4 + 1
	buf := make([]byte, bufSize)
	if corefoundation.CFStringGetCString(ref, &buf[0], bufSize, kCFStringEncodingUTF8) {
		for i, b := range buf {
			if b == 0 {
				return string(buf[:i])
			}
		}
	}
	return ""
}

func cfDictGetInt(dict uintptr, key string) int {
	val := cfDictGet(dict, key)
	if val == nil {
		return 0
	}
	var v int32
	corefoundation.CFNumberGetValue(corefoundation.CFNumberRef(uintptr(val)), corefoundation.CFNumberType(kCFNumberSInt32Type), unsafe.Pointer(&v))
	return int(v)
}

func cfDictGetFloat(dict uintptr, key string) float64 {
	val := cfDictGet(dict, key)
	if val == nil {
		return 0
	}
	var v float64
	corefoundation.CFNumberGetValue(corefoundation.CFNumberRef(uintptr(val)), corefoundation.CFNumberType(kCFNumberFloat64Type), unsafe.Pointer(&v))
	return v
}

func cstrToGo(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var bs []byte
	for {
		b := *(*byte)(unsafe.Pointer(ptr))
		if b == 0 {
			break
		}
		bs = append(bs, b)
		ptr++
	}
	return string(bs)
}

// --- XML output ---

type screenXML struct {
	XMLName xml.Name `xml:"screen"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Apps    []appXML `xml:"app"`
}

type appXML struct {
	Name   string    `xml:"name,attr"`
	PID    int       `xml:"pid,attr"`
	Bounds string    `xml:"bounds,attr"`
	Texts  []textXML `xml:"text"`
}

type textXML struct {
	Confidence float64 `xml:"confidence,attr"`
	Bounds     string  `xml:"bounds,attr"`
	Text       string  `xml:",chardata"`
}

func buildScreen(imgW, imgH int, windows []windowInfo, texts []ocrText) screenXML {
	type appKey struct {
		name string
		pid  int
	}
	appIndex := map[appKey]*appXML{}
	var appOrder []appKey

	for _, w := range windows {
		k := appKey{w.OwnerName, w.OwnerPID}
		if _, ok := appIndex[k]; !ok {
			appIndex[k] = &appXML{
				Name:   w.OwnerName,
				PID:    w.OwnerPID,
				Bounds: formatBounds(w.X, w.Y, w.W, w.H),
			}
			appOrder = append(appOrder, k)
		}
	}

	for _, t := range texts {
		cx := t.X + t.W/2
		cy := t.Y + t.H/2
		best := -1
		bestArea := math.MaxFloat64
		for i, w := range windows {
			if cx >= w.X && cx <= w.X+w.W && cy >= w.Y && cy <= w.Y+w.H {
				area := w.W * w.H
				if area < bestArea {
					bestArea = area
					best = i
				}
			}
		}
		txt := textXML{
			Confidence: math.Round(t.Confidence*100) / 100,
			Bounds:     formatBounds(t.X, t.Y, t.W, t.H),
			Text:       t.Text,
		}
		if best >= 0 {
			w := windows[best]
			k := appKey{w.OwnerName, w.OwnerPID}
			appIndex[k].Texts = append(appIndex[k].Texts, txt)
		} else {
			k := appKey{"(unknown)", 0}
			if _, ok := appIndex[k]; !ok {
				appIndex[k] = &appXML{Name: "(unknown)", PID: 0, Bounds: "0,0,0,0"}
				appOrder = append(appOrder, k)
			}
			appIndex[k].Texts = append(appIndex[k].Texts, txt)
		}
	}

	var apps []appXML
	for _, k := range appOrder {
		a := appIndex[k]
		if len(a.Texts) > 0 {
			apps = append(apps, *a)
		}
	}
	sort.Slice(apps, func(i, j int) bool { return apps[i].PID < apps[j].PID })

	return screenXML{Width: imgW, Height: imgH, Apps: apps}
}

func formatBounds(x, y, w, h float64) string {
	return fmt.Sprintf("%.0f,%.0f,%.0f,%.0f", x, y, w, h)
}

func writeXML(w io.Writer, s screenXML) error {
	fmt.Fprintln(w, xml.Header)
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(s); err != nil {
		return err
	}
	fmt.Fprintln(w)
	return nil
}
