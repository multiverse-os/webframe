package filesystem

import "errors"

var (
	errWD         = errors.New("failed to determine working directory")
	errCachePath  = errors.New("failed to determine user cache")
	errIsLocked   = errors.New("file is already locked")
	errTimeout    = errors.New("timeout exceeded when acquiring lock")
	errConfigPath = errors.New("failed to determine user config path")
	errDataPath   = errors.New("failed to determine user data path")
)
