package setting

import (
	"github.com/astaxie/beego/validation"
	"fmt"
)

// messageTmpls store commond validate template
var messageTmpls = map[string]string{
	"Required":     "必填",
	"Min":          "最小值是 %d",
	"Max":          "最大值是 %d",
	"Range":        "范围在 %d 到 %d 之间",
	"MinSize":      "最小字符个数为 %d",
	"MaxSize":      "最大字符个数为 is %d",
	"Length":       "要求字符个数为 %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "必须是数字字符",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "必须是有效的邮箱地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "必须是有效的手机号",
	"Tel":          "必须是有效的固定电话号",
	"Phone":        "必须是有效的固定电话或手机号",
	"ZipCode":      "Must be valid zipcode",
}

//验证提示语重载
func validate() {
	validation.SetDefaultMessage(messageTmpls)
}


//验证失败提示输出
func ValidError(errs []*validation.Error) error {
	for _, err := range errs {
		return fmt.Errorf("%s%s", err.Key, err.Message)
	}
	return nil
}
