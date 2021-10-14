package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//while defining the new type on struct
//in the curly braces there are no commas or equal signs or semi colons
//just name of the filed and a space and type of the field
type person struct {
	name string
	age  int
}

func main() {

	//while declaring a variable of some custome type on struct
	// in the curly braces you have to write << name of the field <semi colon> value of field <comma> >>

	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	// body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	st := incall{}

	io.Copy(st, resp.Body)

}

type incall struct{}

func (incall) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes:", len(p))
	return len(p), nil

}
