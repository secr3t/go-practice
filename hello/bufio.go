package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var f *os.File
    f = os.Stdin
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println(">", text)
        if text == "exit" {
            break
        }
    }
}
