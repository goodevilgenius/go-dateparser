// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var jgoLocale = LocaleData{
	Name:                  "jgo",
	DateOrder:             "YMD",
	January:               []string{"nduŋmbi saŋ"},
	February:              []string{"pɛsaŋ pɛ́pá"},
	March:                 []string{"pɛsaŋ pɛ́tát"},
	April:                 []string{"pɛsaŋ pɛ́nɛ́kwa"},
	May:                   []string{"pɛsaŋ pataa"},
	June:                  []string{"pɛsaŋ pɛ́nɛ́ntúkú"},
	July:                  []string{"pɛsaŋ saambá"},
	August:                []string{"pɛsaŋ pɛ́nɛ́fɔm"},
	September:             []string{"pɛsaŋ pɛ́nɛ́pfú'ú"},
	October:               []string{"pɛsaŋ nɛgɛ́m"},
	November:              []string{"pɛsaŋ ntsɔ̌pmɔ́"},
	December:              []string{"pɛsaŋ ntsɔ̌ppá"},
	Monday:                []string{"mɔ́ndi"},
	Tuesday:               []string{"ápta mɔ́ndi"},
	Wednesday:             []string{"wɛ́nɛsɛdɛ"},
	Thursday:              []string{"tɔ́sɛdɛ"},
	Friday:                []string{"fɛlâyɛdɛ"},
	Saturday:              []string{"sásidɛ"},
	Sunday:                []string{"sɔ́ndi"},
	AM:                    []string{"mba'mba'"},
	PM:                    []string{"ŋka mbɔ́t nji"},
	Year:                  []string{"year"},
	Month:                 []string{"month"},
	Week:                  []string{"week"},
	Day:                   []string{"day"},
	Hour:                  []string{"hour"},
	Minute:                []string{"minute"},
	Second:                []string{"second"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`lɔꞌɔ`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`yesterday`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`tomorrow`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)ɛ́ gɛ́ mɔ́ lɛ́ꞌ (\d+)`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)ɛ́ gɛ mɔ́ (\d+) háwa`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)ɛ́ gɛ́ mɔ́ minút (\d+)`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)ɛ́ gɛ́ mɔ́ pɛsaŋ (\d+)`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)ɛ́ gɛ́ mɔ (\d+) ŋgap-mbi`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)ɛ́gɛ́ mɔ́ ŋguꞌ (\d+)`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)nǔu lɛ́ꞌ (\d+)`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)nǔu háwa (\d+)`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)nǔu (\d+) minút`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)nǔu (\d+) saŋ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)nǔu ŋgap-mbi (\d+)`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)nǔu ŋguꞌ (\d+)`),
		},
	},
}