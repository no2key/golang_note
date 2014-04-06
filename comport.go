/*GUI + COM port 1*/

import (
	. "github.com/lxn/walk/declarative"
	"github.com/tarm/goserial"
)

func main() {

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{200, 200},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "Hello, serial port - COM_1",
				OnClicked: func() {
					port := &serial.Config{Name: "COM1", Baud: 9600}
					com_1, _ := serial.OpenPort(port)
					com_1.Write([]byte("Hello"))
					com_1.Close()
				},
			},
		},
	}.Run()
}


**********************************
Без манифеста не запускать exe - 
сначала создаем exe, потом manifest,
потом запускаем

**************************************
Для того что-бы программа работала - 
необходимо создать exe файл и Manifest файл. 
**************************************

Build app
In the directory containing test.go run
go build
************************************

To get rid of the cmd window, instead run 
(для запуска без командной строки
Эту команду в bash вместо предыдущей)
go build -ldflags="-H windowsgui"
************************************

Create Manifest test.exe.manifest

<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
        <assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
        <dependency>
            <dependentAssembly>
                <assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
            </dependentAssembly>
        </dependency>
    </assembly>
****************************************************************************************************************