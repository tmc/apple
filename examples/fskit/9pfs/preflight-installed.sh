#!/usr/bin/env bash
set -euo pipefail

app=${NINEPFS_INSTALLED_APP:-/Applications/NinePFSHost.app}
fsbundle=${NINEPFS_INSTALLED_FSBUNDLE:-/Library/Filesystems/9pfs.fs}
host_id=${NINEPFS_HOST_BUNDLE_ID:-dev.tmc.apple.examples.fskit.9pfs}
extension_id=${NINEPFS_BUNDLE_ID:-dev.tmc.apple.examples.fskit.9pfs.extension}

fail=0

check() {
	local label=$1
	shift
	if "$@"; then
		printf 'ok: %s\n' "$label"
	else
		printf 'missing: %s\n' "$label"
		fail=1
	fi
}

profile_app_id() {
	local profile=$1
	local plist

	plist=$(mktemp)
	if security cms -D -i "$profile" >"$plist" 2>/dev/null; then
		/usr/libexec/PlistBuddy -c 'Print :Entitlements:com.apple.application-identifier' "$plist" 2>/dev/null || true
	fi
	rm -f "$plist"
}

profile_has_fskit() {
	local profile=$1
	local plist

	plist=$(mktemp)
	if security cms -D -i "$profile" >"$plist" 2>/dev/null; then
		/usr/libexec/PlistBuddy -c 'Print :Entitlements:com.apple.developer.fskit.fsmodule' "$plist" 2>/dev/null || true
	fi
	rm -f "$plist"
}

check "installed host app" test -d "$app"
check "installed extension" test -d "$app/Contents/Extensions/NinePFSExtension.appex"
if [[ -x "$fsbundle/Contents/Resources/mount_9pfs" ]]; then
	echo "ok: installed mount helper"
else
	echo "note: no installed mount helper at $fsbundle/Contents/Resources/mount_9pfs"
	echo "hint: direct /sbin/mount -F -t 9pfs does not require the helper"
fi

if [[ -d "$app" ]]; then
	if codesign --verify --deep --strict "$app" >/dev/null 2>&1; then
		echo "ok: host signature"
	else
		echo "missing: host signature is not strict-valid"
		fail=1
	fi
fi

host_profile=$app/Contents/embedded.provisionprofile
extension_profile=$app/Contents/Extensions/NinePFSExtension.appex/Contents/embedded.provisionprofile

if [[ -f "$host_profile" ]]; then
	host_app_id=$(profile_app_id "$host_profile")
	if [[ "${host_app_id#*.}" == "$host_id" ]]; then
		echo "ok: host profile matches $host_id"
	else
		echo "missing: host profile app id ${host_app_id:-<none>} does not match $host_id"
		fail=1
	fi
else
	echo "missing: host embedded provisioning profile for $host_id"
	echo "hint: rebuild with NINEPFS_HOST_PROFILE or a matching auto-discovered profile"
	fail=1
fi

if [[ -f "$extension_profile" ]]; then
	extension_app_id=$(profile_app_id "$extension_profile")
	if [[ "${extension_app_id#*.}" == "$extension_id" ]]; then
		echo "ok: extension profile matches $extension_id"
	else
		echo "missing: extension profile app id ${extension_app_id:-<none>} does not match $extension_id"
		fail=1
	fi
	if [[ "$(profile_has_fskit "$extension_profile")" == true ]]; then
		echo "ok: extension profile grants com.apple.developer.fskit.fsmodule"
	else
		echo "missing: extension profile does not grant com.apple.developer.fskit.fsmodule"
		fail=1
	fi
else
	echo "missing: extension embedded provisioning profile for $extension_id"
	echo "hint: rebuild with NINEPFS_EXTENSION_PROFILE granting com.apple.developer.fskit.fsmodule"
	fail=1
fi

if pluginkit -m -A -D -i "$extension_id" 2>/dev/null | grep -q "$extension_id"; then
	echo "ok: PlugInKit registration"
else
	echo "missing: PlugInKit registration for $extension_id"
	fail=1
fi

if [[ -x "$app/Contents/MacOS/NinePFSHost" ]]; then
	probe=$("$app/Contents/MacOS/NinePFSHost" --fskit-probe 2>&1 || true)
	printf '%s\n' "$probe"
	if printf '%s\n' "$probe" | grep -q "fskit: $extension_id enabled=true"; then
		echo "ok: FSKit reports module enabled"
	elif printf '%s\n' "$probe" | grep -q "fskit: $extension_id enabled=false"; then
		echo "missing: FSKit reports module disabled; enable it in System Settings"
		echo "hint: $app/Contents/MacOS/NinePFSHost --open-fskit-settings"
		fail=1
	else
		echo "missing: FSKit does not list $extension_id"
		fail=1
	fi
else
	echo "missing: host probe executable"
	fail=1
fi

if [[ "$fail" -ne 0 ]]; then
	echo "9pfs: installed preflight failed"
	exit 1
fi

echo "9pfs: installed preflight ok"
