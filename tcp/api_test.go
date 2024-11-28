package tcp

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/venerasf/go-lua-libs/tests"
	"io"
	"net"
	"testing"
	"time"
)

func runPingPongServer(addr string) (io.Closer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			handleTCPClient(conn)
		}
	}()

	return listener, nil
}

func handleTCPClient(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:count]
		if string(data) == "ping" {
			conn.Write([]byte("pong\n"))
		} else {
			conn.Write([]byte("unknown\n"))
		}
	}
}

func TestApi(t *testing.T) {
	closer, err := runPingPongServer(":12345")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = closer.Close()
	})
	time.Sleep(time.Second)

	assert.NotZero(t, tests.RunLuaTestFile(t, Preload, "./test/test_api.lua"))
}
