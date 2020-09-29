package regexp

import (
	"regexp"
	"testing"
)

//去除手机号前缀+86
func Test_RemovePhonePrefix(t *testing.T) {
	phone := "+8615763946517"
	reg := regexp.MustCompile("^+86|86|[0-9]*$")

	if reg == nil { //解释失败，返回nil
		t.Log("regexp err")
		return
	}

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(phone, -1)
	t.Log("result = ", result[len(result) - 1])
}

//去除十六进制字符串的前缀0x|0X
func Test_RemoveHexPrefix(t *testing.T) {
	src := "123"
	reg, err := regexp.Compile(`^(0x|0X)(\w+)`)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("is mach prefix:", reg.MatchString(src))
	t.Log("result:", reg.ReplaceAllString(src, "${2}"))
}

func Test_GroupNamed(t *testing.T) {
	str := "0xff"
	regMode := `(?P<g1>0x)(?P<g2>\w+)`
	reg, err := regexp.Compile(regMode)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("is mach prefix:", reg.MatchString(str))
	t.Log("result:", reg.ReplaceAllString(str, "${g1}"))
}

func Test_GroupNotCatch(t *testing.T) {
	str := "0xff"
	regMode := `(?:0x)(?P<g2>\w+)`
	reg, err := regexp.Compile(regMode)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("is mach prefix:", reg.MatchString(str))
	t.Log("result:", reg.ReplaceAllString(str, "${g1}"))
}