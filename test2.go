package main

import (
	"fmt"
	"strings"
)

const big = 1 << 100

func main() {
	var arg any = big
	fmt.Printf("arg=%T", arg)
	fmt.Printf("big=%T", big)
}

func main2() {
	files := []string{"a", "x/b"}
	covs := []string{"/d/a", "/d/x/b"}

	stop := false
	for _, file := range files {
		for _, cov := range covs {
			if strings.HasSuffix(cov, file) {
				SourceCodeFilePrefix, _ := strings.CutSuffix(cov, file)
				fmt.Println("----", SourceCodeFilePrefix)
				stop = true
				break
			}
		}

		if stop {
			break
		}
	}
}
