package utils

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var zhTranslator ut.Translator

func init() {
	universalTranslator := ut.New(zh.New())
	zhTranslator, _ = universalTranslator.GetTranslator("zh")

	// 获取默认的验证引擎。
	validate := binding.Validator.Engine().(*validator.Validate)

	// // 注册一个名为"usernameUnique"的自定义验证函数。
	// err := validate.RegisterValidation("usernameUnique", UsernameUnique)
	// if err != nil {
	// 	panic(err)
	// }

	// 注册默认的中文翻译。
	err := zhTranslations.RegisterDefaultTranslations(validate, zhTranslator)
	if err != nil {
		panic(err)
	}

	// 注册一个函数来获取结构体字段的中文标签。
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		chineseName := field.Tag.Get("chinese")
		if chineseName == "" || chineseName == "-" {
			return ""
		}
		return chineseName
	})

	// // 为"usernameUnique"验证函数注册中文翻译。
	// err = validate.RegisterTranslation("usernameUnique", zhTranslator, func(ut ut.Translator) error {
	// 	return ut.Add("usernameUnique", "{0} {1} 已存在", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("usernameUnique", fe.Field(), fe.Value().(string))
	// 	return t
	// })
	// if err != nil {
	// 	panic(err)
	// }
}

func Translate(err error) string {
	errors := err.(validator.ValidationErrors)

	msg := ""
	// 遍历每个验证错误。
	for _, verr := range errors {
		// 使用 zhTranslator 翻译每个错误信息，并将其累加到 msg 中，每个错误信息后跟一个换行符。
		msg += verr.Translate(zhTranslator) + "\n"
	}
	// 返回累积的所有错误信息。
	return msg
}
