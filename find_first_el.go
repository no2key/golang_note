/* find first element string in file*/



import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	find_text := "l"

	file, _ := ioutil.ReadFile("text.txt")
	text := string(file)

	num_first_el := strings.Index(text, find_text)
	fmt.Println(num_first_el)
  fmt.Println(text[num_first_el : num_first_el + 100])

}