# NetworkExtension provider wrapper demo

This example checks generated `NEPacketTunnelProvider` and
`NEAppProxyProvider` wrapper methods against local Objective-C callback
objects. It invokes provider lifecycle completion blocks, synchronous context
wrappers whose local callbacks complete immediately, cancellation methods,
packet tunnel accessors, app-proxy flow decisions, and provider subclass
class availability for Ethernet tunnel and transparent proxy providers. It
does not install provider extensions or start live network traffic.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-provider-wrapper-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-provider-wrapper-demo
```
