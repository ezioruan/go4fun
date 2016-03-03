package fast_http

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"github.com/valyala/fasthttp"
	"strings"
	"testing"
)

func req(b *testing.B, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		b.Fatal(err)
	}
	return req
}

func BenchmarkJsonHandle(b *testing.B) {
	b.ReportAllocs()
	r := req(b, "GET / HTTP/1.0\r\n\r\n")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		jsonHandle(rw, r)
	}
}

func BenchmarkFastJsonHandle(b *testing.B) {
	b.ReportAllocs()
	r := req(b, "GET / HTTP/1.0\r\n\r\n")
    fasthttp.
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		fastJsonHandle1(ctx)
	}
}
