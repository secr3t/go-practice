package main

import "fmt"

func main() {
    fmt.Println(solution(5, []int{2, 4}, []int{1, 3, 5}))
    fmt.Println(solution(5, []int{2, 4}, []int{3}))
    fmt.Println(solution(3, []int{3}, []int{1}))
}

func solution(n int, lost []int, reserve []int) int {
    students := make([]int, n)

    for i := range students {
        students[i] = 1
    }

    for _, v := range lost {
        students[v-1] = 0
    }

    for _, v := range reserve {
        if students[v-1] == 0 {
            students[v-1] = 1
        } else {
            students[v-1] = 2
        }
    }

    for i, curr := range students {
        if i != 0 && students[i-1] == 0 {
            if curr > 1 {
                students[i-1]++
                students[i]--
                continue
            }
        }

        if i != len(students) - 1 && students[i+1] == 0 {
            if curr > 1 {
                students[i+1]++
                students[i]--
            }
        }
    }

    var answer = 0
    for _, v := range students {
        if v > 0 {
            answer++
        }
    }

    return answer
}
