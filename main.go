package main

import (
	"fmt"
	"net/http"

	"github.com/luisfernandogaido/expediente/sessao"
)

func main() {
	sessao.Init("SESSAO_GO", sessao.NewGerenciadorRam())
	http.HandleFunc("/", index)
	http.HandleFunc("/zera", zera)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, _ := sessao.Inicia(r)
	var contador int
	if c, ok := sess["contador"]; ok {
		contador = c.(int)
	}
	contador++
	sess["contador"] = contador
	sessao.Salva(w, r, sess)
	fmt.Fprintln(w, contador)
}

func zera(w http.ResponseWriter, r *http.Request) {
	sessao.Destroi(w, r)
}
