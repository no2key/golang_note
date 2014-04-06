/*git + github note*/

package main

import (
		"fmt"
		"os"
		"os/exec"
		"bufio"
		"io/ioutil"
		"strings"
)

func main() {
	/*for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}*/ 

	if(os.Args[1] == "add") {
		fmt.Println("add")


		file_gnote, err := os.Open(".gnote")
		buf_read := bufio.NewReader(file_gnote)
		buf := make([]byte, 1024)
		n, err := buf_read.Read(buf)
		file_gnote.Close()

		file_bat, _ := os.Create("gnotecmd.bat")		
		buf_write := bufio.NewWriter(file_bat)
		cmds := "start /wait \"\" sublime_text " + os.Args[2] + "\n" 
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file_bat.Close()

		file, err := os.Create(os.Args[2])
		if(err != nil){
			fmt.Println("Error")
		}
		file.Close()

		cmd := exec.Command("gnotecmd.bat")
		err = cmd.Run()
		if(err != nil){
			fmt.Println("Error bat")
		}

		fmt.Println("test")

		find_start_com := "/*"
		find_end_com := "*/"
		file_com, _ := ioutil.ReadFile(os.Args[2])
		text := string(file_com)
		num_first_el := strings.Index(text, find_start_com)
		num_end_el := strings.Index(text, find_end_com)
		com_text := ""
		if(num_first_el == 0){
			com_text = text[(num_first_el+2):num_end_el]	
		}else {
			com_text = "add " + os.Args[2]
		}
		
		file_bat, _ = os.Create("gnotecmd.bat")		
		buf_write = bufio.NewWriter(file_bat)
		cmds = "git add " + os.Args[2] + "\n"
		cmds = cmds + "git commit -m \"" + com_text + "\"\n"
		cmds = cmds + "git push -u " + string(buf[:n]) + " master"
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file_bat.Close()

		cmd = exec.Command("gnotecmd.bat")
		err = cmd.Run()
		if(err != nil){
			fmt.Println("Error bat")
		}
	}

	if(os.Args[1] == "create") {
		fmt.Println("create")
		
		file, _ := os.Create(".gnote")
		buf_write := bufio.NewWriter(file)
		cmds := os.Args[2]
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file.Close()

		file, _ = os.Create(".gitignore")
		buf_write = bufio.NewWriter(file)
		cmds = ".gitignore\n.gnote\nmain.go"
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file.Close()

		name := os.Args[2] + ".txt"
		file, _ = os.Create(name)
		buf_write = bufio.NewWriter(file)
		cmds = "temporary file"
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file.Close()

		file_bat, _ := os.Create("gnotecmd.bat")		
		buf_write = bufio.NewWriter(file_bat)
		cmds = "git init\n"
		cmds = cmds + "git add \"*.txt\"\n"
		cmds = cmds + "git add \"*.go\"\n"
		cmds = cmds + "git add \"*.py\"\n"
		cmds = cmds + "git add \"*.c\"\n"
		cmds = cmds + "git commit -m \"create gnote\"\n"
		cmds = cmds + "git remote add " + os.Args[2] + " https://e3c80462fdc1374b4f7a884d3985ce6e3f82fa66@github.com/sasha-antonov/"  + os.Args[2] + ".git" + "\n"
		cmds = cmds + "git push -u " + os.Args[2] + " master"
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file_bat.Close()

		cmd := exec.Command("gnotecmd.bat")
		err := cmd.Run()
		if(err != nil){
			fmt.Println("Error bat")
		}		
	}

	if(os.Args[1] == "mod") {
		fmt.Println("mod")

		file_gnote, err := os.Open(".gnote")
		buf_read := bufio.NewReader(file_gnote)
		buf := make([]byte, 1024)
		n, err := buf_read.Read(buf)
		file_gnote.Close()

		file_bat, _ := os.Create("gnotecmd.bat")		
		buf_write := bufio.NewWriter(file_bat)
		cmds := "start /wait \"\" sublime_text " + os.Args[2] + "\n" 
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file_bat.Close()

		cmd := exec.Command("gnotecmd.bat")
		err = cmd.Run()
		if(err != nil){
			fmt.Println("Error bat 1")
		}

		find_start_com := "/*"
		find_end_com := "*/"
		file_com, _ := ioutil.ReadFile(os.Args[2])
		text := string(file_com)
		num_first_el := strings.Index(text, find_start_com)
		num_end_el := strings.Index(text, find_end_com)
		com_text := ""
		if(num_first_el == 0){
			com_text = text[(num_first_el+2):num_end_el]	
		}else {
			com_text = "add " + os.Args[2]
		}

		file_bat, _ = os.Create("gnotecmd.bat")		
		buf_write = bufio.NewWriter(file_bat)
		cmds = "git add " + os.Args[2] + "\n"
		cmds = cmds + "git commit -m \"" + com_text + "\"\n"
		cmds = cmds + "git push -u " + string(buf[:n]) + " master"
		buf_write.WriteString(cmds)
		buf_write.Flush()
		file_bat.Close()

		cmd = exec.Command("gnotecmd.bat")
		err = cmd.Run()
		if(err != nil){
			fmt.Println("Error bat 2")
		}
	}
}
