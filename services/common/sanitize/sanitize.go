package sanitize

import "strings"

func String(value string) string {
	return strings.Trim(value, " ");
}
