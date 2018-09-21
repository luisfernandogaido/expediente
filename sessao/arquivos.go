package sessao

import (
	"encoding/gob"
	"os"
	"path"
	"time"
)

type gerenciadorArquivos struct {
	pasta string
}

func NewGerenciadorArquivos(pasta string) (Gerenciador, error) {
	err := os.Mkdir(pasta, 0644)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	gob.Register(time.Time{})
	return gerenciadorArquivos{pasta}, nil
}

func (g gerenciadorArquivos) Inicia(cookieValue string) (map[string]interface{}, error) {
	sess := make(map[string]interface{})
	f, err := os.Open(path.Join(g.pasta, cookieValue))
	if err != nil {
		return sess, nil
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	err = dec.Decode(&sess)
	return sess, err
}

func (g gerenciadorArquivos) Salva(cookieValue string, sess map[string]interface{}) error {
	f, err := os.OpenFile(path.Join(g.pasta, cookieValue), os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(sess)
}

func (g gerenciadorArquivos) Destroi(cookieValue string) error {
	return os.Remove(path.Join(g.pasta, cookieValue))
}
