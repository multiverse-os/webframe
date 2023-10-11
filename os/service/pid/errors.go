package pid

import "errors"

// Segregated to simplify localization
var (
	errStalePid    = errors.New("stale pid check open failed")
	errFileLocked  = errors.New("pid file locked by running process")
	errOpenFailed  = errors.New("failed to open pid file")
	errWriteFailed = errors.New("failed to write pid to file")
	errShortWrite  = errors.New("short write failure writing pid")
	errCleanFailed = errors.New("failed to clean up pid file")
)
