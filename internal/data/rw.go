// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var rw_Locale = LocaleData{
	Name:      "rw",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)today(\z|\W|_)`):          "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):      "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):    "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):     "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):            "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):      "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):      "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)yesterday(\z|\W|_)`):      "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):     "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):      "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):      "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`):             "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)mat(\z|\W|_)`):            "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)mata(\z|\W|_)`):           "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)kan(\z|\W|_)`):            "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)kanama(\z|\W|_)`):         "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)day(\z|\W|_)`):            "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)uku(\z|\W|_)`):            "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)ukuboza(\z|\W|_)`):        "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)gas(\z|\W|_)`):            "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)gashyantare(\z|\W|_)`):    "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)gnu(\z|\W|_)`):            "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa gatanu(\z|\W|_)`):    "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)hour(\z|\W|_)`):           "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)tomorrow(\z|\W|_)`):       "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):     "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):      "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):      "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)mut(\z|\W|_)`):            "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)mutarama(\z|\W|_)`):       "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)nya(\z|\W|_)`):            "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)nyakanga(\z|\W|_)`):       "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)kam(\z|\W|_)`):            "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)kamena(\z|\W|_)`):         "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)wer(\z|\W|_)`):            "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)werurwe(\z|\W|_)`):        "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)gic(\z|\W|_)`):            "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)gicuransi(\z|\W|_)`):      "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa mbere(\z|\W|_)`):     "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mbe(\z|\W|_)`):            "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)month(\z|\W|_)`):          "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)ugu(\z|\W|_)`):            "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ugushyingo(\z|\W|_)`):     "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ukw(\z|\W|_)`):            "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)ukwakira(\z|\W|_)`):       "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`):             "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)gnd(\z|\W|_)`):            "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa gatandatu(\z|\W|_)`): "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)second(\z|\W|_)`):         "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)nze(\z|\W|_)`):            "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)nzeli(\z|\W|_)`):          "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)cyu(\z|\W|_)`):            "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)ku cyumweru(\z|\W|_)`):    "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa kane(\z|\W|_)`):      "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)kab(\z|\W|_)`):            "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa kabiri(\z|\W|_)`):    "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)gtu(\z|\W|_)`):            "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)kuwa gatatu(\z|\W|_)`):    "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)week(\z|\W|_)`):           "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)year(\z|\W|_)`):           "${1}year${2}",
	},
}
