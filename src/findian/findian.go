package main
import (
    "os"
    "fmt"
    "strings"
    "bufio"
)

func main(){
    fmt.Printf("Please enter a string:\n")
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')
    s = strings.ToLower(strings.Trim(s, "\n"))
    if strings.HasPrefix(s, "i")  && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
        fmt.Printf("Found!")
    } else {
        fmt.Printf("Not Found!")
    }
}
