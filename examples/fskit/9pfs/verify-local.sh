#!/usr/bin/env bash
set -euo pipefail

dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
build_dir=${1:-/tmp/9pfs-verify-local-build}

cd "$dir"

# Shell and plist lint.
/bin/bash -n \
	build-appex.sh \
	install-local.sh \
	preflight-installed.sh \
	prepare-p9-module.sh \
	test-installed-live-9p2000l.sh \
	test-installed-mount.sh \
	test-live-9p2000.sh \
	test-live-9p2000l.sh \
	show-live-mount.sh \
	verify-signed-build.sh \
	verify-local.sh

for plist in appex/Info.plist fsbundle/Info.plist host/Info.plist; do
	plutil -lint "$plist"
done

# Go tests and the in-memory FSKit callback smoke path.
GOWORK=off go test ./... -count=1
GOWORK=off go run . -fskit-smoke

# Swift entrypoint shell and Go operation attachment.
(cd swiftshim && swift build -c release --product NinePFSShim -Xcc -DNINEPFS_OBJC_CONTROL=1)
GOWORK=off go run . -extension-main-probe

# Live 9P client operations against disposable servers.
./test-live-9p2000.sh
./test-live-9p2000l.sh

# Default extension/bundle assembly.
./build-appex.sh "$build_dir"

echo "9pfs: local verification ok"
