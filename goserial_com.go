/*example using goserial(for COM port)*/

import (
	"github.com/tarm/goserial"
	"log"
)

func main() {
	c := &serial.Config{Name: "COM1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 2; i++ {
		j := string(i)
		j = j + "test"
		_, err := s.Write([]byte(j))
		if err != nil {
			log.Fatal(err)
		}
	}

	s.Close()

	/*
		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("%q", buf[:n])
	*/
}