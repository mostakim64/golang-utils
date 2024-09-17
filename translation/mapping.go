package translation

var translations = map[string]map[string]string{
	"validation_nil": {
		"en":    "must be blank",
		"ja":    "空白である必要があります",
		"zh-tw": "必須是空白",
	},
	"validation_empty": {
		"en":    "must be blank",
		"ja":    "空白である必要があります",
		"zh-tw": "必須是空白",
	},
	"validation_date_invalid": {
		"en":    "must be a valid date",
		"ja":    "有効な日程を指定してください",
		"zh-tw": "必須是有效的日期",
	},
	"validation_date_out_of_range": {
		"en":    "the date is out of range",
		"ja":    "日付が範囲外です",
		"zh-tw": "日期超出範圍",
	},
	"validation_in_invalid": {
		"en":    "must be a valid value",
		"ja":    "有効な値である必要があります",
		"zh-tw": "必須是有效值",
	},
	"validation_is_email": {
		"en":    "must be a valid email address",
		"ja":    "有効なEメール アドレスである必要があります",
		"zh-tw": "必須是有效的電子郵件",
	},
	"validation_is_url": {
		"en":    "must be a valid URL",
		"ja":    "有効な URL である必要があります",
		"zh-tw": "必須是有效的 URL",
	},
	"validation_is_digit": {
		"en":    "must contain digits only",
		"ja":    "数字の入力ください",
		"zh-tw": "僅包含數字",
	},
	"validation_is_alphanumeric": {
		"en":    "must contain English letters and digits only",
		"ja":    "英数字のみの入力ください",
		"zh-tw": "僅包含英文字母與數字",
	},
	"validation_is_latitude": {
		"en":    "must be a valid latitude",
		"ja":    "有効な緯度である必要があります",
		"zh-tw": "必須是有效的緯度",
	},
	"validation_is_longitude": {
		"en":    "must be a valid longitude",
		"ja":    "有効な軽度である必要があります",
		"zh-tw": "必須是有效的經度",
	},
	"validation_is_ssn": {
		"en":    "must be a valid social security number",
		"ja":    "有効なマインナンバーである必要があります",
		"zh-tw": "必須是有效的社會安全號碼",
	},
	"validation_is_semver": {
		"en":    "must be a valid semantic version",
		"ja":    "有効なセマンティック バージョンである必要があります",
		"zh-tw": "必須是有效的版本",
	},
	"validation_length_too_long": {
		"en":    "the length must be no more than {{.max}}",
		"ja":    "長さは {{.max}} 以下である必要があります",
		"zh-tw": "長度不得超過 {{.max}}",
	},
	"validation_length_too_short": {
		"en":    "the length must be no less than {{.min}}",
		"ja":    "長さは {{.min}} 以上である必要があります",
		"zh-tw": "長度必須不少於{{.min}}",
	},
	"validation_length_invalid": {
		"en":    "the length must be exactly {{.min}}",
		"ja":    "長さは正確に {{.min}} でなければなりません",
		"zh-tw": "長度必須剛好是 {{.min}}",
	},
	"validation_length_out_of_range": {
		"en":    "the length must be between {{.min}} and {{.max}}",
		"ja":    "長さは {{.min}} から {{.max}} の間である必要があります",
		"zh-tw": "長度必須介於 {{.min}} 和 {{.max}} 之間",
	},
	"validation_length_empty_required": {
		"en":    "the value must be empty",
		"ja":    "値は空でなければなりません",
		"zh-tw": "該值必須空白",
	},
	"validation_key_wrong_type": {
		"en":    "key not the correct type",
		"ja":    "キーの種類が正しくありません",
		"zh-tw": "密鑰類型不正確",
	},
	"validation_key_missing": {
		"en":    "required key is missing",
		"ja":    "必要なキーがありません",
		"zh-tw": "缺少所需的密鑰",
	},
	"validation_key_unexpected": {
		"en":    "key not expected",
		"ja":    "キーが予期されていません",
		"zh-tw": "錯誤密鑰",
	},
	"validation_match_invalid": {
		"en":    "must be in a valid format",
		"ja":    "有効な形式である必要があります",
		"zh-tw": "必須採用有效的格式",
	},
	"validation_min_greater_equal_than_required": {
		"en":    "must be no less than {{.threshold}}",
		"ja":    "{{.threshold}} 以上である必要があります",
		"zh-tw": "必須不小於 {{.threshold}}",
	},
	"validation_max_less_equal_than_required": {
		"en":    "must be no greater than {{.threshold}}",
		"ja":    "{{.threshold}} 以下である必要があります",
		"zh-tw": "不得大於 {{.threshold}}",
	},
	"validation_min_greater_than_required": {
		"en":    "must be greater than {{.threshold}}",
		"ja":    "{{.threshold}} より大きくなければなりません",
		"zh-tw": "必須大於 {{.threshold}}",
	},
	"validation_max_less_than_required": {
		"en":    "must be less than {{.threshold}}",
		"ja":    "{{.threshold}} 未満である必要があります",
		"zh-tw": "必須小於 {{.threshold}}",
	},
	"validation_multiple_of_invalid": {
		"en":    "must be multiple of {{.base}}",
		"ja":    "{{.base}} の倍数でなければなりません",
		"zh-tw": "必須是 {{.base}} 的倍數",
	},
	"validation_not_in_invalid": {
		"en":    "must not be in list",
		"ja":    "リストには適しません",
		"zh-tw": "不得在列表中",
	},
	"validation_not_nil_required": {
		"en":    "is required",
		"ja":    "が必要です",
		"zh-tw": "必須的",
	},
	"validation_required": {
		"en":    "cannot be blank",
		"ja":    "空白にすることはできません",
		"zh-tw": "不能為空白",
	},
	"validation_nil_or_not_empty_required": {
		"en":    "cannot be blank",
		"ja":    "空白にすることはできません",
		"zh-tw": "不能為空白",
	},
}
