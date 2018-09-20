package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisfernandogaido/expediente/sessao"
)

func main() {
	if err := sessao.Ini("SESSAO_GO", "files"); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, _ := sessao.Inicia(r)
	var acessos int
	if a, ok := sess["acessos"]; ok {
		acessos = a.(int)
	}
	acessos++
	sess["acessos"] = acessos
	sessao.Salva(w, r, sess)
	if err := sessao.Destroi(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, acessos)
}
