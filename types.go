// SPDX-License-Identifier: BSD-2-Clause

package httplog

import (
	"time"
)

// Request is the structure any formatter will be handed when time to log comes
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
