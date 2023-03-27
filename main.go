package main

import (
	"os"
	"log"
	"io"
	"strings"
	"github.com/Trym123/funtemps/conv"
)

func main() {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0644)
	log.Println(src)
	

	var buffer []byte
	var linebuf []byte //nil
	buffer = make([]byte,1)
	bytesCount := 0

	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			elementArray := strings.Split(string(linebuf), ",")
		if len(elementArray) > 3 {
			celsius := elementArray[3]
			fahr := conv.CelsiusToFahrenheit(celsius)
			log.Println(elementArray[3])
		}
		linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		if err == io.EOF {
			break
		}
	}
}
