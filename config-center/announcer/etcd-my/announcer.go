package etcd_my

import (
	"context"
	"fmt"
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"go.etcd.io/etcd/clientv3"
	"io"
	"io/ioutil"
	"log"
	"time"
)

var (
	defaultDialTimeout = time.Second * 20
	defaultRequestTimeout = time.Second * 20
)

func init() {
	announcer.Register("etcd", &EtcdDriver{})
}

type EtcdDriver struct{}

func (d *EtcdDriver) Create(c *driver.Config) (driver.IManager, error){
	return &EtcdManager{
		dialTimeout: defaultDialTimeout,
		requestTimeout: defaultRequestTimeout,
		secretKeyring: c.SecretKeyring,
		machines: c.Eps,
	}, nil
}

type EtcdManager struct {
	dialTimeout time.Duration
	requestTimeout time.Duration

	secretKeyring string
	machines []string
}

func (t *EtcdManager) Get(ver driver.IVersion, filename string) (map[string][]byte, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, filename)
	cancel()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]byte)
	for _, ev := range resp.Kvs {
		ret[string(ev.Key)] = ev.Value
		//fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return ret, nil
}

func (t *EtcdManager) SSS () {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, "key", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	// Output:
	// key_2 : value
	// key_1 : value
	// key_0 : value
}

func (t *EtcdManager) SSS2 () {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	compRev := resp.Header.Revision // specify compact revision of your choice

	ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.Compact(ctx, compRev)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}

//服务名称，json字符串
func (t *EtcdManager) Add(ver driver.IVersion, filename string, rd io.Reader) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	b, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.Put(ctx, filename+ver.Flag(), string(b))
	cancel()
	return err
}