// 常用的函数封装,来自于github.com/gogather/com/function.go

package tools

import (
	"io"
	"crypto/rand"
	"encoding/base64"
	"runtime"
	"reflect"
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"math"
	"net/http"
)

//创建GUID
func CreateGUID() string{
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//截取字符串
func SubString(str string, start, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if start < 0 {
		start = 0
	}
	if start >= lth {
		start = lth
	}
	end := start + length
	if end > lth {
		end = lth
	}
	return string(rs[start : end])
}

//创建随机字符串
func RandString(n int) string {
	guid := CreateGUID()
	return SubString(guid, 0, n)
}

//获取当前函数名
func GetCurrentFuncName(i interface{})string{
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//检查用户名是否合法
func CheckUsername(username string)bool {
	//数字不能开头
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}
	//用户名只能是数字\下划线\大小写字母
	for i := 0; i< len(username); i++ {
		if !(username[i] == '_' ||
		(username[i] >= 'A' && username[i] <= 'Z') ||
		(username[i] >= 'a' && username[i] <= 'z') ||
		username[i] >= '0' || username[i] <= '9'){
			return false
		}
	}
	return true
}

//生成MD5串
func Md5(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

//获取头像
func GetGravatar(email string)string{
	return "http://www.gravatar.com/avatar/" + Md5(strings.ToUpper(email))
}

//文件或目录是否存在
func PathExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func FileExist(filename string)bool{
	return PathExist(filename)
}

//读取文件
func ReadFileByte(filename string)([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

//读取文本文件
func ReadFileString(filename string)string{
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	return string(content)
}

//字符串写入文件
func WriteFile(filename, data string) error {
	str := []byte(data)
	return ioutil.WriteFile(filename, str, 0644)
}

//创建目录
func MkDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

//JsonEncode
func JsonEncode(data interface{})(string, error){
	str, err := json.Marshal(data)
	return string(str), err
}

//JsonDecode
func JsonDecode(data string)(interface{}, error){
	dataByte := []byte(data)
	type Json struct {
		Name string
		Age int
	}
	var jsons []Json
	err := json.Unmarshal(dataByte, &jsons)
	return jsons, err
}

//分割字符串
func Explode(s []byte, sep []byte, n int) [][]byte {
	start := 0
	if n < 0 {
		n = bytes.Count(s, sep) + 1
	}
	a := make([][]byte, n)
	na := 0
	for i := 0; i+1 <= len(s) && na+1 < 10; i++ {
		if s[i] == sep[0] {
			a[na] = s[start : i+0]
			na++
			start = i + 1
		}
	}
	a[na] = s[start:]
	return a
}

//切割关键词为HTML片段
func TagSplit(keywords []byte)string{
	if keywords == nil || len(keywords) <= 0 {
		return ""
	}
	content := ""
	tags := Explode(keywords, []byte(","), -1)
	for _, value := range tags {
		value = bytes.TrimSpace(value)
		content = content + fmt.Sprintf(`<a class ="tags" href="/tag/%s/1">%s</a>`, value, value)
	}
	return content
}

//四舍五入
func Round(val float64, places int) float64{
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t-x) > 0.50000000001 {
			t -= 1.0
		}
	}else{
		t = math.Ceil(-x)
		if (t+x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f
	if !math.IsInf(x, 0){
		return x
	}
	return t
}

func Get(reqUrl string) (string, error){
	response, err := http.Get(reqUrl)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		response.Body.Close()
		return "", err
	}
	return string(body), nil
}

func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, " ", "", -1)
}