/*read web-site pinator and write in file*/

import (
    "fmt"
    "io"
    "http"
    "os"
    "strconv"
    "io/ioutil"
    "strings"
)

func main() {
    fmt.Printf("START \n")
    pinator_string := ""
    for num := 1002; num < 2004; num++ {
        num_str := strconv.Itoa(num)
            out, _ := os.Create("out.txt")
            web := "http://pinator.org/" + num_str
            file, _ := http.Get(web)
            io.Copy(out, file.Body)
            out.Close()
            file.Body.Close()

            find_text := "<link rel=\"stylesheet\" type=\"text/css\" href=\"/css/style.css\" />"

            page, _ := ioutil.ReadFile("out.txt")
            text := string(page)

            num_first_el := strings.Index(text, find_text)
            test_text := text[num_first_el + 71 : num_first_el + 1100]

            find_text = "</title>"
            num_first_el = strings.Index(test_text, find_text)
            test_text = test_text[0 : num_first_el]

            pinator_string = pinator_string + test_text + "\n"


    }
    fmt.Println(pinator_string)
    out, _ := os.Create("pinator.txt")
    out.Write([]uint8(pinator_string) )
    out.Close()

}
