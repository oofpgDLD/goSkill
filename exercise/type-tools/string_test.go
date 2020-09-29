package type_tools

import (
	"strconv"
	"strings"
	"testing"
)

func TestIntToString(t *testing.T) {
	t.Log(strconv.Itoa(10))
}

func TestInt64ToString(t *testing.T) {
	t.Log(strconv.FormatInt(10, 10))
	t.Log(strconv.FormatInt(10, 16))
	t.Log(strconv.FormatInt(10, 8))
}

func TestFloatToString(t *testing.T) {
	t.Log(strconv.FormatFloat(10.01, 'e', -1, 32))
	t.Log(strconv.FormatFloat(10.01, 'f', -1, 32))
	t.Log(strconv.FormatFloat(7.0e+6, 'f', -1, 32))
	t.Log(strconv.FormatFloat(10000000000000000000000000000000000000000, 'e', -1, 64))
}

func TestStringToInt(t *testing.T) {
	t.Log(strconv.Atoi("10"))
}

func TestStringToInt64(t *testing.T) {
	t.Log(strconv.ParseInt("0b10", 0 , 0))
	t.Log(strconv.ParseInt("010", 0 , 0))
	t.Log(strconv.ParseInt("0x10", 0 , 0))
}

func TestStringToFloat(t *testing.T) {
	t.Log(strconv.ParseFloat("1.001e+01", 32))
	t.Log(strconv.ParseFloat("7000000", 32))
}

func TestBoolToString(t *testing.T) {
	t.Log(strconv.FormatBool(false))
	t.Log(strconv.FormatBool(true))
}

func TestStringToBool(t *testing.T) {
	t.Log(strconv.ParseBool("false"))
	t.Log(strconv.ParseBool("0"))
}

//-----------

func TestStringCount(t *testing.T) {
	t.Log(strings.Count("cheese", "eee"))
}