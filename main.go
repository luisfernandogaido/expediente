package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/luisfernandogaido/expediente/sessao"
)

func main() {
	ga, _ := sessao.NewGerenciadorArquivos("./sessoes")
	//ga = sessao.NewGerenciadorRam()
	sessao.Init("SESSAO_GO", ga)
	http.HandleFunc("/", index)
	http.HandleFunc("/zera", zera)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := sessao.Inicia(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var contador int
	if c, ok := sess["contador"]; ok {
		contador = c.(int)
	}
	contador++
	sess["contador"] = contador
	sess["agora"] = time.Now()
	if err := sessao.Salva(w, r, sess); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, sess)
}

func zera(w http.ResponseWriter, r *http.Request) {
	if err := sessao.Destroi(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
