package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Scan()
	trimmed := bytes.TrimSpace(scanner.Bytes())

	if len(trimmed) == 0{
		fmt.Print(0)
		return
	}

	count := 1
	for _, ch := range trimmed {
		if ch == 32 {
			count = count + 1
		}
	}
	fmt.Print(count)
	//text := scanner.Text()
	//words := strings.Fields(text)
	//fmt.Println(len(words))
}
// fork-retest
// fork test
// after make pull request
// submodule test
// submodule add commit test
// force push test
// second push 
// after review
// test merge strategy
// 7.6.15.50
// 7.6.15.53
// test merge commit
// test rebase
