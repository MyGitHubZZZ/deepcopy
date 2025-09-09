package main

import (
	"fmt"

	"deepcopy"
)

type User struct {
	Name string
	Tags []string
}

func main() {
	u := User{Name: "A", Tags: []string{"x"}}
	u2 := deepcopy.Copy(u).(User)
	u2.Name = "y"
	fmt.Println(u.Name)
	fmt.Println(u2.Name)
}
