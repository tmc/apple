// Package apple is a collection of Go packages for Apple platform frameworks.
//
// These bindings are cgo-free, using purego for dynamic function loading at
// runtime. They are auto-generated from Apple developer documentation via
// applegen (github.com/tmc/appledocs/cmd/applegen).

//go:generate -command applegen-generate applegen generate --output . --module github.com/tmc/apple
//go:generate -command applegen-bootstrap applegen bootstrap --module github.com/tmc/apple --force --verbose

//go:generate applegen-generate --generate-objc-runtime

//go:generate applegen-bootstrap -f "accelerate"
//go:generate applegen-bootstrap -f "appkit"
//go:generate applegen-bootstrap -f "corefoundation"
//go:generate applegen-bootstrap -f "coregraphics"
//go:generate applegen-bootstrap -f "coreml"
//go:generate applegen-bootstrap -f "corevideo"
//go:generate applegen-bootstrap -f "dispatch"
//go:generate applegen-bootstrap -f "foundation"
//go:generate applegen-bootstrap -f "iosurface"
//go:generate applegen-bootstrap -f "metal"
//go:generate applegen-bootstrap -f "metalkit"
//go:generate applegen-bootstrap -f "objectivec"
//go:generate applegen-bootstrap -f "quartzcore"
//go:generate applegen-bootstrap -f "screencapturekit"
//go:generate applegen-bootstrap -f "security"
//go:generate applegen-bootstrap -f "symbols"
//go:generate applegen-bootstrap -f "uniformtypeidentifiers"
//go:generate applegen-bootstrap -f "virtualization"
//go:generate applegen-bootstrap -f "vision"
//go:generate applegen-bootstrap -f "vmnet"

package apple
