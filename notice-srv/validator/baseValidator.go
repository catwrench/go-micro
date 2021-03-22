package validator

// 当前方案：go-playground/validator/v10 (直接引入的gin框架默认验证器)
// 使用示例:https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
// 文档:https://godoc.org/github.com/go-playground/validator
// 自定义规则示例：https://blog.csdn.net/guyan0319/article/details/105918559

// 备用方案：https://github.com/thedevsaddam/govalidator (一个完全基于laravelValidator思想设计的验证器)

// 常用验证规则：
// required：必填
// email：验证字符串是email格式；例：“email”
// url：这将验证字符串值包含有效的网址;例：“url”
// max：字符串最大长度；例：“max=20”
// min:字符串最小长度；例：“min=6”
// excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
// len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
// eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
// ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
// gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
// gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
// lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
// lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”
// oneof：等于指定的几个值之一，数据以空格分割，特殊符号需要单引号包裹；例：“oneof=1 2”、“oneof=apple `small apple`”

// 和其他字段比较：
// eqfield=Field: 必须等于 Field 的值；
// nefield=Field: 必须不等于 Field 的值；
// gtfield=Field: 必须大于 Field 的值；
// gtefield=Field: 必须大于等于 Field 的值；
// ltfield=Field: 必须小于 Field 的值；
// ltefield=Field: 必须小于等于 Field 的值；
// eqcsfield=Other.Field: 必须等于 struct Other 中 Field 的值；
// necsfield=Other.Field: 必须不等于 struct Other 中 Field 的值；
// gtcsfield=Other.Field: 必须大于 struct Other 中 Field 的值；
// gtecsfield=Other.Field: 必须大于等于 struct Other 中 Field 的值；
// ltcsfield=Other.Field: 必须小于 struct Other 中 Field 的值；
// ltecsfield=Other.Field: 必须小于等于 struct Other 中 Field 的值；

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/micro/go-micro/v2/util/log"
	"reflect"
)

//作为翻译字段名的标签
const ValidateLabel = "label"

// 验证器实例,保存了结构体信息
var Builder *validator.Validate

// 翻译器
var trans ut.Translator

func init() {
	//初始化验证实例
	Builder = validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	Builder.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get(ValidateLabel)
		if name == "" {
			return field.Name
		}
		return name
	})
	//注册翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if err := zhTrans.RegisterDefaultTranslations(Builder, trans); err != nil {
		log.Error("注册默认翻译器失败")
	}
}

//翻译所有错误
func TransErrors(err error) []string {
	var errList []string
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			printErrDetail(err)
			errList = append(errList, err.Translate(trans))
		}
	}
	log.Error("字段验证未通过：", errList)
	return errList
}

//翻译第一条错误
func TransError(err error) string {
	var errMsg string

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return ""
		}

		for _, err := range err.(validator.ValidationErrors) {
			printErrDetail(err)
			errMsg = err.Translate(trans)
			break
		}
		log.Error("字段验证未通过：", errMsg)
	}

	return errMsg
}

//打印错误详情
func printErrDetail(err validator.FieldError) {
	fmt.Println("---------validator." + err.Field() + " 字段错误详情---------")
	fmt.Println("StructNamespace: ", err.StructNamespace())
	fmt.Println("StructField: ", err.StructField())
	fmt.Println("Tag: ", err.Tag())
	fmt.Println("ActualTag: ", err.ActualTag())
	fmt.Println("Kind: ", err.Kind())
	fmt.Println("Type: ", err.Type())
	fmt.Println("Value: ", err.Value())
	fmt.Println("Param: ", err.Param())
	fmt.Println()
}
