package utils

import (
	"encoding/json"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func JSON(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func ValidatePage(page, size int) (int, int) {
	if page <= 0 {
		page = 1
	}

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 20
	}
	return page, size
}
