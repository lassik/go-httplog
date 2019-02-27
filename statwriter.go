// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"net/http"
)

// To conform to the http.ResponseWriter interface we cannot use
// pointer receivers to modify the statResponseWriter struct. That's
// why size and status need to be in this indirect helper struct.
type statResponseWriterStat struct {
	size   int
	status int
}

// statResponseWriter is a http.ResponseWriter wrapper that keeps
// track of the response size and the HTTP status code.
type statResponseWriter struct {
	w    http.ResponseWriter
	stat *statResponseWriterStat
}

func makeStatResponseWriter(w http.ResponseWriter) statResponseWriter {
	return statResponseWriter{w, &statResponseWriterStat{0, http.StatusOK}}
}

func (wStat statResponseWriter) Write(p []byte) (int, error) {
	n, err := wStat.w.Write(p)
	wStat.stat.size += n
	return n, err
}

func (wStat statResponseWriter) WriteHeader(status int) {
	wStat.stat.status = status
	wStat.w.WriteHeader(status)
}

func (wStat statResponseWriter) Header() http.Header {
	return wStat.w.Header()
}

func (wStat statResponseWriter) GetSizeAndStatus() (int, int) {
	return wStat.stat.size, wStat.stat.status
}
