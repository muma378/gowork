package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

type Name struct{
    fname string
    lname string
}

var MAX_NAME_NUM int = 100

func main()  {
    var filepath string
    fmt.Println("Enter the filepath: ")
    fmt.Scan(&filepath)
    content, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Printf("Unable to open: %s\n", filepath)
        return
    }
    names := strings.Split(string(content), "\n")
    nameSli := make([]Name, 0, MAX_NAME_NUM)
    for i := 0; i < len(names); i++ {
        fullName := strings.Split(names[i], " ")
        if len(fullName) != 2 {
            fmt.Printf("Unable to parse name: %s", names[i])
            continue
        }
        name := Name{
            fname: fullName[0],
            lname: fullName[1],
        }
        nameSli = append(nameSli, name)
    }

    for i := 0; i < len(nameSli); i++ {
        fmt.Printf("first name: %s, last name: %s\n", nameSli[i].fname, nameSli[i].lname)
    }

}
