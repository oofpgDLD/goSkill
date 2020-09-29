package json

import (
	"encoding/json"
	"testing"
)

type People struct {
	Name string `json:"name"`
}

func Test_UnmarshalArray(t *testing.T) {
	str := ``
	//ret := make([]uint64, 0)
	var ret []uint64
	err := json.Unmarshal([]byte(str), &ret)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success result:", ret)
}