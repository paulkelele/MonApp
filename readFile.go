package main

import (
    "fmt"
    "log"
    "io"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	
	//ad:=readFile("t.txt")
	// fmt.Println(string(ad))
	wg.Add(1)
	go test()
	wg.Wait()
	 
}

func test(){
	defer wg.Done()
	fmt.Println("ddd")
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