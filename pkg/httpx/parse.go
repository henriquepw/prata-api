package httpx

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
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

func GetBodyRequest[T any](r *http.Request) (T, error) {
	defer r.Body.Close()

	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, errorx.InvalidJSON()
	}

	return data, nil
}

func GetJsonResponse[T any](r *http.Response) (T, error) {
	defer r.Body.Close()

	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, errorx.InvalidJSON()
	}

	return data, nil
}
