// Copyright 2014 Mahmud Ridwan
// Copyright 2019 Lassi Kortela
// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"strconv"
)

// CommonLogLine returns a log entry in Apache Common Log Format (CLF).
//
// See http://httpd.apache.org/docs/2.2/logs.html#common
//
// Always sets the ident field of the log to -
func CommonLogLine(r LogRequest) []byte {
	buf := make([]byte, 0, 128)
	buf = append(buf, r.Host...)
	buf = append(buf, " - "...)
	buf = append(buf, r.User...)
	buf = append(buf, " ["...)
	buf = append(buf, r.StartTime.Format("02/Jan/2006:15:04:05 -0700")...)
	buf = append(buf, `] "`...)
	buf = append(buf, r.Method...)
	buf = append(buf, " "...)
	buf = AppendQuoted(buf, r.URI)
	buf = append(buf, " "...)
	buf = append(buf, r.Proto...)
	buf = append(buf, `" `...)
	buf = append(buf, strconv.Itoa(r.Status)...)
	buf = append(buf, " "...)
	buf = append(buf, strconv.Itoa(r.Size)...)
	return buf
}

// CombinedLogLine returns a log entry in Apache Combined Log Format.
//
// See http://httpd.apache.org/docs/2.2/logs.html#combined
//
// Always sets the ident field of the log to -
func CombinedLogLine(r LogRequest) []byte {
	buf := CommonLogLine(r)
	buf = append(buf, ` "`...)
	buf = AppendQuoted(buf, r.Referer)
	buf = append(buf, `" "`...)
	buf = AppendQuoted(buf, r.UserAgent)
	buf = append(buf, '"')
	return buf
}
