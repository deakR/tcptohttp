package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", nil)
	}

	str := ""

	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		// fmt.Printf("read: %s\n", string(data[:n]))

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i]) //add the content upto \n
			data = data[i+1:] //add the remaining content from that 8 bytes after the \n (this will be added to the string after this if condition block)
			fmt.Printf("read: %s\n", str)
			str = ""
		}

		str += string(data)
	}
	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}
}