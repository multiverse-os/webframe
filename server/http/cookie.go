package http

import (
	"time"
)

func HTMLCookie(key, value string) string {
	return `<meta http-equiv="set-cookie" content="` + key + `=` + value + `;" />`
}

func HTMLCookieWithExpiration(key, value string, length time.Duration) string {
	// TODO: If error set to some reasonable default or return HTMLCookie without
	// expire time
	return `<meta http-equiv="set-cookie" content="` + key + `=` + value + `; expires=` + time.Now().UTC().Add(length).Format("Monday, 02-Jan-2006 15:04:05 MST") + `;" />`
}
