package models
import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

//统计文章总数
func CountArticle()(int, error){
	sql := "select count(*) as number from article"
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)
	return strconv.Atoi(maps[0]["number"].(string))
}

func CountUser()(int, error){
	sql := "select count(*) as number from users"
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)
	return strconv.Atoi(maps[0]["numbers"].(string))
}