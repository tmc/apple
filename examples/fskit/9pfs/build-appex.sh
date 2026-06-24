#!/usr/bin/env bash
set -euo pipefail

# Assemble the 9pfs FSKit extension as a Swift @main app-extension linked
# against the Go filesystem operations built as a c-archive.
#
#   ./build-appex.sh /tmp/9pfs-build
#
# Set CODESIGN_IDENTITY to sign the bundles. For an installed test the script
# also needs development provisioning profiles whose application identifiers
# match dev.tmc.apple.examples.fskit.9pfs (host) and
# dev.tmc.apple.examples.fskit.9pfs.extension (extension); the extension
# profile must grant com.apple.developer.fskit.fsmodule. Matching profiles are
# auto-discovered from ~/Library/MobileDevice/Provisioning Profiles; set
# NINEPFS_REQUIRE_PROFILES=yes to fail early when they are missing.

dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
out=${1:-/tmp/9pfs-build}
bundle_id=${NINEPFS_BUNDLE_ID:-dev.tmc.apple.examples.fskit.9pfs.extension}
product=${NINEPFS_PRODUCT_NAME:-NinePFSExtension}
host_product=${NINEPFS_HOST_PRODUCT_NAME:-NinePFSHost}
host_bundle_id=${NINEPFS_HOST_BUNDLE_ID:-dev.tmc.apple.examples.fskit.9pfs}
identity=${CODESIGN_IDENTITY:-}
extension_profile=${NINEPFS_EXTENSION_PROFILE:-}
host_profile=${NINEPFS_HOST_PROFILE:-}
require_profiles=${NINEPFS_REQUIRE_PROFILES:-}

bundle=$out/$product.appex
fsbundle=$out/9pfs.fs
app=$out/$host_product.app
contents=$bundle/Contents
macos=$contents/MacOS
frameworks=$contents/Frameworks
objdir=$out/obj
extension_entitlements=$out/$product.entitlements.plist
host_entitlements=$out/$host_product.entitlements.plist

rm -rf "$bundle" "$fsbundle" "$app" "$objdir"
mkdir -p "$macos" "$frameworks" "$objdir"

find_profile() {
	local want_bundle_id=$1
	local want_fskit=$2
	local profiles_dir=$HOME/Library/MobileDevice/Provisioning\ Profiles
	local profile app_id has_fskit

	[[ -d "$profiles_dir" ]] || return 1
	for profile in "$profiles_dir"/*; do
		[[ -f "$profile" ]] || continue
		if ! security cms -D -i "$profile" > "$out/profile-search.plist" 2>/dev/null; then
			continue
		fi
		app_id=$(/usr/libexec/PlistBuddy -c 'Print :Entitlements:com.apple.application-identifier' "$out/profile-search.plist" 2>/dev/null || true)
		[[ "${app_id#*.}" == "$want_bundle_id" ]] || continue
		has_fskit=$(/usr/libexec/PlistBuddy -c 'Print :Entitlements:com.apple.developer.fskit.fsmodule' "$out/profile-search.plist" 2>/dev/null || true)
		if [[ "$want_fskit" == yes && "$has_fskit" != true ]]; then
			continue
		fi
		rm -f "$out/profile-search.plist"
		printf '%s\n' "$profile"
		return 0
	done
	rm -f "$out/profile-search.plist"
	return 1
}

profile_entitlements() {
	local profile=$1
	local entitlements=$2

	security cms -D -i "$profile" > "$out/profile.plist"
	/usr/libexec/PlistBuddy -x -c 'Print :Entitlements' "$out/profile.plist" > "$entitlements"
	rm -f "$out/profile.plist"
}

ensure_bool_entitlement() {
	local entitlements=$1
	local key=$2
	if ! /usr/libexec/PlistBuddy -c "Print :$key" "$entitlements" >/dev/null 2>&1; then
		/usr/libexec/PlistBuddy -c "Add :$key bool true" "$entitlements"
	fi
}

ensure_array_string_entitlement() {
	local entitlements=$1
	local key=$2
	local value=$3
	if ! /usr/libexec/PlistBuddy -c "Print :$key" "$entitlements" >/dev/null 2>&1; then
		/usr/libexec/PlistBuddy -c "Add :$key array" "$entitlements"
	fi
	if ! /usr/libexec/PlistBuddy -c "Print :$key" "$entitlements" | grep -qx "    $value"; then
		/usr/libexec/PlistBuddy -c "Add :$key: string $value" "$entitlements"
	fi
}

prepare_entitlements() {
	local profile=$1
	local base=$2
	local entitlements=$3
	local role=$4
	local want_bundle_id=$5

	if [[ -n "$profile" ]]; then
		[[ -f "$profile" ]] || { echo "missing provisioning profile: $profile" >&2; exit 1; }
		profile_entitlements "$profile" "$entitlements"
		local app_id
		app_id=$(/usr/libexec/PlistBuddy -c 'Print :com.apple.application-identifier' "$entitlements" 2>/dev/null || true)
		if [[ -z "$app_id" || "${app_id#*.}" != "$want_bundle_id" ]]; then
			echo "provisioning profile $profile has application identifier ${app_id:-<none>}, want *.$want_bundle_id" >&2
			exit 1
		fi
	else
		cp "$base" "$entitlements"
	fi

	ensure_bool_entitlement "$entitlements" "com.apple.security.app-sandbox"
	case "$role" in
	extension)
		ensure_bool_entitlement "$entitlements" "com.apple.developer.fskit.fsmodule"
		ensure_bool_entitlement "$entitlements" "com.apple.security.network.client"
		;;
	host)
		ensure_array_string_entitlement "$entitlements" \
			"com.apple.security.temporary-exception.mach-lookup.global-name" \
			"com.apple.filesystems.fskitd"
		;;
	esac
}

if [[ -n "$identity" ]]; then
	if [[ -z "$extension_profile" ]]; then
		extension_profile=$(find_profile "$bundle_id" yes || true)
	fi
	if [[ -z "$host_profile" ]]; then
		host_profile=$(find_profile "$host_bundle_id" no || true)
	fi
fi

if [[ "$require_profiles" == yes ]]; then
	[[ -n "$extension_profile" ]] || {
		echo "missing matching extension profile for $bundle_id with com.apple.developer.fskit.fsmodule" >&2
		exit 1
	}
	[[ -n "$host_profile" ]] || {
		echo "missing matching host profile for $host_bundle_id" >&2
		exit 1
	}
fi

# Build the Go filesystem operations as a c-archive. The cshared build tag
# exports NinePFSInit/NinePFSConfigureFileSystem/NinePFS*Resource for the Swift
# entrypoint to call. A throwaway module file lets the build resolve the local
# apple module and the patched p9 client without a checked-in replace.
p9src=$objdir/p9-src
moddir=$objdir/mod
modfile=$moddir/go.mod
go_archive=$objdir/libNinePFS.a
mkdir -p "$moddir"

"$dir/prepare-p9-module.sh" "$p9src"
sed '/^replace github.com\/tmc\/apple /d' "$dir/go.mod" > "$modfile"
cp "$dir/go.sum" "${modfile%.mod}.sum"
{
	echo
	echo "replace github.com/tmc/apple => $(cd "$dir/../../.." && pwd)"
	echo "replace github.com/hugelgupf/p9 => $p9src"
} >> "$modfile"

deploy_target=${MACOSX_DEPLOYMENT_TARGET:-15.4}
(cd "$dir" && GOWORK=off GOFLAGS="-modfile=$modfile" \
	MACOSX_DEPLOYMENT_TARGET="$deploy_target" \
	CGO_CFLAGS="-mmacosx-version-min=$deploy_target ${CGO_CFLAGS:-}" \
	CGO_LDFLAGS="-mmacosx-version-min=$deploy_target ${CGO_LDFLAGS:-}" \
	go build -tags cshared \
		-ldflags "-extldflags=-mmacosx-version-min=$deploy_target" \
		-buildmode=c-archive -o "$go_archive" .)
rm -f "${go_archive%.a}.h"

# Compile the ObjC NinePFileSystem class and the Swift @main entrypoint, and
# link them with the Go archive into the extension executable.
header=$dir/swiftshim/Sources/NinePFSShimObjC/include/NinePFileSystem.h
objc_source=$dir/swiftshim/Sources/NinePFSShimObjC/NinePFileSystem.m
objc_object=$objdir/NinePFileSystem.o
swift_target=${NINEPFS_SWIFT_TARGET:-$(uname -m)-apple-macos15.4}

xcrun clang -fobjc-arc -fmodules \
	-target "$swift_target" \
	-I "$dir/swiftshim/Sources/NinePFSShimObjC/include" \
	-c "$objc_source" -o "$objc_object"
xcrun swiftc -O -parse-as-library \
	-target "$swift_target" \
	-application-extension \
	-framework Foundation \
	-framework ExtensionFoundation \
	-framework FSKit \
	-import-objc-header "$header" \
	"$dir/appex/NinePFSExtension.swift" "$objc_object" "$go_archive" \
	-Xlinker -e -Xlinker _NSExtensionMain \
	-Xlinker -rpath -Xlinker @executable_path/../Frameworks \
	-o "$macos/$product"

/usr/bin/sed \
	-e "s/\$(PRODUCT_BUNDLE_IDENTIFIER)/$bundle_id/g" \
	-e "s/\$(PRODUCT_NAME)/$product/g" \
	-e "s/\$(EXECUTABLE_NAME)/$product/g" \
	-e "s/\$(DEVELOPMENT_LANGUAGE)/en/g" \
	"$dir/appex/Info.plist" > "$contents/Info.plist"

prepare_entitlements "$extension_profile" "$dir/appex/NinePFSExtension.entitlements" "$extension_entitlements" extension "$bundle_id"
if [[ -n "$extension_profile" ]]; then
	cp "$extension_profile" "$contents/embedded.provisionprofile"
fi

if [[ -n "$identity" ]]; then
	codesign --force --timestamp=none --sign "$identity" "$macos/$product"
	codesign --force --timestamp=none --sign "$identity" \
		--entitlements "$extension_entitlements" "$bundle"
fi

# The 9pfs.fs mount-helper bundle is only needed for plain `mount -t 9pfs`;
# direct `/sbin/mount -F -t 9pfs` does not use it.
mkdir -p "$fsbundle/Contents/Resources"
cp "$dir/fsbundle/Info.plist" "$fsbundle/Contents/Info.plist"
cp "$dir/mounthelper/mount_9pfs" "$fsbundle/Contents/Resources/mount_9pfs"
chmod +x "$fsbundle/Contents/Resources/mount_9pfs"

# Build the host app and embed the extension under Contents/Extensions.
mkdir -p "$app/Contents/MacOS" "$app/Contents/Extensions"
xcrun swiftc -O -parse-as-library \
	-target "$swift_target" \
	-application-extension \
	-framework FSKit \
	-framework SwiftUI \
	"$dir/host/App.swift" \
	-o "$app/Contents/MacOS/$host_product"
/usr/bin/sed \
	-e "s/\$(PRODUCT_BUNDLE_IDENTIFIER)/$host_bundle_id/g" \
	-e "s/\$(PRODUCT_NAME)/$host_product/g" \
	-e "s/\$(EXECUTABLE_NAME)/$host_product/g" \
	"$dir/host/Info.plist" > "$app/Contents/Info.plist"
cp -R "$bundle" "$app/Contents/Extensions/$product.appex"
prepare_entitlements "$host_profile" "$dir/host/NinePFSHost.entitlements" "$host_entitlements" host "$host_bundle_id"
if [[ -n "$host_profile" ]]; then
	cp "$host_profile" "$app/Contents/embedded.provisionprofile"
fi
if [[ -n "$identity" ]]; then
	codesign --force --timestamp=none --sign "$identity" \
		--entitlements "$host_entitlements" "$app"
fi

echo "$bundle"
echo "$fsbundle"
echo "$app"
