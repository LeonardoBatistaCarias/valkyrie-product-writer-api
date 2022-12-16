package constants

import "time"

const (
	MAX_CONN            = 50
	HEALTH_CHECK_PERIOD = 1 * time.Minute
	MAX_CONN_IDLE_TIME  = 1 * time.Minute
	MAX_CONN_LIFETIME   = 3 * time.Minute
	MIN_CONN            = 10
	LAZY_CONNECT        = false
)
