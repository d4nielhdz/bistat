package main

import (
	src "bistat/src"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./" + os.Args[1])
	objCode := new(src.ObjCode)
	if err == nil {
		src.RegisterTypes()
		decoder := gob.NewDecoder(file)
		decErr := decoder.Decode(objCode)

		if decErr == nil {
			eCtx := NewEctx(*objCode)
			eCtx.Run()
			eCtx.PrintErrors()
		} else {
			fmt.Println(decErr)
		}
	} else {
		fmt.Println(err)
	}
	file.Close()
}
