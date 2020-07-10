package service

import (
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
)



//添加账户
func (s *Service) AddAccount(username, password string) {
	user := &User{
		username: "root",
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

	err = manager.Account("t1", "123", "srv1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}