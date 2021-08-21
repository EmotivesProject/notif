package api

import (
	"net/http"
	"strconv"
)

const (
	baseInt = 10
	bitSize = 64
)

func findPage(r *http.Request) int64 {
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "0"
	}

	page, err := strconv.ParseInt(pageParam, baseInt, bitSize)
	if err != nil {
		return 0
	}

	return page
}
