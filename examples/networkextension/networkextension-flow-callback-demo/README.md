# NetworkExtension flow callback demo

This example registers a local Objective-C object with selectors shaped like
NetworkExtension flow, session, URL filter, and identity callbacks. It calls
those selectors directly with generated NetworkExtension blocks, including the
identity plus certificate-chain completion block, and calls generated app proxy
flow, packet tunnel flow, and tunnel provider session wrapper methods against
the same local object. It does not create real proxy flows, VPN sessions, URL
filter requests, or connections.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-flow-callback-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-flow-callback-demo
```
