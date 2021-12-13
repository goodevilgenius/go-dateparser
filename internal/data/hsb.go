// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hsb_Locale = LocaleData{
	Name:      "hsb",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)před (\d+) d(\z|\W|_)`):          "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) dnj(\z|\W|_)`):        "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) dnjemi(\z|\W|_)`):     "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) dnjom(\z|\W|_)`):      "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) h(\z|\W|_)`):          "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) hodź(\z|\W|_)`):       "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) hodźinami(\z|\W|_)`):  "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) hodźinu(\z|\W|_)`):    "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) měs(\z|\W|_)`):        "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) měsacami(\z|\W|_)`):   "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) měsacom(\z|\W|_)`):    "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) s(\z|\W|_)`):          "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) sek(\z|\W|_)`):        "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) sekundami(\z|\W|_)`):  "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) sekundu(\z|\W|_)`):    "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) tydź(\z|\W|_)`):       "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) tydźenjemi(\z|\W|_)`): "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) tydźenjom(\z|\W|_)`):  "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) l(\z|\W|_)`):          "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) lětami(\z|\W|_)`):     "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)před (\d+) lětom(\z|\W|_)`):      "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)dźensa(\z|\W|_)`):                "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):             "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):           "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)tutón měsac(\z|\W|_)`):           "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)tutón tydźeń(\z|\W|_)`):          "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)lětsa(\z|\W|_)`):                 "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)wčera(\z|\W|_)`):                 "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)zašły měsac(\z|\W|_)`):           "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)zašły tydźeń(\z|\W|_)`):          "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)loni(\z|\W|_)`):                  "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)dopołdnja(\z|\W|_)`):             "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`):                   "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)apryl(\z|\W|_)`):                 "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)apryla(\z|\W|_)`):                "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)awg(\z|\W|_)`):                   "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)awgust(\z|\W|_)`):                "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)awgusta(\z|\W|_)`):               "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)d(\z|\W|_)`):                     "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dźeń(\z|\W|_)`):                  "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`):                   "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)december(\z|\W|_)`):              "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)decembra(\z|\W|_)`):              "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`):                   "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)februar(\z|\W|_)`):               "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)februara(\z|\W|_)`):              "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)pja(\z|\W|_)`):                   "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)pjatk(\z|\W|_)`):                 "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`):                     "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)hodź(\z|\W|_)`):                  "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)hodźina(\z|\W|_)`):               "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) d(\z|\W|_)`):            "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) dnj(\z|\W|_)`):          "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) dnjow(\z|\W|_)`):        "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) dźeń(\z|\W|_)`):         "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) h(\z|\W|_)`):            "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) hodź(\z|\W|_)`):         "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) hodźin(\z|\W|_)`):       "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) hodźinu(\z|\W|_)`):      "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) m(\z|\W|_)`):            "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) min(\z|\W|_)`):          "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) minutow(\z|\W|_)`):      "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) minutu(\z|\W|_)`):       "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) měs(\z|\W|_)`):          "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) měsac(\z|\W|_)`):        "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) měsacow(\z|\W|_)`):      "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) s(\z|\W|_)`):            "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sek(\z|\W|_)`):          "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sekundow(\z|\W|_)`):     "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) sekundu(\z|\W|_)`):      "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) tydź(\z|\W|_)`):         "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) tydźenjow(\z|\W|_)`):    "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) tydźeń(\z|\W|_)`):       "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) l(\z|\W|_)`):            "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) lět(\z|\W|_)`):          "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)za (\d+) lěto(\z|\W|_)`):         "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)jutře(\z|\W|_)`):                 "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)přichodny měsac(\z|\W|_)`):       "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)přichodny tydźeń(\z|\W|_)`):      "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)klětu(\z|\W|_)`):                 "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`):                   "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)januar(\z|\W|_)`):                "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)januara(\z|\W|_)`):               "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`):                   "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)julij(\z|\W|_)`):                 "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)julija(\z|\W|_)`):                "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`):                   "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)junij(\z|\W|_)`):                 "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)junija(\z|\W|_)`):                "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)měr(\z|\W|_)`):                   "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)měrc(\z|\W|_)`):                  "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)měrca(\z|\W|_)`):                 "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)mej(\z|\W|_)`):                   "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)meja(\z|\W|_)`):                  "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)meje(\z|\W|_)`):                  "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)pón(\z|\W|_)`):                   "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)póndźela(\z|\W|_)`):              "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)měs(\z|\W|_)`):                   "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)měsac(\z|\W|_)`):                 "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):                   "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)nowember(\z|\W|_)`):              "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)nowembra(\z|\W|_)`):              "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`):                   "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)oktober(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)oktobra(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)popołdnju(\z|\W|_)`):             "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)sob(\z|\W|_)`):                   "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)sobota(\z|\W|_)`):                "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`):                     "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`):                   "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sekunda(\z|\W|_)`):               "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`):                   "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)september(\z|\W|_)`):             "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)septembra(\z|\W|_)`):             "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)nje(\z|\W|_)`):                   "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)njedźela(\z|\W|_)`):              "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)štw(\z|\W|_)`):                   "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)štwórtk(\z|\W|_)`):               "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)wut(\z|\W|_)`):                   "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)wutora(\z|\W|_)`):                "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)srj(\z|\W|_)`):                   "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)srjeda(\z|\W|_)`):                "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)tydź(\z|\W|_)`):                  "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)tydźeń(\z|\W|_)`):                "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)l(\z|\W|_)`):                     "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)lěto(\z|\W|_)`):                  "${1}year${2}",
	},
}
