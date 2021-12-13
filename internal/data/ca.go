// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ca_Locale = LocaleData{
	Name:         "ca",
	DateOrder:    "DMY",
	SkipWords:    []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "de", "del", "i", "l'", "|"},
	PertainWords: []string{"de", "del"},
	Simplifications: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)un(\z|\W|_)`):  "${1}1${2}",
		regexp.MustCompile(`(\A|\W|_)una(\z|\W|_)`): "${1}1${2}",
	},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)fa (\d+) dia(\z|\W|_)`):            "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) dies(\z|\W|_)`):           "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) h(\z|\W|_)`):              "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) hora(\z|\W|_)`):           "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) hores(\z|\W|_)`):          "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) mes(\z|\W|_)`):            "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) mesos(\z|\W|_)`):          "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) s(\z|\W|_)`):              "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) segon(\z|\W|_)`):          "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) segons(\z|\W|_)`):         "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) setm(\z|\W|_)`):           "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) setmana(\z|\W|_)`):        "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) setmanes(\z|\W|_)`):       "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) any(\z|\W|_)`):            "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)fa (\d+) anys(\z|\W|_)`):           "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)avui(\z|\W|_)`):                    "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)hui(\z|\W|_)`):                     "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)aquesta hora(\z|\W|_)`):            "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)aquest minut(\z|\W|_)`):            "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)aquest mes(\z|\W|_)`):              "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)ara(\z|\W|_)`):                     "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)aquesta setm(\z|\W|_)`):            "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)aquesta setmana(\z|\W|_)`):         "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)enguany(\z|\W|_)`):                 "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)ahir(\z|\W|_)`):                    "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)el mes passat(\z|\W|_)`):           "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)mes passat(\z|\W|_)`):              "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)la setm passada(\z|\W|_)`):         "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)la setmana passada(\z|\W|_)`):      "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)setm passada(\z|\W|_)`):            "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)l'any passat(\z|\W|_)`):            "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)abans-d’ahir(\z|\W|_)`):            "${1}2 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)dellà-ahir(\z|\W|_)`):              "${1}2 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)despús-ahir(\z|\W|_)`):             "${1}2 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`):                      "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)abr(\z|\W|_)`):                     "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)abril(\z|\W|_)`):                   "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)d'abr(\z|\W|_)`):                   "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)d'abril(\z|\W|_)`):                 "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)ag(\z|\W|_)`):                      "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)agost(\z|\W|_)`):                   "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)d'ag(\z|\W|_)`):                    "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)d'agost(\z|\W|_)`):                 "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)dia(\z|\W|_)`):                     "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)de des(\z|\W|_)`):                  "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)de desembre(\z|\W|_)`):             "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)des(\z|\W|_)`):                     "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)desembre(\z|\W|_)`):                "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)de febr(\z|\W|_)`):                 "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)de febrer(\z|\W|_)`):               "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)febr(\z|\W|_)`):                    "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)febrer(\z|\W|_)`):                  "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)divendres(\z|\W|_)`):               "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)dv(\z|\W|_)`):                      "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`):                       "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)hora(\z|\W|_)`):                    "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) dia(\z|\W|_)`):      "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) dies(\z|\W|_)`):     "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) h(\z|\W|_)`):        "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) hora(\z|\W|_)`):     "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) hores(\z|\W|_)`):    "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)d‘aquí a (\d+) h(\z|\W|_)`):        "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) min(\z|\W|_)`):      "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) minut(\z|\W|_)`):    "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) minuts(\z|\W|_)`):   "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) mes(\z|\W|_)`):      "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) mesos(\z|\W|_)`):    "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) s(\z|\W|_)`):        "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) segon(\z|\W|_)`):    "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) segons(\z|\W|_)`):   "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) setm(\z|\W|_)`):     "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) setmana(\z|\W|_)`):  "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) setmanes(\z|\W|_)`): "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) any(\z|\W|_)`):      "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)d'aquí a (\d+) anys(\z|\W|_)`):     "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)demà(\z|\W|_)`):                    "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)el mes que ve(\z|\W|_)`):           "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)mes vinent(\z|\W|_)`):              "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)la propera setmana(\z|\W|_)`):      "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)la pròxima setmana(\z|\W|_)`):      "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)la setm que ve(\z|\W|_)`):          "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)la setmana que ve(\z|\W|_)`):       "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)la setmana vinent(\z|\W|_)`):       "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)setm vinent(\z|\W|_)`):             "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)l'any que ve(\z|\W|_)`):            "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)demà passat(\z|\W|_)`):             "${1}in 2 day${2}",
		regexp.MustCompile(`(\A|\W|_)despús-demà(\z|\W|_)`):             "${1}in 2 day${2}",
		regexp.MustCompile(`(\A|\W|_)endemà(\z|\W|_)`):                  "${1}in 2 day${2}",
		regexp.MustCompile(`(\A|\W|_)passat demà(\z|\W|_)`):             "${1}in 2 day${2}",
		regexp.MustCompile(`(\A|\W|_)sendemà(\z|\W|_)`):                 "${1}in 2 day${2}",
		regexp.MustCompile(`(\A|\W|_)endemà passat(\z|\W|_)`):           "${1}in 3 day${2}",
		regexp.MustCompile(`(\A|\W|_)en(\z|\W|_)`):                      "${1}in${2}",
		regexp.MustCompile(`(\A|\W|_)de gen(\z|\W|_)`):                  "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)de gener(\z|\W|_)`):                "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)gen(\z|\W|_)`):                     "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)gener(\z|\W|_)`):                   "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)de jul(\z|\W|_)`):                  "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)de juliol(\z|\W|_)`):               "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`):                     "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)juliol(\z|\W|_)`):                  "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)de juny(\z|\W|_)`):                 "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)juny(\z|\W|_)`):                    "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)de març(\z|\W|_)`):                 "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)març(\z|\W|_)`):                    "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)de maig(\z|\W|_)`):                 "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)maig(\z|\W|_)`):                    "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)dilluns(\z|\W|_)`):                 "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)dl(\z|\W|_)`):                      "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)mes(\z|\W|_)`):                     "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)de nov(\z|\W|_)`):                  "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)de novembre(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`):                     "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)novembre(\z|\W|_)`):                "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)d'oct(\z|\W|_)`):                   "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)d'octubre(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)oct(\z|\W|_)`):                     "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)octubre(\z|\W|_)`):                 "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`):                      "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)dissabte(\z|\W|_)`):                "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)ds(\z|\W|_)`):                      "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`):                       "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)segon(\z|\W|_)`):                   "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)de set(\z|\W|_)`):                  "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)de setembre(\z|\W|_)`):             "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)set(\z|\W|_)`):                     "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)setembre(\z|\W|_)`):                "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)dg(\z|\W|_)`):                      "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)diumenge(\z|\W|_)`):                "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)dijous(\z|\W|_)`):                  "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)dj(\z|\W|_)`):                      "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)dimarts(\z|\W|_)`):                 "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)dt(\z|\W|_)`):                      "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)dc(\z|\W|_)`):                      "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)dimecres(\z|\W|_)`):                "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)setm(\z|\W|_)`):                    "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)setmana(\z|\W|_)`):                 "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)any(\z|\W|_)`):                     "${1}year${2}",
	},
}

var ca_AD_Locale = LocaleData{
	Name:      "ca-AD",
	Parent:    &ca_Locale,
	DateOrder: "DMY",
}

var ca_FR_Locale = LocaleData{
	Name:      "ca-FR",
	Parent:    &ca_Locale,
	DateOrder: "DMY",
}

var ca_IT_Locale = LocaleData{
	Name:      "ca-IT",
	Parent:    &ca_Locale,
	DateOrder: "DMY",
}
