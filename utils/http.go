package utils
import (
	"net/http"
	"os"
	"io"
)
func GetImage(url, path string) error {
	res, ok := http.Get(url)
	defer res.Body.Close()
	file, ok := os.Create(path)
	io.Copy(file, res.Body)
	return ok
}