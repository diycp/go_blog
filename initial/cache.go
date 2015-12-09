package initial
import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"github.com/gogather/com/log"
	"github.com/go-errors/errors"
	"bytes"
	"encoding/gob"
)

var cc cache.Cache
var err error
func InitCache(){
	cacheConfig := beego.AppConfig.String("cache")
	cc = nil

	if "redis" == cacheConfig {
		initRedis()
	}else{
		initMemcache()
	}

	log.Greenln("[cache] use ", cacheConfig)
}

func initMemcache(){
	cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)
	if err != nil {
		beego.Info(err)
	}
}

func initRedis(){
	defer func(){
		if r:= recover(); r != nil {
			log.Redf("initial redis error caught: %v\n", r)
			cc = nil
		}
	}()
	cc, err = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redis_host")+`"}`)
	if err != nil {
		log.Redln(err)
	}
}

func SetCache(key string, value interface{}, timeout int64)error{
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func(){
		if r:= recover(); r != nil {
			log.Redf("set cache error caught: %v\n", r)
			cc = nil
		}
	}()
	err = cc.Put(key, data, timeout)
	if err != nil {
		log.Warnln("Cache失败, key", key)
		return err
	}else{
		log.Blueln("Cache成功, key", key)
		return nil
	}
}
func GetCache(key string, to interface{})error{
	if cc == nil{
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r!= nil{
			log.Redf("Get Cache error caught: %v\n", r)
			cc = nil
		}
	}()
	data := cc.Get(key)
	if data == nil{
		return errors.New("Cache不存在")
	}
	err = Decode(data.([]byte), to)
	if err != nil {
		log.Warnln("获取Cache失败", key, err)
	}else{
		log.Greenln("获取Cache成功", key)
	}
	return err
}

func DelCache(key string)error{
	if cc == nil{
		return errors.New("cc is nil")
	}
	defer  func (){
		if r := recover(); r != nil{
			log.Redf("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	err := cc.Delete(key)
	if err != nil {
		return errors.New("Cache删除失败")
	}else{
		log.Pinkln("删除Cache成功: "+key)
		return nil
	}
}
func Encode(data interface{})([]byte, error){
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil{
		return nil, err
	}
	return buf.Bytes(), nil
}
func Decode(data []byte, to interface{}) error{
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}