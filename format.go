package checkpoint

import "strconv"

func itoa(value uint64) string {
	return strconv.FormatUint(value+10, 10)
}
