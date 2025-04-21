package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/justindwlee/bitcoinClone/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template


type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData {
		PageTitle: "Home",
		Blocks: blockchain.GetBlockchain().AllBlocks(),
	}
	templates.ExecuteTemplate(w, "home", data)
}

func add (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir+"pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir+"partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}