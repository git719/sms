// main.go
package main

import (
    "fmt"
    "os"
    "net/http"
    "net/url"
)

// Global constants
const (
    prgname = "smscli"
    prgver  = "1.1.0"

)

// Global variables
var (
    cfgfile = ""    // func ProcessConfigFile sets it to $HOME/.${prgname} + "rc"  
    svckey  = "textbelt"
    svcurl  = "https://textbelt.com/text"
)


func main() {

    tel, msg := "", ""
    argCount := len(os.Args[1:])

    if argCount < 1 || argCount > 2 {
        PrintUsage()
    }
    if os.Args[1] == "-y" {
        CreateSkeltonConfigFile()
    }

    ProcessConfigFile()

    tel = os.Args[1]
    msg = os.Args[2]

    values := url.Values{
        "key": {svckey},
        "phone": {tel},
        "message": {msg},
    }
    fmt.Printf("%s  %s  %s\n", svckey, tel, msg)
    resp, err := http.PostForm(svcurl, values)
    if err != nil {
        panic(err)
    }
    if resp.StatusCode != 200 {
        Die(1, "Error. HTTP error code = " + resp.Status)
    }

    Die(0,"")
}


func PrintUsage() {
    fmt.Printf("SMS CLI utility %s\n", prgver)
    fmt.Printf("%s <CellPhoneNum> <Message>\n", prgname)
    fmt.Printf("%s -y Create skeleton ~/.%src file\n", prgname, prgname)
    fmt.Println("Visit https://textbelt.com for more info.")
    Die(0,"")
}
