package core

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

// var (
// 	validate = validator.New()
// 	once     sync.Once
// )

func MarshalForFiber(v any) ([]byte, error) {
	if v, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(v)
	}
	return json.Marshal(v)
}

func UnmarshalForFiber(data []byte, v any) error {
	var dest any
	if u, ok := v.(easyjson.Unmarshaler); ok {
		dest = easyjson.Unmarshal(data, u)
	} else {
		dest = json.Unmarshal(data, v)
	}
	return Validator.Struct(dest)
}

// func ValidatePhoneNumber(fl validator.FieldLevel) bool {
// 	phone := fl.Field().String()
// 	if phone == "" {
// 		return true
// 	}
// 	// E.164 格式正则表达式：+86 后面跟着 11 位数字的中国手机号
// 	phoneRegex := `^\\+?[1-9]\\d{1,14}$`

// 	// 使用正则表达式进行匹配
// 	re := regexp.MustCompile(phoneRegex)
// 	return re.MatchString(phone)
// }

// func getValidator() *validator.Validate {
// 	once.Do(func() {
// 		validate = validator.New()
// 		validate.RegisterValidation("phone", ValidatePhoneNumber)
// 	})
// 	return validate
// }
