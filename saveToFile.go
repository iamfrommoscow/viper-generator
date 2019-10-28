package main

import "io/ioutil"

func saveToFile(text string) {
	d1 := []byte(text)
	err := ioutil.WriteFile("../hello.swift", d1, 0644)
	if err != nil {
		panic(err)
	}
}
