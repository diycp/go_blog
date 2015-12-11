package utils
import (
	"regexp"
	"fmt"
	"strconv"
	"time"
	"encoding/xml"
	"encoding/base64"
	"io/ioutil"
)

type MethodResponse struct{
	Params []Param `xml:"params>param"`
}
type Param struct {
	Value Value `xml:"value"`
}
type Value struct {
	List []Value `xml:"array>data>value"`
	Object []Member `xml:"struct>member"`
	String string `xml:"string"`
	Int string `xml:"int"`
	Boolen string `xml:"boolean"`
	Base64 string `xml:"base64"`
	DateTime string `xml:"dateTime.ios8601"`
}
type Member struct {
	Name string `xml:"name"`
	Value Value `xml:"value"`
}

func GetMethodName(text string)(string, error){
	r, err := regexp.Compile(`<methodName>([\d\D]+)</methodName>`)
	if err != nil {
		return "", err
	}
	ar := r.FindStringSubmatch(text)
	return ar[1], nil
}

func unserialize(value Value) interface{}{
	if value.List != nil {
		result := make([]interface{}, len(value.List))
		for i, v := range value.List {
			result[i] = unserialize(v)
		}
		return result
	}else if value.Object != nil {
		result := make(map[string]interface{}, len(value.Object))
		for _, member := range value.Object {
			result[member.Name] = unserialize(member.Value)
		}
		return result
	}else if value.String != "" || value.Base64 != "" {
		return fmt.Sprintf("%s", value.String)
	}else if value.Int != "" {
		result, _ := strconv.Atoi(value.Int)
		return result
	}else if value.Boolen != "" {
		return value.Boolen == "1"
	}else if value.DateTime != "" {
		var format = "20060102T15:04:05"
		result, _ := time.Parse(format, value.DateTime)
		return result
	}
	return nil
}

func Unserialize(body []byte) interface{}{
	var response MethodResponse
	xml.Unmarshal(body, &response)
	result := make([]interface{}, len(response.Params))
	for i, param := range response.Params {
		result[i] = unserialize(param.Value)
	}
	return result
}

func ParseMedia(fullpath, str string) error{
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fullpath, data, 0644)
	return err
}