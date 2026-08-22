//go:generate -command applegen-bootstrap go run github.com/tmc/appledocs/cmd/applegen bootstrap --module github.com/tmc/apple --force --verbose

//go:generate applegen-bootstrap --private -f "appleneuralengine"
//go:generate applegen-bootstrap --private -f "diskimages2"
//go:generate applegen-bootstrap --private -f "espresso"
//go:generate applegen-bootstrap --private -f "remotecoreml"
//go:generate applegen-bootstrap --private -f "speechobjects"
//go:generate applegen-bootstrap --private -f "texttospeech"
//go:generate applegen-bootstrap --private -f "skylight"

//go:generate applegen-bootstrap --private-companion -f "avfaudio"
//go:generate applegen-bootstrap --private-companion -f "coreml"
//go:generate applegen-bootstrap --private-companion -f "network" --go-names
//go:generate applegen-bootstrap --private-companion -f "virtualization"

//go:generate applegen-bootstrap --private --extra-flags=--framework-path=/System/Library/Frameworks/ApplicationServices.framework/Versions/A/Frameworks/HIServices.framework/Versions/A/HIServices HIServices

// Package private holds bindings to Apple API that carries no support
// contract: symbols that are undocumented, may change or disappear between
// OS releases, and will not pass App Store review.
//
// The directory groups by that contract, not by framework location.
// It contains two kinds of packages:
//
//   - Companion packages (private/coreml, private/network,
//     private/avfaudio, private/virtualization) extend a public framework
//     that also has documented bindings at the repository root. All four
//     frameworks live in /System/Library/Frameworks; what is private is
//     the selectors, not the framework.
//   - Standalone packages (private/skylight, private/hiservices,
//     private/espresso, private/appleneuralengine, and others) bind
//     frameworks with no documented bindings of their own in this
//     repository.
//
// A companion package declares the same package name as its public twin,
// so importing both requires a local name; examples use priv<framework>:
//
//	import (
//		"github.com/tmc/apple/coreml"
//		privcoreml "github.com/tmc/apple/private/coreml"
//	)
//
// Importing anything under private/ is the signal that a program depends
// on unsupported API; auditing for SPI use is a walk of the import graph.
package private
