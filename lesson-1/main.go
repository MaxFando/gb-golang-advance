package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrorUnknown    = errors.New("unknown error")
	ErrorCreateFile = errors.New("could not create file")
	ErrorCloseFile  = errors.New("can not close file")
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Errorf("%w: %s", ErrorUnknown, r))
		}
	}()

	file, errCreate := os.Create("./lesson-1/app.txt")
	if errCreate != nil {
		log.Println(fmt.Errorf("%w: %s", ErrorCreateFile, errCreate.Error()))
	}

	defer func() {
		errClose := file.Close()
		if errClose != nil {
			log.Println(fmt.Errorf("%w: %s", ErrorCloseFile, errClose.Error()))
		}
	}()

	panic("AAA")
}
