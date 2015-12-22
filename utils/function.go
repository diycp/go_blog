package utils
import (
	"github.com/gogather/com"
	"strings"
	"os"
	"io/ioutil"
	"fmt"
	"time"
)
/**
 * 检查用户名
 * 首位不能是数字,所有位置只能是数字字母下划线
 */
func CheckUsername(username string) bool {
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}
	for i := 0; i < len(username); i++{
		if !(username[i] == '_' ||
			(username[i] >= '0' && username[i] <= '9') ||
 			(username[i] >= 'a' && username[i] <= 'z') ||
			(username[i] >= 'A' && username[i] <= 'Z')) {
			return false
		}
	}
	return true
}
/**
 * 获取用户头像
 */
func GetGravatar(email string) string {
	return "http://www.gravatar.com/avatar/" + com.Md5(strings.ToUpper(email))
}

func ReadFileByte(path string)([]byte, error){
	fi, ok := os.Open(path)
	if ok != nil {
		panic(ok)
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

func TagSplit(keywords string)string{
	if "" == keywords {
		return ""
	}
	content := ""
	tags := strings.Split(keywords, ",")
	for _, v := range tags {
		content = content + fmt.Sprintf(`<a class="tags" href="/tag/%s/1">%s</a>,`, v, v)
	}
	return content
}

func WriteFile(path, str string) error {
	data := []byte(str)
	return ioutil.WriteFile(path, data, 0644)
}
func GetDate(dateStr string) string {
	date, ok := time.Parse("2006-01-02 15:04:05", dateStr)
	if ok != nil {
		return dateStr
	}else{
		return date.Format("2006-01-02")
	}
}
func GetDateCN(dateStr string)string{
	date, ok := time.Parse("2006-01-02 15:04:05", dateStr)
	if ok != nil {
		return dateStr
	}else{
		return date.Format("2006年01月02日")
	}
}