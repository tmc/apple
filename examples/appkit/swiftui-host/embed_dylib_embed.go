//go:build swiftui_embed

package main

import _ "embed"

//go:embed ui/libExampleAppUI.dylib
var exampleAppUIDylib []byte

func embeddedExampleAppUI() []byte {
	return exampleAppUIDylib
}
