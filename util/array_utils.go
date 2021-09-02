package util

import "pairinator/models"

func Contains(list []string, value string) bool {
	for _, k := range list {
		if value == k {
			return true
		}
	}
	return false
}

func RemoveOne(s []models.Member, i int) []models.Member {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveTwo(s []models.Member, i int, j int) []models.Member {
	s[i] = s[len(s)-1]
	s[j] = s[len(s)-1]
	return s[:len(s)-2]
}