# NetworkExtension value object demo

This example checks local value and state objects, including TLS parameters,
IKEv2 post-quantum settings, DNS proxy provider protocol, additional network
rule constructors including Network.framework endpoints, default path/TCP/UDP
state and endpoint getters, and generated array completion blocks. It does
not start network traffic, install extensions, or load/save NetworkExtension
preferences.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-value-object-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-value-object-demo
```
