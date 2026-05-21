# NetworkExtension manager config demo

This example configures VPN and DNS settings manager objects in memory, checks
local setter/getter round trips, and invokes generated manager array completion
blocks locally. It does not load, save, or remove NetworkExtension preferences,
start a provider, install an extension, or require NetworkExtension
entitlements.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-manager-config-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-manager-config-demo
```
