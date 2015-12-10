package utils
import (
	"github.com/gogather/iplocation"
	"github.com/astaxie/beego"
	"github.com/go-errors/errors"
	"github.com/gogather/com"
)
var il *iplocation.IpLocation
func init(){
	key := beego.AppConfig.String("iplocation_key")
	il = iplocation.NewIpLocation(key)
}

func GetLocation(ip string)(string, error){
	json, ok := il.Location(ip)
	if json == nil {
		return "", errors.New("json is nil")
	}
	countryName := json["countryName"].(string)
	regionName := json["regionName"].(string)
	cityName := json["cityName"].(string)
	date := map[string]interface{}{
		"countryName": countryName,
		"regionName": regionName,
		"cityName": cityName,
	}
	str, ok := com.JsonEncode(date)
	return str, ok
}