package main

import (
	"io/ioutil"
	"os"
)

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func createFile(path string, filename string, fileText string) {

	d := []byte(fileText)
	err := ioutil.WriteFile(path+filename+".swift", d, 0644)
	if err != nil {
		panic(err)
	}

}
