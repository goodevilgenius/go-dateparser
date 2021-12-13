// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var os_Locale = LocaleData{
	Name:      "os",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)(\d+) бон раздӕр(\z|\W|_)`):   "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) боны размӕ(\z|\W|_)`):   "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) сахаты размӕ(\z|\W|_)`): "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)абон(\z|\W|_)`):               "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`):          "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`):        "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`):         "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`):                "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`):          "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`):          "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)знон(\z|\W|_)`):               "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`):         "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`):          "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`):          "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`):                 "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)ӕмбисбоны размӕ(\z|\W|_)`):    "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)апр(\z|\W|_)`):                "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)апрелы(\z|\W|_)`):             "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)апрель(\z|\W|_)`):             "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)авг(\z|\W|_)`):                "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)август(\z|\W|_)`):             "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)августы(\z|\W|_)`):            "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)бон(\z|\W|_)`):                "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)дек(\z|\W|_)`):                "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)декабры(\z|\W|_)`):            "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)декабрь(\z|\W|_)`):            "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)фев(\z|\W|_)`):                "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)февр(\z|\W|_)`):               "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)февралы(\z|\W|_)`):            "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)февраль(\z|\W|_)`):            "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)майрӕмбон(\z|\W|_)`):          "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)мрб(\z|\W|_)`):                "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)сахат(\z|\W|_)`):              "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)(\d+) боны фӕстӕ(\z|\W|_)`):   "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) сахаты фӕстӕ(\z|\W|_)`): "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)сом(\z|\W|_)`):                "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`):         "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`):          "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`):          "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)янв(\z|\W|_)`):                "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)январы(\z|\W|_)`):             "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)январь(\z|\W|_)`):             "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)июлы(\z|\W|_)`):               "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)июль(\z|\W|_)`):               "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)июны(\z|\W|_)`):               "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)июнь(\z|\W|_)`):               "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)мар(\z|\W|_)`):                "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)март(\z|\W|_)`):               "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)мартъи(\z|\W|_)`):             "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)мартъийы(\z|\W|_)`):           "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)май(\z|\W|_)`):                "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)майы(\z|\W|_)`):               "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)крс(\z|\W|_)`):                "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)къуырисӕр(\z|\W|_)`):          "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)мӕй(\z|\W|_)`):                "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)ноя(\z|\W|_)`):                "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)нояб(\z|\W|_)`):               "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ноябры(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ноябрь(\z|\W|_)`):             "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)окт(\z|\W|_)`):                "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)октябры(\z|\W|_)`):            "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)октябрь(\z|\W|_)`):            "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`):                 "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)ӕмбисбоны фӕстӕ(\z|\W|_)`):    "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)сабат(\z|\W|_)`):              "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)сбт(\z|\W|_)`):                "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)секунд(\z|\W|_)`):             "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)сен(\z|\W|_)`):                "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)сент(\z|\W|_)`):               "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)сентябры(\z|\W|_)`):           "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)сентябрь(\z|\W|_)`):           "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)хуыцаубон(\z|\W|_)`):          "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)хцб(\z|\W|_)`):                "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)цпр(\z|\W|_)`):                "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)цыппӕрӕм(\z|\W|_)`):           "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)дцг(\z|\W|_)`):                "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)дыццӕг(\z|\W|_)`):             "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)ӕрт(\z|\W|_)`):                "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)ӕртыццӕг(\z|\W|_)`):           "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)къуыри(\z|\W|_)`):             "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)аз(\z|\W|_)`):                 "${1}year${2}",
	},
}

var os_RU_Locale = LocaleData{
	Name:      "os-RU",
	Parent:    &os_Locale,
	DateOrder: "DMY",
}
