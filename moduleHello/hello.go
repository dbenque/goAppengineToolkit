package hello

import (
	"fmt"
	"net/http"

	"github.com/dbenque/goAppengineToolkit/dependencyHello"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, helloTxt.GetHelloTxt())
}
