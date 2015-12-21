package models
import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"regexp"
	"github.com/go-errors/errors"
)

type Users struct {
	Id int
	Username string
	Password string
	Salt string
	Email string
}
type Varify struct{
	Id int
	Username string
	Code string
	Overdue time.Time
}
func (this *Users)TableName()string{
	return "users"
}
func init(){
	orm.RegisterModel(new(Users))
}
func AddUser(username, password string)(int64, error){
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = username
	user.Salt = com.RandString(10)
	user.Password = com.Md5(password + user.Salt)
	return o.Insert(user)
}
func FindUser(username string)(Users, error){
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username:username}
	err := o.Read(&user, "username")
	return user, err
}
func ChangeUsername(oldUsername, newUsername string)error{
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("users").Filter("username", oldUsername).Update(orm.Params{"username": newUsername})
	return err
}
func ChangeEmail(username, email string)error{
	o := orm.NewOrm()
	o.Using("default")
	reg := regexp.MustCompile(`^(\w)+(\.\w)*@(\w)+((\.\w+)+)$`)
	result := reg.MatchString(email)
	if !result {
		return errors.New("不是合法的邮箱地址")
	}
	num, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{"email": email})
	if err != nil {
		return err
	}else if num == 0{
		return errors.New("没有更新")
	}else{
		return nil
	}
}
func AddVerify(username, code string, overdue time.Time)error{
	o := orm.NewOrm()
	o.Using("default")
	overdueTime := overdue.Add(1 * time.Hour).Format("2006-01-02 15:04:05")
	_, err := o.Raw("insert into varify (`username`, `code`, `overdue`) VALUE('"+username+"', '"+code+"', '"+overdueTime+"')").Exec()
	return err
}
func CheckVarify(code string)(bool, string, error){
	o := orm.NewOrm()
	o.Using("default")

	var VarifyItem Varify
	err := o.Raw("select * from varify where code = '"+code+"' AND overdue > now()").QueryRow()
	if code == VarifyItem.Code {
		o.Raw("delete from varify where code = '"+code+"'").Exec()
		return true, VarifyItem.Username, err
	}else{
		return false, VarifyItem.Username, err
	}
}
func SetPassword(username, password string)error{
	o := orm.NewOrm()
	o.Using("default")
	salt := com.RandString(10)
	num, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{
		"salt": salt,
		"password": com.Md5(password+salt),
	})
	if num == 0 {
		return errors.New("没有修改")
	}
	return err
}
func ChangePassword(username, oldPassword, newPassword string)error{
	o := orm.NewOrm()
	o.Using("default")
	salt := com.RandString(10)
	user := Users{Username: username}
	err := o.Read(&user, "username")
	if err != nil{
		return err
	}else{
		if user.Password != com.Md5(oldPassword+user.Salt){
			return errors.New("密码错误")
		}
		_, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{
			"salt": salt,
			"password": com.Md5(newPassword+salt),
		})
		return err
	}
}