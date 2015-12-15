package models
import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/go-errors/errors"
	"fmt"
)

type File struct{
	Id int
	Filename string
	Path string
	Time time.Time
	Store string
	Mime string
}
func (this *File)TableName(){
	return "file"
}
func init(){
	orm.RegisterModel(new(File))
}
func AddFile(filename, path, store, mime string)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	var file File
	file.Filename = filename
	file.Path = path
	file.Mime = mime
	if store == "local" {
		file.Store = "local"
	}else{
		file.Store = "oss"
	}
	sql := "select count(*) as number from file where path=?"
	var maps []orm.Params
	o.Raw(sql, path).Values(&maps)
	num, _ := strconv.Atoi(maps[0]["number"].(string))
	var err error
	var id int64
	if num == 0 {
		id, err = o.Insert(&file)
	}else{
		id, err = o.Update(&file, "path")
	}
	return id, err
}

func Remove(id int)error{
	if id < 1{
		return errors.New("id is illegal")
	}
	o := orm.NewOrm()
	_, err := o.Delete(&File{Id: id})
	return err
}

func GetFileList(page, pageNum int)([]orm.Params, bool, int, err){
	sql1 := "select * from file order by time desc limit ?, " + fmt.Sprintf("%d", pageNum)
	sql2 := "select count(*) as number from file"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, (page-1)*pageNum).Values(&maps)
	o.Raw(sql2).Values(&maps2)

	number, _ := strconv.Atoi(maps2[0]["number"].(string))

	//总页数
	var addFlag int
	if number % pageNum == 0 {
		addFlag = 0
	}else{
		addFlag = 1
	}
	pages := number / pageNum + addFlag

	//是否存在下页
	var flagNextPage bool
	if pages == page {
		flagNextPage = false
	}else{
		flagNextPage = true
	}

	if err == nil && num > 0{
		return maps, flagNextPage, pages, nil
	}else{
		return nil, false, pages, err
	}
}