package main

// Standard library reference implementation. It is the control: whatever
// Network.framework costs relative to a plain socket shows up as the
// difference between this transport and the nw one.

import (
	"io"
	"net"
)

type stdServer struct {
	ln net.Listener
}

func serveStd(port string) (echoServer, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return nil, err
	}
	s := &stdServer{ln: ln}
	go s.accept()
	return s, nil
}

func (s *stdServer) accept() {
	for {
		c, err := s.ln.Accept()
		if err != nil {
			return // listener closed
		}
		go func() {
			defer c.Close()
			if tc, ok := c.(*net.TCPConn); ok {
				tc.SetNoDelay(true)
			}
			io.Copy(c, c)
		}()
	}
}

func (s *stdServer) Port() string {
	_, port, _ := net.SplitHostPort(s.ln.Addr().String())
	return port
}

func (s *stdServer) Close() { s.ln.Close() }

type stdClient struct {
	conn net.Conn
	buf  []byte
}

func dialStd(addr string) (echoClient, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if tc, ok := c.(*net.TCPConn); ok {
		tc.SetNoDelay(true)
	}
	return &stdClient{conn: c}, nil
}

// RoundTrip writes n copies of buf and reads all n echoes back. The writes
// run on their own goroutine so that a batch larger than the socket buffers
// cannot deadlock against the reads.
func (c *stdClient) RoundTrip(buf []byte, n int) error {
	total := len(buf) * n
	if len(c.buf) < total {
		c.buf = make([]byte, total)
	}
	werr := make(chan error, 1)
	go func() {
		for range n {
			if _, err := c.conn.Write(buf); err != nil {
				werr <- err
				return
			}
		}
		werr <- nil
	}()
	if _, err := io.ReadFull(c.conn, c.buf[:total]); err != nil {
		return err
	}
	return <-werr
}

func (c *stdClient) Close() { c.conn.Close() }
