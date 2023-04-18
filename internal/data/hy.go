// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	hy_Locale LocaleData
)

func init() {
	hy_Locale = merge(nil, LocaleData{
		Name:      "hy",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzաբգդեըթժիլխկհղմյնշոչպջսվտրցւփքօ`),
		Translations: map[string][]string{
			"դեկտեմբերի": {"december"},
			"երկուշաբթի": {"monday"},
			"հոկտեմբերի": {"october"},
			"չորեքշաբթի": {"wednesday"},
			"սեպտեմբերի": {"september"},
			"դեկտեմբեր":  {"december"},
			"երեքշաբթի":  {"tuesday"},
			"հինգշաբթի":  {"thursday"},
			"հոկտեմբեր":  {"october"},
			"նոյեմբերի":  {"november"},
			"սեպտեմբեր":  {"september"},
			"հունվարի":   {"january"},
			"նոյեմբեր":   {"november"},
			"վայրկյան":   {"second"},
			"փետրվարի":   {"february"},
			"օգոստոսի":   {"august"},
			"հուլիսի":    {"july"},
			"հունիսի":    {"june"},
			"հունվար":    {"january"},
			"փետրվար":    {"february"},
			"օգոստոս":    {"august"},
			"ապրիլի":     {"april"},
			"կիրակի":     {"sunday"},
			"հուլիս":     {"july"},
			"հունիս":     {"june"},
			"մայիսի":     {"may"},
			"ուրբաթ":     {"friday"},
			"ապրիլ":      {"april"},
			"մայիս":      {"may"},
			"մարտի":      {"march"},
			"շաբաթ":      {"week", "saturday"},
			"ամիս":       {"month"},
			"մարտ":       {"march"},
			"տարի":       {"year"},
			"րոպե":       {"minute"},
			"gmt":        {"gmt"},
			"utc":        {"utc"},
			"ամս":        {"month"},
			"ապր":        {"april"},
			"դեկ":        {"december"},
			"երկ":        {"monday"},
			"երք":        {"tuesday"},
			"ժամ":        {"hour"},
			"կիր":        {"sunday"},
			"հլս":        {"july"},
			"հնգ":        {"thursday"},
			"հնս":        {"june"},
			"հնվ":        {"january"},
			"հոկ":        {"october"},
			"մյս":        {"may"},
			"մրտ":        {"march"},
			"նոյ":        {"november"},
			"շաբ":        {"week"},
			"շբթ":        {"saturday"},
			"ուր":        {"friday"},
			"չրք":        {"wednesday"},
			"սեպ":        {"september"},
			"փտվ":        {"february"},
			"օգս":        {"august"},
			"am":         {"am"},
			"pm":         {"pm"},
			"կա":         {"am"},
			"կհ":         {"pm"},
			"օր":         {"day"},
			" ":          {" "},
			"'":          {""},
			"+":          {"+"},
			",":          {""},
			"-":          {"-"},
			".":          {"."},
			"/":          {"/"},
			":":          {":"},
			";":          {""},
			"@":          {""},
			"[":          {""},
			"]":          {""},
			"z":          {"z"},
			"|":          {""},
			"ժ":          {"hour"},
			"վ":          {"second"},
			"տ":          {"year"},
			"ր":          {"minute"},
		},
		RelativeType: map[string]string{
			"հաջորդ շաբաթ": "in 1 week",
			"նախորդ շաբաթ": "1 week ago",
			"անցյալ ամիս":  "1 month ago",
			"հաջորդ ամիս":  "in 1 month",
			"հաջորդ տարի":  "in 1 year",
			"նախորդ ամիս":  "1 month ago",
			"նախորդ տարի":  "1 year ago",
			"այս րոպեին":   "0 minute ago",
			"այս ժամին":    "0 hour ago",
			"այս շաբաթ":    "0 week ago",
			"այս ամիս":     "0 month ago",
			"այս տարի":     "0 year ago",
			"այսօր":        "0 day ago",
			"այժմ":         "0 second ago",
			"երեկ":         "1 day ago",
			"վաղը":         "in 1 day",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վայրկյան առաջ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) շաբաթ առաջ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վայրկյանից`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ամիս առաջ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) տարի առաջ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) րոպե առաջ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ժամ առաջ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) շաբ առաջ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վրկ առաջ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) շաբ անց`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) շաբաթից`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) օր առաջ`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ժ առաջ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) շաբ-ից`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վ առաջ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վրկ-ից`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) տ առաջ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) տարուց`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ր առաջ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) րոպեից`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ամսից`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ժամից`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ժ-ից`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) վ-ից`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ր-ից`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) օրից`), "in $1 day"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* վայրկյան առաջ|\d+[.,]?\d* շաբաթ առաջ|\d+[.,]?\d* վայրկյանից|\d+[.,]?\d* ամիս առաջ|\d+[.,]?\d* տարի առաջ|\d+[.,]?\d* րոպե առաջ|\d+[.,]?\d* ժամ առաջ|\d+[.,]?\d* շաբ առաջ|\d+[.,]?\d* վրկ առաջ|\d+[.,]?\d* շաբ անց|\d+[.,]?\d* շաբաթից|\d+[.,]?\d* օր առաջ|\d+[.,]?\d* ժ առաջ|\d+[.,]?\d* շաբ-ից|\d+[.,]?\d* վ առաջ|\d+[.,]?\d* վրկ-ից|\d+[.,]?\d* տ առաջ|\d+[.,]?\d* տարուց|\d+[.,]?\d* ր առաջ|\d+[.,]?\d* րոպեից|\d+[.,]?\d* ամսից|\d+[.,]?\d* ժամից|\d+[.,]?\d* ժ-ից|\d+[.,]?\d* վ-ից|\d+[.,]?\d* ր-ից|\d+[.,]?\d* օրից)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* վայրկյան առաջ|\d+[.,]?\d* շաբաթ առաջ|\d+[.,]?\d* վայրկյանից|\d+[.,]?\d* ամիս առաջ|\d+[.,]?\d* տարի առաջ|\d+[.,]?\d* րոպե առաջ|\d+[.,]?\d* ժամ առաջ|\d+[.,]?\d* շաբ առաջ|\d+[.,]?\d* վրկ առաջ|\d+[.,]?\d* շաբ անց|\d+[.,]?\d* շաբաթից|\d+[.,]?\d* օր առաջ|\d+[.,]?\d* ժ առաջ|\d+[.,]?\d* շաբ-ից|\d+[.,]?\d* վ առաջ|\d+[.,]?\d* վրկ-ից|\d+[.,]?\d* տ առաջ|\d+[.,]?\d* տարուց|\d+[.,]?\d* ր առաջ|\d+[.,]?\d* րոպեից|\d+[.,]?\d* ամսից|\d+[.,]?\d* ժամից|\d+[.,]?\d* ժ-ից|\d+[.,]?\d* վ-ից|\d+[.,]?\d* ր-ից|\d+[.,]?\d* օրից)$`),
		KnownWords:      []string{"հաջորդ շաբաթ", "նախորդ շաբաթ", "անցյալ ամիս", "հաջորդ ամիս", "հաջորդ տարի", "նախորդ ամիս", "նախորդ տարի", "այս րոպեին", "դեկտեմբերի", "երկուշաբթի", "հոկտեմբերի", "չորեքշաբթի", "սեպտեմբերի", "այս ժամին", "այս շաբաթ", "դեկտեմբեր", "երեքշաբթի", "հինգշաբթի", "հոկտեմբեր", "նոյեմբերի", "սեպտեմբեր", "այս ամիս", "այս տարի", "հունվարի", "նոյեմբեր", "վայրկյան", "փետրվարի", "օգոստոսի", "հուլիսի", "հունիսի", "հունվար", "փետրվար", "օգոստոս", "ապրիլի", "կիրակի", "հուլիս", "հունիս", "մայիսի", "ուրբաթ", "այսօր", "ապրիլ", "մայիս", "մարտի", "շաբաթ", "ամիս", "այժմ", "երեկ", "մարտ", "վաղը", "տարի", "րոպե", "gmt", "utc", "ամս", "ապր", "դեկ", "երկ", "երք", "ժամ", "կիր", "հլս", "հնգ", "հնս", "հնվ", "հոկ", "մյս", "մրտ", "նոյ", "շաբ", "շբթ", "ուր", "չրք", "սեպ", "փտվ", "օգս", "am", "pm", "կա", "կհ", "օր", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ժ", "վ", "տ", "ր"},
	})
}
