// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ky_Locale = merge(nil, LocaleData{
	Name:      "ky",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzабвгдежзийклмнопрстуфчшыьэюяңүӊө`),
	Translations: map[string]string{
		"түштөн кииинки": "pm",
		"беишемби":       "thursday",
		"дүишөмбү":       "monday",
		"жекшемби":       "sunday",
		"сентябрь":       "september",
		"шаршемби":       "wednesday",
		"шеишемби":       "tuesday",
		"декабрь":        "december",
		"октябрь":        "october",
		"февраль":        "february",
		"август":         "august",
		"апрель":         "april",
		"ишемби":         "saturday",
		"ноябрь":         "november",
		"секунд":         "second",
		"январь":         "january",
		"мүнөт":          "minute",
		"таңкы":          "am",
		"апта":           "week",
		"беиш":           "thursday",
		"жума":           "friday",
		"июль":           "july",
		"июнь":           "june",
		"март":           "march",
		"саат":           "hour",
		"шарш":           "wednesday",
		"шеиш":           "tuesday",
		"gmt":            "gmt",
		"utc":            "utc",
		"авг":            "august",
		"апр":            "april",
		"апт":            "week",
		"дек":            "december",
		"дүи":            "monday",
		"жек":            "sunday",
		"жыл":            "year",
		"ишм":            "saturday",
		"июл":            "july",
		"июн":            "june",
		"күн":            "day",
		"маи":            "may",
		"мар":            "march",
		"мүн":            "minute",
		"ноя":            "november",
		"окт":            "october",
		"сек":            "second",
		"сен":            "september",
		"фев":            "february",
		"янв":            "january",
		"am":             "am",
		"pm":             "pm",
		"аи":             "month",
		"ст":             "hour",
		"тк":             "pm",
		"тң":             "am",
		" ":              " ",
		"'":              "",
		"+":              "+",
		",":              "",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"z":              "z",
		"|":              "",
		"ж":              "year",
		"м":              "minute",
	},
	RelativeType: map[string]string{
		"келерки аптада": "in 1 week",
		"ушул мүнөттө":   "0 minute ago",
		"өткөн аптада":   "1 week ago",
		"келерки апт":    "in 1 week",
		"ушул аптада":    "0 week ago",
		"ушул саатта":    "0 hour ago",
		"эмдиги аида":    "in 1 month",
		"эмдиги жылы":    "in 1 year",
		"өткөн аида":     "1 month ago",
		"өткөн апт":      "1 week ago",
		"бул аида":       "0 month ago",
		"ушул апт":       "0 week ago",
		"былтыр":         "1 year ago",
		"быиыл":          "0 year ago",
		"бүгүн":          "0 day ago",
		"кечээ":          "1 day ago",
		"эртеӊ":          "in 1 day",
		"азыр":           "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секунддан кииин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) мүнөттөн кииин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) аптадан кииин`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) сааттан кииин`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) жылдан кииин`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) күндөн кииин`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) секунд мурун`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) аидан кииин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) мүнөт мурун`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) апта мурун`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) саат кииин`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) саат мурун`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) аид кииин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) апт кииин`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) апт мурун`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) жыл кииин`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) жыл мурун`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) күн кииин`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) күн мурун`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) мүн кииин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) мүн мурун`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек кииин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) сек мурун`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) аи мурун`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) аид киин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) мүн киин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) мүн мурн`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек киин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) сек мурн`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) аи мурн`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) с киин`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) с мурн`), "$1 hour ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ секунддан кииин|\d+ мүнөттөн кииин|\d+ аптадан кииин|\d+ сааттан кииин|\d+ жылдан кииин|\d+ күндөн кииин|\d+ секунд мурун|\d+ аидан кииин|\d+ мүнөт мурун|\d+ апта мурун|\d+ саат кииин|\d+ саат мурун|\d+ аид кииин|\d+ апт кииин|\d+ апт мурун|\d+ жыл кииин|\d+ жыл мурун|\d+ күн кииин|\d+ күн мурун|\d+ мүн кииин|\d+ мүн мурун|\d+ сек кииин|\d+ сек мурун|\d+ аи мурун|\d+ аид киин|\d+ мүн киин|\d+ мүн мурн|\d+ сек киин|\d+ сек мурн|\d+ аи мурн|\d+ с киин|\d+ с мурн)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секунддан кииин|\d+ мүнөттөн кииин|\d+ аптадан кииин|\d+ сааттан кииин|\d+ жылдан кииин|\d+ күндөн кииин|\d+ секунд мурун|\d+ аидан кииин|\d+ мүнөт мурун|\d+ апта мурун|\d+ саат кииин|\d+ саат мурун|\d+ аид кииин|\d+ апт кииин|\d+ апт мурун|\d+ жыл кииин|\d+ жыл мурун|\d+ күн кииин|\d+ күн мурун|\d+ мүн кииин|\d+ мүн мурун|\d+ сек кииин|\d+ сек мурун|\d+ аи мурун|\d+ аид киин|\d+ мүн киин|\d+ мүн мурн|\d+ сек киин|\d+ сек мурн|\d+ аи мурн|\d+ с киин|\d+ с мурн)$`),
	KnownWords:      []string{"келерки аптада", "түштөн кииинки", "ушул мүнөттө", "өткөн аптада", "келерки апт", "ушул аптада", "ушул саатта", "эмдиги аида", "эмдиги жылы", "өткөн аида", "өткөн апт", "беишемби", "бул аида", "дүишөмбү", "жекшемби", "сентябрь", "ушул апт", "шаршемби", "шеишемби", "декабрь", "октябрь", "февраль", "август", "апрель", "былтыр", "ишемби", "ноябрь", "секунд", "январь", "быиыл", "бүгүн", "кечээ", "мүнөт", "таңкы", "эртеӊ", "азыр", "апта", "беиш", "жума", "июль", "июнь", "март", "саат", "шарш", "шеиш", "gmt", "utc", "авг", "апр", "апт", "дек", "дүи", "жек", "жыл", "ишм", "июл", "июн", "күн", "маи", "мар", "мүн", "ноя", "окт", "сек", "сен", "фев", "янв", "am", "pm", "аи", "ст", "тк", "тң", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ж", "м"},
})
