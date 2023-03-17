// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var zh_Locale = merge(nil, LocaleData{
	Name:                  "zh",
	DateOrder:             "YMD",
	NoWordSpacing:         true,
	Charset:               []rune(`/cgtuz一七三上下个九二五今個八六分刚刻前十午去后周四在天小年拜日时明星昨月期本此现礼秒约这钟间`),
	SentenceSplitterGroup: 4,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)半小时前`), "30分前"},
		{regexp.MustCompile(`(?i)(?:中午|下午|(?:晚上?))(?:\s*)(\d+[.,]?\d*)(?:\s*):(?:\s+|:)?(\d+[.,]?\d*)`), "${1}:${2} pm"},
		{regexp.MustCompile(`(?i)(?:上午|早上|凌晨)(?:\s*)(\d+[.,]?\d*)(?:\s*):(?:\s+|:)?(\d+[.,]?\d*)`), "${1}:${2} am"},
		{regexp.MustCompile(`(?i)中午`), "12:00"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年(?:\s+)?(\d+[.,]?\d*)月(?:\s+)?(\d+[.,]?\d*)日(?:\s+)?(\d+[.,]?\d*)时(?:\s+)?(\d+[.,]?\d*)分`), "${1}-${2}-${3} ${4}:${5}"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年(?:\s+)?(\d+[.,]?\d*)月(?:\s+)?(\d{1,2})(?:日)?(?:\s+)?(\d{1,2})(?:点|:)(\d{1,2})`), "${1}-${2}-${3} ${4}:${5}"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年(?:\s+)?(\d+[.,]?\d*)月(?:\s+)?(\d{1,2})(?:日)?`), "${1}-${2}-${3}"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)月(.*[前后])`), "${1} 月${2}"},
	},
	Translations: map[string][]string{
		"10月": {"october"},
		"11月": {"november"},
		"12月": {"december"},
		"gmt": {"gmt"},
		"utc": {"utc"},
		"十一月": {"november"},
		"十二月": {"december"},
		"星期一": {"monday"},
		"星期三": {"wednesday"},
		"星期二": {"tuesday"},
		"星期五": {"friday"},
		"星期六": {"saturday"},
		"星期四": {"thursday"},
		"星期天": {"sunday"},
		"星期日": {"sunday"},
		"礼拜一": {"monday"},
		"礼拜三": {"wednesday"},
		"礼拜二": {"tuesday"},
		"礼拜五": {"friday"},
		"礼拜六": {"saturday"},
		"礼拜四": {"thursday"},
		"礼拜天": {"sunday"},
		"礼拜日": {"sunday"},
		"1月":  {"january"},
		"2月":  {"february"},
		"3月":  {"march"},
		"4月":  {"april"},
		"5月":  {"may"},
		"6月":  {"june"},
		"7月":  {"july"},
		"8月":  {"august"},
		"9月":  {"september"},
		"am":  {"am"},
		"pm":  {"pm"},
		"一月":  {"january"},
		"七月":  {"july"},
		"三月":  {"march"},
		"上午":  {"am"},
		"下午":  {"pm"},
		"个月":  {"month"},
		"九月":  {"september"},
		"二月":  {"february"},
		"五月":  {"may"},
		"個月":  {"month"},
		"八月":  {"august"},
		"六月":  {"june"},
		"分钟":  {"minute"},
		"十月":  {"october"},
		"周一":  {"monday"},
		"周三":  {"wednesday"},
		"周二":  {"tuesday"},
		"周五":  {"friday"},
		"周六":  {"saturday"},
		"周四":  {"thursday"},
		"周日":  {"sunday"},
		"四月":  {"april"},
		"小时":  {"hour"},
		"星期":  {"week"},
		" ":   {" "},
		"'":   {""},
		"+":   {"+"},
		",":   {""},
		"-":   {"-"},
		".":   {"."},
		"/":   {"/"},
		":":   {":"},
		";":   {""},
		"@":   {""},
		"[":   {""},
		"]":   {""},
		"z":   {"z"},
		"|":   {""},
		"分":   {"minute"},
		"前":   {"ago"},
		"周":   {"week"},
		"在":   {"in"},
		"天":   {"day"},
		"年":   {"year"},
		"日":   {"day"},
		"月":   {"month"},
		"秒":   {"second"},
		"约":   {""},
	},
	RelativeType: map[string]string{
		"这一时间 / 此时": "0 hour ago",
		"这一时间":      "0 hour ago",
		"上个月":       "1 month ago",
		"下个月":       "in 1 month",
		"上周":        "1 week ago",
		"下周":        "in 1 week",
		"今天":        "0 day ago",
		"今年":        "0 year ago",
		"刚刚":        "0 second ago",
		"前天":        "2 day ago",
		"去年":        "1 year ago",
		"后天":        "in 2 days",
		"明天":        "in 1 day",
		"明年":        "in 1 year",
		"昨天":        "1 day ago",
		"本周":        "0 week ago",
		"本月":        "0 month ago",
		"此刻":        "0 minute ago",
		"此时":        "0 hour ago",
		"现在":        "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)个月前`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)个月后`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分钟前`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分钟后`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小时前`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小时后`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒钟前`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒钟后`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)周前`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)周后`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)天前`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)天后`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年前`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年后`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒前`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒后`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\d+[.,]?\d*个月前|\d+[.,]?\d*个月后|\d+[.,]?\d*分钟前|\d+[.,]?\d*分钟后|\d+[.,]?\d*小时前|\d+[.,]?\d*小时后|\d+[.,]?\d*秒钟前|\d+[.,]?\d*秒钟后|\d+[.,]?\d*周前|\d+[.,]?\d*周后|\d+[.,]?\d*天前|\d+[.,]?\d*天后|\d+[.,]?\d*年前|\d+[.,]?\d*年后|\d+[.,]?\d*秒前|\d+[.,]?\d*秒后)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d*个月前|\d+[.,]?\d*个月后|\d+[.,]?\d*分钟前|\d+[.,]?\d*分钟后|\d+[.,]?\d*小时前|\d+[.,]?\d*小时后|\d+[.,]?\d*秒钟前|\d+[.,]?\d*秒钟后|\d+[.,]?\d*周前|\d+[.,]?\d*周后|\d+[.,]?\d*天前|\d+[.,]?\d*天后|\d+[.,]?\d*年前|\d+[.,]?\d*年后|\d+[.,]?\d*秒前|\d+[.,]?\d*秒后)$`),
	KnownWords:      []string{"这一时间 / 此时", "这一时间", "10月", "11月", "12月", "gmt", "utc", "上个月", "下个月", "十一月", "十二月", "星期一", "星期三", "星期二", "星期五", "星期六", "星期四", "星期天", "星期日", "礼拜一", "礼拜三", "礼拜二", "礼拜五", "礼拜六", "礼拜四", "礼拜天", "礼拜日", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "am", "pm", "一月", "七月", "三月", "上午", "上周", "下午", "下周", "个月", "九月", "二月", "五月", "今天", "今年", "個月", "八月", "六月", "分钟", "刚刚", "前天", "十月", "去年", "后天", "周一", "周三", "周二", "周五", "周六", "周四", "周日", "四月", "小时", "明天", "明年", "星期", "昨天", "本周", "本月", "此刻", "此时", "现在", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "分", "前", "周", "在", "天", "年", "日", "月", "秒", "约"},
})
