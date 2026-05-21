# NetworkExtension provider callback demo

This example registers a local Objective-C object with the same lifecycle
selectors used by `NEPacketTunnelProvider` and calls them directly. It
exercises provider-style completion blocks without installing or running a
Network Extension.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-provider-callback-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-provider-callback-demo
```
