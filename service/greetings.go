package service

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {
    if name == ""{
        return "", errors.New("no name given")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string)(map[string]string,error){
    messages := make(map[string]string)
    for _, name := range names{
        message,err := Hello(name)
        if err != nil{
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}

func init(){
    rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomFormat() string{
    formats := []string{
        "Hi, %v. Welcome!",
        "Bienvenu, %v ..comment va ?",
        "Salut %v tu vas bien ?",
    }
    return formats[rand.Intn(len((formats)))]
}