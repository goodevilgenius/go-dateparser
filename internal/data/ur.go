// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ur_Locale = merge(nil, LocaleData{
	Name:      "ur",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzآئابتجدرزسشعفلمنويٹپچڈکگھہیے`),
	Translations: map[string]string{
		"اکتوبر": "october",
		"جمعرات": "thursday",
		"جولايی": "july",
		"سوموار": "monday",
		"اتوار":  "sunday",
		"اپریل":  "april",
		"جنوری":  "january",
		"دسمبر":  "december",
		"ستمبر":  "september",
		"سیکنڈ":  "second",
		"فروری":  "february",
		"مہینہ":  "month",
		"نومبر":  "november",
		"گھنٹہ":  "hour",
		"اگست":   "august",
		"جمعہ":   "friday",
		"مارچ":   "march",
		"منگل":   "tuesday",
		"ہفتہ":   "week",
		"gmt":    "gmt",
		"utc":    "utc",
		"بدھ":    "wednesday",
		"جون":    "june",
		"سال":    "year",
		"ماہ":    "month",
		"منٹ":    "minute",
		"ميی":    "may",
		"am":     "am",
		"pm":     "pm",
		"دن":     "day",
		" ":      " ",
		"'":      "",
		"+":      "+",
		",":      "",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		";":      "",
		"@":      "",
		"[":      "",
		"]":      "",
		"z":      "z",
		"|":      "",
	},
	RelativeType: map[string]string{
		"پچھلے مہینہ": "1 month ago",
		"اگلے مہینہ":  "in 1 month",
		"پچھلے ہفتہ":  "1 week ago",
		"اگلے ہفتہ":   "in 1 week",
		"گزشتہ سال":   "1 year ago",
		"اس مہینہ":    "0 month ago",
		"اس گھنٹے":    "0 hour ago",
		"ايندہ کل":    "in 1 day",
		"اگلے سال":    "in 1 year",
		"گزشتہ کل":    "1 day ago",
		"اس ہفتہ":     "0 week ago",
		"اس سال":      "0 year ago",
		"اس منٹ":      "0 minute ago",
		"اب":          "0 second ago",
		"اج":          "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) سیکنڈ پہلے`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) مہینہ پہلے`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) مہینے پہلے`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹوں میں`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹہ پہلے`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹے پہلے`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) دنوں پہلے`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) سیکنڈ میں`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) مہینہ میں`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) مہینے میں`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹہ میں`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹے میں`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ہفتہ پہلے`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ہفتے پہلے`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) دنوں میں`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) سال پہلے`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ماہ پہلے`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) منٹ پہلے`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ہفتہ میں`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ہفتے میں`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) دن پہلے`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) سال میں`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ماہ قبل`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ماہ میں`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) منٹ میں`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) دن میں`), "in $1 day"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ سیکنڈ پہلے|\d+ مہینہ پہلے|\d+ مہینے پہلے|\d+ گھنٹوں میں|\d+ گھنٹہ پہلے|\d+ گھنٹے پہلے|\d+ دنوں پہلے|\d+ سیکنڈ میں|\d+ مہینہ میں|\d+ مہینے میں|\d+ گھنٹہ میں|\d+ گھنٹے میں|\d+ ہفتہ پہلے|\d+ ہفتے پہلے|\d+ دنوں میں|\d+ سال پہلے|\d+ ماہ پہلے|\d+ منٹ پہلے|\d+ ہفتہ میں|\d+ ہفتے میں|\d+ دن پہلے|\d+ سال میں|\d+ ماہ قبل|\d+ ماہ میں|\d+ منٹ میں|\d+ دن میں)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ سیکنڈ پہلے|\d+ مہینہ پہلے|\d+ مہینے پہلے|\d+ گھنٹوں میں|\d+ گھنٹہ پہلے|\d+ گھنٹے پہلے|\d+ دنوں پہلے|\d+ سیکنڈ میں|\d+ مہینہ میں|\d+ مہینے میں|\d+ گھنٹہ میں|\d+ گھنٹے میں|\d+ ہفتہ پہلے|\d+ ہفتے پہلے|\d+ دنوں میں|\d+ سال پہلے|\d+ ماہ پہلے|\d+ منٹ پہلے|\d+ ہفتہ میں|\d+ ہفتے میں|\d+ دن پہلے|\d+ سال میں|\d+ ماہ قبل|\d+ ماہ میں|\d+ منٹ میں|\d+ دن میں)$`),
	KnownWords:      []string{"پچھلے مہینہ", "اگلے مہینہ", "پچھلے ہفتہ", "اگلے ہفتہ", "گزشتہ سال", "اس مہینہ", "اس گھنٹے", "ايندہ کل", "اگلے سال", "گزشتہ کل", "اس ہفتہ", "اس سال", "اس منٹ", "اکتوبر", "جمعرات", "جولايی", "سوموار", "اتوار", "اپریل", "جنوری", "دسمبر", "ستمبر", "سیکنڈ", "فروری", "مہینہ", "نومبر", "گھنٹہ", "اگست", "جمعہ", "مارچ", "منگل", "ہفتہ", "gmt", "utc", "بدھ", "جون", "سال", "ماہ", "منٹ", "ميی", "am", "pm", "اب", "اج", "دن", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})

var ur_IN_Locale = merge(&ur_Locale, LocaleData{
	Name:      "ur-IN",
	DateOrder: "DMY",
	Translations: map[string]string{
		"پیر": "monday",
	},
	RelativeType: map[string]string{
		"گزشتہ ہفتہ": "1 week ago",
		"گزشتہ ماہ":  "1 month ago",
		"اگلے ماہ":   "in 1 month",
		"اس ماہ":     "0 month ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) سالوں پہلے`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) سالوں میں`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) سیکنڈ قبل`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹہ قبل`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) گھنٹے قبل`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ہفتوں میں`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ہفتہ قبل`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ہفتے قبل`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) منٹ قبل`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) دن قبل`), "$1 day ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ سالوں پہلے|\d+ سیکنڈ پہلے|\d+ مہینہ پہلے|\d+ مہینے پہلے|\d+ گھنٹوں میں|\d+ گھنٹہ پہلے|\d+ گھنٹے پہلے|\d+ دنوں پہلے|\d+ سالوں میں|\d+ سیکنڈ قبل|\d+ سیکنڈ میں|\d+ مہینہ میں|\d+ مہینے میں|\d+ گھنٹہ قبل|\d+ گھنٹہ میں|\d+ گھنٹے قبل|\d+ گھنٹے میں|\d+ ہفتوں میں|\d+ ہفتہ پہلے|\d+ ہفتے پہلے|\d+ دنوں میں|\d+ سال پہلے|\d+ ماہ پہلے|\d+ منٹ پہلے|\d+ ہفتہ قبل|\d+ ہفتہ میں|\d+ ہفتے قبل|\d+ ہفتے میں|\d+ دن پہلے|\d+ سال میں|\d+ ماہ قبل|\d+ ماہ میں|\d+ منٹ قبل|\d+ منٹ میں|\d+ دن قبل|\d+ دن میں)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ سالوں پہلے|\d+ سیکنڈ پہلے|\d+ مہینہ پہلے|\d+ مہینے پہلے|\d+ گھنٹوں میں|\d+ گھنٹہ پہلے|\d+ گھنٹے پہلے|\d+ دنوں پہلے|\d+ سالوں میں|\d+ سیکنڈ قبل|\d+ سیکنڈ میں|\d+ مہینہ میں|\d+ مہینے میں|\d+ گھنٹہ قبل|\d+ گھنٹہ میں|\d+ گھنٹے قبل|\d+ گھنٹے میں|\d+ ہفتوں میں|\d+ ہفتہ پہلے|\d+ ہفتے پہلے|\d+ دنوں میں|\d+ سال پہلے|\d+ ماہ پہلے|\d+ منٹ پہلے|\d+ ہفتہ قبل|\d+ ہفتہ میں|\d+ ہفتے قبل|\d+ ہفتے میں|\d+ دن پہلے|\d+ سال میں|\d+ ماہ قبل|\d+ ماہ میں|\d+ منٹ قبل|\d+ منٹ میں|\d+ دن قبل|\d+ دن میں)$`),
	KnownWords:      []string{"پچھلے مہینہ", "اگلے مہینہ", "پچھلے ہفتہ", "گزشتہ ہفتہ", "اگلے ہفتہ", "گزشتہ سال", "گزشتہ ماہ", "اس مہینہ", "اس گھنٹے", "ايندہ کل", "اگلے سال", "اگلے ماہ", "گزشتہ کل", "اس ہفتہ", "اس سال", "اس ماہ", "اس منٹ", "اکتوبر", "جمعرات", "جولايی", "سوموار", "اتوار", "اپریل", "جنوری", "دسمبر", "ستمبر", "سیکنڈ", "فروری", "مہینہ", "نومبر", "گھنٹہ", "اگست", "جمعہ", "مارچ", "منگل", "ہفتہ", "gmt", "utc", "بدھ", "جون", "سال", "ماہ", "منٹ", "ميی", "پیر", "am", "pm", "اب", "اج", "دن", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
