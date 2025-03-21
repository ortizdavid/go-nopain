package main

import "github.com/ortizdavid/go-nopain/reflection"

type Object struct {
	Text string 
	Number int 
	Boolean bool
}

func main() {
	obj := Object{"hello", 123, true}
	slice := []int{1, 2, 3}	
	arr := [4]float32{0.9, 1, -4.6, 90.3}
	map1 := map[string]any{"a": 10, "b": 3.8, "c": true}	
		
	reflection.VarDump(1, "test", 23.4, obj, slice, arr, map1, nil)
}