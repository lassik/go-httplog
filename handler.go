// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"net/http"
	"time"
)

type Logger func(r LogRequest)

// LogHandler is a http.Handler wrapper that calls a user-defined
// logger function after each request. The logger function can call
// the utility functions in this package to help produce its log
// message.
type logHandler struct {
	handler http.Handler
	logger  Logger
}

func LogHandler(handler http.Handler, logger Logger) logHandler {
	return logHandler{handler, logger}
}

func (lh logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lr := LogRequest{StartTime: time.Now()}
	wStat := makeStatResponseWriter(w)
	lh.handler.ServeHTTP(wStat, r)
	lr.EndTime = time.Now()
	lr.User = ParseUser(r)
	lr.Host, lr.URI = ParseHostAndURI(r)
	lr.Method = r.Method
	lr.Proto = r.Proto
	lr.Referer = r.Referer()
	lr.UserAgent = r.UserAgent()
	lr.Size, lr.Status = wStat.GetSizeAndStatus()
	lh.logger(lr)
}
