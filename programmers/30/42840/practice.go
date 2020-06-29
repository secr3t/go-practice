package main

import "fmt"

func main() {
    answer1 := []int{1, 2, 3, 4, 5}
    answer2 := []int{1, 3, 2, 4, 2}

    fmt.Print(solution(answer1))
    fmt.Print(solution(answer2))
}

func solution(answers []int) []int {
    students := []int{0, 0, 0}
    patterns := [][]int{{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
        {2, 1, 2, 3, 2, 4, 2, 5},
        {3, 3, 1, 1, 2, 2, 4, 4, 5, 5}}

    for i, value := range patterns {
        students[i] = count(answers, value)
    }

    _, max := MinMax(students)

    indices := findAllIndices(students, max)

    return indices
}

func count(answers []int, pattern []int) int {
    score := 0
    patternLen := len(pattern)

    for i, answer := range answers {
        if answer == pattern[i%patternLen] {
            score++
        }
    }

    return score
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func findAllIndices(array []int, target int) []int {
    var indices []int

    for i, value := range array {
        if value == target {
            indices = append(indices, i + 1)
        }
    }

    return indices
}
