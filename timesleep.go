/*wait/sleep timeout*/


import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello World!")
		time.Sleep(1000 * time.Millisecond)
	}
}



********************************************
получение даты:

const layout = "2006-01-02"

t := time.Now().Format(layout)
fmt.Println(t)
*******************************************