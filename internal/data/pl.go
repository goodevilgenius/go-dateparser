// Code is generated by script; DO NOT EDIT.

package data

var pl_Locale = LocaleData{
	Name:      "pl",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "i", "o", "r.", "roku", "w", "|"},
	Simplifications: map[string]string{
		`(\A|\W|_)dzisiaj(\z|\W|_)`:      "${1}0 dnia${2}",
		`(\A|\W|_)dziś(\z|\W|_)`:         "${1}0 dnia${2}",
		`(\A|\W|_)jutro(\z|\W|_)`:        "${1}za 1 dnia${2}",
		`(\A|\W|_)onegdaj(\z|\W|_)`:      "${1}2 dnia${2}",
		`(\A|\W|_)pojutrze(\z|\W|_)`:     "${1}za 2 dnia${2}",
		`(\A|\W|_)przedwczoraj(\z|\W|_)`: "${1}2 dnia temu${2}",
	},
	Translations: map[string]string{
		`(\A|\W|_)(\d+) dnia temu(\z|\W|_)`:      "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) dzień temu(\z|\W|_)`:     "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) g temu(\z|\W|_)`:         "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) godz temu(\z|\W|_)`:      "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) godziny temu(\z|\W|_)`:   "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) godzinę temu(\z|\W|_)`:   "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) mies temu(\z|\W|_)`:      "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) miesiąc temu(\z|\W|_)`:   "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) miesiąca temu(\z|\W|_)`:  "${1}$2 month ago${3}",
		`(\A|\W|_)–(\d+) mies(\z|\W|_)`:          "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) s temu(\z|\W|_)`:         "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) sek temu(\z|\W|_)`:       "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) sekundy temu(\z|\W|_)`:   "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) sekundę temu(\z|\W|_)`:   "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) tydz temu(\z|\W|_)`:      "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) tydzień temu(\z|\W|_)`:   "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) tyg temu(\z|\W|_)`:       "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) tygodnia temu(\z|\W|_)`:  "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) rok temu(\z|\W|_)`:       "${1}$2 year ago${3}",
		`(\A|\W|_)(\d+) roku temu(\z|\W|_)`:      "${1}$2 year ago${3}",
		`(\A|\W|_)dzisiaj(\z|\W|_)`:              "${1}0 day ago${2}",
		`(\A|\W|_)ta godzina(\z|\W|_)`:           "${1}0 hour ago${2}",
		`(\A|\W|_)ta minuta(\z|\W|_)`:            "${1}0 minute ago${2}",
		`(\A|\W|_)w tym miesiącu(\z|\W|_)`:       "${1}0 month ago${2}",
		`(\A|\W|_)teraz(\z|\W|_)`:                "${1}0 second ago${2}",
		`(\A|\W|_)w tym tygodniu(\z|\W|_)`:       "${1}0 week ago${2}",
		`(\A|\W|_)w tym roku(\z|\W|_)`:           "${1}0 year ago${2}",
		`(\A|\W|_)wczoraj(\z|\W|_)`:              "${1}1 day ago${2}",
		`(\A|\W|_)w zeszłym miesiącu(\z|\W|_)`:   "${1}1 month ago${2}",
		`(\A|\W|_)w zeszłym tygodniu(\z|\W|_)`:   "${1}1 week ago${2}",
		`(\A|\W|_)w zeszłym roku(\z|\W|_)`:       "${1}1 year ago${2}",
		`(\A|\W|_)temu(\z|\W|_)`:                 "${1}ago${2}",
		`(\A|\W|_)am(\z|\W|_)`:                   "${1}am${2}",
		`(\A|\W|_)kwi(\z|\W|_)`:                  "${1}april${2}",
		`(\A|\W|_)kwie(\z|\W|_)`:                 "${1}april${2}",
		`(\A|\W|_)kwiecien(\z|\W|_)`:             "${1}april${2}",
		`(\A|\W|_)kwiecień(\z|\W|_)`:             "${1}april${2}",
		`(\A|\W|_)kwietnia(\z|\W|_)`:             "${1}april${2}",
		`(\A|\W|_)kwietniu(\z|\W|_)`:             "${1}april${2}",
		`(\A|\W|_)sie(\z|\W|_)`:                  "${1}august${2}",
		`(\A|\W|_)sierpien(\z|\W|_)`:             "${1}august${2}",
		`(\A|\W|_)sierpień(\z|\W|_)`:             "${1}august${2}",
		`(\A|\W|_)sierpnia(\z|\W|_)`:             "${1}august${2}",
		`(\A|\W|_)sierpniu(\z|\W|_)`:             "${1}august${2}",
		`(\A|\W|_)dni(\z|\W|_)`:                  "${1}day${2}",
		`(\A|\W|_)dnia(\z|\W|_)`:                 "${1}day${2}",
		`(\A|\W|_)dniu(\z|\W|_)`:                 "${1}day${2}",
		`(\A|\W|_)dzien(\z|\W|_)`:                "${1}day${2}",
		`(\A|\W|_)dzień(\z|\W|_)`:                "${1}day${2}",
		`(\A|\W|_)gru(\z|\W|_)`:                  "${1}december${2}",
		`(\A|\W|_)grudnia(\z|\W|_)`:              "${1}december${2}",
		`(\A|\W|_)grudniu(\z|\W|_)`:              "${1}december${2}",
		`(\A|\W|_)grudzien(\z|\W|_)`:             "${1}december${2}",
		`(\A|\W|_)grudzień(\z|\W|_)`:             "${1}december${2}",
		`(\A|\W|_)lut(\z|\W|_)`:                  "${1}february${2}",
		`(\A|\W|_)lutego(\z|\W|_)`:               "${1}february${2}",
		`(\A|\W|_)luty(\z|\W|_)`:                 "${1}february${2}",
		`(\A|\W|_)lutym(\z|\W|_)`:                "${1}february${2}",
		`(\A|\W|_)pi(\z|\W|_)`:                   "${1}friday${2}",
		`(\A|\W|_)pia(\z|\W|_)`:                  "${1}friday${2}",
		`(\A|\W|_)piatek(\z|\W|_)`:               "${1}friday${2}",
		`(\A|\W|_)pią(\z|\W|_)`:                  "${1}friday${2}",
		`(\A|\W|_)piątek(\z|\W|_)`:               "${1}friday${2}",
		`(\A|\W|_)pt(\z|\W|_)`:                   "${1}friday${2}",
		`(\A|\W|_)g(\z|\W|_)`:                    "${1}hour${2}",
		`(\A|\W|_)godz(\z|\W|_)`:                 "${1}hour${2}",
		`(\A|\W|_)godzin(\z|\W|_)`:               "${1}hour${2}",
		`(\A|\W|_)godzina(\z|\W|_)`:              "${1}hour${2}",
		`(\A|\W|_)godzine(\z|\W|_)`:              "${1}hour${2}",
		`(\A|\W|_)godziny(\z|\W|_)`:              "${1}hour${2}",
		`(\A|\W|_)godzinę(\z|\W|_)`:              "${1}hour${2}",
		`(\A|\W|_)za (\d+) dnia(\z|\W|_)`:        "${1}in $2 day${3}",
		`(\A|\W|_)za (\d+) dzień(\z|\W|_)`:       "${1}in $2 day${3}",
		`(\A|\W|_)za (\d+) g(\z|\W|_)`:           "${1}in $2 hour${3}",
		`(\A|\W|_)za (\d+) godz(\z|\W|_)`:        "${1}in $2 hour${3}",
		`(\A|\W|_)za (\d+) godziny(\z|\W|_)`:     "${1}in $2 hour${3}",
		`(\A|\W|_)za (\d+) godzinę(\z|\W|_)`:     "${1}in $2 hour${3}",
		`(\A|\W|_)za (\d+) min(\z|\W|_)`:         "${1}in $2 minute${3}",
		`(\A|\W|_)za (\d+) minuty(\z|\W|_)`:      "${1}in $2 minute${3}",
		`(\A|\W|_)za (\d+) minutę(\z|\W|_)`:      "${1}in $2 minute${3}",
		`(\A|\W|_)za (\d+) mies(\z|\W|_)`:        "${1}in $2 month${3}",
		`(\A|\W|_)za (\d+) miesiąc(\z|\W|_)`:     "${1}in $2 month${3}",
		`(\A|\W|_)za (\d+) miesiąca(\z|\W|_)`:    "${1}in $2 month${3}",
		`(\A|\W|_)za (\d+) s(\z|\W|_)`:           "${1}in $2 second${3}",
		`(\A|\W|_)za (\d+) sek(\z|\W|_)`:         "${1}in $2 second${3}",
		`(\A|\W|_)za (\d+) sekundy(\z|\W|_)`:     "${1}in $2 second${3}",
		`(\A|\W|_)za (\d+) sekundę(\z|\W|_)`:     "${1}in $2 second${3}",
		`(\A|\W|_)za (\d+) tydz(\z|\W|_)`:        "${1}in $2 week${3}",
		`(\A|\W|_)za (\d+) tydzień(\z|\W|_)`:     "${1}in $2 week${3}",
		`(\A|\W|_)za (\d+) tyg(\z|\W|_)`:         "${1}in $2 week${3}",
		`(\A|\W|_)za (\d+) tygodnia(\z|\W|_)`:    "${1}in $2 week${3}",
		`(\A|\W|_)za (\d+) rok(\z|\W|_)`:         "${1}in $2 year${3}",
		`(\A|\W|_)za (\d+) roku(\z|\W|_)`:        "${1}in $2 year${3}",
		`(\A|\W|_)jutro(\z|\W|_)`:                "${1}in 1 day${2}",
		`(\A|\W|_)w przyszłym miesiącu(\z|\W|_)`: "${1}in 1 month${2}",
		`(\A|\W|_)w przyszłym tygodniu(\z|\W|_)`: "${1}in 1 week${2}",
		`(\A|\W|_)w przyszłym roku(\z|\W|_)`:     "${1}in 1 year${2}",
		`(\A|\W|_)za(\z|\W|_)`:                   "${1}in${2}",
		`(\A|\W|_)sty(\z|\W|_)`:                  "${1}january${2}",
		`(\A|\W|_)styczen(\z|\W|_)`:              "${1}january${2}",
		`(\A|\W|_)styczeń(\z|\W|_)`:              "${1}january${2}",
		`(\A|\W|_)stycznia(\z|\W|_)`:             "${1}january${2}",
		`(\A|\W|_)styczniu(\z|\W|_)`:             "${1}january${2}",
		`(\A|\W|_)lip(\z|\W|_)`:                  "${1}july${2}",
		`(\A|\W|_)lipca(\z|\W|_)`:                "${1}july${2}",
		`(\A|\W|_)lipcu(\z|\W|_)`:                "${1}july${2}",
		`(\A|\W|_)lipiec(\z|\W|_)`:               "${1}july${2}",
		`(\A|\W|_)cze(\z|\W|_)`:                  "${1}june${2}",
		`(\A|\W|_)czerwca(\z|\W|_)`:              "${1}june${2}",
		`(\A|\W|_)czerwcu(\z|\W|_)`:              "${1}june${2}",
		`(\A|\W|_)czerwiec(\z|\W|_)`:             "${1}june${2}",
		`(\A|\W|_)mar(\z|\W|_)`:                  "${1}march${2}",
		`(\A|\W|_)marca(\z|\W|_)`:                "${1}march${2}",
		`(\A|\W|_)marcu(\z|\W|_)`:                "${1}march${2}",
		`(\A|\W|_)marzec(\z|\W|_)`:               "${1}march${2}",
		`(\A|\W|_)maj(\z|\W|_)`:                  "${1}may${2}",
		`(\A|\W|_)maja(\z|\W|_)`:                 "${1}may${2}",
		`(\A|\W|_)maju(\z|\W|_)`:                 "${1}may${2}",
		`(\A|\W|_)minut(\z|\W|_)`:                "${1}minute${2}",
		`(\A|\W|_)minute(\z|\W|_)`:               "${1}minute${2}",
		`(\A|\W|_)minuty(\z|\W|_)`:               "${1}minute${2}",
		`(\A|\W|_)minutę(\z|\W|_)`:               "${1}minute${2}",
		`(\A|\W|_)pn(\z|\W|_)`:                   "${1}monday${2}",
		`(\A|\W|_)pon(\z|\W|_)`:                  "${1}monday${2}",
		`(\A|\W|_)poniedzialek(\z|\W|_)`:         "${1}monday${2}",
		`(\A|\W|_)poniedziałek(\z|\W|_)`:         "${1}monday${2}",
		`(\A|\W|_)mc(\z|\W|_)`:                   "${1}month${2}",
		`(\A|\W|_)mies(\z|\W|_)`:                 "${1}month${2}",
		`(\A|\W|_)miesiac(\z|\W|_)`:              "${1}month${2}",
		`(\A|\W|_)miesiace(\z|\W|_)`:             "${1}month${2}",
		`(\A|\W|_)miesiąc(\z|\W|_)`:              "${1}month${2}",
		`(\A|\W|_)miesiąca(\z|\W|_)`:             "${1}month${2}",
		`(\A|\W|_)miesiące(\z|\W|_)`:             "${1}month${2}",
		`(\A|\W|_)miesięcy(\z|\W|_)`:             "${1}month${2}",
		`(\A|\W|_)lis(\z|\W|_)`:                  "${1}november${2}",
		`(\A|\W|_)listopad(\z|\W|_)`:             "${1}november${2}",
		`(\A|\W|_)listopada(\z|\W|_)`:            "${1}november${2}",
		`(\A|\W|_)listopadzie(\z|\W|_)`:          "${1}november${2}",
		`(\A|\W|_)paz(\z|\W|_)`:                  "${1}october${2}",
		`(\A|\W|_)pazdziernik(\z|\W|_)`:          "${1}october${2}",
		`(\A|\W|_)pazdziernika(\z|\W|_)`:         "${1}october${2}",
		`(\A|\W|_)pazdzierniku(\z|\W|_)`:         "${1}october${2}",
		`(\A|\W|_)paź(\z|\W|_)`:                  "${1}october${2}",
		`(\A|\W|_)październik(\z|\W|_)`:          "${1}october${2}",
		`(\A|\W|_)października(\z|\W|_)`:         "${1}october${2}",
		`(\A|\W|_)październiku(\z|\W|_)`:         "${1}october${2}",
		`(\A|\W|_)pm(\z|\W|_)`:                   "${1}pm${2}",
		`(\A|\W|_)sb(\z|\W|_)`:                   "${1}saturday${2}",
		`(\A|\W|_)so(\z|\W|_)`:                   "${1}saturday${2}",
		`(\A|\W|_)sob(\z|\W|_)`:                  "${1}saturday${2}",
		`(\A|\W|_)sobota(\z|\W|_)`:               "${1}saturday${2}",
		`(\A|\W|_)sobote(\z|\W|_)`:               "${1}saturday${2}",
		`(\A|\W|_)sobotę(\z|\W|_)`:               "${1}saturday${2}",
		`(\A|\W|_)s(\z|\W|_)`:                    "${1}second${2}",
		`(\A|\W|_)sek(\z|\W|_)`:                  "${1}second${2}",
		`(\A|\W|_)sekund(\z|\W|_)`:               "${1}second${2}",
		`(\A|\W|_)sekunda(\z|\W|_)`:              "${1}second${2}",
		`(\A|\W|_)sekunde(\z|\W|_)`:              "${1}second${2}",
		`(\A|\W|_)sekundy(\z|\W|_)`:              "${1}second${2}",
		`(\A|\W|_)sekundę(\z|\W|_)`:              "${1}second${2}",
		`(\A|\W|_)wrz(\z|\W|_)`:                  "${1}september${2}",
		`(\A|\W|_)wrzesien(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)wrzesień(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)wrzesnia(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)wrzesniu(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)września(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)wrześniu(\z|\W|_)`:             "${1}september${2}",
		`(\A|\W|_)nd(\z|\W|_)`:                   "${1}sunday${2}",
		`(\A|\W|_)nie(\z|\W|_)`:                  "${1}sunday${2}",
		`(\A|\W|_)niedz(\z|\W|_)`:                "${1}sunday${2}",
		`(\A|\W|_)niedziela(\z|\W|_)`:            "${1}sunday${2}",
		`(\A|\W|_)niedziele(\z|\W|_)`:            "${1}sunday${2}",
		`(\A|\W|_)niedzielę(\z|\W|_)`:            "${1}sunday${2}",
		`(\A|\W|_)cz(\z|\W|_)`:                   "${1}thursday${2}",
		`(\A|\W|_)czw(\z|\W|_)`:                  "${1}thursday${2}",
		`(\A|\W|_)czwartek(\z|\W|_)`:             "${1}thursday${2}",
		`(\A|\W|_)wt(\z|\W|_)`:                   "${1}tuesday${2}",
		`(\A|\W|_)wto(\z|\W|_)`:                  "${1}tuesday${2}",
		`(\A|\W|_)wtorek(\z|\W|_)`:               "${1}tuesday${2}",
		`(\A|\W|_)sr(\z|\W|_)`:                   "${1}wednesday${2}",
		`(\A|\W|_)sro(\z|\W|_)`:                  "${1}wednesday${2}",
		`(\A|\W|_)sroda(\z|\W|_)`:                "${1}wednesday${2}",
		`(\A|\W|_)srode(\z|\W|_)`:                "${1}wednesday${2}",
		`(\A|\W|_)śr(\z|\W|_)`:                   "${1}wednesday${2}",
		`(\A|\W|_)śro(\z|\W|_)`:                  "${1}wednesday${2}",
		`(\A|\W|_)środa(\z|\W|_)`:                "${1}wednesday${2}",
		`(\A|\W|_)środę(\z|\W|_)`:                "${1}wednesday${2}",
		`(\A|\W|_)tydz(\z|\W|_)`:                 "${1}week${2}",
		`(\A|\W|_)tydzien(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)tydzień(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)tygodni(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)tygodnie(\z|\W|_)`:             "${1}week${2}",
		`(\A|\W|_)lat(\z|\W|_)`:                  "${1}year${2}",
		`(\A|\W|_)lata(\z|\W|_)`:                 "${1}year${2}",
		`(\A|\W|_)r(\z|\W|_)`:                    "${1}year${2}",
		`(\A|\W|_)rok(\z|\W|_)`:                  "${1}year${2}",
	},
}
