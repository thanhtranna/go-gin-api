package time_parse

import (
	"math"
	"net/http"
	"time"
)

var (
	cst *time.Location
)

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	// if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
	// 	panic(err)
	// }
	if cst, err = time.LoadLocation("Asia/Ho_Chi_Minh"); err != nil {
		panic(err)
	}
}

// RFC3339ToCSTLayout convert rfc3339 value to china standard time layout
// 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

// CSTLayoutString formatting time
// Return the time in "2006-01-02 15:04:05" format
func CSTLayoutString() string {
	ts := time.Now()
	return ts.In(cst).Format(CSTLayout)
}

// ParseCSTInLocation Format time
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

// CSTLayoutStringToUnix returns unix timestamp
// 2020-01-24 21:11:11 => 1579871471
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil
}

// GMTLayoutString formatting time
// Return the time in "Mon, 02 Jan 2006 15:04:05 GMT" format
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation format time
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

// SubInLocation calculate the time difference
func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}
