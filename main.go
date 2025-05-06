package main

import (
	"github.com/justindwlee/bitcoinClone/cli"
	"github.com/justindwlee/bitcoinClone/db"
)



func main(){
	defer db.Close()
	cli.Start()
}

