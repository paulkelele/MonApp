package main

import (
    "fmt"
    "log"
    "io"
	"os"
)

func main() {
	ad:=readFile("t.txt")
	fmt.Println(string(ad))
}

func readFile(file string)(string){
	f, err := os.Open(file)
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    defer f.Close()
    buf := make([]byte, 8192)
	data :=""
    for {
        n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			data += string(buf[:n])
		}
    }
	return data
}