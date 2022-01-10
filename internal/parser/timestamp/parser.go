package timestamp

import (
	"regexp"
	"strconv"
	"time"

	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

var (
	rxTimestamp = regexp.MustCompile(`^(\d{10})(\d{3})?(\d{3})?(?:\.|\s|$)`)
)

func Parse(cfg *setting.Configuration, str string) date.Date {
	parts := rxTimestamp.FindStringSubmatch(str)
	if len(parts) > 0 {
		seconds, _ := strconv.ParseInt(parts[1], 10, 64)
		millis, _ := strconv.ParseInt(parts[2], 10, 64)
		micros, _ := strconv.ParseInt(parts[3], 10, 64)
		nanos := micros*1_000 + millis*1_000_000

		t := time.Unix(seconds, nanos)
		return date.Date{Time: t, Period: date.Time}
	}

	return date.Date{}
}