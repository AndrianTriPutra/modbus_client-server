package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	note:
	//holding register
	reg 1-3 uint16
	reg 4   int16

	//input register
	reg  5 uint16
	reg  6 int16
	7,8  (hex7 + hex8) to uint32
	9,10 (hex9 + hex10) to float32
*/

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("     go run . mode \n")
		fmt.Printf("     mode : \n")
		fmt.Printf("     	    1 for goburrowRead \n")
		fmt.Printf("     	    2 for goburrowWrite \n")
		fmt.Printf("     	    3 for simmonRead \n")
		fmt.Printf("     	    4 for simmonWrite \n")
		fmt.Printf("     example : \n")
		fmt.Printf("               go run . 1 \n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "1":
		goburrowRead()

	case "2":
		goburrowWrite()

	case "3":
		simmonRead()

	case "4":
		simmonWrite()
	}

}
