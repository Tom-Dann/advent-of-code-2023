package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func TimeFunction(function func()) {
	start := time.Now()
	function()
	fmt.Println("Time elapsed:", time.Since(start))
}

func TimeFunctionInput(function func([]string), input []string) {
	start := time.Now()
	function(input)
	fmt.Println("Time elapsed:", time.Since(start))
}

func ReadInput(filename string, delim string) []string {
	raw, err := os.ReadFile(filename)
	Check(err)
	return strings.Split(strings.TrimSpace(string(raw)), delim)
}
