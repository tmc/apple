// swiftservice is the Swift half of direction A: an XPCListener on a Mach
// service that a Go xpc.Session talks to.
//
// It answers the same ops as goservice, plus "typezoo", which sends one value
// of *every* XPC type — including the ones the Go codec has no case for (fd,
// shmem, uuid, date, endpoint, connection) — so the Go client can report what
// it received for each.
//
// Build:  swiftc -O -o swiftservice swiftservice.swift
// Run:    launchd starts it with -service <name>; see README.md.

import Foundation
import XPC

let args = CommandLine.arguments
guard let i = args.firstIndex(of: "-service"), i + 1 < args.count else {
    FileHandle.standardError.write("swiftservice: -service is required\n".data(using: .utf8)!)
    exit(2)
}
let serviceName = args[i + 1]

func log(_ s: String) {
    FileHandle.standardError.write("swiftservice: \(s)\n".data(using: .utf8)!)
}

// describe renders the XPC type name and value of every key in a dictionary,
// as Swift sees it. This is the mirror image of goservice.Describe: it says
// what Swift got, in XPC's own vocabulary rather than Go's.
@Sendable func describe(_ dict: xpc_object_t) -> String {
    var lines: [String] = []
    xpc_dictionary_apply(dict) { key, value in
        lines.append("\(String(cString: key))\t\(describeValue(value))")
        return true
    }
    return lines.sorted().joined(separator: "\n") + "\n"
}

@Sendable func describeValue(_ value: xpc_object_t) -> String {
    let type = xpc_get_type(value)
    let name = String(cString: xpc_type_get_name(type))
    switch type {
    case XPC_TYPE_BOOL:   return "\(name)\t\(xpc_bool_get_value(value))"
    case XPC_TYPE_INT64:  return "\(name)\t\(xpc_int64_get_value(value))"
    case XPC_TYPE_UINT64: return "\(name)\t\(xpc_uint64_get_value(value))"
    case XPC_TYPE_DOUBLE: return "\(name)\t\(xpc_double_get_value(value))"
    case XPC_TYPE_STRING: return "\(name)\t\"\(String(cString: xpc_string_get_string_ptr(value)!))\""
    case XPC_TYPE_DATE:   return "\(name)\t\(xpc_date_get_value(value)) ns since epoch"
    case XPC_TYPE_NULL:   return "\(name)\t<null>"
    case XPC_TYPE_DATA:
        let n = xpc_data_get_length(value)
        var hex = ""
        if let p = xpc_data_get_bytes_ptr(value) {
            let b = p.assumingMemoryBound(to: UInt8.self)
            for k in 0..<Int(n) { hex += String(format: "%02x", b[k]) }
        }
        return "\(name)\t\(n) bytes \(hex)"
    case XPC_TYPE_UUID:
        let p = xpc_uuid_get_bytes(value)!
        var hex = ""
        for k in 0..<16 { hex += String(format: "%02x", p[k]) }
        return "\(name)\t\(hex)"
    case XPC_TYPE_FD:
        let fd = xpc_fd_dup(value)
        defer { if fd >= 0 { close(fd) } }
        return "\(name)\tdup'd to fd \(fd)"
    case XPC_TYPE_ARRAY:
        var parts: [String] = []
        xpc_array_apply(value) { _, element in
            parts.append(describeValue(element).replacingOccurrences(of: "\t", with: " "))
            return true
        }
        return "\(name)\t[\(parts.joined(separator: ", "))]"
    case XPC_TYPE_DICTIONARY:
        let inner = describe(value).trimmingCharacters(in: .newlines)
            .replacingOccurrences(of: "\n", with: " | ")
        return "\(name)\t\(inner)"
    default:
        return "\(name)\t\(String(cString: xpc_copy_description(value)))"
    }
}

// typeZoo builds one value of every XPC type. The types after "dict" are the
// ones the Go high-level codec cannot produce and, on receipt, cannot decode.
// only == nil means every key. Otherwise only the base scalars plus the one
// named exotic key are sent, which is how a crashing key is isolated.
@Sendable func typeZoo(endpoint: XPCEndpoint?, only: String? = nil) -> xpc_object_t {
    func want(_ k: String) -> Bool { only == nil || only == k }
    let d = xpc_dictionary_create_empty()
    xpc_dictionary_set_bool(d, "bool", true)
    xpc_dictionary_set_int64(d, "int64", -42)
    xpc_dictionary_set_uint64(d, "uint64", UInt64(1) << 63)
    xpc_dictionary_set_double(d, "double", 3.5)
    xpc_dictionary_set_string(d, "string", "héllo")
    let bytes: [UInt8] = [0xde, 0xad, 0xbe, 0xef]
    bytes.withUnsafeBytes { xpc_dictionary_set_data(d, "data", $0.baseAddress!, $0.count) }
    xpc_dictionary_set_value(d, "null", xpc_null_create())
    let arr = xpc_array_create_empty()
    xpc_array_append_value(arr, xpc_int64_create(1))
    xpc_array_append_value(arr, xpc_string_create("two"))
    xpc_array_append_value(arr, xpc_bool_create(false))
    xpc_dictionary_set_value(d, "array", arr)
    let nested = xpc_dictionary_create_empty()
    xpc_dictionary_set_int64(nested, "nested", 7)
    xpc_dictionary_set_value(d, "dict", nested)

    // --- beyond the Go codec ---

    // date: nanoseconds since the epoch.
    if want("date") { xpc_dictionary_set_date(d, "date", 1_700_000_000_000_000_000) }

    // uuid: 16 raw bytes.
    if want("uuid") {
        var uu: [UInt8] = [0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
                           0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10]
        xpc_dictionary_set_uuid(d, "uuid", &uu)
    }

    // fd: a real descriptor open on a file with known contents, so the
    // receiver can prove whether it got a usable descriptor or a string.
    if want("fd") {
        let path = "/tmp/swiftinterop-fd-probe.txt"
        try? "FD PAYLOAD\n".write(toFile: path, atomically: true, encoding: .utf8)
        let fd = open(path, O_RDONLY)
        if fd >= 0 {
            xpc_dictionary_set_fd(d, "fd", fd)
            close(fd)
        }
    }

    // shmem: a mapped page with a recognisable prefix.
    if want("shmem") {
        // XPC requires a shared-memory region to be page-aligned memory it
        // can share: vm_allocate/mmap memory, not malloc memory. Passing a
        // posix_memalign region is API misuse and traps the SENDING process
        // with SIGTRAP inside xpc_shmem_create.
        let pageSize = Int(getpagesize())
        let region = mmap(nil, pageSize, PROT_READ | PROT_WRITE,
                          MAP_ANON | MAP_SHARED, -1, 0)
        if region != MAP_FAILED, let region {
            let msg = "SHMEM PAYLOAD"
            _ = msg.withCString { memcpy(region, $0, strlen($0)) }
            xpc_dictionary_set_value(d, "shmem", xpc_shmem_create(region, pageSize))
        }
    }

    // endpoint: this listener's own endpoint, so the Go side can be asked
    // whether it can do anything at all with one.
    if want("endpoint"), let endpoint {
        var wrapper = XPCDictionary(d)
        wrapper["endpoint"] = endpoint
    }
    return d
}

@Sendable func reply(to received: XPCDictionary, listener: XPCListener?) -> XPCDictionary? {
    let raw = received.withUnsafeUnderlyingDictionary { $0 }
    let op = xpc_dictionary_get_string(raw, "op").map { String(cString: $0) } ?? ""
    log("op=\(op) keys=\(received.count)")
    if op == "typezoo" {
        return XPCDictionary(typeZoo(endpoint: anonListener.endpoint))
    }
    if op.hasPrefix("typezoo:") {
        return XPCDictionary(typeZoo(endpoint: anonListener.endpoint,
                                     only: String(op.dropFirst("typezoo:".count))))
    }
    switch op {
    case "describe":
        let out = xpc_dictionary_create_empty()
        xpc_dictionary_set_string(out, "report", describe(raw))
        return XPCDictionary(out)
    case "echoendpoint":
        // The peer is handing back an endpoint that travelled out through it
        // and came back. If a session can be built from it and a message sent
        // through, the relay was lossless.
        let out = xpc_dictionary_create_empty()
        guard let ep = received["endpoint"] as XPCEndpoint? else {
            xpc_dictionary_set_string(out, "report", "no endpoint key survived the round trip\n")
            return XPCDictionary(out)
        }
        do {
            let s = try XPCSession(endpoint: ep)
            let ping = xpc_dictionary_create_empty()
            xpc_dictionary_set_string(ping, "op", "ping")
            let r: XPCDictionary = try s.sendSync(message: XPCDictionary(ping))
            let text = r.withUnsafeUnderlyingDictionary { describe($0) }
            s.cancel(reason: "done")
            xpc_dictionary_set_string(out, "report",
                "relayed endpoint is live; reply through it:\n" + text)
        } catch {
            xpc_dictionary_set_string(out, "report", "relayed endpoint unusable: \(error)\n")
        }
        return XPCDictionary(out)
    case "ping":
        let out = xpc_dictionary_create_empty()
        xpc_dictionary_set_string(out, "pong", "reached through a relayed endpoint")
        return XPCDictionary(out)
    case "silent":
        return nil
    case "errorkey":
        let out = xpc_dictionary_create_empty()
        xpc_dictionary_set_string(out, "error", "this is real data, not a failure")
        xpc_dictionary_set_string(out, "status", "ok")
        return XPCDictionary(out)
    default:
        let out = xpc_dictionary_create_empty()
        xpc_dictionary_set_string(out, "unknownOp", op)
        return XPCDictionary(out)
    }
}

final class Box: @unchecked Sendable {
    var listener: XPCListener?
}
let box = Box()

// The endpoint handed out by "typezoo" belongs to a second, ANONYMOUS
// listener, not to the Mach-service listener. Two reasons: an anonymous
// listener is the case endpoints exist for, and dialling one's own listener
// from inside its own message handler deadlocks, so "echoendpoint" needs a
// listener on a different queue to talk to.
let anonQueue = DispatchQueue(label: "swiftservice.anon")
let anonListener = XPCListener(targetQueue: anonQueue) { request in
    request.accept(incomingMessageHandler: { (message: XPCDictionary) -> XPCDictionary? in
        reply(to: message, listener: nil)
    }, cancellationHandler: { error in
        log("anon session cancelled: \(error)")
    })
}

do {
    let listener = try XPCListener(service: serviceName, options: .inactive) { request in
        request.accept(incomingMessageHandler: { (message: XPCDictionary) -> XPCDictionary? in
            reply(to: message, listener: box.listener)
        }, cancellationHandler: { error in
            log("session cancelled: \(error)")
        })
    }
    box.listener = listener
    try listener.activate()
    log("serving \(serviceName)")
    dispatchMain()
} catch {
    log("listener: \(error)")
    exit(1)
}
