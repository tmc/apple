#!/bin/bash
# run.sh builds both halves of the Swift/Go XPC interoperability examples,
# installs each service as a LaunchAgent, drives every probe in both
# directions, and boots the jobs out again.
#
#	./examples/xpc/swiftinterop/run.sh
#
# Everything it creates is namespaced by the shell's pid, so concurrent runs
# and stale jobs cannot collide. The cleanup trap runs on any exit and asserts
# that no job is left behind.
set -u

here=$(cd "$(dirname "$0")" && pwd)
work=$(mktemp -d "${TMPDIR:-/tmp}/swiftinterop.XXXXXX")
agents="$HOME/Library/LaunchAgents"
domain="gui/$(id -u)"

swift_service="dev.tmc.apple.xpc.interop.swift.$$"
go_service="dev.tmc.apple.xpc.interop.go.$$"

# bootout races a launchd respawn of an on-demand job, so kill and boot out
# twice. A plist removed while its job is loaded leaves the job registered
# until the next login, so the plist goes last.
cleanup() {
	for label in "$swift_service" "$go_service"; do
		launchctl bootout "$domain/$label" >/dev/null 2>&1
		launchctl kill SIGKILL "$domain/$label" >/dev/null 2>&1
		launchctl bootout "$domain/$label" >/dev/null 2>&1
		rm -f "$agents/$label.plist"
	done
	echo
	echo "=== cleanup ==="
	local leaked=0
	for label in "$swift_service" "$go_service"; do
		if launchctl print "$domain/$label" >/dev/null 2>&1; then
			echo "LEAKED job $label"
			leaked=1
		fi
		if [ -e "$agents/$label.plist" ]; then
			echo "LEAKED plist $agents/$label.plist"
			leaked=1
		fi
	done
	[ "$leaked" = 0 ] && echo "no jobs or plists left behind"
	echo "service logs under $work (left in place for inspection)"
}
trap cleanup EXIT

echo "=== build ==="
go build -o "$work/goservice" "$here/goservice" || exit 1
go build -o "$work/goclient" "$here/goclient" || exit 1
# -swift-version 5 keeps the strict-concurrency diagnostics as warnings; the
# handlers here are @Sendable but the XPC closures are not fully audited.
swiftc -swift-version 5 -O -o "$work/swiftservice" "$here/swiftservice.swift" || exit 1
swiftc -swift-version 5 -O -o "$work/swiftclient" "$here/swiftclient.swift" || exit 1
echo "built into $work"

plist() { # label program logfile
	cat >"$agents/$1.plist" <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key><string>$1</string>
	<key>ProgramArguments</key>
	<array>
		<string>$2</string>
		<string>-service</string>
		<string>$1</string>
	</array>
	<key>MachServices</key>
	<dict><key>$1</key><true/></dict>
	<key>StandardOutPath</key><string>$3</string>
	<key>StandardErrorPath</key><string>$3</string>
</dict>
</plist>
EOF
}

mkdir -p "$agents"
plist "$swift_service" "$work/swiftservice" "$work/swiftservice.log"
plist "$go_service" "$work/goservice" "$work/goservice.log"
launchctl bootstrap "$domain" "$agents/$swift_service.plist" || exit 1
launchctl bootstrap "$domain" "$agents/$go_service.plist" || exit 1

run() { # label command...
	echo
	echo "--- $1 ---"
	shift
	"$@"
	echo "(exit $?)"
}

echo
echo "=== A: Go client -> Swift service ($swift_service) ==="
run "Go sends every type it can encode; Swift reports what it received" \
	"$work/goclient" -service "$swift_service" -op describe
run "Swift sends every XPC type; Go reports what its codec produced" \
	"$work/goclient" -service "$swift_service" -op typezoo
run "Swift returns nil (no reply); Go SendSync waits" \
	"$work/goclient" -service "$swift_service" -op silent
run "Swift replies with a legitimate payload containing an \"error\" key" \
	"$work/goclient" -service "$swift_service" -op errorkey
run "Go carries a Swift listener endpoint out and back; Swift dials the relayed copy" \
	"$work/goclient" -service "$swift_service" -op endpointrelay
for key in date uuid fd shmem endpoint; do
	run "Swift sends only \"$key\" beyond the base scalars" \
		"$work/goclient" -service "$swift_service" -op "typezoo:$key"
done

echo
echo "=== B: Swift client -> Go service ($go_service) ==="
run "Swift sends every XPC type; Go reports what its codec produced" \
	"$work/swiftclient" -service "$go_service" -op describe
run "Go sends every type it can encode; Swift reports what it received" \
	"$work/swiftclient" -service "$go_service" -op typezoo
run "Go returns (nil, nil) (no reply); Swift waits with a bounded timeout" \
	"$work/swiftclient" -service "$go_service" -op silent
run "Go handler returns an error, which becomes {\"error\": string}" \
	"$work/swiftclient" -service "$go_service" -op fail
run "Go replies with a legitimate payload containing an \"error\" key" \
	"$work/swiftclient" -service "$go_service" -op errorkey

echo
echo "=== service logs ==="
for f in "$work/swiftservice.log" "$work/goservice.log"; do
	echo "--- $f ---"
	cat "$f" 2>/dev/null
done
