package initial
import (
	"github.com/astaxie/beego"
	"strings"
	"net/smtp"
)

/**
 * 发送邮件
 * Param to string a@a.com;a@a.com;a@a.com
 */
func SendMail(to, subject, body string) error{
	user := beego.AppConfig.String("adminemail")
	pass := beego.AppConfig.String("adminemailPass")
	host := beego.AppConfig.String("adminemailhost")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, pass, hp[0])
	contentType := "Content-type: text/html; charset=utf-8"
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}