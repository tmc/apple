// swiftclient is the Swift half of direction B: an XPCSession talking to a
// Mach service served by Go's xpc.Listener.
//
// Build:  swiftc -O -o swiftclient swiftclient.swift
// Run:    ./swiftclient -service <name> -op <op>; see README.md.

import Foundation
import XPC

let args = CommandLine.arguments
func flagValue(_ name: String, _ fallback: String? = nil) -> String? {
    if let i = args.firstIndex(of: name), i + 1 < args.count { return args[i + 1] }
    return fallback
}
guard let serviceName = flagValue("-service") else {
    FileHandle.standardError.write("swiftclient: -service is required\n".data(using: .utf8)!)
    exit(2)
}
let op = flagValue("-op", "describe")!

@Sendable func describe(_ dict: xpc_object_t) -> String {
    var lines: [String] = []
    xpc_dictionary_apply(dict) { key, value in
        lines.append("\(String(cString: key))\t\(describeValue(value))")
        return true
    }
    if lines.isEmpty { return "" }
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

// request is everything Swift can put in a dictionary, so the Go side can
// report what its codec made of each. The last block is the part the Go codec
// has no case for.
func request(_ op: String) -> xpc_object_t {
    let d = xpc_dictionary_create_empty()
    // "describe:<key>" sends the base scalars plus one exotic key, which is
    // how a key that crashes or traps the receiver is isolated.
    let only: String? = op.hasPrefix("describe:") ? String(op.dropFirst("describe:".count)) : nil
    func want(_ k: String) -> Bool { only == nil || only == k }
    xpc_dictionary_set_string(d, "op", only == nil ? op : "describe")
    guard op == "describe" || only != nil else { return d }
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

    if want("date") { xpc_dictionary_set_date(d, "date", 1_700_000_000_000_000_000) }
    if want("uuid") {
        var uu: [UInt8] = [0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
                           0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10]
        xpc_dictionary_set_uuid(d, "uuid", &uu)
    }
    if want("fd") {
        let path = "/tmp/swiftinterop-fd-probe.txt"
        try? "FD PAYLOAD\n".write(toFile: path, atomically: true, encoding: .utf8)
        let fd = open(path, O_RDONLY)
        if fd >= 0 {
            xpc_dictionary_set_fd(d, "fd", fd)
            close(fd)
        }
    }
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
    return d
}

do {
    let session = try XPCSession(machService: serviceName, options: .inactive)
    try session.activate()
    defer { session.cancel(reason: "done") }

    if op == "silent" {
        // Go's handler returns (nil, nil), which sends no reply. Ask for one
        // anyway and see whether the handler ever fires. A bounded wait is
        // the point: an unbounded sendSync would simply hang, which proves
        // nothing about whether the reply is merely late.
        let sem = DispatchSemaphore(value: 0)
        var outcome = "reply handler never fired"
        session.send(message: XPCDictionary(request(op))) { result in
            switch result {
            case .success(let d):
                outcome = "reply handler fired: \(describe(d.withUnsafeUnderlyingDictionary { $0 }))"
            case .failure(let e):
                outcome = "reply handler fired with error: \(e)"
            }
            sem.signal()
        }
        let timedOut = sem.wait(timeout: .now() + 3) == .timedOut
        print("silent op: timedOut=\(timedOut) \(outcome)")
        exit(0)
    }

    let reply: XPCDictionary = try session.sendSync(message: XPCDictionary(request(op)))
    let raw = reply.withUnsafeUnderlyingDictionary { $0 }
    if let report = xpc_dictionary_get_string(raw, "report"), xpc_dictionary_get_count(raw) == 1 {
        print(String(cString: report), terminator: "")
    } else {
        print(describe(raw), terminator: "")
    }
} catch {
    FileHandle.standardError.write("swiftclient: \(error)\n".data(using: .utf8)!)
    exit(1)
}
