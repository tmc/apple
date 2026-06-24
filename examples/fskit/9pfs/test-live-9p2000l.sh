#!/usr/bin/env bash
set -euo pipefail

dir=$(cd "$(dirname "$0")" && pwd)
root=$(mktemp -d "${TMPDIR:-/tmp}/9pfs-9p2000l-root.XXXXXX")
log=$(mktemp "${TMPDIR:-/tmp}/9pfs-9p2000l-server.XXXXXX")
p9src=$(mktemp -d "${TMPDIR:-/tmp}/9pfs-p9-src.XXXXXX")
moddir=$(mktemp -d "${TMPDIR:-/tmp}/9pfs-mod.XXXXXX")
p9src=$(cd "$p9src" && pwd -P)
moddir=$(cd "$moddir" && pwd -P)
modfile=$moddir/go.mod
server_bin=$moddir/p9ufs
client_bin=$moddir/9pfs
port=$((20000 + RANDOM % 20000))
addr=127.0.0.1:$port

cleanup() {
	set +e
	if [[ -n "${server_pid:-}" ]]; then
		kill "$server_pid" 2>/dev/null || true
		wait "$server_pid" 2>/dev/null || true
	fi
	rm -rf "$root"
	rm -rf "$p9src"
	rm -f "$log"
	rm -rf "$moddir"
}
trap cleanup EXIT

printf 'hello from 9p2000.l\n' > "$root/README"

"$dir/prepare-p9-module.sh" "$p9src"

sed '/^replace github.com\/tmc\/apple /d' "$dir/go.mod" > "$modfile"
cp "$dir/go.sum" "${modfile%.mod}.sum"
{
	echo
	echo "replace github.com/tmc/apple => $(cd "$dir/../../.." && pwd)"
	echo "replace github.com/hugelgupf/p9 => $p9src"
} >> "$modfile"

run9pfs() {
	"$client_bin" "$@"
}

(cd "$p9src" && GOWORK=off go build -o "$server_bin" ./cmd/p9ufs)
(cd "$dir" && GOWORK=off GOFLAGS="-modfile=$modfile" go build -o "$client_bin" .)

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
	if kill -0 "$server_pid" 2>/dev/null; then
		echo "9pfs: p9ufs process is still running; startup may be stuck compiling or blocked before listen" >&2
	fi
	cat "$log" >&2
	exit 1
fi

cd "$dir"

run9pfs -dialect 9p2000l -net tcp -addr "$addr" -ls /
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -cat /README

run9pfs -dialect 9p2000l -net tcp -addr "$addr" \
	-create /codex.txt \
	-write /codex.txt \
	-data 'written through 9p2000.l'
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -cat /codex.txt
echo

run9pfs -dialect 9p2000l -net tcp -addr "$addr" -mkdir /dir
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rename /codex.txt -to /dir/renamed.txt
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -truncate /dir/renamed.txt -size 8
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -cat /dir/renamed.txt | grep -qx 'written '
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -chmod /dir/renamed.txt -mode 0600
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -stat /dir/renamed.txt | grep -Eq '^- /?renamed.txt 8 [0-9]+ 0600$'
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -mtime /dir/renamed.txt -mtime-seconds 1577934240
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -stat /dir/renamed.txt | grep -Eq '^- /?renamed.txt 8 1577934240 0600$'
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -symlink /dir/link.txt -target renamed.txt
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -readlink /dir/link.txt | grep -qx 'renamed.txt'
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -link /dir/renamed.txt -link-to /dir/hard.txt

xattr_err=$(mktemp "${TMPDIR:-/tmp}/9pfs-9p2000l-xattr.XXXXXX")
if run9pfs -dialect 9p2000l -net tcp -addr "$addr" \
	-setxattr /dir/renamed.txt \
	-xattr user.codex \
	-xattr-data 'xattr through 9p2000.l' 2>"$xattr_err"; then
	run9pfs -dialect 9p2000l -net tcp -addr "$addr" -listxattr /dir/renamed.txt | grep -qx 'user.codex'
	run9pfs -dialect 9p2000l -net tcp -addr "$addr" -getxattr /dir/renamed.txt -xattr user.codex | grep -qx 'xattr through 9p2000.l'
	run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rmxattr /dir/renamed.txt -xattr user.codex
	if run9pfs -dialect 9p2000l -net tcp -addr "$addr" -listxattr /dir/renamed.txt | grep -q 'user.codex'; then
		echo "9pfs: xattr still present after removal" >&2
		exit 1
	fi
else
	echo "9pfs: live 9P2000.L xattr set failed" >&2
	sed -n '1p' "$xattr_err"
	exit 1
fi
rm -f "$xattr_err"

run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rm /dir/hard.txt
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rm /dir/link.txt
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rm /dir/renamed.txt
run9pfs -dialect 9p2000l -net tcp -addr "$addr" -rm /dir

echo "9pfs: live 9P2000.L read/write/rename/chmod/mtime/truncate/link/symlink/xattr/remove ok"
