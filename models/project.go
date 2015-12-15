package models
import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
	"github.com/gogather/com/log"
	"fmt"
	"strconv"
)
type Project struct {
	Id int
	Name string
	IconUrl string
	Author string
	Description string
	Time time.Time
}
func init(){
	orm.RegisterModel(new(Project))
}
func (this *Project)TableName()string{
	return "project"
}
func GetProject(id int, name string)(Project, error){
	o := orm.NewOrm()
	o.Using("default")
	pro := Project{}
	var err error
	if id > 0 {
		pro.Id = id
		err = o.Read(&pro, "id")
	}else if len(name) > 0 {
		pro.Name = name
		err = o.Read(&pro, "name")
	}else{
		err = errors.New("参数错误")
	}
	return pro, err
}

func AddProject(name, icon, author, desc string, createTime time.Time)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	pro := new(Project)
	pro.Name = name
	pro.IconUrl = icon
	pro.Author = author
	pro.Description = desc
	pro.Time = createTime
	return o.Insert(pro)
}
func DeleteProject(id int64)error{
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(&Project{Id:id})
	return err
}
func UpdateProject(id int, name, icon, author, desc string, createTime time.Time)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	if id <= 0 {
		return errors.New("参数错误")
	}
	pro, err := GetProject(id, "")
	if err != nil {
		return err
	}
	log.Pinkln(pro)

	pro.Name = name
	pro.IconUrl = icon
	pro.Description = desc
	pro.Author = author
	_, err = o.Update(&pro, "name", "icon_url", "description")
	return err
}

func ListProject(page, pageNum int)(orm.Params, bool, int, error){
	o := orm.NewOrm()
	o.Using("default")
	sql1 := "select * from project order by time desc limit ?, "+fmt.Sprintf("%s", pageNum)
	sql2 := "select count(*) as number from project"
	var maps []orm.Params
	var maps2 []orm.Params
	num, err := o.Raw(sql1, (page-1)*pageNum).Values(&maps)
	if err != nil {
		return nil, false, 0, err
	}
	_, err = o.Raw(sql2).Values(&maps2)
	if err != nil {
		return nil, false, 0, err
	}
	number, err := strconv.Atoi(maps2[0]["number"].(string))
	var addFlag int
	if number / pageNum == 0 {
		addFlag = 0;
	}else{
		addFlag = 1
	}
	pages := number / pageNum + addFlag

	var flagNextPage bool
	if pages == page {
		flagNextPage = false
	}else{
		flagNextPage = true
	}
	if err != nil || num <= 0{
		return nil, false, pages, err
	}
	return maps, flagNextPage, pages, nil
}
