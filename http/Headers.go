package http

import (
	"log"
	"strings"
)

type Headers map[string]string

// 添加一个Header
func (h Headers) Add(key, value string) {
	h[key] = value
}

// 提取头
func (h Headers) ExtractHeaders(str, sep string) bool {
	arr := strings.Split(str, sep)
	for _, line := range arr {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Println("header string is not illegal")
			return false
		}
		h.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}
	return true
}

func (h Headers) Exists(key, value string) bool {
	if v, ok := h[key]; ok {
		if v == value {
			return true
		}
	}
	return false
}
