package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(wd)
	if err != nil {
		panic(err)
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		panic(err)
	}
	
	for _, name := range names {
		fmt.Println(name)
	}
}
