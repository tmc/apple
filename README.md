# apple

[![Go Reference](https://pkg.go.dev/badge/github.com/tmc/apple.svg)](https://pkg.go.dev/github.com/tmc/apple)

Go bindings for Apple platform frameworks. These bindings are cgo-free, using
[purego](https://github.com/ebitengine/purego) for dynamic function loading at
runtime.

Auto-generated from Apple developer documentation via
[applegen](https://github.com/tmc/appledocs/cmd/applegen).

## Private frameworks

Bindings under the `private/` directory wrap undocumented Apple frameworks.
These APIs are unstable, may change or disappear between OS releases, and carry
no compatibility guarantees.

## Disclaimer

This is not an official Apple product. Apple, macOS, and all related
frameworks are trademarks of Apple Inc. This project is an independent,
community-driven effort and is not affiliated with, endorsed by, or sponsored
by Apple Inc.

## License

[MIT](LICENSE)
