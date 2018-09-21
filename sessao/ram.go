package sessao

import (
	"sync"
)

type gerenciadorRAM struct {
	sessoes map[string]map[string]interface{}
	mu      sync.RWMutex
}

func NewGerenciadorRam() Gerenciador {
	return gerenciadorRAM{
		sessoes: make(map[string]map[string]interface{}),
	}
}

func (g gerenciadorRAM) Inicia(cookieValue string) (map[string]interface{}, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	sess, ok := g.sessoes[cookieValue]
	if !ok {
		return make(map[string]interface{}), nil
	}
	return sess, nil
}

func (g gerenciadorRAM) Salva(cookieValue string, sess map[string]interface{}) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.sessoes[cookieValue] = sess
	return nil
}

func (g gerenciadorRAM) Destroi(cookieValue string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.sessoes, cookieValue)
	return nil
}
