package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		std, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(1)
		}
		for _, w := range std {
			z01.PrintRune(rune(w))
		}
	}
	arg := os.Args[1:]
	for _, name := range arg {
		file, err := os.Open(name)
		if err != nil {
			out := "ERROR: open " + name + ": no such file or directory\n"
			for _, w := range out {
				z01.PrintRune(rune(w))
			}
			os.Exit(1)
		}
		arr := make([]byte, 1000)
		n, _ := file.Read(arr)
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(arr[i]))
		}
	}
}
