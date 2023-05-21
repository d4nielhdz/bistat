package main

import (
	src "bistat/src"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./obj.gob")
	objCode := new(src.ObjCode)
	if err == nil {
		fmt.Println("decoding")
		src.RegisterTypes()
		decoder := gob.NewDecoder(file)
		decErr := decoder.Decode(objCode)

		if decErr == nil {
			fmt.Println(len(objCode.Quads))
			for _, quad := range objCode.Quads {
				fmt.Println(src.OpToString(quad.Op))
			}
		} else {
			fmt.Println(decErr)
		}
	} else {
		fmt.Println(err)
	}
	file.Close()
}
