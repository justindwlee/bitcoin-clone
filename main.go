package main

import (
	"fmt"

	"github.com/justindwlee/bitcoinClone/person"
)



func main(){
	nico := person.Person{}
	nico.SetName("nico")
	nico.SetAge(33)
	fmt.Println(nico)
}
