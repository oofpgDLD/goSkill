package etcd_my

import (
	"bytes"
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"io/ioutil"
	"testing"
)


func Test_Set(t *testing.T) {
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	manager, err := d.Create(c)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		t.Error("read file err:", err.Error())
		return
	}

	err = manager.Add("/config", bytes.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}

func Test_Get(t *testing.T) {
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	manager, err := d.Create(c)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := manager.Get("foo")
	if err != nil {
		t.Error(err)
		return
	}

	for k,v := range ret{
		t.Log(k, "-----", string(v))
	}
	t.Log("success")
}