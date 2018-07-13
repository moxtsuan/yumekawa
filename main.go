package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

func main() {
	var daku rune
	daku = 0x3099
	s := ""
	if len(os.Args) > 1 {
		t := os.Args[1:]
		for _, c := range t {
			s += c + " "
		}
		s = strings.TrimSpace(s)
	} else {
		s = "ウェェ"
	}
	s2 := string(norm.NFD.Bytes([]byte(s)))
	t := []rune{}
	for _, r := range s2 {
		if r == daku {
			continue
		}
		t = append(t, r)
		if r != ' ' {
			t = append(t, daku)
		}
	}
	s3 := (string)(t)
	fmt.Println(s3)
}
