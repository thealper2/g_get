package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

func printUsage() {
    fmt.Printf(`
Usage:
    ` + os.Args[0] + ` <URL> <destinationPath>

Example:
    ` + os.Args[0] + ` http://127.0.0.1:4444/backup.zip /home/user/backup.zip

`)
}

func checkArgs() (string, string) {
    if len(os.Args) != 3 {
        printUsage()
        os.Exit(1)
    }
    return os.Args[1], os.Args[2]
}

func main() {
    fileURL, filePath := checkArgs()

    response, err := http.Get(fileURL)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    file, err := os.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("File downloaded and saved successfully: ", filePath)
}
