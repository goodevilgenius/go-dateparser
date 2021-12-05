// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hyLocale = LocaleData{
	Name:                  "hy",
	DateOrder:             "DMY",
	January:               []string{"հնվ", "հունվար", "հունվարի"},
	February:              []string{"փետրվար", "փետրվարի", "փտվ"},
	March:                 []string{"մարտ", "մարտի", "մրտ"},
	April:                 []string{"ապր", "ապրիլ", "ապրիլի"},
	May:                   []string{"մայիս", "մայիսի", "մյս"},
	June:                  []string{"հնս", "հունիս", "հունիսի"},
	July:                  []string{"հլս", "հուլիս", "հուլիսի"},
	August:                []string{"օգոստոս", "օգոստոսի", "օգս"},
	September:             []string{"սեպ", "սեպտեմբեր", "սեպտեմբերի"},
	October:               []string{"հոկ", "հոկտեմբեր", "հոկտեմբերի"},
	November:              []string{"նոյ", "նոյեմբեր", "նոյեմբերի"},
	December:              []string{"դեկ", "դեկտեմբեր", "դեկտեմբերի"},
	Monday:                []string{"երկ", "երկուշաբթի"},
	Tuesday:               []string{"երեքշաբթի", "երք"},
	Wednesday:             []string{"չորեքշաբթի", "չրք"},
	Thursday:              []string{"հինգշաբթի", "հնգ"},
	Friday:                []string{"ուր", "ուրբաթ"},
	Saturday:              []string{"շաբաթ", "շբթ"},
	Sunday:                []string{"կիր", "կիրակի"},
	AM:                    []string{"am", "կա"},
	PM:                    []string{"pm", "կհ"},
	Year:                  []string{"տ", "տարի"},
	Month:                 []string{"ամիս", "ամս"},
	Week:                  []string{"շաբ", "շաբաթ"},
	Day:                   []string{"օր"},
	Hour:                  []string{"ժ", "ժամ"},
	Minute:                []string{"ր", "րոպե"},
	Second:                []string{"վ", "վայրկյան"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`այսօր`},
		`0 hour ago`:   {`այս ժամին`},
		`0 minute ago`: {`այս րոպեին`},
		`0 month ago`:  {`այս ամիս`},
		`0 second ago`: {`այժմ`},
		`0 week ago`:   {`այս շաբաթ`},
		`0 year ago`:   {`այս տարի`},
		`1 day ago`:    {`երեկ`},
		`1 month ago`:  {`անցյալ ամիս`, `նախորդ ամիս`},
		`1 week ago`:   {`նախորդ շաբաթ`},
		`1 year ago`:   {`նախորդ տարի`},
		`in 1 day`:     {`վաղը`},
		`in 1 month`:   {`հաջորդ ամիս`},
		`in 1 week`:    {`հաջորդ շաբաթ`},
		`in 1 year`:    {`հաջորդ տարի`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) օր առաջ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ժ առաջ`),
			regexp.MustCompile(`(?i)(\d+) ժամ առաջ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) ր առաջ`),
			regexp.MustCompile(`(?i)(\d+) րոպե առաջ`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ամիս առաջ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) վ առաջ`),
			regexp.MustCompile(`(?i)(\d+) վայրկյան առաջ`),
			regexp.MustCompile(`(?i)(\d+) վրկ առաջ`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) շաբ առաջ`),
			regexp.MustCompile(`(?i)(\d+) շաբաթ առաջ`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) տ առաջ`),
			regexp.MustCompile(`(?i)(\d+) տարի առաջ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) օրից`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ժ-ից`),
			regexp.MustCompile(`(?i)(\d+) ժամից`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) ր-ից`),
			regexp.MustCompile(`(?i)(\d+) րոպեից`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ամսից`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) վ-ից`),
			regexp.MustCompile(`(?i)(\d+) վայրկյանից`),
			regexp.MustCompile(`(?i)(\d+) վրկ-ից`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) շաբ անց`),
			regexp.MustCompile(`(?i)(\d+) շաբ-ից`),
			regexp.MustCompile(`(?i)(\d+) շաբաթից`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) տարուց`),
		},
	},
}