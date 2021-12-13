// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var rm_Locale = LocaleData{
	Name:      "rm",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)oz(\z|\W|_)`):          "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):   "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`): "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):  "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):         "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):   "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):   "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)ier(\z|\W|_)`):         "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):  "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):   "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):   "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`):          "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)avr(\z|\W|_)`):         "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)avrigl(\z|\W|_)`):      "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)avust(\z|\W|_)`):       "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)tag(\z|\W|_)`):         "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`):         "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)december(\z|\W|_)`):    "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)favr(\z|\W|_)`):        "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)favrer(\z|\W|_)`):      "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)ve(\z|\W|_)`):          "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)venderdi(\z|\W|_)`):    "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)ura(\z|\W|_)`):         "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)damaun(\z|\W|_)`):      "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):  "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):   "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):   "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)schan(\z|\W|_)`):       "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)schaner(\z|\W|_)`):     "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)fan(\z|\W|_)`):         "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)fanadur(\z|\W|_)`):     "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)zercl(\z|\W|_)`):       "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)zercladur(\z|\W|_)`):   "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)mars(\z|\W|_)`):        "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)matg(\z|\W|_)`):        "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)gli(\z|\W|_)`):         "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)glindesdi(\z|\W|_)`):   "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mais(\z|\W|_)`):        "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`):         "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)november(\z|\W|_)`):    "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)oct(\z|\W|_)`):         "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)october(\z|\W|_)`):     "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`):          "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)so(\z|\W|_)`):          "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)sonda(\z|\W|_)`):       "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)secunda(\z|\W|_)`):     "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sett(\z|\W|_)`):        "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)settember(\z|\W|_)`):   "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)du(\z|\W|_)`):          "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)dumengia(\z|\W|_)`):    "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)gie(\z|\W|_)`):         "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)gievgia(\z|\W|_)`):     "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)ma(\z|\W|_)`):          "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)mardi(\z|\W|_)`):       "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)me(\z|\W|_)`):          "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)mesemna(\z|\W|_)`):     "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)emna(\z|\W|_)`):        "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)onn(\z|\W|_)`):         "${1}year${2}",
	},
}
