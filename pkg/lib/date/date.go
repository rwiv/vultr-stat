package date

import "time"

const (
	isoLayout = "2006-01-02T15:04:05"
)

func ByIsoString(string string) (time.Time, error) {
	return time.Parse(isoLayout, string)
}

func ToIsoString(time time.Time) string {
	return time.Format(isoLayout)
}

func ByRFC3339String(string string) (time.Time, error) {
	return time.Parse(time.RFC3339, string)
}

func ToRFC3339String(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func ToPrettyString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
