#!/bin/bash
# run.sh builds both benchmark binaries and runs the client/server matrix,
# then prints one comparison table.
#
# Crossing the implementations is the point: go-nw client against a swift
# server measures the Go bindings' send/receive path against the same
# framework server, and swift client against a go-nw server measures the
# bindings' server path. The std/std row is the socket baseline.
#
# Each pairing runs -repeat times and reports its median, and every result
# JSON is kept: it carries the configuration and the machine it was taken on,
# so a saved run stays interpretable later. Set NETPERFBENCH_OUT to keep the
# results directory somewhere permanent.
#
# Usage: ./run.sh [payload_bytes [round_trips [repetitions]]]

set -euo pipefail

size=${1:-4096}
count=${2:-20000}
reps=${3:-3}
dir=$(cd "$(dirname "$0")" && pwd)
results=${NETPERFBENCH_OUT:-}
if [ -n "$results" ]; then
	mkdir -p "$results"
fi
out=$(mktemp -d "${TMPDIR:-$HOME/tmp}/netperfbench.XXXXXX")
server_pid=""
cleanup() {
	[ -n "$server_pid" ] && kill "$server_pid" 2>/dev/null
	rm -rf "$out"
}
trap cleanup EXIT

echo "building..." >&2
go build -o "$out/netperfbench" "$dir"
swiftc -O "$dir/swift/netperfbench.swift" -o "$out/netperfbench-swift"

port=51000
json_results=()

# run SERVER_KIND CLIENT_KIND
run() {
	local server=$1 client=$2 label="$1 server / $2 client"
	port=$((port + 1))

	case $server in
	swift) "$out/netperfbench-swift" -role server -port $port & ;;
	*) "$out/netperfbench" -role server -impl "$server" -port $port & ;;
	esac
	server_pid=$!
	sleep 0.5

	local json
	case $client in
	swift) json=$("$out/netperfbench-swift" -role client -addr "127.0.0.1:$port" -size "$size" -n "$count" -repeat "$reps" -label "$label" -json) ;;
	*) json=$("$out/netperfbench" -role client -impl "$client" -addr "127.0.0.1:$port" -size "$size" -n "$count" -repeat "$reps" -label "$label" -json) ;;
	esac
	kill "$server_pid" 2>/dev/null || true
	wait "$server_pid" 2>/dev/null || true
	server_pid=""
	json_results+=("$json")
	if [ -n "$results" ]; then
		echo "$json" >"$results/$server-server-$client-client.json"
	fi
}

run std std
run nw nw
run swift swift
run swift nw
run nw swift

printf '\n%d round trips of %d bytes each, median of %d repetitions\n\n' "$count" "$size" "$reps"
printf '%-26s %10s %10s %10s %12s %10s\n' "pairing" "p50 us" "p90 us" "p99 us" "trips/sec" "MB/s"
printf '%-26s %10s %10s %10s %12s %10s\n' "--------------------------" "----------" "----------" "----------" "------------" "----------"
for json in "${json_results[@]}"; do
	echo "$json" | python3 -c '
import json, sys
r = json.load(sys.stdin)
print("%-26s %10.1f %10.1f %10.1f %12.0f %10.1f" % (
    r["label"], r["p50_us"], r["p90_us"], r["p99_us"],
    r["round_trips_per_sec"], r["throughput_mbps"]))'
done
