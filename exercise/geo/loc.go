package geo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	Geo_Loc = "geo_"
)

func NewLocRedis(rp *redis.Pool) *locCaChe {
	return &locCaChe{
		rp: rp,
	}
}

type locCaChe struct {
	rp *redis.Pool
}

func (c *locCaChe) GeoAdd(appId, name string, longitude, latitude string) error{
	key := Geo_Loc + appId
	conn := c.rp.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := conn.Do("GEOADD", key, longitude, latitude, name)
	if err != nil {
		return err
	}
	return err
}

func (c *locCaChe) GeoDIST(appId, name1, name2 string) (string, error){
	key := Geo_Loc + appId
	conn := c.rp.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	reply, err := conn.Do("GEODIST", key, name1, name2)
	if err != nil {
		return "", err
	}
	v, err := redis.String(reply, err)
	if err != nil {
		return "", err
	}

	return v, nil
}

