package etcd_crypt

import (
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"io"
	"os"

	"github.com/bketelsen/crypt/config"
)

func init() {
	announcer.Register("etcd", &EtcdDriver{})
}

type EtcdDriver struct{}

func (d *EtcdDriver) Create(c *driver.Config) (driver.IManager, error){
	return &EtcdManager{
		secretKeyring: c.SecretKeyring,
		machines: c.Eps,
	}, nil
}

//etcd 配置发布者
type EtcdManager struct {
	secretKeyring string
	machines []string
}

func (t *EtcdManager)Get(filename string) (map[string][]byte, error){
	kr, err := os.Open(t.secretKeyring)
	if err != nil {
		return nil, err
	}
	defer kr.Close()
	cm, err := config.NewEtcdConfigManager(t.machines, kr)
	if err != nil {
		return nil, err
	}
	value, err := cm.Get(filename)
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]byte)
	ret[filename] = value
	return ret, nil
}

func (t *EtcdManager)Add(filename string, rdc io.Reader) error{
	/*kr, err := os.Open(t.secretKeyring)
	if err != nil {
		return nil, err
	}
	defer kr.Close()
	cm, err := config.NewEtcdConfigManager(t.machines)
	if err != nil {
		return nil, err
	}
	value, err := cm.Watch()filename)
	if err != nil {
		return nil, err
	}
	return value, nil*/
	return nil
}