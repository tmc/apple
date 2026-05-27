package rdma

import "testing"

func benchmarkPollCQ(RDMACQ, int, *IbvWC) int {
	return 0
}

func benchmarkPostSend(RDMAQP, *IbvSendWR, **IbvSendWR) int {
	return 0
}

func benchmarkPostRecv(RDMAQP, *IbvRecvWR, **IbvRecvWR) int {
	return 0
}

func TestDatapathWrappersAllocs(t *testing.T) {
	var wc IbvWC
	var sendWR IbvSendWR
	var badSend *IbvSendWR
	var recvWR IbvRecvWR
	var badRecv *IbvRecvWR
	poller := IbvCQPoller{cq: 1, fn: benchmarkPollCQ}
	poster := IbvQPPoster{qp: 1, send: benchmarkPostSend, recv: benchmarkPostRecv}

	tests := []struct {
		name string
		call func()
	}{
		{"poll", func() { _ = poller.Poll(1, &wc) }},
		{"post_send", func() { _ = poster.PostSend(&sendWR, &badSend) }},
		{"post_recv", func() { _ = poster.PostRecv(&recvWR, &badRecv) }},
	}
	for _, tt := range tests {
		if allocs := testing.AllocsPerRun(1000, tt.call); allocs != 0 {
			t.Fatalf("%s allocs = %v, want 0", tt.name, allocs)
		}
	}
}

func BenchmarkDatapathWrappers(b *testing.B) {
	var wc IbvWC
	var sendWR IbvSendWR
	var badSend *IbvSendWR
	var recvWR IbvRecvWR
	var badRecv *IbvRecvWR
	poller := IbvCQPoller{cq: 1, fn: benchmarkPollCQ}
	poster := IbvQPPoster{qp: 1, send: benchmarkPostSend, recv: benchmarkPostRecv}

	b.Run("poll", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poller.Poll(1, &wc); got != 0 {
				b.Fatalf("Poll = %d, want 0", got)
			}
		}
	})
	b.Run("post_send", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poster.PostSend(&sendWR, &badSend); got != 0 {
				b.Fatalf("PostSend = %d, want 0", got)
			}
		}
	})
	b.Run("post_recv", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poster.PostRecv(&recvWR, &badRecv); got != 0 {
				b.Fatalf("PostRecv = %d, want 0", got)
			}
		}
	})
}
