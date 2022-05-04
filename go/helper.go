// main.go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "github.com/vaughan0/go-ini"
)

// Graceful exit message
func Die(code int, msg string) {
    if msg != "" {
        fmt.Println(msg)
    }
    os.Exit(code)
}

// Set up global variables as per values in configuration file
func ProcessConfigFile() {
    // Read config file
    cfgfile = filepath.Join(os.Getenv("HOME"), "." + prgname + "rc")
    if _, err := os.Stat(cfgfile); os.IsNotExist(err) {
        Die(1, "Error. Missing " + cfgfile + " file. Run " + prgname + " -y to create a new one.")
    }

    f, _ := ini.LoadFile(cfgfile)
    v1, _ := f.Get("global", "svcurl")
    if v1 == "" {
        Die(1, "Error. svcurl not defined in " + cfgfile)
    }
    svcurl = v1

    v2, _ := f.Get("global", "svckey")
    if v2 == "" {
        Die(1, "Error. svckey not defined in " + cfgfile)
    }
    svckey = v2
}

// Create a skeleton configuration file with default hard-coded values
func CreateSkeltonConfigFile() {
    cfgfile = filepath.Join(os.Getenv("HOME"), "." + prgname + "rc")
    if _, err := os.Stat(cfgfile); os.IsNotExist(err) {
        content := "# Edit below values accordingly\n"
        content += "[global]\n"
        content += "svcurl = https://textbelt.com/text\n"
        content += "svckey = textbelt\n"
        err = ioutil.WriteFile(cfgfile, []byte(content), 0600)
        if err != nil {
            panic(err.Error())
        }
    } else {
        fmt.Printf("There's already a %s file.\n", cfgfile)
    }
}
