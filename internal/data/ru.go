// Code is generated by script; DO NOT EDIT.

package data

var ru_Locale = LocaleData{
	Name:      "ru",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "в", "во", "и", "около", "примерно"},
	Simplifications: map[string]string{
		`(\A|\W|_)((?:через|спустя|в течение)\s+)(день|дня)(\z|\W|_)`: "${1}$21 день${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)года?(\z|\W|_)`:      "${1}$21 год${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)месяца?(\z|\W|_)`:    "${1}$21 месяц${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)минут[уы](\z|\W|_)`:  "${1}$21 минуту${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)недел[юи](\z|\W|_)`:  "${1}$21 неделю${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)секунд[уы](\z|\W|_)`: "${1}$21 секунду${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)сут(ки|ок)(\z|\W|_)`: "${1}$21 сутки${3}",
		`(\A|\W|_)((?:через|спустя|в течение)\s+)часа?(\z|\W|_)`:      "${1}$21 час${3}",
		`(\A|\W|_)([^\d]\s+|^)год(\s+назад)(\z|\W|_)`:                 "${1}$21 год$3${4}",
		`(\A|\W|_)([^\d]\s+|^)день(\s+назад)(\z|\W|_)`:                "${1}$21 день$3${4}",
		`(\A|\W|_)([^\d]\s+|^)месяц(\s+назад)(\z|\W|_)`:               "${1}$21 месяц$3${4}",
		`(\A|\W|_)([^\d]\s+|^)минуту(\s+назад)(\z|\W|_)`:              "${1}$21 минуту$3${4}",
		`(\A|\W|_)([^\d]\s+|^)неделю(\s+назад)(\z|\W|_)`:              "${1}$21 неделю$3${4}",
		`(\A|\W|_)([^\d]\s+|^)секунду(\s+назад)(\z|\W|_)`:             "${1}$21 секунду$3${4}",
		`(\A|\W|_)([^\d]\s+|^)сутки(\s+назад)(\z|\W|_)`:               "${1}$21 сутки$3${4}",
		`(\A|\W|_)([^\d]\s+|^)час(\s+назад)(\z|\W|_)`:                 "${1}$21 час$3${4}",
		`(\A|\W|_)восемь(\z|\W|_)`:                                    "${1}8${2}",
		`(\A|\W|_)дв(а|е|ое|ух)(\z|\W|_)`:                             "${1}2${2}",
		`(\A|\W|_)двадцать(\z|\W|_)`:                                  "${1}20${2}",
		`(\A|\W|_)двенадцать(\z|\W|_)`:                                "${1}12${2}",
		`(\A|\W|_)девять(\z|\W|_)`:                                    "${1}9${2}",
		`(\A|\W|_)десять(\z|\W|_)`:                                    "${1}10${2}",
		`(\A|\W|_)несколько секунд(\z|\W|_)`:                          "${1}44 секунды${2}",
		`(\A|\W|_)од(на|ну|ни|ной|ин)(\z|\W|_)`:                       "${1}1${2}",
		`(\A|\W|_)одиннадцать(\z|\W|_)`:                               "${1}11${2}",
		`(\A|\W|_)пар[ауы](\z|\W|_)`:                                  "${1}2${2}",
		`(\A|\W|_)полгода(\z|\W|_)`:                                   "${1}6 месяцев${2}",
		`(\A|\W|_)полтора года(\z|\W|_)`:                              "${1}18 месяцев${2}",
		`(\A|\W|_)полтора часа(\z|\W|_)`:                              "${1}90 минут${2}",
		`(\A|\W|_)полчаса(\z|\W|_)`:                                   "${1}30 минут${2}",
		`(\A|\W|_)пятнадцать(\z|\W|_)`:                                "${1}15${2}",
		`(\A|\W|_)пять(\z|\W|_)`:                                      "${1}5${2}",
		`(\A|\W|_)пятьдесят(\z|\W|_)`:                                 "${1}50${2}",
		`(\A|\W|_)семь(\z|\W|_)`:                                      "${1}7${2}",
		`(\A|\W|_)сорок(\z|\W|_)`:                                     "${1}40${2}",
		`(\A|\W|_)три(\z|\W|_)`:                                       "${1}3${2}",
		`(\A|\W|_)тридцать(\z|\W|_)`:                                  "${1}30${2}",
		`(\A|\W|_)четыре(\z|\W|_)`:                                    "${1}4${2}",
		`(\A|\W|_)шесть(\z|\W|_)`:                                     "${1}6${2}",
	},
	Translations: map[string]string{
		`(\A|\W|_)(\d+) д назад(\z|\W|_)`:       "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) день назад(\z|\W|_)`:    "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) дн назад(\z|\W|_)`:      "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) дня назад(\z|\W|_)`:     "${1}$2 day ago${3}",
		`(\A|\W|_)(\d+) ч назад(\z|\W|_)`:       "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) час назад(\z|\W|_)`:     "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) часа назад(\z|\W|_)`:    "${1}$2 hour ago${3}",
		`(\A|\W|_)(\d+) мес назад(\z|\W|_)`:     "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) месяц назад(\z|\W|_)`:   "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) месяца назад(\z|\W|_)`:  "${1}$2 month ago${3}",
		`(\A|\W|_)(\d+) сек назад(\z|\W|_)`:     "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) секунду назад(\z|\W|_)`: "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) секунды назад(\z|\W|_)`: "${1}$2 second ago${3}",
		`(\A|\W|_)(\d+) нед назад(\z|\W|_)`:     "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) недели назад(\z|\W|_)`:  "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) неделю назад(\z|\W|_)`:  "${1}$2 week ago${3}",
		`(\A|\W|_)(\d+) г назад(\z|\W|_)`:       "${1}$2 year ago${3}",
		`(\A|\W|_)(\d+) год назад(\z|\W|_)`:     "${1}$2 year ago${3}",
		`(\A|\W|_)(\d+) года назад(\z|\W|_)`:    "${1}$2 year ago${3}",
		`(\A|\W|_)сегодня(\z|\W|_)`:             "${1}0 day ago${2}",
		`(\A|\W|_)в этом часе(\z|\W|_)`:         "${1}0 hour ago${2}",
		`(\A|\W|_)в эту минуту(\z|\W|_)`:        "${1}0 minute ago${2}",
		`(\A|\W|_)в этом месяце(\z|\W|_)`:       "${1}0 month ago${2}",
		`(\A|\W|_)сейчас(\z|\W|_)`:              "${1}0 second ago${2}",
		`(\A|\W|_)на этой неделе(\z|\W|_)`:      "${1}0 week ago${2}",
		`(\A|\W|_)в этом году(\z|\W|_)`:         "${1}0 year ago${2}",
		`(\A|\W|_)вчера(\z|\W|_)`:               "${1}1 day ago${2}",
		`(\A|\W|_)в прошлом месяце(\z|\W|_)`:    "${1}1 month ago${2}",
		`(\A|\W|_)на прошлой неделе(\z|\W|_)`:   "${1}1 week ago${2}",
		`(\A|\W|_)в прошлом году(\z|\W|_)`:      "${1}1 year ago${2}",
		`(\A|\W|_)позавчера(\z|\W|_)`:           "${1}2 day ago${2}",
		`(\A|\W|_)назад(\z|\W|_)`:               "${1}ago${2}",
		`(\A|\W|_)дп(\z|\W|_)`:                  "${1}am${2}",
		`(\A|\W|_)апр(\z|\W|_)`:                 "${1}april${2}",
		`(\A|\W|_)апрель(\z|\W|_)`:              "${1}april${2}",
		`(\A|\W|_)апреля(\z|\W|_)`:              "${1}april${2}",
		`(\A|\W|_)авг(\z|\W|_)`:                 "${1}august${2}",
		`(\A|\W|_)август(\z|\W|_)`:              "${1}august${2}",
		`(\A|\W|_)августа(\z|\W|_)`:             "${1}august${2}",
		`(\A|\W|_)день(\z|\W|_)`:                "${1}day${2}",
		`(\A|\W|_)дн(\z|\W|_)`:                  "${1}day${2}",
		`(\A|\W|_)дней(\z|\W|_)`:                "${1}day${2}",
		`(\A|\W|_)дня(\z|\W|_)`:                 "${1}day${2}",
		`(\A|\W|_)сутки(\z|\W|_)`:               "${1}day${2}",
		`(\A|\W|_)суток(\z|\W|_)`:               "${1}day${2}",
		`(\A|\W|_)дек(\z|\W|_)`:                 "${1}december${2}",
		`(\A|\W|_)декабрь(\z|\W|_)`:             "${1}december${2}",
		`(\A|\W|_)декабря(\z|\W|_)`:             "${1}december${2}",
		`(\A|\W|_)фев(\z|\W|_)`:                 "${1}february${2}",
		`(\A|\W|_)февр(\z|\W|_)`:                "${1}february${2}",
		`(\A|\W|_)февраль(\z|\W|_)`:             "${1}february${2}",
		`(\A|\W|_)февраля(\z|\W|_)`:             "${1}february${2}",
		`(\A|\W|_)пт(\z|\W|_)`:                  "${1}friday${2}",
		`(\A|\W|_)пятница(\z|\W|_)`:             "${1}friday${2}",
		`(\A|\W|_)пятницу(\z|\W|_)`:             "${1}friday${2}",
		`(\A|\W|_)ч(\z|\W|_)`:                   "${1}hour${2}",
		`(\A|\W|_)час(\z|\W|_)`:                 "${1}hour${2}",
		`(\A|\W|_)часа(\z|\W|_)`:                "${1}hour${2}",
		`(\A|\W|_)часов(\z|\W|_)`:               "${1}hour${2}",
		`(\A|\W|_)через (\d+) д(\z|\W|_)`:       "${1}in $2 day${3}",
		`(\A|\W|_)через (\d+) день(\z|\W|_)`:    "${1}in $2 day${3}",
		`(\A|\W|_)через (\d+) дн(\z|\W|_)`:      "${1}in $2 day${3}",
		`(\A|\W|_)через (\d+) дня(\z|\W|_)`:     "${1}in $2 day${3}",
		`(\A|\W|_)через (\d+) ч(\z|\W|_)`:       "${1}in $2 hour${3}",
		`(\A|\W|_)через (\d+) час(\z|\W|_)`:     "${1}in $2 hour${3}",
		`(\A|\W|_)через (\d+) часа(\z|\W|_)`:    "${1}in $2 hour${3}",
		`(\A|\W|_)через (\d+) мин(\z|\W|_)`:     "${1}in $2 minute${3}",
		`(\A|\W|_)через (\d+) минуту(\z|\W|_)`:  "${1}in $2 minute${3}",
		`(\A|\W|_)через (\d+) минуты(\z|\W|_)`:  "${1}in $2 minute${3}",
		`(\A|\W|_)через (\d+) мес(\z|\W|_)`:     "${1}in $2 month${3}",
		`(\A|\W|_)через (\d+) месяц(\z|\W|_)`:   "${1}in $2 month${3}",
		`(\A|\W|_)через (\d+) месяца(\z|\W|_)`:  "${1}in $2 month${3}",
		`(\A|\W|_)через (\d+) сек(\z|\W|_)`:     "${1}in $2 second${3}",
		`(\A|\W|_)через (\d+) секунду(\z|\W|_)`: "${1}in $2 second${3}",
		`(\A|\W|_)через (\d+) секунды(\z|\W|_)`: "${1}in $2 second${3}",
		`(\A|\W|_)через (\d+) нед(\z|\W|_)`:     "${1}in $2 week${3}",
		`(\A|\W|_)через (\d+) недели(\z|\W|_)`:  "${1}in $2 week${3}",
		`(\A|\W|_)через (\d+) неделю(\z|\W|_)`:  "${1}in $2 week${3}",
		`(\A|\W|_)через (\d+) г(\z|\W|_)`:       "${1}in $2 year${3}",
		`(\A|\W|_)через (\d+) год(\z|\W|_)`:     "${1}in $2 year${3}",
		`(\A|\W|_)через (\d+) года(\z|\W|_)`:    "${1}in $2 year${3}",
		`(\A|\W|_)завтра(\z|\W|_)`:              "${1}in 1 day${2}",
		`(\A|\W|_)в следующем месяце(\z|\W|_)`:  "${1}in 1 month${2}",
		`(\A|\W|_)на следующей неделе(\z|\W|_)`: "${1}in 1 week${2}",
		`(\A|\W|_)в следующем году(\z|\W|_)`:    "${1}in 1 year${2}",
		`(\A|\W|_)послезавтра(\z|\W|_)`:         "${1}in 2 day${2}",
		`(\A|\W|_)в течение(\z|\W|_)`:           "${1}in${2}",
		`(\A|\W|_)спустя(\z|\W|_)`:              "${1}in${2}",
		`(\A|\W|_)через(\z|\W|_)`:               "${1}in${2}",
		`(\A|\W|_)янв(\z|\W|_)`:                 "${1}january${2}",
		`(\A|\W|_)январь(\z|\W|_)`:              "${1}january${2}",
		`(\A|\W|_)января(\z|\W|_)`:              "${1}january${2}",
		`(\A|\W|_)июл(\z|\W|_)`:                 "${1}july${2}",
		`(\A|\W|_)июль(\z|\W|_)`:                "${1}july${2}",
		`(\A|\W|_)июля(\z|\W|_)`:                "${1}july${2}",
		`(\A|\W|_)июн(\z|\W|_)`:                 "${1}june${2}",
		`(\A|\W|_)июнь(\z|\W|_)`:                "${1}june${2}",
		`(\A|\W|_)июня(\z|\W|_)`:                "${1}june${2}",
		`(\A|\W|_)мар(\z|\W|_)`:                 "${1}march${2}",
		`(\A|\W|_)март(\z|\W|_)`:                "${1}march${2}",
		`(\A|\W|_)марта(\z|\W|_)`:               "${1}march${2}",
		`(\A|\W|_)май(\z|\W|_)`:                 "${1}may${2}",
		`(\A|\W|_)мая(\z|\W|_)`:                 "${1}may${2}",
		`(\A|\W|_)минут(\z|\W|_)`:               "${1}minute${2}",
		`(\A|\W|_)минуту(\z|\W|_)`:              "${1}minute${2}",
		`(\A|\W|_)минуты(\z|\W|_)`:              "${1}minute${2}",
		`(\A|\W|_)пн(\z|\W|_)`:                  "${1}monday${2}",
		`(\A|\W|_)понедельник(\z|\W|_)`:         "${1}monday${2}",
		`(\A|\W|_)мес(\z|\W|_)`:                 "${1}month${2}",
		`(\A|\W|_)месяц(\z|\W|_)`:               "${1}month${2}",
		`(\A|\W|_)месяца(\z|\W|_)`:              "${1}month${2}",
		`(\A|\W|_)месяцев(\z|\W|_)`:             "${1}month${2}",
		`(\A|\W|_)ноя(\z|\W|_)`:                 "${1}november${2}",
		`(\A|\W|_)нояб(\z|\W|_)`:                "${1}november${2}",
		`(\A|\W|_)ноябрь(\z|\W|_)`:              "${1}november${2}",
		`(\A|\W|_)ноября(\z|\W|_)`:              "${1}november${2}",
		`(\A|\W|_)окт(\z|\W|_)`:                 "${1}october${2}",
		`(\A|\W|_)октябрь(\z|\W|_)`:             "${1}october${2}",
		`(\A|\W|_)октября(\z|\W|_)`:             "${1}october${2}",
		`(\A|\W|_)пп(\z|\W|_)`:                  "${1}pm${2}",
		`(\A|\W|_)сб(\z|\W|_)`:                  "${1}saturday${2}",
		`(\A|\W|_)суббота(\z|\W|_)`:             "${1}saturday${2}",
		`(\A|\W|_)субботу(\z|\W|_)`:             "${1}saturday${2}",
		`(\A|\W|_)с(\z|\W|_)`:                   "${1}second${2}",
		`(\A|\W|_)сек(\z|\W|_)`:                 "${1}second${2}",
		`(\A|\W|_)секунд(\z|\W|_)`:              "${1}second${2}",
		`(\A|\W|_)секунда(\z|\W|_)`:             "${1}second${2}",
		`(\A|\W|_)секунду(\z|\W|_)`:             "${1}second${2}",
		`(\A|\W|_)секунды(\z|\W|_)`:             "${1}second${2}",
		`(\A|\W|_)сен(\z|\W|_)`:                 "${1}september${2}",
		`(\A|\W|_)сент(\z|\W|_)`:                "${1}september${2}",
		`(\A|\W|_)сентябрь(\z|\W|_)`:            "${1}september${2}",
		`(\A|\W|_)сентября(\z|\W|_)`:            "${1}september${2}",
		`(\A|\W|_)воскресение(\z|\W|_)`:         "${1}sunday${2}",
		`(\A|\W|_)воскресенье(\z|\W|_)`:         "${1}sunday${2}",
		`(\A|\W|_)вс(\z|\W|_)`:                  "${1}sunday${2}",
		`(\A|\W|_)четверг(\z|\W|_)`:             "${1}thursday${2}",
		`(\A|\W|_)чт(\z|\W|_)`:                  "${1}thursday${2}",
		`(\A|\W|_)вт(\z|\W|_)`:                  "${1}tuesday${2}",
		`(\A|\W|_)вторник(\z|\W|_)`:             "${1}tuesday${2}",
		`(\A|\W|_)ср(\z|\W|_)`:                  "${1}wednesday${2}",
		`(\A|\W|_)среда(\z|\W|_)`:               "${1}wednesday${2}",
		`(\A|\W|_)среду(\z|\W|_)`:               "${1}wednesday${2}",
		`(\A|\W|_)нед(\z|\W|_)`:                 "${1}week${2}",
		`(\A|\W|_)недели(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)недель(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)неделю(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)неделя(\z|\W|_)`:              "${1}week${2}",
		`(\A|\W|_)г(\z|\W|_)`:                   "${1}year${2}",
		`(\A|\W|_)год(\z|\W|_)`:                 "${1}year${2}",
		`(\A|\W|_)года(\z|\W|_)`:                "${1}year${2}",
		`(\A|\W|_)лет(\z|\W|_)`:                 "${1}year${2}",
	},
}

var ru_BY_Locale = LocaleData{
	Name:      "ru-BY",
	Parent:    &ru_Locale,
	DateOrder: "DMY",
}

var ru_KG_Locale = LocaleData{
	Name:      "ru-KG",
	Parent:    &ru_Locale,
	DateOrder: "DMY",
}

var ru_KZ_Locale = LocaleData{
	Name:      "ru-KZ",
	Parent:    &ru_Locale,
	DateOrder: "DMY",
}

var ru_MD_Locale = LocaleData{
	Name:      "ru-MD",
	Parent:    &ru_Locale,
	DateOrder: "DMY",
}

var ru_UA_Locale = LocaleData{
	Name:      "ru-UA",
	Parent:    &ru_Locale,
	DateOrder: "DMY",
	Translations: map[string]string{
		`(\A|\W|_)am(\z|\W|_)`: "${1}am${2}",
		`(\A|\W|_)pm(\z|\W|_)`: "${1}pm${2}",
	},
}
