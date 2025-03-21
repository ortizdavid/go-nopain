package main

import "github.com/ortizdavid/go-nopain/reflection"

type Object struct {
	Text string 
	Number int 
	Boolean bool
}

func main() {
	obj := Object{
		Text: "hello",
		Number: 123,
		Boolean: true,
	}

	reflection.VarDump(1, "", obj, 99, nil)
}