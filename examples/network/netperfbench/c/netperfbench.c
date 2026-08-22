// netperfbench in C, against the Network.framework C API.
//
// This twin exists to remove a confound from the Go comparison. The Go
// bindings call nw_connection_send and nw_connection_receive directly. The
// Swift twin does not: NWConnection is a higher-level Swift class wrapping
// those same entry points, with its own object graph and completion
// plumbing. Subtracting Swift from Go therefore mixes "purego and the Go
// scheduler" with "Swift wrapper versus C API", in an unknown direction.
//
// This client calls exactly what the generated bindings call, so:
//
//	c - std     the framework's floor
//	nw - c      purego and the Go scheduler, with no wrapper confound
//	swift - c   what the Swift wrapper costs
//
// The measured loop matches the Go and Swift clients exactly: -inflight
// messages are sent without waiting, then read back one receive per echo
// with the minimum and maximum length both set to the payload size, so no
// implementation hides a round trip the others pay for.
//
// Build:
//	clang -O2 -fblocks netperfbench.c -framework Network -framework CoreFoundation -o netperfbench-c

#include <dispatch/dispatch.h>
#include <Network/Network.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/resource.h>
#include <sys/time.h>
#include <time.h>

static const char *opt_role = "both";
static const char *opt_addr = "127.0.0.1";
static const char *opt_port = "51000";
static int opt_size = 4096;
static int opt_count = 10000;
static int opt_warmup = 200;
static int opt_inflight = 1;
static int opt_repeat = 1;
static const char *opt_label = "c";
static bool opt_json = false;
static bool opt_recv_batch = false;

static void die(const char *what) {
	fprintf(stderr, "netperfbench: %s\n", what);
	exit(1);
}

// now returns a monotonic clock reading in microseconds.
static double now_us(void) {
	struct timespec ts;
	clock_gettime(CLOCK_MONOTONIC, &ts);
	return (double)ts.tv_sec * 1e6 + (double)ts.tv_nsec / 1e3;
}

// cpu_us returns process CPU time in microseconds. Process rather than
// thread time is deliberate: Network.framework does much of its work on
// dispatch worker threads, and thread time would miss it.
static double cpu_us(double *user, double *sys) {
	struct rusage ru;
	if (getrusage(RUSAGE_SELF, &ru) != 0) {
		return 0;
	}
	*user = (double)ru.ru_utime.tv_sec * 1e6 + (double)ru.ru_utime.tv_usec;
	*sys = (double)ru.ru_stime.tv_sec * 1e6 + (double)ru.ru_stime.tv_usec;
	return *user + *sys;
}

// plain_tcp returns TCP parameters with Nagle disabled, matching the other
// implementations. Leaving Nagle on would let it coalesce sends at depth and
// measure the algorithm instead of the transport.
static nw_parameters_t plain_tcp(void) {
	return nw_parameters_create_secure_tcp(
	    NW_PARAMETERS_DISABLE_PROTOCOL,
	    ^(nw_protocol_options_t options) {
		    nw_tcp_options_set_no_delay(options, true);
	    });
}

// wait_ready blocks until the connection is usable, and fails loudly rather
// than silently measuring a connection that never came up.
static void wait_ready(nw_connection_t conn) {
	dispatch_semaphore_t ready = dispatch_semaphore_create(0);
	__block bool ok = false;
	nw_connection_set_state_changed_handler(
	    conn, ^(nw_connection_state_t state, nw_error_t error) {
		    if (state == nw_connection_state_ready) {
			    ok = true;
			    dispatch_semaphore_signal(ready);
		    } else if (state == nw_connection_state_failed ||
			       state == nw_connection_state_cancelled) {
			    dispatch_semaphore_signal(ready);
		    }
		    (void)error;
	    });
	nw_connection_start(conn);
	if (dispatch_semaphore_wait(ready,
				    dispatch_time(DISPATCH_TIME_NOW, 10ll * NSEC_PER_SEC)) != 0) {
		die("connection not ready after 10s");
	}
	if (!ok) {
		die("connection failed");
	}
}

// echo_loop reads whatever arrives and writes it straight back. The received
// dispatch_data_t is handed to the send unchanged: it is immutable and
// reference counted, so echoing copies nothing.
static void echo_loop(nw_connection_t conn) {
	nw_connection_receive(
	    conn, 1, 65536,
	    ^(dispatch_data_t content, nw_content_context_t context, bool complete,
	      nw_error_t error) {
		    (void)context;
		    (void)complete;
		    if (error != NULL || content == NULL) {
			    return;
		    }
		    nw_connection_send(conn, content, NW_CONNECTION_DEFAULT_MESSAGE_CONTEXT,
				       false, ^(nw_error_t send_error) {
					       if (send_error == NULL) {
						       echo_loop(conn);
					       }
				       });
	    });
}

static nw_listener_t serve(const char *port, uint16_t *bound) {
	nw_parameters_t params = plain_tcp();
	nw_listener_t listener = nw_listener_create_with_port(port, params);
	nw_release(params);
	if (listener == NULL) {
		die("cannot create listener");
	}
	// A serial queue, matching the Go and Swift clients. nw_connection
	// expects a serial queue; handing it the global concurrent queue lets
	// its callbacks run concurrently and stalls badly.
	dispatch_queue_t queue = dispatch_queue_create("netperfbench.server", DISPATCH_QUEUE_SERIAL);
	nw_listener_set_queue(listener, queue);

	dispatch_semaphore_t up = dispatch_semaphore_create(0);
	nw_listener_set_state_changed_handler(
	    listener, ^(nw_listener_state_t state, nw_error_t error) {
		    (void)error;
		    if (state == nw_listener_state_ready) {
			    dispatch_semaphore_signal(up);
		    }
	    });
	nw_listener_set_new_connection_handler(listener, ^(nw_connection_t conn) {
		nw_connection_set_queue(conn, queue);
		nw_retain(conn);
		nw_connection_set_state_changed_handler(
		    conn, ^(nw_connection_state_t state, nw_error_t error) {
			    (void)error;
			    if (state == nw_connection_state_ready) {
				    echo_loop(conn);
			    }
		    });
		nw_connection_start(conn);
	});
	nw_listener_start(listener);
	if (dispatch_semaphore_wait(up, dispatch_time(DISPATCH_TIME_NOW, 10ll * NSEC_PER_SEC)) != 0) {
		die("listener not ready after 10s");
	}
	*bound = nw_listener_get_port(listener);
	return listener;
}

static nw_connection_t dial(const char *host, const char *port) {
	nw_endpoint_t endpoint = nw_endpoint_create_host(host, port);
	nw_parameters_t params = plain_tcp();
	nw_connection_t conn = nw_connection_create(endpoint, params);
	nw_release(endpoint);
	nw_release(params);
	if (conn == NULL) {
		die("cannot create connection");
	}
	dispatch_queue_t queue = dispatch_queue_create("netperfbench.client", DISPATCH_QUEUE_SERIAL);
	nw_connection_set_queue(conn, queue);
	wait_ready(conn);
	return conn;
}

// round_trip sends n copies of the payload without waiting, then reads all n
// echoes back. Errors are fatal: a benchmark that quietly measures a failed
// transfer is worse than one that stops.
static void round_trip(nw_connection_t conn, dispatch_data_t payload, int n, size_t size,
		       dispatch_semaphore_t sent, dispatch_semaphore_t got) {
	for (int i = 0; i < n; i++) {
		nw_connection_send(conn, payload, NW_CONNECTION_DEFAULT_MESSAGE_CONTEXT, false,
				   ^(nw_error_t error) {
					   if (error != NULL) {
						   die("send failed");
					   }
					   dispatch_semaphore_signal(sent);
				   });
	}
	size_t remaining = size * (size_t)n;
	while (remaining > 0) {
		size_t want = opt_recv_batch ? remaining : size;
		if (want > UINT32_MAX) {
			die("receive length exceeds uint32_t");
		}
		__block size_t received = 0;
		nw_connection_receive(conn, (uint32_t)want, (uint32_t)want,
				      ^(dispatch_data_t content, nw_content_context_t context,
					bool complete, nw_error_t error) {
					      (void)context;
					      (void)complete;
				      if (error != NULL) {
					      die("receive failed");
				      }
				      received = dispatch_data_get_size(content);
				      dispatch_semaphore_signal(got);
				      });
		dispatch_semaphore_wait(got, DISPATCH_TIME_FOREVER);
		if (received != want) {
			die("echo returned wrong length");
		}
		remaining -= want;
	}
	for (int i = 0; i < n; i++) {
		dispatch_semaphore_wait(sent, DISPATCH_TIME_FOREVER);
	}
}

static int cmp_double(const void *a, const void *b) {
	double x = *(const double *)a, y = *(const double *)b;
	return (x > y) - (x < y);
}

static void report(double *lat, int count, double elapsed_us, double cpu, double user,
		   double sys) {
	qsort(lat, (size_t)count, sizeof(double), cmp_double);
	int n = opt_inflight > 1 ? opt_inflight : 1;
	long messages = (long)count * n;
	double pct50 = lat[(int)(0.50 * (count - 1))];
	double mean = 0;
	for (int i = 0; i < count; i++) {
		mean += lat[i];
	}
	mean /= count;
	double var = 0;
	for (int i = 0; i < count; i++) {
		var += (lat[i] - mean) * (lat[i] - mean);
	}
	var /= count;
	double elapsed = elapsed_us / 1e6;
	double mb = (double)messages * opt_size * 2 / elapsed / 1048576.0;

	double loads[3] = {0, 0, 0};
	getloadavg(loads, 3);

	if (opt_json) {
		printf("{\n");
		printf("  \"label\": \"%s\",\n  \"impl\": \"c\",\n", opt_label);
		printf("  \"payload_bytes\": %d,\n  \"round_trips\": %d,\n", opt_size, count);
		printf("  \"inflight\": %d,\n  \"receive_batch\": %s,\n  \"messages\": %ld,\n",
		       n, opt_recv_batch ? "true" : "false", messages);
		printf("  \"repetitions\": %d,\n", opt_repeat);
		printf("  \"p50_us_per_message\": %.4f,\n", pct50 / n);
		printf("  \"messages_per_sec\": %.2f,\n", (double)messages / elapsed);
		printf("  \"elapsed_sec\": %.6f,\n", elapsed);
		printf("  \"throughput_mbps\": %.2f,\n", mb);
		printf("  \"round_trips_per_sec\": %.2f,\n", count / elapsed);
		printf("  \"min_us\": %.4f,\n", lat[0]);
		printf("  \"p50_us\": %.4f,\n", pct50);
		printf("  \"p90_us\": %.4f,\n", lat[(int)(0.90 * (count - 1))]);
		printf("  \"p99_us\": %.4f,\n", lat[(int)(0.99 * (count - 1))]);
		printf("  \"max_us\": %.4f,\n", lat[count - 1]);
		printf("  \"mean_us\": %.4f,\n  \"stddev_us\": %.4f,\n", mean, sqrt(var));
		printf("  \"cpu\": { \"user_us\": %.1f, \"sys_us\": %.1f },\n", user, sys);
		printf("  \"cpu_us_per_message\": %.4f,\n", cpu / (double)messages);
		printf("  \"cpu_busy_fraction\": %.4f,\n", cpu / elapsed_us);
		printf("  \"env\": { \"load_average\": \"{ %.2f %.2f %.2f }\" }\n", loads[0], loads[1],
		       loads[2]);
		printf("}\n");
		return;
	}
	printf("%s: %d batches of %d x %d bytes in %.2fs\n", opt_label, count, n, opt_size, elapsed);
	printf("  latency us   min %.1f  p50 %.1f  max %.1f  mean %.1f\n", lat[0], pct50,
	       lat[count - 1], mean);
	printf("  %.1f us/message, %.1f MB/s\n", pct50 / n, mb);
	printf("  cpu %.1f us/message, %.2f cores busy\n", cpu / (double)messages, cpu / elapsed_us);
}

static void run_client(const char *host, const char *port) {
	nw_connection_t conn = dial(host, port);

	// A non-constant payload keeps any compression or page-dedup shortcut
	// out of the measurement.
	char *buf = malloc((size_t)opt_size);
	for (int i = 0; i < opt_size; i++) {
		buf[i] = (char)(i * 7 + 13);
	}
	// DISPATCH_DATA_DESTRUCTOR_DEFAULT copies the buffer, so the payload
	// owns its bytes and buf is ours to free once the loop is done. The
	// copy happens here, outside the measured loop, and is not part of any
	// reported number.
	dispatch_data_t payload =
	    dispatch_data_create(buf, (size_t)opt_size, NULL, DISPATCH_DATA_DESTRUCTOR_DEFAULT);

	dispatch_semaphore_t sent = dispatch_semaphore_create(0);
	dispatch_semaphore_t got = dispatch_semaphore_create(0);
	int n = opt_inflight > 1 ? opt_inflight : 1;

	for (int i = 0; i < opt_warmup; i++) {
		round_trip(conn, payload, n, (size_t)opt_size, sent, got);
	}

	double *lat = malloc(sizeof(double) * (size_t)opt_count);
	for (int rep = 0; rep < (opt_repeat > 1 ? opt_repeat : 1); rep++) {
		double u0, s0, u1, s1;
		double cpu0 = cpu_us(&u0, &s0);
		double start = now_us();
		for (int i = 0; i < opt_count; i++) {
			double t0 = now_us();
			round_trip(conn, payload, n, (size_t)opt_size, sent, got);
			lat[i] = now_us() - t0;
		}
		double elapsed = now_us() - start;
		double cpu1 = cpu_us(&u1, &s1);
		report(lat, opt_count, elapsed, cpu1 - cpu0, u1 - u0, s1 - s0);
	}

	nw_connection_force_cancel(conn);
	free(lat);
	free(buf);
}

static void usage(void) {
	fprintf(stderr,
		"usage: netperfbench-c [-role both|server|client] [-addr host:port] [-port p]\n"
		"       [-size n] [-n n] [-warmup n] [-inflight n] [-repeat n] [-recv-batch] [-label s] [-json]\n");
	exit(2);
}

int main(int argc, char **argv) {
	static char host[256];
	for (int i = 1; i < argc; i++) {
		const char *a = argv[i];
		bool has = i + 1 < argc;
#define STR(flag, var)                          \
	if (strcmp(a, flag) == 0 && has) {      \
		var = argv[++i];                \
		continue;                       \
	}
#define INT(flag, var)                          \
	if (strcmp(a, flag) == 0 && has) {      \
		var = atoi(argv[++i]);          \
		continue;                       \
	}
		STR("-role", opt_role)
		STR("-port", opt_port)
		STR("-label", opt_label)
		INT("-size", opt_size)
		INT("-n", opt_count)
		INT("-warmup", opt_warmup)
		INT("-inflight", opt_inflight)
		INT("-repeat", opt_repeat)
		if (strcmp(a, "-json") == 0) {
			opt_json = true;
			continue;
		}
		if (strcmp(a, "-recv-batch") == 0) {
			opt_recv_batch = true;
			continue;
		}
		if (strcmp(a, "-addr") == 0 && has) {
			snprintf(host, sizeof host, "%s", argv[++i]);
			char *colon = strrchr(host, ':');
			if (colon != NULL) {
				*colon = '\0';
				opt_port = colon + 1;
			}
			opt_addr = host;
			continue;
		}
		usage();
#undef STR
#undef INT
	}
	if (opt_size <= 0) {
		die("-size must be positive");
	}

	if (strcmp(opt_role, "server") == 0) {
		uint16_t bound = 0;
		serve(opt_port, &bound);
		fprintf(stderr, "listening on port %u (c)\n", bound);
		dispatch_main();
	} else if (strcmp(opt_role, "client") == 0) {
		run_client(opt_addr, opt_port);
	} else {
		uint16_t bound = 0;
		serve("0", &bound);
		char p[16];
		snprintf(p, sizeof p, "%u", bound);
		run_client("127.0.0.1", p);
	}
	return 0;
}
