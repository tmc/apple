//go:generate -command applegen-bootstrap applegen bootstrap --module github.com/tmc/apple --force --verbose

//go:generate applegen-bootstrap --private -f "appleneuralengine"
//go:generate applegen-bootstrap --private -f "diskimages2"
//go:generate applegen-bootstrap --private -f "espresso"
//go:generate applegen-bootstrap --private -f "mlruntime"
//go:generate applegen-bootstrap --private -f "remotecoreml"

// Package private contains the generated code for the private frameworks of Apple MacOS. These frameworks are not intended for public use and may change without notice. Use at your own risk.
package private
