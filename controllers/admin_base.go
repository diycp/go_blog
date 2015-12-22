package controllers
import (
	"github.com/astaxie/beego"
	"github.com/gogather/com/log"
	"strings"
	"github.com/gogather/com"
	"blog/utils"
	"blog/models"
)
type AdminBaseController struct{
	beego.Controller
}
func (this *AdminBaseController)Propare(){
	user := this.GetSession("username")
	if user == nil {
		this.Redirect("/login", 302)
	}else{
		username := user.(string)
		u, err := models.FindUser(username)
		if err != nil {
			log.Warnln(err)
		}else{
			userLog := &models.UserLog{}
			ipPort := this.Ctx.Request.Header.Get("X-Forwarded-For")
			ipPortArr := strings.Split(ipPort, ":")
			ip := ipPortArr[0]

			location := ""
			userLogIp, err := userLog.GetUserLogByIp(ip)
			if err == nil {
				locationData, err := com.JsonDecode(userLog.Location)
				if err == nil {
					locationJson := locationData.(map[string]interface{})
					if userLog.IsValidLocation(locationJson){
						location = userLogIp.Location
					}else{
						location, _ = utils.GetLocation(ip)
					}
				}else{
					location, _ = utils.GetLocation(ip)
				}
			}else{
				location, _ = utils.GetLocation(ip)
			}
			ua := this.Ctx.Request.UserAgent()
			_, err = userLog.AddUserLog(int64(u.Id), ip, ua, location, 0)
			if err != nil {
				log.Warnln(err)
			}
		}
	}
	this.Data["isAdmin"] = true
	this.Data["isDev"] = beego.AppConfig.String("RunModel") == "dev"
}