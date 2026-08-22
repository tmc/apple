#!/bin/sh
# build-bundle.sh assembles a host application around the addservice XPC
# service, which launchd will only start out of a .xpc bundle inside a host
# app's Contents/XPCServices. addclient becomes the app's executable, since
# only a process inside the app can look the service up.
#
#	./examples/xpc/build-bundle.sh ~/tmp/AddDemo.app
#	~/tmp/AddDemo.app/Contents/MacOS/AddDemo -first 23 -second 19
#
# No code signing is needed: the ad-hoc signature the Go linker writes is
# enough for launchd to spawn the service.

set -e

app=${1:?usage: build-bundle.sh path/to/Host.app}
case $app in
	*.app) ;;
	*) echo "build-bundle.sh: destination must end in .app" >&2; exit 2 ;;
esac
case $app in
	/|"$HOME"|.|..) echo "build-bundle.sh: refusing unsafe destination: $app" >&2; exit 2 ;;
esac
service=dev.tmc.sample-xpc-service-ll
here=$(cd "$(dirname "$0")" && pwd)
xpc=$app/Contents/XPCServices/$service.xpc

if [ -e "$app" ] || [ -L "$app" ]; then
	marker=$app/Contents/Info.plist
	if [ -L "$app" ] || [ ! -f "$marker" ] ||
		! grep -q '<string>dev.tmc.adddemo</string>' "$marker"; then
		echo "build-bundle.sh: refusing to replace an app it did not create: $app" >&2
		exit 2
	fi
	rm -rf "$app"
fi
mkdir -p "$app/Contents/MacOS" "$xpc/Contents/MacOS"

go build -o "$app/Contents/MacOS/AddDemo" "$here/addclient"
go build -o "$xpc/Contents/MacOS/$service" "$here/addservice"

cat >"$app/Contents/Info.plist" <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleIdentifier</key><string>dev.tmc.adddemo</string>
	<key>CFBundleExecutable</key><string>AddDemo</string>
	<key>CFBundleName</key><string>AddDemo</string>
	<key>CFBundlePackageType</key><string>APPL</string>
</dict>
</plist>
EOF

# CFBundleIdentifier is the service name: that is what the client asks for.
# ServiceType Application gives each host application its own instance, which
# is what the C template's Info.plist specifies.
cat >"$xpc/Contents/Info.plist" <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleIdentifier</key><string>$service</string>
	<key>CFBundleExecutable</key><string>$service</string>
	<key>CFBundlePackageType</key><string>XPC!</string>
	<key>CFBundleInfoDictionaryVersion</key><string>6.0</string>
	<key>XPCService</key>
	<dict><key>ServiceType</key><string>Application</string></dict>
</dict>
</plist>
EOF

echo "$app"
