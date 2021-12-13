// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var af_Locale = LocaleData{
	Name:      "af",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)(\d+) dae gelede(\z|\W|_)`):      "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) dag gelede(\z|\W|_)`):      "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) uur gelede(\z|\W|_)`):      "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) maand gelede(\z|\W|_)`):    "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) maande gelede(\z|\W|_)`):   "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) md gelede(\z|\W|_)`):       "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) sek gelede(\z|\W|_)`):      "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) sekonde gelede(\z|\W|_)`):  "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) sekondes gelede(\z|\W|_)`): "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) w gelede(\z|\W|_)`):        "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) week gelede(\z|\W|_)`):     "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) weke gelede(\z|\W|_)`):     "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) jaar gelede(\z|\W|_)`):     "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)vandag(\z|\W|_)`):                "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)hierdie uur(\z|\W|_)`):           "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)hierdie minuut(\z|\W|_)`):        "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)vandeesmaand(\z|\W|_)`):          "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)nou(\z|\W|_)`):                   "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)vandeesweek(\z|\W|_)`):           "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)hierdie jaar(\z|\W|_)`):          "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)gister(\z|\W|_)`):                "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)verlede maand(\z|\W|_)`):         "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)verlede week(\z|\W|_)`):          "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)verlede jaar(\z|\W|_)`):          "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)vm(\z|\W|_)`):                    "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`):                   "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)april(\z|\W|_)`):                 "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`):                   "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)augustus(\z|\W|_)`):              "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)d(\z|\W|_)`):                     "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dag(\z|\W|_)`):                   "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)des(\z|\W|_)`):                   "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)desember(\z|\W|_)`):              "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`):                   "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)februarie(\z|\W|_)`):             "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)vr(\z|\W|_)`):                    "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)vrydag(\z|\W|_)`):                "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)u(\z|\W|_)`):                     "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)uur(\z|\W|_)`):                   "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) dae(\z|\W|_)`):         "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) dag(\z|\W|_)`):         "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) uur(\z|\W|_)`):         "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) min(\z|\W|_)`):         "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) md(\z|\W|_)`):          "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) minuut(\z|\W|_)`):      "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) sek(\z|\W|_)`):         "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) sekonde(\z|\W|_)`):     "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) sekondes(\z|\W|_)`):    "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) w(\z|\W|_)`):           "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) week(\z|\W|_)`):        "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) weke(\z|\W|_)`):        "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)oor (\d+) jaar(\z|\W|_)`):        "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)môre(\z|\W|_)`):                  "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)volgende maand(\z|\W|_)`):        "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)volgende week(\z|\W|_)`):         "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)volgende jaar(\z|\W|_)`):         "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`):                   "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)januarie(\z|\W|_)`):              "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`):                   "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)julie(\z|\W|_)`):                 "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`):                   "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)junie(\z|\W|_)`):                 "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)maart(\z|\W|_)`):                 "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mrt(\z|\W|_)`):                   "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mei(\z|\W|_)`):                   "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)ma(\z|\W|_)`):                    "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)maandag(\z|\W|_)`):               "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)maand(\z|\W|_)`):                 "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)md(\z|\W|_)`):                    "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`):                   "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)november(\z|\W|_)`):              "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`):                   "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)oktober(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)nm(\z|\W|_)`):                    "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)sa(\z|\W|_)`):                    "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)saterdag(\z|\W|_)`):              "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`):                     "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`):                   "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sekonde(\z|\W|_)`):               "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`):                   "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)september(\z|\W|_)`):             "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)so(\z|\W|_)`):                    "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)sondag(\z|\W|_)`):                "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)do(\z|\W|_)`):                    "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)donderdag(\z|\W|_)`):             "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)di(\z|\W|_)`):                    "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)dinsdag(\z|\W|_)`):               "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)wo(\z|\W|_)`):                    "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)woensdag(\z|\W|_)`):              "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)week(\z|\W|_)`):                  "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)wk(\z|\W|_)`):                    "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)j(\z|\W|_)`):                     "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)jaar(\z|\W|_)`):                  "${1}year${2}",
	},
}

var af_NA_Locale = LocaleData{
	Name:      "af-NA",
	Parent:    &af_Locale,
	DateOrder: "YMD",
}
