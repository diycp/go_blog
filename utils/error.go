package utils
import (
	"github.com/go-errors/errors"
	"fmt"
	"os"
	"strconv"
)

func New(code, msg string) error{
	err := `[` + code + `]` + msg
	return errors.New(err)
}
func Exit(code int, msg string){
	err := New(strconv.Itoa(code), msg)
	fmt.Println(err)
	os.Exit(1)
}