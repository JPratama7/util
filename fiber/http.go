// Copied from https://github.com/valyala/fasthttp/blob/master/fasthttpadaptor/adaptor.go#L31

package fiber

import (
	"bufio"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"io"
	"net"
	"net/http"
	"sync"
)

func ResponseWriteFromCtx(c *fiber.Ctx) (w http.ResponseWriter) {
	return &netHTTPResponseWriter{
		w:   c.Response().BodyWriter(),
		ctx: c.Context(),
	}
}

type netHTTPResponseWriter struct {
	statusCode int
	h          http.Header
	w          io.Writer
	ctx        *fasthttp.RequestCtx
}

func (w *netHTTPResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *netHTTPResponseWriter) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *netHTTPResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *netHTTPResponseWriter) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

type wrappedConn struct {
	net.Conn

	wg   sync.WaitGroup
	once sync.Once
}

func (c *wrappedConn) Close() (err error) {
	c.once.Do(func() {
		err = c.Conn.Close()
		c.wg.Done()
	})
	return
}

func (w *netHTTPResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// Hijack assumes control of the connection, so we need to prevent fasthttp from closing it or
	// doing anything else with it.
	w.ctx.HijackSetNoResponse(true)

	conn := &wrappedConn{Conn: w.ctx.Conn()}
	conn.wg.Add(1)
	w.ctx.Hijack(func(net.Conn) {
		conn.wg.Wait()
	})

	bufW := bufio.NewWriter(conn)

	// Write any unflushed body to the hijacked connection buffer.
	unflushedBody := w.ctx.Response.Body()
	if len(unflushedBody) > 0 {
		if _, err := bufW.Write(unflushedBody); err != nil {
			conn.Close()
			return nil, nil, err
		}
	}

	return conn, &bufio.ReadWriter{Reader: bufio.NewReader(conn), Writer: bufW}, nil
}
