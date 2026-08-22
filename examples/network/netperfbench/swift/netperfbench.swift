// netperfbench, Swift edition: the same TCP echo benchmark as the Go
// command one directory up, written against Network.framework directly.
// It speaks the identical wire protocol, so a Swift client can measure a Go
// server and the reverse — which is what separates binding overhead from
// framework cost.
//
// Build and run:
//
//	swiftc -O netperfbench.swift -o netperfbench-swift
//	./netperfbench-swift -role both -size 4096 -n 10000

import Foundation
import Network
import Darwin
import os.signpost

// MARK: - Arguments

struct Options {
    var role = "both"
    var addr = "127.0.0.1:51000"
    var port: UInt16 = 51000
    var size = 4096
    var count = 10000
    var warmup = 200
    var inflight = 1
    var repeatCount = 1
    var label = "swift"
    var json = false
    var signpost = false
    var recvBatch = false
}

func parseOptions() -> Options {
    var o = Options()
    var args = Array(CommandLine.arguments.dropFirst())
    while let flag = args.first {
        args.removeFirst()
        // Accept -flag value and -flag=value, and both - and -- prefixes.
        var name = flag
        var inline: String? = nil
        if let eq = flag.firstIndex(of: "=") {
            name = String(flag[flag.startIndex..<eq])
            inline = String(flag[flag.index(after: eq)...])
        }
        name = name.hasPrefix("--") ? String(name.dropFirst(2)) : String(name.dropFirst(1))

        func value() -> String {
            if let v = inline { return v }
            guard let v = args.first else {
                FileHandle.standardError.write("netperfbench: -\(name) needs a value\n".data(using: .utf8)!)
                exit(2)
            }
            args.removeFirst()
            return v
        }

        switch name {
        case "role": o.role = value()
        case "addr": o.addr = value()
        case "port": o.port = UInt16(value()) ?? o.port
        case "size": o.size = Int(value()) ?? o.size
        case "n": o.count = Int(value()) ?? o.count
        case "warmup": o.warmup = Int(value()) ?? o.warmup
        case "inflight": o.inflight = max(Int(value()) ?? 1, 1)
        case "repeat": o.repeatCount = max(Int(value()) ?? 1, 1)
        case "label": o.label = value()
        case "json": o.json = true
        case "signpost": o.signpost = true
        case "recv-batch": o.recvBatch = true
        default:
            FileHandle.standardError.write("netperfbench: unknown flag -\(name)\n".data(using: .utf8)!)
            exit(2)
        }
    }
    return o
}

let opts = parseOptions()
let signpostLog = OSLog(subsystem: "github.com.tmc.apple.netperfbench", category: .pointsOfInterest)

func beginSignpost(_ name: StaticString) -> OSSignpostID? {
    guard opts.signpost else { return nil }
    let id = OSSignpostID(log: signpostLog)
    os_signpost(.begin, log: signpostLog, name: name, signpostID: id)
    return id
}

func endSignpost(_ name: StaticString, _ id: OSSignpostID?) {
    guard let id else { return }
    os_signpost(.end, log: signpostLog, name: name, signpostID: id)
}

func eventSignpost(_ name: StaticString, _ id: OSSignpostID?) {
    guard let id else { return }
    os_signpost(.event, log: signpostLog, name: name, signpostID: id)
}

// MARK: - Parameters

// Cleartext TCP with Nagle disabled, matching both Go transports.
func plainTCP() -> NWParameters {
    let tcp = NWProtocolTCP.Options()
    tcp.noDelay = true
    return NWParameters(tls: nil, tcp: tcp)
}

func die(_ message: String) -> Never {
    FileHandle.standardError.write("netperfbench: \(message)\n".data(using: .utf8)!)
    exit(1)
}

// MARK: - Server

final class EchoServer {
    let listener: NWListener
    let queue = DispatchQueue(label: "netperfbench.server")

    init(port: UInt16) {
        let nwPort = NWEndpoint.Port(rawValue: port) ?? .any
        guard let l = try? NWListener(using: plainTCP(), on: nwPort) else {
            die("cannot listen on port \(port)")
        }
        listener = l
    }

    func start() {
        let ready = DispatchSemaphore(value: 0)
        listener.stateUpdateHandler = { state in
            switch state {
            case .ready: ready.signal()
            case .failed(let error): die("listener failed: \(error)")
            default: break
            }
        }
        listener.newConnectionHandler = { conn in
            conn.start(queue: self.queue)
            self.pump(conn)
        }
        listener.start(queue: queue)
        if ready.wait(timeout: .now() + 10) == .timedOut {
            die("listener not ready after 10s")
        }
    }

    // pump echoes each received chunk and schedules the next receive from
    // the send completion, keeping one operation outstanding per direction.
    private func pump(_ conn: NWConnection) {
        conn.receive(minimumIncompleteLength: 1, maximumLength: 1 << 16) { data, _, isComplete, error in
            if error != nil || (data == nil && isComplete) {
                conn.cancel()
                return
            }
            guard let data, !data.isEmpty else {
                self.pump(conn)
                return
            }
            conn.send(content: data, completion: .contentProcessed { sendError in
                if sendError != nil {
                    conn.cancel()
                    return
                }
                self.pump(conn)
            })
        }
    }

    var port: UInt16 { listener.port?.rawValue ?? 0 }
}

// MARK: - Client

final class EchoClient {
    let conn: NWConnection
    let queue = DispatchQueue(label: "netperfbench.client")
    private let sent = DispatchSemaphore(value: 0)
    private let received = DispatchSemaphore(value: 0)
    private var sendError: Error?
    private var recvError: Error?
    private var recvCount = 0

    init(addr: String) {
        let parts = addr.split(separator: ":")
        guard parts.count == 2, let port = NWEndpoint.Port(rawValue: UInt16(parts[1]) ?? 0) else {
            die("bad address \(addr)")
        }
        conn = NWConnection(host: NWEndpoint.Host(String(parts[0])), port: port, using: plainTCP())
    }

    func connect() {
        let ready = DispatchSemaphore(value: 0)
        conn.stateUpdateHandler = { state in
            switch state {
            case .ready: ready.signal()
            case .failed(let error): die("connection failed: \(error)")
            case .cancelled: die("connection cancelled")
            default: break
            }
        }
        conn.start(queue: queue)
        if ready.wait(timeout: .now() + 10) == .timedOut {
            die("connection not ready after 10s")
        }
    }

    // roundTrip sends n copies of payload before waiting for any of them,
    // then reads all n echoes back. n == 1 is pure latency; n > 1 overlaps
    // the per-operation costs, which is what separates a fixed cost per
    // send from a cost per byte.
    func roundTrip(_ payload: Data, _ n: Int) {
        let roundTripID = beginSignpost("round-trip")
        defer { endSignpost("round-trip", roundTripID) }

        sendError = nil
        for _ in 0..<n {
            let signpostID = beginSignpost("send")
            conn.send(content: payload, completion: .contentProcessed { error in
                endSignpost("send", signpostID)
                eventSignpost("send-callback", signpostID)
                if error != nil { self.sendError = error }
                self.sent.signal()
            })
        }

        // The default asks for one echo at a time. -recv-batch is an
        // experimental whole-batch receive, matching the Go and C clients.
        if opts.recvBatch && UInt64(payload.count) * UInt64(n) > UInt64(UInt32.max) {
            die("receive length exceeds UInt32.max")
        }
        var remaining = payload.count * n
        while remaining > 0 {
            recvError = nil
            recvCount = 0
            let want = opts.recvBatch ? remaining : payload.count
            let receiveID = beginSignpost("receive")
            // This interval starts on the benchmark thread and ends only
            // after it resumes from the semaphore. The receive-signal event
            // is emitted on the Network.framework callback thread, so the
            // trace separates callback completion from wake-up latency.
            let waitID = beginSignpost("receive-wait")
            conn.receive(minimumIncompleteLength: want, maximumLength: want) { data, _, _, error in
                endSignpost("receive", receiveID)
                eventSignpost("receive-callback", receiveID)
                self.recvError = error
                self.recvCount = data?.count ?? 0
                eventSignpost("receive-signal", waitID)
                self.received.signal()
            }
            received.wait()
            endSignpost("receive-wait", waitID)
            if let recvError { die("receive: \(recvError)") }
            if recvCount != want {
                die("echo returned \(recvCount) bytes, want \(want)")
            }
            remaining -= recvCount
        }

        for _ in 0..<n {
            let waitID = beginSignpost("send-wait")
            sent.wait()
            endSignpost("send-wait", waitID)
        }
        if let sendError { die("send: \(sendError)") }
    }
}

// MARK: - Reporting

struct CPUTime {
    var user: Double
    var sys: Double

    func subtracting(_ earlier: CPUTime) -> CPUTime {
        CPUTime(user: user - earlier.user, sys: sys - earlier.sys)
    }

    var total: Double { user + sys }
}

func readCPUTime() -> CPUTime {
    var usage = rusage()
    guard getrusage(RUSAGE_SELF, &usage) == 0 else {
        return CPUTime(user: 0, sys: 0)
    }
    func microseconds(_ value: timeval) -> Double {
        Double(value.tv_sec) * 1e6 + Double(value.tv_usec)
    }
    return CPUTime(user: microseconds(usage.ru_utime), sys: microseconds(usage.ru_stime))
}

func loadAverage() -> String {
    var values = [Double](repeating: 0, count: 3)
    guard getloadavg(&values, Int32(values.count)) > 0 else {
        return ""
    }
    return String(format: "{ %.2f %.2f %.2f }", values[0], values[1], values[2])
}

// Run is one measured repetition. Process CPU time intentionally includes
// dispatch worker threads, where Network.framework performs much of its work.
struct Run {
    var latencies: [Double] // microseconds, sorted
    var elapsed: Double
    var cpu: CPUTime

    var count: Int { latencies.count }
    var messages: Int { count * opts.inflight }
    var throughputMBps: Double { Double(messages) * Double(opts.size) * 2 / elapsed / 1048576 }
    var rps: Double { Double(count) / elapsed }
    var perMessage: Double { percentile(0.50) / Double(opts.inflight) }
    var messageRate: Double { Double(messages) / elapsed }
    var cpuPerMessage: Double { cpu.total / Double(messages) }
    var busy: Double { cpu.total / (elapsed * 1e6) }
    var mean: Double { latencies.reduce(0, +) / Double(latencies.count) }
    var stddev: Double {
        let m = mean
        return (latencies.reduce(0) { $0 + ($1 - m) * ($1 - m) } / Double(latencies.count)).squareRoot()
    }
    func percentile(_ p: Double) -> Double {
        latencies[Int(p * Double(latencies.count - 1))]
    }
}

struct Result {
    var label: String
    var size: Int
    var count: Int
    var inflight: Int
    var receiveBatch: Bool
    var repetitions: Int
    var run: Run
    var runs: [Run]

    func print(json: Bool) {
        if json {
            let fields: [(String, Any)] = [
                ("label", label), ("impl", "swift"),
                ("payload_bytes", size), ("round_trips", count), ("inflight", inflight),
                ("receive_batch", receiveBatch),
                ("messages", run.messages), ("repetitions", repetitions),
                ("p50_us_per_message", run.perMessage), ("messages_per_sec", run.messageRate),
                ("elapsed_sec", run.elapsed), ("throughput_mbps", run.throughputMBps),
                ("round_trips_per_sec", run.rps),
                ("min_us", run.percentile(0)), ("p50_us", run.percentile(0.50)),
                ("p90_us", run.percentile(0.90)), ("p99_us", run.percentile(0.99)),
                ("max_us", run.latencies.last ?? 0),
                ("mean_us", run.mean), ("stddev_us", run.stddev),
                ("cpu_us_per_message", run.cpuPerMessage),
                ("cpu_busy_fraction", run.busy),
            ]
            let body = fields.map { key, value -> String in
                if let s = value as? String { return "  \"\(key)\": \"\(s)\"" }
                if let i = value as? Int { return "  \"\(key)\": \(i)" }
                if let b = value as? Bool { return "  \"\(key)\": \(b)" }
                return "  \"\(key)\": \(String(format: "%.4f", value as! Double))"
            }.joined(separator: ",\n")
            let cpu = String(format: "  \"cpu\": {\"user_us\": %.4f, \"sys_us\": %.4f}", run.cpu.user, run.cpu.sys)
            let env = "  \"env\": {\"load_average\": \"\(loadAverage())\"}"
            let allRuns = runs.map { r in
                String(format: "{\"elapsed_sec\": %.4f, \"p50_us_per_message\": %.4f, \"cpu_us_per_message\": %.4f, \"cpu_busy_fraction\": %.4f, \"cpu\": {\"user_us\": %.4f, \"sys_us\": %.4f}}", r.elapsed, r.perMessage, r.cpuPerMessage, r.busy, r.cpu.user, r.cpu.sys)
            }.joined(separator: ", ")
            Swift.print("{\n\(body),\n\(cpu),\n\(env),\n  \"runs\": [\(allRuns)]\n}")
            return
        }
        Swift.print(String(format: "%@: %d batches of %d x %d bytes in %.2fs", label, count, inflight, size, run.elapsed))
        Swift.print(String(format: "  latency us   min %.1f  p50 %.1f  p90 %.1f  p99 %.1f  max %.1f  mean %.1f ±%.1f",
                           run.percentile(0), run.percentile(0.50), run.percentile(0.90), run.percentile(0.99),
                           run.latencies.last ?? 0, run.mean, run.stddev))
        Swift.print(String(format: "  %.1f us/message, %.0f messages/sec, %.1f MB/s", run.perMessage, run.messageRate, run.throughputMBps))
        Swift.print(String(format: "  cpu %.1f us/message (%.1f user, %.1f sys), %.2f cores busy, %.1f us/message waiting",
                           run.cpuPerMessage, run.cpu.user / Double(run.messages), run.cpu.sys / Double(run.messages),
                           run.busy, run.perMessage - run.cpuPerMessage))
    }
}

func runClient(addr: String) {
    let client = EchoClient(addr: addr)
    client.connect()

    var bytes = [UInt8](repeating: 0, count: opts.size)
    for i in 0..<opts.size { bytes[i] = UInt8(i & 0xff) }
    let payload = Data(bytes)

    for _ in 0..<opts.warmup { client.roundTrip(payload, opts.inflight) }

    // Each repetition is summarized on its own; the median by p50 is
    // reported, which resists the one repetition that caught a scheduler
    // hiccup in a way a mean of means does not.
    var runs = [Run]()
    for _ in 0..<opts.repeatCount {
        var latencies = [Double]()
        latencies.reserveCapacity(opts.count)
        let cpu0 = readCPUTime()
        let start = DispatchTime.now()
        for _ in 0..<opts.count {
            let t0 = DispatchTime.now()
            client.roundTrip(payload, opts.inflight)
            latencies.append(Double(DispatchTime.now().uptimeNanoseconds - t0.uptimeNanoseconds) / 1000)
        }
        let elapsed = Double(DispatchTime.now().uptimeNanoseconds - start.uptimeNanoseconds) / 1e9
        latencies.sort()
        runs.append(Run(latencies: latencies, elapsed: elapsed, cpu: readCPUTime().subtracting(cpu0)))
    }
    runs.sort { $0.percentile(0.50) < $1.percentile(0.50) }
    let median = runs[runs.count / 2]
    Result(label: opts.label, size: opts.size, count: opts.count, inflight: opts.inflight, receiveBatch: opts.recvBatch, repetitions: opts.repeatCount,
           run: median, runs: runs).print(json: opts.json)
    client.conn.cancel()
}

switch opts.role {
case "server":
    let server = EchoServer(port: opts.port)
    server.start()
    FileHandle.standardError.write("listening on port \(server.port) (swift)\n".data(using: .utf8)!)
    dispatchMain()
case "client":
    runClient(addr: opts.addr)
case "both":
    let server = EchoServer(port: 0)
    server.start()
    runClient(addr: "127.0.0.1:\(server.port)")
default:
    die("unknown -role \(opts.role)")
}
