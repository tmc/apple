package main

import (
	"fmt"

	"github.com/tmc/apple/uniformtypeidentifiers"
)

func main() {
	extensions := []string{"pdf", "jpg", "html", "swift"}
	for _, ext := range extensions {
		t := uniformtypeidentifiers.NewTypeWithFilenameExtension(ext)
		fmt.Printf("ext=%s identifier=%s mime=%s description=%s\n",
			ext, t.Identifier(), t.PreferredMIMEType(), t.LocalizedDescription())
	}

	// Check conformance: JPEG conforms to image.
	jpg := uniformtypeidentifiers.NewTypeWithFilenameExtension("jpg")
	image := uniformtypeidentifiers.NewTypeWithIdentifier("public.image")
	fmt.Printf("jpg-conforms-to-image=%t\n", jpg.ConformsToType(image))

	// Look up a type by MIME type.
	json := uniformtypeidentifiers.NewTypeWithMIMEType("application/json")
	fmt.Printf("json-identifier=%s ext=%s\n", json.Identifier(), json.PreferredFilenameExtension())
}
