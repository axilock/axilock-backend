package util

import (
	"strings"
)

func GetDomain(email string) string {
	at := strings.LastIndex(email, "@")
	if at >= 0 {
		_, domain := email[:at], email[at+1:]
		return domain
	}
	return ""
}
