# NetworkExtension TCP auth delegate demo

This example registers a tiny Objective-C object that conforms to
`NWTCPConnectionAuthenticationDelegate` and calls its optional delegate
selectors directly. It is a local runtime smoke test for the Objective-C
helper split; it does not create a connection, install an extension, or require
NetworkExtension entitlements.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-tcp-auth-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-tcp-auth-demo
```
