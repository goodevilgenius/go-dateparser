// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pa_Locale = LocaleData{
	Name:      "pa",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)(\d+) ਦਿਨ ਪਹਿਲਾਂ(\z|\W|_)`):   "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਘੰਟਾ ਪਹਿਲਾਂ(\z|\W|_)`):  "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਘੰਟੇ ਪਹਿਲਾਂ(\z|\W|_)`):  "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਹੀਨਾ ਪਹਿਲਾਂ(\z|\W|_)`): "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਹੀਨੇ ਪਹਿਲਾਂ(\z|\W|_)`): "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਕਿੰਟ ਪਹਿਲਾਂ(\z|\W|_)`): "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਹਫ਼ਤਾ ਪਹਿਲਾਂ(\z|\W|_)`): "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਹਫ਼ਤੇ ਪਹਿਲਾਂ(\z|\W|_)`): "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਾਲ ਪਹਿਲਾਂ(\z|\W|_)`):   "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)ਅੱਜ(\z|\W|_)`):                "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਇਸ ਘੰਟੇ(\z|\W|_)`):            "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਇਸ ਮਿੰਟ(\z|\W|_)`):            "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਇਹ ਮਹੀਨਾ(\z|\W|_)`):           "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਹੁਣ(\z|\W|_)`):                "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਇਹ ਹਫ਼ਤਾ(\z|\W|_)`):           "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਇਹ ਸਾਲ(\z|\W|_)`):             "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਬੀਤਿਆ ਕੱਲ੍ਹ(\z|\W|_)`):        "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਪਿਛਲਾ ਮਹੀਨਾ(\z|\W|_)`):        "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਪਿਛਲਾ ਹਫ਼ਤਾ(\z|\W|_)`):        "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਪਿਛਲਾ ਸਾਲ(\z|\W|_)`):          "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)ਪੂਦੁ(\z|\W|_)`):               "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਪ੍ਰੈ(\z|\W|_)`):              "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਪ੍ਰੈਲ(\z|\W|_)`):             "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਗ(\z|\W|_)`):                 "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਗਸਤ(\z|\W|_)`):               "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)ਦਿਨ(\z|\W|_)`):                "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)ਦਸੰ(\z|\W|_)`):                "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)ਦਸੰਬਰ(\z|\W|_)`):              "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)ਫ਼ਰ(\z|\W|_)`):                "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)ਫ਼ਰਵਰੀ(\z|\W|_)`):             "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)ਸ਼ੁੱਕਰ(\z|\W|_)`):             "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)ਸ਼ੁੱਕਰਵਾਰ(\z|\W|_)`):          "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)ਘੰ(\z|\W|_)`):                 "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)ਘੰਟਾ(\z|\W|_)`):               "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਦਿਨ ਵਿੱਚ(\z|\W|_)`):     "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਦਿਨਾਂ ਵਿੱਚ(\z|\W|_)`):   "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਘੰਟਿਆਂ ਵਿੱਚ(\z|\W|_)`):  "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਘੰਟੇ ਵਿੱਚ(\z|\W|_)`):    "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਿੰਟ ਵਿੱਚ(\z|\W|_)`):    "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਿੰਟਾਂ ਵਿੱਚ(\z|\W|_)`):  "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਹੀਨਿਆਂ ਵਿੱਚ(\z|\W|_)`): "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਮਹੀਨੇ ਵਿੱਚ(\z|\W|_)`):   "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਕਿੰਟ ਵਿੱਚ(\z|\W|_)`):   "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਕਿੰਟਾਂ ਵਿੱਚ(\z|\W|_)`): "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਹਫ਼ਤਿਆਂ ਵਿੱਚ(\z|\W|_)`): "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਹਫ਼ਤੇ ਵਿੱਚ(\z|\W|_)`):   "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਾਲ ਵਿੱਚ(\z|\W|_)`):     "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)(\d+) ਸਾਲਾਂ ਵਿੱਚ(\z|\W|_)`):   "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)ਭਲਕੇ(\z|\W|_)`):               "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਗਲਾ ਮਹੀਨਾ(\z|\W|_)`):         "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਗਲਾ ਹਫ਼ਤਾ(\z|\W|_)`):         "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਗਲਾ ਸਾਲ(\z|\W|_)`):           "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)ਜਨ(\z|\W|_)`):                 "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)ਜਨਵਰੀ(\z|\W|_)`):              "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)ਜੁਲਾ(\z|\W|_)`):               "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)ਜੁਲਾਈ(\z|\W|_)`):              "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)ਜੂਨ(\z|\W|_)`):                "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)ਮਾਰਚ(\z|\W|_)`):               "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)ਮਈ(\z|\W|_)`):                 "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)ਸੋਮ(\z|\W|_)`):                "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)ਸੋਮਵਾਰ(\z|\W|_)`):             "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)ਮਹੀਨਾ(\z|\W|_)`):              "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)ਨਵੰ(\z|\W|_)`):                "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ਨਵੰਬਰ(\z|\W|_)`):              "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਕਤੂ(\z|\W|_)`):               "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)ਅਕਤੂਬਰ(\z|\W|_)`):             "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)ਬਾਦੁ(\z|\W|_)`):               "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)ਸ਼ਨਿੱਚਰ(\z|\W|_)`):            "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)ਸ਼ਨਿੱਚਰਵਾਰ(\z|\W|_)`):         "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)ਸਕਿੰਟ(\z|\W|_)`):              "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)ਸਤੰ(\z|\W|_)`):                "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)ਸਤੰਬਰ(\z|\W|_)`):              "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)ਐਤ(\z|\W|_)`):                 "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)ਐਤਵਾਰ(\z|\W|_)`):              "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)ਵੀਰ(\z|\W|_)`):                "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)ਵੀਰਵਾਰ(\z|\W|_)`):             "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)ਮੰਗਲ(\z|\W|_)`):               "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)ਮੰਗਲਵਾਰ(\z|\W|_)`):            "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)ਬੁੱਧ(\z|\W|_)`):               "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)ਬੁੱਧਵਾਰ(\z|\W|_)`):            "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)ਹਫ਼ਤਾ(\z|\W|_)`):              "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)ਸਾਲ(\z|\W|_)`):                "${1}year${2}",
	},
}
