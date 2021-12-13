// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mgh_Locale = LocaleData{
	Name:      "mgh",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)lel'lo(\z|\W|_)`):                  "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):               "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):             "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):              "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):                     "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):               "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):               "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)n'chana(\z|\W|_)`):                 "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):              "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):               "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):               "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)wichishu(\z|\W|_)`):                "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)che(\z|\W|_)`):                     "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo unecheshe(\z|\W|_)`):      "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo nane(\z|\W|_)`):           "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)nan(\z|\W|_)`):                     "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)nihuku(\z|\W|_)`):                  "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo kumi na yel'li(\z|\W|_)`): "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)yel(\z|\W|_)`):                     "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo unayeli(\z|\W|_)`):        "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)una(\z|\W|_)`):                     "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)iju(\z|\W|_)`):                     "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)ijumaa(\z|\W|_)`):                  "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)isaa(\z|\W|_)`):                    "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)me'llo(\z|\W|_)`):                  "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):              "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):               "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):               "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)kwa(\z|\W|_)`):                     "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo kwanza(\z|\W|_)`):         "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo saba(\z|\W|_)`):           "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)sab(\z|\W|_)`):                     "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)moc(\z|\W|_)`):                     "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo thanu na mocha(\z|\W|_)`): "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo uneraru(\z|\W|_)`):        "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)rar(\z|\W|_)`):                     "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo unethanu(\z|\W|_)`):       "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)tha(\z|\W|_)`):                     "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)jtt(\z|\W|_)`):                     "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)jumatatu(\z|\W|_)`):                "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mweri(\z|\W|_)`):                   "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)moj(\z|\W|_)`):                     "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo kumi na moja(\z|\W|_)`):   "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)kum(\z|\W|_)`):                     "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo kumi(\z|\W|_)`):           "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)mchochil'l(\z|\W|_)`):              "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)jmo(\z|\W|_)`):                     "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)jumamosi(\z|\W|_)`):                "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)isekunde(\z|\W|_)`):                "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)mweri wo tisa(\z|\W|_)`):           "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)tis(\z|\W|_)`):                     "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)sabato(\z|\W|_)`):                  "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)ara(\z|\W|_)`):                     "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)arahamisi(\z|\W|_)`):               "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)jnn(\z|\W|_)`):                     "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)jumanne(\z|\W|_)`):                 "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)jtn(\z|\W|_)`):                     "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)jumatano(\z|\W|_)`):                "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)iwiki mocha(\z|\W|_)`):             "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)yaka(\z|\W|_)`):                    "${1}year${2}",
	},
}
