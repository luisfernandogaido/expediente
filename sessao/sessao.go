package sessao

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	nome    string
	sessoes map[string]map[string]interface{}
	mu      sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
	sessoes = make(map[string]map[string]interface{})
	fmt.Println("aoba")
}

func Nome(n string) {
	nome = n
}

func Inicia(r *http.Request) (map[string]interface{}, error) {
	sess := make(map[string]interface{})
	return sess, nil
}

func Salva(w http.ResponseWriter, r *http.Request, sess map[string]interface{}) error {
	return nil
}
