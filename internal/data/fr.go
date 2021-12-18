// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fr_Locale = merge(nil, LocaleData{
	Name:      "fr",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "a", "environ", "er", "et", "le", "|"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+)\s?(minute|seconde|heure)(\z|[^\pL\pM\d]|_)`), "${1}${2} ${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+)\s?h(\z|[^\pL\pM\d]|_)`), "${1}${2} heure${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+)\s?m(\z|[^\pL\pM\d]|_)`), "${1}${2} minute${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+)\s?s(\z|[^\pL\pM\d]|_)`), "${1}${2} seconde${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)\s+h\s+(\d+)\s+min(\z|[^\pL\pM\d]|_)`), "${1}${2}h${3}m${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)h(\d+)m?(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)quatre(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)d'une(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)douze(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)trois(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)d'un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)onze(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)deux(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)cinq(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)sept(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)huit(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)neuf(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)une(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)dix(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)six(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string]string{
		"la semaine prochaine": "in 1 week",
		"la semaine derniere":  "1 week ago",
		"l'annee prochaine":    "in 1 year",
		"l'annee derniere":     "1 year ago",
		"le mois prochain":     "in 1 month",
		"cette minute-ci":      "0 minute ago",
		"le mois dernier":      "1 month ago",
		"cette heure-ci":       "0 hour ago",
		"cette semaine":        "0 week ago",
		"apres-demain":         "in 2 day",
		"aujourd'hui":          "0 day ago",
		"cette annee":          "0 year ago",
		"ce mois-ci":           "0 month ago",
		"maintenant":           "0 second ago",
		"avant-hier":           "2 day ago",
		"septembre":            "september",
		"decembre":             "december",
		"vendredi":             "friday",
		"novembre":             "november",
		"secondes":             "second",
		"dimanche":             "sunday",
		"mercredi":             "wednesday",
		"semaines":             "week",
		"environ":              "",
		"fevrier":              "february",
		"janvier":              "january",
		"juillet":              "july",
		"minutes":              "minute",
		"octobre":              "october",
		"seconde":              "second",
		"semaine":              "week",
		"il y a":               "ago",
		"heures":               "hour",
		"demain":               "in 1 day",
		"minute":               "minute",
		"samedi":               "saturday",
		"annees":               "year",
		"il ya":                "ago",
		"avril":                "april",
		"jours":                "day",
		"heure":                "hour",
		"apres":                "in",
		"lundi":                "monday",
		"jeudi":                "thursday",
		"mardi":                "tuesday",
		"annee":                "year",
		"hier":                 "1 day ago",
		"aout":                 "august",
		"jour":                 "day",
		"fevr":                 "february",
		"dans":                 "in",
		"janv":                 "january",
		"juil":                 "july",
		"juin":                 "june",
		"mars":                 "march",
		"mois":                 "month",
		"sept":                 "september",
		"avr":                  "april",
		"aou":                  "august",
		"dec":                  "december",
		"fev":                  "february",
		"ven":                  "friday",
		"gmt":                  "gmt",
		"jan":                  "january",
		"jul":                  "july",
		"jun":                  "june",
		"mai":                  "may",
		"min":                  "minute",
		"lun":                  "monday",
		"nov":                  "november",
		"oct":                  "october",
		"sam":                  "saturday",
		"sep":                  "september",
		"dim":                  "sunday",
		"jeu":                  "thursday",
		"mar":                  "tuesday",
		"utc":                  "utc",
		"mer":                  "wednesday",
		"sem":                  "week",
		"ans":                  "year",
		"er":                   "",
		"et":                   "",
		"le":                   "",
		"am":                   "am",
		"ve":                   "friday",
		"en":                   "in",
		"lu":                   "monday",
		"pm":                   "pm",
		"sa":                   "saturday",
		"di":                   "sunday",
		"je":                   "thursday",
		"ma":                   "tuesday",
		"me":                   "wednesday",
		"an":                   "year",
		"'":                    "",
		",":                    "",
		";":                    "",
		"@":                    "",
		"[":                    "",
		"]":                    "",
		"a":                    "",
		"|":                    "",
		" ":                    " ",
		"+":                    "+",
		"-":                    "-",
		".":                    ".",
		"/":                    "/",
		":":                    ":",
		"j":                    "day",
		"h":                    "hour",
		"m":                    "month",
		"s":                    "second",
		"z":                    "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)il y a (\d+) secondes`), "$1 second ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) semaines`), "$1 week ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) minutes`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) seconde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) semaine`), "$1 week ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) heures`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) minute`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)dans (\d+) secondes`), "in $1 second"},
		{regexp.MustCompile(`(?i)dans (\d+) semaines`), "in $1 week"},
		{regexp.MustCompile(`(?i)il y a (\d+) jours`), "$1 day ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) heure`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)dans (\d+) minutes`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dans (\d+) seconde`), "in $1 second"},
		{regexp.MustCompile(`(?i)dans (\d+) semaine`), "in $1 week"},
		{regexp.MustCompile(`(?i)il y a (\d+) jour`), "$1 day ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) mois`), "$1 month ago"},
		{regexp.MustCompile(`(?i)dans (\d+) heures`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dans (\d+) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)il y a (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) sem`), "$1 week ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) ans`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dans (\d+) jours`), "in $1 day"},
		{regexp.MustCompile(`(?i)dans (\d+) heure`), "in $1 hour"},
		{regexp.MustCompile(`(?i)il y a (\d+)min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) an`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dans (\d+) jour`), "in $1 day"},
		{regexp.MustCompile(`(?i)dans (\d+) mois`), "in $1 month"},
		{regexp.MustCompile(`(?i)il y a (\d+) j`), "$1 day ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) m`), "$1 month ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)il y a (\d+) a`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dans (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dans (\d+) sem`), "in $1 week"},
		{regexp.MustCompile(`(?i)dans (\d+) ans`), "in $1 year"},
		{regexp.MustCompile(`(?i)il y a (\d+)h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)dans (\d+) an`), "in $1 year"},
		{regexp.MustCompile(`(?i)dans (\d+) j`), "in $1 day"},
		{regexp.MustCompile(`(?i)dans (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dans (\d+) m`), "in $1 month"},
		{regexp.MustCompile(`(?i)dans (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)dans (\d+) a`), "in $1 year"},
		{regexp.MustCompile(`(?i)dans (\d+)h`), "in $1 hour"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(il y a \d+ secondes|il y a \d+ semaines|il y a \d+ minutes|il y a \d+ seconde|il y a \d+ semaine|dans \d+ secondes|dans \d+ semaines|il y a \d+ heures|il y a \d+ minute|dans \d+ minutes|dans \d+ seconde|dans \d+ semaine|il y a \d+ heure|il y a \d+ jours|dans \d+ heures|dans \d+ minute|il y a \d+ jour|il y a \d+ mois|dans \d+ heure|dans \d+ jours|il y a \d+ ans|il y a \d+ min|il y a \d+ sem|dans \d+ jour|dans \d+ mois|il y a \d+ an|il y a \d+min|dans \d+ ans|dans \d+ min|dans \d+ sem|il y a \d+ a|il y a \d+ h|il y a \d+ j|il y a \d+ m|il y a \d+ s|dans \d+ an|il y a \d+h|dans \d+ a|dans \d+ h|dans \d+ j|dans \d+ m|dans \d+ s|dans \d+h)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(il y a \d+ secondes|il y a \d+ semaines|il y a \d+ minutes|il y a \d+ seconde|il y a \d+ semaine|dans \d+ secondes|dans \d+ semaines|il y a \d+ heures|il y a \d+ minute|dans \d+ minutes|dans \d+ seconde|dans \d+ semaine|il y a \d+ heure|il y a \d+ jours|dans \d+ heures|dans \d+ minute|il y a \d+ jour|il y a \d+ mois|dans \d+ heure|dans \d+ jours|il y a \d+ ans|il y a \d+ min|il y a \d+ sem|dans \d+ jour|dans \d+ mois|il y a \d+ an|il y a \d+min|dans \d+ ans|dans \d+ min|dans \d+ sem|il y a \d+ a|il y a \d+ h|il y a \d+ j|il y a \d+ m|il y a \d+ s|dans \d+ an|il y a \d+h|dans \d+ a|dans \d+ h|dans \d+ j|dans \d+ m|dans \d+ s|dans \d+h)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(la semaine prochaine|la semaine derniere|l'annee prochaine|l'annee derniere|le mois prochain|cette minute-ci|le mois dernier|cette heure-ci|cette semaine|apres-demain|aujourd'hui|cette annee|avant-hier|ce mois-ci|maintenant|septembre|decembre|dimanche|mercredi|novembre|secondes|semaines|vendredi|environ|fevrier|janvier|juillet|minutes|octobre|seconde|semaine|annees|demain|heures|il y a|minute|samedi|annee|apres|avril|heure|il ya|jeudi|jours|lundi|mardi|aout|dans|fevr|hier|janv|jour|juil|juin|mars|mois|sept|ans|aou|avr|dec|dim|fev|gmt|jan|jeu|jul|jun|lun|mai|mar|mer|min|nov|oct|sam|sem|sep|utc|ven|\+|\.|\[|\]|\||am|an|di|en|er|et|je|le|lu|ma|me|pm|sa|ve| |'|,|-|/|:|;|@|a|h|j|m|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fr_BE_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-BE",
	DateOrder: "DMY",
})

var fr_BF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-BF",
	DateOrder: "DMY",
})

var fr_BI_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-BI",
	DateOrder: "DMY",
})

var fr_BJ_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-BJ",
	DateOrder: "DMY",
})

var fr_BL_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-BL",
	DateOrder: "DMY",
})

var fr_CA_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CA",
	DateOrder: "YMD",
	Translations: map[string]string{
		"juill": "july",
		"gmt":   "gmt",
		"utc":   "utc",
		"am":    "am",
		"pm":    "pm",
		" ":     " ",
		"+":     "+",
		"-":     "-",
		".":     ".",
		"/":     "/",
		":":     ":",
		"z":     "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(juill|gmt|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fr_CD_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CD",
	DateOrder: "DMY",
})

var fr_CF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CF",
	DateOrder: "DMY",
})

var fr_CG_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CG",
	DateOrder: "DMY",
})

var fr_CH_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CH",
	DateOrder: "DMY",
})

var fr_CI_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CI",
	DateOrder: "DMY",
})

var fr_CM_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-CM",
	DateOrder: "DMY",
	Translations: map[string]string{
		"matin": "am",
		"soir":  "pm",
		"mat":   "am",
		"gmt":   "gmt",
		"utc":   "utc",
		"am":    "am",
		"pm":    "pm",
		" ":     " ",
		"+":     "+",
		"-":     "-",
		".":     ".",
		"/":     "/",
		":":     ":",
		"z":     "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(matin|soir|gmt|mat|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fr_DJ_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-DJ",
	DateOrder: "DMY",
})

var fr_DZ_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-DZ",
	DateOrder: "DMY",
})

var fr_GA_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-GA",
	DateOrder: "DMY",
})

var fr_GF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-GF",
	DateOrder: "DMY",
})

var fr_GN_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-GN",
	DateOrder: "DMY",
})

var fr_GP_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-GP",
	DateOrder: "DMY",
})

var fr_GQ_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-GQ",
	DateOrder: "DMY",
})

var fr_HT_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-HT",
	DateOrder: "DMY",
	Translations: map[string]string{
		"gmt": "gmt",
		"sec": "second",
		"utc": "utc",
		"am":  "am",
		"jr":  "day",
		"hr":  "hour",
		"pm":  "pm",
		" ":   " ",
		"+":   "+",
		"-":   "-",
		".":   ".",
		"/":   "/",
		":":   ":",
		"z":   "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(gmt|sec|utc|\+|\.|am|hr|jr|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fr_KM_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-KM",
	DateOrder: "DMY",
})

var fr_LU_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-LU",
	DateOrder: "DMY",
})

var fr_MA_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MA",
	DateOrder: "DMY",
	Translations: map[string]string{
		"gmt": "gmt",
		"jui": "june",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(gmt|jui|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fr_MC_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MC",
	DateOrder: "DMY",
})

var fr_MF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MF",
	DateOrder: "DMY",
})

var fr_MG_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MG",
	DateOrder: "DMY",
})

var fr_ML_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-ML",
	DateOrder: "DMY",
})

var fr_MQ_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MQ",
	DateOrder: "DMY",
})

var fr_MR_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MR",
	DateOrder: "DMY",
})

var fr_MU_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-MU",
	DateOrder: "DMY",
})

var fr_NC_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-NC",
	DateOrder: "DMY",
})

var fr_NE_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-NE",
	DateOrder: "DMY",
})

var fr_PF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-PF",
	DateOrder: "DMY",
})

var fr_PM_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-PM",
	DateOrder: "DMY",
})

var fr_RE_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-RE",
	DateOrder: "DMY",
})

var fr_RW_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-RW",
	DateOrder: "DMY",
})

var fr_SC_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-SC",
	DateOrder: "DMY",
})

var fr_SN_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-SN",
	DateOrder: "DMY",
})

var fr_SY_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-SY",
	DateOrder: "DMY",
})

var fr_TD_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-TD",
	DateOrder: "DMY",
})

var fr_TG_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-TG",
	DateOrder: "DMY",
})

var fr_TN_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-TN",
	DateOrder: "DMY",
})

var fr_VU_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-VU",
	DateOrder: "DMY",
})

var fr_WF_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-WF",
	DateOrder: "DMY",
})

var fr_YT_Locale = merge(&fr_Locale, LocaleData{
	Name:      "fr-YT",
	DateOrder: "DMY",
})
