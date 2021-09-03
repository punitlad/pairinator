package util

func Contains(list []string, value string) bool {
	for _, k := range list {
		if value == k {
			return true
		}
	}
	return false
}