package regexp

import (
	"regexp"
	"testing"
)

func Test_RegexpPhone(t *testing.T) {
	phone := "+8615763946517"
	reg1 := regexp.MustCompile("^+86|86|[0-9]*$")

	if reg1 == nil { //解释失败，返回nil
		t.Log("regexp err")
		return
	}

	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(phone, -1)
	t.Log("result1 = ", result1[len(result1) - 1])
}