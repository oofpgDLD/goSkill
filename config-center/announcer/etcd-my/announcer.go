package etcd_my

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/oofpgDLD/goSkill/config-center/announcer"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"strings"
	"time"

	"github.com/coreos/etcd/auth/authpb"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

var (
	filters = make(map[driver.Case]announcer.Filter)
	defaultDialTimeout = time.Second * 20
	defaultRequestTimeout = time.Second * 20
)

func init() {
	filters[driver.FilterNone] = ClearlySearch
	filters[driver.FilterByName] = ByServerName
	announcer.Register("etcd", &EtcdDriver{})
}

type EtcdDriver struct{}

func (d *EtcdDriver) Manager(c *driver.Config) (driver.IManager, error){
	if c == nil {
		panic("etcd config is nil")
	}
	return NewEManager(c.User, c.Eps, c.SecretKeyring), nil
}

func (d *EtcdDriver) Client(c *driver.Config) (driver.IClient, error){
	if c == nil {
		panic("etcd config is nil")
	}
	return NewEClient(c.User, c.Eps, c.SecretKeyring), nil
}

func NewEManager(u driver.User, machines []string, secretKeyring string) *EManager{
	return &EManager{
		dialTimeout: defaultDialTimeout,
		requestTimeout: defaultRequestTimeout,
		secretKeyring: secretKeyring,
		machines: machines,

		u: u.(*User),
	}
}

//root
type EManager struct {
	dialTimeout time.Duration
	requestTimeout time.Duration

	secretKeyring string
	machines []string

	//user
	u *User
}

func (t *EManager) AuthEnable(enabled bool) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	if enabled {
		_, err = cli.AuthEnable(ctx)
	}else {
		_, err = cli.AuthDisable(ctx)
	}
	cancel()
	return err
}

func (t *EManager) AddUser(name, password string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserAdd(ctx, name, password)
	cancel()
	return err
}

func (t *EManager) DelUser(name string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserDelete(ctx, name)
	cancel()
	return err
}

//
func (t *EManager) GetUser(name string) (driver.User, error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.UserGet(ctx, name)
	cancel()
	if err != nil {
		return nil, err
	}
	u :=  &User{
		name: name,
		roles: make(map[string]*Role),
	}
	for _,roleName := range resp.Roles{
		role, err := t.GetRole(roleName)
		if err != nil {
			return nil, err
		}
		u.roles[roleName] = role.(*Role)
	}
	return u, err
}

func (t *EManager) ListUser() ([]string, error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.UserList(ctx)
	cancel()
	if err != nil {
		return nil, err
	}
	return resp.Users, err
}

func (t *EManager) GrantRole(userName, roleName string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserGrantRole(ctx, userName, roleName)
	cancel()
	return err
}

func (t *EManager) RevokeRole(userName, roleName string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserRevokeRole(ctx, userName, roleName)
	cancel()
	return err
}

//role
func (t *EManager) AddRole(name string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.RoleAdd(ctx, name)
	cancel()
	return err
}

func (t *EManager) DelRole(name string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.RoleDelete(ctx, name)
	cancel()
	return err
}

func (t *EManager) GetRole(name string) (driver.Role, error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.RoleGet(ctx, name)
	cancel()
	if err != nil {
		return nil, err
	}
	r :=  &Role{
		name: name,
		perm: resp.Perm,
	}
	return r, nil
}

func (t *EManager) ListRole() ([]string, error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.RoleList(ctx)
	cancel()
	if err != nil {
		return nil, err
	}
	return resp.Roles, nil
}

func (t *EManager) GrantPermission(roleName, key string, pType int) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.RoleGrantPermission(
		ctx,
		roleName,   // role name
		key, // key
		clientv3.GetPrefixRangeEnd(key), // range end
		clientv3.PermissionType(clientv3.PermReadWrite),
	)
	cancel()
	return err
}

func (t *EManager) RevokePermission(roleName, key string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.RoleRevokePermission(
		ctx,
		roleName,   // role name
		key, // key
		clientv3.GetPrefixRangeEnd(key), // range end
	)
	cancel()
	return err
}

func (t *EManager) DelAccount(name string) error{
	err := t.DelUser(name)
	if err != nil {
		return err
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	//添加到列表
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.Delete(ctx, storeTag(name).convert(""), clientv3.WithPrefix())
	cancel()
	return err
}

func (t *EManager) Account(name string) ([]string,error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	//添加到列表
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, storeTag(name).convert(""), clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}

	server := make([]string, len(resp.Kvs))
	for i,kv := range resp.Kvs{
		_,srvName := storeTag(kv.Key).parse()
		server[i] = srvName
	}
	return server, nil
}

//服务账户列表
func (t *EManager) AccountList() (map[string]string, error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	//添加到列表
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, driver.UserStorePrefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}

	accounts := make(map[string]string)
	for _,kv := range resp.Kvs{
		username,_ := storeTag(kv.Key).parse()
		accounts[username] = string(kv.Value)
	}
	return accounts, nil
}

//添加服务账户
func (t *EManager) AddAccount(name, password, serverName string) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()


	roleName := driver.RolePrefix + serverName
	serverPath := "/" + serverName
	// ./etcdctl --user=root:admin user add name password
	// Error: etcdserver: user name already exists
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserAdd(ctx, name, password)
	cancel()
	if err != nil {
		if err != rpctypes.ErrUserAlreadyExist{
			return err
		}
	}

	// ./etcdctl --user=root:admin role add role-serverName
	// Error: etcdserver: role name already exists
	ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.RoleAdd(ctx, roleName)
	cancel()
	if err != nil {
		return err
	}

	// ./etcdctl --user=root:admin user grant-role name role-serverName
	ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.UserGrantRole(ctx, name, roleName)
	cancel()
	if err != nil {
		return err
	}

	exists := false
	// check if exists delete and update
	ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.RoleGet(ctx, roleName)
	cancel()
	if err != nil {
		return err
	}
	for _,perm := range resp.Perm{
		if perm.PermType == clientv3.PermReadWrite &&
			0 == bytes.Compare(perm.Key, []byte(serverPath)) &&
			0 == bytes.Compare(perm.RangeEnd, []byte(clientv3.GetPrefixRangeEnd(serverPath))){
			exists = true
		}
	}
	if !exists{
		// ./etcdctl --user=root:admin role grant-permission --prefix=true role-serverName /serverName
		ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
		_, err = cli.RoleGrantPermission(
			ctx,
			roleName,
			serverPath,
			clientv3.GetPrefixRangeEnd(serverPath),
			clientv3.PermissionType(clientv3.PermReadWrite),
		)
		cancel()
		if err != nil {
			return err
		}
	}

	//添加到列表
	ctx, cancel = context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.Put(ctx, storeTag(name).convert(serverName), password)
	cancel()
	return err
}

type storeTag string

func (t storeTag) convert(srvName string) string{
	return fmt.Sprintf("%s%s:%s", driver.UserStorePrefix, t, srvName)
}

//username;server name
func (t storeTag) parse() (string, string){
	key := strings.Replace(string(t), driver.UserStorePrefix, "", 1)
	ret := strings.Split(key, ":")
	if len(ret) >= 1 {
		if len(ret) == 1 {
			return ret[0], ""
		}
		return ret[0], ret[1]
	}
	return "",""
}


func NewEClient(u driver.User, machines []string, secretKeyring string) *EClient{
	return &EClient{
		dialTimeout: defaultDialTimeout,
		requestTimeout: defaultRequestTimeout,
		secretKeyring: secretKeyring,
		machines: machines,
		u:u.(*User),
	}
}

type EClient struct {
	dialTimeout time.Duration
	requestTimeout time.Duration

	secretKeyring string
	machines []string

	//user
	u *User
}

//服务名称，json字符串: /serverName/release/configName/version
func (t *EClient) Publish(anc *driver.Announcer) error{
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
		Username: t.u.name,
		Password: t.u.password,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	b, err := ioutil.ReadAll(anc)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	_, err = cli.Put(ctx, anc.ConvertPath(), string(b))
	cancel()
	return err
}

//服务名称，json字符串
func (t *EClient) Get(query *driver.Query) (map[string][]byte, error){
	if filter,ok := filters[query.Filter]; ok&&filter!=nil {
		return filter(t, query)
	}
	return nil, errors.New("etcd filter not support")
}

//filter support
func ClearlySearch(c driver.IClient, item *driver.Query) (map[string][]byte, error){
	if _,ok := c.(*EClient); !ok {
		return nil, errors.New("")
	}
	t := c.(*EClient)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, item.Filename)
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

func ByServerName(c driver.IClient, item *driver.Query) (map[string][]byte, error){
	if _,ok := c.(*EClient); !ok {
		return nil, errors.New("")
	}
	t := c.(*EClient)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.machines,
		DialTimeout: t.dialTimeout,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), t.requestTimeout)
	resp, err := cli.Get(ctx, item.Filename)
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

type User struct {
	name string
	password string
	roles map[string]*Role
}

func (u *User) Name() string {
	return u.name
}

type Role struct {
	name string
	perm []*authpb.Permission
}

func (r *Role) Name() string {
	return r.name
}
/*func (t *EtcdManager) Get(ver driver.IVersion, filename string) (map[string][]byte, error) {
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
}*/