// Pbd is a pasteboard tool, an improvement over pbcopy/pbpaste.
//
// Usage:
//
//	echo "hello" | pbd copy
//	pbd paste
//	pbd paste -type public.html
//	pbd types
//	pbd -j paste
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/tmc/apple/appkit"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	jsonOut := flag.Bool("j", false, "JSON output")
	typeFlag := flag.String("type", "public.utf8-plain-text", "pasteboard type (UTI)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pbd [-j] [-type <uti>] <copy|paste|types>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}

	pb := appkit.GetNSPasteboardClass().GeneralPasteboard()

	switch flag.Arg(0) {
	case "copy":
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pbd: read stdin: %v\n", err)
			os.Exit(2)
		}
		text := string(data)
		pb.ClearContents()
		pb.SetStringForType(text, *typeFlag)
		if *jsonOut {
			json.NewEncoder(os.Stdout).Encode(map[string]any{
				"copied": true,
				"bytes":  len(data),
			})
		}
	case "paste":
		content := pb.StringForType(*typeFlag)
		if *jsonOut {
			json.NewEncoder(os.Stdout).Encode(map[string]any{
				"type":    *typeFlag,
				"content": content,
			})
		} else {
			fmt.Print(content)
		}
	case "types":
		types := pb.Types()
		if *jsonOut {
			json.NewEncoder(os.Stdout).Encode(types)
		} else {
			for _, t := range types {
				fmt.Println(t)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "pbd: unknown command %q\n", flag.Arg(0))
		flag.Usage()
		os.Exit(2)
	}
}
