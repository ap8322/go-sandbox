package main

import (
	"fmt"
	"github.com/junegunn/fzf/src"
)

var revision string

func main() {
	var instance *string
	opts := fzf.ParseOptions()

	opts.Printer = func(str string) { instance = &str }

    fzf.Run(opts, revision)

	hoge := *instance

	fmt.Println(hoge)
}
