package utils

import "strings"

func SplitHost(hostport string) string {
	index := strings.Index(hostport, ":")
	if index != -1 {
		return hostport[:index]
	}
	return hostport
}
