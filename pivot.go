package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var (
		delimiter = flag.String("str", "default", "string flag")
	)
	flag.Parse()
	fmt.Println(*delimiter)
	var s, t string
	if sc.Scan() {
		s = sc.Text()
	}
	if sc.Scan() {
		t = sc.Text()
	}
	fmt.Println(s)
	fmt.Println(t)
}
