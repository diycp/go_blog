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

func Test_JsonEncode(t *testing.T){
	arr := [3]string{"a", "b", "c"}
	json, err := tools.JsonEncode(arr)
	if err != nil {
		fmt.Println("Test_JsonEncode error" + ": ", err)
	}
	fmt.Println("Test_JsonEncode" + ": " + json)
}

func Test_JsonDecode(t *testing.T){
	json := `[{"name":"lixuan1", "age":25}, {"name":"lixuan2", "age":26}, {"name":"lixuan3", "age":27}]`
	arr, err := tools.JsonDecode(json)
	if err != nil {
		fmt.Println("Test_JsonDecode error" + ": ", err)
	}
	fmt.Println("Test_JsonDecode" + ": ", arr)
}

func Test_TagSplit(t *testing.T){
	var keywords = []byte("tag1, tag2, tag3")
	tags := tools.TagSplit(keywords)
	fmt.Println("Test_TagSplit" + ": " + tags)
}

