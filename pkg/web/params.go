package web

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetQueryString(q url.Values, name string, defaultVal string) string {
	val := q.Get(name)
	if val == "" {
		return defaultVal
	}

	return val
}

func GetQueryTime(q url.Values, name string) time.Time {
	val := q.Get(name)
	if val == "" {
		return time.Time{}
	}

	date, err := time.Parse(time.DateOnly, val)
	if err != nil {
		return time.Time{}
	}

	return date
}

func GetQuerySlice(q url.Values, name string) []string {
	val := q.Get(name)
	if val == "" {
		return nil
	}

	return strings.Split(val, ",")
}

func GetQueryInt(q url.Values, name string, defaultVal int64) int {
	val, err := strconv.ParseInt(q.Get(name), 10, 64)
	if err != nil {
		val = defaultVal
	}

	return int(val)
}

func GetQueryBool(q url.Values, name string, defaultVal bool) bool {
	val, err := strconv.ParseBool(q.Get(name))
	if err != nil {
		val = defaultVal
	}

	return val
}
