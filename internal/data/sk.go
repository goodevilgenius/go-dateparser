// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sk_Locale = LocaleData{
	Name:      "sk",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)pred (\d+) d(\z|\W|_)`):         "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) dňami(\z|\W|_)`):     "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) dňom(\z|\W|_)`):      "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) h(\z|\W|_)`):         "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) hodinami(\z|\W|_)`):  "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) hodinou(\z|\W|_)`):   "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) mes(\z|\W|_)`):       "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) mesiacmi(\z|\W|_)`):  "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) mesiacom(\z|\W|_)`):  "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) s(\z|\W|_)`):         "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) sekundami(\z|\W|_)`): "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) sekundou(\z|\W|_)`):  "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) týž(\z|\W|_)`):       "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) týždňami(\z|\W|_)`):  "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) týždňom(\z|\W|_)`):   "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) r(\z|\W|_)`):         "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) rokmi(\z|\W|_)`):     "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)pred (\d+) rokom(\z|\W|_)`):     "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)dnes(\z|\W|_)`):                 "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)v tejto hodine(\z|\W|_)`):       "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)v tejto minúte(\z|\W|_)`):       "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)tento mesiac(\z|\W|_)`):         "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)teraz(\z|\W|_)`):                "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)tento týždeň(\z|\W|_)`):         "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)tento rok(\z|\W|_)`):            "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)včera(\z|\W|_)`):                "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)minulý mesiac(\z|\W|_)`):        "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)minulý týždeň(\z|\W|_)`):        "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)minulý rok(\z|\W|_)`):           "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`):                   "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`):                  "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)apríl(\z|\W|_)`):                "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)apríla(\z|\W|_)`):               "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`):                  "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)august(\z|\W|_)`):               "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)augusta(\z|\W|_)`):              "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)d(\z|\W|_)`):                    "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)deň(\z|\W|_)`):                  "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`):                  "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)december(\z|\W|_)`):             "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)decembra(\z|\W|_)`):             "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`):                  "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)február(\z|\W|_)`):              "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)februára(\z|\W|_)`):             "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)pi(\z|\W|_)`):                   "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)piatok(\z|\W|_)`):               "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`):                    "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)hodina(\z|\W|_)`):               "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) d(\z|\W|_)`):            "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) deň(\z|\W|_)`):          "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) dní(\z|\W|_)`):          "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) h(\z|\W|_)`):            "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) hodinu(\z|\W|_)`):       "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) hodín(\z|\W|_)`):        "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) min(\z|\W|_)`):          "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) minút(\z|\W|_)`):        "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) minútu(\z|\W|_)`):       "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) mes(\z|\W|_)`):          "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) mesiac(\z|\W|_)`):       "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) mesiacov(\z|\W|_)`):     "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) s(\z|\W|_)`):            "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) sekundu(\z|\W|_)`):      "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) sekúnd(\z|\W|_)`):       "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) týž(\z|\W|_)`):          "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) týždeň(\z|\W|_)`):       "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) týždňov(\z|\W|_)`):      "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) r(\z|\W|_)`):            "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) rok(\z|\W|_)`):          "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)o (\d+) rokov(\z|\W|_)`):        "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)zajtra(\z|\W|_)`):               "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)budúci mesiac(\z|\W|_)`):        "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)budúci týždeň(\z|\W|_)`):        "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)budúci rok(\z|\W|_)`):           "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`):                  "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)január(\z|\W|_)`):               "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)januára(\z|\W|_)`):              "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)júl(\z|\W|_)`):                  "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)júla(\z|\W|_)`):                 "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)jún(\z|\W|_)`):                  "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)júna(\z|\W|_)`):                 "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`):                  "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)marca(\z|\W|_)`):                "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)marec(\z|\W|_)`):                "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)máj(\z|\W|_)`):                  "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)mája(\z|\W|_)`):                 "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)po(\z|\W|_)`):                   "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)pondelok(\z|\W|_)`):             "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mes(\z|\W|_)`):                  "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)mesiac(\z|\W|_)`):               "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`):                  "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)november(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)novembra(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`):                  "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)október(\z|\W|_)`):              "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)októbra(\z|\W|_)`):              "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`):                   "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)so(\z|\W|_)`):                   "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)sobota(\z|\W|_)`):               "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`):                    "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sekunda(\z|\W|_)`):              "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`):                  "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)september(\z|\W|_)`):            "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)septembra(\z|\W|_)`):            "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)ne(\z|\W|_)`):                   "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)nedeľa(\z|\W|_)`):               "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)št(\z|\W|_)`):                   "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)štvrtok(\z|\W|_)`):              "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)ut(\z|\W|_)`):                   "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)utorok(\z|\W|_)`):               "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)st(\z|\W|_)`):                   "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)streda(\z|\W|_)`):               "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)týž(\z|\W|_)`):                  "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)týždeň(\z|\W|_)`):               "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)r(\z|\W|_)`):                    "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)rok(\z|\W|_)`):                  "${1}year${2}",
	},
}
