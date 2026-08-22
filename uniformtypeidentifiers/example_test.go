//go:build darwin

package uniformtypeidentifiers_test

import (
	"fmt"

	"github.com/tmc/apple/uniformtypeidentifiers"
)

func ExampleUTType() {
	pngType := uniformtypeidentifiers.NewTypeWithFilenameExtension("png")
	fmt.Println(pngType.Identifier())
	fmt.Println(pngType.PreferredFilenameExtension())
	fmt.Println(pngType.PreferredMIMEType())

	// Output:
	// public.png
	// png
	// image/png
}

func ExampleUTType_ConformsToType() {
	pngType := uniformtypeidentifiers.NewTypeWithFilenameExtension("png")
	imageType := uniformtypeidentifiers.NewTypeWithIdentifier("public.image")
	audioType := uniformtypeidentifiers.NewTypeWithIdentifier("public.audio")

	if pngType.ConformsToType(imageType) {
		fmt.Printf("%s conforms to %s\n", pngType.Identifier(), imageType.Identifier())
	}
	if !pngType.ConformsToType(audioType) {
		fmt.Printf("%s does not conform to %s\n", pngType.Identifier(), audioType.Identifier())
	}

	// Output:
	// public.png conforms to public.image
	// public.png does not conform to public.audio
}
