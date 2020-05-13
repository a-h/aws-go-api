package log

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// Responses logs the status, duration and length of URLs.
func Responses(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer Default.Sync()
		var written int64
		var status = -1

		wp := writerProxy{
			h: func() http.Header {
				return w.Header()
			},
			w: func(bytes []byte) (int, error) {
				bw, err := w.Write(bytes)
				written += int64(bw)
				return bw, err
			},
			wh: func(code int) {
				status = code
				w.WriteHeader(code)
			},
		}

		start := time.Now()
		next.ServeHTTP(wp, r)
		duration := time.Now().Sub(start)

		// Use default status.
		if status == -1 {
			status = 200
		}

		logResponse(r.Method, r.URL.String(), status, written, duration)
	})
}

func logResponse(method, url string, status int, length int64, duration time.Duration) {
	Default.Info("response", zap.String("method", method),
		zap.String("url", url),
		zap.Int("status", status),
		zap.Duration("duration", duration),
		zap.Int64("len", length),
		zap.Int("http_"+strconv.Itoa(status/100)+"xx", 1))
}

type writerProxy struct {
	h  func() http.Header
	w  func(bytes []byte) (int, error)
	wh func(status int)
}

func (wp writerProxy) Header() http.Header {
	return wp.h()
}

func (wp writerProxy) Write(bytes []byte) (int, error) {
	return wp.w(bytes)
}

func (wp writerProxy) WriteHeader(status int) {
	wp.wh(status)
}
