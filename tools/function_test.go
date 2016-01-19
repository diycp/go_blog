package tools

import (
	"fmt"
	"testing"
	"blog/tools"
)


func Test_CreateGUID(t *testing.T){
	guid := tools.CreateGUID()
	fmt.Println("Test_CreateGUID" + ": " + guid)
}

func Test_RandString(t *testing.T){
	str := tools.RandString(5)
	fmt.Println("Test_RandString" + ": " + str)
}

func Test_Md5(t *testing.T){
	fmt.Println("Test_Md5: " + tools.Md5("123456"))
}
func Test_Json(t *testing.T){
	arr := [3]map[string]interface{}{{"name":"lixuan1", "age":25}, {"name":"lixuan2", "age":26}, {"name":"lixuan3", "age":27}}
	//encode
	json, err := tools.JsonEncode(arr)
	if err != nil {
		fmt.Println("Test_JsonEncode error" + ": ", err)
	}else{
		fmt.Println("Test_JsonEncode" + ": " + json)
	}
	//decode
	data, err := tools.JsonDecode(json)
	if err != nil {
		fmt.Println("Test_JsonDecode error" + ": ", err)
	}else{
		fmt.Println("Test_JsonDecode" + ": ", data)
	}
}

func Test_TagSplit(t *testing.T){
	var keywords = []byte("tag1, tag2, tag3")
	tags := tools.TagSplit(keywords)
	fmt.Println("Test_TagSplit" + ": " + tags)
}

func Test_WriteFile(t *testing.T){
	path := "/tmp/go/"
	if !tools.PathExist(path) {
		err := tools.MkDir(path)
		if err != nil {
			fmt.Println(err)
		}
	}
	filename := path + "Test_WriteFile.txt"
	data := `hello go` + "\n"
	err := tools.WriteFile(filename, data)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Get(t *testing.T){
	content, _ := tools.Get("http://www.lanecn.com")
	fmt.Println("Test_Get" + ": " + content)
}

