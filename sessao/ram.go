package sessao

import (
	"net/http"
	"sync"
)

type handlerRAM struct {
	sessoes map[string]map[string]interface{}
	mu      sync.RWMutex
}

func newHandlerRAM() handler {
	hr := handlerRAM{
		sessoes: make(map[string]map[string]interface{}),
	}
	return hr
}

func (h handlerRAM) inicia(r *http.Request) (map[string]interface{}, error) {
	novaSess := make(map[string]interface{})
	cookie, err := r.Cookie(nome)
	if err != nil {
		return novaSess, nil
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	sess, ok := h.sessoes[cookie.Value]
	if !ok {
		return novaSess, nil
	}
	return sess, nil
}

func (h handlerRAM) salva(w http.ResponseWriter, r *http.Request, sess map[string]interface{}) error {
	cookie, err := r.Cookie(nome)
	if err != nil {
		cookie = &http.Cookie{
			Name:  nome,
			Value: stringAleatoria(),
		}
		http.SetCookie(w, cookie)
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.sessoes[cookie.Value] = sess
	return nil
}

func (h handlerRAM) destroi(r *http.Request) error {
	cookie, err := r.Cookie(nome)
	if err != nil {
		return nil
	}
	delete(h.sessoes, cookie.Value)
	return nil
}
