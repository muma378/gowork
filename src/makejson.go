package main

import (
    "os"
    "fmt"
    "encoding/json"
)


type Person struct{
    name string
    address string
}

func main(){
    var name string
    fmt.Print("Enter a name: ")
    fmt.Scan(&name)
    var address string
    fmt.Print("Enter an address: ")
    fmt.Scan(&address)
    p := map[string] string{
        "name": name,
        "address": address,
    }
    barr, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Unable to marshal.\n")
	return
    }
    os.Stdout.Write(barr)
}
