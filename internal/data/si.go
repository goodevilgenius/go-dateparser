// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var si_Locale = LocaleData{
	Name:      "si",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)දින (\d+)කට පෙර(\z|\W|_)`):    "${1}$2 day ago${3}",
		regexp.MustCompile(`(\A|\W|_)පැය (\d+)කට පෙර(\z|\W|_)`):    "${1}$2 hour ago${3}",
		regexp.MustCompile(`(\A|\W|_)මාස (\d+)කට පෙර(\z|\W|_)`):    "${1}$2 month ago${3}",
		regexp.MustCompile(`(\A|\W|_)තත්පර (\d+)කට පෙර(\z|\W|_)`):  "${1}$2 second ago${3}",
		regexp.MustCompile(`(\A|\W|_)සති (\d+)කට පෙර(\z|\W|_)`):    "${1}$2 week ago${3}",
		regexp.MustCompile(`(\A|\W|_)වසර (\d+)කට පෙර(\z|\W|_)`):    "${1}$2 year ago${3}",
		regexp.MustCompile(`(\A|\W|_)අද(\z|\W|_)`):                 "${1}0 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම පැය(\z|\W|_)`):            "${1}0 hour ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම මිනිත්තුව(\z|\W|_)`):      "${1}0 minute ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම මාස(\z|\W|_)`):            "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම මාසය(\z|\W|_)`):           "${1}0 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)දැන්(\z|\W|_)`):               "${1}0 second ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම සති(\z|\W|_)`):            "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම සතිය(\z|\W|_)`):           "${1}0 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)මෙම වසර(\z|\W|_)`):            "${1}0 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)ඊයේ(\z|\W|_)`):                "${1}1 day ago${2}",
		regexp.MustCompile(`(\A|\W|_)පසුගිය මාස(\z|\W|_)`):         "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)පසුගිය මාසය(\z|\W|_)`):        "${1}1 month ago${2}",
		regexp.MustCompile(`(\A|\W|_)පසුගිය සති(\z|\W|_)`):         "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)පසුගිය සතිය(\z|\W|_)`):        "${1}1 week ago${2}",
		regexp.MustCompile(`(\A|\W|_)පසුගිය වසර(\z|\W|_)`):         "${1}1 year ago${2}",
		regexp.MustCompile(`(\A|\W|_)පෙව(\z|\W|_)`):                "${1}am${2}",
		regexp.MustCompile(`(\A|\W|_)අප්‍රේල්(\z|\W|_)`):           "${1}april${2}",
		regexp.MustCompile(`(\A|\W|_)අගෝ(\z|\W|_)`):                "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)අගෝස්තු(\z|\W|_)`):            "${1}august${2}",
		regexp.MustCompile(`(\A|\W|_)දිනය(\z|\W|_)`):               "${1}day${2}",
		regexp.MustCompile(`(\A|\W|_)දෙසැ(\z|\W|_)`):               "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)දෙසැම්බර්(\z|\W|_)`):          "${1}december${2}",
		regexp.MustCompile(`(\A|\W|_)පෙබ(\z|\W|_)`):                "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)පෙබරවාරි(\z|\W|_)`):           "${1}february${2}",
		regexp.MustCompile(`(\A|\W|_)සිකු(\z|\W|_)`):               "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)සිකුරාදා(\z|\W|_)`):           "${1}friday${2}",
		regexp.MustCompile(`(\A|\W|_)පැ(\z|\W|_)`):                 "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)පැය(\z|\W|_)`):                "${1}hour${2}",
		regexp.MustCompile(`(\A|\W|_)දින (\d+)න්(\z|\W|_)`):        "${1}in $2 day${3}",
		regexp.MustCompile(`(\A|\W|_)පැය (\d+)කින්(\z|\W|_)`):      "${1}in $2 hour${3}",
		regexp.MustCompile(`(\A|\W|_)මිනිත්තු (\d+)කින්(\z|\W|_)`): "${1}in $2 minute${3}",
		regexp.MustCompile(`(\A|\W|_)මාස (\d+)කින්(\z|\W|_)`):      "${1}in $2 month${3}",
		regexp.MustCompile(`(\A|\W|_)තත්පර (\d+)කින්(\z|\W|_)`):    "${1}in $2 second${3}",
		regexp.MustCompile(`(\A|\W|_)සති (\d+)කින්(\z|\W|_)`):      "${1}in $2 week${3}",
		regexp.MustCompile(`(\A|\W|_)වසර (\d+)කින්(\z|\W|_)`):      "${1}in $2 year${3}",
		regexp.MustCompile(`(\A|\W|_)හෙට(\z|\W|_)`):                "${1}in 1 day${2}",
		regexp.MustCompile(`(\A|\W|_)ඊළඟ මාස(\z|\W|_)`):            "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)ඊළඟ මාසය(\z|\W|_)`):           "${1}in 1 month${2}",
		regexp.MustCompile(`(\A|\W|_)ඊළඟ සති(\z|\W|_)`):            "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)ඊළඟ සතිය(\z|\W|_)`):           "${1}in 1 week${2}",
		regexp.MustCompile(`(\A|\W|_)ඊළඟ වසර(\z|\W|_)`):            "${1}in 1 year${2}",
		regexp.MustCompile(`(\A|\W|_)ජන(\z|\W|_)`):                 "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)ජනවාරි(\z|\W|_)`):             "${1}january${2}",
		regexp.MustCompile(`(\A|\W|_)ජූලි(\z|\W|_)`):               "${1}july${2}",
		regexp.MustCompile(`(\A|\W|_)ජූනි(\z|\W|_)`):               "${1}june${2}",
		regexp.MustCompile(`(\A|\W|_)මාර්(\z|\W|_)`):               "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)මාර්තු(\z|\W|_)`):             "${1}march${2}",
		regexp.MustCompile(`(\A|\W|_)මැයි(\z|\W|_)`):               "${1}may${2}",
		regexp.MustCompile(`(\A|\W|_)සඳුදා(\z|\W|_)`):              "${1}monday${2}",
		regexp.MustCompile(`(\A|\W|_)මාස(\z|\W|_)`):                "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)මාසය(\z|\W|_)`):               "${1}month${2}",
		regexp.MustCompile(`(\A|\W|_)නොවැ(\z|\W|_)`):               "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)නොවැම්බර්(\z|\W|_)`):          "${1}november${2}",
		regexp.MustCompile(`(\A|\W|_)ඔක්(\z|\W|_)`):                "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)ඔක්තෝබර්(\z|\W|_)`):           "${1}october${2}",
		regexp.MustCompile(`(\A|\W|_)පව(\z|\W|_)`):                 "${1}pm${2}",
		regexp.MustCompile(`(\A|\W|_)සෙන(\z|\W|_)`):                "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)සෙනසුරාදා(\z|\W|_)`):          "${1}saturday${2}",
		regexp.MustCompile(`(\A|\W|_)ත(\z|\W|_)`):                  "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)තත්(\z|\W|_)`):                "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)තත්පරය(\z|\W|_)`):             "${1}second${2}",
		regexp.MustCompile(`(\A|\W|_)සැප්(\z|\W|_)`):               "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)සැප්තැම්බර්(\z|\W|_)`):        "${1}september${2}",
		regexp.MustCompile(`(\A|\W|_)ඉරිදා(\z|\W|_)`):              "${1}sunday${2}",
		regexp.MustCompile(`(\A|\W|_)බ්‍රහස්(\z|\W|_)`):            "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)බ්‍රහස්පතින්දා(\z|\W|_)`):     "${1}thursday${2}",
		regexp.MustCompile(`(\A|\W|_)අඟහ(\z|\W|_)`):                "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)අඟහරුවාදා(\z|\W|_)`):          "${1}tuesday${2}",
		regexp.MustCompile(`(\A|\W|_)බදාදා(\z|\W|_)`):              "${1}wednesday${2}",
		regexp.MustCompile(`(\A|\W|_)සති(\z|\W|_)`):                "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)සතිය(\z|\W|_)`):               "${1}week${2}",
		regexp.MustCompile(`(\A|\W|_)වර්(\z|\W|_)`):                "${1}year${2}",
		regexp.MustCompile(`(\A|\W|_)වර්ෂය(\z|\W|_)`):              "${1}year${2}",
	},
}
