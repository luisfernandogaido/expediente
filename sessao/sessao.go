package sessao

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	nome string
	hdl  handler
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Ini(nomeSess string, nomeHandler string) error {
	var err error
	nome = nomeSess
	switch nomeHandler {
	case "ram":
		hdl = newHandlerRAM()
		return nil
	case "files":
		hdl, err = newHandlerFiles()
		return err
	default:
		return errors.New("sessão: nome do handler desconhecido")
	}
	return nil
}

func Inicia(r *http.Request) (map[string]interface{}, error) {
	return hdl.inicia(r)
}

func Salva(w http.ResponseWriter, r *http.Request, sess map[string]interface{}) error {
	return hdl.salva(w, r, sess)
}

func Destroi(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(nome)
	if err != nil {
		return hdl.destroi(r)
	}
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, cookie)
	return hdl.destroi(r)
}

func stringAleatoria() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}

type handler interface {
	inicia(*http.Request) (map[string]interface{}, error)
	salva(http.ResponseWriter, *http.Request, map[string]interface{}) error
	destroi(*http.Request) error
}
