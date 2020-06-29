package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Print(solution([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}))
}

func solution(array []int, commands [][]int) []int {
    var result = make([]int, len(commands))

    for i, value := range commands {
        sliced := append([]int{}, array[value[0]-1 : value[1]]...)
        sort.Ints(sliced)
        result[i] = sliced[value[2]-1]
    }

    return result
}
