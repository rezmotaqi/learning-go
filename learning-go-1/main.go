package main

import (
	"fmt"
)

type bot interface {
	outp() string
}

type p1 struct{}
type p2 struct{}

func main() {

	// in oreder ro define a truct we must firstly define a newType on top of struct
	// and then define another variable and the ... to the newType we hust defined
	// type org struct {
	// 	name string
	// 	noe  int
	// }

	// hf := org{
	// 	name: "hf",
	// 	noe:  50,
	// }

	// fmt.Println(hf)

	// 3 ways of defining a map
	//1
	// colors2 := make(map[string]string)
	// colors2["a"] = "b"
	//2
	// var colors3 map[string]string
	// colors3["a"] = "b"
	//3
	// colors := map[string]string{
	// 	"red":   "fffff",
	// 	"green": "ffasd",
	// 	"white": "fffffffff",
	// }

	// colors["dddd"] = "ssssss"

	// fmt.Println(colors)

	p11 := p1{}
	p22 := p2{}

	out(p11)
	out(p22)

}

func out(b bot) {
	fmt.Println(b.outp())
}

func (p1) outp() string {
	return "p1 output"
}
func (p2) outp() string {
	return "p2 output"
}

// func printMap(m map[string]string) {
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}

// }
