package _defer

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Defer1(t *testing.T) {
	i := 0
	defer func(a int) {
		t.Log(a)
	}(i)

	i++
	t.Log(i)
}

var _ fmt.Stringer = LevelDefault

const (
	LevelDefault MyInt = iota
	LevelReadUncommitted
)

type MyInt int

func (t MyInt) String() string{
	return strconv.Itoa(int(t))
}

func Test_FFF(t *testing.T) {
	//var item MyInt = 1
	t.Log(LevelDefault.String())
	t.Log(LevelReadUncommitted.String())
}