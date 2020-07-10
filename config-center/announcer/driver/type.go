package driver

import(
	"errors"
	"fmt"
	"io"
)

const (
	FilterNone Case = iota
	FilterByName
)

const (
	UserStorePrefix = "store-"
	UserServerPrefix = "server-"
	RolePrefix = "role-"
)

type Case uint

type Role interface {
	Name() string
}

type User interface {
	Name() string
}

type File struct {
	Env string
	SrvName string
	Ver string
	Filename string
}

func (t *File) ConvertPath() string{
	return fmt.Sprintf("/%s/%s/%s/%s", t.SrvName,t.Env,t.Filename,t.Ver)
}

type Announcer struct {
	//file
	*File
	rd io.Reader
}

func (t *Announcer) Read(p []byte) (n int, err error){
	if t.rd == nil {
		return 0, errors.New("reader not init")
	}
	return t.rd.Read(p)
}

func (t *Announcer) Load(rd io.Reader){
	t.rd = rd
}

type Query struct {
	*File
	Filter Case
	Client User
}


