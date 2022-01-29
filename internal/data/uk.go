// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uk_Locale = merge(nil, LocaleData{
	Name:                  "uk",
	DateOrder:             "DMY",
	Charset:               []rune(`cgtuzабвгдежзиклмнопрстухцчьюяєі`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^година(\z|[^\pL\pM\d]|_)`), "${1}1 година${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^годину(\z|[^\pL\pM\d]|_)`), "${1}1 годину${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^хвилину(\z|[^\pL\pM\d]|_)`), "${1}1 хвилину${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^секунду(\z|[^\pL\pM\d]|_)`), "${1}1 секунду${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)кілька секунд(\z|[^\pL\pM\d]|_)`), "${1}44 секунди${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)кілька хвилин(\z|[^\pL\pM\d]|_)`), "${1}2 хвилини${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)опів на(\z|[^\pL\pM\d]|_)`), "${1}30 хвилин${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)за чверть(\z|[^\pL\pM\d]|_)`), "${1}15 хвилин${2}"},
	},
	Translations: map[string]string{
		"понеділок": "monday",
		"листопада": "november",
		"п'ятниця":  "friday",
		"протягом":  "in",
		"березень":  "march",
		"листопад":  "november",
		"вересень":  "september",
		"вівторок":  "tuesday",
		"близько":   "",
		"квітень":   "april",
		"серпень":   "august",
		"грудень":   "december",
		"червень":   "june",
		"березня":   "march",
		"травень":   "may",
		"хвилина":   "minute",
		"хвилини":   "minute",
		"хвилину":   "minute",
		"місяців":   "month",
		"жовтень":   "october",
		"секунда":   "second",
		"секунди":   "second",
		"секунду":   "second",
		"вересня":   "september",
		"тиждень":   "week",
		"квітня":    "april",
		"серпня":    "august",
		"грудня":    "december",
		"лютого":    "february",
		"година":    "hour",
		"години":    "hour",
		"годину":    "hour",
		"січень":    "january",
		"липень":    "july",
		"червня":    "june",
		"травня":    "may",
		"хвилин":    "minute",
		"місяць":    "month",
		"місяці":    "month",
		"жовтня":    "october",
		"субота":    "saturday",
		"секунд":    "second",
		"неділя":    "sunday",
		"четвер":    "thursday",
		"середа":    "wednesday",
		"тижнів":    "week",
		"назад":     "ago",
		"лютии":     "february",
		"годин":     "hour",
		"січня":     "january",
		"липня":     "july",
		"тижні":     "week",
		"років":     "year",
		"тому":      "ago",
		"квіт":      "april",
		"серп":      "august",
		"день":      "day",
		"днів":      "day",
		"груд":      "december",
		"черв":      "june",
		"трав":      "may",
		"лист":      "november",
		"жовт":      "october",
		"роки":      "year",
		"кві":       "april",
		"сер":       "august",
		"дні":       "day",
		"гру":       "december",
		"лют":       "february",
		"gmt":       "gmt",
		"год":       "hour",
		"січ":       "january",
		"лип":       "july",
		"чер":       "june",
		"бер":       "march",
		"тра":       "may",
		"пон":       "monday",
		"міс":       "month",
		"лис":       "november",
		"жов":       "october",
		"суб":       "saturday",
		"сек":       "second",
		"вер":       "september",
		"нед":       "sunday",
		"чет":       "thursday",
		"вів":       "tuesday",
		"utc":       "utc",
		"тиж":       "week",
		"рік":       "year",
		"об":        "",
		"am":        "am",
		"дп":        "am",
		"пт":        "friday",
		"хв":        "minute",
		"пн":        "monday",
		"pm":        "pm",
		"пп":        "pm",
		"сб":        "saturday",
		"нд":        "sunday",
		"чт":        "thursday",
		"вт":        "tuesday",
		"ср":        "wednesday",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		"о":         "",
		"і":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"д":         "day",
		"г":         "hour",
		"с":         "second",
		"р":         "year",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"наступного місяця": "in 1 month",
		"наступного тижня":  "in 1 week",
		"минулого місяця":   "1 month ago",
		"наступного року":   "in 1 year",
		"минулого тижня":    "1 week ago",
		"цієі хвилини":      "0 minute ago",
		"цього місяця":      "0 month ago",
		"цієі години":       "0 hour ago",
		"цього тижня":       "0 week ago",
		"цього року":        "0 year ago",
		"позавчора":         "2 day ago",
		"сьогодні":          "0 day ago",
		"завтра":            "in 1 day",
		"зараз":             "0 second ago",
		"вчора":             "1 day ago",
		"учора":             "1 day ago",
		"торік":             "1 year ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)через (\d+) хвилини`), "in $1 minute"},
		{regexp.MustCompile(`(?i)через (\d+) хвилину`), "in $1 minute"},
		{regexp.MustCompile(`(?i)через (\d+) секунди`), "in $1 second"},
		{regexp.MustCompile(`(?i)через (\d+) секунду`), "in $1 second"},
		{regexp.MustCompile(`(?i)через (\d+) тиждень`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) хвилини тому`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) хвилину тому`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) секунди тому`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) секунду тому`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) тиждень тому`), "$1 week ago"},
		{regexp.MustCompile(`(?i)через (\d+) години`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) годину`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) місяць`), "in $1 month"},
		{regexp.MustCompile(`(?i)через (\d+) місяця`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) години тому`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) годину тому`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) місяць тому`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) місяця тому`), "$1 month ago"},
		{regexp.MustCompile(`(?i)через (\d+) тижня`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) тижня тому`), "$1 week ago"},
		{regexp.MustCompile(`(?i)через (\d+) день`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) року`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) день тому`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) року тому`), "$1 year ago"},
		{regexp.MustCompile(`(?i)через (\d+) дня`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) год`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) міс`), "in $1 month"},
		{regexp.MustCompile(`(?i)через (\d+) тиж`), "in $1 week"},
		{regexp.MustCompile(`(?i)через (\d+) рік`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) дня тому`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) год тому`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) міс тому`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) тиж тому`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) рік тому`), "$1 year ago"},
		{regexp.MustCompile(`(?i)через (\d+) дн`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) хв`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) дн тому`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) хв тому`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)через (\d+) с`), "in $1 second"},
		{regexp.MustCompile(`(?i)через (\d+) р`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) д тому`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) с тому`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) р тому`), "$1 year ago"},
		{regexp.MustCompile(`(?i)за (\d+) год`), "in $1 hour"},
		{regexp.MustCompile(`(?i)за (\d+) міс`), "in $1 month"},
		{regexp.MustCompile(`(?i)за (\d+) тиж`), "in $1 week"},
		{regexp.MustCompile(`(?i)за (\d+) хв`), "in $1 minute"},
		{regexp.MustCompile(`(?i)за (\d+) д`), "in $1 day"},
		{regexp.MustCompile(`(?i)за (\d+) с`), "in $1 second"},
		{regexp.MustCompile(`(?i)за (\d+) р`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(через \d+ секунди|через \d+ секунду|через \d+ тиждень|через \d+ хвилини|через \d+ хвилину|\d+ секунди тому|\d+ секунду тому|\d+ тиждень тому|\d+ хвилини тому|\d+ хвилину тому|через \d+ години|через \d+ годину|через \d+ місяць|через \d+ місяця|\d+ години тому|\d+ годину тому|\d+ місяць тому|\d+ місяця тому|через \d+ тижня|\d+ тижня тому|через \d+ день|через \d+ року|\d+ день тому|\d+ року тому|через \d+ год|через \d+ дня|через \d+ міс|через \d+ рік|через \d+ тиж|\d+ год тому|\d+ дня тому|\d+ міс тому|\d+ рік тому|\d+ тиж тому|через \d+ дн|через \d+ хв|\d+ дн тому|\d+ хв тому|через \d+ р|через \d+ с|\d+ д тому|\d+ р тому|\d+ с тому|за \d+ год|за \d+ міс|за \d+ тиж|за \d+ хв|за \d+ д|за \d+ р|за \d+ с)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(через \d+ секунди|через \d+ секунду|через \d+ тиждень|через \d+ хвилини|через \d+ хвилину|\d+ секунди тому|\d+ секунду тому|\d+ тиждень тому|\d+ хвилини тому|\d+ хвилину тому|через \d+ години|через \d+ годину|через \d+ місяць|через \d+ місяця|\d+ години тому|\d+ годину тому|\d+ місяць тому|\d+ місяця тому|через \d+ тижня|\d+ тижня тому|через \d+ день|через \d+ року|\d+ день тому|\d+ року тому|через \d+ год|через \d+ дня|через \d+ міс|через \d+ рік|через \d+ тиж|\d+ год тому|\d+ дня тому|\d+ міс тому|\d+ рік тому|\d+ тиж тому|через \d+ дн|через \d+ хв|\d+ дн тому|\d+ хв тому|через \d+ р|через \d+ с|\d+ д тому|\d+ р тому|\d+ с тому|за \d+ год|за \d+ міс|за \d+ тиж|за \d+ хв|за \d+ д|за \d+ р|за \d+ с)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(наступного місяця|наступного тижня|минулого місяця|наступного року|минулого тижня|цього місяця|цієі хвилини|цього тижня|цієі години|цього року|листопада|позавчора|понеділок|березень|вересень|вівторок|листопад|п'ятниця|протягом|сьогодні|березня|близько|вересня|грудень|жовтень|квітень|місяців|секунда|секунди|секунду|серпень|тиждень|травень|хвилина|хвилини|хвилину|червень|година|години|годину|грудня|жовтня|завтра|квітня|липень|лютого|місяць|місяці|неділя|секунд|середа|серпня|субота|січень|тижнів|травня|хвилин|червня|четвер|вчора|годин|зараз|липня|лютии|назад|років|січня|тижні|торік|учора|груд|день|днів|жовт|квіт|лист|роки|серп|тому|трав|черв|gmt|utc|бер|вер|вів|год|гру|дні|жов|кві|лип|лис|лют|міс|нед|пон|рік|сек|сер|суб|січ|тиж|тра|чер|чет|\+|\.|\[|\]|\||am|pm|вт|дп|нд|об|пн|пп|пт|сб|ср|хв|чт| |'|,|-|/|:|;|@|z|г|д|о|р|с|і)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
