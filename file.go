package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createDirectory(path string) {
	fmt.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func createFile(path string, filename string, fileText string) {

	d := []byte(fileText)
	err := ioutil.WriteFile(path+filename+".swift", d, 0775)
	if err != nil {
		panic(err)
	}

}
