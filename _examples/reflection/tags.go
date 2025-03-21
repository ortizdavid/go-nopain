package main

import (
	"fmt"

	"github.com/ortizdavid/go-nopain/reflection"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age int `numb:"age"`
}

func main() {
	fmt.Println(reflection.HasTag(User{}, "ID", "json"))
	fmt.Println(reflection.HasTag(User{}, "Name", "json"))
	fmt.Println(reflection.HasTag(User{}, "Age", "json"))

	fmt.Println(reflection.GetTag(User{}, "Email", "json"))
	fmt.Println(reflection.GetTag(User{}, "ID", "json"))
	fmt.Println(reflection.GetTag(User{}, "Name", "json"))
}