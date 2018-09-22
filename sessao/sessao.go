package sessao

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	timeout = 1200
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	nomeSess string
	g        Gerenciador
)

type Gerenciador interface {
	Inicia(cookieValue string) (map[string]interface{}, error)
	Salva(cookieValue string, s map[string]interface{}) error
	Destroi(cookieValue string) error
}

func Init(nome string, gerenciador Gerenciador) {
	nomeSess = nome
	g = gerenciador

}

func Inicia(r *http.Request) (map[string]interface{}, error) {
	if g == nil {
		return nil, errors.New("sessão: gerenciador não definido")
	}
	novaSess := make(map[string]interface{})
	cookie, err := r.Cookie(nomeSess)
	if err != nil {
		return novaSess, nil
	}
	return g.Inicia(cookie.Value)
}

func Salva(w http.ResponseWriter, r *http.Request, sess map[string]interface{}) error {
	cookie, err := r.Cookie(nomeSess)
	if err != nil {
		sa, err := stringAleatoria()
		if err != nil {
			return err
		}
		cookie = &http.Cookie{
			Name:     nomeSess,
			Value:    sa,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	var ultimaAtividade time.Time
	if ua, ok := sess[cookie.Value+":ultima_atividade"]; ok {
		ultimaAtividade = ua.(time.Time)
		if time.Now().Sub(ultimaAtividade).Seconds() > timeout {
			Destroi(w, r)
			return errors.New("sessão: expirada")
		}
	}
	sess[cookie.Value+":ultima_atividade"] = time.Now()
	return g.Salva(cookie.Value, sess)
}

func Destroi(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(nomeSess)
	if err != nil {
		return err
	}
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, cookie)
	return g.Destroi(cookie.Value)
}

func stringAleatoria() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.New("sessao: erro ao gerar string aleatória")
	}
	return fmt.Sprintf("%x", bytes), nil
}
