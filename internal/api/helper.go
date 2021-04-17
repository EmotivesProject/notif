package api

import (
	"net/http"
	"strconv"
)

func findPage(r *http.Request) int64 {
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "0"
	}
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		return 0
	}
	return page
}
