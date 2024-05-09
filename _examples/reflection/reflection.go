package main

import "github.com/ortizdavid/go-nopain/reflection"

func main() {
	reflection.VarDump(1, "", struct{
		Text string 
		Number int 
		Boolean bool
	}{
		Text: "hello",
		Number: 123,
		Boolean: true,
	},  99)
}