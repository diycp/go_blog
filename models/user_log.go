package models
import (
	"time"
	"github.com/astaxie/beego/orm"
)

type UserLog struct {
	Id int64
	User int64
	Ip string
	Ua string
	Location string
	Action int
	CreateTime time.Time
}
func init(){
	orm.RegisterModel(new(UserLog))
}
func (this *UserLog)TableName()string{
	return "user_log"
}
func (this *UserLog)AddUserLog(user int64, ip, ua, location string, action int)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	userLog := new(UserLog)
	userLog.User = user
	userLog.Ip = ip
	userLog.Ua = ua
	userLog.Location = location
	userLog.Action = action
	return o.Insert(userLog)
}
func (this *UserLog)GetUserLogByIp(ip string)(UserLog, error){
	o := orm.NewOrm()
	o.Using("default")
	userLog := UserLog{Ip:ip}
	err := o.Read(&userLog, "ip")
	return userLog, err
}
func (this *UserLog)IsVaildLocation(data map[string]interface{})bool{
	o := orm.NewOrm()
	o.Using("default")
	cityName := data["city_name"].(string)
	countryName := data["country_name"].(string)
	regionName := data["region_name"].(string)
	if len(cityName) == 0 && len(countryName) == 0 && len(regionName) == 0{
		return false
	}else{
		return true
	}
}