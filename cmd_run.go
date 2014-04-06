/*run cmd from console windows (.bat)*/

1. Создаем .bat файл, в нем прописываем команды

2. КОД

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("D:\\GoMe\\src\\Scan_doc\\com_start.bat")
	err := cmd.Run()
	fmt.Println(err)
}




//********************************************
дополнительные возможности
out, _ := exec.Command("date").Output()
out_s := string(out)
fmt.Println(out_s)

cmd := exec.Command("cmd", "/C", "date", "08-08-2013")
cmd.Run()
//********************************************
