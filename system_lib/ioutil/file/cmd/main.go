package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const fpath = ""

func main() {
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf(string(bs))
}
