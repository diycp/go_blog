package initial
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"fmt"
)

func InitSql(){
	user := beego.AppConfig.String("mysql_user")
	password := beego.AppConfig.String("mysql_password")
	host := beego.AppConfig.String("mysql_host")
	port, err := beego.AppConfig.String("mysql_port")
	dbName := beego.AppConfig.String("mysql_db_name")
	if err != err {
		port = 3306
	}

	orm.Debug = true

	if com.FileExist("install.lock") {
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf-8", user, password, host, port, dbName))
	}else{
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf-8", user, password, host, port))
	}
}