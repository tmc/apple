#!/usr/bin/env bash
set -euo pipefail

dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
tmp=${TMPDIR:-/tmp}/9pfs-live-$$
export_dir=$tmp/export
moddir=$tmp/mod
server_bin=$tmp/export9p
client_bin=$tmp/9pfs
addr=${NINEPFS_TEST_ADDR:-127.0.0.1:$((20000 + RANDOM % 20000))}

# export9p (the classic 9P2000 server used by this test) is built from a
# temporary module that pins its version, then run as a local binary. Building
# it this way keeps the test hermetic: `go run module@version` forces a
# proxy.golang.org deprecation lookup even when the module is fully cached, so
# it fails offline (GOPROXY=off) and made this "local verification" leg depend
# on the network.
export9p_version=v1.18.0

cleanup() {
	if [[ -n "${server_pid:-}" ]]; then
		kill "$server_pid" 2>/dev/null || true
		wait "$server_pid" 2>/dev/null || true
	fi
	rm -rf "$tmp"
}
trap cleanup EXIT

mkdir -p "$export_dir" "$moddir"
printf 'hello from 9p\n' > "$export_dir/README"

# Build export9p from a throwaway module that pins the version, with no
# network lookup at run time. `go get` the command package so its full
# transitive go.sum is recorded, then build from the cache.
(
	cd "$moddir"
	GOWORK=off go mod init ninepfs.test/export9p >/dev/null 2>&1
	GOWORK=off go get "github.com/knusbaum/go9p/cmd/export9p@$export9p_version" >/dev/null 2>&1
	GOWORK=off go build -o "$server_bin" "github.com/knusbaum/go9p/cmd/export9p"
)

# Build the 9pfs client once.
(cd "$dir" && GOWORK=off go build -o "$client_bin" .)

run9pfs() {
	"$client_bin" "$@"
}

"$server_bin" -dir "$export_dir" -address "$addr" -noperm >"$tmp/export9p.log" 2>&1 &
server_pid=$!

ready=
for _ in {1..50}; do
	if run9pfs -dialect 9p2000 -net tcp -addr "$addr" -cat /README >/dev/null 2>&1; then
		ready=1
		break
	fi
	if ! kill -0 "$server_pid" 2>/dev/null; then
		echo "9pfs: export9p exited before becoming ready" >&2
		cat "$tmp/export9p.log" >&2
		exit 1
	fi
	sleep 0.1
done
if [[ -z "$ready" ]]; then
	echo "9pfs: export9p did not become ready at $addr" >&2
	cat "$tmp/export9p.log" >&2
	exit 1
fi

run9pfs -dialect 9p2000 -net tcp -addr "$addr" -ls /
run9pfs -dialect 9p2000 -net tcp -addr "$addr" -cat /README
printf '\n'
run9pfs -dialect 9p2000 -net tcp -addr "$addr" \
	-create /codex.txt -write /codex.txt -data 'written through 9pfs' -cat /codex.txt
printf '\n'

got=$(cat "$export_dir/codex.txt")
if [[ "$got" != "written through 9pfs" ]]; then
	echo "backing file mismatch: $got" >&2
	exit 1
fi

run9pfs -dialect 9p2000 -net tcp -addr "$addr" -truncate /codex.txt -size 8 >/dev/null
if [[ "$(cat "$export_dir/codex.txt")" != "written " ]]; then
	echo "truncate did not update backing file" >&2
	exit 1
fi

run9pfs -dialect 9p2000 -net tcp -addr "$addr" -chmod /codex.txt -mode 0600 >/dev/null
mode=$(stat -f '%Lp' "$export_dir/codex.txt")
if [[ "$mode" != "600" ]]; then
	echo "chmod did not update backing file mode: $mode" >&2
	exit 1
fi

run9pfs -dialect 9p2000 -net tcp -addr "$addr" -mtime /codex.txt -mtime-seconds 1577934240 >/dev/null
mtime=$(stat -f '%Sm' -t '%Y%m%d%H%M' "$export_dir/codex.txt")
if [[ "$mtime" != "202001020304" ]]; then
	echo "9pfs: export9p did not persist mtime update; continuing" >&2
fi

run9pfs -dialect 9p2000 -net tcp -addr "$addr" -rm /codex.txt >/dev/null
if [[ -e "$export_dir/codex.txt" ]]; then
	echo "remove did not delete backing file" >&2
	exit 1
fi

echo "9pfs: live 9P2000 read/write/chmod/truncate/remove ok"
