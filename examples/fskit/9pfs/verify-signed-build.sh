#!/usr/bin/env bash
set -euo pipefail

build_dir=${1:?usage: verify-signed-build.sh BUILD_DIR}
app=$build_dir/NinePFSHost.app
appex=$app/Contents/Extensions/NinePFSExtension.appex
extension_bin=$appex/Contents/MacOS/NinePFSExtension

die() {
	echo "verify-signed-build: $*" >&2
	exit 1
}

[[ -d "$app" ]] || die "missing app: $app"
[[ -d "$appex" ]] || die "missing extension: $appex"
[[ -x "$extension_bin" ]] || die "missing extension executable: $extension_bin"

codesign --verify --deep --strict "$app"

host_id=$(/usr/libexec/PlistBuddy -c 'Print :CFBundleIdentifier' "$app/Contents/Info.plist")
extension_id=$(/usr/libexec/PlistBuddy -c 'Print :CFBundleIdentifier' "$appex/Contents/Info.plist")
extension_point=$(/usr/libexec/PlistBuddy -c 'Print :EXAppExtensionAttributes:EXExtensionPointIdentifier' "$appex/Contents/Info.plist")

[[ "$host_id" == "dev.tmc.apple.examples.fskit.9pfs" ]] ||
	die "unexpected host bundle id: $host_id"
[[ "$extension_id" == "dev.tmc.apple.examples.fskit.9pfs.extension" ]] ||
	die "unexpected extension bundle id: $extension_id"
[[ "$extension_point" == "com.apple.fskit.fsmodule" ]] ||
	die "unexpected extension point: $extension_point"

host_entitlements=$(mktemp)
extension_entitlements=$(mktemp)
trap 'rm -f "$host_entitlements" "$extension_entitlements"' EXIT

codesign -d --entitlements :- "$app" >"$host_entitlements" 2>/dev/null
codesign -d --entitlements :- "$appex" >"$extension_entitlements" 2>/dev/null

/usr/libexec/PlistBuddy -c 'Print :com.apple.security.app-sandbox' "$host_entitlements" | grep -qx true ||
	die "host is missing app sandbox entitlement"
/usr/libexec/PlistBuddy -c 'Print :com.apple.security.app-sandbox' "$extension_entitlements" | grep -qx true ||
	die "extension is missing app sandbox entitlement"
/usr/libexec/PlistBuddy -c 'Print :com.apple.developer.fskit.fsmodule' "$extension_entitlements" | grep -qx true ||
	die "extension is missing fskit module entitlement"
/usr/libexec/PlistBuddy -c 'Print :com.apple.security.network.client' "$extension_entitlements" | grep -qx true ||
	die "extension is missing network client entitlement"

echo "9pfs: signed build verification ok"
echo "host: $host_id"
echo "extension: $extension_id"
echo "extension point: $extension_point"
