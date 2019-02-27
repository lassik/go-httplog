// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"time"
)

// LogRequest holds all the logging information pertaining to a single
// HTTP request. Log formatters turn this struct into a log entry in
// the desired format. This struct holds the information in a pre-parsed
// format to make it as easy as possible to write log formatters.
type LogRequest struct {
	StartTime time.Time
	EndTime   time.Time
	User      string
	Host      string
	Method    string
	URI       string
	Proto     string
	Referer   string
	UserAgent string
	Size      int
	Status    int
}
