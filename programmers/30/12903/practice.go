package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(solution("abcde"))
    fmt.Println(solution("qwerty"))
}

func solution(s string) string {
    length := len(s)
    if length %2 == 0 {
        return strings.Join(append([]string{},strings.Split(s, "")[length/2-1:length/2+1]...), "")
    }

    return strings.Split(s, "")[length/2]
}