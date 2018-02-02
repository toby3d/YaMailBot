package main

import "fmt"

func errCheck(err error) {
	if err != nil {
		fmt.Sprintln(err.Error())
		panic(err.Error())
	}
}
