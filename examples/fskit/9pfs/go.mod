module github.com/tmc/apple/examples/fskit/9pfs

go 1.25.0

require (
	9fans.net/go v0.0.7
	github.com/ebitengine/purego v0.10.1
	github.com/tmc/apple v0.0.0
)

require (
	github.com/hugelgupf/p9 v0.4.0
	github.com/u-root/uio v0.0.0-20230305220412-3e8cd9d6bf63 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace github.com/tmc/apple => ../../..
