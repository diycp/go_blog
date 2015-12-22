package index
import (
	"blog/controllers"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"strings"
	"fmt"
	"blog/utils"
)

type InstallController struct {
	controllers.BaseController
}
func (this *InstallController)Get{
	o := orm.NewOrm()
	if com.FileExist("install.lock"){
		this.Abort("404")
	}else{
		sqls := com.ReadFile("etc/blog.sql")
		sqlArr := strings.Split(sqls, ";")
		for index, element := range sqlArr{
			this.Ctx.WriteString(fmt.Sprintf("[%d]", index) + element)
			_, err := o.Raw(element).Exec()
			if err != nil {
				this.Ctx.WriteString("~~ERROR"+element+"\n")
			}
		}
		utils.WriteFile("install.lock", " ")
	}
}