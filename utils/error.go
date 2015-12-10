package utils
import (
	"github.com/go-errors/errors"
	"fmt"
	"os"
)

func New(code, msg string) error{
	err := `[` + code + `]` + msg
	return errors.New(err)
}
func Exit(code, msg string){
	err := New(code, msg)
	fmt.Println(err)
	os.Exit(1)
}