// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"net/http"
	"time"
)

// Logger is any user-supplied function to process one LogRequest.
type Logger func(r LogRequest)

type logHandler struct {
	handler http.Handler
	logger  Logger
}

// LogHandler makes a http.Handler (middleware for the Go HTTP
// server). For each HTTP request, it calls the user-supplied handler.
// At the end of the request, it additionally calls the user-supplied
// logger to log the request.
//
// The logger gets a LogRequest. Usually you'll want to call one of
// the log formatting functions in this package or roll your own log
// formatter. Then write the resulting log entry somewhere.
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
