package driver

type Config struct {
	User User
	Eps []string
	SecretKeyring string
}

type Driver interface {
	Manager(c *Config) (IManager, error)
	Client(c *Config) (IClient, error)
}

//root 账户负责添加用户和权限
type IManager interface {
	AuthEnable(enabled bool) error

	AddUser(name, password string) error
	DelUser(name string) error
	GetUser(name string) (User, error)
	ListUser() ([]string, error)
	GrantRole(userName, roleName string) error
	RevokeRole(userName, roleName string) error

	AddRole(name string) error
	DelRole(name string) error
	GetRole(name string) (Role, error)
	ListRole() ([]string, error)
	GrantPermission(roleName, key string, permission int) error
	RevokePermission(roleName, key string) error

	//创建服务配置发布账户
	Account(name string) ([]string,error)
	AddAccount(name, password, serverName string) error
	DelAccount(name string) error
	AccountList() (map[string]string, error)
}

type IClient interface {
	//发布配置文件 环境-配置文件名称-服务名称-版本号 如：Realease:config.toml:work:2.0.0
	Publish(f *Announcer) error
	//根据条件获取
	Get(query *Query) (map[string][]byte, error)
}