#!/usr/bin/env bash
set -euo pipefail

dest=${1:?"usage: prepare-p9-module.sh dest"}

srcdir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
modcache=$(GOWORK=off go env GOMODCACHE)
rm -rf "$dest"
mkdir -p "$dest"
cp -R "$modcache/github.com/hugelgupf/p9@v0.4.0/." "$dest/"
chmod -R u+w "$dest"

patch -d "$dest" -p1 < "$srcdir/p9-v0.4.0-9pfs.patch"
