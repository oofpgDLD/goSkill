package etcd_my

import (
	"bytes"
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"io/ioutil"
	"testing"
)

func Test_AccountList(t *testing.T) {
	user := &User{
		name: "root",
		password: "admin",
	}
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
		User: user,
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	manager, err := d.Manager(c)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := manager.AccountList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", ret)
}

func Test_AddAccount(t *testing.T) {
	user := &User{
		name: "root",
		password: "admin",
	}
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
		User: user,
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	manager, err := d.Manager(c)
	if err != nil {
		t.Error(err)
		return
	}

	err = manager.AddAccount("t1", "456", "srv2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}

func Test_UserGet(t *testing.T) {
	user := &User{
		name: "root",
		password: "admin",
	}
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
		User: user,
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	manager, err := d.Manager(c)
	if err != nil {
		t.Error(err)
		return
	}
	
	u, err := manager.GetUser("dld")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", u.Name())
}

func Test_Set(t *testing.T) {
	user := &User{
		name: "t1",
		password: "123",
	}
	c := &driver.Config{
		Eps: []string{"172.16.103.31:2379"},
		User: user,
	}
	d, err := announcer.Create("etcd")
	if err != nil {
		t.Error(err)
		return
	}
	client, err := d.Client(c)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		t.Error("read file err:", err.Error())
		return
	}

	anc := driver.Announcer{
		File:&driver.File{
			Env: "release",
			SrvName: "srv1",
			Ver: "1.0.0",
			Filename: "testhhh",
		},
	}
	anc.Load(bytes.NewReader(data))

	err = client.Publish(&anc)
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
	manager, err := d.Client(c)
	if err != nil {
		t.Error(err)
		return
	}

	query := driver.Query{
		File: &driver.File{
			Filename: "testhhh",
		},
		Filter: driver.FilterNone,
	}
	ret, err := manager.Get(&query)
	if err != nil {
		t.Error(err)
		return
	}

	for k,v := range ret{
		t.Log(k, "-----", string(v))
	}
	t.Log("success")
}