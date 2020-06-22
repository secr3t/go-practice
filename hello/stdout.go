package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var args = os.Args
	var stringArr = args[1:]
	n, err := io.WriteString(os.Stdout, strings.Join(stringArr, " ") + "\n")
	fmt.Println(n, err)
}
