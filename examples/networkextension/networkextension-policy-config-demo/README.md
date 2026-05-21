# NetworkExtension policy/config demo

This example builds NetworkExtension policy and configuration objects locally
and checks setter/getter, class-factory, app-rule domain, endpoint, packet,
and default metadata round trips. It does not load, save, or remove
NetworkExtension preferences, start a provider, install an extension, or
require NetworkExtension entitlements.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-policy-config-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-policy-config-demo
```
