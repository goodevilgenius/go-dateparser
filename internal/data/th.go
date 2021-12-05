// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var thLocale = LocaleData{
	Name:                  "th",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 5,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "น", "เมื่อ", "เวลา", "，"},
	January:               []string{"มกรา", "มกราคม", "มค", "เดือนมกราคม"},
	February:              []string{"กพ", "กุมภา", "กุมภาพันธ์", "เดือนกุมภาพันธ์"},
	March:                 []string{"มีค", "มีนา", "มีนาคม", "เดือนมีนาคม"},
	April:                 []string{"เดือนเมษายน", "เมย", "เมษา", "เมษายน"},
	May:                   []string{"พค", "พฤษภา", "พฤษภาคม", "เดือนพฤษภาคม"},
	June:                  []string{"มิถุนา", "มิถุนายน", "มิย", "เดือนมิถุนายน"},
	July:                  []string{"กค", "กรกฎา", "กรกฎาคม", "เดือนกรกฏาคม"},
	August:                []string{"สค", "สิงหา", "สิงหาคม", "เดือนสิงหาคม"},
	September:             []string{"กย", "กันยา", "กันยายน", "เดือนกันยายน"},
	October:               []string{"ตค", "ตุลา", "ตุลาคม", "เดือนตุลาคม"},
	November:              []string{"พย", "พฤศจิ", "พฤศจิกายน", "เดือนพฤศจิกายน"},
	December:              []string{"ธค", "ธันวา", "ธันวาคม", "เดือนธันวาคม"},
	Monday:                []string{"จ", "จันทร์", "วันจันทร์", "วันจันทร์ที่"},
	Tuesday:               []string{"วันอังคาร", "วันอังคารที่", "อ", "อังคาร"},
	Wednesday:             []string{"พ", "พุธ", "วันพุธ", "วันพุธที่"},
	Thursday:              []string{"พฤ", "พฤหัส", "พฤหัสบดี", "วันพฤหัสบดี", "วันพฤหัสบดีที่"},
	Friday:                []string{"วันศุกร์", "วันศุกร์ที่", "ศ", "ศุกร์"},
	Saturday:              []string{"วันเสาร์", "วันเสาร์ที่", "ส", "เสาร์"},
	Sunday:                []string{"วันอาทิตย์", "วันอาทิตย์ที่", "อา", "อาทิตย์"},
	AM:                    []string{"ก่อนเที่ยง"},
	PM:                    []string{"หลังเที่ยง"},
	Year:                  []string{"ปี"},
	Month:                 []string{"เดือน"},
	Week:                  []string{"สัปดาห์"},
	Day:                   []string{"วัน"},
	Hour:                  []string{"ชม", "ชั่วโมง"},
	Minute:                []string{"น", "นาที"},
	Second:                []string{"วิ", "วินาที"},
	In:                    []string{"ใน"},
	Ago:                   []string{"ก่อน", "มาแล้ว", "แต่ก่อน"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`วันนี้`},
		`0 hour ago`:   {`ชั่วโมงนี้`},
		`0 minute ago`: {`นาทีนี้`},
		`0 month ago`:  {`เดือนนี้`},
		`0 second ago`: {`ขณะนี้`},
		`0 week ago`:   {`สัปดาห์นี้`},
		`0 year ago`:   {`ปีนี้`},
		`1 day ago`:    {`เมื่อวาน`},
		`1 month ago`:  {`เดือนที่แล้ว`},
		`1 week ago`:   {`สัปดาห์ที่แล้ว`},
		`1 year ago`:   {`ปีที่แล้ว`},
		`in 1 day`:     {`พรุ่งนี้`},
		`in 1 month`:   {`เดือนหน้า`},
		`in 1 week`:    {`สัปดาห์หน้า`},
		`in 1 year`:    {`ปีหน้า`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) วันที่ผ่านมา`),
			regexp.MustCompile(`(?i)(\d+) วันที่แล้ว`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ชม ที่แล้ว`),
			regexp.MustCompile(`(?i)(\d+) ชั่วโมงที่ผ่านมา`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) นาทีที่ผ่านมา`),
			regexp.MustCompile(`(?i)(\d+) นาทีที่แล้ว`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) เดือนที่ผ่านมา`),
			regexp.MustCompile(`(?i)(\d+) เดือนที่แล้ว`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) วินาทีที่ผ่านมา`),
			regexp.MustCompile(`(?i)(\d+) วินาทีที่แล้ว`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) สัปดาห์ที่ผ่านมา`),
			regexp.MustCompile(`(?i)(\d+) สัปดาห์ที่แล้ว`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ปีที่แล้ว`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)ใน (\d+) วัน`),
			regexp.MustCompile(`(?i)ในอีก (\d+) วัน`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)ใน (\d+) ชม`),
			regexp.MustCompile(`(?i)ในอีก (\d+) ชั่วโมง`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)ใน (\d+) นาที`),
			regexp.MustCompile(`(?i)ในอีก (\d+) นาที`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)ใน (\d+) เดือน`),
			regexp.MustCompile(`(?i)ในอีก (\d+) เดือน`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)ใน (\d+) วินาที`),
			regexp.MustCompile(`(?i)ในอีก (\d+) วินาที`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)ใน (\d+) สัปดาห์`),
			regexp.MustCompile(`(?i)ในอีก (\d+) สัปดาห์`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)ใน (\d+) ปี`),
			regexp.MustCompile(`(?i)ในอีก (\d+) ปี`),
		},
	},
	Simplifications: map[string]string{
		`1 วันที่แล้ว`:    `1 วัน`,
		`2 วันที่แล้ว`:    `2 วัน`,
		`วันนี้`:          `0 วัน`,
		`เมื่อวานซืน`:     `2 วัน`,
		`เมื่อวานนี้`:     `1 วัน`,
		`เมื่อสักครู่นี้`: `0 วินาที`,
	},
}