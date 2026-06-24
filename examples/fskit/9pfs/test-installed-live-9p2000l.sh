#!/usr/bin/env bash
set -euo pipefail

dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
source "$dir/scriptlib.sh"

require_no_active_9pfs_mounts

tmp=$(mktemp -d "${TMPDIR:-/tmp}/9pfs-installed-live.XXXXXX")
root=$tmp/export
log=$tmp/p9ufs.log
p9src=$tmp/p9-src
server_bin=$tmp/p9ufs
port=$((20000 + RANDOM % 20000))
addr=127.0.0.1:$port
mnt=${1:-$tmp/mnt}

cleanup() {
	set +e
	if [[ -n "${mnt:-}" ]] && mount | awk '{print $3}' | grep -qx "$mnt"; then
		umount "$mnt" 2>/dev/null || diskutil unmount force "$mnt" 2>/dev/null || true
	fi
	if [[ -n "${server_pid:-}" ]]; then
		kill "$server_pid" 2>/dev/null || true
		wait "$server_pid" 2>/dev/null || true
	fi
	rm -rf "$tmp"
}
trap cleanup EXIT

mkdir -p "$root" "$mnt"
printf 'hello from mounted 9p2000.l\n' > "$root/README"

"$dir/preflight-installed.sh"

"$dir/prepare-p9-module.sh" "$p9src"
(cd "$p9src" && GOWORK=off go build -o "$server_bin" ./cmd/p9ufs)

("$server_bin" -root "$root" "$addr" >"$log" 2>&1) &
server_pid=$!

ready=
for _ in $(seq 1 300); do
	if nc -z 127.0.0.1 "$port" >/dev/null 2>&1; then
		ready=1
		break
	fi
	if ! kill -0 "$server_pid" 2>/dev/null; then
		echo "9pfs: p9ufs exited before becoming ready" >&2
		cat "$log" >&2
		exit 1
	fi
	sleep 0.1
done
if [[ -z "$ready" ]]; then
	echo "9pfs: p9ufs did not become ready at $addr" >&2
	cat "$log" >&2
	exit 1
fi

url="ninep://$addr?dialect=9p2000l"
echo "9pfs: testing installed FSKit mount against $url"
"$dir/test-installed-mount.sh" "$url" "$mnt"
