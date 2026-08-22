// Command codecs reports the video encoders VideoToolbox exposes on this
// machine, the codecs it can decode in hardware, and the properties an encoder
// supports for a given frame size.
//
// Usage: codecs [-codec fourcc] [-width n] [-height n]
//
// The -codec, -width and -height flags select the encoder whose supported
// property dictionary is printed; the default is avc1 at 1920x1080.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/videotoolbox"
)

func main() {
	codec := flag.String("codec", "avc1", "four-character code of the codec to inspect")
	width := flag.Int("width", 1920, "frame width used when querying encoder properties")
	height := flag.Int("height", 1080, "frame height used when querying encoder properties")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: codecs [-codec fourcc] [-width n] [-height n]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(*codec) != 4 {
		fmt.Fprintf(os.Stderr, "codecs: -codec must be exactly four characters, got %q\n", *codec)
		os.Exit(1)
	}
	codecType := fourCC(*codec)

	if err := listEncoders(); err != nil {
		fmt.Fprintf(os.Stderr, "codecs: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	reportHardwareDecode()
	fmt.Println()
	if err := listEncoderProperties(codecType, int32(*width), int32(*height)); err != nil {
		fmt.Fprintf(os.Stderr, "codecs: %v\n", err)
		os.Exit(1)
	}
}

// listEncoders prints one line per encoder VideoToolbox advertises.
func listEncoders() error {
	var list corefoundation.CFArrayRef
	if status := videotoolbox.VTCopyVideoEncoderList(0, &list); status != 0 {
		return fmt.Errorf("VTCopyVideoEncoderList: OSStatus %d", status)
	}
	if list == 0 {
		return fmt.Errorf("VTCopyVideoEncoderList returned no list")
	}
	defer cfRelease(objc.ID(list))

	n := corefoundation.CFArrayGetCount(list)
	fmt.Printf("video encoders (%d):\n", n)
	for i := 0; i < n; i++ {
		enc := objc.ID(uintptr(corefoundation.CFArrayGetValueAtIndex(list, i)))
		if enc == 0 {
			continue
		}
		name := dictString(enc, videotoolbox.KVTVideoEncoderList_EncoderName)
		if name == "" {
			name = dictString(enc, videotoolbox.KVTVideoEncoderList_DisplayName)
		}
		fmt.Printf("  %-40s %s\n", name, dictString(enc, videotoolbox.KVTVideoEncoderList_EncoderID))
		fmt.Printf("      codec:           %s (%s)\n",
			dictString(enc, videotoolbox.KVTVideoEncoderList_CodecName),
			codecString(enc))
		fmt.Printf("      hardware:        %s\n", dictBool(enc, videotoolbox.KVTVideoEncoderList_IsHardwareAccelerated))
		fmt.Printf("      frame reorder:   %s\n", dictBool(enc, videotoolbox.KVTVideoEncoderList_SupportsFrameReordering))
		fmt.Printf("      multi pass:      %s\n", dictBool(enc, videotoolbox.KVTVideoEncoderList_SupportsMultiPass))
		if v := dictString(enc, videotoolbox.KVTVideoEncoderList_InstanceLimit); v != "" {
			fmt.Printf("      instance limit:  %s\n", v)
		}
		if v := dictString(enc, videotoolbox.KVTVideoEncoderList_GPURegistryID); v != "" {
			fmt.Printf("      gpu registry id: %s\n", v)
		}
	}
	return nil
}

// reportHardwareDecode prints hardware decode support for the common codecs.
func reportHardwareDecode() {
	codecs := []struct {
		name string
		typ  coremedia.KCMVideoCodecType
	}{
		{"H.264", coremedia.KCMVideoCodecType_H264},
		{"HEVC", coremedia.KCMVideoCodecType_HEVC},
		{"AV1", coremedia.KCMVideoCodecType_AV1},
		{"VP9", coremedia.KCMVideoCodecType_VP9},
		{"JPEG", coremedia.KCMVideoCodecType_JPEG},
		{"ProRes 422", coremedia.KCMVideoCodecType_AppleProRes422},
		{"ProRes 4444", coremedia.KCMVideoCodecType_AppleProRes4444},
	}
	fmt.Println("hardware decode support:")
	for _, c := range codecs {
		fmt.Printf("  %-12s %-6s %v\n", c.name, fourCCString(uint32(c.typ)),
			videotoolbox.VTIsHardwareDecodeSupported(uint32(c.typ)))
	}
	fmt.Printf("  %-12s %-6s %v\n", "MV-HEVC", "decode", videotoolbox.VTIsStereoMVHEVCDecodeSupported())
	fmt.Printf("  %-12s %-6s %v\n", "MV-HEVC", "encode", videotoolbox.VTIsStereoMVHEVCEncodeSupported())
}

// listEncoderProperties prints the properties the encoder chosen for codecType
// at the given frame size supports.
func listEncoderProperties(codecType uint32, width, height int32) error {
	var encoderID corefoundation.CFStringRef
	var props corefoundation.CFDictionaryRef
	status := videotoolbox.VTCopySupportedPropertyDictionaryForEncoder(width, height, codecType, 0, &encoderID, &props)
	if status != 0 {
		return fmt.Errorf("VTCopySupportedPropertyDictionaryForEncoder(%s, %dx%d): OSStatus %d",
			fourCCString(codecType), width, height, status)
	}
	if encoderID != 0 {
		defer cfRelease(objc.ID(encoderID))
	}
	if props == 0 {
		return fmt.Errorf("no property dictionary for %s", fourCCString(codecType))
	}
	defer cfRelease(objc.ID(props))

	fmt.Printf("encoder for %s at %dx%d: %s\n", fourCCString(codecType), width, height,
		describe(objc.ID(uintptr(encoderID))))
	keys := dictKeys(objc.ID(uintptr(props)))
	sort.Strings(keys)
	fmt.Printf("supported properties (%d):\n", len(keys))
	for _, k := range keys {
		fmt.Printf("  %s\n", k)
	}
	return nil
}

// cfRelease releases a Core Foundation object. Core Foundation refs are
// Objective-C objects, and the generated CFRelease takes an unsafe.Pointer
// while the ref types are uintptr, so -release is the cleaner spelling here.
func cfRelease(id objc.ID) {
	if id != 0 {
		objc.Send[struct{}](id, objc.Sel("release"))
	}
}

// dictKeys returns the keys of a CFDictionary, which is toll-free bridged to
// NSDictionary, as Go strings.
func dictKeys(dict objc.ID) []string {
	all := objc.Send[objc.ID](dict, objc.Sel("allKeys"))
	if all == 0 {
		return nil
	}
	n := int(objc.Send[uint](all, objc.Sel("count")))
	keys := make([]string, 0, n)
	for i := 0; i < n; i++ {
		k := objc.Send[objc.ID](all, objc.Sel("objectAtIndex:"), uint(i))
		keys = append(keys, describe(k))
	}
	return keys
}

// dictString looks up key in a CFDictionary and returns a printable form of the
// value, or "" if the key is absent.
func dictString(dict objc.ID, key string) string {
	if key == "" {
		return ""
	}
	v := objc.Send[objc.ID](dict, objc.Sel("objectForKey:"), objc.String(key))
	if v == 0 {
		return ""
	}
	return describe(v)
}

// dictBool reports a boolean encoder attribute. VideoToolbox omits these keys
// rather than storing false, so an absent key reads as "no".
func dictBool(dict objc.ID, key string) string {
	if key == "" {
		return "unknown"
	}
	v := objc.Send[objc.ID](dict, objc.Sel("objectForKey:"), objc.String(key))
	if v == 0 {
		return "no"
	}
	if objc.Send[bool](v, objc.Sel("boolValue")) {
		return "yes"
	}
	return "no"
}

// codecString renders an encoder entry's codec type as a four-character code.
func codecString(enc objc.ID) string {
	key := videotoolbox.KVTVideoEncoderList_CodecType
	if key == "" {
		return "?"
	}
	v := objc.Send[objc.ID](enc, objc.Sel("objectForKey:"), objc.String(key))
	if v == 0 {
		return "?"
	}
	return fourCCString(uint32(objc.Send[int64](v, objc.Sel("longLongValue"))))
}

// describe returns -[NSObject description] as a Go string.
func describe(id objc.ID) string {
	if id == 0 {
		return ""
	}
	desc := objc.Send[objc.ID](id, objc.Sel("description"))
	if desc == 0 {
		return ""
	}
	cstr := objc.Send[*byte](desc, objc.Sel("UTF8String"))
	if cstr == nil {
		return ""
	}
	return objc.GoString(cstr)
}

func fourCC(s string) uint32 {
	return uint32(s[0])<<24 | uint32(s[1])<<16 | uint32(s[2])<<8 | uint32(s[3])
}

func fourCCString(v uint32) string {
	b := []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
	for i, c := range b {
		if c < 0x20 || c > 0x7e {
			b[i] = '.'
		}
	}
	return string(b)
}
