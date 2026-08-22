package rdma

import "testing"

func TestDatapathWrappersAllocs(t *testing.T) {
	var wc IbvWC
	var sendWR IbvSendWR
	var badSend *IbvSendWR
	var recvWR IbvRecvWR
	var badRecv *IbvRecvWR
	fn := buildDatapathTestFunction(t)
	poller := IbvCQPoller{cq: 1, fnPtr: fn}
	poster := IbvQPPoster{qp: 1, sendPtr: fn, recvPtr: fn}

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
	fn := buildDatapathTestFunction(b)
	poller := IbvCQPoller{cq: 1, fnPtr: fn}
	poster := IbvQPPoster{qp: 1, sendPtr: fn, recvPtr: fn}

	b.Run("poll", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poller.Poll(1, &wc); got != 2 {
				b.Fatalf("Poll = %d, want 2", got)
			}
		}
	})
	b.Run("post_send", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poster.PostSend(&sendWR, &badSend); got != 1 {
				b.Fatalf("PostSend = %d, want 1", got)
			}
		}
	})
	b.Run("post_recv", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if got := poster.PostRecv(&recvWR, &badRecv); got != 1 {
				b.Fatalf("PostRecv = %d, want 1", got)
			}
		}
	})
}
