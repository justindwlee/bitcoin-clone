package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/justindwlee/bitcoinClone/explorer"
	"github.com/justindwlee/bitcoinClone/rest"
)

func usage(){
	fmt.Printf("Welcomde to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:   Set the PORT of the server\n")
	fmt.Printf("-mode=rest:	    Choose between 'html' and 'rest' and 'both'\n")
	runtime.Goexit()
}

func Start(){
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")

	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' and 'both'")

	flag.Parse()

	switch *mode  {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1000)
	default:
		usage()
	}

}