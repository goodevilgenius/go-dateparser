// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kln_Locale = LocaleData{
	Name:      "kln",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)raini(\z|\W|_)`):                "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):            "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):          "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):           "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):                  "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):            "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):            "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)amut(\z|\W|_)`):                 "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):           "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):            "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):            "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)karoon(\z|\W|_)`):               "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)krn(\z|\W|_)`):                  "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)iwo(\z|\W|_)`):                  "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)iwootkuut(\z|\W|_)`):            "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)roo(\z|\W|_)`):                  "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)rooptui(\z|\W|_)`):              "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)betut(\z|\W|_)`):                "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)kipsuunde nebo aeng'(\z|\W|_)`): "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)kpa(\z|\W|_)`):                  "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)ng'atyaato(\z|\W|_)`):           "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)ngat(\z|\W|_)`):                 "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)kom(\z|\W|_)`):                  "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)komuut(\z|\W|_)`):               "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)sait(\z|\W|_)`):                 "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)mutai(\z|\W|_)`):                "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):           "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):            "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):            "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)mul(\z|\W|_)`):                  "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)mulgul(\z|\W|_)`):               "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)ng'eiyeet(\z|\W|_)`):            "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)nge(\z|\W|_)`):                  "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)paa(\z|\W|_)`):                  "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)paagi(\z|\W|_)`):                "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)kiptaamo(\z|\W|_)`):             "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)taa(\z|\W|_)`):                  "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mam(\z|\W|_)`):                  "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)mamuut(\z|\W|_)`):               "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)kot(\z|\W|_)`):                  "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)kotaai(\z|\W|_)`):               "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)arawet(\z|\W|_)`):               "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)kipsuunde ne taai(\z|\W|_)`):    "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)kpt(\z|\W|_)`):                  "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)epe(\z|\W|_)`):                  "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)epeeso(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)koosk(\z|\W|_)`):                "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)kooskoliny(\z|\W|_)`):           "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)kol(\z|\W|_)`):                  "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)kolo(\z|\W|_)`):                 "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)sekondit(\z|\W|_)`):             "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)bur(\z|\W|_)`):                  "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)bureet(\z|\W|_)`):               "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)kotisap(\z|\W|_)`):              "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)kts(\z|\W|_)`):                  "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)koa(\z|\W|_)`):                  "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)koang'wan(\z|\W|_)`):            "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)koaeng'(\z|\W|_)`):              "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)koo(\z|\W|_)`):                  "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)kos(\z|\W|_)`):                  "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)kosomok(\z|\W|_)`):              "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)wikit(\z|\W|_)`):                "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)kenyit(\z|\W|_)`):               "${1}year${2}",
	},
}
