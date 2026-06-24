#!/usr/bin/env bash
# Build a SIGNED, INSTALLABLE pureentry 9pfs bundle whose extension binary is the
# no-cgo Go entrypoint (LC_MAIN -> main.nsextMainEntry.abi0 -> NSExtensionMain),
# instead of the default Swift @main binary.
#
# Strategy: let the maintained ../9pfs/build-appex.sh produce a correctly-signed
# host-app + .appex + .fs bundle (all entitlement/profile/signing logic reused),
# then rebuild the pureentry Go-entry binary from the current (post-extraction)
# tree and swap it into the .appex, re-signing just the extension binary + bundle.
#
# Output: $OUT (default /tmp/9pfs-pureentry-appex) containing NinePFSHost.app and
# 9pfs.fs, ready for:  CONFIRM_9PFS_INSTALL=yes ../9pfs/install-local.sh $OUT
#
# This does NOT install anything and makes no privileged/system changes.
set -uo pipefail

RESEARCH=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
NINE=$RESEARCH/../9pfs
NINE=$(cd "$NINE" && pwd)
TOOLS=$RESEARCH/pureentry
OUT=${1:-/tmp/9pfs-pureentry-appex}
OBJ=$OUT/pureentry-obj
PURE_SRC=$OBJ/pureentry-src
product=NinePFSExtension
# Match the identity + team the WORKING installed extension uses (Team LJ98655CHY),
# otherwise the LJ98655CHY.* provisioning profiles won't apply and the extension
# fails to load for a signing reason, not the research reason.
identity=${CODESIGN_IDENTITY:-Apple Development: Travis Cline (A6VCBGEHSE)}
# Provisioning profiles matching the bundle IDs. Default to the ones extracted
# from the installed working extension (the source files are not in the standard
# Provisioning Profiles dir on this machine). Override via env if needed.
ext_profile=${NINEPFS_EXTENSION_PROFILE:-/tmp/9pfs-profiles/extension.provisionprofile}
host_profile=${NINEPFS_HOST_PROFILE:-/tmp/9pfs-profiles/host.provisionprofile}

say() { echo "==> $*"; }
die() { echo "build-pureentry-appex: $*" >&2; exit 1; }

[[ -f "$ext_profile" ]] || die "missing extension profile: $ext_profile (extract from /Applications/NinePFSHost.app/Contents/Extensions/NinePFSExtension.appex/Contents/embedded.provisionprofile)"
[[ -f "$host_profile" ]] || die "missing host profile: $host_profile (extract from /Applications/NinePFSHost.app/Contents/embedded.provisionprofile)"

# ---------------------------------------------------------------------------
say "1/5 build the default signed bundle via build-appex.sh -> $OUT"
# Produce $OUT/NinePFSHost.app (+ embedded .appex) and $OUT/9pfs.fs, signed with
# the matching identity + provisioning profiles (so embedded.provisionprofile and
# the fskit entitlement land exactly as in the working installed extension).
# build-appex.sh leaves bundles UNSIGNED unless CODESIGN_IDENTITY is set, and
# auto-discovers profiles; NINEPFS_REQUIRE_PROFILES=yes makes it fail loudly if a
# matching profile is absent rather than silently producing an unloadable bundle.
( cd "$NINE" && CODESIGN_IDENTITY="$identity" \
	NINEPFS_EXTENSION_PROFILE="$ext_profile" NINEPFS_HOST_PROFILE="$host_profile" \
	NINEPFS_REQUIRE_PROFILES=yes ./build-appex.sh "$OUT" ) \
	|| die "default build-appex.sh failed (identity/profile mismatch?)"

app=$OUT/NinePFSHost.app
appex=$app/Contents/Extensions/$product.appex
ext_bin=$appex/Contents/MacOS/$product
ext_entitlements=$OUT/$product.entitlements.plist
[[ -d "$appex" ]] || die "expected signed appex not found at $appex"
[[ -f "$ext_entitlements" ]] || die "expected extension entitlements at $ext_entitlements (build-appex.sh writes these)"
[[ -f "$appex/Contents/embedded.provisionprofile" ]] || die "appex is missing embedded.provisionprofile — profile auto-discovery failed; set NINEPFS_EXTENSION_PROFILE/NINEPFS_HOST_PROFILE"
say "    default bundle built (signed, profiled); extension binary: $(file "$ext_bin" | sed 's/.*: //')"

# ---------------------------------------------------------------------------
say "2/5 assemble pureentry overlay source (9pfs core minus main_darwin.go + research fragments)"
rm -rf "$OBJ"
mkdir -p "$OBJ/linktmp" "$PURE_SRC"
cp "$NINE"/*.go "$PURE_SRC"/ 2>/dev/null
# main_darwin.go (//go:build darwin && !pureentry) is excluded by tag under
# -tags pureentry, but copy-then-remove keeps this robust regardless of guard.
rm -f "$PURE_SRC/main_darwin.go"
cp "$RESEARCH"/pureentry_darwin.go "$PURE_SRC"/
cp "$RESEARCH"/dynimport_pureentry_darwin.go "$PURE_SRC"/
cp "$RESEARCH"/nsext_arm64.s "$PURE_SRC"/
cp -R "$NINE/internal" "$PURE_SRC"/ 2>/dev/null
sed '/^replace github.com\/tmc\/apple /d' "$NINE/go.mod" > "$PURE_SRC/go.mod"
printf '\nreplace github.com/tmc/apple => %s\n' "$(cd "$NINE/../../.." && pwd)" >> "$PURE_SRC/go.mod"
cp "$NINE/go.sum" "$PURE_SRC/go.sum" 2>/dev/null
say "    overlay func main owner(s): $(grep -rln '^func main(' "$PURE_SRC"/*.go | xargs -n1 basename | tr '\n' ' ')"

# ---------------------------------------------------------------------------
say "3/5 emit swiftmeta object from the checked-in fixture (no swiftc), rename _main"
SMB=$OBJ/swiftmeta-bundled
mkdir -p "$SMB"
go run "$TOOLS/write_swiftmeta_fixture.go" -dir "$SMB" -compact || die "write_swiftmeta_fixture failed"
go run "$TOOLS/emit_macho_object.go" -manifest "$SMB/manifest.json" -data-dir "$SMB/sections" -o "$PURE_SRC/swiftmeta.syso" || die "emit_macho_object failed"
go run "$TOOLS/rename_macho_symbol.go" "$PURE_SRC/swiftmeta.syso" "$PURE_SRC/swiftmeta.syso.tmp" _main _swift_unused_main || die "rename_macho_symbol failed"
mv "$PURE_SRC/swiftmeta.syso.tmp" "$PURE_SRC/swiftmeta.syso"
say "    swiftmeta.syso: $(file "$PURE_SRC/swiftmeta.syso" | sed 's/.*: //')"

# ---------------------------------------------------------------------------
say "4/5 go build -tags pureentry, external link, patch LC_MAIN, verify metadata"
( cd "$PURE_SRC" && CGO_ENABLED=0 GOWORK=off go build -a -tags pureentry -work -o "$OBJ/$product.internal" . ) > "$OBJ/build.out" 2>&1
workdir=$(sed -n 's/^WORK=//p' "$OBJ/build.out" | tail -1)
if [[ -z "$workdir" || ! -f "$workdir/b001/importcfg.link" || ! -f "$workdir/b001/_pkg_.a" ]]; then
	tail -30 "$OBJ/build.out" >&2
	die "go build -tags pureentry did not produce link inputs"
fi
gotooldir=$(cd "$PURE_SRC" && GOWORK=off go env GOTOOLDIR)
"$gotooldir/link" \
	-o "$OBJ/$product.goentry" \
	-importcfg "$workdir/b001/importcfg.link" \
	-buildmode=pie -linkmode=external -extld=clang \
	-tmpdir "$OBJ/linktmp" \
	-extldflags '-Wl,-headerpad,1144 -framework ExtensionFoundation -framework FSKit -framework Foundation -lobjc -L/usr/lib/swift -lswiftCore -rpath /usr/lib/swift' \
	"$workdir/b001/_pkg_.a" > "$OBJ/link.out" 2>&1 || { tail -30 "$OBJ/link.out" >&2; die "external link failed"; }
rm -rf "$workdir"
go run "$TOOLS/patch_lcmain.go" -in "$OBJ/$product.goentry" -out "$OBJ/$product.patched" -target main.nsextMainEntry.abi0 || die "patch_lcmain failed"
go run "$TOOLS/check_swift_metadata.go" "$OBJ/$product.patched" || die "check_swift_metadata failed on patched binary"
say "    pureentry binary: $(file "$OBJ/$product.patched" | sed 's/.*: //')"
say "    LC_MAIN: $(otool -l "$OBJ/$product.patched" | awk '/LC_MAIN/{f=1} f&&/entryoff/{print; exit}')"

# ---------------------------------------------------------------------------
say "5/5 swap the pureentry binary into the signed .appex and re-sign (inner -> outer)"
host_product=NinePFSHost
host_entitlements=$OUT/$host_product.entitlements.plist
cp "$OBJ/$product.patched" "$ext_bin"
chmod +x "$ext_bin"
# Re-sign innermost first, then each enclosing bundle (modifying nested contents
# invalidates the outer seal). Sign the appex as a bundle (this seals the binary),
# then the host app deep with its own entitlements.
codesign --force --timestamp=none --sign "$identity" --entitlements "$ext_entitlements" "$appex" || die "codesign appex failed"
codesign --verify --strict "$appex" || die "codesign --verify failed on appex"
if [[ -f "$host_entitlements" ]]; then
	codesign --force --timestamp=none --sign "$identity" --entitlements "$host_entitlements" "$app" || die "codesign host app failed"
else
	codesign --force --timestamp=none --sign "$identity" "$app" || die "codesign host app failed"
fi
codesign --verify --deep --strict "$app" || die "codesign --verify --deep failed on host app"

echo
say "DONE. Signed pureentry bundle at: $OUT"
echo "    host app : $app"
echo "    extension: $ext_bin (Go-entry, $(file "$ext_bin" | sed 's/.*: //'))"
echo
echo "To install (replaces the working Swift-@main extension; needs your password):"
echo "    cd $NINE && CONFIRM_9PFS_INSTALL=yes ./install-local.sh $OUT"
echo "Then watch logs + mount per ../9pfs-research/MOUNT-PROBE-RUNBOOK.md."
echo "To restore the working extension afterward:"
echo "    cd $NINE && ./build-appex.sh /tmp/9pfs-appex && CONFIRM_9PFS_INSTALL=yes ./install-local.sh /tmp/9pfs-appex"
