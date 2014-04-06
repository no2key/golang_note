/*read web page*/


import (
	"io"
	"net/http"
	"os"
)

func main() {
	out, _ := os.Create("out.txt")
	file, _ := http.Get("http://www.gismeteo.ru")
	io.Copy(out, file.Body)

	out.Close()
	file.Body.Close()
}