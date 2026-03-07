// Command typeid resolves Uniform Type Identifiers for file extensions, MIME types, and UTI strings.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tmc/apple/uniformtypeidentifiers"
)

func main() {
	jsonOut := flag.Bool("j", false, "JSON output")
	mime := flag.Bool("mime", false, "lookup by MIME type")
	uti := flag.Bool("uti", false, "lookup by UTI identifier")
	conforms := flag.String("conforms", "", "check conformance to UTI type")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: typeid [-j] [-mime] [-uti] [-conforms type] arg ...")
		os.Exit(2)
	}

	var conformsTo uniformtypeidentifiers.UTType
	if *conforms != "" {
		conformsTo = uniformtypeidentifiers.NewTypeWithIdentifier(*conforms)
	}

	var results []map[string]any
	for _, arg := range args {
		var ut uniformtypeidentifiers.UTType
		switch {
		case *mime:
			ut = uniformtypeidentifiers.NewTypeWithMIMEType(arg)
		case *uti:
			ut = uniformtypeidentifiers.NewTypeWithIdentifier(arg)
		default:
			ext := strings.TrimPrefix(arg, ".")
			ut = uniformtypeidentifiers.NewTypeWithFilenameExtension(ext)
		}

		r := map[string]any{
			"identifier":  ut.Identifier(),
			"mime":        ut.PreferredMIMEType(),
			"extension":   ut.PreferredFilenameExtension(),
			"description": ut.LocalizedDescription(),
		}
		if *conforms != "" {
			r["conforms"] = ut.ConformsToType(conformsTo)
		}
		results = append(results, r)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		for _, r := range results {
			if err := enc.Encode(r); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}
		return
	}

	for i, r := range results {
		arg := args[i]
		if *conforms != "" {
			fmt.Printf("%s\tidentifier=%s\tconforms-to-%s=%v\n", arg, r["identifier"], *conforms, r["conforms"])
		} else {
			fmt.Printf("%s\tidentifier=%s\tmime=%s\text=%s\tdescription=%s\n",
				arg, r["identifier"], r["mime"], r["extension"], r["description"])
		}
	}
}
