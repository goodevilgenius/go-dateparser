// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var yav_Locale = LocaleData{
	Name:      "yav",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)ínaan(\z|\W|_)`):                      "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):                  "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):                "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):                 "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):                        "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):                  "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):                  "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)púyoó(\z|\W|_)`):                      "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):                 "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):                  "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):                  "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)kiɛmɛ́ɛm(\z|\W|_)`):                   "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)mesiŋ, oóli ú kénie(\z|\W|_)`):        "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)o4(\z|\W|_)`):                         "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)o8(\z|\W|_)`):                         "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)pisuyú(\z|\W|_)`):                     "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)puɔ́sɛ́(\z|\W|_)`):                    "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)o12(\z|\W|_)`):                        "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)pilɔndɔ́(\z|\W|_)`):                   "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)o2(\z|\W|_)`):                         "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)siɛyɛ́, oóli ú kándíɛ(\z|\W|_)`):      "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)feléte(\z|\W|_)`):                     "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)fl(\z|\W|_)`):                         "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)kisikɛl,(\z|\W|_)`):                   "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)nakinyám(\z|\W|_)`):                   "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):                 "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):                  "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):                  "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)o1(\z|\W|_)`):                         "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)pikítíkítie, oólí ú kutúan(\z|\W|_)`): "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)efute(\z|\W|_)`):                      "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)o7(\z|\W|_)`):                         "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)o6(\z|\W|_)`):                         "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)ɔsɔn(\z|\W|_)`):                       "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)o3(\z|\W|_)`):                         "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)ɔnsúmbɔl, oóli ú kátátúɛ(\z|\W|_)`):   "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)ensil, oóli ú kátánuɛ(\z|\W|_)`):      "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)o5(\z|\W|_)`):                         "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)md(\z|\W|_)`):                         "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)móndie(\z|\W|_)`):                     "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)oóli(\z|\W|_)`):                       "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)makandikɛ(\z|\W|_)`):                  "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)o11(\z|\W|_)`):                        "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)imɛŋ i putúk,oóli ú kátíɛ(\z|\W|_)`):  "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)o10(\z|\W|_)`):                        "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)kisɛ́ndɛ(\z|\W|_)`):                   "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)ss(\z|\W|_)`):                         "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)séselé(\z|\W|_)`):                     "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)síkɛn(\z|\W|_)`):                      "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)imɛŋ i puɔs(\z|\W|_)`):                "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)o9(\z|\W|_)`):                         "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)sd(\z|\W|_)`):                         "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)sɔ́ndiɛ(\z|\W|_)`):                    "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)kl(\z|\W|_)`):                         "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)kúpélimetúkpiapɛ(\z|\W|_)`):           "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)muányáŋmóndie(\z|\W|_)`):              "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)mw(\z|\W|_)`):                         "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)et(\z|\W|_)`):                         "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)metúkpíápɛ(\z|\W|_)`):                 "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)yɔɔŋ(\z|\W|_)`):                       "${1}year${2}",
	},
}
