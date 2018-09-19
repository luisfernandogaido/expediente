package main

import (
	"fmt"
	"net/http"

	"github.com/luisfernandogaido/expediente/sessao"
)

func main() {
	sessao.Nome("SESSAO_GO")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := sessao.Inicia(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, sess)
	if err := sessao.Salva(w, r, sess); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
