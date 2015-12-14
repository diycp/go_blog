package models
import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/duguying/blog/utils"
	"fmt"
	"github.com/duguying/blog/models"
	"strconv"
)

type Article struct{
	Id int
	Title string
	Uri string
	Keywords string
	Abstract string
	Content string
	Author string
	Time time.Time
	Count int
	Status int
}

func (this *Article)TableName()string{
	return "article"
}
func init(){
	orm.RegisterModel(new(Article))
}
func AddArticle(title, content, keyword, abstract, author string)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	sql := "insert into article(title, uri, keywords, abstract, content, author)values(?,?,?,?,?,?)"
	res, err := o.Raw(sql, title, strings.Replace(title, "/", "-", -1), keyword, abstract, content, author).Exec()
	if err != nil{
		return 0, err
	}else{
		return res.LastInsertId()
	}
}

func GetArticleById(id int)(Article, error){
	var art Article
	err := utils.GetCache("GetArticle.id."+fmt.Sprintf("%d", id), &art)
	if err != nil {
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Id: id}
		err = o.Read(&art, "id")
		utils.SetCache("GetArticle.id."+fmt.Sprintf("%d", id), &art, 600)
	}
	return art, err
}
func GetArticleByUri(uri string)(Article, error){
	var art Article
	err := utils.GetCache("GetArticleByUri.uri."+uri, &art)
	if err == nil {
		count, err := GetArticleViewCount(art.Id)
		if err == nil {
			art.Count = int(count)
		}
		return art, nil
	}else{
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Uri:uri}
		err = o.Read(&art, "uri")
		utils.SetCache("GetArticleByUri.uri."+uri, &art, 600)
	}
	return art, err
}
func GetArticleByTitle(title string)(Article, error){
	var art Article
	err := utils.GetCache("GetArticleByTitle.title."+title, &art)
	if err != nil{
		count, err := GetArticleViewCount(art.Id)
		if err == nil {
			art.Content = int(count)
		}
		return art, nil
	}else{
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Title:title}
		err = o.Read(&art, "title")
		utils.SetCache("GetArticleByTitle.title"+title, art, 600)
	}
}

//获取浏览量
func GetArticleViewCount(id int)(int, error){
	var maps []orm.Params

	sql := `select count from article where id = ?`
	o := orm.NewOrm()
	num, err := o.Raw(sql, id).Values(&maps)
	if err == nil && num > 0 {
		count := maps[0]["count"].(string)
		return strconv.Atoi(count)
	}else{
		return 0, err
	}
}