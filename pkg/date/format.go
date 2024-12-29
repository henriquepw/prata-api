package date

import "time"

const (
	BRHour = "15:04"
	BRDate = "02/01/2006"
)

func LocalFormat(dt time.Time, format string) string {
	return dt.Add(-time.Hour * 3).Format(format)
}

func FormatToISO(dt time.Time) string {
	return dt.UTC().Format(time.RFC3339)
}
