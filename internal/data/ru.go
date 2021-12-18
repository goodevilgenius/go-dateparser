// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ru_Locale = merge(nil, LocaleData{
	Name:      "ru",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "в", "во", "и", "около", "примерно"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)(день|дня)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 день${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)секунд[уы](\z|[^\pL\pM\d]|_)`), "${1}${2}1 секунду${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)сут(ки|ок)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 сутки${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)минут[уы](\z|[^\pL\pM\d]|_)`), "${1}${2}1 минуту${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)недел[юи](\z|[^\pL\pM\d]|_)`), "${1}${2}1 неделю${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)месяца?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 месяц${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)года?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 год${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)часа?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 час${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)секунду(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 секунду${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)минуту(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 минуту${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)неделю(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 неделю${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)месяц(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 месяц${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)сутки(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 сутки${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)день(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 день${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)год(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 год${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)час(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 час${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)од(на|ну|ни|нои|ин)(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)несколько секунд(\z|[^\pL\pM\d]|_)`), "${1}44 секунды${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)дв(а|е|ое|ух)(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полтора года(\z|[^\pL\pM\d]|_)`), "${1}18 месяцев${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полтора часа(\z|[^\pL\pM\d]|_)`), "${1}90 минут${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)одиннадцать(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)двенадцать(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пятнадцать(\z|[^\pL\pM\d]|_)`), "${1}15${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пятьдесят(\z|[^\pL\pM\d]|_)`), "${1}50${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пар[ауы](\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)двадцать(\z|[^\pL\pM\d]|_)`), "${1}20${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)тридцать(\z|[^\pL\pM\d]|_)`), "${1}30${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полчаса(\z|[^\pL\pM\d]|_)`), "${1}30 минут${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полгода(\z|[^\pL\pM\d]|_)`), "${1}6 месяцев${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)десять(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)четыре(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)восемь(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)девять(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)сорок(\z|[^\pL\pM\d]|_)`), "${1}40${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)шесть(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пять(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)семь(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)три(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
	},
	Translations: map[string]string{
		"на следующеи неделе": "in 1 week",
		"в следующем месяце":  "in 1 month",
		"на прошлои неделе":   "1 week ago",
		"в прошлом месяце":    "1 month ago",
		"в следующем году":    "in 1 year",
		"на этои неделе":      "0 week ago",
		"в прошлом году":      "1 year ago",
		"в этом месяце":       "0 month ago",
		"в эту минуту":        "0 minute ago",
		"в этом часе":         "0 hour ago",
		"в этом году":         "0 year ago",
		"послезавтра":         "in 2 day",
		"понедельник":         "monday",
		"воскресение":         "sunday",
		"воскресенье":         "sunday",
		"позавчера":           "2 day ago",
		"в течение":           "in",
		"примерно":            "",
		"сентябрь":            "september",
		"сентября":            "september",
		"сегодня":             "0 day ago",
		"августа":             "august",
		"декабрь":             "december",
		"декабря":             "december",
		"февраль":             "february",
		"февраля":             "february",
		"пятница":             "friday",
		"пятницу":             "friday",
		"месяцев":             "month",
		"октябрь":             "october",
		"октября":             "october",
		"суббота":             "saturday",
		"субботу":             "saturday",
		"секунда":             "second",
		"секунду":             "second",
		"секунды":             "second",
		"четверг":             "thursday",
		"вторник":             "tuesday",
		"сеичас":              "0 second ago",
		"апрель":              "april",
		"апреля":              "april",
		"август":              "august",
		"спустя":              "in",
		"завтра":              "in 1 day",
		"январь":              "january",
		"января":              "january",
		"минута":              "minute",
		"минуту":              "minute",
		"минуты":              "minute",
		"месяца":              "month",
		"ноябрь":              "november",
		"ноября":              "november",
		"секунд":              "second",
		"недели":              "week",
		"недель":              "week",
		"неделю":              "week",
		"неделя":              "week",
		"около":               "",
		"вчера":               "1 day ago",
		"назад":               "ago",
		"сутки":               "day",
		"суток":               "day",
		"часов":               "hour",
		"через":               "in",
		"марта":               "march",
		"минут":               "minute",
		"месяц":               "month",
		"среда":               "wednesday",
		"среду":               "wednesday",
		"день":                "day",
		"днеи":                "day",
		"февр":                "february",
		"часа":                "hour",
		"июль":                "july",
		"июля":                "july",
		"июнь":                "june",
		"июня":                "june",
		"март":                "march",
		"нояб":                "november",
		"сент":                "september",
		"года":                "year",
		"апр":                 "april",
		"авг":                 "august",
		"дня":                 "day",
		"дек":                 "december",
		"фев":                 "february",
		"gmt":                 "gmt",
		"час":                 "hour",
		"янв":                 "january",
		"июл":                 "july",
		"июн":                 "june",
		"мар":                 "march",
		"маи":                 "may",
		"мая":                 "may",
		"мин":                 "minute",
		"мес":                 "month",
		"ноя":                 "november",
		"окт":                 "october",
		"сек":                 "second",
		"сен":                 "september",
		"utc":                 "utc",
		"нед":                 "week",
		"год":                 "year",
		"лет":                 "year",
		"во":                  "",
		"am":                  "am",
		"дп":                  "am",
		"дн":                  "day",
		"пт":                  "friday",
		"пн":                  "monday",
		"pm":                  "pm",
		"пп":                  "pm",
		"сб":                  "saturday",
		"вс":                  "sunday",
		"чт":                  "thursday",
		"вт":                  "tuesday",
		"ср":                  "wednesday",
		"'":                   "",
		",":                   "",
		";":                   "",
		"@":                   "",
		"[":                   "",
		"]":                   "",
		"|":                   "",
		"в":                   "",
		"и":                   "",
		" ":                   " ",
		"+":                   "+",
		"-":                   "-",
		".":                   ".",
		"/":                   "/",
		":":                   ":",
		"ч":                   "hour",
		"с":                   "second",
		"г":                   "year",
		"z":                   "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секунду назад`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) секунды назад`), "$1 second ago"},
		{regexp.MustCompile(`(?i)через (\d+) секунду`), "in $1 second"},
		{regexp.MustCompile(`(?i)через (\d+) секунды`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) минуту назад`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) минуты назад`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) месяца назад`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) недели назад`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) неделю назад`), "$1 week ago"},
		{regexp.MustCompile(`(?i)через (\d+) минуту`), "in $1 minute"},
		{regexp.MustCompile(`(?i)через (\d+) минуты`), "in $1 minute"},
		{regexp.MustCompile(`(?i)через (\d+) месяца`), "in $1 month"},
		{regexp.MustCompile(`(?i)через (\d+) недели`), "in $1 week"},
		{regexp.MustCompile(`(?i)через (\d+) неделю`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) месяц назад`), "$1 month ago"},
		{regexp.MustCompile(`(?i)через (\d+) месяц`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) день назад`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) часа назад`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) года назад`), "$1 year ago"},
		{regexp.MustCompile(`(?i)через (\d+) день`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) часа`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) года`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) дня назад`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) час назад`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) мин назад`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) мес назад`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) сек назад`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) нед назад`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) год назад`), "$1 year ago"},
		{regexp.MustCompile(`(?i)через (\d+) дня`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) час`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) мин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)через (\d+) мес`), "in $1 month"},
		{regexp.MustCompile(`(?i)через (\d+) сек`), "in $1 second"},
		{regexp.MustCompile(`(?i)через (\d+) нед`), "in $1 week"},
		{regexp.MustCompile(`(?i)через (\d+) год`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) дн назад`), "$1 day ago"},
		{regexp.MustCompile(`(?i)через (\d+) дн`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) д назад`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ч назад`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) г назад`), "$1 year ago"},
		{regexp.MustCompile(`(?i)через (\d+) д`), "in $1 day"},
		{regexp.MustCompile(`(?i)через (\d+) ч`), "in $1 hour"},
		{regexp.MustCompile(`(?i)через (\d+) г`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ секунду назад|\d+ секунды назад|через \d+ секунду|через \d+ секунды|\d+ месяца назад|\d+ минуту назад|\d+ минуты назад|\d+ недели назад|\d+ неделю назад|через \d+ месяца|через \d+ минуту|через \d+ минуты|через \d+ недели|через \d+ неделю|\d+ месяц назад|через \d+ месяц|\d+ года назад|\d+ день назад|\d+ часа назад|через \d+ года|через \d+ день|через \d+ часа|\d+ год назад|\d+ дня назад|\d+ мес назад|\d+ мин назад|\d+ нед назад|\d+ сек назад|\d+ час назад|через \d+ год|через \d+ дня|через \d+ мес|через \d+ мин|через \d+ нед|через \d+ сек|через \d+ час|\d+ дн назад|через \d+ дн|\d+ г назад|\d+ д назад|\d+ ч назад|через \d+ г|через \d+ д|через \d+ ч)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секунду назад|\d+ секунды назад|через \d+ секунду|через \d+ секунды|\d+ месяца назад|\d+ минуту назад|\d+ минуты назад|\d+ недели назад|\d+ неделю назад|через \d+ месяца|через \d+ минуту|через \d+ минуты|через \d+ недели|через \d+ неделю|\d+ месяц назад|через \d+ месяц|\d+ года назад|\d+ день назад|\d+ часа назад|через \d+ года|через \d+ день|через \d+ часа|\d+ год назад|\d+ дня назад|\d+ мес назад|\d+ мин назад|\d+ нед назад|\d+ сек назад|\d+ час назад|через \d+ год|через \d+ дня|через \d+ мес|через \d+ мин|через \d+ нед|через \d+ сек|через \d+ час|\d+ дн назад|через \d+ дн|\d+ г назад|\d+ д назад|\d+ ч назад|через \d+ г|через \d+ д|через \d+ ч)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(на следующеи неделе|в следующем месяце|на прошлои неделе|в прошлом месяце|в следующем году|в прошлом году|на этои неделе|в этом месяце|в эту минуту|в этом году|в этом часе|воскресение|воскресенье|понедельник|послезавтра|в течение|позавчера|примерно|сентябрь|сентября|августа|вторник|декабрь|декабря|месяцев|октябрь|октября|пятница|пятницу|сегодня|секунда|секунду|секунды|суббота|субботу|февраль|февраля|четверг|август|апрель|апреля|завтра|месяца|минута|минуту|минуты|недели|недель|неделю|неделя|ноябрь|ноября|сеичас|секунд|спустя|январь|января|вчера|марта|месяц|минут|назад|около|среда|среду|сутки|суток|часов|через|года|день|днеи|июль|июля|июнь|июня|март|нояб|сент|февр|часа|gmt|utc|авг|апр|год|дек|дня|июл|июн|лет|маи|мар|мая|мес|мин|нед|ноя|окт|сек|сен|фев|час|янв|\+|\.|\[|\]|\||am|pm|во|вс|вт|дн|дп|пн|пп|пт|сб|ср|чт| |'|,|-|/|:|;|@|z|в|г|и|с|ч)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ru_BY_Locale = merge(&ru_Locale, LocaleData{
	Name:      "ru-BY",
	DateOrder: "DMY",
})

var ru_KG_Locale = merge(&ru_Locale, LocaleData{
	Name:      "ru-KG",
	DateOrder: "DMY",
})

var ru_KZ_Locale = merge(&ru_Locale, LocaleData{
	Name:      "ru-KZ",
	DateOrder: "DMY",
})

var ru_MD_Locale = merge(&ru_Locale, LocaleData{
	Name:      "ru-MD",
	DateOrder: "DMY",
})

var ru_UA_Locale = merge(&ru_Locale, LocaleData{
	Name:      "ru-UA",
	DateOrder: "DMY",
	Translations: map[string]string{
		"gmt": "gmt",
		"utc": "utc",
		"am":  "am",
		"pm":  "pm",
		" ":   " ",
		"+":   "+",
		"-":   "-",
		".":   ".",
		"/":   "/",
		":":   ":",
		"z":   "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(gmt|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
