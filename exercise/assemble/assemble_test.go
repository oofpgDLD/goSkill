package assemble

import "testing"

type super interface {
	head() string
	end() string
}

type absClass struct {
	su super
}

func (t *absClass) Compose() string{
	return t.su.head() + t.su.end()
}

//imp
type imp1 struct {
	ab absClass
	h string
	e string
}

func (t *imp1) head() string{
	return t.h
}

func (t *imp1) end() string{
	return t.e
}

func Test_F(t *testing.T) {
	i1 := imp1{
		ab: absClass{
		},
		h:"hello",
		e:"world",
	}
	t.Log(i1.ab.Compose())
}