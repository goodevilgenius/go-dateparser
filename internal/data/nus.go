// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var nus_Locale = LocaleData{
	Name:      "nus",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)walɛ(\z|\W|_)`):            "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):       "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):     "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):      "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):             "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):       "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):       "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)pan(\z|\W|_)`):             "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):      "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):       "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):       "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)rw(\z|\W|_)`):              "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)guak(\z|\W|_)`):            "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)thoo(\z|\W|_)`):            "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)tho̱o̱r(\z|\W|_)`):         "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)tid(\z|\W|_)`):             "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)tio̱p in di̱i̱t(\z|\W|_)`): "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)pɛt(\z|\W|_)`):             "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)dhieec lätni(\z|\W|_)`):    "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)dhieec(\z|\W|_)`):          "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)thaak(\z|\W|_)`):           "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)ruun(\z|\W|_)`):            "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):      "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):       "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):       "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)tiop thar pɛt(\z|\W|_)`):   "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)tiop(\z|\W|_)`):            "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)pay yie̱tni(\z|\W|_)`):     "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)pay(\z|\W|_)`):             "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)kor(\z|\W|_)`):             "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)kornyoot(\z|\W|_)`):        "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)duɔ̱ɔ̱(\z|\W|_)`):          "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)duɔ̱ɔ̱ŋ(\z|\W|_)`):         "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)duä(\z|\W|_)`):             "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)duät(\z|\W|_)`):            "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)jiec la̱t(\z|\W|_)`):       "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)jiec(\z|\W|_)`):            "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)kur(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)laa(\z|\W|_)`):             "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)laath(\z|\W|_)`):           "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)tŋ(\z|\W|_)`):              "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)bäkɛl lätni(\z|\W|_)`):     "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)bäkɛl(\z|\W|_)`):           "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)thɛkɛni(\z|\W|_)`):         "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)tɛɛ(\z|\W|_)`):             "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)tɛɛr(\z|\W|_)`):            "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)cäŋ kuɔth(\z|\W|_)`):       "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)cäŋ(\z|\W|_)`):             "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)ŋuaan lätni(\z|\W|_)`):     "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)ŋuaan(\z|\W|_)`):           "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)rɛw lätni(\z|\W|_)`):       "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)rɛw(\z|\W|_)`):             "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)diɔ̱k lätni(\z|\W|_)`):     "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)diɔ̱k(\z|\W|_)`):           "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)jiɔk(\z|\W|_)`):            "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)ruɔ̱n(\z|\W|_)`):           "${1}year${2}",
	},
}
