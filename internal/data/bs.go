// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Locale = LocaleData{
	Name:      "bs",
	DateOrder: "DMY.",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)prije (\d+) d(\z|\W|_)`):       "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) dan(\z|\W|_)`):     "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) dana(\z|\W|_)`):    "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sat(\z|\W|_)`):     "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sati(\z|\W|_)`):    "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) mj(\z|\W|_)`):      "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) mjesec(\z|\W|_)`):  "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) mjeseci(\z|\W|_)`): "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sek(\z|\W|_)`):     "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sekundi(\z|\W|_)`): "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sekundu(\z|\W|_)`): "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sed(\z|\W|_)`):     "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sedmica(\z|\W|_)`): "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) sedmicu(\z|\W|_)`): "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) g(\z|\W|_)`):       "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) god(\z|\W|_)`):     "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) godina(\z|\W|_)`):  "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)prije (\d+) godinu(\z|\W|_)`):  "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)danas(\z|\W|_)`):               "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)ovaj sat(\z|\W|_)`):            "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)ova minuta(\z|\W|_)`):          "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)ovaj mjesec(\z|\W|_)`):         "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)sada(\z|\W|_)`):                "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)ove sedmice(\z|\W|_)`):         "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)ove godine(\z|\W|_)`):          "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)jučer(\z|\W|_)`):               "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)prošli mjesec(\z|\W|_)`):       "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)prošle sedmice(\z|\W|_)`):      "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)prošle godine(\z|\W|_)`):       "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)prijepodne(\z|\W|_)`):          "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`):                 "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)april(\z|\W|_)`):               "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)avg(\z|\W|_)`):                 "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)avgust(\z|\W|_)`):              "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)dan(\z|\W|_)`):                 "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`):                 "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)decembar(\z|\W|_)`):            "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`):                 "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)februar(\z|\W|_)`):             "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)pet(\z|\W|_)`):                 "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)petak(\z|\W|_)`):               "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`):                   "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)sat(\z|\W|_)`):                 "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) d(\z|\W|_)`):          "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) dan(\z|\W|_)`):        "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) dana(\z|\W|_)`):       "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sat(\z|\W|_)`):        "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sati(\z|\W|_)`):       "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) min(\z|\W|_)`):        "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) minuta(\z|\W|_)`):     "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) minutu(\z|\W|_)`):     "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) mj(\z|\W|_)`):         "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) mjesec(\z|\W|_)`):     "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) mjeseci(\z|\W|_)`):    "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sek(\z|\W|_)`):        "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sekundi(\z|\W|_)`):    "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sekundu(\z|\W|_)`):    "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sed(\z|\W|_)`):        "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sedmica(\z|\W|_)`):    "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sedmicu(\z|\W|_)`):    "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) g(\z|\W|_)`):          "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) god(\z|\W|_)`):        "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) godina(\z|\W|_)`):     "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) godinu(\z|\W|_)`):     "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)sutra(\z|\W|_)`):               "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)sljedeći mjesec(\z|\W|_)`):     "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)sljedeće sedmice(\z|\W|_)`):    "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)sljedeće godine(\z|\W|_)`):     "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`):                 "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)januar(\z|\W|_)`):              "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`):                 "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)juli(\z|\W|_)`):                "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`):                 "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)juni(\z|\W|_)`):                "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`):                 "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mart(\z|\W|_)`):                "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)maj(\z|\W|_)`):                 "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)pon(\z|\W|_)`):                 "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)ponedjeljak(\z|\W|_)`):         "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mj(\z|\W|_)`):                  "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)mjesec(\z|\W|_)`):              "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`):                 "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)novembar(\z|\W|_)`):            "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`):                 "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)oktobar(\z|\W|_)`):             "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)popodne(\z|\W|_)`):             "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)sub(\z|\W|_)`):                 "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)subota(\z|\W|_)`):              "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`):                   "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`):                 "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sekunda(\z|\W|_)`):             "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`):                 "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)septembar(\z|\W|_)`):           "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)ned(\z|\W|_)`):                 "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)nedjelja(\z|\W|_)`):            "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)čet(\z|\W|_)`):                 "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)četvrtak(\z|\W|_)`):            "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)uto(\z|\W|_)`):                 "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)utorak(\z|\W|_)`):              "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)sri(\z|\W|_)`):                 "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)srijeda(\z|\W|_)`):             "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)sed(\z|\W|_)`):                 "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)sedmica(\z|\W|_)`):             "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)g(\z|\W|_)`):                   "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)god(\z|\W|_)`):                 "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)godina(\z|\W|_)`):              "${1}year${2}",
	},
}
