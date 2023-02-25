package main

import (
	"beesee/service"
	"fmt"
    "log"

)

func main() {
    log.SetPrefix("Main: ")
    log.SetFlags(0)
    names := []string{"Jean", "Eric", "Paul"}

     messages, err := service.Hellos(names)
     if err != nil{
        log.Fatal(err)
     }
     fmt.Println(messages)
}


