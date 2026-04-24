//go:generate -command applegen-bootstrap applegen bootstrap --module github.com/tmc/apple --force --verbose

//go:generate applegen-bootstrap --private -f "appleneuralengine"
//go:generate applegen-bootstrap --private -f "diskimages2"
//go:generate applegen-bootstrap --private -f "espresso"
//go:generate applegen-bootstrap --private -f "remotecoreml"
//go:generate applegen-bootstrap --private -f "speechobjects"
//go:generate applegen-bootstrap --private -f "texttospeech"
//go:generate applegen-bootstrap --private -f "skylight"

//go:generate applegen-bootstrap --private-companion -f "avfaudio"
//go:generate applegen-bootstrap --private-companion -f "coreml"
//go:generate applegen-bootstrap --private-companion -f "virtualization"

//go:generate applegen-bootstrap --private --extra-flags=--framework-path=/System/Library/Frameworks/ApplicationServices.framework/Versions/A/Frameworks/HIServices.framework/Versions/A/HIServices HIServices

// Package private contains the generated code for the private frameworks of Apple MacOS. These frameworks are not intended for public use and may change without notice. Use at your own risk.
package private
