# NetworkExtension settings demo

This example builds NetworkExtension configuration objects locally and checks
that the generated setters, getters, and local IPv4/IPv6 class factories
round trip, including base tunnel settings, filter-provider credential
references, and tunnel overhead. It also checks proxy PAC/static settings,
DNS identity references, and VPN credential references using in-memory data.
It does not save preferences, start a provider, install an extension, or
require NetworkExtension entitlements.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-settings-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-settings-demo
```
